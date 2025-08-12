// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoInstall(v bool) *DescribeMonitoringConfigResponseBody
	GetAutoInstall() *bool
	SetCode(v string) *DescribeMonitoringConfigResponseBody
	GetCode() *string
	SetEnableInstallAgentNewECS(v bool) *DescribeMonitoringConfigResponseBody
	GetEnableInstallAgentNewECS() *bool
	SetMessage(v string) *DescribeMonitoringConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMonitoringConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMonitoringConfigResponseBody
	GetSuccess() *bool
}

type DescribeMonitoringConfigResponseBody struct {
	// Indicates whether the CloudMonitor agent is automatically installed on existing Elastic Compute Service (ECS) instances. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AutoInstall *bool `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the CloudMonitor agent is automatically installed on new ECS instances. Valid values: Valid values:
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
	// F35654DB-0C9D-4FB3-903F-479BA7663061
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

func (s DescribeMonitoringConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringConfigResponseBody) GetAutoInstall() *bool {
	return s.AutoInstall
}

func (s *DescribeMonitoringConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMonitoringConfigResponseBody) GetEnableInstallAgentNewECS() *bool {
	return s.EnableInstallAgentNewECS
}

func (s *DescribeMonitoringConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitoringConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitoringConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMonitoringConfigResponseBody) SetAutoInstall(v bool) *DescribeMonitoringConfigResponseBody {
	s.AutoInstall = &v
	return s
}

func (s *DescribeMonitoringConfigResponseBody) SetCode(v string) *DescribeMonitoringConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitoringConfigResponseBody) SetEnableInstallAgentNewECS(v bool) *DescribeMonitoringConfigResponseBody {
	s.EnableInstallAgentNewECS = &v
	return s
}

func (s *DescribeMonitoringConfigResponseBody) SetMessage(v string) *DescribeMonitoringConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitoringConfigResponseBody) SetRequestId(v string) *DescribeMonitoringConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitoringConfigResponseBody) SetSuccess(v bool) *DescribeMonitoringConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMonitoringConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
