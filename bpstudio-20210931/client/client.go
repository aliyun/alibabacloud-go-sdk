// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AppFailBackRequest struct {
	// The application ID.
	//
	// example:
	//
	// 61ZW1DY5Y3FSAOO2
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
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The disaster recovery switchback task ID.
	//
	// example:
	//
	// 3309
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message. If the request was successful, a success message is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// OKITHEVRQCN6ULQG
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7036DDBE-0ABA-52D7-A39D-75E511970F07
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
	// The application ID.
	//
	// example:
	//
	// BE68D71ZY5YYIU9R
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The destination zone to which you want to switch the disaster recovery application.
	//
	// example:
	//
	// cn-hangzhou-g
	FailZone *string `json:"FailZone,omitempty" xml:"FailZone,omitempty"`
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
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The disaster recovery switchover task ID.
	//
	// example:
	//
	// 7441
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message. If the request was successful, a success message is returned. If the request failed, an error message is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9656C816-1E9A-58D2-86D5-710678D61AF1
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
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aek2ajbjoloa23q
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// The ID of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// P7RMVSVM9LOVYQOM
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource for which you want to change the resource group. Valid values: APPLICATION and TEMPLATE.
	//
	// This parameter is required.
	//
	// example:
	//
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
	// The HTTP status code. A value of 200 indicates that the request is successful. Other values indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// No business data is returned for this parameter.
	//
	// example:
	//
	// No business data is returned for this parameter.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A3488F1D-C444-17D0-BA4F-5374BA0F3562
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
	//
	// example:
	//
	// cn-hangzhou
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The parameters that are used to configure the application you want to create. For example, enableMonitor specifies whether to automatically create a CloudMonitor task for the application, and enableReport specifies whether to generate reports.
	//
	// example:
	//
	// {"enableMonitor":"0", "enableReport":"1"}
	Configuration map[string]*string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The instances in which you want to create the application. You can create applications in an existing virtual private cloud (VPC).
	Instances []*CreateApplicationRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The name of the application.
	//
	// 	- The application name must be unique. You can call the [ListApplication](https://www.alibabacloud.com/help/en/bp-studio/latest/api-bpstudio-2021-09-31-listapplication) operation to query the existing applications.
	//
	// 	- The application name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http:// or https://`. The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// cadt-application
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the application you want to create belongs.
	//
	// example:
	//
	// rg-acfmyjt3c5om3hi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0KSHPM6SJU03TNZP
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The parameter values that are contained in the template. If the template contains no parameter values, the default values are used.
	//
	// example:
	//
	// {"variable1":"1"}
	Variables map[string]interface{} `json:"Variables,omitempty" xml:"Variables,omitempty"`
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

func (s *CreateApplicationRequest) SetVariables(v map[string]interface{}) *CreateApplicationRequest {
	s.Variables = v
	return s
}

type CreateApplicationRequestInstances struct {
	// The ID of the instance.
	//
	// example:
	//
	// vpc-bp1q56trhtaq40vlq5ojm
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// vpc
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// vpc
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
	//
	// example:
	//
	// cn-hangzhou
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The parameters that are used to configure the application you want to create. For example, enableMonitor specifies whether to automatically create a CloudMonitor task for the application, and enableReport specifies whether to generate reports.
	//
	// example:
	//
	// {"enableMonitor":"0", "enableReport":"1"}
	ConfigurationShrink *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// The instances in which you want to create the application. You can create applications in an existing virtual private cloud (VPC).
	InstancesShrink *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// The name of the application.
	//
	// 	- The application name must be unique. You can call the [ListApplication](https://www.alibabacloud.com/help/en/bp-studio/latest/api-bpstudio-2021-09-31-listapplication) operation to query the existing applications.
	//
	// 	- The application name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http:// or https://`. The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// cadt-application
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the application you want to create belongs.
	//
	// example:
	//
	// rg-acfmyjt3c5om3hi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0KSHPM6SJU03TNZP
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The parameter values that are contained in the template. If the template contains no parameter values, the default values are used.
	//
	// example:
	//
	// {"variable1":"1"}
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 002XWH7MXB8MJRU0
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
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
	//
	// This parameter is required.
	//
	// example:
	//
	// 002XWH7MXB8MJRU0
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyjt3c5om3hi
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
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
	//
	// This parameter is required.
	//
	// example:
	//
	// VVK605ZH00OA4MRT
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyjt3c5om3hi
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

func (s *DeployApplicationRequest) SetClientToken(v string) *DeployApplicationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeployApplicationRequest) SetResourceGroupId(v string) *DeployApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

type DeployApplicationResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the application.
	//
	// example:
	//
	// 123
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
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
	//
	// example:
	//
	// BG**********UQ
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The parameters related to the action. Specify the parameters based on the value of Operation. The parameters are passed in the map format. The following examples show how to specify the parameters if you want to change the specifications of an Elastic Compute Service (ECS) instance:
	//
	// 	- The following common parameters are required: change_type, regionId, instanceId, appId
	//
	// 	- Example values for changing the instance type of the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_instance_type","instance_type":"ecs.hfr7.2xlarge","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// 	- Example values for stopping the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Stopped","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// 	- Example values for starting the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Running","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// 	- Example values for restarting the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Restart","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// Example of enumerating more than one set of parameters:
	//
	// 	- { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{\\\\"change_type\\\\":\\\\"modify_instance_type\\\\",\\\\"instance_type\\\\":\\\\"ecs.hfr7.2xlarge\\\\",\\\\"instanceId\\\\":\\\\"i-xxxxxxxxx\\\\",\\\\"regionId\\\\":\\\\"cn-beijing\\\\",\\\\"appId\\\\":\\\\"xxxxxxxxxxxxx\\\\"}" }
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{\\\\"change_type\\\\":\\\\"modify_instance_type\\\\",\\\\"instance_type\\\\":\\\\"ecs.hfr7.2xlarge\\\\",\\\\"instanceId\\\\":\\\\"i-xxxxxxxxx\\\\",\\\\"regionId\\\\":\\\\"cn-beijing\\\\",\\\\"appId\\\\":\\\\"xxxxxxxxxxxxx\\\\"}" }
	//
	//     <!-- -->
	Attributes map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This operation type is the operation type of modifying the product, some operation types are generic, and some are used alone. The following is an example of ECS deployment:
	//
	// - The name of the ECS: rename
	//
	// - Specificationof ecs: modifyInstanceType
	//
	// - Startup of ecs: modifyInstanceType
	//
	// - Stop of ecs: modifyInstanceType
	//
	// - Restart of ecs: modifyInstanceType
	//
	// - Ecs Tag: addTags
	//
	// - Deletion of ecs: ecsDelete
	//
	// - Paid type for ecs: modifyPayType
	//
	// This parameter is required.
	//
	// example:
	//
	// queryTopo
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// Resource group ID, which is used to verify the permissions of the resource group
	//
	// example:
	//
	// ceshi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the service. If you want to perform operations on an Elastic Compute Service (ECS) instance, set ServiceType to ecs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACK
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

