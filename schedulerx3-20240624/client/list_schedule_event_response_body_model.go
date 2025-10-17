// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListScheduleEventResponseBody
	GetCode() *int32
	SetData(v *ListScheduleEventResponseBodyData) *ListScheduleEventResponseBody
	GetData() *ListScheduleEventResponseBodyData
	SetMessage(v string) *ListScheduleEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListScheduleEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListScheduleEventResponseBody
	GetSuccess() *bool
}

type ListScheduleEventResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListScheduleEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// B880122A-B0E4-52E8-8F54-87DB7779EB74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListScheduleEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduleEventResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListScheduleEventResponseBody) GetData() *ListScheduleEventResponseBodyData {
	return s.Data
}

func (s *ListScheduleEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListScheduleEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScheduleEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListScheduleEventResponseBody) SetCode(v int32) *ListScheduleEventResponseBody {
	s.Code = &v
	return s
}

func (s *ListScheduleEventResponseBody) SetData(v *ListScheduleEventResponseBodyData) *ListScheduleEventResponseBody {
	s.Data = v
	return s
}

func (s *ListScheduleEventResponseBody) SetMessage(v string) *ListScheduleEventResponseBody {
	s.Message = &v
	return s
}

func (s *ListScheduleEventResponseBody) SetRequestId(v string) *ListScheduleEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduleEventResponseBody) SetSuccess(v bool) *ListScheduleEventResponseBody {
	s.Success = &v
	return s
}

func (s *ListScheduleEventResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListScheduleEventResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Records  []*ListScheduleEventResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListScheduleEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListScheduleEventResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListScheduleEventResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScheduleEventResponseBodyData) GetRecords() []*ListScheduleEventResponseBodyDataRecords {
	return s.Records
}

func (s *ListScheduleEventResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListScheduleEventResponseBodyData) SetPageNumber(v int32) *ListScheduleEventResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListScheduleEventResponseBodyData) SetPageSize(v int32) *ListScheduleEventResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListScheduleEventResponseBodyData) SetRecords(v []*ListScheduleEventResponseBodyDataRecords) *ListScheduleEventResponseBodyData {
	s.Records = v
	return s
}

func (s *ListScheduleEventResponseBodyData) SetTotal(v int64) *ListScheduleEventResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListScheduleEventResponseBodyData) Validate() error {
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

type ListScheduleEventResponseBodyDataRecords struct {
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// hello word
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// INFO
	Event *string `json:"Event,omitempty" xml:"Event,omitempty"`
	// 130
	//
	// example:
	//
	// 1310630367761285120
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// test-job
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// example:
	//
	// 2024-10-31 16:43:51
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// http://192.168.1.5:9999/
	WorkerAddr *string `json:"WorkerAddr,omitempty" xml:"WorkerAddr,omitempty"`
}

func (s ListScheduleEventResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleEventResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListScheduleEventResponseBodyDataRecords) GetAppName() *string {
	return s.AppName
}

func (s *ListScheduleEventResponseBodyDataRecords) GetContent() *string {
	return s.Content
}

func (s *ListScheduleEventResponseBodyDataRecords) GetEvent() *string {
	return s.Event
}

func (s *ListScheduleEventResponseBodyDataRecords) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *ListScheduleEventResponseBodyDataRecords) GetJobName() *string {
	return s.JobName
}

func (s *ListScheduleEventResponseBodyDataRecords) GetTime() *string {
	return s.Time
}

func (s *ListScheduleEventResponseBodyDataRecords) GetWorkerAddr() *string {
	return s.WorkerAddr
}

func (s *ListScheduleEventResponseBodyDataRecords) SetAppName(v string) *ListScheduleEventResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *ListScheduleEventResponseBodyDataRecords) SetContent(v string) *ListScheduleEventResponseBodyDataRecords {
	s.Content = &v
	return s
}

func (s *ListScheduleEventResponseBodyDataRecords) SetEvent(v string) *ListScheduleEventResponseBodyDataRecords {
	s.Event = &v
	return s
}

func (s *ListScheduleEventResponseBodyDataRecords) SetJobExecutionId(v string) *ListScheduleEventResponseBodyDataRecords {
	s.JobExecutionId = &v
	return s
}

func (s *ListScheduleEventResponseBodyDataRecords) SetJobName(v string) *ListScheduleEventResponseBodyDataRecords {
	s.JobName = &v
	return s
}

func (s *ListScheduleEventResponseBodyDataRecords) SetTime(v string) *ListScheduleEventResponseBodyDataRecords {
	s.Time = &v
	return s
}

func (s *ListScheduleEventResponseBodyDataRecords) SetWorkerAddr(v string) *ListScheduleEventResponseBodyDataRecords {
	s.WorkerAddr = &v
	return s
}

func (s *ListScheduleEventResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
