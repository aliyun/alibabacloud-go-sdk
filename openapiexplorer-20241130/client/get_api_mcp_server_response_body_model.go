// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalApiDescriptions(v []*GetApiMcpServerResponseBodyAdditionalApiDescriptions) *GetApiMcpServerResponseBody
	GetAdditionalApiDescriptions() []*GetApiMcpServerResponseBodyAdditionalApiDescriptions
	SetApiInfos(v []*GetApiMcpServerResponseBodyApiInfos) *GetApiMcpServerResponseBody
	GetApiInfos() []*GetApiMcpServerResponseBodyApiInfos
	SetApis(v []*GetApiMcpServerResponseBodyApis) *GetApiMcpServerResponseBody
	GetApis() []*GetApiMcpServerResponseBodyApis
	SetAssumeRoleExtraPolicy(v string) *GetApiMcpServerResponseBody
	GetAssumeRoleExtraPolicy() *string
	SetAssumeRoleName(v string) *GetApiMcpServerResponseBody
	GetAssumeRoleName() *string
	SetCreateTime(v string) *GetApiMcpServerResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetApiMcpServerResponseBody
	GetDescription() *string
	SetEnableAssumeRole(v bool) *GetApiMcpServerResponseBody
	GetEnableAssumeRole() *bool
	SetEnableCustomVpcWhitelist(v bool) *GetApiMcpServerResponseBody
	GetEnableCustomVpcWhitelist() *bool
	SetId(v string) *GetApiMcpServerResponseBody
	GetId() *string
	SetInstructions(v string) *GetApiMcpServerResponseBody
	GetInstructions() *string
	SetLanguage(v string) *GetApiMcpServerResponseBody
	GetLanguage() *string
	SetName(v string) *GetApiMcpServerResponseBody
	GetName() *string
	SetOauthClientId(v string) *GetApiMcpServerResponseBody
	GetOauthClientId() *string
	SetPrompts(v []*GetApiMcpServerResponseBodyPrompts) *GetApiMcpServerResponseBody
	GetPrompts() []*GetApiMcpServerResponseBodyPrompts
	SetPublicAccess(v string) *GetApiMcpServerResponseBody
	GetPublicAccess() *string
	SetRequestId(v string) *GetApiMcpServerResponseBody
	GetRequestId() *string
	SetRequiredRAMPolicy(v string) *GetApiMcpServerResponseBody
	GetRequiredRAMPolicy() *string
	SetSourceType(v string) *GetApiMcpServerResponseBody
	GetSourceType() *string
	SetSystemMcpServerInfo(v *GetApiMcpServerResponseBodySystemMcpServerInfo) *GetApiMcpServerResponseBody
	GetSystemMcpServerInfo() *GetApiMcpServerResponseBodySystemMcpServerInfo
	SetSystemTools(v []*string) *GetApiMcpServerResponseBody
	GetSystemTools() []*string
	SetTerraformTools(v []*GetApiMcpServerResponseBodyTerraformTools) *GetApiMcpServerResponseBody
	GetTerraformTools() []*GetApiMcpServerResponseBodyTerraformTools
	SetUpdateTime(v string) *GetApiMcpServerResponseBody
	GetUpdateTime() *string
	SetUrls(v *GetApiMcpServerResponseBodyUrls) *GetApiMcpServerResponseBody
	GetUrls() *GetApiMcpServerResponseBodyUrls
	SetVpcWhitelists(v []*string) *GetApiMcpServerResponseBody
	GetVpcWhitelists() []*string
}

