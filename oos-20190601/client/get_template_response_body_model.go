// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *GetTemplateResponseBody
	GetContent() *string
	SetRequestId(v string) *GetTemplateResponseBody
	GetRequestId() *string
	SetTemplate(v *GetTemplateResponseBodyTemplate) *GetTemplateResponseBody
	GetTemplate() *GetTemplateResponseBodyTemplate
}

type GetTemplateResponseBody struct {
	// The content of the template.
	//
	// example:
	//
	// "FormatVersion: OOS-2019-06-01\\nDescription:\\n  en:  Creates an ECS image\\n  zh-cn: 创建一个ECS镜像\\n  name-en: Create Image\\n  name-zh-cn: 创建镜像\\n  categories:\\n    - image_manage\\n    - application_manage\\nParameters:\\n  regionId:\\n    Type: String\\n    Label:\\n      en: RegionId\\n      zh-cn: 地域ID\\n    AssociationProperty: RegionId\\n    Default: \\"{{ ACS::RegionId }}\\"\\n  instanceId:\\n    Label:\\n      en: InstanceId\\n      zh-cn: ECS实例ID\\n    Type: String\\n    AssociationProperty: ALIYUN::ECS::Instance::InstanceId\\n    AssociationPropertyMetadata:\\n      RegionId: regionId\\n  imageName:\\n    Label:\\n      en: ImageName\\n      zh-cn: 新镜像的名称\\n    Type: String\\n    Description:\\n      en: <p class=\\"p\\"	Note:</p> <ul class=\\"ul\\"> <li class=\\"li\\">Length is 2~128 English or Chinese characters</li> <li class=\\"li\\"><font color=\\"red\\">must start with big or small letters or Chinese, not http:// and https://. </font></li> <li class=\\"li\\">Can contain numbers, colons (:), underscores (_), or dashes (-). </li> </ul>\\n      zh-cn: <p class=\\"p\\">注意：</p> <ul class=\\"ul\\"> <li class=\\"li\\">长度为2~128个英文或中文字符</li> <li class=\\"li\\"><font color=\\"red\\">必须以大小字母或中文开头，不能以http://和https://开头。</font></li> <li class=\\"li\\">可以包含数字、半角冒号（:）、下划线（_）或者短划线（-）。</li> </ul>\\n  tags:\\n    Label:\\n      en: Tags\\n      zh-cn: 镜像标签\\n    Type: Json\\n    AssociationProperty: Tags\\n    AssociationPropertyMetadata:\\n      ShowSystem: false\\n    Default: []\\n  OOSAssumeRole:\\n    Label:\\n      en: OOSAssumeRole\\n      zh-cn: OOS扮演的RAM角色\\n    Type: String\\n    Default: OOSServiceRole\\nRamRole: \\"{{ OOSAssumeRole }}\\"\\nTasks:\\n- Name: createImage\\n  Action: ACS::ECS::CreateImage\\n  Description:\\n    en: Create new image with the specified image name and instance ID\\n    zh-cn: 通过指定实例ID和镜像名称创建新的镜像\\n  Properties:\\n    regionId: \\"{{ regionId }}\\"\\n    imageName: \\"{{ imageName }}__on_{{ ACS::ExecutionId }}_at_{{ Acs::CurrentDate }}\\"\\n    instanceId: \\"{{ instanceId }}\\"\\n    tags: \\"{{tags}}\\"\\n  Outputs:\\n    imageId:\\n      ValueSelector: imageId\\n      Type: String\\nOutputs:\\n  imageId:\\n    Type: String\\n    Value: \\"{{ createImage.imageId }}\\"\\nMetadata:\\n  ALIYUN::OOS::Interface:\\n    ParameterGroups:\\n      - Parameters:\\n          - regionId\\n          - instanceId\\n        Label:\\n          default:\\n            zh-cn: 选择实例\\n            en: Select Ecs Instances\\n      - Parameters:\\n          - imageName\\n          - tags\\n        Label:\\n          default:\\n            zh-cn: 镜像设置\\n            en: Image Configure\\n      - Parameters:\\n          - OOSAssumeRole\\n        Label:\\n          default:\\n            zh-cn: 高级选项\\n            en: Control Options"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5BBE2663-A18E-5261-9BBB-F4832F5294D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metadata of the template.
	Template *GetTemplateResponseBodyTemplate `json:"Template,omitempty" xml:"Template,omitempty" type:"Struct"`
}

func (s GetTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) GetContent() *string {
	return s.Content
}

func (s *GetTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateResponseBody) GetTemplate() *GetTemplateResponseBodyTemplate {
	return s.Template
}

