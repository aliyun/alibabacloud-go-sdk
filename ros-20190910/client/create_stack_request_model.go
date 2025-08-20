// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateStackRequest
	GetClientToken() *string
	SetCreateOption(v string) *CreateStackRequest
	GetCreateOption() *string
	SetCreateOptions(v []*string) *CreateStackRequest
	GetCreateOptions() []*string
	SetDeletionProtection(v string) *CreateStackRequest
	GetDeletionProtection() *string
	SetDisableRollback(v bool) *CreateStackRequest
	GetDisableRollback() *bool
	SetNotificationURLs(v []*string) *CreateStackRequest
	GetNotificationURLs() []*string
	SetParallelism(v int64) *CreateStackRequest
	GetParallelism() *int64
	SetParameters(v []*CreateStackRequestParameters) *CreateStackRequest
	GetParameters() []*CreateStackRequestParameters
	SetRamRoleName(v string) *CreateStackRequest
	GetRamRoleName() *string
	SetRegionId(v string) *CreateStackRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateStackRequest
	GetResourceGroupId() *string
	SetStackName(v string) *CreateStackRequest
	GetStackName() *string
	SetStackPolicyBody(v string) *CreateStackRequest
	GetStackPolicyBody() *string
	SetStackPolicyURL(v string) *CreateStackRequest
	GetStackPolicyURL() *string
	SetTags(v []*CreateStackRequestTags) *CreateStackRequest
	GetTags() []*CreateStackRequestTags
	SetTemplateBody(v string) *CreateStackRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *CreateStackRequest
	GetTemplateId() *string
	SetTemplateScratchId(v string) *CreateStackRequest
	GetTemplateScratchId() *string
	SetTemplateScratchRegionId(v string) *CreateStackRequest
	GetTemplateScratchRegionId() *string
	SetTemplateURL(v string) *CreateStackRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *CreateStackRequest
	GetTemplateVersion() *string
	SetTimeoutInMinutes(v int64) *CreateStackRequest
	GetTimeoutInMinutes() *int64
}

type CreateStackRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The creation option for the stack. Valid values:
	//
	// 	- KeepStackOnCreationComplete (default): After the stack is created, the stack and its resources are retained. The quota for the maximum number of stacks that can be created in ROS is consumed.
	//
	// 	- AbandonStackOnCreationComplete: After the stack is created, the stack is deleted, but its resources are retained. The quota for the maximum number of stacks that can be created in ROS is not consumed. If the stack fails to be created, the stack is retained.
	//
	// 	- AbandonStackOnCreationRollbackComplete: When the resources of the stack are rolled back after the stack fails to be created, the stack is deleted. The quota for the maximum number of stacks that can be created in ROS is not consumed. In other rollback scenarios, the stack is retained.
	//
	// 	- ManuallyPay: When you create the stack, you must manually pay for the subscription resources that are used. The following resource types support manual payment: `ALIYUN::ECS::InstanceGroup`, `ALIYUN::RDS::DBInstance`, `ALIYUN::SLB::LoadBalancer`, `ALIYUN::VPC::EIP`, and `ALIYUN::VPC::VpnGateway`.
	//
	// >  You can specify only one of CreateOption and CreateOptions.
	//
	// example:
	//
	// KeepStackOnCreationComplete
	CreateOption *string `json:"CreateOption,omitempty" xml:"CreateOption,omitempty"`
	// The creation options for the stack.
	CreateOptions []*string `json:"CreateOptions,omitempty" xml:"CreateOptions,omitempty" type:"Repeated"`
	// Specifies whether to enable deletion protection for the stack. Valid values:
	//
	// 	- Enabled.
	//
	// 	- Disabled (default). If deletion protection is disabled, you can delete the stack by using the ROS console or by calling the DeleteStack operation.
	//
	// > The value of DeletionProtection that you specify for the root stack applies to its nested stacks.
	//
	// example:
	//
	// Enabled
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Specifies whether to disable rollback for the resources when the stack fails to be created.
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
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The callback URLs that are used to receive stack events. Valid values:
	//
	// 	- HTTP POST URL
	//
	// Each URL can be up to 1,024 bytes in length.
	//
	// 	- eventbridge
	//
	// When the status of a stack changes, ROS sends notifications to the EventBridge service. You can view the event information in the [EventBridge](https://eventbridge.console.aliyun.com) console.
	//
	// > This feature is supported in the China (Hangzhou), China (Shanghai), China (Beijing), China (Hong Kong), and China (Zhangjiakou) regions.
	//
	// Maximum value of N: 5. When the status of a stack changes, ROS sends a notification to the specified URL. When rollback is enabled for the stack, notifications are sent if the stack is in the CREATE_ROLLBACK or ROLLBACK state, but are not sent if the stack is in the CREATE_FAILED, UPDATE_FAILED, or IN_PROGRESS state.\\
	//
	// ROS sends notifications regardless of whether you specify the Outputs section. The following sample code provides an example on the content of a notification:
	//
	//     {
	//
	//        "Outputs": [
	//
	//            {
	//
	//                "Description": "No description given",
	//
	//                "OutputKey": "InstanceId",
	//
	//                "OutputValue": "i-xxx"
	//
	//            }
	//
	//        ],
	//
	//        "StackId": "80bd6b6c-e888-4573-ae3b-93d29113****",
	//
	//        "StackName": "test-notification-url",
	//
	//        "Status": "CREATE_COMPLETE"
	//
	//     }
	//
	// example:
	//
	// http://my-site.com/ros-event
	NotificationURLs []*string `json:"NotificationURLs,omitempty" xml:"NotificationURLs,omitempty" type:"Repeated"`
	// The maximum number of concurrent operations that can be performed on resources.
	//
	// By default, this parameter is empty. You can set this parameter to an integer that is greater than or equal to 0.
	//
	//
	//
	// > -  If you set this parameter to an integer that is greater than 0, the integer is used. If you set this parameter to 0 or leave this parameter empty, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	//
	// > -  If you set this parameter to a specific value, ROS associates the value with the stack. The value affects subsequent operations on the stack, such as an update operation.
	//
	// example:
	//
	// 1
	Parallelism *int64 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The parameters that are defined in the template.
	Parameters []*CreateStackRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The name of the RAM role. ROS assumes the RAM role to create the stack and uses the credentials of the role to call the APIs of Alibaba Cloud services.\\
	//
	// ROS assumes the RAM role to perform operations on the stack. If you have permissions to perform operations on the stack but do not have permissions to use the RAM role, ROS still assumes the RAM role. You must make sure that the least privileges are granted to the RAM role.
	//
	// If you do not specify this parameter, ROS assumes the existing role that is associated with the stack. If no roles are available, ROS uses a temporary credential that is generated from the credentials of your account.
	//
	// The RAM role name can be up to 64 characters in length.
	//
	// example:
	//
	// test-role
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. If you leave this parameter empty, the stack is added to the default resource group.
	//
	// For more information about resource groups, see the "Resource group" section of the [What is Resource Management?](https://help.aliyun.com/document_detail/94475.html) topic.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the stack.\\
	//
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (_). It must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyStack
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The structure that contains the stack policy body. The policy body must be 1 to 16,384 bytes in length.
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
	StackPolicyURL *string `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
	// The tags that you want to add to the stack.
	Tags         []*CreateStackRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateBody *string                   `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The template ID. This parameter applies to shared templates and private templates.
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
	// 	- Default value: 60.
	//
	// 	- Unit: minutes.
	//
	// 	- Valid values: 10 to 1440.
	//
	// example:
	//
	// 10
	TimeoutInMinutes *int64 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s CreateStackRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStackRequest) GoString() string {
	return s.String()
}

func (s *CreateStackRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateStackRequest) GetCreateOption() *string {
	return s.CreateOption
}

func (s *CreateStackRequest) GetCreateOptions() []*string {
	return s.CreateOptions
}

func (s *CreateStackRequest) GetDeletionProtection() *string {
	return s.DeletionProtection
}

func (s *CreateStackRequest) GetDisableRollback() *bool {
	return s.DisableRollback
}

func (s *CreateStackRequest) GetNotificationURLs() []*string {
	return s.NotificationURLs
}

func (s *CreateStackRequest) GetParallelism() *int64 {
	return s.Parallelism
}

func (s *CreateStackRequest) GetParameters() []*CreateStackRequestParameters {
	return s.Parameters
}

func (s *CreateStackRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateStackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStackRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateStackRequest) GetStackName() *string {
	return s.StackName
}

func (s *CreateStackRequest) GetStackPolicyBody() *string {
	return s.StackPolicyBody
}

func (s *CreateStackRequest) GetStackPolicyURL() *string {
	return s.StackPolicyURL
}

func (s *CreateStackRequest) GetTags() []*CreateStackRequestTags {
	return s.Tags
}

