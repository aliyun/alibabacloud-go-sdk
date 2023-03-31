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

type ChangeResourceGroupRequest struct {
	// rg-aek2ajbjoloa23q
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// P7RMVSVM9LOVYQOM
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// APPLICATION
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
	Code    *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The name of the application that you want to create.
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
	// The name of the application that you want to create.
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeployApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The parameters that you need to specify when you perform an operation. The parameters vary based on the operation and are specified in the map format. The following examples show how to specify the parameters if you perform an operation on an ECS instance:
	//
	// *   The following common parameters need to be specified for operations on an ECS instance:
	//
	// change_type, regionId, instanceId, and appId.
	//
	// *   To change the instance type of an ECS instance, specify the following parameters:
	//
	// { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_instance_type","instance_type":"ecs.hfr7.2xlarge","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// *   To stop an ECS instance, specify the following parameters:
	//
	// { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Stopped","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// *   To start an ECS instance, specify the following parameters:
	//
	// { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Running","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// *   To restart an ECS instance, specify the following parameters:
	//
	// { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Restart","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// Enumeration values
	//
	// <!-- -->
	//
	// :
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
	Attributes map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The type of the operation to be performed to modify the parameters of an instance of the specified service. Some operations are common to different services, and other operations are specific to each service. For example, set this parameter to one of the following values to perform the corresponding operation on an ECS instance:
	//
	// *   rename: modifies the name of an ECS instance.
	// *   modifyInstanceType: changes the instance type of an ECS instance.
	// *   modifyInstanceType: starts an ECS instance.
	// *   modifyInstanceType: stops an ECS instance.
	// *   modifyInstanceType: restarts an ECS instance.
	// *   addTags: adds tags to an ECS instance.
	// *   ecsDelete: deletes an ECS instance.
	// *   modifyPayType: changes the billing method of an ECS instance.
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The ID of the resource group. This parameter is specified to verify the permissions on the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the service. If you want to modify the parameters of an Elastic Compute Service (ECS) instance, set this parameter to ecs.
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ExecuteOperationASyncRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteOperationASyncRequest) GoString() string {
	return s.String()
}

