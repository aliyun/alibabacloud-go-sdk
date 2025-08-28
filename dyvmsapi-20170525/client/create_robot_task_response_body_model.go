// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRobotTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRobotTaskResponseBody
	GetCode() *string
	SetData(v string) *CreateRobotTaskResponseBody
	GetData() *string
	SetMessage(v string) *CreateRobotTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRobotTaskResponseBody
	GetRequestId() *string
}

type CreateRobotTaskResponseBody struct {
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
	// The unique ID of the robocall task.
	//
	// You can call the [QueryRobotTaskDetail](https://help.aliyun.com/document_detail/393538.html) operation to query the details of the task based on the task ID.
	//
	// example:
	//
	// 400111****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
}

func (s CreateRobotTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRobotTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRobotTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateRobotTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRobotTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRobotTaskResponseBody) SetCode(v string) *CreateRobotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRobotTaskResponseBody) SetData(v string) *CreateRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRobotTaskResponseBody) SetMessage(v string) *CreateRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRobotTaskResponseBody) SetRequestId(v string) *CreateRobotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRobotTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
