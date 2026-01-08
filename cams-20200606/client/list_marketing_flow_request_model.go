// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMarketingFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityCode(v string) *ListMarketingFlowRequest
	GetActivityCode() *string
	SetActivityName(v string) *ListMarketingFlowRequest
	GetActivityName() *string
	SetActivityStatus(v string) *ListMarketingFlowRequest
	GetActivityStatus() *string
	SetBizCode(v string) *ListMarketingFlowRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *ListMarketingFlowRequest
	GetBizExtend() map[string]interface{}
	SetOwnerId(v int64) *ListMarketingFlowRequest
	GetOwnerId() *int64
	SetPageIndex(v string) *ListMarketingFlowRequest
	GetPageIndex() *string
	SetPageSize(v string) *ListMarketingFlowRequest
	GetPageSize() *string
	SetRelatedFlowCode(v string) *ListMarketingFlowRequest
	GetRelatedFlowCode() *string
	SetRelatedGroupId(v int64) *ListMarketingFlowRequest
	GetRelatedGroupId() *int64
	SetResourceOwnerAccount(v string) *ListMarketingFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListMarketingFlowRequest
	GetResourceOwnerId() *int64
}

type ListMarketingFlowRequest struct {
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
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	OwnerId   *int64                 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s ListMarketingFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMarketingFlowRequest) GoString() string {
	return s.String()
}

func (s *ListMarketingFlowRequest) GetActivityCode() *string {
	return s.ActivityCode
}

func (s *ListMarketingFlowRequest) GetActivityName() *string {
	return s.ActivityName
}

func (s *ListMarketingFlowRequest) GetActivityStatus() *string {
	return s.ActivityStatus
}

func (s *ListMarketingFlowRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *ListMarketingFlowRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *ListMarketingFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListMarketingFlowRequest) GetPageIndex() *string {
	return s.PageIndex
}

func (s *ListMarketingFlowRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListMarketingFlowRequest) GetRelatedFlowCode() *string {
	return s.RelatedFlowCode
}

func (s *ListMarketingFlowRequest) GetRelatedGroupId() *int64 {
	return s.RelatedGroupId
}

func (s *ListMarketingFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListMarketingFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListMarketingFlowRequest) SetActivityCode(v string) *ListMarketingFlowRequest {
	s.ActivityCode = &v
	return s
}

func (s *ListMarketingFlowRequest) SetActivityName(v string) *ListMarketingFlowRequest {
	s.ActivityName = &v
	return s
}

func (s *ListMarketingFlowRequest) SetActivityStatus(v string) *ListMarketingFlowRequest {
	s.ActivityStatus = &v
	return s
}

func (s *ListMarketingFlowRequest) SetBizCode(v string) *ListMarketingFlowRequest {
	s.BizCode = &v
	return s
}

func (s *ListMarketingFlowRequest) SetBizExtend(v map[string]interface{}) *ListMarketingFlowRequest {
	s.BizExtend = v
	return s
}

func (s *ListMarketingFlowRequest) SetOwnerId(v int64) *ListMarketingFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *ListMarketingFlowRequest) SetPageIndex(v string) *ListMarketingFlowRequest {
	s.PageIndex = &v
	return s
}

func (s *ListMarketingFlowRequest) SetPageSize(v string) *ListMarketingFlowRequest {
	s.PageSize = &v
	return s
}

func (s *ListMarketingFlowRequest) SetRelatedFlowCode(v string) *ListMarketingFlowRequest {
	s.RelatedFlowCode = &v
	return s
}

func (s *ListMarketingFlowRequest) SetRelatedGroupId(v int64) *ListMarketingFlowRequest {
	s.RelatedGroupId = &v
	return s
}

func (s *ListMarketingFlowRequest) SetResourceOwnerAccount(v string) *ListMarketingFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListMarketingFlowRequest) SetResourceOwnerId(v int64) *ListMarketingFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListMarketingFlowRequest) Validate() error {
	return dara.Validate(s)
}