func (s *ExecuteOperationASyncRequest) SetClientToken(v string) *ExecuteOperationASyncRequest {
	s.ClientToken = &v
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
	//
	// example:
	//
	// BG**********UQ
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The parameters related to the action. Specify the parameters based on the value of Operation. The parameters are passed in the map format. The following examples show how to specify the parameters if you want to change the specifications of an Elastic Compute Service (ECS) instance:
	//
	// 	- The following common parameters are required: change_type, regionId, instanceId, appId
	//
	// 	- Example values for changing the instance type of the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_instance_type","instance_type":"ecs.hfr7.2xlarge","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// 	- Example values for stopping the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Stopped","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// 	- Example values for starting the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Running","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// 	- Example values for restarting the ECS instance: { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{"change_type":"modify_status","status":"Restart","instanceId":"i-xxxxxxxxx","regionId":"cn-beijing","appId":"xxxxxxxxxxxxx"}" }
	//
	// Example of enumerating more than one set of parameters:
	//
	// 	- { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{\\\\"change_type\\\\":\\\\"modify_instance_type\\\\",\\\\"instance_type\\\\":\\\\"ecs.hfr7.2xlarge\\\\",\\\\"instanceId\\\\":\\\\"i-xxxxxxxxx\\\\",\\\\"regionId\\\\":\\\\"cn-beijing\\\\",\\\\"appId\\\\":\\\\"xxxxxxxxxxxxx\\\\"}" }
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     { "ServiceType": "ecs", "Operation": "modifyInstanceType", "Attributes": "{\\\\"change_type\\\\":\\\\"modify_instance_type\\\\",\\\\"instance_type\\\\":\\\\"ecs.hfr7.2xlarge\\\\",\\\\"instanceId\\\\":\\\\"i-xxxxxxxxx\\\\",\\\\"regionId\\\\":\\\\"cn-beijing\\\\",\\\\"appId\\\\":\\\\"xxxxxxxxxxxxx\\\\"}" }
	//
	//     <!-- -->
	AttributesShrink *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This operation type is the operation type of modifying the product, some operation types are generic, and some are used alone. The following is an example of ECS deployment:
	//
	// - The name of the ECS: rename
	//
	// - Specificationof ecs: modifyInstanceType
	//
	// - Startup of ecs: modifyInstanceType
	//
	// - Stop of ecs: modifyInstanceType
	//
	// - Restart of ecs: modifyInstanceType
	//
	// - Ecs Tag: addTags
	//
	// - Deletion of ecs: ecsDelete
	//
	// - Paid type for ecs: modifyPayType
	//
	// This parameter is required.
	//
	// example:
	//
	// queryTopo
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// Resource group ID, which is used to verify the permissions of the resource group
	//
	// example:
	//
	// ceshi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the service. If you want to perform operations on an Elastic Compute Service (ECS) instance, set ServiceType to ecs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACK
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

func (s *ExecuteOperationASyncShrinkRequest) SetClientToken(v string) *ExecuteOperationASyncShrinkRequest {
	s.ClientToken = &v
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the operation.
	//
	// example:
	//
	// op_xxxxxxxxxxxxxxxxxx_ecs_modifyInstanceType_BYSOQGWUV6PME412_ERMEZLXNN3K9N3OL
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Error message
	//
	// example:
	//
	// " "
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// SD-WEF-DSW-32ED-323DDSD-2332D
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

type ExecuteOperationSyncRequest struct {
	// example:
	//
	// BE68D71ZY5YYIU9R
	ApplicationId *string                `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Attributes    map[string]interface{} `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rename
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// rg-acfmyjt3c5om3hi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dds
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ExecuteOperationSyncRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteOperationSyncRequest) GoString() string {
	return s.String()
}

func (s *ExecuteOperationSyncRequest) SetApplicationId(v string) *ExecuteOperationSyncRequest {
	s.ApplicationId = &v
	return s
}

func (s *ExecuteOperationSyncRequest) SetAttributes(v map[string]interface{}) *ExecuteOperationSyncRequest {
	s.Attributes = v
	return s
}

func (s *ExecuteOperationSyncRequest) SetClientToken(v string) *ExecuteOperationSyncRequest {
	s.ClientToken = &v
	return s
}

func (s *ExecuteOperationSyncRequest) SetOperation(v string) *ExecuteOperationSyncRequest {
	s.Operation = &v
	return s
}

func (s *ExecuteOperationSyncRequest) SetResourceGroupId(v string) *ExecuteOperationSyncRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ExecuteOperationSyncRequest) SetServiceType(v string) *ExecuteOperationSyncRequest {
	s.ServiceType = &v
	return s
}

type ExecuteOperationSyncShrinkRequest struct {
	// example:
	//
	// BE68D71ZY5YYIU9R
	ApplicationId    *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	AttributesShrink *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rename
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// rg-acfmyjt3c5om3hi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dds
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
}

func (s ExecuteOperationSyncShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExecuteOperationSyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExecuteOperationSyncShrinkRequest) SetApplicationId(v string) *ExecuteOperationSyncShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *ExecuteOperationSyncShrinkRequest) SetAttributesShrink(v string) *ExecuteOperationSyncShrinkRequest {
	s.AttributesShrink = &v
	return s
}

func (s *ExecuteOperationSyncShrinkRequest) SetClientToken(v string) *ExecuteOperationSyncShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ExecuteOperationSyncShrinkRequest) SetOperation(v string) *ExecuteOperationSyncShrinkRequest {
	s.Operation = &v
	return s
}

func (s *ExecuteOperationSyncShrinkRequest) SetResourceGroupId(v string) *ExecuteOperationSyncShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ExecuteOperationSyncShrinkRequest) SetServiceType(v string) *ExecuteOperationSyncShrinkRequest {
	s.ServiceType = &v
	return s
}

type ExecuteOperationSyncResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// op_xxxxxxxxxxxxxxxxxx_dds_modifyInstanceType_BYSOQGWUV6PME412_ERMEZLXNN3K9N3OL
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Access key ID cannot be null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FDC73B05-5331-57AA-BA93-4C9882792FF5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ExecuteOperationSyncResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExecuteOperationSyncResponseBody) GoString() string {
	return s.String()
}

func (s *ExecuteOperationSyncResponseBody) SetCode(v int32) *ExecuteOperationSyncResponseBody {
	s.Code = &v
	return s
}

func (s *ExecuteOperationSyncResponseBody) SetData(v string) *ExecuteOperationSyncResponseBody {
	s.Data = &v
	return s
}

func (s *ExecuteOperationSyncResponseBody) SetMessage(v string) *ExecuteOperationSyncResponseBody {
	s.Message = &v
	return s
}

func (s *ExecuteOperationSyncResponseBody) SetRequestId(v string) *ExecuteOperationSyncResponseBody {
	s.RequestId = &v
	return s
}

type ExecuteOperationSyncResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExecuteOperationSyncResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExecuteOperationSyncResponse) String() string {
	return tea.Prettify(s)
}

func (s ExecuteOperationSyncResponse) GoString() string {
	return s.String()
}

func (s *ExecuteOperationSyncResponse) SetHeaders(v map[string]*string) *ExecuteOperationSyncResponse {
	s.Headers = v
	return s
}

func (s *ExecuteOperationSyncResponse) SetStatusCode(v int32) *ExecuteOperationSyncResponse {
	s.StatusCode = &v
	return s
}

func (s *ExecuteOperationSyncResponse) SetBody(v *ExecuteOperationSyncResponseBody) *ExecuteOperationSyncResponse {
	s.Body = v
	return s
}