type GetApiMcpServerResponseBody struct {
	AdditionalApiDescriptions []*GetApiMcpServerResponseBodyAdditionalApiDescriptions `json:"additionalApiDescriptions,omitempty" xml:"additionalApiDescriptions,omitempty" type:"Repeated"`
	ApiInfos                  []*GetApiMcpServerResponseBodyApiInfos                  `json:"apiInfos,omitempty" xml:"apiInfos,omitempty" type:"Repeated"`
	Apis                      []*GetApiMcpServerResponseBodyApis                      `json:"apis,omitempty" xml:"apis,omitempty" type:"Repeated"`
	// example:
	//
	// {
	//
	//   "Version": "1",
	//
	//   "Statement": [
	//
	//     {
	//
	//       "Effect": "Allow",
	//
	//       "Action": [
	//
	//         "ecs:Describe*",
	//
	//         "vpc:Describe*",
	//
	//         "vpc:List*"
	//
	//       ],
	//
	//       "Resource": "*"
	//
	//     }
	//
	//   ]
	//
	// }
	AssumeRoleExtraPolicy *string `json:"assumeRoleExtraPolicy,omitempty" xml:"assumeRoleExtraPolicy,omitempty"`
	// example:
	//
	// test
	AssumeRoleName *string `json:"assumeRoleName,omitempty" xml:"assumeRoleName,omitempty"`
	// example:
	//
	// 2025-02-07T02:17:46Z
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	EnableAssumeRole         *bool `json:"enableAssumeRole,omitempty" xml:"enableAssumeRole,omitempty"`
	EnableCustomVpcWhitelist *bool `json:"enableCustomVpcWhitelist,omitempty" xml:"enableCustomVpcWhitelist,omitempty"`
	// example:
	//
	// v6ZZ7ftCzEILW***
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// test
	Instructions *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	// example:
	//
	// ZH_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// mcp-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 403*************370
	OauthClientId *string                               `json:"oauthClientId,omitempty" xml:"oauthClientId,omitempty"`
	Prompts       []*GetApiMcpServerResponseBodyPrompts `json:"prompts,omitempty" xml:"prompts,omitempty" type:"Repeated"`
	PublicAccess  *string                               `json:"publicAccess,omitempty" xml:"publicAccess,omitempty"`
	// example:
	//
	// 9BFC4AC1-6BE4-5405-BDEC-CA288D404812
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// {
	//
	//   "Version": "1",
	//
	//   "Statement": [
	//
	//     {
	//
	//       "Effect": "Allow",
	//
	//       "Action": [
	//
	//         "ecs:Describe*",
	//
	//         "vpc:Describe*",
	//
	//         "vpc:List*"
	//
	//       ],
	//
	//       "Resource": "*"
	//
	//     }
	//
	//   ]
	//
	// }
	RequiredRAMPolicy *string `json:"requiredRAMPolicy,omitempty" xml:"requiredRAMPolicy,omitempty"`
	// example:
	//
	// system
	SourceType          *string                                         `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	SystemMcpServerInfo *GetApiMcpServerResponseBodySystemMcpServerInfo `json:"systemMcpServerInfo,omitempty" xml:"systemMcpServerInfo,omitempty" type:"Struct"`
	SystemTools         []*string                                       `json:"systemTools,omitempty" xml:"systemTools,omitempty" type:"Repeated"`
	TerraformTools      []*GetApiMcpServerResponseBodyTerraformTools    `json:"terraformTools,omitempty" xml:"terraformTools,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-02-05T02:26:04Z
	UpdateTime    *string                          `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Urls          *GetApiMcpServerResponseBodyUrls `json:"urls,omitempty" xml:"urls,omitempty" type:"Struct"`
	VpcWhitelists []*string                        `json:"vpcWhitelists,omitempty" xml:"vpcWhitelists,omitempty" type:"Repeated"`
}

func (s GetApiMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerResponseBody) GetAdditionalApiDescriptions() []*GetApiMcpServerResponseBodyAdditionalApiDescriptions {
	return s.AdditionalApiDescriptions
}

func (s *GetApiMcpServerResponseBody) GetApiInfos() []*GetApiMcpServerResponseBodyApiInfos {
	return s.ApiInfos
}

func (s *GetApiMcpServerResponseBody) GetApis() []*GetApiMcpServerResponseBodyApis {
	return s.Apis
}

func (s *GetApiMcpServerResponseBody) GetAssumeRoleExtraPolicy() *string {
	return s.AssumeRoleExtraPolicy
}

func (s *GetApiMcpServerResponseBody) GetAssumeRoleName() *string {
	return s.AssumeRoleName
}

func (s *GetApiMcpServerResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetApiMcpServerResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetApiMcpServerResponseBody) GetEnableAssumeRole() *bool {
	return s.EnableAssumeRole
}

func (s *GetApiMcpServerResponseBody) GetEnableCustomVpcWhitelist() *bool {
	return s.EnableCustomVpcWhitelist
}

func (s *GetApiMcpServerResponseBody) GetId() *string {
	return s.Id
}

func (s *GetApiMcpServerResponseBody) GetInstructions() *string {
	return s.Instructions
}

func (s *GetApiMcpServerResponseBody) GetLanguage() *string {
	return s.Language
}

func (s *GetApiMcpServerResponseBody) GetName() *string {
	return s.Name
}

func (s *GetApiMcpServerResponseBody) GetOauthClientId() *string {
	return s.OauthClientId
}

func (s *GetApiMcpServerResponseBody) GetPrompts() []*GetApiMcpServerResponseBodyPrompts {
	return s.Prompts
}

func (s *GetApiMcpServerResponseBody) GetPublicAccess() *string {
	return s.PublicAccess
}

func (s *GetApiMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApiMcpServerResponseBody) GetRequiredRAMPolicy() *string {
	return s.RequiredRAMPolicy
}

func (s *GetApiMcpServerResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *GetApiMcpServerResponseBody) GetSystemMcpServerInfo() *GetApiMcpServerResponseBodySystemMcpServerInfo {
	return s.SystemMcpServerInfo
}

func (s *GetApiMcpServerResponseBody) GetSystemTools() []*string {
	return s.SystemTools
}

func (s *GetApiMcpServerResponseBody) GetTerraformTools() []*GetApiMcpServerResponseBodyTerraformTools {
	return s.TerraformTools
}

func (s *GetApiMcpServerResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetApiMcpServerResponseBody) GetUrls() *GetApiMcpServerResponseBodyUrls {
	return s.Urls
}

func (s *GetApiMcpServerResponseBody) GetVpcWhitelists() []*string {
	return s.VpcWhitelists
}

func (s *GetApiMcpServerResponseBody) SetAdditionalApiDescriptions(v []*GetApiMcpServerResponseBodyAdditionalApiDescriptions) *GetApiMcpServerResponseBody {
	s.AdditionalApiDescriptions = v
	return s
}

func (s *GetApiMcpServerResponseBody) SetApiInfos(v []*GetApiMcpServerResponseBodyApiInfos) *GetApiMcpServerResponseBody {
	s.ApiInfos = v
	return s
}

func (s *GetApiMcpServerResponseBody) SetApis(v []*GetApiMcpServerResponseBodyApis) *GetApiMcpServerResponseBody {
	s.Apis = v
	return s
}

func (s *GetApiMcpServerResponseBody) SetAssumeRoleExtraPolicy(v string) *GetApiMcpServerResponseBody {
	s.AssumeRoleExtraPolicy = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetAssumeRoleName(v string) *GetApiMcpServerResponseBody {
	s.AssumeRoleName = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetCreateTime(v string) *GetApiMcpServerResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetDescription(v string) *GetApiMcpServerResponseBody {
	s.Description = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetEnableAssumeRole(v bool) *GetApiMcpServerResponseBody {
	s.EnableAssumeRole = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetEnableCustomVpcWhitelist(v bool) *GetApiMcpServerResponseBody {
	s.EnableCustomVpcWhitelist = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetId(v string) *GetApiMcpServerResponseBody {
	s.Id = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetInstructions(v string) *GetApiMcpServerResponseBody {
	s.Instructions = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetLanguage(v string) *GetApiMcpServerResponseBody {
	s.Language = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetName(v string) *GetApiMcpServerResponseBody {
	s.Name = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetOauthClientId(v string) *GetApiMcpServerResponseBody {
	s.OauthClientId = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetPrompts(v []*GetApiMcpServerResponseBodyPrompts) *GetApiMcpServerResponseBody {
	s.Prompts = v
	return s
}

func (s *GetApiMcpServerResponseBody) SetPublicAccess(v string) *GetApiMcpServerResponseBody {
	s.PublicAccess = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetRequestId(v string) *GetApiMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetRequiredRAMPolicy(v string) *GetApiMcpServerResponseBody {
	s.RequiredRAMPolicy = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetSourceType(v string) *GetApiMcpServerResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetSystemMcpServerInfo(v *GetApiMcpServerResponseBodySystemMcpServerInfo) *GetApiMcpServerResponseBody {
	s.SystemMcpServerInfo = v
	return s
}

func (s *GetApiMcpServerResponseBody) SetSystemTools(v []*string) *GetApiMcpServerResponseBody {
	s.SystemTools = v
	return s
}

func (s *GetApiMcpServerResponseBody) SetTerraformTools(v []*GetApiMcpServerResponseBodyTerraformTools) *GetApiMcpServerResponseBody {
	s.TerraformTools = v
	return s
}

func (s *GetApiMcpServerResponseBody) SetUpdateTime(v string) *GetApiMcpServerResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetApiMcpServerResponseBody) SetUrls(v *GetApiMcpServerResponseBodyUrls) *GetApiMcpServerResponseBody {
	s.Urls = v
	return s
}

func (s *GetApiMcpServerResponseBody) SetVpcWhitelists(v []*string) *GetApiMcpServerResponseBody {
	s.VpcWhitelists = v
	return s
}

func (s *GetApiMcpServerResponseBody) Validate() error {
	if s.AdditionalApiDescriptions != nil {
		for _, item := range s.AdditionalApiDescriptions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ApiInfos != nil {
		for _, item := range s.ApiInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Apis != nil {
		for _, item := range s.Apis {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Prompts != nil {
		for _, item := range s.Prompts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SystemMcpServerInfo != nil {
		if err := s.SystemMcpServerInfo.Validate(); err != nil {
			return err
		}
	}
	if s.TerraformTools != nil {
		for _, item := range s.TerraformTools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Urls != nil {
		if err := s.Urls.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApiMcpServerResponseBodyAdditionalApiDescriptions struct {
	// example:
	//
	// DescribeRegions
	ApiName         *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	ApiOverrideJson *string `json:"apiOverrideJson,omitempty" xml:"apiOverrideJson,omitempty"`
	// example:
	//
	// 2014-05-26
	ApiVersion      *string                                                                `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	ConstParameters []*GetApiMcpServerResponseBodyAdditionalApiDescriptionsConstParameters `json:"constParameters,omitempty" xml:"constParameters,omitempty" type:"Repeated"`
	// example:
	//
	// true
	EnableOutputSchema *bool `json:"enableOutputSchema,omitempty" xml:"enableOutputSchema,omitempty"`
	// example:
	//
	// true
	ExecuteCliCommand *bool `json:"executeCliCommand,omitempty" xml:"executeCliCommand,omitempty"`
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s GetApiMcpServerResponseBodyAdditionalApiDescriptions) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerResponseBodyAdditionalApiDescriptions) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) GetApiName() *string {
	return s.ApiName
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) GetApiOverrideJson() *string {
	return s.ApiOverrideJson
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) GetConstParameters() []*GetApiMcpServerResponseBodyAdditionalApiDescriptionsConstParameters {
	return s.ConstParameters
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) GetEnableOutputSchema() *bool {
	return s.EnableOutputSchema
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) GetExecuteCliCommand() *bool {
	return s.ExecuteCliCommand
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) GetProduct() *string {
	return s.Product
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) SetApiName(v string) *GetApiMcpServerResponseBodyAdditionalApiDescriptions {
	s.ApiName = &v
	return s
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) SetApiOverrideJson(v string) *GetApiMcpServerResponseBodyAdditionalApiDescriptions {
	s.ApiOverrideJson = &v
	return s
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) SetApiVersion(v string) *GetApiMcpServerResponseBodyAdditionalApiDescriptions {
	s.ApiVersion = &v
	return s
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) SetConstParameters(v []*GetApiMcpServerResponseBodyAdditionalApiDescriptionsConstParameters) *GetApiMcpServerResponseBodyAdditionalApiDescriptions {
	s.ConstParameters = v
	return s
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) SetEnableOutputSchema(v bool) *GetApiMcpServerResponseBodyAdditionalApiDescriptions {
	s.EnableOutputSchema = &v
	return s
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) SetExecuteCliCommand(v bool) *GetApiMcpServerResponseBodyAdditionalApiDescriptions {
	s.ExecuteCliCommand = &v
	return s
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) SetProduct(v string) *GetApiMcpServerResponseBodyAdditionalApiDescriptions {
	s.Product = &v
	return s
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptions) Validate() error {
	if s.ConstParameters != nil {
		for _, item := range s.ConstParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApiMcpServerResponseBodyAdditionalApiDescriptionsConstParameters struct {
	// example:
	//
	// InstanceId
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 1234
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetApiMcpServerResponseBodyAdditionalApiDescriptionsConstParameters) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerResponseBodyAdditionalApiDescriptionsConstParameters) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptionsConstParameters) GetKey() *string {
	return s.Key
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptionsConstParameters) GetValue() interface{} {
	return s.Value
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptionsConstParameters) SetKey(v string) *GetApiMcpServerResponseBodyAdditionalApiDescriptionsConstParameters {
	s.Key = &v
	return s
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptionsConstParameters) SetValue(v interface{}) *GetApiMcpServerResponseBodyAdditionalApiDescriptionsConstParameters {
	s.Value = v
	return s
}

