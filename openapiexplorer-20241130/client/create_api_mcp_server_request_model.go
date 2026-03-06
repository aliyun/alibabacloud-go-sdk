// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApiMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalApiDescriptions(v []*CreateApiMcpServerRequestAdditionalApiDescriptions) *CreateApiMcpServerRequest
	GetAdditionalApiDescriptions() []*CreateApiMcpServerRequestAdditionalApiDescriptions
	SetApis(v []*CreateApiMcpServerRequestApis) *CreateApiMcpServerRequest
	GetApis() []*CreateApiMcpServerRequestApis
	SetAssumeRoleExtraPolicy(v string) *CreateApiMcpServerRequest
	GetAssumeRoleExtraPolicy() *string
	SetAssumeRoleName(v string) *CreateApiMcpServerRequest
	GetAssumeRoleName() *string
	SetClientToken(v string) *CreateApiMcpServerRequest
	GetClientToken() *string
	SetDescription(v string) *CreateApiMcpServerRequest
	GetDescription() *string
	SetEnableAssumeRole(v bool) *CreateApiMcpServerRequest
	GetEnableAssumeRole() *bool
	SetEnableCustomVpcWhitelist(v bool) *CreateApiMcpServerRequest
	GetEnableCustomVpcWhitelist() *bool
	SetInstructions(v string) *CreateApiMcpServerRequest
	GetInstructions() *string
	SetLanguage(v string) *CreateApiMcpServerRequest
	GetLanguage() *string
	SetName(v string) *CreateApiMcpServerRequest
	GetName() *string
	SetOauthClientId(v string) *CreateApiMcpServerRequest
	GetOauthClientId() *string
	SetPrompts(v []*CreateApiMcpServerRequestPrompts) *CreateApiMcpServerRequest
	GetPrompts() []*CreateApiMcpServerRequestPrompts
	SetPublicAccess(v string) *CreateApiMcpServerRequest
	GetPublicAccess() *string
	SetSystemTools(v []*string) *CreateApiMcpServerRequest
	GetSystemTools() []*string
	SetTerraformTools(v []*CreateApiMcpServerRequestTerraformTools) *CreateApiMcpServerRequest
	GetTerraformTools() []*CreateApiMcpServerRequestTerraformTools
	SetVpcWhitelists(v []*string) *CreateApiMcpServerRequest
	GetVpcWhitelists() []*string
}

