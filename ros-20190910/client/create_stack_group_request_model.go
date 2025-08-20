// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStackGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdministrationRoleName(v string) *CreateStackGroupRequest
	GetAdministrationRoleName() *string
	SetAutoDeployment(v *CreateStackGroupRequestAutoDeployment) *CreateStackGroupRequest
	GetAutoDeployment() *CreateStackGroupRequestAutoDeployment
	SetCapabilities(v []*string) *CreateStackGroupRequest
	GetCapabilities() []*string
	SetClientToken(v string) *CreateStackGroupRequest
	GetClientToken() *string
	SetDescription(v string) *CreateStackGroupRequest
	GetDescription() *string
	SetExecutionRoleName(v string) *CreateStackGroupRequest
	GetExecutionRoleName() *string
	SetParameters(v []*CreateStackGroupRequestParameters) *CreateStackGroupRequest
	GetParameters() []*CreateStackGroupRequestParameters
	SetPermissionModel(v string) *CreateStackGroupRequest
	GetPermissionModel() *string
	SetRegionId(v string) *CreateStackGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateStackGroupRequest
	GetResourceGroupId() *string
	SetStackArn(v string) *CreateStackGroupRequest
	GetStackArn() *string
	SetStackGroupName(v string) *CreateStackGroupRequest
	GetStackGroupName() *string
	SetTags(v []*CreateStackGroupRequestTags) *CreateStackGroupRequest
	GetTags() []*CreateStackGroupRequestTags
	SetTemplateBody(v string) *CreateStackGroupRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *CreateStackGroupRequest
	GetTemplateId() *string
	SetTemplateURL(v string) *CreateStackGroupRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *CreateStackGroupRequest
	GetTemplateVersion() *string
}

type CreateStackGroupRequest struct {
	// The name of the RAM role that you specify for the administrator account when you create a self-managed stack group. ROS assumes the administrator role to perform operations. If you do not specify this parameter, AliyunROSStackGroupAdministrationRole is used as the default value. ROS uses the administrator role to assume the execution role AliyunROSStackGroupExecutionRole to perform operations on the stacks in the stack group.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// AliyunROSStackGroupAdministrationRole
	AdministrationRoleName *string `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	// The information about automatic deployment settings.
	//
	// > You must specify this parameter if PermissionModel is set to SERVICE_MANAGED.
	//
	// example:
	//
	// {"Enabled": true, "RetainStacksOnAccountRemoval": true}
	AutoDeployment *CreateStackGroupRequestAutoDeployment `json:"AutoDeployment,omitempty" xml:"AutoDeployment,omitempty" type:"Struct"`
	// The options for the stack group. You can specify up to one option.
	Capabilities []*string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\\
	//
	// The token can contain letters, digits, underscores (_), and hyphens (-) and cannot exceed 64 characters in length.\\
	//
	// For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the stack group.\\
	//
	// The description must be 1 to 256 characters in length.
	//
	// example:
	//
	// StackGroup Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the RAM role that you specify for the execution account when you create a self-managed stack group. The administrator role AliyunROSStackGroupAdministrationRole assumes the execution role to perform operations. If you do not specify this parameter, AliyunROSStackGroupExecutionRole is used as the default value. ROS assumes the execution role to perform operations on the stacks in the stack group.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// AliyunROSStackGroupExecutionRole
	ExecutionRoleName *string `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	// The parameters of the stack group.
	Parameters []*CreateStackGroupRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The permission model of the stack group.
	//
	// Valid values:
	//
	// 	- SELF_MANAGED (default): the self-managed permission model. If you create a self-managed stack group, you must create RAM roles within the administrator and execution accounts and establish a trust relationship between the accounts. Then, you can deploy stacks within the execution account.
	//
	// 	- SERVICE_MANAGED: the service-managed permission model. If you create a service-managed stack group, ROS creates service-linked roles for the administrator and execution accounts, and the administrator account uses its role to deploy stacks within the execution account.
	//
	// > If you want to use the service-managed permission model to deploy stacks, your account must be the management account or a delegated administrator account of your resource directory and the trusted access feature is enabled for the account. For more information, see [Manage a delegated administrator account](https://help.aliyun.com/document_detail/308253.html) and [Enable trusted access](https://help.aliyun.com/document_detail/298229.html).
	//
	// example:
	//
	// SELF_MANAGED
	PermissionModel *string `json:"PermissionModel,omitempty" xml:"PermissionModel,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. If you do not specify this parameter, the stack group is added to the default resource group.\\
	//
	// For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/94475.html).
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	StackArn        *string `json:"StackArn,omitempty" xml:"StackArn,omitempty"`
	// The name of the stack group. The name must be unique within a region.\\
	//
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The tags of the stack group.
	Tags         []*CreateStackGroupRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateBody *string                        `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Alibaba Cloud Object Storage Service (OSS) bucket. The template body must be 1 to 524,288 bytes in length. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// example:
	//
	// oss://ros-template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. If you do not specify this parameter, the latest version is used.
	//
	// > TemplateVersion takes effect only if you specify TemplateId.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s CreateStackGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStackGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateStackGroupRequest) GetAdministrationRoleName() *string {
	return s.AdministrationRoleName
}

