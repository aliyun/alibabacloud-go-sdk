// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAudienceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountId(v string) *ListCustomAudienceShrinkRequest
	GetAdAccountId() *string
	SetCustSpaceId(v string) *ListCustomAudienceShrinkRequest
	GetCustSpaceId() *string
	SetCustomAudienceId(v string) *ListCustomAudienceShrinkRequest
	GetCustomAudienceId() *string
	SetCustomAudienceName(v string) *ListCustomAudienceShrinkRequest
	GetCustomAudienceName() *string
	SetOwnerId(v int64) *ListCustomAudienceShrinkRequest
	GetOwnerId() *int64
	SetPageShrink(v string) *ListCustomAudienceShrinkRequest
	GetPageShrink() *string
	SetPageId(v string) *ListCustomAudienceShrinkRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *ListCustomAudienceShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListCustomAudienceShrinkRequest
	GetResourceOwnerId() *int64
	SetTokenType(v string) *ListCustomAudienceShrinkRequest
	GetTokenType() *string
}

type ListCustomAudienceShrinkRequest struct {
	// example:
	//
	// 339**
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-**
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// 239**
	CustomAudienceId *string `json:"CustomAudienceId,omitempty" xml:"CustomAudienceId,omitempty"`
	// example:
	//
	// name
	CustomAudienceName *string `json:"CustomAudienceName,omitempty" xml:"CustomAudienceName,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3939**
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// custom
	TokenType *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
}

func (s ListCustomAudienceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAudienceShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCustomAudienceShrinkRequest) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *ListCustomAudienceShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListCustomAudienceShrinkRequest) GetCustomAudienceId() *string {
	return s.CustomAudienceId
}

func (s *ListCustomAudienceShrinkRequest) GetCustomAudienceName() *string {
	return s.CustomAudienceName
}

func (s *ListCustomAudienceShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListCustomAudienceShrinkRequest) GetPageShrink() *string {
	return s.PageShrink
}

func (s *ListCustomAudienceShrinkRequest) GetPageId() *string {
	return s.PageId
}

func (s *ListCustomAudienceShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListCustomAudienceShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListCustomAudienceShrinkRequest) GetTokenType() *string {
	return s.TokenType
}

func (s *ListCustomAudienceShrinkRequest) SetAdAccountId(v string) *ListCustomAudienceShrinkRequest {
	s.AdAccountId = &v
	return s
}

func (s *ListCustomAudienceShrinkRequest) SetCustSpaceId(v string) *ListCustomAudienceShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListCustomAudienceShrinkRequest) SetCustomAudienceId(v string) *ListCustomAudienceShrinkRequest {
	s.CustomAudienceId = &v
	return s
}

func (s *ListCustomAudienceShrinkRequest) SetCustomAudienceName(v string) *ListCustomAudienceShrinkRequest {
	s.CustomAudienceName = &v
	return s
}

func (s *ListCustomAudienceShrinkRequest) SetOwnerId(v int64) *ListCustomAudienceShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCustomAudienceShrinkRequest) SetPageShrink(v string) *ListCustomAudienceShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListCustomAudienceShrinkRequest) SetPageId(v string) *ListCustomAudienceShrinkRequest {
	s.PageId = &v
	return s
}

func (s *ListCustomAudienceShrinkRequest) SetResourceOwnerAccount(v string) *ListCustomAudienceShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCustomAudienceShrinkRequest) SetResourceOwnerId(v int64) *ListCustomAudienceShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCustomAudienceShrinkRequest) SetTokenType(v string) *ListCustomAudienceShrinkRequest {
	s.TokenType = &v
	return s
}

func (s *ListCustomAudienceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
