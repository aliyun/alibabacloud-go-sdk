// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateSnapshot(v bool) *RebootDesktopsRequest
	GetCreateSnapshot() *bool
	SetDesktopId(v []*string) *RebootDesktopsRequest
	GetDesktopId() []*string
	SetOsUpdate(v bool) *RebootDesktopsRequest
	GetOsUpdate() *bool
	SetPatchId(v string) *RebootDesktopsRequest
	GetPatchId() *string
	SetRegionId(v string) *RebootDesktopsRequest
	GetRegionId() *string
}

type RebootDesktopsRequest struct {
	CreateSnapshot *bool `json:"CreateSnapshot,omitempty" xml:"CreateSnapshot,omitempty"`
	// The IDs of the cloud computers. You can specify 1 to 100 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	OsUpdate  *bool     `json:"OsUpdate,omitempty" xml:"OsUpdate,omitempty"`
	// example:
	//
	// KB5082063
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebootDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootDesktopsRequest) GoString() string {
	return s.String()
}

func (s *RebootDesktopsRequest) GetCreateSnapshot() *bool {
	return s.CreateSnapshot
}

func (s *RebootDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *RebootDesktopsRequest) GetOsUpdate() *bool {
	return s.OsUpdate
}

func (s *RebootDesktopsRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *RebootDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebootDesktopsRequest) SetCreateSnapshot(v bool) *RebootDesktopsRequest {
	s.CreateSnapshot = &v
	return s
}

func (s *RebootDesktopsRequest) SetDesktopId(v []*string) *RebootDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *RebootDesktopsRequest) SetOsUpdate(v bool) *RebootDesktopsRequest {
	s.OsUpdate = &v
	return s
}

func (s *RebootDesktopsRequest) SetPatchId(v string) *RebootDesktopsRequest {
	s.PatchId = &v
	return s
}

func (s *RebootDesktopsRequest) SetRegionId(v string) *RebootDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RebootDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
