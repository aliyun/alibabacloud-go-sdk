// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallMonitoringAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InstallMonitoringAgentResponseBody
	GetCode() *string
	SetMessage(v string) *InstallMonitoringAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstallMonitoringAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InstallMonitoringAgentResponseBody
	GetSuccess() *bool
}

type InstallMonitoringAgentResponseBody struct {
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
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0BDAF8A8-04DC-5F0C-90E4-724D42C41945
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

func (s InstallMonitoringAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallMonitoringAgentResponseBody) GoString() string {
	return s.String()
}

func (s *InstallMonitoringAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *InstallMonitoringAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallMonitoringAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallMonitoringAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InstallMonitoringAgentResponseBody) SetCode(v string) *InstallMonitoringAgentResponseBody {
	s.Code = &v
	return s
}

func (s *InstallMonitoringAgentResponseBody) SetMessage(v string) *InstallMonitoringAgentResponseBody {
	s.Message = &v
	return s
}

func (s *InstallMonitoringAgentResponseBody) SetRequestId(v string) *InstallMonitoringAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallMonitoringAgentResponseBody) SetSuccess(v bool) *InstallMonitoringAgentResponseBody {
	s.Success = &v
	return s
}

func (s *InstallMonitoringAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
