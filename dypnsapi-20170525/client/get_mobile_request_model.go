// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMobileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetMobileRequest
	GetAccessToken() *string
	SetOutId(v string) *GetMobileRequest
	GetOutId() *string
	SetOwnerId(v int64) *GetMobileRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetMobileRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetMobileRequest
	GetResourceOwnerId() *int64
}

type GetMobileRequest struct {
	// The logon token obtained by the SDK for your app.
	//
	// This parameter is required.
	//
	// example:
	//
	// Dfafdafad5422****
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// The external ID.
	//
	// example:
	//
	// 22345****
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetMobileRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMobileRequest) GoString() string {
	return s.String()
}

func (s *GetMobileRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetMobileRequest) GetOutId() *string {
	return s.OutId
}

func (s *GetMobileRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetMobileRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetMobileRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetMobileRequest) SetAccessToken(v string) *GetMobileRequest {
	s.AccessToken = &v
	return s
}

func (s *GetMobileRequest) SetOutId(v string) *GetMobileRequest {
	s.OutId = &v
	return s
}

func (s *GetMobileRequest) SetOwnerId(v int64) *GetMobileRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMobileRequest) SetResourceOwnerAccount(v string) *GetMobileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMobileRequest) SetResourceOwnerId(v int64) *GetMobileRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetMobileRequest) Validate() error {
	return dara.Validate(s)
}
