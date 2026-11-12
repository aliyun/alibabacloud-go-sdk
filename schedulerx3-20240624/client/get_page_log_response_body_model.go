// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetPageLogResponseBody
	GetCode() *int32
	SetData(v *GetPageLogResponseBodyData) *GetPageLogResponseBody
	GetData() *GetPageLogResponseBodyData
	SetMessage(v string) *GetPageLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPageLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPageLogResponseBody
	GetSuccess() *bool
}

type GetPageLogResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The paginated log results.
	Data *GetPageLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// CalendarName is already existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5EF879D0-3B43-5AD1-9BF7-52418F9C5E73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPageLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPageLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetPageLogResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetPageLogResponseBody) GetData() *GetPageLogResponseBodyData {
	return s.Data
}

func (s *GetPageLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPageLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPageLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPageLogResponseBody) SetCode(v int32) *GetPageLogResponseBody {
	s.Code = &v
	return s
}

func (s *GetPageLogResponseBody) SetData(v *GetPageLogResponseBodyData) *GetPageLogResponseBody {
	s.Data = v
	return s
}

func (s *GetPageLogResponseBody) SetMessage(v string) *GetPageLogResponseBody {
	s.Message = &v
	return s
}

func (s *GetPageLogResponseBody) SetRequestId(v string) *GetPageLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPageLogResponseBody) SetSuccess(v bool) *GetPageLogResponseBody {
	s.Success = &v
	return s
}

func (s *GetPageLogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPageLogResponseBodyData struct {
	// The current page number.
	//
	// example:
	//
	// 10
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of log records on the current page.
	Records []*GetPageLogResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 65
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetPageLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPageLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPageLogResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetPageLogResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetPageLogResponseBodyData) GetRecords() []*GetPageLogResponseBodyDataRecords {
	return s.Records
}

func (s *GetPageLogResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *GetPageLogResponseBodyData) SetPageNum(v int32) *GetPageLogResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *GetPageLogResponseBodyData) SetPageSize(v int32) *GetPageLogResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetPageLogResponseBodyData) SetRecords(v []*GetPageLogResponseBodyDataRecords) *GetPageLogResponseBodyData {
	s.Records = v
	return s
}

func (s *GetPageLogResponseBodyData) SetTotal(v int32) *GetPageLogResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetPageLogResponseBodyData) Validate() error {
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

type GetPageLogResponseBodyDataRecords struct {
	// The application name.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The job execution ID.
	//
	// example:
	//
	// 101
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// The job name.
	//
	// example:
	//
	// test
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The content output by the current command in the current output stream.
	//
	// example:
	//
	// [2026-02-10 10:23:36]tStart writing to Recall Enginen
	Log *string `json:"Log,omitempty" xml:"Log,omitempty"`
	// The print time in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2024-10-31 16:43:51
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The user ID.
	//
	// example:
	//
	// zhangsan
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The worker address.
	//
	// example:
	//
	// 030225016025_9357_60125@127.0.0.1:51363
	WorkerAddr *string `json:"WorkerAddr,omitempty" xml:"WorkerAddr,omitempty"`
}

func (s GetPageLogResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s GetPageLogResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *GetPageLogResponseBodyDataRecords) GetAppName() *string {
	return s.AppName
}

func (s *GetPageLogResponseBodyDataRecords) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *GetPageLogResponseBodyDataRecords) GetJobName() *string {
	return s.JobName
}

func (s *GetPageLogResponseBodyDataRecords) GetLog() *string {
	return s.Log
}

func (s *GetPageLogResponseBodyDataRecords) GetTime() *string {
	return s.Time
}

func (s *GetPageLogResponseBodyDataRecords) GetUserId() *string {
	return s.UserId
}

func (s *GetPageLogResponseBodyDataRecords) GetWorkerAddr() *string {
	return s.WorkerAddr
}

func (s *GetPageLogResponseBodyDataRecords) SetAppName(v string) *GetPageLogResponseBodyDataRecords {
	s.AppName = &v
	return s
}

func (s *GetPageLogResponseBodyDataRecords) SetJobExecutionId(v string) *GetPageLogResponseBodyDataRecords {
	s.JobExecutionId = &v
	return s
}

func (s *GetPageLogResponseBodyDataRecords) SetJobName(v string) *GetPageLogResponseBodyDataRecords {
	s.JobName = &v
	return s
}

func (s *GetPageLogResponseBodyDataRecords) SetLog(v string) *GetPageLogResponseBodyDataRecords {
	s.Log = &v
	return s
}

func (s *GetPageLogResponseBodyDataRecords) SetTime(v string) *GetPageLogResponseBodyDataRecords {
	s.Time = &v
	return s
}

func (s *GetPageLogResponseBodyDataRecords) SetUserId(v string) *GetPageLogResponseBodyDataRecords {
	s.UserId = &v
	return s
}

func (s *GetPageLogResponseBodyDataRecords) SetWorkerAddr(v string) *GetPageLogResponseBodyDataRecords {
	s.WorkerAddr = &v
	return s
}

func (s *GetPageLogResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