func (s *CreateStackRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *CreateStackRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateStackRequest) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *CreateStackRequest) GetTemplateScratchRegionId() *string {
	return s.TemplateScratchRegionId
}

func (s *CreateStackRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *CreateStackRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *CreateStackRequest) GetTimeoutInMinutes() *int64 {
	return s.TimeoutInMinutes
}

func (s *CreateStackRequest) SetClientToken(v string) *CreateStackRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStackRequest) SetCreateOption(v string) *CreateStackRequest {
	s.CreateOption = &v
	return s
}

func (s *CreateStackRequest) SetCreateOptions(v []*string) *CreateStackRequest {
	s.CreateOptions = v
	return s
}

func (s *CreateStackRequest) SetDeletionProtection(v string) *CreateStackRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateStackRequest) SetDisableRollback(v bool) *CreateStackRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateStackRequest) SetNotificationURLs(v []*string) *CreateStackRequest {
	s.NotificationURLs = v
	return s
}

func (s *CreateStackRequest) SetParallelism(v int64) *CreateStackRequest {
	s.Parallelism = &v
	return s
}

func (s *CreateStackRequest) SetParameters(v []*CreateStackRequestParameters) *CreateStackRequest {
	s.Parameters = v
	return s
}

func (s *CreateStackRequest) SetRamRoleName(v string) *CreateStackRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateStackRequest) SetRegionId(v string) *CreateStackRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStackRequest) SetResourceGroupId(v string) *CreateStackRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateStackRequest) SetStackName(v string) *CreateStackRequest {
	s.StackName = &v
	return s
}

func (s *CreateStackRequest) SetStackPolicyBody(v string) *CreateStackRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *CreateStackRequest) SetStackPolicyURL(v string) *CreateStackRequest {
	s.StackPolicyURL = &v
	return s
}

func (s *CreateStackRequest) SetTags(v []*CreateStackRequestTags) *CreateStackRequest {
	s.Tags = v
	return s
}

func (s *CreateStackRequest) SetTemplateBody(v string) *CreateStackRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateStackRequest) SetTemplateId(v string) *CreateStackRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateStackRequest) SetTemplateScratchId(v string) *CreateStackRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *CreateStackRequest) SetTemplateScratchRegionId(v string) *CreateStackRequest {
	s.TemplateScratchRegionId = &v
	return s
}

func (s *CreateStackRequest) SetTemplateURL(v string) *CreateStackRequest {
	s.TemplateURL = &v
	return s
}

func (s *CreateStackRequest) SetTemplateVersion(v string) *CreateStackRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateStackRequest) SetTimeoutInMinutes(v int64) *CreateStackRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *CreateStackRequest) Validate() error {
	return dara.Validate(s)
}

type CreateStackRequestParameters struct {
	// The key of parameter N that is defined in the template. If you do not specify the name and value of a parameter, ROS uses the default name and value that are specified in the template.
	//
	// Maximum value of N: 200.\\
	//
	// The name must be 1 to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// InstanceId
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N that is defined in the template.
	//
	// Maximum value of N: 200.\\
	//
	// The value can be up to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-xxxxxx
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateStackRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateStackRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateStackRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *CreateStackRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *CreateStackRequestParameters) SetParameterKey(v string) *CreateStackRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackRequestParameters) SetParameterValue(v string) *CreateStackRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *CreateStackRequestParameters) Validate() error {
	return dara.Validate(s)
}

type CreateStackRequestTags struct {
	// The key of tag N that you want to add to the stack.
	//
	// Valid values of N: 1 to 20.
	//
	// > - The Tags parameter is optional. If you specify Tags, you must specify Tags.N.Key.
	//
	// > -  The tag of a stack is propagated to each resource that supports the tag feature in the stack. For more information, see [Propagate tags](https://help.aliyun.com/document_detail/201421.html).
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
	// > The tag of a stack is propagated to each resource that supports the tag feature in the stack. For more information, see [Propagate tags](https://help.aliyun.com/document_detail/201421.html).
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateStackRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateStackRequestTags) GoString() string {
	return s.String()
}

func (s *CreateStackRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateStackRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateStackRequestTags) SetKey(v string) *CreateStackRequestTags {
	s.Key = &v
	return s
}

func (s *CreateStackRequestTags) SetValue(v string) *CreateStackRequestTags {
	s.Value = &v
	return s
}

func (s *CreateStackRequestTags) Validate() error {
	return dara.Validate(s)
}