func (s *GetApiMcpServerResponseBodyAdditionalApiDescriptionsConstParameters) Validate() error {
	return dara.Validate(s)
}

type GetApiMcpServerResponseBodyApiInfos struct {
	// example:
	//
	// DescribeRegions
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// example:
	//
	// 2014-05-26
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s GetApiMcpServerResponseBodyApiInfos) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerResponseBodyApiInfos) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerResponseBodyApiInfos) GetApiName() *string {
	return s.ApiName
}

func (s *GetApiMcpServerResponseBodyApiInfos) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GetApiMcpServerResponseBodyApiInfos) GetProduct() *string {
	return s.Product
}

func (s *GetApiMcpServerResponseBodyApiInfos) SetApiName(v string) *GetApiMcpServerResponseBodyApiInfos {
	s.ApiName = &v
	return s
}

func (s *GetApiMcpServerResponseBodyApiInfos) SetApiVersion(v string) *GetApiMcpServerResponseBodyApiInfos {
	s.ApiVersion = &v
	return s
}

func (s *GetApiMcpServerResponseBodyApiInfos) SetProduct(v string) *GetApiMcpServerResponseBodyApiInfos {
	s.Product = &v
	return s
}

func (s *GetApiMcpServerResponseBodyApiInfos) Validate() error {
	return dara.Validate(s)
}

