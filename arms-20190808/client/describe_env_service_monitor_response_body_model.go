// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvServiceMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeEnvServiceMonitorResponseBody
	GetCode() *int32
	SetData(v *DescribeEnvServiceMonitorResponseBodyData) *DescribeEnvServiceMonitorResponseBody
	GetData() *DescribeEnvServiceMonitorResponseBodyData
	SetMessage(v string) *DescribeEnvServiceMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEnvServiceMonitorResponseBody
	GetRequestId() *string
}

type DescribeEnvServiceMonitorResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *DescribeEnvServiceMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40B10E04-81E8-4643-970D-F1B38F2E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnvServiceMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvServiceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnvServiceMonitorResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeEnvServiceMonitorResponseBody) GetData() *DescribeEnvServiceMonitorResponseBodyData {
	return s.Data
}

func (s *DescribeEnvServiceMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEnvServiceMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnvServiceMonitorResponseBody) SetCode(v int32) *DescribeEnvServiceMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnvServiceMonitorResponseBody) SetData(v *DescribeEnvServiceMonitorResponseBodyData) *DescribeEnvServiceMonitorResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEnvServiceMonitorResponseBody) SetMessage(v string) *DescribeEnvServiceMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEnvServiceMonitorResponseBody) SetRequestId(v string) *DescribeEnvServiceMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnvServiceMonitorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEnvServiceMonitorResponseBodyData struct {
	// The YAML configuration file of the ServiceMonitor.
	//
	// example:
	//
	// Refer to supplementary instructions.
	ConfigYaml *string `json:"ConfigYaml,omitempty" xml:"ConfigYaml,omitempty"`
	// The ID of the environment instance.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The namespace.
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the ServiceMonitor.
	//
	// example:
	//
	// serviceMonitor1
	ServiceMonitorName *string `json:"ServiceMonitorName,omitempty" xml:"ServiceMonitorName,omitempty"`
	// The status. Valid values:
	//
	// 	- run
	//
	// 	- stop
	//
	// example:
	//
	// run
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEnvServiceMonitorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvServiceMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEnvServiceMonitorResponseBodyData) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *DescribeEnvServiceMonitorResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DescribeEnvServiceMonitorResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeEnvServiceMonitorResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnvServiceMonitorResponseBodyData) GetServiceMonitorName() *string {
	return s.ServiceMonitorName
}

func (s *DescribeEnvServiceMonitorResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeEnvServiceMonitorResponseBodyData) SetConfigYaml(v string) *DescribeEnvServiceMonitorResponseBodyData {
	s.ConfigYaml = &v
	return s
}

func (s *DescribeEnvServiceMonitorResponseBodyData) SetEnvironmentId(v string) *DescribeEnvServiceMonitorResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *DescribeEnvServiceMonitorResponseBodyData) SetNamespace(v string) *DescribeEnvServiceMonitorResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *DescribeEnvServiceMonitorResponseBodyData) SetRegionId(v string) *DescribeEnvServiceMonitorResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeEnvServiceMonitorResponseBodyData) SetServiceMonitorName(v string) *DescribeEnvServiceMonitorResponseBodyData {
	s.ServiceMonitorName = &v
	return s
}

func (s *DescribeEnvServiceMonitorResponseBodyData) SetStatus(v string) *DescribeEnvServiceMonitorResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeEnvServiceMonitorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
