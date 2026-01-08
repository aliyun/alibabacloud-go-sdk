// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMarketingFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListMarketingFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListMarketingFlowResponseBody
	GetCode() *string
	SetData(v []*ListMarketingFlowResponseBodyData) *ListMarketingFlowResponseBody
	GetData() []*ListMarketingFlowResponseBodyData
	SetMessage(v string) *ListMarketingFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMarketingFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMarketingFlowResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListMarketingFlowResponseBody
	GetTotalCount() *int64
}

type ListMarketingFlowResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListMarketingFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// NULL
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// CB6122C9-09B5-5926-**476A96CB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMarketingFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMarketingFlowResponseBody) GoString() string {
	return s.String()
}

func (s *ListMarketingFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListMarketingFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMarketingFlowResponseBody) GetData() []*ListMarketingFlowResponseBodyData {
	return s.Data
}

func (s *ListMarketingFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMarketingFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMarketingFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMarketingFlowResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMarketingFlowResponseBody) SetAccessDeniedDetail(v string) *ListMarketingFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListMarketingFlowResponseBody) SetCode(v string) *ListMarketingFlowResponseBody {
	s.Code = &v
	return s
}

func (s *ListMarketingFlowResponseBody) SetData(v []*ListMarketingFlowResponseBodyData) *ListMarketingFlowResponseBody {
	s.Data = v
	return s
}

func (s *ListMarketingFlowResponseBody) SetMessage(v string) *ListMarketingFlowResponseBody {
	s.Message = &v
	return s
}

func (s *ListMarketingFlowResponseBody) SetRequestId(v string) *ListMarketingFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMarketingFlowResponseBody) SetSuccess(v bool) *ListMarketingFlowResponseBody {
	s.Success = &v
	return s
}

func (s *ListMarketingFlowResponseBody) SetTotalCount(v int64) *ListMarketingFlowResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMarketingFlowResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMarketingFlowResponseBodyData struct {
	// example:
	//
	// 3243243***
	ActivityCode *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
	// example:
	//
	// aaa
	ActivityDesc *string `json:"ActivityDesc,omitempty" xml:"ActivityDesc,omitempty"`
	// example:
	//
	// aaa
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// example:
	//
	// active
	ActivityStatus *string `json:"ActivityStatus,omitempty" xml:"ActivityStatus,omitempty"`
	// example:
	//
	// {}
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtend map[string]interface{} `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// example:
	//
	// 0 0 4 1/1 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// 2025-01-01 XX1:11:11
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 示例值
	ExecutionType *string `json:"ExecutionType,omitempty" xml:"ExecutionType,omitempty"`
	// example:
	//
	// 2025-XX-01 11:11:11
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// N/A
	GmtModifier *string `json:"GmtModifier,omitempty" xml:"GmtModifier,omitempty"`
	// example:
	//
	// 99
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Y
	ParamFlag *string `json:"ParamFlag,omitempty" xml:"ParamFlag,omitempty"`
	// example:
	//
	// {\\"CustomerLimit\\":false}
	Params map[string]interface{} `json:"Params,omitempty" xml:"Params,omitempty"`
	// example:
	//
	// dsafdsf***
	RelatedFlowCode *string `json:"RelatedFlowCode,omitempty" xml:"RelatedFlowCode,omitempty"`
	// example:
	//
	// aaa
	RelatedFlowName *string `json:"RelatedFlowName,omitempty" xml:"RelatedFlowName,omitempty"`
	// example:
	//
	// 54354**
	RelatedGroupId *string `json:"RelatedGroupId,omitempty" xml:"RelatedGroupId,omitempty"`
	// example:
	//
	// AAA
	RelatedGroupName *string `json:"RelatedGroupName,omitempty" xml:"RelatedGroupName,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	SpecificTime *string `json:"SpecificTime,omitempty" xml:"SpecificTime,omitempty"`
	// example:
	//
	// 2025-01-XX 11:11:11
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// 1111
	TenantCode *string `json:"TenantCode,omitempty" xml:"TenantCode,omitempty"`
}

func (s ListMarketingFlowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMarketingFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMarketingFlowResponseBodyData) GetActivityCode() *string {
	return s.ActivityCode
}

