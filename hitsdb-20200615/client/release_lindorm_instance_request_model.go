// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseLindormInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImmediately(v bool) *ReleaseLindormInstanceRequest
	GetImmediately() *bool
	SetInstanceId(v string) *ReleaseLindormInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ReleaseLindormInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReleaseLindormInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ReleaseLindormInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReleaseLindormInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ReleaseLindormInstanceRequest
	GetSecurityToken() *string
}

type ReleaseLindormInstanceRequest struct {
	// Specifies whether to release the instance immediately. If you set this parameter to false, data in the released instance is retained for seven days before it is completely deleted. If you set this parameter to true, data in the released instance is immediately deleted. The default value is false.
	//
	// example:
	//
	// false
	Immediately *bool `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// Instance ID, which can be obtained by calling the [GetLindormInstanceList](https://help.aliyun.com/document_detail/426069.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ReleaseLindormInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseLindormInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseLindormInstanceRequest) GetImmediately() *bool {
	return s.Immediately
}

func (s *ReleaseLindormInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseLindormInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReleaseLindormInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReleaseLindormInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReleaseLindormInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReleaseLindormInstanceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ReleaseLindormInstanceRequest) SetImmediately(v bool) *ReleaseLindormInstanceRequest {
	s.Immediately = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetInstanceId(v string) *ReleaseLindormInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetOwnerAccount(v string) *ReleaseLindormInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetOwnerId(v int64) *ReleaseLindormInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetResourceOwnerAccount(v string) *ReleaseLindormInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetResourceOwnerId(v int64) *ReleaseLindormInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) SetSecurityToken(v string) *ReleaseLindormInstanceRequest {
	s.SecurityToken = &v
	return s
}

func (s *ReleaseLindormInstanceRequest) Validate() error {
	return dara.Validate(s)
}
