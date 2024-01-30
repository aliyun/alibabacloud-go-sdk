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

type AppFailBackRequest struct {
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s AppFailBackRequest) String() string {
	return tea.Prettify(s)
}

func (s AppFailBackRequest) GoString() string {
	return s.String()
}

func (s *AppFailBackRequest) SetApplicationId(v string) *AppFailBackRequest {
	s.ApplicationId = &v
	return s
}

type AppFailBackResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int32  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppFailBackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppFailBackResponseBody) GoString() string {
	return s.String()
}

func (s *AppFailBackResponseBody) SetCode(v string) *AppFailBackResponseBody {
	s.Code = &v
	return s
}

func (s *AppFailBackResponseBody) SetData(v int32) *AppFailBackResponseBody {
	s.Data = &v
	return s
}

func (s *AppFailBackResponseBody) SetMessage(v string) *AppFailBackResponseBody {
	s.Message = &v
	return s
}

func (s *AppFailBackResponseBody) SetRequestId(v string) *AppFailBackResponseBody {
	s.RequestId = &v
	return s
}

type AppFailBackResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppFailBackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppFailBackResponse) String() string {
	return tea.Prettify(s)
}

func (s AppFailBackResponse) GoString() string {
	return s.String()
}

func (s *AppFailBackResponse) SetHeaders(v map[string]*string) *AppFailBackResponse {
	s.Headers = v
	return s
}

func (s *AppFailBackResponse) SetStatusCode(v int32) *AppFailBackResponse {
	s.StatusCode = &v
	return s
}

func (s *AppFailBackResponse) SetBody(v *AppFailBackResponseBody) *AppFailBackResponse {
	s.Body = v
	return s
}

type AppFailOverRequest struct {
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	FailZone      *string `json:"FailZone,omitempty" xml:"FailZone,omitempty"`
}

func (s AppFailOverRequest) String() string {
	return tea.Prettify(s)
}

func (s AppFailOverRequest) GoString() string {
	return s.String()
}

func (s *AppFailOverRequest) SetApplicationId(v string) *AppFailOverRequest {
	s.ApplicationId = &v
	return s
}

func (s *AppFailOverRequest) SetFailZone(v string) *AppFailOverRequest {
	s.FailZone = &v
	return s
}

type AppFailOverResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int32  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AppFailOverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AppFailOverResponseBody) GoString() string {
	return s.String()
}

func (s *AppFailOverResponseBody) SetCode(v string) *AppFailOverResponseBody {
	s.Code = &v
	return s
}

func (s *AppFailOverResponseBody) SetData(v int32) *AppFailOverResponseBody {
	s.Data = &v
	return s
}

func (s *AppFailOverResponseBody) SetMessage(v string) *AppFailOverResponseBody {
	s.Message = &v
	return s
}

func (s *AppFailOverResponseBody) SetRequestId(v string) *AppFailOverResponseBody {
	s.RequestId = &v
	return s
}

type AppFailOverResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppFailOverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppFailOverResponse) String() string {
	return tea.Prettify(s)
}

func (s AppFailOverResponse) GoString() string {
	return s.String()
}

func (s *AppFailOverResponse) SetHeaders(v map[string]*string) *AppFailOverResponse {
	s.Headers = v
	return s
}

func (s *AppFailOverResponse) SetStatusCode(v int32) *AppFailOverResponse {
	s.StatusCode = &v
	return s
}

func (s *AppFailOverResponse) SetBody(v *AppFailOverResponseBody) *AppFailOverResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// The ID of the new resource group.
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The ID of the resource.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource for which you want to change the resource group. Valid values: APPLICATION and TEMPLATE.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// The HTTP status code. A value of 200 indicates that the request is successful. Other values indicate that the request failed.
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// No business data is returned for this parameter.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetCode(v int64) *ChangeResourceGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetData(v string) *ChangeResourceGroupResponseBody {
	s.Data = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetMessage(v string) *ChangeResourceGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CreateApplicationRequest struct {
	// The ID of the region.
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The parameters that are used to configure the application you want to create. For example, enableMonitor specifies whether to automatically create a CloudMonitor task for the application, and enableReport specifies whether to generate reports.
	Configuration map[string]*string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The instances in which you want to create the application. You can create applications in an existing virtual private cloud (VPC).
	Instances []*CreateApplicationRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The name of the application.
	//
	// *   The application name must be unique. You can call the [ListApplication](https://www.alibabacloud.com/help/en/bp-studio/latest/api-bpstudio-2021-09-31-listapplication) operation to query the existing applications.
	// *   The application name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http:// or https://`. The name can contain letters, digits, underscores (\_), and hyphens (-).
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the application you want to create belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the template.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The parameter values that are contained in the template. If the template contains no parameter values, the default values are used.
	Variables map[string]*string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) SetAreaId(v string) *CreateApplicationRequest {
	s.AreaId = &v
	return s
}

func (s *CreateApplicationRequest) SetClientToken(v string) *CreateApplicationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApplicationRequest) SetConfiguration(v map[string]*string) *CreateApplicationRequest {
	s.Configuration = v
	return s
}

func (s *CreateApplicationRequest) SetInstances(v []*CreateApplicationRequestInstances) *CreateApplicationRequest {
	s.Instances = v
	return s
}

func (s *CreateApplicationRequest) SetName(v string) *CreateApplicationRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationRequest) SetResourceGroupId(v string) *CreateApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationRequest) SetTemplateId(v string) *CreateApplicationRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateApplicationRequest) SetVariables(v map[string]*string) *CreateApplicationRequest {
	s.Variables = v
	return s
}

type CreateApplicationRequestInstances struct {
	// The ID of the instance.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the instance.
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The type of the instance.
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s CreateApplicationRequestInstances) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequestInstances) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequestInstances) SetId(v string) *CreateApplicationRequestInstances {
	s.Id = &v
	return s
}

func (s *CreateApplicationRequestInstances) SetNodeName(v string) *CreateApplicationRequestInstances {
	s.NodeName = &v
	return s
}

func (s *CreateApplicationRequestInstances) SetNodeType(v string) *CreateApplicationRequestInstances {
	s.NodeType = &v
	return s
}

type CreateApplicationShrinkRequest struct {
	// The ID of the region.
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The parameters that are used to configure the application you want to create. For example, enableMonitor specifies whether to automatically create a CloudMonitor task for the application, and enableReport specifies whether to generate reports.
	ConfigurationShrink *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The instances in which you want to create the application. You can create applications in an existing virtual private cloud (VPC).
	InstancesShrink *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// The name of the application.
	//
	// *   The application name must be unique. You can call the [ListApplication](https://www.alibabacloud.com/help/en/bp-studio/latest/api-bpstudio-2021-09-31-listapplication) operation to query the existing applications.
	// *   The application name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http:// or https://`. The name can contain letters, digits, underscores (\_), and hyphens (-).
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the application you want to create belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the template.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The parameter values that are contained in the template. If the template contains no parameter values, the default values are used.
	VariablesShrink *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s CreateApplicationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationShrinkRequest) SetAreaId(v string) *CreateApplicationShrinkRequest {
	s.AreaId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetClientToken(v string) *CreateApplicationShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetConfigurationShrink(v string) *CreateApplicationShrinkRequest {
	s.ConfigurationShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetInstancesShrink(v string) *CreateApplicationShrinkRequest {
	s.InstancesShrink = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetName(v string) *CreateApplicationShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetResourceGroupId(v string) *CreateApplicationShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetTemplateId(v string) *CreateApplicationShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateApplicationShrinkRequest) SetVariablesShrink(v string) *CreateApplicationShrinkRequest {
	s.VariablesShrink = &v
	return s
}

type CreateApplicationResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the application.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) SetCode(v int32) *CreateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApplicationResponseBody) SetData(v string) *CreateApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *CreateApplicationResponseBody) SetMessage(v string) *CreateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

type CreateApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponse) SetHeaders(v map[string]*string) *CreateApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationResponse) SetStatusCode(v int32) *CreateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationResponse) SetBody(v *CreateApplicationResponseBody) *CreateApplicationResponse {
	s.Body = v
	return s
}

