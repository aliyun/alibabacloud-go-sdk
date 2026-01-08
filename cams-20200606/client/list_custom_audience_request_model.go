// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAudienceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountId(v string) *ListCustomAudienceRequest
	GetAdAccountId() *string
	SetCustSpaceId(v string) *ListCustomAudienceRequest
	GetCustSpaceId() *string
	SetCustomAudienceId(v string) *ListCustomAudienceRequest
	GetCustomAudienceId() *string
	SetCustomAudienceName(v string) *ListCustomAudienceRequest
	GetCustomAudienceName() *string
	SetOwnerId(v int64) *ListCustomAudienceRequest
	GetOwnerId() *int64
	SetPage(v *ListCustomAudienceRequestPage) *ListCustomAudienceRequest
	GetPage() *ListCustomAudienceRequestPage
	SetPageId(v string) *ListCustomAudienceRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *ListCustomAudienceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListCustomAudienceRequest
	GetResourceOwnerId() *int64
	SetTokenType(v string) *ListCustomAudienceRequest
	GetTokenType() *string
}

type ListCustomAudienceRequest struct {
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
	Page *ListCustomAudienceRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
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

func (s ListCustomAudienceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAudienceRequest) GoString() string {
	return s.String()
}

func (s *ListCustomAudienceRequest) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *ListCustomAudienceRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListCustomAudienceRequest) GetCustomAudienceId() *string {
	return s.CustomAudienceId
}

func (s *ListCustomAudienceRequest) GetCustomAudienceName() *string {
	return s.CustomAudienceName
}

func (s *ListCustomAudienceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListCustomAudienceRequest) GetPage() *ListCustomAudienceRequestPage {
	return s.Page
}

func (s *ListCustomAudienceRequest) GetPageId() *string {
	return s.PageId
}

func (s *ListCustomAudienceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListCustomAudienceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListCustomAudienceRequest) GetTokenType() *string {
	return s.TokenType
}

func (s *ListCustomAudienceRequest) SetAdAccountId(v string) *ListCustomAudienceRequest {
	s.AdAccountId = &v
	return s
}

func (s *ListCustomAudienceRequest) SetCustSpaceId(v string) *ListCustomAudienceRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListCustomAudienceRequest) SetCustomAudienceId(v string) *ListCustomAudienceRequest {
	s.CustomAudienceId = &v
	return s
}

func (s *ListCustomAudienceRequest) SetCustomAudienceName(v string) *ListCustomAudienceRequest {
	s.CustomAudienceName = &v
	return s
}

func (s *ListCustomAudienceRequest) SetOwnerId(v int64) *ListCustomAudienceRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCustomAudienceRequest) SetPage(v *ListCustomAudienceRequestPage) *ListCustomAudienceRequest {
	s.Page = v
	return s
}

func (s *ListCustomAudienceRequest) SetPageId(v string) *ListCustomAudienceRequest {
	s.PageId = &v
	return s
}

func (s *ListCustomAudienceRequest) SetResourceOwnerAccount(v string) *ListCustomAudienceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCustomAudienceRequest) SetResourceOwnerId(v int64) *ListCustomAudienceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCustomAudienceRequest) SetTokenType(v string) *ListCustomAudienceRequest {
	s.TokenType = &v
	return s
}

func (s *ListCustomAudienceRequest) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomAudienceRequestPage struct {
	// This parameter is required.
	//
	// example:
	//
	// 40
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 95
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListCustomAudienceRequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAudienceRequestPage) GoString() string {
	return s.String()
}

func (s *ListCustomAudienceRequestPage) GetIndex() *int64 {
	return s.Index
}

func (s *ListCustomAudienceRequestPage) GetSize() *int64 {
	return s.Size
}

func (s *ListCustomAudienceRequestPage) SetIndex(v int64) *ListCustomAudienceRequestPage {
	s.Index = &v
	return s
}

func (s *ListCustomAudienceRequestPage) SetSize(v int64) *ListCustomAudienceRequestPage {
	s.Size = &v
	return s
}

func (s *ListCustomAudienceRequestPage) Validate() error {
	return dara.Validate(s)
}
