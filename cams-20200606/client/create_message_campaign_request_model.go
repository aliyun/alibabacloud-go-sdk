// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMessageCampaignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountId(v string) *CreateMessageCampaignRequest
	GetAdAccountId() *string
	SetBudget(v int64) *CreateMessageCampaignRequest
	GetBudget() *int64
	SetBudgetType(v string) *CreateMessageCampaignRequest
	GetBudgetType() *string
	SetCustSpaceId(v string) *CreateMessageCampaignRequest
	GetCustSpaceId() *string
	SetName(v string) *CreateMessageCampaignRequest
	GetName() *string
	SetOwnerId(v int64) *CreateMessageCampaignRequest
	GetOwnerId() *int64
	SetPageId(v string) *CreateMessageCampaignRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *CreateMessageCampaignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateMessageCampaignRequest
	GetResourceOwnerId() *int64
}

type CreateMessageCampaignRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3993**
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 300
	Budget *int64 `json:"Budget,omitempty" xml:"Budget,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// daily
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-x***
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 238***
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateMessageCampaignRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMessageCampaignRequest) GoString() string {
	return s.String()
}

func (s *CreateMessageCampaignRequest) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *CreateMessageCampaignRequest) GetBudget() *int64 {
	return s.Budget
}

func (s *CreateMessageCampaignRequest) GetBudgetType() *string {
	return s.BudgetType
}

func (s *CreateMessageCampaignRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *CreateMessageCampaignRequest) GetName() *string {
	return s.Name
}

func (s *CreateMessageCampaignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateMessageCampaignRequest) GetPageId() *string {
	return s.PageId
}

func (s *CreateMessageCampaignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateMessageCampaignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateMessageCampaignRequest) SetAdAccountId(v string) *CreateMessageCampaignRequest {
	s.AdAccountId = &v
	return s
}

func (s *CreateMessageCampaignRequest) SetBudget(v int64) *CreateMessageCampaignRequest {
	s.Budget = &v
	return s
}

func (s *CreateMessageCampaignRequest) SetBudgetType(v string) *CreateMessageCampaignRequest {
	s.BudgetType = &v
	return s
}

func (s *CreateMessageCampaignRequest) SetCustSpaceId(v string) *CreateMessageCampaignRequest {
	s.CustSpaceId = &v
	return s
}

func (s *CreateMessageCampaignRequest) SetName(v string) *CreateMessageCampaignRequest {
	s.Name = &v
	return s
}

func (s *CreateMessageCampaignRequest) SetOwnerId(v int64) *CreateMessageCampaignRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMessageCampaignRequest) SetPageId(v string) *CreateMessageCampaignRequest {
	s.PageId = &v
	return s
}

func (s *CreateMessageCampaignRequest) SetResourceOwnerAccount(v string) *CreateMessageCampaignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateMessageCampaignRequest) SetResourceOwnerId(v int64) *CreateMessageCampaignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateMessageCampaignRequest) Validate() error {
	return dara.Validate(s)
}
