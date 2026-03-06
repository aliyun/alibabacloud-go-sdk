// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApiMcpServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalApiDescriptions(v []*UpdateApiMcpServerRequestAdditionalApiDescriptions) *UpdateApiMcpServerRequest
	GetAdditionalApiDescriptions() []*UpdateApiMcpServerRequestAdditionalApiDescriptions
	SetApis(v []*UpdateApiMcpServerRequestApis) *UpdateApiMcpServerRequest
	GetApis() []*UpdateApiMcpServerRequestApis
	SetAssumeRoleExtraPolicy(v string) *UpdateApiMcpServerRequest
	GetAssumeRoleExtraPolicy() *string
	SetAssumeRoleName(v string) *UpdateApiMcpServerRequest
	GetAssumeRoleName() *string
	SetDescription(v string) *UpdateApiMcpServerRequest
	GetDescription() *string
	SetEnableAssumeRole(v bool) *UpdateApiMcpServerRequest
	GetEnableAssumeRole() *bool
	SetEnableCustomVpcWhitelist(v bool) *UpdateApiMcpServerRequest
	GetEnableCustomVpcWhitelist() *bool
	SetInstructions(v string) *UpdateApiMcpServerRequest
	GetInstructions() *string
	SetLanguage(v string) *UpdateApiMcpServerRequest
	GetLanguage() *string
	SetOauthClientId(v string) *UpdateApiMcpServerRequest
	GetOauthClientId() *string
	SetPrompts(v []*UpdateApiMcpServerRequestPrompts) *UpdateApiMcpServerRequest
	GetPrompts() []*UpdateApiMcpServerRequestPrompts
	SetPublicAccess(v string) *UpdateApiMcpServerRequest
	GetPublicAccess() *string
	SetSystemTools(v []*string) *UpdateApiMcpServerRequest
	GetSystemTools() []*string
	SetTerraformTools(v []*UpdateApiMcpServerRequestTerraformTools) *UpdateApiMcpServerRequest
	GetTerraformTools() []*UpdateApiMcpServerRequestTerraformTools
	SetVpcWhitelists(v []*string) *UpdateApiMcpServerRequest
	GetVpcWhitelists() []*string
	SetClientToken(v string) *UpdateApiMcpServerRequest
	GetClientToken() *string
	SetId(v string) *UpdateApiMcpServerRequest
	GetId() *string
}

type UpdateApiMcpServerRequest struct {
	// A list of supplementary API descriptions.
	AdditionalApiDescriptions []*UpdateApiMcpServerRequestAdditionalApiDescriptions `json:"additionalApiDescriptions,omitempty" xml:"additionalApiDescriptions,omitempty" type:"Repeated"`
	// The list of APIs to modify.
	Apis []*UpdateApiMcpServerRequestApis `json:"apis,omitempty" xml:"apis,omitempty" type:"Repeated"`
	// An additional policy for role assumption when multi-account access is enabled. If this policy exists, the permissions for the role assumption are based on this policy, which overwrites the permissions defined for the role itself.
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
	// The description of the API MCP server.
	//
	// example:
	//
	// 这是一个API MCP服务器。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to enable multi-account access.
	//
	// example:
	//
	// true
	EnableAssumeRole *bool `json:"enableAssumeRole,omitempty" xml:"enableAssumeRole,omitempty"`
	// Specifies whether to enable a custom VPC whitelist. If this is disabled, the account-level configuration is used.
	//
	// example:
	//
	// true
	EnableCustomVpcWhitelist *bool `json:"enableCustomVpcWhitelist,omitempty" xml:"enableCustomVpcWhitelist,omitempty"`
	// The MCP instruction. It prompts the large language model on how to use this MCP. The client must support the Instructions field of the standard MCP protocol.
	//
	// example:
	//
	// test
	Instructions *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	// The language of the API documentation for the API MCP service. You can select Chinese or English. Different language prompts may affect the AI\\"s response.
	//
	// example:
	//
	// ZH_CN
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The custom OAuth client ID to use when you select a custom OAuth configuration.
	//
	// `Only web and native applications are supported. The OAuth scope must include /acs/mcp-server.`
	//
	// example:
	//
	// 403*************370
	OauthClientId *string `json:"oauthClientId,omitempty" xml:"oauthClientId,omitempty"`
	// A list of prompt configurations.
	Prompts []*UpdateApiMcpServerRequestPrompts `json:"prompts,omitempty" xml:"prompts,omitempty" type:"Repeated"`
	// Specifies whether to enable public network access.
	//
	// example:
	//
	// on
	PublicAccess *string `json:"publicAccess,omitempty" xml:"publicAccess,omitempty"`
	// A list of system tools.
	SystemTools []*string `json:"systemTools,omitempty" xml:"systemTools,omitempty" type:"Repeated"`
	// A list of Terraform tools.
	TerraformTools []*UpdateApiMcpServerRequestTerraformTools `json:"terraformTools,omitempty" xml:"terraformTools,omitempty" type:"Repeated"`
	// The VPC whitelist that specifies the allowed source VPCs after public network access is disabled. If you do not set this parameter or leave it empty, access from all sources is allowed.
	VpcWhitelists []*string `json:"vpcWhitelists,omitempty" xml:"vpcWhitelists,omitempty" type:"Repeated"`
	// A client token that you can use to ensure the idempotence of the request. Generate a unique value from your client. The token can contain only ASCII characters and must be no more than 64 characters long. We recommend that you use a UUID. The token is valid for three days.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The ID of the API MCP service.
	//
	// This parameter is required.
	//
	// example:
	//
	// v6ZZ7ftCzEILW***
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s UpdateApiMcpServerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiMcpServerRequest) GoString() string {
	return s.String()
}

