// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupNotifyPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMonitorGroupNotifyPolicyResponseBody
	GetCode() *string
	SetMessage(v string) *CreateMonitorGroupNotifyPolicyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMonitorGroupNotifyPolicyResponseBody
	GetRequestId() *string
	SetResult(v int32) *CreateMonitorGroupNotifyPolicyResponseBody
	GetResult() *int32
	SetSuccess(v string) *CreateMonitorGroupNotifyPolicyResponseBody
	GetSuccess() *string
}

type CreateMonitorGroupNotifyPolicyResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
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
	// 13356BCA-3EC3-4748-A771-2064DA69AEF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
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

func (s CreateMonitorGroupNotifyPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupNotifyPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) SetCode(v string) *CreateMonitorGroupNotifyPolicyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) SetMessage(v string) *CreateMonitorGroupNotifyPolicyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) SetRequestId(v string) *CreateMonitorGroupNotifyPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) SetResult(v int32) *CreateMonitorGroupNotifyPolicyResponseBody {
	s.Result = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) SetSuccess(v string) *CreateMonitorGroupNotifyPolicyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMonitorGroupNotifyPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
