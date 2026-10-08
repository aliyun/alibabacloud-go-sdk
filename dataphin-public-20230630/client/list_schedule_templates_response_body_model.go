// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListScheduleTemplatesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListScheduleTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetListScheduleTemplatesResponse(v *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse) *ListScheduleTemplatesResponseBody
	GetListScheduleTemplatesResponse() *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse
	SetMessage(v string) *ListScheduleTemplatesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListScheduleTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListScheduleTemplatesResponseBody
	GetSuccess() *bool
}

type ListScheduleTemplatesResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode                *int32                                                          `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	ListScheduleTemplatesResponse *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse `json:"ListScheduleTemplatesResponse,omitempty" xml:"ListScheduleTemplatesResponse,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListScheduleTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduleTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListScheduleTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListScheduleTemplatesResponseBody) GetListScheduleTemplatesResponse() *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse {
	return s.ListScheduleTemplatesResponse
}

func (s *ListScheduleTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScheduleTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScheduleTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScheduleTemplatesResponseBody) SetCode(v string) *ListScheduleTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListScheduleTemplatesResponseBody) SetHttpStatusCode(v int32) *ListScheduleTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListScheduleTemplatesResponseBody) SetListScheduleTemplatesResponse(v *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse) *ListScheduleTemplatesResponseBody {
	s.ListScheduleTemplatesResponse = v
	return s
}

