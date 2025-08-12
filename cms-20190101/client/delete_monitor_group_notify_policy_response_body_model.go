// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupNotifyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMonitorGroupNotifyPolicyResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteMonitorGroupNotifyPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMonitorGroupNotifyPolicyResponseBody
	GetRequestId() *string
	SetResult(v int32) *DeleteMonitorGroupNotifyPolicyResponseBody
	GetResult() *int32
	SetSuccess(v string) *DeleteMonitorGroupNotifyPolicyResponseBody
	GetSuccess() *string
}

type DeleteMonitorGroupNotifyPolicyResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B7AF834D-D38B-4A46-920B-FE974EB7E135
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of affected rows.
	//
	// example:
	//
	// 1
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMonitorGroupNotifyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupNotifyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) SetCode(v string) *DeleteMonitorGroupNotifyPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) SetMessage(v string) *DeleteMonitorGroupNotifyPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) SetRequestId(v string) *DeleteMonitorGroupNotifyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) SetResult(v int32) *DeleteMonitorGroupNotifyPolicyResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) SetSuccess(v string) *DeleteMonitorGroupNotifyPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMonitorGroupNotifyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
