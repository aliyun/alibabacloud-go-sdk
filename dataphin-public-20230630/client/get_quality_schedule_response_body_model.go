// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetQualityScheduleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetQualityScheduleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetQualityScheduleResponseBody
	GetMessage() *string
	SetQualityScheduleInfo(v *GetQualityScheduleResponseBodyQualityScheduleInfo) *GetQualityScheduleResponseBody
	GetQualityScheduleInfo() *GetQualityScheduleResponseBodyQualityScheduleInfo
	SetRequestId(v string) *GetQualityScheduleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetQualityScheduleResponseBody
	GetSuccess() *bool
}

type GetQualityScheduleResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message             *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	QualityScheduleInfo *GetQualityScheduleResponseBodyQualityScheduleInfo `json:"QualityScheduleInfo,omitempty" xml:"QualityScheduleInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQualityScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQualityScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *GetQualityScheduleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetQualityScheduleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetQualityScheduleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQualityScheduleResponseBody) GetQualityScheduleInfo() *GetQualityScheduleResponseBodyQualityScheduleInfo {
	return s.QualityScheduleInfo
}

func (s *GetQualityScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQualityScheduleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQualityScheduleResponseBody) SetCode(v string) *GetQualityScheduleResponseBody {
	s.Code = &v
	return s
}

func (s *GetQualityScheduleResponseBody) SetHttpStatusCode(v int32) *GetQualityScheduleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetQualityScheduleResponseBody) SetMessage(v string) *GetQualityScheduleResponseBody {
	s.Message = &v
	return s
}

func (s *GetQualityScheduleResponseBody) SetQualityScheduleInfo(v *GetQualityScheduleResponseBodyQualityScheduleInfo) *GetQualityScheduleResponseBody {
	s.QualityScheduleInfo = v
	return s
}

func (s *GetQualityScheduleResponseBody) SetRequestId(v string) *GetQualityScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQualityScheduleResponseBody) SetSuccess(v bool) *GetQualityScheduleResponseBody {
	s.Success = &v
	return s
}

func (s *GetQualityScheduleResponseBody) Validate() error {
	if s.QualityScheduleInfo != nil {
		if err := s.QualityScheduleInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQualityScheduleResponseBodyQualityScheduleInfo struct {
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 30012011
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 	- 	- 1/	- 	- 	- *
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// example:
	//
	// 1
	Id          *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	IsRefByRule *bool  `json:"IsRefByRule,omitempty" xml:"IsRefByRule,omitempty"`
	// example:
	//
	// 30012011
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ds=${yyyyMMdd}
	PartitionExpression *string `json:"PartitionExpression,omitempty" xml:"PartitionExpression,omitempty"`
	// example:
	//
	// CUSTOM
	PartitionType *string `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	// example:
	//
	// DAILY
	PeriodScheduleIntervalType *string   `json:"PeriodScheduleIntervalType,omitempty" xml:"PeriodScheduleIntervalType,omitempty"`
	PeriodScheduleParamList    []*string `json:"PeriodScheduleParamList,omitempty" xml:"PeriodScheduleParamList,omitempty" type:"Repeated"`
	// example:
	//
	// ONE_TASKS_FINISHED
	StaticTaskTriggerType *string   `json:"StaticTaskTriggerType,omitempty" xml:"StaticTaskTriggerType,omitempty"`
	TriggerNodeList       []*string `json:"TriggerNodeList,omitempty" xml:"TriggerNodeList,omitempty" type:"Repeated"`
	// example:
	//
	// STATIC_TASK_TRIGGER
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// example:
	//
	// PERIOD_SCHEDULE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// TASK_REFERRED_PARTITION
	ValidatePartitionType *string `json:"ValidatePartitionType,omitempty" xml:"ValidatePartitionType,omitempty"`
	// example:
	//
	// 22
	WatchId *int64 `json:"WatchId,omitempty" xml:"WatchId,omitempty"`
}

func (s GetQualityScheduleResponseBodyQualityScheduleInfo) String() string {
	return dara.Prettify(s)
}

func (s GetQualityScheduleResponseBodyQualityScheduleInfo) GoString() string {
	return s.String()
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetCronExpression() *string {
	return s.CronExpression
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetId() *int64 {
	return s.Id
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetIsRefByRule() *bool {
	return s.IsRefByRule
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetModifier() *string {
	return s.Modifier
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetName() *string {
	return s.Name
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetPartitionExpression() *string {
	return s.PartitionExpression
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetPartitionType() *string {
	return s.PartitionType
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetPeriodScheduleIntervalType() *string {
	return s.PeriodScheduleIntervalType
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetPeriodScheduleParamList() []*string {
	return s.PeriodScheduleParamList
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetStaticTaskTriggerType() *string {
	return s.StaticTaskTriggerType
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetTriggerNodeList() []*string {
	return s.TriggerNodeList
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetType() *string {
	return s.Type
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetValidatePartitionType() *string {
	return s.ValidatePartitionType
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) GetWatchId() *int64 {
	return s.WatchId
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetCreateTime(v string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.CreateTime = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetCreator(v string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.Creator = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetCronExpression(v string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.CronExpression = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetId(v int64) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.Id = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetIsRefByRule(v bool) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.IsRefByRule = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetModifier(v string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.Modifier = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetModifyTime(v string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.ModifyTime = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetName(v string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.Name = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetPartitionExpression(v string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.PartitionExpression = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetPartitionType(v string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.PartitionType = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetPeriodScheduleIntervalType(v string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.PeriodScheduleIntervalType = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetPeriodScheduleParamList(v []*string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.PeriodScheduleParamList = v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetStaticTaskTriggerType(v string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.StaticTaskTriggerType = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetTriggerNodeList(v []*string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.TriggerNodeList = v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetTriggerType(v string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.TriggerType = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetType(v string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.Type = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetValidatePartitionType(v string) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.ValidatePartitionType = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) SetWatchId(v int64) *GetQualityScheduleResponseBodyQualityScheduleInfo {
	s.WatchId = &v
	return s
}

func (s *GetQualityScheduleResponseBodyQualityScheduleInfo) Validate() error {
	return dara.Validate(s)
}