type CreateApiMcpServerRequest struct {
	// The list of supplementary API descriptions.
	AdditionalApiDescriptions []*CreateApiMcpServerRequestAdditionalApiDescriptions `json:"additionalApiDescriptions,omitempty" xml:"additionalApiDescriptions,omitempty" type:"Repeated"`
	// The list of APIs to add. This parameter cannot be empty.
	//
	// This parameter is required.
	Apis []*CreateApiMcpServerRequestApis `json:"apis,omitempty" xml:"apis,omitempty" type:"Repeated"`
	// An additional policy for role assumption when multi-account access is enabled. If this policy exists, the permissions for role assumption are based on this policy, which overwrites the permissions defined for the role itself.
	//
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
	// The name of the RAM role in the destination account that is assumed for cross-account operations when multi-account access is enabled.
	//
	// example:
	//
	// test
	AssumeRoleName *string `json:"assumeRoleName,omitempty" xml:"assumeRoleName,omitempty"`
	// Ensures the idempotence of the request. Generate a parameter value from your client to make sure that the value is unique among different requests. The client token can contain only ASCII characters and cannot exceed 64 characters in length. Use a universally unique identifier (UUID). The token expires in three days.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The description of the API MCP service.
	//
	// example:
	//
	// 这是一个API MCP服务。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to enable multi-account access.
	//
	// example:
	//
	// true
	EnableAssumeRole *bool `json:"enableAssumeRole,omitempty" xml:"enableAssumeRole,omitempty"`
	// Specifies whether to enable a custom VPC whitelist. If not enabled, the account-level configuration is used.
	//
	// example:
	//
	// true
	EnableCustomVpcWhitelist *bool `json:"enableCustomVpcWhitelist,omitempty" xml:"enableCustomVpcWhitelist,omitempty"`
	// The MCP instruction. It prompts the large language model on how to use the MCP. The client must support the Instructions field of the standard MCP protocol.
	//
	// example:
	//
	// test
	Instructions *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	// The language of the API documentation for the API MCP service. You can select Chinese or English API documentation. The prompts in different languages may affect the AI\\"s response.
	//
	// example:
	//
	// ZH_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The name of the MCP server. The name must be 3 to 64 characters in length and can contain only lowercase letters, digits, underscores (_), and hyphens (-). It cannot start with a digit. The name must be unique within the same Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// mcp-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The custom OAuth client ID when you select a custom OAuth configuration.
	//
	// `Only web and native applications are supported. The OAuth scope must include /acs/mcp-server.`
	//
	// example:
	//
	// 403*************370
	OauthClientId *string `json:"oauthClientId,omitempty" xml:"oauthClientId,omitempty"`
	// The list of prompt configurations.
	Prompts []*CreateApiMcpServerRequestPrompts `json:"prompts,omitempty" xml:"prompts,omitempty" type:"Repeated"`
	// Specifies whether to enable public network access.
	//
	// example:
	//
	// on
	PublicAccess *string `json:"publicAccess,omitempty" xml:"publicAccess,omitempty"`
	// The list of system tools.
	SystemTools []*string `json:"systemTools,omitempty" xml:"systemTools,omitempty" type:"Repeated"`
	// The list of Terraform tools.
	TerraformTools []*CreateApiMcpServerRequestTerraformTools `json:"terraformTools,omitempty" xml:"terraformTools,omitempty" type:"Repeated"`
	// The VPC whitelist that restricts the source of access after public network access is disabled. If you do not set this parameter or leave it empty, the source is not restricted.
	VpcWhitelists []*string `json:"vpcWhitelists,omitempty" xml:"vpcWhitelists,omitempty" type:"Repeated"`
}

func (s CreateApiMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApiMcpServerRequest) GoString() string {
	return s.String()
}

func (s *CreateApiMcpServerRequest) GetAdditionalApiDescriptions() []*CreateApiMcpServerRequestAdditionalApiDescriptions {
	return s.AdditionalApiDescriptions
}

func (s *CreateApiMcpServerRequest) GetApis() []*CreateApiMcpServerRequestApis {
	return s.Apis
}

func (s *CreateApiMcpServerRequest) GetAssumeRoleExtraPolicy() *string {
	return s.AssumeRoleExtraPolicy
}

func (s *CreateApiMcpServerRequest) GetAssumeRoleName() *string {
	return s.AssumeRoleName
}

func (s *CreateApiMcpServerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateApiMcpServerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateApiMcpServerRequest) GetEnableAssumeRole() *bool {
	return s.EnableAssumeRole
}

func (s *CreateApiMcpServerRequest) GetEnableCustomVpcWhitelist() *bool {
	return s.EnableCustomVpcWhitelist
}

func (s *CreateApiMcpServerRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *CreateApiMcpServerRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateApiMcpServerRequest) GetName() *string {
	return s.Name
}

func (s *CreateApiMcpServerRequest) GetOauthClientId() *string {
	return s.OauthClientId
}

func (s *CreateApiMcpServerRequest) GetPrompts() []*CreateApiMcpServerRequestPrompts {
	return s.Prompts
}

func (s *CreateApiMcpServerRequest) GetPublicAccess() *string {
	return s.PublicAccess
}

func (s *CreateApiMcpServerRequest) GetSystemTools() []*string {
	return s.SystemTools
}

func (s *CreateApiMcpServerRequest) GetTerraformTools() []*CreateApiMcpServerRequestTerraformTools {
	return s.TerraformTools
}

