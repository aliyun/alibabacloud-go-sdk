// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvPodMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeEnvPodMonitorResponseBody
	GetCode() *int32
	SetData(v *DescribeEnvPodMonitorResponseBodyData) *DescribeEnvPodMonitorResponseBody
	GetData() *DescribeEnvPodMonitorResponseBodyData
	SetMessage(v string) *DescribeEnvPodMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEnvPodMonitorResponseBody
	GetRequestId() *string
}

type DescribeEnvPodMonitorResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *DescribeEnvPodMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnvPodMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvPodMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnvPodMonitorResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeEnvPodMonitorResponseBody) GetData() *DescribeEnvPodMonitorResponseBodyData {
	return s.Data
}

func (s *DescribeEnvPodMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEnvPodMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnvPodMonitorResponseBody) SetCode(v int32) *DescribeEnvPodMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnvPodMonitorResponseBody) SetData(v *DescribeEnvPodMonitorResponseBodyData) *DescribeEnvPodMonitorResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEnvPodMonitorResponseBody) SetMessage(v string) *DescribeEnvPodMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEnvPodMonitorResponseBody) SetRequestId(v string) *DescribeEnvPodMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnvPodMonitorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEnvPodMonitorResponseBodyData struct {
	// The YAML string of the PodMonitor.
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
	// The name of the PodMonitor.
	//
	// example:
	//
	// podMonitor1
	PodMonitorName *string `json:"PodMonitorName,omitempty" xml:"PodMonitorName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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

func (s DescribeEnvPodMonitorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvPodMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEnvPodMonitorResponseBodyData) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *DescribeEnvPodMonitorResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DescribeEnvPodMonitorResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeEnvPodMonitorResponseBodyData) GetPodMonitorName() *string {
	return s.PodMonitorName
}

func (s *DescribeEnvPodMonitorResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnvPodMonitorResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeEnvPodMonitorResponseBodyData) SetConfigYaml(v string) *DescribeEnvPodMonitorResponseBodyData {
	s.ConfigYaml = &v
	return s
}

func (s *DescribeEnvPodMonitorResponseBodyData) SetEnvironmentId(v string) *DescribeEnvPodMonitorResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *DescribeEnvPodMonitorResponseBodyData) SetNamespace(v string) *DescribeEnvPodMonitorResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *DescribeEnvPodMonitorResponseBodyData) SetPodMonitorName(v string) *DescribeEnvPodMonitorResponseBodyData {
	s.PodMonitorName = &v
	return s
}

func (s *DescribeEnvPodMonitorResponseBodyData) SetRegionId(v string) *DescribeEnvPodMonitorResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeEnvPodMonitorResponseBodyData) SetStatus(v string) *DescribeEnvPodMonitorResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeEnvPodMonitorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
