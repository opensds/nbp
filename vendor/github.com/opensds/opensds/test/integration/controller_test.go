// Copyright (c) 2017 Huawei Technologies Co., Ltd. All Rights Reserved.
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

// +build integration

package integration

import (
	"encoding/json"
	"testing"

	"github.com/opensds/opensds/pkg/controller/volume"
	pb "github.com/opensds/opensds/pkg/dock/proto"
	"github.com/opensds/opensds/pkg/model"
)

var vc = volume.NewController(
	&pb.CreateVolumeOpts{},
	&pb.DeleteVolumeOpts{},
	&pb.CreateVolumeSnapshotOpts{},
	&pb.DeleteVolumeSnapshotOpts{},
	&pb.CreateAttachmentOpts{},
	&pb.DeleteAttachmentOpts{},
)

var dckInfo = &model.DockSpec{
	Endpoint:   "localhost:50050",
	DriverName: "default",
}

func TestControllerCreateVolume(t *testing.T) {
	vc.SetDock(dckInfo)

	vol, err := vc.CreateVolume()
	if err != nil {
		t.Error("create volume in controller failed:", err)
		return
	}

	volBody, _ := json.MarshalIndent(vol, "", "	")
	t.Log(string(volBody))
}

func TestControllerDeleteVolume(t *testing.T) {
	vc.SetDock(dckInfo)

	res := vc.DeleteVolume()
	if err := res.ToError(); err != nil {
		t.Error("delete volume in controller failed:", err)
		return
	}

	resBody, _ := json.MarshalIndent(res, "", "	")
	t.Log(string(resBody))
}

func TestControllerCreateVolumeAttachment(t *testing.T) {
	vc.SetDock(dckInfo)

	atc, err := vc.CreateVolumeAttachment()
	if err != nil {
		t.Error("create volume attachment in controller failed:", err)
		return
	}

	atcBody, _ := json.MarshalIndent(atc, "", "	")
	t.Log(string(atcBody))
}

func TestControllerDeleteVolumeAttachment(t *testing.T) {
	vc.SetDock(dckInfo)

	res := vc.DeleteVolumeAttachment()
	if err := res.ToError(); err != nil {
		t.Error("delete volume attachment in controller failed:", err)
		return
	}

	resBody, _ := json.MarshalIndent(res, "", "	")
	t.Log(string(resBody))
}

func TestControllerCreateVolumeSnapshot(t *testing.T) {
	vc.SetDock(dckInfo)

	snp, err := vc.CreateVolumeSnapshot()
	if err != nil {
		t.Error("create volume snapshot in controller failed:", err)
		return
	}

	snpBody, _ := json.MarshalIndent(snp, "", "	")
	t.Log(string(snpBody))
}

func TestControllerDeleteVolumeSnapshot(t *testing.T) {
	vc.SetDock(dckInfo)

	res := vc.DeleteVolumeSnapshot()
	if err := res.ToError(); err != nil {
		t.Error("delete volume snapshot in controller failed:", err)
		return
	}

	resBody, _ := json.MarshalIndent(res, "", "	")
	t.Log(string(resBody))
}