type GetApplicationRequest struct {
	// The ID of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// VVK605ZH00OA4MRT
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Queries the basic information, verification results, billing results, and deployment results of an application.
	//
	// example:
	//
	// rg-acfmyjt3c5om3hi
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
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the application.
	Data *GetApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Reason for the request failure
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
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
	// App ID
	//
	// example:
	//
	// VVK605ZH00OA4MRT
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The resource tag.
	Checklist []*GetApplicationResponseBodyDataChecklist `json:"Checklist,omitempty" xml:"Checklist,omitempty" type:"Repeated"`
	// The time when the app was created
	//
	// example:
	//
	// 2021-08-09 14:37:16
	CreateTime    *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeployPercent *float64 `json:"DeployPercent,omitempty" xml:"DeployPercent,omitempty"`
	// Application description
	//
	// example:
	//
	// remark
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Success
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The URL of the image in the database.
	//
	// example:
	//
	// The details of the application.
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// App name
	//
	// example:
	//
	// 1411182597819805/sr-8DWU4RUS49NIDII0.png
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The billing results.
	PriceList []*GetApplicationResponseBodyDataPriceList `json:"PriceList,omitempty" xml:"PriceList,omitempty" type:"Repeated"`
	// The ID of the resource group to which the app belongs
	//
	// example:
	//
	// rg-aekzhfgmw4e6fwq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource specification.
	ResourceList []*GetApplicationResponseBodyDataResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// Verification passed
	//
	// example:
	//
	// Deployed_Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the template associated with the application
	//
	// example:
	//
	// FYS9VZ535U20V7HT
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

func (s *GetApplicationResponseBodyData) SetDeployPercent(v float64) *GetApplicationResponseBodyData {
	s.DeployPercent = &v
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
	//
	// example:
	//
	// Create
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The message returned for verification.
	//
	// example:
	//
	// The ID of the region.
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The service code.
	//
	// example:
	//
	// vpc
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// vpc
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The verification result.
	//
	// example:
	//
	// Finish
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The resource specifications.
	//
	// example:
	//
	// 192.168.0.0/16
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
	//
	// example:
	//
	// PayAsYouGo
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The quantity.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// ecs
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Resource Fill Labels.
	//
	// example:
	//
	// Create
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	// The unit price of the instance.
	//
	// example:
	//
	// 0.01
	OnePrice *float64 `json:"OnePrice,omitempty" xml:"OnePrice,omitempty"`
	// The original price of the instance.
	//
	// example:
	//
	// 3.570
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The service duration.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The total price.
	//
	// example:
	//
	// 0.01
	Price *float64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// Unit: USD per hour
	//
	// example:
	//
	// The service duration.
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The error message that is returned when a price query fails.
	//
	// example:
	//
	// ecs.e3.large
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Product code
	//
	// example:
	//
	// ecs
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// The instance type. This parameter indicates the information about the instance type. For example, 192.168.0.0/16 may be returned for a Virtual Private Cloud (VPC) instance, ecs.g5.large may be returned for an Elastic Compute Service (ECS) instance, and slb.s1.small may be returned for a Server Load Balancer (SLB) instance. If the resource does not have a specific type, an empty value is returned.
	//
	// example:
	//
	// The billing method.
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// The creation mode. Valid values:\\
	//
	// 1: creates a new instance.\\
	//
	// 2: imports an instance.
	//
	// example:
	//
	// 1
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
	//
	// example:
	//
	// PayAsYouGo
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The resource tag.
	//
	// example:
	//
	// Create
	Lifecycle *string `json:"Lifecycle,omitempty" xml:"Lifecycle,omitempty"`
	NodeLabel *string `json:"NodeLabel,omitempty" xml:"NodeLabel,omitempty"`
	// The deployment result.
	//
	// example:
	//
	// {"hostName":"iZ2zehnzxqixu1pywsfbx1Z","memory":32768.0,"creationTime":"2021-09-28T11:23:46Z","instanceName":"ecs","internetMaxBandwidthOut":0.0,"description":"","clusterId":"","private_ip":"192.168.0.247","instanceId":"i-2zehnzxqixu1pywsfbx1","requestId":"F1C64344-3723-51A0-855B-5F08B5634323","zoneId":"cn-beijing-b","ioOptimized":"optimized","id":"i-2zehnzxqixu1pywsfbx1","instanceNetworkType":"vpc","instanceChargeType":"PostPaid","imageId":"centos_8_4_x64_20G_alibase_20210824.vhd","serialNumber":"cee246c4-38f3-4bf3-950b-c17e88ff6527","vlanId":"","instanceType":"ecs.e3.large","cpu":4.0,"creditSpecification":"","internetMaxBandwidthIn":-1.0,"expiredTime":"2099-12-31T15:59Z","internetChargeType":"PayByTraffic","regionId":"cn-beijing","refId":"79224644_0","stoppedMode":"Not-applicable","status":"Running"}
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The service code.
	//
	// example:
	//
	// ecs
	ResourceCode *string `json:"ResourceCode,omitempty" xml:"ResourceCode,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-2zehnzxqixu1pywsfbx1
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// ecs
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// ecs
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The resource deployment result.
	//
	// example:
	//
	// Finish
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

func (s *GetApplicationResponseBodyDataResourceList) SetNodeLabel(v string) *GetApplicationResponseBodyDataResourceList {
	s.NodeLabel = &v
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

type GetApplicationVariablesRequest struct {
	// example:
	//
	// Q2P4O9YSOKCQ35L9
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetApplicationVariablesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationVariablesRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationVariablesRequest) SetAppId(v string) *GetApplicationVariablesRequest {
	s.AppId = &v
	return s
}

type GetApplicationVariablesResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetApplicationVariablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7036DDBE-0ABA-52D7-A39D-75E511970F07
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationVariablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationVariablesResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationVariablesResponseBody) SetCode(v int32) *GetApplicationVariablesResponseBody {
	s.Code = &v
	return s
}

func (s *GetApplicationVariablesResponseBody) SetData(v []*GetApplicationVariablesResponseBodyData) *GetApplicationVariablesResponseBody {
	s.Data = v
	return s
}

func (s *GetApplicationVariablesResponseBody) SetMessage(v string) *GetApplicationVariablesResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplicationVariablesResponseBody) SetRequestId(v string) *GetApplicationVariablesResponseBody {
	s.RequestId = &v
	return s
}

type GetApplicationVariablesResponseBodyData struct {
	// example:
	//
	// ob5epf79uv****
	InstanceId   *string                                                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	VariableList []*GetApplicationVariablesResponseBodyDataVariableList `json:"VariableList,omitempty" xml:"VariableList,omitempty" type:"Repeated"`
}

func (s GetApplicationVariablesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationVariablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApplicationVariablesResponseBodyData) SetInstanceId(v string) *GetApplicationVariablesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationVariablesResponseBodyData) SetVariableList(v []*GetApplicationVariablesResponseBodyDataVariableList) *GetApplicationVariablesResponseBodyData {
	s.VariableList = v
	return s
}

type GetApplicationVariablesResponseBodyDataVariableList struct {
	// example:
	//
	// ecs.c6.4xlarge
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// ${instance_type}
	Variable *string `json:"Variable,omitempty" xml:"Variable,omitempty"`
}

func (s GetApplicationVariablesResponseBodyDataVariableList) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationVariablesResponseBodyDataVariableList) GoString() string {
	return s.String()
}

func (s *GetApplicationVariablesResponseBodyDataVariableList) SetValue(v string) *GetApplicationVariablesResponseBodyDataVariableList {
	s.Value = &v
	return s
}

func (s *GetApplicationVariablesResponseBodyDataVariableList) SetVariable(v string) *GetApplicationVariablesResponseBodyDataVariableList {
	s.Variable = &v
	return s
}

type GetApplicationVariablesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationVariablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationVariablesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationVariablesResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationVariablesResponse) SetHeaders(v map[string]*string) *GetApplicationVariablesResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationVariablesResponse) SetStatusCode(v int32) *GetApplicationVariablesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationVariablesResponse) SetBody(v *GetApplicationVariablesResponseBody) *GetApplicationVariablesResponse {
	s.Body = v
	return s
}

type GetApplicationVariables4FailRequest struct {
	// example:
	//
	// Q2P4O9YSOKCT35L9
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s GetApplicationVariables4FailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationVariables4FailRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationVariables4FailRequest) SetAppId(v string) *GetApplicationVariables4FailRequest {
	s.AppId = &v
	return s
}

type GetApplicationVariables4FailResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetApplicationVariables4FailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BFB7F5C8-FE7A-06CA-9F87-ABBF6B848F0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationVariables4FailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationVariables4FailResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationVariables4FailResponseBody) SetCode(v int32) *GetApplicationVariables4FailResponseBody {
	s.Code = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBody) SetData(v []*GetApplicationVariables4FailResponseBodyData) *GetApplicationVariables4FailResponseBody {
	s.Data = v
	return s
}

func (s *GetApplicationVariables4FailResponseBody) SetMessage(v string) *GetApplicationVariables4FailResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBody) SetRequestId(v string) *GetApplicationVariables4FailResponseBody {
	s.RequestId = &v
	return s
}

type GetApplicationVariables4FailResponseBodyData struct {
	// example:
	//
	// instance_name
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// example:
	//
	// cadt-app-01
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// ${name}
	PlaceHolder *string `json:"PlaceHolder,omitempty" xml:"PlaceHolder,omitempty"`
	// example:
	//
	// cn-shenzhen
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// cadt-app-01
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// ${name}
	Variable *string `json:"Variable,omitempty" xml:"Variable,omitempty"`
}

func (s GetApplicationVariables4FailResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationVariables4FailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApplicationVariables4FailResponseBodyData) SetAttribute(v string) *GetApplicationVariables4FailResponseBodyData {
	s.Attribute = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBodyData) SetDefaultValue(v string) *GetApplicationVariables4FailResponseBodyData {
	s.DefaultValue = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBodyData) SetPlaceHolder(v string) *GetApplicationVariables4FailResponseBodyData {
	s.PlaceHolder = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBodyData) SetRegion(v string) *GetApplicationVariables4FailResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBodyData) SetValue(v string) *GetApplicationVariables4FailResponseBodyData {
	s.Value = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBodyData) SetVariable(v string) *GetApplicationVariables4FailResponseBodyData {
	s.Variable = &v
	return s
}

type GetApplicationVariables4FailResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApplicationVariables4FailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApplicationVariables4FailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationVariables4FailResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationVariables4FailResponse) SetHeaders(v map[string]*string) *GetApplicationVariables4FailResponse {
	s.Headers = v
	return s
}

func (s *GetApplicationVariables4FailResponse) SetStatusCode(v int32) *GetApplicationVariables4FailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApplicationVariables4FailResponse) SetBody(v *GetApplicationVariables4FailResponseBody) *GetApplicationVariables4FailResponse {
	s.Body = v
	return s
}

type GetExecuteOperationResultRequest struct {
	// The ID of the operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// op_xxxxxxxxxxxxxxxxxxxxxx
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the resource group. This parameter is specified to verify the permissions on the resource group.
	//
	// example:
	//
	// testResourceId
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed result of the queried operation.
	Data *GetExecuteOperationResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// " "
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
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
	//
	// example:
	//
	// " "
	Arguments *string `json:"Arguments,omitempty" xml:"Arguments,omitempty"`
	// The returned message.
	//
	// example:
	//
	// “ ”
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the operation.
	//
	// example:
	//
	// op_xxxxxxxx
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The status of the operation.
	//
	// example:
	//
	// success
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
	// The disaster recovery switchover task ID.
	//
	// example:
	//
	// 2615
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
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status of the switchover task.
	//
	// example:
	//
	// Running
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message. If the request was successful, a success message is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// OKITHEVRQCN6ULQG
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 923692F0-A15B-58B4-BAF4-2AFA4AF46240
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
	// Specifies whether the value of this parameter is the ID of a disaster recovery set.
	//
	// example:
	//
	// true
	IsPlanId *bool `json:"IsPlanId,omitempty" xml:"IsPlanId,omitempty"`
	// If you set IsPlanId to false, specify the ID of a disaster recovery application. If you set IsPlanId to true, specify the ID of a disaster recovery set.
	//
	// example:
	//
	// FS3ATPTOSC4SE1GG
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
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The zones where the current disaster recovery service can be switched.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// The specified ResourceIds are not found in our records.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BFB7F5C8-FE7A-06CA-9F87-ABBF6B848F0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type GetResource4ModifyRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// BE68D71ZY5YYIU9R
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 1
	NextToken *int64 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s GetResource4ModifyRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResource4ModifyRecordRequest) GoString() string {
	return s.String()
}

func (s *GetResource4ModifyRecordRequest) SetApplicationId(v string) *GetResource4ModifyRecordRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetResource4ModifyRecordRequest) SetMaxResults(v int64) *GetResource4ModifyRecordRequest {
	s.MaxResults = &v
	return s
}

func (s *GetResource4ModifyRecordRequest) SetNextToken(v int64) *GetResource4ModifyRecordRequest {
	s.NextToken = &v
	return s
}

type GetResource4ModifyRecordResponseBody struct {
	// example:
	//
	// {
	//
	//     "PolicyType": "",
	//
	//     "AuthPrincipalOwnerId": "",
	//
	//     "EncodedDiagnosticMessage": "",
	//
	//     "AuthPrincipalType": "",
	//
	//     "AuthPrincipalDisplayName": "",
	//
	//     "NoPermissionType": "",
	//
	//     "AuthAction": ""
	//
	//   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetResource4ModifyRecordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetResource4ModifyRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResource4ModifyRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetResource4ModifyRecordResponseBody) SetAccessDeniedDetail(v string) *GetResource4ModifyRecordResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBody) SetCode(v string) *GetResource4ModifyRecordResponseBody {
	s.Code = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBody) SetData(v []*GetResource4ModifyRecordResponseBodyData) *GetResource4ModifyRecordResponseBody {
	s.Data = v
	return s
}

func (s *GetResource4ModifyRecordResponseBody) SetMessage(v string) *GetResource4ModifyRecordResponseBody {
	s.Message = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBody) SetNextToken(v string) *GetResource4ModifyRecordResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBody) SetRequestId(v string) *GetResource4ModifyRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBody) SetTotalCount(v string) *GetResource4ModifyRecordResponseBody {
	s.TotalCount = &v
	return s
}

type GetResource4ModifyRecordResponseBodyData struct {
	// example:
	//
	// {\\"InstanceId\\": \\"\\", \\"AttributeName\\": \\"drmCommand\\", \\"Id\\": 16800, \\"Desc\\": \\"test\\"}
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// example:
	//
	// 1726645341000
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// rm-uf6308dyal1*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// Finish
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// rds
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetResource4ModifyRecordResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResource4ModifyRecordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResource4ModifyRecordResponseBodyData) SetAttribute(v string) *GetResource4ModifyRecordResponseBodyData {
	s.Attribute = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBodyData) SetModifyTime(v string) *GetResource4ModifyRecordResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBodyData) SetResourceId(v string) *GetResource4ModifyRecordResponseBodyData {
	s.ResourceId = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBodyData) SetStatus(v string) *GetResource4ModifyRecordResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetResource4ModifyRecordResponseBodyData) SetType(v string) *GetResource4ModifyRecordResponseBodyData {
	s.Type = &v
	return s
}

type GetResource4ModifyRecordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResource4ModifyRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResource4ModifyRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResource4ModifyRecordResponse) GoString() string {
	return s.String()
}

func (s *GetResource4ModifyRecordResponse) SetHeaders(v map[string]*string) *GetResource4ModifyRecordResponse {
	s.Headers = v
	return s
}

func (s *GetResource4ModifyRecordResponse) SetStatusCode(v int32) *GetResource4ModifyRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResource4ModifyRecordResponse) SetBody(v *GetResource4ModifyRecordResponseBody) *GetResource4ModifyRecordResponse {
	s.Body = v
	return s
}

type GetResult4QueryInstancePrice4ModifyRequest struct {
	// example:
	//
	// 02S7UU41WKJL7ERR
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 58d5d8c5b5489150417a7cd6caa614bb
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetResult4QueryInstancePrice4ModifyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetResult4QueryInstancePrice4ModifyRequest) GoString() string {
	return s.String()
}

func (s *GetResult4QueryInstancePrice4ModifyRequest) SetApplicationId(v string) *GetResult4QueryInstancePrice4ModifyRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyRequest) SetTaskId(v string) *GetResult4QueryInstancePrice4ModifyRequest {
	s.TaskId = &v
	return s
}

type GetResult4QueryInstancePrice4ModifyResponseBody struct {
	// example:
	//
	// {
	//
	//     "PolicyType": "",
	//
	//     "AuthPrincipalOwnerId": "",
	//
	//     "EncodedDiagnosticMessage": "",
	//
	//     "AuthPrincipalType": "",
	//
	//     "AuthPrincipalDisplayName": "",
	//
	//     "NoPermissionType": "",
	//
	//     "AuthAction": ""
	//
	//   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string                                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetResult4QueryInstancePrice4ModifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResult4QueryInstancePrice4ModifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetResult4QueryInstancePrice4ModifyResponseBody) GoString() string {
	return s.String()
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) SetAccessDeniedDetail(v string) *GetResult4QueryInstancePrice4ModifyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) SetCode(v string) *GetResult4QueryInstancePrice4ModifyResponseBody {
	s.Code = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) SetData(v *GetResult4QueryInstancePrice4ModifyResponseBodyData) *GetResult4QueryInstancePrice4ModifyResponseBody {
	s.Data = v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) SetMessage(v string) *GetResult4QueryInstancePrice4ModifyResponseBody {
	s.Message = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBody) SetRequestId(v string) *GetResult4QueryInstancePrice4ModifyResponseBody {
	s.RequestId = &v
	return s
}

type GetResult4QueryInstancePrice4ModifyResponseBodyData struct {
	PriceList []*GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList `json:"PriceList,omitempty" xml:"PriceList,omitempty" type:"Repeated"`
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// d9a3e99b-6954-4a16-ad51-954db4a528b7
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetResult4QueryInstancePrice4ModifyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetResult4QueryInstancePrice4ModifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyData) SetPriceList(v []*GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) *GetResult4QueryInstancePrice4ModifyResponseBodyData {
	s.PriceList = v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyData) SetStatus(v string) *GetResult4QueryInstancePrice4ModifyResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyData) SetTaskId(v string) *GetResult4QueryInstancePrice4ModifyResponseBodyData {
	s.TaskId = &v
	return s
}

type GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList struct {
	// example:
	//
	// 2
	DiscountAmount *float64 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	Error          *string  `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// vpc
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// example:
	//
	// 3
	OriginalAmount *float64 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	PriceUnit      *string  `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	PromotionName  *string  `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// 1
	TradeAmount *float64 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) String() string {
	return tea.Prettify(s)
}

func (s GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) GoString() string {
	return s.String()
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetDiscountAmount(v float64) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.DiscountAmount = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetError(v string) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.Error = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetNodeType(v string) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.NodeType = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetOriginalAmount(v float64) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.OriginalAmount = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetPriceUnit(v string) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.PriceUnit = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetPromotionName(v string) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.PromotionName = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList) SetTradeAmount(v float64) *GetResult4QueryInstancePrice4ModifyResponseBodyDataPriceList {
	s.TradeAmount = &v
	return s
}

type GetResult4QueryInstancePrice4ModifyResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResult4QueryInstancePrice4ModifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResult4QueryInstancePrice4ModifyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetResult4QueryInstancePrice4ModifyResponse) GoString() string {
	return s.String()
}

