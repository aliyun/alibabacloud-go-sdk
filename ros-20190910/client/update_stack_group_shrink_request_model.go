// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIdsShrink(v string) *UpdateStackGroupShrinkRequest
	GetAccountIdsShrink() *string
	SetAdministrationRoleName(v string) *UpdateStackGroupShrinkRequest
	GetAdministrationRoleName() *string
	SetAutoDeploymentShrink(v string) *UpdateStackGroupShrinkRequest
	GetAutoDeploymentShrink() *string
	SetCapabilities(v []*string) *UpdateStackGroupShrinkRequest
	GetCapabilities() []*string
	SetClientToken(v string) *UpdateStackGroupShrinkRequest
	GetClientToken() *string
	SetDeploymentOptions(v []*string) *UpdateStackGroupShrinkRequest
	GetDeploymentOptions() []*string
	SetDeploymentTargetsShrink(v string) *UpdateStackGroupShrinkRequest
	GetDeploymentTargetsShrink() *string
	SetDescription(v string) *UpdateStackGroupShrinkRequest
	GetDescription() *string
	SetExecutionRoleName(v string) *UpdateStackGroupShrinkRequest
	GetExecutionRoleName() *string
	SetOperationDescription(v string) *UpdateStackGroupShrinkRequest
	GetOperationDescription() *string
	SetOperationPreferencesShrink(v string) *UpdateStackGroupShrinkRequest
	GetOperationPreferencesShrink() *string
	SetParameters(v []*UpdateStackGroupShrinkRequestParameters) *UpdateStackGroupShrinkRequest
	GetParameters() []*UpdateStackGroupShrinkRequestParameters
	SetPermissionModel(v string) *UpdateStackGroupShrinkRequest
	GetPermissionModel() *string
	SetRegionId(v string) *UpdateStackGroupShrinkRequest
	GetRegionId() *string
	SetRegionIdsShrink(v string) *UpdateStackGroupShrinkRequest
	GetRegionIdsShrink() *string
	SetStackGroupName(v string) *UpdateStackGroupShrinkRequest
	GetStackGroupName() *string
	SetTemplateBody(v string) *UpdateStackGroupShrinkRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *UpdateStackGroupShrinkRequest
	GetTemplateId() *string
	SetTemplateURL(v string) *UpdateStackGroupShrinkRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *UpdateStackGroupShrinkRequest
	GetTemplateVersion() *string
}

