// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *RebootDesktopsRequest
	GetDesktopId() []*string
	SetOsUpdate(v bool) *RebootDesktopsRequest
	GetOsUpdate() *bool
	SetRegionId(v string) *RebootDesktopsRequest
	GetRegionId() *string
}

type RebootDesktopsRequest struct {
	// The IDs of the cloud computers. You can specify 1 to 100 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	OsUpdate  *bool     `json:"OsUpdate,omitempty" xml:"OsUpdate,omitempty"`
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

func (s *RebootDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *RebootDesktopsRequest) GetOsUpdate() *bool {
	return s.OsUpdate
}

func (s *RebootDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebootDesktopsRequest) SetDesktopId(v []*string) *RebootDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *RebootDesktopsRequest) SetOsUpdate(v bool) *RebootDesktopsRequest {
	s.OsUpdate = &v
	return s
}

func (s *RebootDesktopsRequest) SetRegionId(v string) *RebootDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *RebootDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
