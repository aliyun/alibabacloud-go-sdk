// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEngineVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeEngineVersionRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeEngineVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeEngineVersionRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeEngineVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEngineVersionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeEngineVersionRequest
	GetSecurityToken() *string
}

type DescribeEngineVersionRequest struct {
	// The instance ID. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the instance ID.
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

func (s DescribeEngineVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEngineVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeEngineVersionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEngineVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeEngineVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeEngineVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEngineVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEngineVersionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeEngineVersionRequest) SetInstanceId(v string) *DescribeEngineVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeEngineVersionRequest) SetOwnerAccount(v string) *DescribeEngineVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEngineVersionRequest) SetOwnerId(v int64) *DescribeEngineVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEngineVersionRequest) SetResourceOwnerAccount(v string) *DescribeEngineVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEngineVersionRequest) SetResourceOwnerId(v int64) *DescribeEngineVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEngineVersionRequest) SetSecurityToken(v string) *DescribeEngineVersionRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeEngineVersionRequest) Validate() error {
	return dara.Validate(s)
}