func (s *CreateApiMcpServerRequest) GetVpcWhitelists() []*string {
	return s.VpcWhitelists
}

func (s *CreateApiMcpServerRequest) SetAdditionalApiDescriptions(v []*CreateApiMcpServerRequestAdditionalApiDescriptions) *CreateApiMcpServerRequest {
	s.AdditionalApiDescriptions = v
	return s
}

func (s *CreateApiMcpServerRequest) SetApis(v []*CreateApiMcpServerRequestApis) *CreateApiMcpServerRequest {
	s.Apis = v
	return s
}

func (s *CreateApiMcpServerRequest) SetAssumeRoleExtraPolicy(v string) *CreateApiMcpServerRequest {
	s.AssumeRoleExtraPolicy = &v
	return s
}

func (s *CreateApiMcpServerRequest) SetAssumeRoleName(v string) *CreateApiMcpServerRequest {
	s.AssumeRoleName = &v
	return s
}

func (s *CreateApiMcpServerRequest) SetClientToken(v string) *CreateApiMcpServerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApiMcpServerRequest) SetDescription(v string) *CreateApiMcpServerRequest {
	s.Description = &v
	return s
}

func (s *CreateApiMcpServerRequest) SetEnableAssumeRole(v bool) *CreateApiMcpServerRequest {
	s.EnableAssumeRole = &v
	return s
}

func (s *CreateApiMcpServerRequest) SetEnableCustomVpcWhitelist(v bool) *CreateApiMcpServerRequest {
	s.EnableCustomVpcWhitelist = &v
	return s
}

func (s *CreateApiMcpServerRequest) SetInstructions(v string) *CreateApiMcpServerRequest {
	s.Instructions = &v
	return s
}

func (s *CreateApiMcpServerRequest) SetLanguage(v string) *CreateApiMcpServerRequest {
	s.Language = &v
	return s
}

func (s *CreateApiMcpServerRequest) SetName(v string) *CreateApiMcpServerRequest {
	s.Name = &v
	return s
}

func (s *CreateApiMcpServerRequest) SetOauthClientId(v string) *CreateApiMcpServerRequest {
	s.OauthClientId = &v
	return s
}

func (s *CreateApiMcpServerRequest) SetPrompts(v []*CreateApiMcpServerRequestPrompts) *CreateApiMcpServerRequest {
	s.Prompts = v
	return s
}

func (s *CreateApiMcpServerRequest) SetPublicAccess(v string) *CreateApiMcpServerRequest {
	s.PublicAccess = &v
	return s
}

func (s *CreateApiMcpServerRequest) SetSystemTools(v []*string) *CreateApiMcpServerRequest {
	s.SystemTools = v
	return s
}

func (s *CreateApiMcpServerRequest) SetTerraformTools(v []*CreateApiMcpServerRequestTerraformTools) *CreateApiMcpServerRequest {
	s.TerraformTools = v
	return s
}

func (s *CreateApiMcpServerRequest) SetVpcWhitelists(v []*string) *CreateApiMcpServerRequest {
	s.VpcWhitelists = v
	return s
}

