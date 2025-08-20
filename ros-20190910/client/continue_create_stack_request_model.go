// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueCreateStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *ContinueCreateStackRequest
	GetDryRun() *bool
	SetMode(v string) *ContinueCreateStackRequest
	GetMode() *string
	SetParallelism(v int64) *ContinueCreateStackRequest
	GetParallelism() *int64
	SetParameters(v []*ContinueCreateStackRequestParameters) *ContinueCreateStackRequest
	GetParameters() []*ContinueCreateStackRequestParameters
	SetRamRoleName(v string) *ContinueCreateStackRequest
	GetRamRoleName() *string
	SetRecreatingOptions(v []*string) *ContinueCreateStackRequest
	GetRecreatingOptions() []*string
	SetRecreatingResources(v []*string) *ContinueCreateStackRequest
	GetRecreatingResources() []*string
	SetRegionId(v string) *ContinueCreateStackRequest
	GetRegionId() *string
	SetStackId(v string) *ContinueCreateStackRequest
	GetStackId() *string
	SetTemplateBody(v string) *ContinueCreateStackRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *ContinueCreateStackRequest
	GetTemplateId() *string
	SetTemplateURL(v string) *ContinueCreateStackRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *ContinueCreateStackRequest
	GetTemplateVersion() *string
}

