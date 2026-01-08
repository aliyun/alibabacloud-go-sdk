// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessengerSubscriptionTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *ListMessengerSubscriptionTokenRequest
	GetCustSpaceId() *string
	SetCustomAudienceId(v string) *ListMessengerSubscriptionTokenRequest
	GetCustomAudienceId() *string
	SetLimit(v int64) *ListMessengerSubscriptionTokenRequest
	GetLimit() *int64
	SetOwnerId(v int64) *ListMessengerSubscriptionTokenRequest
	GetOwnerId() *int64
	SetPageId(v string) *ListMessengerSubscriptionTokenRequest
	GetPageId() *string
	SetPageKey(v string) *ListMessengerSubscriptionTokenRequest
	GetPageKey() *string
	SetResourceOwnerAccount(v string) *ListMessengerSubscriptionTokenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListMessengerSubscriptionTokenRequest
	GetResourceOwnerId() *int64
	SetTokenType(v string) *ListMessengerSubscriptionTokenRequest
	GetTokenType() *string
}

type ListMessengerSubscriptionTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cams-**
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// 233**
	CustomAudienceId *string `json:"CustomAudienceId,omitempty" xml:"CustomAudienceId,omitempty"`
	// example:
	//
	// 67
	Limit   *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 239***
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// example:
	//
	// 2ie**
	PageKey              *string `json:"PageKey,omitempty" xml:"PageKey,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customAudience
	TokenType *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
}

func (s ListMessengerSubscriptionTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessengerSubscriptionTokenRequest) GoString() string {
	return s.String()
}

func (s *ListMessengerSubscriptionTokenRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListMessengerSubscriptionTokenRequest) GetCustomAudienceId() *string {
	return s.CustomAudienceId
}

func (s *ListMessengerSubscriptionTokenRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ListMessengerSubscriptionTokenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListMessengerSubscriptionTokenRequest) GetPageId() *string {
	return s.PageId
}

func (s *ListMessengerSubscriptionTokenRequest) GetPageKey() *string {
	return s.PageKey
}

func (s *ListMessengerSubscriptionTokenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListMessengerSubscriptionTokenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListMessengerSubscriptionTokenRequest) GetTokenType() *string {
	return s.TokenType
}

func (s *ListMessengerSubscriptionTokenRequest) SetCustSpaceId(v string) *ListMessengerSubscriptionTokenRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListMessengerSubscriptionTokenRequest) SetCustomAudienceId(v string) *ListMessengerSubscriptionTokenRequest {
	s.CustomAudienceId = &v
	return s
}

func (s *ListMessengerSubscriptionTokenRequest) SetLimit(v int64) *ListMessengerSubscriptionTokenRequest {
	s.Limit = &v
	return s
}

func (s *ListMessengerSubscriptionTokenRequest) SetOwnerId(v int64) *ListMessengerSubscriptionTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *ListMessengerSubscriptionTokenRequest) SetPageId(v string) *ListMessengerSubscriptionTokenRequest {
	s.PageId = &v
	return s
}

func (s *ListMessengerSubscriptionTokenRequest) SetPageKey(v string) *ListMessengerSubscriptionTokenRequest {
	s.PageKey = &v
	return s
}

func (s *ListMessengerSubscriptionTokenRequest) SetResourceOwnerAccount(v string) *ListMessengerSubscriptionTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListMessengerSubscriptionTokenRequest) SetResourceOwnerId(v int64) *ListMessengerSubscriptionTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListMessengerSubscriptionTokenRequest) SetTokenType(v string) *ListMessengerSubscriptionTokenRequest {
	s.TokenType = &v
	return s
}

func (s *ListMessengerSubscriptionTokenRequest) Validate() error {
	return dara.Validate(s)
}