type GetApiMcpServerResponseBodyApis struct {
	// example:
	//
	// 2014-05-26
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// example:
	//
	// Ecs
	Product   *string   `json:"product,omitempty" xml:"product,omitempty"`
	Selectors []*string `json:"selectors,omitempty" xml:"selectors,omitempty" type:"Repeated"`
}

func (s GetApiMcpServerResponseBodyApis) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerResponseBodyApis) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerResponseBodyApis) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *GetApiMcpServerResponseBodyApis) GetProduct() *string {
	return s.Product
}

func (s *GetApiMcpServerResponseBodyApis) GetSelectors() []*string {
	return s.Selectors
}

func (s *GetApiMcpServerResponseBodyApis) SetApiVersion(v string) *GetApiMcpServerResponseBodyApis {
	s.ApiVersion = &v
	return s
}

func (s *GetApiMcpServerResponseBodyApis) SetProduct(v string) *GetApiMcpServerResponseBodyApis {
	s.Product = &v
	return s
}

func (s *GetApiMcpServerResponseBodyApis) SetSelectors(v []*string) *GetApiMcpServerResponseBodyApis {
	s.Selectors = v
	return s
}

func (s *GetApiMcpServerResponseBodyApis) Validate() error {
	return dara.Validate(s)
}

type GetApiMcpServerResponseBodyPrompts struct {
	Arguments []*GetApiMcpServerResponseBodyPromptsArguments `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	Content   *string                                        `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// prompt description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetApiMcpServerResponseBodyPrompts) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerResponseBodyPrompts) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerResponseBodyPrompts) GetArguments() []*GetApiMcpServerResponseBodyPromptsArguments {
	return s.Arguments
}

