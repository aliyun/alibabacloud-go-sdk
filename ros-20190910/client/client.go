// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type CancelStackOperationRequest struct {
	// The operations that you want to cancel on the stack.
	AllowedStackOperations []*string `json:"AllowedStackOperations,omitempty" xml:"AllowedStackOperations,omitempty" type:"Repeated"`
	// The method that you want to use to cancel the operations. Valid values:
	//
	// *   Quick: cancels the operations on the stack at the earliest opportunity. In this case, Resource Orchestration Service (ROS) stops scheduling new resources and stops running resources at the earliest opportunity. If you use this method, the resource status may become invalid and subsequent stack operations may be affected.
	// *   Safe (default): cancels the operations on the stack in a secure manner. In this case, ROS stops scheduling new resources and waits for running resources to be stopped.
	CancelType *string `json:"CancelType,omitempty" xml:"CancelType,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The stack ID.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s CancelStackOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelStackOperationRequest) GoString() string {
	return s.String()
}

func (s *CancelStackOperationRequest) SetAllowedStackOperations(v []*string) *CancelStackOperationRequest {
	s.AllowedStackOperations = v
	return s
}

func (s *CancelStackOperationRequest) SetCancelType(v string) *CancelStackOperationRequest {
	s.CancelType = &v
	return s
}

func (s *CancelStackOperationRequest) SetRegionId(v string) *CancelStackOperationRequest {
	s.RegionId = &v
	return s
}

func (s *CancelStackOperationRequest) SetStackId(v string) *CancelStackOperationRequest {
	s.StackId = &v
	return s
}

type CancelStackOperationResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelStackOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelStackOperationResponseBody) GoString() string {
	return s.String()
}

func (s *CancelStackOperationResponseBody) SetRequestId(v string) *CancelStackOperationResponseBody {
	s.RequestId = &v
	return s
}

type CancelStackOperationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelStackOperationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelStackOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelStackOperationResponse) GoString() string {
	return s.String()
}

func (s *CancelStackOperationResponse) SetHeaders(v map[string]*string) *CancelStackOperationResponse {
	s.Headers = v
	return s
}

func (s *CancelStackOperationResponse) SetStatusCode(v int32) *CancelStackOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelStackOperationResponse) SetBody(v *CancelStackOperationResponseBody) *CancelStackOperationResponse {
	s.Body = v
	return s
}

type CancelUpdateStackRequest struct {
	// The method to cancel the update operation. Valid values:
	//
	// *   Quick: cancels the update of a stack as soon as possible.
	// *   Safe: cancels the update of a stack as safely as possible.
	CancelType *string `json:"CancelType,omitempty" xml:"CancelType,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s CancelUpdateStackRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelUpdateStackRequest) GoString() string {
	return s.String()
}

func (s *CancelUpdateStackRequest) SetCancelType(v string) *CancelUpdateStackRequest {
	s.CancelType = &v
	return s
}

func (s *CancelUpdateStackRequest) SetRegionId(v string) *CancelUpdateStackRequest {
	s.RegionId = &v
	return s
}

func (s *CancelUpdateStackRequest) SetStackId(v string) *CancelUpdateStackRequest {
	s.StackId = &v
	return s
}

type CancelUpdateStackResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelUpdateStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelUpdateStackResponseBody) GoString() string {
	return s.String()
}

func (s *CancelUpdateStackResponseBody) SetRequestId(v string) *CancelUpdateStackResponseBody {
	s.RequestId = &v
	return s
}

type CancelUpdateStackResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelUpdateStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelUpdateStackResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelUpdateStackResponse) GoString() string {
	return s.String()
}

func (s *CancelUpdateStackResponse) SetHeaders(v map[string]*string) *CancelUpdateStackResponse {
	s.Headers = v
	return s
}

func (s *CancelUpdateStackResponse) SetStatusCode(v int32) *CancelUpdateStackResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelUpdateStackResponse) SetBody(v *CancelUpdateStackResponseBody) *CancelUpdateStackResponse {
	s.Body = v
	return s
}

type ContinueCreateStackRequest struct {
	// Specifies whether only to validate the stack in the request. Valid values:
	//
	// *   true: only validates the stack.
	// *   false (default): validates and continues to create the stack.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The mode in which ROS continues to create the stack. Valid values:
	//
	// *   Recreate (default)
	//
	//     If you set this parameter to Recreate, ROS continues to create only the following types of resources:
	//
	//     *   Resources that fail to be created
	//     *   Resources that you specify for RecreatingResources.N
	//     *   Dependencies of the resources that you specify for RecreatingResources.N
	//     *   Resources that you have not created
	//
	// > RecreatingResources.N, TemplateBody, TemplateURL, and Parameters take effect only when Mode is set to Recreate.
	//
	// *   Ignore
	//
	//     *   ROS ignores and discards resources that fail to be created and you have not created, and considers that the stack is successfully created.
	//     *   The body of the template that you use to create the stack is changed.
	//
	// > This mode is available only for ROS stacks.
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The maximum number of concurrent operations that can be performed on resources.
	//
	// By default, this parameter is empty. You can set this parameter to an integer that is greater than or equal to 0.
	//
	// > - If you set this parameter to an integer that is greater than 0, the integer is used.
	// > - If you set this parameter to 0, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	// > - If you leave this parameter empty, the value that you specified for this parameter in the previous request is used. If you left this parameter empty in the previous request, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	// > - If you set this parameter to a specific value, ROS associates the value with the stack. The value affects subsequent operations on the stack.
	Parallelism *int64 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The template parameters that you want to use to override specific parameters.
	Parameters []*ContinueCreateStackRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The name of the RAM role. Resource Orchestration Service (ROS) assumes the RAM role to create the stack and uses the credentials of the role to call the APIs of Alibaba Cloud services.\
	// ROS assumes the RAM role to perform operations on the stack. If you have permissions to perform operations on the stack but do not have permissions to use the RAM role, ROS still assumes the RAM role. You must make sure that the least privileges are granted to the RAM role.\
	// If you do not specify this parameter, ROS assumes the existing role that is associated with the stack. If no roles are available, ROS uses a temporary credential that is generated from the credentials of your account.\
	// The name of the RAM role can be up to 64 bytes in length.
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The options that ROS adopts when ROS continues to create the stack.
	RecreatingOptions []*string `json:"RecreatingOptions,omitempty" xml:"RecreatingOptions,omitempty" type:"Repeated"`
	// The resources that ROS continues to create after the resources failed to be created. You can add new resources to the resources that ROS continues to create. ROS continues to create all dependencies of the new resources.
	//
	// > This parameter is available only for ROS stacks.
	RecreatingResources []*string `json:"RecreatingResources,omitempty" xml:"RecreatingResources,omitempty" type:"Repeated"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The stack ID.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.\
	// If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.
	//
	// A ROS template is subject to the following limits:
	//
	// *   You can modify only the following sections in the template: Description, Metadata, Resources, and Outputs.
	//
	// *   You cannot add sections to or remove sections from the template.
	//
	// *   The Resources section is subject to the following limits:
	//
	//     *   You cannot delete the resources or change the template body for the resources that you do not want to continue to create.
	//     *   You can delete the resources or change the template body for the resources that you want to continue to create.
	//     *   You can add resources to this section.
	//
	//
	//
	// > - This parameter takes effect only when Mode is set to Recreate.
	// > - You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId. If you do not specify the parameters, the existing template is used.
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The template ID. This parameter applies to shared and private templates.
	//
	// > - This parameter takes effect when `Mode` is set to `Recreate`. When you specify TemplateId of a template, the template is subject to the limits that are described for `TemplateBody` in this topic.
	// > - You can specify only one of the following parameters: `TemplateBody`, `TemplateURL`, and `TemplateId`. If you do not specify the parameters, the existing template is used.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length.
	//
	// If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// > - This parameter takes effect only when Mode is set to Recreate. When you specify TemplateURL of a template, the template is subject to the limits that are described for TemplateBody in this topic.
	// > - You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId. If you do not specify the parameters, the existing template is used.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. This parameter takes effect only when TemplateId is specified.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s ContinueCreateStackRequest) String() string {
	return tea.Prettify(s)
}

func (s ContinueCreateStackRequest) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackRequest) SetDryRun(v bool) *ContinueCreateStackRequest {
	s.DryRun = &v
	return s
}

func (s *ContinueCreateStackRequest) SetMode(v string) *ContinueCreateStackRequest {
	s.Mode = &v
	return s
}

func (s *ContinueCreateStackRequest) SetParallelism(v int64) *ContinueCreateStackRequest {
	s.Parallelism = &v
	return s
}

func (s *ContinueCreateStackRequest) SetParameters(v []*ContinueCreateStackRequestParameters) *ContinueCreateStackRequest {
	s.Parameters = v
	return s
}

func (s *ContinueCreateStackRequest) SetRamRoleName(v string) *ContinueCreateStackRequest {
	s.RamRoleName = &v
	return s
}

func (s *ContinueCreateStackRequest) SetRecreatingOptions(v []*string) *ContinueCreateStackRequest {
	s.RecreatingOptions = v
	return s
}

func (s *ContinueCreateStackRequest) SetRecreatingResources(v []*string) *ContinueCreateStackRequest {
	s.RecreatingResources = v
	return s
}

func (s *ContinueCreateStackRequest) SetRegionId(v string) *ContinueCreateStackRequest {
	s.RegionId = &v
	return s
}

func (s *ContinueCreateStackRequest) SetStackId(v string) *ContinueCreateStackRequest {
	s.StackId = &v
	return s
}

func (s *ContinueCreateStackRequest) SetTemplateBody(v string) *ContinueCreateStackRequest {
	s.TemplateBody = &v
	return s
}

func (s *ContinueCreateStackRequest) SetTemplateId(v string) *ContinueCreateStackRequest {
	s.TemplateId = &v
	return s
}

func (s *ContinueCreateStackRequest) SetTemplateURL(v string) *ContinueCreateStackRequest {
	s.TemplateURL = &v
	return s
}

func (s *ContinueCreateStackRequest) SetTemplateVersion(v string) *ContinueCreateStackRequest {
	s.TemplateVersion = &v
	return s
}

type ContinueCreateStackRequestParameters struct {
	// The name of template parameter N that you want to use to override a specific parameter. If you do not specify the name and value of a template parameter, ROS uses the name and value specified in the previous operation that was performed to create the stack. Maximum value of N: 200.
	//
	// > This parameter takes effect only when Mode is set to Recreate.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of template parameter N that you want to use to override a specific parameter. Maximum value of N: 200.
	//
	// For ROS stacks, the template parameters that you use to override specific parameters are subject to the following limits:
	//
	// *   You cannot change the condition values in the Conditions section of a template from true to false or from false to true.
	// *   The template parameters can be referenced only by resources that ROS continues to create.
	//
	// > This parameter takes effect only when Mode is set to Recreate.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s ContinueCreateStackRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s ContinueCreateStackRequestParameters) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackRequestParameters) SetParameterKey(v string) *ContinueCreateStackRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *ContinueCreateStackRequestParameters) SetParameterValue(v string) *ContinueCreateStackRequestParameters {
	s.ParameterValue = &v
	return s
}

type ContinueCreateStackResponseBody struct {
	// The validation result.
	DryRunResult *ContinueCreateStackResponseBodyDryRunResult `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty" type:"Struct"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stack ID.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s ContinueCreateStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ContinueCreateStackResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackResponseBody) SetDryRunResult(v *ContinueCreateStackResponseBodyDryRunResult) *ContinueCreateStackResponseBody {
	s.DryRunResult = v
	return s
}

func (s *ContinueCreateStackResponseBody) SetRequestId(v string) *ContinueCreateStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinueCreateStackResponseBody) SetStackId(v string) *ContinueCreateStackResponseBody {
	s.StackId = &v
	return s
}

type ContinueCreateStackResponseBodyDryRunResult struct {
	// The parameters that can be modified.
	ParametersAllowedToBeModified []*string `json:"ParametersAllowedToBeModified,omitempty" xml:"ParametersAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that can be modified under specific conditions.
	ParametersConditionallyAllowedToBeModified []*string `json:"ParametersConditionallyAllowedToBeModified,omitempty" xml:"ParametersConditionallyAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that cannot be modified.
	ParametersNotAllowedToBeModified []*string `json:"ParametersNotAllowedToBeModified,omitempty" xml:"ParametersNotAllowedToBeModified,omitempty" type:"Repeated"`
}

func (s ContinueCreateStackResponseBodyDryRunResult) String() string {
	return tea.Prettify(s)
}

func (s ContinueCreateStackResponseBodyDryRunResult) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackResponseBodyDryRunResult) SetParametersAllowedToBeModified(v []*string) *ContinueCreateStackResponseBodyDryRunResult {
	s.ParametersAllowedToBeModified = v
	return s
}

func (s *ContinueCreateStackResponseBodyDryRunResult) SetParametersConditionallyAllowedToBeModified(v []*string) *ContinueCreateStackResponseBodyDryRunResult {
	s.ParametersConditionallyAllowedToBeModified = v
	return s
}

func (s *ContinueCreateStackResponseBodyDryRunResult) SetParametersNotAllowedToBeModified(v []*string) *ContinueCreateStackResponseBodyDryRunResult {
	s.ParametersNotAllowedToBeModified = v
	return s
}

type ContinueCreateStackResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ContinueCreateStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ContinueCreateStackResponse) String() string {
	return tea.Prettify(s)
}

func (s ContinueCreateStackResponse) GoString() string {
	return s.String()
}

func (s *ContinueCreateStackResponse) SetHeaders(v map[string]*string) *ContinueCreateStackResponse {
	s.Headers = v
	return s
}

func (s *ContinueCreateStackResponse) SetStatusCode(v int32) *ContinueCreateStackResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinueCreateStackResponse) SetBody(v *ContinueCreateStackResponseBody) *ContinueCreateStackResponse {
	s.Body = v
	return s
}

type CreateChangeSetRequest struct {
	// The name of the change set.\
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). The name must start with a digit or a letter.
	//
	// > Make sure that the name is unique among all names of change sets that are associated with the specified stack.
	ChangeSetName *string `json:"ChangeSetName,omitempty" xml:"ChangeSetName,omitempty"`
	// The type of the change set. Valid values:
	//
	// *   CREATE: creates a change set for a new stack.
	// *   UPDATE (default): creates a change set for an existing stack.
	// *   IMPORT: creates a change set for a new stack or an existing stack to import resources that are not managed by ROS.
	//
	// If you create a change set for a new stack, ROS generates a unique stack ID for the stack. The stack remains in the REVIEW_IN_PROGRESS state until you execute the change set.\
	// If you want to create a change set for a new stack, do not set ChangeSetType to UPDATE. If you want to create a change set for an existing stack, do not set ChangeSetType to CREATE.
	ChangeSetType *string `json:"ChangeSetType,omitempty" xml:"ChangeSetType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests.\
	// The token can contain letters, digits, hyphens (-), and underscores (\_) and cannot exceed 64 characters in length.\
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the change set. The description can be up to 1,024 bytes in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to disable rollback when the stack fails to be created.\
	// Valid values:
	//
	// *   true: disables rollback for the stack when the stack fails to be created.
	// *   false (default): enables rollback for the stack when the stack fails to be created.
	//
	// > This parameter takes effect only if you set ChangeSetType to CREATE or IMPORT.
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The callback URLs that are used to receive stack events.
	NotificationURLs []*string `json:"NotificationURLs,omitempty" xml:"NotificationURLs,omitempty" type:"Repeated"`
	// The maximum number of concurrent operations that can be performed on resources. By default, this parameter is empty. You can set this parameter to an integer that is greater than or equal to 0. If you set this parameter to a specific value, ROS associates the value with the stack. The value can affect subsequent operations on the stack.
	//
	// This parameter takes effect only if you set ChangeSetType to CREATE or UPDATE.
	//
	// *   Valid values for change sets of the CREATE type:
	//
	//     *   If you set this parameter to an integer that is greater than 0, the integer is used.
	//     *   If you set this parameter to 0 or leave this parameter empty, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	//
	// *   Valid values for change sets of the UPDATE type:
	//
	//     *   If you set this parameter to an integer that is greater than 0, the integer is used.
	//     *   If you set this parameter to 0, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	//     *   If you leave this parameter empty, the value that you specified for this parameter in the previous request is used. If you left this parameter empty in the previous request, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	Parallelism *int64 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The parameters that are defined in the template.
	Parameters []*CreateChangeSetRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The name of the Resource Access Management (RAM) role. Resource Orchestration Service (ROS) assumes the RAM role to create the stack and uses the credentials of the role to call the APIs of Alibaba Cloud services.\
	// ROS assumes the role to perform operations on the stack. If you have permissions to perform operations on the stack but do not have permissions to use the RAM role, ROS still assumes the RAM role. You must make sure that the least privileges are granted to the RAM role.\
	// If you do not specify this parameter, ROS uses the existing role that is associated with the stack. If no roles are available for ROS to assume, ROS uses a temporary credential that is generated from the credentials of your Alibaba Cloud account.\
	// The name of the RAM role can be up to 64 bytes in length.
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the change set. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable replacement update if a resource property is changed and you cannot modify the new resource property. For a change, the physical ID of the resource remains unchanged. For a replacement update, the existing resource is deleted, a new resource is created, and the physical ID of the resource is changed. Valid values:
	//
	// *   Enabled
	// *   Disabled (default)
	//
	// > Operations that you perform to modify the resource properties for an update take precedence over operations you perform to replace the resource properties for an update. This parameter takes effect only if you set ChangeSetType to UPDATE.
	ReplacementOption *string `json:"ReplacementOption,omitempty" xml:"ReplacementOption,omitempty"`
	// The resources that you want to import into the stack.
	ResourcesToImport []*CreateChangeSetRequestResourcesToImport `json:"ResourcesToImport,omitempty" xml:"ResourcesToImport,omitempty" type:"Repeated"`
	// The ID of the stack for which you want to create a change set. ROS compares the stack information with the information that you submit, such as a modified template or a changed parameter value, to generate the change set.
	//
	// > This parameter takes effect only for change sets of the UPDATE or IMPORT.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The name of the stack for which you want to create the change set.\
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). The name must start with a digit or a letter.
	//
	// > This parameter takes effect only if you set ChangeSetType to CREATE or IMPORT.
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The structure of the stack policy body. The policy body must be 1 to 16,384 bytes in length. If you set ChangeSetType to CREATE, you can specify StackPolicyBody or StackPolicyURL. If you set ChangeSetType to UPDATE, you can specify only one of the following parameters:
	//
	// *   StackPolicyBody
	// *   StackPolicyURL
	// *   StackPolicyDuringUpdateBody
	// *   StackPolicyDuringUpdateURL
	StackPolicyBody *string `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	// The structure of the temporary overriding stack policy. The policy body must be 1 to 16,384 bytes in length.\
	// If you need to update protected resources, specify a temporary overriding stack policy for the resources during the update. If you do not specify a temporary overriding stack policy, the existing stack policy that is associated with the stack is used.\
	// This parameter takes effect only if you set ChangeSetType to UPDATE. You can specify only one of the following parameters:
	//
	// *   StackPolicyBody
	// *   StackPolicyURL
	// *   StackPolicyDuringUpdateBody
	// *   StackPolicyDuringUpdateURL
	StackPolicyDuringUpdateBody *string `json:"StackPolicyDuringUpdateBody,omitempty" xml:"StackPolicyDuringUpdateBody,omitempty"`
	// The URL of the stack policy based on which the stack is updated. The URL must point to a policy that is located on an HTTP or HTTPS web server or in an OSS bucket, such as oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy file can be up to 16,384 bytes in length.
	//
	// > If you do not specify the region of the OSS bucket, the value of RegionId is used.
	//
	// The URL can be up to 1,350 bytes in length.\
	// If you need to update protected resources, specify a temporary overriding stack policy for the resources during the update. If you do not specify a stack policy, the existing policy that is associated with the stack is used. This parameter takes effect only if you set ChangeSetType to UPDATE. You can specify only one of the following parameters:
	//
	// *   StackPolicyBody
	// *   StackPolicyURL
	// *   StackPolicyDuringUpdateBody
	// *   StackPolicyDuringUpdateURL
	StackPolicyDuringUpdateURL *string `json:"StackPolicyDuringUpdateURL,omitempty" xml:"StackPolicyDuringUpdateURL,omitempty"`
	// The URL of the file that contains the stack policy. The URL must point to a policy that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy file can be up to 16,384 bytes in length.
	//
	// > If you do not specify the region of the OSS bucket, the value of RegionId is used.
	//
	// You can specify only one of the following parameters: StackPolicyBody and StackPolicyURL.\
	// The URL can be up to 1,350 bytes in length.
	//
	// If you set ChangeSetType to CREATE, you can specify StackPolicyBody or StackPolicyURL. If you set ChangeSetType to UPDATE, you can specify only one of the following parameters:
	//
	// *   StackPolicyBody
	// *   StackPolicyURL
	// *   StackPolicyDuringUpdateBody
	// *   StackPolicyDuringUpdateURL
	StackPolicyURL *string `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
	TemplateBody   *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the scenario template.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an OSS bucket, such as oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length.
	//
	// > If you do not specify the region of the OSS bucket, the value of RegionId is used.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// The URL can be up to 1,024 bytes in length.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template.
	//
	// > This parameter takes effect only if you specify TemplateId.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The amount of time that can elapse before the stack enters the CREATE_FAILED or UPDATE_FAILED state.\
	// If you set ChangeSetType to CREATE, this parameter is required. If you set ChangeSetType to UPDATE, this parameter is optional.
	//
	// *   Unit: minutes.
	// *   Valid values: 10 to 1440.
	// *   Default value: 60.
	TimeoutInMinutes *int64 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	// Specifies whether to use the values that were passed last time for the parameters that you do not specify in the current request. Valid values:
	//
	// *   true
	// *   false (default)
	//
	// > This parameter takes effect only if you set ChangeSetType to UPDATE or IMPORT.
	UsePreviousParameters *bool `json:"UsePreviousParameters,omitempty" xml:"UsePreviousParameters,omitempty"`
}

func (s CreateChangeSetRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChangeSetRequest) GoString() string {
	return s.String()
}

func (s *CreateChangeSetRequest) SetChangeSetName(v string) *CreateChangeSetRequest {
	s.ChangeSetName = &v
	return s
}

func (s *CreateChangeSetRequest) SetChangeSetType(v string) *CreateChangeSetRequest {
	s.ChangeSetType = &v
	return s
}

func (s *CreateChangeSetRequest) SetClientToken(v string) *CreateChangeSetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateChangeSetRequest) SetDescription(v string) *CreateChangeSetRequest {
	s.Description = &v
	return s
}

func (s *CreateChangeSetRequest) SetDisableRollback(v bool) *CreateChangeSetRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateChangeSetRequest) SetNotificationURLs(v []*string) *CreateChangeSetRequest {
	s.NotificationURLs = v
	return s
}

func (s *CreateChangeSetRequest) SetParallelism(v int64) *CreateChangeSetRequest {
	s.Parallelism = &v
	return s
}

func (s *CreateChangeSetRequest) SetParameters(v []*CreateChangeSetRequestParameters) *CreateChangeSetRequest {
	s.Parameters = v
	return s
}

func (s *CreateChangeSetRequest) SetRamRoleName(v string) *CreateChangeSetRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateChangeSetRequest) SetRegionId(v string) *CreateChangeSetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateChangeSetRequest) SetReplacementOption(v string) *CreateChangeSetRequest {
	s.ReplacementOption = &v
	return s
}

func (s *CreateChangeSetRequest) SetResourcesToImport(v []*CreateChangeSetRequestResourcesToImport) *CreateChangeSetRequest {
	s.ResourcesToImport = v
	return s
}

func (s *CreateChangeSetRequest) SetStackId(v string) *CreateChangeSetRequest {
	s.StackId = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackName(v string) *CreateChangeSetRequest {
	s.StackName = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackPolicyBody(v string) *CreateChangeSetRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackPolicyDuringUpdateBody(v string) *CreateChangeSetRequest {
	s.StackPolicyDuringUpdateBody = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackPolicyDuringUpdateURL(v string) *CreateChangeSetRequest {
	s.StackPolicyDuringUpdateURL = &v
	return s
}

func (s *CreateChangeSetRequest) SetStackPolicyURL(v string) *CreateChangeSetRequest {
	s.StackPolicyURL = &v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateBody(v string) *CreateChangeSetRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateId(v string) *CreateChangeSetRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateScratchId(v string) *CreateChangeSetRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateURL(v string) *CreateChangeSetRequest {
	s.TemplateURL = &v
	return s
}

func (s *CreateChangeSetRequest) SetTemplateVersion(v string) *CreateChangeSetRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateChangeSetRequest) SetTimeoutInMinutes(v int64) *CreateChangeSetRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *CreateChangeSetRequest) SetUsePreviousParameters(v bool) *CreateChangeSetRequest {
	s.UsePreviousParameters = &v
	return s
}

type CreateChangeSetRequestParameters struct {
	// The key of parameter N that is defined in the template. If you do not specify the key and value of a parameter, ROS uses the default name and value that are defined in the template. Maximum value of N: 200.
	//
	// >  Parameters is optional. If you specify Parameters, you must also specify Parameters.N.ParameterKey.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N that is defined in the template. Maximum value of N: 200.
	//
	// >  Parameters is optional. If you specify Parameters, you must also specify Parameters.N.ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateChangeSetRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateChangeSetRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateChangeSetRequestParameters) SetParameterKey(v string) *CreateChangeSetRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateChangeSetRequestParameters) SetParameterValue(v string) *CreateChangeSetRequestParameters {
	s.ParameterValue = &v
	return s
}

type CreateChangeSetRequestResourcesToImport struct {
	// The logical ID of resource N that you want to import. The logical ID specifies the name of resource N that is defined in the template.
	//
	// > This parameter takes effect only if you set ChangeSetType to IMPORT. ResourcesToImport is optional. If you specify ResourcesToImport, you must also specify ResourcesToImport.N.LogicalResourceId.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The key-value mapping between strings. The value is a JSON string that identifies resource N that you want to import.\
	// A key is an identifier for a resource and a value is an assignment of data to the key. For example, VpcId is a key that indicates the ID of a virtual private cloud (VPC), and `vpc-2zevx9ios****` is a value that is assigned to VpcId. You can call the [GetTemplateSummary](~~172485~~) operation to obtain the key of a resource.
	//
	// > This parameter takes effect only if you set ChangeSetType to IMPORT. ResourcesToImport is optional. If you specify ResourcesToImport, you must also specify ResourcesToImport.N.ResourceIdentifier.
	ResourceIdentifier *string `json:"ResourceIdentifier,omitempty" xml:"ResourceIdentifier,omitempty"`
	// The type of resource N that you want to import. The resource type must be the same as the resource type that is defined in the template.
	//
	// > This parameter takes effect only if you set ChangeSetType to IMPORT. ResourcesToImport is optional. If you specify ResourcesToImport, you must also specify ResourcesToImport.N.ResourceType.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateChangeSetRequestResourcesToImport) String() string {
	return tea.Prettify(s)
}

func (s CreateChangeSetRequestResourcesToImport) GoString() string {
	return s.String()
}

func (s *CreateChangeSetRequestResourcesToImport) SetLogicalResourceId(v string) *CreateChangeSetRequestResourcesToImport {
	s.LogicalResourceId = &v
	return s
}

func (s *CreateChangeSetRequestResourcesToImport) SetResourceIdentifier(v string) *CreateChangeSetRequestResourcesToImport {
	s.ResourceIdentifier = &v
	return s
}

func (s *CreateChangeSetRequestResourcesToImport) SetResourceType(v string) *CreateChangeSetRequestResourcesToImport {
	s.ResourceType = &v
	return s
}

type CreateChangeSetResponseBody struct {
	// The ID of the change set.
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s CreateChangeSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateChangeSetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChangeSetResponseBody) SetChangeSetId(v string) *CreateChangeSetResponseBody {
	s.ChangeSetId = &v
	return s
}

func (s *CreateChangeSetResponseBody) SetRequestId(v string) *CreateChangeSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChangeSetResponseBody) SetStackId(v string) *CreateChangeSetResponseBody {
	s.StackId = &v
	return s
}

type CreateChangeSetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateChangeSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateChangeSetResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateChangeSetResponse) GoString() string {
	return s.String()
}

func (s *CreateChangeSetResponse) SetHeaders(v map[string]*string) *CreateChangeSetResponse {
	s.Headers = v
	return s
}

func (s *CreateChangeSetResponse) SetStatusCode(v int32) *CreateChangeSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateChangeSetResponse) SetBody(v *CreateChangeSetResponseBody) *CreateChangeSetResponse {
	s.Body = v
	return s
}

type CreateDiagnosticRequest struct {
	// The keyword in the diagnosis.
	DiagnosticKey *string `json:"DiagnosticKey,omitempty" xml:"DiagnosticKey,omitempty"`
	// The type of the item that is diagnosed. Set the value to Stack, which specifies that the stack is diagnosed.
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// The name of the product that is diagonosed.
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
}

func (s CreateDiagnosticRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticRequest) SetDiagnosticKey(v string) *CreateDiagnosticRequest {
	s.DiagnosticKey = &v
	return s
}

func (s *CreateDiagnosticRequest) SetDiagnosticType(v string) *CreateDiagnosticRequest {
	s.DiagnosticType = &v
	return s
}

func (s *CreateDiagnosticRequest) SetProduct(v string) *CreateDiagnosticRequest {
	s.Product = &v
	return s
}

type CreateDiagnosticResponseBody struct {
	// The error code returned.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The type of the item that is diagnosed.
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// The HTTP status code returned.
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the diagnostic report.
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   true
	// *   false
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDiagnosticResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticResponseBody) SetCode(v string) *CreateDiagnosticResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDiagnosticResponseBody) SetDiagnosticType(v string) *CreateDiagnosticResponseBody {
	s.DiagnosticType = &v
	return s
}

func (s *CreateDiagnosticResponseBody) SetHttpStatusCode(v int32) *CreateDiagnosticResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDiagnosticResponseBody) SetMessage(v string) *CreateDiagnosticResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDiagnosticResponseBody) SetReportId(v string) *CreateDiagnosticResponseBody {
	s.ReportId = &v
	return s
}

func (s *CreateDiagnosticResponseBody) SetRequestId(v string) *CreateDiagnosticResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiagnosticResponseBody) SetSuccess(v string) *CreateDiagnosticResponseBody {
	s.Success = &v
	return s
}

type CreateDiagnosticResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDiagnosticResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDiagnosticResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticResponse) SetHeaders(v map[string]*string) *CreateDiagnosticResponse {
	s.Headers = v
	return s
}

func (s *CreateDiagnosticResponse) SetStatusCode(v int32) *CreateDiagnosticResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiagnosticResponse) SetBody(v *CreateDiagnosticResponseBody) *CreateDiagnosticResponse {
	s.Body = v
	return s
}

type CreateStackRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// For more information, see [Ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The creation option for the stack. Valid values:
	//
	// *   KeepStackOnCreationComplete (default): After the stack is created, the stack and its resources are retained. The quota for the maximum number of stacks that can be created in ROS is consumed.
	// *   AbandonStackOnCreationComplete: After the stack is created, the stack is deleted, but its resources are retained. The quota for the maximum number of stacks that can be created in ROS is not consumed. If the stack fails to be created, the stack is retained.
	// *   AbandonStackOnCreationRollbackComplete: When the resources of the stack are rolled back after the stack fails to be created, the stack is deleted. The quota for the maximum number of stacks that can be created in ROS is not consumed. In other rollback scenarios, the stack is retained.
	// *   ManuallyPay: When you create the stack, you must manually pay for the subscription resources that are used. The following resource types support manual payment: `ALIYUN::ECS::InstanceGroup`, `ALIYUN::RDS::DBInstance`, `ALIYUN::SLB::LoadBalancer`, `ALIYUN::VPC::EIP`, and `ALIYUN::VPC::VpnGateway`.
	//
	// >  You can specify only one of CreateOption and CreateOptions.
	CreateOption *string `json:"CreateOption,omitempty" xml:"CreateOption,omitempty"`
	// The creation options for the stack.
	CreateOptions []*string `json:"CreateOptions,omitempty" xml:"CreateOptions,omitempty" type:"Repeated"`
	// Specifies whether to enable deletion protection for the stack. Valid values:
	//
	// *   Enabled.
	// *   Disabled (default). If deletion protection is disabled, you can delete the stack by using the ROS console or by calling the DeleteStack operation.
	//
	// > The value of DeletionProtection that you specify for the root stack applies to its nested stacks.
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Specifies whether to disable rollback for the resources when the stack fails to be created.
	//
	// Valid values:
	//
	// *   true
	// *   false (default)
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The callback URLs that are used to receive stack events. Valid values:
	//
	// *   HTTP POST URL
	//
	// Each URL can be up to 1,024 bytes in length.
	//
	// *   eventbridge
	//
	// When the status of a stack changes, ROS sends notifications to the EventBridge service. You can view the event information in the [EventBridge](https://eventbridge.console.aliyun.com) console.
	//
	// > This feature is supported in the China (Hangzhou), China (Shanghai), China (Beijing), China (Hong Kong), and China (Zhangjiakou) regions.
	//
	// Maximum value of N: 5. When the status of a stack changes, ROS sends a notification to the specified URL. When rollback is enabled for the stack, notifications are sent if the stack is in the CREATE_ROLLBACK or ROLLBACK state, but are not sent if the stack is in the CREATE_FAILED, UPDATE_FAILED, or IN_PROGRESS state.\
	// ROS sends notifications regardless of whether you specify the Outputs section. The following sample code provides an example on the content of a notification:
	//
	//     {
	//        "Outputs": [
	//            {
	//                "Description": "No description given",
	//                "OutputKey": "InstanceId",
	//                "OutputValue": "i-xxx"
	//            }
	//        ],
	//        "StackId": "80bd6b6c-e888-4573-ae3b-93d29113****",
	//        "StackName": "test-notification-url",
	//        "Status": "CREATE_COMPLETE"
	//     }
	NotificationURLs []*string `json:"NotificationURLs,omitempty" xml:"NotificationURLs,omitempty" type:"Repeated"`
	// The maximum number of concurrent operations that can be performed on resources.
	//
	// By default, this parameter is empty. You can set this parameter to an integer that is greater than or equal to 0.
	//
	//
	//
	// > -  If you set this parameter to an integer that is greater than 0, the integer is used. If you set this parameter to 0 or leave this parameter empty, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	// > -  If you set this parameter to a specific value, ROS associates the value with the stack. The value affects subsequent operations on the stack, such as an update operation.
	Parallelism *int64 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The parameters that are defined in the template.
	Parameters []*CreateStackRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The name of the RAM role. ROS assumes the RAM role to create the stack and uses the credentials of the role to call the APIs of Alibaba Cloud services.\
	// ROS assumes the RAM role to perform operations on the stack. If you have permissions to perform operations on the stack but do not have permissions to use the RAM role, ROS still assumes the RAM role. You must make sure that the least privileges are granted to the RAM role.
	//
	// If you do not specify this parameter, ROS assumes the existing role that is associated with the stack. If no roles are available, ROS uses a temporary credential that is generated from the credentials of your account.
	//
	// The RAM role name can be up to 64 characters in length.
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. If you leave this parameter empty, the stack is added to the default resource group.
	//
	// For more information about resource groups, see the "Resource group" section of the [What is Resource Management?](~~94475~~) topic.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the stack.\
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a letter.
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The structure that contains the stack policy body. The policy body must be 1 to 16,384 bytes in length.
	//
	// > You can specify only one of StackPolicyBody and StackPolicyURL.
	StackPolicyBody *string `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	// The URL of the file that contains the stack policy. The URL must point to a policy that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo or oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy file can be up to 16,384 bytes in length. If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// > You can specify only one of StackPolicyBody and StackPolicyURL.
	//
	// The URL can be up to 1,350 bytes in length.
	StackPolicyURL *string `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
	// The tags that you want to add to the stack.
	Tags         []*CreateStackRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateBody *string                   `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The template ID. This parameter applies to shared templates and private templates.
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The scenario ID.
	//
	// For more information about how to query the scenario ID, see [ListTemplateScratches](~~363050~~).
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The region ID of the scenario. The default value is the same as the value of RegionId.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	TemplateScratchRegionId *string `json:"TemplateScratchRegionId,omitempty" xml:"TemplateScratchRegionId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an OSS bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length. If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. This parameter takes effect only when TemplateId is specified.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The timeout period for creating the stack.
	//
	// *   Default value: 60.
	// *   Unit: minutes.
	// *   Valid values: 10 to 1440.
	TimeoutInMinutes *int64 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s CreateStackRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStackRequest) GoString() string {
	return s.String()
}

func (s *CreateStackRequest) SetClientToken(v string) *CreateStackRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStackRequest) SetCreateOption(v string) *CreateStackRequest {
	s.CreateOption = &v
	return s
}

func (s *CreateStackRequest) SetCreateOptions(v []*string) *CreateStackRequest {
	s.CreateOptions = v
	return s
}

func (s *CreateStackRequest) SetDeletionProtection(v string) *CreateStackRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateStackRequest) SetDisableRollback(v bool) *CreateStackRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateStackRequest) SetNotificationURLs(v []*string) *CreateStackRequest {
	s.NotificationURLs = v
	return s
}

func (s *CreateStackRequest) SetParallelism(v int64) *CreateStackRequest {
	s.Parallelism = &v
	return s
}

func (s *CreateStackRequest) SetParameters(v []*CreateStackRequestParameters) *CreateStackRequest {
	s.Parameters = v
	return s
}

func (s *CreateStackRequest) SetRamRoleName(v string) *CreateStackRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateStackRequest) SetRegionId(v string) *CreateStackRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStackRequest) SetResourceGroupId(v string) *CreateStackRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateStackRequest) SetStackName(v string) *CreateStackRequest {
	s.StackName = &v
	return s
}

func (s *CreateStackRequest) SetStackPolicyBody(v string) *CreateStackRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *CreateStackRequest) SetStackPolicyURL(v string) *CreateStackRequest {
	s.StackPolicyURL = &v
	return s
}

func (s *CreateStackRequest) SetTags(v []*CreateStackRequestTags) *CreateStackRequest {
	s.Tags = v
	return s
}

func (s *CreateStackRequest) SetTemplateBody(v string) *CreateStackRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateStackRequest) SetTemplateId(v string) *CreateStackRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateStackRequest) SetTemplateScratchId(v string) *CreateStackRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *CreateStackRequest) SetTemplateScratchRegionId(v string) *CreateStackRequest {
	s.TemplateScratchRegionId = &v
	return s
}

func (s *CreateStackRequest) SetTemplateURL(v string) *CreateStackRequest {
	s.TemplateURL = &v
	return s
}

func (s *CreateStackRequest) SetTemplateVersion(v string) *CreateStackRequest {
	s.TemplateVersion = &v
	return s
}

func (s *CreateStackRequest) SetTimeoutInMinutes(v int64) *CreateStackRequest {
	s.TimeoutInMinutes = &v
	return s
}

type CreateStackRequestParameters struct {
	// The key of parameter N that is defined in the template. If you do not specify the name and value of a parameter, ROS uses the default name and value that are specified in the template.
	//
	// Maximum value of N: 200.\
	// The name must be 1 to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N that is defined in the template.
	//
	// Maximum value of N: 200.\
	// The value can be up to 128 characters in length, and cannot contain `http://` or `https://`. It cannot start with `aliyun` or `acs:`.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateStackRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateStackRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateStackRequestParameters) SetParameterKey(v string) *CreateStackRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackRequestParameters) SetParameterValue(v string) *CreateStackRequestParameters {
	s.ParameterValue = &v
	return s
}

type CreateStackRequestTags struct {
	// The key of tag N that you want to add to the stack.
	//
	// Valid values of N: 1 to 20.
	//
	// > - The Tags parameter is optional. If you specify Tags, you must specify Tags.N.Key.
	// > -  The tag of a stack is propagated to each resource that supports the tag feature in the stack. For more information, see [Propagate tags](~~201421~~).
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that you want to add to the stack.
	//
	// Valid values of N: 1 to 20.
	//
	// > The tag of a stack is propagated to each resource that supports the tag feature in the stack. For more information, see [Propagate tags](~~201421~~).
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateStackRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateStackRequestTags) GoString() string {
	return s.String()
}

func (s *CreateStackRequestTags) SetKey(v string) *CreateStackRequestTags {
	s.Key = &v
	return s
}

func (s *CreateStackRequestTags) SetValue(v string) *CreateStackRequestTags {
	s.Value = &v
	return s
}

type CreateStackResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stack ID.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s CreateStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStackResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStackResponseBody) SetRequestId(v string) *CreateStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStackResponseBody) SetStackId(v string) *CreateStackResponseBody {
	s.StackId = &v
	return s
}

type CreateStackResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStackResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStackResponse) GoString() string {
	return s.String()
}

func (s *CreateStackResponse) SetHeaders(v map[string]*string) *CreateStackResponse {
	s.Headers = v
	return s
}

func (s *CreateStackResponse) SetStatusCode(v int32) *CreateStackResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStackResponse) SetBody(v *CreateStackResponseBody) *CreateStackResponse {
	s.Body = v
	return s
}

type CreateStackGroupRequest struct {
	// The name of the RAM role that you specify for the administrator account when you create a self-managed stack group. ROS assumes the administrator role to perform operations. If you do not specify this parameter, AliyunROSStackGroupAdministrationRole is used as the default value. ROS uses the administrator role to assume the execution role AliyunROSStackGroupExecutionRole to perform operations on the stacks in the stack group.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, and hyphens (-).
	AdministrationRoleName *string `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	// The information about automatic deployment settings.
	//
	// > You must specify this parameter if PermissionModel is set to SERVICE_MANAGED.
	AutoDeployment *CreateStackGroupRequestAutoDeployment `json:"AutoDeployment,omitempty" xml:"AutoDeployment,omitempty" type:"Struct"`
	// The options for the stack group. You can specify up to one option.
	Capabilities []*string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\
	// The token can contain letters, digits, underscores (\_), and hyphens (-) and cannot exceed 64 characters in length.\
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the stack group.\
	// The description must be 1 to 256 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the RAM role that you specify for the execution account when you create a self-managed stack group. The administrator role AliyunROSStackGroupAdministrationRole assumes the execution role to perform operations. If you do not specify this parameter, AliyunROSStackGroupExecutionRole is used as the default value. ROS assumes the execution role to perform operations on the stacks in the stack group.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, and hyphens (-).
	ExecutionRoleName *string `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	// The parameters of the stack group.
	Parameters []*CreateStackGroupRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The permission model of the stack group.
	//
	// Valid values:
	//
	// *   SELF_MANAGED (default): the self-managed permission model. If you create a self-managed stack group, you must create RAM roles within the administrator and execution accounts and establish a trust relationship between the accounts. Then, you can deploy stacks within the execution account.
	// *   SERVICE_MANAGED: the service-managed permission model. If you create a service-managed stack group, ROS creates service-linked roles for the administrator and execution accounts, and the administrator account uses its role to deploy stacks within the execution account.
	//
	// > If you want to use the service-managed permission model to deploy stacks, your account must be the management account or a delegated administrator account of your resource directory and the trusted access feature is enabled for the account. For more information, see [Manage a delegated administrator account](~~308253~~) and [Enable trusted access](~~298229~~).
	PermissionModel *string `json:"PermissionModel,omitempty" xml:"PermissionModel,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. If you do not specify this parameter, the stack group is added to the default resource group.\
	// For more information about resource groups, see [Resource groups](~~94475~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the stack group. The name must be unique within a region.\
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or a letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The tags of the stack group.
	Tags         []*CreateStackGroupRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateBody *string                        `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Alibaba Cloud Object Storage Service (OSS) bucket. The template body must be 1 to 524,288 bytes in length. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. If you do not specify this parameter, the latest version is used.
	//
	// > TemplateVersion takes effect only if you specify TemplateId.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s CreateStackGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStackGroupRequest) GoString() string {
	return s.String()
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

type CreateStackGroupRequestAutoDeployment struct {
	// Indicates whether automatic deployment is enabled.
	//
	// Valid values:
	//
	// *   true: Automatic deployment is enabled. If you add a member account to the folder to which the stack group belongs after you enable automatic deployment, ROS automatically adds the stacks in the stack group to the member account. If you remove a member account from the folder, ROS automatically deletes the stacks from the member account.
	// *   false: Automatic deployment is disabled. After you disable automatic deployment, the stacks remain unchanged when you add member accounts to or remove member accounts from the folder.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether the stacks within a member account are retained when you remove the member account from the folder.
	//
	// Valid values:
	//
	// *   true
	// *   false
	//
	// > You must specify RetainStacksOnAccountRemoval if Enabled is set to true.
	RetainStacksOnAccountRemoval *bool `json:"RetainStacksOnAccountRemoval,omitempty" xml:"RetainStacksOnAccountRemoval,omitempty"`
}

func (s CreateStackGroupRequestAutoDeployment) String() string {
	return tea.Prettify(s)
}

func (s CreateStackGroupRequestAutoDeployment) GoString() string {
	return s.String()
}

func (s *CreateStackGroupRequestAutoDeployment) SetEnabled(v bool) *CreateStackGroupRequestAutoDeployment {
	s.Enabled = &v
	return s
}

func (s *CreateStackGroupRequestAutoDeployment) SetRetainStacksOnAccountRemoval(v bool) *CreateStackGroupRequestAutoDeployment {
	s.RetainStacksOnAccountRemoval = &v
	return s
}

type CreateStackGroupRequestParameters struct {
	// The key of parameter N. If you do not specify the key and value of a parameter, ROS uses the default name and value that are defined in the template.
	//
	// Maximum value of N: 200.
	//
	// > Parameters is optional. If you specify Parameters, you must also specify Parameters.N.ParameterKey.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N.
	//
	// Maximum value of N: 200.
	//
	// > Parameters is optional. If you specify Parameters, you must also specify Parameters.N.ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateStackGroupRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateStackGroupRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateStackGroupRequestParameters) SetParameterKey(v string) *CreateStackGroupRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackGroupRequestParameters) SetParameterValue(v string) *CreateStackGroupRequestParameters {
	s.ParameterValue = &v
	return s
}

type CreateStackGroupRequestTags struct {
	// The tag key of the stack group.
	//
	// > Tags is optional. If you want to specify Tags, you must also specify Tags.N.Key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the stack group.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateStackGroupRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateStackGroupRequestTags) GoString() string {
	return s.String()
}

func (s *CreateStackGroupRequestTags) SetKey(v string) *CreateStackGroupRequestTags {
	s.Key = &v
	return s
}

func (s *CreateStackGroupRequestTags) SetValue(v string) *CreateStackGroupRequestTags {
	s.Value = &v
	return s
}

type CreateStackGroupShrinkRequest struct {
	// The name of the RAM role that you specify for the administrator account when you create a self-managed stack group. ROS assumes the administrator role to perform operations. If you do not specify this parameter, AliyunROSStackGroupAdministrationRole is used as the default value. ROS uses the administrator role to assume the execution role AliyunROSStackGroupExecutionRole to perform operations on the stacks in the stack group.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, and hyphens (-).
	AdministrationRoleName *string `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	// The information about automatic deployment settings.
	//
	// > You must specify this parameter if PermissionModel is set to SERVICE_MANAGED.
	AutoDeploymentShrink *string `json:"AutoDeployment,omitempty" xml:"AutoDeployment,omitempty"`
	// The options for the stack group. You can specify up to one option.
	Capabilities []*string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\
	// The token can contain letters, digits, underscores (\_), and hyphens (-) and cannot exceed 64 characters in length.\
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the stack group.\
	// The description must be 1 to 256 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the RAM role that you specify for the execution account when you create a self-managed stack group. The administrator role AliyunROSStackGroupAdministrationRole assumes the execution role to perform operations. If you do not specify this parameter, AliyunROSStackGroupExecutionRole is used as the default value. ROS assumes the execution role to perform operations on the stacks in the stack group.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, and hyphens (-).
	ExecutionRoleName *string `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	// The parameters of the stack group.
	Parameters []*CreateStackGroupShrinkRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The permission model of the stack group.
	//
	// Valid values:
	//
	// *   SELF_MANAGED (default): the self-managed permission model. If you create a self-managed stack group, you must create RAM roles within the administrator and execution accounts and establish a trust relationship between the accounts. Then, you can deploy stacks within the execution account.
	// *   SERVICE_MANAGED: the service-managed permission model. If you create a service-managed stack group, ROS creates service-linked roles for the administrator and execution accounts, and the administrator account uses its role to deploy stacks within the execution account.
	//
	// > If you want to use the service-managed permission model to deploy stacks, your account must be the management account or a delegated administrator account of your resource directory and the trusted access feature is enabled for the account. For more information, see [Manage a delegated administrator account](~~308253~~) and [Enable trusted access](~~298229~~).
	PermissionModel *string `json:"PermissionModel,omitempty" xml:"PermissionModel,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. If you do not specify this parameter, the stack group is added to the default resource group.\
	// For more information about resource groups, see [Resource groups](~~94475~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the stack group. The name must be unique within a region.\
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or a letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The tags of the stack group.
	Tags         []*CreateStackGroupShrinkRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateBody *string                              `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Alibaba Cloud Object Storage Service (OSS) bucket. The template body must be 1 to 524,288 bytes in length. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. If you do not specify this parameter, the latest version is used.
	//
	// > TemplateVersion takes effect only if you specify TemplateId.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s CreateStackGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStackGroupShrinkRequest) GoString() string {
	return s.String()
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

type CreateStackGroupShrinkRequestParameters struct {
	// The key of parameter N. If you do not specify the key and value of a parameter, ROS uses the default name and value that are defined in the template.
	//
	// Maximum value of N: 200.
	//
	// > Parameters is optional. If you specify Parameters, you must also specify Parameters.N.ParameterKey.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N.
	//
	// Maximum value of N: 200.
	//
	// > Parameters is optional. If you specify Parameters, you must also specify Parameters.N.ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateStackGroupShrinkRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateStackGroupShrinkRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateStackGroupShrinkRequestParameters) SetParameterKey(v string) *CreateStackGroupShrinkRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackGroupShrinkRequestParameters) SetParameterValue(v string) *CreateStackGroupShrinkRequestParameters {
	s.ParameterValue = &v
	return s
}

type CreateStackGroupShrinkRequestTags struct {
	// The tag key of the stack group.
	//
	// > Tags is optional. If you want to specify Tags, you must also specify Tags.N.Key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the stack group.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateStackGroupShrinkRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateStackGroupShrinkRequestTags) GoString() string {
	return s.String()
}

func (s *CreateStackGroupShrinkRequestTags) SetKey(v string) *CreateStackGroupShrinkRequestTags {
	s.Key = &v
	return s
}

func (s *CreateStackGroupShrinkRequestTags) SetValue(v string) *CreateStackGroupShrinkRequestTags {
	s.Value = &v
	return s
}

type CreateStackGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the stack group.
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
}

func (s CreateStackGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStackGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStackGroupResponseBody) SetRequestId(v string) *CreateStackGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStackGroupResponseBody) SetStackGroupId(v string) *CreateStackGroupResponseBody {
	s.StackGroupId = &v
	return s
}

type CreateStackGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateStackGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStackGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStackGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateStackGroupResponse) SetHeaders(v map[string]*string) *CreateStackGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateStackGroupResponse) SetStatusCode(v int32) *CreateStackGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStackGroupResponse) SetBody(v *CreateStackGroupResponseBody) *CreateStackGroupResponse {
	s.Body = v
	return s
}

type CreateStackInstancesRequest struct {
	// The IDs of the execution accounts within which you want to deploy stacks in self-managed mode. You can specify up to 20 execution account IDs.
	//
	// > You must specify one of the following parameters: `AccountIds` and `DeploymentTargets`.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\
	// The token can contain letters, digits, hyphens (-), and underscores (\_), and cannot exceed 64 characters in length.\
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The folders in which ROS deploy stacks in service-managed permission model.
	//
	// > You must specify one of the following parameters: `AccountIds` and `DeploymentTargets`.
	DeploymentTargets *CreateStackInstancesRequestDeploymentTargets `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty" type:"Struct"`
	// Specifies whether to disable rollback when the stack fails to be created.
	//
	// Valid values:
	//
	// *   true
	// *   false (default)
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The description of the stack creation operation.
	//
	// The description must be 1 to 256 characters in length.
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The preference settings of the stack creation operation.
	//
	// The following parameters are available:
	//
	// -  {"FailureToleranceCount": N}
	//
	//     The number of accounts within which stack operation failures are allowed in each region. If the value of this parameter is exceeded in a region, Resource Orchestration Service (ROS) stops the operation in the region. If ROS stops the operation in one region, ROS stops the operation in other regions.
	//
	//     Valid values of N: 0 to 20.
	//
	//     If you do not specify FailureToleranceCount, 0 is used as the default value.
	//
	// -  {"FailureTolerancePercentage": N}
	//
	//     The percentage of the number of accounts within which stack operation failures are allowed to the total number of accounts in each region. If the value of this parameter is exceeded, ROS stops the operation in the region.
	//
	//     Valid values of N: 0 to 100. If the numeric value in the percentage is not an integer, ROS rounds the value down to the nearest integer.
	//
	//     If you do not specify FailureTolerancePercentage, 0 is used as the default value.
	//
	// -  {"MaxConcurrentCount": N}
	//
	//    The maximum number of accounts within which multiple stacks are deployed at the same time in each region.
	//
	//    Valid values of N: 1 to 20.
	//
	//    If you do not specify MaxConcurrentCount, 1 is used as the default value.
	//
	// -  {"MaxConcurrentPercentage": N}
	//
	//     The percentage of the maximum number of accounts within which multiple stacks are deployed at the same time to the total number of accounts in each region.
	//
	//     Valid values: 1 to 100. If the numeric value in the percentage is not an integer, ROS rounds the number down to the nearest integer.
	//
	//     If you do not specify MaxConcurrentPercentage, 1 is used as the default value.
	//
	// -  {"RegionConcurrencyType": N}\
	//     The mode that you want to use to deploy stacks across regions. Valid values:
	//    - SEQUENTIAL (default): deploys stacks in each specified region based on the specified sequence of regions. ROS deploys stacks in one region at a time.
	//    - PARALLEL: deploys stacks in parallel across all specified regions.
	//
	// Separate multiple parameters with commas (,).
	//
	// >-  You can specify only one of the following parameters: MaxConcurrentCount and MaxConcurrentPercentage.
	// >-  You can specify only one of the following parameters: FailureToleranceCount and FailureTolerancePercentage.
	OperationPreferences map[string]interface{} `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// The parameters that are used to override specific parameters.
	ParameterOverrides []*CreateStackInstancesRequestParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the regions where you want to create the stacks. You can specify up to 20 region IDs.
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// The name of the stack group. The name must be unique within a region.\
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or a letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The timeout period within which you can create the stack.
	//
	// *   Default value: 60.
	// *   Unit: minutes.
	TimeoutInMinutes *int64 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s CreateStackInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStackInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesRequest) SetAccountIds(v []*string) *CreateStackInstancesRequest {
	s.AccountIds = v
	return s
}

func (s *CreateStackInstancesRequest) SetClientToken(v string) *CreateStackInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStackInstancesRequest) SetDeploymentTargets(v *CreateStackInstancesRequestDeploymentTargets) *CreateStackInstancesRequest {
	s.DeploymentTargets = v
	return s
}

func (s *CreateStackInstancesRequest) SetDisableRollback(v bool) *CreateStackInstancesRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateStackInstancesRequest) SetOperationDescription(v string) *CreateStackInstancesRequest {
	s.OperationDescription = &v
	return s
}

func (s *CreateStackInstancesRequest) SetOperationPreferences(v map[string]interface{}) *CreateStackInstancesRequest {
	s.OperationPreferences = v
	return s
}

func (s *CreateStackInstancesRequest) SetParameterOverrides(v []*CreateStackInstancesRequestParameterOverrides) *CreateStackInstancesRequest {
	s.ParameterOverrides = v
	return s
}

func (s *CreateStackInstancesRequest) SetRegionId(v string) *CreateStackInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStackInstancesRequest) SetRegionIds(v []*string) *CreateStackInstancesRequest {
	s.RegionIds = v
	return s
}

func (s *CreateStackInstancesRequest) SetStackGroupName(v string) *CreateStackInstancesRequest {
	s.StackGroupName = &v
	return s
}

func (s *CreateStackInstancesRequest) SetTimeoutInMinutes(v int64) *CreateStackInstancesRequest {
	s.TimeoutInMinutes = &v
	return s
}

type CreateStackInstancesRequestDeploymentTargets struct {
	// The folder IDs of the resource directory. You can add up to five folder IDs.
	//
	// You can create stacks within all the member accounts in the specified folders. If you create stacks in the Root folder, the stacks are created within all member accounts in the resource directory.
	//
	// > To view the folder IDs, go to the **Overview** page in the **Resource Management** console. For more information, see [View the basic information about a folder](~~111223~~).
	RdFolderIds []*string `json:"RdFolderIds,omitempty" xml:"RdFolderIds,omitempty" type:"Repeated"`
}

func (s CreateStackInstancesRequestDeploymentTargets) String() string {
	return tea.Prettify(s)
}

func (s CreateStackInstancesRequestDeploymentTargets) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesRequestDeploymentTargets) SetRdFolderIds(v []*string) *CreateStackInstancesRequestDeploymentTargets {
	s.RdFolderIds = v
	return s
}

type CreateStackInstancesRequestParameterOverrides struct {
	// The key of parameter N that you want to use to override a specific parameter. If you do not specify this parameter, ROS uses the name that you specified when you created the stack group.
	//
	// Maximum value of N: 200.
	//
	// >-   ParameterOverrides is optional.
	// >-   If you specify ParameterOverrides, you must specify ParameterOverrides.N.ParameterKey and ParameterOverrides.N.ParameterValue.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N that you want to use to override a specific parameter. If you do not specify this parameter, ROS uses the value that you specify when you create the stack group.
	//
	// Maximum value of N: 200.
	//
	// >-  ParameterOverrides is optional.
	// >-  If you specify ParameterOverrides, you must specify ParameterOverrides.N.ParameterKey and ParameterOverrides.N.ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateStackInstancesRequestParameterOverrides) String() string {
	return tea.Prettify(s)
}

func (s CreateStackInstancesRequestParameterOverrides) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesRequestParameterOverrides) SetParameterKey(v string) *CreateStackInstancesRequestParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackInstancesRequestParameterOverrides) SetParameterValue(v string) *CreateStackInstancesRequestParameterOverrides {
	s.ParameterValue = &v
	return s
}

type CreateStackInstancesShrinkRequest struct {
	// The IDs of the execution accounts within which you want to deploy stacks in self-managed mode. You can specify up to 20 execution account IDs.
	//
	// > You must specify one of the following parameters: `AccountIds` and `DeploymentTargets`.
	AccountIdsShrink *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\
	// The token can contain letters, digits, hyphens (-), and underscores (\_), and cannot exceed 64 characters in length.\
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The folders in which ROS deploy stacks in service-managed permission model.
	//
	// > You must specify one of the following parameters: `AccountIds` and `DeploymentTargets`.
	DeploymentTargetsShrink *string `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty"`
	// Specifies whether to disable rollback when the stack fails to be created.
	//
	// Valid values:
	//
	// *   true
	// *   false (default)
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The description of the stack creation operation.
	//
	// The description must be 1 to 256 characters in length.
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The preference settings of the stack creation operation.
	//
	// The following parameters are available:
	//
	// -  {"FailureToleranceCount": N}
	//
	//     The number of accounts within which stack operation failures are allowed in each region. If the value of this parameter is exceeded in a region, Resource Orchestration Service (ROS) stops the operation in the region. If ROS stops the operation in one region, ROS stops the operation in other regions.
	//
	//     Valid values of N: 0 to 20.
	//
	//     If you do not specify FailureToleranceCount, 0 is used as the default value.
	//
	// -  {"FailureTolerancePercentage": N}
	//
	//     The percentage of the number of accounts within which stack operation failures are allowed to the total number of accounts in each region. If the value of this parameter is exceeded, ROS stops the operation in the region.
	//
	//     Valid values of N: 0 to 100. If the numeric value in the percentage is not an integer, ROS rounds the value down to the nearest integer.
	//
	//     If you do not specify FailureTolerancePercentage, 0 is used as the default value.
	//
	// -  {"MaxConcurrentCount": N}
	//
	//    The maximum number of accounts within which multiple stacks are deployed at the same time in each region.
	//
	//    Valid values of N: 1 to 20.
	//
	//    If you do not specify MaxConcurrentCount, 1 is used as the default value.
	//
	// -  {"MaxConcurrentPercentage": N}
	//
	//     The percentage of the maximum number of accounts within which multiple stacks are deployed at the same time to the total number of accounts in each region.
	//
	//     Valid values: 1 to 100. If the numeric value in the percentage is not an integer, ROS rounds the number down to the nearest integer.
	//
	//     If you do not specify MaxConcurrentPercentage, 1 is used as the default value.
	//
	// -  {"RegionConcurrencyType": N}\
	//     The mode that you want to use to deploy stacks across regions. Valid values:
	//    - SEQUENTIAL (default): deploys stacks in each specified region based on the specified sequence of regions. ROS deploys stacks in one region at a time.
	//    - PARALLEL: deploys stacks in parallel across all specified regions.
	//
	// Separate multiple parameters with commas (,).
	//
	// >-  You can specify only one of the following parameters: MaxConcurrentCount and MaxConcurrentPercentage.
	// >-  You can specify only one of the following parameters: FailureToleranceCount and FailureTolerancePercentage.
	OperationPreferencesShrink *string `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// The parameters that are used to override specific parameters.
	ParameterOverrides []*CreateStackInstancesShrinkRequestParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the regions where you want to create the stacks. You can specify up to 20 region IDs.
	RegionIdsShrink *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	// The name of the stack group. The name must be unique within a region.\
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or a letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The timeout period within which you can create the stack.
	//
	// *   Default value: 60.
	// *   Unit: minutes.
	TimeoutInMinutes *int64 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s CreateStackInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStackInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesShrinkRequest) SetAccountIdsShrink(v string) *CreateStackInstancesShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetClientToken(v string) *CreateStackInstancesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetDeploymentTargetsShrink(v string) *CreateStackInstancesShrinkRequest {
	s.DeploymentTargetsShrink = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetDisableRollback(v bool) *CreateStackInstancesShrinkRequest {
	s.DisableRollback = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetOperationDescription(v string) *CreateStackInstancesShrinkRequest {
	s.OperationDescription = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetOperationPreferencesShrink(v string) *CreateStackInstancesShrinkRequest {
	s.OperationPreferencesShrink = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetParameterOverrides(v []*CreateStackInstancesShrinkRequestParameterOverrides) *CreateStackInstancesShrinkRequest {
	s.ParameterOverrides = v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetRegionId(v string) *CreateStackInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetRegionIdsShrink(v string) *CreateStackInstancesShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetStackGroupName(v string) *CreateStackInstancesShrinkRequest {
	s.StackGroupName = &v
	return s
}

func (s *CreateStackInstancesShrinkRequest) SetTimeoutInMinutes(v int64) *CreateStackInstancesShrinkRequest {
	s.TimeoutInMinutes = &v
	return s
}

type CreateStackInstancesShrinkRequestParameterOverrides struct {
	// The key of parameter N that you want to use to override a specific parameter. If you do not specify this parameter, ROS uses the name that you specified when you created the stack group.
	//
	// Maximum value of N: 200.
	//
	// >-   ParameterOverrides is optional.
	// >-   If you specify ParameterOverrides, you must specify ParameterOverrides.N.ParameterKey and ParameterOverrides.N.ParameterValue.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N that you want to use to override a specific parameter. If you do not specify this parameter, ROS uses the value that you specify when you create the stack group.
	//
	// Maximum value of N: 200.
	//
	// >-  ParameterOverrides is optional.
	// >-  If you specify ParameterOverrides, you must specify ParameterOverrides.N.ParameterKey and ParameterOverrides.N.ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateStackInstancesShrinkRequestParameterOverrides) String() string {
	return tea.Prettify(s)
}

func (s CreateStackInstancesShrinkRequestParameterOverrides) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesShrinkRequestParameterOverrides) SetParameterKey(v string) *CreateStackInstancesShrinkRequestParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *CreateStackInstancesShrinkRequestParameterOverrides) SetParameterValue(v string) *CreateStackInstancesShrinkRequestParameterOverrides {
	s.ParameterValue = &v
	return s
}

type CreateStackInstancesResponseBody struct {
	// The ID of the operation.
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateStackInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesResponseBody) SetOperationId(v string) *CreateStackInstancesResponseBody {
	s.OperationId = &v
	return s
}

func (s *CreateStackInstancesResponseBody) SetRequestId(v string) *CreateStackInstancesResponseBody {
	s.RequestId = &v
	return s
}

type CreateStackInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateStackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStackInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStackInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateStackInstancesResponse) SetHeaders(v map[string]*string) *CreateStackInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateStackInstancesResponse) SetStatusCode(v int32) *CreateStackInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStackInstancesResponse) SetBody(v *CreateStackInstancesResponseBody) *CreateStackInstancesResponse {
	s.Body = v
	return s
}

type CreateTemplateRequest struct {
	// The description of the template. The description can be up to 256 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the resource group.\
	// For more information about resource groups, see [Resource groups](~~94475~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the template.
	Tags []*CreateTemplateRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The structure of the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.
	//
	// > You must specify TemplateBody or TemplateURL.
	//
	// You can create a Terraform template based on your business requirements. The following sample code provides an example on how to create a Terraform template:
	//
	//     {
	//       "ROSTemplateFormatVersion": "2015-09-01",
	//       "Transform": "Aliyun::Terraform-v1.0",
	//       "Workspace": {
	//         "main.tf": "variable  \"name\" {  default = \"auto_provisioning_group\"}"
	//       },
	//       "Outputs": {}
	//     }
	//
	// For more information about Terraform templates, see [Structure of Terraform templates](~~184397~~).
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The name of the template.\
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or a letter.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Alibaba Cloud Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo or oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The template body must be 1 to 1,024 bytes in length. If you do not specify the region of the OSS bucket, the value of RegionId is used.
	//
	// > You must specify TemplateBody or TemplateURL.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) SetDescription(v string) *CreateTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateRequest) SetResourceGroupId(v string) *CreateTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateRequest) SetTags(v []*CreateTemplateRequestTags) *CreateTemplateRequest {
	s.Tags = v
	return s
}

func (s *CreateTemplateRequest) SetTemplateBody(v string) *CreateTemplateRequest {
	s.TemplateBody = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateURL(v string) *CreateTemplateRequest {
	s.TemplateURL = &v
	return s
}

type CreateTemplateRequestTags struct {
	// The tag key of the template.
	//
	// > Tags is optional. If you need to specify Tags, you must also specify Key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the template.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTemplateRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateRequestTags) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequestTags) SetKey(v string) *CreateTemplateRequestTags {
	s.Key = &v
	return s
}

func (s *CreateTemplateRequestTags) SetValue(v string) *CreateTemplateRequestTags {
	s.Value = &v
	return s
}

type CreateTemplateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the template.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponseBody) SetRequestId(v string) *CreateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateResponseBody) SetTemplateId(v string) *CreateTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateResponse) SetHeaders(v map[string]*string) *CreateTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateResponse) SetStatusCode(v int32) *CreateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplateResponse) SetBody(v *CreateTemplateResponseBody) *CreateTemplateResponse {
	s.Body = v
	return s
}

type CreateTemplateScratchRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the scenario.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution mode. Valid values:
	//
	// *   Async (default)
	// *   Sync
	//
	// > If you have a wide scope of resources, Sync takes longer. If you set ExecutionMode to Sync, we recommend that you specify ClientToken to prevent the execution timeout.
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The policy based on which the logical ID is generated. Valid values:
	//
	// *   LongTypePrefixAndIndexSuffix (default): long-type prefix + index-type suffix
	// *   LongTypePrefixAndHashSuffix: long-type prefix + hash-type suffix
	// *   ShortTypePrefixAndHashSuffix: short-type prefix + hash-type suffix
	LogicalIdStrategy *string `json:"LogicalIdStrategy,omitempty" xml:"LogicalIdStrategy,omitempty"`
	// The preference parameters of the scenario.
	PreferenceParameters []*CreateTemplateScratchRequestPreferenceParameters `json:"PreferenceParameters,omitempty" xml:"PreferenceParameters,omitempty" type:"Repeated"`
	// The region ID of the scenario.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The source resource group.
	SourceResourceGroup *CreateTemplateScratchRequestSourceResourceGroup `json:"SourceResourceGroup,omitempty" xml:"SourceResourceGroup,omitempty" type:"Struct"`
	// The source resources.
	SourceResources []*CreateTemplateScratchRequestSourceResources `json:"SourceResources,omitempty" xml:"SourceResources,omitempty" type:"Repeated"`
	// The source tag.
	SourceTag *CreateTemplateScratchRequestSourceTag `json:"SourceTag,omitempty" xml:"SourceTag,omitempty" type:"Struct"`
	// The tags of the scenario.
	Tags []*CreateTemplateScratchRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the scenario. Valid values:
	//
	// *   ResourceImport: resource management
	// *   ArchitectureReplication: resource replication
	// *   ResourceMigration: resource migration
	TemplateScratchType *string `json:"TemplateScratchType,omitempty" xml:"TemplateScratchType,omitempty"`
}

func (s CreateTemplateScratchRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateScratchRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchRequest) SetClientToken(v string) *CreateTemplateScratchRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTemplateScratchRequest) SetDescription(v string) *CreateTemplateScratchRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateScratchRequest) SetExecutionMode(v string) *CreateTemplateScratchRequest {
	s.ExecutionMode = &v
	return s
}

func (s *CreateTemplateScratchRequest) SetLogicalIdStrategy(v string) *CreateTemplateScratchRequest {
	s.LogicalIdStrategy = &v
	return s
}

func (s *CreateTemplateScratchRequest) SetPreferenceParameters(v []*CreateTemplateScratchRequestPreferenceParameters) *CreateTemplateScratchRequest {
	s.PreferenceParameters = v
	return s
}

func (s *CreateTemplateScratchRequest) SetRegionId(v string) *CreateTemplateScratchRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTemplateScratchRequest) SetResourceGroupId(v string) *CreateTemplateScratchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateScratchRequest) SetSourceResourceGroup(v *CreateTemplateScratchRequestSourceResourceGroup) *CreateTemplateScratchRequest {
	s.SourceResourceGroup = v
	return s
}

func (s *CreateTemplateScratchRequest) SetSourceResources(v []*CreateTemplateScratchRequestSourceResources) *CreateTemplateScratchRequest {
	s.SourceResources = v
	return s
}

func (s *CreateTemplateScratchRequest) SetSourceTag(v *CreateTemplateScratchRequestSourceTag) *CreateTemplateScratchRequest {
	s.SourceTag = v
	return s
}

func (s *CreateTemplateScratchRequest) SetTags(v []*CreateTemplateScratchRequestTags) *CreateTemplateScratchRequest {
	s.Tags = v
	return s
}

func (s *CreateTemplateScratchRequest) SetTemplateScratchType(v string) *CreateTemplateScratchRequest {
	s.TemplateScratchType = &v
	return s
}

type CreateTemplateScratchRequestPreferenceParameters struct {
	// The key of the parameter.
	//
	// For information about the valid values of ParameterKey, see the **Additional information about request parameters** section of this topic.
	// > - PreferenceParameters is optional. If you want to specify PreferenceParameters, you must specify ParameterKey and ParameterValue.
	// > -  If you set TemplateScratchType to ResourceImport, you must set ParameterKey to DeletionPolicy.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter. The value of ParameterValue varies based on the value of ParameterKey.
	//
	// For information about the valid values of ParameterValue, see the **Additional information about request parameters** section of this topic.
	//
	// > PreferenceParameters is optional. If you want to specify PreferenceParameters, you must specify ParameterKey and ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s CreateTemplateScratchRequestPreferenceParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateScratchRequestPreferenceParameters) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchRequestPreferenceParameters) SetParameterKey(v string) *CreateTemplateScratchRequestPreferenceParameters {
	s.ParameterKey = &v
	return s
}

func (s *CreateTemplateScratchRequestPreferenceParameters) SetParameterValue(v string) *CreateTemplateScratchRequestPreferenceParameters {
	s.ParameterValue = &v
	return s
}

type CreateTemplateScratchRequestSourceResourceGroup struct {
	// The ID of the source resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource types.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s CreateTemplateScratchRequestSourceResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateScratchRequestSourceResourceGroup) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchRequestSourceResourceGroup) SetResourceGroupId(v string) *CreateTemplateScratchRequestSourceResourceGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateScratchRequestSourceResourceGroup) SetResourceTypeFilter(v []*string) *CreateTemplateScratchRequestSourceResourceGroup {
	s.ResourceTypeFilter = v
	return s
}

type CreateTemplateScratchRequestSourceResources struct {
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateTemplateScratchRequestSourceResources) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateScratchRequestSourceResources) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchRequestSourceResources) SetRegionId(v string) *CreateTemplateScratchRequestSourceResources {
	s.RegionId = &v
	return s
}

func (s *CreateTemplateScratchRequestSourceResources) SetResourceId(v string) *CreateTemplateScratchRequestSourceResources {
	s.ResourceId = &v
	return s
}

func (s *CreateTemplateScratchRequestSourceResources) SetResourceType(v string) *CreateTemplateScratchRequestSourceResources {
	s.ResourceType = &v
	return s
}

type CreateTemplateScratchRequestSourceTag struct {
	// The source tags that consist of key-value pairs. If you want to specify only the tag key, you must leave the tag value empty. Example: `{"TagKey": ""}`.
	//
	// You can add up to 10 source tags.
	ResourceTags map[string]interface{} `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty"`
	// The resource types.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s CreateTemplateScratchRequestSourceTag) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateScratchRequestSourceTag) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchRequestSourceTag) SetResourceTags(v map[string]interface{}) *CreateTemplateScratchRequestSourceTag {
	s.ResourceTags = v
	return s
}

func (s *CreateTemplateScratchRequestSourceTag) SetResourceTypeFilter(v []*string) *CreateTemplateScratchRequestSourceTag {
	s.ResourceTypeFilter = v
	return s
}

type CreateTemplateScratchRequestTags struct {
	// The tag key of the scenario.
	//
	// > Tags is optional. If you want to specify Tags, you must specify Key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the scenario.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTemplateScratchRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateScratchRequestTags) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchRequestTags) SetKey(v string) *CreateTemplateScratchRequestTags {
	s.Key = &v
	return s
}

func (s *CreateTemplateScratchRequestTags) SetValue(v string) *CreateTemplateScratchRequestTags {
	s.Value = &v
	return s
}

type CreateTemplateScratchShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the scenario.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution mode. Valid values:
	//
	// *   Async (default)
	// *   Sync
	//
	// > If you have a wide scope of resources, Sync takes longer. If you set ExecutionMode to Sync, we recommend that you specify ClientToken to prevent the execution timeout.
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The policy based on which the logical ID is generated. Valid values:
	//
	// *   LongTypePrefixAndIndexSuffix (default): long-type prefix + index-type suffix
	// *   LongTypePrefixAndHashSuffix: long-type prefix + hash-type suffix
	// *   ShortTypePrefixAndHashSuffix: short-type prefix + hash-type suffix
	LogicalIdStrategy *string `json:"LogicalIdStrategy,omitempty" xml:"LogicalIdStrategy,omitempty"`
	// The preference parameters of the scenario.
	PreferenceParametersShrink *string `json:"PreferenceParameters,omitempty" xml:"PreferenceParameters,omitempty"`
	// The region ID of the scenario.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The source resource group.
	SourceResourceGroupShrink *string `json:"SourceResourceGroup,omitempty" xml:"SourceResourceGroup,omitempty"`
	// The source resources.
	SourceResourcesShrink *string `json:"SourceResources,omitempty" xml:"SourceResources,omitempty"`
	// The source tag.
	SourceTagShrink *string `json:"SourceTag,omitempty" xml:"SourceTag,omitempty"`
	// The tags of the scenario.
	Tags []*CreateTemplateScratchShrinkRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the scenario. Valid values:
	//
	// *   ResourceImport: resource management
	// *   ArchitectureReplication: resource replication
	// *   ResourceMigration: resource migration
	TemplateScratchType *string `json:"TemplateScratchType,omitempty" xml:"TemplateScratchType,omitempty"`
}

func (s CreateTemplateScratchShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateScratchShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchShrinkRequest) SetClientToken(v string) *CreateTemplateScratchShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetDescription(v string) *CreateTemplateScratchShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetExecutionMode(v string) *CreateTemplateScratchShrinkRequest {
	s.ExecutionMode = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetLogicalIdStrategy(v string) *CreateTemplateScratchShrinkRequest {
	s.LogicalIdStrategy = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetPreferenceParametersShrink(v string) *CreateTemplateScratchShrinkRequest {
	s.PreferenceParametersShrink = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetRegionId(v string) *CreateTemplateScratchShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetResourceGroupId(v string) *CreateTemplateScratchShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetSourceResourceGroupShrink(v string) *CreateTemplateScratchShrinkRequest {
	s.SourceResourceGroupShrink = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetSourceResourcesShrink(v string) *CreateTemplateScratchShrinkRequest {
	s.SourceResourcesShrink = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetSourceTagShrink(v string) *CreateTemplateScratchShrinkRequest {
	s.SourceTagShrink = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetTags(v []*CreateTemplateScratchShrinkRequestTags) *CreateTemplateScratchShrinkRequest {
	s.Tags = v
	return s
}

func (s *CreateTemplateScratchShrinkRequest) SetTemplateScratchType(v string) *CreateTemplateScratchShrinkRequest {
	s.TemplateScratchType = &v
	return s
}

type CreateTemplateScratchShrinkRequestTags struct {
	// The tag key of the scenario.
	//
	// > Tags is optional. If you want to specify Tags, you must specify Key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the scenario.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTemplateScratchShrinkRequestTags) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateScratchShrinkRequestTags) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchShrinkRequestTags) SetKey(v string) *CreateTemplateScratchShrinkRequestTags {
	s.Key = &v
	return s
}

func (s *CreateTemplateScratchShrinkRequestTags) SetValue(v string) *CreateTemplateScratchShrinkRequestTags {
	s.Value = &v
	return s
}

type CreateTemplateScratchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scenario.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
}

func (s CreateTemplateScratchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateScratchResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchResponseBody) SetRequestId(v string) *CreateTemplateScratchResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTemplateScratchResponseBody) SetTemplateScratchId(v string) *CreateTemplateScratchResponseBody {
	s.TemplateScratchId = &v
	return s
}

type CreateTemplateScratchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTemplateScratchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTemplateScratchResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTemplateScratchResponse) GoString() string {
	return s.String()
}

func (s *CreateTemplateScratchResponse) SetHeaders(v map[string]*string) *CreateTemplateScratchResponse {
	s.Headers = v
	return s
}

func (s *CreateTemplateScratchResponse) SetStatusCode(v int32) *CreateTemplateScratchResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTemplateScratchResponse) SetBody(v *CreateTemplateScratchResponseBody) *CreateTemplateScratchResponse {
	s.Body = v
	return s
}

type DeleteChangeSetRequest struct {
	// The ID of the change set.
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The region ID of the change set. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteChangeSetRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteChangeSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteChangeSetRequest) SetChangeSetId(v string) *DeleteChangeSetRequest {
	s.ChangeSetId = &v
	return s
}

func (s *DeleteChangeSetRequest) SetRegionId(v string) *DeleteChangeSetRequest {
	s.RegionId = &v
	return s
}

type DeleteChangeSetResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChangeSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteChangeSetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChangeSetResponseBody) SetRequestId(v string) *DeleteChangeSetResponseBody {
	s.RequestId = &v
	return s
}

type DeleteChangeSetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteChangeSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteChangeSetResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteChangeSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteChangeSetResponse) SetHeaders(v map[string]*string) *DeleteChangeSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteChangeSetResponse) SetStatusCode(v int32) *DeleteChangeSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChangeSetResponse) SetBody(v *DeleteChangeSetResponseBody) *DeleteChangeSetResponse {
	s.Body = v
	return s
}

type DeleteDiagnosticRequest struct {
	// The report ID. You can troubleshoot issues based on the report.
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s DeleteDiagnosticRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiagnosticRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiagnosticRequest) SetReportId(v string) *DeleteDiagnosticRequest {
	s.ReportId = &v
	return s
}

type DeleteDiagnosticResponseBody struct {
	// The error code returned.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   true
	// *   false
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDiagnosticResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiagnosticResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDiagnosticResponseBody) SetCode(v string) *DeleteDiagnosticResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDiagnosticResponseBody) SetHttpStatusCode(v int32) *DeleteDiagnosticResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDiagnosticResponseBody) SetMessage(v string) *DeleteDiagnosticResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDiagnosticResponseBody) SetRequestId(v string) *DeleteDiagnosticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDiagnosticResponseBody) SetSuccess(v string) *DeleteDiagnosticResponseBody {
	s.Success = &v
	return s
}

type DeleteDiagnosticResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDiagnosticResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDiagnosticResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDiagnosticResponse) GoString() string {
	return s.String()
}

func (s *DeleteDiagnosticResponse) SetHeaders(v map[string]*string) *DeleteDiagnosticResponse {
	s.Headers = v
	return s
}

func (s *DeleteDiagnosticResponse) SetStatusCode(v int32) *DeleteDiagnosticResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDiagnosticResponse) SetBody(v *DeleteDiagnosticResponseBody) *DeleteDiagnosticResponse {
	s.Body = v
	return s
}

type DeleteStackRequest struct {
	// The options for deleting the stack.
	DeleteOptions []*string `json:"DeleteOptions,omitempty" xml:"DeleteOptions,omitempty" type:"Repeated"`
	// The name of the RAM role. Resource Orchestration Service (ROS) assumes the RAM role to create the stack and uses the credentials of the role to call the APIs of Alibaba Cloud services.\
	// ROS assumes the role to perform operations on the stack. If you have permissions to perform operations on the stack but do not have permissions to use the RAM role, ROS still assumes the RAM role. You must make sure that the least privileges are granted to the RAM role.\
	// If you leave this parameter empty when you call the DeleteStack operation, ROS cannot assume the existing RAM role that is associated with the stack. If you want ROS to assume a RAM role, you must specify this parameter. If no RAM roles are available, ROS uses a temporary credential that is generated from the credentials of your account.\
	// The name of the RAM role can be up to 64 bytes in length.
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to retain all resources in the stack.
	//
	// Valid values:
	//
	// *   true
	// *   false (default)
	RetainAllResources *bool `json:"RetainAllResources,omitempty" xml:"RetainAllResources,omitempty"`
	// The resources that you want to retain.
	RetainResources []*string `json:"RetainResources,omitempty" xml:"RetainResources,omitempty" type:"Repeated"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s DeleteStackRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackRequest) GoString() string {
	return s.String()
}

func (s *DeleteStackRequest) SetDeleteOptions(v []*string) *DeleteStackRequest {
	s.DeleteOptions = v
	return s
}

func (s *DeleteStackRequest) SetRamRoleName(v string) *DeleteStackRequest {
	s.RamRoleName = &v
	return s
}

func (s *DeleteStackRequest) SetRegionId(v string) *DeleteStackRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStackRequest) SetRetainAllResources(v bool) *DeleteStackRequest {
	s.RetainAllResources = &v
	return s
}

func (s *DeleteStackRequest) SetRetainResources(v []*string) *DeleteStackRequest {
	s.RetainResources = v
	return s
}

func (s *DeleteStackRequest) SetStackId(v string) *DeleteStackRequest {
	s.StackId = &v
	return s
}

type DeleteStackResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStackResponseBody) SetRequestId(v string) *DeleteStackResponseBody {
	s.RequestId = &v
	return s
}

type DeleteStackResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStackResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackResponse) GoString() string {
	return s.String()
}

func (s *DeleteStackResponse) SetHeaders(v map[string]*string) *DeleteStackResponse {
	s.Headers = v
	return s
}

func (s *DeleteStackResponse) SetStatusCode(v int32) *DeleteStackResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStackResponse) SetBody(v *DeleteStackResponseBody) *DeleteStackResponse {
	s.Body = v
	return s
}

type DeleteStackGroupRequest struct {
	// The ID of the region to which the stack group belongs. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group. The name must be unique in a region.
	//
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). The name must start with a digit or a letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s DeleteStackGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteStackGroupRequest) SetRegionId(v string) *DeleteStackGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStackGroupRequest) SetStackGroupName(v string) *DeleteStackGroupRequest {
	s.StackGroupName = &v
	return s
}

type DeleteStackGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStackGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStackGroupResponseBody) SetRequestId(v string) *DeleteStackGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteStackGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteStackGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStackGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteStackGroupResponse) SetHeaders(v map[string]*string) *DeleteStackGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteStackGroupResponse) SetStatusCode(v int32) *DeleteStackGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStackGroupResponse) SetBody(v *DeleteStackGroupResponseBody) *DeleteStackGroupResponse {
	s.Body = v
	return s
}

type DeleteStackInstancesRequest struct {
	// The IDs of the execution accounts within which you want to deploy stacks in self-managed mode. You can specify up to 20 execution account IDs.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\
	// The token can contain letters, digits, hyphens (-), and underscores (\_), and cannot exceed 64 characters in length.\
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The folders in which you want to deploy stacks in service-managed mode.
	DeploymentTargets *DeleteStackInstancesRequestDeploymentTargets `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty" type:"Struct"`
	// The description of the delete operation.
	//
	// The description must be 1 to 256 characters in length.
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The preference settings of the delete operation.
	//
	// The following parameters are available:
	//
	// -  {"FailureToleranceCount": N}
	//
	//     The number of accounts within which stack operation failures are allowed in each region. If the value of this parameter is exceeded in a region, ROS stops the operation in the region. If ROS stops the operation in one region, ROS stops the operation in other regions.
	//
	//     Valid values of N: 0 to 20.
	//
	//     If you do not specify FailureToleranceCount, 0 is used as the default value.
	//
	// -  {"FailureTolerancePercentage": N}
	//
	//     The percentage of the number of accounts within which stack operation failures are allowed to the total number of accounts in each region. If the value of this parameter is exceeded, ROS stops the operation in the region.
	//
	//     Valid values of N: 0 to 100. If the numeric value in the percentage is not an integer, ROS rounds the value down to the nearest integer.
	//
	//     If you do not specify FailureTolerancePercentage, 0 is used as the default value.
	//
	// -  {"MaxConcurrentCount": N}
	//
	//     The maximum number of accounts within which multiple stacks are deployed at the same time in each region.
	//
	//     Valid values of N: 1 to 20.
	//
	//     If you do not specify MaxConcurrentCount, 1 is used as the default value.
	//
	// -  {"MaxConcurrentPercentage": N}
	//
	//     The percentage of the maximum number of accounts within which stacks are deployed at the same time to the total number of accounts in each region.
	//
	//     Valid values of N: 1 to 100. If the numeric value in the percentage is not an integer, ROS rounds the number down to the nearest integer.
	//
	//     If you do not specify MaxConcurrentPercentage, 1 is used as the default value.
	//
	// -   {"RegionConcurrencyType": N}
	//
	//     The mode that you want to use to deploy stacks across regions. Valid values:
	//     - SEQUENTIAL (default): deploys stacks in the specified regions one by one in sequence. This way, ROS deploys stacks in only one region at a time.
	//
	//      - PARALLEL: deploys stacks in all the specified regions in parallel.
	//
	// Separate multiple parameters with commas (,).
	//
	// > - You can specify only one of the following parameters: MaxConcurrentCount and MaxConcurrentPercentage.
	// > - You can specify only one of the following parameters: FailureToleranceCount and FailureTolerancePercentage.
	OperationPreferences map[string]interface{} `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the regions where you want to delete the stacks. You can specify up to 20 region IDs.
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// Specifies whether to delete the stacks.
	//
	// Valid values:
	//
	// *   true: retains the stacks.
	// *   false: deletes the stacks.
	RetainStacks *bool `json:"RetainStacks,omitempty" xml:"RetainStacks,omitempty"`
	// The name of the stack group. The name must be unique within a region.\
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or a letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s DeleteStackInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteStackInstancesRequest) SetAccountIds(v []*string) *DeleteStackInstancesRequest {
	s.AccountIds = v
	return s
}

func (s *DeleteStackInstancesRequest) SetClientToken(v string) *DeleteStackInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteStackInstancesRequest) SetDeploymentTargets(v *DeleteStackInstancesRequestDeploymentTargets) *DeleteStackInstancesRequest {
	s.DeploymentTargets = v
	return s
}

func (s *DeleteStackInstancesRequest) SetOperationDescription(v string) *DeleteStackInstancesRequest {
	s.OperationDescription = &v
	return s
}

func (s *DeleteStackInstancesRequest) SetOperationPreferences(v map[string]interface{}) *DeleteStackInstancesRequest {
	s.OperationPreferences = v
	return s
}

func (s *DeleteStackInstancesRequest) SetRegionId(v string) *DeleteStackInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStackInstancesRequest) SetRegionIds(v []*string) *DeleteStackInstancesRequest {
	s.RegionIds = v
	return s
}

func (s *DeleteStackInstancesRequest) SetRetainStacks(v bool) *DeleteStackInstancesRequest {
	s.RetainStacks = &v
	return s
}

func (s *DeleteStackInstancesRequest) SetStackGroupName(v string) *DeleteStackInstancesRequest {
	s.StackGroupName = &v
	return s
}

type DeleteStackInstancesRequestDeploymentTargets struct {
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The IDs of the folders in the resource directory. You can add up to five folder IDs.
	//
	// You can create stacks within all the member accounts in the specified folders. If you create stacks in the Root folder, the stacks are created within all member accounts in the resource directory.
	//
	// > To view the folder IDs, go to the **Overview** page in the **Resource Management** console. For more information, see [View the basic information about a folder](~~111223~~).
	RdFolderIds []*string `json:"RdFolderIds,omitempty" xml:"RdFolderIds,omitempty" type:"Repeated"`
}

func (s DeleteStackInstancesRequestDeploymentTargets) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackInstancesRequestDeploymentTargets) GoString() string {
	return s.String()
}

func (s *DeleteStackInstancesRequestDeploymentTargets) SetAccountIds(v []*string) *DeleteStackInstancesRequestDeploymentTargets {
	s.AccountIds = v
	return s
}

func (s *DeleteStackInstancesRequestDeploymentTargets) SetRdFolderIds(v []*string) *DeleteStackInstancesRequestDeploymentTargets {
	s.RdFolderIds = v
	return s
}

type DeleteStackInstancesShrinkRequest struct {
	// The IDs of the execution accounts within which you want to deploy stacks in self-managed mode. You can specify up to 20 execution account IDs.
	AccountIdsShrink *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\
	// The token can contain letters, digits, hyphens (-), and underscores (\_), and cannot exceed 64 characters in length.\
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The folders in which you want to deploy stacks in service-managed mode.
	DeploymentTargetsShrink *string `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty"`
	// The description of the delete operation.
	//
	// The description must be 1 to 256 characters in length.
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The preference settings of the delete operation.
	//
	// The following parameters are available:
	//
	// -  {"FailureToleranceCount": N}
	//
	//     The number of accounts within which stack operation failures are allowed in each region. If the value of this parameter is exceeded in a region, ROS stops the operation in the region. If ROS stops the operation in one region, ROS stops the operation in other regions.
	//
	//     Valid values of N: 0 to 20.
	//
	//     If you do not specify FailureToleranceCount, 0 is used as the default value.
	//
	// -  {"FailureTolerancePercentage": N}
	//
	//     The percentage of the number of accounts within which stack operation failures are allowed to the total number of accounts in each region. If the value of this parameter is exceeded, ROS stops the operation in the region.
	//
	//     Valid values of N: 0 to 100. If the numeric value in the percentage is not an integer, ROS rounds the value down to the nearest integer.
	//
	//     If you do not specify FailureTolerancePercentage, 0 is used as the default value.
	//
	// -  {"MaxConcurrentCount": N}
	//
	//     The maximum number of accounts within which multiple stacks are deployed at the same time in each region.
	//
	//     Valid values of N: 1 to 20.
	//
	//     If you do not specify MaxConcurrentCount, 1 is used as the default value.
	//
	// -  {"MaxConcurrentPercentage": N}
	//
	//     The percentage of the maximum number of accounts within which stacks are deployed at the same time to the total number of accounts in each region.
	//
	//     Valid values of N: 1 to 100. If the numeric value in the percentage is not an integer, ROS rounds the number down to the nearest integer.
	//
	//     If you do not specify MaxConcurrentPercentage, 1 is used as the default value.
	//
	// -   {"RegionConcurrencyType": N}
	//
	//     The mode that you want to use to deploy stacks across regions. Valid values:
	//     - SEQUENTIAL (default): deploys stacks in the specified regions one by one in sequence. This way, ROS deploys stacks in only one region at a time.
	//
	//      - PARALLEL: deploys stacks in all the specified regions in parallel.
	//
	// Separate multiple parameters with commas (,).
	//
	// > - You can specify only one of the following parameters: MaxConcurrentCount and MaxConcurrentPercentage.
	// > - You can specify only one of the following parameters: FailureToleranceCount and FailureTolerancePercentage.
	OperationPreferencesShrink *string `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the regions where you want to delete the stacks. You can specify up to 20 region IDs.
	RegionIdsShrink *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	// Specifies whether to delete the stacks.
	//
	// Valid values:
	//
	// *   true: retains the stacks.
	// *   false: deletes the stacks.
	RetainStacks *bool `json:"RetainStacks,omitempty" xml:"RetainStacks,omitempty"`
	// The name of the stack group. The name must be unique within a region.\
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or a letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s DeleteStackInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteStackInstancesShrinkRequest) SetAccountIdsShrink(v string) *DeleteStackInstancesShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetClientToken(v string) *DeleteStackInstancesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetDeploymentTargetsShrink(v string) *DeleteStackInstancesShrinkRequest {
	s.DeploymentTargetsShrink = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetOperationDescription(v string) *DeleteStackInstancesShrinkRequest {
	s.OperationDescription = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetOperationPreferencesShrink(v string) *DeleteStackInstancesShrinkRequest {
	s.OperationPreferencesShrink = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetRegionId(v string) *DeleteStackInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetRegionIdsShrink(v string) *DeleteStackInstancesShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetRetainStacks(v bool) *DeleteStackInstancesShrinkRequest {
	s.RetainStacks = &v
	return s
}

func (s *DeleteStackInstancesShrinkRequest) SetStackGroupName(v string) *DeleteStackInstancesShrinkRequest {
	s.StackGroupName = &v
	return s
}

type DeleteStackInstancesResponseBody struct {
	// The ID of the operation.
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStackInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStackInstancesResponseBody) SetOperationId(v string) *DeleteStackInstancesResponseBody {
	s.OperationId = &v
	return s
}

func (s *DeleteStackInstancesResponseBody) SetRequestId(v string) *DeleteStackInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteStackInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteStackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStackInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStackInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteStackInstancesResponse) SetHeaders(v map[string]*string) *DeleteStackInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteStackInstancesResponse) SetStatusCode(v int32) *DeleteStackInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStackInstancesResponse) SetBody(v *DeleteStackInstancesResponseBody) *DeleteStackInstancesResponse {
	s.Body = v
	return s
}

type DeleteTemplateRequest struct {
	// The ID of the template. This parameter applies to only private templates.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) SetTemplateId(v string) *DeleteTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteTemplateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponse) SetHeaders(v map[string]*string) *DeleteTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateResponse) SetStatusCode(v int32) *DeleteTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateResponse) SetBody(v *DeleteTemplateResponseBody) *DeleteTemplateResponse {
	s.Body = v
	return s
}

type DeleteTemplateScratchRequest struct {
	// The region ID of the scenario.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scenario.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
}

func (s DeleteTemplateScratchRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateScratchRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateScratchRequest) SetRegionId(v string) *DeleteTemplateScratchRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTemplateScratchRequest) SetTemplateScratchId(v string) *DeleteTemplateScratchRequest {
	s.TemplateScratchId = &v
	return s
}

type DeleteTemplateScratchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateScratchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateScratchResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateScratchResponseBody) SetRequestId(v string) *DeleteTemplateScratchResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTemplateScratchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTemplateScratchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTemplateScratchResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTemplateScratchResponse) GoString() string {
	return s.String()
}

func (s *DeleteTemplateScratchResponse) SetHeaders(v map[string]*string) *DeleteTemplateScratchResponse {
	s.Headers = v
	return s
}

func (s *DeleteTemplateScratchResponse) SetStatusCode(v int32) *DeleteTemplateScratchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTemplateScratchResponse) SetBody(v *DeleteTemplateScratchResponseBody) *DeleteTemplateScratchResponse {
	s.Body = v
	return s
}

type DeregisterResourceTypeRequest struct {
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The version ID. If you want to delete a version of the resource type, you must specify this parameter.
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DeregisterResourceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeregisterResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *DeregisterResourceTypeRequest) SetResourceType(v string) *DeregisterResourceTypeRequest {
	s.ResourceType = &v
	return s
}

func (s *DeregisterResourceTypeRequest) SetVersionId(v string) *DeregisterResourceTypeRequest {
	s.VersionId = &v
	return s
}

type DeregisterResourceTypeResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeregisterResourceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeregisterResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *DeregisterResourceTypeResponseBody) SetRequestId(v string) *DeregisterResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

type DeregisterResourceTypeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeregisterResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeregisterResourceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeregisterResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *DeregisterResourceTypeResponse) SetHeaders(v map[string]*string) *DeregisterResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *DeregisterResourceTypeResponse) SetStatusCode(v int32) *DeregisterResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeregisterResourceTypeResponse) SetBody(v *DeregisterResourceTypeResponseBody) *DeregisterResourceTypeResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// The language in which you want to display the results. Valid values:
	//
	// *   zh-CN (default): Chinese
	// *   en-US: English
	// *   ja: Japanese
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// The regions.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// The name of the region.
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The endpoint of the region.
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// The region ID.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionEndpoint(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DetectStackDriftRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// The value can be up to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The logical ID of resource.
	LogicalResourceId []*string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty" type:"Repeated"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s DetectStackDriftRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectStackDriftRequest) GoString() string {
	return s.String()
}

func (s *DetectStackDriftRequest) SetClientToken(v string) *DetectStackDriftRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectStackDriftRequest) SetLogicalResourceId(v []*string) *DetectStackDriftRequest {
	s.LogicalResourceId = v
	return s
}

func (s *DetectStackDriftRequest) SetRegionId(v string) *DetectStackDriftRequest {
	s.RegionId = &v
	return s
}

func (s *DetectStackDriftRequest) SetStackId(v string) *DetectStackDriftRequest {
	s.StackId = &v
	return s
}

type DetectStackDriftResponseBody struct {
	// The ID of the drift detection.
	DriftDetectionId *string `json:"DriftDetectionId,omitempty" xml:"DriftDetectionId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectStackDriftResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectStackDriftResponseBody) GoString() string {
	return s.String()
}

func (s *DetectStackDriftResponseBody) SetDriftDetectionId(v string) *DetectStackDriftResponseBody {
	s.DriftDetectionId = &v
	return s
}

func (s *DetectStackDriftResponseBody) SetRequestId(v string) *DetectStackDriftResponseBody {
	s.RequestId = &v
	return s
}

type DetectStackDriftResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectStackDriftResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectStackDriftResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectStackDriftResponse) GoString() string {
	return s.String()
}

func (s *DetectStackDriftResponse) SetHeaders(v map[string]*string) *DetectStackDriftResponse {
	s.Headers = v
	return s
}

func (s *DetectStackDriftResponse) SetStatusCode(v int32) *DetectStackDriftResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectStackDriftResponse) SetBody(v *DetectStackDriftResponseBody) *DetectStackDriftResponse {
	s.Body = v
	return s
}

type DetectStackGroupDriftRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// The value can be up to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The operation settings, in JSON format. The following fields are supported:
	//
	// *   FailureToleranceCount
	//
	// The maximum number of stack group operation failures that can occur. In a stack group operation, if the total number of failures does not exceed the FailureToleranceCount value, the operation succeeds. Otherwise, the operation fails.
	//
	// If FailureToleranceCount is not specified, the default value 0 is used. You can specify one of FailureToleranceCount or FailureTolerancePercentage parameters, but you cannot specify both of them.
	//
	// Valid values: 0 to 20.
	//
	// *   FailureTolerancePercentage
	//
	// The percentage of stack group operation failures that can occur. In a stack group operation, if the percentage of failures does not exceed the FailureTolerancePercentage value, the operation succeeds. Otherwise, the operation fails.
	//
	// You can specify one of FailureToleranceCount or FailureTolerancePercentage parameters, but you cannot specify both of them.
	//
	// Valid values: 0 to 100.
	//
	// *   MaxConcurrentCount
	//
	// The maximum number of target accounts in which a drift detection operation can be performed at a time.
	//
	// You can specify one of MaxConcurrentCount or MaxConcurrentPercentage parameters, but you cannot specify both of them.
	//
	// Valid values: 1 to 20.
	//
	// *   MaxConcurrentPercentage
	//
	// The maximum percentage of target accounts in which a drift detection operation can be performed at a time.
	//
	// You can specify one of MaxConcurrentCount or MaxConcurrentPercentage parameters, but you cannot specify both of them.
	//
	// Valid values: 1 to 100.
	OperationPreferences map[string]interface{} `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035#doc-api-ROS-DescribeRegions~~ "Queries the DescribeRegions list of a region.") operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group. The name must be unique in a region.
	//
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s DetectStackGroupDriftRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectStackGroupDriftRequest) GoString() string {
	return s.String()
}

func (s *DetectStackGroupDriftRequest) SetClientToken(v string) *DetectStackGroupDriftRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectStackGroupDriftRequest) SetOperationPreferences(v map[string]interface{}) *DetectStackGroupDriftRequest {
	s.OperationPreferences = v
	return s
}

func (s *DetectStackGroupDriftRequest) SetRegionId(v string) *DetectStackGroupDriftRequest {
	s.RegionId = &v
	return s
}

func (s *DetectStackGroupDriftRequest) SetStackGroupName(v string) *DetectStackGroupDriftRequest {
	s.StackGroupName = &v
	return s
}

type DetectStackGroupDriftShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// The value can be up to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The operation settings, in JSON format. The following fields are supported:
	//
	// *   FailureToleranceCount
	//
	// The maximum number of stack group operation failures that can occur. In a stack group operation, if the total number of failures does not exceed the FailureToleranceCount value, the operation succeeds. Otherwise, the operation fails.
	//
	// If FailureToleranceCount is not specified, the default value 0 is used. You can specify one of FailureToleranceCount or FailureTolerancePercentage parameters, but you cannot specify both of them.
	//
	// Valid values: 0 to 20.
	//
	// *   FailureTolerancePercentage
	//
	// The percentage of stack group operation failures that can occur. In a stack group operation, if the percentage of failures does not exceed the FailureTolerancePercentage value, the operation succeeds. Otherwise, the operation fails.
	//
	// You can specify one of FailureToleranceCount or FailureTolerancePercentage parameters, but you cannot specify both of them.
	//
	// Valid values: 0 to 100.
	//
	// *   MaxConcurrentCount
	//
	// The maximum number of target accounts in which a drift detection operation can be performed at a time.
	//
	// You can specify one of MaxConcurrentCount or MaxConcurrentPercentage parameters, but you cannot specify both of them.
	//
	// Valid values: 1 to 20.
	//
	// *   MaxConcurrentPercentage
	//
	// The maximum percentage of target accounts in which a drift detection operation can be performed at a time.
	//
	// You can specify one of MaxConcurrentCount or MaxConcurrentPercentage parameters, but you cannot specify both of them.
	//
	// Valid values: 1 to 100.
	OperationPreferencesShrink *string `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035#doc-api-ROS-DescribeRegions~~ "Queries the DescribeRegions list of a region.") operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group. The name must be unique in a region.
	//
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s DetectStackGroupDriftShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectStackGroupDriftShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetectStackGroupDriftShrinkRequest) SetClientToken(v string) *DetectStackGroupDriftShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectStackGroupDriftShrinkRequest) SetOperationPreferencesShrink(v string) *DetectStackGroupDriftShrinkRequest {
	s.OperationPreferencesShrink = &v
	return s
}

func (s *DetectStackGroupDriftShrinkRequest) SetRegionId(v string) *DetectStackGroupDriftShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DetectStackGroupDriftShrinkRequest) SetStackGroupName(v string) *DetectStackGroupDriftShrinkRequest {
	s.StackGroupName = &v
	return s
}

type DetectStackGroupDriftResponseBody struct {
	// The ID of the operation.
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectStackGroupDriftResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectStackGroupDriftResponseBody) GoString() string {
	return s.String()
}

func (s *DetectStackGroupDriftResponseBody) SetOperationId(v string) *DetectStackGroupDriftResponseBody {
	s.OperationId = &v
	return s
}

func (s *DetectStackGroupDriftResponseBody) SetRequestId(v string) *DetectStackGroupDriftResponseBody {
	s.RequestId = &v
	return s
}

type DetectStackGroupDriftResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectStackGroupDriftResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectStackGroupDriftResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectStackGroupDriftResponse) GoString() string {
	return s.String()
}

func (s *DetectStackGroupDriftResponse) SetHeaders(v map[string]*string) *DetectStackGroupDriftResponse {
	s.Headers = v
	return s
}

func (s *DetectStackGroupDriftResponse) SetStatusCode(v int32) *DetectStackGroupDriftResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectStackGroupDriftResponse) SetBody(v *DetectStackGroupDriftResponseBody) *DetectStackGroupDriftResponse {
	s.Body = v
	return s
}

type DetectStackResourceDriftRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// The value can be up to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of the resource.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s DetectStackResourceDriftRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectStackResourceDriftRequest) GoString() string {
	return s.String()
}

func (s *DetectStackResourceDriftRequest) SetClientToken(v string) *DetectStackResourceDriftRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectStackResourceDriftRequest) SetLogicalResourceId(v string) *DetectStackResourceDriftRequest {
	s.LogicalResourceId = &v
	return s
}

func (s *DetectStackResourceDriftRequest) SetRegionId(v string) *DetectStackResourceDriftRequest {
	s.RegionId = &v
	return s
}

func (s *DetectStackResourceDriftRequest) SetStackId(v string) *DetectStackResourceDriftRequest {
	s.StackId = &v
	return s
}

type DetectStackResourceDriftResponseBody struct {
	// The actual JSON-formatted resource properties.
	ActualProperties *string `json:"ActualProperties,omitempty" xml:"ActualProperties,omitempty"`
	// The time when the resource drift detection was initiated.
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The JSON-formatted resource properties that are defined in the template.
	ExpectedProperties *string `json:"ExpectedProperties,omitempty" xml:"ExpectedProperties,omitempty"`
	// The logical ID of the resource that is defined in the template.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The physical ID of the resource.
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The property drifts of the resource.
	PropertyDifferences []*DetectStackResourceDriftResponseBodyPropertyDifferences `json:"PropertyDifferences,omitempty" xml:"PropertyDifferences,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The drift status of the resource. Valid values:
	//
	// *   DELETED: The actual configuration of the resource differs from its expected template configuration because the resource is deleted.
	// *   MODIFIED: The actual configuration of the resource differs from its expected template configuration.
	// *   NOT_CHECKED: Resource Orchestration Service (ROS) has not checked whether the actual configuration of the resource differs from its expected template configuration.
	// *   IN_SYNC: The actual configuration of the resource matches its expected template configuration.
	ResourceDriftStatus *string `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s DetectStackResourceDriftResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectStackResourceDriftResponseBody) GoString() string {
	return s.String()
}

func (s *DetectStackResourceDriftResponseBody) SetActualProperties(v string) *DetectStackResourceDriftResponseBody {
	s.ActualProperties = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetDriftDetectionTime(v string) *DetectStackResourceDriftResponseBody {
	s.DriftDetectionTime = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetExpectedProperties(v string) *DetectStackResourceDriftResponseBody {
	s.ExpectedProperties = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetLogicalResourceId(v string) *DetectStackResourceDriftResponseBody {
	s.LogicalResourceId = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetPhysicalResourceId(v string) *DetectStackResourceDriftResponseBody {
	s.PhysicalResourceId = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetPropertyDifferences(v []*DetectStackResourceDriftResponseBodyPropertyDifferences) *DetectStackResourceDriftResponseBody {
	s.PropertyDifferences = v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetRequestId(v string) *DetectStackResourceDriftResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetResourceDriftStatus(v string) *DetectStackResourceDriftResponseBody {
	s.ResourceDriftStatus = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetResourceType(v string) *DetectStackResourceDriftResponseBody {
	s.ResourceType = &v
	return s
}

func (s *DetectStackResourceDriftResponseBody) SetStackId(v string) *DetectStackResourceDriftResponseBody {
	s.StackId = &v
	return s
}

type DetectStackResourceDriftResponseBodyPropertyDifferences struct {
	// The actual value of the resource property.
	ActualValue *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	// The drift type of the resource property. Valid values:
	//
	// *   ADD: The property value has been added to a resource property whose data type was Array or List.
	// *   REMOVE: The property has been deleted from the current resource configuration.
	// *   NOT_EQUAL: The current property value differs from the expected value defined in the stack template.
	DifferenceType *string `json:"DifferenceType,omitempty" xml:"DifferenceType,omitempty"`
	// The expected value of the resource property that is defined in the template.
	ExpectedValue *string `json:"ExpectedValue,omitempty" xml:"ExpectedValue,omitempty"`
	// The path of the resource property.
	PropertyPath *string `json:"PropertyPath,omitempty" xml:"PropertyPath,omitempty"`
}

func (s DetectStackResourceDriftResponseBodyPropertyDifferences) String() string {
	return tea.Prettify(s)
}

func (s DetectStackResourceDriftResponseBodyPropertyDifferences) GoString() string {
	return s.String()
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) SetActualValue(v string) *DetectStackResourceDriftResponseBodyPropertyDifferences {
	s.ActualValue = &v
	return s
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) SetDifferenceType(v string) *DetectStackResourceDriftResponseBodyPropertyDifferences {
	s.DifferenceType = &v
	return s
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) SetExpectedValue(v string) *DetectStackResourceDriftResponseBodyPropertyDifferences {
	s.ExpectedValue = &v
	return s
}

func (s *DetectStackResourceDriftResponseBodyPropertyDifferences) SetPropertyPath(v string) *DetectStackResourceDriftResponseBodyPropertyDifferences {
	s.PropertyPath = &v
	return s
}

type DetectStackResourceDriftResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetectStackResourceDriftResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DetectStackResourceDriftResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectStackResourceDriftResponse) GoString() string {
	return s.String()
}

func (s *DetectStackResourceDriftResponse) SetHeaders(v map[string]*string) *DetectStackResourceDriftResponse {
	s.Headers = v
	return s
}

func (s *DetectStackResourceDriftResponse) SetStatusCode(v int32) *DetectStackResourceDriftResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectStackResourceDriftResponse) SetBody(v *DetectStackResourceDriftResponseBody) *DetectStackResourceDriftResponse {
	s.Body = v
	return s
}

type ExecuteChangeSetRequest struct {
	// The ID of the change set.
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests.
	//
	// The token can be up to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the change set. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ExecuteChangeSetRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteChangeSetRequest) GoString() string {
	return s.String()
}

func (s *ExecuteChangeSetRequest) SetChangeSetId(v string) *ExecuteChangeSetRequest {
	s.ChangeSetId = &v
	return s
}

func (s *ExecuteChangeSetRequest) SetClientToken(v string) *ExecuteChangeSetRequest {
	s.ClientToken = &v
	return s
}

func (s *ExecuteChangeSetRequest) SetRegionId(v string) *ExecuteChangeSetRequest {
	s.RegionId = &v
	return s
}

type ExecuteChangeSetResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteChangeSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteChangeSetResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteChangeSetResponseBody) SetRequestId(v string) *ExecuteChangeSetResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteChangeSetResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExecuteChangeSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ExecuteChangeSetResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteChangeSetResponse) GoString() string {
	return s.String()
}

func (s *ExecuteChangeSetResponse) SetHeaders(v map[string]*string) *ExecuteChangeSetResponse {
	s.Headers = v
	return s
}

func (s *ExecuteChangeSetResponse) SetStatusCode(v int32) *ExecuteChangeSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteChangeSetResponse) SetBody(v *ExecuteChangeSetResponseBody) *ExecuteChangeSetResponse {
	s.Body = v
	return s
}

type GenerateTemplateByScratchRequest struct {
	// The region ID of the new node.
	ProvisionRegionId *string `json:"ProvisionRegionId,omitempty" xml:"ProvisionRegionId,omitempty"`
	// The region ID of the scenario.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scenario.
	//
	// For more information about how to query the IDs of scenarios, see [ListTemplateScratches](~~363050~~).
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The type of the template that Resource Orchestration Service (ROS) generates. ROS can generate templates of the ROS and Terraform types. Default value: ROS.
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s GenerateTemplateByScratchRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplateByScratchRequest) GoString() string {
	return s.String()
}

func (s *GenerateTemplateByScratchRequest) SetProvisionRegionId(v string) *GenerateTemplateByScratchRequest {
	s.ProvisionRegionId = &v
	return s
}

func (s *GenerateTemplateByScratchRequest) SetRegionId(v string) *GenerateTemplateByScratchRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateTemplateByScratchRequest) SetTemplateScratchId(v string) *GenerateTemplateByScratchRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *GenerateTemplateByScratchRequest) SetTemplateType(v string) *GenerateTemplateByScratchRequest {
	s.TemplateType = &v
	return s
}

type GenerateTemplateByScratchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources that you want to import into a stack in the scenario.
	//
	// > This parameter is returned only if the scenario is of the Resource Management type.
	ResourcesToImport []*GenerateTemplateByScratchResponseBodyResourcesToImport `json:"ResourcesToImport,omitempty" xml:"ResourcesToImport,omitempty" type:"Repeated"`
	// The template content of the scenario.
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
}

func (s GenerateTemplateByScratchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplateByScratchResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTemplateByScratchResponseBody) SetRequestId(v string) *GenerateTemplateByScratchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateTemplateByScratchResponseBody) SetResourcesToImport(v []*GenerateTemplateByScratchResponseBodyResourcesToImport) *GenerateTemplateByScratchResponseBody {
	s.ResourcesToImport = v
	return s
}

func (s *GenerateTemplateByScratchResponseBody) SetTemplateBody(v string) *GenerateTemplateByScratchResponseBody {
	s.TemplateBody = &v
	return s
}

type GenerateTemplateByScratchResponseBodyResourcesToImport struct {
	// The logical ID of the resource.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The key-value mapping between strings. The value is a JSON string that identifies the resource that you want to import into a stack.\
	// A key is an identifier for a resource, and a value is an assignment of data to the key. For example, VpcId is a key that indicates the ID of a virtual private cloud (VPC), and `vpc-bp1m6fww66xbntjyc****"` is a value that is assigned to VpcId.
	ResourceIdentifier map[string]interface{} `json:"ResourceIdentifier,omitempty" xml:"ResourceIdentifier,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GenerateTemplateByScratchResponseBodyResourcesToImport) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplateByScratchResponseBodyResourcesToImport) GoString() string {
	return s.String()
}

func (s *GenerateTemplateByScratchResponseBodyResourcesToImport) SetLogicalResourceId(v string) *GenerateTemplateByScratchResponseBodyResourcesToImport {
	s.LogicalResourceId = &v
	return s
}

func (s *GenerateTemplateByScratchResponseBodyResourcesToImport) SetResourceIdentifier(v map[string]interface{}) *GenerateTemplateByScratchResponseBodyResourcesToImport {
	s.ResourceIdentifier = v
	return s
}

func (s *GenerateTemplateByScratchResponseBodyResourcesToImport) SetResourceType(v string) *GenerateTemplateByScratchResponseBodyResourcesToImport {
	s.ResourceType = &v
	return s
}

type GenerateTemplateByScratchResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateTemplateByScratchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateTemplateByScratchResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplateByScratchResponse) GoString() string {
	return s.String()
}

func (s *GenerateTemplateByScratchResponse) SetHeaders(v map[string]*string) *GenerateTemplateByScratchResponse {
	s.Headers = v
	return s
}

func (s *GenerateTemplateByScratchResponse) SetStatusCode(v int32) *GenerateTemplateByScratchResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTemplateByScratchResponse) SetBody(v *GenerateTemplateByScratchResponseBody) *GenerateTemplateByScratchResponse {
	s.Body = v
	return s
}

type GenerateTemplatePolicyRequest struct {
	// The type of operation N for which you want to generate the policy information.
	//
	// Valid values:
	//
	// *   CreateStack: creates a stack by calling the CreateStack operation.
	// *   UpdateStack: updates a stack by calling the UpdateStack operation.
	// *   DeleteStack: deletes a stack by calling the DeleteStack operation.
	// *   DetectStackDrift: detects drifts on a stack by calling the DelectStackDrift operation.
	// *   ListStackOperationRisks: lists the risks of a deletion operation on a stack by setting the OperationType parameter to DeleteStack in the ListStackOperationRisks operation.
	// *   GetTemplateEstimateCost: queries the estimated prices of resources that you want to use in the template by calling the GetTemplateEstimateCost operation.
	// *   GetTemplateParameterConstraints: queries the values of parameters in the template by calling the GetTemplateParameterConstraints operation.
	// *   ImportResourcesToStack: imports resources to a stack by setting the ChangeSetType parameter to IMPORT in the CreateChangeSet operation.
	// *   SignalResource: sends a signal to a stack.
	//
	// >  The default value is the combination of all valid values.
	OperationTypes []*string `json:"OperationTypes,omitempty" xml:"OperationTypes,omitempty" type:"Repeated"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.
	//
	// If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared templates and private templates.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length.
	//
	// >  If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	//
	// The URL can be up to 1,024 bytes in length.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. This parameter takes effect only when the TemplateId parameter is specified.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GenerateTemplatePolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplatePolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyRequest) SetOperationTypes(v []*string) *GenerateTemplatePolicyRequest {
	s.OperationTypes = v
	return s
}

func (s *GenerateTemplatePolicyRequest) SetTemplateBody(v string) *GenerateTemplatePolicyRequest {
	s.TemplateBody = &v
	return s
}

func (s *GenerateTemplatePolicyRequest) SetTemplateId(v string) *GenerateTemplatePolicyRequest {
	s.TemplateId = &v
	return s
}

func (s *GenerateTemplatePolicyRequest) SetTemplateURL(v string) *GenerateTemplatePolicyRequest {
	s.TemplateURL = &v
	return s
}

func (s *GenerateTemplatePolicyRequest) SetTemplateVersion(v string) *GenerateTemplatePolicyRequest {
	s.TemplateVersion = &v
	return s
}

type GenerateTemplatePolicyResponseBody struct {
	// The information about the policy.
	Policy *GenerateTemplatePolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateTemplatePolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplatePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponseBody) SetPolicy(v *GenerateTemplatePolicyResponseBodyPolicy) *GenerateTemplatePolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GenerateTemplatePolicyResponseBody) SetRequestId(v string) *GenerateTemplatePolicyResponseBody {
	s.RequestId = &v
	return s
}

type GenerateTemplatePolicyResponseBodyPolicy struct {
	// The statements that are contained in the policy.
	Statement []*GenerateTemplatePolicyResponseBodyPolicyStatement `json:"Statement,omitempty" xml:"Statement,omitempty" type:"Repeated"`
	// The version number.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GenerateTemplatePolicyResponseBodyPolicy) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplatePolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponseBodyPolicy) SetStatement(v []*GenerateTemplatePolicyResponseBodyPolicyStatement) *GenerateTemplatePolicyResponseBodyPolicy {
	s.Statement = v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicy) SetVersion(v string) *GenerateTemplatePolicyResponseBodyPolicy {
	s.Version = &v
	return s
}

type GenerateTemplatePolicyResponseBodyPolicyStatement struct {
	// The operations that are performed on the specified resource.
	Action []*string `json:"Action,omitempty" xml:"Action,omitempty" type:"Repeated"`
	// The condition that is required for the policy to take effect.
	Condition map[string]interface{} `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The effect of the statement. Valid values:
	//
	// *   Allow
	// *   Deny
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The objects that the statement covers. An asterisk (\*) indicates all resources.
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
}

func (s GenerateTemplatePolicyResponseBodyPolicyStatement) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplatePolicyResponseBodyPolicyStatement) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) SetAction(v []*string) *GenerateTemplatePolicyResponseBodyPolicyStatement {
	s.Action = v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) SetCondition(v map[string]interface{}) *GenerateTemplatePolicyResponseBodyPolicyStatement {
	s.Condition = v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) SetEffect(v string) *GenerateTemplatePolicyResponseBodyPolicyStatement {
	s.Effect = &v
	return s
}

func (s *GenerateTemplatePolicyResponseBodyPolicyStatement) SetResource(v string) *GenerateTemplatePolicyResponseBodyPolicyStatement {
	s.Resource = &v
	return s
}

type GenerateTemplatePolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateTemplatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateTemplatePolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateTemplatePolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponse) SetHeaders(v map[string]*string) *GenerateTemplatePolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateTemplatePolicyResponse) SetStatusCode(v int32) *GenerateTemplatePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTemplatePolicyResponse) SetBody(v *GenerateTemplatePolicyResponseBody) *GenerateTemplatePolicyResponse {
	s.Body = v
	return s
}

type GetChangeSetRequest struct {
	// The ID of the change set.
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The region ID of the change set. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to obtain the template. Valid values:
	//
	// *   true
	// *   false (default)
	ShowTemplate *bool `json:"ShowTemplate,omitempty" xml:"ShowTemplate,omitempty"`
}

func (s GetChangeSetRequest) String() string {
	return tea.Prettify(s)
}

func (s GetChangeSetRequest) GoString() string {
	return s.String()
}

func (s *GetChangeSetRequest) SetChangeSetId(v string) *GetChangeSetRequest {
	s.ChangeSetId = &v
	return s
}

func (s *GetChangeSetRequest) SetRegionId(v string) *GetChangeSetRequest {
	s.RegionId = &v
	return s
}

func (s *GetChangeSetRequest) SetShowTemplate(v bool) *GetChangeSetRequest {
	s.ShowTemplate = &v
	return s
}

type GetChangeSetResponseBody struct {
	// The ID of the change set.
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The name of the change set.
	ChangeSetName *string `json:"ChangeSetName,omitempty" xml:"ChangeSetName,omitempty"`
	// The type of the change set.
	ChangeSetType *string `json:"ChangeSetType,omitempty" xml:"ChangeSetType,omitempty"`
	// The changes of the change set.
	Changes []map[string]interface{} `json:"Changes,omitempty" xml:"Changes,omitempty" type:"Repeated"`
	// The time when the change set was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the change set.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether rollback was performed when the stack failed to be created or updated.
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The execution status of the change set.
	ExecutionStatus *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	// The output logs of the change set.
	Log *GetChangeSetResponseBodyLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Struct"`
	// The parameters of the stack.
	Parameters []*GetChangeSetResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region ID of the change set.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the stack with which the change set is associated.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The name of the stack with which the change set is associated.
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The status of the change set.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the change set is in its current state.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The template body of the change set.
	//
	// > This parameter takes effect only if you set ShowTemplate to true.
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The timeout period that is specified for the stack creation or update operation.
	TimeoutInMinutes *int32 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s GetChangeSetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetChangeSetResponseBody) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponseBody) SetChangeSetId(v string) *GetChangeSetResponseBody {
	s.ChangeSetId = &v
	return s
}

func (s *GetChangeSetResponseBody) SetChangeSetName(v string) *GetChangeSetResponseBody {
	s.ChangeSetName = &v
	return s
}

func (s *GetChangeSetResponseBody) SetChangeSetType(v string) *GetChangeSetResponseBody {
	s.ChangeSetType = &v
	return s
}

func (s *GetChangeSetResponseBody) SetChanges(v []map[string]interface{}) *GetChangeSetResponseBody {
	s.Changes = v
	return s
}

func (s *GetChangeSetResponseBody) SetCreateTime(v string) *GetChangeSetResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetChangeSetResponseBody) SetDescription(v string) *GetChangeSetResponseBody {
	s.Description = &v
	return s
}

func (s *GetChangeSetResponseBody) SetDisableRollback(v bool) *GetChangeSetResponseBody {
	s.DisableRollback = &v
	return s
}

func (s *GetChangeSetResponseBody) SetExecutionStatus(v string) *GetChangeSetResponseBody {
	s.ExecutionStatus = &v
	return s
}

func (s *GetChangeSetResponseBody) SetLog(v *GetChangeSetResponseBodyLog) *GetChangeSetResponseBody {
	s.Log = v
	return s
}

func (s *GetChangeSetResponseBody) SetParameters(v []*GetChangeSetResponseBodyParameters) *GetChangeSetResponseBody {
	s.Parameters = v
	return s
}

func (s *GetChangeSetResponseBody) SetRegionId(v string) *GetChangeSetResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetChangeSetResponseBody) SetRequestId(v string) *GetChangeSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChangeSetResponseBody) SetStackId(v string) *GetChangeSetResponseBody {
	s.StackId = &v
	return s
}

func (s *GetChangeSetResponseBody) SetStackName(v string) *GetChangeSetResponseBody {
	s.StackName = &v
	return s
}

func (s *GetChangeSetResponseBody) SetStatus(v string) *GetChangeSetResponseBody {
	s.Status = &v
	return s
}

func (s *GetChangeSetResponseBody) SetStatusReason(v string) *GetChangeSetResponseBody {
	s.StatusReason = &v
	return s
}

func (s *GetChangeSetResponseBody) SetTemplateBody(v string) *GetChangeSetResponseBody {
	s.TemplateBody = &v
	return s
}

func (s *GetChangeSetResponseBody) SetTimeoutInMinutes(v int32) *GetChangeSetResponseBody {
	s.TimeoutInMinutes = &v
	return s
}

type GetChangeSetResponseBodyLog struct {
	// The Terraform logs. This parameter is returned only for change sets of Terraform stacks.
	//
	// > This parameter is not returned for change sets that are in the Creating state. This parameter indicates the logs of the change set creation operation for Terraform stacks.
	TerraformLogs []*GetChangeSetResponseBodyLogTerraformLogs `json:"TerraformLogs,omitempty" xml:"TerraformLogs,omitempty" type:"Repeated"`
}

func (s GetChangeSetResponseBodyLog) String() string {
	return tea.Prettify(s)
}

func (s GetChangeSetResponseBodyLog) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponseBodyLog) SetTerraformLogs(v []*GetChangeSetResponseBodyLogTerraformLogs) *GetChangeSetResponseBodyLog {
	s.TerraformLogs = v
	return s
}

type GetChangeSetResponseBodyLogTerraformLogs struct {
	// The name of the Terraform command that is run. Valid values:
	//
	// *   apply
	// *   plan
	// *   destroy
	// *   version
	//
	// For more information about Terraform commands, see [Command](https://www.terraform.io/cli/commands).
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The content of the output stream that is returned after the command is run.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The output stream. Valid values:
	//
	// *   stdout: standard output stream
	// *   stderr: standard error stream
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s GetChangeSetResponseBodyLogTerraformLogs) String() string {
	return tea.Prettify(s)
}

func (s GetChangeSetResponseBodyLogTerraformLogs) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponseBodyLogTerraformLogs) SetCommand(v string) *GetChangeSetResponseBodyLogTerraformLogs {
	s.Command = &v
	return s
}

func (s *GetChangeSetResponseBodyLogTerraformLogs) SetContent(v string) *GetChangeSetResponseBodyLogTerraformLogs {
	s.Content = &v
	return s
}

func (s *GetChangeSetResponseBodyLogTerraformLogs) SetStream(v string) *GetChangeSetResponseBodyLogTerraformLogs {
	s.Stream = &v
	return s
}

type GetChangeSetResponseBodyParameters struct {
	// The key of the parameter.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetChangeSetResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetChangeSetResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponseBodyParameters) SetParameterKey(v string) *GetChangeSetResponseBodyParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetChangeSetResponseBodyParameters) SetParameterValue(v string) *GetChangeSetResponseBodyParameters {
	s.ParameterValue = &v
	return s
}

type GetChangeSetResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetChangeSetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetChangeSetResponse) String() string {
	return tea.Prettify(s)
}

func (s GetChangeSetResponse) GoString() string {
	return s.String()
}

func (s *GetChangeSetResponse) SetHeaders(v map[string]*string) *GetChangeSetResponse {
	s.Headers = v
	return s
}

func (s *GetChangeSetResponse) SetStatusCode(v int32) *GetChangeSetResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChangeSetResponse) SetBody(v *GetChangeSetResponseBody) *GetChangeSetResponse {
	s.Body = v
	return s
}

type GetDiagnosticRequest struct {
	// The ID of the diagnostic report.
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s GetDiagnosticRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosticRequest) GoString() string {
	return s.String()
}

func (s *GetDiagnosticRequest) SetReportId(v string) *GetDiagnosticRequest {
	s.ReportId = &v
	return s
}

type GetDiagnosticResponseBody struct {
	// The error code returned.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The keyword in the diagnosis.
	DiagnosticKey *string `json:"DiagnosticKey,omitempty" xml:"DiagnosticKey,omitempty"`
	// The name of the diagnostic item.
	DiagnosticProduct *string `json:"DiagnosticProduct,omitempty" xml:"DiagnosticProduct,omitempty"`
	// The diagnosis result.
	DiagnosticResult *GetDiagnosticResponseBodyDiagnosticResult `json:"DiagnosticResult,omitempty" xml:"DiagnosticResult,omitempty" type:"Struct"`
	// The time when the diagnosis was performed.
	DiagnosticTime *string `json:"DiagnosticTime,omitempty" xml:"DiagnosticTime,omitempty"`
	// The HTTP status code
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The HTTP status code returned.
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The suggestion for the diagnosis.
	Recommends map[string]interface{} `json:"Recommends,omitempty" xml:"Recommends,omitempty"`
	// The ID of the diagnostic report.
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The diagnosis status. Valid values:
	//
	// *   Running: The diagnosis is in progress.
	// *   Complete: The diagnosis is complete.
	// *   Failed: The diagnosis failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason for the diagnosis status.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   true
	// *   false
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDiagnosticResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosticResponseBody) GoString() string {
	return s.String()
}

func (s *GetDiagnosticResponseBody) SetCode(v string) *GetDiagnosticResponseBody {
	s.Code = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetDiagnosticKey(v string) *GetDiagnosticResponseBody {
	s.DiagnosticKey = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetDiagnosticProduct(v string) *GetDiagnosticResponseBody {
	s.DiagnosticProduct = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetDiagnosticResult(v *GetDiagnosticResponseBodyDiagnosticResult) *GetDiagnosticResponseBody {
	s.DiagnosticResult = v
	return s
}

func (s *GetDiagnosticResponseBody) SetDiagnosticTime(v string) *GetDiagnosticResponseBody {
	s.DiagnosticTime = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetHttpCode(v string) *GetDiagnosticResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetHttpStatusCode(v int32) *GetDiagnosticResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetMessage(v string) *GetDiagnosticResponseBody {
	s.Message = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetRecommends(v map[string]interface{}) *GetDiagnosticResponseBody {
	s.Recommends = v
	return s
}

func (s *GetDiagnosticResponseBody) SetReportId(v string) *GetDiagnosticResponseBody {
	s.ReportId = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetRequestId(v string) *GetDiagnosticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetStatus(v string) *GetDiagnosticResponseBody {
	s.Status = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetStatusReason(v string) *GetDiagnosticResponseBody {
	s.StatusReason = &v
	return s
}

func (s *GetDiagnosticResponseBody) SetSuccess(v string) *GetDiagnosticResponseBody {
	s.Success = &v
	return s
}

type GetDiagnosticResponseBodyDiagnosticResult struct {
	// The resources that failed to be diagnosed.
	FailedResources map[string]interface{} `json:"FailedResources,omitempty" xml:"FailedResources,omitempty"`
	// The information about Resource Orchestration Service (ROS) calling.
	RosActionMessages map[string]interface{} `json:"RosActionMessages,omitempty" xml:"RosActionMessages,omitempty"`
	// The stack information.
	StackMessages map[string]interface{} `json:"StackMessages,omitempty" xml:"StackMessages,omitempty"`
}

func (s GetDiagnosticResponseBodyDiagnosticResult) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosticResponseBodyDiagnosticResult) GoString() string {
	return s.String()
}

func (s *GetDiagnosticResponseBodyDiagnosticResult) SetFailedResources(v map[string]interface{}) *GetDiagnosticResponseBodyDiagnosticResult {
	s.FailedResources = v
	return s
}

func (s *GetDiagnosticResponseBodyDiagnosticResult) SetRosActionMessages(v map[string]interface{}) *GetDiagnosticResponseBodyDiagnosticResult {
	s.RosActionMessages = v
	return s
}

func (s *GetDiagnosticResponseBodyDiagnosticResult) SetStackMessages(v map[string]interface{}) *GetDiagnosticResponseBodyDiagnosticResult {
	s.StackMessages = v
	return s
}

type GetDiagnosticResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDiagnosticResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDiagnosticResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDiagnosticResponse) GoString() string {
	return s.String()
}

func (s *GetDiagnosticResponse) SetHeaders(v map[string]*string) *GetDiagnosticResponse {
	s.Headers = v
	return s
}

func (s *GetDiagnosticResponse) SetStatusCode(v int32) *GetDiagnosticResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDiagnosticResponse) SetBody(v *GetDiagnosticResponseBody) *GetDiagnosticResponse {
	s.Body = v
	return s
}

type GetFeatureDetailsRequest struct {
	// The one or more features that you want to query. Valid values:
	//
	// *   Terraform: the Terraform hosting feature.
	// *   ResourceCleaner: the resource cleaner feature. You can use ALIYUN::ROS::ResourceCleaner to create a resource cleaner.
	// *   TemplateScratch: the scenario feature.
	// *   All: all features that are supported by ROS.
	Feature *string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetFeatureDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsRequest) SetFeature(v string) *GetFeatureDetailsRequest {
	s.Feature = &v
	return s
}

func (s *GetFeatureDetailsRequest) SetRegionId(v string) *GetFeatureDetailsRequest {
	s.RegionId = &v
	return s
}

type GetFeatureDetailsResponseBody struct {
	// Details of the drift detection feature.
	DriftDetection *GetFeatureDetailsResponseBodyDriftDetection `json:"DriftDetection,omitempty" xml:"DriftDetection,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the resource cleaner feature.
	ResourceCleaner *GetFeatureDetailsResponseBodyResourceCleaner `json:"ResourceCleaner,omitempty" xml:"ResourceCleaner,omitempty" type:"Struct"`
	// Details of the resource import feature.
	ResourceImport *GetFeatureDetailsResponseBodyResourceImport `json:"ResourceImport,omitempty" xml:"ResourceImport,omitempty" type:"Struct"`
	// Details of the template parameter constraint feature.
	TemplateParameterConstraints *GetFeatureDetailsResponseBodyTemplateParameterConstraints `json:"TemplateParameterConstraints,omitempty" xml:"TemplateParameterConstraints,omitempty" type:"Struct"`
	// Details of the scenario feature.
	TemplateScratch *GetFeatureDetailsResponseBodyTemplateScratch `json:"TemplateScratch,omitempty" xml:"TemplateScratch,omitempty" type:"Struct"`
	// Details of the Terraform hosting feature.
	Terraform *GetFeatureDetailsResponseBodyTerraform `json:"Terraform,omitempty" xml:"Terraform,omitempty" type:"Struct"`
}

func (s GetFeatureDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBody) SetDriftDetection(v *GetFeatureDetailsResponseBodyDriftDetection) *GetFeatureDetailsResponseBody {
	s.DriftDetection = v
	return s
}

func (s *GetFeatureDetailsResponseBody) SetRequestId(v string) *GetFeatureDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFeatureDetailsResponseBody) SetResourceCleaner(v *GetFeatureDetailsResponseBodyResourceCleaner) *GetFeatureDetailsResponseBody {
	s.ResourceCleaner = v
	return s
}

func (s *GetFeatureDetailsResponseBody) SetResourceImport(v *GetFeatureDetailsResponseBodyResourceImport) *GetFeatureDetailsResponseBody {
	s.ResourceImport = v
	return s
}

func (s *GetFeatureDetailsResponseBody) SetTemplateParameterConstraints(v *GetFeatureDetailsResponseBodyTemplateParameterConstraints) *GetFeatureDetailsResponseBody {
	s.TemplateParameterConstraints = v
	return s
}

func (s *GetFeatureDetailsResponseBody) SetTemplateScratch(v *GetFeatureDetailsResponseBodyTemplateScratch) *GetFeatureDetailsResponseBody {
	s.TemplateScratch = v
	return s
}

func (s *GetFeatureDetailsResponseBody) SetTerraform(v *GetFeatureDetailsResponseBodyTerraform) *GetFeatureDetailsResponseBody {
	s.Terraform = v
	return s
}

type GetFeatureDetailsResponseBodyDriftDetection struct {
	// The resource types that are supported by the drift detection feature.
	SupportedResourceTypes []*string `json:"SupportedResourceTypes,omitempty" xml:"SupportedResourceTypes,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyDriftDetection) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyDriftDetection) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyDriftDetection) SetSupportedResourceTypes(v []*string) *GetFeatureDetailsResponseBodyDriftDetection {
	s.SupportedResourceTypes = v
	return s
}

type GetFeatureDetailsResponseBodyResourceCleaner struct {
	// The resource types that can be cleaned up.
	SupportedResourceTypes []*GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes `json:"SupportedResourceTypes,omitempty" xml:"SupportedResourceTypes,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyResourceCleaner) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyResourceCleaner) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyResourceCleaner) SetSupportedResourceTypes(v []*GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) *GetFeatureDetailsResponseBodyResourceCleaner {
	s.SupportedResourceTypes = v
	return s
}

type GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes struct {
	// The resource type that supports the resource cleaner feature.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The names of the side effects that may be caused by the cleanup operation performed on the resources of the specified type.
	SideEffects []*string `json:"SideEffects,omitempty" xml:"SideEffects,omitempty" type:"Repeated"`
	// The names of the filters that are supported by the resource type.
	SupportedFilters []*string `json:"SupportedFilters,omitempty" xml:"SupportedFilters,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) SetResourceType(v string) *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) SetSideEffects(v []*string) *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes {
	s.SideEffects = v
	return s
}

func (s *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes) SetSupportedFilters(v []*string) *GetFeatureDetailsResponseBodyResourceCleanerSupportedResourceTypes {
	s.SupportedFilters = v
	return s
}

type GetFeatureDetailsResponseBodyResourceImport struct {
	// The resource types that are supported by the resource import feature.
	SupportedResourceTypes []*GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes `json:"SupportedResourceTypes,omitempty" xml:"SupportedResourceTypes,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyResourceImport) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyResourceImport) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyResourceImport) SetSupportedResourceTypes(v []*GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes) *GetFeatureDetailsResponseBodyResourceImport {
	s.SupportedResourceTypes = v
	return s
}

type GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes struct {
	// The resource identifiers.
	ResourceIdentifiers []*string `json:"ResourceIdentifiers,omitempty" xml:"ResourceIdentifiers,omitempty" type:"Repeated"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes) SetResourceIdentifiers(v []*string) *GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes {
	s.ResourceIdentifiers = v
	return s
}

func (s *GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes) SetResourceType(v string) *GetFeatureDetailsResponseBodyResourceImportSupportedResourceTypes {
	s.ResourceType = &v
	return s
}

type GetFeatureDetailsResponseBodyTemplateParameterConstraints struct {
	// The resource types that support the template parameter constraint feature.
	SupportedResourceTypes []*GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes `json:"SupportedResourceTypes,omitempty" xml:"SupportedResourceTypes,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTemplateParameterConstraints) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTemplateParameterConstraints) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTemplateParameterConstraints) SetSupportedResourceTypes(v []*GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes) *GetFeatureDetailsResponseBodyTemplateParameterConstraints {
	s.SupportedResourceTypes = v
	return s
}

type GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes struct {
	// The names of properties that are supported by the resource type.
	Properties []*string `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes) SetProperties(v []*string) *GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes {
	s.Properties = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes) SetResourceType(v string) *GetFeatureDetailsResponseBodyTemplateParameterConstraintsSupportedResourceTypes {
	s.ResourceType = &v
	return s
}

type GetFeatureDetailsResponseBodyTemplateScratch struct {
	// The resource types that are supported by the scenario feature.
	SupportedResourceTypes []*GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes `json:"SupportedResourceTypes,omitempty" xml:"SupportedResourceTypes,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTemplateScratch) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTemplateScratch) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTemplateScratch) SetSupportedResourceTypes(v []*GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) *GetFeatureDetailsResponseBodyTemplateScratch {
	s.SupportedResourceTypes = v
	return s
}

type GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes struct {
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Indicates whether the resource scope can be specified by source resource group. Valid values:
	//
	// *   true
	// *   false
	SourceResourceGroupSupported *bool `json:"SourceResourceGroupSupported,omitempty" xml:"SourceResourceGroupSupported,omitempty"`
	// Indicates whether the resource scope can be specified by source resource. Valid values:
	//
	// *   true
	// *   false
	SourceResourcesSupported *bool `json:"SourceResourcesSupported,omitempty" xml:"SourceResourcesSupported,omitempty"`
	// Indicates whether the resource scope can be specified by source tag, resource group, or resource. Valid values:
	//
	// *   true
	// *   false
	SourceSupported *bool `json:"SourceSupported,omitempty" xml:"SourceSupported,omitempty"`
	// Indicates whether the resource scope can be specified by source tag. Valid values:
	//
	// *   true
	// *   false
	SourceTagSupported *bool `json:"SourceTagSupported,omitempty" xml:"SourceTagSupported,omitempty"`
	// The scenario types that are supported.
	SupportedTemplateScratchTypes []*string `json:"SupportedTemplateScratchTypes,omitempty" xml:"SupportedTemplateScratchTypes,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) SetResourceType(v string) *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes {
	s.ResourceType = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) SetSourceResourceGroupSupported(v bool) *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes {
	s.SourceResourceGroupSupported = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) SetSourceResourcesSupported(v bool) *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes {
	s.SourceResourcesSupported = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) SetSourceSupported(v bool) *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes {
	s.SourceSupported = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) SetSourceTagSupported(v bool) *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes {
	s.SourceTagSupported = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes) SetSupportedTemplateScratchTypes(v []*string) *GetFeatureDetailsResponseBodyTemplateScratchSupportedResourceTypes {
	s.SupportedTemplateScratchTypes = v
	return s
}

type GetFeatureDetailsResponseBodyTerraform struct {
	// The resource types that support the scenario feature.
	SupportedResourceTypes *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes `json:"SupportedResourceTypes,omitempty" xml:"SupportedResourceTypes,omitempty" type:"Struct"`
	// The Terraform versions.
	SupportedVersions []*GetFeatureDetailsResponseBodyTerraformSupportedVersions `json:"SupportedVersions,omitempty" xml:"SupportedVersions,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTerraform) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTerraform) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTerraform) SetSupportedResourceTypes(v *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) *GetFeatureDetailsResponseBodyTerraform {
	s.SupportedResourceTypes = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraform) SetSupportedVersions(v []*GetFeatureDetailsResponseBodyTerraformSupportedVersions) *GetFeatureDetailsResponseBodyTerraform {
	s.SupportedVersions = v
	return s
}

type GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes struct {
	// The resource types that support the custom tag feature.
	CustomTag []*string `json:"CustomTag,omitempty" xml:"CustomTag,omitempty" type:"Repeated"`
	// The resource types that support the price inquiry feature.
	EstimateCost []*string `json:"EstimateCost,omitempty" xml:"EstimateCost,omitempty" type:"Repeated"`
	// The resource types that support the resource group feature.
	ResourceGroup []*string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Repeated"`
	// The resource type that support the risk check feature.
	StackOperationRisk *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk `json:"StackOperationRisk,omitempty" xml:"StackOperationRisk,omitempty" type:"Struct"`
	// The resource types that support the system tag `acs:ros:stackId`.
	SystemTag []*string `json:"SystemTag,omitempty" xml:"SystemTag,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) SetCustomTag(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes {
	s.CustomTag = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) SetEstimateCost(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes {
	s.EstimateCost = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) SetResourceGroup(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes {
	s.ResourceGroup = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) SetStackOperationRisk(v *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk) *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes {
	s.StackOperationRisk = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes) SetSystemTag(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypes {
	s.SystemTag = v
	return s
}

type GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk struct {
	// The resource types that support the risk check performed to detect risks caused by a stack deletion operation.
	DeleteStack []*string `json:"DeleteStack,omitempty" xml:"DeleteStack,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk) SetDeleteStack(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedResourceTypesStackOperationRisk {
	s.DeleteStack = v
	return s
}

type GetFeatureDetailsResponseBodyTerraformSupportedVersions struct {
	// The names and versions of the providers that correspond to the Terraform versions.
	ProviderVersions []*GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions `json:"ProviderVersions,omitempty" xml:"ProviderVersions,omitempty" type:"Repeated"`
	// The Terraform version.
	TerraformVersion *string `json:"TerraformVersion,omitempty" xml:"TerraformVersion,omitempty"`
	// The Terraform version that is supported by ROS. The parameter value is the same as the value of the Transform parameter in a Terraform template.
	Transform *string `json:"Transform,omitempty" xml:"Transform,omitempty"`
	// The Terraform versions that can be updated in ROS.
	UpdateAllowedTransforms []*string `json:"UpdateAllowedTransforms,omitempty" xml:"UpdateAllowedTransforms,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedVersions) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedVersions) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersions) SetProviderVersions(v []*GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions) *GetFeatureDetailsResponseBodyTerraformSupportedVersions {
	s.ProviderVersions = v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersions) SetTerraformVersion(v string) *GetFeatureDetailsResponseBodyTerraformSupportedVersions {
	s.TerraformVersion = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersions) SetTransform(v string) *GetFeatureDetailsResponseBodyTerraformSupportedVersions {
	s.Transform = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersions) SetUpdateAllowedTransforms(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedVersions {
	s.UpdateAllowedTransforms = v
	return s
}

type GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions struct {
	// The name of the provider.
	ProviderName *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	// The provider versions.
	SupportedVersions []*string `json:"SupportedVersions,omitempty" xml:"SupportedVersions,omitempty" type:"Repeated"`
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions) SetProviderName(v string) *GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions {
	s.ProviderName = &v
	return s
}

func (s *GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions) SetSupportedVersions(v []*string) *GetFeatureDetailsResponseBodyTerraformSupportedVersionsProviderVersions {
	s.SupportedVersions = v
	return s
}

type GetFeatureDetailsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFeatureDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFeatureDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFeatureDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetFeatureDetailsResponse) SetHeaders(v map[string]*string) *GetFeatureDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetFeatureDetailsResponse) SetStatusCode(v int32) *GetFeatureDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFeatureDetailsResponse) SetBody(v *GetFeatureDetailsResponseBody) *GetFeatureDetailsResponse {
	s.Body = v
	return s
}

type GetResourceTypeRequest struct {
	// The ID of the request.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The version ID. If you want to query a specific version of the resource type, you must specify this parameter. If you do not specify this parameter, only the resource type is queried.
	//
	// > This parameter is supported only for modules.
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetResourceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *GetResourceTypeRequest) SetResourceType(v string) *GetResourceTypeRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceTypeRequest) SetVersionId(v string) *GetResourceTypeRequest {
	s.VersionId = &v
	return s
}

type GetResourceTypeResponseBody struct {
	// The type of the resource.
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The creation time. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The default version ID.
	//
	// > This parameter is returned only if the resource type is queried.
	DefaultVersionId *string `json:"DefaultVersionId,omitempty" xml:"DefaultVersionId,omitempty"`
	// The description of the resource type.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The entity type. Valid values:
	//
	// *   Resource: regular resource. For more information, see [Resources](~~28863~~).
	// *   DataSource: DataSource resource. For more information, see [DataSource resources](~~404753~~).
	// *   module: module.
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// Indicates whether the version is the default version. Valid values:
	//
	// *   true
	// *   false
	//
	// > This parameter is returned only if a specific version of the resource type is queried.
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	// The latest version ID.
	//
	// > This parameter is returned only if the resource type is queried.
	LatestVersionId *string `json:"LatestVersionId,omitempty" xml:"LatestVersionId,omitempty"`
	// Indicates whether the resource supports drift detection. Default value: false. Valid values:
	//
	// *   true: Drift detection is supported.
	// *   false: Drift detection is not supported.
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The provider of the resource type. Valid values:
	//
	// *   ROS: The resource type is provided by Resource Orchestration Service (ROS).
	// *   Self: The resource type is provided by you.
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The attributes of the resource.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The properties of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Indicates whether the resource supports scratch detection. Default value: false. Valid values:
	//
	// *   true: Scratch detection is supported.
	// *   false: Scratch detection is not supported.
	SupportDriftDetection *bool `json:"SupportDriftDetection,omitempty" xml:"SupportDriftDetection,omitempty"`
	// The entity type. Valid values:
	//
	// *   Resource: resources other than DataSource resources. For more information, see [Resources](~~28863~~).
	// *   DataSource: DataSource resources.
	SupportScratchDetection *bool `json:"SupportScratchDetection,omitempty" xml:"SupportScratchDetection,omitempty"`
	// The template content in the module.
	//
	// > This parameter is returned only if a specific version of the resource type is queried.
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The total number of versions.
	//
	// > This parameter is returned only if the resource type is queried.
	TotalVersionCount *int32 `json:"TotalVersionCount,omitempty" xml:"TotalVersionCount,omitempty"`
	// The update time. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetResourceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponseBody) SetAttributes(v map[string]interface{}) *GetResourceTypeResponseBody {
	s.Attributes = v
	return s
}

func (s *GetResourceTypeResponseBody) SetCreateTime(v string) *GetResourceTypeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetDefaultVersionId(v string) *GetResourceTypeResponseBody {
	s.DefaultVersionId = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetDescription(v string) *GetResourceTypeResponseBody {
	s.Description = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetEntityType(v string) *GetResourceTypeResponseBody {
	s.EntityType = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetIsDefaultVersion(v bool) *GetResourceTypeResponseBody {
	s.IsDefaultVersion = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetLatestVersionId(v string) *GetResourceTypeResponseBody {
	s.LatestVersionId = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetProperties(v map[string]interface{}) *GetResourceTypeResponseBody {
	s.Properties = v
	return s
}

func (s *GetResourceTypeResponseBody) SetProvider(v string) *GetResourceTypeResponseBody {
	s.Provider = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetRequestId(v string) *GetResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetResourceType(v string) *GetResourceTypeResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetSupportDriftDetection(v bool) *GetResourceTypeResponseBody {
	s.SupportDriftDetection = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetSupportScratchDetection(v bool) *GetResourceTypeResponseBody {
	s.SupportScratchDetection = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetTemplateBody(v string) *GetResourceTypeResponseBody {
	s.TemplateBody = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetTotalVersionCount(v int32) *GetResourceTypeResponseBody {
	s.TotalVersionCount = &v
	return s
}

func (s *GetResourceTypeResponseBody) SetUpdateTime(v string) *GetResourceTypeResponseBody {
	s.UpdateTime = &v
	return s
}

type GetResourceTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *GetResourceTypeResponse) SetHeaders(v map[string]*string) *GetResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *GetResourceTypeResponse) SetStatusCode(v int32) *GetResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceTypeResponse) SetBody(v *GetResourceTypeResponseBody) *GetResourceTypeResponse {
	s.Body = v
	return s
}

type GetResourceTypeTemplateRequest struct {
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The version ID. If you want to query a specific version of the resource type, you must specify this parameter. If you do not specify this parameter, only the resource type is queried.
	//
	// > This parameter is supported only for modules.
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetResourceTypeTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetResourceTypeTemplateRequest) SetResourceType(v string) *GetResourceTypeTemplateRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceTypeTemplateRequest) SetVersionId(v string) *GetResourceTypeTemplateRequest {
	s.VersionId = &v
	return s
}

type GetResourceTypeTemplateResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The structure that contains the template body. The template body must be 1 to 51,200 bytes in length. For more information, see [Template syntax](~~28857~~).
	//
	// > We recommend that use TemplateContent instead of TemplateBody.
	TemplateBody map[string]interface{} `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The JSON-formatted structure of the template body. For more information, see [Template syntax](~~28857~~).
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
}

func (s GetResourceTypeTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceTypeTemplateResponseBody) SetRequestId(v string) *GetResourceTypeTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceTypeTemplateResponseBody) SetTemplateBody(v map[string]interface{}) *GetResourceTypeTemplateResponseBody {
	s.TemplateBody = v
	return s
}

func (s *GetResourceTypeTemplateResponseBody) SetTemplateContent(v string) *GetResourceTypeTemplateResponseBody {
	s.TemplateContent = &v
	return s
}

type GetResourceTypeTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetResourceTypeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetResourceTypeTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResourceTypeTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetResourceTypeTemplateResponse) SetHeaders(v map[string]*string) *GetResourceTypeTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetResourceTypeTemplateResponse) SetStatusCode(v int32) *GetResourceTypeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceTypeTemplateResponse) SetBody(v *GetResourceTypeTemplateResponseBody) *GetResourceTypeTemplateResponse {
	s.Body = v
	return s
}

type GetServiceProvisionsRequest struct {
	// The parameters.
	Parameters []*GetServiceProvisionsRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The services.
	Services []*GetServiceProvisionsRequestServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.
	//
	// You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and Services.
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The template ID. This parameter applies to shared and private templates.
	//
	// You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and Services.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body must be 1 to 524,288 bytes in length. If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and Services.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. If you do not specify this parameter, the latest version is used.
	//
	// This parameter takes effect only when TemplateId is specified.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetServiceProvisionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsRequest) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsRequest) SetParameters(v []*GetServiceProvisionsRequestParameters) *GetServiceProvisionsRequest {
	s.Parameters = v
	return s
}

func (s *GetServiceProvisionsRequest) SetRegionId(v string) *GetServiceProvisionsRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetServices(v []*GetServiceProvisionsRequestServices) *GetServiceProvisionsRequest {
	s.Services = v
	return s
}

func (s *GetServiceProvisionsRequest) SetTemplateBody(v string) *GetServiceProvisionsRequest {
	s.TemplateBody = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetTemplateId(v string) *GetServiceProvisionsRequest {
	s.TemplateId = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetTemplateURL(v string) *GetServiceProvisionsRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetServiceProvisionsRequest) SetTemplateVersion(v string) *GetServiceProvisionsRequest {
	s.TemplateVersion = &v
	return s
}

type GetServiceProvisionsRequestParameters struct {
	// The name of the parameter. If you do not specify the name and value of a parameter, Resource Orchestration Service (ROS) uses the default name and value that are specified in the template.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify ParameterKey.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetServiceProvisionsRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsRequestParameters) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsRequestParameters) SetParameterKey(v string) *GetServiceProvisionsRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetServiceProvisionsRequestParameters) SetParameterValue(v string) *GetServiceProvisionsRequestParameters {
	s.ParameterValue = &v
	return s
}

type GetServiceProvisionsRequestServices struct {
	// The service or feature name. Valid values:
	//
	// *   AHAS: Application High Availability Service
	// *   ARMS: Application Real-Time Monitoring Service (ARMS)
	// *   ApiGateway: API Gateway
	// *   BatchCompute: Batch Compute
	// *   BrainIndustrial: Industrial Brain
	// *   CloudStorageGateway: Cloud Storage Gateway (CSG)
	// *   CMS: CloudMonitor
	// *   CR: Container Registry
	// *   CS: Container Service for Kubernetes (ACK)
	// *   DCDN: Dynamic Content Delivery Network (DCDN)
	// *   DataHub: DataHub
	// *   DataWorks: DataWorks
	// *   EDAS: Enterprise Distributed Application Service (EDAS)
	// *   EHPC: Elastic High Performance Computing (E-HPC)
	// *   EMAS: Enterprise Mobile Application Studio (EMAS)
	// *   FC: Function Compute
	// *   FNF: Serverless Workflow (SWF)
	// *   MaxCompute: MaxCompute
	// *   MNS: Message Service (MNS)
	// *   HBR: Hybrid Backup Recovery (HBR)
	// *   IMM: Intelligent Media Management
	// *   IOT: IoT Platform
	// *   KMS: Key Management Service (KMS)
	// *   NAS: Apsara File Storage NAS (NAS)
	// *   NLP: Natural Language Processing (NLP)
	// *   OSS: OSS
	// *   OTS: Tablestore
	// *   PrivateLink: PrivateLink
	// *   PrivateZone: Alibaba Cloud DNS PrivateZone
	// *   RocketMQ: ApsaraMQ for RocketMQ
	// *   SAE: Serverless App Engine (SAE)
	// *   SLS: Log Service
	// *   TrafficMirror: the traffic mirroring feature
	// *   VS: Video Surveillance System
	// *   Xtrace: Managed Service for OpenTelemetry
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s GetServiceProvisionsRequestServices) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsRequestServices) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsRequestServices) SetServiceName(v string) *GetServiceProvisionsRequestServices {
	s.ServiceName = &v
	return s
}

type GetServiceProvisionsResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the services.
	ServiceProvisions []*GetServiceProvisionsResponseBodyServiceProvisions `json:"ServiceProvisions,omitempty" xml:"ServiceProvisions,omitempty" type:"Repeated"`
}

func (s GetServiceProvisionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBody) SetRequestId(v string) *GetServiceProvisionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceProvisionsResponseBody) SetServiceProvisions(v []*GetServiceProvisionsResponseBodyServiceProvisions) *GetServiceProvisionsResponseBody {
	s.ServiceProvisions = v
	return s
}

type GetServiceProvisionsResponseBodyServiceProvisions struct {
	// Indicates whether automatic activation for the service is defined in the template. Valid values:
	//
	// *   true: Automatic activation for the service is defined in the template.
	// *   false: Manual activation for the service is defined in the template.
	AutoEnableService *bool `json:"AutoEnableService,omitempty" xml:"AutoEnableService,omitempty"`
	// The names of the services on which the service that is queried depends.
	DependentServiceNames []*string `json:"DependentServiceNames,omitempty" xml:"DependentServiceNames,omitempty" type:"Repeated"`
	// The URL that points to the activation page of the service.
	//
	// > This parameter is returned if Status is set to Disabled.
	EnableURL *string `json:"EnableURL,omitempty" xml:"EnableURL,omitempty"`
	// The information about the RAM roles of the service. If this parameter is empty, no RAM role is associated with the service.
	RoleProvision *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision `json:"RoleProvision,omitempty" xml:"RoleProvision,omitempty" type:"Struct"`
	// The service name.
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The activation status of the service. Valid values:
	//
	// *   Enabled: The service is activated.
	// *   Disabled: The service is not activated.
	// *   Unknown: The activation status of the service is unknown.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the service is in the Disabled or Unknown state.
	//
	// > This parameter is returned if Status is set to Disabled or Unknown.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisions) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisions) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetAutoEnableService(v bool) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.AutoEnableService = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetDependentServiceNames(v []*string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.DependentServiceNames = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetEnableURL(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.EnableURL = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetRoleProvision(v *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.RoleProvision = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetServiceName(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.ServiceName = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetStatus(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.Status = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisions) SetStatusReason(v string) *GetServiceProvisionsResponseBodyServiceProvisions {
	s.StatusReason = &v
	return s
}

type GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision struct {
	// The authorization URL of the RAM role.
	//
	// > This parameter is returned if Created is set to false.
	AuthorizationURL *string `json:"AuthorizationURL,omitempty" xml:"AuthorizationURL,omitempty"`
	// The RAM roles of the service.
	Roles []*GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) SetAuthorizationURL(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision {
	s.AuthorizationURL = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision) SetRoles(v []*GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvision {
	s.Roles = v
	return s
}

type GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles struct {
	// The information about the API operation that is used to create the RAM role.
	ApiForCreation *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation `json:"ApiForCreation,omitempty" xml:"ApiForCreation,omitempty" type:"Struct"`
	// Indicates whether the RAM role is created. Valid values:
	//
	// *   true
	// *   false
	Created *bool `json:"Created,omitempty" xml:"Created,omitempty"`
	// The purpose for which the RAM role is used. Default value: Default. A value of Default indicates that the RAM role is the default role of the service.
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The name of the role.
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetApiForCreation(v *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.ApiForCreation = v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetCreated(v bool) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.Created = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetFunction(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.Function = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles) SetRoleName(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRoles {
	s.RoleName = &v
	return s
}

type GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation struct {
	// The name of the API operation.
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The ID of the Alibaba Cloud service to which the API operation belongs.
	ApiProductId *string `json:"ApiProductId,omitempty" xml:"ApiProductId,omitempty"`
	// The type of the API operation. Valid values:
	//
	// *   Open: public
	// *   Inner: private
	ApiType *string `json:"ApiType,omitempty" xml:"ApiType,omitempty"`
	// The parameters of the API operation. If a parameter is a variable, use the ${Variable name} format. Only the following variable is supported: ${RegionId}.
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetApiName(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.ApiName = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetApiProductId(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.ApiProductId = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetApiType(v string) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.ApiType = &v
	return s
}

func (s *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation) SetParameters(v map[string]interface{}) *GetServiceProvisionsResponseBodyServiceProvisionsRoleProvisionRolesApiForCreation {
	s.Parameters = v
	return s
}

type GetServiceProvisionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceProvisionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceProvisionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceProvisionsResponse) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsResponse) SetHeaders(v map[string]*string) *GetServiceProvisionsResponse {
	s.Headers = v
	return s
}

func (s *GetServiceProvisionsResponse) SetStatusCode(v int32) *GetServiceProvisionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceProvisionsResponse) SetBody(v *GetServiceProvisionsResponseBody) *GetServiceProvisionsResponse {
	s.Body = v
	return s
}

type GetStackRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// The token can be up to 64 characters in length,
	//
	// For more information, see [Ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The option for returning logs. Valid values:
	//
	// *   None: does not return logs.
	// *   Stack: returns the logs of the stack. This is the default value.
	// *   Resource: returns the logs of resources in the stack.
	// *   All: returns the logs of all resources.
	LogOption *string `json:"LogOption,omitempty" xml:"LogOption,omitempty"`
	// Specifies whether to return the output parameters of the stack. Valid values:
	//
	// *   Enabled: returns the output parameters. This is the default value.
	//
	// *   Disabled: does not return the output parameters.
	//
	// > The system takes a long period of time to calculate output parameters. If you do not want to query the output parameters, we recommend that you set the OutputOption parameter to Disabled to improve the response speed of the GetStack operation.
	OutputOption *string `json:"OutputOption,omitempty" xml:"OutputOption,omitempty"`
	// The ID of the region in which the stack resides. You can call the [DescribeRegions](~~131035~~) operation to query the most recent list of Alibaba Cloud regions.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to return the ResourceProgress parameter. Valid values:
	//
	// *   Disabled: does not return the ResourceProgress parameter. This is the default value.
	//
	// *   EnabledIfCreateStack: returns the ResourceProgress parameter only if a stack is created.
	//
	// > A stack is in one of the following states when it is created: CREATE_IN_PROGRESS, CREATE_COMPLETE, CREATE_FAILED, CREATE_ROLLBACK_IN_PROGRESS, CREATE_ROLLBACK_COMPLETE, or CREATE_ROLLBACK_FAILED.
	ShowResourceProgress *string `json:"ShowResourceProgress,omitempty" xml:"ShowResourceProgress,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s GetStackRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackRequest) GoString() string {
	return s.String()
}

func (s *GetStackRequest) SetClientToken(v string) *GetStackRequest {
	s.ClientToken = &v
	return s
}

func (s *GetStackRequest) SetLogOption(v string) *GetStackRequest {
	s.LogOption = &v
	return s
}

func (s *GetStackRequest) SetOutputOption(v string) *GetStackRequest {
	s.OutputOption = &v
	return s
}

func (s *GetStackRequest) SetRegionId(v string) *GetStackRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackRequest) SetShowResourceProgress(v string) *GetStackRequest {
	s.ShowResourceProgress = &v
	return s
}

func (s *GetStackRequest) SetStackId(v string) *GetStackRequest {
	s.StackId = &v
	return s
}

type GetStackResponseBody struct {
	// The number of resources on which drift detection is performed.
	//
	// >  This parameter is returned only if the drift detection on the stack is successful.
	CheckedStackResourceCount *int32 `json:"CheckedStackResourceCount,omitempty" xml:"CheckedStackResourceCount,omitempty"`
	// The time when the stack was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether deletion protection is enabled for the stack. Valid values:
	//
	// *   Enabled: Deletion protection is enabled for the stack.
	// *   Disabled: Deletion protection is disabled for the stack. You can delete the stack in the Resource Orchestration Service (ROS) console or by calling the DeleteStack operation.
	//
	// >  Deletion protection of a nested stack works in the same way as that of the root stack.
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The description of the stack.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether rollback is disabled when the stack fails to be created. Valid values:
	//
	// *   true: Rollback is disabled when the stack fails to be created.
	// *   false: Rollback is enabled when the stack fails to be created. This is the default value.
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The time when the last successful drift detection operation was performed.
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The description of the web UI in the ROS console.
	Interface *string `json:"Interface,omitempty" xml:"Interface,omitempty"`
	// The logs of the stack.
	Log *GetStackResponseBodyLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Struct"`
	// The number of resources on which drift detection is not performed.
	//
	// >  This parameter is returned only if the drift detection on the stack is successful.
	NotCheckedStackResourceCount *int32 `json:"NotCheckedStackResourceCount,omitempty" xml:"NotCheckedStackResourceCount,omitempty"`
	// The callback URLs that are used to receive stack events.
	NotificationURLs []*string `json:"NotificationURLs,omitempty" xml:"NotificationURLs,omitempty" type:"Repeated"`
	// The additional information that is displayed when an error occurs on a stack operation.
	//
	// >  This property is returned in specific conditions. At least one sub-property is returned. For example, an error is reported when you call the API of another cloud service.
	OperationInfo *GetStackResponseBodyOperationInfo `json:"OperationInfo,omitempty" xml:"OperationInfo,omitempty" type:"Struct"`
	// The ID of the order. This parameter is returned only if you set the ChargeType parameter to PrePaid.
	OrderIds []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
	// The output parameters of the stack.
	//
	// >  This parameter is returned if the OutputOption parameter is set to Enabled.
	Outputs []map[string]interface{} `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// The parameters of the stack.
	Parameters []*GetStackResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the parent stack.
	ParentStackId *string `json:"ParentStackId,omitempty" xml:"ParentStackId,omitempty"`
	// The name of the RAM role. ROS assumes the RAM role to create the stack and uses credentials of the role to call the APIs of Alibaba Cloud services.
	//
	// ROS assumes the RAM role to perform operations on the stack. If you have permissions to perform operations on the stack but do not have permissions to use the RAM role, ROS still assumes the RAM role. You must make sure that the least privileges are granted to the role.
	//
	// If you do not specify this parameter, ROS assumes an existing role that is associated with the stack. If no roles are available for ROS to assume, ROS uses a temporary credential that is generated from the credentials of your account.
	//
	// The name of the RAM role can be up to 64 bytes in length.
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the region in which the stack is deployed. You can call the [DescribeRegions](~~131035~~) operation to query the most recent list of Alibaba Cloud regions.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the instances belong.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The creation progress of resources.
	ResourceProgress *GetStackResponseBodyResourceProgress `json:"ResourceProgress,omitempty" xml:"ResourceProgress,omitempty" type:"Struct"`
	// The ID of the root stack. This parameter is returned if the specified stack is a nested stack.
	RootStackId *string `json:"RootStackId,omitempty" xml:"RootStackId,omitempty"`
	// Indicates whether the stack is a managed stack. Valid values:
	//
	// - true
	// - false
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The name of the service to which the managed stack belongs.
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The status of the stack in the last successful drift detection. Valid values:
	//
	// *   DRIFTED: The stack has drifted.
	// *   NOT_CHECKED: No successful drift detection is performed on the stack.
	// *   IN_SYNC: The stack is being synchronized.
	StackDriftStatus *string `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The name of the stack.
	//
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or letter.
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The type of the stack. Valid values:
	//
	// *   ROS: The ROS stack, which is created by using an ROS template.
	// *   Terraform: The Terraform stack, which is created by using a Terraform template.
	StackType *string `json:"StackType,omitempty" xml:"StackType,omitempty"`
	// The state of the stack. Valid values:
	//
	// *   CREATE_IN_PROGRESS: The stack is being created.
	// *   CREATE_FAILED: The stack fails to be created.
	// *   CREATE_COMPLETE: The stack is created.
	// *   UPDATE_IN_PROGRESS: The stack is being updated.
	// *   UPDATE_FAILED: The stack fails to be updated.
	// *   UPDATE_COMPLETE: The stack is updated.
	// *   DELETE_IN_PROGRESS: The stack is being deleted.
	// *   DELETE_FAILED: The stack fails to be deleted.
	// *   CREATE_ROLLBACK_IN_PROGRESS: The stack is being rolled back after the stack fails to be created.
	// *   CREATE_ROLLBACK_FAILED: The stack fails to be rolled back after the stack fails to be created.
	// *   CREATE_ROLLBACK_COMPLETE: The stack is rolled back after the stack fails to be created.
	// *   ROLLBACK_IN_PROGRESS: The resources in the stack are being rolled back.
	// *   ROLLBACK_FAILED: The resources in the stack fail to be rolled back.
	// *   ROLLBACK_COMPLETE: The resources in the stack are rolled back.
	// *   CHECK_IN_PROGRESS: The stack is being validated.
	// *   CHECK_FAILED: The stack fails to be validated.
	// *   CHECK_COMPLETE: The stack is validated.
	// *   REVIEW_IN_PROGRESS: The stack is being reviewed.
	// *   IMPORT_CREATE_IN_PROGRESS: The stack is being created by using imported resources.
	// *   IMPORT_CREATE_FAILED: The stack fails to be created by using imported resources.
	// *   IMPORT_CREATE_COMPLETE: The stack is created by using imported resources.
	// *   IMPORT_CREATE_ROLLBACK_IN_PROGRESS: The resources are being rolled back after the stack fails to be created by using imported resources.
	// *   IMPORT_CREATE_ROLLBACK_FAILED: The resources fail to be rolled back after the stack fails to be created by using imported resources.
	// *   IMPORT_CREATE_ROLLBACK_COMPLETE: The resources are rolled back after the stack fails to be created by using imported resources.
	// *   IMPORT_UPDATE_IN_PROGRESS: The stack is being updated by using imported resources.
	// *   IMPORT_UPDATE_FAILED: The stack fails to be updated by using imported resources.
	// *   IMPORT_UPDATE_COMPLETE: The stack is updated by using imported resources.
	// *   IMPORT_UPDATE_ROLLBACK_IN_PROGRESS: The resources are being rolled back after the stack fails to be updated by using imported resources.
	// *   IMPORT_UPDATE_ROLLBACK_FAILED: The resources fail to be rolled back after the stack fails to be updated by using imported resources.
	// *   IMPORT_UPDATE_ROLLBACK_COMPLETE: The resources are rolled back after the stack fails to be updated by using imported resources.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the stack is in its current state.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The tags of the stack.
	Tags []*GetStackResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The description of the template.
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// The ID of the template. This parameter is returned only if the current template of the stack is a custom template or a shared template.
	//
	// If the template is a shared template, the value of this parameter is the same as the value of the TemplateARN parameter.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the scenario. This parameter is returned only if the current template of the stack is generated from a scenario.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The URL of the file that contains the template body. This parameter is returned only if the current template of the stack is from a URL. The URL can point to a template that is located on an HTTP or HTTPS web server or in an Alibaba Cloud Object Storage Service (OSS) bucket.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. This parameter is returned only if the current template of the stack is a custom template or a shared template.
	//
	// If the template is a shared template, this parameter is returned only when the VersionOption parameter is set to AllVersions.
	//
	// Valid values: v1 to v100.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The timeout period within which the stack can be created. Unit: minutes.
	TimeoutInMinutes *int32 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	// The time when the stack was updated. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackResponseBody) SetCheckedStackResourceCount(v int32) *GetStackResponseBody {
	s.CheckedStackResourceCount = &v
	return s
}

func (s *GetStackResponseBody) SetCreateTime(v string) *GetStackResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetStackResponseBody) SetDeletionProtection(v string) *GetStackResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *GetStackResponseBody) SetDescription(v string) *GetStackResponseBody {
	s.Description = &v
	return s
}

func (s *GetStackResponseBody) SetDisableRollback(v bool) *GetStackResponseBody {
	s.DisableRollback = &v
	return s
}

func (s *GetStackResponseBody) SetDriftDetectionTime(v string) *GetStackResponseBody {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackResponseBody) SetInterface(v string) *GetStackResponseBody {
	s.Interface = &v
	return s
}

func (s *GetStackResponseBody) SetLog(v *GetStackResponseBodyLog) *GetStackResponseBody {
	s.Log = v
	return s
}

func (s *GetStackResponseBody) SetNotCheckedStackResourceCount(v int32) *GetStackResponseBody {
	s.NotCheckedStackResourceCount = &v
	return s
}

func (s *GetStackResponseBody) SetNotificationURLs(v []*string) *GetStackResponseBody {
	s.NotificationURLs = v
	return s
}

func (s *GetStackResponseBody) SetOperationInfo(v *GetStackResponseBodyOperationInfo) *GetStackResponseBody {
	s.OperationInfo = v
	return s
}

func (s *GetStackResponseBody) SetOrderIds(v []*string) *GetStackResponseBody {
	s.OrderIds = v
	return s
}

func (s *GetStackResponseBody) SetOutputs(v []map[string]interface{}) *GetStackResponseBody {
	s.Outputs = v
	return s
}

func (s *GetStackResponseBody) SetParameters(v []*GetStackResponseBodyParameters) *GetStackResponseBody {
	s.Parameters = v
	return s
}

func (s *GetStackResponseBody) SetParentStackId(v string) *GetStackResponseBody {
	s.ParentStackId = &v
	return s
}

func (s *GetStackResponseBody) SetRamRoleName(v string) *GetStackResponseBody {
	s.RamRoleName = &v
	return s
}

func (s *GetStackResponseBody) SetRegionId(v string) *GetStackResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetStackResponseBody) SetRequestId(v string) *GetStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackResponseBody) SetResourceGroupId(v string) *GetStackResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetStackResponseBody) SetResourceProgress(v *GetStackResponseBodyResourceProgress) *GetStackResponseBody {
	s.ResourceProgress = v
	return s
}

func (s *GetStackResponseBody) SetRootStackId(v string) *GetStackResponseBody {
	s.RootStackId = &v
	return s
}

func (s *GetStackResponseBody) SetServiceManaged(v bool) *GetStackResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *GetStackResponseBody) SetServiceName(v string) *GetStackResponseBody {
	s.ServiceName = &v
	return s
}

func (s *GetStackResponseBody) SetStackDriftStatus(v string) *GetStackResponseBody {
	s.StackDriftStatus = &v
	return s
}

func (s *GetStackResponseBody) SetStackId(v string) *GetStackResponseBody {
	s.StackId = &v
	return s
}

func (s *GetStackResponseBody) SetStackName(v string) *GetStackResponseBody {
	s.StackName = &v
	return s
}

func (s *GetStackResponseBody) SetStackType(v string) *GetStackResponseBody {
	s.StackType = &v
	return s
}

func (s *GetStackResponseBody) SetStatus(v string) *GetStackResponseBody {
	s.Status = &v
	return s
}

func (s *GetStackResponseBody) SetStatusReason(v string) *GetStackResponseBody {
	s.StatusReason = &v
	return s
}

func (s *GetStackResponseBody) SetTags(v []*GetStackResponseBodyTags) *GetStackResponseBody {
	s.Tags = v
	return s
}

func (s *GetStackResponseBody) SetTemplateDescription(v string) *GetStackResponseBody {
	s.TemplateDescription = &v
	return s
}

func (s *GetStackResponseBody) SetTemplateId(v string) *GetStackResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetStackResponseBody) SetTemplateScratchId(v string) *GetStackResponseBody {
	s.TemplateScratchId = &v
	return s
}

func (s *GetStackResponseBody) SetTemplateURL(v string) *GetStackResponseBody {
	s.TemplateURL = &v
	return s
}

func (s *GetStackResponseBody) SetTemplateVersion(v string) *GetStackResponseBody {
	s.TemplateVersion = &v
	return s
}

func (s *GetStackResponseBody) SetTimeoutInMinutes(v int32) *GetStackResponseBody {
	s.TimeoutInMinutes = &v
	return s
}

func (s *GetStackResponseBody) SetUpdateTime(v string) *GetStackResponseBody {
	s.UpdateTime = &v
	return s
}

type GetStackResponseBodyLog struct {
	// The logs of resources in the stack. This parameter is returned if the LogOption parameter is set to Resource or All.
	//
	// >  The logs are returned for resources of specific types, such as `ALIYUN::ROS::ResourceCleaner`.
	ResourceLogs []*GetStackResponseBodyLogResourceLogs `json:"ResourceLogs,omitempty" xml:"ResourceLogs,omitempty" type:"Repeated"`
	// The logs of the Terraform stack. This parameter is returned only for a Terraform stack.
	//
	// This parameter is returned if the LogOption parameter is left empty or set to Stack or All.
	//
	// >  This parameter is not returned for a running stack. The logs are generated from the last creation, re-creation, update, or deletion operation on the stack.
	TerraformLogs []*GetStackResponseBodyLogTerraformLogs `json:"TerraformLogs,omitempty" xml:"TerraformLogs,omitempty" type:"Repeated"`
}

func (s GetStackResponseBodyLog) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyLog) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyLog) SetResourceLogs(v []*GetStackResponseBodyLogResourceLogs) *GetStackResponseBodyLog {
	s.ResourceLogs = v
	return s
}

func (s *GetStackResponseBodyLog) SetTerraformLogs(v []*GetStackResponseBodyLogTerraformLogs) *GetStackResponseBodyLog {
	s.TerraformLogs = v
	return s
}

type GetStackResponseBodyLogResourceLogs struct {
	// The logs of all resources.
	Logs []*GetStackResponseBodyLogResourceLogsLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The name of the resource that is defined in the template.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s GetStackResponseBodyLogResourceLogs) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyLogResourceLogs) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyLogResourceLogs) SetLogs(v []*GetStackResponseBodyLogResourceLogsLogs) *GetStackResponseBodyLogResourceLogs {
	s.Logs = v
	return s
}

func (s *GetStackResponseBodyLogResourceLogs) SetResourceName(v string) *GetStackResponseBodyLogResourceLogs {
	s.ResourceName = &v
	return s
}

type GetStackResponseBodyLogResourceLogsLogs struct {
	// The content of a resource log.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The keywords of a resource log.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s GetStackResponseBodyLogResourceLogsLogs) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyLogResourceLogsLogs) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyLogResourceLogsLogs) SetContent(v string) *GetStackResponseBodyLogResourceLogsLogs {
	s.Content = &v
	return s
}

func (s *GetStackResponseBodyLogResourceLogsLogs) SetKeys(v []*string) *GetStackResponseBodyLogResourceLogsLogs {
	s.Keys = v
	return s
}

type GetStackResponseBodyLogTerraformLogs struct {
	// The name of the Terraform command that is run. Valid values:
	//
	// *   apply
	// *   plan
	// *   destroy
	// *   version
	//
	// For more information about Terraform commands, see [Basic CLI Features](https://www.terraform.io/cli/commands).
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The content of the output stream that is returned after the command is run.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The output stream. Valid values:
	//
	// *   stdout: standard output stream
	// *   stderr: standard error stream
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s GetStackResponseBodyLogTerraformLogs) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyLogTerraformLogs) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyLogTerraformLogs) SetCommand(v string) *GetStackResponseBodyLogTerraformLogs {
	s.Command = &v
	return s
}

func (s *GetStackResponseBodyLogTerraformLogs) SetContent(v string) *GetStackResponseBodyLogTerraformLogs {
	s.Content = &v
	return s
}

func (s *GetStackResponseBodyLogTerraformLogs) SetStream(v string) *GetStackResponseBodyLogTerraformLogs {
	s.Stream = &v
	return s
}

type GetStackResponseBodyOperationInfo struct {
	// The name of the API of another cloud service.
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The error code returned.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The logical ID of the resource on which the operation error occurred.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request to call the API of another cloud service.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the resource on which the operation error occurred.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetStackResponseBodyOperationInfo) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyOperationInfo) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyOperationInfo) SetAction(v string) *GetStackResponseBodyOperationInfo {
	s.Action = &v
	return s
}

func (s *GetStackResponseBodyOperationInfo) SetCode(v string) *GetStackResponseBodyOperationInfo {
	s.Code = &v
	return s
}

func (s *GetStackResponseBodyOperationInfo) SetLogicalResourceId(v string) *GetStackResponseBodyOperationInfo {
	s.LogicalResourceId = &v
	return s
}

func (s *GetStackResponseBodyOperationInfo) SetMessage(v string) *GetStackResponseBodyOperationInfo {
	s.Message = &v
	return s
}

func (s *GetStackResponseBodyOperationInfo) SetRequestId(v string) *GetStackResponseBodyOperationInfo {
	s.RequestId = &v
	return s
}

func (s *GetStackResponseBodyOperationInfo) SetResourceType(v string) *GetStackResponseBodyOperationInfo {
	s.ResourceType = &v
	return s
}

type GetStackResponseBodyParameters struct {
	// The name of the parameter.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetStackResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyParameters) SetParameterKey(v string) *GetStackResponseBodyParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetStackResponseBodyParameters) SetParameterValue(v string) *GetStackResponseBodyParameters {
	s.ParameterValue = &v
	return s
}

type GetStackResponseBodyResourceProgress struct {
	// The number of resources that fail to be created.
	FailedResourceCount *int32 `json:"FailedResourceCount,omitempty" xml:"FailedResourceCount,omitempty"`
	// The number of resources that are being created.
	InProgressResourceCount *int32 `json:"InProgressResourceCount,omitempty" xml:"InProgressResourceCount,omitempty"`
	// The progress details of resources that are being created.
	InProgressResourceDetails []*GetStackResponseBodyResourceProgressInProgressResourceDetails `json:"InProgressResourceDetails,omitempty" xml:"InProgressResourceDetails,omitempty" type:"Repeated"`
	// The number of resources to be created.
	PendingResourceCount   *int32   `json:"PendingResourceCount,omitempty" xml:"PendingResourceCount,omitempty"`
	StackActionProgress    *float32 `json:"StackActionProgress,omitempty" xml:"StackActionProgress,omitempty"`
	StackOperationProgress *float32 `json:"StackOperationProgress,omitempty" xml:"StackOperationProgress,omitempty"`
	// The number of resources that are created.
	SuccessResourceCount *int32 `json:"SuccessResourceCount,omitempty" xml:"SuccessResourceCount,omitempty"`
	// The total number of resources.
	TotalResourceCount *int32 `json:"TotalResourceCount,omitempty" xml:"TotalResourceCount,omitempty"`
}

func (s GetStackResponseBodyResourceProgress) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyResourceProgress) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyResourceProgress) SetFailedResourceCount(v int32) *GetStackResponseBodyResourceProgress {
	s.FailedResourceCount = &v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetInProgressResourceCount(v int32) *GetStackResponseBodyResourceProgress {
	s.InProgressResourceCount = &v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetInProgressResourceDetails(v []*GetStackResponseBodyResourceProgressInProgressResourceDetails) *GetStackResponseBodyResourceProgress {
	s.InProgressResourceDetails = v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetPendingResourceCount(v int32) *GetStackResponseBodyResourceProgress {
	s.PendingResourceCount = &v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetStackActionProgress(v float32) *GetStackResponseBodyResourceProgress {
	s.StackActionProgress = &v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetStackOperationProgress(v float32) *GetStackResponseBodyResourceProgress {
	s.StackOperationProgress = &v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetSuccessResourceCount(v int32) *GetStackResponseBodyResourceProgress {
	s.SuccessResourceCount = &v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetTotalResourceCount(v int32) *GetStackResponseBodyResourceProgress {
	s.TotalResourceCount = &v
	return s
}

type GetStackResponseBodyResourceProgressInProgressResourceDetails struct {
	// The desired progress value of the resource.
	ProgressTargetValue *float32 `json:"ProgressTargetValue,omitempty" xml:"ProgressTargetValue,omitempty"`
	// The current progress value of the resource.
	ProgressValue *float32 `json:"ProgressValue,omitempty" xml:"ProgressValue,omitempty"`
	// The resource name.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetStackResponseBodyResourceProgressInProgressResourceDetails) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyResourceProgressInProgressResourceDetails) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyResourceProgressInProgressResourceDetails) SetProgressTargetValue(v float32) *GetStackResponseBodyResourceProgressInProgressResourceDetails {
	s.ProgressTargetValue = &v
	return s
}

func (s *GetStackResponseBodyResourceProgressInProgressResourceDetails) SetProgressValue(v float32) *GetStackResponseBodyResourceProgressInProgressResourceDetails {
	s.ProgressValue = &v
	return s
}

func (s *GetStackResponseBodyResourceProgressInProgressResourceDetails) SetResourceName(v string) *GetStackResponseBodyResourceProgressInProgressResourceDetails {
	s.ResourceName = &v
	return s
}

func (s *GetStackResponseBodyResourceProgressInProgressResourceDetails) SetResourceType(v string) *GetStackResponseBodyResourceProgressInProgressResourceDetails {
	s.ResourceType = &v
	return s
}

type GetStackResponseBodyTags struct {
	// The tag key of the stack.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the stack.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetStackResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyTags) SetKey(v string) *GetStackResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetStackResponseBodyTags) SetValue(v string) *GetStackResponseBodyTags {
	s.Value = &v
	return s
}

type GetStackResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackResponse) GoString() string {
	return s.String()
}

func (s *GetStackResponse) SetHeaders(v map[string]*string) *GetStackResponse {
	s.Headers = v
	return s
}

func (s *GetStackResponse) SetStatusCode(v int32) *GetStackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackResponse) SetBody(v *GetStackResponseBody) *GetStackResponse {
	s.Body = v
	return s
}

type GetStackDriftDetectionStatusRequest struct {
	// The ID of the drift detection operation.
	//
	// You can call the [ListStackResourceDrifts](~~155098~~) operation to obtain the ID of the drift detection operation.
	DriftDetectionId *string `json:"DriftDetectionId,omitempty" xml:"DriftDetectionId,omitempty"`
	// The region ID of the stack to be detected for drift.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetStackDriftDetectionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackDriftDetectionStatusRequest) GoString() string {
	return s.String()
}

func (s *GetStackDriftDetectionStatusRequest) SetDriftDetectionId(v string) *GetStackDriftDetectionStatusRequest {
	s.DriftDetectionId = &v
	return s
}

func (s *GetStackDriftDetectionStatusRequest) SetRegionId(v string) *GetStackDriftDetectionStatusRequest {
	s.RegionId = &v
	return s
}

type GetStackDriftDetectionStatusResponseBody struct {
	// The ID of the drift detection operation.
	DriftDetectionId *string `json:"DriftDetectionId,omitempty" xml:"DriftDetectionId,omitempty"`
	// The drift detection status. Valid values:
	//
	// *   DETECTION_COMPLETE: The drift detection operation has been completed for all resources that support drift detection in the stack.
	// *   DETECTION_FAILED: The stack drift detection operation has failed for at least one resource in the stack.
	// *   DETECTION_IN_PROGRESS: The stack drift detection operation is in progress.
	DriftDetectionStatus *string `json:"DriftDetectionStatus,omitempty" xml:"DriftDetectionStatus,omitempty"`
	// The reason why the stack drift detection operation has its current status.
	DriftDetectionStatusReason *string `json:"DriftDetectionStatusReason,omitempty" xml:"DriftDetectionStatusReason,omitempty"`
	// The time when the stack drift detection operation was initiated.
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The total number of stack resources that have drifted.
	DriftedStackResourceCount *int32 `json:"DriftedStackResourceCount,omitempty" xml:"DriftedStackResourceCount,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The drift status of the stack. Valid values:
	//
	// *   DRIFTED: The actual configuration of the stack differs, or has drifted, from its expected template configuration. A stack is considered to have drifted if one or more of its resources have drifted.
	// *   NOT_CHECKED: Resource Orchestration Service (ROS) has not checked whether the actual configuration of the resource differs from its expected template configuration.
	// *   IN_SYNC: The current configuration of each supported resource matches its expected template configuration. A stack with no resources that support drift detection also has a status of IN_SYNC.
	StackDriftStatus *string `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s GetStackDriftDetectionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackDriftDetectionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftDetectionId(v string) *GetStackDriftDetectionStatusResponseBody {
	s.DriftDetectionId = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftDetectionStatus(v string) *GetStackDriftDetectionStatusResponseBody {
	s.DriftDetectionStatus = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftDetectionStatusReason(v string) *GetStackDriftDetectionStatusResponseBody {
	s.DriftDetectionStatusReason = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftDetectionTime(v string) *GetStackDriftDetectionStatusResponseBody {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetDriftedStackResourceCount(v int32) *GetStackDriftDetectionStatusResponseBody {
	s.DriftedStackResourceCount = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetRequestId(v string) *GetStackDriftDetectionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetStackDriftStatus(v string) *GetStackDriftDetectionStatusResponseBody {
	s.StackDriftStatus = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponseBody) SetStackId(v string) *GetStackDriftDetectionStatusResponseBody {
	s.StackId = &v
	return s
}

type GetStackDriftDetectionStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStackDriftDetectionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackDriftDetectionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackDriftDetectionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetStackDriftDetectionStatusResponse) SetHeaders(v map[string]*string) *GetStackDriftDetectionStatusResponse {
	s.Headers = v
	return s
}

func (s *GetStackDriftDetectionStatusResponse) SetStatusCode(v int32) *GetStackDriftDetectionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackDriftDetectionStatusResponse) SetBody(v *GetStackDriftDetectionStatusResponseBody) *GetStackDriftDetectionStatusResponse {
	s.Body = v
	return s
}

type GetStackGroupRequest struct {
	// The name of the stack group. The name must be unique within a region.
	//
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or letter.
	//
	// >  You must specify one of the StackGroupName and StackGroupId parameters.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	// The ID of the stack group.
	//
	// >  You must specify one of the StackGroupName and StackGroupId parameters.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s GetStackGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupRequest) GoString() string {
	return s.String()
}

func (s *GetStackGroupRequest) SetRegionId(v string) *GetStackGroupRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackGroupRequest) SetStackGroupId(v string) *GetStackGroupRequest {
	s.StackGroupId = &v
	return s
}

func (s *GetStackGroupRequest) SetStackGroupName(v string) *GetStackGroupRequest {
	s.StackGroupName = &v
	return s
}

type GetStackGroupResponseBody struct {
	// The details of the stack group.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the stack group.
	StackGroup *GetStackGroupResponseBodyStackGroup `json:"StackGroup,omitempty" xml:"StackGroup,omitempty" type:"Struct"`
}

func (s GetStackGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBody) SetRequestId(v string) *GetStackGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackGroupResponseBody) SetStackGroup(v *GetStackGroupResponseBodyStackGroup) *GetStackGroupResponseBody {
	s.StackGroup = v
	return s
}

type GetStackGroupResponseBodyStackGroup struct {
	// The parameters of the stack group.
	AdministrationRoleName *string `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	// Indicates whether automatic deployment is enabled.
	//
	// Valid values:
	//
	// *   true: Automatic deployment is enabled. If a member account is added to the folder to which the stack group belongs after automatic deployment is enabled, the stack group deploys its stack instances in the specified region where the added account is deployed. If the account is deleted from the folder, the stack instances in the specified region are deleted from the stack group.
	// *   false: Automatic deployment is disabled. After automatic deployment is disabled, the stack instances remain unchanged when the member account in the folder is changed.
	AutoDeployment *GetStackGroupResponseBodyStackGroupAutoDeployment `json:"AutoDeployment,omitempty" xml:"AutoDeployment,omitempty" type:"Struct"`
	// The name of the stack group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The template body.
	ExecutionRoleName *string `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	// The key of the parameter.
	Parameters []*GetStackGroupResponseBodyStackGroupParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The information about automatic deployment settings.
	//
	// >  This parameter is returned only when the PermissionModel parameter is set to SERVICE_MANAGED.
	PermissionModel *string `json:"PermissionModel,omitempty" xml:"PermissionModel,omitempty"`
	// The folder IDs of the resource directory. This parameter is used to deploy stack instances within all the accounts in the folders.
	//
	// >  This parameter is returned only when the PermissionModel parameter is set to SERVICE_MANAGED.
	RdFolderIds []*string `json:"RdFolderIds,omitempty" xml:"RdFolderIds,omitempty" type:"Repeated"`
	// The permission model.
	//
	// Valid values:
	//
	// *   SELF_MANAGED: the self-managed permission model
	// *   SERVICE_MANAGED: the service-managed permission model
	//
	// >  For more information about the permission models of stack groups, see [Overview](~~154578~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The time when drift detection was performed on the stack group.
	StackGroupDriftDetectionDetail *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail `json:"StackGroupDriftDetectionDetail,omitempty" xml:"StackGroupDriftDetectionDetail,omitempty" type:"Struct"`
	// The status of the stack group.
	//
	// Valid values:
	//
	// *   ACTIVE
	// *   DELETED
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	// The name of the RAM role that is specified for the execution account when you create the self-managed stack group. The administrator role AliyunROSStackGroupAdministrationRole assumes the execution role. If this parameter is not specified, the default value AliyunROSStackGroupExecutionRole is returned.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The name of the RAM role that is specified for the administrator account in Resource Orchestration Service (ROS) when you create the self-managed stack group. If this parameter is not specified, the default value AliyunROSStackGroupAdministrationRole is returned.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The structure that contains the template body.
	//
	// > We recommend that you use TemplateContent instead of TemplateBody.
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The JSON-formatted structure that contains the template body. For more information, see [Template syntax](~~28857~~).
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
}

func (s GetStackGroupResponseBodyStackGroup) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupResponseBodyStackGroup) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBodyStackGroup) SetAdministrationRoleName(v string) *GetStackGroupResponseBodyStackGroup {
	s.AdministrationRoleName = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetAutoDeployment(v *GetStackGroupResponseBodyStackGroupAutoDeployment) *GetStackGroupResponseBodyStackGroup {
	s.AutoDeployment = v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetDescription(v string) *GetStackGroupResponseBodyStackGroup {
	s.Description = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetExecutionRoleName(v string) *GetStackGroupResponseBodyStackGroup {
	s.ExecutionRoleName = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetParameters(v []*GetStackGroupResponseBodyStackGroupParameters) *GetStackGroupResponseBodyStackGroup {
	s.Parameters = v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetPermissionModel(v string) *GetStackGroupResponseBodyStackGroup {
	s.PermissionModel = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetRdFolderIds(v []*string) *GetStackGroupResponseBodyStackGroup {
	s.RdFolderIds = v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetResourceGroupId(v string) *GetStackGroupResponseBodyStackGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetStackGroupDriftDetectionDetail(v *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) *GetStackGroupResponseBodyStackGroup {
	s.StackGroupDriftDetectionDetail = v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetStackGroupId(v string) *GetStackGroupResponseBodyStackGroup {
	s.StackGroupId = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetStackGroupName(v string) *GetStackGroupResponseBodyStackGroup {
	s.StackGroupName = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetStatus(v string) *GetStackGroupResponseBodyStackGroup {
	s.Status = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetTemplateBody(v string) *GetStackGroupResponseBodyStackGroup {
	s.TemplateBody = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroup) SetTemplateContent(v string) *GetStackGroupResponseBodyStackGroup {
	s.TemplateContent = &v
	return s
}

type GetStackGroupResponseBodyStackGroupAutoDeployment struct {
	// Indicates whether stacks in the member account are retained when the member account is deleted from the folder.
	//
	// Valid values:
	//
	// *   true: The stacks are retained.
	// *   false: The stacks are deleted.
	//
	// >  This parameter is returned only when the Enabled parameter is set to true.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The folder IDs of the resource directory. This parameter is used to deploy stack instances within all the accounts in the folders.
	//
	// >  This parameter is returned only when the PermissionModel parameter is set to SERVICE_MANAGED.
	RetainStacksOnAccountRemoval *bool `json:"RetainStacksOnAccountRemoval,omitempty" xml:"RetainStacksOnAccountRemoval,omitempty"`
}

func (s GetStackGroupResponseBodyStackGroupAutoDeployment) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupResponseBodyStackGroupAutoDeployment) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBodyStackGroupAutoDeployment) SetEnabled(v bool) *GetStackGroupResponseBodyStackGroupAutoDeployment {
	s.Enabled = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupAutoDeployment) SetRetainStacksOnAccountRemoval(v bool) *GetStackGroupResponseBodyStackGroupAutoDeployment {
	s.RetainStacksOnAccountRemoval = &v
	return s
}

type GetStackGroupResponseBodyStackGroupParameters struct {
	// The name of the parameter.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetStackGroupResponseBodyStackGroupParameters) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupResponseBodyStackGroupParameters) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBodyStackGroupParameters) SetParameterKey(v string) *GetStackGroupResponseBodyStackGroupParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupParameters) SetParameterValue(v string) *GetStackGroupResponseBodyStackGroupParameters {
	s.ParameterValue = &v
	return s
}

type GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail struct {
	// The number of stack instances that have drifted.
	CancelledStackInstancesCount *int32 `json:"CancelledStackInstancesCount,omitempty" xml:"CancelledStackInstancesCount,omitempty"`
	// The drift status of the stack group.
	//
	// Valid values:
	//
	// *   DRIFTED: At least one stack instance in the stack group has drifted.
	// *   NOT_CHECKED: No drift detection is completed on the stack group.
	// *   IN_SYNC: All the stack instances in the stack group are being synchronized.
	DriftDetectionStatus *string `json:"DriftDetectionStatus,omitempty" xml:"DriftDetectionStatus,omitempty"`
	// The number of stack instances.
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The ID of the resource group. This parameter is specified when you create the stack group.
	DriftedStackInstancesCount *int32 `json:"DriftedStackInstancesCount,omitempty" xml:"DriftedStackInstancesCount,omitempty"`
	// The status of drift detection on the stack group.
	//
	// Valid values:
	//
	// *   COMPLETED: Drift detection is performed and completed on all stack instances.
	// *   FAILED: Drift detection is performed. The number of stack instances that failed the drift detection exceeds the specified threshold.
	// *   PARTIAL_SUCCESS: Drift detection is performed. The number of stack instances that failed the drift detection does not exceed the specified threshold.
	// *   IN_PROGRESS: Drift detection is being performed on the stack group.
	// *   STOPPED: Drift detection is canceled for the stack group.
	FailedStackInstancesCount *int32 `json:"FailedStackInstancesCount,omitempty" xml:"FailedStackInstancesCount,omitempty"`
	// The number of stack instances that were being synchronized.
	InProgressStackInstancesCount *int32 `json:"InProgressStackInstancesCount,omitempty" xml:"InProgressStackInstancesCount,omitempty"`
	// The number of stack instances for which drift detection was canceled.
	InSyncStackInstancesCount *int32 `json:"InSyncStackInstancesCount,omitempty" xml:"InSyncStackInstancesCount,omitempty"`
	// The number of stack instances on which drift detection was being performed.
	StackGroupDriftStatus *string `json:"StackGroupDriftStatus,omitempty" xml:"StackGroupDriftStatus,omitempty"`
	// The number of stack instances that failed drift detection.
	TotalStackInstancesCount *int32 `json:"TotalStackInstancesCount,omitempty" xml:"TotalStackInstancesCount,omitempty"`
}

func (s GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetCancelledStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.CancelledStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetDriftDetectionStatus(v string) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.DriftDetectionStatus = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetDriftDetectionTime(v string) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetDriftedStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.DriftedStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetFailedStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.FailedStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetInProgressStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.InProgressStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetInSyncStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.InSyncStackInstancesCount = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetStackGroupDriftStatus(v string) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.StackGroupDriftStatus = &v
	return s
}

func (s *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail) SetTotalStackInstancesCount(v int32) *GetStackGroupResponseBodyStackGroupStackGroupDriftDetectionDetail {
	s.TotalStackInstancesCount = &v
	return s
}

type GetStackGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStackGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupResponse) GoString() string {
	return s.String()
}

func (s *GetStackGroupResponse) SetHeaders(v map[string]*string) *GetStackGroupResponse {
	s.Headers = v
	return s
}

func (s *GetStackGroupResponse) SetStatusCode(v int32) *GetStackGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackGroupResponse) SetBody(v *GetStackGroupResponseBody) *GetStackGroupResponse {
	s.Body = v
	return s
}

type GetStackGroupOperationRequest struct {
	// The operation ID. You can call the [ListStackGroupOperations](~~151342~~) operation to query the operation ID.
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetStackGroupOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupOperationRequest) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationRequest) SetOperationId(v string) *GetStackGroupOperationRequest {
	s.OperationId = &v
	return s
}

func (s *GetStackGroupOperationRequest) SetRegionId(v string) *GetStackGroupOperationRequest {
	s.RegionId = &v
	return s
}

type GetStackGroupOperationResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the stack group operation.
	StackGroupOperation *GetStackGroupOperationResponseBodyStackGroupOperation `json:"StackGroupOperation,omitempty" xml:"StackGroupOperation,omitempty" type:"Struct"`
}

func (s GetStackGroupOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupOperationResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBody) SetRequestId(v string) *GetStackGroupOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackGroupOperationResponseBody) SetStackGroupOperation(v *GetStackGroupOperationResponseBodyStackGroupOperation) *GetStackGroupOperationResponseBody {
	s.StackGroupOperation = v
	return s
}

type GetStackGroupOperationResponseBodyStackGroupOperation struct {
	// The operation type.
	//
	// Valid values:
	//
	// *   CREATE
	// *   UPDATE
	// *   DELETE
	// *   DETECT_DRIFT
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The name of the RAM role that you specify for the administrator account when you create the self-managed stack group. ROS assumes the administrator role to perform operations. If this parameter is not specified, the default value AliyunROSStackGroupAdministrationRole is returned.
	AdministrationRoleName *string `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	// The time when the operation was initiated.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The destinations to deploy stack instances when the stack is granted service-managed permissions.
	DeploymentTargets *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty" type:"Struct"`
	// The time when the operation ended.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the RAM role that you specify for the execution account when you create the self-managed stack group. The administrator role AliyunROSStackGroupAdministrationRole assumes the execution role to perform operations. If this parameter is not specified, the default value AliyunROSStackGroupExecutionRole is returned.
	ExecutionRoleName *string `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	// The description of the operation.
	//
	// > This parameter is returned only if OperationDescription is specified when the [CreateStackInstances](~~151338~~) operation is called to create stack instances.
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The operation ID.
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The operation settings.
	OperationPreferences *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty" type:"Struct"`
	// Indicates whether stacks are retained when the associated stack instances are deleted. When you delete a stack instance, you can choose to delete or retain the stack with which the stack instance is associated.
	//
	// Valid values:
	//
	// *   true: Stacks are retained when the associated stack instances are deleted.
	// *   false: Stacks are deleted when the associated stack instances are deleted. Proceed with caution.
	//
	// > This parameter is returned only if you delete stack instances.
	RetainStacks *bool `json:"RetainStacks,omitempty" xml:"RetainStacks,omitempty"`
	// The information about drift detection.
	//
	// > This parameter is returned only if drift detection is performed.
	StackGroupDriftDetectionDetail *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail `json:"StackGroupDriftDetectionDetail,omitempty" xml:"StackGroupDriftDetectionDetail,omitempty" type:"Struct"`
	// The ID of the stack group.
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	// The name of the stack group.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The state of the operation.
	//
	// Valid values:
	//
	// *   RUNNING
	// *   SUCCEEDED
	// *   FAILED
	// *   STOPPING
	// *   STOPPED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetStackGroupOperationResponseBodyStackGroupOperation) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupOperationResponseBodyStackGroupOperation) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetAction(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.Action = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetAdministrationRoleName(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.AdministrationRoleName = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetCreateTime(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.CreateTime = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetDeploymentTargets(v *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.DeploymentTargets = v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetEndTime(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.EndTime = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetExecutionRoleName(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.ExecutionRoleName = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetOperationDescription(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.OperationDescription = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetOperationId(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.OperationId = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetOperationPreferences(v *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.OperationPreferences = v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetRetainStacks(v bool) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.RetainStacks = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetStackGroupDriftDetectionDetail(v *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.StackGroupDriftDetectionDetail = v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetStackGroupId(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.StackGroupId = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetStackGroupName(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.StackGroupName = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetStatus(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.Status = &v
	return s
}

type GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets struct {
	// The IDs of the members in the resource directory.
	//
	// > This parameter is returned only if AccountIds is specified when the [UpdateStackInstances](~~151716~~) operation is called to update stack instances.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The IDs of the folders in the resource directory.
	RdFolderIds []*string `json:"RdFolderIds,omitempty" xml:"RdFolderIds,omitempty" type:"Repeated"`
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets) SetAccountIds(v []*string) *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets {
	s.AccountIds = v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets) SetRdFolderIds(v []*string) *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets {
	s.RdFolderIds = v
	return s
}

type GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences struct {
	// The number of accounts within which stack operation failures are allowed to occur in each region. If the value of this parameter is exceeded in a region, Resource Orchestration Service (ROS) stops the operation in the region. If the operation is stopped in one region, the operation is no longer performed in other regions.
	//
	// Valid values: 0 to 20.
	//
	// > Only one of FailureToleranceCount and FailureTolerancePercentage can be returned.
	FailureToleranceCount *int32 `json:"FailureToleranceCount,omitempty" xml:"FailureToleranceCount,omitempty"`
	// The percentage of the number of accounts within which stack operation failures are allowed to occur to the total number of accounts in each region. If the value of this parameter is exceeded in a region, ROS stops the operation in the region.
	//
	// Valid values: 0 to 100.
	//
	// > Only one of FailureToleranceCount and FailureTolerancePercentage can be returned.
	FailureTolerancePercentage *int32 `json:"FailureTolerancePercentage,omitempty" xml:"FailureTolerancePercentage,omitempty"`
	// The maximum number of accounts within which stacks are deployed at the same time in each region.
	//
	// Valid values: 1 to 20.
	//
	// > Only one of MaxConcurrentCount and MaxConcurrentPercentage can be returned.
	MaxConcurrentCount *int32 `json:"MaxConcurrentCount,omitempty" xml:"MaxConcurrentCount,omitempty"`
	// The percentage of the maximum number of accounts within which stacks are deployed at the same time to the total number of accounts in each region.
	//
	// Valid values: 1 to 100.
	//
	// > Only one of MaxConcurrentCount and MaxConcurrentPercentage can be returned.
	MaxConcurrentPercentage *int32 `json:"MaxConcurrentPercentage,omitempty" xml:"MaxConcurrentPercentage,omitempty"`
	// The regions in the order of operation execution.
	RegionIdsOrder []*string `json:"RegionIdsOrder,omitempty" xml:"RegionIdsOrder,omitempty" type:"Repeated"`
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetFailureToleranceCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.FailureToleranceCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetFailureTolerancePercentage(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.FailureTolerancePercentage = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetMaxConcurrentCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.MaxConcurrentCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetMaxConcurrentPercentage(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.MaxConcurrentPercentage = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetRegionIdsOrder(v []*string) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.RegionIdsOrder = v
	return s
}

type GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail struct {
	// The number of stack instances for which drift detection was canceled.
	CancelledStackInstancesCount *int32 `json:"CancelledStackInstancesCount,omitempty" xml:"CancelledStackInstancesCount,omitempty"`
	// The drift detection state.
	//
	// Valid values:
	//
	// *   COMPLETED: Drift detection is performed on the stack group and all stack instances passed the drift detection.
	// *   FAILED: Drift detection is performed on the stack group. The number of stack instances that failed the drift detection exceeds the specified threshold.
	// *   PARTIAL_SUCCESS: Drift detection is performed on the stack group. The number of stack instances that failed the drift detection does not exceed the specified threshold.
	// *   IN_PROGRESS: Drift detection is being performed on the stack group.
	// *   STOPPED: Drift detection is canceled for the stack group.
	DriftDetectionStatus *string `json:"DriftDetectionStatus,omitempty" xml:"DriftDetectionStatus,omitempty"`
	// The time when drift detection was performed.
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The number of stack instances that have drifted.
	DriftedStackInstancesCount *int32 `json:"DriftedStackInstancesCount,omitempty" xml:"DriftedStackInstancesCount,omitempty"`
	// The number of stack instances that failed drift detection.
	FailedStackInstancesCount *int32 `json:"FailedStackInstancesCount,omitempty" xml:"FailedStackInstancesCount,omitempty"`
	// The number of stack instances on which drift detection was being performed.
	InProgressStackInstancesCount *int32 `json:"InProgressStackInstancesCount,omitempty" xml:"InProgressStackInstancesCount,omitempty"`
	// The number of stack instances that were being synchronized.
	InSyncStackInstancesCount *int32 `json:"InSyncStackInstancesCount,omitempty" xml:"InSyncStackInstancesCount,omitempty"`
	// The drift state of the stack group.
	//
	// Valid values:
	//
	// *   DRIFTED: At least one stack instance in the stack group has drifted.
	// *   NOT_CHECKED: No successful drift detection is performed in the stack group.
	// *   IN_SYNC: All the stack instances in the stack group are being synchronized.
	StackGroupDriftStatus *string `json:"StackGroupDriftStatus,omitempty" xml:"StackGroupDriftStatus,omitempty"`
	// The number of stack instances.
	TotalStackInstancesCount *int32 `json:"TotalStackInstancesCount,omitempty" xml:"TotalStackInstancesCount,omitempty"`
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetCancelledStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.CancelledStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetDriftDetectionStatus(v string) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.DriftDetectionStatus = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetDriftDetectionTime(v string) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetDriftedStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.DriftedStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetFailedStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.FailedStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetInProgressStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.InProgressStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetInSyncStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.InSyncStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetStackGroupDriftStatus(v string) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.StackGroupDriftStatus = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetTotalStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.TotalStackInstancesCount = &v
	return s
}

type GetStackGroupOperationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStackGroupOperationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackGroupOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackGroupOperationResponse) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponse) SetHeaders(v map[string]*string) *GetStackGroupOperationResponse {
	s.Headers = v
	return s
}

func (s *GetStackGroupOperationResponse) SetStatusCode(v int32) *GetStackGroupOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackGroupOperationResponse) SetBody(v *GetStackGroupOperationResponseBody) *GetStackGroupOperationResponse {
	s.Body = v
	return s
}

type GetStackInstanceRequest struct {
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group. The name must be unique within a region.\
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The ID of the destination account to which the stack belongs.
	//
	// *   If the stack group is granted self-managed permissions, the stack belongs to an Alibaba Cloud account.
	// *   If the stack group is granted service-managed permissions, the stack belongs to a member in a resource directory.
	//
	// > For more information about the destination account, see [Overview](~~154578~~).
	StackInstanceAccountId *string `json:"StackInstanceAccountId,omitempty" xml:"StackInstanceAccountId,omitempty"`
	// The region ID of the stack.
	StackInstanceRegionId *string `json:"StackInstanceRegionId,omitempty" xml:"StackInstanceRegionId,omitempty"`
}

func (s GetStackInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetStackInstanceRequest) SetRegionId(v string) *GetStackInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackInstanceRequest) SetStackGroupName(v string) *GetStackInstanceRequest {
	s.StackGroupName = &v
	return s
}

func (s *GetStackInstanceRequest) SetStackInstanceAccountId(v string) *GetStackInstanceRequest {
	s.StackInstanceAccountId = &v
	return s
}

func (s *GetStackInstanceRequest) SetStackInstanceRegionId(v string) *GetStackInstanceRequest {
	s.StackInstanceRegionId = &v
	return s
}

type GetStackInstanceResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the stack.
	StackInstance *GetStackInstanceResponseBodyStackInstance `json:"StackInstance,omitempty" xml:"StackInstance,omitempty" type:"Struct"`
}

func (s GetStackInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackInstanceResponseBody) SetRequestId(v string) *GetStackInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackInstanceResponseBody) SetStackInstance(v *GetStackInstanceResponseBodyStackInstance) *GetStackInstanceResponseBody {
	s.StackInstance = v
	return s
}

type GetStackInstanceResponseBodyStackInstance struct {
	// The ID of the destination account to which the stack belongs.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the most recent successful drift detection was performed on the stack group.
	//
	// > This parameter is returned only if drift detection is performed on the stack group.
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The parameters that are used to override specific parameters.
	ParameterOverrides []*GetStackInstanceResponseBodyStackInstanceParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
	// The ID of the folder in the resource directory.
	//
	// > This parameter is returned only if the stack group is granted service-managed permissions.
	RdFolderId *string `json:"RdFolderId,omitempty" xml:"RdFolderId,omitempty"`
	// The region ID of the stack.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The state of the stack when the most recent successful drift detection was performed on the stack group.
	//
	// Valid values:
	//
	// *   DRIFTED: The stack has drifted.
	// *   NOT_CHECKED: No successful drift detection is performed on the stack.
	// *   IN_SYNC: The stack is being synchronized.
	//
	// > This parameter is returned only if drift detection is performed on the stack group.
	StackDriftStatus *string `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	// The ID of the stack group.
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	// The name of the stack group.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The stack ID.
	//
	// > This parameter is returned only if the stack is in the CURRENT state.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The state of the stack.
	//
	// Valid values:
	//
	// *   CURRENT: The stack is up-to-date with the stack group.
	//
	// *   OUTDATED: The stack is not up-to-date with the stack group. Stacks are in the OUTDATED state due to the following possible reasons:
	//
	//     *   When the CreateStackInstances operation is called to create stacks, the stacks fail to be created.
	//     *   When the UpdateStackInstances or UpdateStackGroup operation is called to update stacks, the stacks fail to be updated, or only specific stacks are updated.
	//     *   The creation or update operation is not complete.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the stack instance is in the OUTDATED state.
	//
	// > This parameter is returned only if the stack instance is in the OUTDATED state.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s GetStackInstanceResponseBodyStackInstance) String() string {
	return tea.Prettify(s)
}

func (s GetStackInstanceResponseBodyStackInstance) GoString() string {
	return s.String()
}

func (s *GetStackInstanceResponseBodyStackInstance) SetAccountId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.AccountId = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetDriftDetectionTime(v string) *GetStackInstanceResponseBodyStackInstance {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetParameterOverrides(v []*GetStackInstanceResponseBodyStackInstanceParameterOverrides) *GetStackInstanceResponseBodyStackInstance {
	s.ParameterOverrides = v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetRdFolderId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.RdFolderId = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetRegionId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.RegionId = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStackDriftStatus(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StackDriftStatus = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStackGroupId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StackGroupId = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStackGroupName(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StackGroupName = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStackId(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StackId = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStatus(v string) *GetStackInstanceResponseBodyStackInstance {
	s.Status = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstance) SetStatusReason(v string) *GetStackInstanceResponseBodyStackInstance {
	s.StatusReason = &v
	return s
}

type GetStackInstanceResponseBodyStackInstanceParameterOverrides struct {
	// The name of the parameter that is used to override a specific parameter.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter that is used to override a specific parameter.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetStackInstanceResponseBodyStackInstanceParameterOverrides) String() string {
	return tea.Prettify(s)
}

func (s GetStackInstanceResponseBodyStackInstanceParameterOverrides) GoString() string {
	return s.String()
}

func (s *GetStackInstanceResponseBodyStackInstanceParameterOverrides) SetParameterKey(v string) *GetStackInstanceResponseBodyStackInstanceParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *GetStackInstanceResponseBodyStackInstanceParameterOverrides) SetParameterValue(v string) *GetStackInstanceResponseBodyStackInstanceParameterOverrides {
	s.ParameterValue = &v
	return s
}

type GetStackInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStackInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetStackInstanceResponse) SetHeaders(v map[string]*string) *GetStackInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetStackInstanceResponse) SetStatusCode(v int32) *GetStackInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackInstanceResponse) SetBody(v *GetStackInstanceResponseBody) *GetStackInstanceResponse {
	s.Body = v
	return s
}

type GetStackPolicyRequest struct {
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s GetStackPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetStackPolicyRequest) SetRegionId(v string) *GetStackPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackPolicyRequest) SetStackId(v string) *GetStackPolicyRequest {
	s.StackId = &v
	return s
}

type GetStackPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
	StackPolicyBody map[string]interface{} `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
}

func (s GetStackPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackPolicyResponseBody) SetRequestId(v string) *GetStackPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackPolicyResponseBody) SetStackPolicyBody(v map[string]interface{}) *GetStackPolicyResponseBody {
	s.StackPolicyBody = v
	return s
}

type GetStackPolicyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStackPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetStackPolicyResponse) SetHeaders(v map[string]*string) *GetStackPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetStackPolicyResponse) SetStatusCode(v int32) *GetStackPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackPolicyResponse) SetBody(v *GetStackPolicyResponseBody) *GetStackPolicyResponse {
	s.Body = v
	return s
}

type GetStackResourceRequest struct {
	// Specifies whether to query the resource properties. Valid values:
	//
	// *   true
	// *   false
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of resource property N that you want to query.
	//
	// >  Maximum value of N: 20.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The logical ID of the resource defined in the template.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the resource. Valid values:
	//
	// *   CREATE_COMPLETE
	// *   CREATE_FAILED
	// *   CREATE_IN_PROGRESS
	// *   UPDATE_IN_PROGRESS
	// *   UPDATE_FAILED
	// *   UPDATE_COMPLETE
	// *   DELETE_IN_PROGRESS
	// *   DELETE_FAILED
	// *   CHECK_IN_PROGRESS
	// *   CHECK_FAILED
	// *   CHECK_COMPLETE
	// *   IMPORT_IN_PROGRESS
	// *   IMPORT_FAILED
	// *   IMPORT_COMPLETE
	ResourceAttributes []*string `json:"ResourceAttributes,omitempty" xml:"ResourceAttributes,omitempty" type:"Repeated"`
	// The name of resource property N that you want to query.
	ShowResourceAttributes *bool `json:"ShowResourceAttributes,omitempty" xml:"ShowResourceAttributes,omitempty"`
	// The ID of the region to which the stack belongs. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s GetStackResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStackResourceRequest) GoString() string {
	return s.String()
}

func (s *GetStackResourceRequest) SetClientToken(v string) *GetStackResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *GetStackResourceRequest) SetLogicalResourceId(v string) *GetStackResourceRequest {
	s.LogicalResourceId = &v
	return s
}

func (s *GetStackResourceRequest) SetRegionId(v string) *GetStackResourceRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackResourceRequest) SetResourceAttributes(v []*string) *GetStackResourceRequest {
	s.ResourceAttributes = v
	return s
}

func (s *GetStackResourceRequest) SetShowResourceAttributes(v bool) *GetStackResourceRequest {
	s.ShowResourceAttributes = &v
	return s
}

func (s *GetStackResourceRequest) SetStackId(v string) *GetStackResourceRequest {
	s.StackId = &v
	return s
}

type GetStackResourceResponseBody struct {
	// The resource type.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The reason why the resource is in its current state.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the stack.
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The time when the resource was updated.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The list of the resource properties.
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The information about the modules from which the resource is created. This parameter is returned only if the resource is created from modules.
	ModuleInfo *GetStackResourceResponseBodyModuleInfo `json:"ModuleInfo,omitempty" xml:"ModuleInfo,omitempty" type:"Struct"`
	// The metadata.
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The physical ID of the resource.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the resource in the last successful drift detection. Valid values:
	//
	// *   DELETED: The actual configuration of the resource differs from its expected template configuration because the resource is deleted.
	// *   MODIFIED: The actual configuration of the resource differs from its expected template configuration.
	// *   NOT_CHECKED: ROS has not checked whether the actual configuration of the resource differs from its expected template configuration.
	// *   IN_SYNC: The actual configuration of the resource matches its expected template configuration.
	ResourceAttributes []map[string]interface{} `json:"ResourceAttributes,omitempty" xml:"ResourceAttributes,omitempty" type:"Repeated"`
	// The time when the last successful drift detection was performed on the stack.
	ResourceDriftStatus *string `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty"`
	// The logical ID of the resource defined in the template.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The name of the stack.
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The ID of the request.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the resource was created.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The name of the stack.
	//
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (\_). The name must start with a digit or letter.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetStackResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetStackResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackResourceResponseBody) SetCreateTime(v string) *GetStackResourceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetStackResourceResponseBody) SetDescription(v string) *GetStackResourceResponseBody {
	s.Description = &v
	return s
}

func (s *GetStackResourceResponseBody) SetDriftDetectionTime(v string) *GetStackResourceResponseBody {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackResourceResponseBody) SetLogicalResourceId(v string) *GetStackResourceResponseBody {
	s.LogicalResourceId = &v
	return s
}

func (s *GetStackResourceResponseBody) SetMetadata(v map[string]interface{}) *GetStackResourceResponseBody {
	s.Metadata = v
	return s
}

func (s *GetStackResourceResponseBody) SetModuleInfo(v *GetStackResourceResponseBodyModuleInfo) *GetStackResourceResponseBody {
	s.ModuleInfo = v
	return s
}

func (s *GetStackResourceResponseBody) SetPhysicalResourceId(v string) *GetStackResourceResponseBody {
	s.PhysicalResourceId = &v
	return s
}

func (s *GetStackResourceResponseBody) SetRequestId(v string) *GetStackResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackResourceResponseBody) SetResourceAttributes(v []map[string]interface{}) *GetStackResourceResponseBody {
	s.ResourceAttributes = v
	return s
}

func (s *GetStackResourceResponseBody) SetResourceDriftStatus(v string) *GetStackResourceResponseBody {
	s.ResourceDriftStatus = &v
	return s
}

func (s *GetStackResourceResponseBody) SetResourceType(v string) *GetStackResourceResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetStackResourceResponseBody) SetStackId(v string) *GetStackResourceResponseBody {
	s.StackId = &v
	return s
}

func (s *GetStackResourceResponseBody) SetStackName(v string) *GetStackResourceResponseBody {
	s.StackName = &v
	return s
}

func (s *GetStackResourceResponseBody) SetStatus(v string) *GetStackResourceResponseBody {
	s.Status = &v
	return s
}

func (s *GetStackResourceResponseBody) SetStatusReason(v string) *GetStackResourceResponseBody {
	s.StatusReason = &v
	return s
}

func (s *GetStackResourceResponseBody) SetUpdateTime(v string) *GetStackResourceResponseBody {
	s.UpdateTime = &v
	return s
}

type GetStackResourceResponseBodyModuleInfo struct {
	// The concatenated logical IDs of one or more modules that contain the resource. The modules are listed from the outermost layer and separated by forward slashes (`/`).
	//
	// In the following example, the resource is created from Module B nested within Parent Module A:
	//
	// `moduleA/moduleB`
	LogicalIdHierarchy *string `json:"LogicalIdHierarchy,omitempty" xml:"LogicalIdHierarchy,omitempty"`
	// The concatenated types of one or more modules that contain the resource. The module types are listed from the outermost layer and separated by forward slashes (`/`).
	//
	// In the following example, the resource is created from a module of the `MODULE::ROS::Child::Example` type that is nested within a parent module of the `MODULE::ROS::Parent::Example` type:
	//
	// `MODULE::ROS::Parent::Example/MODULE::ROS::Child::Example`
	TypeHierarchy *string `json:"TypeHierarchy,omitempty" xml:"TypeHierarchy,omitempty"`
}

func (s GetStackResourceResponseBodyModuleInfo) String() string {
	return tea.Prettify(s)
}

func (s GetStackResourceResponseBodyModuleInfo) GoString() string {
	return s.String()
}

func (s *GetStackResourceResponseBodyModuleInfo) SetLogicalIdHierarchy(v string) *GetStackResourceResponseBodyModuleInfo {
	s.LogicalIdHierarchy = &v
	return s
}

func (s *GetStackResourceResponseBodyModuleInfo) SetTypeHierarchy(v string) *GetStackResourceResponseBodyModuleInfo {
	s.TypeHierarchy = &v
	return s
}

type GetStackResourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetStackResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetStackResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStackResourceResponse) GoString() string {
	return s.String()
}

func (s *GetStackResourceResponse) SetHeaders(v map[string]*string) *GetStackResourceResponse {
	s.Headers = v
	return s
}

func (s *GetStackResourceResponse) SetStatusCode(v int32) *GetStackResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStackResourceResponse) SetBody(v *GetStackResourceResponseBody) *GetStackResourceResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	// The ID of the change set.
	//
	// > You must specify one of the following parameters: StackId, ChangeSetId, StackGroupName, and TemplateId.
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// Specifies whether to query the shared information about the template. Valid values:
	//
	// *   Enabled
	// *   Disabled (default)
	//
	// > Only the template owner can query the shared information of a template.
	IncludePermission *string `json:"IncludePermission,omitempty" xml:"IncludePermission,omitempty"`
	// Specifies whether to query the information about tags. Valid values:
	//
	// *   Enabled
	// *   Disabled (default)
	//
	// > This parameter takes effect only if you specify TemplateId.
	IncludeTags *string `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The region ID of the stack or stack group that uses the template. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group.
	//
	// > You must specify one of the following parameters: StackId, ChangeSetId, StackGroupName, and TemplateId.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The ID of the stack.
	//
	// > You must specify one of the following parameters: StackId, ChangeSetId, StackGroupName, and TemplateId.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The ID of the template.
	//
	// This parameter applies to shared and private templates. If the template is a shared template, the value of TemplateId is the same as the value of TemplateARN. You can use the template ID to query a shared template.
	//
	// > You must specify one of the following parameters: StackId, ChangeSetId, StackGroupName, and TemplateId.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The stage of the template. This parameter takes effect only if you specify StackId, ChangeSetId, or StackGroupName.
	//
	// Valid values:
	//
	// *   Processed (default): returns the processed template.
	// *   Original: returns the original template.
	TemplateStage *string `json:"TemplateStage,omitempty" xml:"TemplateStage,omitempty"`
	// The version of the template. This parameter takes effect only if you specify TemplateId.\
	// If the template is a shared template, you can specify this parameter only if VersionOption is set to AllVersions. For more information, see [SetTemplatePermission](~~194768~~).
	//
	// Valid values: v1 to v100.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetChangeSetId(v string) *GetTemplateRequest {
	s.ChangeSetId = &v
	return s
}

func (s *GetTemplateRequest) SetIncludePermission(v string) *GetTemplateRequest {
	s.IncludePermission = &v
	return s
}

func (s *GetTemplateRequest) SetIncludeTags(v string) *GetTemplateRequest {
	s.IncludeTags = &v
	return s
}

func (s *GetTemplateRequest) SetRegionId(v string) *GetTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateRequest) SetStackGroupName(v string) *GetTemplateRequest {
	s.StackGroupName = &v
	return s
}

func (s *GetTemplateRequest) SetStackId(v string) *GetTemplateRequest {
	s.StackId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateId(v string) *GetTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateStage(v string) *GetTemplateRequest {
	s.TemplateStage = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateVersion(v string) *GetTemplateRequest {
	s.TemplateVersion = &v
	return s
}

type GetTemplateResponseBody struct {
	// The ID of the change set. This parameter is returned only if you specify ChangeSetId.
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The time when the template was created. This parameter is returned only if you specify TemplateId.
	//
	// > - If you specify TemplateVersion, the creation time of the template whose version is specified by TemplateVersion is returned.
	// > - If you do not specify TemplateVersion, the creation time of the template whose version is the default version is returned.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the template. This parameter is returned only if you specify TemplateId.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The description of the web UI in the ROS console.
	Interface *string `json:"Interface,omitempty" xml:"Interface,omitempty"`
	// The ID of the Alibaba Cloud account to which the template belongs. This parameter is returned only if you specify TemplateId.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Details of the sharing status of the template. This parameter is returned only if you specify TemplateId and set IncludePermission to Enabled.
	//
	// > - If TemplateVersion is not specified or does not take effect, the details of the sharing status of the template whose version is the default version is returned.
	// > - If TemplateVersion is specified and takes effect, the details of the sharing status of the template whose version is specified by TemplateVersion is returned.
	Permissions []*GetTemplateResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The region ID of the stack or stack group that uses the template. This parameter is returned only if you specify StackId, ChangeSetId, or StackGroupName.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The sharing type of the template. This parameter is returned only if you specify TemplateId.
	//
	// Valid values:
	//
	// *   Private: The template belongs to the template owner.
	// *   Shared: The template is shared by other users.
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The name of the stack group. This parameter is returned only if you specify StackGroupName.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The ID of the stack. This parameter is returned only if you specify StackId.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The tags of the template.
	Tags []*GetTemplateResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The Alibaba Cloud Resource Name (ARN) of the template. This parameter is returned only if you specify TemplateId.
	TemplateARN *string `json:"TemplateARN,omitempty" xml:"TemplateARN,omitempty"`
	// The content of the template.
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter is returned only if you specify TemplateId.
	//
	// If the template is a shared template, the value of this parameter is the same as the value of TemplateARN.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template. This parameter is returned only if you specify TemplateId.
	//
	// > -   If you specify TemplateVersion, the name of the template whose version is specified by TemplateVersion is returned.
	// > -  If you not specify TemplateVersion, the name of the template whose version is the default version is returned.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version of the template. This parameter is returned only if you specify TemplateId.\
	// If TemplateVersion is not specified or does not take effect, the default version is used.
	//
	// If the template is a shared template, this parameter is returned only if you set VersionOption to AllVersions.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The time when the template was last updated. This parameter is returned only if you specify TemplateId.
	//
	// > - If you specify TemplateVersion, the last update time of the template whose version is specified by TemplateVersion is returned.
	// > - If you do not specify TemplateVersion, the last update time of the template whose version is the default version is returned.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
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

type GetTemplateResponseBodyPermissions struct {
	// The ID of the Alibaba Cloud account with which the template is shared.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The sharing option.
	//
	// The value ShareToAccounts indicates that the template is shared with one or more Alibaba Cloud accounts.
	ShareOption *string `json:"ShareOption,omitempty" xml:"ShareOption,omitempty"`
	// The service that is used for resource sharing. Valid values:
	//
	// - ROS: Resources are shared from ROS by using the ROS console or calling the ROS API.
	// - ResourceDirectory: Resources are shared with accounts in a resource directory from Resource Management by using the resource sharing feature.
	// > -  The number of accounts with which resources are shared from ROS is independent of the number of accounts with which resources are shared from the resource directory.
	// > -  The shared resources from ROS cannot override or overwrite the shared resources from the resource directory.
	// > -  The shared resources from the resource directory can overwrite the shared resources from ROS.
	ShareSource *string `json:"ShareSource,omitempty" xml:"ShareSource,omitempty"`
	// The version of the shared template. This parameter is returned only if you set ShareOption to ShareToAccounts and set VersionOption to Specified or Current.
	//
	// Valid values: v1 to v100.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The version option for the shared template. This parameter is returned only if you set ShareOption to ShareToAccounts.
	//
	// Valid values:
	//
	// *   AllVersions: All template versions are shared.
	// *   Latest: Only the latest template version is shared. When the version of the template is updated, Resource Orchestration Service (ROS) updates the shared version to the latest version.
	// *   Current: Only the latest template version is shared. When the version of the template is updated, ROS does not update the shared version.
	// *   Specified: Only the specified template version is shared.
	VersionOption *string `json:"VersionOption,omitempty" xml:"VersionOption,omitempty"`
}

func (s GetTemplateResponseBodyPermissions) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyPermissions) GoString() string {
	return s.String()
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

type GetTemplateResponseBodyTags struct {
	// The tag key of the template.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the template.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTemplateResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyTags) SetKey(v string) *GetTemplateResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetTemplateResponseBodyTags) SetValue(v string) *GetTemplateResponseBodyTags {
	s.Value = &v
	return s
}

type GetTemplateResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateResponse) SetHeaders(v map[string]*string) *GetTemplateResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateResponse) SetStatusCode(v int32) *GetTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateResponse) SetBody(v *GetTemplateResponseBody) *GetTemplateResponse {
	s.Body = v
	return s
}

type GetTemplateEstimateCostRequest struct {
	// The name of parameter N. If you do not specify the name and value of a parameter, ROS uses the default name and value that are specified in the template.
	//
	// Maximum value of N: 200.
	//
	// Examples:
	//
	// *   Parameters.1.ParameterKey: `Name`
	// *   Parameters.2.ParameterKey: `Netmode`
	//
	// >  The Parameters parameter is optional. If you want to specify Parameters, you must specify both Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the scenario. The default value is the same as the value of the RegionId parameter.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	Parameters []*GetTemplateEstimateCostRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The stack ID.
	//
	// This parameter is used to query the estimated price in a configuration change scenario.
	//
	// Assume that the specified stack contains only one Elastic Compute Service (ECS) instance and the instance type is ecs.s6-c1m2.large. You downgrade the instance type to ecs.s6-c1m1.small and specify a new ApsaraDB RDS instance in the template that is used for the price inquiry. The queried result is the sum of the downgrade refund of the ECS instance and the price of the new ApsaraDB RDS instance.
	StackId      *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The value of parameter N.
	//
	// Maximum value of N: 200.
	//
	// Examples:
	//
	// *   Parameters.1.ParameterValue: `DemoEip`
	// *   Parameters.2.ParameterValue: `public`
	//
	// >  The Parameters parameter is optional. If you want to specify Parameters, you must specify both Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the scenario.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The region ID of the scenario. The default value is the same as the value of the RegionId parameter.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	TemplateScratchRegionId *string `json:"TemplateScratchRegionId,omitempty" xml:"TemplateScratchRegionId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests.
	//
	// The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// For more information, see [Ensure idempotence](~~134212~~).
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The ID of the scenario.
	//
	// For more information about how to query the IDs of scenarios, see [ListTemplateScratches](~~363050~~).
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateEstimateCostRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateEstimateCostRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateEstimateCostRequest) SetClientToken(v string) *GetTemplateEstimateCostRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetParameters(v []*GetTemplateEstimateCostRequestParameters) *GetTemplateEstimateCostRequest {
	s.Parameters = v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetRegionId(v string) *GetTemplateEstimateCostRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetStackId(v string) *GetTemplateEstimateCostRequest {
	s.StackId = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateBody(v string) *GetTemplateEstimateCostRequest {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateId(v string) *GetTemplateEstimateCostRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateScratchId(v string) *GetTemplateEstimateCostRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateScratchRegionId(v string) *GetTemplateEstimateCostRequest {
	s.TemplateScratchRegionId = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateURL(v string) *GetTemplateEstimateCostRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetTemplateEstimateCostRequest) SetTemplateVersion(v string) *GetTemplateEstimateCostRequest {
	s.TemplateVersion = &v
	return s
}

type GetTemplateEstimateCostRequestParameters struct {
	// The ID of the request.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// Details of the resource.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTemplateEstimateCostRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateEstimateCostRequestParameters) GoString() string {
	return s.String()
}

func (s *GetTemplateEstimateCostRequestParameters) SetParameterKey(v string) *GetTemplateEstimateCostRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateEstimateCostRequestParameters) SetParameterValue(v string) *GetTemplateEstimateCostRequestParameters {
	s.ParameterValue = &v
	return s
}

type GetTemplateEstimateCostResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource details.
	Resources map[string]interface{} `json:"Resources,omitempty" xml:"Resources,omitempty"`
}

func (s GetTemplateEstimateCostResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateEstimateCostResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateEstimateCostResponseBody) SetRequestId(v string) *GetTemplateEstimateCostResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateEstimateCostResponseBody) SetResources(v map[string]interface{}) *GetTemplateEstimateCostResponseBody {
	s.Resources = v
	return s
}

type GetTemplateEstimateCostResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTemplateEstimateCostResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateEstimateCostResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateEstimateCostResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateEstimateCostResponse) SetHeaders(v map[string]*string) *GetTemplateEstimateCostResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateEstimateCostResponse) SetStatusCode(v int32) *GetTemplateEstimateCostResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateEstimateCostResponse) SetBody(v *GetTemplateEstimateCostResponseBody) *GetTemplateEstimateCostResponse {
	s.Body = v
	return s
}

type GetTemplateParameterConstraintsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// For more information, see [Ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of parameter N in the template.
	Parameters []*GetTemplateParameterConstraintsRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The parameters whose values you want to query.
	ParametersKeyFilter []*string `json:"ParametersKeyFilter,omitempty" xml:"ParametersKeyFilter,omitempty" type:"Repeated"`
	// The order in which associated parameters are arranged.
	//
	// >  By default, the order of the associated parameters specified in the `Metadata` section of the template is used.
	ParametersOrder []*string `json:"ParametersOrder,omitempty" xml:"ParametersOrder,omitempty" type:"Repeated"`
	// The region ID of the template.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	StackId      *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo or oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length. If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. If you do not specify this parameter, the latest version is used.
	//
	// >  This parameter takes effect only if the TemplateId parameter is specified.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateParameterConstraintsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateParameterConstraintsRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsRequest) SetClientToken(v string) *GetTemplateParameterConstraintsRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetParameters(v []*GetTemplateParameterConstraintsRequestParameters) *GetTemplateParameterConstraintsRequest {
	s.Parameters = v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetParametersKeyFilter(v []*string) *GetTemplateParameterConstraintsRequest {
	s.ParametersKeyFilter = v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetParametersOrder(v []*string) *GetTemplateParameterConstraintsRequest {
	s.ParametersOrder = v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetRegionId(v string) *GetTemplateParameterConstraintsRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetStackId(v string) *GetTemplateParameterConstraintsRequest {
	s.StackId = &v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetTemplateBody(v string) *GetTemplateParameterConstraintsRequest {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetTemplateId(v string) *GetTemplateParameterConstraintsRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetTemplateURL(v string) *GetTemplateParameterConstraintsRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetTemplateParameterConstraintsRequest) SetTemplateVersion(v string) *GetTemplateParameterConstraintsRequest {
	s.TemplateVersion = &v
	return s
}

type GetTemplateParameterConstraintsRequestParameters struct {
	// The name of parameter N in the template.
	//
	// >  The Parameters parameter is optional. If you specify the Parameters parameter, you must specify the Parameters.N.ParameterKey parameter.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N in the template.
	//
	// >  The Parameters parameter is optional. If you specify the Parameters parameter, you must specify the Parameters.N.ParameterValue parameter.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTemplateParameterConstraintsRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateParameterConstraintsRequestParameters) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsRequestParameters) SetParameterKey(v string) *GetTemplateParameterConstraintsRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateParameterConstraintsRequestParameters) SetParameterValue(v string) *GetTemplateParameterConstraintsRequestParameters {
	s.ParameterValue = &v
	return s
}

type GetTemplateParameterConstraintsShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// For more information, see [Ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of parameter N in the template.
	Parameters []*GetTemplateParameterConstraintsShrinkRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The parameters whose values you want to query.
	ParametersKeyFilterShrink *string `json:"ParametersKeyFilter,omitempty" xml:"ParametersKeyFilter,omitempty"`
	// The order in which associated parameters are arranged.
	//
	// >  By default, the order of the associated parameters specified in the `Metadata` section of the template is used.
	ParametersOrderShrink *string `json:"ParametersOrder,omitempty" xml:"ParametersOrder,omitempty"`
	// The region ID of the template.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	StackId      *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo or oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length. If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. If you do not specify this parameter, the latest version is used.
	//
	// >  This parameter takes effect only if the TemplateId parameter is specified.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateParameterConstraintsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateParameterConstraintsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetClientToken(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetParameters(v []*GetTemplateParameterConstraintsShrinkRequestParameters) *GetTemplateParameterConstraintsShrinkRequest {
	s.Parameters = v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetParametersKeyFilterShrink(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.ParametersKeyFilterShrink = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetParametersOrderShrink(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.ParametersOrderShrink = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetRegionId(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetStackId(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.StackId = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetTemplateBody(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetTemplateId(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetTemplateURL(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequest) SetTemplateVersion(v string) *GetTemplateParameterConstraintsShrinkRequest {
	s.TemplateVersion = &v
	return s
}

type GetTemplateParameterConstraintsShrinkRequestParameters struct {
	// The name of parameter N in the template.
	//
	// >  The Parameters parameter is optional. If you specify the Parameters parameter, you must specify the Parameters.N.ParameterKey parameter.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N in the template.
	//
	// >  The Parameters parameter is optional. If you specify the Parameters parameter, you must specify the Parameters.N.ParameterValue parameter.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTemplateParameterConstraintsShrinkRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateParameterConstraintsShrinkRequestParameters) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsShrinkRequestParameters) SetParameterKey(v string) *GetTemplateParameterConstraintsShrinkRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateParameterConstraintsShrinkRequestParameters) SetParameterValue(v string) *GetTemplateParameterConstraintsShrinkRequestParameters {
	s.ParameterValue = &v
	return s
}

type GetTemplateParameterConstraintsResponseBody struct {
	// The constraints of the parameters.
	ParameterConstraints []*GetTemplateParameterConstraintsResponseBodyParameterConstraints `json:"ParameterConstraints,omitempty" xml:"ParameterConstraints,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTemplateParameterConstraintsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponseBody) SetParameterConstraints(v []*GetTemplateParameterConstraintsResponseBodyParameterConstraints) *GetTemplateParameterConstraintsResponseBody {
	s.ParameterConstraints = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBody) SetRequestId(v string) *GetTemplateParameterConstraintsResponseBody {
	s.RequestId = &v
	return s
}

type GetTemplateParameterConstraintsResponseBodyParameterConstraints struct {
	// The values of the parameter.
	AllowedValues []*string `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	// The names of the associated parameters.
	AssociationParameterNames []*string `json:"AssociationParameterNames,omitempty" xml:"AssociationParameterNames,omitempty" type:"Repeated"`
	// The behavior of the parameter. Valid values:
	//
	// *   NoLimit: No limit is imposed on the value of this parameter.
	// *   NotSupport: The value of this parameter cannot be queried.
	// *   QueryError: This parameter failed to be queried.
	//
	// > If AllowedValues is not returned, Behavior and BehaviorReason are returned.
	Behavior *string `json:"Behavior,omitempty" xml:"Behavior,omitempty"`
	// The reason why the behavior of the parameter is returned.
	BehaviorReason *string `json:"BehaviorReason,omitempty" xml:"BehaviorReason,omitempty"`
	// The values that do not conform to the parameter constraints.
	//
	// > If AllowedValues is returned, IllegalValueByParameterConstraints and IllegalValueByRules are returned at the same time.
	IllegalValueByParameterConstraints []interface{} `json:"IllegalValueByParameterConstraints,omitempty" xml:"IllegalValueByParameterConstraints,omitempty" type:"Repeated"`
	// The values that do not match the rules in the template.
	//
	// > If AllowedValues is returned, IllegalValueByParameterConstraints and IllegalValueByRules are returned at the same time.
	IllegalValueByRules []interface{} `json:"IllegalValueByRules,omitempty" xml:"IllegalValueByRules,omitempty" type:"Repeated"`
	// The unsupported resource in the template.
	NotSupportResources []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources `json:"NotSupportResources,omitempty" xml:"NotSupportResources,omitempty" type:"Repeated"`
	// The original constraint information.
	OriginalConstraints []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints `json:"OriginalConstraints,omitempty" xml:"OriginalConstraints,omitempty" type:"Repeated"`
	// The name of the parameter.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The error that is returned when the request fails.
	QueryErrors []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors `json:"QueryErrors,omitempty" xml:"QueryErrors,omitempty" type:"Repeated"`
	// The data type of the parameter.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraints) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraints) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetAllowedValues(v []*string) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetAssociationParameterNames(v []*string) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.AssociationParameterNames = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehavior(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Behavior = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetBehaviorReason(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.BehaviorReason = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetIllegalValueByParameterConstraints(v []interface{}) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.IllegalValueByParameterConstraints = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetIllegalValueByRules(v []interface{}) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.IllegalValueByRules = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetNotSupportResources(v []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.NotSupportResources = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetOriginalConstraints(v []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.OriginalConstraints = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetParameterKey(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetQueryErrors(v []*GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.QueryErrors = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraints) SetType(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraints {
	s.Type = &v
	return s
}

type GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources struct {
	// The name of the resource property.
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources) SetPropertyName(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources {
	s.PropertyName = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources) SetResourceType(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsNotSupportResources {
	s.ResourceType = &v
	return s
}

type GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints struct {
	// The values of the parameter.
	AllowedValues []interface{} `json:"AllowedValues,omitempty" xml:"AllowedValues,omitempty" type:"Repeated"`
	// The name of the resource property.
	PropertyName *string `json:"PropertyName,omitempty" xml:"PropertyName,omitempty"`
	// The name of the resource that is defined in the template.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetAllowedValues(v []interface{}) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.AllowedValues = v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetPropertyName(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.PropertyName = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceName(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceName = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints) SetResourceType(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsOriginalConstraints {
	s.ResourceType = &v
	return s
}

type GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors struct {
	// The error message.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The resource name.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) SetErrorMessage(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors {
	s.ErrorMessage = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) SetResourceName(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors {
	s.ResourceName = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors) SetResourceType(v string) *GetTemplateParameterConstraintsResponseBodyParameterConstraintsQueryErrors {
	s.ResourceType = &v
	return s
}

type GetTemplateParameterConstraintsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTemplateParameterConstraintsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateParameterConstraintsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateParameterConstraintsResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateParameterConstraintsResponse) SetHeaders(v map[string]*string) *GetTemplateParameterConstraintsResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateParameterConstraintsResponse) SetStatusCode(v int32) *GetTemplateParameterConstraintsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateParameterConstraintsResponse) SetBody(v *GetTemplateParameterConstraintsResponseBody) *GetTemplateParameterConstraintsResponse {
	s.Body = v
	return s
}

type GetTemplateRecommendParametersRequest struct {
	ClientToken     *string                                            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Parameters      []*GetTemplateRecommendParametersRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	RegionId        *string                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateBody    *string                                            `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	TemplateId      *string                                            `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateURL     *string                                            `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	TemplateVersion *string                                            `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateRecommendParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRecommendParametersRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRecommendParametersRequest) SetClientToken(v string) *GetTemplateRecommendParametersRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTemplateRecommendParametersRequest) SetParameters(v []*GetTemplateRecommendParametersRequestParameters) *GetTemplateRecommendParametersRequest {
	s.Parameters = v
	return s
}

func (s *GetTemplateRecommendParametersRequest) SetRegionId(v string) *GetTemplateRecommendParametersRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateRecommendParametersRequest) SetTemplateBody(v string) *GetTemplateRecommendParametersRequest {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateRecommendParametersRequest) SetTemplateId(v string) *GetTemplateRecommendParametersRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateRecommendParametersRequest) SetTemplateURL(v string) *GetTemplateRecommendParametersRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetTemplateRecommendParametersRequest) SetTemplateVersion(v string) *GetTemplateRecommendParametersRequest {
	s.TemplateVersion = &v
	return s
}

type GetTemplateRecommendParametersRequestParameters struct {
	ParameterCandidateValues []*string `json:"ParameterCandidateValues,omitempty" xml:"ParameterCandidateValues,omitempty" type:"Repeated"`
	ParameterKey             *string   `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue           *string   `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTemplateRecommendParametersRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRecommendParametersRequestParameters) GoString() string {
	return s.String()
}

func (s *GetTemplateRecommendParametersRequestParameters) SetParameterCandidateValues(v []*string) *GetTemplateRecommendParametersRequestParameters {
	s.ParameterCandidateValues = v
	return s
}

func (s *GetTemplateRecommendParametersRequestParameters) SetParameterKey(v string) *GetTemplateRecommendParametersRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateRecommendParametersRequestParameters) SetParameterValue(v string) *GetTemplateRecommendParametersRequestParameters {
	s.ParameterValue = &v
	return s
}

type GetTemplateRecommendParametersResponseBody struct {
	RecommendParameterValues []*GetTemplateRecommendParametersResponseBodyRecommendParameterValues `json:"RecommendParameterValues,omitempty" xml:"RecommendParameterValues,omitempty" type:"Repeated"`
	RequestId                *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTemplateRecommendParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRecommendParametersResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateRecommendParametersResponseBody) SetRecommendParameterValues(v []*GetTemplateRecommendParametersResponseBodyRecommendParameterValues) *GetTemplateRecommendParametersResponseBody {
	s.RecommendParameterValues = v
	return s
}

func (s *GetTemplateRecommendParametersResponseBody) SetRequestId(v string) *GetTemplateRecommendParametersResponseBody {
	s.RequestId = &v
	return s
}

type GetTemplateRecommendParametersResponseBodyRecommendParameterValues struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	RecommendValue *string `json:"RecommendValue,omitempty" xml:"RecommendValue,omitempty"`
}

func (s GetTemplateRecommendParametersResponseBodyRecommendParameterValues) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRecommendParametersResponseBodyRecommendParameterValues) GoString() string {
	return s.String()
}

func (s *GetTemplateRecommendParametersResponseBodyRecommendParameterValues) SetParameterKey(v string) *GetTemplateRecommendParametersResponseBodyRecommendParameterValues {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateRecommendParametersResponseBodyRecommendParameterValues) SetRecommendValue(v string) *GetTemplateRecommendParametersResponseBodyRecommendParameterValues {
	s.RecommendValue = &v
	return s
}

type GetTemplateRecommendParametersResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTemplateRecommendParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateRecommendParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRecommendParametersResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateRecommendParametersResponse) SetHeaders(v map[string]*string) *GetTemplateRecommendParametersResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateRecommendParametersResponse) SetStatusCode(v int32) *GetTemplateRecommendParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateRecommendParametersResponse) SetBody(v *GetTemplateRecommendParametersResponseBody) *GetTemplateRecommendParametersResponse {
	s.Body = v
	return s
}

type GetTemplateScratchRequest struct {
	// The region ID of the scenario.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The data display option. Valid values:
	//
	// *   Sources: displays only the data of source nodes. This setting takes effect only when TemplateScratchType is set to ArchitectureDetection.
	// *   Source: displays only the data of the source node. This setting takes effect only when TemplateScratchType is not set to ArchitectureDetection.
	// *   Provisions: displays only the data of new nodes. This setting takes effect only when TemplateScratchType is not set to ArchitectureDetection.
	// *   All: displays all data.
	//
	// For more information about source nodes and new nodes, see [Overview](~~352074~~).
	//
	// >  If you do not specify this parameter, the node data is not displayed.
	ShowDataOption *string `json:"ShowDataOption,omitempty" xml:"ShowDataOption,omitempty"`
	// The ID of the scenario.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
}

func (s GetTemplateScratchRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateScratchRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchRequest) SetRegionId(v string) *GetTemplateScratchRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateScratchRequest) SetShowDataOption(v string) *GetTemplateScratchRequest {
	s.ShowDataOption = &v
	return s
}

func (s *GetTemplateScratchRequest) SetTemplateScratchId(v string) *GetTemplateScratchRequest {
	s.TemplateScratchId = &v
	return s
}

type GetTemplateScratchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource scenario.
	TemplateScratch *GetTemplateScratchResponseBodyTemplateScratch `json:"TemplateScratch,omitempty" xml:"TemplateScratch,omitempty" type:"Struct"`
}

func (s GetTemplateScratchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateScratchResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBody) SetRequestId(v string) *GetTemplateScratchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateScratchResponseBody) SetTemplateScratch(v *GetTemplateScratchResponseBodyTemplateScratch) *GetTemplateScratchResponseBody {
	s.TemplateScratch = v
	return s
}

type GetTemplateScratchResponseBodyTemplateScratch struct {
	// The time at which the scenario was created.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the scenario.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The status code of the scenario that fails to be created.
	//
	// > This parameter is returned only if you set Status to GENERATE_FAILED.
	FailedCode *string `json:"FailedCode,omitempty" xml:"FailedCode,omitempty"`
	// The policy based on which the logical ID is generated. Valid values:
	//
	// *   LongTypePrefixAndIndexSuffix (default): long-type prefix + index-type suffix
	// *   LongTypePrefixAndHashSuffix: long-type prefix + hash-type suffix
	// *   ShortTypePrefixAndHashSuffix: short-type prefix + hash-type suffix
	LogicalIdStrategy *string `json:"LogicalIdStrategy,omitempty" xml:"LogicalIdStrategy,omitempty"`
	// The preference parameters of the scenario.
	PreferenceParameters []*GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters `json:"PreferenceParameters,omitempty" xml:"PreferenceParameters,omitempty" type:"Repeated"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The source resource group.
	SourceResourceGroup *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup `json:"SourceResourceGroup,omitempty" xml:"SourceResourceGroup,omitempty" type:"Struct"`
	// The source resources.
	SourceResources []*GetTemplateScratchResponseBodyTemplateScratchSourceResources `json:"SourceResources,omitempty" xml:"SourceResources,omitempty" type:"Repeated"`
	// The source tag.
	SourceTag *GetTemplateScratchResponseBodyTemplateScratchSourceTag `json:"SourceTag,omitempty" xml:"SourceTag,omitempty" type:"Struct"`
	// The preset information of the stack.
	StackProvision *GetTemplateScratchResponseBodyTemplateScratchStackProvision `json:"StackProvision,omitempty" xml:"StackProvision,omitempty" type:"Struct"`
	// The stacks that are associated with the scenario.
	Stacks []*GetTemplateScratchResponseBodyTemplateScratchStacks `json:"Stacks,omitempty" xml:"Stacks,omitempty" type:"Repeated"`
	// The state of the scenario. Valid values:
	//
	// *   GENERATE_IN_PROGRESS: The scenario is being created.
	// *   GENERATE_COMPLETE: The scenario is created.
	// *   GENERATE_FAILED: The scenario fails to be created.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the scenario fails to be created.
	//
	// > This parameter is returned only if you set Status to GENERATE_FAILED.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The scenario data.
	TemplateScratchData map[string]interface{} `json:"TemplateScratchData,omitempty" xml:"TemplateScratchData,omitempty"`
	// The ID of the scenario.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The type of the scenario. Valid values:
	//
	// *   ResourceImport: resource management
	// *   ArchitectureReplication: resource replication
	TemplateScratchType *string `json:"TemplateScratchType,omitempty" xml:"TemplateScratchType,omitempty"`
	// The time at which the scenario was updated.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetTemplateScratchResponseBodyTemplateScratch) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratch) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetCreateTime(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.CreateTime = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetDescription(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.Description = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetFailedCode(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.FailedCode = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetLogicalIdStrategy(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.LogicalIdStrategy = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetPreferenceParameters(v []*GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters) *GetTemplateScratchResponseBodyTemplateScratch {
	s.PreferenceParameters = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetResourceGroupId(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetSourceResourceGroup(v *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup) *GetTemplateScratchResponseBodyTemplateScratch {
	s.SourceResourceGroup = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetSourceResources(v []*GetTemplateScratchResponseBodyTemplateScratchSourceResources) *GetTemplateScratchResponseBodyTemplateScratch {
	s.SourceResources = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetSourceTag(v *GetTemplateScratchResponseBodyTemplateScratchSourceTag) *GetTemplateScratchResponseBodyTemplateScratch {
	s.SourceTag = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetStackProvision(v *GetTemplateScratchResponseBodyTemplateScratchStackProvision) *GetTemplateScratchResponseBodyTemplateScratch {
	s.StackProvision = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetStacks(v []*GetTemplateScratchResponseBodyTemplateScratchStacks) *GetTemplateScratchResponseBodyTemplateScratch {
	s.Stacks = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetStatus(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.Status = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetStatusReason(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.StatusReason = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetTemplateScratchData(v map[string]interface{}) *GetTemplateScratchResponseBodyTemplateScratch {
	s.TemplateScratchData = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetTemplateScratchId(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.TemplateScratchId = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetTemplateScratchType(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.TemplateScratchType = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratch) SetUpdateTime(v string) *GetTemplateScratchResponseBodyTemplateScratch {
	s.UpdateTime = &v
	return s
}

type GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters struct {
	// The parameter name.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The parameter value.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters) SetParameterKey(v string) *GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters) SetParameterValue(v string) *GetTemplateScratchResponseBodyTemplateScratchPreferenceParameters {
	s.ParameterValue = &v
	return s
}

type GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup struct {
	// The ID of the source resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource types.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup) SetResourceGroupId(v string) *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup) SetResourceTypeFilter(v []*string) *GetTemplateScratchResponseBodyTemplateScratchSourceResourceGroup {
	s.ResourceTypeFilter = v
	return s
}

type GetTemplateScratchResponseBodyTemplateScratchSourceResources struct {
	// The resource ID.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetTemplateScratchResponseBodyTemplateScratchSourceResources) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratchSourceResources) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResources) SetResourceId(v string) *GetTemplateScratchResponseBodyTemplateScratchSourceResources {
	s.ResourceId = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceResources) SetResourceType(v string) *GetTemplateScratchResponseBodyTemplateScratchSourceResources {
	s.ResourceType = &v
	return s
}

type GetTemplateScratchResponseBodyTemplateScratchSourceTag struct {
	// The source tags.
	ResourceTags map[string]interface{} `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty"`
	// The resource types.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s GetTemplateScratchResponseBodyTemplateScratchSourceTag) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratchSourceTag) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceTag) SetResourceTags(v map[string]interface{}) *GetTemplateScratchResponseBodyTemplateScratchSourceTag {
	s.ResourceTags = v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchSourceTag) SetResourceTypeFilter(v []*string) *GetTemplateScratchResponseBodyTemplateScratchSourceTag {
	s.ResourceTypeFilter = v
	return s
}

type GetTemplateScratchResponseBodyTemplateScratchStackProvision struct {
	// Indicates whether the resource is replicated by calling the [CreateStack](~~132086~~) operation. Valid values:
	//
	// *   true
	// *   false
	Creatable *bool `json:"Creatable,omitempty" xml:"Creatable,omitempty"`
	// Indicates whether the resource is managed by calling the [CreateChangeSet](~~131051~~) operation. Valid values:
	//
	// *   true
	// *   false
	Importable *bool `json:"Importable,omitempty" xml:"Importable,omitempty"`
}

func (s GetTemplateScratchResponseBodyTemplateScratchStackProvision) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratchStackProvision) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStackProvision) SetCreatable(v bool) *GetTemplateScratchResponseBodyTemplateScratchStackProvision {
	s.Creatable = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStackProvision) SetImportable(v bool) *GetTemplateScratchResponseBodyTemplateScratchStackProvision {
	s.Importable = &v
	return s
}

type GetTemplateScratchResponseBodyTemplateScratchStacks struct {
	// The region ID of the stack.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The stack ID.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The purpose of the stack. Valid values:
	//
	// *   ResourceImport: resource management
	// *   ArchitectureReplication: resource replication
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
}

func (s GetTemplateScratchResponseBodyTemplateScratchStacks) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateScratchResponseBodyTemplateScratchStacks) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStacks) SetRegionId(v string) *GetTemplateScratchResponseBodyTemplateScratchStacks {
	s.RegionId = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStacks) SetStackId(v string) *GetTemplateScratchResponseBodyTemplateScratchStacks {
	s.StackId = &v
	return s
}

func (s *GetTemplateScratchResponseBodyTemplateScratchStacks) SetUsageType(v string) *GetTemplateScratchResponseBodyTemplateScratchStacks {
	s.UsageType = &v
	return s
}

type GetTemplateScratchResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTemplateScratchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateScratchResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateScratchResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateScratchResponse) SetHeaders(v map[string]*string) *GetTemplateScratchResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateScratchResponse) SetStatusCode(v int32) *GetTemplateScratchResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateScratchResponse) SetBody(v *GetTemplateScratchResponseBody) *GetTemplateScratchResponse {
	s.Body = v
	return s
}

type GetTemplateSummaryRequest struct {
	// The ID of the change set.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, StackId, ChangeSetId, and StackGroupName.
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\
	// The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (\_).\
	// For more information, see [Ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The parameters that are defined in the template.
	Parameters []*GetTemplateSummaryRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region ID of the stack or stack group that uses the template. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	//
	// This parameter takes effect only when one of the following parameters are specified: StackId, ChangeSetId, and StackGroupName.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, StackId, ChangeSetId, and StackGroupName.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The stack ID.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, StackId, ChangeSetId, and StackGroupName.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.\
	// If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.\
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, StackId, ChangeSetId, and StackGroupName.
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The template ID. This parameter applies to shared and private templates.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, StackId, ChangeSetId, and StackGroupName.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length.
	//
	// > If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// You can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, StackId, ChangeSetId, and StackGroupName.
	//
	// The URL can be up to 1,024 bytes in length.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. This parameter takes effect when TemplateId is specified.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryRequest) SetChangeSetId(v string) *GetTemplateSummaryRequest {
	s.ChangeSetId = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetClientToken(v string) *GetTemplateSummaryRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetParameters(v []*GetTemplateSummaryRequestParameters) *GetTemplateSummaryRequest {
	s.Parameters = v
	return s
}

func (s *GetTemplateSummaryRequest) SetRegionId(v string) *GetTemplateSummaryRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetStackGroupName(v string) *GetTemplateSummaryRequest {
	s.StackGroupName = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetStackId(v string) *GetTemplateSummaryRequest {
	s.StackId = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetTemplateBody(v string) *GetTemplateSummaryRequest {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetTemplateId(v string) *GetTemplateSummaryRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetTemplateURL(v string) *GetTemplateSummaryRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetTemplateSummaryRequest) SetTemplateVersion(v string) *GetTemplateSummaryRequest {
	s.TemplateVersion = &v
	return s
}

type GetTemplateSummaryRequestParameters struct {
	// The name of parameter N that is defined in the template. If you do not specify the name and value of a parameter, Resource Orchestration Service (ROS) uses the default name and value that are defined in the template.
	//
	// Maximum value of N: 200.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify both Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N that is defined in the template.
	//
	// Maximum value of N: 200.
	//
	// > The Parameters parameter is optional. If you specify Parameters, you must specify both Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTemplateSummaryRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateSummaryRequestParameters) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryRequestParameters) SetParameterKey(v string) *GetTemplateSummaryRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateSummaryRequestParameters) SetParameterValue(v string) *GetTemplateSummaryRequestParameters {
	s.ParameterValue = &v
	return s
}

type GetTemplateSummaryResponseBody struct {
	// The description of the stack template.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The metadata that is defined in the template.
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The declarations of the parameters in the template.
	Parameters []map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource identifier summaries.\
	// A summary describes the resource that you want to import and the properties that are used to identify the resource during the import. For example, VpcId is an identifier property of ALIYUN::ECS::VPC.
	ResourceIdentifierSummaries []*GetTemplateSummaryResponseBodyResourceIdentifierSummaries `json:"ResourceIdentifierSummaries,omitempty" xml:"ResourceIdentifierSummaries,omitempty" type:"Repeated"`
	// All resource types that are used in the template.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
	// The version of the template.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetTemplateSummaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryResponseBody) SetDescription(v string) *GetTemplateSummaryResponseBody {
	s.Description = &v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetMetadata(v map[string]interface{}) *GetTemplateSummaryResponseBody {
	s.Metadata = v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetParameters(v []map[string]interface{}) *GetTemplateSummaryResponseBody {
	s.Parameters = v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetRequestId(v string) *GetTemplateSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetResourceIdentifierSummaries(v []*GetTemplateSummaryResponseBodyResourceIdentifierSummaries) *GetTemplateSummaryResponseBody {
	s.ResourceIdentifierSummaries = v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetResourceTypes(v []*string) *GetTemplateSummaryResponseBody {
	s.ResourceTypes = v
	return s
}

func (s *GetTemplateSummaryResponseBody) SetVersion(v string) *GetTemplateSummaryResponseBody {
	s.Version = &v
	return s
}

type GetTemplateSummaryResponseBodyResourceIdentifierSummaries struct {
	// The logical IDs of all resources of the type that is specified by ResouceType in the template.
	LogicalResourceIds []*string `json:"LogicalResourceIds,omitempty" xml:"LogicalResourceIds,omitempty" type:"Repeated"`
	// The resource properties. You can use a resource property to identify the resource that you want to manage. For example, VpcId is an identifier property of ALIYUN::ECS::VPC.
	ResourceIdentifiers []*string `json:"ResourceIdentifiers,omitempty" xml:"ResourceIdentifiers,omitempty" type:"Repeated"`
	// The resource type.
	//
	// > The resource import feature is supported for the resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetTemplateSummaryResponseBodyResourceIdentifierSummaries) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateSummaryResponseBodyResourceIdentifierSummaries) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryResponseBodyResourceIdentifierSummaries) SetLogicalResourceIds(v []*string) *GetTemplateSummaryResponseBodyResourceIdentifierSummaries {
	s.LogicalResourceIds = v
	return s
}

func (s *GetTemplateSummaryResponseBodyResourceIdentifierSummaries) SetResourceIdentifiers(v []*string) *GetTemplateSummaryResponseBodyResourceIdentifierSummaries {
	s.ResourceIdentifiers = v
	return s
}

func (s *GetTemplateSummaryResponseBodyResourceIdentifierSummaries) SetResourceType(v string) *GetTemplateSummaryResponseBodyResourceIdentifierSummaries {
	s.ResourceType = &v
	return s
}

type GetTemplateSummaryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTemplateSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTemplateSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetTemplateSummaryResponse) SetHeaders(v map[string]*string) *GetTemplateSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetTemplateSummaryResponse) SetStatusCode(v int32) *GetTemplateSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTemplateSummaryResponse) SetBody(v *GetTemplateSummaryResponseBody) *GetTemplateSummaryResponse {
	s.Body = v
	return s
}

type ListChangeSetsRequest struct {
	// The ID of the change set. If detailed information about the change set is not required, you can use this parameter to replace the GetChangeSet operation.
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The name of change set N. Maximum value of N: 5. You can use an asterisk (\*) as a wildcard for fuzzy search.
	ChangeSetName []*string `json:"ChangeSetName,omitempty" xml:"ChangeSetName,omitempty" type:"Repeated"`
	// The execution status of change set N. Maximum value of N: 5. Valid values:
	//
	// *   UNAVAILABLE
	// *   AVAILABLE
	// *   EXECUTE_IN_PROGRESS
	// *   EXECUTE_COMPLETE
	// *   EXECUTE_FAILED
	// *   OBSOLETE
	ExecutionStatus []*string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty" type:"Repeated"`
	// The page number.\
	// Pages start from page 1.\
	// Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\
	// Valid values: 1 to 50.\
	// Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the change set. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The status of change set N. Maximum value of N: 5. Valid values:
	//
	// *   CREATE_PENDING
	// *   CREATE_IN_PROGRESS
	// *   CREATE_COMPLETE
	// *   CREATE_FAILED
	// *   DELETE_FAILED
	// *   DELETE_COMPLETE
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s ListChangeSetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListChangeSetsRequest) GoString() string {
	return s.String()
}

func (s *ListChangeSetsRequest) SetChangeSetId(v string) *ListChangeSetsRequest {
	s.ChangeSetId = &v
	return s
}

func (s *ListChangeSetsRequest) SetChangeSetName(v []*string) *ListChangeSetsRequest {
	s.ChangeSetName = v
	return s
}

func (s *ListChangeSetsRequest) SetExecutionStatus(v []*string) *ListChangeSetsRequest {
	s.ExecutionStatus = v
	return s
}

func (s *ListChangeSetsRequest) SetPageNumber(v int64) *ListChangeSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListChangeSetsRequest) SetPageSize(v int64) *ListChangeSetsRequest {
	s.PageSize = &v
	return s
}

func (s *ListChangeSetsRequest) SetRegionId(v string) *ListChangeSetsRequest {
	s.RegionId = &v
	return s
}

func (s *ListChangeSetsRequest) SetStackId(v string) *ListChangeSetsRequest {
	s.StackId = &v
	return s
}

func (s *ListChangeSetsRequest) SetStatus(v []*string) *ListChangeSetsRequest {
	s.Status = v
	return s
}

type ListChangeSetsResponseBody struct {
	// The change sets.
	ChangeSets []*ListChangeSetsResponseBodyChangeSets `json:"ChangeSets,omitempty" xml:"ChangeSets,omitempty" type:"Repeated"`
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of change sets returned.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListChangeSetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListChangeSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChangeSetsResponseBody) SetChangeSets(v []*ListChangeSetsResponseBodyChangeSets) *ListChangeSetsResponseBody {
	s.ChangeSets = v
	return s
}

func (s *ListChangeSetsResponseBody) SetPageNumber(v int32) *ListChangeSetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListChangeSetsResponseBody) SetPageSize(v int32) *ListChangeSetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChangeSetsResponseBody) SetRequestId(v string) *ListChangeSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChangeSetsResponseBody) SetTotalCount(v int32) *ListChangeSetsResponseBody {
	s.TotalCount = &v
	return s
}

type ListChangeSetsResponseBodyChangeSets struct {
	// The ID of the change set.
	ChangeSetId *string `json:"ChangeSetId,omitempty" xml:"ChangeSetId,omitempty"`
	// The name of the change set.
	ChangeSetName *string `json:"ChangeSetName,omitempty" xml:"ChangeSetName,omitempty"`
	// The type of the change set.
	ChangeSetType *string `json:"ChangeSetType,omitempty" xml:"ChangeSetType,omitempty"`
	// The time when the change set was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the change set.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution status of the change set.
	ExecutionStatus *string `json:"ExecutionStatus,omitempty" xml:"ExecutionStatus,omitempty"`
	// The region ID of the change set.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack with which the change set is associated.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The name of the stack with which the change set is associated.
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The status of the change set.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the change set is in its current state.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s ListChangeSetsResponseBodyChangeSets) String() string {
	return tea.Prettify(s)
}

func (s ListChangeSetsResponseBodyChangeSets) GoString() string {
	return s.String()
}

func (s *ListChangeSetsResponseBodyChangeSets) SetChangeSetId(v string) *ListChangeSetsResponseBodyChangeSets {
	s.ChangeSetId = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetChangeSetName(v string) *ListChangeSetsResponseBodyChangeSets {
	s.ChangeSetName = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetChangeSetType(v string) *ListChangeSetsResponseBodyChangeSets {
	s.ChangeSetType = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetCreateTime(v string) *ListChangeSetsResponseBodyChangeSets {
	s.CreateTime = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetDescription(v string) *ListChangeSetsResponseBodyChangeSets {
	s.Description = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetExecutionStatus(v string) *ListChangeSetsResponseBodyChangeSets {
	s.ExecutionStatus = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetRegionId(v string) *ListChangeSetsResponseBodyChangeSets {
	s.RegionId = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetStackId(v string) *ListChangeSetsResponseBodyChangeSets {
	s.StackId = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetStackName(v string) *ListChangeSetsResponseBodyChangeSets {
	s.StackName = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetStatus(v string) *ListChangeSetsResponseBodyChangeSets {
	s.Status = &v
	return s
}

func (s *ListChangeSetsResponseBodyChangeSets) SetStatusReason(v string) *ListChangeSetsResponseBodyChangeSets {
	s.StatusReason = &v
	return s
}

type ListChangeSetsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListChangeSetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListChangeSetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListChangeSetsResponse) GoString() string {
	return s.String()
}

func (s *ListChangeSetsResponse) SetHeaders(v map[string]*string) *ListChangeSetsResponse {
	s.Headers = v
	return s
}

func (s *ListChangeSetsResponse) SetStatusCode(v int32) *ListChangeSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListChangeSetsResponse) SetBody(v *ListChangeSetsResponseBody) *ListChangeSetsResponse {
	s.Body = v
	return s
}

type ListDiagnosticsRequest struct {
	// The keyword in the diagnosis.
	DiagnosticKey *string `json:"DiagnosticKey,omitempty" xml:"DiagnosticKey,omitempty"`
	// The product that is diagnosed.
	DiagnosticProduct *string `json:"DiagnosticProduct,omitempty" xml:"DiagnosticProduct,omitempty"`
	// The maximum number of results to be returned in a single call when NextToken is used for the query.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 50.
	MaxResults *string `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The diagnosis status.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDiagnosticsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosticsRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnosticsRequest) SetDiagnosticKey(v string) *ListDiagnosticsRequest {
	s.DiagnosticKey = &v
	return s
}

func (s *ListDiagnosticsRequest) SetDiagnosticProduct(v string) *ListDiagnosticsRequest {
	s.DiagnosticProduct = &v
	return s
}

func (s *ListDiagnosticsRequest) SetMaxResults(v string) *ListDiagnosticsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDiagnosticsRequest) SetNextToken(v string) *ListDiagnosticsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDiagnosticsRequest) SetStatus(v string) *ListDiagnosticsRequest {
	s.Status = &v
	return s
}

type ListDiagnosticsResponseBody struct {
	// The items that are diagnosed.
	Diagnostics []*ListDiagnosticsResponseBodyDiagnostics `json:"Diagnostics,omitempty" xml:"Diagnostics,omitempty" type:"Repeated"`
	// The HTTP status code returned. The value 200 indicates that the request was successful.
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// *   true
	// *   false
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDiagnosticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnosticsResponseBody) SetDiagnostics(v []*ListDiagnosticsResponseBodyDiagnostics) *ListDiagnosticsResponseBody {
	s.Diagnostics = v
	return s
}

func (s *ListDiagnosticsResponseBody) SetHttpStatusCode(v int32) *ListDiagnosticsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDiagnosticsResponseBody) SetMessage(v string) *ListDiagnosticsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDiagnosticsResponseBody) SetNextToken(v string) *ListDiagnosticsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDiagnosticsResponseBody) SetRequestId(v string) *ListDiagnosticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDiagnosticsResponseBody) SetSuccess(v string) *ListDiagnosticsResponseBody {
	s.Success = &v
	return s
}

type ListDiagnosticsResponseBodyDiagnostics struct {
	// The time when the diagnostic report was generated.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The keyword in the diagnosis.
	DiagnosticKey *string `json:"DiagnosticKey,omitempty" xml:"DiagnosticKey,omitempty"`
	// The product that is diagnosed.
	DiagnosticProduct *string `json:"DiagnosticProduct,omitempty" xml:"DiagnosticProduct,omitempty"`
	// The ID of the diagnostic report.
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The diagnosis status.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDiagnosticsResponseBodyDiagnostics) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosticsResponseBodyDiagnostics) GoString() string {
	return s.String()
}

func (s *ListDiagnosticsResponseBodyDiagnostics) SetCreateTime(v string) *ListDiagnosticsResponseBodyDiagnostics {
	s.CreateTime = &v
	return s
}

func (s *ListDiagnosticsResponseBodyDiagnostics) SetDiagnosticKey(v string) *ListDiagnosticsResponseBodyDiagnostics {
	s.DiagnosticKey = &v
	return s
}

func (s *ListDiagnosticsResponseBodyDiagnostics) SetDiagnosticProduct(v string) *ListDiagnosticsResponseBodyDiagnostics {
	s.DiagnosticProduct = &v
	return s
}

func (s *ListDiagnosticsResponseBodyDiagnostics) SetReportId(v string) *ListDiagnosticsResponseBodyDiagnostics {
	s.ReportId = &v
	return s
}

func (s *ListDiagnosticsResponseBodyDiagnostics) SetStatus(v string) *ListDiagnosticsResponseBodyDiagnostics {
	s.Status = &v
	return s
}

type ListDiagnosticsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDiagnosticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDiagnosticsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosticsResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnosticsResponse) SetHeaders(v map[string]*string) *ListDiagnosticsResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnosticsResponse) SetStatusCode(v int32) *ListDiagnosticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiagnosticsResponse) SetBody(v *ListDiagnosticsResponseBody) *ListDiagnosticsResponse {
	s.Body = v
	return s
}

type ListResourceTypeRegistrationsRequest struct {
	// The entity type. Set the value to Module.
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the registration record.
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The resource type. The resource type can contain letters, digits, colons (:), and asterisks (\*). You can use an asterisk (\*) to perform a fuzzy match.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The registration state. Valid values:
	//
	// *   IN_PROGRESS
	// *   COMPLETE
	// *   FAILED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListResourceTypeRegistrationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypeRegistrationsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypeRegistrationsRequest) SetEntityType(v string) *ListResourceTypeRegistrationsRequest {
	s.EntityType = &v
	return s
}

func (s *ListResourceTypeRegistrationsRequest) SetPageNumber(v int32) *ListResourceTypeRegistrationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceTypeRegistrationsRequest) SetPageSize(v int32) *ListResourceTypeRegistrationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceTypeRegistrationsRequest) SetRegistrationId(v string) *ListResourceTypeRegistrationsRequest {
	s.RegistrationId = &v
	return s
}

func (s *ListResourceTypeRegistrationsRequest) SetResourceType(v string) *ListResourceTypeRegistrationsRequest {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypeRegistrationsRequest) SetStatus(v string) *ListResourceTypeRegistrationsRequest {
	s.Status = &v
	return s
}

type ListResourceTypeRegistrationsResponseBody struct {
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The registration records.
	Registrations []*ListResourceTypeRegistrationsResponseBodyRegistrations `json:"Registrations,omitempty" xml:"Registrations,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of registration records.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceTypeRegistrationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypeRegistrationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypeRegistrationsResponseBody) SetPageNumber(v int32) *ListResourceTypeRegistrationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBody) SetRegistrations(v []*ListResourceTypeRegistrationsResponseBodyRegistrations) *ListResourceTypeRegistrationsResponseBody {
	s.Registrations = v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBody) SetRequestId(v string) *ListResourceTypeRegistrationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBody) SetTotalCount(v int32) *ListResourceTypeRegistrationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceTypeRegistrationsResponseBodyRegistrations struct {
	// The time when the version was created. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The entity type. Only Module may be returned.
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The ID of the registration record.
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The registration state. Valid values:
	//
	// *   IN_PROGRESS
	// *   COMPLETE
	// *   FAILED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason for the state.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The version ID.
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListResourceTypeRegistrationsResponseBodyRegistrations) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypeRegistrationsResponseBodyRegistrations) GoString() string {
	return s.String()
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetCreateTime(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.CreateTime = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetEntityType(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.EntityType = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetRegistrationId(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.RegistrationId = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetResourceType(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetStatus(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.Status = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetStatusReason(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.StatusReason = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponseBodyRegistrations) SetVersionId(v string) *ListResourceTypeRegistrationsResponseBodyRegistrations {
	s.VersionId = &v
	return s
}

type ListResourceTypeRegistrationsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceTypeRegistrationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceTypeRegistrationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypeRegistrationsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTypeRegistrationsResponse) SetHeaders(v map[string]*string) *ListResourceTypeRegistrationsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTypeRegistrationsResponse) SetStatusCode(v int32) *ListResourceTypeRegistrationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceTypeRegistrationsResponse) SetBody(v *ListResourceTypeRegistrationsResponseBody) *ListResourceTypeRegistrationsResponse {
	s.Body = v
	return s
}

type ListResourceTypeVersionsRequest struct {
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceTypeVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypeVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypeVersionsRequest) SetResourceType(v string) *ListResourceTypeVersionsRequest {
	s.ResourceType = &v
	return s
}

type ListResourceTypeVersionsResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The versions of the resource types.
	ResourceTypeVersions []*ListResourceTypeVersionsResponseBodyResourceTypeVersions `json:"ResourceTypeVersions,omitempty" xml:"ResourceTypeVersions,omitempty" type:"Repeated"`
}

func (s ListResourceTypeVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypeVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypeVersionsResponseBody) SetRequestId(v string) *ListResourceTypeVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBody) SetResourceTypeVersions(v []*ListResourceTypeVersionsResponseBodyResourceTypeVersions) *ListResourceTypeVersionsResponseBody {
	s.ResourceTypeVersions = v
	return s
}

type ListResourceTypeVersionsResponseBodyResourceTypeVersions struct {
	// The time when the version was created. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the version.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The entity type. Only Module may be returned.
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// Indicates whether the version is the default version. Valid values:
	//
	// *   true
	// *   false
	IsDefaultVersion *bool `json:"IsDefaultVersion,omitempty" xml:"IsDefaultVersion,omitempty"`
	// The provider of the resource type. Valid values:
	//
	// *   ROS: ROS
	// *   Self: yourself
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The time when the version was updated. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version ID.
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListResourceTypeVersionsResponseBodyResourceTypeVersions) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypeVersionsResponseBodyResourceTypeVersions) GoString() string {
	return s.String()
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetCreateTime(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.CreateTime = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetDescription(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.Description = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetEntityType(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.EntityType = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetIsDefaultVersion(v bool) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.IsDefaultVersion = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetProvider(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.Provider = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetResourceType(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetUpdateTime(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.UpdateTime = &v
	return s
}

func (s *ListResourceTypeVersionsResponseBodyResourceTypeVersions) SetVersionId(v string) *ListResourceTypeVersionsResponseBodyResourceTypeVersions {
	s.VersionId = &v
	return s
}

type ListResourceTypeVersionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceTypeVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceTypeVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypeVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTypeVersionsResponse) SetHeaders(v map[string]*string) *ListResourceTypeVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTypeVersionsResponse) SetStatusCode(v int32) *ListResourceTypeVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceTypeVersionsResponse) SetBody(v *ListResourceTypeVersionsResponseBody) *ListResourceTypeVersionsResponse {
	s.Body = v
	return s
}

type ListResourceTypesRequest struct {
	// The entity type. Valid values:
	//
	// *   All: all types of resources.
	// *   Resource (default): regular resources. For more information, see [Resources](~~28863~~).
	// *   DataSource: DataSource resources. For more information, see [DataSource resources](~~404753~~).
	// *   Module: modules.
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The provider of the resource type. Valid values:
	//
	// *   ROS (default): The resource type is provided by Resource Orchestration Service (ROS).
	// *   Self: The resource type is provided by you.
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The resource type. The resource type can contain letters, digits, colons (:), and asterisks (\*). You can use an asterisk (\*) to perform a fuzzy match.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListResourceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesRequest) SetEntityType(v string) *ListResourceTypesRequest {
	s.EntityType = &v
	return s
}

func (s *ListResourceTypesRequest) SetProvider(v string) *ListResourceTypesRequest {
	s.Provider = &v
	return s
}

func (s *ListResourceTypesRequest) SetResourceType(v string) *ListResourceTypesRequest {
	s.ResourceType = &v
	return s
}

type ListResourceTypesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource type summaries.
	ResourceTypeSummaries []*ListResourceTypesResponseBodyResourceTypeSummaries `json:"ResourceTypeSummaries,omitempty" xml:"ResourceTypeSummaries,omitempty" type:"Repeated"`
	// The array of resource types.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s ListResourceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBody) SetRequestId(v string) *ListResourceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceTypesResponseBody) SetResourceTypeSummaries(v []*ListResourceTypesResponseBodyResourceTypeSummaries) *ListResourceTypesResponseBody {
	s.ResourceTypeSummaries = v
	return s
}

func (s *ListResourceTypesResponseBody) SetResourceTypes(v []*string) *ListResourceTypesResponseBody {
	s.ResourceTypes = v
	return s
}

type ListResourceTypesResponseBodyResourceTypeSummaries struct {
	// The creation time. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the default version.
	DefaultVersionId *string `json:"DefaultVersionId,omitempty" xml:"DefaultVersionId,omitempty"`
	// The description of the resource type.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The entity type. Valid values:
	//
	// *   Resource: regular resources.
	// *   DataSource: DataSource resources.
	// *   Module: modules.
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The ID of the latest version.
	LatestVersionId *string `json:"LatestVersionId,omitempty" xml:"LatestVersionId,omitempty"`
	// The provider of the resource type. Valid values:
	//
	// *   ROS: The resource type is provided by ROS.
	// *   Self: The resource type is provided by you.
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The number of versions.
	TotalVersionCount *int32 `json:"TotalVersionCount,omitempty" xml:"TotalVersionCount,omitempty"`
	// The update time. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListResourceTypesResponseBodyResourceTypeSummaries) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponseBodyResourceTypeSummaries) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetCreateTime(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.CreateTime = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetDefaultVersionId(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.DefaultVersionId = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetDescription(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.Description = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetEntityType(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.EntityType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetLatestVersionId(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.LatestVersionId = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetProvider(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.Provider = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetResourceType(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.ResourceType = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetTotalVersionCount(v int32) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.TotalVersionCount = &v
	return s
}

func (s *ListResourceTypesResponseBodyResourceTypeSummaries) SetUpdateTime(v string) *ListResourceTypesResponseBodyResourceTypeSummaries {
	s.UpdateTime = &v
	return s
}

type ListResourceTypesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTypesResponse) SetHeaders(v map[string]*string) *ListResourceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceTypesResponse) SetStatusCode(v int32) *ListResourceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceTypesResponse) SetBody(v *ListResourceTypesResponseBody) *ListResourceTypesResponse {
	s.Body = v
	return s
}

type ListStackEventsRequest struct {
	// The logical IDs of the resources.
	LogicalResourceId []*string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty" type:"Repeated"`
	// The number of the page to return.\
	// Pages start from page 1.\
	// Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.\
	// Maximum value: 50.\
	// Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource types.
	ResourceType []*string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" type:"Repeated"`
	// The stack ID.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The status of the resource.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
}

func (s ListStackEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackEventsRequest) GoString() string {
	return s.String()
}

func (s *ListStackEventsRequest) SetLogicalResourceId(v []*string) *ListStackEventsRequest {
	s.LogicalResourceId = v
	return s
}

func (s *ListStackEventsRequest) SetPageNumber(v int64) *ListStackEventsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStackEventsRequest) SetPageSize(v int64) *ListStackEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackEventsRequest) SetRegionId(v string) *ListStackEventsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackEventsRequest) SetResourceType(v []*string) *ListStackEventsRequest {
	s.ResourceType = v
	return s
}

func (s *ListStackEventsRequest) SetStackId(v string) *ListStackEventsRequest {
	s.StackId = &v
	return s
}

func (s *ListStackEventsRequest) SetStatus(v []*string) *ListStackEventsRequest {
	s.Status = v
	return s
}

type ListStackEventsResponseBody struct {
	// The events.
	Events []*ListStackEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The page number of the returned page.\
	// Pages start from page 1.\
	// Default value: 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.\
	// Maximum value: 50.\
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned events.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListStackEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackEventsResponseBody) SetEvents(v []*ListStackEventsResponseBodyEvents) *ListStackEventsResponseBody {
	s.Events = v
	return s
}

func (s *ListStackEventsResponseBody) SetPageNumber(v int32) *ListStackEventsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackEventsResponseBody) SetPageSize(v int32) *ListStackEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackEventsResponseBody) SetRequestId(v string) *ListStackEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackEventsResponseBody) SetTotalCount(v int32) *ListStackEventsResponseBody {
	s.TotalCount = &v
	return s
}

type ListStackEventsResponseBodyEvents struct {
	// The time when the event was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The event ID.
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The logical ID of the resource. The logical ID indicates the name of the resource that is defined in the template.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The physical ID of the resource.
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The stack ID.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The stack name.
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The state of the resource.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the resource is in the current state.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s ListStackEventsResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s ListStackEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *ListStackEventsResponseBodyEvents) SetCreateTime(v string) *ListStackEventsResponseBodyEvents {
	s.CreateTime = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetEventId(v string) *ListStackEventsResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetLogicalResourceId(v string) *ListStackEventsResponseBodyEvents {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetPhysicalResourceId(v string) *ListStackEventsResponseBodyEvents {
	s.PhysicalResourceId = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetResourceType(v string) *ListStackEventsResponseBodyEvents {
	s.ResourceType = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetStackId(v string) *ListStackEventsResponseBodyEvents {
	s.StackId = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetStackName(v string) *ListStackEventsResponseBodyEvents {
	s.StackName = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetStatus(v string) *ListStackEventsResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *ListStackEventsResponseBodyEvents) SetStatusReason(v string) *ListStackEventsResponseBodyEvents {
	s.StatusReason = &v
	return s
}

type ListStackEventsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListStackEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackEventsResponse) GoString() string {
	return s.String()
}

func (s *ListStackEventsResponse) SetHeaders(v map[string]*string) *ListStackEventsResponse {
	s.Headers = v
	return s
}

func (s *ListStackEventsResponse) SetStatusCode(v int32) *ListStackEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackEventsResponse) SetBody(v *ListStackEventsResponseBody) *ListStackEventsResponse {
	s.Body = v
	return s
}

type ListStackGroupOperationResultsRequest struct {
	// The ID of the operation.
	//
	// You can call the [ListStackGroupOperations](~~151342~~) operation to query the operation ID.
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The number of the page to return.
	//
	// *   Pages start from page 1.
	// *   Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// *   Valid values: 1 to 50.
	// *   Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the stack group.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListStackGroupOperationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationResultsRequest) SetOperationId(v string) *ListStackGroupOperationResultsRequest {
	s.OperationId = &v
	return s
}

func (s *ListStackGroupOperationResultsRequest) SetPageNumber(v int64) *ListStackGroupOperationResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupOperationResultsRequest) SetPageSize(v int64) *ListStackGroupOperationResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupOperationResultsRequest) SetRegionId(v string) *ListStackGroupOperationResultsRequest {
	s.RegionId = &v
	return s
}

type ListStackGroupOperationResultsResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the results of the operation.
	StackGroupOperationResults []*ListStackGroupOperationResultsResponseBodyStackGroupOperationResults `json:"StackGroupOperationResults,omitempty" xml:"StackGroupOperationResults,omitempty" type:"Repeated"`
	// The total number of results.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListStackGroupOperationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationResultsResponseBody) SetPageNumber(v int32) *ListStackGroupOperationResultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBody) SetPageSize(v int32) *ListStackGroupOperationResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBody) SetRequestId(v string) *ListStackGroupOperationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBody) SetStackGroupOperationResults(v []*ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) *ListStackGroupOperationResultsResponseBody {
	s.StackGroupOperationResults = v
	return s
}

func (s *ListStackGroupOperationResultsResponseBody) SetTotalCount(v int32) *ListStackGroupOperationResultsResponseBody {
	s.TotalCount = &v
	return s
}

type ListStackGroupOperationResultsResponseBodyStackGroupOperationResults struct {
	// The ID of the account to which the stack instance belongs.
	//
	// *   If the stack group has self-managed permissions, the stack instance belongs to an Alibaba Cloud account.
	// *   If the stack group has service-managed permissions, the stack instance belongs to a member account in the resource directory.
	//
	// >  For more information about the account, see [Overview](~~154578~~).
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The folder ID of the resource directory.
	//
	// >  This parameter is returned only when the stack group is granted service-managed permissions.
	RdFolderId *string `json:"RdFolderId,omitempty" xml:"RdFolderId,omitempty"`
	// The region ID of the stack instance.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the operation.
	//
	// Valid values:
	//
	// *   RUNNING: The operation is being performed.
	// *   SUCCEEDED: The operation succeeded.
	// *   FAILED: The operation failed.
	// *   STOPPING: The operation is being stopped.
	// *   STOPPED: The operation is stopped.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the operation is in a specific state.
	//
	// >  This parameter is returned only when stack instances are in the OUTDATED state.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetAccountId(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.AccountId = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetRdFolderId(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.RdFolderId = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetRegionId(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.RegionId = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetStatus(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.Status = &v
	return s
}

func (s *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults) SetStatusReason(v string) *ListStackGroupOperationResultsResponseBodyStackGroupOperationResults {
	s.StatusReason = &v
	return s
}

type ListStackGroupOperationResultsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListStackGroupOperationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackGroupOperationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationResultsResponse) SetHeaders(v map[string]*string) *ListStackGroupOperationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListStackGroupOperationResultsResponse) SetStatusCode(v int32) *ListStackGroupOperationResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackGroupOperationResultsResponse) SetBody(v *ListStackGroupOperationResultsResponseBody) *ListStackGroupOperationResultsResponse {
	s.Body = v
	return s
}

type ListStackGroupOperationsRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group. The name must be unique within a region.
	//
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
}

func (s ListStackGroupOperationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationsRequest) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationsRequest) SetPageNumber(v int64) *ListStackGroupOperationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupOperationsRequest) SetPageSize(v int64) *ListStackGroupOperationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupOperationsRequest) SetRegionId(v string) *ListStackGroupOperationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackGroupOperationsRequest) SetStackGroupName(v string) *ListStackGroupOperationsRequest {
	s.StackGroupName = &v
	return s
}

type ListStackGroupOperationsResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stack group operations.
	StackGroupOperations []*ListStackGroupOperationsResponseBodyStackGroupOperations `json:"StackGroupOperations,omitempty" xml:"StackGroupOperations,omitempty" type:"Repeated"`
	// The total number of stack group operations.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListStackGroupOperationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationsResponseBody) SetPageNumber(v int32) *ListStackGroupOperationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupOperationsResponseBody) SetPageSize(v int32) *ListStackGroupOperationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupOperationsResponseBody) SetRequestId(v string) *ListStackGroupOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackGroupOperationsResponseBody) SetStackGroupOperations(v []*ListStackGroupOperationsResponseBodyStackGroupOperations) *ListStackGroupOperationsResponseBody {
	s.StackGroupOperations = v
	return s
}

func (s *ListStackGroupOperationsResponseBody) SetTotalCount(v int32) *ListStackGroupOperationsResponseBody {
	s.TotalCount = &v
	return s
}

type ListStackGroupOperationsResponseBodyStackGroupOperations struct {
	// The operation type.
	//
	// Valid values:
	//
	// *   CREATE
	// *   UPDATE
	// *   DELETE
	// *   DETECT_DRIFT
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The time when the operation was initiated.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the operation ended.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The description of the operation.
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The operation ID.
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the stack group.
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	// The name of the stack group.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The state of the operation.
	//
	// Valid values:
	//
	// *   RUNNING
	// *   SUCCEEDED
	// *   FAILED
	// *   STOPPING
	// *   STOPPED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListStackGroupOperationsResponseBodyStackGroupOperations) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationsResponseBodyStackGroupOperations) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetAction(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.Action = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetCreateTime(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.CreateTime = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetEndTime(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.EndTime = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetOperationDescription(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.OperationDescription = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetOperationId(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.OperationId = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetStackGroupId(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.StackGroupId = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetStackGroupName(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.StackGroupName = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetStatus(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.Status = &v
	return s
}

type ListStackGroupOperationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListStackGroupOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackGroupOperationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupOperationsResponse) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationsResponse) SetHeaders(v map[string]*string) *ListStackGroupOperationsResponse {
	s.Headers = v
	return s
}

func (s *ListStackGroupOperationsResponse) SetStatusCode(v int32) *ListStackGroupOperationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackGroupOperationsResponse) SetBody(v *ListStackGroupOperationsResponseBody) *ListStackGroupOperationsResponse {
	s.Body = v
	return s
}

type ListStackGroupsRequest struct {
	// The number of the page to return.
	//
	// *   Pages start from page 1.
	// *   Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// *   Valid values: 1 to 50.
	// *   Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. If you do not specify this parameter, the stack groups in all the resource groups are queried.
	//
	// > To obtain the resource group ID, go to the **Resource Group** page in the **Resource Management** console. For more information, see [View the basic information about a resource group](~~151181~~).
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The state of the stack group. If you do not specify this parameter, the stack groups in all states in the specified region are queried.
	//
	// Valid values:
	//
	// *   ACTIVE
	// *   DELETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the stack group.
	Tags []*ListStackGroupsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListStackGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListStackGroupsRequest) SetPageNumber(v int64) *ListStackGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupsRequest) SetPageSize(v int64) *ListStackGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupsRequest) SetRegionId(v string) *ListStackGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackGroupsRequest) SetResourceGroupId(v string) *ListStackGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStackGroupsRequest) SetStatus(v string) *ListStackGroupsRequest {
	s.Status = &v
	return s
}

func (s *ListStackGroupsRequest) SetTags(v []*ListStackGroupsRequestTags) *ListStackGroupsRequest {
	s.Tags = v
	return s
}

type ListStackGroupsRequestTags struct {
	// The key of the tag that is added to the stack group.
	//
	// > Tags is optional. If you specify Tags, you must specify Tags.N.Key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is added to the stack group.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListStackGroupsRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupsRequestTags) GoString() string {
	return s.String()
}

func (s *ListStackGroupsRequestTags) SetKey(v string) *ListStackGroupsRequestTags {
	s.Key = &v
	return s
}

func (s *ListStackGroupsRequestTags) SetValue(v string) *ListStackGroupsRequestTags {
	s.Value = &v
	return s
}

type ListStackGroupsResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stack groups.
	StackGroups []*ListStackGroupsResponseBodyStackGroups `json:"StackGroups,omitempty" xml:"StackGroups,omitempty" type:"Repeated"`
	// The total number of stack groups.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListStackGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackGroupsResponseBody) SetPageNumber(v int32) *ListStackGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupsResponseBody) SetPageSize(v int32) *ListStackGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupsResponseBody) SetRequestId(v string) *ListStackGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackGroupsResponseBody) SetStackGroups(v []*ListStackGroupsResponseBodyStackGroups) *ListStackGroupsResponseBody {
	s.StackGroups = v
	return s
}

func (s *ListStackGroupsResponseBody) SetTotalCount(v int32) *ListStackGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListStackGroupsResponseBodyStackGroups struct {
	// The information about automatic deployment settings.
	AutoDeployment *ListStackGroupsResponseBodyStackGroupsAutoDeployment `json:"AutoDeployment,omitempty" xml:"AutoDeployment,omitempty" type:"Struct"`
	// The description of the stack group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the most recent successful drift detection was performed on the stack group.
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The permission model of the stack group.
	//
	// Valid values:
	//
	// *   SELF_MANAGED
	// *   SERVICE_MANAGED
	//
	// > For more information about the permission models of stack groups, see [Overview](~~154578~~).
	PermissionModel *string `json:"PermissionModel,omitempty" xml:"PermissionModel,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The drift state of the stack group on which the most recent successful drift detection was performed.
	//
	// Valid values:
	//
	// *   DRIFTED: The stack group has drifted.
	// *   NOT_CHECKED: No drift detection is performed on the stack group.
	// *   IN_SYNC: No drifts are detected on the stack group.
	StackGroupDriftStatus *string `json:"StackGroupDriftStatus,omitempty" xml:"StackGroupDriftStatus,omitempty"`
	// The ID of the stack group.
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	// The name of the stack group.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The state of the stack group.
	//
	// Valid values:
	//
	// *   ACTIVE
	// *   DELETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the stack group.
	Tags []*ListStackGroupsResponseBodyStackGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListStackGroupsResponseBodyStackGroups) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupsResponseBodyStackGroups) GoString() string {
	return s.String()
}

func (s *ListStackGroupsResponseBodyStackGroups) SetAutoDeployment(v *ListStackGroupsResponseBodyStackGroupsAutoDeployment) *ListStackGroupsResponseBodyStackGroups {
	s.AutoDeployment = v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetDescription(v string) *ListStackGroupsResponseBodyStackGroups {
	s.Description = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetDriftDetectionTime(v string) *ListStackGroupsResponseBodyStackGroups {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetPermissionModel(v string) *ListStackGroupsResponseBodyStackGroups {
	s.PermissionModel = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetResourceGroupId(v string) *ListStackGroupsResponseBodyStackGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetStackGroupDriftStatus(v string) *ListStackGroupsResponseBodyStackGroups {
	s.StackGroupDriftStatus = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetStackGroupId(v string) *ListStackGroupsResponseBodyStackGroups {
	s.StackGroupId = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetStackGroupName(v string) *ListStackGroupsResponseBodyStackGroups {
	s.StackGroupName = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetStatus(v string) *ListStackGroupsResponseBodyStackGroups {
	s.Status = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetTags(v []*ListStackGroupsResponseBodyStackGroupsTags) *ListStackGroupsResponseBodyStackGroups {
	s.Tags = v
	return s
}

type ListStackGroupsResponseBodyStackGroupsAutoDeployment struct {
	// Indicates whether automatic deployment is enabled.
	//
	// Valid values:
	//
	// *   true: Automatic deployment is enabled. If you add a member to the folder to which the stack group belongs after automatic deployment is enabled, Resource Orchestration Service (ROS) automatically adds the stack instances in the stack group to the specified region of the member. If you delete the member from the folder, ROS automatically deletes the stack instances in the stack group from the specified region of the member.
	// *   false: Automatic deployment is disabled. After you disable automatic deployment, the stack instances remain unchanged when you change the member in the folder.
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether the stacks within a member are retained when you delete the member from the folder.
	//
	// Valid values:
	//
	// *   true
	// *   false
	//
	// > This parameter is returned only if Enabled is set to true.
	RetainStacksOnAccountRemoval *bool `json:"RetainStacksOnAccountRemoval,omitempty" xml:"RetainStacksOnAccountRemoval,omitempty"`
}

func (s ListStackGroupsResponseBodyStackGroupsAutoDeployment) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupsResponseBodyStackGroupsAutoDeployment) GoString() string {
	return s.String()
}

func (s *ListStackGroupsResponseBodyStackGroupsAutoDeployment) SetEnabled(v bool) *ListStackGroupsResponseBodyStackGroupsAutoDeployment {
	s.Enabled = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroupsAutoDeployment) SetRetainStacksOnAccountRemoval(v bool) *ListStackGroupsResponseBodyStackGroupsAutoDeployment {
	s.RetainStacksOnAccountRemoval = &v
	return s
}

type ListStackGroupsResponseBodyStackGroupsTags struct {
	// The key of the tag that is added to the stack group.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is added to the stack group.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListStackGroupsResponseBodyStackGroupsTags) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupsResponseBodyStackGroupsTags) GoString() string {
	return s.String()
}

func (s *ListStackGroupsResponseBodyStackGroupsTags) SetKey(v string) *ListStackGroupsResponseBodyStackGroupsTags {
	s.Key = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroupsTags) SetValue(v string) *ListStackGroupsResponseBodyStackGroupsTags {
	s.Value = &v
	return s
}

type ListStackGroupsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListStackGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListStackGroupsResponse) SetHeaders(v map[string]*string) *ListStackGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListStackGroupsResponse) SetStatusCode(v int32) *ListStackGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackGroupsResponse) SetBody(v *ListStackGroupsResponseBody) *ListStackGroupsResponse {
	s.Body = v
	return s
}

type ListStackInstancesRequest struct {
	// The number of the page to return.
	//
	// *   Pages start from page 1.
	// *   Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// *   Valid values: 1 to 50.
	// *   Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the stack group. The name must be unique within a region.\
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The ID of the destination account to which the stack belongs.
	//
	// *   If the stack group is granted self-managed permissions, the stack belongs to an Alibaba Cloud account.
	// *   If the stack group is granted service-managed permissions, the stack belongs to a member in a resource directory.
	//
	// > For more information about the destination account, see [Overview](~~154578~~).
	StackInstanceAccountId *string `json:"StackInstanceAccountId,omitempty" xml:"StackInstanceAccountId,omitempty"`
	// The region ID of the stack.
	StackInstanceRegionId *string `json:"StackInstanceRegionId,omitempty" xml:"StackInstanceRegionId,omitempty"`
}

func (s ListStackInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListStackInstancesRequest) SetPageNumber(v int64) *ListStackInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStackInstancesRequest) SetPageSize(v int64) *ListStackInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackInstancesRequest) SetRegionId(v string) *ListStackInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackInstancesRequest) SetStackGroupName(v string) *ListStackInstancesRequest {
	s.StackGroupName = &v
	return s
}

func (s *ListStackInstancesRequest) SetStackInstanceAccountId(v string) *ListStackInstancesRequest {
	s.StackInstanceAccountId = &v
	return s
}

func (s *ListStackInstancesRequest) SetStackInstanceRegionId(v string) *ListStackInstancesRequest {
	s.StackInstanceRegionId = &v
	return s
}

type ListStackInstancesResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stacks.
	StackInstances []*ListStackInstancesResponseBodyStackInstances `json:"StackInstances,omitempty" xml:"StackInstances,omitempty" type:"Repeated"`
	// The total number of stacks.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListStackInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackInstancesResponseBody) SetPageNumber(v int32) *ListStackInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackInstancesResponseBody) SetPageSize(v int32) *ListStackInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackInstancesResponseBody) SetRequestId(v string) *ListStackInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackInstancesResponseBody) SetStackInstances(v []*ListStackInstancesResponseBodyStackInstances) *ListStackInstancesResponseBody {
	s.StackInstances = v
	return s
}

func (s *ListStackInstancesResponseBody) SetTotalCount(v int32) *ListStackInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListStackInstancesResponseBodyStackInstances struct {
	// The ID of the destination account to which the stack belongs.
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the last successful drift detection was performed on the stack.
	//
	// > This parameter is returned only if drift detection is performed on the stack group.
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The ID of the folder in the resource directory.
	//
	// > This parameter is returned only if the stack group is granted service-managed permissions.
	RdFolderId *string `json:"RdFolderId,omitempty" xml:"RdFolderId,omitempty"`
	// The region ID of the stack.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The state of the stack when the last successful drift detection was performed on the stack group.
	//
	// Valid values:
	//
	// *   DRIFTED: The stack has drifted.
	// *   NOT_CHECKED: No successful drift detection is performed on the stack.
	// *   IN_SYNC: The stack is being synchronized.
	//
	// > This parameter is returned only if drift detection is performed on the stack group.
	StackDriftStatus *string `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	// The ID of the stack group.
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	// The name of the stack group.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The stack ID.
	//
	// > This parameter is returned only if the stack is in the CURRENT state.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The state of the stack.
	//
	// Valid values:
	//
	// *   CURRENT: The stack is up-to-date with the stack group.
	//
	// *   OUTDATED: The stack is not up-to-date with the stack group. Stacks are in the OUTDATED state due to the following possible reasons:
	//
	//     *   When the CreateStackInstances operation is called to create stacks, the stacks fail to be created.
	//     *   When the UpdateStackInstances or UpdateStackGroup operation is called to update stacks, the stacks fail to be updated, or only specific stacks are updated.
	//     *   The creation or update operation is not complete.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the stack instance is in the OUTDATED state.
	//
	// > This parameter is returned only if the stack instance is in the OUTDATED state.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
}

func (s ListStackInstancesResponseBodyStackInstances) String() string {
	return tea.Prettify(s)
}

func (s ListStackInstancesResponseBodyStackInstances) GoString() string {
	return s.String()
}

func (s *ListStackInstancesResponseBodyStackInstances) SetAccountId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.AccountId = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetDriftDetectionTime(v string) *ListStackInstancesResponseBodyStackInstances {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetRdFolderId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.RdFolderId = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetRegionId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.RegionId = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStackDriftStatus(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StackDriftStatus = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStackGroupId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StackGroupId = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStackGroupName(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StackGroupName = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStackId(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StackId = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStatus(v string) *ListStackInstancesResponseBodyStackInstances {
	s.Status = &v
	return s
}

func (s *ListStackInstancesResponseBodyStackInstances) SetStatusReason(v string) *ListStackInstancesResponseBodyStackInstances {
	s.StatusReason = &v
	return s
}

type ListStackInstancesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListStackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListStackInstancesResponse) SetHeaders(v map[string]*string) *ListStackInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListStackInstancesResponse) SetStatusCode(v int32) *ListStackInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackInstancesResponse) SetBody(v *ListStackInstancesResponseBody) *ListStackInstancesResponse {
	s.Body = v
	return s
}

type ListStackOperationRisksRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (\_). For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The type of the operation of which you want to detect risks. Valid values:
	//
	// *   DeleteStack: detects high risks that may arise in resources when you delete a stack.
	// *   CreateStack: detects the missing permissions when you fail to create a stack.
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// The name of the RAM role.
	//
	// *   If you specify a RAM role, ROS creates stacks based on the permissions that are granted to the RAM role and uses the credentials of the RAM role to call the API operations of Alibaba Cloud services.
	// *   If you do not specify a RAM role, ROS creates stacks based on the permissions of your Alibaba Cloud account.
	//
	// The name of the RAM role can be up to 64 bytes in length.
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to retain all resources in the stack. Valid values:
	//
	// *   true
	// *   false (default)
	//
	// > This parameter takes effect only if you set OperationType to DeleteStack.
	RetainAllResources *bool `json:"RetainAllResources,omitempty" xml:"RetainAllResources,omitempty"`
	// The list of resources to retain.
	//
	// > This parameter takes effect only if you set OperationType to DeleteStack.
	RetainResources []*string `json:"RetainResources,omitempty" xml:"RetainResources,omitempty" type:"Repeated"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.
	//
	// > You must specify one of TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// > You must specify one of TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo and oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length. If you do not specify RegionId in the URL, the region ID of the stack is used.
	//
	// > You must specify one of TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template.
	//
	// > This parameter takes effect only if you specify TemplateId.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s ListStackOperationRisksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackOperationRisksRequest) GoString() string {
	return s.String()
}

func (s *ListStackOperationRisksRequest) SetClientToken(v string) *ListStackOperationRisksRequest {
	s.ClientToken = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetOperationType(v string) *ListStackOperationRisksRequest {
	s.OperationType = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetRamRoleName(v string) *ListStackOperationRisksRequest {
	s.RamRoleName = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetRegionId(v string) *ListStackOperationRisksRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetRetainAllResources(v bool) *ListStackOperationRisksRequest {
	s.RetainAllResources = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetRetainResources(v []*string) *ListStackOperationRisksRequest {
	s.RetainResources = v
	return s
}

func (s *ListStackOperationRisksRequest) SetStackId(v string) *ListStackOperationRisksRequest {
	s.StackId = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetTemplateBody(v string) *ListStackOperationRisksRequest {
	s.TemplateBody = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetTemplateId(v string) *ListStackOperationRisksRequest {
	s.TemplateId = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetTemplateURL(v string) *ListStackOperationRisksRequest {
	s.TemplateURL = &v
	return s
}

func (s *ListStackOperationRisksRequest) SetTemplateVersion(v string) *ListStackOperationRisksRequest {
	s.TemplateVersion = &v
	return s
}

type ListStackOperationRisksResponseBody struct {
	// The operations on which the permissions are not granted to the Alibaba Cloud account of the caller.
	MissingPolicyActions []*string `json:"MissingPolicyActions,omitempty" xml:"MissingPolicyActions,omitempty" type:"Repeated"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources that are at risk.
	RiskResources []*ListStackOperationRisksResponseBodyRiskResources `json:"RiskResources,omitempty" xml:"RiskResources,omitempty" type:"Repeated"`
}

func (s ListStackOperationRisksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackOperationRisksResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackOperationRisksResponseBody) SetMissingPolicyActions(v []*string) *ListStackOperationRisksResponseBody {
	s.MissingPolicyActions = v
	return s
}

func (s *ListStackOperationRisksResponseBody) SetRequestId(v string) *ListStackOperationRisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackOperationRisksResponseBody) SetRiskResources(v []*ListStackOperationRisksResponseBodyRiskResources) *ListStackOperationRisksResponseBody {
	s.RiskResources = v
	return s
}

type ListStackOperationRisksResponseBodyRiskResources struct {
	// The error code that is returned when the risk detection fails.
	//
	// >  This parameter is not returned if the risk detection is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The logical ID of the resource. The logical ID is the resource name that is defined in the template.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The error message that is returned when the risk detection fails.
	//
	// >  This parameter is not returned if the risk detection is successful.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The physical ID of the resource. The physical ID is the actual ID of the resource.
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The cause of the risk.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the request when the risk detection fails.
	//
	// >  This parameter is not returned if the risk detection is successful.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The type of the risk. Valid values:
	//
	// *   Referenced: The resource is referenced by other resources.
	// *   MaybeReferenced: The resource may be referenced by other resources.
	// *   AdditionalRiskCheckRequired: An additional risk detection is required for a nested stack.
	// *   OperationIgnored: The operation does not take effect for the resource.
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
}

func (s ListStackOperationRisksResponseBodyRiskResources) String() string {
	return tea.Prettify(s)
}

func (s ListStackOperationRisksResponseBodyRiskResources) GoString() string {
	return s.String()
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetCode(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.Code = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetLogicalResourceId(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetMessage(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.Message = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetPhysicalResourceId(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.PhysicalResourceId = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetReason(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.Reason = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetRequestId(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.RequestId = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetResourceType(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.ResourceType = &v
	return s
}

func (s *ListStackOperationRisksResponseBodyRiskResources) SetRiskType(v string) *ListStackOperationRisksResponseBodyRiskResources {
	s.RiskType = &v
	return s
}

type ListStackOperationRisksResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListStackOperationRisksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackOperationRisksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackOperationRisksResponse) GoString() string {
	return s.String()
}

func (s *ListStackOperationRisksResponse) SetHeaders(v map[string]*string) *ListStackOperationRisksResponse {
	s.Headers = v
	return s
}

func (s *ListStackOperationRisksResponse) SetStatusCode(v int32) *ListStackOperationRisksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackOperationRisksResponse) SetBody(v *ListStackOperationRisksResponseBody) *ListStackOperationRisksResponse {
	s.Body = v
	return s
}

type ListStackResourceDriftsRequest struct {
	// The time when the resource drift detection operation was initiated.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The type of the resource.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The physical ID of the resource.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource properties as defined in the template, in JSON format.
	ResourceDriftStatus []*string `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty" type:"Repeated"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s ListStackResourceDriftsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourceDriftsRequest) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsRequest) SetMaxResults(v int64) *ListStackResourceDriftsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListStackResourceDriftsRequest) SetNextToken(v string) *ListStackResourceDriftsRequest {
	s.NextToken = &v
	return s
}

func (s *ListStackResourceDriftsRequest) SetRegionId(v string) *ListStackResourceDriftsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackResourceDriftsRequest) SetResourceDriftStatus(v []*string) *ListStackResourceDriftsRequest {
	s.ResourceDriftStatus = v
	return s
}

func (s *ListStackResourceDriftsRequest) SetStackId(v string) *ListStackResourceDriftsRequest {
	s.StackId = &v
	return s
}

type ListStackResourceDriftsResponseBody struct {
	// The query token returned in this call.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource drifts.
	ResourceDrifts []*ListStackResourceDriftsResponseBodyResourceDrifts `json:"ResourceDrifts,omitempty" xml:"ResourceDrifts,omitempty" type:"Repeated"`
}

func (s ListStackResourceDriftsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourceDriftsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponseBody) SetNextToken(v string) *ListStackResourceDriftsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListStackResourceDriftsResponseBody) SetRequestId(v string) *ListStackResourceDriftsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackResourceDriftsResponseBody) SetResourceDrifts(v []*ListStackResourceDriftsResponseBodyResourceDrifts) *ListStackResourceDriftsResponseBody {
	s.ResourceDrifts = v
	return s
}

type ListStackResourceDriftsResponseBodyResourceDrifts struct {
	// The actual JSON-formatted resource properties.
	ActualProperties *string `json:"ActualProperties,omitempty" xml:"ActualProperties,omitempty"`
	// The time when the drift detection operation was performed on the resource.
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The JSON-formatted resource properties that are defined in the template.
	ExpectedProperties *string `json:"ExpectedProperties,omitempty" xml:"ExpectedProperties,omitempty"`
	// The logical ID of the resource. The logical ID indicates the name of the resource that is defined in the template.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The information about the modules from which the resource was created. This parameter is returned only if the resource is created from modules.
	ModuleInfo *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo `json:"ModuleInfo,omitempty" xml:"ModuleInfo,omitempty" type:"Struct"`
	// The physical ID of the resource.
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The property drifts of the resource.
	PropertyDifferences []*ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences `json:"PropertyDifferences,omitempty" xml:"PropertyDifferences,omitempty" type:"Repeated"`
	// The drift state of the resource. Valid values:
	//
	// *   DELETED: The actual configuration of the resource differs from its expected template configuration because the resource is deleted.
	// *   MODIFIED: The actual configuration of the resource differs from its expected template configuration.
	// *   NOT_CHECKED: Resource Orchestration Service (ROS) has not checked whether the actual configuration of the resource differs from its expected template configuration.
	// *   IN_SYNC: The actual configuration of the resource matches its expected template configuration.
	ResourceDriftStatus *string `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The stack ID.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s ListStackResourceDriftsResponseBodyResourceDrifts) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourceDriftsResponseBodyResourceDrifts) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetActualProperties(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ActualProperties = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetDriftDetectionTime(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetExpectedProperties(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ExpectedProperties = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetLogicalResourceId(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetModuleInfo(v *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ModuleInfo = v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetPhysicalResourceId(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.PhysicalResourceId = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetPropertyDifferences(v []*ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.PropertyDifferences = v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetResourceDriftStatus(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ResourceDriftStatus = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetResourceType(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.ResourceType = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDrifts) SetStackId(v string) *ListStackResourceDriftsResponseBodyResourceDrifts {
	s.StackId = &v
	return s
}

type ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo struct {
	// The concatenated logical IDs of one or more modules that contain the resource. The modules are listed from the outermost layer and separated by forward slashes (`/`).
	//
	// In the following example, the resource is created from Module B nested within Parent Module A:
	//
	// `moduleA/moduleB`
	LogicalIdHierarchy *string `json:"LogicalIdHierarchy,omitempty" xml:"LogicalIdHierarchy,omitempty"`
	// The concatenated types of one or more modules that contain the resource. The module types are listed from the outermost layer and separated by forward slashes (`/`).
	//
	// In the following example, the resource is created from a module of the `MODULE::ROS::Child::Example` type that is nested within a parent module of the `MODULE::ROS::Parent::Example` type:
	//
	// `MODULE::ROS::Parent::Example/MODULE::ROS::Child::Example`
	TypeHierarchy *string `json:"TypeHierarchy,omitempty" xml:"TypeHierarchy,omitempty"`
}

func (s ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo) SetLogicalIdHierarchy(v string) *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo {
	s.LogicalIdHierarchy = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo) SetTypeHierarchy(v string) *ListStackResourceDriftsResponseBodyResourceDriftsModuleInfo {
	s.TypeHierarchy = &v
	return s
}

type ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences struct {
	// The actual value of the resource property.
	ActualValue *string `json:"ActualValue,omitempty" xml:"ActualValue,omitempty"`
	// The drift type of the resource property. Valid values:
	//
	// *   ADD: The value is added to a resource property whose data type is Array or List.
	// *   REMOVE: The property is deleted from the current resource configuration.
	// *   NOT_EQUAL: The current property value differs from the expected value that is defined in the stack template.
	DifferenceType *string `json:"DifferenceType,omitempty" xml:"DifferenceType,omitempty"`
	// The expected value of the resource property that is defined in the template.
	ExpectedValue *string `json:"ExpectedValue,omitempty" xml:"ExpectedValue,omitempty"`
	// The path of the resource property.
	PropertyPath *string `json:"PropertyPath,omitempty" xml:"PropertyPath,omitempty"`
}

func (s ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) SetActualValue(v string) *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences {
	s.ActualValue = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) SetDifferenceType(v string) *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences {
	s.DifferenceType = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) SetExpectedValue(v string) *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences {
	s.ExpectedValue = &v
	return s
}

func (s *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences) SetPropertyPath(v string) *ListStackResourceDriftsResponseBodyResourceDriftsPropertyDifferences {
	s.PropertyPath = &v
	return s
}

type ListStackResourceDriftsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListStackResourceDriftsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackResourceDriftsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourceDriftsResponse) GoString() string {
	return s.String()
}

func (s *ListStackResourceDriftsResponse) SetHeaders(v map[string]*string) *ListStackResourceDriftsResponse {
	s.Headers = v
	return s
}

func (s *ListStackResourceDriftsResponse) SetStatusCode(v int32) *ListStackResourceDriftsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackResourceDriftsResponse) SetBody(v *ListStackResourceDriftsResponseBody) *ListStackResourceDriftsResponse {
	s.Body = v
	return s
}

type ListStackResourcesRequest struct {
	// The ID of the request.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the region to which the stack belongs. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s ListStackResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListStackResourcesRequest) SetRegionId(v string) *ListStackResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackResourcesRequest) SetStackId(v string) *ListStackResourcesRequest {
	s.StackId = &v
	return s
}

type ListStackResourcesResponseBody struct {
	// Details about resources.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources.
	Resources []*ListStackResourcesResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ListStackResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackResourcesResponseBody) SetRequestId(v string) *ListStackResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackResourcesResponseBody) SetResources(v []*ListStackResourcesResponseBodyResources) *ListStackResourcesResponseBody {
	s.Resources = v
	return s
}

type ListStackResourcesResponseBodyResources struct {
	// The time when the resource was created. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the most recent successful drift detection was performed on the stack.
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The logical ID of the resource. The logical ID is the resource name that is defined in the template.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The information about the modules from which the resource is created. This parameter is returned only if the resource is created from modules.
	ModuleInfo *ListStackResourcesResponseBodyResourcesModuleInfo `json:"ModuleInfo,omitempty" xml:"ModuleInfo,omitempty" type:"Struct"`
	// The physical ID of the resource.
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The drift state of the resource in the most recent successful drift detection. Valid values:
	//
	// *   DELETED: The actual configuration of the resource differs from its expected template configuration because the resource is deleted.
	// *   MODIFIED: The actual configuration of the resource differs from its expected template configuration.
	// *   NOT_CHECKED: Resource Orchestration Service (ROS) has not checked whether the actual configuration of the resource differs from its expected template configuration.
	// *   IN_SYNC: The actual configuration of the resource matches its expected template configuration.
	ResourceDriftStatus *string `json:"ResourceDriftStatus,omitempty" xml:"ResourceDriftStatus,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The stack ID.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The stack name.\
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or letter.
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The state of the resource. Valid values:
	//
	// *   INIT_COMPLETE: The resource is pending to be created.
	// *   CREATE_COMPLETE: The resource is created.
	// *   CREATE_FAILED: The resource failed to be created.
	// *   CREATE_IN_PROGRESS: The resource is being created.
	// *   UPDATE_IN_PROGRESS: The resource is being updated.
	// *   UPDATE_FAILED: The resource failed to be updated.
	// *   UPDATE_COMPLETE: The resource is updated.
	// *   DELETE_IN_PROGRESS: The resource is being deleted.
	// *   DELETE_FAILED: The resource failed to be deleted.
	// *   DELETE_COMPLETE: The resource is deleted.
	// *   CHECK_IN_PROGRESS: The resource is being validated.
	// *   CHECK_FAILED: The resource failed to be validated.
	// *   CHECK_COMPLETE: The resource is validated.
	// *   IMPORT_IN_PROGRESS: The resource is being imported.
	// *   IMPORT_FAILED: The resource failed to be imported.
	// *   IMPORT_COMPLETE: The resource is imported.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the resource is in its current state.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The time when the resource was updated. The time is displayed in UTC. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListStackResourcesResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourcesResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ListStackResourcesResponseBodyResources) SetCreateTime(v string) *ListStackResourcesResponseBodyResources {
	s.CreateTime = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetDriftDetectionTime(v string) *ListStackResourcesResponseBodyResources {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetLogicalResourceId(v string) *ListStackResourcesResponseBodyResources {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetModuleInfo(v *ListStackResourcesResponseBodyResourcesModuleInfo) *ListStackResourcesResponseBodyResources {
	s.ModuleInfo = v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetPhysicalResourceId(v string) *ListStackResourcesResponseBodyResources {
	s.PhysicalResourceId = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetResourceDriftStatus(v string) *ListStackResourcesResponseBodyResources {
	s.ResourceDriftStatus = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetResourceType(v string) *ListStackResourcesResponseBodyResources {
	s.ResourceType = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetStackId(v string) *ListStackResourcesResponseBodyResources {
	s.StackId = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetStackName(v string) *ListStackResourcesResponseBodyResources {
	s.StackName = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetStatus(v string) *ListStackResourcesResponseBodyResources {
	s.Status = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetStatusReason(v string) *ListStackResourcesResponseBodyResources {
	s.StatusReason = &v
	return s
}

func (s *ListStackResourcesResponseBodyResources) SetUpdateTime(v string) *ListStackResourcesResponseBodyResources {
	s.UpdateTime = &v
	return s
}

type ListStackResourcesResponseBodyResourcesModuleInfo struct {
	// The concatenated logical IDs of one or more modules that contain the resource. The modules are listed from the outermost layer and separated by forward slashes (`/`).
	//
	// In the following example, the resource is created from Module B nested within Parent Module A:
	//
	// `moduleA/moduleB`
	LogicalIdHierarchy *string `json:"LogicalIdHierarchy,omitempty" xml:"LogicalIdHierarchy,omitempty"`
	// The concatenated types of one or more modules that contain the resource. The module types are listed from the outermost layer and separated by forward slashes (`/`).
	//
	// In the following example, the resource is created from a module of the `MODULE::ROS::Child::Example` type that is nested within a parent module of the `MODULE::ROS::Parent::Example` type:
	//
	// `MODULE::ROS::Parent::Example/MODULE::ROS::Child::Example`
	TypeHierarchy *string `json:"TypeHierarchy,omitempty" xml:"TypeHierarchy,omitempty"`
}

func (s ListStackResourcesResponseBodyResourcesModuleInfo) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourcesResponseBodyResourcesModuleInfo) GoString() string {
	return s.String()
}

func (s *ListStackResourcesResponseBodyResourcesModuleInfo) SetLogicalIdHierarchy(v string) *ListStackResourcesResponseBodyResourcesModuleInfo {
	s.LogicalIdHierarchy = &v
	return s
}

func (s *ListStackResourcesResponseBodyResourcesModuleInfo) SetTypeHierarchy(v string) *ListStackResourcesResponseBodyResourcesModuleInfo {
	s.TypeHierarchy = &v
	return s
}

type ListStackResourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListStackResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStackResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStackResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListStackResourcesResponse) SetHeaders(v map[string]*string) *ListStackResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListStackResourcesResponse) SetStatusCode(v int32) *ListStackResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStackResourcesResponse) SetBody(v *ListStackResourcesResponseBody) *ListStackResourcesResponse {
	s.Body = v
	return s
}

type ListStacksRequest struct {
	// The end of the time range during which data was queried. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent stack.
	ParentStackId *string `json:"ParentStackId,omitempty" xml:"ParentStackId,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.\
	// For more information about resource groups, see the "Resource Group" section of the [What is Resource Management?](~~94475~~) topic.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether to return nested stacks. Valid values:
	//
	// *   true
	// *   false (default)
	//
	// > If you specify ParentStackId, you must set ShowNestedStack to true.
	ShowNestedStack *bool `json:"ShowNestedStack,omitempty" xml:"ShowNestedStack,omitempty"`
	// The stack ID. You can specify this parameter to query only the stack ID. If you want to query the detailed information about the stack, call the GetStack operation.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The IDs of the stacks.
	StackIds []*string `json:"StackIds,omitempty" xml:"StackIds,omitempty" type:"Repeated"`
	// The names of the stacks.
	StackName []*string `json:"StackName,omitempty" xml:"StackName,omitempty" type:"Repeated"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time must be in UTC.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the stack.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The tags of the stack.
	Tag []*ListStacksRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListStacksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStacksRequest) GoString() string {
	return s.String()
}

func (s *ListStacksRequest) SetEndTime(v string) *ListStacksRequest {
	s.EndTime = &v
	return s
}

func (s *ListStacksRequest) SetPageNumber(v int64) *ListStacksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStacksRequest) SetPageSize(v int64) *ListStacksRequest {
	s.PageSize = &v
	return s
}

func (s *ListStacksRequest) SetParentStackId(v string) *ListStacksRequest {
	s.ParentStackId = &v
	return s
}

func (s *ListStacksRequest) SetRegionId(v string) *ListStacksRequest {
	s.RegionId = &v
	return s
}

func (s *ListStacksRequest) SetResourceGroupId(v string) *ListStacksRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStacksRequest) SetShowNestedStack(v bool) *ListStacksRequest {
	s.ShowNestedStack = &v
	return s
}

func (s *ListStacksRequest) SetStackId(v string) *ListStacksRequest {
	s.StackId = &v
	return s
}

func (s *ListStacksRequest) SetStackIds(v []*string) *ListStacksRequest {
	s.StackIds = v
	return s
}

func (s *ListStacksRequest) SetStackName(v []*string) *ListStacksRequest {
	s.StackName = v
	return s
}

func (s *ListStacksRequest) SetStartTime(v string) *ListStacksRequest {
	s.StartTime = &v
	return s
}

func (s *ListStacksRequest) SetStatus(v []*string) *ListStacksRequest {
	s.Status = v
	return s
}

func (s *ListStacksRequest) SetTag(v []*ListStacksRequestTag) *ListStacksRequest {
	s.Tag = v
	return s
}

type ListStacksRequestTag struct {
	// The key of tag N.\
	// Valid values of N: 1 to 20.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.\
	// Valid values of N: 1 to 20.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListStacksRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListStacksRequestTag) GoString() string {
	return s.String()
}

func (s *ListStacksRequestTag) SetKey(v string) *ListStacksRequestTag {
	s.Key = &v
	return s
}

func (s *ListStacksRequestTag) SetValue(v string) *ListStacksRequestTag {
	s.Value = &v
	return s
}

type ListStacksResponseBody struct {
	// The page number.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the stacks.
	Stacks []*ListStacksResponseBodyStacks `json:"Stacks,omitempty" xml:"Stacks,omitempty" type:"Repeated"`
	// The total number of stacks.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListStacksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStacksResponseBody) GoString() string {
	return s.String()
}

func (s *ListStacksResponseBody) SetPageNumber(v int32) *ListStacksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStacksResponseBody) SetPageSize(v int32) *ListStacksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStacksResponseBody) SetRequestId(v string) *ListStacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStacksResponseBody) SetStacks(v []*ListStacksResponseBodyStacks) *ListStacksResponseBody {
	s.Stacks = v
	return s
}

func (s *ListStacksResponseBody) SetTotalCount(v int32) *ListStacksResponseBody {
	s.TotalCount = &v
	return s
}

type ListStacksResponseBodyStacks struct {
	// The time when the stack was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether deletion protection is enabled for the stack. Valid values:
	//
	// *   Enabled: Deletion protection is enabled for the stack.
	// *   Disabled: Deletion protection is disabled for the stack. In this case, you can delete the stack by using the console or calling the [DeleteStack](~~610812~~) operation.
	//
	// >  Deletion protection of a nested stack is the same as that of its root stack.
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Indicates whether rollback is disabled when the stack fails to be created. Valid values:
	//
	// *   true
	// *   false (default)
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The time when the most recent successful drift detection was performed on the stack.
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The supplementary information that is returned if an error occurs on a stack operation.
	//
	// >  This parameter is returned only under specific conditions, and is returned together with at least one sub-parameter. For example, an error occurred when an API operation of another Alibaba Cloud service was called.
	OperationInfo *ListStacksResponseBodyStacksOperationInfo `json:"OperationInfo,omitempty" xml:"OperationInfo,omitempty" type:"Struct"`
	// The ID of the parent stack.
	ParentStackId *string `json:"ParentStackId,omitempty" xml:"ParentStackId,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the stack is a managed stack. Valid values:
	//
	// *   true
	// *   false
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The name of the service to which the managed stack belongs.
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The state of the stack on which the most recent successful drift detection was performed. Valid values:
	//
	// *   DRIFTED: The stack has drifted.
	// *   NOT_CHECKED: No successful drift detection is performed on the stack.
	// *   IN_SYNC: The stack is being synchronized.
	StackDriftStatus *string `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	// The stack ID.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The stack name.
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The stack type. Valid values:
	//
	// *   ROS: ROS stack. The stack is created by using a ROS template.
	// *   Terraform: Terraform stack. The stack is created by using a Terraform template.
	StackType *string `json:"StackType,omitempty" xml:"StackType,omitempty"`
	// The state of the stack.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the stack is in its current state.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The tags of the stack.
	Tags []*ListStacksResponseBodyStacksTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period for creating the stack. Unit: minutes.
	TimeoutInMinutes *int32 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	// The time when the stack was updated. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListStacksResponseBodyStacks) String() string {
	return tea.Prettify(s)
}

func (s ListStacksResponseBodyStacks) GoString() string {
	return s.String()
}

func (s *ListStacksResponseBodyStacks) SetCreateTime(v string) *ListStacksResponseBodyStacks {
	s.CreateTime = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetDeletionProtection(v string) *ListStacksResponseBodyStacks {
	s.DeletionProtection = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetDisableRollback(v bool) *ListStacksResponseBodyStacks {
	s.DisableRollback = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetDriftDetectionTime(v string) *ListStacksResponseBodyStacks {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetOperationInfo(v *ListStacksResponseBodyStacksOperationInfo) *ListStacksResponseBodyStacks {
	s.OperationInfo = v
	return s
}

func (s *ListStacksResponseBodyStacks) SetParentStackId(v string) *ListStacksResponseBodyStacks {
	s.ParentStackId = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetRegionId(v string) *ListStacksResponseBodyStacks {
	s.RegionId = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetResourceGroupId(v string) *ListStacksResponseBodyStacks {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetServiceManaged(v bool) *ListStacksResponseBodyStacks {
	s.ServiceManaged = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetServiceName(v string) *ListStacksResponseBodyStacks {
	s.ServiceName = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackDriftStatus(v string) *ListStacksResponseBodyStacks {
	s.StackDriftStatus = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackId(v string) *ListStacksResponseBodyStacks {
	s.StackId = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackName(v string) *ListStacksResponseBodyStacks {
	s.StackName = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackType(v string) *ListStacksResponseBodyStacks {
	s.StackType = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStatus(v string) *ListStacksResponseBodyStacks {
	s.Status = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStatusReason(v string) *ListStacksResponseBodyStacks {
	s.StatusReason = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetTags(v []*ListStacksResponseBodyStacksTags) *ListStacksResponseBodyStacks {
	s.Tags = v
	return s
}

func (s *ListStacksResponseBodyStacks) SetTimeoutInMinutes(v int32) *ListStacksResponseBodyStacks {
	s.TimeoutInMinutes = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetUpdateTime(v string) *ListStacksResponseBodyStacks {
	s.UpdateTime = &v
	return s
}

type ListStacksResponseBodyStacksOperationInfo struct {
	// The name of the API operation that belongs to another Alibaba Cloud service.
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The error code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The logical ID of the resource on which the operation error occurred.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request that is initiated to call the API operation of another Alibaba Cloud service.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the resource on which the operation error occurred.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListStacksResponseBodyStacksOperationInfo) String() string {
	return tea.Prettify(s)
}

func (s ListStacksResponseBodyStacksOperationInfo) GoString() string {
	return s.String()
}

func (s *ListStacksResponseBodyStacksOperationInfo) SetAction(v string) *ListStacksResponseBodyStacksOperationInfo {
	s.Action = &v
	return s
}

func (s *ListStacksResponseBodyStacksOperationInfo) SetCode(v string) *ListStacksResponseBodyStacksOperationInfo {
	s.Code = &v
	return s
}

func (s *ListStacksResponseBodyStacksOperationInfo) SetLogicalResourceId(v string) *ListStacksResponseBodyStacksOperationInfo {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStacksResponseBodyStacksOperationInfo) SetMessage(v string) *ListStacksResponseBodyStacksOperationInfo {
	s.Message = &v
	return s
}

func (s *ListStacksResponseBodyStacksOperationInfo) SetRequestId(v string) *ListStacksResponseBodyStacksOperationInfo {
	s.RequestId = &v
	return s
}

func (s *ListStacksResponseBodyStacksOperationInfo) SetResourceType(v string) *ListStacksResponseBodyStacksOperationInfo {
	s.ResourceType = &v
	return s
}

type ListStacksResponseBodyStacksTags struct {
	// The tag key of the stack.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the stack.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListStacksResponseBodyStacksTags) String() string {
	return tea.Prettify(s)
}

func (s ListStacksResponseBodyStacksTags) GoString() string {
	return s.String()
}

func (s *ListStacksResponseBodyStacksTags) SetKey(v string) *ListStacksResponseBodyStacksTags {
	s.Key = &v
	return s
}

func (s *ListStacksResponseBodyStacksTags) SetValue(v string) *ListStacksResponseBodyStacksTags {
	s.Value = &v
	return s
}

type ListStacksResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListStacksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStacksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStacksResponse) GoString() string {
	return s.String()
}

func (s *ListStacksResponse) SetHeaders(v map[string]*string) *ListStacksResponse {
	s.Headers = v
	return s
}

func (s *ListStacksResponse) SetStatusCode(v int32) *ListStacksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStacksResponse) SetBody(v *ListStacksResponseBody) *ListStacksResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the tag key. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource. Valid values:
	//
	// *   stack: stack
	// *   stackgroup: stack group
	// *   template: template
	// *   templatescratch: scenario
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	// The tag keys.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetKeys(v []*string) *ListTagKeysResponseBody {
	s.Keys = v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the tag. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Valid values:
	//
	// *   stack: stack
	// *   stackgroup: stack group
	// *   template: template
	// *   templatescratch: scenario
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resources. You can specify up to 20 tags.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.\
	// The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.\
	// The tag value can be up to 128 characters in length, and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the tags that are added to the resources.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key of the resource.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the resource.
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the tag value. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource. Valid values:
	//
	// *   stack: stack
	// *   stackgroup: stack group
	// *   template: template
	// *   templatescratch: scenario
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

type ListTagValuesResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tag values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetValues(v []*string) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ListTemplateScratchesRequest struct {
	// The number of the page to return.
	//
	// Pages start from page 1.
	//
	// Default value: 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the scenario.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the scenario. Valid values:
	//
	// *   GENERATE_IN_PROGRESS: The scenario is being created.
	// *   GENERATE_COMPLETE: The scenario is created.
	// *   GENERATE_FAILED: The scenario fails to be created.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the scenario.
	Tags []*ListTemplateScratchesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the scenario.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The type of the scenario. Valid values:
	//
	// *   ArchitectureReplication: resource replication
	// *   ArchitectureDetection: resource detection
	// *   ResourceImport: resource management
	// *   ResourceMigration: resource migration
	TemplateScratchType *string `json:"TemplateScratchType,omitempty" xml:"TemplateScratchType,omitempty"`
}

func (s ListTemplateScratchesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateScratchesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesRequest) SetPageNumber(v int32) *ListTemplateScratchesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplateScratchesRequest) SetPageSize(v int32) *ListTemplateScratchesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplateScratchesRequest) SetRegionId(v string) *ListTemplateScratchesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTemplateScratchesRequest) SetResourceGroupId(v string) *ListTemplateScratchesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplateScratchesRequest) SetStatus(v string) *ListTemplateScratchesRequest {
	s.Status = &v
	return s
}

func (s *ListTemplateScratchesRequest) SetTags(v []*ListTemplateScratchesRequestTags) *ListTemplateScratchesRequest {
	s.Tags = v
	return s
}

func (s *ListTemplateScratchesRequest) SetTemplateScratchId(v string) *ListTemplateScratchesRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *ListTemplateScratchesRequest) SetTemplateScratchType(v string) *ListTemplateScratchesRequest {
	s.TemplateScratchType = &v
	return s
}

type ListTemplateScratchesRequestTags struct {
	// The tag key of the scenario.
	//
	// > Tags is optional. If you want to specify Tags, you must specify Key.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the scenario.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTemplateScratchesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateScratchesRequestTags) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesRequestTags) SetKey(v string) *ListTemplateScratchesRequestTags {
	s.Key = &v
	return s
}

func (s *ListTemplateScratchesRequestTags) SetValue(v string) *ListTemplateScratchesRequestTags {
	s.Value = &v
	return s
}

type ListTemplateScratchesResponseBody struct {
	// The page number of the returned page.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scenarios.
	TemplateScratches []*ListTemplateScratchesResponseBodyTemplateScratches `json:"TemplateScratches,omitempty" xml:"TemplateScratches,omitempty" type:"Repeated"`
	// The total number of scenarios.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplateScratchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateScratchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBody) SetPageNumber(v int32) *ListTemplateScratchesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTemplateScratchesResponseBody) SetPageSize(v int32) *ListTemplateScratchesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTemplateScratchesResponseBody) SetRequestId(v string) *ListTemplateScratchesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateScratchesResponseBody) SetTemplateScratches(v []*ListTemplateScratchesResponseBodyTemplateScratches) *ListTemplateScratchesResponseBody {
	s.TemplateScratches = v
	return s
}

func (s *ListTemplateScratchesResponseBody) SetTotalCount(v int32) *ListTemplateScratchesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTemplateScratchesResponseBodyTemplateScratches struct {
	// The time when the scenario was created.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the scenario.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The status code that is returned if the scenario failed to be generated.
	//
	// >  This parameter is returned only if Status is set to GENERATE_FAILED.
	FailedCode *string `json:"FailedCode,omitempty" xml:"FailedCode,omitempty"`
	// The policy based on which the logical ID is generated. Valid values:
	//
	// *   LongTypePrefixAndIndexSuffix (default): long-type prefix + index-type suffix
	// *   LongTypePrefixAndHashSuffix: long-type prefix + hash-type suffix
	// *   ShortTypePrefixAndHashSuffix: short-type prefix + hash-type suffix
	LogicalIdStrategy *string `json:"LogicalIdStrategy,omitempty" xml:"LogicalIdStrategy,omitempty"`
	// The parameters that are configured for the scenario.
	PreferenceParameters []*ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters `json:"PreferenceParameters,omitempty" xml:"PreferenceParameters,omitempty" type:"Repeated"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The source resource group.
	SourceResourceGroup *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup `json:"SourceResourceGroup,omitempty" xml:"SourceResourceGroup,omitempty" type:"Struct"`
	// The source resources.
	SourceResources []*ListTemplateScratchesResponseBodyTemplateScratchesSourceResources `json:"SourceResources,omitempty" xml:"SourceResources,omitempty" type:"Repeated"`
	// The source tag.
	SourceTag *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag `json:"SourceTag,omitempty" xml:"SourceTag,omitempty" type:"Struct"`
	// The state of the scenario.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the scenario failed to be generated.
	//
	// >  This parameter is returned only if Status is set to GENERATE_FAILED.
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The tags of the scenario.
	Tags []*ListTemplateScratchesResponseBodyTemplateScratchesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The scenario ID.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The type of the scenario. Valid values:
	//
	// *   ResourceImport: resource management
	// *   ArchitectureReplication: resource replication
	TemplateScratchType *string `json:"TemplateScratchType,omitempty" xml:"TemplateScratchType,omitempty"`
	// The time when the scenario was updated.
	//
	// The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTemplateScratchesResponseBodyTemplateScratches) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateScratchesResponseBodyTemplateScratches) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetCreateTime(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.CreateTime = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetDescription(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.Description = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetFailedCode(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.FailedCode = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetLogicalIdStrategy(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.LogicalIdStrategy = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetPreferenceParameters(v []*ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.PreferenceParameters = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetResourceGroupId(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetSourceResourceGroup(v *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.SourceResourceGroup = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetSourceResources(v []*ListTemplateScratchesResponseBodyTemplateScratchesSourceResources) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.SourceResources = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetSourceTag(v *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.SourceTag = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetStatus(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.Status = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetStatusReason(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.StatusReason = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetTags(v []*ListTemplateScratchesResponseBodyTemplateScratchesTags) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.Tags = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetTemplateScratchId(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.TemplateScratchId = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetTemplateScratchType(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.TemplateScratchType = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratches) SetUpdateTime(v string) *ListTemplateScratchesResponseBodyTemplateScratches {
	s.UpdateTime = &v
	return s
}

type ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters struct {
	// The name of the parameter.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters) SetParameterKey(v string) *ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters {
	s.ParameterKey = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters) SetParameterValue(v string) *ListTemplateScratchesResponseBodyTemplateScratchesPreferenceParameters {
	s.ParameterValue = &v
	return s
}

type ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup struct {
	// The ID of the source resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The filters for resource types.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup) SetResourceGroupId(v string) *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup) SetResourceTypeFilter(v []*string) *ListTemplateScratchesResponseBodyTemplateScratchesSourceResourceGroup {
	s.ResourceTypeFilter = v
	return s
}

type ListTemplateScratchesResponseBodyTemplateScratchesSourceResources struct {
	// The resource ID.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesSourceResources) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesSourceResources) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResources) SetResourceId(v string) *ListTemplateScratchesResponseBodyTemplateScratchesSourceResources {
	s.ResourceId = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceResources) SetResourceType(v string) *ListTemplateScratchesResponseBodyTemplateScratchesSourceResources {
	s.ResourceType = &v
	return s
}

type ListTemplateScratchesResponseBodyTemplateScratchesSourceTag struct {
	// The source tags.
	ResourceTags map[string]interface{} `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty"`
	// The filters for resource types.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesSourceTag) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesSourceTag) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag) SetResourceTags(v map[string]interface{}) *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag {
	s.ResourceTags = v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag) SetResourceTypeFilter(v []*string) *ListTemplateScratchesResponseBodyTemplateScratchesSourceTag {
	s.ResourceTypeFilter = v
	return s
}

type ListTemplateScratchesResponseBodyTemplateScratchesTags struct {
	// The tag key of the scenario.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the scenario.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesTags) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateScratchesResponseBodyTemplateScratchesTags) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesTags) SetKey(v string) *ListTemplateScratchesResponseBodyTemplateScratchesTags {
	s.Key = &v
	return s
}

func (s *ListTemplateScratchesResponseBodyTemplateScratchesTags) SetValue(v string) *ListTemplateScratchesResponseBodyTemplateScratchesTags {
	s.Value = &v
	return s
}

type ListTemplateScratchesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTemplateScratchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplateScratchesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateScratchesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateScratchesResponse) SetHeaders(v map[string]*string) *ListTemplateScratchesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateScratchesResponse) SetStatusCode(v int32) *ListTemplateScratchesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplateScratchesResponse) SetBody(v *ListTemplateScratchesResponseBody) *ListTemplateScratchesResponse {
	s.Body = v
	return s
}

type ListTemplateVersionsRequest struct {
	// The maximum number of results to be returned in a single call when NextToken is used for the query.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 50.
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The template ID. This parameter applies to shared and private templates.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTemplateVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsRequest) SetMaxResults(v int64) *ListTemplateVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetNextToken(v string) *ListTemplateVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateVersionsRequest) SetTemplateId(v string) *ListTemplateVersionsRequest {
	s.TemplateId = &v
	return s
}

type ListTemplateVersionsResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The versions.
	Versions []*ListTemplateVersionsResponseBodyVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListTemplateVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponseBody) SetNextToken(v string) *ListTemplateVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetRequestId(v string) *ListTemplateVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateVersionsResponseBody) SetVersions(v []*ListTemplateVersionsResponseBodyVersions) *ListTemplateVersionsResponseBody {
	s.Versions = v
	return s
}

type ListTemplateVersionsResponseBodyVersions struct {
	// The time when the version was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the version.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The template ID. This parameter applies to shared and private templates. For a shared template, the template ID is the same as the Alibaba Cloud Resource Name (ARN) of the template.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name that corresponds to the specified version.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The version number.
	//
	// For a shared template, this parameter is returned only if VersionOption is set to AllVersions.
	//
	// Valid values: v1 to v100.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The time when the version was last updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTemplateVersionsResponseBodyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponseBodyVersions) SetCreateTime(v string) *ListTemplateVersionsResponseBodyVersions {
	s.CreateTime = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetDescription(v string) *ListTemplateVersionsResponseBodyVersions {
	s.Description = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetTemplateId(v string) *ListTemplateVersionsResponseBodyVersions {
	s.TemplateId = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetTemplateName(v string) *ListTemplateVersionsResponseBodyVersions {
	s.TemplateName = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetTemplateVersion(v string) *ListTemplateVersionsResponseBodyVersions {
	s.TemplateVersion = &v
	return s
}

func (s *ListTemplateVersionsResponseBodyVersions) SetUpdateTime(v string) *ListTemplateVersionsResponseBodyVersions {
	s.UpdateTime = &v
	return s
}

type ListTemplateVersionsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTemplateVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplateVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateVersionsResponse) SetHeaders(v map[string]*string) *ListTemplateVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateVersionsResponse) SetStatusCode(v int32) *ListTemplateVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplateVersionsResponse) SetBody(v *ListTemplateVersionsResponseBody) *ListTemplateVersionsResponse {
	s.Body = v
	return s
}

type ListTemplatesRequest struct {
	// Specifies whether to query the tag information. Valid values:
	//
	// *   Enabled
	// *   Disabled (default)
	IncludeTags *string `json:"IncludeTags,omitempty" xml:"IncludeTags,omitempty"`
	// The page number.\
	// Pages start from page 1.\
	// Default value: 1.
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\
	// Valid values: 1 to 50.\
	// Default value: 10.
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the resource group.\
	// For more information about resource groups, see the "Resource Group" section of [What is Resource Management?](~~94475~~)
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The sharing type of the template.
	//
	// Valid values:
	//
	// *   Private (default): The template belongs to the template owner.
	// *   Shared: The template is shared with other users.
	// *   Official: The template is the shared template of the official version.
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags. You can specify up to 20 tags.
	Tag []*ListTemplatesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The template name. This parameter takes effect only when ShareType is set to Private. The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (\_). The name must start with a digit or letter.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) SetIncludeTags(v string) *ListTemplatesRequest {
	s.IncludeTags = &v
	return s
}

func (s *ListTemplatesRequest) SetPageNumber(v int64) *ListTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int64) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetResourceGroupId(v string) *ListTemplatesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplatesRequest) SetShareType(v string) *ListTemplatesRequest {
	s.ShareType = &v
	return s
}

func (s *ListTemplatesRequest) SetTag(v []*ListTemplatesRequestTag) *ListTemplatesRequest {
	s.Tag = v
	return s
}

func (s *ListTemplatesRequest) SetTemplateName(v string) *ListTemplatesRequest {
	s.TemplateName = &v
	return s
}

type ListTemplatesRequestTag struct {
	// The key of the tag. This parameter takes effect only when ShareType is set to Private.
	//
	// You can specify up to 20 tag keys.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. This parameter takes effect only when ShareType is set to Private.
	//
	// You can specify up to 20 tag values.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTemplatesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequestTag) SetKey(v string) *ListTemplatesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTemplatesRequestTag) SetValue(v string) *ListTemplatesRequestTag {
	s.Value = &v
	return s
}

type ListTemplatesResponseBody struct {
	// The page number.\
	// Pages start from page 1.
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The templates.
	Templates []*ListTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// The total number of templates.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBody) SetPageNumber(v int32) *ListTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesResponseBody) SetPageSize(v int32) *ListTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesResponseBody) SetRequestId(v string) *ListTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplatesResponseBody) SetTemplates(v []*ListTemplatesResponseBodyTemplates) *ListTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListTemplatesResponseBody) SetTotalCount(v int32) *ListTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

type ListTemplatesResponseBodyTemplates struct {
	// The time when the template was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the template.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Alibaba Cloud account to which the template belongs.
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The sharing type of the template.
	//
	// Valid values:
	//
	// *   Private: The template belongs to the template owner.
	// *   Shared: The template is shared with other users.
	ShareType *string `json:"ShareType,omitempty" xml:"ShareType,omitempty"`
	// The tags of the template.
	Tags []*ListTemplatesResponseBodyTemplatesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The Alibaba Cloud Resource Name (ARN) of the template.
	TemplateARN *string `json:"TemplateARN,omitempty" xml:"TemplateARN,omitempty"`
	// The template ID.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The latest version of the template.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The time when the template was last updated.
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplates) SetCreateTime(v string) *ListTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetDescription(v string) *ListTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetOwnerId(v string) *ListTemplatesResponseBodyTemplates {
	s.OwnerId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetResourceGroupId(v string) *ListTemplatesResponseBodyTemplates {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetShareType(v string) *ListTemplatesResponseBodyTemplates {
	s.ShareType = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTags(v []*ListTemplatesResponseBodyTemplatesTags) *ListTemplatesResponseBodyTemplates {
	s.Tags = v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateARN(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateARN = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateId(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateName(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateName = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetTemplateVersion(v string) *ListTemplatesResponseBodyTemplates {
	s.TemplateVersion = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplates) SetUpdateTime(v string) *ListTemplatesResponseBodyTemplates {
	s.UpdateTime = &v
	return s
}

type ListTemplatesResponseBodyTemplatesTags struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTemplatesResponseBodyTemplatesTags) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponseBodyTemplatesTags) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponseBodyTemplatesTags) SetKey(v string) *ListTemplatesResponseBodyTemplatesTags {
	s.Key = &v
	return s
}

func (s *ListTemplatesResponseBodyTemplatesTags) SetValue(v string) *ListTemplatesResponseBodyTemplatesTags {
	s.Value = &v
	return s
}

type ListTemplatesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListTemplatesResponse) SetHeaders(v map[string]*string) *ListTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListTemplatesResponse) SetStatusCode(v int32) *ListTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplatesResponse) SetBody(v *ListTemplatesResponseBody) *ListTemplatesResponse {
	s.Body = v
	return s
}

type MoveResourceGroupRequest struct {
	// The ID of the resource group to which you want to move the resource. For more information about resource groups, see the "Resource Group" section of the [What is Resource Management?](~~94475~~) topic.
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The region ID of the resource.\
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Valid values:
	//
	// *   stack: stack
	// *   stackgroup: stack group
	// *   template: template
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s MoveResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupRequest) SetNewResourceGroupId(v string) *MoveResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetRegionId(v string) *MoveResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceId(v string) *MoveResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveResourceGroupRequest) SetResourceType(v string) *MoveResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type MoveResourceGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponseBody) SetRequestId(v string) *MoveResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type MoveResourceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MoveResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *MoveResourceGroupResponse) SetHeaders(v map[string]*string) *MoveResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *MoveResourceGroupResponse) SetStatusCode(v int32) *MoveResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveResourceGroupResponse) SetBody(v *MoveResourceGroupResponseBody) *MoveResourceGroupResponse {
	s.Body = v
	return s
}

type PreviewStackRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\
	// The token can be up to 64 characters in length, and can contain letters, digits, underscores (\_), and hyphens (-).\
	// For more information, see [Ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to disable rollback for the resources when the stack fails to be created. Valid values:
	//
	// *   true
	// *   false (default)
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// Specifies whether to query the parameters that you want to use in compliance precheck.
	//
	// Valid values:
	//
	// *   true
	// *   false (default)
	EnablePreConfig *bool `json:"EnablePreConfig,omitempty" xml:"EnablePreConfig,omitempty"`
	// The maximum number of concurrent operations that can be performed on resources. This parameter takes effect only for Terraform stacks.
	//
	// By default, this parameter is empty. You can set this parameter to an integer that is greater than or equal to 0.
	//
	// > If you set this parameter to an integer greater than 0, the integer is used. If you set this parameter to 0 or leave this parameter empty, the default value of Terraform is used. In most cases, the default value of Terraform is 10.
	Parallelism *int64 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The parameters of the stack.
	Parameters []*PreviewStackRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The stack ID. You can use this parameter to preview a stack that you want to update.
	//
	//
	//
	// > -  You must and can specify only one of StackName and StackId.
	// > - In the scenario in which you preview a stack that you want to create or update, you cannot preview the resources in its nested stacks.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The stack name. You can use this parameter to preview the stack that you want to create. The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or letter.
	//
	// > You must and can specify only one of StackName and StackId.
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
	//
	// > You can specify only one of StackPolicyBody and StackPolicyURL.
	StackPolicyBody *string `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	// The URL of the file that contains the stack policy. The URL must point to a policy that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo or oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy file can be up to 16,384 bytes in length. If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// > You can specify only one of StackPolicyBody and StackPolicyURL.
	//
	// The URL can be up to 1,350 bytes in length.
	StackPolicyURL *string `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
	TemplateBody   *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The template ID. This parameter applies to shared and private templates.
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The scenario ID.
	//
	// For more information about how to query the scenario ID, see [ListTemplateScratches](~~363050~~).
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The region ID of the scenario. The default value is the same as the value of RegionId.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	TemplateScratchRegionId *string `json:"TemplateScratchRegionId,omitempty" xml:"TemplateScratchRegionId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an OSS bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length. If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// > You must and can specify only one of the following parameters: TemplateBody, TemplateURL, TemplateId, and TemplateScratchId.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. This parameter takes effect only when TemplateId is specified.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The timeout period for creating the stack.
	//
	// Unit: minutes.
	//
	// Default value: 60.
	TimeoutInMinutes *int64 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s PreviewStackRequest) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackRequest) GoString() string {
	return s.String()
}

func (s *PreviewStackRequest) SetClientToken(v string) *PreviewStackRequest {
	s.ClientToken = &v
	return s
}

func (s *PreviewStackRequest) SetDisableRollback(v bool) *PreviewStackRequest {
	s.DisableRollback = &v
	return s
}

func (s *PreviewStackRequest) SetEnablePreConfig(v bool) *PreviewStackRequest {
	s.EnablePreConfig = &v
	return s
}

func (s *PreviewStackRequest) SetParallelism(v int64) *PreviewStackRequest {
	s.Parallelism = &v
	return s
}

func (s *PreviewStackRequest) SetParameters(v []*PreviewStackRequestParameters) *PreviewStackRequest {
	s.Parameters = v
	return s
}

func (s *PreviewStackRequest) SetRegionId(v string) *PreviewStackRequest {
	s.RegionId = &v
	return s
}

func (s *PreviewStackRequest) SetStackId(v string) *PreviewStackRequest {
	s.StackId = &v
	return s
}

func (s *PreviewStackRequest) SetStackName(v string) *PreviewStackRequest {
	s.StackName = &v
	return s
}

func (s *PreviewStackRequest) SetStackPolicyBody(v string) *PreviewStackRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *PreviewStackRequest) SetStackPolicyURL(v string) *PreviewStackRequest {
	s.StackPolicyURL = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateBody(v string) *PreviewStackRequest {
	s.TemplateBody = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateId(v string) *PreviewStackRequest {
	s.TemplateId = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateScratchId(v string) *PreviewStackRequest {
	s.TemplateScratchId = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateScratchRegionId(v string) *PreviewStackRequest {
	s.TemplateScratchRegionId = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateURL(v string) *PreviewStackRequest {
	s.TemplateURL = &v
	return s
}

func (s *PreviewStackRequest) SetTemplateVersion(v string) *PreviewStackRequest {
	s.TemplateVersion = &v
	return s
}

func (s *PreviewStackRequest) SetTimeoutInMinutes(v int64) *PreviewStackRequest {
	s.TimeoutInMinutes = &v
	return s
}

type PreviewStackRequestParameters struct {
	// The name of the parameter N. If you do not specify the name and value of a parameter, Resource Orchestration Service (ROS) uses the default name and value that are specified in the template. Maximum value of N: 200.
	//
	// > If you specify Parameters, you must specify Parameters.N.ParameterKey.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N. Maximum value of N: 200.
	//
	// > If you specify Parameters, you must specify Parameters.N.ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s PreviewStackRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackRequestParameters) GoString() string {
	return s.String()
}

func (s *PreviewStackRequestParameters) SetParameterKey(v string) *PreviewStackRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *PreviewStackRequestParameters) SetParameterValue(v string) *PreviewStackRequestParameters {
	s.ParameterValue = &v
	return s
}

type PreviewStackResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the stack that is previewed.
	Stack *PreviewStackResponseBodyStack `json:"Stack,omitempty" xml:"Stack,omitempty" type:"Struct"`
}

func (s PreviewStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBody) SetRequestId(v string) *PreviewStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewStackResponseBody) SetStack(v *PreviewStackResponseBodyStack) *PreviewStackResponseBody {
	s.Stack = v
	return s
}

type PreviewStackResponseBodyStack struct {
	// The description of the stack.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether rollback is disabled for the resources when the stack fails to be created.
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The log that is generated when the stack is run.
	Log *PreviewStackResponseBodyStackLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Struct"`
	// The parameters of the stack.
	Parameters []*PreviewStackResponseBodyStackParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The region where the stack resides.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resources in the stack.
	Resources []*PreviewStackResponseBodyStackResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The stack name.
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The structure that contains the stack policy body.
	StackPolicyBody map[string]interface{} `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	// The description of the template.
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// The timeout period for creating the stack.
	//
	// Unit: minutes.
	TimeoutInMinutes *int32 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s PreviewStackResponseBodyStack) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackResponseBodyStack) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBodyStack) SetDescription(v string) *PreviewStackResponseBodyStack {
	s.Description = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetDisableRollback(v bool) *PreviewStackResponseBodyStack {
	s.DisableRollback = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetLog(v *PreviewStackResponseBodyStackLog) *PreviewStackResponseBodyStack {
	s.Log = v
	return s
}

func (s *PreviewStackResponseBodyStack) SetParameters(v []*PreviewStackResponseBodyStackParameters) *PreviewStackResponseBodyStack {
	s.Parameters = v
	return s
}

func (s *PreviewStackResponseBodyStack) SetRegionId(v string) *PreviewStackResponseBodyStack {
	s.RegionId = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetResources(v []*PreviewStackResponseBodyStackResources) *PreviewStackResponseBodyStack {
	s.Resources = v
	return s
}

func (s *PreviewStackResponseBodyStack) SetStackName(v string) *PreviewStackResponseBodyStack {
	s.StackName = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetStackPolicyBody(v map[string]interface{}) *PreviewStackResponseBodyStack {
	s.StackPolicyBody = v
	return s
}

func (s *PreviewStackResponseBodyStack) SetTemplateDescription(v string) *PreviewStackResponseBodyStack {
	s.TemplateDescription = &v
	return s
}

func (s *PreviewStackResponseBodyStack) SetTimeoutInMinutes(v int32) *PreviewStackResponseBodyStack {
	s.TimeoutInMinutes = &v
	return s
}

type PreviewStackResponseBodyStackLog struct {
	// The Terraform logs. This parameter is returned only if the stack is a Terraform stack.
	//
	// > This parameter contains the logs of previewing the stack.
	TerraformLogs []*PreviewStackResponseBodyStackLogTerraformLogs `json:"TerraformLogs,omitempty" xml:"TerraformLogs,omitempty" type:"Repeated"`
}

func (s PreviewStackResponseBodyStackLog) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackResponseBodyStackLog) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBodyStackLog) SetTerraformLogs(v []*PreviewStackResponseBodyStackLogTerraformLogs) *PreviewStackResponseBodyStackLog {
	s.TerraformLogs = v
	return s
}

type PreviewStackResponseBodyStackLogTerraformLogs struct {
	// The name of the Terraform command that is run. Valid values:
	//
	// *   apply
	// *   plan
	// *   destroy
	// *   version
	//
	// For more information about Terraform commands, see [Basic CLI Features](https://www.terraform.io/cli/commands).
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The content of the output stream that is returned after the command is run.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The output stream. Valid values:
	//
	// *   stdout: standard output stream
	// *   stderr: standard error stream
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s PreviewStackResponseBodyStackLogTerraformLogs) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackResponseBodyStackLogTerraformLogs) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBodyStackLogTerraformLogs) SetCommand(v string) *PreviewStackResponseBodyStackLogTerraformLogs {
	s.Command = &v
	return s
}

func (s *PreviewStackResponseBodyStackLogTerraformLogs) SetContent(v string) *PreviewStackResponseBodyStackLogTerraformLogs {
	s.Content = &v
	return s
}

func (s *PreviewStackResponseBodyStackLogTerraformLogs) SetStream(v string) *PreviewStackResponseBodyStackLogTerraformLogs {
	s.Stream = &v
	return s
}

type PreviewStackResponseBodyStackParameters struct {
	// The name of the parameter.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s PreviewStackResponseBodyStackParameters) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackResponseBodyStackParameters) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBodyStackParameters) SetParameterKey(v string) *PreviewStackResponseBodyStackParameters {
	s.ParameterKey = &v
	return s
}

func (s *PreviewStackResponseBodyStackParameters) SetParameterValue(v string) *PreviewStackResponseBodyStackParameters {
	s.ParameterValue = &v
	return s
}

type PreviewStackResponseBodyStackResources struct {
	// The resource type of an Alibaba Cloud service.
	AcsResourceType *string `json:"AcsResourceType,omitempty" xml:"AcsResourceType,omitempty"`
	// The action that is performed on the resource. Valid values:
	//
	// *   Add
	// *   Modify
	// *   Remove
	// *   None
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The description of the resource.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The logical ID of the resource.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The physical ID of the resource.
	//
	// This parameter is returned only if Action is set to Modify or Remove.
	PhysicalResourceId *string `json:"PhysicalResourceId,omitempty" xml:"PhysicalResourceId,omitempty"`
	// The resource properties.
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// Indicates whether a replacement update is performed on the template. Valid values:
	//
	// *   True: A replacement update is performed on the template.
	// *   False: A change is made on the template.
	// *   Conditional: A replacement update may be performed on the template. You can check whether a replacement update is performed when the template is in use.
	//
	// > This parameter is returned only if Action is set to Modify.
	Replacement *string `json:"Replacement,omitempty" xml:"Replacement,omitempty"`
	// The resources on which the stack depends.
	RequiredBy []*string `json:"RequiredBy,omitempty" xml:"RequiredBy,omitempty" type:"Repeated"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The information about the nested stack. The data structure of the value is the same as the data structure of the entire response.
	Stack map[string]interface{} `json:"Stack,omitempty" xml:"Stack,omitempty"`
}

func (s PreviewStackResponseBodyStackResources) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackResponseBodyStackResources) GoString() string {
	return s.String()
}

func (s *PreviewStackResponseBodyStackResources) SetAcsResourceType(v string) *PreviewStackResponseBodyStackResources {
	s.AcsResourceType = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetAction(v string) *PreviewStackResponseBodyStackResources {
	s.Action = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetDescription(v string) *PreviewStackResponseBodyStackResources {
	s.Description = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetLogicalResourceId(v string) *PreviewStackResponseBodyStackResources {
	s.LogicalResourceId = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetPhysicalResourceId(v string) *PreviewStackResponseBodyStackResources {
	s.PhysicalResourceId = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetProperties(v map[string]interface{}) *PreviewStackResponseBodyStackResources {
	s.Properties = v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetReplacement(v string) *PreviewStackResponseBodyStackResources {
	s.Replacement = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetRequiredBy(v []*string) *PreviewStackResponseBodyStackResources {
	s.RequiredBy = v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetResourceType(v string) *PreviewStackResponseBodyStackResources {
	s.ResourceType = &v
	return s
}

func (s *PreviewStackResponseBodyStackResources) SetStack(v map[string]interface{}) *PreviewStackResponseBodyStackResources {
	s.Stack = v
	return s
}

type PreviewStackResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PreviewStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PreviewStackResponse) String() string {
	return tea.Prettify(s)
}

func (s PreviewStackResponse) GoString() string {
	return s.String()
}

func (s *PreviewStackResponse) SetHeaders(v map[string]*string) *PreviewStackResponse {
	s.Headers = v
	return s
}

func (s *PreviewStackResponse) SetStatusCode(v int32) *PreviewStackResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewStackResponse) SetBody(v *PreviewStackResponseBody) *PreviewStackResponse {
	s.Body = v
	return s
}

type RegisterResourceTypeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\
	// The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (\_).\
	// For more information, see [Ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the resource type. The description can be up to 512 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The entity type. Set the value to Module.
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. The template body is used as the module content. If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.
	//
	//
	// > - This parameter takes effect only when EntityType is set to Module.
	// > - You can specify only one of the TemplateBody and TemplateURL parameters.
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length. The template body is used as the module content.
	//
	// > - If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	// > -  This parameter takes effect only when EntityType is set to Module.
	// > -  You can specify only one of the TemplateBody and TemplateURL parameters.
	//
	// The URL can be up to 1,024 bytes in length.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
}

func (s RegisterResourceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *RegisterResourceTypeRequest) SetClientToken(v string) *RegisterResourceTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *RegisterResourceTypeRequest) SetDescription(v string) *RegisterResourceTypeRequest {
	s.Description = &v
	return s
}

func (s *RegisterResourceTypeRequest) SetEntityType(v string) *RegisterResourceTypeRequest {
	s.EntityType = &v
	return s
}

func (s *RegisterResourceTypeRequest) SetResourceType(v string) *RegisterResourceTypeRequest {
	s.ResourceType = &v
	return s
}

func (s *RegisterResourceTypeRequest) SetTemplateBody(v string) *RegisterResourceTypeRequest {
	s.TemplateBody = &v
	return s
}

func (s *RegisterResourceTypeRequest) SetTemplateURL(v string) *RegisterResourceTypeRequest {
	s.TemplateURL = &v
	return s
}

type RegisterResourceTypeResponseBody struct {
	// The ID of the registration record. You can call the [ListResourceTypeRegistrations](~~2330740~~) operation to query registration records.
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterResourceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterResourceTypeResponseBody) SetRegistrationId(v string) *RegisterResourceTypeResponseBody {
	s.RegistrationId = &v
	return s
}

func (s *RegisterResourceTypeResponseBody) SetRequestId(v string) *RegisterResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

type RegisterResourceTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RegisterResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterResourceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *RegisterResourceTypeResponse) SetHeaders(v map[string]*string) *RegisterResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *RegisterResourceTypeResponse) SetStatusCode(v int32) *RegisterResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterResourceTypeResponse) SetBody(v *RegisterResourceTypeResponseBody) *RegisterResourceTypeResponse {
	s.Body = v
	return s
}

type SetDeletionProtectionRequest struct {
	// Indicates whether stack deletion protection is enabled. Valid values:
	//
	// *   Enabled: enables the stack deletion protection.
	// *   Disabled (default): Resource stack deletion protection is Disabled. You can use the console or API(DeleteStack) to release the stack resources.
	//
	// >  The deletion of nested stacks is the same as the root stack.
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	//
	// The delete protection attribute of a nested stack is determined by the root stack and remains unchanged from the root stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s SetDeletionProtectionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDeletionProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetDeletionProtectionRequest) SetDeletionProtection(v string) *SetDeletionProtectionRequest {
	s.DeletionProtection = &v
	return s
}

func (s *SetDeletionProtectionRequest) SetRegionId(v string) *SetDeletionProtectionRequest {
	s.RegionId = &v
	return s
}

func (s *SetDeletionProtectionRequest) SetStackId(v string) *SetDeletionProtectionRequest {
	s.StackId = &v
	return s
}

type SetDeletionProtectionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDeletionProtectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *SetDeletionProtectionResponseBody) SetRequestId(v string) *SetDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

type SetDeletionProtectionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetDeletionProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetDeletionProtectionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDeletionProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetDeletionProtectionResponse) SetHeaders(v map[string]*string) *SetDeletionProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetDeletionProtectionResponse) SetStatusCode(v int32) *SetDeletionProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDeletionProtectionResponse) SetBody(v *SetDeletionProtectionResponseBody) *SetDeletionProtectionResponse {
	s.Body = v
	return s
}

type SetResourceTypeRequest struct {
	// The ID of the default version. You can use this parameter to specify the default version of the resource type.
	//
	// > You can specify only one of the VersionId and DefaultVersionId parameters.
	DefaultVersionId *string `json:"DefaultVersionId,omitempty" xml:"DefaultVersionId,omitempty"`
	// The description of the resource type or resource type version. The description can be up to 512 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The version ID. If you want to modify a version of the resource type, you must specify this parameter. If you do not specify this parameter, only the resource type is modified.
	//
	// > You can specify only one of the VersionId and DefaultVersionId parameters.
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s SetResourceTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetResourceTypeRequest) GoString() string {
	return s.String()
}

func (s *SetResourceTypeRequest) SetDefaultVersionId(v string) *SetResourceTypeRequest {
	s.DefaultVersionId = &v
	return s
}

func (s *SetResourceTypeRequest) SetDescription(v string) *SetResourceTypeRequest {
	s.Description = &v
	return s
}

func (s *SetResourceTypeRequest) SetResourceType(v string) *SetResourceTypeRequest {
	s.ResourceType = &v
	return s
}

func (s *SetResourceTypeRequest) SetVersionId(v string) *SetResourceTypeRequest {
	s.VersionId = &v
	return s
}

type SetResourceTypeResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetResourceTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *SetResourceTypeResponseBody) SetRequestId(v string) *SetResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

type SetResourceTypeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetResourceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *SetResourceTypeResponse) SetHeaders(v map[string]*string) *SetResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *SetResourceTypeResponse) SetStatusCode(v int32) *SetResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetResourceTypeResponse) SetBody(v *SetResourceTypeResponseBody) *SetResourceTypeResponse {
	s.Body = v
	return s
}

type SetStackPolicyRequest struct {
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The structure that contains the stack policy body. The stack policy body must be 1 to 16,384 bytes in length.
	//
	// You can specify one of the StackPolicyBody and StackPolicyURL parameters, but you cannot specify both of them.
	StackPolicyBody *string `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	// The URL for the file that contains the stack policy. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template can be up to 16,384 bytes in length, and the URL can be up to 1,350 bytes in length.
	//
	// >  If the region of the OSS bucket is not specified, the RegionId value is used.
	//
	// You can specify one of the StackPolicyBody and StackPolicyURL parameters, but you cannot specify both of them.
	StackPolicyURL *string `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
}

func (s SetStackPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetStackPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetStackPolicyRequest) SetRegionId(v string) *SetStackPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *SetStackPolicyRequest) SetStackId(v string) *SetStackPolicyRequest {
	s.StackId = &v
	return s
}

func (s *SetStackPolicyRequest) SetStackPolicyBody(v string) *SetStackPolicyRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *SetStackPolicyRequest) SetStackPolicyURL(v string) *SetStackPolicyRequest {
	s.StackPolicyURL = &v
	return s
}

type SetStackPolicyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetStackPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetStackPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetStackPolicyResponseBody) SetRequestId(v string) *SetStackPolicyResponseBody {
	s.RequestId = &v
	return s
}

type SetStackPolicyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetStackPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetStackPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetStackPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetStackPolicyResponse) SetHeaders(v map[string]*string) *SetStackPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetStackPolicyResponse) SetStatusCode(v int32) *SetStackPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetStackPolicyResponse) SetBody(v *SetStackPolicyResponseBody) *SetStackPolicyResponse {
	s.Body = v
	return s
}

type SetTemplatePermissionRequest struct {
	// The Alibaba Cloud accounts with or from which you want to share or unshare the template.\
	// Valid values of N: 1, 2, 3, 4, and 5.
	//
	// > - This parameter cannot be set to the ID of the Alibaba Cloud account that owns the template, or the RAM users of this Alibaba Cloud account.
	// > - When ShareOption is set to CancelSharing, you can unshare the template from all the specified Alibaba Cloud accounts by using an asterisk (\*).
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The sharing option.
	//
	// Valid values:
	//
	// *   ShareToAccounts: shares the template with other Alibaba Cloud accounts.
	// *   CancelSharing: unshares the template.
	ShareOption *string `json:"ShareOption,omitempty" xml:"ShareOption,omitempty"`
	// The ID of the template.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The version of the shared template. This parameter takes effect only if you set ShareOption to ShareToAccounts and set VersionOption to Specified.
	//
	// Valid values: v1 to v100.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The version option for the shared template. This parameter takes effect only if you set ShareOption to ShareToAccounts.
	//
	// Valid values:
	//
	// *   AllVersions (default): shares all versions of the template.
	// *   Latest: shares only the latest version of template. When the version of the template is updated, ROS updates the shared version to the latest version.
	// *   Current: shares only the current version of the template. When the version of the template is updated, ROS does not update the shared version.
	// *   Specified: shares only the specified version of the template.
	VersionOption *string `json:"VersionOption,omitempty" xml:"VersionOption,omitempty"`
}

func (s SetTemplatePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s SetTemplatePermissionRequest) GoString() string {
	return s.String()
}

func (s *SetTemplatePermissionRequest) SetAccountIds(v []*string) *SetTemplatePermissionRequest {
	s.AccountIds = v
	return s
}

func (s *SetTemplatePermissionRequest) SetShareOption(v string) *SetTemplatePermissionRequest {
	s.ShareOption = &v
	return s
}

func (s *SetTemplatePermissionRequest) SetTemplateId(v string) *SetTemplatePermissionRequest {
	s.TemplateId = &v
	return s
}

func (s *SetTemplatePermissionRequest) SetTemplateVersion(v string) *SetTemplatePermissionRequest {
	s.TemplateVersion = &v
	return s
}

func (s *SetTemplatePermissionRequest) SetVersionOption(v string) *SetTemplatePermissionRequest {
	s.VersionOption = &v
	return s
}

type SetTemplatePermissionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetTemplatePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetTemplatePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *SetTemplatePermissionResponseBody) SetRequestId(v string) *SetTemplatePermissionResponseBody {
	s.RequestId = &v
	return s
}

type SetTemplatePermissionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SetTemplatePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetTemplatePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s SetTemplatePermissionResponse) GoString() string {
	return s.String()
}

func (s *SetTemplatePermissionResponse) SetHeaders(v map[string]*string) *SetTemplatePermissionResponse {
	s.Headers = v
	return s
}

func (s *SetTemplatePermissionResponse) SetStatusCode(v int32) *SetTemplatePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetTemplatePermissionResponse) SetBody(v *SetTemplatePermissionResponseBody) *SetTemplatePermissionResponse {
	s.Body = v
	return s
}

type SignalResourceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// The token can be up to 64 characters in length and can contain letters, digits, hyphens (-) and underscores (\_).
	//
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The logical ID of the resource as defined in the template.
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The status of the signal. Failure signals can cause stack creation or update to fail. If all signals are warnings, the stack cannot be created or updated. Valid values:
	//
	// *   SUCCESS
	// *   FAILURE
	// *   WARNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The unique ID of the signal. The ID must be 1 to 64 characters in length. If multiple signals are sent to a single resource, each signal must have a unique ID.
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s SignalResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s SignalResourceRequest) GoString() string {
	return s.String()
}

func (s *SignalResourceRequest) SetClientToken(v string) *SignalResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *SignalResourceRequest) SetLogicalResourceId(v string) *SignalResourceRequest {
	s.LogicalResourceId = &v
	return s
}

func (s *SignalResourceRequest) SetRegionId(v string) *SignalResourceRequest {
	s.RegionId = &v
	return s
}

func (s *SignalResourceRequest) SetStackId(v string) *SignalResourceRequest {
	s.StackId = &v
	return s
}

func (s *SignalResourceRequest) SetStatus(v string) *SignalResourceRequest {
	s.Status = &v
	return s
}

func (s *SignalResourceRequest) SetUniqueId(v string) *SignalResourceRequest {
	s.UniqueId = &v
	return s
}

type SignalResourceResponseBody struct {
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SignalResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SignalResourceResponseBody) GoString() string {
	return s.String()
}

func (s *SignalResourceResponseBody) SetRequestId(v string) *SignalResourceResponseBody {
	s.RequestId = &v
	return s
}

type SignalResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SignalResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SignalResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s SignalResourceResponse) GoString() string {
	return s.String()
}

func (s *SignalResourceResponse) SetHeaders(v map[string]*string) *SignalResourceResponse {
	s.Headers = v
	return s
}

func (s *SignalResourceResponse) SetStatusCode(v int32) *SignalResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *SignalResourceResponse) SetBody(v *SignalResourceResponseBody) *SignalResourceResponse {
	s.Body = v
	return s
}

type StopStackGroupOperationRequest struct {
	// The ID of the operation.
	//
	// You can call the [ListStackGroupOperations](~~151342~~) operation to obtain the operation ID.
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopStackGroupOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s StopStackGroupOperationRequest) GoString() string {
	return s.String()
}

func (s *StopStackGroupOperationRequest) SetOperationId(v string) *StopStackGroupOperationRequest {
	s.OperationId = &v
	return s
}

func (s *StopStackGroupOperationRequest) SetRegionId(v string) *StopStackGroupOperationRequest {
	s.RegionId = &v
	return s
}

type StopStackGroupOperationResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopStackGroupOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopStackGroupOperationResponseBody) GoString() string {
	return s.String()
}

func (s *StopStackGroupOperationResponseBody) SetRequestId(v string) *StopStackGroupOperationResponseBody {
	s.RequestId = &v
	return s
}

type StopStackGroupOperationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopStackGroupOperationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopStackGroupOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s StopStackGroupOperationResponse) GoString() string {
	return s.String()
}

func (s *StopStackGroupOperationResponse) SetHeaders(v map[string]*string) *StopStackGroupOperationResponse {
	s.Headers = v
	return s
}

func (s *StopStackGroupOperationResponse) SetStatusCode(v int32) *StopStackGroupOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *StopStackGroupOperationResponse) SetBody(v *StopStackGroupOperationResponseBody) *StopStackGroupOperationResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The region ID of the tag that you want to create. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resources.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Valid values:
	//
	// *   stack: stack
	// *   stackgroup: stack group
	// *   template: template
	// *   templatescratch: scenario
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags of the resource. You can specify up to 20 tags.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys.
	//
	// The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `aliyun` or `acs:`.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the resource. This parameter takes effect when TagKey is not specified in the request. Valid values:
	//
	// *   true
	// *   false (default)
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The region ID of the tag. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the resource.
	//
	// > If you set ResourceType to stackgroup, you must set ResourceId to the name of the stack group.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The type of the resource. Valid values:
	//
	// *   stack: stack
	// *   stackgroup: stack group
	// *   template: template
	// *   templatescratch: scenario
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag keys of the resource. You can specify up to 20 tag keys.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateStackRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests.
	//
	// The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// For more information, see [Ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to roll back the resources in the stack when the stack fails to be updated.
	//
	// Default value: false. Valid values:
	//
	// *   true
	// *   false
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// Specifies whether only to validate the stack in the request. Default value: false. Valid values:
	//
	// *   true: only validates the stack.
	// *   false: validates and updates the stack.
	//
	// >  When no changes are made during the update, the following rules apply: If you set the DryRun parameter to false, the NotSupported error code is returned. If you set the DryRun parameter to true, no error is reported.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The dry run option in the list format. You can specify only one dry run option.
	//
	// > This parameter takes effect only when DryRun is set to true.
	DryRunOptions []*string `json:"DryRunOptions,omitempty" xml:"DryRunOptions,omitempty" type:"Repeated"`
	// The maximum number of concurrent operations that can be performed on resources.
	//
	// By default, this parameter is empty. You can set this parameter to an integer that is greater than or equal to 0.
	//
	// > - If you set this parameter to an integer that is greater than 0, the integer is used.
	// > -  If you set this parameter to 0, no limit is imposed on Resource Orchestration Service (ROS) stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	// > -  If you leave this parameter empty, the value that you specified for this parameter in the previous request is used. If you left this parameter empty in the previous request, no limit is imposed on ROS stacks. However, the default value in Terraform is used for Terraform stacks. In most cases, the default value in Terraform is 10.
	// > - If you set this parameter to a specific value, ROS associates the value with the stack. The value affects subsequent operations on the stack.
	Parallelism *int64 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The parameters.
	Parameters []*UpdateStackRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The name of the RAM role. Resource Orchestration Service (ROS) assumes the RAM role to create the stack and uses the credentials of the role to call the APIs of Alibaba Cloud services.
	//
	// ROS assumes the RAM role to perform operations on the stack. If you have permissions to perform operations on the stack but do not have permissions to use the RAM role, ROS still assumes the RAM role. You must make sure that the least privileges are granted to the RAM role.
	//
	// If you do not specify this parameter, ROS assumes the existing RAM role that is associated with the stack. If no RAM roles are available, ROS uses a temporary credential that is generated from the credentials of your account.
	//
	// The name of the RAM role can be up to 64 bytes in length.
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the region in which the stack is deployed. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to enable the replacement update feature. If you cannot change resource properties, you can enable the replacement update feature to replace the resource properties. If the replacement update feature is used, the existing resource is deleted and a new resource is created. The physical ID of the new resource is different from the physical ID of the deleted resource.
	//
	// Default value: Disabled. Valid values:
	//
	// *   Enabled
	// *   Disabled
	//
	// >  Changes have higher priorities than replacement updates.
	ReplacementOption *string `json:"ReplacementOption,omitempty" xml:"ReplacementOption,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The structure that contains the stack policy body. The policy body must be 1 to 16,384 bytes in length.
	//
	// >  You can specify only one of the StackPolicyBody and StackPolicyURL parameters.
	StackPolicyBody *string `json:"StackPolicyBody,omitempty" xml:"StackPolicyBody,omitempty"`
	// The structure that contains the body of the temporary overriding stack policy. The policy body must be 1 to 16,384 bytes in length.
	//
	// If you want to update protected resources, you must specify a temporary overriding stack policy during the update. If you do not specify a temporary overriding stack policy, the existing policy that is associated with the stack is used.
	//
	// This parameter takes effect only when the ChangeSetType parameter is set to UPDATE. You can specify only one of the following parameters:
	//
	// *   StackPolicyBody
	// *   StackPolicyURL
	// *   StackPolicyDuringUpdateBody
	// *   StackPolicyDuringUpdateURL
	StackPolicyDuringUpdateBody *string `json:"StackPolicyDuringUpdateBody,omitempty" xml:"StackPolicyDuringUpdateBody,omitempty"`
	// The URL of the file that contains the temporary overriding stack policy. The URL must point to a policy that is located on an HTTP or HTTPS web server or in an OSS bucket, such as oss://ros/stack-policy/demo or oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy file can be up to 16,384 bytes in length.
	//
	// >  If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// The URL can be up to 1,350 bytes in length.
	//
	// If you want to update protected resources, you must specify a temporary overriding stack policy during the update. If you do not specify a temporary overriding stack policy, the existing policy that is associated with the stack is used. This parameter takes effect only when the ChangeSetType parameter is set to UPDATE. You can specify only one of the following parameters:
	//
	// *   StackPolicyBody
	// *   StackPolicyURL
	// *   StackPolicyDuringUpdateBody
	// *   StackPolicyDuringUpdateURL
	StackPolicyDuringUpdateURL *string `json:"StackPolicyDuringUpdateURL,omitempty" xml:"StackPolicyDuringUpdateURL,omitempty"`
	// The URL of the file that contains the stack policy. The URL must point to a policy that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket, such as oss://ros/stack-policy/demo or oss://ros/stack-policy/demo?RegionId=cn-hangzhou. The policy file can be up to 16,384 bytes in length. If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// >  You can specify only one of the StackPolicyBody and StackPolicyURL parameters.
	//
	// The URL can be up to 1,350 bytes in length.
	StackPolicyURL *string `json:"StackPolicyURL,omitempty" xml:"StackPolicyURL,omitempty"`
	// The value of tag N that you want to add to the template.
	Tags         []*UpdateStackRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TemplateBody *string                   `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared templates and private templates.
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an OSS bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body must be 1 to 524,288 bytes in length. If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// >  You must specify only one of the following parameters: TemplateBody, TemplateURL, and TemplateId.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. This parameter takes effect only when the TemplateId parameter is specified.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The timeout period for the update operation on the stack.
	//
	// *   Default value: 60.
	// *   Unit: minutes.
	TimeoutInMinutes *int64 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	// Specifies whether to use the values specified in the previous request for the parameters that you do not specify in the current request.
	//
	// Valid values:
	//
	// *   true
	// *   false
	UsePreviousParameters *bool `json:"UsePreviousParameters,omitempty" xml:"UsePreviousParameters,omitempty"`
}

func (s UpdateStackRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackRequest) SetClientToken(v string) *UpdateStackRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackRequest) SetDisableRollback(v bool) *UpdateStackRequest {
	s.DisableRollback = &v
	return s
}

func (s *UpdateStackRequest) SetDryRun(v bool) *UpdateStackRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateStackRequest) SetDryRunOptions(v []*string) *UpdateStackRequest {
	s.DryRunOptions = v
	return s
}

func (s *UpdateStackRequest) SetParallelism(v int64) *UpdateStackRequest {
	s.Parallelism = &v
	return s
}

func (s *UpdateStackRequest) SetParameters(v []*UpdateStackRequestParameters) *UpdateStackRequest {
	s.Parameters = v
	return s
}

func (s *UpdateStackRequest) SetRamRoleName(v string) *UpdateStackRequest {
	s.RamRoleName = &v
	return s
}

func (s *UpdateStackRequest) SetRegionId(v string) *UpdateStackRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackRequest) SetReplacementOption(v string) *UpdateStackRequest {
	s.ReplacementOption = &v
	return s
}

func (s *UpdateStackRequest) SetResourceGroupId(v string) *UpdateStackRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateStackRequest) SetStackId(v string) *UpdateStackRequest {
	s.StackId = &v
	return s
}

func (s *UpdateStackRequest) SetStackPolicyBody(v string) *UpdateStackRequest {
	s.StackPolicyBody = &v
	return s
}

func (s *UpdateStackRequest) SetStackPolicyDuringUpdateBody(v string) *UpdateStackRequest {
	s.StackPolicyDuringUpdateBody = &v
	return s
}

func (s *UpdateStackRequest) SetStackPolicyDuringUpdateURL(v string) *UpdateStackRequest {
	s.StackPolicyDuringUpdateURL = &v
	return s
}

func (s *UpdateStackRequest) SetStackPolicyURL(v string) *UpdateStackRequest {
	s.StackPolicyURL = &v
	return s
}

func (s *UpdateStackRequest) SetTags(v []*UpdateStackRequestTags) *UpdateStackRequest {
	s.Tags = v
	return s
}

func (s *UpdateStackRequest) SetTemplateBody(v string) *UpdateStackRequest {
	s.TemplateBody = &v
	return s
}

func (s *UpdateStackRequest) SetTemplateId(v string) *UpdateStackRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateStackRequest) SetTemplateURL(v string) *UpdateStackRequest {
	s.TemplateURL = &v
	return s
}

func (s *UpdateStackRequest) SetTemplateVersion(v string) *UpdateStackRequest {
	s.TemplateVersion = &v
	return s
}

func (s *UpdateStackRequest) SetTimeoutInMinutes(v int64) *UpdateStackRequest {
	s.TimeoutInMinutes = &v
	return s
}

func (s *UpdateStackRequest) SetUsePreviousParameters(v bool) *UpdateStackRequest {
	s.UsePreviousParameters = &v
	return s
}

type UpdateStackRequestParameters struct {
	// The name of parameter N. If you do not specify the name and value of a parameter, ROS uses the default name and value in the template.
	//
	// Maximum value of N: 200.
	//
	// >  The Parameters parameter is optional. If you specify Parameters, you must specify both Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N. Maximum value of N: 200.
	//
	// >  The Parameters parameter is optional. If you specify Parameters, you must specify both Parameters.N.ParameterKey and Parameters.N.ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateStackRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateStackRequestParameters) SetParameterKey(v string) *UpdateStackRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackRequestParameters) SetParameterValue(v string) *UpdateStackRequestParameters {
	s.ParameterValue = &v
	return s
}

type UpdateStackRequestTags struct {
	// The key of tag N that you want to add to the stack.
	//
	// Valid values of N: 1 to 20.
	//
	// > - The Tags parameter is optional. If you specify Tags, you must specify Tags.N.Key.
	// > - The tag of a stack is propagated to each resource that supports the tag feature in the stack. For more information, see [Propagate tags](~~201421~~).
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that you want to add to the stack.
	//
	// Valid values of N: 1 to 20.
	//
	// >  The tag of a stack is propagated to each resource that supports the tag feature in the stack. For more information, see [Propagate tags](~~201421~~).
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateStackRequestTags) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackRequestTags) GoString() string {
	return s.String()
}

func (s *UpdateStackRequestTags) SetKey(v string) *UpdateStackRequestTags {
	s.Key = &v
	return s
}

func (s *UpdateStackRequestTags) SetValue(v string) *UpdateStackRequestTags {
	s.Value = &v
	return s
}

type UpdateStackResponseBody struct {
	// The dry run result. This parameter is returned only if DryRun is set to true.
	DryRunResult *UpdateStackResponseBodyDryRunResult `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty" type:"Struct"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s UpdateStackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStackResponseBody) SetDryRunResult(v *UpdateStackResponseBodyDryRunResult) *UpdateStackResponseBody {
	s.DryRunResult = v
	return s
}

func (s *UpdateStackResponseBody) SetRequestId(v string) *UpdateStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStackResponseBody) SetStackId(v string) *UpdateStackResponseBody {
	s.StackId = &v
	return s
}

type UpdateStackResponseBodyDryRunResult struct {
	// The parameters that can be modified. If you change only values of the parameters in a stack template and use the template to update the stack, no validation errors are caused.
	ParametersAllowedToBeModified []*string `json:"ParametersAllowedToBeModified,omitempty" xml:"ParametersAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters whose changes cause service interruptions.
	// > - This parameter is supported only for a small number of resource types.
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersCauseInterruptionIfModified []*string `json:"ParametersCauseInterruptionIfModified,omitempty" xml:"ParametersCauseInterruptionIfModified,omitempty" type:"Repeated"`
	// The parameters whose changes trigger replacement updates for resources.
	//
	// > -  This parameter can be returned only if ReplacementOption is set to Enabled.
	// > -  This parameter is valid only for updates on ROS stacks.
	ParametersCauseReplacementIfModified []*string `json:"ParametersCauseReplacementIfModified,omitempty" xml:"ParametersCauseReplacementIfModified,omitempty" type:"Repeated"`
	// The parameters that can be modified under specific conditions. If you change only values of the parameters in a stack template and use the template to update the stack, the new values of the parameters determine whether validation errors are caused.
	ParametersConditionallyAllowedToBeModified []*string `json:"ParametersConditionallyAllowedToBeModified,omitempty" xml:"ParametersConditionallyAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters whose changes cause service interruptions under specific conditions.
	//
	// > - This parameter is supported only for a small number of resource types.
	// > -  This parameter is valid only for updates on ROS stacks.
	ParametersConditionallyCauseInterruptionIfModified []*string `json:"ParametersConditionallyCauseInterruptionIfModified,omitempty" xml:"ParametersConditionallyCauseInterruptionIfModified,omitempty" type:"Repeated"`
	// The parameters whose changes trigger replacement updates for resources under specific conditions.
	//
	// > - This parameter can be returned only if ReplacementOption is set to Enabled.
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersConditionallyCauseReplacementIfModified []*string `json:"ParametersConditionallyCauseReplacementIfModified,omitempty" xml:"ParametersConditionallyCauseReplacementIfModified,omitempty" type:"Repeated"`
	// The parameters that cannot be modified. If you change only values of the parameters in a stack template and use the template to update the stack, validation errors are caused.
	ParametersNotAllowedToBeModified []*string `json:"ParametersNotAllowedToBeModified,omitempty" xml:"ParametersNotAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that can be modified under uncertain conditions. If you change only values of the parameters in a stack template and use the template to update the stack, the actual running environment determines whether validation errors are caused.
	ParametersUncertainlyAllowedToBeModified []*string `json:"ParametersUncertainlyAllowedToBeModified,omitempty" xml:"ParametersUncertainlyAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters whose changes cause service interruptions under uncertain conditions.
	//
	// > - This parameter is supported only for a small number of resource types.
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersUncertainlyCauseInterruptionIfModified []*string `json:"ParametersUncertainlyCauseInterruptionIfModified,omitempty" xml:"ParametersUncertainlyCauseInterruptionIfModified,omitempty" type:"Repeated"`
	// The parameters whose changes trigger replacement updates for resources under uncertain conditions.
	//
	// > - This parameter can be returned only if ReplacementOption is set to Enabled.
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersUncertainlyCauseReplacementIfModified []*string `json:"ParametersUncertainlyCauseReplacementIfModified,omitempty" xml:"ParametersUncertainlyCauseReplacementIfModified,omitempty" type:"Repeated"`
}

func (s UpdateStackResponseBodyDryRunResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackResponseBodyDryRunResult) GoString() string {
	return s.String()
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersAllowedToBeModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersAllowedToBeModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersCauseInterruptionIfModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersCauseInterruptionIfModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersCauseReplacementIfModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersCauseReplacementIfModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersConditionallyAllowedToBeModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersConditionallyAllowedToBeModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersConditionallyCauseInterruptionIfModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersConditionallyCauseInterruptionIfModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersConditionallyCauseReplacementIfModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersConditionallyCauseReplacementIfModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersNotAllowedToBeModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersNotAllowedToBeModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersUncertainlyAllowedToBeModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersUncertainlyAllowedToBeModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersUncertainlyCauseInterruptionIfModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersUncertainlyCauseInterruptionIfModified = v
	return s
}

func (s *UpdateStackResponseBodyDryRunResult) SetParametersUncertainlyCauseReplacementIfModified(v []*string) *UpdateStackResponseBodyDryRunResult {
	s.ParametersUncertainlyCauseReplacementIfModified = v
	return s
}

type UpdateStackResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateStackResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateStackResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackResponse) GoString() string {
	return s.String()
}

func (s *UpdateStackResponse) SetHeaders(v map[string]*string) *UpdateStackResponse {
	s.Headers = v
	return s
}

func (s *UpdateStackResponse) SetStatusCode(v int32) *UpdateStackResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStackResponse) SetBody(v *UpdateStackResponseBody) *UpdateStackResponse {
	s.Body = v
	return s
}

type UpdateStackGroupRequest struct {
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Alibaba Cloud Object Storage Service (OSS) bucket. The template body must be 1 to 524,288 bytes in length. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// >  You must specify only one of the TemplateBody, TemplateURL, and TemplateId parameters.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The key of parameter N. If you do not specify the key and value of the parameter, ROS uses the default key and value in the template.
	//
	// Maximum value of N: 200.
	//
	// >  The Parameters parameter is optional. If you set the Parameters parameter, you must set the Parameters.N.ParameterKey parameter.
	AdministrationRoleName *string `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	// The IDs of the folders in the resource directory. You can specify up to five folder IDs.
	//
	// You can create stacks within all members in the specified folders. If you create stacks in the Root folder, the stacks are created within all members in the resource directory.
	//
	// >  To view the folder IDs, go to the **Overview** page in the **Resource Management** console. For more information, see [View the basic information of a folder](~~111223~~).
	AutoDeployment *UpdateStackGroupRequestAutoDeployment `json:"AutoDeployment,omitempty" xml:"AutoDeployment,omitempty" type:"Struct"`
	// The option for the stack group. You can specify up to one option.
	Capabilities []*string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// >  You must specify only one of the TemplateBody, TemplateURL, and TemplateId parameters.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the request.
	DeploymentTargets *UpdateStackGroupRequestDeploymentTargets `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty" type:"Struct"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.
	//
	// >  You must specify only one of the TemplateBody, TemplateURL, and TemplateId parameters.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The value of parameter N.
	//
	// Maximum value of N: 200.
	//
	// >  The Parameters parameter is optional. If you set the Parameters parameter, you must set the Parameters.N.ParameterValue parameter.
	ExecutionRoleName *string `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	// The version of the template. If you do not specify a version, the latest version is used.
	//
	// >  This parameter takes effect only if the TemplateId parameter is set.
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The list of parameters.
	OperationPreferences map[string]interface{} `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// Specifies whether to enable automatic deployment.
	//
	// Valid values:
	//
	// *   true: enables automatic deployment. If you add a member to the folder to which the stack group belongs after you enable automatic deployment, the stack group deploys its stack instances within the member. If you remove a member from the folder, the stack group deletes stack instances that are deployed within the member.
	// *   false: disables automatic deployment. After you disable automatic deployment, the stack instances remain unchanged even if members in the folder change.
	Parameters []*UpdateStackGroupRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The folder IDs in the resource directory. You can specify a maximum of five folder IDs.
	//
	// You must set at least one of the RdFolderIds and AccountIds parameters. The parameters are subject to the following items:
	//
	// *   If you set only the RdFolderIds parameter, stacks are deployed within all the members in the specified folders. If you specify the Root folder, ROS deploys the stacks within all the members in the resource directory.
	// *   If you set only the AccountIds parameter, stacks are deployed within the specified members.
	// *   If you set both parameters, the accounts specified by AccountIds must be contained in the folders specified by RdFolderIds.
	//
	// >  To view the folder IDs, go to the **Overview** page in the **Resource Management** console. For more information, see [View the basic information of a folder](~~111223~~).
	PermissionModel *string `json:"PermissionModel,omitempty" xml:"PermissionModel,omitempty"`
	// The region IDs of stack instances. You can specify a maximum of 20 region IDs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the operation to update the stack group.
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// The region IDs of stack instances. You can specify a maximum of 20 region IDs.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	TemplateBody   *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The permission model.
	//
	// Valid values:
	//
	// *   SELF_MANAGED: the self-managed permission model. This is the default value. If you use the self-managed model for the stack group, you must create RAM roles for the administrator and execution accounts, and establish a trust relationship between the accounts to deploy stacks within the execution account.
	// *   SERVICE_MANAGED: the service-managed permission model. If you use the service-managed model for the stack group, ROS creates service-linked roles for the administrator and execution accounts, and the administrator account uses its role to deploy stacks within the execution account.
	//
	// >- If stack instances have been created in the stack group, you cannot switch the permission mode of the stack group.
	// >- If you want to use the service-managed permission model to deploy stacks, your account must be the management account or a delegated administrator account of your resource directory and the trusted access feature is enabled for the account. For more information, see [Step 1: (Optional) Create a delegated administrator account](~~308253~~) and [Step 2: Enable trusted access](~~298229~~).
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the RAM role to be assumed by the administrator role AliyunROSStackGroupAdministrationRole. This parameter is required if you want to grant self-managed permissions to the stack group. If you do not specify a value for this parameter, the default value AliyunROSStackGroupExecutionRole is used. You can use this role in ROS to perform operations on the stacks that correspond to stack instances in the stack group.
	//
	// The name must be 1 to 64 characters in length, and can contain letters, digits, and hyphens (-).
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The information about automatic deployment settings.
	//
	// >  This parameter is required only if the PermissionModel parameter is set to SERVICE_MANAGED.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s UpdateStackGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupRequest) GoString() string {
	return s.String()
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

type UpdateStackGroupRequestAutoDeployment struct {
	// The IDs of the members in the resource directory. You can specify a maximum of 20 member IDs.
	//
	// >  To view the member IDs, go to the **Overview** page in the **Resource Management** console. For more information, see [View the detailed information of a member](~~111624~~).
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The IDs of the members in the resource directory. You can specify a maximum of 20 member IDs.
	//
	// >  To view the member IDs, go to the **Overview** page in the **Resource Management** console. For more information, see [View the detailed information of a member](~~111624~~).
	RetainStacksOnAccountRemoval *bool `json:"RetainStacksOnAccountRemoval,omitempty" xml:"RetainStacksOnAccountRemoval,omitempty"`
}

func (s UpdateStackGroupRequestAutoDeployment) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupRequestAutoDeployment) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupRequestAutoDeployment) SetEnabled(v bool) *UpdateStackGroupRequestAutoDeployment {
	s.Enabled = &v
	return s
}

func (s *UpdateStackGroupRequestAutoDeployment) SetRetainStacksOnAccountRemoval(v bool) *UpdateStackGroupRequestAutoDeployment {
	s.RetainStacksOnAccountRemoval = &v
	return s
}

type UpdateStackGroupRequestDeploymentTargets struct {
	// The list of one or more Alibaba Cloud accounts with which you want to share or unshare the template.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The ID of the operation.
	RdFolderIds []*string `json:"RdFolderIds,omitempty" xml:"RdFolderIds,omitempty" type:"Repeated"`
}

func (s UpdateStackGroupRequestDeploymentTargets) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupRequestDeploymentTargets) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupRequestDeploymentTargets) SetAccountIds(v []*string) *UpdateStackGroupRequestDeploymentTargets {
	s.AccountIds = v
	return s
}

func (s *UpdateStackGroupRequestDeploymentTargets) SetRdFolderIds(v []*string) *UpdateStackGroupRequestDeploymentTargets {
	s.RdFolderIds = v
	return s
}

type UpdateStackGroupRequestParameters struct {
	// Specifies whether to retain stacks in a member when you remove the member from the folder.
	//
	// Valid values:
	//
	// *   true: retains the stacks.
	// *   false: deletes the stacks.
	//
	// >  This parameter is required if the Enabled parameter is set to true.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The folders in which you want to use service-managed permissions to update stacks.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateStackGroupRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupRequestParameters) SetParameterKey(v string) *UpdateStackGroupRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackGroupRequestParameters) SetParameterValue(v string) *UpdateStackGroupRequestParameters {
	s.ParameterValue = &v
	return s
}

type UpdateStackGroupShrinkRequest struct {
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP or HTTPS web server or in an Alibaba Cloud Object Storage Service (OSS) bucket. The template body must be 1 to 524,288 bytes in length. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. If you do not specify the region ID of the OSS bucket, the value of the RegionId parameter is used.
	//
	// >  You must specify only one of the TemplateBody, TemplateURL, and TemplateId parameters.
	AccountIdsShrink *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	// The key of parameter N. If you do not specify the key and value of the parameter, ROS uses the default key and value in the template.
	//
	// Maximum value of N: 200.
	//
	// >  The Parameters parameter is optional. If you set the Parameters parameter, you must set the Parameters.N.ParameterKey parameter.
	AdministrationRoleName *string `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	// The IDs of the folders in the resource directory. You can specify up to five folder IDs.
	//
	// You can create stacks within all members in the specified folders. If you create stacks in the Root folder, the stacks are created within all members in the resource directory.
	//
	// >  To view the folder IDs, go to the **Overview** page in the **Resource Management** console. For more information, see [View the basic information of a folder](~~111223~~).
	AutoDeploymentShrink *string `json:"AutoDeployment,omitempty" xml:"AutoDeployment,omitempty"`
	// The option for the stack group. You can specify up to one option.
	Capabilities []*string `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Repeated"`
	// The ID of the template. This parameter applies to shared and private templates.
	//
	// >  You must specify only one of the TemplateBody, TemplateURL, and TemplateId parameters.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the request.
	DeploymentTargetsShrink *string `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length. If the length of the template body exceeds the upper limit, we recommend that you add parameters to the HTTP POST request body to prevent request failures caused by excessively long URLs.
	//
	// >  You must specify only one of the TemplateBody, TemplateURL, and TemplateId parameters.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The value of parameter N.
	//
	// Maximum value of N: 200.
	//
	// >  The Parameters parameter is optional. If you set the Parameters parameter, you must set the Parameters.N.ParameterValue parameter.
	ExecutionRoleName *string `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	// The version of the template. If you do not specify a version, the latest version is used.
	//
	// >  This parameter takes effect only if the TemplateId parameter is set.
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The list of parameters.
	OperationPreferencesShrink *string `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// Specifies whether to enable automatic deployment.
	//
	// Valid values:
	//
	// *   true: enables automatic deployment. If you add a member to the folder to which the stack group belongs after you enable automatic deployment, the stack group deploys its stack instances within the member. If you remove a member from the folder, the stack group deletes stack instances that are deployed within the member.
	// *   false: disables automatic deployment. After you disable automatic deployment, the stack instances remain unchanged even if members in the folder change.
	Parameters []*UpdateStackGroupShrinkRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The folder IDs in the resource directory. You can specify a maximum of five folder IDs.
	//
	// You must set at least one of the RdFolderIds and AccountIds parameters. The parameters are subject to the following items:
	//
	// *   If you set only the RdFolderIds parameter, stacks are deployed within all the members in the specified folders. If you specify the Root folder, ROS deploys the stacks within all the members in the resource directory.
	// *   If you set only the AccountIds parameter, stacks are deployed within the specified members.
	// *   If you set both parameters, the accounts specified by AccountIds must be contained in the folders specified by RdFolderIds.
	//
	// >  To view the folder IDs, go to the **Overview** page in the **Resource Management** console. For more information, see [View the basic information of a folder](~~111223~~).
	PermissionModel *string `json:"PermissionModel,omitempty" xml:"PermissionModel,omitempty"`
	// The region IDs of stack instances. You can specify a maximum of 20 region IDs.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the operation to update the stack group.
	RegionIdsShrink *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	// The region IDs of stack instances. You can specify a maximum of 20 region IDs.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	TemplateBody   *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The permission model.
	//
	// Valid values:
	//
	// *   SELF_MANAGED: the self-managed permission model. This is the default value. If you use the self-managed model for the stack group, you must create RAM roles for the administrator and execution accounts, and establish a trust relationship between the accounts to deploy stacks within the execution account.
	// *   SERVICE_MANAGED: the service-managed permission model. If you use the service-managed model for the stack group, ROS creates service-linked roles for the administrator and execution accounts, and the administrator account uses its role to deploy stacks within the execution account.
	//
	// >- If stack instances have been created in the stack group, you cannot switch the permission mode of the stack group.
	// >- If you want to use the service-managed permission model to deploy stacks, your account must be the management account or a delegated administrator account of your resource directory and the trusted access feature is enabled for the account. For more information, see [Step 1: (Optional) Create a delegated administrator account](~~308253~~) and [Step 2: Enable trusted access](~~298229~~).
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the RAM role to be assumed by the administrator role AliyunROSStackGroupAdministrationRole. This parameter is required if you want to grant self-managed permissions to the stack group. If you do not specify a value for this parameter, the default value AliyunROSStackGroupExecutionRole is used. You can use this role in ROS to perform operations on the stacks that correspond to stack instances in the stack group.
	//
	// The name must be 1 to 64 characters in length, and can contain letters, digits, and hyphens (-).
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The information about automatic deployment settings.
	//
	// >  This parameter is required only if the PermissionModel parameter is set to SERVICE_MANAGED.
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s UpdateStackGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupShrinkRequest) GoString() string {
	return s.String()
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

type UpdateStackGroupShrinkRequestParameters struct {
	// Specifies whether to retain stacks in a member when you remove the member from the folder.
	//
	// Valid values:
	//
	// *   true: retains the stacks.
	// *   false: deletes the stacks.
	//
	// >  This parameter is required if the Enabled parameter is set to true.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The folders in which you want to use service-managed permissions to update stacks.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateStackGroupShrinkRequestParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupShrinkRequestParameters) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupShrinkRequestParameters) SetParameterKey(v string) *UpdateStackGroupShrinkRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackGroupShrinkRequestParameters) SetParameterValue(v string) *UpdateStackGroupShrinkRequestParameters {
	s.ParameterValue = &v
	return s
}

type UpdateStackGroupResponseBody struct {
	// The ID of the operation.
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateStackGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupResponseBody) SetOperationId(v string) *UpdateStackGroupResponseBody {
	s.OperationId = &v
	return s
}

func (s *UpdateStackGroupResponseBody) SetRequestId(v string) *UpdateStackGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateStackGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateStackGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateStackGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateStackGroupResponse) SetHeaders(v map[string]*string) *UpdateStackGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateStackGroupResponse) SetStatusCode(v int32) *UpdateStackGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStackGroupResponse) SetBody(v *UpdateStackGroupResponseBody) *UpdateStackGroupResponse {
	s.Body = v
	return s
}

type UpdateStackInstancesRequest struct {
	// The IDs of the execution accounts within which you want to deploy stacks in self-managed mode. You can specify up to 20 execution account IDs.
	//
	// > If you want to update stacks in self-managed permission mode, you must specify this parameter.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\
	// The token can contain letters, digits, hyphens (-), and underscores (\_), and cannot exceed 64 characters in length.\
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The folders in which you want to deploy stacks in service-managed mode.
	//
	// > If you want to update stacks in service-managed permission mode, you must specify this parameter.
	DeploymentTargets *UpdateStackInstancesRequestDeploymentTargets `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty" type:"Struct"`
	// The description of the update operation.
	//
	// The description must be 1 to 256 characters in length.
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The preference settings of the update operation.
	//
	// The following parameters are available:
	// -  {"FailureToleranceCount": N}
	//
	//     The number of accounts within which stack operation failures are allowed in each region. If the value of this parameter is exceeded in a region, ROS stops the operation in the region. If ROS stops the operation in one region, ROS stops the operation in other regions.
	//
	//     Valid values of N: 0 to 20.
	//
	//     If you do not specify FailureToleranceCount, 0 is used as the default value.
	//
	// -  {"FailureTolerancePercentage": N}
	//
	//     The percentage of the number of accounts within which stack operation failures are allowed to the total number of accounts in each region. If the value of this parameter is exceeded, ROS stops the operation in the region.
	//
	//     Valid values of N: 0 to 100. If the numeric value in the percentage is not an integer, ROS rounds the value down to the nearest integer.
	//
	//     If you do not specify FailureTolerancePercentage, 0 is used as the default value.
	//
	// -  {"MaxConcurrentCount": N}
	//
	//     The maximum number of accounts within which multiple stacks are deployed at the same time in each region.
	//
	//     Valid values of N: 1 to 20.
	//
	//     If you do not specify MaxConcurrentCount, 1 is used as the default value.
	//
	// - {"MaxConcurrentPercentage": N}
	//
	//     The percentage of the maximum number of accounts within which stacks are deployed at the same time to the total number of accounts in each region.
	//
	//     Valid values: 1 to 100. If the numeric value in the percentage is not an integer, ROS rounds the value down to the nearest integer.
	//
	//     If you do not specify MaxConcurrentPercentage, 1 is used as the default value.
	//
	// - {"RegionConcurrencyType": N}
	//
	//   The mode that you want to use to deploy stacks across regions. Valid values:
	//   - SEQUENTIAL (default): deploys stacks in the specified regions one by one in sequence. This way, ROS deploys stacks in only one region at a time.
	//
	//    - PARALLEL: deploys stacks in all the specified regions in parallel.
	//
	// Separate multiple parameters with commas (,).
	//
	// > - You can specify only one of the following parameters: MaxConcurrentCount and MaxConcurrentPercentage.
	// > - You can specify only one of the following parameters: FailureToleranceCount and FailureTolerancePercentage.
	OperationPreferences map[string]interface{} `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// The parameters that are used to override specific parameters.
	ParameterOverrides []*UpdateStackInstancesRequestParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the regions where you want to update the stacks. You can specify up to 20 region IDs.
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// The name of the stack group. The name must be unique within a region.\
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or a letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The timeout period for the update operation.
	//
	// *   Default value: 60.
	// *   Unit: minutes.
	TimeoutInMinutes *int64 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s UpdateStackInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackInstancesRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesRequest) SetAccountIds(v []*string) *UpdateStackInstancesRequest {
	s.AccountIds = v
	return s
}

func (s *UpdateStackInstancesRequest) SetClientToken(v string) *UpdateStackInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackInstancesRequest) SetDeploymentTargets(v *UpdateStackInstancesRequestDeploymentTargets) *UpdateStackInstancesRequest {
	s.DeploymentTargets = v
	return s
}

func (s *UpdateStackInstancesRequest) SetOperationDescription(v string) *UpdateStackInstancesRequest {
	s.OperationDescription = &v
	return s
}

func (s *UpdateStackInstancesRequest) SetOperationPreferences(v map[string]interface{}) *UpdateStackInstancesRequest {
	s.OperationPreferences = v
	return s
}

func (s *UpdateStackInstancesRequest) SetParameterOverrides(v []*UpdateStackInstancesRequestParameterOverrides) *UpdateStackInstancesRequest {
	s.ParameterOverrides = v
	return s
}

func (s *UpdateStackInstancesRequest) SetRegionId(v string) *UpdateStackInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackInstancesRequest) SetRegionIds(v []*string) *UpdateStackInstancesRequest {
	s.RegionIds = v
	return s
}

func (s *UpdateStackInstancesRequest) SetStackGroupName(v string) *UpdateStackInstancesRequest {
	s.StackGroupName = &v
	return s
}

func (s *UpdateStackInstancesRequest) SetTimeoutInMinutes(v int64) *UpdateStackInstancesRequest {
	s.TimeoutInMinutes = &v
	return s
}

type UpdateStackInstancesRequestDeploymentTargets struct {
	// The IDs of the member accounts in the resource directory. You can specify up to 20 member account IDs.
	//
	// > To view the member account IDs, go to the **Overview** page in the **Resource Management** console. For more information, see [View the details of a member](~~111624~~).
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The folder IDs of the resource directory.
	RdFolderIds []*string `json:"RdFolderIds,omitempty" xml:"RdFolderIds,omitempty" type:"Repeated"`
}

func (s UpdateStackInstancesRequestDeploymentTargets) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackInstancesRequestDeploymentTargets) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesRequestDeploymentTargets) SetAccountIds(v []*string) *UpdateStackInstancesRequestDeploymentTargets {
	s.AccountIds = v
	return s
}

func (s *UpdateStackInstancesRequestDeploymentTargets) SetRdFolderIds(v []*string) *UpdateStackInstancesRequestDeploymentTargets {
	s.RdFolderIds = v
	return s
}

type UpdateStackInstancesRequestParameterOverrides struct {
	// The key of parameter N that you want to use to override a specific parameter. If you do not specify this parameter, ROS uses the name that you specified when you created the stack group.
	//
	// Maximum value of N: 200.
	//
	// > -  ParameterOverrides is optional.
	// > - If you specify ParameterOverrides, you must specify ParameterOverrides.N.ParameterKey and ParameterOverrides.N.ParameterValue.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N that you want to use to override a specific parameter. If you do not specify this parameter, ROS uses the value that you specified when you created the stack group.
	//
	// Maximum value of N: 200.
	//
	// > -  ParameterOverrides is optional.
	// > - If you specify ParameterOverrides, you must specify ParameterOverrides.N.ParameterKey and ParameterOverrides.N.ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateStackInstancesRequestParameterOverrides) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackInstancesRequestParameterOverrides) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesRequestParameterOverrides) SetParameterKey(v string) *UpdateStackInstancesRequestParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackInstancesRequestParameterOverrides) SetParameterValue(v string) *UpdateStackInstancesRequestParameterOverrides {
	s.ParameterValue = &v
	return s
}

type UpdateStackInstancesShrinkRequest struct {
	// The IDs of the execution accounts within which you want to deploy stacks in self-managed mode. You can specify up to 20 execution account IDs.
	//
	// > If you want to update stacks in self-managed permission mode, you must specify this parameter.
	AccountIdsShrink *string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.\
	// The token can contain letters, digits, hyphens (-), and underscores (\_), and cannot exceed 64 characters in length.\
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The folders in which you want to deploy stacks in service-managed mode.
	//
	// > If you want to update stacks in service-managed permission mode, you must specify this parameter.
	DeploymentTargetsShrink *string `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty"`
	// The description of the update operation.
	//
	// The description must be 1 to 256 characters in length.
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The preference settings of the update operation.
	//
	// The following parameters are available:
	// -  {"FailureToleranceCount": N}
	//
	//     The number of accounts within which stack operation failures are allowed in each region. If the value of this parameter is exceeded in a region, ROS stops the operation in the region. If ROS stops the operation in one region, ROS stops the operation in other regions.
	//
	//     Valid values of N: 0 to 20.
	//
	//     If you do not specify FailureToleranceCount, 0 is used as the default value.
	//
	// -  {"FailureTolerancePercentage": N}
	//
	//     The percentage of the number of accounts within which stack operation failures are allowed to the total number of accounts in each region. If the value of this parameter is exceeded, ROS stops the operation in the region.
	//
	//     Valid values of N: 0 to 100. If the numeric value in the percentage is not an integer, ROS rounds the value down to the nearest integer.
	//
	//     If you do not specify FailureTolerancePercentage, 0 is used as the default value.
	//
	// -  {"MaxConcurrentCount": N}
	//
	//     The maximum number of accounts within which multiple stacks are deployed at the same time in each region.
	//
	//     Valid values of N: 1 to 20.
	//
	//     If you do not specify MaxConcurrentCount, 1 is used as the default value.
	//
	// - {"MaxConcurrentPercentage": N}
	//
	//     The percentage of the maximum number of accounts within which stacks are deployed at the same time to the total number of accounts in each region.
	//
	//     Valid values: 1 to 100. If the numeric value in the percentage is not an integer, ROS rounds the value down to the nearest integer.
	//
	//     If you do not specify MaxConcurrentPercentage, 1 is used as the default value.
	//
	// - {"RegionConcurrencyType": N}
	//
	//   The mode that you want to use to deploy stacks across regions. Valid values:
	//   - SEQUENTIAL (default): deploys stacks in the specified regions one by one in sequence. This way, ROS deploys stacks in only one region at a time.
	//
	//    - PARALLEL: deploys stacks in all the specified regions in parallel.
	//
	// Separate multiple parameters with commas (,).
	//
	// > - You can specify only one of the following parameters: MaxConcurrentCount and MaxConcurrentPercentage.
	// > - You can specify only one of the following parameters: FailureToleranceCount and FailureTolerancePercentage.
	OperationPreferencesShrink *string `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty"`
	// The parameters that are used to override specific parameters.
	ParameterOverrides []*UpdateStackInstancesShrinkRequestParameterOverrides `json:"ParameterOverrides,omitempty" xml:"ParameterOverrides,omitempty" type:"Repeated"`
	// The region ID of the stack group. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the regions where you want to update the stacks. You can specify up to 20 region IDs.
	RegionIdsShrink *string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty"`
	// The name of the stack group. The name must be unique within a region.\
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or a letter.
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The timeout period for the update operation.
	//
	// *   Default value: 60.
	// *   Unit: minutes.
	TimeoutInMinutes *int64 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
}

func (s UpdateStackInstancesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackInstancesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesShrinkRequest) SetAccountIdsShrink(v string) *UpdateStackInstancesShrinkRequest {
	s.AccountIdsShrink = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetClientToken(v string) *UpdateStackInstancesShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetDeploymentTargetsShrink(v string) *UpdateStackInstancesShrinkRequest {
	s.DeploymentTargetsShrink = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetOperationDescription(v string) *UpdateStackInstancesShrinkRequest {
	s.OperationDescription = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetOperationPreferencesShrink(v string) *UpdateStackInstancesShrinkRequest {
	s.OperationPreferencesShrink = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetParameterOverrides(v []*UpdateStackInstancesShrinkRequestParameterOverrides) *UpdateStackInstancesShrinkRequest {
	s.ParameterOverrides = v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetRegionId(v string) *UpdateStackInstancesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetRegionIdsShrink(v string) *UpdateStackInstancesShrinkRequest {
	s.RegionIdsShrink = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetStackGroupName(v string) *UpdateStackInstancesShrinkRequest {
	s.StackGroupName = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequest) SetTimeoutInMinutes(v int64) *UpdateStackInstancesShrinkRequest {
	s.TimeoutInMinutes = &v
	return s
}

type UpdateStackInstancesShrinkRequestParameterOverrides struct {
	// The key of parameter N that you want to use to override a specific parameter. If you do not specify this parameter, ROS uses the name that you specified when you created the stack group.
	//
	// Maximum value of N: 200.
	//
	// > -  ParameterOverrides is optional.
	// > - If you specify ParameterOverrides, you must specify ParameterOverrides.N.ParameterKey and ParameterOverrides.N.ParameterValue.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of parameter N that you want to use to override a specific parameter. If you do not specify this parameter, ROS uses the value that you specified when you created the stack group.
	//
	// Maximum value of N: 200.
	//
	// > -  ParameterOverrides is optional.
	// > - If you specify ParameterOverrides, you must specify ParameterOverrides.N.ParameterKey and ParameterOverrides.N.ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateStackInstancesShrinkRequestParameterOverrides) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackInstancesShrinkRequestParameterOverrides) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesShrinkRequestParameterOverrides) SetParameterKey(v string) *UpdateStackInstancesShrinkRequestParameterOverrides {
	s.ParameterKey = &v
	return s
}

func (s *UpdateStackInstancesShrinkRequestParameterOverrides) SetParameterValue(v string) *UpdateStackInstancesShrinkRequestParameterOverrides {
	s.ParameterValue = &v
	return s
}

type UpdateStackInstancesResponseBody struct {
	// The ID of the operation.
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateStackInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesResponseBody) SetOperationId(v string) *UpdateStackInstancesResponseBody {
	s.OperationId = &v
	return s
}

func (s *UpdateStackInstancesResponseBody) SetRequestId(v string) *UpdateStackInstancesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateStackInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateStackInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateStackInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackInstancesResponse) GoString() string {
	return s.String()
}

func (s *UpdateStackInstancesResponse) SetHeaders(v map[string]*string) *UpdateStackInstancesResponse {
	s.Headers = v
	return s
}

func (s *UpdateStackInstancesResponse) SetStatusCode(v int32) *UpdateStackInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStackInstancesResponse) SetBody(v *UpdateStackInstancesResponseBody) *UpdateStackInstancesResponse {
	s.Body = v
	return s
}

type UpdateStackTemplateByResourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests.
	//
	// The token can be up to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only preview the corrected template in this request. Default value: false. Valid values:
	//
	// *   true: returns the content of the corrected template and does not correct the template. After Resource Orchestration Service (ROS) compares the corrected template with the original template, ROS determines whether to execute the correction.
	// *   false: corrects the template to eliminate drift.
	//
	// >  We recommend that you set the DryRun parameter to true to preview the corrected template. If the template content meets expectations, set the DryRun parameter to false to execute the correction.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The logical ID of resource.
	LogicalResourceId []*string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty" type:"Repeated"`
	// The region ID of the stack. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The format of the returned template. Default value: JSON. Valid values:
	//
	// *   JSON
	// *   YAML
	TemplateFormat *string `json:"TemplateFormat,omitempty" xml:"TemplateFormat,omitempty"`
}

func (s UpdateStackTemplateByResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackTemplateByResourcesRequest) GoString() string {
	return s.String()
}

func (s *UpdateStackTemplateByResourcesRequest) SetClientToken(v string) *UpdateStackTemplateByResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetDryRun(v bool) *UpdateStackTemplateByResourcesRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetLogicalResourceId(v []*string) *UpdateStackTemplateByResourcesRequest {
	s.LogicalResourceId = v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetRegionId(v string) *UpdateStackTemplateByResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetStackId(v string) *UpdateStackTemplateByResourcesRequest {
	s.StackId = &v
	return s
}

func (s *UpdateStackTemplateByResourcesRequest) SetTemplateFormat(v string) *UpdateStackTemplateByResourcesRequest {
	s.TemplateFormat = &v
	return s
}

type UpdateStackTemplateByResourcesResponseBody struct {
	// The template content after correction.
	NewTemplateBody *string `json:"NewTemplateBody,omitempty" xml:"NewTemplateBody,omitempty"`
	// The template content before correction.
	OldTemplateBody *string `json:"OldTemplateBody,omitempty" xml:"OldTemplateBody,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateStackTemplateByResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackTemplateByResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStackTemplateByResourcesResponseBody) SetNewTemplateBody(v string) *UpdateStackTemplateByResourcesResponseBody {
	s.NewTemplateBody = &v
	return s
}

func (s *UpdateStackTemplateByResourcesResponseBody) SetOldTemplateBody(v string) *UpdateStackTemplateByResourcesResponseBody {
	s.OldTemplateBody = &v
	return s
}

func (s *UpdateStackTemplateByResourcesResponseBody) SetRequestId(v string) *UpdateStackTemplateByResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateStackTemplateByResourcesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateStackTemplateByResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateStackTemplateByResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStackTemplateByResourcesResponse) GoString() string {
	return s.String()
}

func (s *UpdateStackTemplateByResourcesResponse) SetHeaders(v map[string]*string) *UpdateStackTemplateByResourcesResponse {
	s.Headers = v
	return s
}

func (s *UpdateStackTemplateByResourcesResponse) SetStatusCode(v int32) *UpdateStackTemplateByResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStackTemplateByResourcesResponse) SetBody(v *UpdateStackTemplateByResourcesResponseBody) *UpdateStackTemplateByResourcesResponse {
	s.Body = v
	return s
}

type UpdateTemplateRequest struct {
	// The description of the template. It can be up to 256 characters in length.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.
	//
	// If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.
	//
	// You can specify only one of the TemplateBody and TemplateURL parameters.
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The ID of the template. This parameter applies to shared and private templates.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The name of the template.
	//
	// The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\_). It must start with a digit or letter.
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template located in an HTTP or HTTPS web server or an Alibaba Cloud OSS bucket. Examples: oss://ros/template/demo and oss://ros/template/demo?RegionId=cn-hangzhou. The template can be up to 524,288 bytes in length, and the URL can be up to 1,024 bytes in length.
	//
	// >  If the region of the OSS bucket is not specified, the RegionId value is used.
	//
	// You can specify only one of the TemplateBody and TemplateURL parameters.
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
}

func (s UpdateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateRequest) SetDescription(v string) *UpdateTemplateRequest {
	s.Description = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateBody(v string) *UpdateTemplateRequest {
	s.TemplateBody = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateId(v string) *UpdateTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateName(v string) *UpdateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateTemplateRequest) SetTemplateURL(v string) *UpdateTemplateRequest {
	s.TemplateURL = &v
	return s
}

type UpdateTemplateResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the template.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponseBody) SetRequestId(v string) *UpdateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateResponseBody) SetTemplateId(v string) *UpdateTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type UpdateTemplateResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateResponse) SetHeaders(v map[string]*string) *UpdateTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateResponse) SetStatusCode(v int32) *UpdateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTemplateResponse) SetBody(v *UpdateTemplateResponseBody) *UpdateTemplateResponse {
	s.Body = v
	return s
}

type UpdateTemplateScratchRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the scenario.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution mode. Valid values:
	//
	// *   Async (default)
	// *   Sync
	//
	// > If you have a wide scope of resources, Sync takes longer. If you set ExecutionMode to Sync, we recommend that you specify ClientToken to prevent the execution timeout.
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The policy based on which the logical ID is generated. Valid values:
	//
	// *   LongTypePrefixAndIndexSuffix: long-type prefix + index-type suffix
	// *   LongTypePrefixAndHashSuffix: long-type prefix + hash-type suffix
	// *   ShortTypePrefixAndHashSuffix: short-type prefix + hash-type suffix
	//
	// >  If you set TemplateScratchType to ArchitectureDetection, the default value of LogicalIdStrategy is LongTypePrefixAndHashSuffix. In other cases, the default value of LogicalIdStrategy is LongTypePrefixAndIndexSuffix.
	LogicalIdStrategy *string `json:"LogicalIdStrategy,omitempty" xml:"LogicalIdStrategy,omitempty"`
	// The parameters that you want to configure for the scenario.
	PreferenceParameters []*UpdateTemplateScratchRequestPreferenceParameters `json:"PreferenceParameters,omitempty" xml:"PreferenceParameters,omitempty" type:"Repeated"`
	// The region ID of the scenario.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The source resource group.
	SourceResourceGroup *UpdateTemplateScratchRequestSourceResourceGroup `json:"SourceResourceGroup,omitempty" xml:"SourceResourceGroup,omitempty" type:"Struct"`
	// The source resources.
	//
	// If you specify source resources as the value of SourceResources when TemplateScratchType is set to ArchitectureDetection, the system detects the schema data of all resources that are associated with the specified source resources. For example, if you specify the ID of a Classic Load Balancer (CLB) instance as the value of SourceResources, the system detects the schema data of resources, such as Elastic Compute Service (ECS) instances, vSwitches, and VPCs, that are associated with the CLB instance.
	//
	// If you set TemplateScratchType to ArchitectureDetection, you can specify up to 20 source resources for SourceResources. In other cases, you can specify up to 200 source resources.
	SourceResources []*UpdateTemplateScratchRequestSourceResources `json:"SourceResources,omitempty" xml:"SourceResources,omitempty" type:"Repeated"`
	// The source tag.
	SourceTag *UpdateTemplateScratchRequestSourceTag `json:"SourceTag,omitempty" xml:"SourceTag,omitempty" type:"Struct"`
	// The ID of the scenario.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
}

func (s UpdateTemplateScratchRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateScratchRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchRequest) SetClientToken(v string) *UpdateTemplateScratchRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTemplateScratchRequest) SetDescription(v string) *UpdateTemplateScratchRequest {
	s.Description = &v
	return s
}

func (s *UpdateTemplateScratchRequest) SetExecutionMode(v string) *UpdateTemplateScratchRequest {
	s.ExecutionMode = &v
	return s
}

func (s *UpdateTemplateScratchRequest) SetLogicalIdStrategy(v string) *UpdateTemplateScratchRequest {
	s.LogicalIdStrategy = &v
	return s
}

func (s *UpdateTemplateScratchRequest) SetPreferenceParameters(v []*UpdateTemplateScratchRequestPreferenceParameters) *UpdateTemplateScratchRequest {
	s.PreferenceParameters = v
	return s
}

func (s *UpdateTemplateScratchRequest) SetRegionId(v string) *UpdateTemplateScratchRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTemplateScratchRequest) SetResourceGroupId(v string) *UpdateTemplateScratchRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTemplateScratchRequest) SetSourceResourceGroup(v *UpdateTemplateScratchRequestSourceResourceGroup) *UpdateTemplateScratchRequest {
	s.SourceResourceGroup = v
	return s
}

func (s *UpdateTemplateScratchRequest) SetSourceResources(v []*UpdateTemplateScratchRequestSourceResources) *UpdateTemplateScratchRequest {
	s.SourceResources = v
	return s
}

func (s *UpdateTemplateScratchRequest) SetSourceTag(v *UpdateTemplateScratchRequestSourceTag) *UpdateTemplateScratchRequest {
	s.SourceTag = v
	return s
}

func (s *UpdateTemplateScratchRequest) SetTemplateScratchId(v string) *UpdateTemplateScratchRequest {
	s.TemplateScratchId = &v
	return s
}

type UpdateTemplateScratchRequestPreferenceParameters struct {
	// The name of the parameter.
	//
	// For more information about the valid values of ParameterKey, see the "**Additional information about request parameters**" section of this topic.
	//
	// > - PreferenceParameters is optional. If you specify PreferenceParameters, you must specify both ParameterKey and ParameterValue.
	// > - If you set TemplateScratchType to ResourceImport, you must set ParameterKey to DeletionPolicy.
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The value of the parameter. The value of ParameterValue varies based on the value of ParameterKey.
	//
	// For more information about the valid values of ParameterValue, see the "**Additional information about request parameters**" section of this topic.
	//
	// >  PreferenceParameters is optional. If you specify PreferenceParameters, you must specify both ParameterKey and ParameterValue.
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s UpdateTemplateScratchRequestPreferenceParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateScratchRequestPreferenceParameters) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchRequestPreferenceParameters) SetParameterKey(v string) *UpdateTemplateScratchRequestPreferenceParameters {
	s.ParameterKey = &v
	return s
}

func (s *UpdateTemplateScratchRequestPreferenceParameters) SetParameterValue(v string) *UpdateTemplateScratchRequestPreferenceParameters {
	s.ParameterValue = &v
	return s
}

type UpdateTemplateScratchRequestSourceResourceGroup struct {
	// The ID of the source resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource types.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s UpdateTemplateScratchRequestSourceResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateScratchRequestSourceResourceGroup) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchRequestSourceResourceGroup) SetResourceGroupId(v string) *UpdateTemplateScratchRequestSourceResourceGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTemplateScratchRequestSourceResourceGroup) SetResourceTypeFilter(v []*string) *UpdateTemplateScratchRequestSourceResourceGroup {
	s.ResourceTypeFilter = v
	return s
}

type UpdateTemplateScratchRequestSourceResources struct {
	// The resource ID.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s UpdateTemplateScratchRequestSourceResources) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateScratchRequestSourceResources) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchRequestSourceResources) SetResourceId(v string) *UpdateTemplateScratchRequestSourceResources {
	s.ResourceId = &v
	return s
}

func (s *UpdateTemplateScratchRequestSourceResources) SetResourceType(v string) *UpdateTemplateScratchRequestSourceResources {
	s.ResourceType = &v
	return s
}

type UpdateTemplateScratchRequestSourceTag struct {
	// The source tags that consist of key-value pairs.
	//
	// If you want to specify only the tag key, you must set the tag value to an empty string. Example: {"TagKey": ""}.
	//
	// If you set TemplateScratchType to ArchitectureDetection, you can add up to five source tags. In other cases, you can add up to 10 source tags.
	ResourceTags map[string]interface{} `json:"ResourceTags,omitempty" xml:"ResourceTags,omitempty"`
	// The filters for resource types.
	ResourceTypeFilter []*string `json:"ResourceTypeFilter,omitempty" xml:"ResourceTypeFilter,omitempty" type:"Repeated"`
}

func (s UpdateTemplateScratchRequestSourceTag) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateScratchRequestSourceTag) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchRequestSourceTag) SetResourceTags(v map[string]interface{}) *UpdateTemplateScratchRequestSourceTag {
	s.ResourceTags = v
	return s
}

func (s *UpdateTemplateScratchRequestSourceTag) SetResourceTypeFilter(v []*string) *UpdateTemplateScratchRequestSourceTag {
	s.ResourceTypeFilter = v
	return s
}

type UpdateTemplateScratchShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// For more information, see [How to ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the scenario.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution mode. Valid values:
	//
	// *   Async (default)
	// *   Sync
	//
	// > If you have a wide scope of resources, Sync takes longer. If you set ExecutionMode to Sync, we recommend that you specify ClientToken to prevent the execution timeout.
	ExecutionMode *string `json:"ExecutionMode,omitempty" xml:"ExecutionMode,omitempty"`
	// The policy based on which the logical ID is generated. Valid values:
	//
	// *   LongTypePrefixAndIndexSuffix: long-type prefix + index-type suffix
	// *   LongTypePrefixAndHashSuffix: long-type prefix + hash-type suffix
	// *   ShortTypePrefixAndHashSuffix: short-type prefix + hash-type suffix
	//
	// >  If you set TemplateScratchType to ArchitectureDetection, the default value of LogicalIdStrategy is LongTypePrefixAndHashSuffix. In other cases, the default value of LogicalIdStrategy is LongTypePrefixAndIndexSuffix.
	LogicalIdStrategy *string `json:"LogicalIdStrategy,omitempty" xml:"LogicalIdStrategy,omitempty"`
	// The parameters that you want to configure for the scenario.
	PreferenceParametersShrink *string `json:"PreferenceParameters,omitempty" xml:"PreferenceParameters,omitempty"`
	// The region ID of the scenario.
	//
	// You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The source resource group.
	SourceResourceGroupShrink *string `json:"SourceResourceGroup,omitempty" xml:"SourceResourceGroup,omitempty"`
	// The source resources.
	//
	// If you specify source resources as the value of SourceResources when TemplateScratchType is set to ArchitectureDetection, the system detects the schema data of all resources that are associated with the specified source resources. For example, if you specify the ID of a Classic Load Balancer (CLB) instance as the value of SourceResources, the system detects the schema data of resources, such as Elastic Compute Service (ECS) instances, vSwitches, and VPCs, that are associated with the CLB instance.
	//
	// If you set TemplateScratchType to ArchitectureDetection, you can specify up to 20 source resources for SourceResources. In other cases, you can specify up to 200 source resources.
	SourceResourcesShrink *string `json:"SourceResources,omitempty" xml:"SourceResources,omitempty"`
	// The source tag.
	SourceTagShrink *string `json:"SourceTag,omitempty" xml:"SourceTag,omitempty"`
	// The ID of the scenario.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
}

func (s UpdateTemplateScratchShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateScratchShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchShrinkRequest) SetClientToken(v string) *UpdateTemplateScratchShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetDescription(v string) *UpdateTemplateScratchShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetExecutionMode(v string) *UpdateTemplateScratchShrinkRequest {
	s.ExecutionMode = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetLogicalIdStrategy(v string) *UpdateTemplateScratchShrinkRequest {
	s.LogicalIdStrategy = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetPreferenceParametersShrink(v string) *UpdateTemplateScratchShrinkRequest {
	s.PreferenceParametersShrink = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetRegionId(v string) *UpdateTemplateScratchShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetResourceGroupId(v string) *UpdateTemplateScratchShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetSourceResourceGroupShrink(v string) *UpdateTemplateScratchShrinkRequest {
	s.SourceResourceGroupShrink = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetSourceResourcesShrink(v string) *UpdateTemplateScratchShrinkRequest {
	s.SourceResourcesShrink = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetSourceTagShrink(v string) *UpdateTemplateScratchShrinkRequest {
	s.SourceTagShrink = &v
	return s
}

func (s *UpdateTemplateScratchShrinkRequest) SetTemplateScratchId(v string) *UpdateTemplateScratchShrinkRequest {
	s.TemplateScratchId = &v
	return s
}

type UpdateTemplateScratchResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scenario.
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
}

func (s UpdateTemplateScratchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateScratchResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchResponseBody) SetRequestId(v string) *UpdateTemplateScratchResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTemplateScratchResponseBody) SetTemplateScratchId(v string) *UpdateTemplateScratchResponseBody {
	s.TemplateScratchId = &v
	return s
}

type UpdateTemplateScratchResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTemplateScratchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTemplateScratchResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTemplateScratchResponse) GoString() string {
	return s.String()
}

func (s *UpdateTemplateScratchResponse) SetHeaders(v map[string]*string) *UpdateTemplateScratchResponse {
	s.Headers = v
	return s
}

func (s *UpdateTemplateScratchResponse) SetStatusCode(v int32) *UpdateTemplateScratchResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTemplateScratchResponse) SetBody(v *UpdateTemplateScratchResponseBody) *UpdateTemplateScratchResponse {
	s.Body = v
	return s
}

type ValidateTemplateRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can be up to 64 characters in length, and can contain letters, digits, hyphens (-), and underscores (\_).
	//
	// For more information, see [Ensure idempotence](~~134212~~).
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the template. You can call the [DescribeRegions](~~131035~~) operation to query the most recent region list.
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateBody *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	// The URL of the file that contains the template body. The URL must point to a template that is located on an HTTP web server or in an Object Storage Service (OSS) bucket, such as oss://ros/template/demo or oss://ros/template/demo?RegionId=cn-hangzhou. The template body can be up to 524,288 bytes in length.
	//
	// > If you do not specify the region ID of the OSS bucket, the value of RegionId is used.
	//
	// You can specify one of TemplateBody and TemplateURL, but not both of them. The URL can be up to 1,024 bytes in length.\
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The options that are used to control the generation of information about the stack update. You can specify up to two options.
	UpdateInfoOptions []*string `json:"UpdateInfoOptions,omitempty" xml:"UpdateInfoOptions,omitempty" type:"Repeated"`
	// Specifies whether to enable additional validation for the template. Valid values:
	//
	// *   None (default): does not enable additional validation.
	// *   EnableTerraformValidation: runs the `terraform validate` command in the Terraform CLI to enable additional validation for a Terraform template.
	// *   EnableFastTerraformValidation: runs a command that is similar to the `terraform validate` command in the Terraform CLI to enable additional validation for a Terraform template.
	//
	// > Compared with the EnableTerraformValidation method, the EnableFastTerraformValidation method validates a template at a faster speed but a lower integrity level.
	ValidationOption *string `json:"ValidationOption,omitempty" xml:"ValidationOption,omitempty"`
}

func (s ValidateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateRequest) GoString() string {
	return s.String()
}

func (s *ValidateTemplateRequest) SetClientToken(v string) *ValidateTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *ValidateTemplateRequest) SetRegionId(v string) *ValidateTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *ValidateTemplateRequest) SetTemplateBody(v string) *ValidateTemplateRequest {
	s.TemplateBody = &v
	return s
}

func (s *ValidateTemplateRequest) SetTemplateURL(v string) *ValidateTemplateRequest {
	s.TemplateURL = &v
	return s
}

func (s *ValidateTemplateRequest) SetUpdateInfoOptions(v []*string) *ValidateTemplateRequest {
	s.UpdateInfoOptions = v
	return s
}

func (s *ValidateTemplateRequest) SetValidationOption(v string) *ValidateTemplateRequest {
	s.ValidationOption = &v
	return s
}

type ValidateTemplateResponseBody struct {
	// The description of the template.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The outputs of the template.
	Outputs []*ValidateTemplateResponseBodyOutputs `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// The parameters that are defined in the Parameters section of the template.
	Parameters []map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource types that are used in the template.
	ResourceTypes *ValidateTemplateResponseBodyResourceTypes `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Struct"`
	// The regular resources that are defined in the template.
	//
	// > - For a Resource Orchestration Service (ROS) template, the resource whose definition contains `Count` is not displayed as a list.
	// > -  For a Terraform template, the resource whose definition contains `count` or `for_each` is not displayed as a list.
	Resources []*ValidateTemplateResponseBodyResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The information about the stack update. This parameter cannot be returned if the value of UpdateInfoOptions contains Disabled.
	UpdateInfo *ValidateTemplateResponseBodyUpdateInfo `json:"UpdateInfo,omitempty" xml:"UpdateInfo,omitempty" type:"Struct"`
}

func (s ValidateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponseBody) SetDescription(v string) *ValidateTemplateResponseBody {
	s.Description = &v
	return s
}

func (s *ValidateTemplateResponseBody) SetOutputs(v []*ValidateTemplateResponseBodyOutputs) *ValidateTemplateResponseBody {
	s.Outputs = v
	return s
}

func (s *ValidateTemplateResponseBody) SetParameters(v []map[string]interface{}) *ValidateTemplateResponseBody {
	s.Parameters = v
	return s
}

func (s *ValidateTemplateResponseBody) SetRequestId(v string) *ValidateTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateTemplateResponseBody) SetResourceTypes(v *ValidateTemplateResponseBodyResourceTypes) *ValidateTemplateResponseBody {
	s.ResourceTypes = v
	return s
}

func (s *ValidateTemplateResponseBody) SetResources(v []*ValidateTemplateResponseBodyResources) *ValidateTemplateResponseBody {
	s.Resources = v
	return s
}

func (s *ValidateTemplateResponseBody) SetUpdateInfo(v *ValidateTemplateResponseBodyUpdateInfo) *ValidateTemplateResponseBody {
	s.UpdateInfo = v
	return s
}

type ValidateTemplateResponseBodyOutputs struct {
	// The description of the template output.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The alias of the template output.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The name of the template output.
	OutputKey *string `json:"OutputKey,omitempty" xml:"OutputKey,omitempty"`
}

func (s ValidateTemplateResponseBodyOutputs) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateResponseBodyOutputs) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponseBodyOutputs) SetDescription(v string) *ValidateTemplateResponseBodyOutputs {
	s.Description = &v
	return s
}

func (s *ValidateTemplateResponseBodyOutputs) SetLabel(v string) *ValidateTemplateResponseBodyOutputs {
	s.Label = &v
	return s
}

func (s *ValidateTemplateResponseBodyOutputs) SetOutputKey(v string) *ValidateTemplateResponseBodyOutputs {
	s.OutputKey = &v
	return s
}

type ValidateTemplateResponseBodyResourceTypes struct {
	// The DataSource resource types that are used in the template. The value is deduplicated.
	DataSources []*string `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The regular resource types that are used in the template. The value is deduplicated.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
}

func (s ValidateTemplateResponseBodyResourceTypes) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateResponseBodyResourceTypes) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponseBodyResourceTypes) SetDataSources(v []*string) *ValidateTemplateResponseBodyResourceTypes {
	s.DataSources = v
	return s
}

func (s *ValidateTemplateResponseBodyResourceTypes) SetResources(v []*string) *ValidateTemplateResponseBodyResourceTypes {
	s.Resources = v
	return s
}

type ValidateTemplateResponseBodyResources struct {
	// The pattern in which the logical IDs of regular resources are formed.
	//
	// If resources are defined in a ROS template, the following rules apply:
	//
	// *   Resource whose definition does not contain `Count`: If the resource name defined in the template is `server`, the values of LogicalResourceIdPattern and `ResourcePath` are both `server`.``
	// *   Resource whose definition contains `Count`: If the resource name defined in the template is `server`, the value of LogicalResourceIdPattern is `server[*]`, and the value of `ResourcePath` is `server`.
	//
	// If resources and [modules](https://www.terraform.io/language/modules) are defined in a Terraform template, the following rules apply:
	//
	// *   Resource and module whose definitions do not contain [`count`](https://www.terraform.io/language/meta-arguments/count) or [`for_each`](https://www.terraform.io/language/meta-arguments/for_each): If the resource name defined in the template is `server`, the values of LogicalResourceIdPattern and `ResourcePath` are both `server`.``
	// *   Resource and module whose definitions contain [`count`](https://www.terraform.io/language/meta-arguments/count) or [`for_each`](https://www.terraform.io/language/meta-arguments/for_each): If the resource name defined in the template is `server`, the value of LogicalResourceIdPattern is `server[*]`, and the value of `ResourcePath` is `server`.
	//
	// Examples of LogicalResourceIdPattern for resources in a Terraform template:
	//
	// *   Valid values of LogicalResourceIdPattern if a resource belongs to the root module:
	//
	//     *   `server`: In this case, `count` and `for_each` are not contained in the resource. The value of `ResourcePath` is `server`.
	//     *   `server[*]`: In this case, `count` or `for_each` is contained in the resource. The value of `ResourcePath` is `server`.
	//
	// *   Valid values of LogicalResourceIdPattern if a resource belongs to a child module:
	//
	//     *   `app.server`: In this case, `count` and `for_each` are not contained in the `app` module and the `server` resource. The value of `ResourcePath` is `app.server`.````
	//     *   `app.server[*]`: In this case, `count` or `for_each` is contained in the `server` resource, but `count` and `for_each` are not contained in the `app` module. The value of `ResourcePath` is `app.server`.
	//     *   `app[*].server`: In this case, `count` or `for_each` is contained in the `app` module, but `count` and `for_each` are not contained in the `server` resource. The value of `ResourcePath` is `app.server`.
	//     *   `app[*].server[*]`: In this case, `count` or `for_each` is contained in the `app` module and the `server` resource. The value of `ResourcePath` is `app.server`.````
	//     *   `app.app_group[*].server`: In this case, `count` or `for_each` is contained in the `app_group` module, but `count` and `for_each` are not contained in the `app` module and the `server` resource. The value of `ResourcePath` is `app.app_group.server`. The `app_group` module is a child module of the `app` module.````
	LogicalResourceIdPattern *string `json:"LogicalResourceIdPattern,omitempty" xml:"LogicalResourceIdPattern,omitempty"`
	// The path of the regular resource. In most cases, the path of a regular resource is the same as the resource name.
	ResourcePath *string `json:"ResourcePath,omitempty" xml:"ResourcePath,omitempty"`
	// The regular resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ValidateTemplateResponseBodyResources) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateResponseBodyResources) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponseBodyResources) SetLogicalResourceIdPattern(v string) *ValidateTemplateResponseBodyResources {
	s.LogicalResourceIdPattern = &v
	return s
}

func (s *ValidateTemplateResponseBodyResources) SetResourcePath(v string) *ValidateTemplateResponseBodyResources {
	s.ResourcePath = &v
	return s
}

func (s *ValidateTemplateResponseBodyResources) SetResourceType(v string) *ValidateTemplateResponseBodyResources {
	s.ResourceType = &v
	return s
}

type ValidateTemplateResponseBodyUpdateInfo struct {
	// The parameters that can be modified.
	ParametersAllowedToBeModified []*string `json:"ParametersAllowedToBeModified,omitempty" xml:"ParametersAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters whose changes cause service interruptions.
	//
	// > - This parameter is supported only for a small number of resource types.
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersCauseInterruptionIfModified []*string `json:"ParametersCauseInterruptionIfModified,omitempty" xml:"ParametersCauseInterruptionIfModified,omitempty" type:"Repeated"`
	// The parameters whose changes trigger replacement updates for resources.
	//
	// > -  This parameter can be returned only if the value of UpdateInfoOptions contains EnableReplacement.
	// > -  This parameter is valid only for updates on ROS stacks.
	ParametersCauseReplacementIfModified []*string `json:"ParametersCauseReplacementIfModified,omitempty" xml:"ParametersCauseReplacementIfModified,omitempty" type:"Repeated"`
	// The parameters that can be modified under specific conditions.
	ParametersConditionallyAllowedToBeModified []*string `json:"ParametersConditionallyAllowedToBeModified,omitempty" xml:"ParametersConditionallyAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters whose changes cause service interruptions under specific conditions.
	//
	// > - This parameter is supported only for a small number of resource types.
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersConditionallyCauseInterruptionIfModified []*string `json:"ParametersConditionallyCauseInterruptionIfModified,omitempty" xml:"ParametersConditionallyCauseInterruptionIfModified,omitempty" type:"Repeated"`
	// The parameters whose changes trigger replacement updates for resources under specific conditions.
	//
	// > - This parameter can be returned only if the value of UpdateInfoOptions contains EnableReplacement.
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersConditionallyCauseReplacementIfModified []*string `json:"ParametersConditionallyCauseReplacementIfModified,omitempty" xml:"ParametersConditionallyCauseReplacementIfModified,omitempty" type:"Repeated"`
	// The parameters that cannot be modified.
	ParametersNotAllowedToBeModified []*string `json:"ParametersNotAllowedToBeModified,omitempty" xml:"ParametersNotAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters that can be modified under uncertain conditions.
	ParametersUncertainlyAllowedToBeModified []*string `json:"ParametersUncertainlyAllowedToBeModified,omitempty" xml:"ParametersUncertainlyAllowedToBeModified,omitempty" type:"Repeated"`
	// The parameters whose changes cause service interruptions under uncertain conditions.
	//
	// > - This parameter is supported only for a small number of resource types.
	// > - This parameter is valid only for updates on ROS stacks.
	ParametersUncertainlyCauseInterruptionIfModified []*string `json:"ParametersUncertainlyCauseInterruptionIfModified,omitempty" xml:"ParametersUncertainlyCauseInterruptionIfModified,omitempty" type:"Repeated"`
	// The parameters whose changes trigger replacement updates for resources under uncertain conditions.
	//
	// > -  This parameter can be returned only if the value of UpdateInfoOptions contains EnableReplacement.
	// > -  This parameter is valid only for updates on ROS stacks.
	ParametersUncertainlyCauseReplacementIfModified []*string `json:"ParametersUncertainlyCauseReplacementIfModified,omitempty" xml:"ParametersUncertainlyCauseReplacementIfModified,omitempty" type:"Repeated"`
}

func (s ValidateTemplateResponseBodyUpdateInfo) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateResponseBodyUpdateInfo) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersAllowedToBeModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersAllowedToBeModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersCauseInterruptionIfModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersCauseInterruptionIfModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersCauseReplacementIfModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersCauseReplacementIfModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersConditionallyAllowedToBeModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersConditionallyAllowedToBeModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersConditionallyCauseInterruptionIfModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersConditionallyCauseInterruptionIfModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersConditionallyCauseReplacementIfModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersConditionallyCauseReplacementIfModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersNotAllowedToBeModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersNotAllowedToBeModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersUncertainlyAllowedToBeModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersUncertainlyAllowedToBeModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersUncertainlyCauseInterruptionIfModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersUncertainlyCauseInterruptionIfModified = v
	return s
}

func (s *ValidateTemplateResponseBodyUpdateInfo) SetParametersUncertainlyCauseReplacementIfModified(v []*string) *ValidateTemplateResponseBodyUpdateInfo {
	s.ParametersUncertainlyCauseReplacementIfModified = v
	return s
}

type ValidateTemplateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValidateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateTemplateResponse) GoString() string {
	return s.String()
}

func (s *ValidateTemplateResponse) SetHeaders(v map[string]*string) *ValidateTemplateResponse {
	s.Headers = v
	return s
}

func (s *ValidateTemplateResponse) SetStatusCode(v int32) *ValidateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateTemplateResponse) SetBody(v *ValidateTemplateResponseBody) *ValidateTemplateResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ros"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelStackOperationWithOptions(request *CancelStackOperationRequest, runtime *util.RuntimeOptions) (_result *CancelStackOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowedStackOperations)) {
		query["AllowedStackOperations"] = request.AllowedStackOperations
	}

	if !tea.BoolValue(util.IsUnset(request.CancelType)) {
		query["CancelType"] = request.CancelType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelStackOperation"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelStackOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelStackOperation(request *CancelStackOperationRequest) (_result *CancelStackOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelStackOperationResponse{}
	_body, _err := client.CancelStackOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelUpdateStackWithOptions(request *CancelUpdateStackRequest, runtime *util.RuntimeOptions) (_result *CancelUpdateStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CancelType)) {
		query["CancelType"] = request.CancelType
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelUpdateStack"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelUpdateStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelUpdateStack(request *CancelUpdateStackRequest) (_result *CancelUpdateStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelUpdateStackResponse{}
	_body, _err := client.CancelUpdateStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to continue to create a stack after the stack fails to be created. In this example, the stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` is created in the China (Hangzhou) region.
 *
 * @param request ContinueCreateStackRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ContinueCreateStackResponse
 */
func (client *Client) ContinueCreateStackWithOptions(request *ContinueCreateStackRequest, runtime *util.RuntimeOptions) (_result *ContinueCreateStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.Mode)) {
		query["Mode"] = request.Mode
	}

	if !tea.BoolValue(util.IsUnset(request.Parallelism)) {
		query["Parallelism"] = request.Parallelism
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RamRoleName)) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.RecreatingOptions)) {
		query["RecreatingOptions"] = request.RecreatingOptions
	}

	if !tea.BoolValue(util.IsUnset(request.RecreatingResources)) {
		query["RecreatingResources"] = request.RecreatingResources
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		query["TemplateBody"] = request.TemplateBody
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ContinueCreateStack"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ContinueCreateStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to continue to create a stack after the stack fails to be created. In this example, the stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` is created in the China (Hangzhou) region.
 *
 * @param request ContinueCreateStackRequest
 * @return ContinueCreateStackResponse
 */
func (client *Client) ContinueCreateStack(request *ContinueCreateStackRequest) (_result *ContinueCreateStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ContinueCreateStackResponse{}
	_body, _err := client.ContinueCreateStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Limits
 * You can create and execute change sets to update running stacks. For more information about change sets, see [Overview](~~156038~~).
 * ### Usage notes
 * *   A stack can contain up to 20 change sets.
 * *   Change sets reflect only the changes of stacks. Change sets do not reflect whether stacks are updated.
 * *   You cannot use change sets to check the following items: whether the upper limit of your Alibaba Cloud account is reached, whether resources that cannot be updated are updated, and whether your Alibaba Cloud account has permissions to modify resources. These items may cause stack updates to fail. If stacks fail to be updated, Resource Orchestration Service (ROS) rolls back the resources to the original status.
 * In this topic, a change set named `MyChangeSet` is created in the `China (Hangzhou)` region. In this example, the template of the stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` is updated to `{"ROSTemplateFormatVersion":"2015-09-01"}`.
 *
 * @param request CreateChangeSetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateChangeSetResponse
 */
func (client *Client) CreateChangeSetWithOptions(request *CreateChangeSetRequest, runtime *util.RuntimeOptions) (_result *CreateChangeSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeSetName)) {
		query["ChangeSetName"] = request.ChangeSetName
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeSetType)) {
		query["ChangeSetType"] = request.ChangeSetType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisableRollback)) {
		query["DisableRollback"] = request.DisableRollback
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationURLs)) {
		query["NotificationURLs"] = request.NotificationURLs
	}

	if !tea.BoolValue(util.IsUnset(request.Parallelism)) {
		query["Parallelism"] = request.Parallelism
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RamRoleName)) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplacementOption)) {
		query["ReplacementOption"] = request.ReplacementOption
	}

	if !tea.BoolValue(util.IsUnset(request.ResourcesToImport)) {
		query["ResourcesToImport"] = request.ResourcesToImport
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.StackName)) {
		query["StackName"] = request.StackName
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyBody)) {
		query["StackPolicyBody"] = request.StackPolicyBody
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyDuringUpdateBody)) {
		query["StackPolicyDuringUpdateBody"] = request.StackPolicyDuringUpdateBody
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyDuringUpdateURL)) {
		query["StackPolicyDuringUpdateURL"] = request.StackPolicyDuringUpdateURL
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyURL)) {
		query["StackPolicyURL"] = request.StackPolicyURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchId)) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TimeoutInMinutes)) {
		query["TimeoutInMinutes"] = request.TimeoutInMinutes
	}

	if !tea.BoolValue(util.IsUnset(request.UsePreviousParameters)) {
		query["UsePreviousParameters"] = request.UsePreviousParameters
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateChangeSet"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateChangeSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Limits
 * You can create and execute change sets to update running stacks. For more information about change sets, see [Overview](~~156038~~).
 * ### Usage notes
 * *   A stack can contain up to 20 change sets.
 * *   Change sets reflect only the changes of stacks. Change sets do not reflect whether stacks are updated.
 * *   You cannot use change sets to check the following items: whether the upper limit of your Alibaba Cloud account is reached, whether resources that cannot be updated are updated, and whether your Alibaba Cloud account has permissions to modify resources. These items may cause stack updates to fail. If stacks fail to be updated, Resource Orchestration Service (ROS) rolls back the resources to the original status.
 * In this topic, a change set named `MyChangeSet` is created in the `China (Hangzhou)` region. In this example, the template of the stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` is updated to `{"ROSTemplateFormatVersion":"2015-09-01"}`.
 *
 * @param request CreateChangeSetRequest
 * @return CreateChangeSetResponse
 */
func (client *Client) CreateChangeSet(request *CreateChangeSetRequest) (_result *CreateChangeSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateChangeSetResponse{}
	_body, _err := client.CreateChangeSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDiagnosticWithOptions(request *CreateDiagnosticRequest, runtime *util.RuntimeOptions) (_result *CreateDiagnosticResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiagnosticKey)) {
		query["DiagnosticKey"] = request.DiagnosticKey
	}

	if !tea.BoolValue(util.IsUnset(request.DiagnosticType)) {
		query["DiagnosticType"] = request.DiagnosticType
	}

	if !tea.BoolValue(util.IsUnset(request.Product)) {
		query["Product"] = request.Product
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDiagnostic"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDiagnosticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDiagnostic(request *CreateDiagnosticRequest) (_result *CreateDiagnosticResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDiagnosticResponse{}
	_body, _err := client.CreateDiagnosticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * A stack is a collection of ROS resources that you can manage as a single unit. To create a collection of resources, you can create a stack. For more information about stacks, see [Overview](~~172973~~).\\
 * When you call the operation, take note of the following limits:
 * *   You can create up to 200 stacks within an Alibaba Cloud account.
 * *   You can create up to 200 resources in a stack.
 * This topic provides an example on how to create a stack named `MyStack` in the China (Hangzhou) region by using a template. In this example, `TemplateBody` is set to `{"ROSTemplateFormatVersion":"2015-09-01"}`.
 *
 * @param request CreateStackRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateStackResponse
 */
func (client *Client) CreateStackWithOptions(request *CreateStackRequest, runtime *util.RuntimeOptions) (_result *CreateStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CreateOption)) {
		query["CreateOption"] = request.CreateOption
	}

	if !tea.BoolValue(util.IsUnset(request.CreateOptions)) {
		query["CreateOptions"] = request.CreateOptions
	}

	if !tea.BoolValue(util.IsUnset(request.DeletionProtection)) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !tea.BoolValue(util.IsUnset(request.DisableRollback)) {
		query["DisableRollback"] = request.DisableRollback
	}

	if !tea.BoolValue(util.IsUnset(request.NotificationURLs)) {
		query["NotificationURLs"] = request.NotificationURLs
	}

	if !tea.BoolValue(util.IsUnset(request.Parallelism)) {
		query["Parallelism"] = request.Parallelism
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RamRoleName)) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StackName)) {
		query["StackName"] = request.StackName
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyBody)) {
		query["StackPolicyBody"] = request.StackPolicyBody
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyURL)) {
		query["StackPolicyURL"] = request.StackPolicyURL
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchId)) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchRegionId)) {
		query["TemplateScratchRegionId"] = request.TemplateScratchRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TimeoutInMinutes)) {
		query["TimeoutInMinutes"] = request.TimeoutInMinutes
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStack"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * A stack is a collection of ROS resources that you can manage as a single unit. To create a collection of resources, you can create a stack. For more information about stacks, see [Overview](~~172973~~).\\
 * When you call the operation, take note of the following limits:
 * *   You can create up to 200 stacks within an Alibaba Cloud account.
 * *   You can create up to 200 resources in a stack.
 * This topic provides an example on how to create a stack named `MyStack` in the China (Hangzhou) region by using a template. In this example, `TemplateBody` is set to `{"ROSTemplateFormatVersion":"2015-09-01"}`.
 *
 * @param request CreateStackRequest
 * @return CreateStackResponse
 */
func (client *Client) CreateStack(request *CreateStackRequest) (_result *CreateStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStackResponse{}
	_body, _err := client.CreateStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * A stack group is a collection of ROS stacks that you can manage as a unit. You can use an ROS template of a stack group to create stacks within Alibaba Cloud accounts across regions.
 * You can create a stack group that is granted self-managed or service-managed permissions:
 * *   If you use an Alibaba Cloud account to create a self-managed stack group, the administrator account and the execution account are Alibaba Cloud accounts.
 * *   If you enable a resource directory and use the management account or a delegated administrator account of the resource directory to create a service-managed stack group, the administrator account is the management account or delegated administrator account, and the execution account is a member account of the resource directory.
 * For more information about stack groups, see [Overview](~~154578~~).
 * In this topic, a stack group named `MyStackGroup` is created in the `China (Hangzhou)` region and granted the self-managed permissions. In this example, the template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****` is used.
 *
 * @param tmpReq CreateStackGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateStackGroupResponse
 */
func (client *Client) CreateStackGroupWithOptions(tmpReq *CreateStackGroupRequest, runtime *util.RuntimeOptions) (_result *CreateStackGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateStackGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AutoDeployment)) {
		request.AutoDeploymentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AutoDeployment, tea.String("AutoDeployment"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdministrationRoleName)) {
		query["AdministrationRoleName"] = request.AdministrationRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.AutoDeploymentShrink)) {
		query["AutoDeployment"] = request.AutoDeploymentShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Capabilities)) {
		query["Capabilities"] = request.Capabilities
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionRoleName)) {
		query["ExecutionRoleName"] = request.ExecutionRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionModel)) {
		query["PermissionModel"] = request.PermissionModel
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupName)) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStackGroup"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStackGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * A stack group is a collection of ROS stacks that you can manage as a unit. You can use an ROS template of a stack group to create stacks within Alibaba Cloud accounts across regions.
 * You can create a stack group that is granted self-managed or service-managed permissions:
 * *   If you use an Alibaba Cloud account to create a self-managed stack group, the administrator account and the execution account are Alibaba Cloud accounts.
 * *   If you enable a resource directory and use the management account or a delegated administrator account of the resource directory to create a service-managed stack group, the administrator account is the management account or delegated administrator account, and the execution account is a member account of the resource directory.
 * For more information about stack groups, see [Overview](~~154578~~).
 * In this topic, a stack group named `MyStackGroup` is created in the `China (Hangzhou)` region and granted the self-managed permissions. In this example, the template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****` is used.
 *
 * @param request CreateStackGroupRequest
 * @return CreateStackGroupResponse
 */
func (client *Client) CreateStackGroup(request *CreateStackGroupRequest) (_result *CreateStackGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStackGroupResponse{}
	_body, _err := client.CreateStackGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you call this operation, make sure that a stack group is created. For more information, see [CreateStackGroup](~~151333~~).
 * In this topic, the stack group named `MyStackGroup` is used. The stack group is created in the China (Hangzhou) region and granted the self-managed permissions. In this example, stacks are created by using Alibaba Cloud accounts whose IDs are `151266687691****` and `141261387191****` in the China (Hangzhou) region and China (Beijing) region.
 *
 * @param tmpReq CreateStackInstancesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateStackInstancesResponse
 */
func (client *Client) CreateStackInstancesWithOptions(tmpReq *CreateStackInstancesRequest, runtime *util.RuntimeOptions) (_result *CreateStackInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateStackInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AccountIds)) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, tea.String("AccountIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DeploymentTargets)) {
		request.DeploymentTargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeploymentTargets, tea.String("DeploymentTargets"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OperationPreferences)) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, tea.String("OperationPreferences"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("RegionIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountIdsShrink)) {
		query["AccountIds"] = request.AccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeploymentTargetsShrink)) {
		query["DeploymentTargets"] = request.DeploymentTargetsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DisableRollback)) {
		query["DisableRollback"] = request.DisableRollback
	}

	if !tea.BoolValue(util.IsUnset(request.OperationDescription)) {
		query["OperationDescription"] = request.OperationDescription
	}

	if !tea.BoolValue(util.IsUnset(request.OperationPreferencesShrink)) {
		query["OperationPreferences"] = request.OperationPreferencesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ParameterOverrides)) {
		query["ParameterOverrides"] = request.ParameterOverrides
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionIdsShrink)) {
		query["RegionIds"] = request.RegionIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupName)) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.TimeoutInMinutes)) {
		query["TimeoutInMinutes"] = request.TimeoutInMinutes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStackInstances"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStackInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you call this operation, make sure that a stack group is created. For more information, see [CreateStackGroup](~~151333~~).
 * In this topic, the stack group named `MyStackGroup` is used. The stack group is created in the China (Hangzhou) region and granted the self-managed permissions. In this example, stacks are created by using Alibaba Cloud accounts whose IDs are `151266687691****` and `141261387191****` in the China (Hangzhou) region and China (Beijing) region.
 *
 * @param request CreateStackInstancesRequest
 * @return CreateStackInstancesResponse
 */
func (client *Client) CreateStackInstances(request *CreateStackInstancesRequest) (_result *CreateStackInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateStackInstancesResponse{}
	_body, _err := client.CreateStackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this topic, a custom template named `MyTemplate` is created in the `cn-hangzhou` region. The `TemplateBody` parameter of the template is set to `{"ROSTemplateFormatVersion": "2015-09-01"}`.
 *
 * @param request CreateTemplateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateTemplateResponse
 */
func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		query["TemplateBody"] = request.TemplateBody
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTemplate"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this topic, a custom template named `MyTemplate` is created in the `cn-hangzhou` region. The `TemplateBody` parameter of the template is set to `{"ROSTemplateFormatVersion": "2015-09-01"}`.
 *
 * @param request CreateTemplateRequest
 * @return CreateTemplateResponse
 */
func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * You can call this operation to create a scenario that allows you to specify a resource scope on a visualized interface and easily replicate and manage the resources that you specify. For more information about scenarios, see [Overview](~~352074~~).
 * In this example, a scenario of the Resource Replication type is created in the China (Hangzhou) region. In the scenario, the virtual private cloud (VPC) whose ID is `vpc-bp1m6fww66xbntjyc****` is replicated.
 *
 * @param tmpReq CreateTemplateScratchRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateTemplateScratchResponse
 */
func (client *Client) CreateTemplateScratchWithOptions(tmpReq *CreateTemplateScratchRequest, runtime *util.RuntimeOptions) (_result *CreateTemplateScratchResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateTemplateScratchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PreferenceParameters)) {
		request.PreferenceParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PreferenceParameters, tea.String("PreferenceParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceResourceGroup)) {
		request.SourceResourceGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceResourceGroup, tea.String("SourceResourceGroup"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceResources)) {
		request.SourceResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceResources, tea.String("SourceResources"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceTag)) {
		request.SourceTagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceTag, tea.String("SourceTag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionMode)) {
		query["ExecutionMode"] = request.ExecutionMode
	}

	if !tea.BoolValue(util.IsUnset(request.LogicalIdStrategy)) {
		query["LogicalIdStrategy"] = request.LogicalIdStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.PreferenceParametersShrink)) {
		query["PreferenceParameters"] = request.PreferenceParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceResourceGroupShrink)) {
		query["SourceResourceGroup"] = request.SourceResourceGroupShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceResourcesShrink)) {
		query["SourceResources"] = request.SourceResourcesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTagShrink)) {
		query["SourceTag"] = request.SourceTagShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchType)) {
		query["TemplateScratchType"] = request.TemplateScratchType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTemplateScratch"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTemplateScratchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * You can call this operation to create a scenario that allows you to specify a resource scope on a visualized interface and easily replicate and manage the resources that you specify. For more information about scenarios, see [Overview](~~352074~~).
 * In this example, a scenario of the Resource Replication type is created in the China (Hangzhou) region. In the scenario, the virtual private cloud (VPC) whose ID is `vpc-bp1m6fww66xbntjyc****` is replicated.
 *
 * @param request CreateTemplateScratchRequest
 * @return CreateTemplateScratchResponse
 */
func (client *Client) CreateTemplateScratch(request *CreateTemplateScratchRequest) (_result *CreateTemplateScratchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTemplateScratchResponse{}
	_body, _err := client.CreateTemplateScratchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Before you call this operation, make sure that the following requirements are met:
 *     *   The status of the change set is CREATE_COMPLETE, CREATE_FAILED, or DELETE_FAILED.
 *     *   The execution status is UNAVAILABLE or AVAILABLE.
 * *   After a change set is executed, other change sets associated with the same stack as this change set are also deleted.
 * *   After a stack is deleted, change sets associated with the stack are deleted.
 * *   If a change set of the CREATE type is deleted, you must delete stacks associated with the change set.
 * In this example, a change set whose ID is `1f6521a4-05af-4975-afe9-bc4b45ad****` is deleted. The change set is created in the China (Hangzhou) region.
 *
 * @param request DeleteChangeSetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteChangeSetResponse
 */
func (client *Client) DeleteChangeSetWithOptions(request *DeleteChangeSetRequest, runtime *util.RuntimeOptions) (_result *DeleteChangeSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeSetId)) {
		query["ChangeSetId"] = request.ChangeSetId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteChangeSet"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteChangeSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Before you call this operation, make sure that the following requirements are met:
 *     *   The status of the change set is CREATE_COMPLETE, CREATE_FAILED, or DELETE_FAILED.
 *     *   The execution status is UNAVAILABLE or AVAILABLE.
 * *   After a change set is executed, other change sets associated with the same stack as this change set are also deleted.
 * *   After a stack is deleted, change sets associated with the stack are deleted.
 * *   If a change set of the CREATE type is deleted, you must delete stacks associated with the change set.
 * In this example, a change set whose ID is `1f6521a4-05af-4975-afe9-bc4b45ad****` is deleted. The change set is created in the China (Hangzhou) region.
 *
 * @param request DeleteChangeSetRequest
 * @return DeleteChangeSetResponse
 */
func (client *Client) DeleteChangeSet(request *DeleteChangeSetRequest) (_result *DeleteChangeSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteChangeSetResponse{}
	_body, _err := client.DeleteChangeSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDiagnosticWithOptions(request *DeleteDiagnosticRequest, runtime *util.RuntimeOptions) (_result *DeleteDiagnosticResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReportId)) {
		query["ReportId"] = request.ReportId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDiagnostic"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDiagnosticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDiagnostic(request *DeleteDiagnosticRequest) (_result *DeleteDiagnosticResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDiagnosticResponse{}
	_body, _err := client.DeleteDiagnosticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStackWithOptions(request *DeleteStackRequest, runtime *util.RuntimeOptions) (_result *DeleteStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeleteOptions)) {
		query["DeleteOptions"] = request.DeleteOptions
	}

	if !tea.BoolValue(util.IsUnset(request.RamRoleName)) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RetainAllResources)) {
		query["RetainAllResources"] = request.RetainAllResources
	}

	if !tea.BoolValue(util.IsUnset(request.RetainResources)) {
		query["RetainResources"] = request.RetainResources
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteStack"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStack(request *DeleteStackRequest) (_result *DeleteStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStackResponse{}
	_body, _err := client.DeleteStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * A stack group can be deleted only when the stack group does not contain stacks. You can call the [DeleteStackInstances](~~151715~~) operation to delete stacks.
 * This topic provides an example on how to delete a stack group. In this example, a stack group named `MyStackGroup` in the China (Hangzhou) region is deleted.
 *
 * @param request DeleteStackGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteStackGroupResponse
 */
func (client *Client) DeleteStackGroupWithOptions(request *DeleteStackGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteStackGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupName)) {
		query["StackGroupName"] = request.StackGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteStackGroup"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteStackGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * A stack group can be deleted only when the stack group does not contain stacks. You can call the [DeleteStackInstances](~~151715~~) operation to delete stacks.
 * This topic provides an example on how to delete a stack group. In this example, a stack group named `MyStackGroup` in the China (Hangzhou) region is deleted.
 *
 * @param request DeleteStackGroupRequest
 * @return DeleteStackGroupResponse
 */
func (client *Client) DeleteStackGroup(request *DeleteStackGroupRequest) (_result *DeleteStackGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStackGroupResponse{}
	_body, _err := client.DeleteStackGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this topic, the stack group named `MyStackGroup` that is created in the China (Hangzhou) region is used. In this example, the stacks of the stack group that are deployed in the China (Beijing) region by using the Alibaba Cloud account whose ID is `151266687691****` are deleted.
 *
 * @param tmpReq DeleteStackInstancesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteStackInstancesResponse
 */
func (client *Client) DeleteStackInstancesWithOptions(tmpReq *DeleteStackInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteStackInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteStackInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AccountIds)) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, tea.String("AccountIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DeploymentTargets)) {
		request.DeploymentTargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeploymentTargets, tea.String("DeploymentTargets"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OperationPreferences)) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, tea.String("OperationPreferences"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("RegionIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountIdsShrink)) {
		query["AccountIds"] = request.AccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeploymentTargetsShrink)) {
		query["DeploymentTargets"] = request.DeploymentTargetsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OperationDescription)) {
		query["OperationDescription"] = request.OperationDescription
	}

	if !tea.BoolValue(util.IsUnset(request.OperationPreferencesShrink)) {
		query["OperationPreferences"] = request.OperationPreferencesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionIdsShrink)) {
		query["RegionIds"] = request.RegionIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RetainStacks)) {
		query["RetainStacks"] = request.RetainStacks
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupName)) {
		query["StackGroupName"] = request.StackGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteStackInstances"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteStackInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this topic, the stack group named `MyStackGroup` that is created in the China (Hangzhou) region is used. In this example, the stacks of the stack group that are deployed in the China (Beijing) region by using the Alibaba Cloud account whose ID is `151266687691****` are deleted.
 *
 * @param request DeleteStackInstancesRequest
 * @return DeleteStackInstancesResponse
 */
func (client *Client) DeleteStackInstances(request *DeleteStackInstancesRequest) (_result *DeleteStackInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteStackInstancesResponse{}
	_body, _err := client.DeleteStackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If a template is shared with other Alibaba Cloud accounts, you must unshare the template before you delete it.
 *
 * @param request DeleteTemplateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteTemplateResponse
 */
func (client *Client) DeleteTemplateWithOptions(request *DeleteTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTemplate"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If a template is shared with other Alibaba Cloud accounts, you must unshare the template before you delete it.
 *
 * @param request DeleteTemplateRequest
 * @return DeleteTemplateResponse
 */
func (client *Client) DeleteTemplate(request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this topic, a scenario whose ID is `ts-4f83704400994409****` is deleted in the China (Hangzhou) region.
 *
 * @param request DeleteTemplateScratchRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteTemplateScratchResponse
 */
func (client *Client) DeleteTemplateScratchWithOptions(request *DeleteTemplateScratchRequest, runtime *util.RuntimeOptions) (_result *DeleteTemplateScratchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchId)) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTemplateScratch"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTemplateScratchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this topic, a scenario whose ID is `ts-4f83704400994409****` is deleted in the China (Hangzhou) region.
 *
 * @param request DeleteTemplateScratchRequest
 * @return DeleteTemplateScratchResponse
 */
func (client *Client) DeleteTemplateScratch(request *DeleteTemplateScratchRequest) (_result *DeleteTemplateScratchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTemplateScratchResponse{}
	_body, _err := client.DeleteTemplateScratchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   If you delete a resource type, you can no longer use the resource type in Resource Orchestration Service (ROS).
 * *   If you delete a version of a resource type, you can no longer use the version in ROS.
 * *   If a resource type has only one version, you can delete the version by calling the operation. If a resource type has more than one version, you must manually delete the remaining versions.
 * *   When a resource type has more than one version, you cannot delete the default version by calling the operation.
 * *   When a resource type has only one version, you can delete the resource type and the version by calling the operation.
 *
 * @param request DeregisterResourceTypeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeregisterResourceTypeResponse
 */
func (client *Client) DeregisterResourceTypeWithOptions(request *DeregisterResourceTypeRequest, runtime *util.RuntimeOptions) (_result *DeregisterResourceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeregisterResourceType"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeregisterResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   If you delete a resource type, you can no longer use the resource type in Resource Orchestration Service (ROS).
 * *   If you delete a version of a resource type, you can no longer use the version in ROS.
 * *   If a resource type has only one version, you can delete the version by calling the operation. If a resource type has more than one version, you must manually delete the remaining versions.
 * *   When a resource type has more than one version, you cannot delete the default version by calling the operation.
 * *   When a resource type has only one version, you can delete the resource type and the version by calling the operation.
 *
 * @param request DeregisterResourceTypeRequest
 * @return DeregisterResourceTypeResponse
 */
func (client *Client) DeregisterResourceType(request *DeregisterResourceTypeRequest) (_result *DeregisterResourceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeregisterResourceTypeResponse{}
	_body, _err := client.DeregisterResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		query["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectStackDriftWithOptions(request *DetectStackDriftRequest, runtime *util.RuntimeOptions) (_result *DetectStackDriftResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.LogicalResourceId)) {
		query["LogicalResourceId"] = request.LogicalResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectStackDrift"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectStackDriftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectStackDrift(request *DetectStackDriftRequest) (_result *DetectStackDriftResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectStackDriftResponse{}
	_body, _err := client.DetectStackDriftWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectStackGroupDriftWithOptions(tmpReq *DetectStackGroupDriftRequest, runtime *util.RuntimeOptions) (_result *DetectStackGroupDriftResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetectStackGroupDriftShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OperationPreferences)) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, tea.String("OperationPreferences"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperationPreferencesShrink)) {
		query["OperationPreferences"] = request.OperationPreferencesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupName)) {
		query["StackGroupName"] = request.StackGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectStackGroupDrift"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectStackGroupDriftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectStackGroupDrift(request *DetectStackGroupDriftRequest) (_result *DetectStackGroupDriftResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectStackGroupDriftResponse{}
	_body, _err := client.DetectStackGroupDriftWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectStackResourceDriftWithOptions(request *DetectStackResourceDriftRequest, runtime *util.RuntimeOptions) (_result *DetectStackResourceDriftResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.LogicalResourceId)) {
		query["LogicalResourceId"] = request.LogicalResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DetectStackResourceDrift"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectStackResourceDriftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectStackResourceDrift(request *DetectStackResourceDriftRequest) (_result *DetectStackResourceDriftResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetectStackResourceDriftResponse{}
	_body, _err := client.DetectStackResourceDriftWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the change set whose ID is `1f6521a4-05af-4975-afe9-bc4b45ad****` is executed. The change set is created in the `China (Hangzhou)` region.
 *
 * @param request ExecuteChangeSetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ExecuteChangeSetResponse
 */
func (client *Client) ExecuteChangeSetWithOptions(request *ExecuteChangeSetRequest, runtime *util.RuntimeOptions) (_result *ExecuteChangeSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeSetId)) {
		query["ChangeSetId"] = request.ChangeSetId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteChangeSet"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteChangeSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the change set whose ID is `1f6521a4-05af-4975-afe9-bc4b45ad****` is executed. The change set is created in the `China (Hangzhou)` region.
 *
 * @param request ExecuteChangeSetRequest
 * @return ExecuteChangeSetResponse
 */
func (client *Client) ExecuteChangeSet(request *ExecuteChangeSetRequest) (_result *ExecuteChangeSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteChangeSetResponse{}
	_body, _err := client.ExecuteChangeSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, a template is created for the scenario whose ID is `ts-aa9c62feab844a6b****`. The scenario is of the Resource Management type and resides in the China (Hangzhou) region.
 *
 * @param request GenerateTemplateByScratchRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GenerateTemplateByScratchResponse
 */
func (client *Client) GenerateTemplateByScratchWithOptions(request *GenerateTemplateByScratchRequest, runtime *util.RuntimeOptions) (_result *GenerateTemplateByScratchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ProvisionRegionId)) {
		query["ProvisionRegionId"] = request.ProvisionRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchId)) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateType)) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateTemplateByScratch"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateTemplateByScratchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, a template is created for the scenario whose ID is `ts-aa9c62feab844a6b****`. The scenario is of the Resource Management type and resides in the China (Hangzhou) region.
 *
 * @param request GenerateTemplateByScratchRequest
 * @return GenerateTemplateByScratchResponse
 */
func (client *Client) GenerateTemplateByScratch(request *GenerateTemplateByScratchRequest) (_result *GenerateTemplateByScratchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateTemplateByScratchResponse{}
	_body, _err := client.GenerateTemplateByScratchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * If the policy information is related to Enterprise Distributed Application Service (EDAS), you must log on to your Alibaba Cloud account and grant the required permissions to the relevant RAM users.
 * In this example, a policy is generated for a template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****`.
 *
 * @param request GenerateTemplatePolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GenerateTemplatePolicyResponse
 */
func (client *Client) GenerateTemplatePolicyWithOptions(request *GenerateTemplatePolicyRequest, runtime *util.RuntimeOptions) (_result *GenerateTemplatePolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationTypes)) {
		query["OperationTypes"] = request.OperationTypes
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		query["TemplateBody"] = request.TemplateBody
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateTemplatePolicy"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateTemplatePolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * If the policy information is related to Enterprise Distributed Application Service (EDAS), you must log on to your Alibaba Cloud account and grant the required permissions to the relevant RAM users.
 * In this example, a policy is generated for a template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****`.
 *
 * @param request GenerateTemplatePolicyRequest
 * @return GenerateTemplatePolicyResponse
 */
func (client *Client) GenerateTemplatePolicy(request *GenerateTemplatePolicyRequest) (_result *GenerateTemplatePolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateTemplatePolicyResponse{}
	_body, _err := client.GenerateTemplatePolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the details of a change set whose ID is `4c11658d-bd47-4dd0-ba64-727edc62****` is queried. The change set is created in the China (Hangzhou) region.
 *
 * @param request GetChangeSetRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetChangeSetResponse
 */
func (client *Client) GetChangeSetWithOptions(request *GetChangeSetRequest, runtime *util.RuntimeOptions) (_result *GetChangeSetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeSetId)) {
		query["ChangeSetId"] = request.ChangeSetId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowTemplate)) {
		query["ShowTemplate"] = request.ShowTemplate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetChangeSet"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetChangeSetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the details of a change set whose ID is `4c11658d-bd47-4dd0-ba64-727edc62****` is queried. The change set is created in the China (Hangzhou) region.
 *
 * @param request GetChangeSetRequest
 * @return GetChangeSetResponse
 */
func (client *Client) GetChangeSet(request *GetChangeSetRequest) (_result *GetChangeSetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetChangeSetResponse{}
	_body, _err := client.GetChangeSetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDiagnosticWithOptions(request *GetDiagnosticRequest, runtime *util.RuntimeOptions) (_result *GetDiagnosticResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ReportId)) {
		query["ReportId"] = request.ReportId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDiagnostic"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDiagnosticResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDiagnostic(request *GetDiagnosticRequest) (_result *GetDiagnosticResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDiagnosticResponse{}
	_body, _err := client.GetDiagnosticWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to query the Terraform hosting, resource cleaner, and scenario features.
 * This topic provides an example on how to query the details of features supported by ROS in the China (Hangzhou) region. The details include Terraform versions, provider versions, and supported resource types.
 * >  In the Examples section, only part of the sample code is provided.
 *
 * @param request GetFeatureDetailsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetFeatureDetailsResponse
 */
func (client *Client) GetFeatureDetailsWithOptions(request *GetFeatureDetailsRequest, runtime *util.RuntimeOptions) (_result *GetFeatureDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Feature)) {
		query["Feature"] = request.Feature
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFeatureDetails"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFeatureDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to query the Terraform hosting, resource cleaner, and scenario features.
 * This topic provides an example on how to query the details of features supported by ROS in the China (Hangzhou) region. The details include Terraform versions, provider versions, and supported resource types.
 * >  In the Examples section, only part of the sample code is provided.
 *
 * @param request GetFeatureDetailsRequest
 * @return GetFeatureDetailsResponse
 */
func (client *Client) GetFeatureDetails(request *GetFeatureDetailsRequest) (_result *GetFeatureDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFeatureDetailsResponse{}
	_body, _err := client.GetFeatureDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * For more information about common request parameters, see [Common parameters](~~131957~~).
 *
 * @param request GetResourceTypeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetResourceTypeResponse
 */
func (client *Client) GetResourceTypeWithOptions(request *GetResourceTypeRequest, runtime *util.RuntimeOptions) (_result *GetResourceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceType"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * For more information about common request parameters, see [Common parameters](~~131957~~).
 *
 * @param request GetResourceTypeRequest
 * @return GetResourceTypeResponse
 */
func (client *Client) GetResourceType(request *GetResourceTypeRequest) (_result *GetResourceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceTypeResponse{}
	_body, _err := client.GetResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetResourceTypeTemplateWithOptions(request *GetResourceTypeTemplateRequest, runtime *util.RuntimeOptions) (_result *GetResourceTypeTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResourceTypeTemplate"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResourceTypeTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetResourceTypeTemplate(request *GetResourceTypeTemplateRequest) (_result *GetResourceTypeTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResourceTypeTemplateResponse{}
	_body, _err := client.GetResourceTypeTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ### Description
 * This topic describes how to query the activation status and the RAM roles of an Alibaba Cloud service. In this example, the Elastic High Performance Computing (E-HPC) service that is deployed in the China (Hangzhou) region is queried.
 * > Make sure that you have the permissions to call the [GetRole](~~28711~~) operation.
 *
 * @param request GetServiceProvisionsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetServiceProvisionsResponse
 */
func (client *Client) GetServiceProvisionsWithOptions(request *GetServiceProvisionsRequest, runtime *util.RuntimeOptions) (_result *GetServiceProvisionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Services)) {
		query["Services"] = request.Services
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		query["TemplateBody"] = request.TemplateBody
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceProvisions"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceProvisionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ### Description
 * This topic describes how to query the activation status and the RAM roles of an Alibaba Cloud service. In this example, the Elastic High Performance Computing (E-HPC) service that is deployed in the China (Hangzhou) region is queried.
 * > Make sure that you have the permissions to call the [GetRole](~~28711~~) operation.
 *
 * @param request GetServiceProvisionsRequest
 * @return GetServiceProvisionsResponse
 */
func (client *Client) GetServiceProvisions(request *GetServiceProvisionsRequest) (_result *GetServiceProvisionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetServiceProvisionsResponse{}
	_body, _err := client.GetServiceProvisionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this topic, the information about a stack whose ID is `c754d2a4-28f1-46df-b557-9586173a****` is queried. The stack is deployed in the China (Hangzhou) region.
 *
 * @param request GetStackRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetStackResponse
 */
func (client *Client) GetStackWithOptions(request *GetStackRequest, runtime *util.RuntimeOptions) (_result *GetStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.LogOption)) {
		query["LogOption"] = request.LogOption
	}

	if !tea.BoolValue(util.IsUnset(request.OutputOption)) {
		query["OutputOption"] = request.OutputOption
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowResourceProgress)) {
		query["ShowResourceProgress"] = request.ShowResourceProgress
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStack"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this topic, the information about a stack whose ID is `c754d2a4-28f1-46df-b557-9586173a****` is queried. The stack is deployed in the China (Hangzhou) region.
 *
 * @param request GetStackRequest
 * @return GetStackResponse
 */
func (client *Client) GetStack(request *GetStackRequest) (_result *GetStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackResponse{}
	_body, _err := client.GetStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this topic, the status of a drift detection operation whose ID is `a7044f0d-6f2e-4128-a307-4524ef88****` is queried. The operation is performed in the China (Hangzhou) region.
 *
 * @param request GetStackDriftDetectionStatusRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetStackDriftDetectionStatusResponse
 */
func (client *Client) GetStackDriftDetectionStatusWithOptions(request *GetStackDriftDetectionStatusRequest, runtime *util.RuntimeOptions) (_result *GetStackDriftDetectionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriftDetectionId)) {
		query["DriftDetectionId"] = request.DriftDetectionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStackDriftDetectionStatus"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStackDriftDetectionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this topic, the status of a drift detection operation whose ID is `a7044f0d-6f2e-4128-a307-4524ef88****` is queried. The operation is performed in the China (Hangzhou) region.
 *
 * @param request GetStackDriftDetectionStatusRequest
 * @return GetStackDriftDetectionStatusResponse
 */
func (client *Client) GetStackDriftDetectionStatus(request *GetStackDriftDetectionStatusRequest) (_result *GetStackDriftDetectionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackDriftDetectionStatusResponse{}
	_body, _err := client.GetStackDriftDetectionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * For more information about common request parameters, see [Common parameters](~~131957~~).
 *
 * @param request GetStackGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetStackGroupResponse
 */
func (client *Client) GetStackGroupWithOptions(request *GetStackGroupRequest, runtime *util.RuntimeOptions) (_result *GetStackGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupId)) {
		query["StackGroupId"] = request.StackGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupName)) {
		query["StackGroupName"] = request.StackGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStackGroup"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStackGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * For more information about common request parameters, see [Common parameters](~~131957~~).
 *
 * @param request GetStackGroupRequest
 * @return GetStackGroupResponse
 */
func (client *Client) GetStackGroup(request *GetStackGroupRequest) (_result *GetStackGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackGroupResponse{}
	_body, _err := client.GetStackGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the information about the stack group operation whose ID is `6da106ca-1784-4a6f-a7e1-e723863d****` is queried. The stack group named `MyStackGroup` is granted self-managed permissions and deployed in the China (Hangzhou) region.
 *
 * @param request GetStackGroupOperationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetStackGroupOperationResponse
 */
func (client *Client) GetStackGroupOperationWithOptions(request *GetStackGroupOperationRequest, runtime *util.RuntimeOptions) (_result *GetStackGroupOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationId)) {
		query["OperationId"] = request.OperationId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStackGroupOperation"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStackGroupOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the information about the stack group operation whose ID is `6da106ca-1784-4a6f-a7e1-e723863d****` is queried. The stack group named `MyStackGroup` is granted self-managed permissions and deployed in the China (Hangzhou) region.
 *
 * @param request GetStackGroupOperationRequest
 * @return GetStackGroupOperationResponse
 */
func (client *Client) GetStackGroupOperation(request *GetStackGroupOperationRequest) (_result *GetStackGroupOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackGroupOperationResponse{}
	_body, _err := client.GetStackGroupOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the information about a stack instance associated with a stack group named `MyStackGroup` is queried. The stack instance is deployed in the China (Beijing) region within the `151266687691****` Alibaba Cloud account. The stack group is granted self-managed permissions and deployed in the China (Hangzhou) region.
 *
 * @param request GetStackInstanceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetStackInstanceResponse
 */
func (client *Client) GetStackInstanceWithOptions(request *GetStackInstanceRequest, runtime *util.RuntimeOptions) (_result *GetStackInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupName)) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.StackInstanceAccountId)) {
		query["StackInstanceAccountId"] = request.StackInstanceAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.StackInstanceRegionId)) {
		query["StackInstanceRegionId"] = request.StackInstanceRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStackInstance"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStackInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the information about a stack instance associated with a stack group named `MyStackGroup` is queried. The stack instance is deployed in the China (Beijing) region within the `151266687691****` Alibaba Cloud account. The stack group is granted self-managed permissions and deployed in the China (Hangzhou) region.
 *
 * @param request GetStackInstanceRequest
 * @return GetStackInstanceResponse
 */
func (client *Client) GetStackInstance(request *GetStackInstanceRequest) (_result *GetStackInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackInstanceResponse{}
	_body, _err := client.GetStackInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the stack policy of a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` is queried. The stack is deployed in the China (Hangzhou) region.
 *
 * @param request GetStackPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetStackPolicyResponse
 */
func (client *Client) GetStackPolicyWithOptions(request *GetStackPolicyRequest, runtime *util.RuntimeOptions) (_result *GetStackPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStackPolicy"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStackPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the stack policy of a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` is queried. The stack is deployed in the China (Hangzhou) region.
 *
 * @param request GetStackPolicyRequest
 * @return GetStackPolicyResponse
 */
func (client *Client) GetStackPolicy(request *GetStackPolicyRequest) (_result *GetStackPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackPolicyResponse{}
	_body, _err := client.GetStackPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * | Http status code | Error code | Error message | Description |
 * | ---------------- | ---------- | ------------- | ----------- |
 * | 404 | ResourceNotFound | The Resource ({name}) could not be found in Stack {stack}. | The error message returned because the specified resource does not exist in the stack. name indicates the resource name. stack indicates the stack name or ID. |
 * | 404 | StackNotFound | The Stack ({name}) could not be found. | The error message returned because the stack does not exist. name indicates the name or ID of the stack. |
 *
 * @param request GetStackResourceRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetStackResourceResponse
 */
func (client *Client) GetStackResourceWithOptions(request *GetStackResourceRequest, runtime *util.RuntimeOptions) (_result *GetStackResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.LogicalResourceId)) {
		query["LogicalResourceId"] = request.LogicalResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceAttributes)) {
		query["ResourceAttributes"] = request.ResourceAttributes
	}

	if !tea.BoolValue(util.IsUnset(request.ShowResourceAttributes)) {
		query["ShowResourceAttributes"] = request.ShowResourceAttributes
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStackResource"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStackResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * | Http status code | Error code | Error message | Description |
 * | ---------------- | ---------- | ------------- | ----------- |
 * | 404 | ResourceNotFound | The Resource ({name}) could not be found in Stack {stack}. | The error message returned because the specified resource does not exist in the stack. name indicates the resource name. stack indicates the stack name or ID. |
 * | 404 | StackNotFound | The Stack ({name}) could not be found. | The error message returned because the stack does not exist. name indicates the name or ID of the stack. |
 *
 * @param request GetStackResourceRequest
 * @return GetStackResourceResponse
 */
func (client *Client) GetStackResource(request *GetStackResourceRequest) (_result *GetStackResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetStackResourceResponse{}
	_body, _err := client.GetStackResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the details of a template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****` is queried. The region ID of the template is `cn-hangzhou`.
 *
 * @param request GetTemplateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetTemplateResponse
 */
func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeSetId)) {
		query["ChangeSetId"] = request.ChangeSetId
	}

	if !tea.BoolValue(util.IsUnset(request.IncludePermission)) {
		query["IncludePermission"] = request.IncludePermission
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeTags)) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupName)) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateStage)) {
		query["TemplateStage"] = request.TemplateStage
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplate"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the details of a template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****` is queried. The region ID of the template is `cn-hangzhou`.
 *
 * @param request GetTemplateRequest
 * @return GetTemplateResponse
 */
func (client *Client) GetTemplate(request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *   For more information about the resources that support price inquiry in Resource Orchestration Service (ROS) templates, see the "**Resource types that support price inquiry**" section of the [Estimate resource prices](~~203165~~) topic.
 * *   For more information about the resources that support price inquiry in Terraform templates, see the "**ROS resources supported by Terraform**" section of the [ROS features and resources supported by Terraform](~~184389~~) topic.
 * The following sample code provides an example on how to query the estimated price of an elastic IP address (EIP) that you want to create based on a template. In this example, the following template is used:
 *     {
 *       "ROSTemplateFormatVersion": "2015-09-01",
 *       "Parameters": {
 *         "Isp": {
 *           "Type": "String",
 *           "Default": "BGP"
 *         },
 *         "Name": {
 *           "Type": "String",
 *           "Default": "test"
 *         },
 *         "Netmode": {
 *           "Type": "String",
 *           "Default": "public"
 *         },
 *         "Bandwidth": {
 *           "Type": "Number",
 *           "Default": 5
 *         }
 *       },
 *       "Resources": {
 *         "NewEip": {
 *           "Type": "ALIYUN::VPC::EIP",
 *           "Properties": {
 *             "InstanceChargeType": "Prepaid",
 *             "PricingCycle": "Month",
 *             "Isp": {
 *               "Ref": "Isp"
 *             },
 *             "Period": 1,
 *             "DeletionProtection": false,
 *             "AutoPay": false,
 *             "Name": {
 *               "Ref": "Name"
 *             },
 *             "InternetChargeType": "PayByTraffic",
 *             "Netmode": {
 *               "Ref": "Netmode"
 *             },
 *             "Bandwidth": {
 *               "Ref": "Bandwidth"
 *             }
 *           }
 *         }
 *       }
 *     }
 *
 * @param request GetTemplateEstimateCostRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetTemplateEstimateCostResponse
 */
func (client *Client) GetTemplateEstimateCostWithOptions(request *GetTemplateEstimateCostRequest, runtime *util.RuntimeOptions) (_result *GetTemplateEstimateCostResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchId)) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchRegionId)) {
		query["TemplateScratchRegionId"] = request.TemplateScratchRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplateEstimateCost"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateEstimateCostResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *   For more information about the resources that support price inquiry in Resource Orchestration Service (ROS) templates, see the "**Resource types that support price inquiry**" section of the [Estimate resource prices](~~203165~~) topic.
 * *   For more information about the resources that support price inquiry in Terraform templates, see the "**ROS resources supported by Terraform**" section of the [ROS features and resources supported by Terraform](~~184389~~) topic.
 * The following sample code provides an example on how to query the estimated price of an elastic IP address (EIP) that you want to create based on a template. In this example, the following template is used:
 *     {
 *       "ROSTemplateFormatVersion": "2015-09-01",
 *       "Parameters": {
 *         "Isp": {
 *           "Type": "String",
 *           "Default": "BGP"
 *         },
 *         "Name": {
 *           "Type": "String",
 *           "Default": "test"
 *         },
 *         "Netmode": {
 *           "Type": "String",
 *           "Default": "public"
 *         },
 *         "Bandwidth": {
 *           "Type": "Number",
 *           "Default": 5
 *         }
 *       },
 *       "Resources": {
 *         "NewEip": {
 *           "Type": "ALIYUN::VPC::EIP",
 *           "Properties": {
 *             "InstanceChargeType": "Prepaid",
 *             "PricingCycle": "Month",
 *             "Isp": {
 *               "Ref": "Isp"
 *             },
 *             "Period": 1,
 *             "DeletionProtection": false,
 *             "AutoPay": false,
 *             "Name": {
 *               "Ref": "Name"
 *             },
 *             "InternetChargeType": "PayByTraffic",
 *             "Netmode": {
 *               "Ref": "Netmode"
 *             },
 *             "Bandwidth": {
 *               "Ref": "Bandwidth"
 *             }
 *           }
 *         }
 *       }
 *     }
 *
 * @param request GetTemplateEstimateCostRequest
 * @return GetTemplateEstimateCostResponse
 */
func (client *Client) GetTemplateEstimateCost(request *GetTemplateEstimateCostRequest) (_result *GetTemplateEstimateCostResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateEstimateCostResponse{}
	_body, _err := client.GetTemplateEstimateCostWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to query the values of a parameter. In this example, the values of the `ZoneInfo` parameter in a template that is created in the China (Hangzhou) region are queried. The template body is `{"Parameters":{"ZoneInfo":{"Type": "String"},"InstanceType": {"Type": "String"}},"ROSTemplateFormatVersion": "2015-09-01","Resources":{"ECS":{"Properties":{"ZoneId":{"Ref": "ZoneInfo"},"InstanceType": {"Ref": "InstanceType"}},"Type": "ALIYUN::ECS::Instance"}}}`.
 * For more information about the template parameters whose values you can query by calling this operation and the sample code of the template, see [Query the constraints of parameters](~~432820~~).
 *
 * @param tmpReq GetTemplateParameterConstraintsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetTemplateParameterConstraintsResponse
 */
func (client *Client) GetTemplateParameterConstraintsWithOptions(tmpReq *GetTemplateParameterConstraintsRequest, runtime *util.RuntimeOptions) (_result *GetTemplateParameterConstraintsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetTemplateParameterConstraintsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ParametersKeyFilter)) {
		request.ParametersKeyFilterShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParametersKeyFilter, tea.String("ParametersKeyFilter"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ParametersOrder)) {
		request.ParametersOrderShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ParametersOrder, tea.String("ParametersOrder"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersKeyFilterShrink)) {
		query["ParametersKeyFilter"] = request.ParametersKeyFilterShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersOrderShrink)) {
		query["ParametersOrder"] = request.ParametersOrderShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplateParameterConstraints"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateParameterConstraintsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to query the values of a parameter. In this example, the values of the `ZoneInfo` parameter in a template that is created in the China (Hangzhou) region are queried. The template body is `{"Parameters":{"ZoneInfo":{"Type": "String"},"InstanceType": {"Type": "String"}},"ROSTemplateFormatVersion": "2015-09-01","Resources":{"ECS":{"Properties":{"ZoneId":{"Ref": "ZoneInfo"},"InstanceType": {"Ref": "InstanceType"}},"Type": "ALIYUN::ECS::Instance"}}}`.
 * For more information about the template parameters whose values you can query by calling this operation and the sample code of the template, see [Query the constraints of parameters](~~432820~~).
 *
 * @param request GetTemplateParameterConstraintsRequest
 * @return GetTemplateParameterConstraintsResponse
 */
func (client *Client) GetTemplateParameterConstraints(request *GetTemplateParameterConstraintsRequest) (_result *GetTemplateParameterConstraintsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateParameterConstraintsResponse{}
	_body, _err := client.GetTemplateParameterConstraintsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateRecommendParametersWithOptions(request *GetTemplateRecommendParametersRequest, runtime *util.RuntimeOptions) (_result *GetTemplateRecommendParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		query["TemplateBody"] = request.TemplateBody
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplateRecommendParameters"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateRecommendParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplateRecommendParameters(request *GetTemplateRecommendParametersRequest) (_result *GetTemplateRecommendParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateRecommendParametersResponse{}
	_body, _err := client.GetTemplateRecommendParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the details of the scenario whose ID is `ts-7f7a704cf71c49a6****` is queried. In the response, the source node data is displayed.
 *
 * @param request GetTemplateScratchRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetTemplateScratchResponse
 */
func (client *Client) GetTemplateScratchWithOptions(request *GetTemplateScratchRequest, runtime *util.RuntimeOptions) (_result *GetTemplateScratchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowDataOption)) {
		query["ShowDataOption"] = request.ShowDataOption
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchId)) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplateScratch"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateScratchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the details of the scenario whose ID is `ts-7f7a704cf71c49a6****` is queried. In the response, the source node data is displayed.
 *
 * @param request GetTemplateScratchRequest
 * @return GetTemplateScratchResponse
 */
func (client *Client) GetTemplateScratch(request *GetTemplateScratchRequest) (_result *GetTemplateScratchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateScratchResponse{}
	_body, _err := client.GetTemplateScratchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateSummaryWithOptions(request *GetTemplateSummaryRequest, runtime *util.RuntimeOptions) (_result *GetTemplateSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeSetId)) {
		query["ChangeSetId"] = request.ChangeSetId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupName)) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		query["TemplateBody"] = request.TemplateBody
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplateSummary"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTemplateSummaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTemplateSummary(request *GetTemplateSummaryRequest) (_result *GetTemplateSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTemplateSummaryResponse{}
	_body, _err := client.GetTemplateSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListChangeSetsWithOptions(request *ListChangeSetsRequest, runtime *util.RuntimeOptions) (_result *ListChangeSetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChangeSetId)) {
		query["ChangeSetId"] = request.ChangeSetId
	}

	if !tea.BoolValue(util.IsUnset(request.ChangeSetName)) {
		query["ChangeSetName"] = request.ChangeSetName
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionStatus)) {
		query["ExecutionStatus"] = request.ExecutionStatus
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListChangeSets"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListChangeSetsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListChangeSets(request *ListChangeSetsRequest) (_result *ListChangeSetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListChangeSetsResponse{}
	_body, _err := client.ListChangeSetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDiagnosticsWithOptions(request *ListDiagnosticsRequest, runtime *util.RuntimeOptions) (_result *ListDiagnosticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiagnosticKey)) {
		query["DiagnosticKey"] = request.DiagnosticKey
	}

	if !tea.BoolValue(util.IsUnset(request.DiagnosticProduct)) {
		query["DiagnosticProduct"] = request.DiagnosticProduct
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDiagnostics"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDiagnosticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDiagnostics(request *ListDiagnosticsRequest) (_result *ListDiagnosticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDiagnosticsResponse{}
	_body, _err := client.ListDiagnosticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceTypeRegistrationsWithOptions(request *ListResourceTypeRegistrationsRequest, runtime *util.RuntimeOptions) (_result *ListResourceTypeRegistrationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegistrationId)) {
		query["RegistrationId"] = request.RegistrationId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceTypeRegistrations"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceTypeRegistrationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceTypeRegistrations(request *ListResourceTypeRegistrationsRequest) (_result *ListResourceTypeRegistrationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceTypeRegistrationsResponse{}
	_body, _err := client.ListResourceTypeRegistrationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceTypeVersionsWithOptions(request *ListResourceTypeVersionsRequest, runtime *util.RuntimeOptions) (_result *ListResourceTypeVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceTypeVersions"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceTypeVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceTypeVersions(request *ListResourceTypeVersionsRequest) (_result *ListResourceTypeVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceTypeVersionsResponse{}
	_body, _err := client.ListResourceTypeVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * For more information about errors common to all operations, see [Common error codes](/help/en/resource-orchestration-service/latest/common-error-codes).
 *
 * @param request ListResourceTypesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListResourceTypesResponse
 */
func (client *Client) ListResourceTypesWithOptions(request *ListResourceTypesRequest, runtime *util.RuntimeOptions) (_result *ListResourceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.Provider)) {
		query["Provider"] = request.Provider
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceTypes"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * For more information about errors common to all operations, see [Common error codes](/help/en/resource-orchestration-service/latest/common-error-codes).
 *
 * @param request ListResourceTypesRequest
 * @return ListResourceTypesResponse
 */
func (client *Client) ListResourceTypes(request *ListResourceTypesRequest) (_result *ListResourceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceTypesResponse{}
	_body, _err := client.ListResourceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStackEventsWithOptions(request *ListStackEventsRequest, runtime *util.RuntimeOptions) (_result *ListStackEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogicalResourceId)) {
		query["LogicalResourceId"] = request.LogicalResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStackEvents"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStackEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStackEvents(request *ListStackEventsRequest) (_result *ListStackEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackEventsResponse{}
	_body, _err := client.ListStackEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the operation ID `6da106ca-1784-4a6f-a7e1-e723863d∗∗∗∗` is set to query the results of an operation on a stack group named `MyStackGroup`. The stack group is granted self-managed permissions and created in the China (Hangzhou) region.
 *
 * @param request ListStackGroupOperationResultsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListStackGroupOperationResultsResponse
 */
func (client *Client) ListStackGroupOperationResultsWithOptions(request *ListStackGroupOperationResultsRequest, runtime *util.RuntimeOptions) (_result *ListStackGroupOperationResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationId)) {
		query["OperationId"] = request.OperationId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStackGroupOperationResults"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStackGroupOperationResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the operation ID `6da106ca-1784-4a6f-a7e1-e723863d∗∗∗∗` is set to query the results of an operation on a stack group named `MyStackGroup`. The stack group is granted self-managed permissions and created in the China (Hangzhou) region.
 *
 * @param request ListStackGroupOperationResultsRequest
 * @return ListStackGroupOperationResultsResponse
 */
func (client *Client) ListStackGroupOperationResults(request *ListStackGroupOperationResultsRequest) (_result *ListStackGroupOperationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackGroupOperationResultsResponse{}
	_body, _err := client.ListStackGroupOperationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStackGroupOperationsWithOptions(request *ListStackGroupOperationsRequest, runtime *util.RuntimeOptions) (_result *ListStackGroupOperationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupName)) {
		query["StackGroupName"] = request.StackGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStackGroupOperations"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStackGroupOperationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStackGroupOperations(request *ListStackGroupOperationsRequest) (_result *ListStackGroupOperationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackGroupOperationsResponse{}
	_body, _err := client.ListStackGroupOperationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the list of stack groups that are in the ACTIVE state and deployed in the China (Hangzhou) region is queried.
 *
 * @param request ListStackGroupsRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListStackGroupsResponse
 */
func (client *Client) ListStackGroupsWithOptions(request *ListStackGroupsRequest, runtime *util.RuntimeOptions) (_result *ListStackGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStackGroups"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStackGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the list of stack groups that are in the ACTIVE state and deployed in the China (Hangzhou) region is queried.
 *
 * @param request ListStackGroupsRequest
 * @return ListStackGroupsResponse
 */
func (client *Client) ListStackGroups(request *ListStackGroupsRequest) (_result *ListStackGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackGroupsResponse{}
	_body, _err := client.ListStackGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the list of stack instances that are associated with a stack group named `MyStackGroup` is queried. The stack group is granted self-managed permissions and deployed in the China (Hangzhou) region.
 *
 * @param request ListStackInstancesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListStackInstancesResponse
 */
func (client *Client) ListStackInstancesWithOptions(request *ListStackInstancesRequest, runtime *util.RuntimeOptions) (_result *ListStackInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupName)) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.StackInstanceAccountId)) {
		query["StackInstanceAccountId"] = request.StackInstanceAccountId
	}

	if !tea.BoolValue(util.IsUnset(request.StackInstanceRegionId)) {
		query["StackInstanceRegionId"] = request.StackInstanceRegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStackInstances"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStackInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the list of stack instances that are associated with a stack group named `MyStackGroup` is queried. The stack group is granted self-managed permissions and deployed in the China (Hangzhou) region.
 *
 * @param request ListStackInstancesRequest
 * @return ListStackInstancesResponse
 */
func (client *Client) ListStackInstances(request *ListStackInstancesRequest) (_result *ListStackInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackInstancesResponse{}
	_body, _err := client.ListStackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The ListStackOperationRisks operation is suitable for the following scenarios:
 * *   You want to detect high risks that may arise in resources when you delete a stack that contains the resources, and query the cause of each risk in a resource.
 * *   When you create a stack, the creation may fail. In this case, you can call this operation to check which types of permissions that are required to create stacks are missing.
 *
 * @param request ListStackOperationRisksRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListStackOperationRisksResponse
 */
func (client *Client) ListStackOperationRisksWithOptions(request *ListStackOperationRisksRequest, runtime *util.RuntimeOptions) (_result *ListStackOperationRisksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		query["OperationType"] = request.OperationType
	}

	if !tea.BoolValue(util.IsUnset(request.RamRoleName)) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RetainAllResources)) {
		query["RetainAllResources"] = request.RetainAllResources
	}

	if !tea.BoolValue(util.IsUnset(request.RetainResources)) {
		query["RetainResources"] = request.RetainResources
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		query["TemplateBody"] = request.TemplateBody
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStackOperationRisks"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStackOperationRisksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The ListStackOperationRisks operation is suitable for the following scenarios:
 * *   You want to detect high risks that may arise in resources when you delete a stack that contains the resources, and query the cause of each risk in a resource.
 * *   When you create a stack, the creation may fail. In this case, you can call this operation to check which types of permissions that are required to create stacks are missing.
 *
 * @param request ListStackOperationRisksRequest
 * @return ListStackOperationRisksResponse
 */
func (client *Client) ListStackOperationRisks(request *ListStackOperationRisksRequest) (_result *ListStackOperationRisksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackOperationRisksResponse{}
	_body, _err := client.ListStackOperationRisksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStackResourceDriftsWithOptions(request *ListStackResourceDriftsRequest, runtime *util.RuntimeOptions) (_result *ListStackResourceDriftsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceDriftStatus)) {
		query["ResourceDriftStatus"] = request.ResourceDriftStatus
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStackResourceDrifts"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStackResourceDriftsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStackResourceDrifts(request *ListStackResourceDriftsRequest) (_result *ListStackResourceDriftsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackResourceDriftsResponse{}
	_body, _err := client.ListStackResourceDriftsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * For more information about common request parameters, see [Common parameters](~~131957~~).
 *
 * @param request ListStackResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListStackResourcesResponse
 */
func (client *Client) ListStackResourcesWithOptions(request *ListStackResourcesRequest, runtime *util.RuntimeOptions) (_result *ListStackResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStackResources"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStackResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * For more information about common request parameters, see [Common parameters](~~131957~~).
 *
 * @param request ListStackResourcesRequest
 * @return ListStackResourcesResponse
 */
func (client *Client) ListStackResources(request *ListStackResourcesRequest) (_result *ListStackResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStackResourcesResponse{}
	_body, _err := client.ListStackResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * This topic provides an example on how to query a list of stacks. In this example, the stacks that are deployed in the China (Hangzhou) region are queried.
 *
 * @param request ListStacksRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListStacksResponse
 */
func (client *Client) ListStacksWithOptions(request *ListStacksRequest, runtime *util.RuntimeOptions) (_result *ListStacksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentStackId)) {
		query["ParentStackId"] = request.ParentStackId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowNestedStack)) {
		query["ShowNestedStack"] = request.ShowNestedStack
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.StackIds)) {
		query["StackIds"] = request.StackIds
	}

	if !tea.BoolValue(util.IsUnset(request.StackName)) {
		query["StackName"] = request.StackName
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStacks"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStacksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * This topic provides an example on how to query a list of stacks. In this example, the stacks that are deployed in the China (Hangzhou) region are queried.
 *
 * @param request ListStacksRequest
 * @return ListStacksResponse
 */
func (client *Client) ListStacks(request *ListStacksRequest) (_result *ListStacksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListStacksResponse{}
	_body, _err := client.ListStacksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the tag keys that are added to a stack in the China (Hangzhou) region are queried.
 *
 * @param request ListTagKeysRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTagKeysResponse
 */
func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the tag keys that are added to a stack in the China (Hangzhou) region are queried.
 *
 * @param request ListTagKeysRequest
 * @return ListTagKeysResponse
 */
func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ###
 * *   To specify the query object, specify ResourceId or Tag in the request. Tag consists of Key and Value.
 * *   If you specify Tag and ResourceId, ROS resources that match both the parameters are returned.
 * This topic provides an example on how to query the tags that are added to a stack. In this example, the stack ID is `6bc589b5-9c02-4944-8fc3-f3624234****`. The stack is deployed in the China (Hangzhou) region.
 *
 * @param request ListTagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTagResourcesResponse
 */
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ###
 * *   To specify the query object, specify ResourceId or Tag in the request. Tag consists of Key and Value.
 * *   If you specify Tag and ResourceId, ROS resources that match both the parameters are returned.
 * This topic provides an example on how to query the tags that are added to a stack. In this example, the stack ID is `6bc589b5-9c02-4944-8fc3-f3624234****`. The stack is deployed in the China (Hangzhou) region.
 *
 * @param request ListTagResourcesRequest
 * @return ListTagResourcesResponse
 */
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the tag values of `TagKey1` that is added to a stack in the China (Hangzhou) region are queried.
 *
 * @param request ListTagValuesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTagValuesResponse
 */
func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the tag values of `TagKey1` that is added to a stack in the China (Hangzhou) region are queried.
 *
 * @param request ListTagValuesRequest
 * @return ListTagValuesResponse
 */
func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the scenarios that are created in the China (Hangzhou) region are queried. In the response, a scenario of the Resource Management and a scenario of the Resource Replication type are returned.
 *
 * @param request ListTemplateScratchesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListTemplateScratchesResponse
 */
func (client *Client) ListTemplateScratchesWithOptions(request *ListTemplateScratchesRequest, runtime *util.RuntimeOptions) (_result *ListTemplateScratchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchId)) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchType)) {
		query["TemplateScratchType"] = request.TemplateScratchType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTemplateScratches"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTemplateScratchesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the scenarios that are created in the China (Hangzhou) region are queried. In the response, a scenario of the Resource Management and a scenario of the Resource Replication type are returned.
 *
 * @param request ListTemplateScratchesRequest
 * @return ListTemplateScratchesResponse
 */
func (client *Client) ListTemplateScratches(request *ListTemplateScratchesRequest) (_result *ListTemplateScratchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplateScratchesResponse{}
	_body, _err := client.ListTemplateScratchesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplateVersionsWithOptions(request *ListTemplateVersionsRequest, runtime *util.RuntimeOptions) (_result *ListTemplateVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTemplateVersions"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTemplateVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplateVersions(request *ListTemplateVersionsRequest) (_result *ListTemplateVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplateVersionsResponse{}
	_body, _err := client.ListTemplateVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IncludeTags)) {
		query["IncludeTags"] = request.IncludeTags
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareType)) {
		query["ShareType"] = request.ShareType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTemplates"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, a stack deployed in the `China (Hangzhou)` region is moved to a specific resource group. The ID of the stack is `4e8611cb-251e-42b7-b9cb-3496362c****` and the ID of the resource group is `rg-acfm3peow3k****`.
 *
 * @param request MoveResourceGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return MoveResourceGroupResponse
 */
func (client *Client) MoveResourceGroupWithOptions(request *MoveResourceGroupRequest, runtime *util.RuntimeOptions) (_result *MoveResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		query["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveResourceGroup"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, a stack deployed in the `China (Hangzhou)` region is moved to a specific resource group. The ID of the stack is `4e8611cb-251e-42b7-b9cb-3496362c****` and the ID of the resource group is `rg-acfm3peow3k****`.
 *
 * @param request MoveResourceGroupRequest
 * @return MoveResourceGroupResponse
 */
func (client *Client) MoveResourceGroup(request *MoveResourceGroupRequest) (_result *MoveResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveResourceGroupResponse{}
	_body, _err := client.MoveResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to create a stack named `MyStack` in the China (Hangzhou) region by using a template and preview the information about the stack. In this example, the `template body` is `{"ROSTemplateFormatVersion":"2015-09-01"}`.
 *
 * @param request PreviewStackRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return PreviewStackResponse
 */
func (client *Client) PreviewStackWithOptions(request *PreviewStackRequest, runtime *util.RuntimeOptions) (_result *PreviewStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DisableRollback)) {
		query["DisableRollback"] = request.DisableRollback
	}

	if !tea.BoolValue(util.IsUnset(request.EnablePreConfig)) {
		query["EnablePreConfig"] = request.EnablePreConfig
	}

	if !tea.BoolValue(util.IsUnset(request.Parallelism)) {
		query["Parallelism"] = request.Parallelism
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.StackName)) {
		query["StackName"] = request.StackName
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyBody)) {
		query["StackPolicyBody"] = request.StackPolicyBody
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyURL)) {
		query["StackPolicyURL"] = request.StackPolicyURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchId)) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchRegionId)) {
		query["TemplateScratchRegionId"] = request.TemplateScratchRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TimeoutInMinutes)) {
		query["TimeoutInMinutes"] = request.TimeoutInMinutes
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PreviewStack"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &PreviewStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to create a stack named `MyStack` in the China (Hangzhou) region by using a template and preview the information about the stack. In this example, the `template body` is `{"ROSTemplateFormatVersion":"2015-09-01"}`.
 *
 * @param request PreviewStackRequest
 * @return PreviewStackResponse
 */
func (client *Client) PreviewStack(request *PreviewStackRequest) (_result *PreviewStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PreviewStackResponse{}
	_body, _err := client.PreviewStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   Versions increase from v1.
 * *   If you create a new resource type, v1 is used as the default version of the resource type. You can call the SetResourceType operation to change the default version of a resource type.
 *
 * @param request RegisterResourceTypeRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return RegisterResourceTypeResponse
 */
func (client *Client) RegisterResourceTypeWithOptions(request *RegisterResourceTypeRequest, runtime *util.RuntimeOptions) (_result *RegisterResourceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EntityType)) {
		query["EntityType"] = request.EntityType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterResourceType"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   Versions increase from v1.
 * *   If you create a new resource type, v1 is used as the default version of the resource type. You can call the SetResourceType operation to change the default version of a resource type.
 *
 * @param request RegisterResourceTypeRequest
 * @return RegisterResourceTypeResponse
 */
func (client *Client) RegisterResourceType(request *RegisterResourceTypeRequest) (_result *RegisterResourceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterResourceTypeResponse{}
	_body, _err := client.RegisterResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDeletionProtectionWithOptions(request *SetDeletionProtectionRequest, runtime *util.RuntimeOptions) (_result *SetDeletionProtectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeletionProtection)) {
		query["DeletionProtection"] = request.DeletionProtection
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetDeletionProtection"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetDeletionProtectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDeletionProtection(request *SetDeletionProtectionRequest) (_result *SetDeletionProtectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDeletionProtectionResponse{}
	_body, _err := client.SetDeletionProtectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetResourceTypeWithOptions(request *SetResourceTypeRequest, runtime *util.RuntimeOptions) (_result *SetResourceTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DefaultVersionId)) {
		query["DefaultVersionId"] = request.DefaultVersionId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetResourceType"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetResourceTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetResourceType(request *SetResourceTypeRequest) (_result *SetResourceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetResourceTypeResponse{}
	_body, _err := client.SetResourceTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, a stack policy is configured for a stack deployed in the `China (Hangzhou)` region whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****`. The URL to the stack policy body is `oss://ros/stack-policy/demo`.
 *
 * @param request SetStackPolicyRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetStackPolicyResponse
 */
func (client *Client) SetStackPolicyWithOptions(request *SetStackPolicyRequest, runtime *util.RuntimeOptions) (_result *SetStackPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyBody)) {
		query["StackPolicyBody"] = request.StackPolicyBody
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyURL)) {
		query["StackPolicyURL"] = request.StackPolicyURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetStackPolicy"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetStackPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, a stack policy is configured for a stack deployed in the `China (Hangzhou)` region whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****`. The URL to the stack policy body is `oss://ros/stack-policy/demo`.
 *
 * @param request SetStackPolicyRequest
 * @return SetStackPolicyResponse
 */
func (client *Client) SetStackPolicy(request *SetStackPolicyRequest) (_result *SetStackPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetStackPolicyResponse{}
	_body, _err := client.SetStackPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, the template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****` is shared with an Alibaba Cloud account. The ID of the Alibaba Cloud account is `151266687691****`.
 * > The recipient Alibaba Cloud account (ID: `151266687691****`) can authorize RAM users to use the shared template.
 *
 * @param request SetTemplatePermissionRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return SetTemplatePermissionResponse
 */
func (client *Client) SetTemplatePermissionWithOptions(request *SetTemplatePermissionRequest, runtime *util.RuntimeOptions) (_result *SetTemplatePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountIds)) {
		query["AccountIds"] = request.AccountIds
	}

	if !tea.BoolValue(util.IsUnset(request.ShareOption)) {
		query["ShareOption"] = request.ShareOption
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.VersionOption)) {
		query["VersionOption"] = request.VersionOption
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetTemplatePermission"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetTemplatePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, the template whose ID is `5ecd1e10-b0e9-4389-a565-e4c15efc****` is shared with an Alibaba Cloud account. The ID of the Alibaba Cloud account is `151266687691****`.
 * > The recipient Alibaba Cloud account (ID: `151266687691****`) can authorize RAM users to use the shared template.
 *
 * @param request SetTemplatePermissionRequest
 * @return SetTemplatePermissionResponse
 */
func (client *Client) SetTemplatePermission(request *SetTemplatePermissionRequest) (_result *SetTemplatePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetTemplatePermissionResponse{}
	_body, _err := client.SetTemplatePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SignalResourceWithOptions(request *SignalResourceRequest, runtime *util.RuntimeOptions) (_result *SignalResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.LogicalResourceId)) {
		query["LogicalResourceId"] = request.LogicalResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UniqueId)) {
		query["UniqueId"] = request.UniqueId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SignalResource"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SignalResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SignalResource(request *SignalResourceRequest) (_result *SignalResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SignalResourceResponse{}
	_body, _err := client.SignalResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to stop a stack group operation whose ID is `6da106ca-1784-4a6f-a7e1-e723863****` in the China (Hangzhou) region.
 *
 * @param request StopStackGroupOperationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return StopStackGroupOperationResponse
 */
func (client *Client) StopStackGroupOperationWithOptions(request *StopStackGroupOperationRequest, runtime *util.RuntimeOptions) (_result *StopStackGroupOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationId)) {
		query["OperationId"] = request.OperationId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopStackGroupOperation"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopStackGroupOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to stop a stack group operation whose ID is `6da106ca-1784-4a6f-a7e1-e723863****` in the China (Hangzhou) region.
 *
 * @param request StopStackGroupOperationRequest
 * @return StopStackGroupOperationResponse
 */
func (client *Client) StopStackGroupOperation(request *StopStackGroupOperationRequest) (_result *StopStackGroupOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopStackGroupOperationResponse{}
	_body, _err := client.StopStackGroupOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to create a tag and add the tag to a stack. In this example, the stack ID is `7fee80e1-8c48-4c2f-8300-0f6dc40b****`, the tag key is `FinanceDept`, and the tag value is `FinanceJoshua`.
 *
 * @param request TagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return TagResourcesResponse
 */
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to create a tag and add the tag to a stack. In this example, the stack ID is `7fee80e1-8c48-4c2f-8300-0f6dc40b****`, the tag key is `FinanceDept`, and the tag value is `FinanceJoshua`.
 *
 * @param request TagResourcesRequest
 * @return TagResourcesResponse
 */
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * This topic provides an example on how to remove all tags from a stack that is deployed in the China (Hangzhou) region. In this example, the stack ID is `46ec7b78-9d5e-4b21-aefd-448c90aa****`.
 *
 * @param request UntagResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UntagResourcesResponse
 */
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * This topic provides an example on how to remove all tags from a stack that is deployed in the China (Hangzhou) region. In this example, the stack ID is `46ec7b78-9d5e-4b21-aefd-448c90aa****`.
 *
 * @param request UntagResourcesRequest
 * @return UntagResourcesResponse
 */
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The values of parameters in the Parameters section vary based on the value that you specify for the UsePreviousParameters parameter in the request. If you do not add the parameters that are defined in the template to the Parameters section, take note of the following items:
 * *   UsePreviousParameters is set to false: If the template parameters have default values, the default values are used. Otherwise, you must specify values for the template parameters in the Parameters section.
 * *   UsePreviousParameters is set to true: If you specify values for the template parameters when you create a stack, the values are used. If you leave the template parameters empty when you create a stack but the template parameters have default values, the default values are used.
 * This topic describes how to update a stack. In this example, the template body of a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` in the China (Beijing) region is updated to `{"ROSTemplateFormatVersion": "2015-09-01"}`.
 *
 * @param request UpdateStackRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateStackResponse
 */
func (client *Client) UpdateStackWithOptions(request *UpdateStackRequest, runtime *util.RuntimeOptions) (_result *UpdateStackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DisableRollback)) {
		query["DisableRollback"] = request.DisableRollback
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.DryRunOptions)) {
		query["DryRunOptions"] = request.DryRunOptions
	}

	if !tea.BoolValue(util.IsUnset(request.Parallelism)) {
		query["Parallelism"] = request.Parallelism
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.RamRoleName)) {
		query["RamRoleName"] = request.RamRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ReplacementOption)) {
		query["ReplacementOption"] = request.ReplacementOption
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyBody)) {
		query["StackPolicyBody"] = request.StackPolicyBody
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyDuringUpdateBody)) {
		query["StackPolicyDuringUpdateBody"] = request.StackPolicyDuringUpdateBody
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyDuringUpdateURL)) {
		query["StackPolicyDuringUpdateURL"] = request.StackPolicyDuringUpdateURL
	}

	if !tea.BoolValue(util.IsUnset(request.StackPolicyURL)) {
		query["StackPolicyURL"] = request.StackPolicyURL
	}

	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	if !tea.BoolValue(util.IsUnset(request.TimeoutInMinutes)) {
		query["TimeoutInMinutes"] = request.TimeoutInMinutes
	}

	if !tea.BoolValue(util.IsUnset(request.UsePreviousParameters)) {
		query["UsePreviousParameters"] = request.UsePreviousParameters
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStack"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The values of parameters in the Parameters section vary based on the value that you specify for the UsePreviousParameters parameter in the request. If you do not add the parameters that are defined in the template to the Parameters section, take note of the following items:
 * *   UsePreviousParameters is set to false: If the template parameters have default values, the default values are used. Otherwise, you must specify values for the template parameters in the Parameters section.
 * *   UsePreviousParameters is set to true: If you specify values for the template parameters when you create a stack, the values are used. If you leave the template parameters empty when you create a stack but the template parameters have default values, the default values are used.
 * This topic describes how to update a stack. In this example, the template body of a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****` in the China (Beijing) region is updated to `{"ROSTemplateFormatVersion": "2015-09-01"}`.
 *
 * @param request UpdateStackRequest
 * @return UpdateStackResponse
 */
func (client *Client) UpdateStack(request *UpdateStackRequest) (_result *UpdateStackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStackResponse{}
	_body, _err := client.UpdateStackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The name of the stack group. The name must be unique within a region.
 * The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\\_). The name must start with a digit or a letter.
 *
 * @param tmpReq UpdateStackGroupRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateStackGroupResponse
 */
func (client *Client) UpdateStackGroupWithOptions(tmpReq *UpdateStackGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateStackGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateStackGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AccountIds)) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, tea.String("AccountIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.AutoDeployment)) {
		request.AutoDeploymentShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AutoDeployment, tea.String("AutoDeployment"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DeploymentTargets)) {
		request.DeploymentTargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeploymentTargets, tea.String("DeploymentTargets"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OperationPreferences)) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, tea.String("OperationPreferences"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("RegionIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountIdsShrink)) {
		query["AccountIds"] = request.AccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.AdministrationRoleName)) {
		query["AdministrationRoleName"] = request.AdministrationRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.AutoDeploymentShrink)) {
		query["AutoDeployment"] = request.AutoDeploymentShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Capabilities)) {
		query["Capabilities"] = request.Capabilities
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeploymentTargetsShrink)) {
		query["DeploymentTargets"] = request.DeploymentTargetsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionRoleName)) {
		query["ExecutionRoleName"] = request.ExecutionRoleName
	}

	if !tea.BoolValue(util.IsUnset(request.OperationDescription)) {
		query["OperationDescription"] = request.OperationDescription
	}

	if !tea.BoolValue(util.IsUnset(request.OperationPreferencesShrink)) {
		query["OperationPreferences"] = request.OperationPreferencesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		query["Parameters"] = request.Parameters
	}

	if !tea.BoolValue(util.IsUnset(request.PermissionModel)) {
		query["PermissionModel"] = request.PermissionModel
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionIdsShrink)) {
		query["RegionIds"] = request.RegionIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupName)) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateVersion)) {
		query["TemplateVersion"] = request.TemplateVersion
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStackGroup"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStackGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The name of the stack group. The name must be unique within a region.
 * The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (\\_). The name must start with a digit or a letter.
 *
 * @param request UpdateStackGroupRequest
 * @return UpdateStackGroupResponse
 */
func (client *Client) UpdateStackGroup(request *UpdateStackGroupRequest) (_result *UpdateStackGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStackGroupResponse{}
	_body, _err := client.UpdateStackGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this topic, the stack group named `MyStackGroup` that is created in the China (Hangzhou) region is used. The stack group is granted the self-managed permissions. In this example, stacks of the stack group are updated by using the Alibaba Cloud accounts whose IDs are `151266687691****` and `141261387191****` in the China (Hangzhou) region and China (Beijing) region.
 *
 * @param tmpReq UpdateStackInstancesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateStackInstancesResponse
 */
func (client *Client) UpdateStackInstancesWithOptions(tmpReq *UpdateStackInstancesRequest, runtime *util.RuntimeOptions) (_result *UpdateStackInstancesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateStackInstancesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AccountIds)) {
		request.AccountIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AccountIds, tea.String("AccountIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DeploymentTargets)) {
		request.DeploymentTargetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DeploymentTargets, tea.String("DeploymentTargets"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.OperationPreferences)) {
		request.OperationPreferencesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OperationPreferences, tea.String("OperationPreferences"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.RegionIds)) {
		request.RegionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RegionIds, tea.String("RegionIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AccountIdsShrink)) {
		query["AccountIds"] = request.AccountIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DeploymentTargetsShrink)) {
		query["DeploymentTargets"] = request.DeploymentTargetsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OperationDescription)) {
		query["OperationDescription"] = request.OperationDescription
	}

	if !tea.BoolValue(util.IsUnset(request.OperationPreferencesShrink)) {
		query["OperationPreferences"] = request.OperationPreferencesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ParameterOverrides)) {
		query["ParameterOverrides"] = request.ParameterOverrides
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionIdsShrink)) {
		query["RegionIds"] = request.RegionIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StackGroupName)) {
		query["StackGroupName"] = request.StackGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.TimeoutInMinutes)) {
		query["TimeoutInMinutes"] = request.TimeoutInMinutes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStackInstances"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStackInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this topic, the stack group named `MyStackGroup` that is created in the China (Hangzhou) region is used. The stack group is granted the self-managed permissions. In this example, stacks of the stack group are updated by using the Alibaba Cloud accounts whose IDs are `151266687691****` and `141261387191****` in the China (Hangzhou) region and China (Beijing) region.
 *
 * @param request UpdateStackInstancesRequest
 * @return UpdateStackInstancesResponse
 */
func (client *Client) UpdateStackInstances(request *UpdateStackInstancesRequest) (_result *UpdateStackInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStackInstancesResponse{}
	_body, _err := client.UpdateStackInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Limits: You can eliminate only drift on stacks that have drifted. You must call the [DetectStackDrift](~~155094~~) operation to perform drift detection on a stack, call the [GetStackDriftDetectionStatus](~~155097~~) operation to query the drift status of the stack to make sure that the stack has drifted, and then call the UpdateStackTemplateByResources operation to eliminate drift.
 * In this topic, drift is eliminated for a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****`. The stack is deployed in the China (Hangzhou) region.
 *
 * @param request UpdateStackTemplateByResourcesRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateStackTemplateByResourcesResponse
 */
func (client *Client) UpdateStackTemplateByResourcesWithOptions(request *UpdateStackTemplateByResourcesRequest, runtime *util.RuntimeOptions) (_result *UpdateStackTemplateByResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["DryRun"] = request.DryRun
	}

	if !tea.BoolValue(util.IsUnset(request.LogicalResourceId)) {
		query["LogicalResourceId"] = request.LogicalResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.StackId)) {
		query["StackId"] = request.StackId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateFormat)) {
		query["TemplateFormat"] = request.TemplateFormat
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStackTemplateByResources"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStackTemplateByResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Limits: You can eliminate only drift on stacks that have drifted. You must call the [DetectStackDrift](~~155094~~) operation to perform drift detection on a stack, call the [GetStackDriftDetectionStatus](~~155097~~) operation to query the drift status of the stack to make sure that the stack has drifted, and then call the UpdateStackTemplateByResources operation to eliminate drift.
 * In this topic, drift is eliminated for a stack whose ID is `4a6c9851-3b0f-4f5f-b4ca-a14bf691****`. The stack is deployed in the China (Hangzhou) region.
 *
 * @param request UpdateStackTemplateByResourcesRequest
 * @return UpdateStackTemplateByResourcesResponse
 */
func (client *Client) UpdateStackTemplateByResources(request *UpdateStackTemplateByResourcesRequest) (_result *UpdateStackTemplateByResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStackTemplateByResourcesResponse{}
	_body, _err := client.UpdateStackTemplateByResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * When you update a template, take note of the following items:
 * *   If you specify TemplateBody or TemplateURL, the existing version number is automatically incremented by 1 after the template is updated. For example, the version is changed from v1 to v2.
 * *   If you do not specify TemplateBody or TemplateURL, the version number remains unchanged.
 * *   A template can have up to 100 versions. If the number of templates reaches the upper limit, the template fails to be updated. You must create another template.
 *
 * @param request UpdateTemplateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateTemplateResponse
 */
func (client *Client) UpdateTemplateWithOptions(request *UpdateTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		query["TemplateBody"] = request.TemplateBody
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateName)) {
		query["TemplateName"] = request.TemplateName
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTemplate"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * When you update a template, take note of the following items:
 * *   If you specify TemplateBody or TemplateURL, the existing version number is automatically incremented by 1 after the template is updated. For example, the version is changed from v1 to v2.
 * *   If you do not specify TemplateBody or TemplateURL, the version number remains unchanged.
 * *   A template can have up to 100 versions. If the number of templates reaches the upper limit, the template fails to be updated. You must create another template.
 *
 * @param request UpdateTemplateRequest
 * @return UpdateTemplateResponse
 */
func (client *Client) UpdateTemplate(request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, a scenario whose ID is `ts-7f7a704cf71c49a6****` is used. The scenario is created in the China (Hangzhou) region. In the scenario, the ID of a virtual private cloud (VPC) is updated to `vpc-bp1m6fww66xbntjyc****`.
 *
 * @param tmpReq UpdateTemplateScratchRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateTemplateScratchResponse
 */
func (client *Client) UpdateTemplateScratchWithOptions(tmpReq *UpdateTemplateScratchRequest, runtime *util.RuntimeOptions) (_result *UpdateTemplateScratchResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateTemplateScratchShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.PreferenceParameters)) {
		request.PreferenceParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PreferenceParameters, tea.String("PreferenceParameters"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceResourceGroup)) {
		request.SourceResourceGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceResourceGroup, tea.String("SourceResourceGroup"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceResources)) {
		request.SourceResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceResources, tea.String("SourceResources"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceTag)) {
		request.SourceTagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceTag, tea.String("SourceTag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ExecutionMode)) {
		query["ExecutionMode"] = request.ExecutionMode
	}

	if !tea.BoolValue(util.IsUnset(request.LogicalIdStrategy)) {
		query["LogicalIdStrategy"] = request.LogicalIdStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.PreferenceParametersShrink)) {
		query["PreferenceParameters"] = request.PreferenceParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceResourceGroupShrink)) {
		query["SourceResourceGroup"] = request.SourceResourceGroupShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceResourcesShrink)) {
		query["SourceResources"] = request.SourceResourcesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SourceTagShrink)) {
		query["SourceTag"] = request.SourceTagShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateScratchId)) {
		query["TemplateScratchId"] = request.TemplateScratchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTemplateScratch"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTemplateScratchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, a scenario whose ID is `ts-7f7a704cf71c49a6****` is used. The scenario is created in the China (Hangzhou) region. In the scenario, the ID of a virtual private cloud (VPC) is updated to `vpc-bp1m6fww66xbntjyc****`.
 *
 * @param request UpdateTemplateScratchRequest
 * @return UpdateTemplateScratchResponse
 */
func (client *Client) UpdateTemplateScratch(request *UpdateTemplateScratchRequest) (_result *UpdateTemplateScratchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateTemplateScratchResponse{}
	_body, _err := client.UpdateTemplateScratchWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * In this example, a template that you want to use to create a stack is validated. `TemplateURL` is set to `oss://ros/template/demo`.
 *
 * @param request ValidateTemplateRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return ValidateTemplateResponse
 */
func (client *Client) ValidateTemplateWithOptions(request *ValidateTemplateRequest, runtime *util.RuntimeOptions) (_result *ValidateTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateURL)) {
		query["TemplateURL"] = request.TemplateURL
	}

	if !tea.BoolValue(util.IsUnset(request.UpdateInfoOptions)) {
		query["UpdateInfoOptions"] = request.UpdateInfoOptions
	}

	if !tea.BoolValue(util.IsUnset(request.ValidationOption)) {
		query["ValidationOption"] = request.ValidationOption
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TemplateBody)) {
		body["TemplateBody"] = request.TemplateBody
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateTemplate"),
		Version:     tea.String("2019-09-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * In this example, a template that you want to use to create a stack is validated. `TemplateURL` is set to `oss://ros/template/demo`.
 *
 * @param request ValidateTemplateRequest
 * @return ValidateTemplateResponse
 */
func (client *Client) ValidateTemplate(request *ValidateTemplateRequest) (_result *ValidateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateTemplateResponse{}
	_body, _err := client.ValidateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
