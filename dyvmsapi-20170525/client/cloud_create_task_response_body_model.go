// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudCreateTaskResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudCreateTaskResponseBody
	GetCode() *string
	SetData(v *CloudCreateTaskResponseBodyData) *CloudCreateTaskResponseBody
	GetData() *CloudCreateTaskResponseBodyData
	SetMessage(v string) *CloudCreateTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudCreateTaskResponseBody
	GetRequestId() *string
}

type CloudCreateTaskResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudCreateTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudCreateTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CloudCreateTaskResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudCreateTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudCreateTaskResponseBody) GetData() *CloudCreateTaskResponseBodyData {
	return s.Data
}

func (s *CloudCreateTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudCreateTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudCreateTaskResponseBody) SetAccessDeniedDetail(v string) *CloudCreateTaskResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudCreateTaskResponseBody) SetCode(v string) *CloudCreateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CloudCreateTaskResponseBody) SetData(v *CloudCreateTaskResponseBodyData) *CloudCreateTaskResponseBody {
	s.Data = v
	return s
}

func (s *CloudCreateTaskResponseBody) SetMessage(v string) *CloudCreateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CloudCreateTaskResponseBody) SetRequestId(v string) *CloudCreateTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudCreateTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudCreateTaskResponseBodyData struct {
	// example:
	//
	// 示例值示例值示例值
	AgentTimeout *string `json:"AgentTimeout,omitempty" xml:"AgentTimeout,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	AnswerRate *string `json:"AnswerRate,omitempty" xml:"AnswerRate,omitempty"`
	// example:
	//
	// 1
	AutoComplete *int64 `json:"AutoComplete,omitempty" xml:"AutoComplete,omitempty"`
	// example:
	//
	// 1
	AutoStart *string `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	// example:
	//
	// 2016-04-11
	AutoStartDay *string `json:"AutoStartDay,omitempty" xml:"AutoStartDay,omitempty"`
	// example:
	//
	// 08:00:00
	AutoStartTime *string `json:"AutoStartTime,omitempty" xml:"AutoStartTime,omitempty"`
	// example:
	//
	// 1
	AutoStop *string `json:"AutoStop,omitempty" xml:"AutoStop,omitempty"`
	// example:
	//
	// 示例值示例值
	AutoStopDay *string `json:"AutoStopDay,omitempty" xml:"AutoStopDay,omitempty"`
	// example:
	//
	// 17:00:00
	AutoStopTime *string `json:"AutoStopTime,omitempty" xml:"AutoStopTime,omitempty"`
	// example:
	//
	// 示例值示例值
	Cnos *string `json:"Cnos,omitempty" xml:"Cnos,omitempty"`
	// example:
	//
	// 示例值示例值
	Concurrency *string `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	CustomerClids *string `json:"CustomerClids,omitempty" xml:"CustomerClids,omitempty"`
	// example:
	//
	// 示例值示例值
	CustomerMoh *string `json:"CustomerMoh,omitempty" xml:"CustomerMoh,omitempty"`
	// example:
	//
	// 示例值示例值
	CustomerTimeout *string `json:"CustomerTimeout,omitempty" xml:"CustomerTimeout,omitempty"`
	// example:
	//
	// 示例值示例值
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 示例值示例值
	EnterpriseId *string `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// example:
	//
	// 11
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 0
	IsRandom *string `json:"IsRandom,omitempty" xml:"IsRandom,omitempty"`
	// example:
	//
	// 330
	IvrId *string `json:"IvrId,omitempty" xml:"IvrId,omitempty"`
	// example:
	//
	// 10
	MaxWaitTime *string `json:"MaxWaitTime,omitempty" xml:"MaxWaitTime,omitempty"`
	// example:
	//
	// 示例值示例值
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 示例值示例值
	OverTime *string `json:"OverTime,omitempty" xml:"OverTime,omitempty"`
	// example:
	//
	// 示例值示例值
	PredictAdjust *string `json:"PredictAdjust,omitempty" xml:"PredictAdjust,omitempty"`
	// example:
	//
	// 0.33
	Quotiety *string `json:"Quotiety,omitempty" xml:"Quotiety,omitempty"`
	// example:
	//
	// 示例值
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 3
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// [{"key1":"value1"},{"key2":"value2"}]
	UserFields *string `json:"UserFields,omitempty" xml:"UserFields,omitempty"`
	// example:
	//
	// 10
	Wrapup *string `json:"Wrapup,omitempty" xml:"Wrapup,omitempty"`
}

func (s CloudCreateTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudCreateTaskResponseBodyData) GetAgentTimeout() *string {
	return s.AgentTimeout
}

func (s *CloudCreateTaskResponseBodyData) GetAnswerRate() *string {
	return s.AnswerRate
}

func (s *CloudCreateTaskResponseBodyData) GetAutoComplete() *int64 {
	return s.AutoComplete
}

func (s *CloudCreateTaskResponseBodyData) GetAutoStart() *string {
	return s.AutoStart
}

func (s *CloudCreateTaskResponseBodyData) GetAutoStartDay() *string {
	return s.AutoStartDay
}

func (s *CloudCreateTaskResponseBodyData) GetAutoStartTime() *string {
	return s.AutoStartTime
}

