// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFbInstagramPagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetFbInstagramPagesRequest
	GetAccessToken() *string
	SetInstanceId(v string) *GetFbInstagramPagesRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *GetFbInstagramPagesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetFbInstagramPagesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetFbInstagramPagesRequest
	GetResourceOwnerId() *int64
}

type GetFbInstagramPagesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ftete3254***
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// inst-8t47hh***5pyj0ps
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetFbInstagramPagesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFbInstagramPagesRequest) GoString() string {
	return s.String()
}

func (s *GetFbInstagramPagesRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetFbInstagramPagesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetFbInstagramPagesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetFbInstagramPagesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetFbInstagramPagesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetFbInstagramPagesRequest) SetAccessToken(v string) *GetFbInstagramPagesRequest {
	s.AccessToken = &v
	return s
}

func (s *GetFbInstagramPagesRequest) SetInstanceId(v string) *GetFbInstagramPagesRequest {
	s.InstanceId = &v
	return s
}

func (s *GetFbInstagramPagesRequest) SetOwnerId(v int64) *GetFbInstagramPagesRequest {
	s.OwnerId = &v
	return s
}

func (s *GetFbInstagramPagesRequest) SetResourceOwnerAccount(v string) *GetFbInstagramPagesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetFbInstagramPagesRequest) SetResourceOwnerId(v int64) *GetFbInstagramPagesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetFbInstagramPagesRequest) Validate() error {
	return dara.Validate(s)
}