func (s *CreateStackGroupRequest) GetAutoDeployment() *CreateStackGroupRequestAutoDeployment {
	return s.AutoDeployment
}

func (s *CreateStackGroupRequest) GetCapabilities() []*string {
	return s.Capabilities
}

func (s *CreateStackGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateStackGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateStackGroupRequest) GetExecutionRoleName() *string {
	return s.ExecutionRoleName
}

func (s *CreateStackGroupRequest) GetParameters() []*CreateStackGroupRequestParameters {
	return s.Parameters
}

func (s *CreateStackGroupRequest) GetPermissionModel() *string {
	return s.PermissionModel
}

func (s *CreateStackGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStackGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateStackGroupRequest) GetStackArn() *string {
	return s.StackArn
}

func (s *CreateStackGroupRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *CreateStackGroupRequest) GetTags() []*CreateStackGroupRequestTags {
	return s.Tags
}

func (s *CreateStackGroupRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *CreateStackGroupRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateStackGroupRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *CreateStackGroupRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *CreateStackGroupRequest) SetAdministrationRoleName(v string) *CreateStackGroupRequest {
	s.AdministrationRoleName = &v
	return s
}

func (s *CreateStackGroupRequest) SetAutoDeployment(v *CreateStackGroupRequestAutoDeployment) *CreateStackGroupRequest {
	s.AutoDeployment = v
	return s
}

func (s *CreateStackGroupRequest) SetCapabilities(v []*string) *CreateStackGroupRequest {
	s.Capabilities = v
	return s
}

func (s *CreateStackGroupRequest) SetClientToken(v string) *CreateStackGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStackGroupRequest) SetDescription(v string) *CreateStackGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateStackGroupRequest) SetExecutionRoleName(v string) *CreateStackGroupRequest {
	s.ExecutionRoleName = &v
	return s
}

func (s *CreateStackGroupRequest) SetParameters(v []*CreateStackGroupRequestParameters) *CreateStackGroupRequest {
	s.Parameters = v
	return s
}

func (s *CreateStackGroupRequest) SetPermissionModel(v string) *CreateStackGroupRequest {
	s.PermissionModel = &v
	return s
}

func (s *CreateStackGroupRequest) SetRegionId(v string) *CreateStackGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStackGroupRequest) SetResourceGroupId(v string) *CreateStackGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateStackGroupRequest) SetStackArn(v string) *CreateStackGroupRequest {
	s.StackArn = &v
	return s
}

func (s *CreateStackGroupRequest) SetStackGroupName(v string) *CreateStackGroupRequest {
	s.StackGroupName = &v
	return s
}

func (s *CreateStackGroupRequest) SetTags(v []*CreateStackGroupRequestTags) *CreateStackGroupRequest {
	s.Tags = v
	return s
}

