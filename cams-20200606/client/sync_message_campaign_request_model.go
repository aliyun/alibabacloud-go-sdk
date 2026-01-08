// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncMessageCampaignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountId(v string) *SyncMessageCampaignRequest
	GetAdAccountId() *string
	SetOwnerId(v int64) *SyncMessageCampaignRequest
	GetOwnerId() *int64
	SetPageId(v string) *SyncMessageCampaignRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *SyncMessageCampaignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SyncMessageCampaignRequest
	GetResourceOwnerId() *int64
}

type SyncMessageCampaignRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 244**
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 238**
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SyncMessageCampaignRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncMessageCampaignRequest) GoString() string {
	return s.String()
}

func (s *SyncMessageCampaignRequest) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *SyncMessageCampaignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SyncMessageCampaignRequest) GetPageId() *string {
	return s.PageId
}

func (s *SyncMessageCampaignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SyncMessageCampaignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SyncMessageCampaignRequest) SetAdAccountId(v string) *SyncMessageCampaignRequest {
	s.AdAccountId = &v
	return s
}

func (s *SyncMessageCampaignRequest) SetOwnerId(v int64) *SyncMessageCampaignRequest {
	s.OwnerId = &v
	return s
}

func (s *SyncMessageCampaignRequest) SetPageId(v string) *SyncMessageCampaignRequest {
	s.PageId = &v
	return s
}

func (s *SyncMessageCampaignRequest) SetResourceOwnerAccount(v string) *SyncMessageCampaignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SyncMessageCampaignRequest) SetResourceOwnerId(v int64) *SyncMessageCampaignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SyncMessageCampaignRequest) Validate() error {
	return dara.Validate(s)
}
