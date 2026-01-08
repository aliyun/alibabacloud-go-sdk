// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessengerPageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountIdsShrink(v string) *CreateMessengerPageShrinkRequest
	GetAdAccountIdsShrink() *string
	SetAuthenticationCode(v string) *CreateMessengerPageShrinkRequest
	GetAuthenticationCode() *string
	SetBusinessId(v string) *CreateMessengerPageShrinkRequest
	GetBusinessId() *string
	SetCustSpaceId(v string) *CreateMessengerPageShrinkRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *CreateMessengerPageShrinkRequest
	GetOwnerId() *int64
	SetPageId(v string) *CreateMessengerPageShrinkRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *CreateMessengerPageShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateMessengerPageShrinkRequest
	GetResourceOwnerId() *int64
}

type CreateMessengerPageShrinkRequest struct {
	// This parameter is required.
	AdAccountIdsShrink *string `json:"AdAccountIds,omitempty" xml:"AdAccountIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// EAA****
	AuthenticationCode *string `json:"AuthenticationCode,omitempty" xml:"AuthenticationCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 293***
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-s***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 19283***
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateMessengerPageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMessengerPageShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMessengerPageShrinkRequest) GetAdAccountIdsShrink() *string {
	return s.AdAccountIdsShrink
}

func (s *CreateMessengerPageShrinkRequest) GetAuthenticationCode() *string {
	return s.AuthenticationCode
}

func (s *CreateMessengerPageShrinkRequest) GetBusinessId() *string {
	return s.BusinessId
}

func (s *CreateMessengerPageShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *CreateMessengerPageShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateMessengerPageShrinkRequest) GetPageId() *string {
	return s.PageId
}

func (s *CreateMessengerPageShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateMessengerPageShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateMessengerPageShrinkRequest) SetAdAccountIdsShrink(v string) *CreateMessengerPageShrinkRequest {
	s.AdAccountIdsShrink = &v
	return s
}

func (s *CreateMessengerPageShrinkRequest) SetAuthenticationCode(v string) *CreateMessengerPageShrinkRequest {
	s.AuthenticationCode = &v
	return s
}

func (s *CreateMessengerPageShrinkRequest) SetBusinessId(v string) *CreateMessengerPageShrinkRequest {
	s.BusinessId = &v
	return s
}

func (s *CreateMessengerPageShrinkRequest) SetCustSpaceId(v string) *CreateMessengerPageShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateMessengerPageShrinkRequest) SetOwnerId(v int64) *CreateMessengerPageShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMessengerPageShrinkRequest) SetPageId(v string) *CreateMessengerPageShrinkRequest {
	s.PageId = &v
	return s
}

func (s *CreateMessengerPageShrinkRequest) SetResourceOwnerAccount(v string) *CreateMessengerPageShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMessengerPageShrinkRequest) SetResourceOwnerId(v int64) *CreateMessengerPageShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMessengerPageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
