// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudUpdateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudUpdateTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudUpdateTaskResponseBody
	GetCode() *string
	SetData(v *CloudUpdateTaskResponseBodyData) *CloudUpdateTaskResponseBody
	GetData() *CloudUpdateTaskResponseBodyData
	SetMessage(v string) *CloudUpdateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudUpdateTaskResponseBody
	GetRequestId() *string
}

type CloudUpdateTaskResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudUpdateTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudUpdateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudUpdateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CloudUpdateTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudUpdateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudUpdateTaskResponseBody) GetData() *CloudUpdateTaskResponseBodyData {
	return s.Data
}

func (s *CloudUpdateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudUpdateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudUpdateTaskResponseBody) SetAccessDeniedDetail(v string) *CloudUpdateTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudUpdateTaskResponseBody) SetCode(v string) *CloudUpdateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CloudUpdateTaskResponseBody) SetData(v *CloudUpdateTaskResponseBodyData) *CloudUpdateTaskResponseBody {
	s.Data = v
	return s
}

func (s *CloudUpdateTaskResponseBody) SetMessage(v string) *CloudUpdateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CloudUpdateTaskResponseBody) SetRequestId(v string) *CloudUpdateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudUpdateTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudUpdateTaskResponseBodyData struct {
	// example:
	//
	// 30
	AgentTimeout *string `json:"AgentTimeout,omitempty" xml:"AgentTimeout,omitempty"`
	// example:
	//
	// 55
	AnswerRate *string `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// example:
	//
	// 0
	AutoComplete *int64 `json:"AutoComplete,omitempty" xml:"AutoComplete,omitempty"`
	// example:
	//
	// 0
	AutoStart *string `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// example:
	//
	// 2026-02-11
	AutoStartDay *string `json:"AutoStartDay,omitempty" xml:"AutoStartDay,omitempty"`
	// example:
	//
	// 08:00:00
	AutoStartTime *string `json:"AutoStartTime,omitempty" xml:"AutoStartTime,omitempty"`
	// example:
	//
	// 0
	AutoStop *string `json:"AutoStop,omitempty" xml:"AutoStop,omitempty"`
	// example:
	//
	// 2026-02-11
	AutoStopDay *string `json:"AutoStopDay,omitempty" xml:"AutoStopDay,omitempty"`
	// example:
	//
	// 17:00:00
	AutoStopTime *string `json:"AutoStopTime,omitempty" xml:"AutoStopTime,omitempty"`
	// example:
	//
	// 1111
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
	// example:
	//
	// 10
	Concurrency *string `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// example:
	//
	// 2026-02-10 14:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 42366453,74565327
	CustomerClids *string `json:"CustomerClids,omitempty" xml:"CustomerClids,omitempty"`
	// example:
	//
	// 示例值示例值
	CustomerMoh *string `json:"CustomerMoh,omitempty" xml:"CustomerMoh,omitempty"`
	// example:
	//
	// 5
	CustomerTimeout *string `json:"CustomerTimeout,omitempty" xml:"CustomerTimeout,omitempty"`
	// example:
	//
	// 示例值示例值
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 8001654
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// example:
	//
	// 123
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	IsRandom *string `json:"IsRandom,omitempty" xml:"IsRandom,omitempty"`
	// example:
	//
	// 133
	IvrId *string `json:"IvrId,omitempty" xml:"IvrId,omitempty"`
	// example:
	//
	// 10
	MaxWaitTime *string `json:"MaxWaitTime,omitempty" xml:"MaxWaitTime,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2026-01-0114:00:00
	OverTime *string `json:"OverTime,omitempty" xml:"OverTime,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	PredictAdjust *string `json:"PredictAdjust,omitempty" xml:"PredictAdjust,omitempty"`
	// example:
	//
	// 0.33
	Quotiety *string `json:"Quotiety,omitempty" xml:"Quotiety,omitempty"`
	// example:
	//
	// 2026-02-10 15:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 3
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// [{"key1":"value1"},{"key2":"value2"}]
	UserFields *string `json:"UserFields,omitempty" xml:"UserFields,omitempty"`
	// example:
	//
	// 5
	Wrapup *string `json:"Wrapup,omitempty" xml:"Wrapup,omitempty"`
}

func (s CloudUpdateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudUpdateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudUpdateTaskResponseBodyData) GetAgentTimeout() *string {
	return s.AgentTimeout
}

func (s *CloudUpdateTaskResponseBodyData) GetAnswerRate() *string {
	return s.AnswerRate
}

func (s *CloudUpdateTaskResponseBodyData) GetAutoComplete() *int64 {
	return s.AutoComplete
}

func (s *CloudUpdateTaskResponseBodyData) GetAutoStart() *string {
	return s.AutoStart
}

func (s *CloudUpdateTaskResponseBodyData) GetAutoStartDay() *string {
	return s.AutoStartDay
}

func (s *CloudUpdateTaskResponseBodyData) GetAutoStartTime() *string {
	return s.AutoStartTime
}

func (s *CloudUpdateTaskResponseBodyData) GetAutoStop() *string {
	return s.AutoStop
}

func (s *CloudUpdateTaskResponseBodyData) GetAutoStopDay() *string {
	return s.AutoStopDay
}

func (s *CloudUpdateTaskResponseBodyData) GetAutoStopTime() *string {
	return s.AutoStopTime
}

func (s *CloudUpdateTaskResponseBodyData) GetCnos() *string {
	return s.Cnos
}

func (s *CloudUpdateTaskResponseBodyData) GetConcurrency() *string {
	return s.Concurrency
}

func (s *CloudUpdateTaskResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudUpdateTaskResponseBodyData) GetCustomerClids() *string {
	return s.CustomerClids
}

func (s *CloudUpdateTaskResponseBodyData) GetCustomerMoh() *string {
	return s.CustomerMoh
}

func (s *CloudUpdateTaskResponseBodyData) GetCustomerTimeout() *string {
	return s.CustomerTimeout
}

func (s *CloudUpdateTaskResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CloudUpdateTaskResponseBodyData) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudUpdateTaskResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CloudUpdateTaskResponseBodyData) GetIsRandom() *string {
	return s.IsRandom
}

func (s *CloudUpdateTaskResponseBodyData) GetIvrId() *string {
	return s.IvrId
}

func (s *CloudUpdateTaskResponseBodyData) GetMaxWaitTime() *string {
	return s.MaxWaitTime
}

func (s *CloudUpdateTaskResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CloudUpdateTaskResponseBodyData) GetOverTime() *string {
	return s.OverTime
}

func (s *CloudUpdateTaskResponseBodyData) GetPredictAdjust() *string {
	return s.PredictAdjust
}

func (s *CloudUpdateTaskResponseBodyData) GetQuotiety() *string {
	return s.Quotiety
}

func (s *CloudUpdateTaskResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudUpdateTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CloudUpdateTaskResponseBodyData) GetType() *string {
	return s.Type
}

func (s *CloudUpdateTaskResponseBodyData) GetUserFields() *string {
	return s.UserFields
}

func (s *CloudUpdateTaskResponseBodyData) GetWrapup() *string {
	return s.Wrapup
}

func (s *CloudUpdateTaskResponseBodyData) SetAgentTimeout(v string) *CloudUpdateTaskResponseBodyData {
	s.AgentTimeout = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetAnswerRate(v string) *CloudUpdateTaskResponseBodyData {
	s.AnswerRate = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetAutoComplete(v int64) *CloudUpdateTaskResponseBodyData {
	s.AutoComplete = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetAutoStart(v string) *CloudUpdateTaskResponseBodyData {
	s.AutoStart = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetAutoStartDay(v string) *CloudUpdateTaskResponseBodyData {
	s.AutoStartDay = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetAutoStartTime(v string) *CloudUpdateTaskResponseBodyData {
	s.AutoStartTime = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetAutoStop(v string) *CloudUpdateTaskResponseBodyData {
	s.AutoStop = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetAutoStopDay(v string) *CloudUpdateTaskResponseBodyData {
	s.AutoStopDay = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetAutoStopTime(v string) *CloudUpdateTaskResponseBodyData {
	s.AutoStopTime = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetCnos(v string) *CloudUpdateTaskResponseBodyData {
	s.Cnos = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetConcurrency(v string) *CloudUpdateTaskResponseBodyData {
	s.Concurrency = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetCreateTime(v string) *CloudUpdateTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetCustomerClids(v string) *CloudUpdateTaskResponseBodyData {
	s.CustomerClids = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetCustomerMoh(v string) *CloudUpdateTaskResponseBodyData {
	s.CustomerMoh = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetCustomerTimeout(v string) *CloudUpdateTaskResponseBodyData {
	s.CustomerTimeout = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetDescription(v string) *CloudUpdateTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetEnterpriseId(v string) *CloudUpdateTaskResponseBodyData {
	s.EnterpriseId = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetId(v string) *CloudUpdateTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetIsRandom(v string) *CloudUpdateTaskResponseBodyData {
	s.IsRandom = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetIvrId(v string) *CloudUpdateTaskResponseBodyData {
	s.IvrId = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetMaxWaitTime(v string) *CloudUpdateTaskResponseBodyData {
	s.MaxWaitTime = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetName(v string) *CloudUpdateTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetOverTime(v string) *CloudUpdateTaskResponseBodyData {
	s.OverTime = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetPredictAdjust(v string) *CloudUpdateTaskResponseBodyData {
	s.PredictAdjust = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetQuotiety(v string) *CloudUpdateTaskResponseBodyData {
	s.Quotiety = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetStartTime(v string) *CloudUpdateTaskResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetStatus(v string) *CloudUpdateTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetType(v string) *CloudUpdateTaskResponseBodyData {
	s.Type = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetUserFields(v string) *CloudUpdateTaskResponseBodyData {
	s.UserFields = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) SetWrapup(v string) *CloudUpdateTaskResponseBodyData {
	s.Wrapup = &v
	return s
}

func (s *CloudUpdateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
