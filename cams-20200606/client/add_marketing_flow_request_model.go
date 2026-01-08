// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMarketingFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityDesc(v string) *AddMarketingFlowRequest
	GetActivityDesc() *string
	SetActivityName(v string) *AddMarketingFlowRequest
	GetActivityName() *string
	SetBizCode(v string) *AddMarketingFlowRequest
	GetBizCode() *string
	SetBizExtend(v map[string]interface{}) *AddMarketingFlowRequest
	GetBizExtend() map[string]interface{}
	SetCronExpression(v string) *AddMarketingFlowRequest
	GetCronExpression() *string
	SetEndDate(v string) *AddMarketingFlowRequest
	GetEndDate() *string
	SetExecutionType(v string) *AddMarketingFlowRequest
	GetExecutionType() *string
	SetOwnerId(v int64) *AddMarketingFlowRequest
	GetOwnerId() *int64
	SetParamFlag(v string) *AddMarketingFlowRequest
	GetParamFlag() *string
	SetParams(v map[string]interface{}) *AddMarketingFlowRequest
	GetParams() map[string]interface{}
	SetRelatedFlowCode(v string) *AddMarketingFlowRequest
	GetRelatedFlowCode() *string
	SetRelatedFlowName(v string) *AddMarketingFlowRequest
	GetRelatedFlowName() *string
	SetRelatedGroupId(v int64) *AddMarketingFlowRequest
	GetRelatedGroupId() *int64
	SetResourceOwnerAccount(v string) *AddMarketingFlowRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddMarketingFlowRequest
	GetResourceOwnerId() *int64
	SetStartDate(v string) *AddMarketingFlowRequest
	GetStartDate() *string
}

type AddMarketingFlowRequest struct {
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
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
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
	Params map[string]interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
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

func (s AddMarketingFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMarketingFlowRequest) GoString() string {
	return s.String()
}

func (s *AddMarketingFlowRequest) GetActivityDesc() *string {
	return s.ActivityDesc
}

func (s *AddMarketingFlowRequest) GetActivityName() *string {
	return s.ActivityName
}

func (s *AddMarketingFlowRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *AddMarketingFlowRequest) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *AddMarketingFlowRequest) GetCronExpression() *string {
	return s.CronExpression
}

func (s *AddMarketingFlowRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *AddMarketingFlowRequest) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *AddMarketingFlowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddMarketingFlowRequest) GetParamFlag() *string {
	return s.ParamFlag
}

func (s *AddMarketingFlowRequest) GetParams() map[string]interface{} {
	return s.Params
}

func (s *AddMarketingFlowRequest) GetRelatedFlowCode() *string {
	return s.RelatedFlowCode
}

func (s *AddMarketingFlowRequest) GetRelatedFlowName() *string {
	return s.RelatedFlowName
}

func (s *AddMarketingFlowRequest) GetRelatedGroupId() *int64 {
	return s.RelatedGroupId
}

func (s *AddMarketingFlowRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddMarketingFlowRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddMarketingFlowRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *AddMarketingFlowRequest) SetActivityDesc(v string) *AddMarketingFlowRequest {
	s.ActivityDesc = &v
	return s
}

func (s *AddMarketingFlowRequest) SetActivityName(v string) *AddMarketingFlowRequest {
	s.ActivityName = &v
	return s
}

func (s *AddMarketingFlowRequest) SetBizCode(v string) *AddMarketingFlowRequest {
	s.BizCode = &v
	return s
}

func (s *AddMarketingFlowRequest) SetBizExtend(v map[string]interface{}) *AddMarketingFlowRequest {
	s.BizExtend = v
	return s
}

func (s *AddMarketingFlowRequest) SetCronExpression(v string) *AddMarketingFlowRequest {
	s.CronExpression = &v
	return s
}

func (s *AddMarketingFlowRequest) SetEndDate(v string) *AddMarketingFlowRequest {
	s.EndDate = &v
	return s
}

func (s *AddMarketingFlowRequest) SetExecutionType(v string) *AddMarketingFlowRequest {
	s.ExecutionType = &v
	return s
}

func (s *AddMarketingFlowRequest) SetOwnerId(v int64) *AddMarketingFlowRequest {
	s.OwnerId = &v
	return s
}

func (s *AddMarketingFlowRequest) SetParamFlag(v string) *AddMarketingFlowRequest {
	s.ParamFlag = &v
	return s
}

func (s *AddMarketingFlowRequest) SetParams(v map[string]interface{}) *AddMarketingFlowRequest {
	s.Params = v
	return s
}

func (s *AddMarketingFlowRequest) SetRelatedFlowCode(v string) *AddMarketingFlowRequest {
	s.RelatedFlowCode = &v
	return s
}

func (s *AddMarketingFlowRequest) SetRelatedFlowName(v string) *AddMarketingFlowRequest {
	s.RelatedFlowName = &v
	return s
}

func (s *AddMarketingFlowRequest) SetRelatedGroupId(v int64) *AddMarketingFlowRequest {
	s.RelatedGroupId = &v
	return s
}

func (s *AddMarketingFlowRequest) SetResourceOwnerAccount(v string) *AddMarketingFlowRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddMarketingFlowRequest) SetResourceOwnerId(v int64) *AddMarketingFlowRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddMarketingFlowRequest) SetStartDate(v string) *AddMarketingFlowRequest {
	s.StartDate = &v
	return s
}

func (s *AddMarketingFlowRequest) Validate() error {
	return dara.Validate(s)
}
