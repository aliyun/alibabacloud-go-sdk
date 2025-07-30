// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeAuditLogConfigRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeAuditLogConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAuditLogConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAuditLogConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAuditLogConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAuditLogConfigRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeAuditLogConfigRequest
	GetSecurityToken() *string
}

type DescribeAuditLogConfigRequest struct {
	// The ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeInstanceAttribute](https://help.aliyun.com/document_detail/473779.html) operation to query the region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hanghzou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAuditLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAuditLogConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAuditLogConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAuditLogConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAuditLogConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAuditLogConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAuditLogConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeAuditLogConfigRequest) SetInstanceId(v string) *DescribeAuditLogConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetOwnerAccount(v string) *DescribeAuditLogConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetOwnerId(v int64) *DescribeAuditLogConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetRegionId(v string) *DescribeAuditLogConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetResourceOwnerAccount(v string) *DescribeAuditLogConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetResourceOwnerId(v int64) *DescribeAuditLogConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) SetSecurityToken(v string) *DescribeAuditLogConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAuditLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