func (s *CloudCreateTaskResponseBodyData) GetAutoStop() *string {
	return s.AutoStop
}

func (s *CloudCreateTaskResponseBodyData) GetAutoStopDay() *string {
	return s.AutoStopDay
}

func (s *CloudCreateTaskResponseBodyData) GetAutoStopTime() *string {
	return s.AutoStopTime
}

func (s *CloudCreateTaskResponseBodyData) GetCnos() *string {
	return s.Cnos
}

func (s *CloudCreateTaskResponseBodyData) GetConcurrency() *string {
	return s.Concurrency
}

func (s *CloudCreateTaskResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CloudCreateTaskResponseBodyData) GetCustomerClids() *string {
	return s.CustomerClids
}

func (s *CloudCreateTaskResponseBodyData) GetCustomerMoh() *string {
	return s.CustomerMoh
}

func (s *CloudCreateTaskResponseBodyData) GetCustomerTimeout() *string {
	return s.CustomerTimeout
}

func (s *CloudCreateTaskResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CloudCreateTaskResponseBodyData) GetEnterpriseId() *string {
	return s.EnterpriseId
}

func (s *CloudCreateTaskResponseBodyData) GetId() *string {
	return s.Id
}

func (s *CloudCreateTaskResponseBodyData) GetIsRandom() *string {
	return s.IsRandom
}

func (s *CloudCreateTaskResponseBodyData) GetIvrId() *string {
	return s.IvrId
}

func (s *CloudCreateTaskResponseBodyData) GetMaxWaitTime() *string {
	return s.MaxWaitTime
}

func (s *CloudCreateTaskResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CloudCreateTaskResponseBodyData) GetOverTime() *string {
	return s.OverTime
}

func (s *CloudCreateTaskResponseBodyData) GetPredictAdjust() *string {
	return s.PredictAdjust
}

func (s *CloudCreateTaskResponseBodyData) GetQuotiety() *string {
	return s.Quotiety
}

func (s *CloudCreateTaskResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *CloudCreateTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CloudCreateTaskResponseBodyData) GetType() *string {
	return s.Type
}

func (s *CloudCreateTaskResponseBodyData) GetUserFields() *string {
	return s.UserFields
}

func (s *CloudCreateTaskResponseBodyData) GetWrapup() *string {
	return s.Wrapup
}

func (s *CloudCreateTaskResponseBodyData) SetAgentTimeout(v string) *CloudCreateTaskResponseBodyData {
	s.AgentTimeout = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetAnswerRate(v string) *CloudCreateTaskResponseBodyData {
	s.AnswerRate = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetAutoComplete(v int64) *CloudCreateTaskResponseBodyData {
	s.AutoComplete = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetAutoStart(v string) *CloudCreateTaskResponseBodyData {
	s.AutoStart = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetAutoStartDay(v string) *CloudCreateTaskResponseBodyData {
	s.AutoStartDay = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetAutoStartTime(v string) *CloudCreateTaskResponseBodyData {
	s.AutoStartTime = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetAutoStop(v string) *CloudCreateTaskResponseBodyData {
	s.AutoStop = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetAutoStopDay(v string) *CloudCreateTaskResponseBodyData {
	s.AutoStopDay = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetAutoStopTime(v string) *CloudCreateTaskResponseBodyData {
	s.AutoStopTime = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetCnos(v string) *CloudCreateTaskResponseBodyData {
	s.Cnos = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetConcurrency(v string) *CloudCreateTaskResponseBodyData {
	s.Concurrency = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetCreateTime(v string) *CloudCreateTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetCustomerClids(v string) *CloudCreateTaskResponseBodyData {
	s.CustomerClids = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetCustomerMoh(v string) *CloudCreateTaskResponseBodyData {
	s.CustomerMoh = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetCustomerTimeout(v string) *CloudCreateTaskResponseBodyData {
	s.CustomerTimeout = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetDescription(v string) *CloudCreateTaskResponseBodyData {
	s.Description = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetEnterpriseId(v string) *CloudCreateTaskResponseBodyData {
	s.EnterpriseId = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetId(v string) *CloudCreateTaskResponseBodyData {
	s.Id = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetIsRandom(v string) *CloudCreateTaskResponseBodyData {
	s.IsRandom = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetIvrId(v string) *CloudCreateTaskResponseBodyData {
	s.IvrId = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetMaxWaitTime(v string) *CloudCreateTaskResponseBodyData {
	s.MaxWaitTime = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetName(v string) *CloudCreateTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetOverTime(v string) *CloudCreateTaskResponseBodyData {
	s.OverTime = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetPredictAdjust(v string) *CloudCreateTaskResponseBodyData {
	s.PredictAdjust = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetQuotiety(v string) *CloudCreateTaskResponseBodyData {
	s.Quotiety = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetStartTime(v string) *CloudCreateTaskResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetStatus(v string) *CloudCreateTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetType(v string) *CloudCreateTaskResponseBodyData {
	s.Type = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetUserFields(v string) *CloudCreateTaskResponseBodyData {
	s.UserFields = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) SetWrapup(v string) *CloudCreateTaskResponseBodyData {
	s.Wrapup = &v
	return s
}

func (s *CloudCreateTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