type DeleteApplicationRequest struct {
	// The ID of the application.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeleteApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) SetApplicationId(v string) *DeleteApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeleteApplicationRequest) SetResourceGroupId(v string) *DeleteApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

type DeleteApplicationResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponseBody) SetCode(v int32) *DeleteApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetMessage(v string) *DeleteApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApplicationResponseBody) SetRequestId(v string) *DeleteApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DeleteApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponse) SetHeaders(v map[string]*string) *DeleteApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationResponse) SetStatusCode(v int32) *DeleteApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationResponse) SetBody(v *DeleteApplicationResponseBody) *DeleteApplicationResponse {
	s.Body = v
	return s
}

type DeployApplicationRequest struct {
	// The ID of the application.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DeployApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeployApplicationRequest) SetApplicationId(v string) *DeployApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *DeployApplicationRequest) SetResourceGroupId(v string) *DeployApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

type DeployApplicationResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the application.
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeployApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeployApplicationResponseBody) SetCode(v int32) *DeployApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *DeployApplicationResponseBody) SetData(v int64) *DeployApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *DeployApplicationResponseBody) SetMessage(v string) *DeployApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *DeployApplicationResponseBody) SetRequestId(v string) *DeployApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DeployApplicationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeployApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeployApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeployApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeployApplicationResponse) SetHeaders(v map[string]*string) *DeployApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeployApplicationResponse) SetStatusCode(v int32) *DeployApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeployApplicationResponse) SetBody(v *DeployApplicationResponseBody) *DeployApplicationResponse {
	s.Body = v
	return s
}

type ExecuteOperationASyncRequest struct {
	// The ID of the Cloud Architect Design Tools (CADT) application.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The parameters related to the action. Specify the parameters based on the value of Operation. The parameters are passed in the map format. The following examples show how to specify the parameters if you want to change the specifications of an Elastic Compute Service (ECS) instance:
	//
	// *   The following common parameters are required: change_type, regionId, instanceId, appId
	// *   Example values for changing the instance type of the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_instance_type","instance_type":"ecs.hfr7.2xlarge","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	// *   Example values for stopping the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Stopped","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	// *   Example values for starting the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Running","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	// *   Example values for restarting the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Restart","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// Example of enumerating more than one set of parameters:
	//
	// *   { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{\\"change_type\\":\\"modify_instance_type\\",\\"instance_type\\":\\"ecs.hfr7.2xlarge\\",\\"instanceId\\":\\"i-xxxxxxxxx\\",\\"regionId\\":\\"cn-beijing\\",\\"appId\\":\\"xxxxxxxxxxxxx\\"}" }
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{\\"change_type\\":\\"modify_instance_type\\",\\"instance_type\\":\\"ecs.hfr7.2xlarge\\",\\"instanceId\\":\\"i-xxxxxxxxx\\",\\"regionId\\":\\"cn-beijing\\",\\"appId\\":\\"xxxxxxxxxxxxx\\"}" }
	//
	//     <!-- -->
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// This operation type is the operation type of modifying the product, some operation types are generic, and some are used alone. The following is an example of ECS deployment:
	// - The name of the ECS: rename
	// - Specificationof ecs: modifyInstanceType
	// - Startup of ecs: modifyInstanceType
	// - Stop of ecs: modifyInstanceType
	// - Restart of ecs: modifyInstanceType
	// - Ecs Tag: addTags
	// - Deletion of ecs: ecsDelete
	// - Paid type for ecs: modifyPayType
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// Resource group ID, which is used to verify the permissions of the resource group
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the service. If you want to perform operations on an Elastic Compute Service (ECS) instance, set ServiceType to ecs.
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ExecuteOperationASyncRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteOperationASyncRequest) GoString() string {
	return s.String()
}

func (s *ExecuteOperationASyncRequest) SetApplicationId(v string) *ExecuteOperationASyncRequest {
	s.ApplicationId = &v
	return s
}

func (s *ExecuteOperationASyncRequest) SetAttributes(v map[string]interface{}) *ExecuteOperationASyncRequest {
	s.Attributes = v
	return s
}

func (s *ExecuteOperationASyncRequest) SetOperation(v string) *ExecuteOperationASyncRequest {
	s.Operation = &v
	return s
}

func (s *ExecuteOperationASyncRequest) SetResourceGroupId(v string) *ExecuteOperationASyncRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ExecuteOperationASyncRequest) SetServiceType(v string) *ExecuteOperationASyncRequest {
	s.ServiceType = &v
	return s
}

type ExecuteOperationASyncShrinkRequest struct {
	// The ID of the Cloud Architect Design Tools (CADT) application.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The parameters related to the action. Specify the parameters based on the value of Operation. The parameters are passed in the map format. The following examples show how to specify the parameters if you want to change the specifications of an Elastic Compute Service (ECS) instance:
	//
	// *   The following common parameters are required: change_type, regionId, instanceId, appId
	// *   Example values for changing the instance type of the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_instance_type","instance_type":"ecs.hfr7.2xlarge","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	// *   Example values for stopping the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Stopped","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	// *   Example values for starting the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Running","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	// *   Example values for restarting the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Restart","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// Example of enumerating more than one set of parameters:
	//
	// *   { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{\\"change_type\\":\\"modify_instance_type\\",\\"instance_type\\":\\"ecs.hfr7.2xlarge\\",\\"instanceId\\":\\"i-xxxxxxxxx\\",\\"regionId\\":\\"cn-beijing\\",\\"appId\\":\\"xxxxxxxxxxxxx\\"}" }
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{\\"change_type\\":\\"modify_instance_type\\",\\"instance_type\\":\\"ecs.hfr7.2xlarge\\",\\"instanceId\\":\\"i-xxxxxxxxx\\",\\"regionId\\":\\"cn-beijing\\",\\"appId\\":\\"xxxxxxxxxxxxx\\"}" }
	//
	//     <!-- -->
	AttributesShrink *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// This operation type is the operation type of modifying the product, some operation types are generic, and some are used alone. The following is an example of ECS deployment:
	// - The name of the ECS: rename
	// - Specificationof ecs: modifyInstanceType
	// - Startup of ecs: modifyInstanceType
	// - Stop of ecs: modifyInstanceType
	// - Restart of ecs: modifyInstanceType
	// - Ecs Tag: addTags
	// - Deletion of ecs: ecsDelete
	// - Paid type for ecs: modifyPayType
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// Resource group ID, which is used to verify the permissions of the resource group
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the service. If you want to perform operations on an Elastic Compute Service (ECS) instance, set ServiceType to ecs.
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ExecuteOperationASyncShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteOperationASyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExecuteOperationASyncShrinkRequest) SetApplicationId(v string) *ExecuteOperationASyncShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *ExecuteOperationASyncShrinkRequest) SetAttributesShrink(v string) *ExecuteOperationASyncShrinkRequest {
	s.AttributesShrink = &v
	return s
}

func (s *ExecuteOperationASyncShrinkRequest) SetOperation(v string) *ExecuteOperationASyncShrinkRequest {
	s.Operation = &v
	return s
}

func (s *ExecuteOperationASyncShrinkRequest) SetResourceGroupId(v string) *ExecuteOperationASyncShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ExecuteOperationASyncShrinkRequest) SetServiceType(v string) *ExecuteOperationASyncShrinkRequest {
	s.ServiceType = &v
	return s
}

type ExecuteOperationASyncResponseBody struct {
	// Result code, 200 for success; Other representatives fail.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the operation.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Error message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteOperationASyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteOperationASyncResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteOperationASyncResponseBody) SetCode(v int32) *ExecuteOperationASyncResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteOperationASyncResponseBody) SetData(v string) *ExecuteOperationASyncResponseBody {
	s.Data = &v
	return s
}

func (s *ExecuteOperationASyncResponseBody) SetMessage(v string) *ExecuteOperationASyncResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteOperationASyncResponseBody) SetRequestId(v string) *ExecuteOperationASyncResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteOperationASyncResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteOperationASyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteOperationASyncResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteOperationASyncResponse) GoString() string {
	return s.String()
}