func (s *UpdateApiMcpServerRequest) GetAdditionalApiDescriptions() []*UpdateApiMcpServerRequestAdditionalApiDescriptions {
	return s.AdditionalApiDescriptions
}

func (s *UpdateApiMcpServerRequest) GetApis() []*UpdateApiMcpServerRequestApis {
	return s.Apis
}

func (s *UpdateApiMcpServerRequest) GetAssumeRoleExtraPolicy() *string {
	return s.AssumeRoleExtraPolicy
}

func (s *UpdateApiMcpServerRequest) GetAssumeRoleName() *string {
	return s.AssumeRoleName
}

func (s *UpdateApiMcpServerRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateApiMcpServerRequest) GetEnableAssumeRole() *bool {
	return s.EnableAssumeRole
}

func (s *UpdateApiMcpServerRequest) GetEnableCustomVpcWhitelist() *bool {
	return s.EnableCustomVpcWhitelist
}

func (s *UpdateApiMcpServerRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *UpdateApiMcpServerRequest) GetLanguage() *string {
	return s.Language
}

func (s *UpdateApiMcpServerRequest) GetOauthClientId() *string {
	return s.OauthClientId
}

func (s *UpdateApiMcpServerRequest) GetPrompts() []*UpdateApiMcpServerRequestPrompts {
	return s.Prompts
}

func (s *UpdateApiMcpServerRequest) GetPublicAccess() *string {
	return s.PublicAccess
}

func (s *UpdateApiMcpServerRequest) GetSystemTools() []*string {
	return s.SystemTools
}

func (s *UpdateApiMcpServerRequest) GetTerraformTools() []*UpdateApiMcpServerRequestTerraformTools {
	return s.TerraformTools
}

func (s *UpdateApiMcpServerRequest) GetVpcWhitelists() []*string {
	return s.VpcWhitelists
}

func (s *UpdateApiMcpServerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateApiMcpServerRequest) GetId() *string {
	return s.Id
}

func (s *UpdateApiMcpServerRequest) SetAdditionalApiDescriptions(v []*UpdateApiMcpServerRequestAdditionalApiDescriptions) *UpdateApiMcpServerRequest {
	s.AdditionalApiDescriptions = v
	return s
}

func (s *UpdateApiMcpServerRequest) SetApis(v []*UpdateApiMcpServerRequestApis) *UpdateApiMcpServerRequest {
	s.Apis = v
	return s
}

func (s *UpdateApiMcpServerRequest) SetAssumeRoleExtraPolicy(v string) *UpdateApiMcpServerRequest {
	s.AssumeRoleExtraPolicy = &v
	return s
}

func (s *UpdateApiMcpServerRequest) SetAssumeRoleName(v string) *UpdateApiMcpServerRequest {
	s.AssumeRoleName = &v
	return s
}

func (s *UpdateApiMcpServerRequest) SetDescription(v string) *UpdateApiMcpServerRequest {
	s.Description = &v
	return s
}

func (s *UpdateApiMcpServerRequest) SetEnableAssumeRole(v bool) *UpdateApiMcpServerRequest {
	s.EnableAssumeRole = &v
	return s
}

func (s *UpdateApiMcpServerRequest) SetEnableCustomVpcWhitelist(v bool) *UpdateApiMcpServerRequest {
	s.EnableCustomVpcWhitelist = &v
	return s
}

