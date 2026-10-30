// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAifgJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAifgJobsResponseBody
	GetCode() *string
	SetData(v []*ListAifgJobsResponseBodyData) *ListAifgJobsResponseBody
	GetData() []*ListAifgJobsResponseBodyData
	SetMaxResults(v int32) *ListAifgJobsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListAifgJobsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListAifgJobsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAifgJobsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListAifgJobsResponseBody
	GetTotal() *int64
}

type ListAifgJobsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data under the root node.
	Data []*ListAifgJobsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The maximum number of records to retrieve at a time.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// Indicates the read position returned by the current call. An empty value indicates that all data has been read.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hello, world!
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination cursor. A non-empty value indicates that more data is available.
	//
	// example:
	//
	// U+w1wv2R4ZWR5oZLXD0+Dp4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The unique identifier of the request.
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total count.
	//
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAifgJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAifgJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAifgJobsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAifgJobsResponseBody) GetData() []*ListAifgJobsResponseBodyData {
	return s.Data
}

func (s *ListAifgJobsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAifgJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAifgJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAifgJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAifgJobsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAifgJobsResponseBody) SetCode(v string) *ListAifgJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAifgJobsResponseBody) SetData(v []*ListAifgJobsResponseBodyData) *ListAifgJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListAifgJobsResponseBody) SetMaxResults(v int32) *ListAifgJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAifgJobsResponseBody) SetMessage(v string) *ListAifgJobsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAifgJobsResponseBody) SetNextToken(v string) *ListAifgJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAifgJobsResponseBody) SetRequestId(v string) *ListAifgJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAifgJobsResponseBody) SetTotal(v int64) *ListAifgJobsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAifgJobsResponseBody) Validate() error {
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

type ListAifgJobsResponseBodyData struct {
	// The first command in the data list.
	//
	// example:
	//
	// start
	Command interface{} `json:"command,omitempty" xml:"command,omitempty"`
	// The ID of the first element in the data list.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The first parameter in the root data.
	//
	// example:
	//
	// {\\"key1\\": \\"value1\\"}
	Params interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// The first result in the root data.
	//
	// example:
	//
	// success
	Result interface{} `json:"result,omitempty" xml:"result,omitempty"`
	// The service name.
	//
	// example:
	//
	// user-service
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// The status of the data item.
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The sub-status of data item 0.
	//
	// example:
	//
	// processing
	SubStatus *string `json:"subStatus,omitempty" xml:"subStatus,omitempty"`
	// The time when the data was submitted.
	//
	// example:
	//
	// 2023-09-01 12:00:00
	SubmitTime *string `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 123456
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The unique identifier of the trace task.
	//
	// example:
	//
	// T123456789
	TraceTaskId *string `json:"traceTaskId,omitempty" xml:"traceTaskId,omitempty"`
}

func (s ListAifgJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAifgJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAifgJobsResponseBodyData) GetCommand() interface{} {
	return s.Command
}

func (s *ListAifgJobsResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListAifgJobsResponseBodyData) GetParams() interface{} {
	return s.Params
}

func (s *ListAifgJobsResponseBodyData) GetResult() interface{} {
	return s.Result
}

func (s *ListAifgJobsResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListAifgJobsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListAifgJobsResponseBodyData) GetSubStatus() *string {
	return s.SubStatus
}

func (s *ListAifgJobsResponseBodyData) GetSubmitTime() *string {
	return s.SubmitTime
}

func (s *ListAifgJobsResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAifgJobsResponseBodyData) GetTraceTaskId() *string {
	return s.TraceTaskId
}

func (s *ListAifgJobsResponseBodyData) SetCommand(v interface{}) *ListAifgJobsResponseBodyData {
	s.Command = v
	return s
}

func (s *ListAifgJobsResponseBodyData) SetId(v int64) *ListAifgJobsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAifgJobsResponseBodyData) SetParams(v interface{}) *ListAifgJobsResponseBodyData {
	s.Params = v
	return s
}

func (s *ListAifgJobsResponseBodyData) SetResult(v interface{}) *ListAifgJobsResponseBodyData {
	s.Result = v
	return s
}

func (s *ListAifgJobsResponseBodyData) SetServiceName(v string) *ListAifgJobsResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListAifgJobsResponseBodyData) SetStatus(v string) *ListAifgJobsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAifgJobsResponseBodyData) SetSubStatus(v string) *ListAifgJobsResponseBodyData {
	s.SubStatus = &v
	return s
}

func (s *ListAifgJobsResponseBodyData) SetSubmitTime(v string) *ListAifgJobsResponseBodyData {
	s.SubmitTime = &v
	return s
}

func (s *ListAifgJobsResponseBodyData) SetTaskId(v string) *ListAifgJobsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListAifgJobsResponseBodyData) SetTraceTaskId(v string) *ListAifgJobsResponseBodyData {
	s.TraceTaskId = &v
	return s
}

func (s *ListAifgJobsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
