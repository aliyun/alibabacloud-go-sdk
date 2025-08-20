// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChangeSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeSetName(v string) *CreateChangeSetRequest
	GetChangeSetName() *string
	SetChangeSetType(v string) *CreateChangeSetRequest
	GetChangeSetType() *string
	SetClientToken(v string) *CreateChangeSetRequest
	GetClientToken() *string
	SetDescription(v string) *CreateChangeSetRequest
	GetDescription() *string
	SetDisableRollback(v bool) *CreateChangeSetRequest
	GetDisableRollback() *bool
	SetNotificationURLs(v []*string) *CreateChangeSetRequest
	GetNotificationURLs() []*string
	SetParallelism(v int64) *CreateChangeSetRequest
	GetParallelism() *int64
	SetParameters(v []*CreateChangeSetRequestParameters) *CreateChangeSetRequest
	GetParameters() []*CreateChangeSetRequestParameters
	SetRamRoleName(v string) *CreateChangeSetRequest
	GetRamRoleName() *string
	SetRegionId(v string) *CreateChangeSetRequest
	GetRegionId() *string
	SetReplacementOption(v string) *CreateChangeSetRequest
	GetReplacementOption() *string
	SetResourceGroupId(v string) *CreateChangeSetRequest
	GetResourceGroupId() *string
	SetResourcesToImport(v []*CreateChangeSetRequestResourcesToImport) *CreateChangeSetRequest
	GetResourcesToImport() []*CreateChangeSetRequestResourcesToImport
	SetStackId(v string) *CreateChangeSetRequest
	GetStackId() *string
	SetStackName(v string) *CreateChangeSetRequest
	GetStackName() *string
	SetStackPolicyBody(v string) *CreateChangeSetRequest
	GetStackPolicyBody() *string
	SetStackPolicyDuringUpdateBody(v string) *CreateChangeSetRequest
	GetStackPolicyDuringUpdateBody() *string
	SetStackPolicyDuringUpdateURL(v string) *CreateChangeSetRequest
	GetStackPolicyDuringUpdateURL() *string
	SetStackPolicyURL(v string) *CreateChangeSetRequest
	GetStackPolicyURL() *string
	SetTags(v []*CreateChangeSetRequestTags) *CreateChangeSetRequest
	GetTags() []*CreateChangeSetRequestTags
	SetTaintResources(v []*string) *CreateChangeSetRequest
	GetTaintResources() []*string
	SetTemplateBody(v string) *CreateChangeSetRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *CreateChangeSetRequest
	GetTemplateId() *string
	SetTemplateScratchId(v string) *CreateChangeSetRequest
	GetTemplateScratchId() *string
	SetTemplateURL(v string) *CreateChangeSetRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *CreateChangeSetRequest
	GetTemplateVersion() *string
	SetTimeoutInMinutes(v int64) *CreateChangeSetRequest
	GetTimeoutInMinutes() *int64
	SetUsePreviousParameters(v bool) *CreateChangeSetRequest
	GetUsePreviousParameters() *bool
}