func (s *GetResult4QueryInstancePrice4ModifyResponse) SetHeaders(v map[string]*string) *GetResult4QueryInstancePrice4ModifyResponse {
	s.Headers = v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponse) SetStatusCode(v int32) *GetResult4QueryInstancePrice4ModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResult4QueryInstancePrice4ModifyResponse) SetBody(v *GetResult4QueryInstancePrice4ModifyResponseBody) *GetResult4QueryInstancePrice4ModifyResponse {
	s.Body = v
	return s
}

type GetTemplateRequest struct {
	// Template Area
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// ResourceGroup ID
	//
	// example:
	//
	// rg-aekzhfgmw4e6fwq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Template ID
	//
	// This parameter is required.
	//
	// example:
	//
	// XFKR6WYRVS24S07R
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the template.
	Data *GetTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The interface returns information
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
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
	//
	// example:
	//
	// 2020-09-22 17:08:31
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Template DescriptionD
	//
	// example:
	//
	// 无
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The path to the template schema image file
	//
	// example:
	//
	// bp-studio-template/sr-U37UD2YQCRJ75X5V.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// The name of the template
	//
	// example:
	//
	// cadt-template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzhfgmw4e6fwq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Template ID
	//
	// example:
	//
	// XFKR6WYRVS24S07R
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
	//
	// example:
	//
	// instance_name
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The type of the variable.
	//
	// example:
	//
	// String
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The default value of the variable.
	//
	// example:
	//
	// cadt-app-01
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The value of the variable.
	//
	// example:
	//
	// ${name}
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
	//
	// example:
	//
	// rg-aekzhfgmw4e6fwq
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the token.
	Data *GetTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
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
	//
	// example:
	//
	// STS.NTm*****8tu
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret used to access OSS.
	//
	// example:
	//
	// 9NG*****K4X
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The OSS bucket that is used to store the architecture image.
	//
	// example:
	//
	// bucket-1
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS endpoint.
	//
	// example:
	//
	// https://oss-cn-beijing.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The token that is used to access the Object Storage Service (OSS) bucket that stores the architecture image.
	//
	// example:
	//
	// ABCD
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The OSS bucket that is used to save data snapshots.
	//
	// example:
	//
	// bucket-2
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
	// The application ID.
	//
	// example:
	//
	// 002XWH7MXB8MJRU0
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
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The switchover task ID.
	//
	// example:
	//
	// 7250
	Data *int32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message. If the request was successful, a success message is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// Unsupported Operation PrepareEvent->FailOverPrepareSuccess FoApp_DDLJK2WM8ETU9JAC
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
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
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The pagination size of the resulting value cannot be less than the minimum value of 1 and cannot be greater than the maximum value of 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination page number of the resulting value cannot be less than the minimum value of 1 and cannot be greater than the maximum value of 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// 1 update time,<br>2 creation time
	//
	// example:
	//
	// 1
	OrderType *int64 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Resource Id
	//
	// example:
	//
	// vsw-xxxxxxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The status of the applications to be returned.
	//
	// example:
	//
	// The following values are "success" and "release".
	//
	// If the input value is "success", the returned application list includes all applications in the Deployed_Success state of successful deployment.
	//
	// If the input value is release, the returned application list includes all applications in the release success (Destroyed_Success) and release failure (Destroyed_Failure) status.
	//
	// If this parameter is left blank, the returned app list includes apps in all states.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Template Id
	//
	// example:
	//
	// 0KSXXX6SJU03TXXX
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// App listing information
	Data []*ListApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The interface returns information
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The query token returned in this call.
	//
	// example:
	//
	// 2
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// BFB7F5C8-FE7A-06CA-9F87-ABBF6B848F0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 123
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
	//
	// example:
	//
	// JIX9NEZUALGS46UI
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The time when the application was created.
	//
	// example:
	//
	// 2021-09-15  08:30:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The URL of the application architecture image.
	//
	// example:
	//
	// https://bp-studio-daily.oss-cn-beijing.aliyuncs.com/1411182597819805/sr-Y3KR7ZSQZR2F0YX3.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// cadt-appName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group to which the application belongs.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the application.
	//
	// example:
	//
	// Deployed_Success
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
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about disaster recovery plans.
	Data []*ListFoCreatedAppsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message. If the request was successful, a success message is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// Cannot find region according to your domain.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40F63F07-3AB6-53B3-8825-0580C130E3EA
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
	// 应用ID
	//
	// example:
	//
	// JIX9NEZUALGS46UI
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The URL of an error report.
	//
	// example:
	//
	// https://api.aliyun.com/troubleshoot?q=ServiceUnavailable&product=BPStudio&requestId=4CDA03A3-C652-1408-8ABD-7E652A7CBFB6
	ReportUrl *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
	// The status of the disaster recovery plan.
	//
	// example:
	//
	// Deployed_Success
	//
	// Destroyed_Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The title.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify a maximum number of 50 IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// Application
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
	//
	// example:
	//
	// Label 1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// Label value 1
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
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// bp.java.nopowerContact
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned value of NextToken is a pagination token, which can be used in the next request to retrieve a new page of results. If the NextToken parameter is empty, no next page exists.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 847C9D0A-BABD-589C-8A9C-6464409EDED9
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
	//
	// example:
	//
	// MW7T1820ACJX9BI7
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type. Valid values: application and template.
	//
	// example:
	//
	// application
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// Key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// Value
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
	//
	// example:
	//
	// cadt
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The criterion by which the returned templates are sorted. Valid values:
	//
	// 	- 1: The templates are sorted by the time when they are updated.
	//
	// 	- 2: The templates are sorted by the time when they are created.
	//
	// 	- 3: The templates are sorted by the system.
	//
	// 	- 4: The templates are sorted by the number of times that they are used.
	//
	// 	- If you specify an integer other than 1, 2, 3, and 4 or do not specify any value, the templates are sorted by the system.
	//
	// example:
	//
	// 1
	OrderType *int64 `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyjt3c5om3hi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tag that you want to use to query templates.
	//
	// example:
	//
	// 1
	TagList *int32 `json:"TagList,omitempty" xml:"TagList,omitempty"`
	// The type of the templates to be returned. Valid values: public and private
	//
	// This parameter is required.
	//
	// example:
	//
	// public
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
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details about templates.
	Data []*ListTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 100
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
	//
	// example:
	//
	// 2021-03-18 16:41:31
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The URL of the architecture image.
	//
	// example:
	//
	// bp-studio-template/sr-U37UD2YQCRJ75X5V.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// cadt-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyjt3c5om3hi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the tag that is added to the template.
	//
	// example:
	//
	// 1
	TagId *int32 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The name of the tag that is added to the template.
	//
	// example:
	//
	// Official template
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// CJQ6H0XUEQ20IYJQ
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

type ModifyApplicationSpecRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 02S7UU41WKJL7ERR
	ApplicationId *string                                     `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	InstanceSpec  []*ModifyApplicationSpecRequestInstanceSpec `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" type:"Repeated"`
}

