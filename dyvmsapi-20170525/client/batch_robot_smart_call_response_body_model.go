// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRobotSmartCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchRobotSmartCallResponseBody
	GetCode() *string
	SetMessage(v string) *BatchRobotSmartCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchRobotSmartCallResponseBody
	GetRequestId() *string
	SetTaskId(v string) *BatchRobotSmartCallResponseBody
	GetTaskId() *string
}

type BatchRobotSmartCallResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The unique ID of the robocall task. You can call the [QueryCallDetailByTaskId](https://help.aliyun.com/document_detail/393537.html) operation to query the details of the task based on the task ID.
	//
	// example:
	//
	// 4001112222
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BatchRobotSmartCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchRobotSmartCallResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRobotSmartCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchRobotSmartCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchRobotSmartCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchRobotSmartCallResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *BatchRobotSmartCallResponseBody) SetCode(v string) *BatchRobotSmartCallResponseBody {
	s.Code = &v
	return s
}

func (s *BatchRobotSmartCallResponseBody) SetMessage(v string) *BatchRobotSmartCallResponseBody {
	s.Message = &v
	return s
}

func (s *BatchRobotSmartCallResponseBody) SetRequestId(v string) *BatchRobotSmartCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchRobotSmartCallResponseBody) SetTaskId(v string) *BatchRobotSmartCallResponseBody {
	s.TaskId = &v
	return s
}

func (s *BatchRobotSmartCallResponseBody) Validate() error {
	return dara.Validate(s)
}