type CreateChangeSetRequest struct {
	// The name of the change set.\\
	//
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). The name must start with a digit or a letter.
	//
	// > Make sure that the name is unique among all names of change sets that are associated with the specified stack.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyChangeSet
	ChangeSetName *string `json:"ChangeSetName,omitempty" xml:"ChangeSetName,omitempty"`
	// The type of the change set. Valid values:
	//
	// 	- CREATE: creates a change set for a new stack.
	//
	// 	- UPDATE (default): creates a change set for an existing stack.
	//
	// 	- IMPORT: creates a change set for a new stack or an existing stack to import resources that are not managed by ROS.
	//
	// If you set ChangeSetType to CREATE, ROS creates a stack. The stack remains in the `REVIEW_IN_PROGRESS` state until you execute the change set.
	//
	// >
	//
	// 	- You cannot set ChangeSetType to UPDATE when you create a change set for a new stack. You cannot set ChangeSetType to CREATE when you create a change set for an existing stack.
	//
	// 	- If you set ChangeSetType to Import, you cannot configure a stack policy. You can specify ChangeSetType only when you create or update a stack.
	//
	// example:
	//
	// UPDATE
	ChangeSetType *string `json:"ChangeSetType,omitempty" xml:"ChangeSetType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests.\\
	//
	// The token can contain letters, digits, hyphens (-), and underscores (_) and cannot exceed 64 characters in length.\\
	//
	// For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the change set. The description can be up to 1,024 bytes in length.
	//
	// example:
	//
	// It is a demo.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to disable rollback when the stack fails to be created.\\
	//
	// Valid values:
	//
	// 	- true: disables rollback for the stack when the stack fails to be created.
	//
	// 	- false (default): enables rollback for the stack when the stack fails to be created.
	//
	// > This parameter takes effect only if you set ChangeSetType to CREATE or IMPORT.
	//
	// example:
	//
	// false
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The callback URLs that are used to receive stack events.
	//
	// example:
	//
	// http://my-site.com/ros-notify
	NotificationURLs []*string `json:"NotificationURLs,omitempty" xml:"NotificationURLs,omitempty" type:"Repeated"`
	// The maximum number of concurrent operations that can be performed on resources. By default, this parameter is empty. You can set this parameter to an integer that is greater than or equal to 0. If you set this parameter to a specific value, ROS associates the value with the stack. The value can affect subsequent operations on the stack.
	//
	// This parameter takes effect only if you set ChangeSetType to CREATE or UPDATE.
	//
	// 	- Valid values for change sets of the CREATE type:
	//
	//     	- If you set this parameter to an integer that is greater than 0, the integer is used.
	//
	//     	- If you set this parameter to 0 or leave this parameter empty, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	//
	// 	- Valid values for change sets of the UPDATE type:
	//
	//     	- If you set this parameter to an integer that is greater than 0, the integer is used.
	//
	//     	- If you set this parameter to 0, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	//
	//     	- If you leave this parameter empty, the value that you specified for this parameter in the previous request is used. If you left this parameter empty in the previous request, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	//
	// example:
	//
	// 1
	Parallelism *int64 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The parameters that are defined in the template.
	Parameters []*CreateChangeSetRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The name of the Resource Access Management (RAM) role. ROS assumes the RAM role to create the stack and uses the credentials of the role to call the APIs of Alibaba Cloud services.\\
	//
	// ROS assumes the RAM role to perform operations on the stack. If you have permissions to perform operations on the stack, ROS assumes the RAM role even if you do not have permissions to use the RAM role. You must make sure that permissions are granted to the RAM role based on the principle of least privilege.\\
	//
	// If you do not specify this parameter, ROS assumes the existing role of the stack. If no roles are available, ROS uses a temporary credential that is generated from the credentials of your account.\\
	//
	// The RAM role name can be up to 64 characters in length.
	//
	// For more information about RAM roles, see [Use a stack role](https://help.aliyun.com/document_detail/2568025.html).
	//
	// example:
	//
	// test-role
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the change set.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable replacement update if a resource property is changed and you cannot modify the new resource property. For a change, the physical ID of the resource remains unchanged. For a replacement update, the existing resource is deleted, a new resource is created, and the physical ID of the resource is changed. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled (default)
	//
	// > Operations that you perform to modify the resource properties for an update take precedence over operations you perform to replace the resource properties for an update. This parameter takes effect only if you set ChangeSetType to UPDATE.
	//
	// example:
	//
	// Disabled
	ReplacementOption *string `json:"ReplacementOption,omitempty" xml:"ReplacementOption,omitempty"`
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resources that you want to import to the stack.
	ResourcesToImport []*CreateChangeSetRequestResourcesToImport `json:"ResourcesToImport,omitempty" xml:"ResourcesToImport,omitempty" type:"Repeated"`
	// The ID of the stack for which you want to create the change set. ROS compares the stack information with the information that you submit, such as an updated template or parameter value, to generate the change set.\\
	//
	// You can call the [ListStacks](https://help.aliyun.com/document_detail/610818.html) operation to query the stack ID.
	//
	// >  This parameter takes effect only when ChangeSetType is set to UPDATE or IMPORT.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The name of the stack for which you want to create the change set.\\
	//
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). The name must start with a digit or a letter.
	//
	// > This parameter takes effect only if you set ChangeSetType to CREATE or IMPORT.
	//
	// example:
	//
	// MyStack
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The structure that contains the stack policy body. The policy body must be 1 to 16,384 bytes in length.
	//
	// If you set ChangeSetType to **CREATE**, you can specify StackPolicyBody or StackPolicyURL.
	//
	// If you set ChangeSetType to **UPDATE**, you can specify only one of the following parameters:
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
	// {"Statement":[{"Effect":"Allow","Action":"Update:*","Principal":"*","Resource":"*"}]}
	StackPolicyBody *string `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	// The structure of the temporary overriding stack policy. The policy body must be 1 to 16,384 bytes in length.\\
	//
	// If you need to update protected resources, specify a temporary overriding stack policy for the resources during the update. If you do not specify a temporary overriding stack policy, the existing stack policy that is associated with the stack is used.\\
	//
	// This parameter takes effect only if you set ChangeSetType to UPDATE. You can specify only one of the following parameters:
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
	// {"Statement":[{"Effect":"Allow","Action":"Update:*","Principal":"*","Resource":"*"}]}
	StackPolicyDuringUpdateBody *string `json:"StackPolicyDuringUpdateBody,omitempty" xml:"StackPolicyDuringUpdateBody,omitempty"`
	// The URL of the stack policy based on which the stack is updated. The URL must point to a policy that is located on an HTTP or HTTPS web server or in an OSS bucket, such as oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy file can be up to 16,384 bytes in length.
	//
	// > If you do not specify the region of the OSS bucket, the value of RegionId is used.
	//
	// The URL can be up to 1,350 bytes in length.\\
	//
	// If you need to update protected resources, specify a temporary overriding stack policy for the resources during the update. If you do not specify a stack policy, the existing policy that is associated with the stack is used. This parameter takes effect only if you set ChangeSetType to UPDATE. You can specify only one of the following parameters:
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
	// oss://ros/stack-policy/demo
	StackPolicyDuringUpdateURL *string `json:"StackPolicyDuringUpdateURL,omitempty" xml:"StackPolicyDuringUpdateURL,omitempty"`
	// The URL of the file that contains the stack policy. The URL must point to a policy that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo or oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy file can be up to 16,384 bytes in length.
	//
	// The URL can be up to 1,350 bytes in length.
	//
	// >  If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// If you set ChangeSetType to **CREATE**, you can specify StackPolicyBody or StackPolicyURL.
	//
	// If you set ChangeSetType to **UPDATE**, you can specify only one of the following parameters:
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
	// oss://ros/stack-policy/demo
	StackPolicyURL *string                       `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
	Tags           []*CreateChangeSetRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TaintResources []*string                     `json:"TaintResources,omitempty" xml:"TaintResources,omitempty" type:"Repeated"`
	TemplateBody   *string                       `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The template ID. This parameter applies to shared templates and private templates.
	//
	// You can call the [ListTemplates](https://help.aliyun.com/document_detail/610842.html) operation to query the template ID.
	//
	// >  You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the resource scenario. In this example, this parameter specifies the ID of a resource management scenario.
	//
	// This parameter takes effect only when ChangeSetType is set to IMPORT. TemplateScratchId is supported only when you import resources to create a new stack.
	//
	// If you want to use a resource management scenario to import resources, you can specify only TemplateScratchId rather than configuring parameters related to templates.
	//
	// You can call the [ListTemplateScratches](https://help.aliyun.com/document_detail/610832.html) operation to query the ID of the resource management scenario.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an OSS bucket, such as oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length.
	//
	// > If you do not specify the region of the OSS bucket, the value of RegionId is used.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// The URL can be up to 1,024 bytes in length.
	//
	// example:
	//
	// oss://ros/template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template.
	//
	// > This parameter takes effect only if you specify TemplateId.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The amount of time that can elapse before the stack enters the CREATE_FAILED or UPDATE_FAILED state.\\
	//
	// If you set ChangeSetType to CREATE, this parameter is required. If you set ChangeSetType to UPDATE, this parameter is optional.
	//
	// 	- Unit: minutes.
	//
	// 	- Valid values: 10 to 1440.
	//
	// 	- Default value: 60.
	//
	// example:
	//
	// 12
	TimeoutInMinutes *int64 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	// Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// > This parameter takes effect only if you set ChangeSetType to UPDATE or IMPORT.
	//
	// example:
	//
	// true
	UsePreviousParameters *bool `json:"UsePreviousParameters,omitempty" xml:"UsePreviousParameters,omitempty"`
}