func (s *ExecuteOperationASyncResponse) SetHeaders(v map[string]*string) *ExecuteOperationASyncResponse {
	s.Headers = v
	return s
}

func (s *ExecuteOperationASyncResponse) SetStatusCode(v int32) *ExecuteOperationASyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteOperationASyncResponse) SetBody(v *ExecuteOperationASyncResponseBody) *ExecuteOperationASyncResponse {
	s.Body = v
	return s
}

type GetApplicationRequest struct {
	// The ID of the request.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Queries the basic information, verification results, billing results, and deployment results of an application.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) SetApplicationId(v string) *GetApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationRequest) SetResourceGroupId(v string) *GetApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

type GetApplicationResponseBody struct {
	// The response code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the application.
	Data *GetApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Possible application states:
	//
	// *   Creating: The application is being created.
	// *   Modified: The application has been modified.
	// *   Verifying: The application is being verified.
	// *   Verified_Failure: The application failed to pass the verification.
	// *   Verified_Success: The application has passed the verification.
	// *   Valuating: Fees are being calculated for the application.
	// *   Valuating_Failure: Fees failed to be calculated for the application.
	// *   Valuating_Success: Fees are calculated for the application.
	// *   Deploying: The application is being deployed.
	// *   Deployed_Failure: The application failed to be deployed.
	// *   Partially_Deployed_Success: Some resources of the application are deployed.
	// *   Deployed_Success: The application is deployed.
	// *   Destroying: The application is being released.
	// *   Delayed_Destroy: The application release is delayed.
	// *   Destroyed_Failure: The application failed to be released.
	// *   Partially_Destroyed_Success: Some resources of the application are released.
	// *   Destroyed_Success: The application is released.
	// *   Revised: The application architecture is adjusted.
	// *   Verifying_In_Revision: The application resources are being verified during architecture adjustment.
	// *   Verified_Failure_In_Revision: The application resources failed to pass the verification during architecture adjustment.
	// *   Verified_Success_In_Revision: The application resources are verified during architecture adjustment.
	// *   Valuating_In_Revision: Fees are being calculated for the application during architecture adjustment.
	// *   Valuating_Failure_In_Revision: Fees failed to be calculated for the application during architecture adjustment.
	// *   Valuating_Success_In_Revision: Fees are calculated for the application during architecture adjustment.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the application.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) SetCode(v string) *GetApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *GetApplicationResponseBody) SetData(v *GetApplicationResponseBodyData) *GetApplicationResponseBody {
	s.Data = v
	return s
}

func (s *GetApplicationResponseBody) SetMessage(v string) *GetApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplicationResponseBody) SetRequestId(v string) *GetApplicationResponseBody {
	s.RequestId = &v
	return s
}

type GetApplicationResponseBodyData struct {
	// The description of the application.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The resource tag.
	Checklist []*GetApplicationResponseBodyDataChecklist `json:"Checklist,omitempty" xml:"Checklist,omitempty" type:"Repeated"`
	// The URL of the application topology image.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The message returned for the request.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The resource type.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The URL of the image in the database.
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// The URL of the image in the database.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The billing results.
	PriceList []*GetApplicationResponseBodyDataPriceList `json:"PriceList,omitempty" xml:"PriceList,omitempty" type:"Repeated"`
	// 1411182597819805/topo-MCEXDI5EL2OM10NY.json
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource specification.
	ResourceList []*GetApplicationResponseBodyDataResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// Verification passed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// CADT application
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyData) SetApplicationId(v string) *GetApplicationResponseBodyData {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetChecklist(v []*GetApplicationResponseBodyDataChecklist) *GetApplicationResponseBodyData {
	s.Checklist = v
	return s
}

func (s *GetApplicationResponseBodyData) SetCreateTime(v string) *GetApplicationResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetDescription(v string) *GetApplicationResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetError(v string) *GetApplicationResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetImageURL(v string) *GetApplicationResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetName(v string) *GetApplicationResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetPriceList(v []*GetApplicationResponseBodyDataPriceList) *GetApplicationResponseBodyData {
	s.PriceList = v
	return s
}

func (s *GetApplicationResponseBodyData) SetResourceGroupId(v string) *GetApplicationResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetResourceList(v []*GetApplicationResponseBodyDataResourceList) *GetApplicationResponseBodyData {
	s.ResourceList = v
	return s
}

func (s *GetApplicationResponseBodyData) SetStatus(v string) *GetApplicationResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetApplicationResponseBodyData) SetTemplateId(v string) *GetApplicationResponseBodyData {
	s.TemplateId = &v
	return s
}

type GetApplicationResponseBodyDataChecklist struct {
	// The resource tag.
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The region in which the instance resides.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The message returned for verification.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The service code.
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// The name of the instance.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The verification result.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The resource specifications.
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
}

func (s GetApplicationResponseBodyDataChecklist) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyDataChecklist) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataChecklist) SetLifecycle(v string) *GetApplicationResponseBodyDataChecklist {
	s.Lifecycle = &v
	return s
}

func (s *GetApplicationResponseBodyDataChecklist) SetRegion(v string) *GetApplicationResponseBodyDataChecklist {
	s.Region = &v
	return s
}

func (s *GetApplicationResponseBodyDataChecklist) SetRemark(v string) *GetApplicationResponseBodyDataChecklist {
	s.Remark = &v
	return s
}

func (s *GetApplicationResponseBodyDataChecklist) SetResourceCode(v string) *GetApplicationResponseBodyDataChecklist {
	s.ResourceCode = &v
	return s
}

func (s *GetApplicationResponseBodyDataChecklist) SetResourceName(v string) *GetApplicationResponseBodyDataChecklist {
	s.ResourceName = &v
	return s
}

func (s *GetApplicationResponseBodyDataChecklist) SetResult(v string) *GetApplicationResponseBodyDataChecklist {
	s.Result = &v
	return s
}

func (s *GetApplicationResponseBodyDataChecklist) SetSpecification(v string) *GetApplicationResponseBodyDataChecklist {
	s.Specification = &v
	return s
}

