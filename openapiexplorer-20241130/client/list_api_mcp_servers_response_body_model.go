// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiMcpServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiMcpServers(v []*ListApiMcpServersResponseBodyApiMcpServers) *ListApiMcpServersResponseBody
	GetApiMcpServers() []*ListApiMcpServersResponseBodyApiMcpServers
	SetMaxResults(v int32) *ListApiMcpServersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApiMcpServersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApiMcpServersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListApiMcpServersResponseBody
	GetTotalCount() *int32
}

type ListApiMcpServersResponseBody struct {
	ApiMcpServers []*ListApiMcpServersResponseBodyApiMcpServers `json:"apiMcpServers,omitempty" xml:"apiMcpServers,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// AAAAAZjtYxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 9BFC4AC1-6BE4-5405-BDEC-CA288D404812
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListApiMcpServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServersResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiMcpServersResponseBody) GetApiMcpServers() []*ListApiMcpServersResponseBodyApiMcpServers {
	return s.ApiMcpServers
}

func (s *ListApiMcpServersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApiMcpServersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApiMcpServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiMcpServersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApiMcpServersResponseBody) SetApiMcpServers(v []*ListApiMcpServersResponseBodyApiMcpServers) *ListApiMcpServersResponseBody {
	s.ApiMcpServers = v
	return s
}

func (s *ListApiMcpServersResponseBody) SetMaxResults(v int32) *ListApiMcpServersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApiMcpServersResponseBody) SetNextToken(v string) *ListApiMcpServersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApiMcpServersResponseBody) SetRequestId(v string) *ListApiMcpServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiMcpServersResponseBody) SetTotalCount(v int32) *ListApiMcpServersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApiMcpServersResponseBody) Validate() error {
	if s.ApiMcpServers != nil {
		for _, item := range s.ApiMcpServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApiMcpServersResponseBodyApiMcpServers struct {
	AdditionalApiDescriptions []*ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions `json:"additionalApiDescriptions,omitempty" xml:"additionalApiDescriptions,omitempty" type:"Repeated"`
	Apis                      []*ListApiMcpServersResponseBodyApiMcpServersApis                      `json:"apis,omitempty" xml:"apis,omitempty" type:"Repeated"`
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
	// 2024-12-10T03:20:21Z
	CreateTime  *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// true
	EnableAssumeRole *bool `json:"enableAssumeRole,omitempty" xml:"enableAssumeRole,omitempty"`
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
	OauthClientId *string                                              `json:"oauthClientId,omitempty" xml:"oauthClientId,omitempty"`
	Prompts       []*ListApiMcpServersResponseBodyApiMcpServersPrompts `json:"prompts,omitempty" xml:"prompts,omitempty" type:"Repeated"`
	// example:
	//
	// system
	SourceType          *string                                                        `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	SystemMcpServerInfo *ListApiMcpServersResponseBodyApiMcpServersSystemMcpServerInfo `json:"systemMcpServerInfo,omitempty" xml:"systemMcpServerInfo,omitempty" type:"Struct"`
	SystemTools         []*string                                                      `json:"systemTools,omitempty" xml:"systemTools,omitempty" type:"Repeated"`
	TerraformTools      []*ListApiMcpServersResponseBodyApiMcpServersTerraformTools    `json:"terraformTools,omitempty" xml:"terraformTools,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-01-10T02:11:43Z
	UpdateTime *string                                         `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Urls       *ListApiMcpServersResponseBodyApiMcpServersUrls `json:"urls,omitempty" xml:"urls,omitempty" type:"Struct"`
}