func (s CreateChangeSetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateChangeSetRequest) GoString() string {
	return s.String()
}

func (s *CreateChangeSetRequest) GetChangeSetName() *string {
	return s.ChangeSetName
}

func (s *CreateChangeSetRequest) GetChangeSetType() *string {
	return s.ChangeSetType
}

func (s *CreateChangeSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateChangeSetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateChangeSetRequest) GetDisableRollback() *bool {
	return s.DisableRollback
}

func (s *CreateChangeSetRequest) GetNotificationURLs() []*string {
	return s.NotificationURLs
}

func (s *CreateChangeSetRequest) GetParallelism() *int64 {
	return s.Parallelism
}

func (s *CreateChangeSetRequest) GetParameters() []*CreateChangeSetRequestParameters {
	return s.Parameters
}

func (s *CreateChangeSetRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateChangeSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateChangeSetRequest) GetReplacementOption() *string {
	return s.ReplacementOption
}

func (s *CreateChangeSetRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateChangeSetRequest) GetResourcesToImport() []*CreateChangeSetRequestResourcesToImport {
	return s.ResourcesToImport
}

func (s *CreateChangeSetRequest) GetStackId() *string {
	return s.StackId
}

func (s *CreateChangeSetRequest) GetStackName() *string {
	return s.StackName
}

