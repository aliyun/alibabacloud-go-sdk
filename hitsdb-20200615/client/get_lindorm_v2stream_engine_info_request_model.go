// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2StreamEngineInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLindormV2StreamEngineInfoRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetLindormV2StreamEngineInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetLindormV2StreamEngineInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetLindormV2StreamEngineInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetLindormV2StreamEngineInfoRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *GetLindormV2StreamEngineInfoRequest
	GetSecurityToken() *string
}

type GetLindormV2StreamEngineInfoRequest struct {
	// This parameter is required.
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetLindormV2StreamEngineInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2StreamEngineInfoRequest) GoString() string {
	return s.String()
}

func (s *GetLindormV2StreamEngineInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLindormV2StreamEngineInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetLindormV2StreamEngineInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetLindormV2StreamEngineInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetLindormV2StreamEngineInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetLindormV2StreamEngineInfoRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetLindormV2StreamEngineInfoRequest) SetInstanceId(v string) *GetLindormV2StreamEngineInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoRequest) SetOwnerAccount(v string) *GetLindormV2StreamEngineInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoRequest) SetOwnerId(v int64) *GetLindormV2StreamEngineInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoRequest) SetResourceOwnerAccount(v string) *GetLindormV2StreamEngineInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoRequest) SetResourceOwnerId(v int64) *GetLindormV2StreamEngineInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoRequest) SetSecurityToken(v string) *GetLindormV2StreamEngineInfoRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetLindormV2StreamEngineInfoRequest) Validate() error {
	return dara.Validate(s)
}