type GetApplicationResponseBodyDataPriceList struct {
	// The billing method.
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The quantity.
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the instance.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Resource Fill Labels.
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The unit price of the instance.
	OnePrice *float64 `json:"OnePrice,omitempty" xml:"OnePrice,omitempty"`
	// The original price of the instance.
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The service duration.
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The total price.
	Price *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// Unit: USD per hour
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The region in which the instance resides.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The error message that is returned when a price query fails.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The service code.
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// The instance type. This parameter indicates the information about the instance type. For example, 192.168.0.0/16 may be returned for a Virtual Private Cloud (VPC) instance, ecs.g5.large may be returned for an Elastic Compute Service (ECS) instance, and slb.s1.small may be returned for a Server Load Balancer (SLB) instance. If the resource does not have a specific type, an empty value is returned.
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The creation mode. Valid values:\
	// 1: creates a new instance.\
	// 2: imports an instance.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetApplicationResponseBodyDataPriceList) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyDataPriceList) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataPriceList) SetChargeType(v string) *GetApplicationResponseBodyDataPriceList {
	s.ChargeType = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetCount(v int32) *GetApplicationResponseBodyDataPriceList {
	s.Count = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetInstanceName(v string) *GetApplicationResponseBodyDataPriceList {
	s.InstanceName = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetLifecycle(v string) *GetApplicationResponseBodyDataPriceList {
	s.Lifecycle = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetOnePrice(v float64) *GetApplicationResponseBodyDataPriceList {
	s.OnePrice = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetOriginalPrice(v float64) *GetApplicationResponseBodyDataPriceList {
	s.OriginalPrice = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetPeriod(v int32) *GetApplicationResponseBodyDataPriceList {
	s.Period = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetPrice(v float64) *GetApplicationResponseBodyDataPriceList {
	s.Price = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetPriceUnit(v string) *GetApplicationResponseBodyDataPriceList {
	s.PriceUnit = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetRegion(v string) *GetApplicationResponseBodyDataPriceList {
	s.Region = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetRemark(v string) *GetApplicationResponseBodyDataPriceList {
	s.Remark = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetResourceCode(v string) *GetApplicationResponseBodyDataPriceList {
	s.ResourceCode = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetSpecification(v string) *GetApplicationResponseBodyDataPriceList {
	s.Specification = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetType(v string) *GetApplicationResponseBodyDataPriceList {
	s.Type = &v
	return s
}

type GetApplicationResponseBodyDataResourceList struct {
	// The billing method.
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The resource tag.
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The deployment result.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The service code.
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// The instance ID.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the instance.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The resource deployment result.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetApplicationResponseBodyDataResourceList) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBodyDataResourceList) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBodyDataResourceList) SetChargeType(v string) *GetApplicationResponseBodyDataResourceList {
	s.ChargeType = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetLifecycle(v string) *GetApplicationResponseBodyDataResourceList {
	s.Lifecycle = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetRemark(v string) *GetApplicationResponseBodyDataResourceList {
	s.Remark = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetResourceCode(v string) *GetApplicationResponseBodyDataResourceList {
	s.ResourceCode = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetResourceId(v string) *GetApplicationResponseBodyDataResourceList {
	s.ResourceId = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetResourceName(v string) *GetApplicationResponseBodyDataResourceList {
	s.ResourceName = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetResourceType(v string) *GetApplicationResponseBodyDataResourceList {
	s.ResourceType = &v
	return s
}

func (s *GetApplicationResponseBodyDataResourceList) SetStatus(v string) *GetApplicationResponseBodyDataResourceList {
	s.Status = &v
	return s
}

type GetApplicationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationResponse) SetHeaders(v map[string]*string) *GetApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationResponse) SetStatusCode(v int32) *GetApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationResponse) SetBody(v *GetApplicationResponseBody) *GetApplicationResponse {
	s.Body = v
	return s
}

type GetExecuteOperationResultRequest struct {
	// The ID of the operation.
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the resource group. This parameter is specified to verify the permissions on the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetExecuteOperationResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetExecuteOperationResultRequest) GoString() string {
	return s.String()
}

func (s *GetExecuteOperationResultRequest) SetOperationId(v string) *GetExecuteOperationResultRequest {
	s.OperationId = &v
	return s
}

func (s *GetExecuteOperationResultRequest) SetResourceGroupId(v string) *GetExecuteOperationResultRequest {
	s.ResourceGroupId = &v
	return s
}

type GetExecuteOperationResultResponseBody struct {
	// The HTTP status code. A value of 200 indicates that the request is successful.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed result of the queried operation.
	Data *GetExecuteOperationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExecuteOperationResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetExecuteOperationResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetExecuteOperationResultResponseBody) SetCode(v int32) *GetExecuteOperationResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetExecuteOperationResultResponseBody) SetData(v *GetExecuteOperationResultResponseBodyData) *GetExecuteOperationResultResponseBody {
	s.Data = v
	return s
}

func (s *GetExecuteOperationResultResponseBody) SetMessage(v string) *GetExecuteOperationResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetExecuteOperationResultResponseBody) SetRequestId(v string) *GetExecuteOperationResultResponseBody {
	s.RequestId = &v
	return s
}

type GetExecuteOperationResultResponseBodyData struct {
	// The output of the operation.
	Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the operation.
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The status of the operation.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetExecuteOperationResultResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetExecuteOperationResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExecuteOperationResultResponseBodyData) SetArguments(v string) *GetExecuteOperationResultResponseBodyData {
	s.Arguments = &v
	return s
}

func (s *GetExecuteOperationResultResponseBodyData) SetMessage(v string) *GetExecuteOperationResultResponseBodyData {
	s.Message = &v
	return s
}

func (s *GetExecuteOperationResultResponseBodyData) SetOperationId(v string) *GetExecuteOperationResultResponseBodyData {
	s.OperationId = &v
	return s
}

func (s *GetExecuteOperationResultResponseBodyData) SetStatus(v string) *GetExecuteOperationResultResponseBodyData {
	s.Status = &v
	return s
}

type GetExecuteOperationResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExecuteOperationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExecuteOperationResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetExecuteOperationResultResponse) GoString() string {
	return s.String()
}

func (s *GetExecuteOperationResultResponse) SetHeaders(v map[string]*string) *GetExecuteOperationResultResponse {
	s.Headers = v
	return s
}

func (s *GetExecuteOperationResultResponse) SetStatusCode(v int32) *GetExecuteOperationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExecuteOperationResultResponse) SetBody(v *GetExecuteOperationResultResponseBody) *GetExecuteOperationResultResponse {
	s.Body = v
	return s
}

type GetFoTaskStatusRequest struct {
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetFoTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFoTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetFoTaskStatusRequest) SetTaskId(v int32) *GetFoTaskStatusRequest {
	s.TaskId = &v
	return s
}

type GetFoTaskStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFoTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFoTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetFoTaskStatusResponseBody) SetCode(v string) *GetFoTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetFoTaskStatusResponseBody) SetData(v string) *GetFoTaskStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetFoTaskStatusResponseBody) SetMessage(v string) *GetFoTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetFoTaskStatusResponseBody) SetRequestId(v string) *GetFoTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetFoTaskStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFoTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFoTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFoTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetFoTaskStatusResponse) SetHeaders(v map[string]*string) *GetFoTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetFoTaskStatusResponse) SetStatusCode(v int32) *GetFoTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFoTaskStatusResponse) SetBody(v *GetFoTaskStatusResponseBody) *GetFoTaskStatusResponse {
	s.Body = v
	return s
}

type GetPotentialFailZonesRequest struct {
	IsPlanId *bool   `json:"IsPlanId,omitempty" xml:"IsPlanId,omitempty"`
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
}

func (s GetPotentialFailZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPotentialFailZonesRequest) GoString() string {
	return s.String()
}

func (s *GetPotentialFailZonesRequest) SetIsPlanId(v bool) *GetPotentialFailZonesRequest {
	s.IsPlanId = &v
	return s
}

func (s *GetPotentialFailZonesRequest) SetObjectId(v string) *GetPotentialFailZonesRequest {
	s.ObjectId = &v
	return s
}

type GetPotentialFailZonesResponseBody struct {
	Code      *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPotentialFailZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPotentialFailZonesResponseBody) GoString() string {
	return s.String()
}

func (s *GetPotentialFailZonesResponseBody) SetCode(v string) *GetPotentialFailZonesResponseBody {
	s.Code = &v
	return s
}

func (s *GetPotentialFailZonesResponseBody) SetData(v []*string) *GetPotentialFailZonesResponseBody {
	s.Data = v
	return s
}

func (s *GetPotentialFailZonesResponseBody) SetMessage(v string) *GetPotentialFailZonesResponseBody {
	s.Message = &v
	return s
}

func (s *GetPotentialFailZonesResponseBody) SetRequestId(v string) *GetPotentialFailZonesResponseBody {
	s.RequestId = &v
	return s
}

type GetPotentialFailZonesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPotentialFailZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPotentialFailZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPotentialFailZonesResponse) GoString() string {
	return s.String()
}

func (s *GetPotentialFailZonesResponse) SetHeaders(v map[string]*string) *GetPotentialFailZonesResponse {
	s.Headers = v
	return s
}

func (s *GetPotentialFailZonesResponse) SetStatusCode(v int32) *GetPotentialFailZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPotentialFailZonesResponse) SetBody(v *GetPotentialFailZonesResponseBody) *GetPotentialFailZonesResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	// Template Area
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// ResourceGroup ID
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Template ID
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) SetRegion(v string) *GetTemplateRequest {
	s.Region = &v
	return s
}

func (s *GetTemplateRequest) SetResourceGroupId(v string) *GetTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateId(v string) *GetTemplateRequest {
	s.TemplateId = &v
	return s
}

type GetTemplateResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the template.
	Data *GetTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The interface returns information
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) SetCode(v int32) *GetTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetTemplateResponseBody) SetData(v *GetTemplateResponseBodyData) *GetTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetTemplateResponseBody) SetMessage(v string) *GetTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

