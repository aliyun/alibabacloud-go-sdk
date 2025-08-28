// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRobotTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StopRobotTaskResponseBody
	GetCode() *string
	SetData(v string) *StopRobotTaskResponseBody
	GetData() *string
	SetMessage(v string) *StopRobotTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopRobotTaskResponseBody
	GetRequestId() *string
}

type StopRobotTaskResponseBody struct {
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
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
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

func (s StopRobotTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopRobotTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopRobotTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *StopRobotTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *StopRobotTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopRobotTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopRobotTaskResponseBody) SetCode(v string) *StopRobotTaskResponseBody {
	s.Code = &v
	return s
}

func (s *StopRobotTaskResponseBody) SetData(v string) *StopRobotTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StopRobotTaskResponseBody) SetMessage(v string) *StopRobotTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopRobotTaskResponseBody) SetRequestId(v string) *StopRobotTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopRobotTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