func (s *ExecuteOperationASyncRequest) SetAttributes(v map[string]*string) *ExecuteOperationASyncRequest {
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
	// The parameters that you need to specify when you perform an operation. The parameters vary based on the operation and are specified in the map format. The following examples show how to specify the parameters if you perform an operation on an ECS instance:
	//
	// *   The following common parameters need to be specified for operations on an ECS instance:
	//
	// change_type, regionId, instanceId, and appId.
	//
	// *   To change the instance type of an ECS instance, specify the following parameters:
	//
	// { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_instance_type","instance_type":"ecs.hfr7.2xlarge","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// *   To stop an ECS instance, specify the following parameters:
	//
	// { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Stopped","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// *   To start an ECS instance, specify the following parameters:
	//
	// { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Running","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// *   To restart an ECS instance, specify the following parameters:
	//
	// { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Restart","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// Enumeration values
	//
	// <!-- -->
	//
	// :
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
	AttributesShrink *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The type of the operation to be performed to modify the parameters of an instance of the specified service. Some operations are common to different services, and other operations are specific to each service. For example, set this parameter to one of the following values to perform the corresponding operation on an ECS instance:
	//
	// *   rename: modifies the name of an ECS instance.
	// *   modifyInstanceType: changes the instance type of an ECS instance.
	// *   modifyInstanceType: starts an ECS instance.
	// *   modifyInstanceType: stops an ECS instance.
	// *   modifyInstanceType: restarts an ECS instance.
	// *   addTags: adds tags to an ECS instance.
	// *   ecsDelete: deletes an ECS instance.
	// *   modifyPayType: changes the billing method of an ECS instance.
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The ID of the resource group. This parameter is specified to verify the permissions on the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the service. If you want to modify the parameters of an Elastic Compute Service (ECS) instance, set this parameter to ecs.
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ExecuteOperationASyncShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteOperationASyncShrinkRequest) GoString() string {
	return s.String()
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
	// The HTTP status code. A value of 200 indicates that the request is successful, and other values indicate that the request fails.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the operation.
	Data map[string]*string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
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

func (s *ExecuteOperationASyncResponseBody) SetData(v map[string]*string) *ExecuteOperationASyncResponseBody {
	s.Data = v
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ExecuteOperationASyncResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The ID of the application.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the resource group.
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
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the application.
	Data *GetApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned for the request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseBody) SetCode(v int32) *GetApplicationResponseBody {
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
	// The ID of the application.
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The verification results returned.
	Checklist []*GetApplicationResponseBodyDataChecklist `json:"Checklist,omitempty" xml:"Checklist,omitempty" type:"Repeated"`
	// The time when the application was created.
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the application.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The message returned for deployment.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The URL of the image in the database.
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// The name of the application.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The billing results.
	PriceList []*GetApplicationResponseBodyDataPriceList `json:"PriceList,omitempty" xml:"PriceList,omitempty" type:"Repeated"`
	// The ID of the resource group to which the application belongs.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resources.
	ResourceList []*GetApplicationResponseBodyDataResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// The status of the application.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the template associated with the application.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the application topology image.
	TopoURL *string `json:"TopoURL,omitempty" xml:"TopoURL,omitempty"`
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

func (s *GetApplicationResponseBodyData) SetTopoURL(v string) *GetApplicationResponseBodyData {
	s.TopoURL = &v
	return s
}

type GetApplicationResponseBodyDataChecklist struct {
	// The resource tag.
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The ID of the region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The message returned for verification.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The service code.
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// The name of the instance.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The verification result.
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The resource specification.
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
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the instance.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The resource tag.
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The unit price.
	OnePrice *float32 `json:"OnePrice,omitempty" xml:"OnePrice,omitempty"`
	// The original price.
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The service duration.
	Period *float32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The total price.
	Price *float32 `json:"Price,omitempty" xml:"Price,omitempty"`
	// The price unit.
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The ID of the region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The error message that is returned when a price query fails.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The service code.
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// The instance type.
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
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

func (s *GetApplicationResponseBodyDataPriceList) SetCount(v int64) *GetApplicationResponseBodyDataPriceList {
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

func (s *GetApplicationResponseBodyDataPriceList) SetOnePrice(v float32) *GetApplicationResponseBodyDataPriceList {
	s.OnePrice = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetOriginalPrice(v float32) *GetApplicationResponseBodyDataPriceList {
	s.OriginalPrice = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetPeriod(v float32) *GetApplicationResponseBodyDataPriceList {
	s.Period = &v
	return s
}

func (s *GetApplicationResponseBodyDataPriceList) SetPrice(v float32) *GetApplicationResponseBodyDataPriceList {
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

type GetApplicationResponseBodyDataResourceList struct {
	// The billing method.
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The resource tag.
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The deployment result.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The service code.
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// The ID of the instance.
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the instance.
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	// The details of the operation.
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetExecuteOperationResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetTemplateRequest struct {
	// The region in which the template resides.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the template.
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
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
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
	// The description of the template.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The URL of the architecture image.
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// The name of the template.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the template.
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The URL of the architecture image file.
	TopoURL   *string                                 `json:"TopoURL,omitempty" xml:"TopoURL,omitempty"`
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

func (s *GetTemplateResponseBodyData) SetTopoURL(v string) *GetTemplateResponseBodyData {
	s.TopoURL = &v
	return s
}

func (s *GetTemplateResponseBodyData) SetVariables(v []*GetTemplateResponseBodyDataVariables) *GetTemplateResponseBodyData {
	s.Variables = v
	return s
}

type GetTemplateResponseBodyDataVariables struct {
	Attribute    *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	DataType     *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Variable     *string `json:"Variable,omitempty" xml:"Variable,omitempty"`
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
	// The AccessKey secret that is used to access OSS.
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListApplicationRequest struct {
	// The keyword in the application name.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries to return on each page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The number of the page to return.
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The field by which the list is sorted. Valid values: 1 and 2. The value 1 specifies that the list is sorted in descending order of application update time, and the value 2 specifies that the list is sorted in descending order of application creation time.
	OrderType *int64 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the application.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s *ListApplicationRequest) SetStatus(v string) *ListApplicationRequest {
	s.Status = &v
	return s
}

type ListApplicationResponseBody struct {
	// The HTTP status code.
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the applications.
	Data []*ListApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the next page.
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned applications.
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
	// The ID of the application.
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
	// The URL of the application topology image.
	TopoURL *string `json:"TopoURL,omitempty" xml:"TopoURL,omitempty"`
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

func (s *ListApplicationResponseBodyData) SetTopoURL(v string) *ListApplicationResponseBodyData {
	s.TopoURL = &v
	return s
}

type ListApplicationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListTagResourcesRequest struct {
	ClientToken  *string                       `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// TagResource
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
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
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

type ListTemplateRequest struct {
	// The keyword that is used to search for templates.
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries to return on each page.
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The number of the page to return.
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The criterion by which the templates are sorted. Valid values:
	//
	// *   1: modification time
	// *   2: creation time
	OrderType *int64 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag of the template.
	TagList *int32 `json:"TagList,omitempty" xml:"TagList,omitempty"`
	// The type of the template.
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
	// The URL of the template topology image.
	TopoURL *string `json:"TopoURL,omitempty" xml:"TopoURL,omitempty"`
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

func (s *ListTemplateResponseBodyData) SetTopoURL(v string) *ListTemplateResponseBodyData {
	s.TopoURL = &v
	return s
}

type ListTemplateResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValidateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValuateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
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
