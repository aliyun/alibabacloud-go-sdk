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
	// The response message.
	//
	// - If the request is successful, the value is **success**.
	//
	// - If the request fails, the value is a specific error code.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01CF26C7-00A3-4AA6-BA76-7E95F2A3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The trace ID used to query the details of the request.
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
	// The application description.
	//
	// example:
	//
	// Test
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
	// The base application ID.
	//
	// example:
	//
	// ee99cce6-1c8e-4bfa-96c3-3e2fa9******
	BaseAppId *string `json:"BaseAppId,omitempty" xml:"BaseAppId,omitempty"`
	// The CPU required for each instance, in millicores. This value cannot be 0. Valid values:
	//
	// - **500**
	//
	// - **1000**
	//
	// - **2000**
	//
	// - **4000**
	//
	// - **8000**
	//
	// - **12000**
	//
	// - **16000**
	//
	// - **32000**
	//
	// example:
	//
	// 2000
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The total number of application instances.
	//
	// example:
	//
	// 6
	Instances *int32 `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// Indicates whether the application is stateful.
	IsStateful *bool `json:"IsStateful,omitempty" xml:"IsStateful,omitempty"`
	// The memory required for each instance, in MB. This value cannot be 0. The memory specification is coupled with the CPU specification. The following configurations are supported:
	//
	// - **1024**: corresponds to 500 or 1,000 millicores of CPU.
	//
	// - **2048**: corresponds to 500, 1,000, or 2,000 millicores of CPU.
	//
	// - **4096**: corresponds to 1,000, 2,000, or 4,000 millicores of CPU.
	//
	// - **8192**: corresponds to 2,000, 4,000, or 8,000 millicores of CPU.
	//
	// - **12288**: corresponds to 12,000 millicores of CPU.
	//
	// - **16384**: corresponds to 4,000, 8,000, or 16,000 millicores of CPU.
	//
	// - **24576**: corresponds to 12,000 millicores of CPU.
	//
	// - **32768**: corresponds to 16,000 millicores of CPU.
	//
	// - **65536**: corresponds to 8,000, 16,000, or 32,000 millicores of CPU.
	//
	// - **131072**: corresponds to 32,000 millicores of CPU.
	//
	// example:
	//
	// 4096
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// Indicates whether WebAssemblyFilter is enabled. Valid values:
	//
	// - **true**: enabled.
	//
	// - **false**: disabled.
	//
	// example:
	//
	// true
	MseEnabled *bool `json:"MseEnabled,omitempty" xml:"MseEnabled,omitempty"`
	// The namespace ID of the MSE instance.
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
	// The programming language of the application. Valid values:
	//
	// - **java**: Java.
	//
	// - **php**: PHP.
	//
	// - **other**: other languages, such as Python, C++, Go, .NET, and Node.js.
	//
	// example:
	//
	// java
	ProgrammingLanguage *string `json:"ProgrammingLanguage,omitempty" xml:"ProgrammingLanguage,omitempty"`
	// The number of running instances.
	//
	// example:
	//
	// 6
	RunningInstances *int32 `json:"RunningInstances,omitempty" xml:"RunningInstances,omitempty"`
	// Indicates whether the auto scaling policy is enabled. Valid values:
	//
	// - **true**: The policy is enabled.
	//
	// - **false**: The policy is disabled.
	//
	// example:
	//
	// true
	ScaleRuleEnabled *string `json:"ScaleRuleEnabled,omitempty" xml:"ScaleRuleEnabled,omitempty"`
	// The type of the auto scaling policy. Valid values:
	//
	// - **timing**: scheduled auto scaling.
	//
	// - **metric**: metric-based auto scaling.
	//
	// - **mix**: hybrid auto scaling.
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
