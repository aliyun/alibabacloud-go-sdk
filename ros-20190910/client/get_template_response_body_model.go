// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalInfo(v map[string]interface{}) *GetTemplateResponseBody
	GetAdditionalInfo() map[string]interface{}
	SetChangeSetId(v string) *GetTemplateResponseBody
	GetChangeSetId() *string
	SetCreateTime(v string) *GetTemplateResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetTemplateResponseBody
	GetDescription() *string
	SetInterface(v string) *GetTemplateResponseBody
	GetInterface() *string
	SetOwnerId(v string) *GetTemplateResponseBody
	GetOwnerId() *string
	SetPermissions(v []*GetTemplateResponseBodyPermissions) *GetTemplateResponseBody
	GetPermissions() []*GetTemplateResponseBodyPermissions
	SetRegionId(v string) *GetTemplateResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetTemplateResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetTemplateResponseBody
	GetResourceGroupId() *string
	SetShareType(v string) *GetTemplateResponseBody
	GetShareType() *string
	SetStackGroupName(v string) *GetTemplateResponseBody
	GetStackGroupName() *string
	SetStackId(v string) *GetTemplateResponseBody
	GetStackId() *string
	SetTags(v []*GetTemplateResponseBodyTags) *GetTemplateResponseBody
	GetTags() []*GetTemplateResponseBodyTags
	SetTemplateARN(v string) *GetTemplateResponseBody
	GetTemplateARN() *string
	SetTemplateBody(v string) *GetTemplateResponseBody
	GetTemplateBody() *string
	SetTemplateId(v string) *GetTemplateResponseBody
	GetTemplateId() *string
	SetTemplateName(v string) *GetTemplateResponseBody
	GetTemplateName() *string
	SetTemplateVersion(v string) *GetTemplateResponseBody
	GetTemplateVersion() *string
	SetUpdateTime(v string) *GetTemplateResponseBody
	GetUpdateTime() *string
}

type GetTemplateResponseBody struct {
	// Supplementary information for the public template.
	//
	// example:
	//
	// {"DeploymentDuration":null,"Title":"Self-Built_ElasticSearch_Snapshot_Saved_To_OSS","Labels":{"ResourceTypes":["ALIYUN::ECS::Instance","ALIYUN::ECS::SecurityGroup","ALIYUN::ECS::VPC","ALIYUN::ECS::VSwitch","ALIYUN::OSS::Bucket","ALIYUN::ROS::WaitCondition","ALIYUN::ROS::WaitConditionHandle"],"DeployTypes":["ROS"],"ApplicationScenes":["其他"]},"Provider":"ROS","Categories":["Solution"]}
	AdditionalInfo map[string]interface{} `json:"AdditionalInfo,omitempty" xml:"AdditionalInfo,omitempty"`
	// The ID of the change set. This parameter is returned only if you specify ChangeSetId.
	//
	// example:
	//
	// e85abe0c-6528-43fb-ae93-fdf8de22****
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The time when the template was created. This parameter is returned only if you specify TemplateId.
	//
	// > - If you specify TemplateVersion, the creation time of the template whose version is specified by TemplateVersion is returned.
	//
	// > - If you do not specify TemplateVersion, the creation time of the template whose version is the default version is returned.
	//
	// example:
	//
	// 2020-11-18T08:49:26
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the template. This parameter is returned only if you specify TemplateId.
	//
	// example:
	//
	// ROS template for create ECS instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The description of the web UI in the ROS console.
	//
	// example:
	//
	// {}
	Interface *string `json:"Interface,omitempty" xml:"Interface,omitempty"`
	// The ID of the Alibaba Cloud account to which the template belongs. This parameter is returned only if you specify TemplateId.
	//
	// example:
	//
	// 151266687691****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Details of the sharing status of the template. This parameter is returned only if you specify TemplateId and set IncludePermission to Enabled.
	//
	// > - If TemplateVersion is not specified or does not take effect, the details of the sharing status of the template whose version is the default version is returned.
	//
	// > - If TemplateVersion is specified and takes effect, the details of the sharing status of the template whose version is specified by TemplateVersion is returned.
	Permissions []*GetTemplateResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The region ID of the stack or stack group that uses the template. This parameter is returned only if you specify StackId, ChangeSetId, or StackGroupName.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The sharing type of the template. This parameter is returned only if you specify TemplateId.
	//
	// Valid values:
	//
	// 	- Private: The template belongs to the template owner.
	//
	// 	- Shared: The template is shared by other users.
	//
	// example:
	//
	// Private
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The name of the stack group. This parameter is returned only if you specify StackGroupName.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The ID of the stack. This parameter is returned only if you specify StackId.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The tags of the template.
	Tags []*GetTemplateResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The Alibaba Cloud Resource Name (ARN) of the template. This parameter is returned only if you specify TemplateId.
	//
	// example:
	//
	// acs:ros:*:151266687691****:template/a52f81be-496f-4e1c-a286-8852ab54****
	TemplateARN *string `json:"TemplateARN,omitempty" xml:"TemplateARN,omitempty"`
	// The content of the template.
	//
	// example:
	//
	// {"ROSTemplateFormatVersion": "2015-09-01"}
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter is returned only if you specify TemplateId.
	//
	// If the template is a shared template, the value of this parameter is the same as the value of TemplateARN.
	//
	// example:
	//
	// a52f81be-496f-4e1c-a286-8852ab54****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template. This parameter is returned only if you specify TemplateId.
	//
	// > -   If you specify TemplateVersion, the name of the template whose version is specified by TemplateVersion is returned.
	//
	// > -  If you not specify TemplateVersion, the name of the template whose version is the default version is returned.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version of the template. This parameter is returned only if you specify TemplateId.\\
	//
	// If TemplateVersion is not specified or does not take effect, the default version is used.
	//
	// If the template is a shared template, this parameter is returned only if you set VersionOption to AllVersions.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The time when the template was last updated. This parameter is returned only if you specify TemplateId.
	//
	// > - If you specify TemplateVersion, the last update time of the template whose version is specified by TemplateVersion is returned.
	//
	// > - If you do not specify TemplateVersion, the last update time of the template whose version is the default version is returned.
	//
	// example:
	//
	// 2020-12-07T06:11:48
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) GetAdditionalInfo() map[string]interface{} {
	return s.AdditionalInfo
}