type GetTemplateResponseBodyData struct {
	// The time when the template was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Template DescriptionD
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The path to the template schema image file
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// The name of the template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Template ID
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The details of the template variables.
	Variables []*GetTemplateResponseBodyDataVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s GetTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyData) SetCreateTime(v string) *GetTemplateResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetDescription(v string) *GetTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetImageURL(v string) *GetTemplateResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetName(v string) *GetTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetResourceGroupId(v string) *GetTemplateResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetTemplateId(v string) *GetTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetVariables(v []*GetTemplateResponseBodyDataVariables) *GetTemplateResponseBodyData {
	s.Variables = v
	return s
}

type GetTemplateResponseBodyDataVariables struct {
	// The name of the variable.
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The type of the variable.
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The default value of the variable.
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The value of the variable.
	Variable *string `json:"Variable,omitempty" xml:"Variable,omitempty"`
}

func (s GetTemplateResponseBodyDataVariables) String() string {
	return tea.Prettify(s)
}

func (s GetTemplateResponseBodyDataVariables) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBodyDataVariables) SetAttribute(v string) *GetTemplateResponseBodyDataVariables {
	s.Attribute = &v
	return s
}

func (s *GetTemplateResponseBodyDataVariables) SetDataType(v string) *GetTemplateResponseBodyDataVariables {
	s.DataType = &v
	return s
}

func (s *GetTemplateResponseBodyDataVariables) SetDefaultValue(v string) *GetTemplateResponseBodyDataVariables {
	s.DefaultValue = &v
	return s
}

func (s *GetTemplateResponseBodyDataVariables) SetVariable(v string) *GetTemplateResponseBodyDataVariables {
	s.Variable = &v
	return s
}

type GetTemplateResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetTokenRequest struct {
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetResourceGroupId(v string) *GetTokenRequest {
	s.ResourceGroupId = &v
	return s
}

type GetTokenResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the token.
	Data *GetTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) SetCode(v int32) *GetTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenResponseBody) SetData(v *GetTokenResponseBodyData) *GetTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenResponseBody) SetMessage(v string) *GetTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

type GetTokenResponseBodyData struct {
	// The AccessKey ID that is used to access OSS.
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret used to access OSS.
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The OSS bucket that is used to store the architecture image.
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS endpoint.
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The token that is used to access the Object Storage Service (OSS) bucket that stores the architecture image.
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The OSS bucket that is used to save data snapshots.
	SnapshotBucket *string `json:"SnapshotBucket,omitempty" xml:"SnapshotBucket,omitempty"`
}

func (s GetTokenResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBodyData) SetAccessKeyId(v string) *GetTokenResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetTokenResponseBodyData) SetAccessKeySecret(v string) *GetTokenResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetTokenResponseBodyData) SetBucket(v string) *GetTokenResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *GetTokenResponseBodyData) SetEndpoint(v string) *GetTokenResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetTokenResponseBodyData) SetSecurityToken(v string) *GetTokenResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetTokenResponseBodyData) SetSnapshotBucket(v string) *GetTokenResponseBodyData {
	s.SnapshotBucket = &v
	return s
}

type GetTokenResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetStatusCode(v int32) *GetTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
	s.Body = v
	return s
}

type InitAppFailOverRequest struct {
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s InitAppFailOverRequest) String() string {
	return tea.Prettify(s)
}

func (s InitAppFailOverRequest) GoString() string {
	return s.String()
}

func (s *InitAppFailOverRequest) SetApplicationId(v string) *InitAppFailOverRequest {
	s.ApplicationId = &v
	return s
}

type InitAppFailOverResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int32  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitAppFailOverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitAppFailOverResponseBody) GoString() string {
	return s.String()
}

func (s *InitAppFailOverResponseBody) SetCode(v string) *InitAppFailOverResponseBody {
	s.Code = &v
	return s
}

func (s *InitAppFailOverResponseBody) SetData(v int32) *InitAppFailOverResponseBody {
	s.Data = &v
	return s
}

func (s *InitAppFailOverResponseBody) SetMessage(v string) *InitAppFailOverResponseBody {
	s.Message = &v
	return s
}

func (s *InitAppFailOverResponseBody) SetRequestId(v string) *InitAppFailOverResponseBody {
	s.RequestId = &v
	return s
}

type InitAppFailOverResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitAppFailOverResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitAppFailOverResponse) String() string {
	return tea.Prettify(s)
}

func (s InitAppFailOverResponse) GoString() string {
	return s.String()
}

func (s *InitAppFailOverResponse) SetHeaders(v map[string]*string) *InitAppFailOverResponse {
	s.Headers = v
	return s
}

func (s *InitAppFailOverResponse) SetStatusCode(v int32) *InitAppFailOverResponse {
	s.StatusCode = &v
	return s
}

func (s *InitAppFailOverResponse) SetBody(v *InitAppFailOverResponseBody) *InitAppFailOverResponse {
	s.Body = v
	return s
}

type ListApplicationRequest struct {
	// Keywords in the app name
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The HTTP status code.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The ID of the resource group to which the application belongs.
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 1 update time,<br>2 creation time
	OrderType *int64 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceId      *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The status of the applications to be returned.
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationRequest) SetKeyword(v string) *ListApplicationRequest {
	s.Keyword = &v
	return s
}

func (s *ListApplicationRequest) SetMaxResults(v int32) *ListApplicationRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationRequest) SetNextToken(v int32) *ListApplicationRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationRequest) SetOrderType(v int64) *ListApplicationRequest {
	s.OrderType = &v
	return s
}

func (s *ListApplicationRequest) SetResourceGroupId(v string) *ListApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListApplicationRequest) SetResourceId(v string) *ListApplicationRequest {
	s.ResourceId = &v
	return s
}

func (s *ListApplicationRequest) SetStatus(v string) *ListApplicationRequest {
	s.Status = &v
	return s
}

func (s *ListApplicationRequest) SetTemplateId(v string) *ListApplicationRequest {
	s.TemplateId = &v
	return s
}

type ListApplicationResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// App listing information
	Data []*ListApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The interface returns information
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The query token returned in this call.
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the application.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBody) SetCode(v int32) *ListApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ListApplicationResponseBody) SetData(v []*ListApplicationResponseBodyData) *ListApplicationResponseBody {
	s.Data = v
	return s
}

func (s *ListApplicationResponseBody) SetMessage(v string) *ListApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ListApplicationResponseBody) SetNextToken(v int32) *ListApplicationResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationResponseBody) SetRequestId(v string) *ListApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationResponseBody) SetTotalCount(v int32) *ListApplicationResponseBody {
	s.TotalCount = &v
	return s
}

type ListApplicationResponseBodyData struct {
	// The application ID.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the application was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The URL of the application architecture image.
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// The name of the application.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the application belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the application.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListApplicationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBodyData) SetApplicationId(v string) *ListApplicationResponseBodyData {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetCreateTime(v string) *ListApplicationResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetImageURL(v string) *ListApplicationResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetName(v string) *ListApplicationResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetResourceGroupId(v string) *ListApplicationResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetStatus(v string) *ListApplicationResponseBodyData {
	s.Status = &v
	return s
}

type ListApplicationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationResponse) SetHeaders(v map[string]*string) *ListApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationResponse) SetStatusCode(v int32) *ListApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationResponse) SetBody(v *ListApplicationResponseBody) *ListApplicationResponse {
	s.Body = v
	return s
}

type ListFoCreatedAppsResponseBody struct {
	Code    *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    []*ListFoCreatedAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFoCreatedAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFoCreatedAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFoCreatedAppsResponseBody) SetCode(v string) *ListFoCreatedAppsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFoCreatedAppsResponseBody) SetData(v []*ListFoCreatedAppsResponseBodyData) *ListFoCreatedAppsResponseBody {
	s.Data = v
	return s
}

func (s *ListFoCreatedAppsResponseBody) SetMessage(v string) *ListFoCreatedAppsResponseBody {
	s.Message = &v
	return s
}

func (s *ListFoCreatedAppsResponseBody) SetRequestId(v string) *ListFoCreatedAppsResponseBody {
	s.RequestId = &v
	return s
}

type ListFoCreatedAppsResponseBodyData struct {
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ReportUrl     *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title         *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListFoCreatedAppsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListFoCreatedAppsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFoCreatedAppsResponseBodyData) SetApplicationId(v string) *ListFoCreatedAppsResponseBodyData {
	s.ApplicationId = &v
	return s
}