func (s *CreateStackGroupRequest) SetTemplateBody(v string) *CreateStackGroupRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateStackGroupRequest) SetTemplateId(v string) *CreateStackGroupRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateStackGroupRequest) SetTemplateURL(v string) *CreateStackGroupRequest {
	s.TemplateURL = &v
	return s
}

func (s *CreateStackGroupRequest) SetTemplateVersion(v string) *CreateStackGroupRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateStackGroupRequest) Validate() error {
	return dara.Validate(s)
}

type CreateStackGroupRequestAutoDeployment struct {
	// Indicates whether automatic deployment is enabled.
	//
	// Valid values:
	//
	// 	- true: Automatic deployment is enabled. If you add a member account to the folder to which the stack group belongs after you enable automatic deployment, ROS automatically adds the stacks in the stack group to the member account. If you remove a member account from the folder, ROS automatically deletes the stacks from the member account.
	//
	// 	- false: Automatic deployment is disabled. After you disable automatic deployment, the stacks remain unchanged when you add member accounts to or remove member accounts from the folder.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether the stacks within a member account are retained when you remove the member account from the folder.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// > You must specify RetainStacksOnAccountRemoval if Enabled is set to true.
	//
	// example:
	//
	// true
	RetainStacksOnAccountRemoval *bool `json:"RetainStacksOnAccountRemoval,omitempty" xml:"RetainStacksOnAccountRemoval,omitempty"`
}

func (s CreateStackGroupRequestAutoDeployment) String() string {
	return dara.Prettify(s)
}

func (s CreateStackGroupRequestAutoDeployment) GoString() string {
	return s.String()
}

func (s *CreateStackGroupRequestAutoDeployment) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateStackGroupRequestAutoDeployment) GetRetainStacksOnAccountRemoval() *bool {
	return s.RetainStacksOnAccountRemoval
}

func (s *CreateStackGroupRequestAutoDeployment) SetEnabled(v bool) *CreateStackGroupRequestAutoDeployment {
	s.Enabled = &v
	return s
}

func (s *CreateStackGroupRequestAutoDeployment) SetRetainStacksOnAccountRemoval(v bool) *CreateStackGroupRequestAutoDeployment {
	s.RetainStacksOnAccountRemoval = &v
	return s
}

func (s *CreateStackGroupRequestAutoDeployment) Validate() error {
	return dara.Validate(s)
}

type CreateStackGroupRequestParameters struct {
	// The key of parameter N. If you do not specify the key and value of a parameter, ROS uses the default name and value that are defined in the template.
	//
	// Maximum value of N: 200.
	//
	// > Parameters is optional. If you specify Parameters, you must also specify Parameters.N.ParameterKey.
	//
	// This parameter is required.
	//
	// example:
	//
	// Amount
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N.
	//
	// Maximum value of N: 200.
	//
	// > Parameters is optional. If you specify Parameters, you must also specify Parameters.N.ParameterValue.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateStackGroupRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateStackGroupRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateStackGroupRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *CreateStackGroupRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *CreateStackGroupRequestParameters) SetParameterKey(v string) *CreateStackGroupRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackGroupRequestParameters) SetParameterValue(v string) *CreateStackGroupRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *CreateStackGroupRequestParameters) Validate() error {
	return dara.Validate(s)
}

type CreateStackGroupRequestTags struct {
	// The tag key of the stack group.
	//
	// > Tags is optional. If you want to specify Tags, you must also specify Tags.N.Key.
	//
	// This parameter is required.
	//
	// example:
	//
	// usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the stack group.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateStackGroupRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateStackGroupRequestTags) GoString() string {
	return s.String()
}

func (s *CreateStackGroupRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateStackGroupRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateStackGroupRequestTags) SetKey(v string) *CreateStackGroupRequestTags {
	s.Key = &v
	return s
}

func (s *CreateStackGroupRequestTags) SetValue(v string) *CreateStackGroupRequestTags {
	s.Value = &v
	return s
}

func (s *CreateStackGroupRequestTags) Validate() error {
	return dara.Validate(s)
}
