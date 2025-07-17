// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvCustomJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeEnvCustomJobResponseBody
	GetCode() *int32
	SetData(v *DescribeEnvCustomJobResponseBodyData) *DescribeEnvCustomJobResponseBody
	GetData() *DescribeEnvCustomJobResponseBodyData
	SetMessage(v string) *DescribeEnvCustomJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEnvCustomJobResponseBody
	GetRequestId() *string
}

type DescribeEnvCustomJobResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *DescribeEnvCustomJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 6A9AEA84-7186-4D8D-B498-4585C6A2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEnvCustomJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvCustomJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnvCustomJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeEnvCustomJobResponseBody) GetData() *DescribeEnvCustomJobResponseBodyData {
	return s.Data
}

func (s *DescribeEnvCustomJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEnvCustomJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnvCustomJobResponseBody) SetCode(v int32) *DescribeEnvCustomJobResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEnvCustomJobResponseBody) SetData(v *DescribeEnvCustomJobResponseBodyData) *DescribeEnvCustomJobResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEnvCustomJobResponseBody) SetMessage(v string) *DescribeEnvCustomJobResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEnvCustomJobResponseBody) SetRequestId(v string) *DescribeEnvCustomJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnvCustomJobResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEnvCustomJobResponseBodyData struct {
	// If the **encryptYaml*	- parameter is set to `true`, a Base64-encoded YAML string is returned. Otherwise, a plaintext YAML string is returned.
	//
	// example:
	//
	// Refer to supplementary instructions.
	ConfigYaml *string `json:"ConfigYaml,omitempty" xml:"ConfigYaml,omitempty"`
	// The name of the custom job.
	//
	// example:
	//
	// customJob1
	CustomJobName *string `json:"CustomJobName,omitempty" xml:"CustomJobName,omitempty"`
	// The ID of the environment instance.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
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

func (s DescribeEnvCustomJobResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvCustomJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEnvCustomJobResponseBodyData) GetConfigYaml() *string {
	return s.ConfigYaml
}

func (s *DescribeEnvCustomJobResponseBodyData) GetCustomJobName() *string {
	return s.CustomJobName
}

func (s *DescribeEnvCustomJobResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DescribeEnvCustomJobResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnvCustomJobResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeEnvCustomJobResponseBodyData) SetConfigYaml(v string) *DescribeEnvCustomJobResponseBodyData {
	s.ConfigYaml = &v
	return s
}

func (s *DescribeEnvCustomJobResponseBodyData) SetCustomJobName(v string) *DescribeEnvCustomJobResponseBodyData {
	s.CustomJobName = &v
	return s
}

func (s *DescribeEnvCustomJobResponseBodyData) SetEnvironmentId(v string) *DescribeEnvCustomJobResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *DescribeEnvCustomJobResponseBodyData) SetRegionId(v string) *DescribeEnvCustomJobResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeEnvCustomJobResponseBodyData) SetStatus(v string) *DescribeEnvCustomJobResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeEnvCustomJobResponseBodyData) Validate() error {
	return dara.Validate(s)
}
