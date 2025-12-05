// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserVpcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *GetUserVpcsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetUserVpcsRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetUserVpcsRequest
	GetRegionId() *string
	SetVpcId(v string) *GetUserVpcsRequest
	GetVpcId() *string
}

type GetUserVpcsRequest struct {
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-2ze22asdfuwiea2ebjkqhf4pyj
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetUserVpcsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserVpcsRequest) GoString() string {
	return s.String()
}

func (s *GetUserVpcsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetUserVpcsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetUserVpcsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUserVpcsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *GetUserVpcsRequest) SetPageNumber(v int32) *GetUserVpcsRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserVpcsRequest) SetPageSize(v int32) *GetUserVpcsRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserVpcsRequest) SetRegionId(v string) *GetUserVpcsRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserVpcsRequest) SetVpcId(v string) *GetUserVpcsRequest {
	s.VpcId = &v
	return s
}

func (s *GetUserVpcsRequest) Validate() error {
	return dara.Validate(s)
}