func (s ModifyApplicationSpecRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyApplicationSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationSpecRequest) SetApplicationId(v string) *ModifyApplicationSpecRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationSpecRequest) SetInstanceSpec(v []*ModifyApplicationSpecRequestInstanceSpec) *ModifyApplicationSpecRequest {
	s.InstanceSpec = v
	return s
}

type ModifyApplicationSpecRequestInstanceSpec struct {
	Configuration map[string]interface{} `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// example:
	//
	// rm-2ze8f4ah378a*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyApplicationSpecRequestInstanceSpec) String() string {
	return tea.Prettify(s)
}

func (s ModifyApplicationSpecRequestInstanceSpec) GoString() string {
	return s.String()
}

func (s *ModifyApplicationSpecRequestInstanceSpec) SetConfiguration(v map[string]interface{}) *ModifyApplicationSpecRequestInstanceSpec {
	s.Configuration = v
	return s
}

func (s *ModifyApplicationSpecRequestInstanceSpec) SetInstanceId(v string) *ModifyApplicationSpecRequestInstanceSpec {
	s.InstanceId = &v
	return s
}

type ModifyApplicationSpecShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 02S7UU41WKJL7ERR
	ApplicationId      *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	InstanceSpecShrink *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
}

func (s ModifyApplicationSpecShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyApplicationSpecShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationSpecShrinkRequest) SetApplicationId(v string) *ModifyApplicationSpecShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationSpecShrinkRequest) SetInstanceSpecShrink(v string) *ModifyApplicationSpecShrinkRequest {
	s.InstanceSpecShrink = &v
	return s
}

type ModifyApplicationSpecResponseBody struct {
	// example:
	//
	// {
	//
	//     "PolicyType": "",
	//
	//     "AuthPrincipalOwnerId": "",
	//
	//     "EncodedDiagnosticMessage": "",
	//
	//     "AuthPrincipalType": "",
	//
	//     "AuthPrincipalDisplayName": "",
	//
	//     "NoPermissionType": "",
	//
	//     "AuthAction": ""
	//
	//   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// -1
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A3488F1D-xxxx-xxxx-xxxx-5374BA0F3562
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyApplicationSpecResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyApplicationSpecResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApplicationSpecResponseBody) SetAccessDeniedDetail(v string) *ModifyApplicationSpecResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ModifyApplicationSpecResponseBody) SetCode(v string) *ModifyApplicationSpecResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyApplicationSpecResponseBody) SetData(v string) *ModifyApplicationSpecResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyApplicationSpecResponseBody) SetMessage(v string) *ModifyApplicationSpecResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyApplicationSpecResponseBody) SetRequestId(v string) *ModifyApplicationSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApplicationSpecResponseBody) SetSuccess(v string) *ModifyApplicationSpecResponseBody {
	s.Success = &v
	return s
}

type ModifyApplicationSpecResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApplicationSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApplicationSpecResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyApplicationSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyApplicationSpecResponse) SetHeaders(v map[string]*string) *ModifyApplicationSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyApplicationSpecResponse) SetStatusCode(v int32) *ModifyApplicationSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApplicationSpecResponse) SetBody(v *ModifyApplicationSpecResponseBody) *ModifyApplicationSpecResponse {
	s.Body = v
	return s
}

type QueryInstancePrice4ModifyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 002XWH7MXB8MJRU0
	ApplicationId *string                `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Configuration map[string]interface{} `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rm-uf66k9143r2ch*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryInstancePrice4ModifyRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePrice4ModifyRequest) GoString() string {
	return s.String()
}

func (s *QueryInstancePrice4ModifyRequest) SetApplicationId(v string) *QueryInstancePrice4ModifyRequest {
	s.ApplicationId = &v
	return s
}

func (s *QueryInstancePrice4ModifyRequest) SetConfiguration(v map[string]interface{}) *QueryInstancePrice4ModifyRequest {
	s.Configuration = v
	return s
}

func (s *QueryInstancePrice4ModifyRequest) SetInstanceId(v string) *QueryInstancePrice4ModifyRequest {
	s.InstanceId = &v
	return s
}

type QueryInstancePrice4ModifyShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 002XWH7MXB8MJRU0
	ApplicationId       *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ConfigurationShrink *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rm-uf66k9143r2ch*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryInstancePrice4ModifyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePrice4ModifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryInstancePrice4ModifyShrinkRequest) SetApplicationId(v string) *QueryInstancePrice4ModifyShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *QueryInstancePrice4ModifyShrinkRequest) SetConfigurationShrink(v string) *QueryInstancePrice4ModifyShrinkRequest {
	s.ConfigurationShrink = &v
	return s
}

func (s *QueryInstancePrice4ModifyShrinkRequest) SetInstanceId(v string) *QueryInstancePrice4ModifyShrinkRequest {
	s.InstanceId = &v
	return s
}

type QueryInstancePrice4ModifyResponseBody struct {
	// example:
	//
	// {
	//
	//     "PolicyType": "",
	//
	//     "AuthPrincipalOwnerId": "",
	//
	//     "EncodedDiagnosticMessage": "",
	//
	//     "AuthPrincipalType": "",
	//
	//     "AuthPrincipalDisplayName": "",
	//
	//     "NoPermissionType": "",
	//
	//     "AuthAction": ""
	//
	//   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// taskId
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9656C816-1E9A-58D2-86D5-710678D61AF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInstancePrice4ModifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePrice4ModifyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstancePrice4ModifyResponseBody) SetAccessDeniedDetail(v string) *QueryInstancePrice4ModifyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryInstancePrice4ModifyResponseBody) SetCode(v string) *QueryInstancePrice4ModifyResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInstancePrice4ModifyResponseBody) SetData(v string) *QueryInstancePrice4ModifyResponseBody {
	s.Data = &v
	return s
}

func (s *QueryInstancePrice4ModifyResponseBody) SetMessage(v string) *QueryInstancePrice4ModifyResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstancePrice4ModifyResponseBody) SetRequestId(v string) *QueryInstancePrice4ModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstancePrice4ModifyResponseBody) SetSuccess(v string) *QueryInstancePrice4ModifyResponseBody {
	s.Success = &v
	return s
}

type QueryInstancePrice4ModifyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstancePrice4ModifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstancePrice4ModifyResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInstancePrice4ModifyResponse) GoString() string {
	return s.String()
}

func (s *QueryInstancePrice4ModifyResponse) SetHeaders(v map[string]*string) *QueryInstancePrice4ModifyResponse {
	s.Headers = v
	return s
}

func (s *QueryInstancePrice4ModifyResponse) SetStatusCode(v int32) *QueryInstancePrice4ModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstancePrice4ModifyResponse) SetBody(v *QueryInstancePrice4ModifyResponseBody) *QueryInstancePrice4ModifyResponse {
	s.Body = v
	return s
}

type QueryInstanceSpec4ModifyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// BE68D71ZY5YYIU9R
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rm-uf66k9143r2ch*****
	InstanceId *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MethodName *string                `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s QueryInstanceSpec4ModifyRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceSpec4ModifyRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceSpec4ModifyRequest) SetApplicationId(v string) *QueryInstanceSpec4ModifyRequest {
	s.ApplicationId = &v
	return s
}

