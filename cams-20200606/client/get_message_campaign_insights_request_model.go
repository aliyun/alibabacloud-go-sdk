// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCampaignInsightsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountId(v string) *GetMessageCampaignInsightsRequest
	GetAdAccountId() *string
	SetCampaignId(v string) *GetMessageCampaignInsightsRequest
	GetCampaignId() *string
	SetCustSpaceId(v string) *GetMessageCampaignInsightsRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *GetMessageCampaignInsightsRequest
	GetOwnerId() *int64
	SetPageId(v string) *GetMessageCampaignInsightsRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *GetMessageCampaignInsightsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetMessageCampaignInsightsRequest
	GetResourceOwnerId() *int64
}

type GetMessageCampaignInsightsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23***
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23**
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-x**
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 323**
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetMessageCampaignInsightsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCampaignInsightsRequest) GoString() string {
	return s.String()
}

func (s *GetMessageCampaignInsightsRequest) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *GetMessageCampaignInsightsRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *GetMessageCampaignInsightsRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetMessageCampaignInsightsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetMessageCampaignInsightsRequest) GetPageId() *string {
	return s.PageId
}

func (s *GetMessageCampaignInsightsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetMessageCampaignInsightsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetMessageCampaignInsightsRequest) SetAdAccountId(v string) *GetMessageCampaignInsightsRequest {
	s.AdAccountId = &v
	return s
}

func (s *GetMessageCampaignInsightsRequest) SetCampaignId(v string) *GetMessageCampaignInsightsRequest {
	s.CampaignId = &v
	return s
}

func (s *GetMessageCampaignInsightsRequest) SetCustSpaceId(v string) *GetMessageCampaignInsightsRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetMessageCampaignInsightsRequest) SetOwnerId(v int64) *GetMessageCampaignInsightsRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMessageCampaignInsightsRequest) SetPageId(v string) *GetMessageCampaignInsightsRequest {
	s.PageId = &v
	return s
}

func (s *GetMessageCampaignInsightsRequest) SetResourceOwnerAccount(v string) *GetMessageCampaignInsightsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMessageCampaignInsightsRequest) SetResourceOwnerId(v int64) *GetMessageCampaignInsightsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetMessageCampaignInsightsRequest) Validate() error {
	return dara.Validate(s)
}
