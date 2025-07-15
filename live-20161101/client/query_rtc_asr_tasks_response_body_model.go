// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRtcAsrTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *QueryRtcAsrTasksResponseBody
	GetDescription() *string
	SetRequestId(v string) *QueryRtcAsrTasksResponseBody
	GetRequestId() *string
	SetRetCode(v int64) *QueryRtcAsrTasksResponseBody
	GetRetCode() *int64
	SetTasks(v map[string]interface{}) *QueryRtcAsrTasksResponseBody
	GetTasks() map[string]interface{}
}

type QueryRtcAsrTasksResponseBody struct {
	// The result of the request. If success is returned, the request was successful. If an error message is returned, the request failed.
	//
	// example:
	//
	// success
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D8ADAB55-1BB8-5C01-8434-C45D353BB1FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP status code. HTTP status code 2000 indicates that the request was successful. Other HTTP status codes indicate that the request failed.
	//
	// example:
	//
	// 2000
	RetCode *int64 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
	// The results returned for the tasks.
	Tasks map[string]interface{} `json:"Tasks,omitempty" xml:"Tasks,omitempty"`
}

func (s QueryRtcAsrTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRtcAsrTasksResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRtcAsrTasksResponseBody) GetDescription() *string {
	return s.Description
}

func (s *QueryRtcAsrTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRtcAsrTasksResponseBody) GetRetCode() *int64 {
	return s.RetCode
}

func (s *QueryRtcAsrTasksResponseBody) GetTasks() map[string]interface{} {
	return s.Tasks
}

func (s *QueryRtcAsrTasksResponseBody) SetDescription(v string) *QueryRtcAsrTasksResponseBody {
	s.Description = &v
	return s
}

func (s *QueryRtcAsrTasksResponseBody) SetRequestId(v string) *QueryRtcAsrTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRtcAsrTasksResponseBody) SetRetCode(v int64) *QueryRtcAsrTasksResponseBody {
	s.RetCode = &v
	return s
}

func (s *QueryRtcAsrTasksResponseBody) SetTasks(v map[string]interface{}) *QueryRtcAsrTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *QueryRtcAsrTasksResponseBody) Validate() error {
	return dara.Validate(s)
}
