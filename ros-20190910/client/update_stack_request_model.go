// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateStackRequest
	GetClientToken() *string
	SetDisableRollback(v bool) *UpdateStackRequest
	GetDisableRollback() *bool
	SetDryRun(v bool) *UpdateStackRequest
	GetDryRun() *bool
	SetDryRunOptions(v []*string) *UpdateStackRequest
	GetDryRunOptions() []*string
	SetParallelism(v int64) *UpdateStackRequest
	GetParallelism() *int64
	SetParameters(v []*UpdateStackRequestParameters) *UpdateStackRequest
	GetParameters() []*UpdateStackRequestParameters
	SetRamRoleName(v string) *UpdateStackRequest
	GetRamRoleName() *string
	SetRegionId(v string) *UpdateStackRequest
	GetRegionId() *string
	SetReplacementOption(v string) *UpdateStackRequest
	GetReplacementOption() *string
	SetResourceGroupId(v string) *UpdateStackRequest
	GetResourceGroupId() *string
	SetStackId(v string) *UpdateStackRequest
	GetStackId() *string
	SetStackPolicyBody(v string) *UpdateStackRequest
	GetStackPolicyBody() *string
	SetStackPolicyDuringUpdateBody(v string) *UpdateStackRequest
	GetStackPolicyDuringUpdateBody() *string
	SetStackPolicyDuringUpdateURL(v string) *UpdateStackRequest
	GetStackPolicyDuringUpdateURL() *string
	SetStackPolicyURL(v string) *UpdateStackRequest
	GetStackPolicyURL() *string
	SetTags(v []*UpdateStackRequestTags) *UpdateStackRequest
	GetTags() []*UpdateStackRequestTags
	SetTaintResources(v []*string) *UpdateStackRequest
	GetTaintResources() []*string
	SetTemplateBody(v string) *UpdateStackRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *UpdateStackRequest
	GetTemplateId() *string
	SetTemplateURL(v string) *UpdateStackRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *UpdateStackRequest
	GetTemplateVersion() *string
	SetTimeoutInMinutes(v int64) *UpdateStackRequest
	GetTimeoutInMinutes() *int64
	SetUsePreviousParameters(v bool) *UpdateStackRequest
	GetUsePreviousParameters() *bool
}

type UpdateStackRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests.
	//
	// The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Deprecated
	//
	// Specifies whether to roll back the resources in the stack when the stack fails to be updated.
	//
	// Default value: false. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// Specifies whether only to validate the stack in the request. Default value: false. Valid values:
	//
	// 	- true: only validates the stack.
	//
	// 	- false: validates and updates the stack.
	//
	// >  When no changes are made during the update, the following rules apply: If you set the DryRun parameter to false, the NotSupported error code is returned. If you set the DryRun parameter to true, no error is reported.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The dry run option in the list format. You can specify only one dry run option.
	//
	// > This parameter takes effect only when DryRun is set to true.
	DryRunOptions []*string `json:"DryRunOptions,omitempty" xml:"DryRunOptions,omitempty" type:"Repeated"`
	// The maximum number of concurrent operations that can be performed on resources.
	//
	// By default, this parameter is empty. You can set this parameter to an integer that is greater than or equal to 0.
	//
	// > - If you set this parameter to an integer that is greater than 0, the integer is used.
	//
	// > -  If you set this parameter to 0, no limit is imposed on Resource Orchestration Service (ROS) stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	//
	// > -  If you leave this parameter empty, the value that you specified for this parameter in the previous request is used. If you left this parameter empty in the previous request, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	//
	// > - If you set this parameter to a specific value, ROS associates the value with the stack. The value affects subsequent operations on the stack.
	//
	// example:
	//
	// 1
	Parallelism *int64 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The parameters.
	Parameters []*UpdateStackRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The name of the RAM role. Resource Orchestration Service (ROS) assumes the RAM role to create the stack and uses the credentials of the role to call the APIs of Alibaba Cloud services.
	//
	// ROS assumes the RAM role to perform operations on the stack. If you have permissions to perform operations on the stack but do not have permissions to use the RAM role, ROS still assumes the RAM role. You must make sure that the least privileges are granted to the RAM role.
	//
	// If you do not specify this parameter, ROS assumes the existing RAM role that is associated with the stack. If no RAM roles are available, ROS uses a temporary credential that is generated from the credentials of your account.
	//
	// The name of the RAM role can be up to 64 bytes in length.
	//
	// example:
	//
	// test-role
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the region in which the stack is deployed. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable the replacement update feature. If you cannot change resource properties, you can enable the replacement update feature to replace the resource properties. If the replacement update feature is used, the existing resource is deleted and a new resource is created. The physical ID of the new resource is different from the physical ID of the deleted resource.
	//
	// Default value: Disabled. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// >  Changes have higher priorities than replacement updates.
	//
	// example:
	//
	// Disabled
	ReplacementOption *string `json:"ReplacementOption,omitempty" xml:"ReplacementOption,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the stack.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The structure that contains the stack policy body. The policy body must be 1 to 16,384 bytes in length.
	//
	// >  You can specify only one of the StackPolicyBody and StackPolicyURL parameters.
	//
	// example:
	//
	// {"Statement": [{"Action": "Update:*", "Resource": "*", "Effect": "Allow", "Principal": "*"}]}
	StackPolicyBody *string `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	// The structure that contains the body of the temporary overriding stack policy. The policy body must be 1 to 16,384 bytes in length.
	//
	// If you want to update protected resources, you must specify a temporary overriding stack policy during the update. If you do not specify a temporary overriding stack policy, the existing policy that is associated with the stack is used.
	//
	// This parameter takes effect only when the ChangeSetType parameter is set to UPDATE. You can specify only one of the following parameters:
	//
	// 	- StackPolicyBody
	//
	// 	- StackPolicyURL
	//
	// 	- StackPolicyDuringUpdateBody
	//
	// 	- StackPolicyDuringUpdateURL
	//
	// example:
	//
	// {"Statement": [{"Effect": "Allow", "Action": "Update:*", "Principal": "*", "Resource": "*"}]}
	StackPolicyDuringUpdateBody *string `json:"StackPolicyDuringUpdateBody,omitempty" xml:"StackPolicyDuringUpdateBody,omitempty"`
	// The URL of the file that contains the temporary overriding stack policy. The URL must point to a policy that is located on an HTTP or HTTPS web server or in an OSS bucket, such as oss://ros/stack-policy/demo or oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy file can be up to 16,384 bytes in length.
	//
	// >  If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// The URL can be up to 1,350 bytes in length.
	//
	// If you want to update protected resources, you must specify a temporary overriding stack policy during the update. If you do not specify a temporary overriding stack policy, the existing policy that is associated with the stack is used. This parameter takes effect only when the ChangeSetType parameter is set to UPDATE. You can specify only one of the following parameters:
	//
	// 	- StackPolicyBody
	//
	// 	- StackPolicyURL
	//
	// 	- StackPolicyDuringUpdateBody
	//
	// 	- StackPolicyDuringUpdateURL
	//
	// example:
	//
	// oss://ros-stack-policy/demo
	StackPolicyDuringUpdateURL *string `json:"StackPolicyDuringUpdateURL,omitempty" xml:"StackPolicyDuringUpdateURL,omitempty"`
	// The URL of the file that contains the stack policy. The URL must point to a policy that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo or oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy file can be up to 16,384 bytes in length. If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// >  You can specify only one of the StackPolicyBody and StackPolicyURL parameters.
	//
	// The URL can be up to 1,350 bytes in length.
	//
	// example:
	//
	// oss://ros-stack-policy/demo
	StackPolicyURL *string `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
	// The value of tag N that you want to add to the template.
	Tags           []*UpdateStackRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TaintResources []*string                 `json:"TaintResources,omitempty" xml:"TaintResources,omitempty" type:"Repeated"`
	TemplateBody   *string                   `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared templates and private templates.
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an OSS bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body must be 1 to 524,288 bytes in length. If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// example:
	//
	// oss://ros-template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. This parameter takes effect only when the TemplateId parameter is specified.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The timeout period for the update operation on the stack.
	//
	// 	- Default value: 60.
	//
	// 	- Unit: minutes.
	//
	// example:
	//
	// 10
	TimeoutInMinutes *int64 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	// Specifies whether to use the values specified in the previous request for the parameters that you do not specify in the current request.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	UsePreviousParameters *bool `json:"UsePreviousParameters,omitempty" xml:"UsePreviousParameters,omitempty"`
}

func (s UpdateStackRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateStackRequest) GetDisableRollback() *bool {
	return s.DisableRollback
}

func (s *UpdateStackRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateStackRequest) GetDryRunOptions() []*string {
	return s.DryRunOptions
}

func (s *UpdateStackRequest) GetParallelism() *int64 {
	return s.Parallelism
}

func (s *UpdateStackRequest) GetParameters() []*UpdateStackRequestParameters {
	return s.Parameters
}

func (s *UpdateStackRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *UpdateStackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateStackRequest) GetReplacementOption() *string {
	return s.ReplacementOption
}

func (s *UpdateStackRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateStackRequest) GetStackId() *string {
	return s.StackId
}

func (s *UpdateStackRequest) GetStackPolicyBody() *string {
	return s.StackPolicyBody
}

func (s *UpdateStackRequest) GetStackPolicyDuringUpdateBody() *string {
	return s.StackPolicyDuringUpdateBody
}

func (s *UpdateStackRequest) GetStackPolicyDuringUpdateURL() *string {
	return s.StackPolicyDuringUpdateURL
}

func (s *UpdateStackRequest) GetStackPolicyURL() *string {
	return s.StackPolicyURL
}

func (s *UpdateStackRequest) GetTags() []*UpdateStackRequestTags {
	return s.Tags
}

func (s *UpdateStackRequest) GetTaintResources() []*string {
	return s.TaintResources
}