func (s *ListFoCreatedAppsResponseBodyData) SetReportUrl(v string) *ListFoCreatedAppsResponseBodyData {
	s.ReportUrl = &v
	return s
}

func (s *ListFoCreatedAppsResponseBodyData) SetStatus(v string) *ListFoCreatedAppsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListFoCreatedAppsResponseBodyData) SetTitle(v string) *ListFoCreatedAppsResponseBodyData {
	s.Title = &v
	return s
}

type ListFoCreatedAppsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFoCreatedAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFoCreatedAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFoCreatedAppsResponse) GoString() string {
	return s.String()
}

func (s *ListFoCreatedAppsResponse) SetHeaders(v map[string]*string) *ListFoCreatedAppsResponse {
	s.Headers = v
	return s
}

func (s *ListFoCreatedAppsResponse) SetStatusCode(v int32) *ListFoCreatedAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFoCreatedAppsResponse) SetBody(v *ListFoCreatedAppsResponseBody) *ListFoCreatedAppsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify a maximum number of 50 IDs.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. A maximum of 20 tags are supported.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetClientToken(v string) *ListTagResourcesRequest {
	s.ClientToken = &v
	return s
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
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
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
	// The HTTP status code. A value of 200 indicates that the request is successful.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. If the NextToken parameter is empty, no next page exists.
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags that are added to the resources.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetCode(v string) *ListTagResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetMessage(v string) *ListTagResourcesResponseBody {
	s.Message = &v
	return s
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
	// The resource type. Valid values: application and template.
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag.
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListTemplateRequest struct {
	// The keyword that is used to search for templates.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries to return on each page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The number of the page to return.
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The criterion by which the returned templates are sorted. Valid values:
	//
	// *   1: The templates are sorted by the time when they are updated.
	// *   2: The templates are sorted by the time when they are created.
	// *   3: The templates are sorted by the system.
	// *   4: The templates are sorted by the number of times that they are used.
	// *   If you specify an integer other than 1, 2, 3, and 4 or do not specify any value, the templates are sorted by the system.
	OrderType *int64 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag that you want to use to query templates.
	TagList *int32 `json:"TagList,omitempty" xml:"TagList,omitempty"`
	// The type of the templates to be returned. Valid values: public and private
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateRequest) GoString() string {
	return s.String()
}

func (s *ListTemplateRequest) SetKeyword(v string) *ListTemplateRequest {
	s.Keyword = &v
	return s
}

func (s *ListTemplateRequest) SetMaxResults(v int32) *ListTemplateRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTemplateRequest) SetNextToken(v int32) *ListTemplateRequest {
	s.NextToken = &v
	return s
}

func (s *ListTemplateRequest) SetOrderType(v int64) *ListTemplateRequest {
	s.OrderType = &v
	return s
}

func (s *ListTemplateRequest) SetResourceGroupId(v string) *ListTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplateRequest) SetTagList(v int32) *ListTemplateRequest {
	s.TagList = &v
	return s
}

func (s *ListTemplateRequest) SetType(v string) *ListTemplateRequest {
	s.Type = &v
	return s
}

type ListTemplateResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details about templates.
	Data []*ListTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateResponseBody) SetCode(v int32) *ListTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ListTemplateResponseBody) SetData(v []*ListTemplateResponseBodyData) *ListTemplateResponseBody {
	s.Data = v
	return s
}

func (s *ListTemplateResponseBody) SetMessage(v string) *ListTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ListTemplateResponseBody) SetNextToken(v int32) *ListTemplateResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplateResponseBody) SetRequestId(v string) *ListTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateResponseBody) SetTotalCount(v int32) *ListTemplateResponseBody {
	s.TotalCount = &v
	return s
}

type ListTemplateResponseBodyData struct {
	// The time when the template was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The URL of the architecture image.
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// The name of the template.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tag that is added to the template.
	TagId *int32 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The name of the tag that is added to the template.
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The ID of the template.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTemplateResponseBodyData) SetCreateTime(v string) *ListTemplateResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetImageURL(v string) *ListTemplateResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetName(v string) *ListTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetResourceGroupId(v string) *ListTemplateResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetTagId(v int32) *ListTemplateResponseBodyData {
	s.TagId = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetTagName(v string) *ListTemplateResponseBodyData {
	s.TagName = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetTemplateId(v string) *ListTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

type ListTemplateResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTemplateResponse) GoString() string {
	return s.String()
}

func (s *ListTemplateResponse) SetHeaders(v map[string]*string) *ListTemplateResponse {
	s.Headers = v
	return s
}

func (s *ListTemplateResponse) SetStatusCode(v int32) *ListTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTemplateResponse) SetBody(v *ListTemplateResponseBody) *ListTemplateResponse {
	s.Body = v
	return s
}

type ReleaseApplicationRequest struct {
	// The ID of the application.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the resource.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ReleaseApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseApplicationRequest) GoString() string {
	return s.String()
}

func (s *ReleaseApplicationRequest) SetApplicationId(v string) *ReleaseApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ReleaseApplicationRequest) SetResourceGroupId(v string) *ReleaseApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

type ReleaseApplicationResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value.
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseApplicationResponseBody) SetCode(v int32) *ReleaseApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseApplicationResponseBody) SetData(v int64) *ReleaseApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *ReleaseApplicationResponseBody) SetMessage(v string) *ReleaseApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseApplicationResponseBody) SetRequestId(v string) *ReleaseApplicationResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseApplicationResponse) GoString() string {
	return s.String()
}

func (s *ReleaseApplicationResponse) SetHeaders(v map[string]*string) *ReleaseApplicationResponse {
	s.Headers = v
	return s
}

func (s *ReleaseApplicationResponse) SetStatusCode(v int32) *ReleaseApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseApplicationResponse) SetBody(v *ReleaseApplicationResponseBody) *ReleaseApplicationResponse {
	s.Body = v
	return s
}

type ValidateApplicationRequest struct {
	// The ID of the application.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ValidateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationRequest) GoString() string {
	return s.String()
}

func (s *ValidateApplicationRequest) SetApplicationId(v string) *ValidateApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ValidateApplicationRequest) SetResourceGroupId(v string) *ValidateApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

type ValidateApplicationResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the application.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValidateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateApplicationResponseBody) SetCode(v int32) *ValidateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateApplicationResponseBody) SetData(v string) *ValidateApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *ValidateApplicationResponseBody) SetMessage(v string) *ValidateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateApplicationResponseBody) SetRequestId(v string) *ValidateApplicationResponseBody {
	s.RequestId = &v
	return s
}

type ValidateApplicationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateApplicationResponse) GoString() string {
	return s.String()
}

func (s *ValidateApplicationResponse) SetHeaders(v map[string]*string) *ValidateApplicationResponse {
	s.Headers = v
	return s
}

func (s *ValidateApplicationResponse) SetStatusCode(v int32) *ValidateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateApplicationResponse) SetBody(v *ValidateApplicationResponseBody) *ValidateApplicationResponse {
	s.Body = v
	return s
}

type ValuateApplicationRequest struct {
	// The ID of the application.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ValuateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ValuateApplicationRequest) GoString() string {
	return s.String()
}

func (s *ValuateApplicationRequest) SetApplicationId(v string) *ValuateApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ValuateApplicationRequest) SetResourceGroupId(v string) *ValuateApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

type ValuateApplicationResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the application.
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValuateApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValuateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ValuateApplicationResponseBody) SetCode(v int32) *ValuateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ValuateApplicationResponseBody) SetData(v int64) *ValuateApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *ValuateApplicationResponseBody) SetMessage(v string) *ValuateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ValuateApplicationResponseBody) SetRequestId(v string) *ValuateApplicationResponseBody {
	s.RequestId = &v
	return s
}

type ValuateApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValuateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValuateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ValuateApplicationResponse) GoString() string {
	return s.String()
}

func (s *ValuateApplicationResponse) SetHeaders(v map[string]*string) *ValuateApplicationResponse {
	s.Headers = v
	return s
}

