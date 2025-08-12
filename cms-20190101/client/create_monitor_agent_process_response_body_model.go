// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorAgentProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMonitorAgentProcessResponseBody
	GetCode() *string
	SetId(v int64) *CreateMonitorAgentProcessResponseBody
	GetId() *int64
	SetMessage(v string) *CreateMonitorAgentProcessResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMonitorAgentProcessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMonitorAgentProcessResponseBody
	GetSuccess() *bool
}

type CreateMonitorAgentProcessResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The process ID.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 971CC023-5A96-452A-BB7C-2483F948BCFD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMonitorAgentProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorAgentProcessResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMonitorAgentProcessResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateMonitorAgentProcessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMonitorAgentProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMonitorAgentProcessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMonitorAgentProcessResponseBody) SetCode(v string) *CreateMonitorAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMonitorAgentProcessResponseBody) SetId(v int64) *CreateMonitorAgentProcessResponseBody {
	s.Id = &v
	return s
}

func (s *CreateMonitorAgentProcessResponseBody) SetMessage(v string) *CreateMonitorAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMonitorAgentProcessResponseBody) SetRequestId(v string) *CreateMonitorAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitorAgentProcessResponseBody) SetSuccess(v bool) *CreateMonitorAgentProcessResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMonitorAgentProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