func (s *GetApiMcpServerResponseBodyPrompts) GetContent() *string {
	return s.Content
}

func (s *GetApiMcpServerResponseBodyPrompts) GetDescription() *string {
	return s.Description
}

func (s *GetApiMcpServerResponseBodyPrompts) GetName() *string {
	return s.Name
}

func (s *GetApiMcpServerResponseBodyPrompts) SetArguments(v []*GetApiMcpServerResponseBodyPromptsArguments) *GetApiMcpServerResponseBodyPrompts {
	s.Arguments = v
	return s
}

func (s *GetApiMcpServerResponseBodyPrompts) SetContent(v string) *GetApiMcpServerResponseBodyPrompts {
	s.Content = &v
	return s
}

func (s *GetApiMcpServerResponseBodyPrompts) SetDescription(v string) *GetApiMcpServerResponseBodyPrompts {
	s.Description = &v
	return s
}

func (s *GetApiMcpServerResponseBodyPrompts) SetName(v string) *GetApiMcpServerResponseBodyPrompts {
	s.Name = &v
	return s
}

func (s *GetApiMcpServerResponseBodyPrompts) Validate() error {
	if s.Arguments != nil {
		for _, item := range s.Arguments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetApiMcpServerResponseBodyPromptsArguments struct {
	// example:
	//
	// argument description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// test
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	Required *bool   `json:"required,omitempty" xml:"required,omitempty"`
}

func (s GetApiMcpServerResponseBodyPromptsArguments) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerResponseBodyPromptsArguments) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerResponseBodyPromptsArguments) GetDescription() *string {
	return s.Description
}

