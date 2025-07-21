// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatappUploadAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetChatappUploadAuthorizationRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *GetChatappUploadAuthorizationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetChatappUploadAuthorizationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetChatappUploadAuthorizationRequest
	GetResourceOwnerId() *int64
}

type GetChatappUploadAuthorizationRequest struct {
	// The space ID of the RAM user within the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId          *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetChatappUploadAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatappUploadAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *GetChatappUploadAuthorizationRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetChatappUploadAuthorizationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetChatappUploadAuthorizationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetChatappUploadAuthorizationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetChatappUploadAuthorizationRequest) SetCustSpaceId(v string) *GetChatappUploadAuthorizationRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetChatappUploadAuthorizationRequest) SetOwnerId(v int64) *GetChatappUploadAuthorizationRequest {
	s.OwnerId = &v
	return s
}

func (s *GetChatappUploadAuthorizationRequest) SetResourceOwnerAccount(v string) *GetChatappUploadAuthorizationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetChatappUploadAuthorizationRequest) SetResourceOwnerId(v int64) *GetChatappUploadAuthorizationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetChatappUploadAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