func (s *GetTemplateResponseBody) GetChangeSetId() *string {
	return s.ChangeSetId
}

func (s *GetTemplateResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTemplateResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetTemplateResponseBody) GetInterface() *string {
	return s.Interface
}

func (s *GetTemplateResponseBody) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetTemplateResponseBody) GetPermissions() []*GetTemplateResponseBodyPermissions {
	return s.Permissions
}

func (s *GetTemplateResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetTemplateResponseBody) GetShareType() *string {
	return s.ShareType
}

func (s *GetTemplateResponseBody) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *GetTemplateResponseBody) GetStackId() *string {
	return s.StackId
}

func (s *GetTemplateResponseBody) GetTags() []*GetTemplateResponseBodyTags {
	return s.Tags
}

func (s *GetTemplateResponseBody) GetTemplateARN() *string {
	return s.TemplateARN
}

func (s *GetTemplateResponseBody) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *GetTemplateResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateResponseBody) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetTemplateResponseBody) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetTemplateResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetTemplateResponseBody) SetAdditionalInfo(v map[string]interface{}) *GetTemplateResponseBody {
	s.AdditionalInfo = v
	return s
}

func (s *GetTemplateResponseBody) SetChangeSetId(v string) *GetTemplateResponseBody {
	s.ChangeSetId = &v
	return s
}

func (s *GetTemplateResponseBody) SetCreateTime(v string) *GetTemplateResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetTemplateResponseBody) SetDescription(v string) *GetTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *GetTemplateResponseBody) SetInterface(v string) *GetTemplateResponseBody {
	s.Interface = &v
	return s
}

func (s *GetTemplateResponseBody) SetOwnerId(v string) *GetTemplateResponseBody {
	s.OwnerId = &v
	return s
}

func (s *GetTemplateResponseBody) SetPermissions(v []*GetTemplateResponseBodyPermissions) *GetTemplateResponseBody {
	s.Permissions = v
	return s
}

func (s *GetTemplateResponseBody) SetRegionId(v string) *GetTemplateResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetResourceGroupId(v string) *GetTemplateResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTemplateResponseBody) SetShareType(v string) *GetTemplateResponseBody {
	s.ShareType = &v
	return s
}

func (s *GetTemplateResponseBody) SetStackGroupName(v string) *GetTemplateResponseBody {
	s.StackGroupName = &v
	return s
}