func (s *CreateChangeSetRequest) GetStackPolicyBody() *string {
	return s.StackPolicyBody
}

func (s *CreateChangeSetRequest) GetStackPolicyDuringUpdateBody() *string {
	return s.StackPolicyDuringUpdateBody
}

func (s *CreateChangeSetRequest) GetStackPolicyDuringUpdateURL() *string {
	return s.StackPolicyDuringUpdateURL
}

func (s *CreateChangeSetRequest) GetStackPolicyURL() *string {
	return s.StackPolicyURL
}

func (s *CreateChangeSetRequest) GetTags() []*CreateChangeSetRequestTags {
	return s.Tags
}

func (s *CreateChangeSetRequest) GetTaintResources() []*string {
	return s.TaintResources
}

func (s *CreateChangeSetRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *CreateChangeSetRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateChangeSetRequest) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *CreateChangeSetRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *CreateChangeSetRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *CreateChangeSetRequest) GetTimeoutInMinutes() *int64 {
	return s.TimeoutInMinutes
}

func (s *CreateChangeSetRequest) GetUsePreviousParameters() *bool {
	return s.UsePreviousParameters
}

func (s *CreateChangeSetRequest) SetChangeSetName(v string) *CreateChangeSetRequest {
	s.ChangeSetName = &v
	return s
}

func (s *CreateChangeSetRequest) SetChangeSetType(v string) *CreateChangeSetRequest {
	s.ChangeSetType = &v
	return s
}

func (s *CreateChangeSetRequest) SetClientToken(v string) *CreateChangeSetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateChangeSetRequest) SetDescription(v string) *CreateChangeSetRequest {
	s.Description = &v
	return s
}

func (s *CreateChangeSetRequest) SetDisableRollback(v bool) *CreateChangeSetRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateChangeSetRequest) SetNotificationURLs(v []*string) *CreateChangeSetRequest {
	s.NotificationURLs = v
	return s
}

func (s *CreateChangeSetRequest) SetParallelism(v int64) *CreateChangeSetRequest {
	s.Parallelism = &v
	return s
}

func (s *CreateChangeSetRequest) SetParameters(v []*CreateChangeSetRequestParameters) *CreateChangeSetRequest {
	s.Parameters = v
	return s
}

func (s *CreateChangeSetRequest) SetRamRoleName(v string) *CreateChangeSetRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateChangeSetRequest) SetRegionId(v string) *CreateChangeSetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateChangeSetRequest) SetReplacementOption(v string) *CreateChangeSetRequest {
	s.ReplacementOption = &v
	return s
}

func (s *CreateChangeSetRequest) SetResourceGroupId(v string) *CreateChangeSetRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateChangeSetRequest) SetResourcesToImport(v []*CreateChangeSetRequestResourcesToImport) *CreateChangeSetRequest {
	s.ResourcesToImport = v
	return s
}

