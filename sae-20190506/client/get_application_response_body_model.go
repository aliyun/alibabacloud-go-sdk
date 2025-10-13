// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody
	GetApplication() *GetApplicationResponseBodyApplication
	SetMessage(v string) *GetApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApplicationResponseBody
	GetRequestId() *string
	SetTraceId(v string) *GetApplicationResponseBody
	GetTraceId() *string
}

type GetApplicationResponseBody struct {
	// The details of the application.
	Application *GetApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The additional information returned. Valid values:
	//
	// 	- When a request is successful, **success**is returned.
	//
	// 	- An error code is returned when a request failed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// example:
	//
	// ac1a0b2215622920113732501e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) GetApplication() *GetApplicationResponseBodyApplication {
	return s.Application
}

func (s *GetApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GetApplicationResponseBody) SetApplication(v *GetApplicationResponseBodyApplication) *GetApplicationResponseBody {
	s.Application = v
	return s
}

func (s *GetApplicationResponseBody) SetMessage(v string) *GetApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationResponseBody) SetTraceId(v string) *GetApplicationResponseBody {
	s.TraceId = &v
	return s
}

func (s *GetApplicationResponseBody) Validate() error {
	if s.Application != nil {
		if err := s.Application.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationResponseBodyApplication struct {
	// The description of the application.
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	// The application ID.
	//
	// example:
	//
	// 443d638a-ef76-47c4-b707-61197d******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the basic application.
	//
	// example:
	//
	// ee99cce6-1c8e-4bfa-96c3-3e2fa9******
	BaseAppId *string `json:"BaseAppId,omitempty" xml:"BaseAppId,omitempty"`
	// The CPU specifications that are required for each instance. Unit: millicores. This parameter cannot be set to 0. Valid values:
	//
	// 	- **500**
	//
	// 	- **1000**
	//
	// 	- **2000**
	//
	// 	- **4000**
	//
	// 	- **8000**
	//
	// 	- **12000**
	//
	// 	- **16000**
	//
	// 	- **32000**
	//
	// example:
	//
	// 2000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of application instances.
	//
	// example:
	//
	// i-8ps2o182102o1jv05bys
	Instances  *int32 `json:"Instances,omitempty" xml:"Instances,omitempty"`
	IsStateful *bool  `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// The memory size that is required by each instance. Unit: MB. This parameter cannot be set to 0. The values of this parameter correspond to the values of the Cpu parameter:
	//
	// 	- This parameter is set to **1024*	- if the Cpu parameter is set to 500 or 1000.
	//
	// 	- This parameter is set to **2048*	- if the Cpu parameter is set to 500, 1000, or 2000.
	//
	// 	- This parameter is set to **4096*	- if the Cpu parameter is set to 1000, 2000, or 4000.
	//
	// 	- This parameter is set to **8192*	- if the Cpu parameter is set to 2000, 4000, or 8000.
	//
	// 	- This parameter is set to **12288*	- if the Cpu parameter is set to 12000.
	//
	// 	- This parameter is set to **16384*	- if the Cpu parameter is set to 4000, 8000, or 16000.
	//
	// 	- This parameter is set to **24576*	- if the Cpu parameter is set to 12000.
	//
	// 	- This parameter is set to **32768*	- if the Cpu parameter is set to 16000.
	//
	// 	- This parameter is set to **65536*	- if the Cpu parameter is set to 8000, 16000, or 32000.
	//
	// 	- This parameter is set to **131072*	- if the Cpu parameter is set to 32000.
	//
	// example:
	//
	// 4096
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// Specifies whether to enable WebAssembly Filter. Valid values:
	//
	// 	- true: enables this parameter.
	//
	// 	- false: disables this parameter.
	//
	// example:
	//
	// true
	MseEnabled *bool `json:"MseEnabled,omitempty" xml:"MseEnabled,omitempty"`
	// The ID of the namespace to which the MSE instance belongs.
	//
	// example:
	//
	// test
	MseNamespaceId *string `json:"MseNamespaceId,omitempty" xml:"MseNamespaceId,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-shenzhen
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The programming language that is used to create the application. Valid values:
	//
	// 	- **java*	- :Java.
	//
	// 	- **php**: PHP.
	//
	// 	- **other**: other programming languages, such as Python, C++, Go, .NET, and Node.js
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// The number of application instances that are running.
	//
	// example:
	//
	// 1
	RunningInstances *int32 `json:"RunningInstances,omitempty" xml:"RunningInstances,omitempty"`
	// Indicates whether the auto scaling policy is enabled. Valid values:
	//
	// 	- **true**: The auto scaling policy is enabled.
	//
	// 	- **false**: The auto scaling policy is disabled.
	//
	// example:
	//
	// true
	ScaleRuleEnabled *string `json:"ScaleRuleEnabled,omitempty" xml:"ScaleRuleEnabled,omitempty"`
	// The type of the auto scaling policy. Valid values:
	//
	// 	- **timing**: a scheduled auto scaling policy.
	//
	// 	- **metric**: a metric-based auto scaling policy.
	//
	// 	- **mix**: a hybrid auto scaling policy.
	//
	// example:
	//
	// timing
	ScaleRuleType *string `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
}

func (s GetApplicationResponseBodyApplication) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyApplication) GetAppDescription() *string {
	return s.AppDescription
}

func (s *GetApplicationResponseBodyApplication) GetAppId() *string {
	return s.AppId
}

func (s *GetApplicationResponseBodyApplication) GetAppName() *string {
	return s.AppName
}

func (s *GetApplicationResponseBodyApplication) GetBaseAppId() *string {
	return s.BaseAppId
}

func (s *GetApplicationResponseBodyApplication) GetCpu() *int32 {
	return s.Cpu
}

func (s *GetApplicationResponseBodyApplication) GetInstances() *int32 {
	return s.Instances
}

func (s *GetApplicationResponseBodyApplication) GetIsStateful() *bool {
	return s.IsStateful
}

func (s *GetApplicationResponseBodyApplication) GetMem() *int32 {
	return s.Mem
}

func (s *GetApplicationResponseBodyApplication) GetMseEnabled() *bool {
	return s.MseEnabled
}

func (s *GetApplicationResponseBodyApplication) GetMseNamespaceId() *string {
	return s.MseNamespaceId
}

func (s *GetApplicationResponseBodyApplication) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetApplicationResponseBodyApplication) GetProgrammingLanguage() *string {
	return s.ProgrammingLanguage
}

func (s *GetApplicationResponseBodyApplication) GetRunningInstances() *int32 {
	return s.RunningInstances
}

func (s *GetApplicationResponseBodyApplication) GetScaleRuleEnabled() *string {
	return s.ScaleRuleEnabled
}

func (s *GetApplicationResponseBodyApplication) GetScaleRuleType() *string {
	return s.ScaleRuleType
}

func (s *GetApplicationResponseBodyApplication) SetAppDescription(v string) *GetApplicationResponseBodyApplication {
	s.AppDescription = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetAppId(v string) *GetApplicationResponseBodyApplication {
	s.AppId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetAppName(v string) *GetApplicationResponseBodyApplication {
	s.AppName = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetBaseAppId(v string) *GetApplicationResponseBodyApplication {
	s.BaseAppId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetCpu(v int32) *GetApplicationResponseBodyApplication {
	s.Cpu = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetInstances(v int32) *GetApplicationResponseBodyApplication {
	s.Instances = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetIsStateful(v bool) *GetApplicationResponseBodyApplication {
	s.IsStateful = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetMem(v int32) *GetApplicationResponseBodyApplication {
	s.Mem = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetMseEnabled(v bool) *GetApplicationResponseBodyApplication {
	s.MseEnabled = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetMseNamespaceId(v string) *GetApplicationResponseBodyApplication {
	s.MseNamespaceId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetNamespaceId(v string) *GetApplicationResponseBodyApplication {
	s.NamespaceId = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetProgrammingLanguage(v string) *GetApplicationResponseBodyApplication {
	s.ProgrammingLanguage = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetRunningInstances(v int32) *GetApplicationResponseBodyApplication {
	s.RunningInstances = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetScaleRuleEnabled(v string) *GetApplicationResponseBodyApplication {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) SetScaleRuleType(v string) *GetApplicationResponseBodyApplication {
	s.ScaleRuleType = &v
	return s
}

func (s *GetApplicationResponseBodyApplication) Validate() error {
	return dara.Validate(s)
}
