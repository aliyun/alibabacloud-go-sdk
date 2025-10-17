// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlarmEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListAlarmEventResponseBody
	GetCode() *int32
	SetData(v *ListAlarmEventResponseBodyData) *ListAlarmEventResponseBody
	GetData() *ListAlarmEventResponseBodyData
	SetMessage(v string) *ListAlarmEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAlarmEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAlarmEventResponseBody
	GetSuccess() *bool
}

type ListAlarmEventResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListAlarmEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 27B1345D-5F71-5972-8E4C-AABA6C6232F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAlarmEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlarmEventResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListAlarmEventResponseBody) GetData() *ListAlarmEventResponseBodyData {
	return s.Data
}

func (s *ListAlarmEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAlarmEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAlarmEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAlarmEventResponseBody) SetCode(v int32) *ListAlarmEventResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlarmEventResponseBody) SetData(v *ListAlarmEventResponseBodyData) *ListAlarmEventResponseBody {
	s.Data = v
	return s
}

func (s *ListAlarmEventResponseBody) SetMessage(v string) *ListAlarmEventResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlarmEventResponseBody) SetRequestId(v string) *ListAlarmEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlarmEventResponseBody) SetSuccess(v bool) *ListAlarmEventResponseBody {
	s.Success = &v
	return s
}

func (s *ListAlarmEventResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAlarmEventResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*ListAlarmEventResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 64
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAlarmEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAlarmEventResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAlarmEventResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlarmEventResponseBodyData) GetRecords() []*ListAlarmEventResponseBodyDataRecords {
	return s.Records
}

func (s *ListAlarmEventResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListAlarmEventResponseBodyData) SetPageNumber(v int32) *ListAlarmEventResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAlarmEventResponseBodyData) SetPageSize(v int32) *ListAlarmEventResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListAlarmEventResponseBodyData) SetRecords(v []*ListAlarmEventResponseBodyDataRecords) *ListAlarmEventResponseBodyData {
	s.Records = v
	return s
}

func (s *ListAlarmEventResponseBodyData) SetTotal(v int64) *ListAlarmEventResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListAlarmEventResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAlarmEventResponseBodyDataRecords struct {
	// example:
	//
	// webhook
	AlarmChannel *string `json:"AlarmChannel,omitempty" xml:"AlarmChannel,omitempty"`
	// example:
	//
	// zhangsan
	AlarmContacts *string `json:"AlarmContacts,omitempty" xml:"AlarmContacts,omitempty"`
	AlarmMessage  *string `json:"AlarmMessage,omitempty" xml:"AlarmMessage,omitempty"`
	// example:
	//
	// true
	AlarmStatus *string `json:"AlarmStatus,omitempty" xml:"AlarmStatus,omitempty"`
	// example:
	//
	// schedulerx3_fail_alarm
	AlarmType *string `json:"AlarmType,omitempty" xml:"AlarmType,omitempty"`
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 2024-10-31 16:43:51
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s ListAlarmEventResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListAlarmEventResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListAlarmEventResponseBodyDataRecords) GetAlarmChannel() *string {
	return s.AlarmChannel
}

func (s *ListAlarmEventResponseBodyDataRecords) GetAlarmContacts() *string {
	return s.AlarmContacts
}

func (s *ListAlarmEventResponseBodyDataRecords) GetAlarmMessage() *string {
	return s.AlarmMessage
}

func (s *ListAlarmEventResponseBodyDataRecords) GetAlarmStatus() *string {
	return s.AlarmStatus
}

func (s *ListAlarmEventResponseBodyDataRecords) GetAlarmType() *string {
	return s.AlarmType
}

func (s *ListAlarmEventResponseBodyDataRecords) GetAppName() *string {
	return s.AppName
}

func (s *ListAlarmEventResponseBodyDataRecords) GetJobName() *string {
	return s.JobName
}

func (s *ListAlarmEventResponseBodyDataRecords) GetTime() *string {
	return s.Time
}

func (s *ListAlarmEventResponseBodyDataRecords) SetAlarmChannel(v string) *ListAlarmEventResponseBodyDataRecords {
	s.AlarmChannel = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetAlarmContacts(v string) *ListAlarmEventResponseBodyDataRecords {
	s.AlarmContacts = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetAlarmMessage(v string) *ListAlarmEventResponseBodyDataRecords {
	s.AlarmMessage = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetAlarmStatus(v string) *ListAlarmEventResponseBodyDataRecords {
	s.AlarmStatus = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetAlarmType(v string) *ListAlarmEventResponseBodyDataRecords {
	s.AlarmType = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetAppName(v string) *ListAlarmEventResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetJobName(v string) *ListAlarmEventResponseBodyDataRecords {
	s.JobName = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) SetTime(v string) *ListAlarmEventResponseBodyDataRecords {
	s.Time = &v
	return s
}

func (s *ListAlarmEventResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