func (s *ValuateApplicationResponse) SetStatusCode(v int32) *ValuateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ValuateApplicationResponse) SetBody(v *ValuateApplicationResponseBody) *ValuateApplicationResponse {
	s.Body = v
	return s
}

type ValuateTemplateRequest struct {
	// The region ID.
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instances to be replaced.
	Instances []*ValuateTemplateRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The ID of the resource group to which the application belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The template ID.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The parameter values that are contained in the template. If the template contains no parameter values, the default values are used.
	Variables map[string]*string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s ValuateTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ValuateTemplateRequest) GoString() string {
	return s.String()
}

func (s *ValuateTemplateRequest) SetAreaId(v string) *ValuateTemplateRequest {
	s.AreaId = &v
	return s
}

func (s *ValuateTemplateRequest) SetClientToken(v string) *ValuateTemplateRequest {
	s.ClientToken = &v
	return s
}

func (s *ValuateTemplateRequest) SetInstances(v []*ValuateTemplateRequestInstances) *ValuateTemplateRequest {
	s.Instances = v
	return s
}

func (s *ValuateTemplateRequest) SetResourceGroupId(v string) *ValuateTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ValuateTemplateRequest) SetTemplateId(v string) *ValuateTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ValuateTemplateRequest) SetVariables(v map[string]*string) *ValuateTemplateRequest {
	s.Variables = v
	return s
}

type ValuateTemplateRequestInstances struct {
	// The instance ID.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the application instance that is displayed on the diagram.
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The instance type.
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s ValuateTemplateRequestInstances) String() string {
	return tea.Prettify(s)
}

func (s ValuateTemplateRequestInstances) GoString() string {
	return s.String()
}

func (s *ValuateTemplateRequestInstances) SetId(v string) *ValuateTemplateRequestInstances {
	s.Id = &v
	return s
}

func (s *ValuateTemplateRequestInstances) SetNodeName(v string) *ValuateTemplateRequestInstances {
	s.NodeName = &v
	return s
}

func (s *ValuateTemplateRequestInstances) SetNodeType(v string) *ValuateTemplateRequestInstances {
	s.NodeType = &v
	return s
}

type ValuateTemplateShrinkRequest struct {
	// The region ID.
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instances to be replaced.
	InstancesShrink *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// The ID of the resource group to which the application belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The template ID.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The parameter values that are contained in the template. If the template contains no parameter values, the default values are used.
	VariablesShrink *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s ValuateTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ValuateTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ValuateTemplateShrinkRequest) SetAreaId(v string) *ValuateTemplateShrinkRequest {
	s.AreaId = &v
	return s
}

func (s *ValuateTemplateShrinkRequest) SetClientToken(v string) *ValuateTemplateShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ValuateTemplateShrinkRequest) SetInstancesShrink(v string) *ValuateTemplateShrinkRequest {
	s.InstancesShrink = &v
	return s
}

func (s *ValuateTemplateShrinkRequest) SetResourceGroupId(v string) *ValuateTemplateShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ValuateTemplateShrinkRequest) SetTemplateId(v string) *ValuateTemplateShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *ValuateTemplateShrinkRequest) SetVariablesShrink(v string) *ValuateTemplateShrinkRequest {
	s.VariablesShrink = &v
	return s
}

type ValuateTemplateResponseBody struct {
	// The HTTP status code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the inquiry.
	Data *ValuateTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the request failed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ValuateTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValuateTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ValuateTemplateResponseBody) SetCode(v string) *ValuateTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ValuateTemplateResponseBody) SetData(v *ValuateTemplateResponseBodyData) *ValuateTemplateResponseBody {
	s.Data = v
	return s
}

func (s *ValuateTemplateResponseBody) SetMessage(v string) *ValuateTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ValuateTemplateResponseBody) SetRequestId(v string) *ValuateTemplateResponseBody {
	s.RequestId = &v
	return s
}

type ValuateTemplateResponseBodyData struct {
	// The result set of the inquiry.
	ResourceList []*ValuateTemplateResponseBodyDataResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
}

func (s ValuateTemplateResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ValuateTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ValuateTemplateResponseBodyData) SetResourceList(v []*ValuateTemplateResponseBodyDataResourceList) *ValuateTemplateResponseBodyData {
	s.ResourceList = v
	return s
}

type ValuateTemplateResponseBodyDataResourceList struct {
	// The discount amount.
	DiscountAmount *float64 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The error message that is returned.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The resource type.
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The original price.
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The information about the price.
	PriceList []*ValuateTemplateResponseBodyDataResourceListPriceList `json:"PriceList,omitempty" xml:"PriceList,omitempty" type:"Repeated"`
	// The pricing unit.
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The discount information.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// The price at which the transaction is made.
	TradePrice *float64 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s ValuateTemplateResponseBodyDataResourceList) String() string {
	return tea.Prettify(s)
}