func (s *QueryInstanceSpec4ModifyRequest) SetInstanceId(v string) *QueryInstanceSpec4ModifyRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryInstanceSpec4ModifyRequest) SetMethodName(v string) *QueryInstanceSpec4ModifyRequest {
	s.MethodName = &v
	return s
}

func (s *QueryInstanceSpec4ModifyRequest) SetParameters(v map[string]interface{}) *QueryInstanceSpec4ModifyRequest {
	s.Parameters = v
	return s
}

type QueryInstanceSpec4ModifyShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// BE68D71ZY5YYIU9R
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rm-uf66k9143r2ch*****
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MethodName       *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s QueryInstanceSpec4ModifyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceSpec4ModifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceSpec4ModifyShrinkRequest) SetApplicationId(v string) *QueryInstanceSpec4ModifyShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *QueryInstanceSpec4ModifyShrinkRequest) SetInstanceId(v string) *QueryInstanceSpec4ModifyShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryInstanceSpec4ModifyShrinkRequest) SetMethodName(v string) *QueryInstanceSpec4ModifyShrinkRequest {
	s.MethodName = &v
	return s
}

func (s *QueryInstanceSpec4ModifyShrinkRequest) SetParametersShrink(v string) *QueryInstanceSpec4ModifyShrinkRequest {
	s.ParametersShrink = &v
	return s
}

type QueryInstanceSpec4ModifyResponseBody struct {
	// example:
	//
	// {
	//
	//     "PolicyType": "",
	//
	//     "AuthPrincipalOwnerId": "",
	//
	//     "EncodedDiagnosticMessage": "",
	//
	//     "AuthPrincipalType": "",
	//
	//     "AuthPrincipalDisplayName": "",
	//
	//     "NoPermissionType": "",
	//
	//     "AuthAction": ""
	//
	//   }
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QueryInstanceSpec4ModifyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 847C9D0A-BABD-589C-8A9C-6464409EDED9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInstanceSpec4ModifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceSpec4ModifyResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInstanceSpec4ModifyResponseBody) SetAccessDeniedDetail(v string) *QueryInstanceSpec4ModifyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBody) SetCode(v string) *QueryInstanceSpec4ModifyResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBody) SetData(v *QueryInstanceSpec4ModifyResponseBodyData) *QueryInstanceSpec4ModifyResponseBody {
	s.Data = v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBody) SetMessage(v string) *QueryInstanceSpec4ModifyResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBody) SetRequestId(v string) *QueryInstanceSpec4ModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBody) SetSuccess(v bool) *QueryInstanceSpec4ModifyResponseBody {
	s.Success = &v
	return s
}

type QueryInstanceSpec4ModifyResponseBodyData struct {
	OptionalValues []*QueryInstanceSpec4ModifyResponseBodyDataOptionalValues `json:"OptionalValues,omitempty" xml:"OptionalValues,omitempty" type:"Repeated"`
}

func (s QueryInstanceSpec4ModifyResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceSpec4ModifyResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInstanceSpec4ModifyResponseBodyData) SetOptionalValues(v []*QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) *QueryInstanceSpec4ModifyResponseBodyData {
	s.OptionalValues = v
	return s
}

