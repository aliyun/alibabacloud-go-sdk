// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *PreviewStackRequest
	GetClientToken() *string
	SetDisableRollback(v bool) *PreviewStackRequest
	GetDisableRollback() *bool
	SetEnablePreConfig(v bool) *PreviewStackRequest
	GetEnablePreConfig() *bool
	SetParallelism(v int64) *PreviewStackRequest
	GetParallelism() *int64
	SetParameters(v []*PreviewStackRequestParameters) *PreviewStackRequest
	GetParameters() []*PreviewStackRequestParameters
	SetRegionId(v string) *PreviewStackRequest
	GetRegionId() *string
	SetStackId(v string) *PreviewStackRequest
	GetStackId() *string
	SetStackName(v string) *PreviewStackRequest
	GetStackName() *string
	SetStackPolicyBody(v string) *PreviewStackRequest
	GetStackPolicyBody() *string
	SetStackPolicyURL(v string) *PreviewStackRequest
	GetStackPolicyURL() *string
	SetTaintResources(v []*string) *PreviewStackRequest
	GetTaintResources() []*string
	SetTemplateBody(v string) *PreviewStackRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *PreviewStackRequest
	GetTemplateId() *string
	SetTemplateScratchId(v string) *PreviewStackRequest
	GetTemplateScratchId() *string
	SetTemplateScratchRegionId(v string) *PreviewStackRequest
	GetTemplateScratchRegionId() *string
	SetTemplateURL(v string) *PreviewStackRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *PreviewStackRequest
	GetTemplateVersion() *string
	SetTimeoutInMinutes(v int64) *PreviewStackRequest
	GetTimeoutInMinutes() *int64
	SetUsePreviousParameters(v bool) *PreviewStackRequest
	GetUsePreviousParameters() *bool
}

type PreviewStackRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\\
	//
	// The token can be up to 64 characters in length, and can contain letters, digits, underscores (_), and hyphens (-).\\
	//
	// For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to disable rollback for the resources when the stack fails to be created. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// Specifies whether to query the parameters that you want to use in compliance precheck.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	EnablePreConfig *bool `json:"EnablePreConfig,omitempty" xml:"EnablePreConfig,omitempty"`
	// The maximum number of concurrent operations that can be performed on resources. This parameter takes effect only for Terraform stacks.
	//
	// By default, this parameter is empty. You can set this parameter to an integer that is greater than or equal to 0.
	//
	// > If you set this parameter to an integer greater than 0, the integer is used. If you set this parameter to 0 or leave this parameter empty, the default value of Terraform is used. In most cases, the default value of Terraform is 10.
	//
	// example:
	//
	// 1
	Parallelism *int64 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The parameters of the stack.
	Parameters []*PreviewStackRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The stack ID. You can use this parameter to preview a stack that you want to update.
	//
	//
	//
	// > -  You must and can specify only one of StackName and StackId.
	//
	// > - In the scenario in which you preview a stack that you want to create or update, you cannot preview the resources in its nested stacks.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The stack name. You can use this parameter to preview the stack that you want to create. The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
	//
	// > You must and can specify only one of StackName and StackId.
	//
	// example:
	//
	// MyStack
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
	//
	// > You can specify only one of StackPolicyBody and StackPolicyURL.
	//
	// example:
	//
	// {"Statement": [{"Action": "Update:*", "Resource": "*", "Effect": "Allow", "Principal": "*"}]}
	StackPolicyBody *string `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	// The URL of the file that contains the stack policy. The URL must point to a policy that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo or oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy file can be up to 16,384 bytes in length. If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// > You can specify only one of StackPolicyBody and StackPolicyURL.
	//
	// The URL can be up to 1,350 bytes in length.
	//
	// example:
	//
	// oss://ros-stack-policy/demo
	StackPolicyURL *string   `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
	TaintResources []*string `json:"TaintResources,omitempty" xml:"TaintResources,omitempty" type:"Repeated"`
	TemplateBody   *string   `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The template ID. This parameter applies to shared and private templates.
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The scenario ID.
	//
	// For more information about how to query the scenario ID, see [ListTemplateScratches](https://help.aliyun.com/document_detail/363050.html).
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	//
	// example:
	//
	// ts-aa9c62feab844a6b****
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The region ID of the scenario. The default value is the same as the value of RegionId.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	TemplateScratchRegionId *string `json:"TemplateScratchRegionId,omitempty" xml:"TemplateScratchRegionId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an OSS bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length. If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	//
	// example:
	//
	// oss://ros-template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. This parameter takes effect only when TemplateId is specified.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The timeout period for creating the stack.
	//
	// Unit: minutes.
	//
	// Default value: 60.
	//
	// example:
	//
	// 60
	TimeoutInMinutes      *int64 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	UsePreviousParameters *bool  `json:"UsePreviousParameters,omitempty" xml:"UsePreviousParameters,omitempty"`
}

