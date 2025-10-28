// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStackGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdministrationRoleName(v string) *CreateStackGroupShrinkRequest
	GetAdministrationRoleName() *string
	SetAutoDeploymentShrink(v string) *CreateStackGroupShrinkRequest
	GetAutoDeploymentShrink() *string
	SetCapabilities(v []*string) *CreateStackGroupShrinkRequest
	GetCapabilities() []*string
	SetClientToken(v string) *CreateStackGroupShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreateStackGroupShrinkRequest
	GetDescription() *string
	SetExecutionRoleName(v string) *CreateStackGroupShrinkRequest
	GetExecutionRoleName() *string
	SetParameters(v []*CreateStackGroupShrinkRequestParameters) *CreateStackGroupShrinkRequest
	GetParameters() []*CreateStackGroupShrinkRequestParameters
	SetPermissionModel(v string) *CreateStackGroupShrinkRequest
	GetPermissionModel() *string
	SetRegionId(v string) *CreateStackGroupShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateStackGroupShrinkRequest
	GetResourceGroupId() *string
	SetStackArn(v string) *CreateStackGroupShrinkRequest
	GetStackArn() *string
	SetStackGroupName(v string) *CreateStackGroupShrinkRequest
	GetStackGroupName() *string
	SetTags(v []*CreateStackGroupShrinkRequestTags) *CreateStackGroupShrinkRequest
	GetTags() []*CreateStackGroupShrinkRequestTags
	SetTemplateBody(v string) *CreateStackGroupShrinkRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *CreateStackGroupShrinkRequest
	GetTemplateId() *string
	SetTemplateURL(v string) *CreateStackGroupShrinkRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *CreateStackGroupShrinkRequest
	GetTemplateVersion() *string
}

type CreateStackGroupShrinkRequest struct {
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
	AutoDeploymentShrink *string `json:"AutoDeployment,omitempty" xml:"AutoDeployment,omitempty"`
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
	Parameters []*CreateStackGroupShrinkRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
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
	Tags         []*CreateStackGroupShrinkRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateBody *string                              `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
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

func (s CreateStackGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStackGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStackGroupShrinkRequest) GetAdministrationRoleName() *string {
	return s.AdministrationRoleName
}

func (s *CreateStackGroupShrinkRequest) GetAutoDeploymentShrink() *string {
	return s.AutoDeploymentShrink
}

func (s *CreateStackGroupShrinkRequest) GetCapabilities() []*string {
	return s.Capabilities
}

func (s *CreateStackGroupShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateStackGroupShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateStackGroupShrinkRequest) GetExecutionRoleName() *string {
	return s.ExecutionRoleName
}

func (s *CreateStackGroupShrinkRequest) GetParameters() []*CreateStackGroupShrinkRequestParameters {
	return s.Parameters
}

func (s *CreateStackGroupShrinkRequest) GetPermissionModel() *string {
	return s.PermissionModel
}

func (s *CreateStackGroupShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateStackGroupShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateStackGroupShrinkRequest) GetStackArn() *string {
	return s.StackArn
}

func (s *CreateStackGroupShrinkRequest) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *CreateStackGroupShrinkRequest) GetTags() []*CreateStackGroupShrinkRequestTags {
	return s.Tags
}

func (s *CreateStackGroupShrinkRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *CreateStackGroupShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateStackGroupShrinkRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *CreateStackGroupShrinkRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *CreateStackGroupShrinkRequest) SetAdministrationRoleName(v string) *CreateStackGroupShrinkRequest {
	s.AdministrationRoleName = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetAutoDeploymentShrink(v string) *CreateStackGroupShrinkRequest {
	s.AutoDeploymentShrink = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetCapabilities(v []*string) *CreateStackGroupShrinkRequest {
	s.Capabilities = v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetClientToken(v string) *CreateStackGroupShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetDescription(v string) *CreateStackGroupShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetExecutionRoleName(v string) *CreateStackGroupShrinkRequest {
	s.ExecutionRoleName = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetParameters(v []*CreateStackGroupShrinkRequestParameters) *CreateStackGroupShrinkRequest {
	s.Parameters = v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetPermissionModel(v string) *CreateStackGroupShrinkRequest {
	s.PermissionModel = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetRegionId(v string) *CreateStackGroupShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetResourceGroupId(v string) *CreateStackGroupShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetStackArn(v string) *CreateStackGroupShrinkRequest {
	s.StackArn = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetStackGroupName(v string) *CreateStackGroupShrinkRequest {
	s.StackGroupName = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetTags(v []*CreateStackGroupShrinkRequestTags) *CreateStackGroupShrinkRequest {
	s.Tags = v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetTemplateBody(v string) *CreateStackGroupShrinkRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetTemplateId(v string) *CreateStackGroupShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetTemplateURL(v string) *CreateStackGroupShrinkRequest {
	s.TemplateURL = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) SetTemplateVersion(v string) *CreateStackGroupShrinkRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateStackGroupShrinkRequest) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateStackGroupShrinkRequestParameters struct {
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

func (s CreateStackGroupShrinkRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateStackGroupShrinkRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateStackGroupShrinkRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *CreateStackGroupShrinkRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *CreateStackGroupShrinkRequestParameters) SetParameterKey(v string) *CreateStackGroupShrinkRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackGroupShrinkRequestParameters) SetParameterValue(v string) *CreateStackGroupShrinkRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *CreateStackGroupShrinkRequestParameters) Validate() error {
	return dara.Validate(s)
}

type CreateStackGroupShrinkRequestTags struct {
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

func (s CreateStackGroupShrinkRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateStackGroupShrinkRequestTags) GoString() string {
	return s.String()
}

func (s *CreateStackGroupShrinkRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateStackGroupShrinkRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateStackGroupShrinkRequestTags) SetKey(v string) *CreateStackGroupShrinkRequestTags {
	s.Key = &v
	return s
}

func (s *CreateStackGroupShrinkRequestTags) SetValue(v string) *CreateStackGroupShrinkRequestTags {
	s.Value = &v
	return s
}

func (s *CreateStackGroupShrinkRequestTags) Validate() error {
	return dara.Validate(s)
}
