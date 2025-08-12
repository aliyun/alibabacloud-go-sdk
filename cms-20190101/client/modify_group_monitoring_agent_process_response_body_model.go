// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGroupMonitoringAgentProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyGroupMonitoringAgentProcessResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyGroupMonitoringAgentProcessResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyGroupMonitoringAgentProcessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyGroupMonitoringAgentProcessResponseBody
	GetSuccess() *bool
}

type ModifyGroupMonitoringAgentProcessResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
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
	// The ID of the request.
	//
	// example:
	//
	// 7985D471-3FA8-4EE9-8F4B-45C19DF3D36F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- true: The call was successful.
	//
	// 	- false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyGroupMonitoringAgentProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGroupMonitoringAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGroupMonitoringAgentProcessResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyGroupMonitoringAgentProcessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyGroupMonitoringAgentProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGroupMonitoringAgentProcessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyGroupMonitoringAgentProcessResponseBody) SetCode(v string) *ModifyGroupMonitoringAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessResponseBody) SetMessage(v string) *ModifyGroupMonitoringAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessResponseBody) SetRequestId(v string) *ModifyGroupMonitoringAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessResponseBody) SetSuccess(v bool) *ModifyGroupMonitoringAgentProcessResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyGroupMonitoringAgentProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
