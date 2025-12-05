// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAppPlayKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SetAppPlayKeyRequest
	GetAppId() *string
	SetOwnerId(v int64) *SetAppPlayKeyRequest
	GetOwnerId() *int64
	SetPlayKey(v string) *SetAppPlayKeyRequest
	GetPlayKey() *string
	SetResourceOwnerAccount(v string) *SetAppPlayKeyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetAppPlayKeyRequest
	GetResourceOwnerId() *int64
}

type SetAppPlayKeyRequest struct {
	// example:
	//
	// app-1000231
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// yzNgTUtAl6HAuosIA
	PlayKey              *string `json:"PlayKey,omitempty" xml:"PlayKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetAppPlayKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAppPlayKeyRequest) GoString() string {
	return s.String()
}

func (s *SetAppPlayKeyRequest) GetAppId() *string {
	return s.AppId
}

func (s *SetAppPlayKeyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetAppPlayKeyRequest) GetPlayKey() *string {
	return s.PlayKey
}

func (s *SetAppPlayKeyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetAppPlayKeyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetAppPlayKeyRequest) SetAppId(v string) *SetAppPlayKeyRequest {
	s.AppId = &v
	return s
}

func (s *SetAppPlayKeyRequest) SetOwnerId(v int64) *SetAppPlayKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *SetAppPlayKeyRequest) SetPlayKey(v string) *SetAppPlayKeyRequest {
	s.PlayKey = &v
	return s
}

func (s *SetAppPlayKeyRequest) SetResourceOwnerAccount(v string) *SetAppPlayKeyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetAppPlayKeyRequest) SetResourceOwnerId(v int64) *SetAppPlayKeyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetAppPlayKeyRequest) Validate() error {
	return dara.Validate(s)
}
