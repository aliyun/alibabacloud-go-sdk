// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMarketingFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityDesc(v string) *AddMarketingFlowShrinkRequest
	GetActivityDesc() *string
	SetActivityName(v string) *AddMarketingFlowShrinkRequest
	GetActivityName() *string
	SetBizCode(v string) *AddMarketingFlowShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *AddMarketingFlowShrinkRequest
	GetBizExtendShrink() *string
	SetCronExpression(v string) *AddMarketingFlowShrinkRequest
	GetCronExpression() *string
	SetEndDate(v string) *AddMarketingFlowShrinkRequest
	GetEndDate() *string
	SetExecutionType(v string) *AddMarketingFlowShrinkRequest
	GetExecutionType() *string
	SetOwnerId(v int64) *AddMarketingFlowShrinkRequest
	GetOwnerId() *int64
	SetParamFlag(v string) *AddMarketingFlowShrinkRequest
	GetParamFlag() *string
	SetParamsShrink(v string) *AddMarketingFlowShrinkRequest
	GetParamsShrink() *string
	SetRelatedFlowCode(v string) *AddMarketingFlowShrinkRequest
	GetRelatedFlowCode() *string
	SetRelatedFlowName(v string) *AddMarketingFlowShrinkRequest
	GetRelatedFlowName() *string
	SetRelatedGroupId(v int64) *AddMarketingFlowShrinkRequest
	GetRelatedGroupId() *int64
	SetResourceOwnerAccount(v string) *AddMarketingFlowShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddMarketingFlowShrinkRequest
	GetResourceOwnerId() *int64
	SetStartDate(v string) *AddMarketingFlowShrinkRequest
	GetStartDate() *string
}

type AddMarketingFlowShrinkRequest struct {
	// example:
	//
	// 示例值示例值
	ActivityDesc *string `json:"ActivityDesc,omitempty" xml:"ActivityDesc,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// example:
	//
	// 示例值
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	ExecutionType *string `json:"ExecutionType,omitempty" xml:"ExecutionType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 示例值
	ParamFlag *string `json:"ParamFlag,omitempty" xml:"ParamFlag,omitempty"`
	// example:
	//
	// {}
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// 示例值
	RelatedFlowCode *string `json:"RelatedFlowCode,omitempty" xml:"RelatedFlowCode,omitempty"`
	// example:
	//
	// 示例值
	RelatedFlowName *string `json:"RelatedFlowName,omitempty" xml:"RelatedFlowName,omitempty"`
	// example:
	//
	// 43
	RelatedGroupId       *int64  `json:"RelatedGroupId,omitempty" xml:"RelatedGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 示例值
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s AddMarketingFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMarketingFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddMarketingFlowShrinkRequest) GetActivityDesc() *string {
	return s.ActivityDesc
}

func (s *AddMarketingFlowShrinkRequest) GetActivityName() *string {
	return s.ActivityName
}

func (s *AddMarketingFlowShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *AddMarketingFlowShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *AddMarketingFlowShrinkRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *AddMarketingFlowShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *AddMarketingFlowShrinkRequest) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *AddMarketingFlowShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddMarketingFlowShrinkRequest) GetParamFlag() *string {
	return s.ParamFlag
}

func (s *AddMarketingFlowShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *AddMarketingFlowShrinkRequest) GetRelatedFlowCode() *string {
	return s.RelatedFlowCode
}

func (s *AddMarketingFlowShrinkRequest) GetRelatedFlowName() *string {
	return s.RelatedFlowName
}

func (s *AddMarketingFlowShrinkRequest) GetRelatedGroupId() *int64 {
	return s.RelatedGroupId
}

func (s *AddMarketingFlowShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddMarketingFlowShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddMarketingFlowShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *AddMarketingFlowShrinkRequest) SetActivityDesc(v string) *AddMarketingFlowShrinkRequest {
	s.ActivityDesc = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetActivityName(v string) *AddMarketingFlowShrinkRequest {
	s.ActivityName = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetBizCode(v string) *AddMarketingFlowShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetBizExtendShrink(v string) *AddMarketingFlowShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetCronExpression(v string) *AddMarketingFlowShrinkRequest {
	s.CronExpression = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetEndDate(v string) *AddMarketingFlowShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetExecutionType(v string) *AddMarketingFlowShrinkRequest {
	s.ExecutionType = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetOwnerId(v int64) *AddMarketingFlowShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetParamFlag(v string) *AddMarketingFlowShrinkRequest {
	s.ParamFlag = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetParamsShrink(v string) *AddMarketingFlowShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetRelatedFlowCode(v string) *AddMarketingFlowShrinkRequest {
	s.RelatedFlowCode = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetRelatedFlowName(v string) *AddMarketingFlowShrinkRequest {
	s.RelatedFlowName = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetRelatedGroupId(v int64) *AddMarketingFlowShrinkRequest {
	s.RelatedGroupId = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetResourceOwnerAccount(v string) *AddMarketingFlowShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetResourceOwnerId(v int64) *AddMarketingFlowShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) SetStartDate(v string) *AddMarketingFlowShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *AddMarketingFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