func (s ValuateTemplateResponseBodyDataResourceList) GoString() string {
	return s.String()
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetDiscountAmount(v float64) *ValuateTemplateResponseBodyDataResourceList {
	s.DiscountAmount = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetError(v string) *ValuateTemplateResponseBodyDataResourceList {
	s.Error = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetNodeType(v string) *ValuateTemplateResponseBodyDataResourceList {
	s.NodeType = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetOriginalPrice(v float64) *ValuateTemplateResponseBodyDataResourceList {
	s.OriginalPrice = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetPriceList(v []*ValuateTemplateResponseBodyDataResourceListPriceList) *ValuateTemplateResponseBodyDataResourceList {
	s.PriceList = v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetPriceUnit(v string) *ValuateTemplateResponseBodyDataResourceList {
	s.PriceUnit = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetPromotionName(v string) *ValuateTemplateResponseBodyDataResourceList {
	s.PromotionName = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceList) SetTradePrice(v float64) *ValuateTemplateResponseBodyDataResourceList {
	s.TradePrice = &v
	return s
}

type ValuateTemplateResponseBodyDataResourceListPriceList struct {
	// The discount amount.
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The error message that is returned.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The resource type.
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The original price.
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The pricing unit.
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The discount information.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// The resource ID.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The price at which the transaction is made.
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
	// Indicates whether the instance is newly created. Valid values:\
	// 1: The instance is newly created.\
	// 2: The instance already exists.\
	// 0: The price of the instance is not included.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ValuateTemplateResponseBodyDataResourceListPriceList) String() string {
	return tea.Prettify(s)
}

func (s ValuateTemplateResponseBodyDataResourceListPriceList) GoString() string {
	return s.String()
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetDiscountAmount(v float32) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.DiscountAmount = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetError(v string) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.Error = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetNodeType(v string) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.NodeType = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetOriginalPrice(v float32) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.OriginalPrice = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetPriceUnit(v string) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.PriceUnit = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetPromotionName(v string) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.PromotionName = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetResourceId(v string) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.ResourceId = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetTradePrice(v float32) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.TradePrice = &v
	return s
}

func (s *ValuateTemplateResponseBodyDataResourceListPriceList) SetType(v string) *ValuateTemplateResponseBodyDataResourceListPriceList {
	s.Type = &v
	return s
}

type ValuateTemplateResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValuateTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValuateTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ValuateTemplateResponse) GoString() string {
	return s.String()
}

func (s *ValuateTemplateResponse) SetHeaders(v map[string]*string) *ValuateTemplateResponse {
	s.Headers = v
	return s
}

func (s *ValuateTemplateResponse) SetStatusCode(v int32) *ValuateTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ValuateTemplateResponse) SetBody(v *ValuateTemplateResponseBody) *ValuateTemplateResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("bpstudio"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AppFailBackWithOptions(request *AppFailBackRequest, runtime *util.RuntimeOptions) (_result *AppFailBackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppFailBack"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppFailBackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppFailBack(request *AppFailBackRequest) (_result *AppFailBackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppFailBackResponse{}
	_body, _err := client.AppFailBackWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppFailOverWithOptions(request *AppFailOverRequest, runtime *util.RuntimeOptions) (_result *AppFailOverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.FailZone)) {
		body["FailZone"] = request.FailZone
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AppFailOver"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AppFailOverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppFailOver(request *AppFailOverRequest) (_result *AppFailOverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AppFailOverResponse{}
	_body, _err := client.AppFailOverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewResourceGroupId)) {
		body["NewResourceGroupId"] = request.NewResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApplicationWithOptions(tmpReq *CreateApplicationRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Configuration)) {
		request.ConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configuration, tea.String("Configuration"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Instances)) {
		request.InstancesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Instances, tea.String("Instances"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Variables)) {
		request.VariablesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Variables, tea.String("Variables"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AreaId)) {
		body["AreaId"] = request.AreaId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigurationShrink)) {
		body["Configuration"] = request.ConfigurationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstancesShrink)) {
		body["Instances"] = request.InstancesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.VariablesShrink)) {
		body["Variables"] = request.VariablesShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApplication"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApplication(request *CreateApplicationRequest) (_result *CreateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateApplicationResponse{}
	_body, _err := client.CreateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Before you call this operation to delete an application, make sure that the application is in the `Destroyed_Success` state. Otherwise, the application fails to be deleted.`` You can call the [GetApplication](https://www.alibabacloud.com/help/en/bp-studio/latest/api-bpstudio-2021-09-31-getapplication) operation to query the status of an application.
 *
 * @param request DeleteApplicationRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return DeleteApplicationResponse
 */
func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteApplication"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Before you call this operation to delete an application, make sure that the application is in the `Destroyed_Success` state. Otherwise, the application fails to be deleted.`` You can call the [GetApplication](https://www.alibabacloud.com/help/en/bp-studio/latest/api-bpstudio-2021-09-31-getapplication) operation to query the status of an application.
 *
 * @param request DeleteApplicationRequest
 * @return DeleteApplicationResponse
 */
func (client *Client) DeleteApplication(request *DeleteApplicationRequest) (_result *DeleteApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DeleteApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeployApplicationWithOptions(request *DeployApplicationRequest, runtime *util.RuntimeOptions) (_result *DeployApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeployApplication"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeployApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeployApplication(request *DeployApplicationRequest) (_result *DeployApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeployApplicationResponse{}
	_body, _err := client.DeployApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExecuteOperationASyncWithOptions(tmpReq *ExecuteOperationASyncRequest, runtime *util.RuntimeOptions) (_result *ExecuteOperationASyncResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExecuteOperationASyncShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Attributes)) {
		request.AttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Attributes, tea.String("Attributes"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.AttributesShrink)) {
		body["Attributes"] = request.AttributesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		body["Operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceType)) {
		body["ServiceType"] = request.ServiceType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExecuteOperationASync"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteOperationASyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExecuteOperationASync(request *ExecuteOperationASyncRequest) (_result *ExecuteOperationASyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteOperationASyncResponse{}
	_body, _err := client.ExecuteOperationASyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetApplicationWithOptions(request *GetApplicationRequest, runtime *util.RuntimeOptions) (_result *GetApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplication"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApplication(request *GetApplicationRequest) (_result *GetApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationResponse{}
	_body, _err := client.GetApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetExecuteOperationResultWithOptions(request *GetExecuteOperationResultRequest, runtime *util.RuntimeOptions) (_result *GetExecuteOperationResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OperationId)) {
		body["OperationId"] = request.OperationId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetExecuteOperationResult"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetExecuteOperationResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetExecuteOperationResult(request *GetExecuteOperationResultRequest) (_result *GetExecuteOperationResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetExecuteOperationResultResponse{}
	_body, _err := client.GetExecuteOperationResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFoTaskStatusWithOptions(request *GetFoTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetFoTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFoTaskStatus"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFoTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFoTaskStatus(request *GetFoTaskStatusRequest) (_result *GetFoTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFoTaskStatusResponse{}
	_body, _err := client.GetFoTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPotentialFailZonesWithOptions(request *GetPotentialFailZonesRequest, runtime *util.RuntimeOptions) (_result *GetPotentialFailZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsPlanId)) {
		body["IsPlanId"] = request.IsPlanId
	}

	if !tea.BoolValue(util.IsUnset(request.ObjectId)) {
		body["ObjectId"] = request.ObjectId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPotentialFailZones"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPotentialFailZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPotentialFailZones(request *GetPotentialFailZonesRequest) (_result *GetPotentialFailZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPotentialFailZonesResponse{}
	_body, _err := client.GetPotentialFailZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTemplateWithOptions(request *GetTemplateRequest, runtime *util.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Region)) {
		body["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTemplate"),
		Version:     tea.String("2021-09-31"),
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
 * @deprecated : GetToken is deprecated, please use BPStudio::2021-09-31::GetApplication instead.
 * ><danger> This API is no longer recommended, and the image related to the Application has included access authorization in the GetApplication property.></danger>
 *
 * @param request GetTokenRequest
 * @param runtime runtime options for this request RuntimeOptions
 * @return GetTokenResponse
 */
// Deprecated
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetToken"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * @deprecated : GetToken is deprecated, please use BPStudio::2021-09-31::GetApplication instead.
 * ><danger> This API is no longer recommended, and the image related to the Application has included access authorization in the GetApplication property.></danger>
 *
 * @param request GetTokenRequest
 * @return GetTokenResponse
 */
// Deprecated
func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitAppFailOverWithOptions(request *InitAppFailOverRequest, runtime *util.RuntimeOptions) (_result *InitAppFailOverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InitAppFailOver"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitAppFailOverResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitAppFailOver(request *InitAppFailOverRequest) (_result *InitAppFailOverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitAppFailOverResponse{}
	_body, _err := client.InitAppFailOverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationWithOptions(request *ListApplicationRequest, runtime *util.RuntimeOptions) (_result *ListApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		body["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplication"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplication(request *ListApplicationRequest) (_result *ListApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationResponse{}
	_body, _err := client.ListApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFoCreatedAppsWithOptions(runtime *util.RuntimeOptions) (_result *ListFoCreatedAppsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListFoCreatedApps"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFoCreatedAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFoCreatedApps() (_result *ListFoCreatedAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFoCreatedAppsResponse{}
	_body, _err := client.ListFoCreatedAppsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		body["RegionId"] = request.RegionId
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		bodyFlat["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		bodyFlat["Tag"] = request.Tag
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2021-09-31"),
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

func (client *Client) ListTemplateWithOptions(request *ListTemplateRequest, runtime *util.RuntimeOptions) (_result *ListTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		body["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		body["OrderType"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TagList)) {
		body["TagList"] = request.TagList
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTemplate"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTemplate(request *ListTemplateRequest) (_result *ListTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTemplateResponse{}
	_body, _err := client.ListTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseApplicationWithOptions(request *ReleaseApplicationRequest, runtime *util.RuntimeOptions) (_result *ReleaseApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseApplication"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseApplication(request *ReleaseApplicationRequest) (_result *ReleaseApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseApplicationResponse{}
	_body, _err := client.ReleaseApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateApplicationWithOptions(request *ValidateApplicationRequest, runtime *util.RuntimeOptions) (_result *ValidateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		query["ApplicationId"] = request.ApplicationId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateApplication"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateApplication(request *ValidateApplicationRequest) (_result *ValidateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValidateApplicationResponse{}
	_body, _err := client.ValidateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValuateApplicationWithOptions(request *ValuateApplicationRequest, runtime *util.RuntimeOptions) (_result *ValuateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ValuateApplication"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValuateApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValuateApplication(request *ValuateApplicationRequest) (_result *ValuateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValuateApplicationResponse{}
	_body, _err := client.ValuateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValuateTemplateWithOptions(tmpReq *ValuateTemplateRequest, runtime *util.RuntimeOptions) (_result *ValuateTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ValuateTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Instances)) {
		request.InstancesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Instances, tea.String("Instances"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Variables)) {
		request.VariablesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Variables, tea.String("Variables"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AreaId)) {
		body["AreaId"] = request.AreaId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstancesShrink)) {
		body["Instances"] = request.InstancesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.VariablesShrink)) {
		body["Variables"] = request.VariablesShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ValuateTemplate"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ValuateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValuateTemplate(request *ValuateTemplateRequest) (_result *ValuateTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ValuateTemplateResponse{}
	_body, _err := client.ValuateTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
