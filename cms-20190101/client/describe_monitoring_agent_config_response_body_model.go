// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoInstall(v bool) *DescribeMonitoringAgentConfigResponseBody
	GetAutoInstall() *bool
	SetCode(v string) *DescribeMonitoringAgentConfigResponseBody
	GetCode() *string
	SetEnableActiveAlert(v string) *DescribeMonitoringAgentConfigResponseBody
	GetEnableActiveAlert() *string
	SetEnableInstallAgentNewECS(v bool) *DescribeMonitoringAgentConfigResponseBody
	GetEnableInstallAgentNewECS() *bool
	SetMessage(v string) *DescribeMonitoringAgentConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMonitoringAgentConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMonitoringAgentConfigResponseBody
	GetSuccess() *bool
}

type DescribeMonitoringAgentConfigResponseBody struct {
	// Indicates whether the CloudMonitor agent is automatically installed on existing Elastic Compute Service (ECS) instances. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	AutoInstall *bool `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The cloud services for which proactive alerting is enabled.
	//
	// example:
	//
	// redis,rds,ecs
	EnableActiveAlert *string `json:"EnableActiveAlert,omitempty" xml:"EnableActiveAlert,omitempty"`
	// Indicates whether the CloudMonitor agent is automatically installed on newly purchased ECS instances. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableInstallAgentNewECS *bool `json:"EnableInstallAgentNewECS,omitempty" xml:"EnableInstallAgentNewECS,omitempty"`
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
	// E9F4FA2A-54BE-4EF9-9D1D-1A0B1DC86B8D
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

func (s DescribeMonitoringAgentConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentConfigResponseBody) GetAutoInstall() *bool {
	return s.AutoInstall
}

func (s *DescribeMonitoringAgentConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMonitoringAgentConfigResponseBody) GetEnableActiveAlert() *string {
	return s.EnableActiveAlert
}

func (s *DescribeMonitoringAgentConfigResponseBody) GetEnableInstallAgentNewECS() *bool {
	return s.EnableInstallAgentNewECS
}

func (s *DescribeMonitoringAgentConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitoringAgentConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitoringAgentConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetAutoInstall(v bool) *DescribeMonitoringAgentConfigResponseBody {
	s.AutoInstall = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetCode(v string) *DescribeMonitoringAgentConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetEnableActiveAlert(v string) *DescribeMonitoringAgentConfigResponseBody {
	s.EnableActiveAlert = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetEnableInstallAgentNewECS(v bool) *DescribeMonitoringAgentConfigResponseBody {
	s.EnableInstallAgentNewECS = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetMessage(v string) *DescribeMonitoringAgentConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetRequestId(v string) *DescribeMonitoringAgentConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponseBody) SetSuccess(v bool) *DescribeMonitoringAgentConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMonitoringAgentConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
