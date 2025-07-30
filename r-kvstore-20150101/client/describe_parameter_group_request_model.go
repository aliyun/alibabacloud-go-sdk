// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeParameterGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeParameterGroupRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *DescribeParameterGroupRequest
	GetParameterGroupId() *string
	SetRegionId(v string) *DescribeParameterGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeParameterGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParameterGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeParameterGroupRequest
	GetSecurityToken() *string
}

type DescribeParameterGroupRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the parameter template.
	//
	// This parameter is required.
	//
	// example:
	//
	// rpg-sys-00*****
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeParameterGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParameterGroupRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DescribeParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeParameterGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParameterGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParameterGroupRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeParameterGroupRequest) SetOwnerAccount(v string) *DescribeParameterGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetOwnerId(v int64) *DescribeParameterGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetParameterGroupId(v string) *DescribeParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetRegionId(v string) *DescribeParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetResourceOwnerAccount(v string) *DescribeParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetResourceOwnerId(v int64) *DescribeParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetSecurityToken(v string) *DescribeParameterGroupRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
