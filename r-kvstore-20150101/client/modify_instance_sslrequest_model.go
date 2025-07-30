// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceSSLRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceSSLRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceSSLRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceSSLRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceSSLRequest
	GetResourceOwnerId() *int64
	SetSSLEnabled(v string) *ModifyInstanceSSLRequest
	GetSSLEnabled() *string
	SetSecurityToken(v string) *ModifyInstanceSSLRequest
	GetSecurityToken() *string
}

type ModifyInstanceSSLRequest struct {
	// The instance ID.
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
	// Specifies whether to enable TLS (SSL) encryption. Valid values:
	//
	// 	- **Disable**: disables SSL encryption.
	//
	// 	- **Enable**: enables SSL encryption.
	//
	// 	- **Update**: updates the SSL certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	SSLEnabled    *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyInstanceSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceSSLRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceSSLRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceSSLRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceSSLRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceSSLRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceSSLRequest) GetSSLEnabled() *string {
	return s.SSLEnabled
}

func (s *ModifyInstanceSSLRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyInstanceSSLRequest) SetInstanceId(v string) *ModifyInstanceSSLRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetOwnerAccount(v string) *ModifyInstanceSSLRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetOwnerId(v int64) *ModifyInstanceSSLRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetResourceOwnerAccount(v string) *ModifyInstanceSSLRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetResourceOwnerId(v int64) *ModifyInstanceSSLRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetSSLEnabled(v string) *ModifyInstanceSSLRequest {
	s.SSLEnabled = &v
	return s
}

func (s *ModifyInstanceSSLRequest) SetSecurityToken(v string) *ModifyInstanceSSLRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyInstanceSSLRequest) Validate() error {
	return dara.Validate(s)
}