func (s *GetApiMcpServerResponseBodyPromptsArguments) GetName() *string {
	return s.Name
}

func (s *GetApiMcpServerResponseBodyPromptsArguments) GetRequired() *bool {
	return s.Required
}

func (s *GetApiMcpServerResponseBodyPromptsArguments) SetDescription(v string) *GetApiMcpServerResponseBodyPromptsArguments {
	s.Description = &v
	return s
}

func (s *GetApiMcpServerResponseBodyPromptsArguments) SetName(v string) *GetApiMcpServerResponseBodyPromptsArguments {
	s.Name = &v
	return s
}

func (s *GetApiMcpServerResponseBodyPromptsArguments) SetRequired(v bool) *GetApiMcpServerResponseBodyPromptsArguments {
	s.Required = &v
	return s
}

func (s *GetApiMcpServerResponseBodyPromptsArguments) Validate() error {
	return dara.Validate(s)
}

type GetApiMcpServerResponseBodySystemMcpServerInfo struct {
	// example:
	//
	// mcp-system
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s GetApiMcpServerResponseBodySystemMcpServerInfo) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerResponseBodySystemMcpServerInfo) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerResponseBodySystemMcpServerInfo) GetName() *string {
	return s.Name
}

func (s *GetApiMcpServerResponseBodySystemMcpServerInfo) GetProduct() *string {
	return s.Product
}

func (s *GetApiMcpServerResponseBodySystemMcpServerInfo) SetName(v string) *GetApiMcpServerResponseBodySystemMcpServerInfo {
	s.Name = &v
	return s
}

func (s *GetApiMcpServerResponseBodySystemMcpServerInfo) SetProduct(v string) *GetApiMcpServerResponseBodySystemMcpServerInfo {
	s.Product = &v
	return s
}

func (s *GetApiMcpServerResponseBodySystemMcpServerInfo) Validate() error {
	return dara.Validate(s)
}