func (s PreviewStackRequest) String() string {
	return dara.Prettify(s)
}

func (s PreviewStackRequest) GoString() string {
	return s.String()
}

func (s *PreviewStackRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PreviewStackRequest) GetDisableRollback() *bool {
	return s.DisableRollback
}

func (s *PreviewStackRequest) GetEnablePreConfig() *bool {
	return s.EnablePreConfig
}

func (s *PreviewStackRequest) GetParallelism() *int64 {
	return s.Parallelism
}

func (s *PreviewStackRequest) GetParameters() []*PreviewStackRequestParameters {
	return s.Parameters
}

func (s *PreviewStackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PreviewStackRequest) GetStackId() *string {
	return s.StackId
}

func (s *PreviewStackRequest) GetStackName() *string {
	return s.StackName
}

func (s *PreviewStackRequest) GetStackPolicyBody() *string {
	return s.StackPolicyBody
}

func (s *PreviewStackRequest) GetStackPolicyURL() *string {
	return s.StackPolicyURL
}

func (s *PreviewStackRequest) GetTaintResources() []*string {
	return s.TaintResources
}

func (s *PreviewStackRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *PreviewStackRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *PreviewStackRequest) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *PreviewStackRequest) GetTemplateScratchRegionId() *string {
	return s.TemplateScratchRegionId
}

func (s *PreviewStackRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *PreviewStackRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *PreviewStackRequest) GetTimeoutInMinutes() *int64 {
	return s.TimeoutInMinutes
}

func (s *PreviewStackRequest) GetUsePreviousParameters() *bool {
	return s.UsePreviousParameters
}

func (s *PreviewStackRequest) SetClientToken(v string) *PreviewStackRequest {
	s.ClientToken = &v
	return s
}

func (s *PreviewStackRequest) SetDisableRollback(v bool) *PreviewStackRequest {
	s.DisableRollback = &v
	return s
}

func (s *PreviewStackRequest) SetEnablePreConfig(v bool) *PreviewStackRequest {
	s.EnablePreConfig = &v
	return s
}

func (s *PreviewStackRequest) SetParallelism(v int64) *PreviewStackRequest {
	s.Parallelism = &v
	return s
}

func (s *PreviewStackRequest) SetParameters(v []*PreviewStackRequestParameters) *PreviewStackRequest {
	s.Parameters = v
	return s
}

func (s *PreviewStackRequest) SetRegionId(v string) *PreviewStackRequest {
	s.RegionId = &v
	return s
}

func (s *PreviewStackRequest) SetStackId(v string) *PreviewStackRequest {
	s.StackId = &v
	return s
}

func (s *PreviewStackRequest) SetStackName(v string) *PreviewStackRequest {
	s.StackName = &v
	return s
}

func (s *PreviewStackRequest) SetStackPolicyBody(v string) *PreviewStackRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *PreviewStackRequest) SetStackPolicyURL(v string) *PreviewStackRequest {
	s.StackPolicyURL = &v
	return s
}

func (s *PreviewStackRequest) SetTaintResources(v []*string) *PreviewStackRequest {
	s.TaintResources = v
	return s
}

func (s *PreviewStackRequest) SetTemplateBody(v string) *PreviewStackRequest {
	s.TemplateBody = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateId(v string) *PreviewStackRequest {
	s.TemplateId = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateScratchId(v string) *PreviewStackRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateScratchRegionId(v string) *PreviewStackRequest {
	s.TemplateScratchRegionId = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateURL(v string) *PreviewStackRequest {
	s.TemplateURL = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateVersion(v string) *PreviewStackRequest {
	s.TemplateVersion = &v
	return s
}

func (s *PreviewStackRequest) SetTimeoutInMinutes(v int64) *PreviewStackRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *PreviewStackRequest) SetUsePreviousParameters(v bool) *PreviewStackRequest {
	s.UsePreviousParameters = &v
	return s
}

func (s *PreviewStackRequest) Validate() error {
	return dara.Validate(s)
}

type PreviewStackRequestParameters struct {
	// The name of the parameter N. If you do not specify the name and value of a parameter, Resource Orchestration Service (ROS) uses the default name and value that are specified in the template. Maximum value of N: 200.
	//
	// > If you specify Parameters, you must specify Parameters.N.ParameterKey.
	//
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN::AccountId
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N. Maximum value of N: 200.
	//
	// > If you specify Parameters, you must specify Parameters.N.ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// 151266687691****
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s PreviewStackRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s PreviewStackRequestParameters) GoString() string {
	return s.String()
}

func (s *PreviewStackRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *PreviewStackRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *PreviewStackRequestParameters) SetParameterKey(v string) *PreviewStackRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *PreviewStackRequestParameters) SetParameterValue(v string) *PreviewStackRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *PreviewStackRequestParameters) Validate() error {
	return dara.Validate(s)
}