func (s ListApiMcpServersResponseBodyApiMcpServers) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServersResponseBodyApiMcpServers) GoString() string {
	return s.String()
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetAdditionalApiDescriptions() []*ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions {
	return s.AdditionalApiDescriptions
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetApis() []*ListApiMcpServersResponseBodyApiMcpServersApis {
	return s.Apis
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetAssumeRoleExtraPolicy() *string {
	return s.AssumeRoleExtraPolicy
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetAssumeRoleName() *string {
	return s.AssumeRoleName
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetDescription() *string {
	return s.Description
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetEnableAssumeRole() *bool {
	return s.EnableAssumeRole
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetId() *string {
	return s.Id
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetInstructions() *string {
	return s.Instructions
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetLanguage() *string {
	return s.Language
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetName() *string {
	return s.Name
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetOauthClientId() *string {
	return s.OauthClientId
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetPrompts() []*ListApiMcpServersResponseBodyApiMcpServersPrompts {
	return s.Prompts
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetSourceType() *string {
	return s.SourceType
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetSystemMcpServerInfo() *ListApiMcpServersResponseBodyApiMcpServersSystemMcpServerInfo {
	return s.SystemMcpServerInfo
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetSystemTools() []*string {
	return s.SystemTools
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetTerraformTools() []*ListApiMcpServersResponseBodyApiMcpServersTerraformTools {
	return s.TerraformTools
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) GetUrls() *ListApiMcpServersResponseBodyApiMcpServersUrls {
	return s.Urls
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetAdditionalApiDescriptions(v []*ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) *ListApiMcpServersResponseBodyApiMcpServers {
	s.AdditionalApiDescriptions = v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetApis(v []*ListApiMcpServersResponseBodyApiMcpServersApis) *ListApiMcpServersResponseBodyApiMcpServers {
	s.Apis = v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetAssumeRoleExtraPolicy(v string) *ListApiMcpServersResponseBodyApiMcpServers {
	s.AssumeRoleExtraPolicy = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetAssumeRoleName(v string) *ListApiMcpServersResponseBodyApiMcpServers {
	s.AssumeRoleName = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetCreateTime(v string) *ListApiMcpServersResponseBodyApiMcpServers {
	s.CreateTime = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetDescription(v string) *ListApiMcpServersResponseBodyApiMcpServers {
	s.Description = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetEnableAssumeRole(v bool) *ListApiMcpServersResponseBodyApiMcpServers {
	s.EnableAssumeRole = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetId(v string) *ListApiMcpServersResponseBodyApiMcpServers {
	s.Id = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetInstructions(v string) *ListApiMcpServersResponseBodyApiMcpServers {
	s.Instructions = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetLanguage(v string) *ListApiMcpServersResponseBodyApiMcpServers {
	s.Language = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetName(v string) *ListApiMcpServersResponseBodyApiMcpServers {
	s.Name = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetOauthClientId(v string) *ListApiMcpServersResponseBodyApiMcpServers {
	s.OauthClientId = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetPrompts(v []*ListApiMcpServersResponseBodyApiMcpServersPrompts) *ListApiMcpServersResponseBodyApiMcpServers {
	s.Prompts = v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetSourceType(v string) *ListApiMcpServersResponseBodyApiMcpServers {
	s.SourceType = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetSystemMcpServerInfo(v *ListApiMcpServersResponseBodyApiMcpServersSystemMcpServerInfo) *ListApiMcpServersResponseBodyApiMcpServers {
	s.SystemMcpServerInfo = v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetSystemTools(v []*string) *ListApiMcpServersResponseBodyApiMcpServers {
	s.SystemTools = v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetTerraformTools(v []*ListApiMcpServersResponseBodyApiMcpServersTerraformTools) *ListApiMcpServersResponseBodyApiMcpServers {
	s.TerraformTools = v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetUpdateTime(v string) *ListApiMcpServersResponseBodyApiMcpServers {
	s.UpdateTime = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) SetUrls(v *ListApiMcpServersResponseBodyApiMcpServersUrls) *ListApiMcpServersResponseBodyApiMcpServers {
	s.Urls = v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServers) Validate() error {
	if s.AdditionalApiDescriptions != nil {
		for _, item := range s.AdditionalApiDescriptions {
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

type ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions struct {
	// example:
	//
	// DescribeRegions
	ApiName         *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	ApiOverrideJson *string `json:"apiOverrideJson,omitempty" xml:"apiOverrideJson,omitempty"`
	// example:
	//
	// 2014-05-26
	ApiVersion      *string                                                                               `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	ConstParameters []*ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptionsConstParameters `json:"constParameters,omitempty" xml:"constParameters,omitempty" type:"Repeated"`
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

func (s ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) GoString() string {
	return s.String()
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) GetApiName() *string {
	return s.ApiName
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) GetApiOverrideJson() *string {
	return s.ApiOverrideJson
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) GetConstParameters() []*ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptionsConstParameters {
	return s.ConstParameters
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) GetEnableOutputSchema() *bool {
	return s.EnableOutputSchema
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) GetExecuteCliCommand() *bool {
	return s.ExecuteCliCommand
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) GetProduct() *string {
	return s.Product
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) SetApiName(v string) *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions {
	s.ApiName = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) SetApiOverrideJson(v string) *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions {
	s.ApiOverrideJson = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) SetApiVersion(v string) *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions {
	s.ApiVersion = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) SetConstParameters(v []*ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptionsConstParameters) *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions {
	s.ConstParameters = v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) SetEnableOutputSchema(v bool) *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions {
	s.EnableOutputSchema = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) SetExecuteCliCommand(v bool) *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions {
	s.ExecuteCliCommand = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) SetProduct(v string) *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions {
	s.Product = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptions) Validate() error {
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

type ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptionsConstParameters struct {
	// example:
	//
	// InstanceId
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 1234
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptionsConstParameters) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptionsConstParameters) GoString() string {
	return s.String()
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptionsConstParameters) GetKey() *string {
	return s.Key
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptionsConstParameters) GetValue() interface{} {
	return s.Value
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptionsConstParameters) SetKey(v string) *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptionsConstParameters {
	s.Key = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptionsConstParameters) SetValue(v interface{}) *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptionsConstParameters {
	s.Value = v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersAdditionalApiDescriptionsConstParameters) Validate() error {
	return dara.Validate(s)
}

type ListApiMcpServersResponseBodyApiMcpServersApis struct {
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

func (s ListApiMcpServersResponseBodyApiMcpServersApis) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServersResponseBodyApiMcpServersApis) GoString() string {
	return s.String()
}

func (s *ListApiMcpServersResponseBodyApiMcpServersApis) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *ListApiMcpServersResponseBodyApiMcpServersApis) GetProduct() *string {
	return s.Product
}

func (s *ListApiMcpServersResponseBodyApiMcpServersApis) GetSelectors() []*string {
	return s.Selectors
}

func (s *ListApiMcpServersResponseBodyApiMcpServersApis) SetApiVersion(v string) *ListApiMcpServersResponseBodyApiMcpServersApis {
	s.ApiVersion = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersApis) SetProduct(v string) *ListApiMcpServersResponseBodyApiMcpServersApis {
	s.Product = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersApis) SetSelectors(v []*string) *ListApiMcpServersResponseBodyApiMcpServersApis {
	s.Selectors = v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersApis) Validate() error {
	return dara.Validate(s)
}

type ListApiMcpServersResponseBodyApiMcpServersPrompts struct {
	Arguments []*ListApiMcpServersResponseBodyApiMcpServersPromptsArguments `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	Content   *string                                                       `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// prompt description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListApiMcpServersResponseBodyApiMcpServersPrompts) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServersResponseBodyApiMcpServersPrompts) GoString() string {
	return s.String()
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPrompts) GetArguments() []*ListApiMcpServersResponseBodyApiMcpServersPromptsArguments {
	return s.Arguments
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPrompts) GetContent() *string {
	return s.Content
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPrompts) GetDescription() *string {
	return s.Description
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPrompts) GetName() *string {
	return s.Name
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPrompts) SetArguments(v []*ListApiMcpServersResponseBodyApiMcpServersPromptsArguments) *ListApiMcpServersResponseBodyApiMcpServersPrompts {
	s.Arguments = v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPrompts) SetContent(v string) *ListApiMcpServersResponseBodyApiMcpServersPrompts {
	s.Content = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPrompts) SetDescription(v string) *ListApiMcpServersResponseBodyApiMcpServersPrompts {
	s.Description = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPrompts) SetName(v string) *ListApiMcpServersResponseBodyApiMcpServersPrompts {
	s.Name = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPrompts) Validate() error {
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

type ListApiMcpServersResponseBodyApiMcpServersPromptsArguments struct {
	// example:
	//
	// argument description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
}

func (s ListApiMcpServersResponseBodyApiMcpServersPromptsArguments) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServersResponseBodyApiMcpServersPromptsArguments) GoString() string {
	return s.String()
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPromptsArguments) GetDescription() *string {
	return s.Description
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPromptsArguments) GetName() *string {
	return s.Name
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPromptsArguments) GetRequired() *bool {
	return s.Required
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPromptsArguments) SetDescription(v string) *ListApiMcpServersResponseBodyApiMcpServersPromptsArguments {
	s.Description = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPromptsArguments) SetName(v string) *ListApiMcpServersResponseBodyApiMcpServersPromptsArguments {
	s.Name = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPromptsArguments) SetRequired(v bool) *ListApiMcpServersResponseBodyApiMcpServersPromptsArguments {
	s.Required = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersPromptsArguments) Validate() error {
	return dara.Validate(s)
}

type ListApiMcpServersResponseBodyApiMcpServersSystemMcpServerInfo struct {
	// example:
	//
	// mcp-system
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s ListApiMcpServersResponseBodyApiMcpServersSystemMcpServerInfo) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServersResponseBodyApiMcpServersSystemMcpServerInfo) GoString() string {
	return s.String()
}

func (s *ListApiMcpServersResponseBodyApiMcpServersSystemMcpServerInfo) GetName() *string {
	return s.Name
}

func (s *ListApiMcpServersResponseBodyApiMcpServersSystemMcpServerInfo) GetProduct() *string {
	return s.Product
}

func (s *ListApiMcpServersResponseBodyApiMcpServersSystemMcpServerInfo) SetName(v string) *ListApiMcpServersResponseBodyApiMcpServersSystemMcpServerInfo {
	s.Name = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersSystemMcpServerInfo) SetProduct(v string) *ListApiMcpServersResponseBodyApiMcpServersSystemMcpServerInfo {
	s.Product = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersSystemMcpServerInfo) Validate() error {
	return dara.Validate(s)
}

type ListApiMcpServersResponseBodyApiMcpServersTerraformTools struct {
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

func (s ListApiMcpServersResponseBodyApiMcpServersTerraformTools) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServersResponseBodyApiMcpServersTerraformTools) GoString() string {
	return s.String()
}

func (s *ListApiMcpServersResponseBodyApiMcpServersTerraformTools) GetAsync() *bool {
	return s.Async
}

func (s *ListApiMcpServersResponseBodyApiMcpServersTerraformTools) GetCode() *string {
	return s.Code
}

func (s *ListApiMcpServersResponseBodyApiMcpServersTerraformTools) GetDescription() *string {
	return s.Description
}

func (s *ListApiMcpServersResponseBodyApiMcpServersTerraformTools) GetDestroyPolicy() *string {
	return s.DestroyPolicy
}

func (s *ListApiMcpServersResponseBodyApiMcpServersTerraformTools) GetName() *string {
	return s.Name
}

func (s *ListApiMcpServersResponseBodyApiMcpServersTerraformTools) SetAsync(v bool) *ListApiMcpServersResponseBodyApiMcpServersTerraformTools {
	s.Async = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersTerraformTools) SetCode(v string) *ListApiMcpServersResponseBodyApiMcpServersTerraformTools {
	s.Code = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersTerraformTools) SetDescription(v string) *ListApiMcpServersResponseBodyApiMcpServersTerraformTools {
	s.Description = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersTerraformTools) SetDestroyPolicy(v string) *ListApiMcpServersResponseBodyApiMcpServersTerraformTools {
	s.DestroyPolicy = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersTerraformTools) SetName(v string) *ListApiMcpServersResponseBodyApiMcpServersTerraformTools {
	s.Name = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersTerraformTools) Validate() error {
	return dara.Validate(s)
}

type ListApiMcpServersResponseBodyApiMcpServersUrls struct {
	// example:
	//
	// https://mcpserverinner-pre.cn-zhangjiakou.aliyuncs.com/accounts/xxxx/custom/xxx/id/xxxx/mcp
	Mcp *string `json:"mcp,omitempty" xml:"mcp,omitempty"`
	// example:
	//
	// https://mcpserverinner-pre.cn-zhangjiakou.aliyuncs.com/accounts/xxxx/custom/xxx/id/xxxx/sse
	Sse *string `json:"sse,omitempty" xml:"sse,omitempty"`
}

func (s ListApiMcpServersResponseBodyApiMcpServersUrls) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServersResponseBodyApiMcpServersUrls) GoString() string {
	return s.String()
}

func (s *ListApiMcpServersResponseBodyApiMcpServersUrls) GetMcp() *string {
	return s.Mcp
}

func (s *ListApiMcpServersResponseBodyApiMcpServersUrls) GetSse() *string {
	return s.Sse
}

func (s *ListApiMcpServersResponseBodyApiMcpServersUrls) SetMcp(v string) *ListApiMcpServersResponseBodyApiMcpServersUrls {
	s.Mcp = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersUrls) SetSse(v string) *ListApiMcpServersResponseBodyApiMcpServersUrls {
	s.Sse = &v
	return s
}

func (s *ListApiMcpServersResponseBodyApiMcpServersUrls) Validate() error {
	return dara.Validate(s)
}