func (s *ListMarketingFlowResponseBodyData) GetActivityDesc() *string {
	return s.ActivityDesc
}

func (s *ListMarketingFlowResponseBodyData) GetActivityName() *string {
	return s.ActivityName
}

func (s *ListMarketingFlowResponseBodyData) GetActivityStatus() *string {
	return s.ActivityStatus
}

func (s *ListMarketingFlowResponseBodyData) GetBizCode() *string {
	return s.BizCode
}

func (s *ListMarketingFlowResponseBodyData) GetBizExtend() map[string]interface{} {
	return s.BizExtend
}

func (s *ListMarketingFlowResponseBodyData) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ListMarketingFlowResponseBodyData) GetEndDate() *string {
	return s.EndDate
}

func (s *ListMarketingFlowResponseBodyData) GetExecutionType() *string {
	return s.ExecutionType
}

func (s *ListMarketingFlowResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMarketingFlowResponseBodyData) GetGmtModifier() *string {
	return s.GmtModifier
}

func (s *ListMarketingFlowResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListMarketingFlowResponseBodyData) GetParamFlag() *string {
	return s.ParamFlag
}

func (s *ListMarketingFlowResponseBodyData) GetParams() map[string]interface{} {
	return s.Params
}

func (s *ListMarketingFlowResponseBodyData) GetRelatedFlowCode() *string {
	return s.RelatedFlowCode
}

func (s *ListMarketingFlowResponseBodyData) GetRelatedFlowName() *string {
	return s.RelatedFlowName
}

func (s *ListMarketingFlowResponseBodyData) GetRelatedGroupId() *string {
	return s.RelatedGroupId
}

func (s *ListMarketingFlowResponseBodyData) GetRelatedGroupName() *string {
	return s.RelatedGroupName
}

func (s *ListMarketingFlowResponseBodyData) GetSpecificTime() *string {
	return s.SpecificTime
}

func (s *ListMarketingFlowResponseBodyData) GetStartDate() *string {
	return s.StartDate
}

func (s *ListMarketingFlowResponseBodyData) GetTenantCode() *string {
	return s.TenantCode
}

func (s *ListMarketingFlowResponseBodyData) SetActivityCode(v string) *ListMarketingFlowResponseBodyData {
	s.ActivityCode = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetActivityDesc(v string) *ListMarketingFlowResponseBodyData {
	s.ActivityDesc = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetActivityName(v string) *ListMarketingFlowResponseBodyData {
	s.ActivityName = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetActivityStatus(v string) *ListMarketingFlowResponseBodyData {
	s.ActivityStatus = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetBizCode(v string) *ListMarketingFlowResponseBodyData {
	s.BizCode = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetBizExtend(v map[string]interface{}) *ListMarketingFlowResponseBodyData {
	s.BizExtend = v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetCronExpression(v string) *ListMarketingFlowResponseBodyData {
	s.CronExpression = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetEndDate(v string) *ListMarketingFlowResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetExecutionType(v string) *ListMarketingFlowResponseBodyData {
	s.ExecutionType = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetGmtCreate(v string) *ListMarketingFlowResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetGmtModifier(v string) *ListMarketingFlowResponseBodyData {
	s.GmtModifier = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetId(v int64) *ListMarketingFlowResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetParamFlag(v string) *ListMarketingFlowResponseBodyData {
	s.ParamFlag = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetParams(v map[string]interface{}) *ListMarketingFlowResponseBodyData {
	s.Params = v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetRelatedFlowCode(v string) *ListMarketingFlowResponseBodyData {
	s.RelatedFlowCode = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetRelatedFlowName(v string) *ListMarketingFlowResponseBodyData {
	s.RelatedFlowName = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetRelatedGroupId(v string) *ListMarketingFlowResponseBodyData {
	s.RelatedGroupId = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetRelatedGroupName(v string) *ListMarketingFlowResponseBodyData {
	s.RelatedGroupName = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetSpecificTime(v string) *ListMarketingFlowResponseBodyData {
	s.SpecificTime = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetStartDate(v string) *ListMarketingFlowResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) SetTenantCode(v string) *ListMarketingFlowResponseBodyData {
	s.TenantCode = &v
	return s
}

func (s *ListMarketingFlowResponseBodyData) Validate() error {
	return dara.Validate(s)
}