func (s *UpdateStackRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *UpdateStackRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateStackRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *UpdateStackRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *UpdateStackRequest) GetTimeoutInMinutes() *int64 {
	return s.TimeoutInMinutes
}

func (s *UpdateStackRequest) GetUsePreviousParameters() *bool {
	return s.UsePreviousParameters
}

func (s *UpdateStackRequest) SetClientToken(v string) *UpdateStackRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackRequest) SetDisableRollback(v bool) *UpdateStackRequest {
	s.DisableRollback = &v
	return s
}

func (s *UpdateStackRequest) SetDryRun(v bool) *UpdateStackRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateStackRequest) SetDryRunOptions(v []*string) *UpdateStackRequest {
	s.DryRunOptions = v
	return s
}

func (s *UpdateStackRequest) SetParallelism(v int64) *UpdateStackRequest {
	s.Parallelism = &v
	return s
}

func (s *UpdateStackRequest) SetParameters(v []*UpdateStackRequestParameters) *UpdateStackRequest {
	s.Parameters = v
	return s
}

func (s *UpdateStackRequest) SetRamRoleName(v string) *UpdateStackRequest {
	s.RamRoleName = &v
	return s
}

func (s *UpdateStackRequest) SetRegionId(v string) *UpdateStackRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackRequest) SetReplacementOption(v string) *UpdateStackRequest {
	s.ReplacementOption = &v
	return s
}

func (s *UpdateStackRequest) SetResourceGroupId(v string) *UpdateStackRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateStackRequest) SetStackId(v string) *UpdateStackRequest {
	s.StackId = &v
	return s
}

func (s *UpdateStackRequest) SetStackPolicyBody(v string) *UpdateStackRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *UpdateStackRequest) SetStackPolicyDuringUpdateBody(v string) *UpdateStackRequest {
	s.StackPolicyDuringUpdateBody = &v
	return s
}

func (s *UpdateStackRequest) SetStackPolicyDuringUpdateURL(v string) *UpdateStackRequest {
	s.StackPolicyDuringUpdateURL = &v
	return s
}

func (s *UpdateStackRequest) SetStackPolicyURL(v string) *UpdateStackRequest {
	s.StackPolicyURL = &v
	return s
}

func (s *UpdateStackRequest) SetTags(v []*UpdateStackRequestTags) *UpdateStackRequest {
	s.Tags = v
	return s
}

func (s *UpdateStackRequest) SetTaintResources(v []*string) *UpdateStackRequest {
	s.TaintResources = v
	return s
}

func (s *UpdateStackRequest) SetTemplateBody(v string) *UpdateStackRequest {
	s.TemplateBody = &v
	return s
}

func (s *UpdateStackRequest) SetTemplateId(v string) *UpdateStackRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateStackRequest) SetTemplateURL(v string) *UpdateStackRequest {
	s.TemplateURL = &v
	return s
}

func (s *UpdateStackRequest) SetTemplateVersion(v string) *UpdateStackRequest {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateStackRequest) SetTimeoutInMinutes(v int64) *UpdateStackRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *UpdateStackRequest) SetUsePreviousParameters(v bool) *UpdateStackRequest {
	s.UsePreviousParameters = &v
	return s
}

func (s *UpdateStackRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateStackRequestParameters struct {
	// The name of parameter N. If you do not specify the name and value of a parameter, ROS uses the default name and value in the template.
	//
	// Maximum value of N: 200.
	//
	// >  The Parameters parameter is optional. If you specify Parameters, you must specify both Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// Amount
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N. Maximum value of N: 200.
	//
	// >  The Parameters parameter is optional. If you specify Parameters, you must specify both Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateStackRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateStackRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *UpdateStackRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *UpdateStackRequestParameters) SetParameterKey(v string) *UpdateStackRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackRequestParameters) SetParameterValue(v string) *UpdateStackRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *UpdateStackRequestParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateStackRequestTags struct {
	// The key of tag N that you want to add to the stack.
	//
	// Valid values of N: 1 to 20.
	//
	// > - The Tags parameter is optional. If you specify Tags, you must specify Tags.N.Key.
	//
	// > - The tag of a stack is propagated to each resource that supports the tag feature in the stack. For more information, see [Propagate tags](https://help.aliyun.com/document_detail/201421.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that you want to add to the stack.
	//
	// Valid values of N: 1 to 20.
	//
	// >  The tag of a stack is propagated to each resource that supports the tag feature in the stack. For more information, see [Propagate tags](https://help.aliyun.com/document_detail/201421.html).
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateStackRequestTags) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateStackRequestTags) GetKey() *string {
	return s.Key
}

func (s *UpdateStackRequestTags) GetValue() *string {
	return s.Value
}

func (s *UpdateStackRequestTags) SetKey(v string) *UpdateStackRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateStackRequestTags) SetValue(v string) *UpdateStackRequestTags {
	s.Value = &v
	return s
}

func (s *UpdateStackRequestTags) Validate() error {
	return dara.Validate(s)
}