func (s *ListScheduleTemplatesResponseBody) SetMessage(v string) *ListScheduleTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListScheduleTemplatesResponseBody) SetRequestId(v string) *ListScheduleTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduleTemplatesResponseBody) SetSuccess(v bool) *ListScheduleTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListScheduleTemplatesResponseBody) Validate() error {
	if s.ListScheduleTemplatesResponse != nil {
		if err := s.ListScheduleTemplatesResponse.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse struct {
	// example:
	//
	// 1
	Count      *int32                                                                      `json:"Count,omitempty" xml:"Count,omitempty"`
	ResultData []*ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData `json:"ResultData,omitempty" xml:"ResultData,omitempty" type:"Repeated"`
}

func (s ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse) GetCount() *int32 {
	return s.Count
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse) GetResultData() []*ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	return s.ResultData
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse) SetCount(v int32) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse {
	s.Count = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse) SetResultData(v []*ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse {
	s.ResultData = v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponse) Validate() error {
	if s.ResultData != nil {
		for _, item := range s.ResultData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData struct {
	ConditionScheduleParamList []*ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList `json:"ConditionScheduleParamList,omitempty" xml:"ConditionScheduleParamList,omitempty" type:"Repeated"`
	// example:
	//
	// 0 0 1 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// true
	CustomCronExpression *bool                                                                                         `json:"CustomCronExpression,omitempty" xml:"CustomCronExpression,omitempty"`
	CustomIntervalConfig *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig `json:"CustomIntervalConfig,omitempty" xml:"CustomIntervalConfig,omitempty" type:"Struct"`
	// example:
	//
	// CUSTOM_TIME_PERIOD
	CustomIntervalConfigType *string                                                                                          `json:"CustomIntervalConfigType,omitempty" xml:"CustomIntervalConfigType,omitempty"`
	CustomIntervalConfigs    []*ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs `json:"CustomIntervalConfigs,omitempty" xml:"CustomIntervalConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// 1704153600000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 1709516800000
	GmtModify *int64 `json:"GmtModify,omitempty" xml:"GmtModify,omitempty"`
	// example:
	//
	// true
	HasReference *bool `json:"HasReference,omitempty" xml:"HasReference,omitempty"`
	// example:
	//
	// 30001012
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// 李四
	ModifierName *string `json:"ModifierName,omitempty" xml:"ModifierName,omitempty"`
	// example:
	//
	// DAILY
	ScheduleIntervalType *string `json:"ScheduleIntervalType,omitempty" xml:"ScheduleIntervalType,omitempty"`
	// example:
	//
	// 工作日每天凌晨1点调度
	ScheduleTemplateDesc *string `json:"ScheduleTemplateDesc,omitempty" xml:"ScheduleTemplateDesc,omitempty"`
	// example:
	//
	// 12345
	ScheduleTemplateId *int64 `json:"ScheduleTemplateId,omitempty" xml:"ScheduleTemplateId,omitempty"`
	// example:
	//
	// 每天凌晨1点
	ScheduleTemplateName *string `json:"ScheduleTemplateName,omitempty" xml:"ScheduleTemplateName,omitempty"`
	// example:
	//
	// BASE_SCHEDULE_TEMPLATE
	ScheduleTemplateType *string `json:"ScheduleTemplateType,omitempty" xml:"ScheduleTemplateType,omitempty"`
	// example:
	//
	// 1
	ScheduleType *int32 `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// example:
	//
	// 30001011
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 30001011
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 张三
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// 9999-01-01
	ValidEndDate *string `json:"ValidEndDate,omitempty" xml:"ValidEndDate,omitempty"`
	// example:
	//
	// 2024-01-01
	ValidStartDate *string `json:"ValidStartDate,omitempty" xml:"ValidStartDate,omitempty"`
}

func (s ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GoString() string {
	return s.String()
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetConditionScheduleParamList() []*ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList {
	return s.ConditionScheduleParamList
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetCustomCronExpression() *bool {
	return s.CustomCronExpression
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetCustomIntervalConfig() *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig {
	return s.CustomIntervalConfig
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetCustomIntervalConfigType() *string {
	return s.CustomIntervalConfigType
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetCustomIntervalConfigs() []*ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs {
	return s.CustomIntervalConfigs
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetGmtModify() *int64 {
	return s.GmtModify
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetHasReference() *bool {
	return s.HasReference
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetModifierId() *string {
	return s.ModifierId
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetModifierName() *string {
	return s.ModifierName
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetScheduleIntervalType() *string {
	return s.ScheduleIntervalType
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetScheduleTemplateDesc() *string {
	return s.ScheduleTemplateDesc
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetScheduleTemplateId() *int64 {
	return s.ScheduleTemplateId
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetScheduleTemplateName() *string {
	return s.ScheduleTemplateName
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetScheduleTemplateType() *string {
	return s.ScheduleTemplateType
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetScheduleType() *int32 {
	return s.ScheduleType
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetUserId() *string {
	return s.UserId
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetUserName() *string {
	return s.UserName
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetValidEndDate() *string {
	return s.ValidEndDate
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) GetValidStartDate() *string {
	return s.ValidStartDate
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetConditionScheduleParamList(v []*ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.ConditionScheduleParamList = v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetCronExpression(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.CronExpression = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetCustomCronExpression(v bool) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.CustomCronExpression = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetCustomIntervalConfig(v *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.CustomIntervalConfig = v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetCustomIntervalConfigType(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.CustomIntervalConfigType = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetCustomIntervalConfigs(v []*ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.CustomIntervalConfigs = v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetGmtCreate(v int64) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.GmtCreate = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetGmtModify(v int64) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.GmtModify = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetHasReference(v bool) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.HasReference = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetModifierId(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.ModifierId = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetModifierName(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.ModifierName = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetScheduleIntervalType(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.ScheduleIntervalType = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetScheduleTemplateDesc(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.ScheduleTemplateDesc = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetScheduleTemplateId(v int64) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.ScheduleTemplateId = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetScheduleTemplateName(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.ScheduleTemplateName = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetScheduleTemplateType(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.ScheduleTemplateType = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetScheduleType(v int32) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.ScheduleType = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetTenantId(v int64) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.TenantId = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetUserId(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.UserId = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetUserName(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.UserName = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetValidEndDate(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.ValidEndDate = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) SetValidStartDate(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData {
	s.ValidStartDate = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultData) Validate() error {
	if s.ConditionScheduleParamList != nil {
		for _, item := range s.ConditionScheduleParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CustomIntervalConfig != nil {
		if err := s.CustomIntervalConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CustomIntervalConfigs != nil {
		for _, item := range s.CustomIntervalConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList struct {
	// example:
	//
	// 失败重跑条件
	ConditionName *string `json:"ConditionName,omitempty" xml:"ConditionName,omitempty"`
	// example:
	//
	// 0 30 	- 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// false
	FollowScheduleParam *bool `json:"FollowScheduleParam,omitempty" xml:"FollowScheduleParam,omitempty"`
	// example:
	//
	// 1
	NodeStatus *int32 `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
	// example:
	//
	// {"type":"EXPRESSION_GROUP","operator":"or"}
	ScheduleConditionJson *string `json:"ScheduleConditionJson,omitempty" xml:"ScheduleConditionJson,omitempty"`
	// example:
	//
	// 00:30
	ScheduleTime *string `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
}

func (s ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) GoString() string {
	return s.String()
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) GetConditionName() *string {
	return s.ConditionName
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) GetEnable() *bool {
	return s.Enable
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) GetFollowScheduleParam() *bool {
	return s.FollowScheduleParam
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) GetNodeStatus() *int32 {
	return s.NodeStatus
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) GetScheduleConditionJson() *string {
	return s.ScheduleConditionJson
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) GetScheduleTime() *string {
	return s.ScheduleTime
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) SetConditionName(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList {
	s.ConditionName = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) SetCronExpression(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList {
	s.CronExpression = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) SetEnable(v bool) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList {
	s.Enable = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) SetFollowScheduleParam(v bool) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList {
	s.FollowScheduleParam = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) SetNodeStatus(v int32) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList {
	s.NodeStatus = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) SetScheduleConditionJson(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList {
	s.ScheduleConditionJson = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) SetScheduleTime(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList {
	s.ScheduleTime = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataConditionScheduleParamList) Validate() error {
	return dara.Validate(s)
}

type ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig struct {
	// example:
	//
	// 23:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 30
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// MINUTE
	IntervalUnit *string `json:"IntervalUnit,omitempty" xml:"IntervalUnit,omitempty"`
	// example:
	//
	// DAY_INTERVAL
	SchedulePeriod *string `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
	// example:
	//
	// 00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) GoString() string {
	return s.String()
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) GetEndTime() *string {
	return s.EndTime
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) GetInterval() *int32 {
	return s.Interval
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) GetIntervalUnit() *string {
	return s.IntervalUnit
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) GetSchedulePeriod() *string {
	return s.SchedulePeriod
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) GetStartTime() *string {
	return s.StartTime
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) SetEndTime(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig {
	s.EndTime = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) SetInterval(v int32) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig {
	s.Interval = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) SetIntervalUnit(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig {
	s.IntervalUnit = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) SetSchedulePeriod(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig {
	s.SchedulePeriod = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) SetStartTime(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig {
	s.StartTime = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfig) Validate() error {
	return dara.Validate(s)
}

type ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs struct {
	// example:
	//
	// 23:59
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 30
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// MINUTE
	IntervalUnit *string `json:"IntervalUnit,omitempty" xml:"IntervalUnit,omitempty"`
	// example:
	//
	// DAY_INTERVAL
	SchedulePeriod *string `json:"SchedulePeriod,omitempty" xml:"SchedulePeriod,omitempty"`
	// example:
	//
	// 00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) GoString() string {
	return s.String()
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) GetEndTime() *string {
	return s.EndTime
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) GetInterval() *int32 {
	return s.Interval
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) GetIntervalUnit() *string {
	return s.IntervalUnit
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) GetSchedulePeriod() *string {
	return s.SchedulePeriod
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) GetStartTime() *string {
	return s.StartTime
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) SetEndTime(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs {
	s.EndTime = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) SetInterval(v int32) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs {
	s.Interval = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) SetIntervalUnit(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs {
	s.IntervalUnit = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) SetSchedulePeriod(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs {
	s.SchedulePeriod = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) SetStartTime(v string) *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs {
	s.StartTime = &v
	return s
}

func (s *ListScheduleTemplatesResponseBodyListScheduleTemplatesResponseResultDataCustomIntervalConfigs) Validate() error {
	return dara.Validate(s)
}
