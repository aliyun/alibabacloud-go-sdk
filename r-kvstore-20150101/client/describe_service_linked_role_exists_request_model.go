// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLinkedRoleExistsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeServiceLinkedRoleExistsRequest
	GetEngine() *string
	SetOwnerAccount(v string) *DescribeServiceLinkedRoleExistsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeServiceLinkedRoleExistsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeServiceLinkedRoleExistsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeServiceLinkedRoleExistsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeServiceLinkedRoleExistsRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeServiceLinkedRoleExistsRequest
	GetSecurityToken() *string
}

type DescribeServiceLinkedRoleExistsRequest struct {
	// The database engine of the instance. Only Redis is supported.
	//
	// example:
	//
	// Redis
	Engine       *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// > This parameter does not affect the query results. You can specify any region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeServiceLinkedRoleExistsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleExistsRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleExistsRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeServiceLinkedRoleExistsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeServiceLinkedRoleExistsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeServiceLinkedRoleExistsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServiceLinkedRoleExistsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeServiceLinkedRoleExistsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeServiceLinkedRoleExistsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeServiceLinkedRoleExistsRequest) SetEngine(v string) *DescribeServiceLinkedRoleExistsRequest {
	s.Engine = &v
	return s
}

func (s *DescribeServiceLinkedRoleExistsRequest) SetOwnerAccount(v string) *DescribeServiceLinkedRoleExistsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeServiceLinkedRoleExistsRequest) SetOwnerId(v int64) *DescribeServiceLinkedRoleExistsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeServiceLinkedRoleExistsRequest) SetRegionId(v string) *DescribeServiceLinkedRoleExistsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeServiceLinkedRoleExistsRequest) SetResourceOwnerAccount(v string) *DescribeServiceLinkedRoleExistsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeServiceLinkedRoleExistsRequest) SetResourceOwnerId(v int64) *DescribeServiceLinkedRoleExistsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeServiceLinkedRoleExistsRequest) SetSecurityToken(v string) *DescribeServiceLinkedRoleExistsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeServiceLinkedRoleExistsRequest) Validate() error {
	return dara.Validate(s)
}
