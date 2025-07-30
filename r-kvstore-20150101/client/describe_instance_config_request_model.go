// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeInstanceConfigRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeInstanceConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeInstanceConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceConfigRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeInstanceConfigRequest
	GetSecurityToken() *string
}

type DescribeInstanceConfigRequest struct {
	// The instance ID. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the ID of the instance.
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

func (s DescribeInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeInstanceConfigRequest) SetInstanceId(v string) *DescribeInstanceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceConfigRequest) SetOwnerAccount(v string) *DescribeInstanceConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceConfigRequest) SetOwnerId(v int64) *DescribeInstanceConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceConfigRequest) SetResourceOwnerAccount(v string) *DescribeInstanceConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceConfigRequest) SetResourceOwnerId(v int64) *DescribeInstanceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceConfigRequest) SetSecurityToken(v string) *DescribeInstanceConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