func (s *CreateChangeSetRequest) SetStackId(v string) *CreateChangeSetRequest {
	s.StackId = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackName(v string) *CreateChangeSetRequest {
	s.StackName = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackPolicyBody(v string) *CreateChangeSetRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackPolicyDuringUpdateBody(v string) *CreateChangeSetRequest {
	s.StackPolicyDuringUpdateBody = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackPolicyDuringUpdateURL(v string) *CreateChangeSetRequest {
	s.StackPolicyDuringUpdateURL = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackPolicyURL(v string) *CreateChangeSetRequest {
	s.StackPolicyURL = &v
	return s
}

func (s *CreateChangeSetRequest) SetTags(v []*CreateChangeSetRequestTags) *CreateChangeSetRequest {
	s.Tags = v
	return s
}

func (s *CreateChangeSetRequest) SetTaintResources(v []*string) *CreateChangeSetRequest {
	s.TaintResources = v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateBody(v string) *CreateChangeSetRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateId(v string) *CreateChangeSetRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateScratchId(v string) *CreateChangeSetRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateURL(v string) *CreateChangeSetRequest {
	s.TemplateURL = &v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateVersion(v string) *CreateChangeSetRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateChangeSetRequest) SetTimeoutInMinutes(v int64) *CreateChangeSetRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *CreateChangeSetRequest) SetUsePreviousParameters(v bool) *CreateChangeSetRequest {
	s.UsePreviousParameters = &v
	return s
}

func (s *CreateChangeSetRequest) Validate() error {
	return dara.Validate(s)
}

type CreateChangeSetRequestParameters struct {
	// The key of parameter N that is defined in the template. If you do not specify the key and value of a parameter, ROS uses the default name and value that are defined in the template. Maximum value of N: 200.
	//
	// >  Parameters is optional. If you specify Parameters, you must also specify Parameters.N.ParameterKey.
	//
	// This parameter is required.
	//
	// example:
	//
	// Amount
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N that is defined in the template. Maximum value of N: 200.
	//
	// >  Parameters is optional. If you specify Parameters, you must also specify Parameters.N.ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateChangeSetRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateChangeSetRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateChangeSetRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *CreateChangeSetRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *CreateChangeSetRequestParameters) SetParameterKey(v string) *CreateChangeSetRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateChangeSetRequestParameters) SetParameterValue(v string) *CreateChangeSetRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *CreateChangeSetRequestParameters) Validate() error {
	return dara.Validate(s)
}

type CreateChangeSetRequestResourcesToImport struct {
	// The logical ID of resource N. The logical ID is the name of the resource defined in the template.
	//
	// >  This parameter takes effect only when ChangeSetType is set to IMPORT. ResourcesToImport is optional. If you specify ResourcesToImport, you must specify ResourcesToImport.N.LogicalResourceId.
	//
	// example:
	//
	// Vpc
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The key-value mapping between strings. The key-value mapping is used to identify resource N that you want to import. The key-value mapping must be a JSON string.\\
	//
	// A key is an identifier property of a resource and a value is the property value. For example, the key of the ALIYUN::ECS::VPC resource is VpcId and the value is `vpc-2zevx9ios****`.
	//
	// You can call the [GetTemplateSummary](https://help.aliyun.com/document_detail/172485.html) operation to query the identifier property of the resource.
	//
	// >  This parameter takes effect only when ChangeSetType is set to IMPORT. ResourcesToImport is optional. If you specify ResourcesToImport, you must specify ResourcesToImport.N.ResourceIdentifier.
	//
	// example:
	//
	// {"VpcId": "vpc-2zevx9ios******"}
	ResourceIdentifier *string `json:"ResourceIdentifier,omitempty" xml:"ResourceIdentifier,omitempty"`
	// The type of resource N. The resource type must be the same as the resource type that is defined in the template.
	//
	// >  This parameter takes effect only when ChangeSetType is set to IMPORT. ResourcesToImport is optional. If you specify ResourcesToImport, you must specify ResourcesToImport.N.ResourceType.
	//
	// example:
	//
	// ALIYUN::ECS::VPC
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateChangeSetRequestResourcesToImport) String() string {
	return dara.Prettify(s)
}

func (s CreateChangeSetRequestResourcesToImport) GoString() string {
	return s.String()
}

func (s *CreateChangeSetRequestResourcesToImport) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *CreateChangeSetRequestResourcesToImport) GetResourceIdentifier() *string {
	return s.ResourceIdentifier
}

func (s *CreateChangeSetRequestResourcesToImport) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateChangeSetRequestResourcesToImport) SetLogicalResourceId(v string) *CreateChangeSetRequestResourcesToImport {
	s.LogicalResourceId = &v
	return s
}

func (s *CreateChangeSetRequestResourcesToImport) SetResourceIdentifier(v string) *CreateChangeSetRequestResourcesToImport {
	s.ResourceIdentifier = &v
	return s
}

func (s *CreateChangeSetRequestResourcesToImport) SetResourceType(v string) *CreateChangeSetRequestResourcesToImport {
	s.ResourceType = &v
	return s
}

func (s *CreateChangeSetRequestResourcesToImport) Validate() error {
	return dara.Validate(s)
}

type CreateChangeSetRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateChangeSetRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateChangeSetRequestTags) GoString() string {
	return s.String()
}

func (s *CreateChangeSetRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateChangeSetRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateChangeSetRequestTags) SetKey(v string) *CreateChangeSetRequestTags {
	s.Key = &v
	return s
}

func (s *CreateChangeSetRequestTags) SetValue(v string) *CreateChangeSetRequestTags {
	s.Value = &v
	return s
}

func (s *CreateChangeSetRequestTags) Validate() error {
	return dara.Validate(s)
}