type UpdateStackGroupShrinkRequest struct {
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Alibaba Cloud Object Storage Service (OSS) bucket. The template body must be 1 to 524,288 bytes in length. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// >  You must specify only one of the TemplateBody, TemplateURL, and TemplateId parameters.
	//
	// example:
	//
	// ["12****"]
	AccountIdsShrink *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	// The key of parameter N. If you do not specify the key and value of the parameter, ROS uses the default key and value in the template.
	//
	// Maximum value of N: 200.
	//
	// >  The Parameters parameter is optional. If you set the Parameters parameter, you must set the Parameters.N.ParameterKey parameter.
	//
	// example:
	//
	// AliyunROSStackGroupAdministrationRole
	AdministrationRoleName *string `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	// The IDs of the folders in the resource directory. You can specify up to five folder IDs.
	//
	// You can create stacks within all members in the specified folders. If you create stacks in the Root folder, the stacks are created within all members in the resource directory.
	//
	// >  To view the folder IDs, go to the **Overview*	- page in the **Resource Management*	- console. For more information, see [View the basic information of a folder](https://help.aliyun.com/document_detail/111223.html).
	//
	// example:
	//
	// {"Enabled": true, "RetainStacksOnAccountRemoval": true}
	AutoDeploymentShrink *string `json:"AutoDeployment,omitempty" xml:"AutoDeployment,omitempty"`
	// The option for the stack group. You can specify up to one option.
	Capabilities []*string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// >  You must specify only one of the TemplateBody, TemplateURL, and TemplateId parameters.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken       *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DeploymentOptions []*string `json:"DeploymentOptions,omitempty" xml:"DeploymentOptions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// {"RdFolderIds": ["fd-4PvlVLOL8v"]}
	DeploymentTargetsShrink *string `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.
	//
	// >  You must specify only one of the TemplateBody, TemplateURL, and TemplateId parameters.
	//
	// example:
	//
	// My Stack Group
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The value of parameter N.
	//
	// Maximum value of N: 200.
	//
	// >  The Parameters parameter is optional. If you set the Parameters parameter, you must set the Parameters.N.ParameterValue parameter.
	//
	// example:
	//
	// AliyunROSStackGroupExecutionRole
	ExecutionRoleName *string `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	// The version of the template. If you do not specify a version, the latest version is used.
	//
	// >  This parameter takes effect only if the TemplateId parameter is set.
	//
	// example:
	//
	// Update stack instances in hangzhou
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The list of parameters.
	//
	// example:
	//
	// {"FailureToleranceCount": 1,"MaxConcurrentCount": 2}
	OperationPreferencesShrink *string `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// Specifies whether to enable automatic deployment.
	//
	// Valid values:
	//
	// 	- true: enables automatic deployment. If you add a member to the folder to which the stack group belongs after you enable automatic deployment, the stack group deploys its stack instances within the member. If you remove a member from the folder, the stack group deletes stack instances that are deployed within the member.
	//
	// 	- false: disables automatic deployment. After you disable automatic deployment, the stack instances remain unchanged even if members in the folder change.
	Parameters []*UpdateStackGroupShrinkRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The folder IDs in the resource directory. You can specify a maximum of five folder IDs.
	//
	// You must set at least one of the RdFolderIds and AccountIds parameters. The parameters are subject to the following items:
	//
	// 	- If you set only the RdFolderIds parameter, stacks are deployed within all the members in the specified folders. If you specify the Root folder, ROS deploys the stacks within all the members in the resource directory.
	//
	// 	- If you set only the AccountIds parameter, stacks are deployed within the specified members.
	//
	// 	- If you set both parameters, the accounts specified by AccountIds must be contained in the folders specified by RdFolderIds.
	//
	// >  To view the folder IDs, go to the **Overview*	- page in the **Resource Management*	- console. For more information, see [View the basic information of a folder](https://help.aliyun.com/document_detail/111223.html).
	//
	// example:
	//
	// SELF_MANAGED
	PermissionModel *string `json:"PermissionModel,omitempty" xml:"PermissionModel,omitempty"`
	// The region IDs of stack instances. You can specify a maximum of 20 region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the operation to update the stack group.
	//
	// example:
	//
	// ["cn-hangzhou", "cn-beijing"]
	RegionIdsShrink *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	// The region IDs of stack instances. You can specify a maximum of 20 region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	TemplateBody   *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The permission model.
	//
	// Valid values:
	//
	// 	- SELF_MANAGED: the self-managed permission model. This is the default value. If you use the self-managed model for the stack group, you must create RAM roles for the administrator and execution accounts, and establish a trust relationship between the accounts to deploy stacks within the execution account.
	//
	// 	- SERVICE_MANAGED: the service-managed permission model. If you use the service-managed model for the stack group, ROS creates service-linked roles for the administrator and execution accounts, and the administrator account uses its role to deploy stacks within the execution account.
	//
	// >- If stack instances have been created in the stack group, you cannot switch the permission mode of the stack group.
	//
	// >- If you want to use the service-managed permission model to deploy stacks, your account must be the management account or a delegated administrator account of your resource directory and the trusted access feature is enabled for the account. For more information, see [Step 1: (Optional) Create a delegated administrator account](https://help.aliyun.com/document_detail/308253.html) and [Step 2: Enable trusted access](https://help.aliyun.com/document_detail/298229.html).
	//
	// example:
	//
	// 5ecd1e10-b0e9-4389-a565-e4c15efc****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the RAM role to be assumed by the administrator role AliyunROSStackGroupAdministrationRole. This parameter is required if you want to grant self-managed permissions to the stack group. If you do not specify a value for this parameter, the default value AliyunROSStackGroupExecutionRole is used. You can use this role in ROS to perform operations on the stacks that correspond to stack instances in the stack group.
	//
	// The name must be 1 to 64 characters in length, and can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// oss://ros-template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The information about automatic deployment settings.
	//
	// >  This parameter is required only if the PermissionModel parameter is set to SERVICE_MANAGED.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s UpdateStackGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupShrinkRequest) GetAccountIdsShrink() *string {
	return s.AccountIdsShrink
}

func (s *UpdateStackGroupShrinkRequest) GetAdministrationRoleName() *string {
	return s.AdministrationRoleName
}

func (s *UpdateStackGroupShrinkRequest) GetAutoDeploymentShrink() *string {
	return s.AutoDeploymentShrink
}

func (s *UpdateStackGroupShrinkRequest) GetCapabilities() []*string {
	return s.Capabilities
}

func (s *UpdateStackGroupShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateStackGroupShrinkRequest) GetDeploymentOptions() []*string {
	return s.DeploymentOptions
}

func (s *UpdateStackGroupShrinkRequest) GetDeploymentTargetsShrink() *string {
	return s.DeploymentTargetsShrink
}

func (s *UpdateStackGroupShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateStackGroupShrinkRequest) GetExecutionRoleName() *string {
	return s.ExecutionRoleName
}

func (s *UpdateStackGroupShrinkRequest) GetOperationDescription() *string {
	return s.OperationDescription
}

func (s *UpdateStackGroupShrinkRequest) GetOperationPreferencesShrink() *string {
	return s.OperationPreferencesShrink
}

func (s *UpdateStackGroupShrinkRequest) GetParameters() []*UpdateStackGroupShrinkRequestParameters {
	return s.Parameters
}

func (s *UpdateStackGroupShrinkRequest) GetPermissionModel() *string {
	return s.PermissionModel
}

func (s *UpdateStackGroupShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateStackGroupShrinkRequest) GetRegionIdsShrink() *string {
	return s.RegionIdsShrink
}

func (s *UpdateStackGroupShrinkRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *UpdateStackGroupShrinkRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *UpdateStackGroupShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateStackGroupShrinkRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *UpdateStackGroupShrinkRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *UpdateStackGroupShrinkRequest) SetAccountIdsShrink(v string) *UpdateStackGroupShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetAdministrationRoleName(v string) *UpdateStackGroupShrinkRequest {
	s.AdministrationRoleName = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetAutoDeploymentShrink(v string) *UpdateStackGroupShrinkRequest {
	s.AutoDeploymentShrink = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetCapabilities(v []*string) *UpdateStackGroupShrinkRequest {
	s.Capabilities = v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetClientToken(v string) *UpdateStackGroupShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetDeploymentOptions(v []*string) *UpdateStackGroupShrinkRequest {
	s.DeploymentOptions = v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetDeploymentTargetsShrink(v string) *UpdateStackGroupShrinkRequest {
	s.DeploymentTargetsShrink = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetDescription(v string) *UpdateStackGroupShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetExecutionRoleName(v string) *UpdateStackGroupShrinkRequest {
	s.ExecutionRoleName = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetOperationDescription(v string) *UpdateStackGroupShrinkRequest {
	s.OperationDescription = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetOperationPreferencesShrink(v string) *UpdateStackGroupShrinkRequest {
	s.OperationPreferencesShrink = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetParameters(v []*UpdateStackGroupShrinkRequestParameters) *UpdateStackGroupShrinkRequest {
	s.Parameters = v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetPermissionModel(v string) *UpdateStackGroupShrinkRequest {
	s.PermissionModel = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetRegionId(v string) *UpdateStackGroupShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetRegionIdsShrink(v string) *UpdateStackGroupShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetStackGroupName(v string) *UpdateStackGroupShrinkRequest {
	s.StackGroupName = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetTemplateBody(v string) *UpdateStackGroupShrinkRequest {
	s.TemplateBody = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetTemplateId(v string) *UpdateStackGroupShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetTemplateURL(v string) *UpdateStackGroupShrinkRequest {
	s.TemplateURL = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) SetTemplateVersion(v string) *UpdateStackGroupShrinkRequest {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateStackGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateStackGroupShrinkRequestParameters struct {
	// Specifies whether to retain stacks in a member when you remove the member from the folder.
	//
	// Valid values:
	//
	// 	- true: retains the stacks.
	//
	// 	- false: deletes the stacks.
	//
	// >  This parameter is required if the Enabled parameter is set to true.
	//
	// This parameter is required.
	//
	// example:
	//
	// Amount
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The folders in which you want to use service-managed permissions to update stacks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateStackGroupShrinkRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackGroupShrinkRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupShrinkRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *UpdateStackGroupShrinkRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *UpdateStackGroupShrinkRequestParameters) SetParameterKey(v string) *UpdateStackGroupShrinkRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackGroupShrinkRequestParameters) SetParameterValue(v string) *UpdateStackGroupShrinkRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *UpdateStackGroupShrinkRequestParameters) Validate() error {
	return dara.Validate(s)
}