func (s *GetTemplateResponseBody) SetContent(v string) *GetTemplateResponseBody {
	s.Content = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplate(v *GetTemplateResponseBodyTemplate) *GetTemplateResponseBody {
	s.Template = v
	return s
}

func (s *GetTemplateResponseBody) Validate() error {
	if s.Template != nil {
		if err := s.Template.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTemplateResponseBodyTemplate struct {
	// The creator of the template.
	//
	// example:
	//
	// ACS
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2019-05-16T10:26:14Z
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// The description of the template.
	//
	// example:
	//
	// "{\\"en\\": \\"Creates an ECS image\\", \\"zh-cn\\": \\"创建一个ECS镜像\\", \\"name-en\\": \\"Create Image\\", \\"name-zh-cn\\": \\"创建镜像\\", \\"categories\\": [\\"image_manage\\", \\"application_manage\\"]}"
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the template was configured with a trigger.
	//
	// example:
	//
	// false
	HasTrigger *bool `json:"HasTrigger,omitempty" xml:"HasTrigger,omitempty"`
	// The SHA-256 value of the template content.
	//
	// example:
	//
	// 40fb5e3e08ef6c8a499ff7cd8441194f518028ad08338a84cb70c023a64576f1
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxsn4m4******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The share type of the template. The share type of a user-created template is **Private**.
	//
	// example:
	//
	// Public
	ShareType      *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	SharedAccounts *string `json:"SharedAccounts,omitempty" xml:"SharedAccounts,omitempty"`
	// The tag keys and values. The number of key-value pairs ranges from 1 to 20.
	//
	// example:
	//
	// {"k1":"k2","k2":"v2"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The format of the template. The system automatically determines whether the format is JSON or YAML.
	//
	// example:
	//
	// YAML
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// t-4bdb1745c171401883a2
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// ACS-ECS-CreateImage
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the template.
	//
	// example:
	//
	// Automation
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The version of the template. The name of the version consists of the letter v and a number. The number starts from 1.
	//
	// example:
	//
	// v15
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The user who last updated the template.
	//
	// example:
	//
	// ACS
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The time when the template was last updated.
	//
	// example:
	//
	// 2022-04-26T08:37:07Z
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
	// The name of the version of the template.
	//
	// example:
	//
	// version15
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetTemplateResponseBodyTemplate) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBodyTemplate) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyTemplate) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *GetTemplateResponseBodyTemplate) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *GetTemplateResponseBodyTemplate) GetDescription() *string {
	return s.Description
}

func (s *GetTemplateResponseBodyTemplate) GetHasTrigger() *bool {
	return s.HasTrigger
}

func (s *GetTemplateResponseBodyTemplate) GetHash() *string {
	return s.Hash
}

func (s *GetTemplateResponseBodyTemplate) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetTemplateResponseBodyTemplate) GetShareType() *string {
	return s.ShareType
}

func (s *GetTemplateResponseBodyTemplate) GetSharedAccounts() *string {
	return s.SharedAccounts
}

func (s *GetTemplateResponseBodyTemplate) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *GetTemplateResponseBodyTemplate) GetTemplateFormat() *string {
	return s.TemplateFormat
}

func (s *GetTemplateResponseBodyTemplate) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateResponseBodyTemplate) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetTemplateResponseBodyTemplate) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetTemplateResponseBodyTemplate) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetTemplateResponseBodyTemplate) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *GetTemplateResponseBodyTemplate) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *GetTemplateResponseBodyTemplate) GetVersionName() *string {
	return s.VersionName
}

func (s *GetTemplateResponseBodyTemplate) SetCreatedBy(v string) *GetTemplateResponseBodyTemplate {
	s.CreatedBy = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetCreatedDate(v string) *GetTemplateResponseBodyTemplate {
	s.CreatedDate = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetDescription(v string) *GetTemplateResponseBodyTemplate {
	s.Description = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetHasTrigger(v bool) *GetTemplateResponseBodyTemplate {
	s.HasTrigger = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetHash(v string) *GetTemplateResponseBodyTemplate {
	s.Hash = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetResourceGroupId(v string) *GetTemplateResponseBodyTemplate {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetShareType(v string) *GetTemplateResponseBodyTemplate {
	s.ShareType = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetSharedAccounts(v string) *GetTemplateResponseBodyTemplate {
	s.SharedAccounts = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTags(v map[string]interface{}) *GetTemplateResponseBodyTemplate {
	s.Tags = v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateFormat(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateFormat = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateId(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateName(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateName = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateType(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateType = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetTemplateVersion(v string) *GetTemplateResponseBodyTemplate {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetUpdatedBy(v string) *GetTemplateResponseBodyTemplate {
	s.UpdatedBy = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetUpdatedDate(v string) *GetTemplateResponseBodyTemplate {
	s.UpdatedDate = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) SetVersionName(v string) *GetTemplateResponseBodyTemplate {
	s.VersionName = &v
	return s
}

func (s *GetTemplateResponseBodyTemplate) Validate() error {
	return dara.Validate(s)
}
