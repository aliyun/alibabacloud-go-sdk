// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomAudienceUserShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdAccountId(v string) *AddCustomAudienceUserShrinkRequest
	GetAdAccountId() *string
	SetBatchLastFlag(v bool) *AddCustomAudienceUserShrinkRequest
	GetBatchLastFlag() *bool
	SetCustSpaceId(v string) *AddCustomAudienceUserShrinkRequest
	GetCustSpaceId() *string
	SetCustomAudienceId(v string) *AddCustomAudienceUserShrinkRequest
	GetCustomAudienceId() *string
	SetEstimatedNumTotal(v int64) *AddCustomAudienceUserShrinkRequest
	GetEstimatedNumTotal() *int64
	SetOwnerId(v int64) *AddCustomAudienceUserShrinkRequest
	GetOwnerId() *int64
	SetPageId(v string) *AddCustomAudienceUserShrinkRequest
	GetPageId() *string
	SetResourceOwnerAccount(v string) *AddCustomAudienceUserShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddCustomAudienceUserShrinkRequest
	GetResourceOwnerId() *int64
	SetUsersShrink(v string) *AddCustomAudienceUserShrinkRequest
	GetUsersShrink() *string
}

type AddCustomAudienceUserShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3939**
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// example:
	//
	// false
	BatchLastFlag *bool `json:"BatchLastFlag,omitempty" xml:"BatchLastFlag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cams-**
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 399**
	CustomAudienceId *string `json:"CustomAudienceId,omitempty" xml:"CustomAudienceId,omitempty"`
	// example:
	//
	// 26
	EstimatedNumTotal *int64 `json:"EstimatedNumTotal,omitempty" xml:"EstimatedNumTotal,omitempty"`
	OwnerId           *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 239**
	PageId               *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	UsersShrink *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s AddCustomAudienceUserShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomAudienceUserShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddCustomAudienceUserShrinkRequest) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *AddCustomAudienceUserShrinkRequest) GetBatchLastFlag() *bool {
	return s.BatchLastFlag
}

func (s *AddCustomAudienceUserShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *AddCustomAudienceUserShrinkRequest) GetCustomAudienceId() *string {
	return s.CustomAudienceId
}

func (s *AddCustomAudienceUserShrinkRequest) GetEstimatedNumTotal() *int64 {
	return s.EstimatedNumTotal
}

func (s *AddCustomAudienceUserShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddCustomAudienceUserShrinkRequest) GetPageId() *string {
	return s.PageId
}

func (s *AddCustomAudienceUserShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddCustomAudienceUserShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddCustomAudienceUserShrinkRequest) GetUsersShrink() *string {
	return s.UsersShrink
}

func (s *AddCustomAudienceUserShrinkRequest) SetAdAccountId(v string) *AddCustomAudienceUserShrinkRequest {
	s.AdAccountId = &v
	return s
}

func (s *AddCustomAudienceUserShrinkRequest) SetBatchLastFlag(v bool) *AddCustomAudienceUserShrinkRequest {
	s.BatchLastFlag = &v
	return s
}

func (s *AddCustomAudienceUserShrinkRequest) SetCustSpaceId(v string) *AddCustomAudienceUserShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *AddCustomAudienceUserShrinkRequest) SetCustomAudienceId(v string) *AddCustomAudienceUserShrinkRequest {
	s.CustomAudienceId = &v
	return s
}

func (s *AddCustomAudienceUserShrinkRequest) SetEstimatedNumTotal(v int64) *AddCustomAudienceUserShrinkRequest {
	s.EstimatedNumTotal = &v
	return s
}

func (s *AddCustomAudienceUserShrinkRequest) SetOwnerId(v int64) *AddCustomAudienceUserShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddCustomAudienceUserShrinkRequest) SetPageId(v string) *AddCustomAudienceUserShrinkRequest {
	s.PageId = &v
	return s
}

func (s *AddCustomAudienceUserShrinkRequest) SetResourceOwnerAccount(v string) *AddCustomAudienceUserShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddCustomAudienceUserShrinkRequest) SetResourceOwnerId(v int64) *AddCustomAudienceUserShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddCustomAudienceUserShrinkRequest) SetUsersShrink(v string) *AddCustomAudienceUserShrinkRequest {
	s.UsersShrink = &v
	return s
}

func (s *AddCustomAudienceUserShrinkRequest) Validate() error {
	return dara.Validate(s)
}
