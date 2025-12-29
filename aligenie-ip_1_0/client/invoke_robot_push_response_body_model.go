// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeRobotPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *InvokeRobotPushResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvokeRobotPushResponseBody
	GetRequestId() *string
	SetResult(v bool) *InvokeRobotPushResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *InvokeRobotPushResponseBody
	GetStatusCode() *int32
}

type InvokeRobotPushResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C6***E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s InvokeRobotPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvokeRobotPushResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeRobotPushResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvokeRobotPushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvokeRobotPushResponseBody) GetResult() *bool {
	return s.Result
}

func (s *InvokeRobotPushResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InvokeRobotPushResponseBody) SetMessage(v string) *InvokeRobotPushResponseBody {
	s.Message = &v
	return s
}

func (s *InvokeRobotPushResponseBody) SetRequestId(v string) *InvokeRobotPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeRobotPushResponseBody) SetResult(v bool) *InvokeRobotPushResponseBody {
	s.Result = &v
	return s
}

func (s *InvokeRobotPushResponseBody) SetStatusCode(v int32) *InvokeRobotPushResponseBody {
	s.StatusCode = &v
	return s
}

func (s *InvokeRobotPushResponseBody) Validate() error {
	return dara.Validate(s)
}
