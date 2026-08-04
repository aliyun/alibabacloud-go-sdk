// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetScheduleTaskResponseBody
	GetCode() *int32
	SetMessage(v string) *GetScheduleTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetScheduleTaskResponseBody
	GetRequestId() *string
	SetResult(v *GetScheduleTaskResponseBodyResult) *GetScheduleTaskResponseBody
	GetResult() *GetScheduleTaskResponseBodyResult
}

type GetScheduleTaskResponseBody struct {
	// Response code
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response message
	//
	// example:
	//
	// 调用成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// F7E21065-6C21-1158-A2F9-AEFE5CAB7C94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Service response parameters
	Result *GetScheduleTaskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetScheduleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetScheduleTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetScheduleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScheduleTaskResponseBody) GetResult() *GetScheduleTaskResponseBodyResult {
	return s.Result
}

func (s *GetScheduleTaskResponseBody) SetCode(v int32) *GetScheduleTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetScheduleTaskResponseBody) SetMessage(v string) *GetScheduleTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetScheduleTaskResponseBody) SetRequestId(v string) *GetScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduleTaskResponseBody) SetResult(v *GetScheduleTaskResponseBodyResult) *GetScheduleTaskResponseBody {
	s.Result = v
	return s
}

func (s *GetScheduleTaskResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScheduleTaskResponseBodyResult struct {
	// Trigger behavior
	ActionTopicList []*GetScheduleTaskResponseBodyResultActionTopicList `json:"ActionTopicList,omitempty" xml:"ActionTopicList,omitempty" type:"Repeated"`
	// Trigger Cron Expression
	//
	// example:
	//
	// 0 10 20 30 6 ? 2022
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// Validity Period - End Time
	//
	// example:
	//
	// 1659169473000
	ScheduleEndTime *string `json:"ScheduleEndTime,omitempty" xml:"ScheduleEndTime,omitempty"`
	// Job ID
	//
	// example:
	//
	// 1234567
	ScheduleId *int64 `json:"ScheduleId,omitempty" xml:"ScheduleId,omitempty"`
	// Validity Period - Start Time
	//
	// example:
	//
	// 1656577473000
	ScheduleStartTime *string `json:"ScheduleStartTime,omitempty" xml:"ScheduleStartTime,omitempty"`
	// Schedule Type
	//
	// example:
	//
	// ONCE
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
}

func (s GetScheduleTaskResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskResponseBodyResult) GetActionTopicList() []*GetScheduleTaskResponseBodyResultActionTopicList {
	return s.ActionTopicList
}

func (s *GetScheduleTaskResponseBodyResult) GetCron() *string {
	return s.Cron
}

func (s *GetScheduleTaskResponseBodyResult) GetScheduleEndTime() *string {
	return s.ScheduleEndTime
}

func (s *GetScheduleTaskResponseBodyResult) GetScheduleId() *int64 {
	return s.ScheduleId
}

func (s *GetScheduleTaskResponseBodyResult) GetScheduleStartTime() *string {
	return s.ScheduleStartTime
}

func (s *GetScheduleTaskResponseBodyResult) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *GetScheduleTaskResponseBodyResult) SetActionTopicList(v []*GetScheduleTaskResponseBodyResultActionTopicList) *GetScheduleTaskResponseBodyResult {
	s.ActionTopicList = v
	return s
}

func (s *GetScheduleTaskResponseBodyResult) SetCron(v string) *GetScheduleTaskResponseBodyResult {
	s.Cron = &v
	return s
}

func (s *GetScheduleTaskResponseBodyResult) SetScheduleEndTime(v string) *GetScheduleTaskResponseBodyResult {
	s.ScheduleEndTime = &v
	return s
}

func (s *GetScheduleTaskResponseBodyResult) SetScheduleId(v int64) *GetScheduleTaskResponseBodyResult {
	s.ScheduleId = &v
	return s
}

func (s *GetScheduleTaskResponseBodyResult) SetScheduleStartTime(v string) *GetScheduleTaskResponseBodyResult {
	s.ScheduleStartTime = &v
	return s
}

func (s *GetScheduleTaskResponseBodyResult) SetScheduleType(v string) *GetScheduleTaskResponseBodyResult {
	s.ScheduleType = &v
	return s
}

func (s *GetScheduleTaskResponseBodyResult) Validate() error {
	if s.ActionTopicList != nil {
		for _, item := range s.ActionTopicList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetScheduleTaskResponseBodyResultActionTopicList struct {
	// Vendor-defined command
	//
	// example:
	//
	// {"k1":"v1","k2":{"key":1}}
	CustomAction map[string]interface{} `json:"CustomAction,omitempty" xml:"CustomAction,omitempty"`
}

func (s GetScheduleTaskResponseBodyResultActionTopicList) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleTaskResponseBodyResultActionTopicList) GoString() string {
	return s.String()
}

func (s *GetScheduleTaskResponseBodyResultActionTopicList) GetCustomAction() map[string]interface{} {
	return s.CustomAction
}

func (s *GetScheduleTaskResponseBodyResultActionTopicList) SetCustomAction(v map[string]interface{}) *GetScheduleTaskResponseBodyResultActionTopicList {
	s.CustomAction = v
	return s
}

func (s *GetScheduleTaskResponseBodyResultActionTopicList) Validate() error {
	return dara.Validate(s)
}
