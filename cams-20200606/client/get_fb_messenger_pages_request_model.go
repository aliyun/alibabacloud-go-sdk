// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFbMessengerPagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetFbMessengerPagesRequest
	GetAccessToken() *string
	SetInstanceId(v string) *GetFbMessengerPagesRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *GetFbMessengerPagesRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetFbMessengerPagesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetFbMessengerPagesRequest
	GetResourceOwnerId() *int64
}

type GetFbMessengerPagesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// EARDGFDGF12****
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

func (s GetFbMessengerPagesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFbMessengerPagesRequest) GoString() string {
	return s.String()
}

func (s *GetFbMessengerPagesRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetFbMessengerPagesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetFbMessengerPagesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetFbMessengerPagesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetFbMessengerPagesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetFbMessengerPagesRequest) SetAccessToken(v string) *GetFbMessengerPagesRequest {
	s.AccessToken = &v
	return s
}

func (s *GetFbMessengerPagesRequest) SetInstanceId(v string) *GetFbMessengerPagesRequest {
	s.InstanceId = &v
	return s
}

func (s *GetFbMessengerPagesRequest) SetOwnerId(v int64) *GetFbMessengerPagesRequest {
	s.OwnerId = &v
	return s
}

func (s *GetFbMessengerPagesRequest) SetResourceOwnerAccount(v string) *GetFbMessengerPagesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetFbMessengerPagesRequest) SetResourceOwnerId(v int64) *GetFbMessengerPagesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetFbMessengerPagesRequest) Validate() error {
	return dara.Validate(s)
}
