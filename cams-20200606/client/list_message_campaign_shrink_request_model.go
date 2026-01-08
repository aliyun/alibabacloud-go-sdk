// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageCampaignShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountId(v string) *ListMessageCampaignShrinkRequest
	GetAdAccountId() *string
	SetCampaignId(v string) *ListMessageCampaignShrinkRequest
	GetCampaignId() *string
	SetCampaignName(v string) *ListMessageCampaignShrinkRequest
	GetCampaignName() *string
	SetCustSpaceId(v string) *ListMessageCampaignShrinkRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ListMessageCampaignShrinkRequest
	GetOwnerId() *int64
	SetPageShrink(v string) *ListMessageCampaignShrinkRequest
	GetPageShrink() *string
	SetPageId(v string) *ListMessageCampaignShrinkRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *ListMessageCampaignShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListMessageCampaignShrinkRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ListMessageCampaignShrinkRequest
	GetStatus() *string
}

type ListMessageCampaignShrinkRequest struct {
	// example:
	//
	// 239**
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// example:
	//
	// 233**
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// example:
	//
	// test-campaign
	CampaignName *string `json:"CampaignName,omitempty" xml:"CampaignName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-**
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	PageShrink *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 239***
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMessageCampaignShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessageCampaignShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMessageCampaignShrinkRequest) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *ListMessageCampaignShrinkRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ListMessageCampaignShrinkRequest) GetCampaignName() *string {
	return s.CampaignName
}

func (s *ListMessageCampaignShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListMessageCampaignShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListMessageCampaignShrinkRequest) GetPageShrink() *string {
	return s.PageShrink
}

func (s *ListMessageCampaignShrinkRequest) GetPageId() *string {
	return s.PageId
}

func (s *ListMessageCampaignShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListMessageCampaignShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListMessageCampaignShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMessageCampaignShrinkRequest) SetAdAccountId(v string) *ListMessageCampaignShrinkRequest {
	s.AdAccountId = &v
	return s
}

func (s *ListMessageCampaignShrinkRequest) SetCampaignId(v string) *ListMessageCampaignShrinkRequest {
	s.CampaignId = &v
	return s
}

func (s *ListMessageCampaignShrinkRequest) SetCampaignName(v string) *ListMessageCampaignShrinkRequest {
	s.CampaignName = &v
	return s
}

func (s *ListMessageCampaignShrinkRequest) SetCustSpaceId(v string) *ListMessageCampaignShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListMessageCampaignShrinkRequest) SetOwnerId(v int64) *ListMessageCampaignShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListMessageCampaignShrinkRequest) SetPageShrink(v string) *ListMessageCampaignShrinkRequest {
	s.PageShrink = &v
	return s
}

func (s *ListMessageCampaignShrinkRequest) SetPageId(v string) *ListMessageCampaignShrinkRequest {
	s.PageId = &v
	return s
}

func (s *ListMessageCampaignShrinkRequest) SetResourceOwnerAccount(v string) *ListMessageCampaignShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListMessageCampaignShrinkRequest) SetResourceOwnerId(v int64) *ListMessageCampaignShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListMessageCampaignShrinkRequest) SetStatus(v string) *ListMessageCampaignShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListMessageCampaignShrinkRequest) Validate() error {
	return dara.Validate(s)
}