type GetApiMcpServerResponseBodyTerraformTools struct {
	// example:
	//
	// true
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// example:
	//
	// variable "name" {
	//
	//   default = "terraform-example"
	//
	// }
	//
	// provider "alicloud" {
	//
	//   region = "cn-beijing"
	//
	// }
	//
	// resource "alicloud_vpc" "default" {
	//
	//   ipv6_isp    = "BGP"
	//
	//   description = "test"
	//
	//   cidr_block  = "10.0.0.0/8"
	//
	//   vpc_name    = var.name
	//
	//   enable_ipv6 = true
	//
	// }
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// Terraform Tool description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// ON_FAILURE
	DestroyPolicy *string `json:"destroyPolicy,omitempty" xml:"destroyPolicy,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetApiMcpServerResponseBodyTerraformTools) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerResponseBodyTerraformTools) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerResponseBodyTerraformTools) GetAsync() *bool {
	return s.Async
}

func (s *GetApiMcpServerResponseBodyTerraformTools) GetCode() *string {
	return s.Code
}

func (s *GetApiMcpServerResponseBodyTerraformTools) GetDescription() *string {
	return s.Description
}

func (s *GetApiMcpServerResponseBodyTerraformTools) GetDestroyPolicy() *string {
	return s.DestroyPolicy
}

func (s *GetApiMcpServerResponseBodyTerraformTools) GetName() *string {
	return s.Name
}

func (s *GetApiMcpServerResponseBodyTerraformTools) SetAsync(v bool) *GetApiMcpServerResponseBodyTerraformTools {
	s.Async = &v
	return s
}

func (s *GetApiMcpServerResponseBodyTerraformTools) SetCode(v string) *GetApiMcpServerResponseBodyTerraformTools {
	s.Code = &v
	return s
}

func (s *GetApiMcpServerResponseBodyTerraformTools) SetDescription(v string) *GetApiMcpServerResponseBodyTerraformTools {
	s.Description = &v
	return s
}

func (s *GetApiMcpServerResponseBodyTerraformTools) SetDestroyPolicy(v string) *GetApiMcpServerResponseBodyTerraformTools {
	s.DestroyPolicy = &v
	return s
}

func (s *GetApiMcpServerResponseBodyTerraformTools) SetName(v string) *GetApiMcpServerResponseBodyTerraformTools {
	s.Name = &v
	return s
}

func (s *GetApiMcpServerResponseBodyTerraformTools) Validate() error {
	return dara.Validate(s)
}

type GetApiMcpServerResponseBodyUrls struct {
	// example:
	//
	// https://mcpserverinner-pre.cn-zhangjiakou.aliyuncs.com/accounts/xxxx/custom/xxx/id/xxxx/mcp
	Mcp *string `json:"mcp,omitempty" xml:"mcp,omitempty"`
	// example:
	//
	// https://mcpserverinner-pre.cn-zhangjiakou.aliyuncs.com/accounts/xxxx/custom/xxx/id/xxxx/sse
	Sse    *string `json:"sse,omitempty" xml:"sse,omitempty"`
	VpcMcp *string `json:"vpcMcp,omitempty" xml:"vpcMcp,omitempty"`
	VpcSse *string `json:"vpcSse,omitempty" xml:"vpcSse,omitempty"`
}

func (s GetApiMcpServerResponseBodyUrls) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerResponseBodyUrls) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerResponseBodyUrls) GetMcp() *string {
	return s.Mcp
}

func (s *GetApiMcpServerResponseBodyUrls) GetSse() *string {
	return s.Sse
}

func (s *GetApiMcpServerResponseBodyUrls) GetVpcMcp() *string {
	return s.VpcMcp
}

func (s *GetApiMcpServerResponseBodyUrls) GetVpcSse() *string {
	return s.VpcSse
}

func (s *GetApiMcpServerResponseBodyUrls) SetMcp(v string) *GetApiMcpServerResponseBodyUrls {
	s.Mcp = &v
	return s
}

func (s *GetApiMcpServerResponseBodyUrls) SetSse(v string) *GetApiMcpServerResponseBodyUrls {
	s.Sse = &v
	return s
}

func (s *GetApiMcpServerResponseBodyUrls) SetVpcMcp(v string) *GetApiMcpServerResponseBodyUrls {
	s.VpcMcp = &v
	return s
}

func (s *GetApiMcpServerResponseBodyUrls) SetVpcSse(v string) *GetApiMcpServerResponseBodyUrls {
	s.VpcSse = &v
	return s
}

func (s *GetApiMcpServerResponseBodyUrls) Validate() error {
	return dara.Validate(s)
}
