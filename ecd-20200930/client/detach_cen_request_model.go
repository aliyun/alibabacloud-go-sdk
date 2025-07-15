// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOfficeSiteId(v string) *DetachCenRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *DetachCenRequest
	GetRegionId() *string
}

type DetachCenRequest struct {
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachCenRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachCenRequest) GoString() string {
	return s.String()
}

func (s *DetachCenRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *DetachCenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachCenRequest) SetOfficeSiteId(v string) *DetachCenRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *DetachCenRequest) SetRegionId(v string) *DetachCenRequest {
	s.RegionId = &v
	return s
}

func (s *DetachCenRequest) Validate() error {
	return dara.Validate(s)
}
