// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMarketingFLowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityCode(v string) *UpdateMarketingFLowRequest
	GetActivityCode() *string
	SetActivityDesc(v string) *UpdateMarketingFLowRequest
	GetActivityDesc() *string
	SetActivityId(v string) *UpdateMarketingFLowRequest
	GetActivityId() *string
	SetActivityName(v string) *UpdateMarketingFLowRequest
	GetActivityName() *string
	SetCronExpression(v string) *UpdateMarketingFLowRequest
	GetCronExpression() *string
	SetEndDate(v string) *UpdateMarketingFLowRequest
	GetEndDate() *string
	SetExecutionType(v string) *UpdateMarketingFLowRequest
	GetExecutionType() *string
	SetOwnerId(v int64) *UpdateMarketingFLowRequest
	GetOwnerId() *int64
	SetParamFlag(v string) *UpdateMarketingFLowRequest
	GetParamFlag() *string
	SetParams(v map[string]interface{}) *UpdateMarketingFLowRequest
	GetParams() map[string]interface{}
	SetRelatedFlowCode(v string) *UpdateMarketingFLowRequest
	GetRelatedFlowCode() *string
	SetRelatedFlowName(v string) *UpdateMarketingFLowRequest
	GetRelatedFlowName() *string
	SetRelatedGroupId(v int64) *UpdateMarketingFLowRequest
	GetRelatedGroupId() *int64
	SetResourceOwnerAccount(v string) *UpdateMarketingFLowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateMarketingFLowRequest
	GetResourceOwnerId() *int64
	SetStartDate(v string) *UpdateMarketingFLowRequest
	GetStartDate() *string
}

type UpdateMarketingFLowRequest struct {
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
	Params map[string]interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
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

func (s UpdateMarketingFLowRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMarketingFLowRequest) GoString() string {
	return s.String()
}

func (s *UpdateMarketingFLowRequest) GetActivityCode() *string {
	return s.ActivityCode
}

func (s *UpdateMarketingFLowRequest) GetActivityDesc() *string {
	return s.ActivityDesc
}

func (s *UpdateMarketingFLowRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *UpdateMarketingFLowRequest) GetActivityName() *string {
	return s.ActivityName
}

func (s *UpdateMarketingFLowRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *UpdateMarketingFLowRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *UpdateMarketingFLowRequest) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *UpdateMarketingFLowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateMarketingFLowRequest) GetParamFlag() *string {
	return s.ParamFlag
}

func (s *UpdateMarketingFLowRequest) GetParams() map[string]interface{} {
	return s.Params
}

func (s *UpdateMarketingFLowRequest) GetRelatedFlowCode() *string {
	return s.RelatedFlowCode
}

func (s *UpdateMarketingFLowRequest) GetRelatedFlowName() *string {
	return s.RelatedFlowName
}

func (s *UpdateMarketingFLowRequest) GetRelatedGroupId() *int64 {
	return s.RelatedGroupId
}

func (s *UpdateMarketingFLowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateMarketingFLowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateMarketingFLowRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *UpdateMarketingFLowRequest) SetActivityCode(v string) *UpdateMarketingFLowRequest {
	s.ActivityCode = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetActivityDesc(v string) *UpdateMarketingFLowRequest {
	s.ActivityDesc = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetActivityId(v string) *UpdateMarketingFLowRequest {
	s.ActivityId = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetActivityName(v string) *UpdateMarketingFLowRequest {
	s.ActivityName = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetCronExpression(v string) *UpdateMarketingFLowRequest {
	s.CronExpression = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetEndDate(v string) *UpdateMarketingFLowRequest {
	s.EndDate = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetExecutionType(v string) *UpdateMarketingFLowRequest {
	s.ExecutionType = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetOwnerId(v int64) *UpdateMarketingFLowRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetParamFlag(v string) *UpdateMarketingFLowRequest {
	s.ParamFlag = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetParams(v map[string]interface{}) *UpdateMarketingFLowRequest {
	s.Params = v
	return s
}

func (s *UpdateMarketingFLowRequest) SetRelatedFlowCode(v string) *UpdateMarketingFLowRequest {
	s.RelatedFlowCode = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetRelatedFlowName(v string) *UpdateMarketingFLowRequest {
	s.RelatedFlowName = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetRelatedGroupId(v int64) *UpdateMarketingFLowRequest {
	s.RelatedGroupId = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetResourceOwnerAccount(v string) *UpdateMarketingFLowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetResourceOwnerId(v int64) *UpdateMarketingFLowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateMarketingFLowRequest) SetStartDate(v string) *UpdateMarketingFLowRequest {
	s.StartDate = &v
	return s
}

func (s *UpdateMarketingFLowRequest) Validate() error {
	return dara.Validate(s)
}