type QueryInstanceSpec4ModifyResponseBodyDataOptionalValues struct {
	Label *string  `json:"Label,omitempty" xml:"Label,omitempty"`
	Max   *float64 `json:"Max,omitempty" xml:"Max,omitempty"`
	Min   *float64 `json:"Min,omitempty" xml:"Min,omitempty"`
	Step  *float64 `json:"Step,omitempty" xml:"Step,omitempty"`
	Value *string  `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) GoString() string {
	return s.String()
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) SetLabel(v string) *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues {
	s.Label = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) SetMax(v float64) *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues {
	s.Max = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) SetMin(v float64) *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues {
	s.Min = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) SetStep(v float64) *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues {
	s.Step = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues) SetValue(v string) *QueryInstanceSpec4ModifyResponseBodyDataOptionalValues {
	s.Value = &v
	return s
}

type QueryInstanceSpec4ModifyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstanceSpec4ModifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstanceSpec4ModifyResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryInstanceSpec4ModifyResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceSpec4ModifyResponse) SetHeaders(v map[string]*string) *QueryInstanceSpec4ModifyResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceSpec4ModifyResponse) SetStatusCode(v int32) *QueryInstanceSpec4ModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponse) SetBody(v *QueryInstanceSpec4ModifyResponseBody) *QueryInstanceSpec4ModifyResponse {
	s.Body = v
	return s
}

type ReConfigApplicationRequest struct {
	// example:
	//
	// Q2P4O9YSOKCT35L9
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// {"${instance_type}":"ecs.c6.3xlarge"}
	Variables *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s ReConfigApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ReConfigApplicationRequest) GoString() string {
	return s.String()
}

func (s *ReConfigApplicationRequest) SetAppId(v string) *ReConfigApplicationRequest {
	s.AppId = &v
	return s
}

func (s *ReConfigApplicationRequest) SetVariables(v string) *ReConfigApplicationRequest {
	s.Variables = &v
	return s
}

type ReConfigApplicationResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9656C816-1E9A-58D2-86D5-710678D61AF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReConfigApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReConfigApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ReConfigApplicationResponseBody) SetCode(v int32) *ReConfigApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *ReConfigApplicationResponseBody) SetMessage(v string) *ReConfigApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *ReConfigApplicationResponseBody) SetRequestId(v string) *ReConfigApplicationResponseBody {
	s.RequestId = &v
	return s
}

type ReConfigApplicationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReConfigApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReConfigApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ReConfigApplicationResponse) GoString() string {
	return s.String()
}

func (s *ReConfigApplicationResponse) SetHeaders(v map[string]*string) *ReConfigApplicationResponse {
	s.Headers = v
	return s
}

func (s *ReConfigApplicationResponse) SetStatusCode(v int32) *ReConfigApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ReConfigApplicationResponse) SetBody(v *ReConfigApplicationResponseBody) *ReConfigApplicationResponse {
	s.Body = v
	return s
}

type ReleaseApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7QSXFQW46ZNGOUDM
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// default
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

func (s *ReleaseApplicationRequest) SetClientToken(v string) *ReleaseApplicationRequest {
	s.ClientToken = &v
	return s
}

func (s *ReleaseApplicationRequest) SetResourceGroupId(v string) *ReleaseApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

type ReleaseApplicationResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return value.
	//
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// The resource does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BFB7F5C8-FE7A-06CA-9F87-ABBF6B848F0C
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
	//
	// This parameter is required.
	//
	// example:
	//
	// 02S7UU41WKJL7ERR
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyjt3c5om3hi
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

func (s *ValidateApplicationRequest) SetClientToken(v string) *ValidateApplicationRequest {
	s.ClientToken = &v
	return s
}

func (s *ValidateApplicationRequest) SetResourceGroupId(v string) *ValidateApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

type ValidateApplicationResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data of the application.
	//
	// example:
	//
	// 123
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
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
	// The operation that you want to perform. Set the value to ValuateApplication.
	//
	// This parameter is required.
	//
	// example:
	//
	// 02S7UU41WKJL7ERR
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the resource group to which the application you want to query belongs.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// rg-acfmyjt3c5om3hi
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

func (s *ValuateApplicationRequest) SetClientToken(v string) *ValuateApplicationRequest {
	s.ClientToken = &v
	return s
}

func (s *ValuateApplicationRequest) SetResourceGroupId(v string) *ValuateApplicationRequest {
	s.ResourceGroupId = &v
	return s
}

type ValuateApplicationResponseBody struct {
	// The code of the query task.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 123
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// Idempotent notation
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned message.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
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
	//
	// example:
	//
	// cn-hangzhou
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instances to be replaced.
	Instances []*ValuateTemplateRequestInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The ID of the resource group to which the application belongs.
	//
	// example:
	//
	// rg-acfmyjt3c5om3fi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0KSHPM6SJU03TNZP
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
	//
	// example:
	//
	// vpc-bp1q56trhtaq40vlq5oj
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the application instance that is displayed on the diagram.
	//
	// example:
	//
	// ecs
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs
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
	//
	// example:
	//
	// cn-hangzhou
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instances to be replaced.
	InstancesShrink *string `json:"Instances,omitempty" xml:"Instances,omitempty"`
	// The ID of the resource group to which the application belongs.
	//
	// example:
	//
	// rg-acfmyjt3c5om3fi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0KSHPM6SJU03TNZP
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
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the inquiry.
	Data *ValuateTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 847C9D0A-BABD-589C-8A9C-6464409EDED9
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
	//
	// example:
	//
	// 73
	DiscountAmount *float64 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// InvalidSaleComponentFault : The request not refer to the correct order sale component.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ecs
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The original price.
	//
	// example:
	//
	// 83.0
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The information about the price.
	PriceList []*ValuateTemplateResponseBodyDataResourceListPriceList `json:"PriceList,omitempty" xml:"PriceList,omitempty" type:"Repeated"`
	// The pricing unit.
	//
	// example:
	//
	// USD
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The discount information.
	//
	// example:
	//
	// The discount information.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// The price at which the transaction is made.
	//
	// example:
	//
	// 10.0
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
	//
	// example:
	//
	// 82.99
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// ServiceUnavailable : The request has failed due to a temporary failure of the server.\\r\\nRequestId : 4AA302DB-3286-5589-8637-FF6D8507B7A9.
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The resource type.
	//
	// example:
	//
	// eip
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The original price.
	//
	// example:
	//
	// 83.0
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The pricing unit.
	//
	// example:
	//
	// USD
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// The discount information.
	//
	// example:
	//
	// The discount information.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// The resource ID.
	//
	// example:
	//
	// 1687225092
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The price at which the transaction is made.
	//
	// example:
	//
	// 0.01
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
	// Indicates whether the instance is newly created. Valid values:\\
	//
	// 1: The instance is newly created.\\
	//
	// 2: The instance already exists.\\
	//
	// 0: The price of the instance is not included.
	//
	// example:
	//
	// "1"
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

// Summary:
//
// Switches a disaster recovery application back to the primary zone.
//
// Description:
//
// You can call this operation to switch a disaster recovery application back to the primary zone.
//
// @param request - AppFailBackRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppFailBackResponse
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

// Summary:
//
// Switches a disaster recovery application back to the primary zone.
//
// Description:
//
// You can call this operation to switch a disaster recovery application back to the primary zone.
//
// @param request - AppFailBackRequest
//
// @return AppFailBackResponse
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

// Summary:
//
// Switches a disaster recovery application to another supported zone.
//
// Description:
//
// You can call this operation to switch a disaster recovery application to another supported zone.
//
// @param request - AppFailOverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AppFailOverResponse
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

// Summary:
//
// Switches a disaster recovery application to another supported zone.
//
// Description:
//
// You can call this operation to switch a disaster recovery application to another supported zone.
//
// @param request - AppFailOverRequest
//
// @return AppFailOverResponse
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

// Summary:
//
// Changes the resource group to which an application or template belongs.
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
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

// Summary:
//
// Changes the resource group to which an application or template belongs.
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
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

// Summary:
//
// Creates an application based on an official template or private template in Cloud Architect Design Tool (CADT). Before you call this operation, make sure that you understand the billing methods and prices of Alibaba Cloud services to be used in the application.
//
// @param tmpReq - CreateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApplicationResponse
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

// Summary:
//
// Creates an application based on an official template or private template in Cloud Architect Design Tool (CADT). Before you call this operation, make sure that you understand the billing methods and prices of Alibaba Cloud services to be used in the application.
//
// @param request - CreateApplicationRequest
//
// @return CreateApplicationResponse
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

// Summary:
//
// Deletes an application.
//
// Description:
//
// Before you call this operation to delete an application, make sure that the application is in the `Destroyed_Success` state. Otherwise, the application fails to be deleted.`` You can call the [GetApplication](https://www.alibabacloud.com/help/en/bp-studio/latest/api-bpstudio-2021-09-31-getapplication) operation to query the status of an application.
//
// @param request - DeleteApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApplicationResponse
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

// Summary:
//
// Deletes an application.
//
// Description:
//
// Before you call this operation to delete an application, make sure that the application is in the `Destroyed_Success` state. Otherwise, the application fails to be deleted.`` You can call the [GetApplication](https://www.alibabacloud.com/help/en/bp-studio/latest/api-bpstudio-2021-09-31-getapplication) operation to query the status of an application.
//
// @param request - DeleteApplicationRequest
//
// @return DeleteApplicationResponse
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

// Summary:
//
// Deploys an application after the payment.
//
// @param request - DeployApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeployApplicationResponse
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
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

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

// Summary:
//
// Deploys an application after the payment.
//
// @param request - DeployApplicationRequest
//
// @return DeployApplicationResponse
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

// Summary:
//
// Asynchronous execution of product operation functions.
//
// @param tmpReq - ExecuteOperationASyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteOperationASyncResponse
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

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
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

// Summary:
//
// Asynchronous execution of product operation functions.
//
// @param request - ExecuteOperationASyncRequest
//
// @return ExecuteOperationASyncResponse
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

// Summary:
//
// 维护应用下资源API（同步操作）
//
// @param tmpReq - ExecuteOperationSyncRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteOperationSyncResponse
func (client *Client) ExecuteOperationSyncWithOptions(tmpReq *ExecuteOperationSyncRequest, runtime *util.RuntimeOptions) (_result *ExecuteOperationSyncResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExecuteOperationSyncShrinkRequest{}
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

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
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
		Action:      tea.String("ExecuteOperationSync"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExecuteOperationSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 维护应用下资源API（同步操作）
//
// @param request - ExecuteOperationSyncRequest
//
// @return ExecuteOperationSyncResponse
func (client *Client) ExecuteOperationSync(request *ExecuteOperationSyncRequest) (_result *ExecuteOperationSyncResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExecuteOperationSyncResponse{}
	_body, _err := client.ExecuteOperationSyncWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The URL of the application topology image.
//
// @param request - GetApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationResponse
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

// Summary:
//
// The URL of the application topology image.
//
// @param request - GetApplicationRequest
//
// @return GetApplicationResponse
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

// Summary:
//
// 获取应用输入参数
//
// @param request - GetApplicationVariablesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationVariablesResponse
func (client *Client) GetApplicationVariablesWithOptions(request *GetApplicationVariablesRequest, runtime *util.RuntimeOptions) (_result *GetApplicationVariablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplicationVariables"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationVariablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取应用输入参数
//
// @param request - GetApplicationVariablesRequest
//
// @return GetApplicationVariablesResponse
func (client *Client) GetApplicationVariables(request *GetApplicationVariablesRequest) (_result *GetApplicationVariablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationVariablesResponse{}
	_body, _err := client.GetApplicationVariablesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取需要重新配置的变量列表
//
// @param request - GetApplicationVariables4FailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApplicationVariables4FailResponse
func (client *Client) GetApplicationVariables4FailWithOptions(request *GetApplicationVariables4FailRequest, runtime *util.RuntimeOptions) (_result *GetApplicationVariables4FailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetApplicationVariables4Fail"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetApplicationVariables4FailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取需要重新配置的变量列表
//
// @param request - GetApplicationVariables4FailRequest
//
// @return GetApplicationVariables4FailResponse
func (client *Client) GetApplicationVariables4Fail(request *GetApplicationVariables4FailRequest) (_result *GetApplicationVariables4FailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetApplicationVariables4FailResponse{}
	_body, _err := client.GetApplicationVariables4FailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Asynchronously queries the result of an operation that is performed on a service instance.
//
// @param request - GetExecuteOperationResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExecuteOperationResultResponse
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

// Summary:
//
// Asynchronously queries the result of an operation that is performed on a service instance.
//
// @param request - GetExecuteOperationResultRequest
//
// @return GetExecuteOperationResultResponse
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

// Summary:
//
// Queries the status of a disaster recovery switchover task by task ID.
//
// Description:
//
// You can call this operation to query the status of a disaster recovery switchover task by task ID.
//
// @param request - GetFoTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFoTaskStatusResponse
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

// Summary:
//
// Queries the status of a disaster recovery switchover task by task ID.
//
// Description:
//
// You can call this operation to query the status of a disaster recovery switchover task by task ID.
//
// @param request - GetFoTaskStatusRequest
//
// @return GetFoTaskStatusResponse
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

// Summary:
//
// Queries the zones where the specified disaster recovery service can be switched.
//
// Description:
//
// You can call this operation to query the zones where the specified disaster recovery service can be switched.
//
// @param request - GetPotentialFailZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPotentialFailZonesResponse
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

// Summary:
//
// Queries the zones where the specified disaster recovery service can be switched.
//
// Description:
//
// You can call this operation to query the zones where the specified disaster recovery service can be switched.
//
// @param request - GetPotentialFailZonesRequest
//
// @return GetPotentialFailZonesResponse
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

// Summary:
//
// 获取询价应用变配记录
//
// @param request - GetResource4ModifyRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResource4ModifyRecordResponse
func (client *Client) GetResource4ModifyRecordWithOptions(request *GetResource4ModifyRecordRequest, runtime *util.RuntimeOptions) (_result *GetResource4ModifyRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResource4ModifyRecord"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResource4ModifyRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取询价应用变配记录
//
// @param request - GetResource4ModifyRecordRequest
//
// @return GetResource4ModifyRecordResponse
func (client *Client) GetResource4ModifyRecord(request *GetResource4ModifyRecordRequest) (_result *GetResource4ModifyRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResource4ModifyRecordResponse{}
	_body, _err := client.GetResource4ModifyRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取询价结果
//
// @param request - GetResult4QueryInstancePrice4ModifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetResult4QueryInstancePrice4ModifyResponse
func (client *Client) GetResult4QueryInstancePrice4ModifyWithOptions(request *GetResult4QueryInstancePrice4ModifyRequest, runtime *util.RuntimeOptions) (_result *GetResult4QueryInstancePrice4ModifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetResult4QueryInstancePrice4Modify"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetResult4QueryInstancePrice4ModifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取询价结果
//
// @param request - GetResult4QueryInstancePrice4ModifyRequest
//
// @return GetResult4QueryInstancePrice4ModifyResponse
func (client *Client) GetResult4QueryInstancePrice4Modify(request *GetResult4QueryInstancePrice4ModifyRequest) (_result *GetResult4QueryInstancePrice4ModifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetResult4QueryInstancePrice4ModifyResponse{}
	_body, _err := client.GetResult4QueryInstancePrice4ModifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get template images and information about architecture diagrams.
//
// @param request - GetTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
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

// Summary:
//
// Get template images and information about architecture diagrams.
//
// @param request - GetTemplateRequest
//
// @return GetTemplateResponse
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

// Deprecated: OpenAPI GetToken is deprecated, please use BPStudio::2021-09-31::GetApplication instead.
//
// Summary:
//
// Obtains a temporary token that is used to read the architecture diagram. The validity period of the token is 30 minutes.
//
// Description:
//
// 	Danger:  This API is no longer recommended, and the image related to the Application has included access authorization in the GetApplication property.
//
// @param request - GetTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
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

// Deprecated: OpenAPI GetToken is deprecated, please use BPStudio::2021-09-31::GetApplication instead.
//
// Summary:
//
// Obtains a temporary token that is used to read the architecture diagram. The validity period of the token is 30 minutes.
//
// Description:
//
// 	Danger:  This API is no longer recommended, and the image related to the Application has included access authorization in the GetApplication property.
//
// @param request - GetTokenRequest
//
// @return GetTokenResponse
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

// Summary:
//
// Prepares for application switchover and initiates a switchover task.
//
// Description:
//
// You can call this operation to prepare for application switchover and initiate a switchover task.
//
// @param request - InitAppFailOverRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitAppFailOverResponse
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

// Summary:
//
// Prepares for application switchover and initiates a switchover task.
//
// Description:
//
// You can call this operation to prepare for application switchover and initiate a switchover task.
//
// @param request - InitAppFailOverRequest
//
// @return InitAppFailOverResponse
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

// Summary:
//
// This API provides a list of all applications under the current user. The optional keyword parameter defines the keywords contained in the application name.
//
// @param request - ListApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationResponse
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

// Summary:
//
// This API provides a list of all applications under the current user. The optional keyword parameter defines the keywords contained in the application name.
//
// @param request - ListApplicationRequest
//
// @return ListApplicationResponse
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

// Summary:
//
// Queries disaster recovery plans.
//
// Description:
//
// You can call this operation to query all disaster recovery plans.
//
// @param request - ListFoCreatedAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFoCreatedAppsResponse
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

// Summary:
//
// Queries disaster recovery plans.
//
// Description:
//
// You can call this operation to query all disaster recovery plans.
//
// @return ListFoCreatedAppsResponse
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

// Summary:
//
// Queries the tags of one or more applications or templates.
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
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

// Summary:
//
// Queries the tags of one or more applications or templates.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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

// Summary:
//
// Queries templates, including information such as the template name, architecture image URL, and URL of the serialized architecture image file.
//
// @param request - ListTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplateResponse
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

// Summary:
//
// Queries templates, including information such as the template name, architecture image URL, and URL of the serialized architecture image file.
//
// @param request - ListTemplateRequest
//
// @return ListTemplateResponse
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

// Summary:
//
// 提交应用变配
//
// @param tmpReq - ModifyApplicationSpecRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyApplicationSpecResponse
func (client *Client) ModifyApplicationSpecWithOptions(tmpReq *ModifyApplicationSpecRequest, runtime *util.RuntimeOptions) (_result *ModifyApplicationSpecResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyApplicationSpecShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.InstanceSpec)) {
		request.InstanceSpecShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.InstanceSpec, tea.String("InstanceSpec"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceSpecShrink)) {
		body["InstanceSpec"] = request.InstanceSpecShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApplicationSpec"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyApplicationSpecResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交应用变配
//
// @param request - ModifyApplicationSpecRequest
//
// @return ModifyApplicationSpecResponse
func (client *Client) ModifyApplicationSpec(request *ModifyApplicationSpecRequest) (_result *ModifyApplicationSpecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyApplicationSpecResponse{}
	_body, _err := client.ModifyApplicationSpecWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询变配价格
//
// @param tmpReq - QueryInstancePrice4ModifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstancePrice4ModifyResponse
func (client *Client) QueryInstancePrice4ModifyWithOptions(tmpReq *QueryInstancePrice4ModifyRequest, runtime *util.RuntimeOptions) (_result *QueryInstancePrice4ModifyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryInstancePrice4ModifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Configuration)) {
		request.ConfigurationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Configuration, tea.String("Configuration"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.ConfigurationShrink)) {
		body["Configuration"] = request.ConfigurationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryInstancePrice4Modify"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInstancePrice4ModifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询变配价格
//
// @param request - QueryInstancePrice4ModifyRequest
//
// @return QueryInstancePrice4ModifyResponse
func (client *Client) QueryInstancePrice4Modify(request *QueryInstancePrice4ModifyRequest) (_result *QueryInstancePrice4ModifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInstancePrice4ModifyResponse{}
	_body, _err := client.QueryInstancePrice4ModifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询变配规格列表
//
// @param tmpReq - QueryInstanceSpec4ModifyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInstanceSpec4ModifyResponse
func (client *Client) QueryInstanceSpec4ModifyWithOptions(tmpReq *QueryInstanceSpec4ModifyRequest, runtime *util.RuntimeOptions) (_result *QueryInstanceSpec4ModifyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &QueryInstanceSpec4ModifyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MethodName)) {
		body["MethodName"] = request.MethodName
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		body["Parameters"] = request.ParametersShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryInstanceSpec4Modify"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryInstanceSpec4ModifyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询变配规格列表
//
// @param request - QueryInstanceSpec4ModifyRequest
//
// @return QueryInstanceSpec4ModifyResponse
func (client *Client) QueryInstanceSpec4Modify(request *QueryInstanceSpec4ModifyRequest) (_result *QueryInstanceSpec4ModifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryInstanceSpec4ModifyResponse{}
	_body, _err := client.QueryInstanceSpec4ModifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重新配置应用
//
// @param request - ReConfigApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReConfigApplicationResponse
func (client *Client) ReConfigApplicationWithOptions(request *ReConfigApplicationRequest, runtime *util.RuntimeOptions) (_result *ReConfigApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.Variables)) {
		body["Variables"] = request.Variables
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReConfigApplication"),
		Version:     tea.String("2021-09-31"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReConfigApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新配置应用
//
// @param request - ReConfigApplicationRequest
//
// @return ReConfigApplicationResponse
func (client *Client) ReConfigApplication(request *ReConfigApplicationRequest) (_result *ReConfigApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReConfigApplicationResponse{}
	_body, _err := client.ReConfigApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Releases the resources of an application.
//
// @param request - ReleaseApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseApplicationResponse
func (client *Client) ReleaseApplicationWithOptions(request *ReleaseApplicationRequest, runtime *util.RuntimeOptions) (_result *ReleaseApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
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

// Summary:
//
// Releases the resources of an application.
//
// @param request - ReleaseApplicationRequest
//
// @return ReleaseApplicationResponse
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

// Summary:
//
// Verifies the resources of an application. ValidateApplication is an asynchronous operation. You can call the GetApplication operation to query the verification result.
//
// @param request - ValidateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateApplicationResponse
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
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

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

// Summary:
//
// Verifies the resources of an application. ValidateApplication is an asynchronous operation. You can call the GetApplication operation to query the verification result.
//
// @param request - ValidateApplicationRequest
//
// @return ValidateApplicationResponse
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

// Summary:
//
// Queries the prices of resources of an application. You can call the GetApplication operation to obtain the query results.
//
// @param request - ValuateApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValuateApplicationResponse
func (client *Client) ValuateApplicationWithOptions(request *ValuateApplicationRequest, runtime *util.RuntimeOptions) (_result *ValuateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
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

// Summary:
//
// Queries the prices of resources of an application. You can call the GetApplication operation to obtain the query results.
//
// @param request - ValuateApplicationRequest
//
// @return ValuateApplicationResponse
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

// Summary:
//
// Queries the price of a template.
//
// @param tmpReq - ValuateTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValuateTemplateResponse
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

// Summary:
//
// Queries the price of a template.
//
// @param request - ValuateTemplateRequest
//
// @return ValuateTemplateResponse
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
