// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *UpdateStackGroupRequest
	GetAccountIds() []*string
	SetAdministrationRoleName(v string) *UpdateStackGroupRequest
	GetAdministrationRoleName() *string
	SetAutoDeployment(v *UpdateStackGroupRequestAutoDeployment) *UpdateStackGroupRequest
	GetAutoDeployment() *UpdateStackGroupRequestAutoDeployment
	SetCapabilities(v []*string) *UpdateStackGroupRequest
	GetCapabilities() []*string
	SetClientToken(v string) *UpdateStackGroupRequest
	GetClientToken() *string
	SetDeploymentOptions(v []*string) *UpdateStackGroupRequest
	GetDeploymentOptions() []*string
	SetDeploymentTargets(v *UpdateStackGroupRequestDeploymentTargets) *UpdateStackGroupRequest
	GetDeploymentTargets() *UpdateStackGroupRequestDeploymentTargets
	SetDescription(v string) *UpdateStackGroupRequest
	GetDescription() *string
	SetExecutionRoleName(v string) *UpdateStackGroupRequest
	GetExecutionRoleName() *string
	SetOperationDescription(v string) *UpdateStackGroupRequest
	GetOperationDescription() *string
	SetOperationPreferences(v map[string]interface{}) *UpdateStackGroupRequest
	GetOperationPreferences() map[string]interface{}
	SetParameters(v []*UpdateStackGroupRequestParameters) *UpdateStackGroupRequest
	GetParameters() []*UpdateStackGroupRequestParameters
	SetPermissionModel(v string) *UpdateStackGroupRequest
	GetPermissionModel() *string
	SetRegionId(v string) *UpdateStackGroupRequest
	GetRegionId() *string
	SetRegionIds(v []*string) *UpdateStackGroupRequest
	GetRegionIds() []*string
	SetStackGroupName(v string) *UpdateStackGroupRequest
	GetStackGroupName() *string
	SetTemplateBody(v string) *UpdateStackGroupRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *UpdateStackGroupRequest
	GetTemplateId() *string
	SetTemplateURL(v string) *UpdateStackGroupRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *UpdateStackGroupRequest
	GetTemplateVersion() *string
}

