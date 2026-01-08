// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMarketingFLowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityCode(v string) *UpdateMarketingFLowShrinkRequest
	GetActivityCode() *string
	SetActivityDesc(v string) *UpdateMarketingFLowShrinkRequest
	GetActivityDesc() *string
	SetActivityId(v string) *UpdateMarketingFLowShrinkRequest
	GetActivityId() *string
	SetActivityName(v string) *UpdateMarketingFLowShrinkRequest
	GetActivityName() *string
	SetCronExpression(v string) *UpdateMarketingFLowShrinkRequest
	GetCronExpression() *string
	SetEndDate(v string) *UpdateMarketingFLowShrinkRequest
	GetEndDate() *string
	SetExecutionType(v string) *UpdateMarketingFLowShrinkRequest
	GetExecutionType() *string
	SetOwnerId(v int64) *UpdateMarketingFLowShrinkRequest
	GetOwnerId() *int64
	SetParamFlag(v string) *UpdateMarketingFLowShrinkRequest
	GetParamFlag() *string
	SetParamsShrink(v string) *UpdateMarketingFLowShrinkRequest
	GetParamsShrink() *string
	SetRelatedFlowCode(v string) *UpdateMarketingFLowShrinkRequest
	GetRelatedFlowCode() *string
	SetRelatedFlowName(v string) *UpdateMarketingFLowShrinkRequest
	GetRelatedFlowName() *string
	SetRelatedGroupId(v int64) *UpdateMarketingFLowShrinkRequest
	GetRelatedGroupId() *int64
	SetResourceOwnerAccount(v string) *UpdateMarketingFLowShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateMarketingFLowShrinkRequest
	GetResourceOwnerId() *int64
	SetStartDate(v string) *UpdateMarketingFLowShrinkRequest
	GetStartDate() *string
}

type UpdateMarketingFLowShrinkRequest struct {
	// example:
	//
	// 115311507XXXX49888
	ActivityCode *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
	// example:
	//
	// 示例值
	ActivityDesc *string `json:"ActivityDesc,omitempty" xml:"ActivityDesc,omitempty"`
	// example:
	//
	// 674574575658675XX
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// test
	ActivityName   *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// 2025-11-26 09:59:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	ExecutionType *string `json:"ExecutionType,omitempty" xml:"ExecutionType,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 示例值
	ParamFlag *string `json:"ParamFlag,omitempty" xml:"ParamFlag,omitempty"`
	// example:
	//
	// {"testEmail":"wy****999@alibaba-inc.com"}
	ParamsShrink *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// cf-kr3k31**mfeir8w
	RelatedFlowCode *string `json:"RelatedFlowCode,omitempty" xml:"RelatedFlowCode,omitempty"`
	// example:
	//
	// testflow
	RelatedFlowName *string `json:"RelatedFlowName,omitempty" xml:"RelatedFlowName,omitempty"`
	// example:
	//
	// 114345654645XX
	RelatedGroupId       *int64  `json:"RelatedGroupId,omitempty" xml:"RelatedGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// 2025-11-25 09:59:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s UpdateMarketingFLowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMarketingFLowShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMarketingFLowShrinkRequest) GetActivityCode() *string {
	return s.ActivityCode
}

func (s *UpdateMarketingFLowShrinkRequest) GetActivityDesc() *string {
	return s.ActivityDesc
}

func (s *UpdateMarketingFLowShrinkRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *UpdateMarketingFLowShrinkRequest) GetActivityName() *string {
	return s.ActivityName
}

func (s *UpdateMarketingFLowShrinkRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *UpdateMarketingFLowShrinkRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *UpdateMarketingFLowShrinkRequest) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *UpdateMarketingFLowShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateMarketingFLowShrinkRequest) GetParamFlag() *string {
	return s.ParamFlag
}

func (s *UpdateMarketingFLowShrinkRequest) GetParamsShrink() *string {
	return s.ParamsShrink
}

func (s *UpdateMarketingFLowShrinkRequest) GetRelatedFlowCode() *string {
	return s.RelatedFlowCode
}

func (s *UpdateMarketingFLowShrinkRequest) GetRelatedFlowName() *string {
	return s.RelatedFlowName
}

func (s *UpdateMarketingFLowShrinkRequest) GetRelatedGroupId() *int64 {
	return s.RelatedGroupId
}

func (s *UpdateMarketingFLowShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateMarketingFLowShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateMarketingFLowShrinkRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *UpdateMarketingFLowShrinkRequest) SetActivityCode(v string) *UpdateMarketingFLowShrinkRequest {
	s.ActivityCode = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetActivityDesc(v string) *UpdateMarketingFLowShrinkRequest {
	s.ActivityDesc = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetActivityId(v string) *UpdateMarketingFLowShrinkRequest {
	s.ActivityId = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetActivityName(v string) *UpdateMarketingFLowShrinkRequest {
	s.ActivityName = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetCronExpression(v string) *UpdateMarketingFLowShrinkRequest {
	s.CronExpression = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetEndDate(v string) *UpdateMarketingFLowShrinkRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetExecutionType(v string) *UpdateMarketingFLowShrinkRequest {
	s.ExecutionType = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetOwnerId(v int64) *UpdateMarketingFLowShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetParamFlag(v string) *UpdateMarketingFLowShrinkRequest {
	s.ParamFlag = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetParamsShrink(v string) *UpdateMarketingFLowShrinkRequest {
	s.ParamsShrink = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetRelatedFlowCode(v string) *UpdateMarketingFLowShrinkRequest {
	s.RelatedFlowCode = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetRelatedFlowName(v string) *UpdateMarketingFLowShrinkRequest {
	s.RelatedFlowName = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetRelatedGroupId(v int64) *UpdateMarketingFLowShrinkRequest {
	s.RelatedGroupId = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetResourceOwnerAccount(v string) *UpdateMarketingFLowShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetResourceOwnerId(v int64) *UpdateMarketingFLowShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) SetStartDate(v string) *UpdateMarketingFLowShrinkRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateMarketingFLowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