type ContinueCreateStackRequest struct {
	// Specifies whether only to validate the stack in the request. Valid values:
	//
	// 	- true: only validates the stack.
	//
	// 	- false (default): validates and continues to create the stack.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The mode in which ROS continues to create the stack. Valid values:
	//
	// 	- Recreate (default)
	//
	//     If you set this parameter to Recreate, ROS continues to create only the following types of resources:
	//
	//     	- Resources that fail to be created
	//
	//     	- Resources that you specify for RecreatingResources.N
	//
	//     	- Dependencies of the resources that you specify for RecreatingResources.N
	//
	//     	- Resources that you have not created
	//
	// > RecreatingResources.N, TemplateBody, TemplateURL, and Parameters take effect only when Mode is set to Recreate.
	//
	// 	- Ignore
	//
	//     	- ROS ignores and discards resources that fail to be created and you have not created, and considers that the stack is successfully created.
	//
	//     	- The body of the template that you use to create the stack is changed.
	//
	// > This mode is available only for ROS stacks.
	//
	// example:
	//
	// Recreate
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The maximum number of concurrent operations that can be performed on resources.
	//
	// By default, this parameter is empty. You can set this parameter to an integer that is greater than or equal to 0.
	//
	// > - If you set this parameter to an integer that is greater than 0, the integer is used.
	//
	// > - If you set this parameter to 0, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	//
	// > - If you leave this parameter empty, the value that you specified for this parameter in the previous request is used. If you left this parameter empty in the previous request, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	//
	// > - If you set this parameter to a specific value, ROS associates the value with the stack. The value affects subsequent operations on the stack.
	//
	// example:
	//
	// 1
	Parallelism *int64 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The template parameters that you want to use to override specific parameters.
	Parameters []*ContinueCreateStackRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The name of the RAM role. Resource Orchestration Service (ROS) assumes the RAM role to create the stack and uses the credentials of the role to call the APIs of Alibaba Cloud services.\\
	//
	// ROS assumes the RAM role to perform operations on the stack. If you have permissions to perform operations on the stack but do not have permissions to use the RAM role, ROS still assumes the RAM role. You must make sure that the least privileges are granted to the RAM role.\\
	//
	// If you do not specify this parameter, ROS assumes the existing role that is associated with the stack. If no roles are available, ROS uses a temporary credential that is generated from the credentials of your account.\\
	//
	// The name of the RAM role can be up to 64 bytes in length.
	//
	// example:
	//
	// test-role
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The options that ROS adopts when ROS continues to create the stack.
	RecreatingOptions []*string `json:"RecreatingOptions,omitempty" xml:"RecreatingOptions,omitempty" type:"Repeated"`
	// The resources that ROS continues to create after the resources failed to be created. You can add new resources to the resources that ROS continues to create. ROS continues to create all dependencies of the new resources.
	//
	// > This parameter is available only for ROS stacks.
	//
	// example:
	//
	// WebServer
	RecreatingResources []*string `json:"RecreatingResources,omitempty" xml:"RecreatingResources,omitempty" type:"Repeated"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The stack ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.\\
	//
	// If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.
	//
	// A ROS template is subject to the following limits:
	//
	// 	- You can modify only the following sections in the template: Description, Metadata, Resources, and Outputs.
	//
	// 	- You cannot add sections to or remove sections from the template.
	//
	// 	- The Resources section is subject to the following limits:
	//
	//     	- You cannot delete the resources or change the template body for the resources that you do not want to continue to create.
	//
	//     	- You can delete the resources or change the template body for the resources that you want to continue to create.
	//
	//     	- You can add resources to this section.
	//
	//
	//
	// > - This parameter takes effect only when Mode is set to Recreate.
	//
	// > - You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId. If you do not specify the parameters, the existing template is used.
	//
	// example:
	//
	// {"ROSTemplateFormatVersion": "2015-09-01"}
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The template ID. This parameter applies to shared and private templates.
	//
	// > - This parameter takes effect when `Mode` is set to `Recreate`. When you specify TemplateId of a template, the template is subject to the limits that are described for `TemplateBody` in this topic.
	//
	// > - You can specify only one of the following parameters: `TemplateBody`, `TemplateURL`, and `TemplateId`. If you do not specify the parameters, the existing template is used.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length.
	//
	// If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// > - This parameter takes effect only when Mode is set to Recreate. When you specify TemplateURL of a template, the template is subject to the limits that are described for TemplateBody in this topic.
	//
	// > - You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId. If you do not specify the parameters, the existing template is used.
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
}

func (s ContinueCreateStackRequest) String() string {
	return dara.Prettify(s)
}

func (s ContinueCreateStackRequest) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ContinueCreateStackRequest) GetMode() *string {
	return s.Mode
}

func (s *ContinueCreateStackRequest) GetParallelism() *int64 {
	return s.Parallelism
}

func (s *ContinueCreateStackRequest) GetParameters() []*ContinueCreateStackRequestParameters {
	return s.Parameters
}

func (s *ContinueCreateStackRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *ContinueCreateStackRequest) GetRecreatingOptions() []*string {
	return s.RecreatingOptions
}

func (s *ContinueCreateStackRequest) GetRecreatingResources() []*string {
	return s.RecreatingResources
}

func (s *ContinueCreateStackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ContinueCreateStackRequest) GetStackId() *string {
	return s.StackId
}

func (s *ContinueCreateStackRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *ContinueCreateStackRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ContinueCreateStackRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *ContinueCreateStackRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *ContinueCreateStackRequest) SetDryRun(v bool) *ContinueCreateStackRequest {
	s.DryRun = &v
	return s
}

func (s *ContinueCreateStackRequest) SetMode(v string) *ContinueCreateStackRequest {
	s.Mode = &v
	return s
}

func (s *ContinueCreateStackRequest) SetParallelism(v int64) *ContinueCreateStackRequest {
	s.Parallelism = &v
	return s
}

func (s *ContinueCreateStackRequest) SetParameters(v []*ContinueCreateStackRequestParameters) *ContinueCreateStackRequest {
	s.Parameters = v
	return s
}

func (s *ContinueCreateStackRequest) SetRamRoleName(v string) *ContinueCreateStackRequest {
	s.RamRoleName = &v
	return s
}

func (s *ContinueCreateStackRequest) SetRecreatingOptions(v []*string) *ContinueCreateStackRequest {
	s.RecreatingOptions = v
	return s
}

func (s *ContinueCreateStackRequest) SetRecreatingResources(v []*string) *ContinueCreateStackRequest {
	s.RecreatingResources = v
	return s
}

func (s *ContinueCreateStackRequest) SetRegionId(v string) *ContinueCreateStackRequest {
	s.RegionId = &v
	return s
}

func (s *ContinueCreateStackRequest) SetStackId(v string) *ContinueCreateStackRequest {
	s.StackId = &v
	return s
}

func (s *ContinueCreateStackRequest) SetTemplateBody(v string) *ContinueCreateStackRequest {
	s.TemplateBody = &v
	return s
}

func (s *ContinueCreateStackRequest) SetTemplateId(v string) *ContinueCreateStackRequest {
	s.TemplateId = &v
	return s
}

func (s *ContinueCreateStackRequest) SetTemplateURL(v string) *ContinueCreateStackRequest {
	s.TemplateURL = &v
	return s
}

func (s *ContinueCreateStackRequest) SetTemplateVersion(v string) *ContinueCreateStackRequest {
	s.TemplateVersion = &v
	return s
}

func (s *ContinueCreateStackRequest) Validate() error {
	return dara.Validate(s)
}

type ContinueCreateStackRequestParameters struct {
	// The name of template parameter N that you want to use to override a specific parameter. If you do not specify the name and value of a template parameter, ROS uses the name and value specified in the previous operation that was performed to create the stack. Maximum value of N: 200.
	//
	// > This parameter takes effect only when Mode is set to Recreate.
	//
	// This parameter is required.
	//
	// example:
	//
	// Amount
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of template parameter N that you want to use to override a specific parameter. Maximum value of N: 200.
	//
	// For ROS stacks, the template parameters that you use to override specific parameters are subject to the following limits:
	//
	// 	- You cannot change the condition values in the Conditions section of a template from true to false or from false to true.
	//
	// 	- The template parameters can be referenced only by resources that ROS continues to create.
	//
	// > This parameter takes effect only when Mode is set to Recreate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s ContinueCreateStackRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s ContinueCreateStackRequestParameters) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *ContinueCreateStackRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *ContinueCreateStackRequestParameters) SetParameterKey(v string) *ContinueCreateStackRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *ContinueCreateStackRequestParameters) SetParameterValue(v string) *ContinueCreateStackRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *ContinueCreateStackRequestParameters) Validate() error {
	return dara.Validate(s)
}