func (s *UpdateApiMcpServerRequest) SetInstructions(v string) *UpdateApiMcpServerRequest {
	s.Instructions = &v
	return s
}

func (s *UpdateApiMcpServerRequest) SetLanguage(v string) *UpdateApiMcpServerRequest {
	s.Language = &v
	return s
}

func (s *UpdateApiMcpServerRequest) SetOauthClientId(v string) *UpdateApiMcpServerRequest {
	s.OauthClientId = &v
	return s
}

func (s *UpdateApiMcpServerRequest) SetPrompts(v []*UpdateApiMcpServerRequestPrompts) *UpdateApiMcpServerRequest {
	s.Prompts = v
	return s
}

func (s *UpdateApiMcpServerRequest) SetPublicAccess(v string) *UpdateApiMcpServerRequest {
	s.PublicAccess = &v
	return s
}

func (s *UpdateApiMcpServerRequest) SetSystemTools(v []*string) *UpdateApiMcpServerRequest {
	s.SystemTools = v
	return s
}

func (s *UpdateApiMcpServerRequest) SetTerraformTools(v []*UpdateApiMcpServerRequestTerraformTools) *UpdateApiMcpServerRequest {
	s.TerraformTools = v
	return s
}

func (s *UpdateApiMcpServerRequest) SetVpcWhitelists(v []*string) *UpdateApiMcpServerRequest {
	s.VpcWhitelists = v
	return s
}

func (s *UpdateApiMcpServerRequest) SetClientToken(v string) *UpdateApiMcpServerRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateApiMcpServerRequest) SetId(v string) *UpdateApiMcpServerRequest {
	s.Id = &v
	return s
}

func (s *UpdateApiMcpServerRequest) Validate() error {
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

type UpdateApiMcpServerRequestAdditionalApiDescriptions struct {
	// The API name.
	//
	// example:
	//
	// DescribeRegions
	ApiName *string `json:"apiName,omitempty" xml:"apiName,omitempty"`
	// The API metadata in JSON format. For more information about the format, see https\\://api.aliyun.com/meta/v1/products/Ecs/versions/2014-05-26/apis/DescribeInstances/api.json. You can overwrite the summary and parameters fields.
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
	// A list of constant input parameters. These parameters are not included in the output during API parameter parsing.
	ConstParameters []*UpdateApiMcpServerRequestAdditionalApiDescriptionsConstParameters `json:"constParameters,omitempty" xml:"constParameters,omitempty" type:"Repeated"`
	// Specifies whether to return the schema of the response parameters. Returning the response parameter schema increases the overall size of the API MCP server. The default value is null, which means the response parameter schema is not returned.
	//
	// example:
	//
	// true
	EnableOutputSchema *bool `json:"enableOutputSchema,omitempty" xml:"enableOutputSchema,omitempty"`
	// Specifies whether to return the command-line interface (CLI) command for execution. In this mode, the API call is not actually executed. Instead, the corresponding CLI command is returned. This is suitable for long-running tasks that need to be executed using Alibaba Cloud CLI.
	//
	// example:
	//
	// true
	ExecuteCliCommand *bool `json:"executeCliCommand,omitempty" xml:"executeCliCommand,omitempty"`
	// The product code.
	//
	// - Call the GetRequestLog operation to obtain the product code from the response.
	//
	// - Find the product code from the URL of the product in OpenAPI Explorer. For example, the URL for Short Message Service is https\\://api.alibabacloud.com/product/Dysmsapi. The product code is \\`Dysmsapi\\`.
	//
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
}

func (s UpdateApiMcpServerRequestAdditionalApiDescriptions) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiMcpServerRequestAdditionalApiDescriptions) GoString() string {
	return s.String()
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) GetApiName() *string {
	return s.ApiName
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) GetApiOverrideJson() *string {
	return s.ApiOverrideJson
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) GetConstParameters() []*UpdateApiMcpServerRequestAdditionalApiDescriptionsConstParameters {
	return s.ConstParameters
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) GetEnableOutputSchema() *bool {
	return s.EnableOutputSchema
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) GetExecuteCliCommand() *bool {
	return s.ExecuteCliCommand
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) GetProduct() *string {
	return s.Product
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) SetApiName(v string) *UpdateApiMcpServerRequestAdditionalApiDescriptions {
	s.ApiName = &v
	return s
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) SetApiOverrideJson(v string) *UpdateApiMcpServerRequestAdditionalApiDescriptions {
	s.ApiOverrideJson = &v
	return s
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) SetApiVersion(v string) *UpdateApiMcpServerRequestAdditionalApiDescriptions {
	s.ApiVersion = &v
	return s
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) SetConstParameters(v []*UpdateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) *UpdateApiMcpServerRequestAdditionalApiDescriptions {
	s.ConstParameters = v
	return s
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) SetEnableOutputSchema(v bool) *UpdateApiMcpServerRequestAdditionalApiDescriptions {
	s.EnableOutputSchema = &v
	return s
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) SetExecuteCliCommand(v bool) *UpdateApiMcpServerRequestAdditionalApiDescriptions {
	s.ExecuteCliCommand = &v
	return s
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) SetProduct(v string) *UpdateApiMcpServerRequestAdditionalApiDescriptions {
	s.Product = &v
	return s
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptions) Validate() error {
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

type UpdateApiMcpServerRequestAdditionalApiDescriptionsConstParameters struct {
	// The parameter name. Only first-level parameter names are supported. For ROA-style APIs, you can set parameters such as body.xx. You cannot set values for parameters beyond the top level.
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

func (s UpdateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) GoString() string {
	return s.String()
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) GetKey() *string {
	return s.Key
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) GetValue() interface{} {
	return s.Value
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) SetKey(v string) *UpdateApiMcpServerRequestAdditionalApiDescriptionsConstParameters {
	s.Key = &v
	return s
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) SetValue(v interface{}) *UpdateApiMcpServerRequestAdditionalApiDescriptionsConstParameters {
	s.Value = v
	return s
}

func (s *UpdateApiMcpServerRequestAdditionalApiDescriptionsConstParameters) Validate() error {
	return dara.Validate(s)
}

type UpdateApiMcpServerRequestApis struct {
	// The POP version of the API that is exposed to the MCP server.
	//
	// example:
	//
	// 2014-05-26
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// The product code.
	//
	// - Call the GetRequestLog operation to obtain the product code from the response.
	//
	// - Find the product code from the URL of the product in OpenAPI Explorer. For example, the URL for Short Message Service is https\\://api.alibabacloud.com/product/Dysmsapi. The product code is \\`Dysmsapi\\`.
	//
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// A list of API name matching rules.
	Selectors []*string `json:"selectors,omitempty" xml:"selectors,omitempty" type:"Repeated"`
}

func (s UpdateApiMcpServerRequestApis) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiMcpServerRequestApis) GoString() string {
	return s.String()
}

func (s *UpdateApiMcpServerRequestApis) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *UpdateApiMcpServerRequestApis) GetProduct() *string {
	return s.Product
}

