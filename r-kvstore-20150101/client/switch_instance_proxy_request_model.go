// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchInstanceProxyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SwitchInstanceProxyRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *SwitchInstanceProxyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SwitchInstanceProxyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SwitchInstanceProxyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SwitchInstanceProxyRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *SwitchInstanceProxyRequest
	GetSecurityToken() *string
}

type SwitchInstanceProxyRequest struct {
	// The ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the ID of the instance.
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

func (s SwitchInstanceProxyRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchInstanceProxyRequest) GoString() string {
	return s.String()
}

func (s *SwitchInstanceProxyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchInstanceProxyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SwitchInstanceProxyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SwitchInstanceProxyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SwitchInstanceProxyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SwitchInstanceProxyRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SwitchInstanceProxyRequest) SetInstanceId(v string) *SwitchInstanceProxyRequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchInstanceProxyRequest) SetOwnerAccount(v string) *SwitchInstanceProxyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchInstanceProxyRequest) SetOwnerId(v int64) *SwitchInstanceProxyRequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchInstanceProxyRequest) SetResourceOwnerAccount(v string) *SwitchInstanceProxyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchInstanceProxyRequest) SetResourceOwnerId(v int64) *SwitchInstanceProxyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchInstanceProxyRequest) SetSecurityToken(v string) *SwitchInstanceProxyRequest {
	s.SecurityToken = &v
	return s
}

func (s *SwitchInstanceProxyRequest) Validate() error {
	return dara.Validate(s)
}
