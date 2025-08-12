// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallMonitoringAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UninstallMonitoringAgentResponseBody
	GetCode() *string
	SetMessage(v string) *UninstallMonitoringAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UninstallMonitoringAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UninstallMonitoringAgentResponseBody
	GetSuccess() *bool
}

type UninstallMonitoringAgentResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 466902B9-2842-40B0-B796-00FE772B6EF3
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

func (s UninstallMonitoringAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallMonitoringAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallMonitoringAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UninstallMonitoringAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UninstallMonitoringAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallMonitoringAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UninstallMonitoringAgentResponseBody) SetCode(v string) *UninstallMonitoringAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallMonitoringAgentResponseBody) SetMessage(v string) *UninstallMonitoringAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallMonitoringAgentResponseBody) SetRequestId(v string) *UninstallMonitoringAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallMonitoringAgentResponseBody) SetSuccess(v bool) *UninstallMonitoringAgentResponseBody {
	s.Success = &v
	return s
}

func (s *UninstallMonitoringAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
