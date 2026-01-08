// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncMessengerSubscriptionTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *SyncMessengerSubscriptionTokenRequest
	GetCustSpaceId() *string
	SetCustomAudienceId(v string) *SyncMessengerSubscriptionTokenRequest
	GetCustomAudienceId() *string
	SetOwnerId(v int64) *SyncMessengerSubscriptionTokenRequest
	GetOwnerId() *int64
	SetPageId(v string) *SyncMessengerSubscriptionTokenRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *SyncMessengerSubscriptionTokenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SyncMessengerSubscriptionTokenRequest
	GetResourceOwnerId() *int64
	SetTokenType(v string) *SyncMessengerSubscriptionTokenRequest
	GetTokenType() *string
}

type SyncMessengerSubscriptionTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cams-xi****
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// example:
	//
	// 20399***
	CustomAudienceId *string `json:"CustomAudienceId,omitempty" xml:"CustomAudienceId,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2030***
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// customAudience
	TokenType *string `json:"TokenType,omitempty" xml:"TokenType,omitempty"`
}

func (s SyncMessengerSubscriptionTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncMessengerSubscriptionTokenRequest) GoString() string {
	return s.String()
}

func (s *SyncMessengerSubscriptionTokenRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *SyncMessengerSubscriptionTokenRequest) GetCustomAudienceId() *string {
	return s.CustomAudienceId
}

func (s *SyncMessengerSubscriptionTokenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SyncMessengerSubscriptionTokenRequest) GetPageId() *string {
	return s.PageId
}

func (s *SyncMessengerSubscriptionTokenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SyncMessengerSubscriptionTokenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SyncMessengerSubscriptionTokenRequest) GetTokenType() *string {
	return s.TokenType
}

func (s *SyncMessengerSubscriptionTokenRequest) SetCustSpaceId(v string) *SyncMessengerSubscriptionTokenRequest {
	s.CustSpaceId = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenRequest) SetCustomAudienceId(v string) *SyncMessengerSubscriptionTokenRequest {
	s.CustomAudienceId = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenRequest) SetOwnerId(v int64) *SyncMessengerSubscriptionTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenRequest) SetPageId(v string) *SyncMessengerSubscriptionTokenRequest {
	s.PageId = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenRequest) SetResourceOwnerAccount(v string) *SyncMessengerSubscriptionTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenRequest) SetResourceOwnerId(v int64) *SyncMessengerSubscriptionTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenRequest) SetTokenType(v string) *SyncMessengerSubscriptionTokenRequest {
	s.TokenType = &v
	return s
}

func (s *SyncMessengerSubscriptionTokenRequest) Validate() error {
	return dara.Validate(s)
}
