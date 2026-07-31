// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFormationTaskByIDResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryFormationTaskByIDResponseBody
	GetCode() *string
	SetData(v string) *QueryFormationTaskByIDResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *QueryFormationTaskByIDResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryFormationTaskByIDResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryFormationTaskByIDResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryFormationTaskByIDResponseBody
	GetSuccess() *bool
}

type QueryFormationTaskByIDResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The task details.
	//
	// example:
	//
	// {
	//
	//     "id": 123,
	//
	//     "taskName": "sale_db",
	//
	//     "scheduleState": "NORMAL",
	//
	//     "frequency": {"type": "custom", "cron": "0 0/1 	- 	- 	- ?"},
	//
	//     "failedCount": 0,
	//
	//     "lastTaskInstContent": "{\\"task_inst_id\\":67890,\\"state\\":\\"SUCCESS\\"}"
	//
	//   }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5DC10091-348D-12B1-906D-AB49D658012E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryFormationTaskByIDResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryFormationTaskByIDResponseBody) GoString() string {
	return s.String()
}

func (s *QueryFormationTaskByIDResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryFormationTaskByIDResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryFormationTaskByIDResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryFormationTaskByIDResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryFormationTaskByIDResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryFormationTaskByIDResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryFormationTaskByIDResponseBody) SetCode(v string) *QueryFormationTaskByIDResponseBody {
	s.Code = &v
	return s
}

func (s *QueryFormationTaskByIDResponseBody) SetData(v string) *QueryFormationTaskByIDResponseBody {
	s.Data = &v
	return s
}

func (s *QueryFormationTaskByIDResponseBody) SetHttpStatusCode(v int32) *QueryFormationTaskByIDResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryFormationTaskByIDResponseBody) SetMessage(v string) *QueryFormationTaskByIDResponseBody {
	s.Message = &v
	return s
}

func (s *QueryFormationTaskByIDResponseBody) SetRequestId(v string) *QueryFormationTaskByIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryFormationTaskByIDResponseBody) SetSuccess(v bool) *QueryFormationTaskByIDResponseBody {
	s.Success = &v
	return s
}

func (s *QueryFormationTaskByIDResponseBody) Validate() error {
	return dara.Validate(s)
}
