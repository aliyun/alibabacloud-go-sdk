// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageCampaignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountId(v string) *ListMessageCampaignRequest
	GetAdAccountId() *string
	SetCampaignId(v string) *ListMessageCampaignRequest
	GetCampaignId() *string
	SetCampaignName(v string) *ListMessageCampaignRequest
	GetCampaignName() *string
	SetCustSpaceId(v string) *ListMessageCampaignRequest
	GetCustSpaceId() *string
	SetOwnerId(v int64) *ListMessageCampaignRequest
	GetOwnerId() *int64
	SetPage(v *ListMessageCampaignRequestPage) *ListMessageCampaignRequest
	GetPage() *ListMessageCampaignRequestPage
	SetPageId(v string) *ListMessageCampaignRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *ListMessageCampaignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListMessageCampaignRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *ListMessageCampaignRequest
	GetStatus() *string
}

type ListMessageCampaignRequest struct {
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
	Page *ListMessageCampaignRequestPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
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

func (s ListMessageCampaignRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessageCampaignRequest) GoString() string {
	return s.String()
}

func (s *ListMessageCampaignRequest) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *ListMessageCampaignRequest) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ListMessageCampaignRequest) GetCampaignName() *string {
	return s.CampaignName
}

func (s *ListMessageCampaignRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *ListMessageCampaignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListMessageCampaignRequest) GetPage() *ListMessageCampaignRequestPage {
	return s.Page
}

func (s *ListMessageCampaignRequest) GetPageId() *string {
	return s.PageId
}

func (s *ListMessageCampaignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListMessageCampaignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListMessageCampaignRequest) GetStatus() *string {
	return s.Status
}

func (s *ListMessageCampaignRequest) SetAdAccountId(v string) *ListMessageCampaignRequest {
	s.AdAccountId = &v
	return s
}

func (s *ListMessageCampaignRequest) SetCampaignId(v string) *ListMessageCampaignRequest {
	s.CampaignId = &v
	return s
}

func (s *ListMessageCampaignRequest) SetCampaignName(v string) *ListMessageCampaignRequest {
	s.CampaignName = &v
	return s
}

func (s *ListMessageCampaignRequest) SetCustSpaceId(v string) *ListMessageCampaignRequest {
	s.CustSpaceId = &v
	return s
}

func (s *ListMessageCampaignRequest) SetOwnerId(v int64) *ListMessageCampaignRequest {
	s.OwnerId = &v
	return s
}

func (s *ListMessageCampaignRequest) SetPage(v *ListMessageCampaignRequestPage) *ListMessageCampaignRequest {
	s.Page = v
	return s
}

func (s *ListMessageCampaignRequest) SetPageId(v string) *ListMessageCampaignRequest {
	s.PageId = &v
	return s
}

func (s *ListMessageCampaignRequest) SetResourceOwnerAccount(v string) *ListMessageCampaignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListMessageCampaignRequest) SetResourceOwnerId(v int64) *ListMessageCampaignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListMessageCampaignRequest) SetStatus(v string) *ListMessageCampaignRequest {
	s.Status = &v
	return s
}

func (s *ListMessageCampaignRequest) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMessageCampaignRequestPage struct {
	// This parameter is required.
	//
	// example:
	//
	// 7
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 55
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListMessageCampaignRequestPage) String() string {
	return dara.Prettify(s)
}

func (s ListMessageCampaignRequestPage) GoString() string {
	return s.String()
}

func (s *ListMessageCampaignRequestPage) GetIndex() *int64 {
	return s.Index
}

func (s *ListMessageCampaignRequestPage) GetSize() *int64 {
	return s.Size
}

func (s *ListMessageCampaignRequestPage) SetIndex(v int64) *ListMessageCampaignRequestPage {
	s.Index = &v
	return s
}

func (s *ListMessageCampaignRequestPage) SetSize(v int64) *ListMessageCampaignRequestPage {
	s.Size = &v
	return s
}

func (s *ListMessageCampaignRequestPage) Validate() error {
	return dara.Validate(s)
}