func (s *GetTemplateResponseBody) SetStackId(v string) *GetTemplateResponseBody {
	s.StackId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTags(v []*GetTemplateResponseBodyTags) *GetTemplateResponseBody {
	s.Tags = v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateARN(v string) *GetTemplateResponseBody {
	s.TemplateARN = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateBody(v string) *GetTemplateResponseBody {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateId(v string) *GetTemplateResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateName(v string) *GetTemplateResponseBody {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateVersion(v string) *GetTemplateResponseBody {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateResponseBody) SetUpdateTime(v string) *GetTemplateResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetTemplateResponseBody) Validate() error {
	if s.Permissions != nil {
		for _, item := range s.Permissions {
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

type GetTemplateResponseBodyPermissions struct {
	// The ID of the Alibaba Cloud account with which the template is shared.
	//
	// example:
	//
	// 142437958638****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The sharing option.
	//
	// The value ShareToAccounts indicates that the template is shared with one or more Alibaba Cloud accounts.
	//
	// example:
	//
	// ShareToAccounts
	ShareOption *string `json:"ShareOption,omitempty" xml:"ShareOption,omitempty"`
	// The service that is used for resource sharing. Valid values:
	//
	// - ROS: Resources are shared from ROS by using the ROS console or calling the ROS API.
	//
	// - ResourceDirectory: Resources are shared with accounts in a resource directory from Resource Management by using the resource sharing feature.
	//
	// > -  The number of accounts with which resources are shared from ROS is independent of the number of accounts with which resources are shared from the resource directory.
	//
	// > -  The shared resources from ROS cannot override or overwrite the shared resources from the resource directory.
	//
	// > -  The shared resources from the resource directory can overwrite the shared resources from ROS.
	//
	// example:
	//
	// ROS
	ShareSource *string `json:"ShareSource,omitempty" xml:"ShareSource,omitempty"`
	// The version of the shared template. This parameter is returned only if you set ShareOption to ShareToAccounts and set VersionOption to Specified or Current.
	//
	// Valid values: v1 to v100.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The version option for the shared template. This parameter is returned only if you set ShareOption to ShareToAccounts.
	//
	// Valid values:
	//
	// 	- AllVersions: All template versions are shared.
	//
	// 	- Latest: Only the latest template version is shared. When the version of the template is updated, Resource Orchestration Service (ROS) updates the shared version to the latest version.
	//
	// 	- Current: Only the latest template version is shared. When the version of the template is updated, ROS does not update the shared version.
	//
	// 	- Specified: Only the specified template version is shared.
	//
	// example:
	//
	// AllVersions
	VersionOption *string `json:"VersionOption,omitempty" xml:"VersionOption,omitempty"`
}

func (s GetTemplateResponseBodyPermissions) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyPermissions) GetAccountId() *string {
	return s.AccountId
}

func (s *GetTemplateResponseBodyPermissions) GetShareOption() *string {
	return s.ShareOption
}

func (s *GetTemplateResponseBodyPermissions) GetShareSource() *string {
	return s.ShareSource
}

func (s *GetTemplateResponseBodyPermissions) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetTemplateResponseBodyPermissions) GetVersionOption() *string {
	return s.VersionOption
}

func (s *GetTemplateResponseBodyPermissions) SetAccountId(v string) *GetTemplateResponseBodyPermissions {
	s.AccountId = &v
	return s
}

func (s *GetTemplateResponseBodyPermissions) SetShareOption(v string) *GetTemplateResponseBodyPermissions {
	s.ShareOption = &v
	return s
}

func (s *GetTemplateResponseBodyPermissions) SetShareSource(v string) *GetTemplateResponseBodyPermissions {
	s.ShareSource = &v
	return s
}

func (s *GetTemplateResponseBodyPermissions) SetTemplateVersion(v string) *GetTemplateResponseBodyPermissions {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateResponseBodyPermissions) SetVersionOption(v string) *GetTemplateResponseBodyPermissions {
	s.VersionOption = &v
	return s
}

func (s *GetTemplateResponseBodyPermissions) Validate() error {
	return dara.Validate(s)
}

type GetTemplateResponseBodyTags struct {
	// The tag key of the template.
	//
	// example:
	//
	// usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the template.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTemplateResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetTemplateResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetTemplateResponseBodyTags) SetKey(v string) *GetTemplateResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetTemplateResponseBodyTags) SetValue(v string) *GetTemplateResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetTemplateResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
