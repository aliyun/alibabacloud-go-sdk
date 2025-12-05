// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPlayKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAppPlayKeyRequest
	GetAppId() *string
	SetOwnerId(v int64) *GetAppPlayKeyRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetAppPlayKeyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetAppPlayKeyRequest
	GetResourceOwnerId() *int64
}

type GetAppPlayKeyRequest struct {
	// example:
	//
	// app-1000000
	AppId                *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetAppPlayKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppPlayKeyRequest) GoString() string {
	return s.String()
}

func (s *GetAppPlayKeyRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetAppPlayKeyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetAppPlayKeyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetAppPlayKeyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetAppPlayKeyRequest) SetAppId(v string) *GetAppPlayKeyRequest {
	s.AppId = &v
	return s
}

func (s *GetAppPlayKeyRequest) SetOwnerId(v int64) *GetAppPlayKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *GetAppPlayKeyRequest) SetResourceOwnerAccount(v string) *GetAppPlayKeyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetAppPlayKeyRequest) SetResourceOwnerId(v int64) *GetAppPlayKeyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetAppPlayKeyRequest) Validate() error {
	return dara.Validate(s)
}