func (s *UpdateApiMcpServerRequestApis) GetSelectors() []*string {
	return s.Selectors
}

func (s *UpdateApiMcpServerRequestApis) SetApiVersion(v string) *UpdateApiMcpServerRequestApis {
	s.ApiVersion = &v
	return s
}

func (s *UpdateApiMcpServerRequestApis) SetProduct(v string) *UpdateApiMcpServerRequestApis {
	s.Product = &v
	return s
}

func (s *UpdateApiMcpServerRequestApis) SetSelectors(v []*string) *UpdateApiMcpServerRequestApis {
	s.Selectors = v
	return s
}

func (s *UpdateApiMcpServerRequestApis) Validate() error {
	return dara.Validate(s)
}

type UpdateApiMcpServerRequestPrompts struct {
	// A list of parameters that the prompt supports.
	Arguments []*UpdateApiMcpServerRequestPromptsArguments `json:"arguments,omitempty" xml:"arguments,omitempty" type:"Repeated"`
	// The content of the prompt. Variables are specified in the {{xxx}} format. The xxx variable must be defined in the arguments parameter.
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

func (s UpdateApiMcpServerRequestPrompts) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiMcpServerRequestPrompts) GoString() string {
	return s.String()
}

func (s *UpdateApiMcpServerRequestPrompts) GetArguments() []*UpdateApiMcpServerRequestPromptsArguments {
	return s.Arguments
}

func (s *UpdateApiMcpServerRequestPrompts) GetContent() *string {
	return s.Content
}

func (s *UpdateApiMcpServerRequestPrompts) GetDescription() *string {
	return s.Description
}

func (s *UpdateApiMcpServerRequestPrompts) GetName() *string {
	return s.Name
}

func (s *UpdateApiMcpServerRequestPrompts) SetArguments(v []*UpdateApiMcpServerRequestPromptsArguments) *UpdateApiMcpServerRequestPrompts {
	s.Arguments = v
	return s
}

