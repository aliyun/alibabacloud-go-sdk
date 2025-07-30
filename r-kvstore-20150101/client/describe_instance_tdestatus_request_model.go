// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTDEStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceTDEStatusRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeInstanceTDEStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceTDEStatusRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeInstanceTDEStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceTDEStatusRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeInstanceTDEStatusRequest
	GetSecurityToken() *string
}

type DescribeInstanceTDEStatusRequest struct {
	// The ID of the Tair (Redis OSS-compatible) instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeInstanceTDEStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTDEStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTDEStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceTDEStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceTDEStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceTDEStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceTDEStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceTDEStatusRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceTDEStatusRequest) SetInstanceId(v string) *DescribeInstanceTDEStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceTDEStatusRequest) SetOwnerAccount(v string) *DescribeInstanceTDEStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceTDEStatusRequest) SetOwnerId(v int64) *DescribeInstanceTDEStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceTDEStatusRequest) SetResourceOwnerAccount(v string) *DescribeInstanceTDEStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceTDEStatusRequest) SetResourceOwnerId(v int64) *DescribeInstanceTDEStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceTDEStatusRequest) SetSecurityToken(v string) *DescribeInstanceTDEStatusRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceTDEStatusRequest) Validate() error {
	return dara.Validate(s)
}