func (s *CreateApiMcpServerRequest) Validate() error {
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
	if s.TerraformTools != nil {
		for _, item := range s.TerraformTools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateApiMcpServerRequestAdditionalApiDescriptions struct {
	// The API name.
	//
	// example:
	//
	// DescribeRegions
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// The API metadata in JSON format. For more information about the format, see https\\://api.aliyun.com/meta/v1/products/Ecs/versions/2014-05-26/apis/DescribeInstances/api.json. You can overwrite the summary and parameters.
	//
	// example:
	//
	// {
	//
	//   "summary": "本接口支持根据不同请求条件查询实例列表，并关联查询实例的详细信息。"
	//
	// }
	ApiOverrideJson *string `json:"apiOverrideJson,omitempty" xml:"apiOverrideJson,omitempty"`
	// The POP version of the API that is exposed to the MCP server.
	//
	// example:
	//
	// 2014-05-26
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// The list of constant input parameters. These parameters are not included in the output during API parameter parsing.
	ConstParameters []*CreateApiMcpServerRequestAdditionalApiDescriptionsConstParameters `json:"constParameters,omitempty" xml:"constParameters,omitempty" type:"Repeated"`
	// Specifies whether to return the schema of the response parameters. Returning the schema increases the size of the entire API MCP server. The default value is null, which means the schema is not returned.
	//
	// example:
	//
	// true
	EnableOutputSchema *bool `json:"enableOutputSchema,omitempty" xml:"enableOutputSchema,omitempty"`
	// Specifies whether to return the command-line interface (CLI) command for execution. In this mode, the API call is not executed. Instead, the corresponding CLI command is returned. This is suitable for long-running tasks that need to be executed using Alibaba Cloud CLI.
	//
	// example:
	//
	// true
	ExecuteCliCommand *bool `json:"executeCliCommand,omitempty" xml:"executeCliCommand,omitempty"`
	// The product code.
	//
	// - Call the GetRequestLog operation to obtain the product code from the response.
	//
	// - Find the product code from the URL of the product in OpenAPI Portal. For example, the URL of Short Message Service in OpenAPI Portal is https\\://api.alibabacloud.com/product/Dysmsapi. The product code is Dysmsapi.
	//
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s CreateApiMcpServerRequestAdditionalApiDescriptions) String() string {
	return dara.Prettify(s)
}

func (s CreateApiMcpServerRequestAdditionalApiDescriptions) GoString() string {
	return s.String()
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) GetApiName() *string {
	return s.ApiName
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) GetApiOverrideJson() *string {
	return s.ApiOverrideJson
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) GetConstParameters() []*CreateApiMcpServerRequestAdditionalApiDescriptionsConstParameters {
	return s.ConstParameters
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) GetEnableOutputSchema() *bool {
	return s.EnableOutputSchema
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) GetExecuteCliCommand() *bool {
	return s.ExecuteCliCommand
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) GetProduct() *string {
	return s.Product
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) SetApiName(v string) *CreateApiMcpServerRequestAdditionalApiDescriptions {
	s.ApiName = &v
	return s
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) SetApiOverrideJson(v string) *CreateApiMcpServerRequestAdditionalApiDescriptions {
	s.ApiOverrideJson = &v
	return s
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) SetApiVersion(v string) *CreateApiMcpServerRequestAdditionalApiDescriptions {
	s.ApiVersion = &v
	return s
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) SetConstParameters(v []*CreateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) *CreateApiMcpServerRequestAdditionalApiDescriptions {
	s.ConstParameters = v
	return s
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) SetEnableOutputSchema(v bool) *CreateApiMcpServerRequestAdditionalApiDescriptions {
	s.EnableOutputSchema = &v
	return s
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) SetExecuteCliCommand(v bool) *CreateApiMcpServerRequestAdditionalApiDescriptions {
	s.ExecuteCliCommand = &v
	return s
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) SetProduct(v string) *CreateApiMcpServerRequestAdditionalApiDescriptions {
	s.Product = &v
	return s
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptions) Validate() error {
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

type CreateApiMcpServerRequestAdditionalApiDescriptionsConstParameters struct {
	// The parameter name. Only first-level parameter names are supported. For ROA-style APIs, you can set the parameter to body.xx. You cannot set values that exceed the top-level parameters.
	//
	// example:
	//
	// InstanceId
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// 1234
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) GoString() string {
	return s.String()
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) GetKey() *string {
	return s.Key
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) GetValue() interface{} {
	return s.Value
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) SetKey(v string) *CreateApiMcpServerRequestAdditionalApiDescriptionsConstParameters {
	s.Key = &v
	return s
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) SetValue(v interface{}) *CreateApiMcpServerRequestAdditionalApiDescriptionsConstParameters {
	s.Value = v
	return s
}

func (s *CreateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) Validate() error {
	return dara.Validate(s)
}

type CreateApiMcpServerRequestApis struct {
	// The POP version of the API that is exposed to the MCP server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2014-05-26
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// The product code.
	//
	// - Call the GetRequestLog operation to obtain the product code from the response.
	//
	// - Find the product code from the URL of the product in OpenAPI Portal. For example, the URL of Short Message Service in OpenAPI Portal is https\\://api.alibabacloud.com/product/Dysmsapi. The product code is Dysmsapi.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// The list of API name matching rules. This parameter cannot be empty.
	//
	// This parameter is required.
	Selectors []*string `json:"selectors,omitempty" xml:"selectors,omitempty" type:"Repeated"`
}

func (s CreateApiMcpServerRequestApis) String() string {
	return dara.Prettify(s)
}

func (s CreateApiMcpServerRequestApis) GoString() string {
	return s.String()
}

func (s *CreateApiMcpServerRequestApis) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *CreateApiMcpServerRequestApis) GetProduct() *string {
	return s.Product
}

func (s *CreateApiMcpServerRequestApis) GetSelectors() []*string {
	return s.Selectors
}

func (s *CreateApiMcpServerRequestApis) SetApiVersion(v string) *CreateApiMcpServerRequestApis {
	s.ApiVersion = &v
	return s
}

func (s *CreateApiMcpServerRequestApis) SetProduct(v string) *CreateApiMcpServerRequestApis {
	s.Product = &v
	return s
}

func (s *CreateApiMcpServerRequestApis) SetSelectors(v []*string) *CreateApiMcpServerRequestApis {
	s.Selectors = v
	return s
}

func (s *CreateApiMcpServerRequestApis) Validate() error {
	return dara.Validate(s)
}

type CreateApiMcpServerRequestPrompts struct {
	// The list of parameters supported by the prompt.
	Arguments []*CreateApiMcpServerRequestPromptsArguments `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	// The content of the prompt. Variables are specified in the {{xxx}} format. xxx is a variable that must be defined in the arguments parameter.
	//
	// example:
	//
	// prompt正文，{{name}}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The description.
	//
	// example:
	//
	// prompt description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The prompt name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateApiMcpServerRequestPrompts) String() string {
	return dara.Prettify(s)
}

func (s CreateApiMcpServerRequestPrompts) GoString() string {
	return s.String()
}

func (s *CreateApiMcpServerRequestPrompts) GetArguments() []*CreateApiMcpServerRequestPromptsArguments {
	return s.Arguments
}

func (s *CreateApiMcpServerRequestPrompts) GetContent() *string {
	return s.Content
}

func (s *CreateApiMcpServerRequestPrompts) GetDescription() *string {
	return s.Description
}

func (s *CreateApiMcpServerRequestPrompts) GetName() *string {
	return s.Name
}

func (s *CreateApiMcpServerRequestPrompts) SetArguments(v []*CreateApiMcpServerRequestPromptsArguments) *CreateApiMcpServerRequestPrompts {
	s.Arguments = v
	return s
}

func (s *CreateApiMcpServerRequestPrompts) SetContent(v string) *CreateApiMcpServerRequestPrompts {
	s.Content = &v
	return s
}

func (s *CreateApiMcpServerRequestPrompts) SetDescription(v string) *CreateApiMcpServerRequestPrompts {
	s.Description = &v
	return s
}

func (s *CreateApiMcpServerRequestPrompts) SetName(v string) *CreateApiMcpServerRequestPrompts {
	s.Name = &v
	return s
}

func (s *CreateApiMcpServerRequestPrompts) Validate() error {
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

type CreateApiMcpServerRequestPromptsArguments struct {
	// The description of the parameter.
	//
	// example:
	//
	// argument description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether the parameter is required.
	//
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
}

func (s CreateApiMcpServerRequestPromptsArguments) String() string {
	return dara.Prettify(s)
}

func (s CreateApiMcpServerRequestPromptsArguments) GoString() string {
	return s.String()
}

func (s *CreateApiMcpServerRequestPromptsArguments) GetDescription() *string {
	return s.Description
}

func (s *CreateApiMcpServerRequestPromptsArguments) GetName() *string {
	return s.Name
}

func (s *CreateApiMcpServerRequestPromptsArguments) GetRequired() *bool {
	return s.Required
}

func (s *CreateApiMcpServerRequestPromptsArguments) SetDescription(v string) *CreateApiMcpServerRequestPromptsArguments {
	s.Description = &v
	return s
}

func (s *CreateApiMcpServerRequestPromptsArguments) SetName(v string) *CreateApiMcpServerRequestPromptsArguments {
	s.Name = &v
	return s
}

func (s *CreateApiMcpServerRequestPromptsArguments) SetRequired(v bool) *CreateApiMcpServerRequestPromptsArguments {
	s.Required = &v
	return s
}

func (s *CreateApiMcpServerRequestPromptsArguments) Validate() error {
	return dara.Validate(s)
}

type CreateApiMcpServerRequestTerraformTools struct {
	// Specifies whether to execute the task asynchronously. If set to true, the system immediately proceeds to the next task after initiating a task execution, without waiting for each resource operation to complete.
	//
	// example:
	//
	// true
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// The code of the Terraform tool. For more information, see [HCL language overview](https://www.alibabacloud.com/help/en/terraform/terraform-configuration-and-hcl-language-overview).
	//
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
	// The description of the Terraform tool.
	//
	// example:
	//
	// terraform tool description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The deletion policy. After a task is executed, the system applies the following cleanup policy to temporary resources based on the task execution status.
	//
	// - NEVER: Does not delete any created resources, regardless of whether the task execution succeeds or fails.
	//
	// - ALWAYS: Immediately destroys all related resources after the task is executed, regardless of whether the execution succeeds or fails.
	//
	// - ON_FAILURE: Deletes related resources only when the task execution fails. If the task execution succeeds, the resources are retained.
	//
	// example:
	//
	// ON_FAILURE
	DestroyPolicy *string `json:"destroyPolicy,omitempty" xml:"destroyPolicy,omitempty"`
	// The name of the Terraform tool.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateApiMcpServerRequestTerraformTools) String() string {
	return dara.Prettify(s)
}

func (s CreateApiMcpServerRequestTerraformTools) GoString() string {
	return s.String()
}

func (s *CreateApiMcpServerRequestTerraformTools) GetAsync() *bool {
	return s.Async
}

func (s *CreateApiMcpServerRequestTerraformTools) GetCode() *string {
	return s.Code
}

func (s *CreateApiMcpServerRequestTerraformTools) GetDescription() *string {
	return s.Description
}

func (s *CreateApiMcpServerRequestTerraformTools) GetDestroyPolicy() *string {
	return s.DestroyPolicy
}

func (s *CreateApiMcpServerRequestTerraformTools) GetName() *string {
	return s.Name
}

func (s *CreateApiMcpServerRequestTerraformTools) SetAsync(v bool) *CreateApiMcpServerRequestTerraformTools {
	s.Async = &v
	return s
}

func (s *CreateApiMcpServerRequestTerraformTools) SetCode(v string) *CreateApiMcpServerRequestTerraformTools {
	s.Code = &v
	return s
}

func (s *CreateApiMcpServerRequestTerraformTools) SetDescription(v string) *CreateApiMcpServerRequestTerraformTools {
	s.Description = &v
	return s
}

func (s *CreateApiMcpServerRequestTerraformTools) SetDestroyPolicy(v string) *CreateApiMcpServerRequestTerraformTools {
	s.DestroyPolicy = &v
	return s
}

func (s *CreateApiMcpServerRequestTerraformTools) SetName(v string) *CreateApiMcpServerRequestTerraformTools {
	s.Name = &v
	return s
}

func (s *CreateApiMcpServerRequestTerraformTools) Validate() error {
	return dara.Validate(s)
}
