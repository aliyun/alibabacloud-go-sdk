// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenLdpsColumnarIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OpenLdpsColumnarIndexRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *OpenLdpsColumnarIndexRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *OpenLdpsColumnarIndexRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *OpenLdpsColumnarIndexRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OpenLdpsColumnarIndexRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *OpenLdpsColumnarIndexRequest
	GetSecurityToken() *string
}

type OpenLdpsColumnarIndexRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s OpenLdpsColumnarIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenLdpsColumnarIndexRequest) GoString() string {
	return s.String()
}

func (s *OpenLdpsColumnarIndexRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OpenLdpsColumnarIndexRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *OpenLdpsColumnarIndexRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenLdpsColumnarIndexRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenLdpsColumnarIndexRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OpenLdpsColumnarIndexRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *OpenLdpsColumnarIndexRequest) SetInstanceId(v string) *OpenLdpsColumnarIndexRequest {
	s.InstanceId = &v
	return s
}

func (s *OpenLdpsColumnarIndexRequest) SetOwnerAccount(v string) *OpenLdpsColumnarIndexRequest {
	s.OwnerAccount = &v
	return s
}

func (s *OpenLdpsColumnarIndexRequest) SetOwnerId(v int64) *OpenLdpsColumnarIndexRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenLdpsColumnarIndexRequest) SetResourceOwnerAccount(v string) *OpenLdpsColumnarIndexRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenLdpsColumnarIndexRequest) SetResourceOwnerId(v int64) *OpenLdpsColumnarIndexRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenLdpsColumnarIndexRequest) SetSecurityToken(v string) *OpenLdpsColumnarIndexRequest {
	s.SecurityToken = &v
	return s
}

func (s *OpenLdpsColumnarIndexRequest) Validate() error {
	return dara.Validate(s)
}