func (s *UpdateApiMcpServerRequestPrompts) SetContent(v string) *UpdateApiMcpServerRequestPrompts {
	s.Content = &v
	return s
}

func (s *UpdateApiMcpServerRequestPrompts) SetDescription(v string) *UpdateApiMcpServerRequestPrompts {
	s.Description = &v
	return s
}

func (s *UpdateApiMcpServerRequestPrompts) SetName(v string) *UpdateApiMcpServerRequestPrompts {
	s.Name = &v
	return s
}

func (s *UpdateApiMcpServerRequestPrompts) Validate() error {
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

type UpdateApiMcpServerRequestPromptsArguments struct {
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

func (s UpdateApiMcpServerRequestPromptsArguments) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiMcpServerRequestPromptsArguments) GoString() string {
	return s.String()
}

func (s *UpdateApiMcpServerRequestPromptsArguments) GetDescription() *string {
	return s.Description
}

func (s *UpdateApiMcpServerRequestPromptsArguments) GetName() *string {
	return s.Name
}

func (s *UpdateApiMcpServerRequestPromptsArguments) GetRequired() *bool {
	return s.Required
}

func (s *UpdateApiMcpServerRequestPromptsArguments) SetDescription(v string) *UpdateApiMcpServerRequestPromptsArguments {
	s.Description = &v
	return s
}

func (s *UpdateApiMcpServerRequestPromptsArguments) SetName(v string) *UpdateApiMcpServerRequestPromptsArguments {
	s.Name = &v
	return s
}

func (s *UpdateApiMcpServerRequestPromptsArguments) SetRequired(v bool) *UpdateApiMcpServerRequestPromptsArguments {
	s.Required = &v
	return s
}

func (s *UpdateApiMcpServerRequestPromptsArguments) Validate() error {
	return dara.Validate(s)
}

type UpdateApiMcpServerRequestTerraformTools struct {
	// Specifies whether to execute the task asynchronously. If this is enabled, the system immediately proceeds to the next task after it initiates a task execution, without waiting for each resource operation to complete.
	//
	// example:
	//
	// true
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
	// The Terraform tool code. For more information, see [HCL language overview](https://www.alibabacloud.com/help/en/terraform/latest/hcl-language-overview).
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
	// Terraform Tool  description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The deletion policy. After a task is executed, the system applies one of the following cleanup policies to the temporary resources based on the task execution status.
	//
	// - NEVER: Does not delete any created resources, regardless of whether the task succeeds or fails.
	//
	// - ALWAYS: Immediately destroys all related resources after execution, regardless of whether the task succeeds or fails.
	//
	// - ON_FAILURE: Deletes related resources only if the task fails. If the task succeeds, the resources are retained.
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

func (s UpdateApiMcpServerRequestTerraformTools) String() string {
	return dara.Prettify(s)
}

func (s UpdateApiMcpServerRequestTerraformTools) GoString() string {
	return s.String()
}

func (s *UpdateApiMcpServerRequestTerraformTools) GetAsync() *bool {
	return s.Async
}

func (s *UpdateApiMcpServerRequestTerraformTools) GetCode() *string {
	return s.Code
}

func (s *UpdateApiMcpServerRequestTerraformTools) GetDescription() *string {
	return s.Description
}

func (s *UpdateApiMcpServerRequestTerraformTools) GetDestroyPolicy() *string {
	return s.DestroyPolicy
}

func (s *UpdateApiMcpServerRequestTerraformTools) GetName() *string {
	return s.Name
}

func (s *UpdateApiMcpServerRequestTerraformTools) SetAsync(v bool) *UpdateApiMcpServerRequestTerraformTools {
	s.Async = &v
	return s
}

func (s *UpdateApiMcpServerRequestTerraformTools) SetCode(v string) *UpdateApiMcpServerRequestTerraformTools {
	s.Code = &v
	return s
}

func (s *UpdateApiMcpServerRequestTerraformTools) SetDescription(v string) *UpdateApiMcpServerRequestTerraformTools {
	s.Description = &v
	return s
}

func (s *UpdateApiMcpServerRequestTerraformTools) SetDestroyPolicy(v string) *UpdateApiMcpServerRequestTerraformTools {
	s.DestroyPolicy = &v
	return s
}

func (s *UpdateApiMcpServerRequestTerraformTools) SetName(v string) *UpdateApiMcpServerRequestTerraformTools {
	s.Name = &v
	return s
}

func (s *UpdateApiMcpServerRequestTerraformTools) Validate() error {
	return dara.Validate(s)
}
