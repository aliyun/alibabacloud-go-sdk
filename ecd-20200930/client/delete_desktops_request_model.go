// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDesktopsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v []*string) *DeleteDesktopsRequest
	GetDesktopId() []*string
	SetRegionId(v string) *DeleteDesktopsRequest
	GetRegionId() *string
	SetResellerOwnerUid(v int64) *DeleteDesktopsRequest
	GetResellerOwnerUid() *int64
}

type DeleteDesktopsRequest struct {
	// A list of 1 to 100 cloud computer IDs.
	//
	// This parameter is required.
	DesktopId []*string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty" type:"Repeated"`
	// The region ID. Call [](t2167755.xdita#)to get a list of regions that Elastic Desktop Service (EDS) supports.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResellerOwnerUid *int64  `json:"ResellerOwnerUid,omitempty" xml:"ResellerOwnerUid,omitempty"`
}

func (s DeleteDesktopsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDesktopsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDesktopsRequest) GetDesktopId() []*string {
	return s.DesktopId
}

func (s *DeleteDesktopsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDesktopsRequest) GetResellerOwnerUid() *int64 {
	return s.ResellerOwnerUid
}

func (s *DeleteDesktopsRequest) SetDesktopId(v []*string) *DeleteDesktopsRequest {
	s.DesktopId = v
	return s
}

func (s *DeleteDesktopsRequest) SetRegionId(v string) *DeleteDesktopsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDesktopsRequest) SetResellerOwnerUid(v int64) *DeleteDesktopsRequest {
	s.ResellerOwnerUid = &v
	return s
}

func (s *DeleteDesktopsRequest) Validate() error {
	return dara.Validate(s)
}
