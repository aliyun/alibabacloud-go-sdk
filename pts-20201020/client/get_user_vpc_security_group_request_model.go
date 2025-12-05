// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserVpcSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *GetUserVpcSecurityGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetUserVpcSecurityGroupRequest
	GetPageSize() *int32
	SetRegionId(v string) *GetUserVpcSecurityGroupRequest
	GetRegionId() *string
	SetVpcId(v string) *GetUserVpcSecurityGroupRequest
	GetVpcId() *string
}

type GetUserVpcSecurityGroupRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp10xjz7c7oqjgasodihj1kx7t
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetUserVpcSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserVpcSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *GetUserVpcSecurityGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetUserVpcSecurityGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetUserVpcSecurityGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUserVpcSecurityGroupRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *GetUserVpcSecurityGroupRequest) SetPageNumber(v int32) *GetUserVpcSecurityGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *GetUserVpcSecurityGroupRequest) SetPageSize(v int32) *GetUserVpcSecurityGroupRequest {
	s.PageSize = &v
	return s
}

func (s *GetUserVpcSecurityGroupRequest) SetRegionId(v string) *GetUserVpcSecurityGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserVpcSecurityGroupRequest) SetVpcId(v string) *GetUserVpcSecurityGroupRequest {
	s.VpcId = &v
	return s
}

func (s *GetUserVpcSecurityGroupRequest) Validate() error {
	return dara.Validate(s)
}