type UpdateStackGroupRequest struct {
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Alibaba Cloud Object Storage Service (OSS) bucket. The template body must be 1 to 524,288 bytes in length. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// >  You must specify only one of the TemplateBody, TemplateURL, and TemplateId parameters.
	//
	// example:
	//
	// ["12****"]
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
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
	AutoDeployment *UpdateStackGroupRequestAutoDeployment `json:"AutoDeployment,omitempty" xml:"AutoDeployment,omitempty" type:"Struct"`
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
	DeploymentTargets *UpdateStackGroupRequestDeploymentTargets `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty" type:"Struct"`
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
	OperationPreferences map[string]interface{} `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// Specifies whether to enable automatic deployment.
	//
	// Valid values:
	//
	// 	- true: enables automatic deployment. If you add a member to the folder to which the stack group belongs after you enable automatic deployment, the stack group deploys its stack instances within the member. If you remove a member from the folder, the stack group deletes stack instances that are deployed within the member.
	//
	// 	- false: disables automatic deployment. After you disable automatic deployment, the stack instances remain unchanged even if members in the folder change.
	Parameters []*UpdateStackGroupRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
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
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
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

func (s UpdateStackGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *UpdateStackGroupRequest) GetAdministrationRoleName() *string {
	return s.AdministrationRoleName
}

func (s *UpdateStackGroupRequest) GetAutoDeployment() *UpdateStackGroupRequestAutoDeployment {
	return s.AutoDeployment
}

func (s *UpdateStackGroupRequest) GetCapabilities() []*string {
	return s.Capabilities
}

func (s *UpdateStackGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateStackGroupRequest) GetDeploymentOptions() []*string {
	return s.DeploymentOptions
}

func (s *UpdateStackGroupRequest) GetDeploymentTargets() *UpdateStackGroupRequestDeploymentTargets {
	return s.DeploymentTargets
}

func (s *UpdateStackGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateStackGroupRequest) GetExecutionRoleName() *string {
	return s.ExecutionRoleName
}

func (s *UpdateStackGroupRequest) GetOperationDescription() *string {
	return s.OperationDescription
}

func (s *UpdateStackGroupRequest) GetOperationPreferences() map[string]interface{} {
	return s.OperationPreferences
}

func (s *UpdateStackGroupRequest) GetParameters() []*UpdateStackGroupRequestParameters {
	return s.Parameters
}

func (s *UpdateStackGroupRequest) GetPermissionModel() *string {
	return s.PermissionModel
}

func (s *UpdateStackGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateStackGroupRequest) GetRegionIds() []*string {
	return s.RegionIds
}

func (s *UpdateStackGroupRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *UpdateStackGroupRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *UpdateStackGroupRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateStackGroupRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *UpdateStackGroupRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *UpdateStackGroupRequest) SetAccountIds(v []*string) *UpdateStackGroupRequest {
	s.AccountIds = v
	return s
}

func (s *UpdateStackGroupRequest) SetAdministrationRoleName(v string) *UpdateStackGroupRequest {
	s.AdministrationRoleName = &v
	return s
}

func (s *UpdateStackGroupRequest) SetAutoDeployment(v *UpdateStackGroupRequestAutoDeployment) *UpdateStackGroupRequest {
	s.AutoDeployment = v
	return s
}

func (s *UpdateStackGroupRequest) SetCapabilities(v []*string) *UpdateStackGroupRequest {
	s.Capabilities = v
	return s
}

func (s *UpdateStackGroupRequest) SetClientToken(v string) *UpdateStackGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackGroupRequest) SetDeploymentOptions(v []*string) *UpdateStackGroupRequest {
	s.DeploymentOptions = v
	return s
}

func (s *UpdateStackGroupRequest) SetDeploymentTargets(v *UpdateStackGroupRequestDeploymentTargets) *UpdateStackGroupRequest {
	s.DeploymentTargets = v
	return s
}

func (s *UpdateStackGroupRequest) SetDescription(v string) *UpdateStackGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateStackGroupRequest) SetExecutionRoleName(v string) *UpdateStackGroupRequest {
	s.ExecutionRoleName = &v
	return s
}

func (s *UpdateStackGroupRequest) SetOperationDescription(v string) *UpdateStackGroupRequest {
	s.OperationDescription = &v
	return s
}

func (s *UpdateStackGroupRequest) SetOperationPreferences(v map[string]interface{}) *UpdateStackGroupRequest {
	s.OperationPreferences = v
	return s
}

func (s *UpdateStackGroupRequest) SetParameters(v []*UpdateStackGroupRequestParameters) *UpdateStackGroupRequest {
	s.Parameters = v
	return s
}

func (s *UpdateStackGroupRequest) SetPermissionModel(v string) *UpdateStackGroupRequest {
	s.PermissionModel = &v
	return s
}

func (s *UpdateStackGroupRequest) SetRegionId(v string) *UpdateStackGroupRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackGroupRequest) SetRegionIds(v []*string) *UpdateStackGroupRequest {
	s.RegionIds = v
	return s
}

func (s *UpdateStackGroupRequest) SetStackGroupName(v string) *UpdateStackGroupRequest {
	s.StackGroupName = &v
	return s
}

func (s *UpdateStackGroupRequest) SetTemplateBody(v string) *UpdateStackGroupRequest {
	s.TemplateBody = &v
	return s
}

func (s *UpdateStackGroupRequest) SetTemplateId(v string) *UpdateStackGroupRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateStackGroupRequest) SetTemplateURL(v string) *UpdateStackGroupRequest {
	s.TemplateURL = &v
	return s
}

func (s *UpdateStackGroupRequest) SetTemplateVersion(v string) *UpdateStackGroupRequest {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateStackGroupRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateStackGroupRequestAutoDeployment struct {
	// The IDs of the members in the resource directory. You can specify a maximum of 20 member IDs.
	//
	// >  To view the member IDs, go to the **Overview*	- page in the **Resource Management*	- console. For more information, see [View the detailed information of a member](https://help.aliyun.com/document_detail/111624.html).
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The IDs of the members in the resource directory. You can specify a maximum of 20 member IDs.
	//
	// >  To view the member IDs, go to the **Overview*	- page in the **Resource Management*	- console. For more information, see [View the detailed information of a member](https://help.aliyun.com/document_detail/111624.html).
	//
	// example:
	//
	// true
	RetainStacksOnAccountRemoval *bool `json:"RetainStacksOnAccountRemoval,omitempty" xml:"RetainStacksOnAccountRemoval,omitempty"`
}

func (s UpdateStackGroupRequestAutoDeployment) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackGroupRequestAutoDeployment) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupRequestAutoDeployment) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateStackGroupRequestAutoDeployment) GetRetainStacksOnAccountRemoval() *bool {
	return s.RetainStacksOnAccountRemoval
}

func (s *UpdateStackGroupRequestAutoDeployment) SetEnabled(v bool) *UpdateStackGroupRequestAutoDeployment {
	s.Enabled = &v
	return s
}

func (s *UpdateStackGroupRequestAutoDeployment) SetRetainStacksOnAccountRemoval(v bool) *UpdateStackGroupRequestAutoDeployment {
	s.RetainStacksOnAccountRemoval = &v
	return s
}

func (s *UpdateStackGroupRequestAutoDeployment) Validate() error {
	return dara.Validate(s)
}

type UpdateStackGroupRequestDeploymentTargets struct {
	// The list of one or more Alibaba Cloud accounts with which you want to share or unshare the template.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The ID of the operation.
	RdFolderIds []*string `json:"RdFolderIds,omitempty" xml:"RdFolderIds,omitempty" type:"Repeated"`
}

func (s UpdateStackGroupRequestDeploymentTargets) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackGroupRequestDeploymentTargets) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupRequestDeploymentTargets) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *UpdateStackGroupRequestDeploymentTargets) GetRdFolderIds() []*string {
	return s.RdFolderIds
}

func (s *UpdateStackGroupRequestDeploymentTargets) SetAccountIds(v []*string) *UpdateStackGroupRequestDeploymentTargets {
	s.AccountIds = v
	return s
}

func (s *UpdateStackGroupRequestDeploymentTargets) SetRdFolderIds(v []*string) *UpdateStackGroupRequestDeploymentTargets {
	s.RdFolderIds = v
	return s
}

func (s *UpdateStackGroupRequestDeploymentTargets) Validate() error {
	return dara.Validate(s)
}

type UpdateStackGroupRequestParameters struct {
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

func (s UpdateStackGroupRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackGroupRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *UpdateStackGroupRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *UpdateStackGroupRequestParameters) SetParameterKey(v string) *UpdateStackGroupRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackGroupRequestParameters) SetParameterValue(v string) *UpdateStackGroupRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *UpdateStackGroupRequestParameters) Validate() error {
	return dara.Validate(s)
}
