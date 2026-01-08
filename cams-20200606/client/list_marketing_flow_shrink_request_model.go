// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMarketingFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityCode(v string) *ListMarketingFlowShrinkRequest
	GetActivityCode() *string
	SetActivityName(v string) *ListMarketingFlowShrinkRequest
	GetActivityName() *string
	SetActivityStatus(v string) *ListMarketingFlowShrinkRequest
	GetActivityStatus() *string
	SetBizCode(v string) *ListMarketingFlowShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *ListMarketingFlowShrinkRequest
	GetBizExtendShrink() *string
	SetOwnerId(v int64) *ListMarketingFlowShrinkRequest
	GetOwnerId() *int64
	SetPageIndex(v string) *ListMarketingFlowShrinkRequest
	GetPageIndex() *string
	SetPageSize(v string) *ListMarketingFlowShrinkRequest
	GetPageSize() *string
	SetRelatedFlowCode(v string) *ListMarketingFlowShrinkRequest
	GetRelatedFlowCode() *string
	SetRelatedGroupId(v int64) *ListMarketingFlowShrinkRequest
	GetRelatedGroupId() *int64
	SetResourceOwnerAccount(v string) *ListMarketingFlowShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListMarketingFlowShrinkRequest
	GetResourceOwnerId() *int64
}

type ListMarketingFlowShrinkRequest struct {
	// example:
	//
	// rewrwerw
	ActivityCode *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
	// example:
	//
	// werewew
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// example:
	//
	// sucess
	ActivityStatus *string `json:"ActivityStatus,omitempty" xml:"ActivityStatus,omitempty"`
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageIndex *string `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 1
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// asdfsdfdsfs
	RelatedFlowCode *string `json:"RelatedFlowCode,omitempty" xml:"RelatedFlowCode,omitempty"`
	// example:
	//
	// 68
	RelatedGroupId       *int64  `json:"RelatedGroupId,omitempty" xml:"RelatedGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListMarketingFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMarketingFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMarketingFlowShrinkRequest) GetActivityCode() *string {
	return s.ActivityCode
}

func (s *ListMarketingFlowShrinkRequest) GetActivityName() *string {
	return s.ActivityName
}

func (s *ListMarketingFlowShrinkRequest) GetActivityStatus() *string {
	return s.ActivityStatus
}

func (s *ListMarketingFlowShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *ListMarketingFlowShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *ListMarketingFlowShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListMarketingFlowShrinkRequest) GetPageIndex() *string {
	return s.PageIndex
}

func (s *ListMarketingFlowShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListMarketingFlowShrinkRequest) GetRelatedFlowCode() *string {
	return s.RelatedFlowCode
}

func (s *ListMarketingFlowShrinkRequest) GetRelatedGroupId() *int64 {
	return s.RelatedGroupId
}

func (s *ListMarketingFlowShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListMarketingFlowShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListMarketingFlowShrinkRequest) SetActivityCode(v string) *ListMarketingFlowShrinkRequest {
	s.ActivityCode = &v
	return s
}

func (s *ListMarketingFlowShrinkRequest) SetActivityName(v string) *ListMarketingFlowShrinkRequest {
	s.ActivityName = &v
	return s
}

func (s *ListMarketingFlowShrinkRequest) SetActivityStatus(v string) *ListMarketingFlowShrinkRequest {
	s.ActivityStatus = &v
	return s
}

func (s *ListMarketingFlowShrinkRequest) SetBizCode(v string) *ListMarketingFlowShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *ListMarketingFlowShrinkRequest) SetBizExtendShrink(v string) *ListMarketingFlowShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *ListMarketingFlowShrinkRequest) SetOwnerId(v int64) *ListMarketingFlowShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListMarketingFlowShrinkRequest) SetPageIndex(v string) *ListMarketingFlowShrinkRequest {
	s.PageIndex = &v
	return s
}

func (s *ListMarketingFlowShrinkRequest) SetPageSize(v string) *ListMarketingFlowShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListMarketingFlowShrinkRequest) SetRelatedFlowCode(v string) *ListMarketingFlowShrinkRequest {
	s.RelatedFlowCode = &v
	return s
}

func (s *ListMarketingFlowShrinkRequest) SetRelatedGroupId(v int64) *ListMarketingFlowShrinkRequest {
	s.RelatedGroupId = &v
	return s
}

func (s *ListMarketingFlowShrinkRequest) SetResourceOwnerAccount(v string) *ListMarketingFlowShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListMarketingFlowShrinkRequest) SetResourceOwnerId(v int64) *ListMarketingFlowShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListMarketingFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
