// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageCampaignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountId(v string) *DeleteMessageCampaignRequest
	GetAdAccountId() *string
	SetCampaignId(v string) *DeleteMessageCampaignRequest
	GetCampaignId() *string
	SetCustSpaceId(v string) *DeleteMessageCampaignRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *DeleteMessageCampaignRequest
	GetOwnerId() *int64
	SetPageId(v string) *DeleteMessageCampaignRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *DeleteMessageCampaignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteMessageCampaignRequest
	GetResourceOwnerId() *int64
}

type DeleteMessageCampaignRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 293**
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 329**
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
	// 293***
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteMessageCampaignRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageCampaignRequest) GoString() string {
	return s.String()
}

func (s *DeleteMessageCampaignRequest) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *DeleteMessageCampaignRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *DeleteMessageCampaignRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *DeleteMessageCampaignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteMessageCampaignRequest) GetPageId() *string {
	return s.PageId
}

func (s *DeleteMessageCampaignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteMessageCampaignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteMessageCampaignRequest) SetAdAccountId(v string) *DeleteMessageCampaignRequest {
	s.AdAccountId = &v
	return s
}

func (s *DeleteMessageCampaignRequest) SetCampaignId(v string) *DeleteMessageCampaignRequest {
	s.CampaignId = &v
	return s
}

func (s *DeleteMessageCampaignRequest) SetCustSpaceId(v string) *DeleteMessageCampaignRequest {
	s.CustSpaceId = &v
	return s
}

func (s *DeleteMessageCampaignRequest) SetOwnerId(v int64) *DeleteMessageCampaignRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMessageCampaignRequest) SetPageId(v string) *DeleteMessageCampaignRequest {
	s.PageId = &v
	return s
}

func (s *DeleteMessageCampaignRequest) SetResourceOwnerAccount(v string) *DeleteMessageCampaignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteMessageCampaignRequest) SetResourceOwnerId(v int64) *DeleteMessageCampaignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteMessageCampaignRequest) Validate() error {
	return dara.Validate(s)
}
