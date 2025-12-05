// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserVpcVSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *GetUserVpcVSwitchRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetUserVpcVSwitchRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetUserVpcVSwitchRequest
	GetRegionId() *string
	SetVpcId(v string) *GetUserVpcVSwitchRequest
	GetVpcId() *string
}

type GetUserVpcVSwitchRequest struct {
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-2ze22scdz2ebdfjasdfjkqhf4pyj
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetUserVpcVSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserVpcVSwitchRequest) GoString() string {
	return s.String()
}

func (s *GetUserVpcVSwitchRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetUserVpcVSwitchRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetUserVpcVSwitchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUserVpcVSwitchRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *GetUserVpcVSwitchRequest) SetPageNumber(v int32) *GetUserVpcVSwitchRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserVpcVSwitchRequest) SetPageSize(v int32) *GetUserVpcVSwitchRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserVpcVSwitchRequest) SetRegionId(v string) *GetUserVpcVSwitchRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserVpcVSwitchRequest) SetVpcId(v string) *GetUserVpcVSwitchRequest {
	s.VpcId = &v
	return s
}

func (s *GetUserVpcVSwitchRequest) Validate() error {
	return dara.Validate(s)
}
