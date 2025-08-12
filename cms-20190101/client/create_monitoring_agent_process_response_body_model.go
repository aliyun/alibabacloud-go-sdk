// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitoringAgentProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMonitoringAgentProcessResponseBody
	GetCode() *string
	SetId(v int64) *CreateMonitoringAgentProcessResponseBody
	GetId() *int64
	SetMessage(v string) *CreateMonitoringAgentProcessResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMonitoringAgentProcessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMonitoringAgentProcessResponseBody
	GetSuccess() *bool
}

type CreateMonitoringAgentProcessResponseBody struct {
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
	// 12345
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
	// 0DFCB47D-E7E1-4CBE-A381-8339F7B300EF
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

func (s CreateMonitoringAgentProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitoringAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitoringAgentProcessResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMonitoringAgentProcessResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateMonitoringAgentProcessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMonitoringAgentProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMonitoringAgentProcessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMonitoringAgentProcessResponseBody) SetCode(v string) *CreateMonitoringAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMonitoringAgentProcessResponseBody) SetId(v int64) *CreateMonitoringAgentProcessResponseBody {
	s.Id = &v
	return s
}

func (s *CreateMonitoringAgentProcessResponseBody) SetMessage(v string) *CreateMonitoringAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMonitoringAgentProcessResponseBody) SetRequestId(v string) *CreateMonitoringAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitoringAgentProcessResponseBody) SetSuccess(v bool) *CreateMonitoringAgentProcessResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMonitoringAgentProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
