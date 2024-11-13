// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AttachApplication2ConnectorRequest struct {
	// This parameter is required.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// ConnectorID。
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s AttachApplication2ConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachApplication2ConnectorRequest) GoString() string {
	return s.String()
}

func (s *AttachApplication2ConnectorRequest) SetApplicationIds(v []*string) *AttachApplication2ConnectorRequest {
	s.ApplicationIds = v
	return s
}

func (s *AttachApplication2ConnectorRequest) SetConnectorId(v string) *AttachApplication2ConnectorRequest {
	s.ConnectorId = &v
	return s
}

type AttachApplication2ConnectorShrinkRequest struct {
	// This parameter is required.
	ApplicationIdsShrink *string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty"`
	// ConnectorID。
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s AttachApplication2ConnectorShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AttachApplication2ConnectorShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachApplication2ConnectorShrinkRequest) SetApplicationIdsShrink(v string) *AttachApplication2ConnectorShrinkRequest {
	s.ApplicationIdsShrink = &v
	return s
}

func (s *AttachApplication2ConnectorShrinkRequest) SetConnectorId(v string) *AttachApplication2ConnectorShrinkRequest {
	s.ConnectorId = &v
	return s
}

type AttachApplication2ConnectorResponseBody struct {
	// example:
	//
	// 7E9D7ACD-53D5-56EF-A913-79D148D06299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachApplication2ConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AttachApplication2ConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *AttachApplication2ConnectorResponseBody) SetRequestId(v string) *AttachApplication2ConnectorResponseBody {
	s.RequestId = &v
	return s
}

type AttachApplication2ConnectorResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachApplication2ConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachApplication2ConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s AttachApplication2ConnectorResponse) GoString() string {
	return s.String()
}

func (s *AttachApplication2ConnectorResponse) SetHeaders(v map[string]*string) *AttachApplication2ConnectorResponse {
	s.Headers = v
	return s
}

func (s *AttachApplication2ConnectorResponse) SetStatusCode(v int32) *AttachApplication2ConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachApplication2ConnectorResponse) SetBody(v *AttachApplication2ConnectorResponseBody) *AttachApplication2ConnectorResponse {
	s.Body = v
	return s
}

type CreateClientUserRequest struct {
	// example:
	//
	// 10797
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// johndoe@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 727
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	// example:
	//
	// 13641966835
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	// example:
	//
	// kehudiyi
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateClientUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClientUserRequest) GoString() string {
	return s.String()
}

func (s *CreateClientUserRequest) SetDepartmentId(v string) *CreateClientUserRequest {
	s.DepartmentId = &v
	return s
}

func (s *CreateClientUserRequest) SetDescription(v string) *CreateClientUserRequest {
	s.Description = &v
	return s
}

func (s *CreateClientUserRequest) SetEmail(v string) *CreateClientUserRequest {
	s.Email = &v
	return s
}

func (s *CreateClientUserRequest) SetIdpConfigId(v string) *CreateClientUserRequest {
	s.IdpConfigId = &v
	return s
}

func (s *CreateClientUserRequest) SetMobileNumber(v string) *CreateClientUserRequest {
	s.MobileNumber = &v
	return s
}

func (s *CreateClientUserRequest) SetPassword(v string) *CreateClientUserRequest {
	s.Password = &v
	return s
}

func (s *CreateClientUserRequest) SetUsername(v string) *CreateClientUserRequest {
	s.Username = &v
	return s
}

type CreateClientUserResponseBody struct {
	// example:
	//
	// 726
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateClientUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClientUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClientUserResponseBody) SetData(v string) *CreateClientUserResponseBody {
	s.Data = &v
	return s
}

func (s *CreateClientUserResponseBody) SetRequestId(v string) *CreateClientUserResponseBody {
	s.RequestId = &v
	return s
}

type CreateClientUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClientUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClientUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClientUserResponse) GoString() string {
	return s.String()
}

func (s *CreateClientUserResponse) SetHeaders(v map[string]*string) *CreateClientUserResponse {
	s.Headers = v
	return s
}

func (s *CreateClientUserResponse) SetStatusCode(v int32) *CreateClientUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClientUserResponse) SetBody(v *CreateClientUserResponseBody) *CreateClientUserResponse {
	s.Body = v
	return s
}

type CreateDynamicRouteRequest struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// connector
	DynamicRouteType *string `json:"DynamicRouteType,omitempty" xml:"DynamicRouteType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dynamic_route_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// connector-8ccb13b6f52c****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 99
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// This parameter is required.
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Disabled
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s CreateDynamicRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDynamicRouteRequest) GoString() string {
	return s.String()
}

func (s *CreateDynamicRouteRequest) SetApplicationIds(v []*string) *CreateDynamicRouteRequest {
	s.ApplicationIds = v
	return s
}

func (s *CreateDynamicRouteRequest) SetApplicationType(v string) *CreateDynamicRouteRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetDescription(v string) *CreateDynamicRouteRequest {
	s.Description = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetDynamicRouteType(v string) *CreateDynamicRouteRequest {
	s.DynamicRouteType = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetName(v string) *CreateDynamicRouteRequest {
	s.Name = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetNextHop(v string) *CreateDynamicRouteRequest {
	s.NextHop = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetPriority(v int32) *CreateDynamicRouteRequest {
	s.Priority = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetRegionIds(v []*string) *CreateDynamicRouteRequest {
	s.RegionIds = v
	return s
}

func (s *CreateDynamicRouteRequest) SetStatus(v string) *CreateDynamicRouteRequest {
	s.Status = &v
	return s
}

func (s *CreateDynamicRouteRequest) SetTagIds(v []*string) *CreateDynamicRouteRequest {
	s.TagIds = v
	return s
}

type CreateDynamicRouteResponseBody struct {
	// example:
	//
	// dr-ca9fddfac7c6****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDynamicRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDynamicRouteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDynamicRouteResponseBody) SetDynamicRouteId(v string) *CreateDynamicRouteResponseBody {
	s.DynamicRouteId = &v
	return s
}

func (s *CreateDynamicRouteResponseBody) SetRequestId(v string) *CreateDynamicRouteResponseBody {
	s.RequestId = &v
	return s
}

type CreateDynamicRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDynamicRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDynamicRouteResponse) GoString() string {
	return s.String()
}

func (s *CreateDynamicRouteResponse) SetHeaders(v map[string]*string) *CreateDynamicRouteResponse {
	s.Headers = v
	return s
}

func (s *CreateDynamicRouteResponse) SetStatusCode(v int32) *CreateDynamicRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDynamicRouteResponse) SetBody(v *CreateDynamicRouteResponseBody) *CreateDynamicRouteResponse {
	s.Body = v
	return s
}

type CreateIdpDepartmentRequest struct {
	// This parameter is required.
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1222
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
}

func (s CreateIdpDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIdpDepartmentRequest) GoString() string {
	return s.String()
}

func (s *CreateIdpDepartmentRequest) SetDepartmentName(v string) *CreateIdpDepartmentRequest {
	s.DepartmentName = &v
	return s
}

func (s *CreateIdpDepartmentRequest) SetIdpConfigId(v string) *CreateIdpDepartmentRequest {
	s.IdpConfigId = &v
	return s
}

type CreateIdpDepartmentResponseBody struct {
	// example:
	//
	// 726
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIdpDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateIdpDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIdpDepartmentResponseBody) SetData(v string) *CreateIdpDepartmentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateIdpDepartmentResponseBody) SetRequestId(v string) *CreateIdpDepartmentResponseBody {
	s.RequestId = &v
	return s
}

type CreateIdpDepartmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIdpDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIdpDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIdpDepartmentResponse) GoString() string {
	return s.String()
}

func (s *CreateIdpDepartmentResponse) SetHeaders(v map[string]*string) *CreateIdpDepartmentResponse {
	s.Headers = v
	return s
}

func (s *CreateIdpDepartmentResponse) SetStatusCode(v int32) *CreateIdpDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIdpDepartmentResponse) SetBody(v *CreateIdpDepartmentResponseBody) *CreateIdpDepartmentResponse {
	s.Body = v
	return s
}

type CreatePrivateAccessApplicationRequest struct {
	// This parameter is required.
	Addresses                    []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	BrowserAccessStatus          *string   `json:"BrowserAccessStatus,omitempty" xml:"BrowserAccessStatus,omitempty"`
	Description                  *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	L7ProxyDomainAutomaticPrefix *string   `json:"L7ProxyDomainAutomaticPrefix,omitempty" xml:"L7ProxyDomainAutomaticPrefix,omitempty"`
	L7ProxyDomainCustom          *string   `json:"L7ProxyDomainCustom,omitempty" xml:"L7ProxyDomainCustom,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// private_access_application_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	PortRanges []*CreatePrivateAccessApplicationRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// All
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s CreatePrivateAccessApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessApplicationRequest) SetAddresses(v []*string) *CreatePrivateAccessApplicationRequest {
	s.Addresses = v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetBrowserAccessStatus(v string) *CreatePrivateAccessApplicationRequest {
	s.BrowserAccessStatus = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetDescription(v string) *CreatePrivateAccessApplicationRequest {
	s.Description = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetL7ProxyDomainAutomaticPrefix(v string) *CreatePrivateAccessApplicationRequest {
	s.L7ProxyDomainAutomaticPrefix = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetL7ProxyDomainCustom(v string) *CreatePrivateAccessApplicationRequest {
	s.L7ProxyDomainCustom = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetName(v string) *CreatePrivateAccessApplicationRequest {
	s.Name = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetPortRanges(v []*CreatePrivateAccessApplicationRequestPortRanges) *CreatePrivateAccessApplicationRequest {
	s.PortRanges = v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetProtocol(v string) *CreatePrivateAccessApplicationRequest {
	s.Protocol = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetStatus(v string) *CreatePrivateAccessApplicationRequest {
	s.Status = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequest) SetTagIds(v []*string) *CreatePrivateAccessApplicationRequest {
	s.TagIds = v
	return s
}

type CreatePrivateAccessApplicationRequestPortRanges struct {
	// This parameter is required.
	//
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 81
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s CreatePrivateAccessApplicationRequestPortRanges) String() string {
	return tea.Prettify(s)
}

func (s CreatePrivateAccessApplicationRequestPortRanges) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessApplicationRequestPortRanges) SetBegin(v int32) *CreatePrivateAccessApplicationRequestPortRanges {
	s.Begin = &v
	return s
}

func (s *CreatePrivateAccessApplicationRequestPortRanges) SetEnd(v int32) *CreatePrivateAccessApplicationRequestPortRanges {
	s.End = &v
	return s
}

type CreatePrivateAccessApplicationResponseBody struct {
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePrivateAccessApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePrivateAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessApplicationResponseBody) SetApplicationId(v string) *CreatePrivateAccessApplicationResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *CreatePrivateAccessApplicationResponseBody) SetRequestId(v string) *CreatePrivateAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

type CreatePrivateAccessApplicationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrivateAccessApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePrivateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessApplicationResponse) SetHeaders(v map[string]*string) *CreatePrivateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *CreatePrivateAccessApplicationResponse) SetStatusCode(v int32) *CreatePrivateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrivateAccessApplicationResponse) SetBody(v *CreatePrivateAccessApplicationResponseBody) *CreatePrivateAccessApplicationResponse {
	s.Body = v
	return s
}

type CreatePrivateAccessPolicyRequest struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Application
	ApplicationType       *string                                                 `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	CustomUserAttributes  []*CreatePrivateAccessPolicyRequestCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	Description           *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceAttributeAction *string                                                 `json:"DeviceAttributeAction,omitempty" xml:"DeviceAttributeAction,omitempty"`
	DeviceAttributeId     *string                                                 `json:"DeviceAttributeId,omitempty" xml:"DeviceAttributeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 内网访问标签ID集合。最多可输入100个内网访问标签ID。当**ApplicationType**为**Tag时**，必填。和**ApplicationIds**互斥。
	TagIds       []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// 内网访问策略的用户组类型。取值：
	//
	// - **Normal**：普通用户组。
	//
	// - **Custom**：自定义用户组。
	//
	// This parameter is required.
	//
	// example:
	//
	// Normal
	UserGroupMode *string `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
}

func (s CreatePrivateAccessPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessPolicyRequest) SetApplicationIds(v []*string) *CreatePrivateAccessPolicyRequest {
	s.ApplicationIds = v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetApplicationType(v string) *CreatePrivateAccessPolicyRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetCustomUserAttributes(v []*CreatePrivateAccessPolicyRequestCustomUserAttributes) *CreatePrivateAccessPolicyRequest {
	s.CustomUserAttributes = v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetDescription(v string) *CreatePrivateAccessPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetDeviceAttributeAction(v string) *CreatePrivateAccessPolicyRequest {
	s.DeviceAttributeAction = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetDeviceAttributeId(v string) *CreatePrivateAccessPolicyRequest {
	s.DeviceAttributeId = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetName(v string) *CreatePrivateAccessPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetPolicyAction(v string) *CreatePrivateAccessPolicyRequest {
	s.PolicyAction = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetPriority(v int32) *CreatePrivateAccessPolicyRequest {
	s.Priority = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetStatus(v string) *CreatePrivateAccessPolicyRequest {
	s.Status = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetTagIds(v []*string) *CreatePrivateAccessPolicyRequest {
	s.TagIds = v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetUserGroupIds(v []*string) *CreatePrivateAccessPolicyRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreatePrivateAccessPolicyRequest) SetUserGroupMode(v string) *CreatePrivateAccessPolicyRequest {
	s.UserGroupMode = &v
	return s
}

type CreatePrivateAccessPolicyRequestCustomUserAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePrivateAccessPolicyRequestCustomUserAttributes) String() string {
	return tea.Prettify(s)
}

func (s CreatePrivateAccessPolicyRequestCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessPolicyRequestCustomUserAttributes) SetIdpId(v int32) *CreatePrivateAccessPolicyRequestCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequestCustomUserAttributes) SetRelation(v string) *CreatePrivateAccessPolicyRequestCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequestCustomUserAttributes) SetUserGroupType(v string) *CreatePrivateAccessPolicyRequestCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *CreatePrivateAccessPolicyRequestCustomUserAttributes) SetValue(v string) *CreatePrivateAccessPolicyRequestCustomUserAttributes {
	s.Value = &v
	return s
}

type CreatePrivateAccessPolicyResponseBody struct {
	// example:
	//
	// pa-policy-867ef4007c8a****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// EFE7EBB2-449D-5BBB-B381-CA7839BC1649
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePrivateAccessPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessPolicyResponseBody) SetPolicyId(v string) *CreatePrivateAccessPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *CreatePrivateAccessPolicyResponseBody) SetRequestId(v string) *CreatePrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreatePrivateAccessPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrivateAccessPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *CreatePrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreatePrivateAccessPolicyResponse) SetStatusCode(v int32) *CreatePrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrivateAccessPolicyResponse) SetBody(v *CreatePrivateAccessPolicyResponseBody) *CreatePrivateAccessPolicyResponse {
	s.Body = v
	return s
}

type CreatePrivateAccessTagRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// tag_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePrivateAccessTagRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePrivateAccessTagRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessTagRequest) SetDescription(v string) *CreatePrivateAccessTagRequest {
	s.Description = &v
	return s
}

func (s *CreatePrivateAccessTagRequest) SetName(v string) *CreatePrivateAccessTagRequest {
	s.Name = &v
	return s
}

type CreatePrivateAccessTagResponseBody struct {
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s CreatePrivateAccessTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePrivateAccessTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessTagResponseBody) SetRequestId(v string) *CreatePrivateAccessTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrivateAccessTagResponseBody) SetTagId(v string) *CreatePrivateAccessTagResponseBody {
	s.TagId = &v
	return s
}

type CreatePrivateAccessTagResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrivateAccessTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrivateAccessTagResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePrivateAccessTagResponse) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessTagResponse) SetHeaders(v map[string]*string) *CreatePrivateAccessTagResponse {
	s.Headers = v
	return s
}

func (s *CreatePrivateAccessTagResponse) SetStatusCode(v int32) *CreatePrivateAccessTagResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrivateAccessTagResponse) SetBody(v *CreatePrivateAccessTagResponseBody) *CreatePrivateAccessTagResponse {
	s.Body = v
	return s
}

type CreateRegistrationPolicyRequest struct {
	CompanyLimitCount *CreateRegistrationPolicyRequestCompanyLimitCount `json:"CompanyLimitCount,omitempty" xml:"CompanyLimitCount,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// LimitAll
	CompanyLimitType *string `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// registration_policy_name
	Name               *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PersonalLimitCount *CreateRegistrationPolicyRequestPersonalLimitCount `json:"PersonalLimitCount,omitempty" xml:"PersonalLimitCount,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// LimitDiff
	PersonalLimitType *string `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	// example:
	//
	// 99
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s CreateRegistrationPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRegistrationPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyRequest) SetCompanyLimitCount(v *CreateRegistrationPolicyRequestCompanyLimitCount) *CreateRegistrationPolicyRequest {
	s.CompanyLimitCount = v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetCompanyLimitType(v string) *CreateRegistrationPolicyRequest {
	s.CompanyLimitType = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetDescription(v string) *CreateRegistrationPolicyRequest {
	s.Description = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetMatchMode(v string) *CreateRegistrationPolicyRequest {
	s.MatchMode = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetName(v string) *CreateRegistrationPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetPersonalLimitCount(v *CreateRegistrationPolicyRequestPersonalLimitCount) *CreateRegistrationPolicyRequest {
	s.PersonalLimitCount = v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetPersonalLimitType(v string) *CreateRegistrationPolicyRequest {
	s.PersonalLimitType = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetPriority(v int64) *CreateRegistrationPolicyRequest {
	s.Priority = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetStatus(v string) *CreateRegistrationPolicyRequest {
	s.Status = &v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetUserGroupIds(v []*string) *CreateRegistrationPolicyRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateRegistrationPolicyRequest) SetWhitelist(v []*string) *CreateRegistrationPolicyRequest {
	s.Whitelist = v
	return s
}

type CreateRegistrationPolicyRequestCompanyLimitCount struct {
	// example:
	//
	// 1
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 0
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 0
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s CreateRegistrationPolicyRequestCompanyLimitCount) String() string {
	return tea.Prettify(s)
}

func (s CreateRegistrationPolicyRequestCompanyLimitCount) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyRequestCompanyLimitCount) SetAll(v int32) *CreateRegistrationPolicyRequestCompanyLimitCount {
	s.All = &v
	return s
}

func (s *CreateRegistrationPolicyRequestCompanyLimitCount) SetMobile(v int32) *CreateRegistrationPolicyRequestCompanyLimitCount {
	s.Mobile = &v
	return s
}

func (s *CreateRegistrationPolicyRequestCompanyLimitCount) SetPC(v int32) *CreateRegistrationPolicyRequestCompanyLimitCount {
	s.PC = &v
	return s
}

type CreateRegistrationPolicyRequestPersonalLimitCount struct {
	// example:
	//
	// 0
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 3
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 2
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s CreateRegistrationPolicyRequestPersonalLimitCount) String() string {
	return tea.Prettify(s)
}

func (s CreateRegistrationPolicyRequestPersonalLimitCount) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyRequestPersonalLimitCount) SetAll(v int32) *CreateRegistrationPolicyRequestPersonalLimitCount {
	s.All = &v
	return s
}

func (s *CreateRegistrationPolicyRequestPersonalLimitCount) SetMobile(v int32) *CreateRegistrationPolicyRequestPersonalLimitCount {
	s.Mobile = &v
	return s
}

func (s *CreateRegistrationPolicyRequestPersonalLimitCount) SetPC(v int32) *CreateRegistrationPolicyRequestPersonalLimitCount {
	s.PC = &v
	return s
}

type CreateRegistrationPolicyShrinkRequest struct {
	CompanyLimitCountShrink *string `json:"CompanyLimitCount,omitempty" xml:"CompanyLimitCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// LimitAll
	CompanyLimitType *string `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// registration_policy_name
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PersonalLimitCountShrink *string `json:"PersonalLimitCount,omitempty" xml:"PersonalLimitCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// LimitDiff
	PersonalLimitType *string `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	// example:
	//
	// 99
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s CreateRegistrationPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRegistrationPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyShrinkRequest) SetCompanyLimitCountShrink(v string) *CreateRegistrationPolicyShrinkRequest {
	s.CompanyLimitCountShrink = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetCompanyLimitType(v string) *CreateRegistrationPolicyShrinkRequest {
	s.CompanyLimitType = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetDescription(v string) *CreateRegistrationPolicyShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetMatchMode(v string) *CreateRegistrationPolicyShrinkRequest {
	s.MatchMode = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetName(v string) *CreateRegistrationPolicyShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetPersonalLimitCountShrink(v string) *CreateRegistrationPolicyShrinkRequest {
	s.PersonalLimitCountShrink = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetPersonalLimitType(v string) *CreateRegistrationPolicyShrinkRequest {
	s.PersonalLimitType = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetPriority(v int64) *CreateRegistrationPolicyShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetStatus(v string) *CreateRegistrationPolicyShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetUserGroupIds(v []*string) *CreateRegistrationPolicyShrinkRequest {
	s.UserGroupIds = v
	return s
}

func (s *CreateRegistrationPolicyShrinkRequest) SetWhitelist(v []*string) *CreateRegistrationPolicyShrinkRequest {
	s.Whitelist = v
	return s
}

type CreateRegistrationPolicyResponseBody struct {
	Policy *CreateRegistrationPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// example:
	//
	// FEF1144C-95D1-5F7C-81EF-9DB70EA49FCE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRegistrationPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRegistrationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyResponseBody) SetPolicy(v *CreateRegistrationPolicyResponseBodyPolicy) *CreateRegistrationPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *CreateRegistrationPolicyResponseBody) SetRequestId(v string) *CreateRegistrationPolicyResponseBody {
	s.RequestId = &v
	return s
}

type CreateRegistrationPolicyResponseBodyPolicy struct {
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime  *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail []*CreateRegistrationPolicyResponseBodyPolicyLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// example:
	//
	// registration_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// reg-policy-dcbfd33cb004****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s CreateRegistrationPolicyResponseBodyPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateRegistrationPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetCreateTime(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.CreateTime = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetDescription(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetLimitDetail(v []*CreateRegistrationPolicyResponseBodyPolicyLimitDetail) *CreateRegistrationPolicyResponseBodyPolicy {
	s.LimitDetail = v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetMatchMode(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.MatchMode = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetName(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.Name = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetPolicyId(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetPriority(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.Priority = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetStatus(v string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.Status = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetUserGroupIds(v []*string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.UserGroupIds = v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicy) SetWhitelist(v []*string) *CreateRegistrationPolicyResponseBodyPolicy {
	s.Whitelist = v
	return s
}

type CreateRegistrationPolicyResponseBodyPolicyLimitDetail struct {
	// example:
	//
	// Company
	DeviceBelong *string                                                          `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	// example:
	//
	// LimitDiff
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s CreateRegistrationPolicyResponseBodyPolicyLimitDetail) String() string {
	return tea.Prettify(s)
}

func (s CreateRegistrationPolicyResponseBodyPolicyLimitDetail) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetail) SetDeviceBelong(v string) *CreateRegistrationPolicyResponseBodyPolicyLimitDetail {
	s.DeviceBelong = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetail) SetLimitCount(v *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) *CreateRegistrationPolicyResponseBodyPolicyLimitDetail {
	s.LimitCount = v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetail) SetLimitType(v string) *CreateRegistrationPolicyResponseBodyPolicyLimitDetail {
	s.LimitType = &v
	return s
}

type CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount struct {
	// example:
	//
	// 0
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 3
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 2
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) String() string {
	return tea.Prettify(s)
}

func (s CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) SetAll(v int32) *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	s.All = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) SetMobile(v int32) *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	s.Mobile = &v
	return s
}

func (s *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) SetPC(v int32) *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	s.PC = &v
	return s
}

type CreateRegistrationPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRegistrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRegistrationPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRegistrationPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateRegistrationPolicyResponse) SetHeaders(v map[string]*string) *CreateRegistrationPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateRegistrationPolicyResponse) SetStatusCode(v int32) *CreateRegistrationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRegistrationPolicyResponse) SetBody(v *CreateRegistrationPolicyResponseBody) *CreateRegistrationPolicyResponse {
	s.Body = v
	return s
}

type CreateUserGroupRequest struct {
	// This parameter is required.
	Attributes  []*CreateUserGroupRequestAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Description *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_group_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequest) SetAttributes(v []*CreateUserGroupRequestAttributes) *CreateUserGroupRequest {
	s.Attributes = v
	return s
}

func (s *CreateUserGroupRequest) SetDescription(v string) *CreateUserGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateUserGroupRequest) SetName(v string) *CreateUserGroupRequest {
	s.Name = &v
	return s
}

type CreateUserGroupRequestAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateUserGroupRequestAttributes) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupRequestAttributes) GoString() string {
	return s.String()
}

func (s *CreateUserGroupRequestAttributes) SetIdpId(v int32) *CreateUserGroupRequestAttributes {
	s.IdpId = &v
	return s
}

func (s *CreateUserGroupRequestAttributes) SetRelation(v string) *CreateUserGroupRequestAttributes {
	s.Relation = &v
	return s
}

func (s *CreateUserGroupRequestAttributes) SetUserGroupType(v string) *CreateUserGroupRequestAttributes {
	s.UserGroupType = &v
	return s
}

func (s *CreateUserGroupRequestAttributes) SetValue(v string) *CreateUserGroupRequestAttributes {
	s.Value = &v
	return s
}

type CreateUserGroupResponseBody struct {
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s CreateUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponseBody) SetRequestId(v string) *CreateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetUserGroupId(v string) *CreateUserGroupResponseBody {
	s.UserGroupId = &v
	return s
}

type CreateUserGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponse) SetHeaders(v map[string]*string) *CreateUserGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateUserGroupResponse) SetStatusCode(v int32) *CreateUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserGroupResponse) SetBody(v *CreateUserGroupResponseBody) *CreateUserGroupResponse {
	s.Body = v
	return s
}

type CreateWmBaseImageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1080
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 255
	Opacity *int32 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Scale *int32 `json:"Scale,omitempty" xml:"Scale,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1920
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// aGVsbG8gc2F*****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// example:
	//
	// 12*****
	WmInfoUint *string `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PureWebappInvisible
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s CreateWmBaseImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWmBaseImageRequest) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageRequest) SetHeight(v int32) *CreateWmBaseImageRequest {
	s.Height = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetOpacity(v int32) *CreateWmBaseImageRequest {
	s.Opacity = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetScale(v int32) *CreateWmBaseImageRequest {
	s.Scale = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetWidth(v int32) *CreateWmBaseImageRequest {
	s.Width = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetWmInfoBytesB64(v string) *CreateWmBaseImageRequest {
	s.WmInfoBytesB64 = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetWmInfoSize(v int64) *CreateWmBaseImageRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetWmInfoUint(v string) *CreateWmBaseImageRequest {
	s.WmInfoUint = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetWmType(v string) *CreateWmBaseImageRequest {
	s.WmType = &v
	return s
}

type CreateWmBaseImageResponseBody struct {
	Data *CreateWmBaseImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWmBaseImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWmBaseImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageResponseBody) SetData(v *CreateWmBaseImageResponseBodyData) *CreateWmBaseImageResponseBody {
	s.Data = v
	return s
}

func (s *CreateWmBaseImageResponseBody) SetRequestId(v string) *CreateWmBaseImageResponseBody {
	s.RequestId = &v
	return s
}

type CreateWmBaseImageResponseBodyData struct {
	// example:
	//
	// fafb432cdede9b20640e12105845386e-496883833-8242409229217337*****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// https://example.com/test-*****.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// example:
	//
	// 17185*****
	ImageUrlExp *int64 `json:"ImageUrlExp,omitempty" xml:"ImageUrlExp,omitempty"`
}

func (s CreateWmBaseImageResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateWmBaseImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageResponseBodyData) SetImageId(v string) *CreateWmBaseImageResponseBodyData {
	s.ImageId = &v
	return s
}

func (s *CreateWmBaseImageResponseBodyData) SetImageUrl(v string) *CreateWmBaseImageResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *CreateWmBaseImageResponseBodyData) SetImageUrlExp(v int64) *CreateWmBaseImageResponseBodyData {
	s.ImageUrlExp = &v
	return s
}

type CreateWmBaseImageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWmBaseImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWmBaseImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWmBaseImageResponse) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageResponse) SetHeaders(v map[string]*string) *CreateWmBaseImageResponse {
	s.Headers = v
	return s
}

func (s *CreateWmBaseImageResponse) SetStatusCode(v int32) *CreateWmBaseImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWmBaseImageResponse) SetBody(v *CreateWmBaseImageResponseBody) *CreateWmBaseImageResponse {
	s.Body = v
	return s
}

type CreateWmEmbedTaskRequest struct {
	CsvControl      *CreateWmEmbedTaskRequestCsvControl      `json:"CsvControl,omitempty" xml:"CsvControl,omitempty" type:"Struct"`
	DocumentControl *CreateWmEmbedTaskRequestDocumentControl `json:"DocumentControl,omitempty" xml:"DocumentControl,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/abc****.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc****.pdf
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// example:
	//
	// 95
	ImageEmbedJpegQuality *int64 `json:"ImageEmbedJpegQuality,omitempty" xml:"ImageEmbedJpegQuality,omitempty"`
	// example:
	//
	// 2
	ImageEmbedLevel *int64 `json:"ImageEmbedLevel,omitempty" xml:"ImageEmbedLevel,omitempty"`
	// example:
	//
	// 3000k
	VideoBitrate *string `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// example:
	//
	// false
	VideoIsLong *bool `json:"VideoIsLong,omitempty" xml:"VideoIsLong,omitempty"`
	// example:
	//
	// aGVsbG8gc2F*****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// example:
	//
	// 123***
	WmInfoUint *string `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PureDocument
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s CreateWmEmbedTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWmEmbedTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequest) SetCsvControl(v *CreateWmEmbedTaskRequestCsvControl) *CreateWmEmbedTaskRequest {
	s.CsvControl = v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetDocumentControl(v *CreateWmEmbedTaskRequestDocumentControl) *CreateWmEmbedTaskRequest {
	s.DocumentControl = v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetFileUrl(v string) *CreateWmEmbedTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetFilename(v string) *CreateWmEmbedTaskRequest {
	s.Filename = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetImageEmbedJpegQuality(v int64) *CreateWmEmbedTaskRequest {
	s.ImageEmbedJpegQuality = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetImageEmbedLevel(v int64) *CreateWmEmbedTaskRequest {
	s.ImageEmbedLevel = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetVideoBitrate(v string) *CreateWmEmbedTaskRequest {
	s.VideoBitrate = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetVideoIsLong(v bool) *CreateWmEmbedTaskRequest {
	s.VideoIsLong = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetWmInfoBytesB64(v string) *CreateWmEmbedTaskRequest {
	s.WmInfoBytesB64 = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetWmInfoSize(v int64) *CreateWmEmbedTaskRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetWmInfoUint(v string) *CreateWmEmbedTaskRequest {
	s.WmInfoUint = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetWmType(v string) *CreateWmEmbedTaskRequest {
	s.WmType = &v
	return s
}

type CreateWmEmbedTaskRequestCsvControl struct {
	EmbedColumn    *int64  `json:"EmbedColumn,omitempty" xml:"EmbedColumn,omitempty"`
	EmbedPrecision *int64  `json:"EmbedPrecision,omitempty" xml:"EmbedPrecision,omitempty"`
	Method         *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s CreateWmEmbedTaskRequestCsvControl) String() string {
	return tea.Prettify(s)
}

func (s CreateWmEmbedTaskRequestCsvControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetEmbedColumn(v int64) *CreateWmEmbedTaskRequestCsvControl {
	s.EmbedColumn = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetEmbedPrecision(v int64) *CreateWmEmbedTaskRequestCsvControl {
	s.EmbedPrecision = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetMethod(v string) *CreateWmEmbedTaskRequestCsvControl {
	s.Method = &v
	return s
}

type CreateWmEmbedTaskRequestDocumentControl struct {
	BackgroundControl *CreateWmEmbedTaskRequestDocumentControlBackgroundControl `json:"BackgroundControl,omitempty" xml:"BackgroundControl,omitempty" type:"Struct"`
	// example:
	//
	// true
	InvisibleAntiAllCopy *bool `json:"InvisibleAntiAllCopy,omitempty" xml:"InvisibleAntiAllCopy,omitempty"`
	// example:
	//
	// true
	InvisibleAntiTextCopy *bool `json:"InvisibleAntiTextCopy,omitempty" xml:"InvisibleAntiTextCopy,omitempty"`
}

func (s CreateWmEmbedTaskRequestDocumentControl) String() string {
	return tea.Prettify(s)
}

func (s CreateWmEmbedTaskRequestDocumentControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestDocumentControl) SetBackgroundControl(v *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) *CreateWmEmbedTaskRequestDocumentControl {
	s.BackgroundControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControl) SetInvisibleAntiAllCopy(v bool) *CreateWmEmbedTaskRequestDocumentControl {
	s.InvisibleAntiAllCopy = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControl) SetInvisibleAntiTextCopy(v bool) *CreateWmEmbedTaskRequestDocumentControl {
	s.InvisibleAntiTextCopy = &v
	return s
}

type CreateWmEmbedTaskRequestDocumentControlBackgroundControl struct {
	// example:
	//
	// true
	BgAddInvisible *bool `json:"BgAddInvisible,omitempty" xml:"BgAddInvisible,omitempty"`
	// example:
	//
	// true
	BgAddVisible       *bool                                                                       `json:"BgAddVisible,omitempty" xml:"BgAddVisible,omitempty"`
	BgInvisibleControl *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl `json:"BgInvisibleControl,omitempty" xml:"BgInvisibleControl,omitempty" type:"Struct"`
	BgVisibleControl   *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl   `json:"BgVisibleControl,omitempty" xml:"BgVisibleControl,omitempty" type:"Struct"`
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControl) String() string {
	return tea.Prettify(s)
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) SetBgAddInvisible(v bool) *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	s.BgAddInvisible = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) SetBgAddVisible(v bool) *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	s.BgAddVisible = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) SetBgInvisibleControl(v *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	s.BgInvisibleControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) SetBgVisibleControl(v *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	s.BgVisibleControl = v
	return s
}

type CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl struct {
	// example:
	//
	// 10
	Opacity *int64 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) String() string {
	return tea.Prettify(s)
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) SetOpacity(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl {
	s.Opacity = &v
	return s
}

type CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl struct {
	// example:
	//
	// 30
	Angle *int64 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	// example:
	//
	// 0x000000
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// example:
	//
	// 30
	FontSize *int64 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// example:
	//
	// 3
	HorizontalNumber *int64 `json:"HorizontalNumber,omitempty" xml:"HorizontalNumber,omitempty"`
	// example:
	//
	// pos
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// 100
	Opacity *int64 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	// example:
	//
	// 0.5
	PosX *string `json:"PosX,omitempty" xml:"PosX,omitempty"`
	// example:
	//
	// 0.5
	PosY *string `json:"PosY,omitempty" xml:"PosY,omitempty"`
	// example:
	//
	// 3
	VerticalNumber *int64 `json:"VerticalNumber,omitempty" xml:"VerticalNumber,omitempty"`
	// example:
	//
	// hello ****
	VisibleText *string `json:"VisibleText,omitempty" xml:"VisibleText,omitempty"`
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) String() string {
	return tea.Prettify(s)
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetAngle(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.Angle = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetFontColor(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.FontColor = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetFontSize(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.FontSize = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetHorizontalNumber(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.HorizontalNumber = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetMode(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.Mode = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetOpacity(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.Opacity = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetPosX(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.PosX = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetPosY(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.PosY = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetVerticalNumber(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.VerticalNumber = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetVisibleText(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.VisibleText = &v
	return s
}

type CreateWmEmbedTaskShrinkRequest struct {
	CsvControlShrink      *string `json:"CsvControl,omitempty" xml:"CsvControl,omitempty"`
	DocumentControlShrink *string `json:"DocumentControl,omitempty" xml:"DocumentControl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/abc****.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc****.pdf
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// example:
	//
	// 95
	ImageEmbedJpegQuality *int64 `json:"ImageEmbedJpegQuality,omitempty" xml:"ImageEmbedJpegQuality,omitempty"`
	// example:
	//
	// 2
	ImageEmbedLevel *int64 `json:"ImageEmbedLevel,omitempty" xml:"ImageEmbedLevel,omitempty"`
	// example:
	//
	// 3000k
	VideoBitrate *string `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// example:
	//
	// false
	VideoIsLong *bool `json:"VideoIsLong,omitempty" xml:"VideoIsLong,omitempty"`
	// example:
	//
	// aGVsbG8gc2F*****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// example:
	//
	// 123***
	WmInfoUint *string `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PureDocument
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s CreateWmEmbedTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWmEmbedTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskShrinkRequest) SetCsvControlShrink(v string) *CreateWmEmbedTaskShrinkRequest {
	s.CsvControlShrink = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetDocumentControlShrink(v string) *CreateWmEmbedTaskShrinkRequest {
	s.DocumentControlShrink = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetFileUrl(v string) *CreateWmEmbedTaskShrinkRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetFilename(v string) *CreateWmEmbedTaskShrinkRequest {
	s.Filename = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetImageEmbedJpegQuality(v int64) *CreateWmEmbedTaskShrinkRequest {
	s.ImageEmbedJpegQuality = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetImageEmbedLevel(v int64) *CreateWmEmbedTaskShrinkRequest {
	s.ImageEmbedLevel = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetVideoBitrate(v string) *CreateWmEmbedTaskShrinkRequest {
	s.VideoBitrate = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetVideoIsLong(v bool) *CreateWmEmbedTaskShrinkRequest {
	s.VideoIsLong = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetWmInfoBytesB64(v string) *CreateWmEmbedTaskShrinkRequest {
	s.WmInfoBytesB64 = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetWmInfoSize(v int64) *CreateWmEmbedTaskShrinkRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetWmInfoUint(v string) *CreateWmEmbedTaskShrinkRequest {
	s.WmInfoUint = &v
	return s
}

func (s *CreateWmEmbedTaskShrinkRequest) SetWmType(v string) *CreateWmEmbedTaskShrinkRequest {
	s.WmType = &v
	return s
}

type CreateWmEmbedTaskResponseBody struct {
	Data *CreateWmEmbedTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWmEmbedTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWmEmbedTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskResponseBody) SetData(v *CreateWmEmbedTaskResponseBodyData) *CreateWmEmbedTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateWmEmbedTaskResponseBody) SetRequestId(v string) *CreateWmEmbedTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateWmEmbedTaskResponseBodyData struct {
	// example:
	//
	// job:5GfrJYsoaffmCE7Z5bZtjUefzxfd****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateWmEmbedTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateWmEmbedTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskResponseBodyData) SetTaskId(v string) *CreateWmEmbedTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type CreateWmEmbedTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWmEmbedTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWmEmbedTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWmEmbedTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskResponse) SetHeaders(v map[string]*string) *CreateWmEmbedTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateWmEmbedTaskResponse) SetStatusCode(v int32) *CreateWmEmbedTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWmEmbedTaskResponse) SetBody(v *CreateWmEmbedTaskResponseBody) *CreateWmEmbedTaskResponse {
	s.Body = v
	return s
}

type CreateWmExtractTaskRequest struct {
	CsvControl *CreateWmExtractTaskRequestCsvControl `json:"CsvControl,omitempty" xml:"CsvControl,omitempty" type:"Struct"`
	// example:
	//
	// false
	DocumentIsCapture *bool `json:"DocumentIsCapture,omitempty" xml:"DocumentIsCapture,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/test-****.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-****.pdf
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// example:
	//
	// false
	VideoIsLong *bool `json:"VideoIsLong,omitempty" xml:"VideoIsLong,omitempty"`
	// example:
	//
	// 1
	VideoSpeed *string `json:"VideoSpeed,omitempty" xml:"VideoSpeed,omitempty"`
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PureDocument
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s CreateWmExtractTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWmExtractTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateWmExtractTaskRequest) SetCsvControl(v *CreateWmExtractTaskRequestCsvControl) *CreateWmExtractTaskRequest {
	s.CsvControl = v
	return s
}

func (s *CreateWmExtractTaskRequest) SetDocumentIsCapture(v bool) *CreateWmExtractTaskRequest {
	s.DocumentIsCapture = &v
	return s
}

func (s *CreateWmExtractTaskRequest) SetFileUrl(v string) *CreateWmExtractTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateWmExtractTaskRequest) SetFilename(v string) *CreateWmExtractTaskRequest {
	s.Filename = &v
	return s
}

func (s *CreateWmExtractTaskRequest) SetVideoIsLong(v bool) *CreateWmExtractTaskRequest {
	s.VideoIsLong = &v
	return s
}

func (s *CreateWmExtractTaskRequest) SetVideoSpeed(v string) *CreateWmExtractTaskRequest {
	s.VideoSpeed = &v
	return s
}

func (s *CreateWmExtractTaskRequest) SetWmInfoSize(v int64) *CreateWmExtractTaskRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmExtractTaskRequest) SetWmType(v string) *CreateWmExtractTaskRequest {
	s.WmType = &v
	return s
}

type CreateWmExtractTaskRequestCsvControl struct {
	EmbedColumn    *int64  `json:"EmbedColumn,omitempty" xml:"EmbedColumn,omitempty"`
	EmbedPrecision *int64  `json:"EmbedPrecision,omitempty" xml:"EmbedPrecision,omitempty"`
	Method         *string `json:"Method,omitempty" xml:"Method,omitempty"`
}

func (s CreateWmExtractTaskRequestCsvControl) String() string {
	return tea.Prettify(s)
}

func (s CreateWmExtractTaskRequestCsvControl) GoString() string {
	return s.String()
}

func (s *CreateWmExtractTaskRequestCsvControl) SetEmbedColumn(v int64) *CreateWmExtractTaskRequestCsvControl {
	s.EmbedColumn = &v
	return s
}

func (s *CreateWmExtractTaskRequestCsvControl) SetEmbedPrecision(v int64) *CreateWmExtractTaskRequestCsvControl {
	s.EmbedPrecision = &v
	return s
}

func (s *CreateWmExtractTaskRequestCsvControl) SetMethod(v string) *CreateWmExtractTaskRequestCsvControl {
	s.Method = &v
	return s
}

type CreateWmExtractTaskShrinkRequest struct {
	CsvControlShrink *string `json:"CsvControl,omitempty" xml:"CsvControl,omitempty"`
	// example:
	//
	// false
	DocumentIsCapture *bool `json:"DocumentIsCapture,omitempty" xml:"DocumentIsCapture,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/test-****.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-****.pdf
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// example:
	//
	// false
	VideoIsLong *bool `json:"VideoIsLong,omitempty" xml:"VideoIsLong,omitempty"`
	// example:
	//
	// 1
	VideoSpeed *string `json:"VideoSpeed,omitempty" xml:"VideoSpeed,omitempty"`
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PureDocument
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s CreateWmExtractTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWmExtractTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWmExtractTaskShrinkRequest) SetCsvControlShrink(v string) *CreateWmExtractTaskShrinkRequest {
	s.CsvControlShrink = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetDocumentIsCapture(v bool) *CreateWmExtractTaskShrinkRequest {
	s.DocumentIsCapture = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetFileUrl(v string) *CreateWmExtractTaskShrinkRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetFilename(v string) *CreateWmExtractTaskShrinkRequest {
	s.Filename = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetVideoIsLong(v bool) *CreateWmExtractTaskShrinkRequest {
	s.VideoIsLong = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetVideoSpeed(v string) *CreateWmExtractTaskShrinkRequest {
	s.VideoSpeed = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetWmInfoSize(v int64) *CreateWmExtractTaskShrinkRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmExtractTaskShrinkRequest) SetWmType(v string) *CreateWmExtractTaskShrinkRequest {
	s.WmType = &v
	return s
}

type CreateWmExtractTaskResponseBody struct {
	Data *CreateWmExtractTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWmExtractTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWmExtractTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWmExtractTaskResponseBody) SetData(v *CreateWmExtractTaskResponseBodyData) *CreateWmExtractTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateWmExtractTaskResponseBody) SetRequestId(v string) *CreateWmExtractTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateWmExtractTaskResponseBodyData struct {
	// example:
	//
	// wmt-9648c22d2eb2cb57bb855dcae7898464********
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateWmExtractTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateWmExtractTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWmExtractTaskResponseBodyData) SetTaskId(v string) *CreateWmExtractTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type CreateWmExtractTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWmExtractTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWmExtractTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWmExtractTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateWmExtractTaskResponse) SetHeaders(v map[string]*string) *CreateWmExtractTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateWmExtractTaskResponse) SetStatusCode(v int32) *CreateWmExtractTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWmExtractTaskResponse) SetBody(v *CreateWmExtractTaskResponseBody) *CreateWmExtractTaskResponse {
	s.Body = v
	return s
}

type CreateWmInfoMappingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aGVsbG8gc2F*****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PureDocument
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s CreateWmInfoMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWmInfoMappingRequest) GoString() string {
	return s.String()
}

func (s *CreateWmInfoMappingRequest) SetWmInfoBytesB64(v string) *CreateWmInfoMappingRequest {
	s.WmInfoBytesB64 = &v
	return s
}

func (s *CreateWmInfoMappingRequest) SetWmInfoSize(v int64) *CreateWmInfoMappingRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmInfoMappingRequest) SetWmType(v string) *CreateWmInfoMappingRequest {
	s.WmType = &v
	return s
}

type CreateWmInfoMappingResponseBody struct {
	Data *CreateWmInfoMappingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 7E9D7ACD-53D5-56EF-A913-79D148D06299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateWmInfoMappingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateWmInfoMappingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWmInfoMappingResponseBody) SetData(v *CreateWmInfoMappingResponseBodyData) *CreateWmInfoMappingResponseBody {
	s.Data = v
	return s
}

func (s *CreateWmInfoMappingResponseBody) SetRequestId(v string) *CreateWmInfoMappingResponseBody {
	s.RequestId = &v
	return s
}

type CreateWmInfoMappingResponseBodyData struct {
	// example:
	//
	// 123***
	WmInfoUint *int64 `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
}

func (s CreateWmInfoMappingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateWmInfoMappingResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWmInfoMappingResponseBodyData) SetWmInfoUint(v int64) *CreateWmInfoMappingResponseBodyData {
	s.WmInfoUint = &v
	return s
}

type CreateWmInfoMappingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateWmInfoMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateWmInfoMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWmInfoMappingResponse) GoString() string {
	return s.String()
}

func (s *CreateWmInfoMappingResponse) SetHeaders(v map[string]*string) *CreateWmInfoMappingResponse {
	s.Headers = v
	return s
}

func (s *CreateWmInfoMappingResponse) SetStatusCode(v int32) *CreateWmInfoMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateWmInfoMappingResponse) SetBody(v *CreateWmInfoMappingResponseBody) *CreateWmInfoMappingResponse {
	s.Body = v
	return s
}

type DeleteClientUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 27058
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteClientUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClientUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientUserRequest) SetId(v string) *DeleteClientUserRequest {
	s.Id = &v
	return s
}

type DeleteClientUserResponseBody struct {
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClientUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClientUserResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClientUserResponseBody) SetRequestId(v string) *DeleteClientUserResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClientUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClientUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClientUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClientUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteClientUserResponse) SetHeaders(v map[string]*string) *DeleteClientUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteClientUserResponse) SetStatusCode(v int32) *DeleteClientUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClientUserResponse) SetBody(v *DeleteClientUserResponseBody) *DeleteClientUserResponse {
	s.Body = v
	return s
}

type DeleteDynamicRouteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dr-ca9fddfac7c6****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
}

func (s DeleteDynamicRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDynamicRouteRequest) GoString() string {
	return s.String()
}

func (s *DeleteDynamicRouteRequest) SetDynamicRouteId(v string) *DeleteDynamicRouteRequest {
	s.DynamicRouteId = &v
	return s
}

type DeleteDynamicRouteResponseBody struct {
	// example:
	//
	// 748CFDC7-1EB6-5B8B-9405-DA76ED5BB60D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDynamicRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDynamicRouteResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDynamicRouteResponseBody) SetRequestId(v string) *DeleteDynamicRouteResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDynamicRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDynamicRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDynamicRouteResponse) GoString() string {
	return s.String()
}

func (s *DeleteDynamicRouteResponse) SetHeaders(v map[string]*string) *DeleteDynamicRouteResponse {
	s.Headers = v
	return s
}

func (s *DeleteDynamicRouteResponse) SetStatusCode(v int32) *DeleteDynamicRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDynamicRouteResponse) SetBody(v *DeleteDynamicRouteResponseBody) *DeleteDynamicRouteResponse {
	s.Body = v
	return s
}

type DeleteIdpDepartmentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10829
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 507
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
}

func (s DeleteIdpDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteIdpDepartmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteIdpDepartmentRequest) SetDepartmentId(v string) *DeleteIdpDepartmentRequest {
	s.DepartmentId = &v
	return s
}

func (s *DeleteIdpDepartmentRequest) SetIdpConfigId(v string) *DeleteIdpDepartmentRequest {
	s.IdpConfigId = &v
	return s
}

type DeleteIdpDepartmentResponseBody struct {
	// example:
	//
	// FEF1144C-95D1-5F7C-81EF-9DB70EA49FCE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIdpDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteIdpDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIdpDepartmentResponseBody) SetRequestId(v string) *DeleteIdpDepartmentResponseBody {
	s.RequestId = &v
	return s
}

type DeleteIdpDepartmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIdpDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIdpDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteIdpDepartmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteIdpDepartmentResponse) SetHeaders(v map[string]*string) *DeleteIdpDepartmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteIdpDepartmentResponse) SetStatusCode(v int32) *DeleteIdpDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIdpDepartmentResponse) SetBody(v *DeleteIdpDepartmentResponseBody) *DeleteIdpDepartmentResponse {
	s.Body = v
	return s
}

type DeletePrivateAccessApplicationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s DeletePrivateAccessApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessApplicationRequest) SetApplicationId(v string) *DeletePrivateAccessApplicationRequest {
	s.ApplicationId = &v
	return s
}

type DeletePrivateAccessApplicationResponseBody struct {
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrivateAccessApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePrivateAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessApplicationResponseBody) SetRequestId(v string) *DeletePrivateAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

type DeletePrivateAccessApplicationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrivateAccessApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePrivateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessApplicationResponse) SetHeaders(v map[string]*string) *DeletePrivateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivateAccessApplicationResponse) SetStatusCode(v int32) *DeletePrivateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivateAccessApplicationResponse) SetBody(v *DeletePrivateAccessApplicationResponseBody) *DeletePrivateAccessApplicationResponse {
	s.Body = v
	return s
}

type DeletePrivateAccessPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-policy-867ef4007c8a****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DeletePrivateAccessPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessPolicyRequest) SetPolicyId(v string) *DeletePrivateAccessPolicyRequest {
	s.PolicyId = &v
	return s
}

type DeletePrivateAccessPolicyResponseBody struct {
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrivateAccessPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessPolicyResponseBody) SetRequestId(v string) *DeletePrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

type DeletePrivateAccessPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrivateAccessPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *DeletePrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivateAccessPolicyResponse) SetStatusCode(v int32) *DeletePrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivateAccessPolicyResponse) SetBody(v *DeletePrivateAccessPolicyResponseBody) *DeletePrivateAccessPolicyResponse {
	s.Body = v
	return s
}

type DeletePrivateAccessTagRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s DeletePrivateAccessTagRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePrivateAccessTagRequest) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessTagRequest) SetTagId(v string) *DeletePrivateAccessTagRequest {
	s.TagId = &v
	return s
}

type DeletePrivateAccessTagResponseBody struct {
	// example:
	//
	// FD724DBC-CD76-5235-BF76-59C51B73296D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrivateAccessTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePrivateAccessTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessTagResponseBody) SetRequestId(v string) *DeletePrivateAccessTagResponseBody {
	s.RequestId = &v
	return s
}

type DeletePrivateAccessTagResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePrivateAccessTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePrivateAccessTagResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePrivateAccessTagResponse) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessTagResponse) SetHeaders(v map[string]*string) *DeletePrivateAccessTagResponse {
	s.Headers = v
	return s
}

func (s *DeletePrivateAccessTagResponse) SetStatusCode(v int32) *DeletePrivateAccessTagResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePrivateAccessTagResponse) SetBody(v *DeletePrivateAccessTagResponseBody) *DeletePrivateAccessTagResponse {
	s.Body = v
	return s
}

type DeleteRegistrationPoliciesRequest struct {
	// This parameter is required.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s DeleteRegistrationPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRegistrationPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRegistrationPoliciesRequest) SetPolicyIds(v []*string) *DeleteRegistrationPoliciesRequest {
	s.PolicyIds = v
	return s
}

type DeleteRegistrationPoliciesResponseBody struct {
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRegistrationPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRegistrationPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRegistrationPoliciesResponseBody) SetRequestId(v string) *DeleteRegistrationPoliciesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRegistrationPoliciesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRegistrationPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRegistrationPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRegistrationPoliciesResponse) GoString() string {
	return s.String()
}

func (s *DeleteRegistrationPoliciesResponse) SetHeaders(v map[string]*string) *DeleteRegistrationPoliciesResponse {
	s.Headers = v
	return s
}

func (s *DeleteRegistrationPoliciesResponse) SetStatusCode(v int32) *DeleteRegistrationPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRegistrationPoliciesResponse) SetBody(v *DeleteRegistrationPoliciesResponseBody) *DeleteRegistrationPoliciesResponse {
	s.Body = v
	return s
}

type DeleteUserDevicesRequest struct {
	DeviceTags []*string `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty" type:"Repeated"`
}

func (s DeleteUserDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDevicesRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserDevicesRequest) SetDeviceTags(v []*string) *DeleteUserDevicesRequest {
	s.DeviceTags = v
	return s
}

type DeleteUserDevicesResponseBody struct {
	// example:
	//
	// 5FEF5CFA-14CC-5DE5-BD1F-AFFE0996E71D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserDevicesResponseBody) SetRequestId(v string) *DeleteUserDevicesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserDevicesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserDevicesResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserDevicesResponse) SetHeaders(v map[string]*string) *DeleteUserDevicesResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserDevicesResponse) SetStatusCode(v int32) *DeleteUserDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserDevicesResponse) SetBody(v *DeleteUserDevicesResponseBody) *DeleteUserDevicesResponse {
	s.Body = v
	return s
}

type DeleteUserGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s DeleteUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupRequest) SetUserGroupId(v string) *DeleteUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type DeleteUserGroupResponseBody struct {
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponseBody) SetRequestId(v string) *DeleteUserGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserGroupResponse) SetHeaders(v map[string]*string) *DeleteUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserGroupResponse) SetStatusCode(v int32) *DeleteUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserGroupResponse) SetBody(v *DeleteUserGroupResponseBody) *DeleteUserGroupResponse {
	s.Body = v
	return s
}

type DetachApplication2ConnectorRequest struct {
	// This parameter is required.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// ConnectorID。
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s DetachApplication2ConnectorRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachApplication2ConnectorRequest) GoString() string {
	return s.String()
}

func (s *DetachApplication2ConnectorRequest) SetApplicationIds(v []*string) *DetachApplication2ConnectorRequest {
	s.ApplicationIds = v
	return s
}

func (s *DetachApplication2ConnectorRequest) SetConnectorId(v string) *DetachApplication2ConnectorRequest {
	s.ConnectorId = &v
	return s
}

type DetachApplication2ConnectorShrinkRequest struct {
	// This parameter is required.
	ApplicationIdsShrink *string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty"`
	// ConnectorID。
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s DetachApplication2ConnectorShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DetachApplication2ConnectorShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetachApplication2ConnectorShrinkRequest) SetApplicationIdsShrink(v string) *DetachApplication2ConnectorShrinkRequest {
	s.ApplicationIdsShrink = &v
	return s
}

func (s *DetachApplication2ConnectorShrinkRequest) SetConnectorId(v string) *DetachApplication2ConnectorShrinkRequest {
	s.ConnectorId = &v
	return s
}

type DetachApplication2ConnectorResponseBody struct {
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachApplication2ConnectorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetachApplication2ConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *DetachApplication2ConnectorResponseBody) SetRequestId(v string) *DetachApplication2ConnectorResponseBody {
	s.RequestId = &v
	return s
}

type DetachApplication2ConnectorResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachApplication2ConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachApplication2ConnectorResponse) String() string {
	return tea.Prettify(s)
}

func (s DetachApplication2ConnectorResponse) GoString() string {
	return s.String()
}

func (s *DetachApplication2ConnectorResponse) SetHeaders(v map[string]*string) *DetachApplication2ConnectorResponse {
	s.Headers = v
	return s
}

func (s *DetachApplication2ConnectorResponse) SetStatusCode(v int32) *DetachApplication2ConnectorResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachApplication2ConnectorResponse) SetBody(v *DetachApplication2ConnectorResponseBody) *DetachApplication2ConnectorResponse {
	s.Body = v
	return s
}

type ExportUserDevicesRequest struct {
	AppStatuses []*string `json:"AppStatuses,omitempty" xml:"AppStatuses,omitempty" type:"Repeated"`
	Department  *string   `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// Company
	DeviceBelong   *string   `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	DeviceStatuses []*string `json:"DeviceStatuses,omitempty" xml:"DeviceStatuses,omitempty" type:"Repeated"`
	DeviceTags     []*string `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty" type:"Repeated"`
	DeviceTypes    []*string `json:"DeviceTypes,omitempty" xml:"DeviceTypes,omitempty" type:"Repeated"`
	DlpStatuses    []*string `json:"DlpStatuses,omitempty" xml:"DlpStatuses,omitempty" type:"Repeated"`
	// example:
	//
	// win10-64bit
	Hostname   *string   `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	IaStatuses []*string `json:"IaStatuses,omitempty" xml:"IaStatuses,omitempty" type:"Repeated"`
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac         *string   `json:"Mac,omitempty" xml:"Mac,omitempty"`
	NacStatuses []*string `json:"NacStatuses,omitempty" xml:"NacStatuses,omitempty" type:"Repeated"`
	PaStatuses  []*string `json:"PaStatuses,omitempty" xml:"PaStatuses,omitempty" type:"Repeated"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// example:
	//
	// true
	SharingStatus *bool   `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	Username      *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ExportUserDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ExportUserDevicesRequest) GoString() string {
	return s.String()
}

func (s *ExportUserDevicesRequest) SetAppStatuses(v []*string) *ExportUserDevicesRequest {
	s.AppStatuses = v
	return s
}

func (s *ExportUserDevicesRequest) SetDepartment(v string) *ExportUserDevicesRequest {
	s.Department = &v
	return s
}

func (s *ExportUserDevicesRequest) SetDeviceBelong(v string) *ExportUserDevicesRequest {
	s.DeviceBelong = &v
	return s
}

func (s *ExportUserDevicesRequest) SetDeviceStatuses(v []*string) *ExportUserDevicesRequest {
	s.DeviceStatuses = v
	return s
}

func (s *ExportUserDevicesRequest) SetDeviceTags(v []*string) *ExportUserDevicesRequest {
	s.DeviceTags = v
	return s
}

func (s *ExportUserDevicesRequest) SetDeviceTypes(v []*string) *ExportUserDevicesRequest {
	s.DeviceTypes = v
	return s
}

func (s *ExportUserDevicesRequest) SetDlpStatuses(v []*string) *ExportUserDevicesRequest {
	s.DlpStatuses = v
	return s
}

func (s *ExportUserDevicesRequest) SetHostname(v string) *ExportUserDevicesRequest {
	s.Hostname = &v
	return s
}

func (s *ExportUserDevicesRequest) SetIaStatuses(v []*string) *ExportUserDevicesRequest {
	s.IaStatuses = v
	return s
}

func (s *ExportUserDevicesRequest) SetMac(v string) *ExportUserDevicesRequest {
	s.Mac = &v
	return s
}

func (s *ExportUserDevicesRequest) SetNacStatuses(v []*string) *ExportUserDevicesRequest {
	s.NacStatuses = v
	return s
}

func (s *ExportUserDevicesRequest) SetPaStatuses(v []*string) *ExportUserDevicesRequest {
	s.PaStatuses = v
	return s
}

func (s *ExportUserDevicesRequest) SetSaseUserId(v string) *ExportUserDevicesRequest {
	s.SaseUserId = &v
	return s
}

func (s *ExportUserDevicesRequest) SetSharingStatus(v bool) *ExportUserDevicesRequest {
	s.SharingStatus = &v
	return s
}

func (s *ExportUserDevicesRequest) SetUsername(v string) *ExportUserDevicesRequest {
	s.Username = &v
	return s
}

type ExportUserDevicesResponseBody struct {
	// example:
	//
	// 748CFDC7-1EB6-5B8B-9405-DA76ED5BB60D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// https://sase-export.oss-cn-hangzhou.aliyuncs.com/export%2Fapp-device%2F20240607154831.xlsx?Expires=1717746571&OSSAccessKeyId=********************
	SignedUrl *string `json:"SignedUrl,omitempty" xml:"SignedUrl,omitempty"`
}

func (s ExportUserDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExportUserDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ExportUserDevicesResponseBody) SetRequestId(v string) *ExportUserDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExportUserDevicesResponseBody) SetSignedUrl(v string) *ExportUserDevicesResponseBody {
	s.SignedUrl = &v
	return s
}

type ExportUserDevicesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExportUserDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExportUserDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ExportUserDevicesResponse) GoString() string {
	return s.String()
}

func (s *ExportUserDevicesResponse) SetHeaders(v map[string]*string) *ExportUserDevicesResponse {
	s.Headers = v
	return s
}

func (s *ExportUserDevicesResponse) SetStatusCode(v int32) *ExportUserDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ExportUserDevicesResponse) SetBody(v *ExportUserDevicesResponseBody) *ExportUserDevicesResponse {
	s.Body = v
	return s
}

type GetActiveIdpConfigResponseBody struct {
	Data *GetActiveIdpConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetActiveIdpConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetActiveIdpConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetActiveIdpConfigResponseBody) SetData(v *GetActiveIdpConfigResponseBodyData) *GetActiveIdpConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetActiveIdpConfigResponseBody) SetRequestId(v string) *GetActiveIdpConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetActiveIdpConfigResponseBodyData struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// idp-cfg001
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DingTalk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetActiveIdpConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetActiveIdpConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetActiveIdpConfigResponseBodyData) SetDescription(v string) *GetActiveIdpConfigResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetActiveIdpConfigResponseBodyData) SetId(v string) *GetActiveIdpConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetActiveIdpConfigResponseBodyData) SetName(v string) *GetActiveIdpConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetActiveIdpConfigResponseBodyData) SetType(v string) *GetActiveIdpConfigResponseBodyData {
	s.Type = &v
	return s
}

type GetActiveIdpConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetActiveIdpConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetActiveIdpConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetActiveIdpConfigResponse) GoString() string {
	return s.String()
}

func (s *GetActiveIdpConfigResponse) SetHeaders(v map[string]*string) *GetActiveIdpConfigResponse {
	s.Headers = v
	return s
}

func (s *GetActiveIdpConfigResponse) SetStatusCode(v int32) *GetActiveIdpConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetActiveIdpConfigResponse) SetBody(v *GetActiveIdpConfigResponseBody) *GetActiveIdpConfigResponse {
	s.Body = v
	return s
}

type GetClientUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 598
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	// This parameter is required.
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetClientUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClientUserRequest) GoString() string {
	return s.String()
}

func (s *GetClientUserRequest) SetIdpConfigId(v string) *GetClientUserRequest {
	s.IdpConfigId = &v
	return s
}

func (s *GetClientUserRequest) SetUsername(v string) *GetClientUserRequest {
	s.Username = &v
	return s
}

type GetClientUserResponseBody struct {
	Data *GetClientUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClientUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClientUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientUserResponseBody) SetData(v *GetClientUserResponseBodyData) *GetClientUserResponseBody {
	s.Data = v
	return s
}

func (s *GetClientUserResponseBody) SetRequestId(v string) *GetClientUserResponseBody {
	s.RequestId = &v
	return s
}

type GetClientUserResponseBodyData struct {
	Department *GetClientUserResponseBodyDataDepartment `json:"Department,omitempty" xml:"Department,omitempty" type:"Struct"`
	// example:
	//
	// 10713
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// johndoe@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 83
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 598
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	// example:
	//
	// 13641966835
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	// example:
	//
	// Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// su_abcd7215****
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetClientUserResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetClientUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetClientUserResponseBodyData) SetDepartment(v *GetClientUserResponseBodyDataDepartment) *GetClientUserResponseBodyData {
	s.Department = v
	return s
}

func (s *GetClientUserResponseBodyData) SetDepartmentId(v string) *GetClientUserResponseBodyData {
	s.DepartmentId = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetDescription(v string) *GetClientUserResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetEmail(v string) *GetClientUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetId(v string) *GetClientUserResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetIdpConfigId(v string) *GetClientUserResponseBodyData {
	s.IdpConfigId = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetMobileNumber(v string) *GetClientUserResponseBodyData {
	s.MobileNumber = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetStatus(v string) *GetClientUserResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetUserId(v string) *GetClientUserResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetClientUserResponseBodyData) SetUsername(v string) *GetClientUserResponseBodyData {
	s.Username = &v
	return s
}

type GetClientUserResponseBodyDataDepartment struct {
	// example:
	//
	// 107
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetClientUserResponseBodyDataDepartment) String() string {
	return tea.Prettify(s)
}

func (s GetClientUserResponseBodyDataDepartment) GoString() string {
	return s.String()
}

func (s *GetClientUserResponseBodyDataDepartment) SetId(v string) *GetClientUserResponseBodyDataDepartment {
	s.Id = &v
	return s
}

func (s *GetClientUserResponseBodyDataDepartment) SetName(v string) *GetClientUserResponseBodyDataDepartment {
	s.Name = &v
	return s
}

type GetClientUserResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClientUserResponse) GoString() string {
	return s.String()
}

func (s *GetClientUserResponse) SetHeaders(v map[string]*string) *GetClientUserResponse {
	s.Headers = v
	return s
}

func (s *GetClientUserResponse) SetStatusCode(v int32) *GetClientUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientUserResponse) SetBody(v *GetClientUserResponseBody) *GetClientUserResponse {
	s.Body = v
	return s
}

type GetDynamicRouteRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dr-16ff07c8207d****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
}

func (s GetDynamicRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDynamicRouteRequest) GoString() string {
	return s.String()
}

func (s *GetDynamicRouteRequest) SetDynamicRouteId(v string) *GetDynamicRouteRequest {
	s.DynamicRouteId = &v
	return s
}

type GetDynamicRouteResponseBody struct {
	DynamicRoute *GetDynamicRouteResponseBodyDynamicRoute `json:"DynamicRoute,omitempty" xml:"DynamicRoute,omitempty" type:"Struct"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDynamicRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDynamicRouteResponseBody) GoString() string {
	return s.String()
}

func (s *GetDynamicRouteResponseBody) SetDynamicRoute(v *GetDynamicRouteResponseBodyDynamicRoute) *GetDynamicRouteResponseBody {
	s.DynamicRoute = v
	return s
}

func (s *GetDynamicRouteResponseBody) SetRequestId(v string) *GetDynamicRouteResponseBody {
	s.RequestId = &v
	return s
}

type GetDynamicRouteResponseBodyDynamicRoute struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// example:
	//
	// 2023-02-09 10:31:47
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// dr-16ff07c8207d****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	// example:
	//
	// connector
	DynamicRouteType *string `json:"DynamicRouteType,omitempty" xml:"DynamicRouteType,omitempty"`
	// example:
	//
	// dynamic_route_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// connector-8ccb13b6f52c****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// example:
	//
	// 1
	Priority  *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s GetDynamicRouteResponseBodyDynamicRoute) String() string {
	return tea.Prettify(s)
}

func (s GetDynamicRouteResponseBodyDynamicRoute) GoString() string {
	return s.String()
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetApplicationIds(v []*string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.ApplicationIds = v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetApplicationType(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.ApplicationType = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetCreateTime(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.CreateTime = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetDescription(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.Description = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetDynamicRouteId(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.DynamicRouteId = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetDynamicRouteType(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.DynamicRouteType = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetName(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.Name = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetNextHop(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.NextHop = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetPriority(v int32) *GetDynamicRouteResponseBodyDynamicRoute {
	s.Priority = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetRegionIds(v []*string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.RegionIds = v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetStatus(v string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.Status = &v
	return s
}

func (s *GetDynamicRouteResponseBodyDynamicRoute) SetTagIds(v []*string) *GetDynamicRouteResponseBodyDynamicRoute {
	s.TagIds = v
	return s
}

type GetDynamicRouteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDynamicRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDynamicRouteResponse) GoString() string {
	return s.String()
}

func (s *GetDynamicRouteResponse) SetHeaders(v map[string]*string) *GetDynamicRouteResponse {
	s.Headers = v
	return s
}

func (s *GetDynamicRouteResponse) SetStatusCode(v int32) *GetDynamicRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDynamicRouteResponse) SetBody(v *GetDynamicRouteResponseBody) *GetDynamicRouteResponse {
	s.Body = v
	return s
}

type GetIdpConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1465
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetIdpConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIdpConfigRequest) GoString() string {
	return s.String()
}

func (s *GetIdpConfigRequest) SetId(v string) *GetIdpConfigRequest {
	s.Id = &v
	return s
}

type GetIdpConfigResponseBody struct {
	Data *GetIdpConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIdpConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIdpConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdpConfigResponseBody) SetData(v *GetIdpConfigResponseBodyData) *GetIdpConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetIdpConfigResponseBody) SetRequestId(v string) *GetIdpConfigResponseBody {
	s.RequestId = &v
	return s
}

type GetIdpConfigResponseBodyData struct {
	// AccessKey ID
	//
	// example:
	//
	// LTAI5tJVztnh6Nn***
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// AccessKey Secret
	//
	// example:
	//
	// E75ktr5jENiR3ssjC***
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// https://172.10.10.2:4321/getGroup?name=%s&pass=%s
	GetGroupUrl *string `json:"GetGroupUrl,omitempty" xml:"GetGroupUrl,omitempty"`
	// example:
	//
	// 1465
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// <?xml version="1.0" encoding="utf-8"?>***
	IdpMetadata *string `json:"IdpMetadata,omitempty" xml:"IdpMetadata,omitempty"`
	// example:
	//
	// totp
	MfaConfigType *string `json:"MfaConfigType,omitempty" xml:"MfaConfigType,omitempty"`
	// example:
	//
	// password
	MobileLoginType *string `json:"MobileLoginType,omitempty" xml:"MobileLoginType,omitempty"`
	// example:
	//
	// totp
	MobileMfaConfigType *string `json:"MobileMfaConfigType,omitempty" xml:"MobileMfaConfigType,omitempty"`
	// example:
	//
	// 1482,1355
	MultiIdpInfo *string `json:"MultiIdpInfo,omitempty" xml:"MultiIdpInfo,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// password
	PcLoginType *string `json:"PcLoginType,omitempty" xml:"PcLoginType,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// CSAS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2024-02-26T02:02:42Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// QVhaU0RDR0JIWVV***
	VerifyAesKey *string `json:"VerifyAesKey,omitempty" xml:"VerifyAesKey,omitempty"`
	// example:
	//
	// 7JAr3fYtnl***
	VerifyToken *string `json:"VerifyToken,omitempty" xml:"VerifyToken,omitempty"`
	// example:
	//
	// http://172.10.10.1:1234/otp_verify
	VerifyUrl *string `json:"VerifyUrl,omitempty" xml:"VerifyUrl,omitempty"`
}

func (s GetIdpConfigResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetIdpConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIdpConfigResponseBodyData) SetAccessKey(v string) *GetIdpConfigResponseBodyData {
	s.AccessKey = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetAccessKeySecret(v string) *GetIdpConfigResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetDescription(v string) *GetIdpConfigResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetGetGroupUrl(v string) *GetIdpConfigResponseBodyData {
	s.GetGroupUrl = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetId(v string) *GetIdpConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetIdpMetadata(v string) *GetIdpConfigResponseBodyData {
	s.IdpMetadata = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetMfaConfigType(v string) *GetIdpConfigResponseBodyData {
	s.MfaConfigType = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetMobileLoginType(v string) *GetIdpConfigResponseBodyData {
	s.MobileLoginType = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetMobileMfaConfigType(v string) *GetIdpConfigResponseBodyData {
	s.MobileMfaConfigType = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetMultiIdpInfo(v string) *GetIdpConfigResponseBodyData {
	s.MultiIdpInfo = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetName(v string) *GetIdpConfigResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetPcLoginType(v string) *GetIdpConfigResponseBodyData {
	s.PcLoginType = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetStatus(v string) *GetIdpConfigResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetType(v string) *GetIdpConfigResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetUpdateTime(v string) *GetIdpConfigResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetVerifyAesKey(v string) *GetIdpConfigResponseBodyData {
	s.VerifyAesKey = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetVerifyToken(v string) *GetIdpConfigResponseBodyData {
	s.VerifyToken = &v
	return s
}

func (s *GetIdpConfigResponseBodyData) SetVerifyUrl(v string) *GetIdpConfigResponseBodyData {
	s.VerifyUrl = &v
	return s
}

type GetIdpConfigResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdpConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdpConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIdpConfigResponse) GoString() string {
	return s.String()
}

func (s *GetIdpConfigResponse) SetHeaders(v map[string]*string) *GetIdpConfigResponse {
	s.Headers = v
	return s
}

func (s *GetIdpConfigResponse) SetStatusCode(v int32) *GetIdpConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdpConfigResponse) SetBody(v *GetIdpConfigResponseBody) *GetIdpConfigResponse {
	s.Body = v
	return s
}

type GetPrivateAccessApplicationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s GetPrivateAccessApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessApplicationRequest) SetApplicationId(v string) *GetPrivateAccessApplicationRequest {
	s.ApplicationId = &v
	return s
}

type GetPrivateAccessApplicationResponseBody struct {
	Application *GetPrivateAccessApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// example:
	//
	// 3ACC5EDC-2B7D-5032-8C58-D7615D66C1D4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPrivateAccessApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessApplicationResponseBody) SetApplication(v *GetPrivateAccessApplicationResponseBodyApplication) *GetPrivateAccessApplicationResponseBody {
	s.Application = v
	return s
}

func (s *GetPrivateAccessApplicationResponseBody) SetRequestId(v string) *GetPrivateAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

type GetPrivateAccessApplicationResponseBodyApplication struct {
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId       *string   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	BrowserAccessStatus *string   `json:"BrowserAccessStatus,omitempty" xml:"BrowserAccessStatus,omitempty"`
	ConnectorIds        []*string `json:"ConnectorIds,omitempty" xml:"ConnectorIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-08-30 16:50:32
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	L7ProxyDomainAutomatic *string `json:"L7ProxyDomainAutomatic,omitempty" xml:"L7ProxyDomainAutomatic,omitempty"`
	L7ProxyDomainCustom    *string `json:"L7ProxyDomainCustom,omitempty" xml:"L7ProxyDomainCustom,omitempty"`
	// example:
	//
	// private_access_application_name
	Name       *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyIds  []*string                                                       `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	PortRanges []*GetPrivateAccessApplicationResponseBodyApplicationPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// example:
	//
	// All
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// Enabled
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s GetPrivateAccessApplicationResponseBodyApplication) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetAddresses(v []*string) *GetPrivateAccessApplicationResponseBodyApplication {
	s.Addresses = v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetApplicationId(v string) *GetPrivateAccessApplicationResponseBodyApplication {
	s.ApplicationId = &v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetBrowserAccessStatus(v string) *GetPrivateAccessApplicationResponseBodyApplication {
	s.BrowserAccessStatus = &v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetConnectorIds(v []*string) *GetPrivateAccessApplicationResponseBodyApplication {
	s.ConnectorIds = v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetCreateTime(v string) *GetPrivateAccessApplicationResponseBodyApplication {
	s.CreateTime = &v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetDescription(v string) *GetPrivateAccessApplicationResponseBodyApplication {
	s.Description = &v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetL7ProxyDomainAutomatic(v string) *GetPrivateAccessApplicationResponseBodyApplication {
	s.L7ProxyDomainAutomatic = &v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetL7ProxyDomainCustom(v string) *GetPrivateAccessApplicationResponseBodyApplication {
	s.L7ProxyDomainCustom = &v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetName(v string) *GetPrivateAccessApplicationResponseBodyApplication {
	s.Name = &v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetPolicyIds(v []*string) *GetPrivateAccessApplicationResponseBodyApplication {
	s.PolicyIds = v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetPortRanges(v []*GetPrivateAccessApplicationResponseBodyApplicationPortRanges) *GetPrivateAccessApplicationResponseBodyApplication {
	s.PortRanges = v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetProtocol(v string) *GetPrivateAccessApplicationResponseBodyApplication {
	s.Protocol = &v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetStatus(v string) *GetPrivateAccessApplicationResponseBodyApplication {
	s.Status = &v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplication) SetTagIds(v []*string) *GetPrivateAccessApplicationResponseBodyApplication {
	s.TagIds = v
	return s
}

type GetPrivateAccessApplicationResponseBodyApplicationPortRanges struct {
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 81
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s GetPrivateAccessApplicationResponseBodyApplicationPortRanges) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessApplicationResponseBodyApplicationPortRanges) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessApplicationResponseBodyApplicationPortRanges) SetBegin(v int32) *GetPrivateAccessApplicationResponseBodyApplicationPortRanges {
	s.Begin = &v
	return s
}

func (s *GetPrivateAccessApplicationResponseBodyApplicationPortRanges) SetEnd(v int32) *GetPrivateAccessApplicationResponseBodyApplicationPortRanges {
	s.End = &v
	return s
}

type GetPrivateAccessApplicationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrivateAccessApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessApplicationResponse) SetHeaders(v map[string]*string) *GetPrivateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetPrivateAccessApplicationResponse) SetStatusCode(v int32) *GetPrivateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrivateAccessApplicationResponse) SetBody(v *GetPrivateAccessApplicationResponseBody) *GetPrivateAccessApplicationResponse {
	s.Body = v
	return s
}

type GetPrivateAccessPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-policy-63b2f1844b86****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s GetPrivateAccessPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessPolicyRequest) SetPolicyId(v string) *GetPrivateAccessPolicyRequest {
	s.PolicyId = &v
	return s
}

type GetPrivateAccessPolicyResponseBody struct {
	Policy *GetPrivateAccessPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// example:
	//
	// 7E9D7ACD-53D5-56EF-A913-79D148D06299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPrivateAccessPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessPolicyResponseBody) SetPolicy(v *GetPrivateAccessPolicyResponseBodyPolicy) *GetPrivateAccessPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *GetPrivateAccessPolicyResponseBody) SetRequestId(v string) *GetPrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

type GetPrivateAccessPolicyResponseBodyPolicy struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// example:
	//
	// 2021-07-29 11:26:02
	CreateTime            *string                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomUserAttributes  []*GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	Description           *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceAttributeAction *string                                                         `json:"DeviceAttributeAction,omitempty" xml:"DeviceAttributeAction,omitempty"`
	DeviceAttributeId     *string                                                         `json:"DeviceAttributeId,omitempty" xml:"DeviceAttributeId,omitempty"`
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// example:
	//
	// pa-policy-63b2f1844b86****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds       []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// Normal
	UserGroupMode *string `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
}

func (s GetPrivateAccessPolicyResponseBodyPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetApplicationIds(v []*string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.ApplicationIds = v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetApplicationType(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.ApplicationType = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetCreateTime(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.CreateTime = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetCustomUserAttributes(v []*GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.CustomUserAttributes = v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetDescription(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetDeviceAttributeAction(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.DeviceAttributeAction = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetDeviceAttributeId(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.DeviceAttributeId = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetName(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.Name = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetPolicyAction(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.PolicyAction = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetPolicyId(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetPriority(v int32) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.Priority = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetStatus(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.Status = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetTagIds(v []*string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.TagIds = v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetUserGroupIds(v []*string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.UserGroupIds = v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicy) SetUserGroupMode(v string) *GetPrivateAccessPolicyResponseBodyPolicy {
	s.UserGroupMode = &v
	return s
}

type GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) SetIdpId(v int32) *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) SetRelation(v string) *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) SetUserGroupType(v string) *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes) SetValue(v string) *GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes {
	s.Value = &v
	return s
}

type GetPrivateAccessPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPrivateAccessPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *GetPrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetPrivateAccessPolicyResponse) SetStatusCode(v int32) *GetPrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPrivateAccessPolicyResponse) SetBody(v *GetPrivateAccessPolicyResponseBody) *GetPrivateAccessPolicyResponse {
	s.Body = v
	return s
}

type GetRegistrationPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// reg-policy-dcbfd33cb004****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s GetRegistrationPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRegistrationPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetRegistrationPolicyRequest) SetPolicyId(v string) *GetRegistrationPolicyRequest {
	s.PolicyId = &v
	return s
}

type GetRegistrationPolicyResponseBody struct {
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime  *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail []*GetRegistrationPolicyResponseBodyLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// example:
	//
	// registration_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// reg-policy-dcbfd33cb004****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 99
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 47363C2B-1AAA-5954-8847-0E50FCC54117
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s GetRegistrationPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRegistrationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegistrationPolicyResponseBody) SetCreateTime(v string) *GetRegistrationPolicyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetDescription(v string) *GetRegistrationPolicyResponseBody {
	s.Description = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetLimitDetail(v []*GetRegistrationPolicyResponseBodyLimitDetail) *GetRegistrationPolicyResponseBody {
	s.LimitDetail = v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetMatchMode(v string) *GetRegistrationPolicyResponseBody {
	s.MatchMode = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetName(v string) *GetRegistrationPolicyResponseBody {
	s.Name = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetPolicyId(v string) *GetRegistrationPolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetPriority(v int64) *GetRegistrationPolicyResponseBody {
	s.Priority = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetRequestId(v string) *GetRegistrationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetStatus(v string) *GetRegistrationPolicyResponseBody {
	s.Status = &v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetUserGroupIds(v []*string) *GetRegistrationPolicyResponseBody {
	s.UserGroupIds = v
	return s
}

func (s *GetRegistrationPolicyResponseBody) SetWhitelist(v []*string) *GetRegistrationPolicyResponseBody {
	s.Whitelist = v
	return s
}

type GetRegistrationPolicyResponseBodyLimitDetail struct {
	// example:
	//
	// Personal
	DeviceBelong *string                                                 `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *GetRegistrationPolicyResponseBodyLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	// example:
	//
	// LimitDiff
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s GetRegistrationPolicyResponseBodyLimitDetail) String() string {
	return tea.Prettify(s)
}

func (s GetRegistrationPolicyResponseBodyLimitDetail) GoString() string {
	return s.String()
}

func (s *GetRegistrationPolicyResponseBodyLimitDetail) SetDeviceBelong(v string) *GetRegistrationPolicyResponseBodyLimitDetail {
	s.DeviceBelong = &v
	return s
}

func (s *GetRegistrationPolicyResponseBodyLimitDetail) SetLimitCount(v *GetRegistrationPolicyResponseBodyLimitDetailLimitCount) *GetRegistrationPolicyResponseBodyLimitDetail {
	s.LimitCount = v
	return s
}

func (s *GetRegistrationPolicyResponseBodyLimitDetail) SetLimitType(v string) *GetRegistrationPolicyResponseBodyLimitDetail {
	s.LimitType = &v
	return s
}

type GetRegistrationPolicyResponseBodyLimitDetailLimitCount struct {
	// example:
	//
	// 0
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 2
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 2
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s GetRegistrationPolicyResponseBodyLimitDetailLimitCount) String() string {
	return tea.Prettify(s)
}

func (s GetRegistrationPolicyResponseBodyLimitDetailLimitCount) GoString() string {
	return s.String()
}

func (s *GetRegistrationPolicyResponseBodyLimitDetailLimitCount) SetAll(v int32) *GetRegistrationPolicyResponseBodyLimitDetailLimitCount {
	s.All = &v
	return s
}

func (s *GetRegistrationPolicyResponseBodyLimitDetailLimitCount) SetMobile(v int32) *GetRegistrationPolicyResponseBodyLimitDetailLimitCount {
	s.Mobile = &v
	return s
}

func (s *GetRegistrationPolicyResponseBodyLimitDetailLimitCount) SetPC(v int32) *GetRegistrationPolicyResponseBodyLimitDetailLimitCount {
	s.PC = &v
	return s
}

type GetRegistrationPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegistrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegistrationPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRegistrationPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetRegistrationPolicyResponse) SetHeaders(v map[string]*string) *GetRegistrationPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetRegistrationPolicyResponse) SetStatusCode(v int32) *GetRegistrationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegistrationPolicyResponse) SetBody(v *GetRegistrationPolicyResponseBody) *GetRegistrationPolicyResponse {
	s.Body = v
	return s
}

type GetUserDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
}

func (s GetUserDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserDeviceRequest) GoString() string {
	return s.String()
}

func (s *GetUserDeviceRequest) SetDeviceTag(v string) *GetUserDeviceRequest {
	s.DeviceTag = &v
	return s
}

type GetUserDeviceResponseBody struct {
	Device *GetUserDeviceResponseBodyDevice `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	// example:
	//
	// EFE7EBB2-449D-5BBB-B381-CA7839BC1649
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserDeviceResponseBody) SetDevice(v *GetUserDeviceResponseBodyDevice) *GetUserDeviceResponseBody {
	s.Device = v
	return s
}

func (s *GetUserDeviceResponseBody) SetRequestId(v string) *GetUserDeviceResponseBody {
	s.RequestId = &v
	return s
}

type GetUserDeviceResponseBodyDevice struct {
	// example:
	//
	// Online
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// example:
	//
	// 2.2.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// Apple M1
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// Company
	DeviceBelong *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	// example:
	//
	// MacBookPro17,1
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	// example:
	//
	// Online
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// example:
	//
	// Windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 3.5.1
	DeviceVersion *string `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	// example:
	//
	// APPLE SSD AP0512Q Media
	Disk *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// example:
	//
	// Unauthorized
	DlpStatus    *string                                        `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	HistoryUsers []*GetUserDeviceResponseBodyDeviceHistoryUsers `json:"HistoryUsers,omitempty" xml:"HistoryUsers,omitempty" type:"Repeated"`
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// Disabled
	IaStatus *string `json:"IaStatus,omitempty" xml:"IaStatus,omitempty"`
	// example:
	//
	// 172.16.XX.XX
	InnerIP *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	// example:
	//
	// 48:9e:XX:XX:02:80
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 16
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// Unprovisioned
	NacStatus *string `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	// example:
	//
	// Enabled
	PaStatus *string `json:"PaStatus,omitempty" xml:"PaStatus,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// example:
	//
	// true
	SharingStatus *bool `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	// example:
	//
	// 106.14.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// example:
	//
	// 2023-08-24 19:04:42
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetUserDeviceResponseBodyDevice) String() string {
	return tea.Prettify(s)
}

func (s GetUserDeviceResponseBodyDevice) GoString() string {
	return s.String()
}

func (s *GetUserDeviceResponseBodyDevice) SetAppStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.AppStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetAppVersion(v string) *GetUserDeviceResponseBodyDevice {
	s.AppVersion = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetCPU(v string) *GetUserDeviceResponseBodyDevice {
	s.CPU = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetCreateTime(v string) *GetUserDeviceResponseBodyDevice {
	s.CreateTime = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDepartment(v string) *GetUserDeviceResponseBodyDevice {
	s.Department = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDeviceBelong(v string) *GetUserDeviceResponseBodyDevice {
	s.DeviceBelong = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDeviceModel(v string) *GetUserDeviceResponseBodyDevice {
	s.DeviceModel = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDeviceStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.DeviceStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDeviceTag(v string) *GetUserDeviceResponseBodyDevice {
	s.DeviceTag = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDeviceType(v string) *GetUserDeviceResponseBodyDevice {
	s.DeviceType = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDeviceVersion(v string) *GetUserDeviceResponseBodyDevice {
	s.DeviceVersion = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDisk(v string) *GetUserDeviceResponseBodyDevice {
	s.Disk = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetDlpStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.DlpStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetHistoryUsers(v []*GetUserDeviceResponseBodyDeviceHistoryUsers) *GetUserDeviceResponseBodyDevice {
	s.HistoryUsers = v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetHostname(v string) *GetUserDeviceResponseBodyDevice {
	s.Hostname = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetIaStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.IaStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetInnerIP(v string) *GetUserDeviceResponseBodyDevice {
	s.InnerIP = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetMac(v string) *GetUserDeviceResponseBodyDevice {
	s.Mac = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetMemory(v string) *GetUserDeviceResponseBodyDevice {
	s.Memory = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetNacStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.NacStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetPaStatus(v string) *GetUserDeviceResponseBodyDevice {
	s.PaStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetSaseUserId(v string) *GetUserDeviceResponseBodyDevice {
	s.SaseUserId = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetSharingStatus(v bool) *GetUserDeviceResponseBodyDevice {
	s.SharingStatus = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetSrcIP(v string) *GetUserDeviceResponseBodyDevice {
	s.SrcIP = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetUpdateTime(v string) *GetUserDeviceResponseBodyDevice {
	s.UpdateTime = &v
	return s
}

func (s *GetUserDeviceResponseBodyDevice) SetUsername(v string) *GetUserDeviceResponseBodyDevice {
	s.Username = &v
	return s
}

type GetUserDeviceResponseBodyDeviceHistoryUsers struct {
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s GetUserDeviceResponseBodyDeviceHistoryUsers) String() string {
	return tea.Prettify(s)
}

func (s GetUserDeviceResponseBodyDeviceHistoryUsers) GoString() string {
	return s.String()
}

func (s *GetUserDeviceResponseBodyDeviceHistoryUsers) SetSaseUserId(v string) *GetUserDeviceResponseBodyDeviceHistoryUsers {
	s.SaseUserId = &v
	return s
}

func (s *GetUserDeviceResponseBodyDeviceHistoryUsers) SetUsername(v string) *GetUserDeviceResponseBodyDeviceHistoryUsers {
	s.Username = &v
	return s
}

type GetUserDeviceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserDeviceResponse) GoString() string {
	return s.String()
}

func (s *GetUserDeviceResponse) SetHeaders(v map[string]*string) *GetUserDeviceResponse {
	s.Headers = v
	return s
}

func (s *GetUserDeviceResponse) SetStatusCode(v int32) *GetUserDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserDeviceResponse) SetBody(v *GetUserDeviceResponseBody) *GetUserDeviceResponse {
	s.Body = v
	return s
}

type GetUserGroupRequest struct {
	// This parameter is required.
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s GetUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupRequest) GoString() string {
	return s.String()
}

func (s *GetUserGroupRequest) SetUserGroupId(v string) *GetUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type GetUserGroupResponseBody struct {
	// example:
	//
	// 1310DBC7-7E1F-55D3-B4B4-E4BE912517FB
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserGroup *GetUserGroupResponseBodyUserGroup `json:"UserGroup,omitempty" xml:"UserGroup,omitempty" type:"Struct"`
}

func (s GetUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBody) SetRequestId(v string) *GetUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserGroupResponseBody) SetUserGroup(v *GetUserGroupResponseBodyUserGroup) *GetUserGroupResponseBody {
	s.UserGroup = v
	return s
}

type GetUserGroupResponseBodyUserGroup struct {
	Attributes []*GetUserGroupResponseBodyUserGroupAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-10-10 11:39:22
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// user_group_name
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s GetUserGroupResponseBodyUserGroup) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupResponseBodyUserGroup) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBodyUserGroup) SetAttributes(v []*GetUserGroupResponseBodyUserGroupAttributes) *GetUserGroupResponseBodyUserGroup {
	s.Attributes = v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) SetCreateTime(v string) *GetUserGroupResponseBodyUserGroup {
	s.CreateTime = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) SetDescription(v string) *GetUserGroupResponseBodyUserGroup {
	s.Description = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) SetName(v string) *GetUserGroupResponseBodyUserGroup {
	s.Name = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroup) SetUserGroupId(v string) *GetUserGroupResponseBodyUserGroup {
	s.UserGroupId = &v
	return s
}

type GetUserGroupResponseBodyUserGroupAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetUserGroupResponseBodyUserGroupAttributes) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupResponseBodyUserGroupAttributes) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponseBodyUserGroupAttributes) SetIdpId(v int32) *GetUserGroupResponseBodyUserGroupAttributes {
	s.IdpId = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupAttributes) SetRelation(v string) *GetUserGroupResponseBodyUserGroupAttributes {
	s.Relation = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupAttributes) SetUserGroupType(v string) *GetUserGroupResponseBodyUserGroupAttributes {
	s.UserGroupType = &v
	return s
}

func (s *GetUserGroupResponseBodyUserGroupAttributes) SetValue(v string) *GetUserGroupResponseBodyUserGroupAttributes {
	s.Value = &v
	return s
}

type GetUserGroupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserGroupResponse) GoString() string {
	return s.String()
}

func (s *GetUserGroupResponse) SetHeaders(v map[string]*string) *GetUserGroupResponse {
	s.Headers = v
	return s
}

func (s *GetUserGroupResponse) SetStatusCode(v int32) *GetUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserGroupResponse) SetBody(v *GetUserGroupResponseBody) *GetUserGroupResponse {
	s.Body = v
	return s
}

type GetWmEmbedTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// job:5GfrJYsoaffmCE7Z5bZtjU********
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetWmEmbedTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWmEmbedTaskRequest) GoString() string {
	return s.String()
}

func (s *GetWmEmbedTaskRequest) SetTaskId(v string) *GetWmEmbedTaskRequest {
	s.TaskId = &v
	return s
}

type GetWmEmbedTaskResponseBody struct {
	Data *GetWmEmbedTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWmEmbedTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWmEmbedTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetWmEmbedTaskResponseBody) SetData(v *GetWmEmbedTaskResponseBodyData) *GetWmEmbedTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetWmEmbedTaskResponseBody) SetRequestId(v string) *GetWmEmbedTaskResponseBody {
	s.RequestId = &v
	return s
}

type GetWmEmbedTaskResponseBodyData struct {
	// example:
	//
	// https://example.com/embed-****.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// 171859****
	FileUrlExp *string `json:"FileUrlExp,omitempty" xml:"FileUrlExp,omitempty"`
	// example:
	//
	// embed-****.pdf
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// example:
	//
	// d41d8cd98f00b204e9800998ecf8****
	OutFileHashMd5 *string `json:"OutFileHashMd5,omitempty" xml:"OutFileHashMd5,omitempty"`
	// example:
	//
	// 123**
	OutFileSize *int64 `json:"OutFileSize,omitempty" xml:"OutFileSize,omitempty"`
	// example:
	//
	// job:5GfrJYsoaffmCE7Z5bZtjUxxxxxxxxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// Success
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetWmEmbedTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWmEmbedTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWmEmbedTaskResponseBodyData) SetFileUrl(v string) *GetWmEmbedTaskResponseBodyData {
	s.FileUrl = &v
	return s
}

func (s *GetWmEmbedTaskResponseBodyData) SetFileUrlExp(v string) *GetWmEmbedTaskResponseBodyData {
	s.FileUrlExp = &v
	return s
}

func (s *GetWmEmbedTaskResponseBodyData) SetFilename(v string) *GetWmEmbedTaskResponseBodyData {
	s.Filename = &v
	return s
}

func (s *GetWmEmbedTaskResponseBodyData) SetOutFileHashMd5(v string) *GetWmEmbedTaskResponseBodyData {
	s.OutFileHashMd5 = &v
	return s
}

func (s *GetWmEmbedTaskResponseBodyData) SetOutFileSize(v int64) *GetWmEmbedTaskResponseBodyData {
	s.OutFileSize = &v
	return s
}

func (s *GetWmEmbedTaskResponseBodyData) SetTaskId(v string) *GetWmEmbedTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetWmEmbedTaskResponseBodyData) SetTaskStatus(v string) *GetWmEmbedTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

type GetWmEmbedTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWmEmbedTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWmEmbedTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWmEmbedTaskResponse) GoString() string {
	return s.String()
}

func (s *GetWmEmbedTaskResponse) SetHeaders(v map[string]*string) *GetWmEmbedTaskResponse {
	s.Headers = v
	return s
}

func (s *GetWmEmbedTaskResponse) SetStatusCode(v int32) *GetWmEmbedTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWmEmbedTaskResponse) SetBody(v *GetWmEmbedTaskResponseBody) *GetWmEmbedTaskResponse {
	s.Body = v
	return s
}

type GetWmExtractTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wmt-9648c22d2eb2cb57bb855dcae7898464********
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetWmExtractTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWmExtractTaskRequest) GoString() string {
	return s.String()
}

func (s *GetWmExtractTaskRequest) SetTaskId(v string) *GetWmExtractTaskRequest {
	s.TaskId = &v
	return s
}

type GetWmExtractTaskResponseBody struct {
	Data *GetWmExtractTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWmExtractTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWmExtractTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetWmExtractTaskResponseBody) SetData(v *GetWmExtractTaskResponseBodyData) *GetWmExtractTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetWmExtractTaskResponseBody) SetRequestId(v string) *GetWmExtractTaskResponseBody {
	s.RequestId = &v
	return s
}

type GetWmExtractTaskResponseBodyData struct {
	// example:
	//
	// 2024-01-01 11:22:33
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// test-****.pdf
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// wmt-9648c22d2eb2cb57bb855dcae7898464********
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// aGVsbG8gc2Fz****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// example:
	//
	// 123**
	WmInfoUint *int64 `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
	// example:
	//
	// PureDocument
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s GetWmExtractTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetWmExtractTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWmExtractTaskResponseBodyData) SetCreateTime(v string) *GetWmExtractTaskResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetFilename(v string) *GetWmExtractTaskResponseBodyData {
	s.Filename = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetStatus(v string) *GetWmExtractTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetTaskId(v string) *GetWmExtractTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetWmInfoBytesB64(v string) *GetWmExtractTaskResponseBodyData {
	s.WmInfoBytesB64 = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetWmInfoSize(v int64) *GetWmExtractTaskResponseBodyData {
	s.WmInfoSize = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetWmInfoUint(v int64) *GetWmExtractTaskResponseBodyData {
	s.WmInfoUint = &v
	return s
}

func (s *GetWmExtractTaskResponseBodyData) SetWmType(v string) *GetWmExtractTaskResponseBodyData {
	s.WmType = &v
	return s
}

type GetWmExtractTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWmExtractTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWmExtractTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWmExtractTaskResponse) GoString() string {
	return s.String()
}

func (s *GetWmExtractTaskResponse) SetHeaders(v map[string]*string) *GetWmExtractTaskResponse {
	s.Headers = v
	return s
}

func (s *GetWmExtractTaskResponse) SetStatusCode(v int32) *GetWmExtractTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWmExtractTaskResponse) SetBody(v *GetWmExtractTaskResponseBody) *GetWmExtractTaskResponse {
	s.Body = v
	return s
}

type ListApplicationsForPrivateAccessPolicyRequest struct {
	// This parameter is required.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s ListApplicationsForPrivateAccessPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForPrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessPolicyRequest) SetPolicyIds(v []*string) *ListApplicationsForPrivateAccessPolicyRequest {
	s.PolicyIds = v
	return s
}

type ListApplicationsForPrivateAccessPolicyResponseBody struct {
	Polices []*ListApplicationsForPrivateAccessPolicyResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationsForPrivateAccessPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForPrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBody) SetPolices(v []*ListApplicationsForPrivateAccessPolicyResponseBodyPolices) *ListApplicationsForPrivateAccessPolicyResponseBody {
	s.Polices = v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBody) SetRequestId(v string) *ListApplicationsForPrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ListApplicationsForPrivateAccessPolicyResponseBodyPolices struct {
	Applications []*ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// pa-policy-1b0d0e8b4bcf****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s ListApplicationsForPrivateAccessPolicyResponseBodyPolices) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForPrivateAccessPolicyResponseBodyPolices) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolices) SetApplications(v []*ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) *ListApplicationsForPrivateAccessPolicyResponseBodyPolices {
	s.Applications = v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolices) SetPolicyId(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolices {
	s.PolicyId = &v
	return s
}

type ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications struct {
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// example:
	//
	// pa-application-7a9243dd02f4****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 2022-09-27 18:10:25
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// application_name
	Name       *string                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges []*ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetAddresses(v []*string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.Addresses = v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetApplicationId(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetCreateTime(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetDescription(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.Description = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetName(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.Name = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetPortRanges(v []*ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.PortRanges = v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetProtocol(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.Protocol = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications) SetStatus(v string) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplications {
	s.Status = &v
	return s
}

type ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges struct {
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 81
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges) SetBegin(v int32) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges {
	s.Begin = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges) SetEnd(v int32) *ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges {
	s.End = &v
	return s
}

type ListApplicationsForPrivateAccessPolicyResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsForPrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsForPrivateAccessPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForPrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *ListApplicationsForPrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponse) SetStatusCode(v int32) *ListApplicationsForPrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsForPrivateAccessPolicyResponse) SetBody(v *ListApplicationsForPrivateAccessPolicyResponseBody) *ListApplicationsForPrivateAccessPolicyResponse {
	s.Body = v
	return s
}

type ListApplicationsForPrivateAccessTagRequest struct {
	// This parameter is required.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s ListApplicationsForPrivateAccessTagRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForPrivateAccessTagRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessTagRequest) SetTagIds(v []*string) *ListApplicationsForPrivateAccessTagRequest {
	s.TagIds = v
	return s
}

type ListApplicationsForPrivateAccessTagResponseBody struct {
	// example:
	//
	// B608C6AE-623D-55C4-9454-601B88AE937E
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*ListApplicationsForPrivateAccessTagResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListApplicationsForPrivateAccessTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForPrivateAccessTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessTagResponseBody) SetRequestId(v string) *ListApplicationsForPrivateAccessTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBody) SetTags(v []*ListApplicationsForPrivateAccessTagResponseBodyTags) *ListApplicationsForPrivateAccessTagResponseBody {
	s.Tags = v
	return s
}

type ListApplicationsForPrivateAccessTagResponseBodyTags struct {
	Applications []*ListApplicationsForPrivateAccessTagResponseBodyTagsApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// tag-7ffc82853476****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListApplicationsForPrivateAccessTagResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForPrivateAccessTagResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTags) SetApplications(v []*ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) *ListApplicationsForPrivateAccessTagResponseBodyTags {
	s.Applications = v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTags) SetTagId(v string) *ListApplicationsForPrivateAccessTagResponseBodyTags {
	s.TagId = &v
	return s
}

type ListApplicationsForPrivateAccessTagResponseBodyTagsApplications struct {
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// example:
	//
	// pa-application-7a9243dd02f4****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 内网访问应用创建时间。
	//
	// example:
	//
	// 2022-09-27 18:10:25
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// private_access_application_name
	Name       *string                                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges []*ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// example:
	//
	// All
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetAddresses(v []*string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.Addresses = v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetApplicationId(v string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetCreateTime(v string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.CreateTime = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetDescription(v string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.Description = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetName(v string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.Name = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetPortRanges(v []*ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.PortRanges = v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetProtocol(v string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.Protocol = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications) SetStatus(v string) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplications {
	s.Status = &v
	return s
}

type ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges struct {
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 81
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges) SetBegin(v int32) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges {
	s.Begin = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges) SetEnd(v int32) *ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges {
	s.End = &v
	return s
}

type ListApplicationsForPrivateAccessTagResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsForPrivateAccessTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsForPrivateAccessTagResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsForPrivateAccessTagResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessTagResponse) SetHeaders(v map[string]*string) *ListApplicationsForPrivateAccessTagResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponse) SetStatusCode(v int32) *ListApplicationsForPrivateAccessTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponse) SetBody(v *ListApplicationsForPrivateAccessTagResponseBody) *ListApplicationsForPrivateAccessTagResponse {
	s.Body = v
	return s
}

type ListClientUsersRequest struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10785
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// example:
	//
	// johndoe@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1071
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	// example:
	//
	// 18980976559
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Enabled
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListClientUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClientUsersRequest) GoString() string {
	return s.String()
}

func (s *ListClientUsersRequest) SetCurrentPage(v int64) *ListClientUsersRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListClientUsersRequest) SetDepartmentId(v string) *ListClientUsersRequest {
	s.DepartmentId = &v
	return s
}

func (s *ListClientUsersRequest) SetEmail(v string) *ListClientUsersRequest {
	s.Email = &v
	return s
}

func (s *ListClientUsersRequest) SetIdpConfigId(v string) *ListClientUsersRequest {
	s.IdpConfigId = &v
	return s
}

func (s *ListClientUsersRequest) SetMobileNumber(v string) *ListClientUsersRequest {
	s.MobileNumber = &v
	return s
}

func (s *ListClientUsersRequest) SetPageSize(v int64) *ListClientUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListClientUsersRequest) SetStatus(v string) *ListClientUsersRequest {
	s.Status = &v
	return s
}

func (s *ListClientUsersRequest) SetUsername(v string) *ListClientUsersRequest {
	s.Username = &v
	return s
}

type ListClientUsersResponseBody struct {
	Data *ListClientUsersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// FEF1144C-95D1-5F7C-81EF-9DB70EA49FCE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClientUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClientUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientUsersResponseBody) SetData(v *ListClientUsersResponseBodyData) *ListClientUsersResponseBody {
	s.Data = v
	return s
}

func (s *ListClientUsersResponseBody) SetRequestId(v string) *ListClientUsersResponseBody {
	s.RequestId = &v
	return s
}

type ListClientUsersResponseBodyData struct {
	DataList []*ListClientUsersResponseBodyDataDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListClientUsersResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListClientUsersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClientUsersResponseBodyData) SetDataList(v []*ListClientUsersResponseBodyDataDataList) *ListClientUsersResponseBodyData {
	s.DataList = v
	return s
}

func (s *ListClientUsersResponseBodyData) SetTotalNum(v int64) *ListClientUsersResponseBodyData {
	s.TotalNum = &v
	return s
}

type ListClientUsersResponseBodyDataDataList struct {
	Department *ListClientUsersResponseBodyDataDataListDepartment `json:"Department,omitempty" xml:"Department,omitempty" type:"Struct"`
	// example:
	//
	// 10800
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// johndoe@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 1970
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1026
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	// example:
	//
	// 15800820468
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
	// example:
	//
	// Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// su_dead7216****
	UserId   *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListClientUsersResponseBodyDataDataList) String() string {
	return tea.Prettify(s)
}

func (s ListClientUsersResponseBodyDataDataList) GoString() string {
	return s.String()
}

func (s *ListClientUsersResponseBodyDataDataList) SetDepartment(v *ListClientUsersResponseBodyDataDataListDepartment) *ListClientUsersResponseBodyDataDataList {
	s.Department = v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetDepartmentId(v string) *ListClientUsersResponseBodyDataDataList {
	s.DepartmentId = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetDescription(v string) *ListClientUsersResponseBodyDataDataList {
	s.Description = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetEmail(v string) *ListClientUsersResponseBodyDataDataList {
	s.Email = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetId(v string) *ListClientUsersResponseBodyDataDataList {
	s.Id = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetIdpConfigId(v string) *ListClientUsersResponseBodyDataDataList {
	s.IdpConfigId = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetMobileNumber(v string) *ListClientUsersResponseBodyDataDataList {
	s.MobileNumber = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetStatus(v string) *ListClientUsersResponseBodyDataDataList {
	s.Status = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetUserId(v string) *ListClientUsersResponseBodyDataDataList {
	s.UserId = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataList) SetUsername(v string) *ListClientUsersResponseBodyDataDataList {
	s.Username = &v
	return s
}

type ListClientUsersResponseBodyDataDataListDepartment struct {
	// example:
	//
	// 105
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListClientUsersResponseBodyDataDataListDepartment) String() string {
	return tea.Prettify(s)
}

func (s ListClientUsersResponseBodyDataDataListDepartment) GoString() string {
	return s.String()
}

func (s *ListClientUsersResponseBodyDataDataListDepartment) SetId(v string) *ListClientUsersResponseBodyDataDataListDepartment {
	s.Id = &v
	return s
}

func (s *ListClientUsersResponseBodyDataDataListDepartment) SetName(v string) *ListClientUsersResponseBodyDataDataListDepartment {
	s.Name = &v
	return s
}

type ListClientUsersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClientUsersResponse) GoString() string {
	return s.String()
}

func (s *ListClientUsersResponse) SetHeaders(v map[string]*string) *ListClientUsersResponse {
	s.Headers = v
	return s
}

func (s *ListClientUsersResponse) SetStatusCode(v int32) *ListClientUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientUsersResponse) SetBody(v *ListClientUsersResponseBody) *ListClientUsersResponse {
	s.Body = v
	return s
}

type ListConnectorsRequest struct {
	ConnectorIds []*string `json:"ConnectorIds,omitempty" xml:"ConnectorIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// connector_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SwitchStatus *string `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
}

func (s ListConnectorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectorsRequest) SetConnectorIds(v []*string) *ListConnectorsRequest {
	s.ConnectorIds = v
	return s
}

func (s *ListConnectorsRequest) SetCurrentPage(v int32) *ListConnectorsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListConnectorsRequest) SetName(v string) *ListConnectorsRequest {
	s.Name = &v
	return s
}

func (s *ListConnectorsRequest) SetPageSize(v int32) *ListConnectorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConnectorsRequest) SetStatus(v string) *ListConnectorsRequest {
	s.Status = &v
	return s
}

func (s *ListConnectorsRequest) SetSwitchStatus(v string) *ListConnectorsRequest {
	s.SwitchStatus = &v
	return s
}

type ListConnectorsResponseBody struct {
	Connectors []*ListConnectorsResponseBodyConnectors `json:"Connectors,omitempty" xml:"Connectors,omitempty" type:"Repeated"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListConnectorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBody) SetConnectors(v []*ListConnectorsResponseBodyConnectors) *ListConnectorsResponseBody {
	s.Connectors = v
	return s
}

func (s *ListConnectorsResponseBody) SetRequestId(v string) *ListConnectorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectorsResponseBody) SetTotalNum(v int32) *ListConnectorsResponseBody {
	s.TotalNum = &v
	return s
}

type ListConnectorsResponseBodyConnectors struct {
	Applications     []*ListConnectorsResponseBodyConnectorsApplications     `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	ConnectorClients []*ListConnectorsResponseBodyConnectorsConnectorClients `json:"ConnectorClients,omitempty" xml:"ConnectorClients,omitempty" type:"Repeated"`
	// ConnectorID。
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// example:
	//
	// 2022-09-27 18:10:25
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// connector_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Enabled
	SwitchStatus *string                                          `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
	UpgradeTime  *ListConnectorsResponseBodyConnectorsUpgradeTime `json:"UpgradeTime,omitempty" xml:"UpgradeTime,omitempty" type:"Struct"`
}

func (s ListConnectorsResponseBodyConnectors) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorsResponseBodyConnectors) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyConnectors) SetApplications(v []*ListConnectorsResponseBodyConnectorsApplications) *ListConnectorsResponseBodyConnectors {
	s.Applications = v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetConnectorClients(v []*ListConnectorsResponseBodyConnectorsConnectorClients) *ListConnectorsResponseBodyConnectors {
	s.ConnectorClients = v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetConnectorId(v string) *ListConnectorsResponseBodyConnectors {
	s.ConnectorId = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetCreateTime(v string) *ListConnectorsResponseBodyConnectors {
	s.CreateTime = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetName(v string) *ListConnectorsResponseBodyConnectors {
	s.Name = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetRegionId(v string) *ListConnectorsResponseBodyConnectors {
	s.RegionId = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetStatus(v string) *ListConnectorsResponseBodyConnectors {
	s.Status = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetSwitchStatus(v string) *ListConnectorsResponseBodyConnectors {
	s.SwitchStatus = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectors) SetUpgradeTime(v *ListConnectorsResponseBodyConnectorsUpgradeTime) *ListConnectorsResponseBodyConnectors {
	s.UpgradeTime = v
	return s
}

type ListConnectorsResponseBodyConnectorsApplications struct {
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// application_name
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
}

func (s ListConnectorsResponseBodyConnectorsApplications) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorsResponseBodyConnectorsApplications) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyConnectorsApplications) SetApplicationId(v string) *ListConnectorsResponseBodyConnectorsApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsApplications) SetApplicationName(v string) *ListConnectorsResponseBodyConnectorsApplications {
	s.ApplicationName = &v
	return s
}

type ListConnectorsResponseBodyConnectorsConnectorClients struct {
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	DevTag           *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	Hostname         *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	PublicIp         *string `json:"PublicIp,omitempty" xml:"PublicIp,omitempty"`
}

func (s ListConnectorsResponseBodyConnectorsConnectorClients) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorsResponseBodyConnectorsConnectorClients) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyConnectorsConnectorClients) SetConnectionStatus(v string) *ListConnectorsResponseBodyConnectorsConnectorClients {
	s.ConnectionStatus = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsConnectorClients) SetDevTag(v string) *ListConnectorsResponseBodyConnectorsConnectorClients {
	s.DevTag = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsConnectorClients) SetHostname(v string) *ListConnectorsResponseBodyConnectorsConnectorClients {
	s.Hostname = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsConnectorClients) SetPublicIp(v string) *ListConnectorsResponseBodyConnectorsConnectorClients {
	s.PublicIp = &v
	return s
}

type ListConnectorsResponseBodyConnectorsUpgradeTime struct {
	// example:
	//
	// 23:00
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// 20:00
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s ListConnectorsResponseBodyConnectorsUpgradeTime) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorsResponseBodyConnectorsUpgradeTime) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponseBodyConnectorsUpgradeTime) SetEnd(v string) *ListConnectorsResponseBodyConnectorsUpgradeTime {
	s.End = &v
	return s
}

func (s *ListConnectorsResponseBodyConnectorsUpgradeTime) SetStart(v string) *ListConnectorsResponseBodyConnectorsUpgradeTime {
	s.Start = &v
	return s
}

type ListConnectorsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConnectorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListConnectorsResponse) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponse) SetHeaders(v map[string]*string) *ListConnectorsResponse {
	s.Headers = v
	return s
}

func (s *ListConnectorsResponse) SetStatusCode(v int32) *ListConnectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConnectorsResponse) SetBody(v *ListConnectorsResponseBody) *ListConnectorsResponse {
	s.Body = v
	return s
}

type ListDynamicRouteRegionsResponseBody struct {
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListDynamicRouteRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDynamicRouteRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDynamicRouteRegionsResponseBody) SetRegions(v []*string) *ListDynamicRouteRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListDynamicRouteRegionsResponseBody) SetRequestId(v string) *ListDynamicRouteRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDynamicRouteRegionsResponseBody) SetTotalNum(v int32) *ListDynamicRouteRegionsResponseBody {
	s.TotalNum = &v
	return s
}

type ListDynamicRouteRegionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDynamicRouteRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDynamicRouteRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDynamicRouteRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListDynamicRouteRegionsResponse) SetHeaders(v map[string]*string) *ListDynamicRouteRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListDynamicRouteRegionsResponse) SetStatusCode(v int32) *ListDynamicRouteRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDynamicRouteRegionsResponse) SetBody(v *ListDynamicRouteRegionsResponseBody) *ListDynamicRouteRegionsResponse {
	s.Body = v
	return s
}

type ListDynamicRoutesRequest struct {
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage     *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DynamicRouteIds []*string `json:"DynamicRouteIds,omitempty" xml:"DynamicRouteIds,omitempty" type:"Repeated"`
	// example:
	//
	// dynamic_route_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// connector-8ccb13b6f52c****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListDynamicRoutesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDynamicRoutesRequest) GoString() string {
	return s.String()
}

func (s *ListDynamicRoutesRequest) SetApplicationId(v string) *ListDynamicRoutesRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListDynamicRoutesRequest) SetCurrentPage(v int32) *ListDynamicRoutesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListDynamicRoutesRequest) SetDynamicRouteIds(v []*string) *ListDynamicRoutesRequest {
	s.DynamicRouteIds = v
	return s
}

func (s *ListDynamicRoutesRequest) SetName(v string) *ListDynamicRoutesRequest {
	s.Name = &v
	return s
}

func (s *ListDynamicRoutesRequest) SetNextHop(v string) *ListDynamicRoutesRequest {
	s.NextHop = &v
	return s
}

func (s *ListDynamicRoutesRequest) SetPageSize(v int32) *ListDynamicRoutesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDynamicRoutesRequest) SetRegionIds(v []*string) *ListDynamicRoutesRequest {
	s.RegionIds = v
	return s
}

func (s *ListDynamicRoutesRequest) SetStatus(v string) *ListDynamicRoutesRequest {
	s.Status = &v
	return s
}

func (s *ListDynamicRoutesRequest) SetTagId(v string) *ListDynamicRoutesRequest {
	s.TagId = &v
	return s
}

type ListDynamicRoutesResponseBody struct {
	DynamicRoutes []*ListDynamicRoutesResponseBodyDynamicRoutes `json:"DynamicRoutes,omitempty" xml:"DynamicRoutes,omitempty" type:"Repeated"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListDynamicRoutesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDynamicRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDynamicRoutesResponseBody) SetDynamicRoutes(v []*ListDynamicRoutesResponseBodyDynamicRoutes) *ListDynamicRoutesResponseBody {
	s.DynamicRoutes = v
	return s
}

func (s *ListDynamicRoutesResponseBody) SetRequestId(v string) *ListDynamicRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDynamicRoutesResponseBody) SetTotalNum(v int32) *ListDynamicRoutesResponseBody {
	s.TotalNum = &v
	return s
}

type ListDynamicRoutesResponseBodyDynamicRoutes struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// example:
	//
	// 2023-03-21 11:50:03
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// dr-a0ca843f53cf****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	// example:
	//
	// connector
	DynamicRouteType *string `json:"DynamicRouteType,omitempty" xml:"DynamicRouteType,omitempty"`
	// example:
	//
	// dynamic_route_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// connector-8ccb13b6f52c****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// example:
	//
	// 1
	Priority  *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s ListDynamicRoutesResponseBodyDynamicRoutes) String() string {
	return tea.Prettify(s)
}

func (s ListDynamicRoutesResponseBodyDynamicRoutes) GoString() string {
	return s.String()
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetApplicationIds(v []*string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.ApplicationIds = v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetApplicationType(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.ApplicationType = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetCreateTime(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.CreateTime = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetDescription(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.Description = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetDynamicRouteId(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.DynamicRouteId = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetDynamicRouteType(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.DynamicRouteType = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetName(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.Name = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetNextHop(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.NextHop = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetPriority(v int32) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.Priority = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetRegionIds(v []*string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.RegionIds = v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetStatus(v string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.Status = &v
	return s
}

func (s *ListDynamicRoutesResponseBodyDynamicRoutes) SetTagIds(v []*string) *ListDynamicRoutesResponseBodyDynamicRoutes {
	s.TagIds = v
	return s
}

type ListDynamicRoutesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDynamicRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDynamicRoutesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDynamicRoutesResponse) GoString() string {
	return s.String()
}

func (s *ListDynamicRoutesResponse) SetHeaders(v map[string]*string) *ListDynamicRoutesResponse {
	s.Headers = v
	return s
}

func (s *ListDynamicRoutesResponse) SetStatusCode(v int32) *ListDynamicRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDynamicRoutesResponse) SetBody(v *ListDynamicRoutesResponseBody) *ListDynamicRoutesResponse {
	s.Body = v
	return s
}

type ListExcessiveDeviceRegistrationApplicationsRequest struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Department  *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string   `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Statuses   []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	Username   *string   `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListExcessiveDeviceRegistrationApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListExcessiveDeviceRegistrationApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetApplicationIds(v []*string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetCurrentPage(v int64) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetDepartment(v string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.Department = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetDeviceTag(v string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.DeviceTag = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetHostname(v string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.Hostname = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetMac(v string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.Mac = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetPageSize(v int64) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetSaseUserId(v string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.SaseUserId = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetStatuses(v []*string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.Statuses = v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsRequest) SetUsername(v string) *ListExcessiveDeviceRegistrationApplicationsRequest {
	s.Username = &v
	return s
}

type ListExcessiveDeviceRegistrationApplicationsResponseBody struct {
	Applications []*ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListExcessiveDeviceRegistrationApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListExcessiveDeviceRegistrationApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBody) SetApplications(v []*ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) *ListExcessiveDeviceRegistrationApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBody) SetRequestId(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBody) SetTotalNum(v int64) *ListExcessiveDeviceRegistrationApplicationsResponseBody {
	s.TotalNum = &v
	return s
}

type ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications struct {
	// example:
	//
	// reg-application-0f4a127b7e78****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 2023-07-17 18:46:55
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department  *string `json:"Department,omitempty" xml:"Department,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// example:
	//
	// Windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// false
	IsUsed *bool `json:"IsUsed,omitempty" xml:"IsUsed,omitempty"`
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// example:
	//
	// Approved
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetApplicationId(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetCreateTime(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.CreateTime = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetDepartment(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.Department = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetDescription(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.Description = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetDeviceTag(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.DeviceTag = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetDeviceType(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.DeviceType = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetHostname(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.Hostname = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetIsUsed(v bool) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.IsUsed = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetMac(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.Mac = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetSaseUserId(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.SaseUserId = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetStatus(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.Status = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications) SetUsername(v string) *ListExcessiveDeviceRegistrationApplicationsResponseBodyApplications {
	s.Username = &v
	return s
}

type ListExcessiveDeviceRegistrationApplicationsResponse struct {
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExcessiveDeviceRegistrationApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExcessiveDeviceRegistrationApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListExcessiveDeviceRegistrationApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponse) SetHeaders(v map[string]*string) *ListExcessiveDeviceRegistrationApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponse) SetStatusCode(v int32) *ListExcessiveDeviceRegistrationApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExcessiveDeviceRegistrationApplicationsResponse) SetBody(v *ListExcessiveDeviceRegistrationApplicationsResponseBody) *ListExcessiveDeviceRegistrationApplicationsResponse {
	s.Body = v
	return s
}

type ListIdpConfigsRequest struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// CSAS,DingTalk,LDAP
	Include *string `json:"Include,omitempty" xml:"Include,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIdpConfigsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIdpConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListIdpConfigsRequest) SetCurrentPage(v int64) *ListIdpConfigsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListIdpConfigsRequest) SetInclude(v string) *ListIdpConfigsRequest {
	s.Include = &v
	return s
}

func (s *ListIdpConfigsRequest) SetPageSize(v int64) *ListIdpConfigsRequest {
	s.PageSize = &v
	return s
}

type ListIdpConfigsResponseBody struct {
	Data *ListIdpConfigsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// FD724DBC-CD76-5235-BF76-59C51B73296D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIdpConfigsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIdpConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdpConfigsResponseBody) SetData(v *ListIdpConfigsResponseBodyData) *ListIdpConfigsResponseBody {
	s.Data = v
	return s
}

func (s *ListIdpConfigsResponseBody) SetRequestId(v string) *ListIdpConfigsResponseBody {
	s.RequestId = &v
	return s
}

type ListIdpConfigsResponseBodyData struct {
	DataList []*ListIdpConfigsResponseBodyDataDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListIdpConfigsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIdpConfigsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIdpConfigsResponseBodyData) SetDataList(v []*ListIdpConfigsResponseBodyDataDataList) *ListIdpConfigsResponseBodyData {
	s.DataList = v
	return s
}

func (s *ListIdpConfigsResponseBodyData) SetTotalNum(v int64) *ListIdpConfigsResponseBodyData {
	s.TotalNum = &v
	return s
}

type ListIdpConfigsResponseBodyDataDataList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 277
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// totp
	Mfa *string `json:"Mfa,omitempty" xml:"Mfa,omitempty"`
	// example:
	//
	// password
	MobileLoginType *string `json:"MobileLoginType,omitempty" xml:"MobileLoginType,omitempty"`
	// example:
	//
	// password
	MobileMfaConfigType *string `json:"MobileMfaConfigType,omitempty" xml:"MobileMfaConfigType,omitempty"`
	// example:
	//
	// 1482,1355
	MultiIdpInfo *string `json:"MultiIdpInfo,omitempty" xml:"MultiIdpInfo,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// password
	PcLoginType *string `json:"PcLoginType,omitempty" xml:"PcLoginType,omitempty"`
	// example:
	//
	// Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// DingTalk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2023-05-09T02:22:41.430Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListIdpConfigsResponseBodyDataDataList) String() string {
	return tea.Prettify(s)
}

func (s ListIdpConfigsResponseBodyDataDataList) GoString() string {
	return s.String()
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetDescription(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.Description = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetId(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.Id = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetMfa(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.Mfa = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetMobileLoginType(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.MobileLoginType = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetMobileMfaConfigType(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.MobileMfaConfigType = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetMultiIdpInfo(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.MultiIdpInfo = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetName(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.Name = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetPcLoginType(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.PcLoginType = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetStatus(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.Status = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetType(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.Type = &v
	return s
}

func (s *ListIdpConfigsResponseBodyDataDataList) SetUpdateTime(v string) *ListIdpConfigsResponseBodyDataDataList {
	s.UpdateTime = &v
	return s
}

type ListIdpConfigsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIdpConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIdpConfigsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIdpConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListIdpConfigsResponse) SetHeaders(v map[string]*string) *ListIdpConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListIdpConfigsResponse) SetStatusCode(v int32) *ListIdpConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIdpConfigsResponse) SetBody(v *ListIdpConfigsResponseBody) *ListIdpConfigsResponse {
	s.Body = v
	return s
}

type ListIdpDepartmentsRequest struct {
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1440
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIdpDepartmentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIdpDepartmentsRequest) GoString() string {
	return s.String()
}

func (s *ListIdpDepartmentsRequest) SetCurrentPage(v int64) *ListIdpDepartmentsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListIdpDepartmentsRequest) SetIdpConfigId(v string) *ListIdpDepartmentsRequest {
	s.IdpConfigId = &v
	return s
}

func (s *ListIdpDepartmentsRequest) SetPageSize(v int64) *ListIdpDepartmentsRequest {
	s.PageSize = &v
	return s
}

type ListIdpDepartmentsResponseBody struct {
	Data *ListIdpDepartmentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIdpDepartmentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIdpDepartmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdpDepartmentsResponseBody) SetData(v *ListIdpDepartmentsResponseBodyData) *ListIdpDepartmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListIdpDepartmentsResponseBody) SetRequestId(v string) *ListIdpDepartmentsResponseBody {
	s.RequestId = &v
	return s
}

type ListIdpDepartmentsResponseBodyData struct {
	DataList []*ListIdpDepartmentsResponseBodyDataDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListIdpDepartmentsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListIdpDepartmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIdpDepartmentsResponseBodyData) SetDataList(v []*ListIdpDepartmentsResponseBodyDataDataList) *ListIdpDepartmentsResponseBodyData {
	s.DataList = v
	return s
}

func (s *ListIdpDepartmentsResponseBodyData) SetTotalNum(v int64) *ListIdpDepartmentsResponseBodyData {
	s.TotalNum = &v
	return s
}

type ListIdpDepartmentsResponseBodyDataDataList struct {
	// example:
	//
	// 30520
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1440
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListIdpDepartmentsResponseBodyDataDataList) String() string {
	return tea.Prettify(s)
}

func (s ListIdpDepartmentsResponseBodyDataDataList) GoString() string {
	return s.String()
}

func (s *ListIdpDepartmentsResponseBodyDataDataList) SetId(v string) *ListIdpDepartmentsResponseBodyDataDataList {
	s.Id = &v
	return s
}

func (s *ListIdpDepartmentsResponseBodyDataDataList) SetIdpConfigId(v string) *ListIdpDepartmentsResponseBodyDataDataList {
	s.IdpConfigId = &v
	return s
}

func (s *ListIdpDepartmentsResponseBodyDataDataList) SetName(v string) *ListIdpDepartmentsResponseBodyDataDataList {
	s.Name = &v
	return s
}

type ListIdpDepartmentsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIdpDepartmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIdpDepartmentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIdpDepartmentsResponse) GoString() string {
	return s.String()
}

func (s *ListIdpDepartmentsResponse) SetHeaders(v map[string]*string) *ListIdpDepartmentsResponse {
	s.Headers = v
	return s
}

func (s *ListIdpDepartmentsResponse) SetStatusCode(v int32) *ListIdpDepartmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIdpDepartmentsResponse) SetBody(v *ListIdpDepartmentsResponseBody) *ListIdpDepartmentsResponse {
	s.Body = v
	return s
}

type ListNacUserCertRequest struct {
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Department  *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 1702770400
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1702260834
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// zhang**
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListNacUserCertRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNacUserCertRequest) GoString() string {
	return s.String()
}

func (s *ListNacUserCertRequest) SetCurrentPage(v string) *ListNacUserCertRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListNacUserCertRequest) SetDepartment(v string) *ListNacUserCertRequest {
	s.Department = &v
	return s
}

func (s *ListNacUserCertRequest) SetDeviceType(v string) *ListNacUserCertRequest {
	s.DeviceType = &v
	return s
}

func (s *ListNacUserCertRequest) SetEndTime(v int64) *ListNacUserCertRequest {
	s.EndTime = &v
	return s
}

func (s *ListNacUserCertRequest) SetPageSize(v string) *ListNacUserCertRequest {
	s.PageSize = &v
	return s
}

func (s *ListNacUserCertRequest) SetStartTime(v int64) *ListNacUserCertRequest {
	s.StartTime = &v
	return s
}

func (s *ListNacUserCertRequest) SetStatus(v string) *ListNacUserCertRequest {
	s.Status = &v
	return s
}

func (s *ListNacUserCertRequest) SetUsername(v string) *ListNacUserCertRequest {
	s.Username = &v
	return s
}

type ListNacUserCertResponseBody struct {
	// example:
	//
	// 200
	Code     *int64                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	DataList []*ListNacUserCertResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListNacUserCertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNacUserCertResponseBody) GoString() string {
	return s.String()
}

func (s *ListNacUserCertResponseBody) SetCode(v int64) *ListNacUserCertResponseBody {
	s.Code = &v
	return s
}

func (s *ListNacUserCertResponseBody) SetDataList(v []*ListNacUserCertResponseBodyDataList) *ListNacUserCertResponseBody {
	s.DataList = v
	return s
}

func (s *ListNacUserCertResponseBody) SetMessage(v string) *ListNacUserCertResponseBody {
	s.Message = &v
	return s
}

func (s *ListNacUserCertResponseBody) SetRequestId(v string) *ListNacUserCertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNacUserCertResponseBody) SetTotalNum(v int64) *ListNacUserCertResponseBody {
	s.TotalNum = &v
	return s
}

type ListNacUserCertResponseBodyDataList struct {
	// example:
	//
	// 1
	Aliuid     *string `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// example:
	//
	// windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 2029-06-30 09:31:54
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// example:
	//
	// MS-XU****
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// 08:f8:**:**:**:5e
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// Disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// zhang**
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListNacUserCertResponseBodyDataList) String() string {
	return tea.Prettify(s)
}

func (s ListNacUserCertResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListNacUserCertResponseBodyDataList) SetAliuid(v string) *ListNacUserCertResponseBodyDataList {
	s.Aliuid = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetDepartment(v string) *ListNacUserCertResponseBodyDataList {
	s.Department = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetDevTag(v string) *ListNacUserCertResponseBodyDataList {
	s.DevTag = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetDeviceType(v string) *ListNacUserCertResponseBodyDataList {
	s.DeviceType = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetExpiredTime(v string) *ListNacUserCertResponseBodyDataList {
	s.ExpiredTime = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetHostname(v string) *ListNacUserCertResponseBodyDataList {
	s.Hostname = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetMac(v string) *ListNacUserCertResponseBodyDataList {
	s.Mac = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetStatus(v string) *ListNacUserCertResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetUserId(v string) *ListNacUserCertResponseBodyDataList {
	s.UserId = &v
	return s
}

func (s *ListNacUserCertResponseBodyDataList) SetUsername(v string) *ListNacUserCertResponseBodyDataList {
	s.Username = &v
	return s
}

type ListNacUserCertResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNacUserCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNacUserCertResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNacUserCertResponse) GoString() string {
	return s.String()
}

func (s *ListNacUserCertResponse) SetHeaders(v map[string]*string) *ListNacUserCertResponse {
	s.Headers = v
	return s
}

func (s *ListNacUserCertResponse) SetStatusCode(v int32) *ListNacUserCertResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNacUserCertResponse) SetBody(v *ListNacUserCertResponseBody) *ListNacUserCertResponse {
	s.Body = v
	return s
}

type ListPolicesForPrivateAccessApplicationRequest struct {
	// This parameter is required.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
}

func (s ListPolicesForPrivateAccessApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForPrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessApplicationRequest) SetApplicationIds(v []*string) *ListPolicesForPrivateAccessApplicationRequest {
	s.ApplicationIds = v
	return s
}

type ListPolicesForPrivateAccessApplicationResponseBody struct {
	Applications []*ListPolicesForPrivateAccessApplicationResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// 4AB972E2-D702-5464-B132-B1911498B8BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPolicesForPrivateAccessApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForPrivateAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessApplicationResponseBody) SetApplications(v []*ListPolicesForPrivateAccessApplicationResponseBodyApplications) *ListPolicesForPrivateAccessApplicationResponseBody {
	s.Applications = v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBody) SetRequestId(v string) *ListPolicesForPrivateAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

type ListPolicesForPrivateAccessApplicationResponseBodyApplications struct {
	// example:
	//
	// pa-application-b927baf3e592****
	ApplicationId *string                                                                   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Policies      []*ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
}

func (s ListPolicesForPrivateAccessApplicationResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForPrivateAccessApplicationResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplications) SetApplicationId(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplications) SetPolicies(v []*ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) *ListPolicesForPrivateAccessApplicationResponseBodyApplications {
	s.Policies = v
	return s
}

type ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies struct {
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// example:
	//
	// 2022-09-27 18:10:25
	CreateTime           *string                                                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomUserAttributes []*ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	Description          *string                                                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// example:
	//
	// pa-policy-867ef4007c8a****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Normal
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
}

func (s ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetApplicationType(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.ApplicationType = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetCreateTime(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetCustomUserAttributes(v []*ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.CustomUserAttributes = v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetDescription(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.Description = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetName(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.Name = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetPolicyAction(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.PolicyAction = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetPolicyId(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetPriority(v int32) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.Priority = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetStatus(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.Status = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies) SetUserGroupType(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPolicies {
	s.UserGroupType = &v
	return s
}

type ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) SetIdpId(v int32) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) SetRelation(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) SetUserGroupType(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes) SetValue(v string) *ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes {
	s.Value = &v
	return s
}

type ListPolicesForPrivateAccessApplicationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicesForPrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicesForPrivateAccessApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForPrivateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessApplicationResponse) SetHeaders(v map[string]*string) *ListPolicesForPrivateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponse) SetStatusCode(v int32) *ListPolicesForPrivateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicesForPrivateAccessApplicationResponse) SetBody(v *ListPolicesForPrivateAccessApplicationResponseBody) *ListPolicesForPrivateAccessApplicationResponse {
	s.Body = v
	return s
}

type ListPolicesForPrivateAccessTagRequest struct {
	// This parameter is required.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s ListPolicesForPrivateAccessTagRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForPrivateAccessTagRequest) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessTagRequest) SetTagIds(v []*string) *ListPolicesForPrivateAccessTagRequest {
	s.TagIds = v
	return s
}

type ListPolicesForPrivateAccessTagResponseBody struct {
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*ListPolicesForPrivateAccessTagResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListPolicesForPrivateAccessTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForPrivateAccessTagResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessTagResponseBody) SetRequestId(v string) *ListPolicesForPrivateAccessTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBody) SetTags(v []*ListPolicesForPrivateAccessTagResponseBodyTags) *ListPolicesForPrivateAccessTagResponseBody {
	s.Tags = v
	return s
}

type ListPolicesForPrivateAccessTagResponseBodyTags struct {
	Polices []*ListPolicesForPrivateAccessTagResponseBodyTagsPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// example:
	//
	// tag-b927baf3e592****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListPolicesForPrivateAccessTagResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForPrivateAccessTagResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTags) SetPolices(v []*ListPolicesForPrivateAccessTagResponseBodyTagsPolices) *ListPolicesForPrivateAccessTagResponseBodyTags {
	s.Polices = v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTags) SetTagId(v string) *ListPolicesForPrivateAccessTagResponseBodyTags {
	s.TagId = &v
	return s
}

type ListPolicesForPrivateAccessTagResponseBodyTagsPolices struct {
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// 内网访问策略创建时间。
	//
	// example:
	//
	// 2023-02-21 14:10:16
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 自定义用户组属性集合。多个自定义用户组属性之间是或的关系，按照合集生效。
	CustomUserAttributes []*ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	Description          *string                                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// example:
	//
	// pa-policy-867ef4007c8a****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Normal
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
}

func (s ListPolicesForPrivateAccessTagResponseBodyTagsPolices) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForPrivateAccessTagResponseBodyTagsPolices) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetApplicationType(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.ApplicationType = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetCreateTime(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.CreateTime = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetCustomUserAttributes(v []*ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.CustomUserAttributes = v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetDescription(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.Description = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetName(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.Name = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetPolicyAction(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.PolicyAction = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetPolicyId(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.PolicyId = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetPriority(v int32) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.Priority = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetStatus(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.Status = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolices) SetUserGroupType(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolices {
	s.UserGroupType = &v
	return s
}

type ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes struct {
	// 用户组的身份源ID。当自定义用户组类型为**department**时，存在该值。
	//
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// 用户组的关系。取值：
	//
	// - **Equal**：等于。
	//
	// - **Unequal**：不等于。
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// 用户组的类型。取值：
	//
	// - **username**：用户名。
	//
	// - **department**：部门。
	//
	// - **email**：邮箱。
	//
	// - **telephone**：手机。
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// 用户组属性的值。
	//
	// - 当用户组类型为**username**时，表示用户名的值。长度为1~128个字符，支持中文和大小写英文字母，可包含数字、半角句号（.）、下划线（_）和短划线（-）。
	//
	// - 当用户组类型为**department**时，表示部门的值。如：OU=部门1,OU=SASE钉钉。
	//
	// - 当用户组类型为**email**时，表示邮箱的值。如：username@example.com。
	//
	// - 当用户组类型为**telephone**时，表示手机的值。如：13900001234。
	//
	// example:
	//
	// OU=部门1,OU=SASE钉钉
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) SetIdpId(v int32) *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) SetRelation(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) SetUserGroupType(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes) SetValue(v string) *ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes {
	s.Value = &v
	return s
}

type ListPolicesForPrivateAccessTagResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicesForPrivateAccessTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicesForPrivateAccessTagResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForPrivateAccessTagResponse) GoString() string {
	return s.String()
}

func (s *ListPolicesForPrivateAccessTagResponse) SetHeaders(v map[string]*string) *ListPolicesForPrivateAccessTagResponse {
	s.Headers = v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponse) SetStatusCode(v int32) *ListPolicesForPrivateAccessTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicesForPrivateAccessTagResponse) SetBody(v *ListPolicesForPrivateAccessTagResponseBody) *ListPolicesForPrivateAccessTagResponse {
	s.Body = v
	return s
}

type ListPolicesForUserGroupRequest struct {
	// This parameter is required.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
}

func (s ListPolicesForUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListPolicesForUserGroupRequest) SetUserGroupIds(v []*string) *ListPolicesForUserGroupRequest {
	s.UserGroupIds = v
	return s
}

type ListPolicesForUserGroupResponseBody struct {
	// example:
	//
	// 5F04DFBD-3F48-5F70-AE72-474026670128
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserGroups []*ListPolicesForUserGroupResponseBodyUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListPolicesForUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicesForUserGroupResponseBody) SetRequestId(v string) *ListPolicesForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicesForUserGroupResponseBody) SetUserGroups(v []*ListPolicesForUserGroupResponseBodyUserGroups) *ListPolicesForUserGroupResponseBody {
	s.UserGroups = v
	return s
}

type ListPolicesForUserGroupResponseBodyUserGroups struct {
	Polices []*ListPolicesForUserGroupResponseBodyUserGroupsPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListPolicesForUserGroupResponseBodyUserGroups) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForUserGroupResponseBodyUserGroups) GoString() string {
	return s.String()
}

func (s *ListPolicesForUserGroupResponseBodyUserGroups) SetPolices(v []*ListPolicesForUserGroupResponseBodyUserGroupsPolices) *ListPolicesForUserGroupResponseBodyUserGroups {
	s.Polices = v
	return s
}

func (s *ListPolicesForUserGroupResponseBodyUserGroups) SetUserGroupId(v string) *ListPolicesForUserGroupResponseBodyUserGroups {
	s.UserGroupId = &v
	return s
}

type ListPolicesForUserGroupResponseBodyUserGroupsPolices struct {
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// pa-policy-ce2bf7236fab****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// PrivateAccess
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListPolicesForUserGroupResponseBodyUserGroupsPolices) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForUserGroupResponseBodyUserGroupsPolices) GoString() string {
	return s.String()
}

func (s *ListPolicesForUserGroupResponseBodyUserGroupsPolices) SetName(v string) *ListPolicesForUserGroupResponseBodyUserGroupsPolices {
	s.Name = &v
	return s
}

func (s *ListPolicesForUserGroupResponseBodyUserGroupsPolices) SetPolicyId(v string) *ListPolicesForUserGroupResponseBodyUserGroupsPolices {
	s.PolicyId = &v
	return s
}

func (s *ListPolicesForUserGroupResponseBodyUserGroupsPolices) SetPolicyType(v string) *ListPolicesForUserGroupResponseBodyUserGroupsPolices {
	s.PolicyType = &v
	return s
}

type ListPolicesForUserGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPolicesForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPolicesForUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPolicesForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListPolicesForUserGroupResponse) SetHeaders(v map[string]*string) *ListPolicesForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListPolicesForUserGroupResponse) SetStatusCode(v int32) *ListPolicesForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPolicesForUserGroupResponse) SetBody(v *ListPolicesForUserGroupResponseBody) *ListPolicesForUserGroupResponse {
	s.Body = v
	return s
}

type ListPopTrafficStatisticsRequest struct {
	// example:
	//
	// 1681293719
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 1681035708
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListPopTrafficStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPopTrafficStatisticsRequest) GoString() string {
	return s.String()
}

func (s *ListPopTrafficStatisticsRequest) SetEndTime(v string) *ListPopTrafficStatisticsRequest {
	s.EndTime = &v
	return s
}

func (s *ListPopTrafficStatisticsRequest) SetRegion(v string) *ListPopTrafficStatisticsRequest {
	s.Region = &v
	return s
}

func (s *ListPopTrafficStatisticsRequest) SetStartTime(v string) *ListPopTrafficStatisticsRequest {
	s.StartTime = &v
	return s
}

type ListPopTrafficStatisticsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// EFE7EBB2-449D-5BBB-B381-CA7839BC1649
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TrafficData []*ListPopTrafficStatisticsResponseBodyTrafficData `json:"TrafficData,omitempty" xml:"TrafficData,omitempty" type:"Repeated"`
}

func (s ListPopTrafficStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPopTrafficStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPopTrafficStatisticsResponseBody) SetRequestId(v string) *ListPopTrafficStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPopTrafficStatisticsResponseBody) SetTrafficData(v []*ListPopTrafficStatisticsResponseBodyTrafficData) *ListPopTrafficStatisticsResponseBody {
	s.TrafficData = v
	return s
}

type ListPopTrafficStatisticsResponseBodyTrafficData struct {
	Datapoints []*ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Repeated"`
	// example:
	//
	// InternetTx
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
}

func (s ListPopTrafficStatisticsResponseBodyTrafficData) String() string {
	return tea.Prettify(s)
}

func (s ListPopTrafficStatisticsResponseBodyTrafficData) GoString() string {
	return s.String()
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficData) SetDatapoints(v []*ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints) *ListPopTrafficStatisticsResponseBodyTrafficData {
	s.Datapoints = v
	return s
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficData) SetMetricName(v string) *ListPopTrafficStatisticsResponseBodyTrafficData {
	s.MetricName = &v
	return s
}

type ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints struct {
	// example:
	//
	// 15325
	Average *float64 `json:"Average,omitempty" xml:"Average,omitempty"`
	// example:
	//
	// 2023-12-06 15:29:00
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
}

func (s ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints) String() string {
	return tea.Prettify(s)
}

func (s ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints) GoString() string {
	return s.String()
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints) SetAverage(v float64) *ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints {
	s.Average = &v
	return s
}

func (s *ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints) SetDateTime(v string) *ListPopTrafficStatisticsResponseBodyTrafficDataDatapoints {
	s.DateTime = &v
	return s
}

type ListPopTrafficStatisticsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPopTrafficStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPopTrafficStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPopTrafficStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListPopTrafficStatisticsResponse) SetHeaders(v map[string]*string) *ListPopTrafficStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListPopTrafficStatisticsResponse) SetStatusCode(v int32) *ListPopTrafficStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPopTrafficStatisticsResponse) SetBody(v *ListPopTrafficStatisticsResponseBody) *ListPopTrafficStatisticsResponse {
	s.Body = v
	return s
}

type ListPrivateAccessApplicationsRequest struct {
	AccessModes *string `json:"AccessModes,omitempty" xml:"AccessModes,omitempty"`
	// example:
	//
	// 192.168.0.0/16
	Address        *string   `json:"Address,omitempty" xml:"Address,omitempty"`
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	ConnectorId    *string   `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// private_access_application_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// pa-policy-54a7838a48bf****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListPrivateAccessApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsRequest) SetAccessModes(v string) *ListPrivateAccessApplicationsRequest {
	s.AccessModes = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetAddress(v string) *ListPrivateAccessApplicationsRequest {
	s.Address = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetApplicationIds(v []*string) *ListPrivateAccessApplicationsRequest {
	s.ApplicationIds = v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetConnectorId(v string) *ListPrivateAccessApplicationsRequest {
	s.ConnectorId = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetCurrentPage(v int32) *ListPrivateAccessApplicationsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetName(v string) *ListPrivateAccessApplicationsRequest {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetPageSize(v int32) *ListPrivateAccessApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetPolicyId(v string) *ListPrivateAccessApplicationsRequest {
	s.PolicyId = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetStatus(v string) *ListPrivateAccessApplicationsRequest {
	s.Status = &v
	return s
}

func (s *ListPrivateAccessApplicationsRequest) SetTagId(v string) *ListPrivateAccessApplicationsRequest {
	s.TagId = &v
	return s
}

type ListPrivateAccessApplicationsResponseBody struct {
	Applications []*ListPrivateAccessApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// 748CFDC7-1EB6-5B8B-9405-DA76ED5BB60D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListPrivateAccessApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsResponseBody) SetApplications(v []*ListPrivateAccessApplicationsResponseBodyApplications) *ListPrivateAccessApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBody) SetRequestId(v string) *ListPrivateAccessApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBody) SetTotalNum(v int32) *ListPrivateAccessApplicationsResponseBody {
	s.TotalNum = &v
	return s
}

type ListPrivateAccessApplicationsResponseBodyApplications struct {
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId       *string   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	BrowserAccessStatus *string   `json:"BrowserAccessStatus,omitempty" xml:"BrowserAccessStatus,omitempty"`
	ConnectorIds        []*string `json:"ConnectorIds,omitempty" xml:"ConnectorIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-12-16 15:03:42
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	L7ProxyDomainAutomatic *string `json:"L7ProxyDomainAutomatic,omitempty" xml:"L7ProxyDomainAutomatic,omitempty"`
	L7ProxyDomainCustom    *string `json:"L7ProxyDomainCustom,omitempty" xml:"L7ProxyDomainCustom,omitempty"`
	// example:
	//
	// private_access_application_name
	Name       *string                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyIds  []*string                                                          `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	PortRanges []*ListPrivateAccessApplicationsResponseBodyApplicationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// example:
	//
	// All
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// Enabled
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s ListPrivateAccessApplicationsResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetAddresses(v []*string) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.Addresses = v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetApplicationId(v string) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetBrowserAccessStatus(v string) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.BrowserAccessStatus = &v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetConnectorIds(v []*string) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.ConnectorIds = v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetCreateTime(v string) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.CreateTime = &v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetDescription(v string) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.Description = &v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetL7ProxyDomainAutomatic(v string) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.L7ProxyDomainAutomatic = &v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetL7ProxyDomainCustom(v string) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.L7ProxyDomainCustom = &v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetName(v string) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetPolicyIds(v []*string) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.PolicyIds = v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetPortRanges(v []*ListPrivateAccessApplicationsResponseBodyApplicationsPortRanges) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.PortRanges = v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetProtocol(v string) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.Protocol = &v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetStatus(v string) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.Status = &v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplications) SetTagIds(v []*string) *ListPrivateAccessApplicationsResponseBodyApplications {
	s.TagIds = v
	return s
}

type ListPrivateAccessApplicationsResponseBodyApplicationsPortRanges struct {
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 81
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s ListPrivateAccessApplicationsResponseBodyApplicationsPortRanges) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessApplicationsResponseBodyApplicationsPortRanges) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsResponseBodyApplicationsPortRanges) SetBegin(v int32) *ListPrivateAccessApplicationsResponseBodyApplicationsPortRanges {
	s.Begin = &v
	return s
}

func (s *ListPrivateAccessApplicationsResponseBodyApplicationsPortRanges) SetEnd(v int32) *ListPrivateAccessApplicationsResponseBodyApplicationsPortRanges {
	s.End = &v
	return s
}

type ListPrivateAccessApplicationsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateAccessApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateAccessApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsResponse) SetHeaders(v map[string]*string) *ListPrivateAccessApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateAccessApplicationsResponse) SetStatusCode(v int32) *ListPrivateAccessApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateAccessApplicationsResponse) SetBody(v *ListPrivateAccessApplicationsResponseBody) *ListPrivateAccessApplicationsResponse {
	s.Body = v
	return s
}

type ListPrivateAccessApplicationsForDynamicRouteRequest struct {
	// This parameter is required.
	DynamicRouteIds []*string `json:"DynamicRouteIds,omitempty" xml:"DynamicRouteIds,omitempty" type:"Repeated"`
}

func (s ListPrivateAccessApplicationsForDynamicRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessApplicationsForDynamicRouteRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsForDynamicRouteRequest) SetDynamicRouteIds(v []*string) *ListPrivateAccessApplicationsForDynamicRouteRequest {
	s.DynamicRouteIds = v
	return s
}

type ListPrivateAccessApplicationsForDynamicRouteResponseBody struct {
	DynamicRoutes []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes `json:"DynamicRoutes,omitempty" xml:"DynamicRoutes,omitempty" type:"Repeated"`
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBody) SetDynamicRoutes(v []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) *ListPrivateAccessApplicationsForDynamicRouteResponseBody {
	s.DynamicRoutes = v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBody) SetRequestId(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBody {
	s.RequestId = &v
	return s
}

type ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes struct {
	Applications []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// dr-ca9fddfac7c6****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) SetApplications(v []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes {
	s.Applications = v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes) SetDynamicRouteId(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutes {
	s.DynamicRouteId = &v
	return s
}

type ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications struct {
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// example:
	//
	// pa-application-7a9243dd02f4****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 2022-04-13 13:33:24
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// application_name
	Name       *string                                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// example:
	//
	// All
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetAddresses(v []*string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.Addresses = v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetApplicationId(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetCreateTime(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.CreateTime = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetDescription(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.Description = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetName(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetPortRanges(v []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.PortRanges = v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetProtocol(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.Protocol = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications) SetStatus(v string) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications {
	s.Status = &v
	return s
}

type ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges struct {
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 81
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges) SetBegin(v int32) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges {
	s.Begin = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges) SetEnd(v int32) *ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges {
	s.End = &v
	return s
}

type ListPrivateAccessApplicationsForDynamicRouteResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateAccessApplicationsForDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessApplicationsForDynamicRouteResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponse) SetHeaders(v map[string]*string) *ListPrivateAccessApplicationsForDynamicRouteResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponse) SetStatusCode(v int32) *ListPrivateAccessApplicationsForDynamicRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateAccessApplicationsForDynamicRouteResponse) SetBody(v *ListPrivateAccessApplicationsForDynamicRouteResponseBody) *ListPrivateAccessApplicationsForDynamicRouteResponse {
	s.Body = v
	return s
}

type ListPrivateAccessPolicesRequest struct {
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId   *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// Allow
	PolicyAction *string   `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	PolicyIds    []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// tag-c0cb77857a99****
	TagId   *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// 用户组ID。取值来源：
	//
	// - [ListUserGroups](~~ListUserGroups~~)：批量查询用户组。
	//
	// - [CreateUserGroup](~~CreateUserGroup~~)：创建用户组。
	//
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListPrivateAccessPolicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessPolicesRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessPolicesRequest) SetApplicationId(v string) *ListPrivateAccessPolicesRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetApplicationName(v string) *ListPrivateAccessPolicesRequest {
	s.ApplicationName = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetCurrentPage(v int32) *ListPrivateAccessPolicesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetName(v string) *ListPrivateAccessPolicesRequest {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetPageSize(v int32) *ListPrivateAccessPolicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetPolicyAction(v string) *ListPrivateAccessPolicesRequest {
	s.PolicyAction = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetPolicyIds(v []*string) *ListPrivateAccessPolicesRequest {
	s.PolicyIds = v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetStatus(v string) *ListPrivateAccessPolicesRequest {
	s.Status = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetTagId(v string) *ListPrivateAccessPolicesRequest {
	s.TagId = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetTagName(v string) *ListPrivateAccessPolicesRequest {
	s.TagName = &v
	return s
}

func (s *ListPrivateAccessPolicesRequest) SetUserGroupId(v string) *ListPrivateAccessPolicesRequest {
	s.UserGroupId = &v
	return s
}

type ListPrivateAccessPolicesResponseBody struct {
	Polices []*ListPrivateAccessPolicesResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// example:
	//
	// 748CFDC7-1EB6-5B8B-9405-DA76ED5BB60D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListPrivateAccessPolicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessPolicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessPolicesResponseBody) SetPolices(v []*ListPrivateAccessPolicesResponseBodyPolices) *ListPrivateAccessPolicesResponseBody {
	s.Polices = v
	return s
}

func (s *ListPrivateAccessPolicesResponseBody) SetRequestId(v string) *ListPrivateAccessPolicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBody) SetTotalNum(v int32) *ListPrivateAccessPolicesResponseBody {
	s.TotalNum = &v
	return s
}

type ListPrivateAccessPolicesResponseBodyPolices struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// example:
	//
	// 2022-07-10 15:50:23
	CreateTime            *string                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomUserAttributes  []*ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	Description           *string                                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceAttributeAction *string                                                            `json:"DeviceAttributeAction,omitempty" xml:"DeviceAttributeAction,omitempty"`
	DeviceAttributeId     *string                                                            `json:"DeviceAttributeId,omitempty" xml:"DeviceAttributeId,omitempty"`
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// example:
	//
	// pa-policy-63b2f1844b86****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds       []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// Normal
	UserGroupMode *string `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
}

func (s ListPrivateAccessPolicesResponseBodyPolices) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessPolicesResponseBodyPolices) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetApplicationIds(v []*string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.ApplicationIds = v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetApplicationType(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.ApplicationType = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetCreateTime(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.CreateTime = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetCustomUserAttributes(v []*ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) *ListPrivateAccessPolicesResponseBodyPolices {
	s.CustomUserAttributes = v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetDescription(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.Description = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetDeviceAttributeAction(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.DeviceAttributeAction = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetDeviceAttributeId(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.DeviceAttributeId = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetName(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetPolicyAction(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.PolicyAction = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetPolicyId(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.PolicyId = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetPriority(v int32) *ListPrivateAccessPolicesResponseBodyPolices {
	s.Priority = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetStatus(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.Status = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetTagIds(v []*string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.TagIds = v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetUserGroupIds(v []*string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.UserGroupIds = v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolices) SetUserGroupMode(v string) *ListPrivateAccessPolicesResponseBodyPolices {
	s.UserGroupMode = &v
	return s
}

type ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) SetIdpId(v int32) *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) SetRelation(v string) *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) SetUserGroupType(v string) *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes) SetValue(v string) *ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes {
	s.Value = &v
	return s
}

type ListPrivateAccessPolicesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateAccessPolicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateAccessPolicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessPolicesResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessPolicesResponse) SetHeaders(v map[string]*string) *ListPrivateAccessPolicesResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateAccessPolicesResponse) SetStatusCode(v int32) *ListPrivateAccessPolicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateAccessPolicesResponse) SetBody(v *ListPrivateAccessPolicesResponseBody) *ListPrivateAccessPolicesResponse {
	s.Body = v
	return s
}

type ListPrivateAccessTagsRequest struct {
	// The ID of the internal access application. You can obtain the application ID by calling the following operations:
	//
	// 	- [ListPrivateAccessApplications](~~ListPrivateAccessApplications~~): queries all internal access applications.
	//
	// 	- [CreatePrivateAccessApplication](~~CreatePrivateAccessApplication~~): creates an internal access application.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The page number. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The name of the internal access tag. The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// tag_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of entries per page. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the internal access policy. You can obtain the policy ID by calling the following operations:
	//
	// 	- [ListPrivateAccessPolices](~~ListPrivateAccessPolices~~): queries all internal access policies.
	//
	// 	- [CreatePrivateAccessPolicy](~~CreatePrivateAccessPolicy~~): creates an internal access policy.
	//
	// example:
	//
	// pa-policy-54a7838a48bf****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// Specifies whether to enable the simple query mode. A value of true specifies that policy IDs are not queried.
	//
	// example:
	//
	// true
	SimpleMode *bool `json:"SimpleMode,omitempty" xml:"SimpleMode,omitempty"`
	// The IDs of internal access tags. You can specify up to 100 tag IDs.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s ListPrivateAccessTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessTagsRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsRequest) SetApplicationId(v string) *ListPrivateAccessTagsRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListPrivateAccessTagsRequest) SetCurrentPage(v int32) *ListPrivateAccessTagsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListPrivateAccessTagsRequest) SetName(v string) *ListPrivateAccessTagsRequest {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessTagsRequest) SetPageSize(v int32) *ListPrivateAccessTagsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPrivateAccessTagsRequest) SetPolicyId(v string) *ListPrivateAccessTagsRequest {
	s.PolicyId = &v
	return s
}

func (s *ListPrivateAccessTagsRequest) SetSimpleMode(v bool) *ListPrivateAccessTagsRequest {
	s.SimpleMode = &v
	return s
}

func (s *ListPrivateAccessTagsRequest) SetTagIds(v []*string) *ListPrivateAccessTagsRequest {
	s.TagIds = v
	return s
}

type ListPrivateAccessTagsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54C1D236-CDB9-586C-B44D-AFDCEA195545
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The internal access tags.
	Tags []*ListPrivateAccessTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The total number of internal access tags.
	//
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListPrivateAccessTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsResponseBody) SetRequestId(v string) *ListPrivateAccessTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrivateAccessTagsResponseBody) SetTags(v []*ListPrivateAccessTagsResponseBodyTags) *ListPrivateAccessTagsResponseBody {
	s.Tags = v
	return s
}

func (s *ListPrivateAccessTagsResponseBody) SetTotalNum(v int32) *ListPrivateAccessTagsResponseBody {
	s.TotalNum = &v
	return s
}

type ListPrivateAccessTagsResponseBodyTags struct {
	// The IDs of the internal access applications.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// The time when the internal access tag was created.
	//
	// example:
	//
	// 2022-10-10 11:39:34
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the internal access tag.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the internal access tag.
	//
	// example:
	//
	// tag_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IDs of the internal access policies.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// The ID of the internal access tag.
	//
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// The type of the internal access tag. Valid values:
	//
	// 	- **Default**
	//
	// 	- **Custom**
	//
	// example:
	//
	// Default
	TagType *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
}

func (s ListPrivateAccessTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetApplicationIds(v []*string) *ListPrivateAccessTagsResponseBodyTags {
	s.ApplicationIds = v
	return s
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetCreateTime(v string) *ListPrivateAccessTagsResponseBodyTags {
	s.CreateTime = &v
	return s
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetDescription(v string) *ListPrivateAccessTagsResponseBodyTags {
	s.Description = &v
	return s
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetName(v string) *ListPrivateAccessTagsResponseBodyTags {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetPolicyIds(v []*string) *ListPrivateAccessTagsResponseBodyTags {
	s.PolicyIds = v
	return s
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetTagId(v string) *ListPrivateAccessTagsResponseBodyTags {
	s.TagId = &v
	return s
}

func (s *ListPrivateAccessTagsResponseBodyTags) SetTagType(v string) *ListPrivateAccessTagsResponseBodyTags {
	s.TagType = &v
	return s
}

type ListPrivateAccessTagsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateAccessTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateAccessTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessTagsResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsResponse) SetHeaders(v map[string]*string) *ListPrivateAccessTagsResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateAccessTagsResponse) SetStatusCode(v int32) *ListPrivateAccessTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateAccessTagsResponse) SetBody(v *ListPrivateAccessTagsResponseBody) *ListPrivateAccessTagsResponse {
	s.Body = v
	return s
}

type ListPrivateAccessTagsForDynamicRouteRequest struct {
	// This parameter is required.
	DynamicRouteIds []*string `json:"DynamicRouteIds,omitempty" xml:"DynamicRouteIds,omitempty" type:"Repeated"`
}

func (s ListPrivateAccessTagsForDynamicRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessTagsForDynamicRouteRequest) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsForDynamicRouteRequest) SetDynamicRouteIds(v []*string) *ListPrivateAccessTagsForDynamicRouteRequest {
	s.DynamicRouteIds = v
	return s
}

type ListPrivateAccessTagsForDynamicRouteResponseBody struct {
	DynamicRoutes []*ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes `json:"DynamicRoutes,omitempty" xml:"DynamicRoutes,omitempty" type:"Repeated"`
	// example:
	//
	// B608C6AE-623D-55C4-9454-601B88AE937E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrivateAccessTagsForDynamicRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessTagsForDynamicRouteResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBody) SetDynamicRoutes(v []*ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) *ListPrivateAccessTagsForDynamicRouteResponseBody {
	s.DynamicRoutes = v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBody) SetRequestId(v string) *ListPrivateAccessTagsForDynamicRouteResponseBody {
	s.RequestId = &v
	return s
}

type ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes struct {
	// example:
	//
	// dr-ca9fddfac7c6****
	DynamicRouteId *string                                                              `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	Tags           []*ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) SetDynamicRouteId(v string) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes {
	s.DynamicRouteId = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes) SetTags(v []*ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutes {
	s.Tags = v
	return s
}

type ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags struct {
	// example:
	//
	// 2022-10-23 14:02:56
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// tag_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// Custom
	TagType *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
}

func (s ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) SetCreateTime(v string) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags {
	s.CreateTime = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) SetDescription(v string) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags {
	s.Description = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) SetName(v string) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags {
	s.Name = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) SetTagId(v string) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags {
	s.TagId = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags) SetTagType(v string) *ListPrivateAccessTagsForDynamicRouteResponseBodyDynamicRoutesTags {
	s.TagType = &v
	return s
}

type ListPrivateAccessTagsForDynamicRouteResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateAccessTagsForDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateAccessTagsForDynamicRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessTagsForDynamicRouteResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessTagsForDynamicRouteResponse) SetHeaders(v map[string]*string) *ListPrivateAccessTagsForDynamicRouteResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponse) SetStatusCode(v int32) *ListPrivateAccessTagsForDynamicRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateAccessTagsForDynamicRouteResponse) SetBody(v *ListPrivateAccessTagsForDynamicRouteResponseBody) *ListPrivateAccessTagsForDynamicRouteResponse {
	s.Body = v
	return s
}

type ListRegistrationPoliciesRequest struct {
	// example:
	//
	// LimitAll
	CompanyLimitType *string `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// example:
	//
	// registration_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// LimitDiff
	PersonalLimitType *string   `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	PolicyIds         []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListRegistrationPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRegistrationPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesRequest) SetCompanyLimitType(v string) *ListRegistrationPoliciesRequest {
	s.CompanyLimitType = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetCurrentPage(v int64) *ListRegistrationPoliciesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetMatchMode(v string) *ListRegistrationPoliciesRequest {
	s.MatchMode = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetName(v string) *ListRegistrationPoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetPageSize(v int64) *ListRegistrationPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetPersonalLimitType(v string) *ListRegistrationPoliciesRequest {
	s.PersonalLimitType = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetPolicyIds(v []*string) *ListRegistrationPoliciesRequest {
	s.PolicyIds = v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetStatus(v string) *ListRegistrationPoliciesRequest {
	s.Status = &v
	return s
}

func (s *ListRegistrationPoliciesRequest) SetUserGroupId(v string) *ListRegistrationPoliciesRequest {
	s.UserGroupId = &v
	return s
}

type ListRegistrationPoliciesResponseBody struct {
	Policies []*ListRegistrationPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// example:
	//
	// 7A8FE38A-E29C-5678-B84A-FEDBCB83552F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *string `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListRegistrationPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegistrationPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesResponseBody) SetPolicies(v []*ListRegistrationPoliciesResponseBodyPolicies) *ListRegistrationPoliciesResponseBody {
	s.Policies = v
	return s
}

func (s *ListRegistrationPoliciesResponseBody) SetRequestId(v string) *ListRegistrationPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBody) SetTotalNum(v string) *ListRegistrationPoliciesResponseBody {
	s.TotalNum = &v
	return s
}

type ListRegistrationPoliciesResponseBodyPolicies struct {
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime  *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail []*ListRegistrationPoliciesResponseBodyPoliciesLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// example:
	//
	// registration_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// reg-policy-dcbfd33cb004****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s ListRegistrationPoliciesResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListRegistrationPoliciesResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetCreateTime(v string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetDescription(v string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.Description = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetLimitDetail(v []*ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) *ListRegistrationPoliciesResponseBodyPolicies {
	s.LimitDetail = v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetMatchMode(v string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.MatchMode = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetName(v string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.Name = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetPolicyId(v string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetPriority(v int64) *ListRegistrationPoliciesResponseBodyPolicies {
	s.Priority = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetStatus(v string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.Status = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetUserGroupIds(v []*string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.UserGroupIds = v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPolicies) SetWhitelist(v []*string) *ListRegistrationPoliciesResponseBodyPolicies {
	s.Whitelist = v
	return s
}

type ListRegistrationPoliciesResponseBodyPoliciesLimitDetail struct {
	// example:
	//
	// Company
	DeviceBelong *string                                                            `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	// example:
	//
	// LimitAll
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) String() string {
	return tea.Prettify(s)
}

func (s ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) SetDeviceBelong(v string) *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail {
	s.DeviceBelong = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) SetLimitCount(v *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail {
	s.LimitCount = v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail) SetLimitType(v string) *ListRegistrationPoliciesResponseBodyPoliciesLimitDetail {
	s.LimitType = &v
	return s
}

type ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount struct {
	// example:
	//
	// 3
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 0
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 0
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) String() string {
	return tea.Prettify(s)
}

func (s ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) SetAll(v int32) *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount {
	s.All = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) SetMobile(v int32) *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount {
	s.Mobile = &v
	return s
}

func (s *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount) SetPC(v int32) *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount {
	s.PC = &v
	return s
}

type ListRegistrationPoliciesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegistrationPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegistrationPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegistrationPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesResponse) SetHeaders(v map[string]*string) *ListRegistrationPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListRegistrationPoliciesResponse) SetStatusCode(v int32) *ListRegistrationPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegistrationPoliciesResponse) SetBody(v *ListRegistrationPoliciesResponseBody) *ListRegistrationPoliciesResponse {
	s.Body = v
	return s
}

type ListRegistrationPoliciesForUserGroupRequest struct {
	// This parameter is required.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
}

func (s ListRegistrationPoliciesForUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupRequest) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupRequest) SetUserGroupIds(v []*string) *ListRegistrationPoliciesForUserGroupRequest {
	s.UserGroupIds = v
	return s
}

type ListRegistrationPoliciesForUserGroupResponseBody struct {
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId  *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserGroups []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListRegistrationPoliciesForUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupResponseBody) SetRequestId(v string) *ListRegistrationPoliciesForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBody) SetUserGroups(v []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) *ListRegistrationPoliciesForUserGroupResponseBody {
	s.UserGroups = v
	return s
}

type ListRegistrationPoliciesForUserGroupResponseBodyUserGroups struct {
	Policies []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) String() string {
	return tea.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) SetPolicies(v []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroups {
	s.Policies = v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroups) SetUserGroupId(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroups {
	s.UserGroupId = &v
	return s
}

type ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies struct {
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime  *string                                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// example:
	//
	// registration_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// reg-policy-dcbfd33cb004****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status    *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	Whitelist []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetCreateTime(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetDescription(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.Description = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetLimitDetail(v []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.LimitDetail = v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetMatchMode(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.MatchMode = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetName(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.Name = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetPolicyId(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetPriority(v int64) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.Priority = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetStatus(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.Status = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies) SetWhitelist(v []*string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies {
	s.Whitelist = v
	return s
}

type ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail struct {
	// example:
	//
	// Company
	DeviceBelong *string                                                                                  `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	// example:
	//
	// LimitAll
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) String() string {
	return tea.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) SetDeviceBelong(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail {
	s.DeviceBelong = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) SetLimitCount(v *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail {
	s.LimitCount = v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail) SetLimitType(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail {
	s.LimitType = &v
	return s
}

type ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount struct {
	// example:
	//
	// 3
	All *string `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 0
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 0
	PC *string `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) String() string {
	return tea.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) SetAll(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount {
	s.All = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) SetMobile(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount {
	s.Mobile = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount) SetPC(v string) *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount {
	s.PC = &v
	return s
}

type ListRegistrationPoliciesForUserGroupResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRegistrationPoliciesForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRegistrationPoliciesForUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegistrationPoliciesForUserGroupResponse) GoString() string {
	return s.String()
}

func (s *ListRegistrationPoliciesForUserGroupResponse) SetHeaders(v map[string]*string) *ListRegistrationPoliciesForUserGroupResponse {
	s.Headers = v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponse) SetStatusCode(v int32) *ListRegistrationPoliciesForUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegistrationPoliciesForUserGroupResponse) SetBody(v *ListRegistrationPoliciesForUserGroupResponseBody) *ListRegistrationPoliciesForUserGroupResponse {
	s.Body = v
	return s
}

type ListSoftwareForUserDeviceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSoftwareForUserDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwareForUserDeviceRequest) GoString() string {
	return s.String()
}

func (s *ListSoftwareForUserDeviceRequest) SetCurrentPage(v int64) *ListSoftwareForUserDeviceRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListSoftwareForUserDeviceRequest) SetDeviceTag(v string) *ListSoftwareForUserDeviceRequest {
	s.DeviceTag = &v
	return s
}

func (s *ListSoftwareForUserDeviceRequest) SetPageSize(v int64) *ListSoftwareForUserDeviceRequest {
	s.PageSize = &v
	return s
}

type ListSoftwareForUserDeviceResponseBody struct {
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Software  []*ListSoftwareForUserDeviceResponseBodySoftware `json:"Software,omitempty" xml:"Software,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListSoftwareForUserDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwareForUserDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ListSoftwareForUserDeviceResponseBody) SetRequestId(v string) *ListSoftwareForUserDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSoftwareForUserDeviceResponseBody) SetSoftware(v []*ListSoftwareForUserDeviceResponseBodySoftware) *ListSoftwareForUserDeviceResponseBody {
	s.Software = v
	return s
}

func (s *ListSoftwareForUserDeviceResponseBody) SetTotalNum(v int64) *ListSoftwareForUserDeviceResponseBody {
	s.TotalNum = &v
	return s
}

type ListSoftwareForUserDeviceResponseBodySoftware struct {
	// example:
	//
	// Alibaba (China) Network Technology Co.,Ltd.
	Inc *string `json:"Inc,omitempty" xml:"Inc,omitempty"`
	// example:
	//
	// 2023-08-18 02:43:02
	InstallTime *string   `json:"InstallTime,omitempty" xml:"InstallTime,omitempty"`
	Name        *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Versions    []*string `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s ListSoftwareForUserDeviceResponseBodySoftware) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwareForUserDeviceResponseBodySoftware) GoString() string {
	return s.String()
}

func (s *ListSoftwareForUserDeviceResponseBodySoftware) SetInc(v string) *ListSoftwareForUserDeviceResponseBodySoftware {
	s.Inc = &v
	return s
}

func (s *ListSoftwareForUserDeviceResponseBodySoftware) SetInstallTime(v string) *ListSoftwareForUserDeviceResponseBodySoftware {
	s.InstallTime = &v
	return s
}

func (s *ListSoftwareForUserDeviceResponseBodySoftware) SetName(v string) *ListSoftwareForUserDeviceResponseBodySoftware {
	s.Name = &v
	return s
}

func (s *ListSoftwareForUserDeviceResponseBodySoftware) SetVersions(v []*string) *ListSoftwareForUserDeviceResponseBodySoftware {
	s.Versions = v
	return s
}

type ListSoftwareForUserDeviceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSoftwareForUserDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSoftwareForUserDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwareForUserDeviceResponse) GoString() string {
	return s.String()
}

func (s *ListSoftwareForUserDeviceResponse) SetHeaders(v map[string]*string) *ListSoftwareForUserDeviceResponse {
	s.Headers = v
	return s
}

func (s *ListSoftwareForUserDeviceResponse) SetStatusCode(v int32) *ListSoftwareForUserDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSoftwareForUserDeviceResponse) SetBody(v *ListSoftwareForUserDeviceResponseBody) *ListSoftwareForUserDeviceResponse {
	s.Body = v
	return s
}

type ListTagsForPrivateAccessApplicationRequest struct {
	// This parameter is required.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
}

func (s ListTagsForPrivateAccessApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagsForPrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessApplicationRequest) SetApplicationIds(v []*string) *ListTagsForPrivateAccessApplicationRequest {
	s.ApplicationIds = v
	return s
}

type ListTagsForPrivateAccessApplicationResponseBody struct {
	Applications []*ListTagsForPrivateAccessApplicationResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// 7241F45B-E8D3-5BA3-8172-8A58AC2AB0FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagsForPrivateAccessApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagsForPrivateAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessApplicationResponseBody) SetApplications(v []*ListTagsForPrivateAccessApplicationResponseBodyApplications) *ListTagsForPrivateAccessApplicationResponseBody {
	s.Applications = v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBody) SetRequestId(v string) *ListTagsForPrivateAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

type ListTagsForPrivateAccessApplicationResponseBodyApplications struct {
	// example:
	//
	// pa-application-7a4445897856****
	ApplicationId *string                                                            `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Tags          []*ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagsForPrivateAccessApplicationResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s ListTagsForPrivateAccessApplicationResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplications) SetApplicationId(v string) *ListTagsForPrivateAccessApplicationResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplications) SetTags(v []*ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) *ListTagsForPrivateAccessApplicationResponseBodyApplications {
	s.Tags = v
	return s
}

type ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags struct {
	// example:
	//
	// 2022-07-01 16:05:26
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// tag_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// tag-c0cb77857a99****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// Default
	TagType *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
}

func (s ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) SetCreateTime(v string) *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags {
	s.CreateTime = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) SetDescription(v string) *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags {
	s.Description = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) SetName(v string) *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags {
	s.Name = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) SetTagId(v string) *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags {
	s.TagId = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags) SetTagType(v string) *ListTagsForPrivateAccessApplicationResponseBodyApplicationsTags {
	s.TagType = &v
	return s
}

type ListTagsForPrivateAccessApplicationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagsForPrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagsForPrivateAccessApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagsForPrivateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessApplicationResponse) SetHeaders(v map[string]*string) *ListTagsForPrivateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponse) SetStatusCode(v int32) *ListTagsForPrivateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagsForPrivateAccessApplicationResponse) SetBody(v *ListTagsForPrivateAccessApplicationResponseBody) *ListTagsForPrivateAccessApplicationResponse {
	s.Body = v
	return s
}

type ListTagsForPrivateAccessPolicyRequest struct {
	// This parameter is required.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s ListTagsForPrivateAccessPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagsForPrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessPolicyRequest) SetPolicyIds(v []*string) *ListTagsForPrivateAccessPolicyRequest {
	s.PolicyIds = v
	return s
}

type ListTagsForPrivateAccessPolicyResponseBody struct {
	Polices []*ListTagsForPrivateAccessPolicyResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// example:
	//
	// 9EE61139-A6A8-5E13-80AF-83435C21B26B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagsForPrivateAccessPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagsForPrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessPolicyResponseBody) SetPolices(v []*ListTagsForPrivateAccessPolicyResponseBodyPolices) *ListTagsForPrivateAccessPolicyResponseBody {
	s.Polices = v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBody) SetRequestId(v string) *ListTagsForPrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ListTagsForPrivateAccessPolicyResponseBodyPolices struct {
	// example:
	//
	// pa-policy-1b0d0e8b4bcf****
	PolicyId *string                                                  `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Tags     []*ListTagsForPrivateAccessPolicyResponseBodyPolicesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagsForPrivateAccessPolicyResponseBodyPolices) String() string {
	return tea.Prettify(s)
}

func (s ListTagsForPrivateAccessPolicyResponseBodyPolices) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolices) SetPolicyId(v string) *ListTagsForPrivateAccessPolicyResponseBodyPolices {
	s.PolicyId = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolices) SetTags(v []*ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) *ListTagsForPrivateAccessPolicyResponseBodyPolices {
	s.Tags = v
	return s
}

type ListTagsForPrivateAccessPolicyResponseBodyPolicesTags struct {
	// 内网访问标签创建时间。
	//
	// example:
	//
	// 2023-02-21 14:10:16
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// tag_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// example:
	//
	// Default
	TagType *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
}

func (s ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) SetCreateTime(v string) *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags {
	s.CreateTime = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) SetDescription(v string) *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags {
	s.Description = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) SetName(v string) *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags {
	s.Name = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) SetTagId(v string) *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags {
	s.TagId = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags) SetTagType(v string) *ListTagsForPrivateAccessPolicyResponseBodyPolicesTags {
	s.TagType = &v
	return s
}

type ListTagsForPrivateAccessPolicyResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagsForPrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagsForPrivateAccessPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagsForPrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListTagsForPrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *ListTagsForPrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponse) SetStatusCode(v int32) *ListTagsForPrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagsForPrivateAccessPolicyResponse) SetBody(v *ListTagsForPrivateAccessPolicyResponseBody) *ListTagsForPrivateAccessPolicyResponse {
	s.Body = v
	return s
}

type ListUserApplicationsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// private_access_application_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
}

func (s ListUserApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListUserApplicationsRequest) SetCurrentPage(v int32) *ListUserApplicationsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserApplicationsRequest) SetName(v string) *ListUserApplicationsRequest {
	s.Name = &v
	return s
}

func (s *ListUserApplicationsRequest) SetPageSize(v int32) *ListUserApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserApplicationsRequest) SetSaseUserId(v string) *ListUserApplicationsRequest {
	s.SaseUserId = &v
	return s
}

type ListUserApplicationsResponseBody struct {
	Applications []*ListUserApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListUserApplicationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserApplicationsResponseBody) SetApplications(v []*ListUserApplicationsResponseBodyApplications) *ListUserApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListUserApplicationsResponseBody) SetRequestId(v string) *ListUserApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserApplicationsResponseBody) SetTotalNum(v int32) *ListUserApplicationsResponseBody {
	s.TotalNum = &v
	return s
}

type ListUserApplicationsResponseBodyApplications struct {
	// example:
	//
	// Block
	Action    *string   `json:"Action,omitempty" xml:"Action,omitempty"`
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// example:
	//
	// pa-application-b927baf3e592****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// private_access_application_name
	Name       *string                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges []*ListUserApplicationsResponseBodyApplicationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s ListUserApplicationsResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s ListUserApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListUserApplicationsResponseBodyApplications) SetAction(v string) *ListUserApplicationsResponseBodyApplications {
	s.Action = &v
	return s
}

func (s *ListUserApplicationsResponseBodyApplications) SetAddresses(v []*string) *ListUserApplicationsResponseBodyApplications {
	s.Addresses = v
	return s
}

func (s *ListUserApplicationsResponseBodyApplications) SetApplicationId(v string) *ListUserApplicationsResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListUserApplicationsResponseBodyApplications) SetName(v string) *ListUserApplicationsResponseBodyApplications {
	s.Name = &v
	return s
}

func (s *ListUserApplicationsResponseBodyApplications) SetPortRanges(v []*ListUserApplicationsResponseBodyApplicationsPortRanges) *ListUserApplicationsResponseBodyApplications {
	s.PortRanges = v
	return s
}

func (s *ListUserApplicationsResponseBodyApplications) SetProtocol(v string) *ListUserApplicationsResponseBodyApplications {
	s.Protocol = &v
	return s
}

type ListUserApplicationsResponseBodyApplicationsPortRanges struct {
	// example:
	//
	// 80
	Begin *string `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 81
	End *string `json:"End,omitempty" xml:"End,omitempty"`
}

func (s ListUserApplicationsResponseBodyApplicationsPortRanges) String() string {
	return tea.Prettify(s)
}

func (s ListUserApplicationsResponseBodyApplicationsPortRanges) GoString() string {
	return s.String()
}

func (s *ListUserApplicationsResponseBodyApplicationsPortRanges) SetBegin(v string) *ListUserApplicationsResponseBodyApplicationsPortRanges {
	s.Begin = &v
	return s
}

func (s *ListUserApplicationsResponseBodyApplicationsPortRanges) SetEnd(v string) *ListUserApplicationsResponseBodyApplicationsPortRanges {
	s.End = &v
	return s
}

type ListUserApplicationsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListUserApplicationsResponse) SetHeaders(v map[string]*string) *ListUserApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListUserApplicationsResponse) SetStatusCode(v int32) *ListUserApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserApplicationsResponse) SetBody(v *ListUserApplicationsResponseBody) *ListUserApplicationsResponse {
	s.Body = v
	return s
}

type ListUserDevicesRequest struct {
	AppStatuses []*string `json:"AppStatuses,omitempty" xml:"AppStatuses,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Department  *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// Company
	DeviceBelong   *string   `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	DeviceStatuses []*string `json:"DeviceStatuses,omitempty" xml:"DeviceStatuses,omitempty" type:"Repeated"`
	DeviceTags     []*string `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty" type:"Repeated"`
	DeviceTypes    []*string `json:"DeviceTypes,omitempty" xml:"DeviceTypes,omitempty" type:"Repeated"`
	DlpStatuses    []*string `json:"DlpStatuses,omitempty" xml:"DlpStatuses,omitempty" type:"Repeated"`
	// example:
	//
	// win10-64bit
	Hostname   *string   `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	IaStatuses []*string `json:"IaStatuses,omitempty" xml:"IaStatuses,omitempty" type:"Repeated"`
	InnerIp    *string   `json:"InnerIp,omitempty" xml:"InnerIp,omitempty"`
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac         *string   `json:"Mac,omitempty" xml:"Mac,omitempty"`
	NacStatuses []*string `json:"NacStatuses,omitempty" xml:"NacStatuses,omitempty" type:"Repeated"`
	PaStatuses  []*string `json:"PaStatuses,omitempty" xml:"PaStatuses,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// example:
	//
	// true
	SharingStatus *bool   `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	SortBy        *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Username      *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListUserDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserDevicesRequest) GoString() string {
	return s.String()
}

func (s *ListUserDevicesRequest) SetAppStatuses(v []*string) *ListUserDevicesRequest {
	s.AppStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetCurrentPage(v int64) *ListUserDevicesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserDevicesRequest) SetDepartment(v string) *ListUserDevicesRequest {
	s.Department = &v
	return s
}

func (s *ListUserDevicesRequest) SetDeviceBelong(v string) *ListUserDevicesRequest {
	s.DeviceBelong = &v
	return s
}

func (s *ListUserDevicesRequest) SetDeviceStatuses(v []*string) *ListUserDevicesRequest {
	s.DeviceStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetDeviceTags(v []*string) *ListUserDevicesRequest {
	s.DeviceTags = v
	return s
}

func (s *ListUserDevicesRequest) SetDeviceTypes(v []*string) *ListUserDevicesRequest {
	s.DeviceTypes = v
	return s
}

func (s *ListUserDevicesRequest) SetDlpStatuses(v []*string) *ListUserDevicesRequest {
	s.DlpStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetHostname(v string) *ListUserDevicesRequest {
	s.Hostname = &v
	return s
}

func (s *ListUserDevicesRequest) SetIaStatuses(v []*string) *ListUserDevicesRequest {
	s.IaStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetInnerIp(v string) *ListUserDevicesRequest {
	s.InnerIp = &v
	return s
}

func (s *ListUserDevicesRequest) SetMac(v string) *ListUserDevicesRequest {
	s.Mac = &v
	return s
}

func (s *ListUserDevicesRequest) SetNacStatuses(v []*string) *ListUserDevicesRequest {
	s.NacStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetPaStatuses(v []*string) *ListUserDevicesRequest {
	s.PaStatuses = v
	return s
}

func (s *ListUserDevicesRequest) SetPageSize(v int64) *ListUserDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserDevicesRequest) SetSaseUserId(v string) *ListUserDevicesRequest {
	s.SaseUserId = &v
	return s
}

func (s *ListUserDevicesRequest) SetSharingStatus(v bool) *ListUserDevicesRequest {
	s.SharingStatus = &v
	return s
}

func (s *ListUserDevicesRequest) SetSortBy(v string) *ListUserDevicesRequest {
	s.SortBy = &v
	return s
}

func (s *ListUserDevicesRequest) SetUsername(v string) *ListUserDevicesRequest {
	s.Username = &v
	return s
}

type ListUserDevicesResponseBody struct {
	Devices []*ListUserDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// example:
	//
	// 5FEF5CFA-14CC-5DE5-BD1F-AFFE0996E71D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListUserDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDevicesResponseBody) SetDevices(v []*ListUserDevicesResponseBodyDevices) *ListUserDevicesResponseBody {
	s.Devices = v
	return s
}

func (s *ListUserDevicesResponseBody) SetRequestId(v string) *ListUserDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDevicesResponseBody) SetTotalNum(v int64) *ListUserDevicesResponseBody {
	s.TotalNum = &v
	return s
}

type ListUserDevicesResponseBodyDevices struct {
	// example:
	//
	// Online
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// example:
	//
	// 2.2.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// Apple M1
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 2023-07-17 18:46:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// Company
	DeviceBelong *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	// example:
	//
	// MacBookPro17,1
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	// example:
	//
	// Online
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// example:
	//
	// Windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 3.5.1
	DeviceVersion *string `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	// example:
	//
	// APPLE SSD AP0512Q Media
	Disk *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// example:
	//
	// Enabled
	DlpStatus *string `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// Enabled
	IaStatus *string `json:"IaStatus,omitempty" xml:"IaStatus,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	InnerIP *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 16
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// Enabled
	NacStatus *string `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	// example:
	//
	// Enabled
	PaStatus *string `json:"PaStatus,omitempty" xml:"PaStatus,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// example:
	//
	// true
	SharingStatus *bool `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	// example:
	//
	// 11.49.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// example:
	//
	// 2023-08-24 19:04:42
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListUserDevicesResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s ListUserDevicesResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *ListUserDevicesResponseBodyDevices) SetAppStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.AppStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetAppVersion(v string) *ListUserDevicesResponseBodyDevices {
	s.AppVersion = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetCPU(v string) *ListUserDevicesResponseBodyDevices {
	s.CPU = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetCreateTime(v string) *ListUserDevicesResponseBodyDevices {
	s.CreateTime = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDepartment(v string) *ListUserDevicesResponseBodyDevices {
	s.Department = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDeviceBelong(v string) *ListUserDevicesResponseBodyDevices {
	s.DeviceBelong = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDeviceModel(v string) *ListUserDevicesResponseBodyDevices {
	s.DeviceModel = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDeviceStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.DeviceStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDeviceTag(v string) *ListUserDevicesResponseBodyDevices {
	s.DeviceTag = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDeviceType(v string) *ListUserDevicesResponseBodyDevices {
	s.DeviceType = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDeviceVersion(v string) *ListUserDevicesResponseBodyDevices {
	s.DeviceVersion = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDisk(v string) *ListUserDevicesResponseBodyDevices {
	s.Disk = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetDlpStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.DlpStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetHostname(v string) *ListUserDevicesResponseBodyDevices {
	s.Hostname = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetIaStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.IaStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetInnerIP(v string) *ListUserDevicesResponseBodyDevices {
	s.InnerIP = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetMac(v string) *ListUserDevicesResponseBodyDevices {
	s.Mac = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetMemory(v string) *ListUserDevicesResponseBodyDevices {
	s.Memory = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetNacStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.NacStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetPaStatus(v string) *ListUserDevicesResponseBodyDevices {
	s.PaStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetSaseUserId(v string) *ListUserDevicesResponseBodyDevices {
	s.SaseUserId = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetSharingStatus(v bool) *ListUserDevicesResponseBodyDevices {
	s.SharingStatus = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetSrcIP(v string) *ListUserDevicesResponseBodyDevices {
	s.SrcIP = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetUpdateTime(v string) *ListUserDevicesResponseBodyDevices {
	s.UpdateTime = &v
	return s
}

func (s *ListUserDevicesResponseBodyDevices) SetUsername(v string) *ListUserDevicesResponseBodyDevices {
	s.Username = &v
	return s
}

type ListUserDevicesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserDevicesResponse) GoString() string {
	return s.String()
}

func (s *ListUserDevicesResponse) SetHeaders(v map[string]*string) *ListUserDevicesResponse {
	s.Headers = v
	return s
}

func (s *ListUserDevicesResponse) SetStatusCode(v int32) *ListUserDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserDevicesResponse) SetBody(v *ListUserDevicesResponseBody) *ListUserDevicesResponse {
	s.Body = v
	return s
}

type ListUserGroupsRequest struct {
	// example:
	//
	// username
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 用户组名称。长度为1~128个字符，支持中文和大小写英文字母，可包含数字、半角句号（.）、下划线（_）和短划线（-）。
	//
	// example:
	//
	// user_group_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// pa-policy-54a7838a48bf****
	PAPolicyId *string `json:"PAPolicyId,omitempty" xml:"PAPolicyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
}

func (s ListUserGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsRequest) SetAttributeValue(v string) *ListUserGroupsRequest {
	s.AttributeValue = &v
	return s
}

func (s *ListUserGroupsRequest) SetCurrentPage(v int32) *ListUserGroupsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserGroupsRequest) SetName(v string) *ListUserGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListUserGroupsRequest) SetPAPolicyId(v string) *ListUserGroupsRequest {
	s.PAPolicyId = &v
	return s
}

func (s *ListUserGroupsRequest) SetPageSize(v int32) *ListUserGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserGroupsRequest) SetUserGroupIds(v []*string) *ListUserGroupsRequest {
	s.UserGroupIds = v
	return s
}

type ListUserGroupsResponseBody struct {
	// example:
	//
	// 4AB972E2-D702-5464-B132-B1911498B8BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum   *int32                                  `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	UserGroups []*ListUserGroupsResponseBodyUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListUserGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBody) SetRequestId(v string) *ListUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetTotalNum(v int32) *ListUserGroupsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListUserGroupsResponseBody) SetUserGroups(v []*ListUserGroupsResponseBodyUserGroups) *ListUserGroupsResponseBody {
	s.UserGroups = v
	return s
}

type ListUserGroupsResponseBodyUserGroups struct {
	Attributes []*ListUserGroupsResponseBodyUserGroupsAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-10-10 11:39:22
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// user_group_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListUserGroupsResponseBodyUserGroups) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseBodyUserGroups) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyUserGroups) SetAttributes(v []*ListUserGroupsResponseBodyUserGroupsAttributes) *ListUserGroupsResponseBodyUserGroups {
	s.Attributes = v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetCreateTime(v string) *ListUserGroupsResponseBodyUserGroups {
	s.CreateTime = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetDescription(v string) *ListUserGroupsResponseBodyUserGroups {
	s.Description = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetName(v string) *ListUserGroupsResponseBodyUserGroups {
	s.Name = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroups) SetUserGroupId(v string) *ListUserGroupsResponseBodyUserGroups {
	s.UserGroupId = &v
	return s
}

type ListUserGroupsResponseBodyUserGroupsAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListUserGroupsResponseBodyUserGroupsAttributes) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponseBodyUserGroupsAttributes) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponseBodyUserGroupsAttributes) SetIdpId(v int32) *ListUserGroupsResponseBodyUserGroupsAttributes {
	s.IdpId = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroupsAttributes) SetRelation(v string) *ListUserGroupsResponseBodyUserGroupsAttributes {
	s.Relation = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroupsAttributes) SetUserGroupType(v string) *ListUserGroupsResponseBodyUserGroupsAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListUserGroupsResponseBodyUserGroupsAttributes) SetValue(v string) *ListUserGroupsResponseBodyUserGroupsAttributes {
	s.Value = &v
	return s
}

type ListUserGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsResponse) SetHeaders(v map[string]*string) *ListUserGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsResponse) SetStatusCode(v int32) *ListUserGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupsResponse) SetBody(v *ListUserGroupsResponseBody) *ListUserGroupsResponse {
	s.Body = v
	return s
}

type ListUserGroupsForPrivateAccessPolicyRequest struct {
	// This parameter is required.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s ListUserGroupsForPrivateAccessPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsForPrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForPrivateAccessPolicyRequest) SetPolicyIds(v []*string) *ListUserGroupsForPrivateAccessPolicyRequest {
	s.PolicyIds = v
	return s
}

type ListUserGroupsForPrivateAccessPolicyResponseBody struct {
	Polices []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// example:
	//
	// 58D6B23E-E5DA-5418-8F61-51A3B5A30049
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBody) SetPolices(v []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) *ListUserGroupsForPrivateAccessPolicyResponseBody {
	s.Polices = v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBody) SetRequestId(v string) *ListUserGroupsForPrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ListUserGroupsForPrivateAccessPolicyResponseBodyPolices struct {
	// example:
	//
	// pa-policy-1b0d0e8b4bcf****
	PolicyId   *string                                                              `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	UserGroups []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) SetPolicyId(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolices {
	s.PolicyId = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolices) SetUserGroups(v []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolices {
	s.UserGroups = v
	return s
}

type ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups struct {
	Attributes []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// 用户组创建时间。
	//
	// example:
	//
	// 2022-09-27 18:10:25
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// user_group_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) SetAttributes(v []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups {
	s.Attributes = v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) SetCreateTime(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups {
	s.CreateTime = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) SetDescription(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups {
	s.Description = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) SetName(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups {
	s.Name = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups) SetUserGroupId(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroups {
	s.UserGroupId = &v
	return s
}

type ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) SetIdpId(v int32) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes {
	s.IdpId = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) SetRelation(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes {
	s.Relation = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) SetUserGroupType(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes) SetValue(v string) *ListUserGroupsForPrivateAccessPolicyResponseBodyPolicesUserGroupsAttributes {
	s.Value = &v
	return s
}

type ListUserGroupsForPrivateAccessPolicyResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupsForPrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserGroupsForPrivateAccessPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsForPrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForPrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *ListUserGroupsForPrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponse) SetStatusCode(v int32) *ListUserGroupsForPrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupsForPrivateAccessPolicyResponse) SetBody(v *ListUserGroupsForPrivateAccessPolicyResponseBody) *ListUserGroupsForPrivateAccessPolicyResponse {
	s.Body = v
	return s
}

type ListUserGroupsForRegistrationPolicyRequest struct {
	// This parameter is required.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s ListUserGroupsForRegistrationPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsForRegistrationPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForRegistrationPolicyRequest) SetPolicyIds(v []*string) *ListUserGroupsForRegistrationPolicyRequest {
	s.PolicyIds = v
	return s
}

type ListUserGroupsForRegistrationPolicyResponseBody struct {
	Policies []*ListUserGroupsForRegistrationPolicyResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// example:
	//
	// D89009C7-54C6-51B6-BAE7-3F373920C6BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserGroupsForRegistrationPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsForRegistrationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForRegistrationPolicyResponseBody) SetPolicies(v []*ListUserGroupsForRegistrationPolicyResponseBodyPolicies) *ListUserGroupsForRegistrationPolicyResponseBody {
	s.Policies = v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBody) SetRequestId(v string) *ListUserGroupsForRegistrationPolicyResponseBody {
	s.RequestId = &v
	return s
}

type ListUserGroupsForRegistrationPolicyResponseBodyPolicies struct {
	// example:
	//
	// reg-policy-f25c9e5872e5****
	PolicyId   *string                                                              `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	UserGroups []*ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups `json:"UserGroups,omitempty" xml:"UserGroups,omitempty" type:"Repeated"`
}

func (s ListUserGroupsForRegistrationPolicyResponseBodyPolicies) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsForRegistrationPolicyResponseBodyPolicies) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPolicies) SetPolicyId(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPolicies {
	s.PolicyId = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPolicies) SetUserGroups(v []*ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) *ListUserGroupsForRegistrationPolicyResponseBodyPolicies {
	s.UserGroups = v
	return s
}

type ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups struct {
	Attributes []*ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-09-27 18:10:25
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// user_group_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) SetAttributes(v []*ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups {
	s.Attributes = v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) SetCreateTime(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups {
	s.CreateTime = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) SetDescription(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups {
	s.Description = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) SetName(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups {
	s.Name = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups) SetUserGroupId(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroups {
	s.UserGroupId = &v
	return s
}

type ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) SetIdpId(v int32) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes {
	s.IdpId = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) SetRelation(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes {
	s.Relation = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) SetUserGroupType(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes) SetValue(v string) *ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes {
	s.Value = &v
	return s
}

type ListUserGroupsForRegistrationPolicyResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserGroupsForRegistrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserGroupsForRegistrationPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserGroupsForRegistrationPolicyResponse) GoString() string {
	return s.String()
}

func (s *ListUserGroupsForRegistrationPolicyResponse) SetHeaders(v map[string]*string) *ListUserGroupsForRegistrationPolicyResponse {
	s.Headers = v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponse) SetStatusCode(v int32) *ListUserGroupsForRegistrationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserGroupsForRegistrationPolicyResponse) SetBody(v *ListUserGroupsForRegistrationPolicyResponseBody) *ListUserGroupsForRegistrationPolicyResponse {
	s.Body = v
	return s
}

type ListUserPrivateAccessPoliciesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
}

func (s ListUserPrivateAccessPoliciesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserPrivateAccessPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListUserPrivateAccessPoliciesRequest) SetCurrentPage(v int32) *ListUserPrivateAccessPoliciesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesRequest) SetName(v string) *ListUserPrivateAccessPoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesRequest) SetPageSize(v int32) *ListUserPrivateAccessPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesRequest) SetSaseUserId(v string) *ListUserPrivateAccessPoliciesRequest {
	s.SaseUserId = &v
	return s
}

type ListUserPrivateAccessPoliciesResponseBody struct {
	Polices []*ListUserPrivateAccessPoliciesResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 9D852F87-AFB5-51B8-AACD-F7D0EFB8277D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListUserPrivateAccessPoliciesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserPrivateAccessPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserPrivateAccessPoliciesResponseBody) SetPolices(v []*ListUserPrivateAccessPoliciesResponseBodyPolices) *ListUserPrivateAccessPoliciesResponseBody {
	s.Polices = v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBody) SetRequestId(v string) *ListUserPrivateAccessPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBody) SetTotalNum(v int32) *ListUserPrivateAccessPoliciesResponseBody {
	s.TotalNum = &v
	return s
}

type ListUserPrivateAccessPoliciesResponseBodyPolices struct {
	CustomUserAttributes []*ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	// example:
	//
	// device_attribute_name
	DeviceAttributeName *string `json:"DeviceAttributeName,omitempty" xml:"DeviceAttributeName,omitempty"`
	// example:
	//
	// user_group_name
	MatchedUserGroup *string `json:"MatchedUserGroup,omitempty" xml:"MatchedUserGroup,omitempty"`
	// example:
	//
	// private_access_policy_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// example:
	//
	// pa-policy-1b0d0e8b4bcf****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Custom
	UserGroupMode *string `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
}

func (s ListUserPrivateAccessPoliciesResponseBodyPolices) String() string {
	return tea.Prettify(s)
}

func (s ListUserPrivateAccessPoliciesResponseBodyPolices) GoString() string {
	return s.String()
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetCustomUserAttributes(v []*ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.CustomUserAttributes = v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetDeviceAttributeName(v string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.DeviceAttributeName = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetMatchedUserGroup(v string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.MatchedUserGroup = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetName(v string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.Name = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetPolicyAction(v string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.PolicyAction = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetPolicyId(v string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.PolicyId = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetPriority(v int64) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.Priority = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolices) SetUserGroupMode(v string) *ListUserPrivateAccessPoliciesResponseBodyPolices {
	s.UserGroupMode = &v
	return s
}

type ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) String() string {
	return tea.Prettify(s)
}

func (s ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) SetIdpId(v int32) *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) SetRelation(v string) *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) SetUserGroupType(v string) *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes) SetValue(v string) *ListUserPrivateAccessPoliciesResponseBodyPolicesCustomUserAttributes {
	s.Value = &v
	return s
}

type ListUserPrivateAccessPoliciesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserPrivateAccessPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserPrivateAccessPoliciesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserPrivateAccessPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListUserPrivateAccessPoliciesResponse) SetHeaders(v map[string]*string) *ListUserPrivateAccessPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponse) SetStatusCode(v int32) *ListUserPrivateAccessPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserPrivateAccessPoliciesResponse) SetBody(v *ListUserPrivateAccessPoliciesResponseBody) *ListUserPrivateAccessPoliciesResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage   *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Department    *string `json:"Department,omitempty" xml:"Department,omitempty"`
	FuzzyUsername *string `json:"FuzzyUsername,omitempty" xml:"FuzzyUsername,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize        *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PreciseUsername *string   `json:"PreciseUsername,omitempty" xml:"PreciseUsername,omitempty"`
	SaseUserIds     []*string `json:"SaseUserIds,omitempty" xml:"SaseUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetCurrentPage(v int64) *ListUsersRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUsersRequest) SetDepartment(v string) *ListUsersRequest {
	s.Department = &v
	return s
}

func (s *ListUsersRequest) SetFuzzyUsername(v string) *ListUsersRequest {
	s.FuzzyUsername = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int64) *ListUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListUsersRequest) SetPreciseUsername(v string) *ListUsersRequest {
	s.PreciseUsername = &v
	return s
}

func (s *ListUsersRequest) SetSaseUserIds(v []*string) *ListUsersRequest {
	s.SaseUserIds = v
	return s
}

func (s *ListUsersRequest) SetStatus(v string) *ListUsersRequest {
	s.Status = &v
	return s
}

type ListUsersResponseBody struct {
	// example:
	//
	// 5FEF5CFA-14CC-5DE5-BD1F-AFFE0996E71D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *string                       `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	Users    []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetTotalNum(v string) *ListUsersResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponseBodyUsers struct {
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// a***@example.net
	Email   *string `json:"Email,omitempty" xml:"Email,omitempty"`
	IdpName *string `json:"IdpName,omitempty" xml:"IdpName,omitempty"`
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// example:
	//
	// Enabled
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) SetDepartment(v string) *ListUsersResponseBodyUsers {
	s.Department = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetEmail(v string) *ListUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetIdpName(v string) *ListUsersResponseBodyUsers {
	s.IdpName = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetPhone(v string) *ListUsersResponseBodyUsers {
	s.Phone = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetSaseUserId(v string) *ListUsersResponseBodyUsers {
	s.SaseUserId = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetStatus(v string) *ListUsersResponseBodyUsers {
	s.Status = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetUsername(v string) *ListUsersResponseBodyUsers {
	s.Username = &v
	return s
}

type ListUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetStatusCode(v int32) *ListUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type LookupWmInfoMappingRequest struct {
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123***
	WmInfoUint *string `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PureDocument
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s LookupWmInfoMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s LookupWmInfoMappingRequest) GoString() string {
	return s.String()
}

func (s *LookupWmInfoMappingRequest) SetWmInfoSize(v int64) *LookupWmInfoMappingRequest {
	s.WmInfoSize = &v
	return s
}

func (s *LookupWmInfoMappingRequest) SetWmInfoUint(v string) *LookupWmInfoMappingRequest {
	s.WmInfoUint = &v
	return s
}

func (s *LookupWmInfoMappingRequest) SetWmType(v string) *LookupWmInfoMappingRequest {
	s.WmType = &v
	return s
}

type LookupWmInfoMappingResponseBody struct {
	Data *LookupWmInfoMappingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 7E9D7ACD-53D5-56EF-A913-79D148D06299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LookupWmInfoMappingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LookupWmInfoMappingResponseBody) GoString() string {
	return s.String()
}

func (s *LookupWmInfoMappingResponseBody) SetData(v *LookupWmInfoMappingResponseBodyData) *LookupWmInfoMappingResponseBody {
	s.Data = v
	return s
}

func (s *LookupWmInfoMappingResponseBody) SetRequestId(v string) *LookupWmInfoMappingResponseBody {
	s.RequestId = &v
	return s
}

type LookupWmInfoMappingResponseBodyData struct {
	// example:
	//
	// aGVsbG8gc2F*****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
}

func (s LookupWmInfoMappingResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s LookupWmInfoMappingResponseBodyData) GoString() string {
	return s.String()
}

func (s *LookupWmInfoMappingResponseBodyData) SetWmInfoBytesB64(v string) *LookupWmInfoMappingResponseBodyData {
	s.WmInfoBytesB64 = &v
	return s
}

type LookupWmInfoMappingResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LookupWmInfoMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LookupWmInfoMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s LookupWmInfoMappingResponse) GoString() string {
	return s.String()
}

func (s *LookupWmInfoMappingResponse) SetHeaders(v map[string]*string) *LookupWmInfoMappingResponse {
	s.Headers = v
	return s
}

func (s *LookupWmInfoMappingResponse) SetStatusCode(v int32) *LookupWmInfoMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *LookupWmInfoMappingResponse) SetBody(v *LookupWmInfoMappingResponseBody) *LookupWmInfoMappingResponse {
	s.Body = v
	return s
}

type RevokeUserSessionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	ExternalIds *string `json:"ExternalIds,omitempty" xml:"ExternalIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// idp-cfg9vcrqylo39c39uxnw
	IdpId *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
}

func (s RevokeUserSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeUserSessionRequest) GoString() string {
	return s.String()
}

func (s *RevokeUserSessionRequest) SetExternalIds(v string) *RevokeUserSessionRequest {
	s.ExternalIds = &v
	return s
}

func (s *RevokeUserSessionRequest) SetIdpId(v string) *RevokeUserSessionRequest {
	s.IdpId = &v
	return s
}

type RevokeUserSessionResponseBody struct {
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeUserSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeUserSessionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeUserSessionResponseBody) SetRequestId(v string) *RevokeUserSessionResponseBody {
	s.RequestId = &v
	return s
}

type RevokeUserSessionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeUserSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeUserSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeUserSessionResponse) GoString() string {
	return s.String()
}

func (s *RevokeUserSessionResponse) SetHeaders(v map[string]*string) *RevokeUserSessionResponse {
	s.Headers = v
	return s
}

func (s *RevokeUserSessionResponse) SetStatusCode(v int32) *RevokeUserSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeUserSessionResponse) SetBody(v *RevokeUserSessionResponseBody) *RevokeUserSessionResponse {
	s.Body = v
	return s
}

type UpdateClientUserRequest struct {
	// example:
	//
	// 10701
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// johndoe@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20644
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 13641966835
	MobileNumber *string `json:"MobileNumber,omitempty" xml:"MobileNumber,omitempty"`
}

func (s UpdateClientUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateClientUserRequest) SetDepartmentId(v string) *UpdateClientUserRequest {
	s.DepartmentId = &v
	return s
}

func (s *UpdateClientUserRequest) SetDescription(v string) *UpdateClientUserRequest {
	s.Description = &v
	return s
}

func (s *UpdateClientUserRequest) SetEmail(v string) *UpdateClientUserRequest {
	s.Email = &v
	return s
}

func (s *UpdateClientUserRequest) SetId(v string) *UpdateClientUserRequest {
	s.Id = &v
	return s
}

func (s *UpdateClientUserRequest) SetMobileNumber(v string) *UpdateClientUserRequest {
	s.MobileNumber = &v
	return s
}

type UpdateClientUserResponseBody struct {
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClientUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClientUserResponseBody) SetRequestId(v string) *UpdateClientUserResponseBody {
	s.RequestId = &v
	return s
}

type UpdateClientUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClientUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClientUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateClientUserResponse) SetHeaders(v map[string]*string) *UpdateClientUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateClientUserResponse) SetStatusCode(v int32) *UpdateClientUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClientUserResponse) SetBody(v *UpdateClientUserResponseBody) *UpdateClientUserResponse {
	s.Body = v
	return s
}

type UpdateClientUserPasswordRequest struct {
	// example:
	//
	// 1128
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// kehudiyidj
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateClientUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *UpdateClientUserPasswordRequest) SetId(v string) *UpdateClientUserPasswordRequest {
	s.Id = &v
	return s
}

func (s *UpdateClientUserPasswordRequest) SetPassword(v string) *UpdateClientUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *UpdateClientUserPasswordRequest) SetUsername(v string) *UpdateClientUserPasswordRequest {
	s.Username = &v
	return s
}

type UpdateClientUserPasswordResponseBody struct {
	// example:
	//
	// EFE7EBB2-449D-5BBB-B381-CA7839BC1649
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClientUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClientUserPasswordResponseBody) SetRequestId(v string) *UpdateClientUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

type UpdateClientUserPasswordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClientUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClientUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *UpdateClientUserPasswordResponse) SetHeaders(v map[string]*string) *UpdateClientUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *UpdateClientUserPasswordResponse) SetStatusCode(v int32) *UpdateClientUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClientUserPasswordResponse) SetBody(v *UpdateClientUserPasswordResponseBody) *UpdateClientUserPasswordResponse {
	s.Body = v
	return s
}

type UpdateClientUserStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1495
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateClientUserStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientUserStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateClientUserStatusRequest) SetId(v string) *UpdateClientUserStatusRequest {
	s.Id = &v
	return s
}

func (s *UpdateClientUserStatusRequest) SetStatus(v string) *UpdateClientUserStatusRequest {
	s.Status = &v
	return s
}

type UpdateClientUserStatusResponseBody struct {
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClientUserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClientUserStatusResponseBody) SetRequestId(v string) *UpdateClientUserStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateClientUserStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateClientUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateClientUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClientUserStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateClientUserStatusResponse) SetHeaders(v map[string]*string) *UpdateClientUserStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateClientUserStatusResponse) SetStatusCode(v int32) *UpdateClientUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateClientUserStatusResponse) SetBody(v *UpdateClientUserStatusResponseBody) *UpdateClientUserStatusResponse {
	s.Body = v
	return s
}

type UpdateDynamicRouteRequest struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// example:
	//
	// Application
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dr-ca9fddfac7c6****
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	// example:
	//
	// connector
	DynamicRouteType *string `json:"DynamicRouteType,omitempty" xml:"DynamicRouteType,omitempty"`
	// example:
	//
	// Cover
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// example:
	//
	// dynamic_route_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// connector-8ccb13b6f52c****
	NextHop *string `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	// example:
	//
	// 99
	Priority  *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RegionIds []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	// example:
	//
	// Disabled
	Status *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s UpdateDynamicRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDynamicRouteRequest) GoString() string {
	return s.String()
}

func (s *UpdateDynamicRouteRequest) SetApplicationIds(v []*string) *UpdateDynamicRouteRequest {
	s.ApplicationIds = v
	return s
}

func (s *UpdateDynamicRouteRequest) SetApplicationType(v string) *UpdateDynamicRouteRequest {
	s.ApplicationType = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetDescription(v string) *UpdateDynamicRouteRequest {
	s.Description = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetDynamicRouteId(v string) *UpdateDynamicRouteRequest {
	s.DynamicRouteId = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetDynamicRouteType(v string) *UpdateDynamicRouteRequest {
	s.DynamicRouteType = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetModifyType(v string) *UpdateDynamicRouteRequest {
	s.ModifyType = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetName(v string) *UpdateDynamicRouteRequest {
	s.Name = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetNextHop(v string) *UpdateDynamicRouteRequest {
	s.NextHop = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetPriority(v int32) *UpdateDynamicRouteRequest {
	s.Priority = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetRegionIds(v []*string) *UpdateDynamicRouteRequest {
	s.RegionIds = v
	return s
}

func (s *UpdateDynamicRouteRequest) SetStatus(v string) *UpdateDynamicRouteRequest {
	s.Status = &v
	return s
}

func (s *UpdateDynamicRouteRequest) SetTagIds(v []*string) *UpdateDynamicRouteRequest {
	s.TagIds = v
	return s
}

type UpdateDynamicRouteResponseBody struct {
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDynamicRouteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDynamicRouteResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDynamicRouteResponseBody) SetRequestId(v string) *UpdateDynamicRouteResponseBody {
	s.RequestId = &v
	return s
}

type UpdateDynamicRouteResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDynamicRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDynamicRouteResponse) GoString() string {
	return s.String()
}

func (s *UpdateDynamicRouteResponse) SetHeaders(v map[string]*string) *UpdateDynamicRouteResponse {
	s.Headers = v
	return s
}

func (s *UpdateDynamicRouteResponse) SetStatusCode(v int32) *UpdateDynamicRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDynamicRouteResponse) SetBody(v *UpdateDynamicRouteResponseBody) *UpdateDynamicRouteResponse {
	s.Body = v
	return s
}

type UpdateExcessiveDeviceRegistrationApplicationsStatusRequest struct {
	// This parameter is required.
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Approved
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) SetApplicationIds(v []*string) *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest {
	s.ApplicationIds = v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) SetStatus(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest {
	s.Status = &v
	return s
}

type UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody struct {
	Applications []*UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) SetApplications(v []*UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody {
	s.Applications = v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) SetRequestId(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications struct {
	// example:
	//
	// reg-application-0f4a127b7e78****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// 2023-07-17 18:46:55
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department  *string `json:"Department,omitempty" xml:"Department,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// example:
	//
	// Windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// false
	IsUsed *bool `json:"IsUsed,omitempty" xml:"IsUsed,omitempty"`
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// example:
	//
	// Approved
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) String() string {
	return tea.Prettify(s)
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetApplicationId(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetCreateTime(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.CreateTime = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetDepartment(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.Department = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetDescription(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.Description = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetDeviceTag(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.DeviceTag = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetDeviceType(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.DeviceType = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetHostname(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.Hostname = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetIsUsed(v bool) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.IsUsed = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetMac(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.Mac = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetSaseUserId(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.SaseUserId = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetStatus(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.Status = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications) SetUsername(v string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBodyApplications {
	s.Username = &v
	return s
}

type UpdateExcessiveDeviceRegistrationApplicationsStatusResponse struct {
	Headers    map[string]*string                                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) SetHeaders(v map[string]*string) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) SetStatusCode(v int32) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse) SetBody(v *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody) *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse {
	s.Body = v
	return s
}

type UpdateIdpDepartmentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10653
	DepartmentId *string `json:"DepartmentId,omitempty" xml:"DepartmentId,omitempty"`
	// This parameter is required.
	DepartmentName *string `json:"DepartmentName,omitempty" xml:"DepartmentName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 598
	IdpConfigId *string `json:"IdpConfigId,omitempty" xml:"IdpConfigId,omitempty"`
}

func (s UpdateIdpDepartmentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIdpDepartmentRequest) GoString() string {
	return s.String()
}

func (s *UpdateIdpDepartmentRequest) SetDepartmentId(v string) *UpdateIdpDepartmentRequest {
	s.DepartmentId = &v
	return s
}

func (s *UpdateIdpDepartmentRequest) SetDepartmentName(v string) *UpdateIdpDepartmentRequest {
	s.DepartmentName = &v
	return s
}

func (s *UpdateIdpDepartmentRequest) SetIdpConfigId(v string) *UpdateIdpDepartmentRequest {
	s.IdpConfigId = &v
	return s
}

type UpdateIdpDepartmentResponseBody struct {
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIdpDepartmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateIdpDepartmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIdpDepartmentResponseBody) SetRequestId(v string) *UpdateIdpDepartmentResponseBody {
	s.RequestId = &v
	return s
}

type UpdateIdpDepartmentResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIdpDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIdpDepartmentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIdpDepartmentResponse) GoString() string {
	return s.String()
}

func (s *UpdateIdpDepartmentResponse) SetHeaders(v map[string]*string) *UpdateIdpDepartmentResponse {
	s.Headers = v
	return s
}

func (s *UpdateIdpDepartmentResponse) SetStatusCode(v int32) *UpdateIdpDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIdpDepartmentResponse) SetBody(v *UpdateIdpDepartmentResponseBody) *UpdateIdpDepartmentResponse {
	s.Body = v
	return s
}

type UpdateNacUserCertStatusRequest struct {
	IdList []*UpdateNacUserCertStatusRequestIdList `json:"IdList,omitempty" xml:"IdList,omitempty" type:"Repeated"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateNacUserCertStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNacUserCertStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateNacUserCertStatusRequest) SetIdList(v []*UpdateNacUserCertStatusRequestIdList) *UpdateNacUserCertStatusRequest {
	s.IdList = v
	return s
}

func (s *UpdateNacUserCertStatusRequest) SetStatus(v string) *UpdateNacUserCertStatusRequest {
	s.Status = &v
	return s
}

type UpdateNacUserCertStatusRequestIdList struct {
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DevTag *string `json:"DevTag,omitempty" xml:"DevTag,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateNacUserCertStatusRequestIdList) String() string {
	return tea.Prettify(s)
}

func (s UpdateNacUserCertStatusRequestIdList) GoString() string {
	return s.String()
}

func (s *UpdateNacUserCertStatusRequestIdList) SetDevTag(v string) *UpdateNacUserCertStatusRequestIdList {
	s.DevTag = &v
	return s
}

func (s *UpdateNacUserCertStatusRequestIdList) SetUserId(v string) *UpdateNacUserCertStatusRequestIdList {
	s.UserId = &v
	return s
}

type UpdateNacUserCertStatusResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNacUserCertStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNacUserCertStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNacUserCertStatusResponseBody) SetCode(v string) *UpdateNacUserCertStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNacUserCertStatusResponseBody) SetMessage(v string) *UpdateNacUserCertStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNacUserCertStatusResponseBody) SetRequestId(v string) *UpdateNacUserCertStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateNacUserCertStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNacUserCertStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNacUserCertStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNacUserCertStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateNacUserCertStatusResponse) SetHeaders(v map[string]*string) *UpdateNacUserCertStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateNacUserCertStatusResponse) SetStatusCode(v int32) *UpdateNacUserCertStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNacUserCertStatusResponse) SetBody(v *UpdateNacUserCertStatusResponseBody) *UpdateNacUserCertStatusResponse {
	s.Body = v
	return s
}

type UpdatePrivateAccessApplicationRequest struct {
	Addresses []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// if can be null:
	// true
	Description                  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	L7ProxyDomainAutomaticPrefix *string `json:"L7ProxyDomainAutomaticPrefix,omitempty" xml:"L7ProxyDomainAutomaticPrefix,omitempty"`
	L7ProxyDomainCustom          *string `json:"L7ProxyDomainCustom,omitempty" xml:"L7ProxyDomainCustom,omitempty"`
	L7ProxyDomainPrivate         *string `json:"L7ProxyDomainPrivate,omitempty" xml:"L7ProxyDomainPrivate,omitempty"`
	// example:
	//
	// Cover
	ModifyType *string                                            `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	PortRanges []*UpdatePrivateAccessApplicationRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	// example:
	//
	// All
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// if can be null:
	// true
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s UpdatePrivateAccessApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationRequest) SetAddresses(v []*string) *UpdatePrivateAccessApplicationRequest {
	s.Addresses = v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetApplicationId(v string) *UpdatePrivateAccessApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetDescription(v string) *UpdatePrivateAccessApplicationRequest {
	s.Description = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetL7ProxyDomainAutomaticPrefix(v string) *UpdatePrivateAccessApplicationRequest {
	s.L7ProxyDomainAutomaticPrefix = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetL7ProxyDomainCustom(v string) *UpdatePrivateAccessApplicationRequest {
	s.L7ProxyDomainCustom = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetL7ProxyDomainPrivate(v string) *UpdatePrivateAccessApplicationRequest {
	s.L7ProxyDomainPrivate = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetModifyType(v string) *UpdatePrivateAccessApplicationRequest {
	s.ModifyType = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetPortRanges(v []*UpdatePrivateAccessApplicationRequestPortRanges) *UpdatePrivateAccessApplicationRequest {
	s.PortRanges = v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetProtocol(v string) *UpdatePrivateAccessApplicationRequest {
	s.Protocol = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetStatus(v string) *UpdatePrivateAccessApplicationRequest {
	s.Status = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequest) SetTagIds(v []*string) *UpdatePrivateAccessApplicationRequest {
	s.TagIds = v
	return s
}

type UpdatePrivateAccessApplicationRequestPortRanges struct {
	// example:
	//
	// 80
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// example:
	//
	// 81
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s UpdatePrivateAccessApplicationRequestPortRanges) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateAccessApplicationRequestPortRanges) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationRequestPortRanges) SetBegin(v int32) *UpdatePrivateAccessApplicationRequestPortRanges {
	s.Begin = &v
	return s
}

func (s *UpdatePrivateAccessApplicationRequestPortRanges) SetEnd(v int32) *UpdatePrivateAccessApplicationRequestPortRanges {
	s.End = &v
	return s
}

type UpdatePrivateAccessApplicationResponseBody struct {
	// example:
	//
	// FD724DBC-CD76-5235-BF76-59C51B73296D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrivateAccessApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateAccessApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationResponseBody) SetRequestId(v string) *UpdatePrivateAccessApplicationResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePrivateAccessApplicationResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrivateAccessApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateAccessApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationResponse) SetHeaders(v map[string]*string) *UpdatePrivateAccessApplicationResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivateAccessApplicationResponse) SetStatusCode(v int32) *UpdatePrivateAccessApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrivateAccessApplicationResponse) SetBody(v *UpdatePrivateAccessApplicationResponseBody) *UpdatePrivateAccessApplicationResponse {
	s.Body = v
	return s
}

type UpdatePrivateAccessPolicyRequest struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// example:
	//
	// Application
	ApplicationType      *string                                                 `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	CustomUserAttributes []*UpdatePrivateAccessPolicyRequestCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	// if can be null:
	// true
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceAttributeAction *string `json:"DeviceAttributeAction,omitempty" xml:"DeviceAttributeAction,omitempty"`
	DeviceAttributeId     *string `json:"DeviceAttributeId,omitempty" xml:"DeviceAttributeId,omitempty"`
	// example:
	//
	// Cover
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// example:
	//
	// Allow
	PolicyAction *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pa-policy-63b2f1844b86****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 内网访问标签ID集合。一条策略最多支持100个内网访问标签ID。
	TagIds       []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// 内网访问策略的用户组类型。取值：
	//
	// - **Normal**：普通用户组。
	//
	// - **Custom**：自定义用户组。
	//
	// example:
	//
	// Normal
	UserGroupMode *string `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
}

func (s UpdatePrivateAccessPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessPolicyRequest) SetApplicationIds(v []*string) *UpdatePrivateAccessPolicyRequest {
	s.ApplicationIds = v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetApplicationType(v string) *UpdatePrivateAccessPolicyRequest {
	s.ApplicationType = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetCustomUserAttributes(v []*UpdatePrivateAccessPolicyRequestCustomUserAttributes) *UpdatePrivateAccessPolicyRequest {
	s.CustomUserAttributes = v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetDescription(v string) *UpdatePrivateAccessPolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetDeviceAttributeAction(v string) *UpdatePrivateAccessPolicyRequest {
	s.DeviceAttributeAction = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetDeviceAttributeId(v string) *UpdatePrivateAccessPolicyRequest {
	s.DeviceAttributeId = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetModifyType(v string) *UpdatePrivateAccessPolicyRequest {
	s.ModifyType = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetPolicyAction(v string) *UpdatePrivateAccessPolicyRequest {
	s.PolicyAction = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetPolicyId(v string) *UpdatePrivateAccessPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetPriority(v int32) *UpdatePrivateAccessPolicyRequest {
	s.Priority = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetStatus(v string) *UpdatePrivateAccessPolicyRequest {
	s.Status = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetTagIds(v []*string) *UpdatePrivateAccessPolicyRequest {
	s.TagIds = v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetUserGroupIds(v []*string) *UpdatePrivateAccessPolicyRequest {
	s.UserGroupIds = v
	return s
}

func (s *UpdatePrivateAccessPolicyRequest) SetUserGroupMode(v string) *UpdatePrivateAccessPolicyRequest {
	s.UserGroupMode = &v
	return s
}

type UpdatePrivateAccessPolicyRequestCustomUserAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdatePrivateAccessPolicyRequestCustomUserAttributes) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateAccessPolicyRequestCustomUserAttributes) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessPolicyRequestCustomUserAttributes) SetIdpId(v int32) *UpdatePrivateAccessPolicyRequestCustomUserAttributes {
	s.IdpId = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequestCustomUserAttributes) SetRelation(v string) *UpdatePrivateAccessPolicyRequestCustomUserAttributes {
	s.Relation = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequestCustomUserAttributes) SetUserGroupType(v string) *UpdatePrivateAccessPolicyRequestCustomUserAttributes {
	s.UserGroupType = &v
	return s
}

func (s *UpdatePrivateAccessPolicyRequestCustomUserAttributes) SetValue(v string) *UpdatePrivateAccessPolicyRequestCustomUserAttributes {
	s.Value = &v
	return s
}

type UpdatePrivateAccessPolicyResponseBody struct {
	// example:
	//
	// 5FEF5CFA-14CC-5DE5-BD1F-AFFE0996E71D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrivateAccessPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateAccessPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessPolicyResponseBody) SetRequestId(v string) *UpdatePrivateAccessPolicyResponseBody {
	s.RequestId = &v
	return s
}

type UpdatePrivateAccessPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrivateAccessPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *UpdatePrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivateAccessPolicyResponse) SetStatusCode(v int32) *UpdatePrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrivateAccessPolicyResponse) SetBody(v *UpdatePrivateAccessPolicyResponseBody) *UpdatePrivateAccessPolicyResponse {
	s.Body = v
	return s
}

type UpdateRegistrationPolicyRequest struct {
	CompanyLimitCount *UpdateRegistrationPolicyRequestCompanyLimitCount `json:"CompanyLimitCount,omitempty" xml:"CompanyLimitCount,omitempty" type:"Struct"`
	// example:
	//
	// LimitAll
	CompanyLimitType *string `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// example:
	//
	// registration_policy_name
	Name               *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PersonalLimitCount *UpdateRegistrationPolicyRequestPersonalLimitCount `json:"PersonalLimitCount,omitempty" xml:"PersonalLimitCount,omitempty" type:"Struct"`
	// example:
	//
	// LimitDiff
	PersonalLimitType *string `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// reg-policy-63b2f1844b86****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 0
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s UpdateRegistrationPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegistrationPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyRequest) SetCompanyLimitCount(v *UpdateRegistrationPolicyRequestCompanyLimitCount) *UpdateRegistrationPolicyRequest {
	s.CompanyLimitCount = v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetCompanyLimitType(v string) *UpdateRegistrationPolicyRequest {
	s.CompanyLimitType = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetDescription(v string) *UpdateRegistrationPolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetMatchMode(v string) *UpdateRegistrationPolicyRequest {
	s.MatchMode = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetName(v string) *UpdateRegistrationPolicyRequest {
	s.Name = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetPersonalLimitCount(v *UpdateRegistrationPolicyRequestPersonalLimitCount) *UpdateRegistrationPolicyRequest {
	s.PersonalLimitCount = v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetPersonalLimitType(v string) *UpdateRegistrationPolicyRequest {
	s.PersonalLimitType = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetPolicyId(v string) *UpdateRegistrationPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetPriority(v int64) *UpdateRegistrationPolicyRequest {
	s.Priority = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetStatus(v string) *UpdateRegistrationPolicyRequest {
	s.Status = &v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetUserGroupIds(v []*string) *UpdateRegistrationPolicyRequest {
	s.UserGroupIds = v
	return s
}

func (s *UpdateRegistrationPolicyRequest) SetWhitelist(v []*string) *UpdateRegistrationPolicyRequest {
	s.Whitelist = v
	return s
}

type UpdateRegistrationPolicyRequestCompanyLimitCount struct {
	// example:
	//
	// 1
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 0
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 0
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s UpdateRegistrationPolicyRequestCompanyLimitCount) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegistrationPolicyRequestCompanyLimitCount) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyRequestCompanyLimitCount) SetAll(v int32) *UpdateRegistrationPolicyRequestCompanyLimitCount {
	s.All = &v
	return s
}

func (s *UpdateRegistrationPolicyRequestCompanyLimitCount) SetMobile(v int32) *UpdateRegistrationPolicyRequestCompanyLimitCount {
	s.Mobile = &v
	return s
}

func (s *UpdateRegistrationPolicyRequestCompanyLimitCount) SetPC(v int32) *UpdateRegistrationPolicyRequestCompanyLimitCount {
	s.PC = &v
	return s
}

type UpdateRegistrationPolicyRequestPersonalLimitCount struct {
	// example:
	//
	// 0
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 1
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 2
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s UpdateRegistrationPolicyRequestPersonalLimitCount) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegistrationPolicyRequestPersonalLimitCount) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyRequestPersonalLimitCount) SetAll(v int32) *UpdateRegistrationPolicyRequestPersonalLimitCount {
	s.All = &v
	return s
}

func (s *UpdateRegistrationPolicyRequestPersonalLimitCount) SetMobile(v int32) *UpdateRegistrationPolicyRequestPersonalLimitCount {
	s.Mobile = &v
	return s
}

func (s *UpdateRegistrationPolicyRequestPersonalLimitCount) SetPC(v int32) *UpdateRegistrationPolicyRequestPersonalLimitCount {
	s.PC = &v
	return s
}

type UpdateRegistrationPolicyShrinkRequest struct {
	CompanyLimitCountShrink *string `json:"CompanyLimitCount,omitempty" xml:"CompanyLimitCount,omitempty"`
	// example:
	//
	// LimitAll
	CompanyLimitType *string `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// example:
	//
	// registration_policy_name
	Name                     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PersonalLimitCountShrink *string `json:"PersonalLimitCount,omitempty" xml:"PersonalLimitCount,omitempty"`
	// example:
	//
	// LimitDiff
	PersonalLimitType *string `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// reg-policy-63b2f1844b86****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 0
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s UpdateRegistrationPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegistrationPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetCompanyLimitCountShrink(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.CompanyLimitCountShrink = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetCompanyLimitType(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.CompanyLimitType = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetDescription(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetMatchMode(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.MatchMode = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetName(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetPersonalLimitCountShrink(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.PersonalLimitCountShrink = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetPersonalLimitType(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.PersonalLimitType = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetPolicyId(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetPriority(v int64) *UpdateRegistrationPolicyShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetStatus(v string) *UpdateRegistrationPolicyShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetUserGroupIds(v []*string) *UpdateRegistrationPolicyShrinkRequest {
	s.UserGroupIds = v
	return s
}

func (s *UpdateRegistrationPolicyShrinkRequest) SetWhitelist(v []*string) *UpdateRegistrationPolicyShrinkRequest {
	s.Whitelist = v
	return s
}

type UpdateRegistrationPolicyResponseBody struct {
	Policy *UpdateRegistrationPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	// example:
	//
	// 27064ECA-0936-59F3-8A98-EC821E5BD08F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRegistrationPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegistrationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyResponseBody) SetPolicy(v *UpdateRegistrationPolicyResponseBodyPolicy) *UpdateRegistrationPolicyResponseBody {
	s.Policy = v
	return s
}

func (s *UpdateRegistrationPolicyResponseBody) SetRequestId(v string) *UpdateRegistrationPolicyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRegistrationPolicyResponseBodyPolicy struct {
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime  *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail []*UpdateRegistrationPolicyResponseBodyPolicyLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// example:
	//
	// registration_policy_name
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
}

func (s UpdateRegistrationPolicyResponseBodyPolicy) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegistrationPolicyResponseBodyPolicy) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetCreateTime(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.CreateTime = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetDescription(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.Description = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetLimitDetail(v []*UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.LimitDetail = v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetMatchMode(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.MatchMode = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetName(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.Name = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetPolicyId(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.PolicyId = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetPriority(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.Priority = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetStatus(v string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.Status = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetUserGroupIds(v []*string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.UserGroupIds = v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicy) SetWhitelist(v []*string) *UpdateRegistrationPolicyResponseBodyPolicy {
	s.Whitelist = v
	return s
}

type UpdateRegistrationPolicyResponseBodyPolicyLimitDetail struct {
	// example:
	//
	// Company
	DeviceBelong *string                                                          `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	// example:
	//
	// LimitAll
	LimitType *string `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
}

func (s UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) SetDeviceBelong(v string) *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail {
	s.DeviceBelong = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) SetLimitCount(v *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail {
	s.LimitCount = v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail) SetLimitType(v string) *UpdateRegistrationPolicyResponseBodyPolicyLimitDetail {
	s.LimitType = &v
	return s
}

type UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount struct {
	// example:
	//
	// 1
	All *int32 `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 0
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// example:
	//
	// 0
	PC *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
}

func (s UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) SetAll(v int32) *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	s.All = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) SetMobile(v int32) *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	s.Mobile = &v
	return s
}

func (s *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount) SetPC(v int32) *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount {
	s.PC = &v
	return s
}

type UpdateRegistrationPolicyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRegistrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRegistrationPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegistrationPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateRegistrationPolicyResponse) SetHeaders(v map[string]*string) *UpdateRegistrationPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateRegistrationPolicyResponse) SetStatusCode(v int32) *UpdateRegistrationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRegistrationPolicyResponse) SetBody(v *UpdateRegistrationPolicyResponseBody) *UpdateRegistrationPolicyResponse {
	s.Body = v
	return s
}

type UpdateUserDevicesSharingStatusRequest struct {
	// This parameter is required.
	DeviceTags []*string `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// true
	SharingStatus *bool `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
}

func (s UpdateUserDevicesSharingStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDevicesSharingStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesSharingStatusRequest) SetDeviceTags(v []*string) *UpdateUserDevicesSharingStatusRequest {
	s.DeviceTags = v
	return s
}

func (s *UpdateUserDevicesSharingStatusRequest) SetSharingStatus(v bool) *UpdateUserDevicesSharingStatusRequest {
	s.SharingStatus = &v
	return s
}

type UpdateUserDevicesSharingStatusResponseBody struct {
	Devices []*UpdateUserDevicesSharingStatusResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// example:
	//
	// 4D169859-A4F2-5EC8-853B-8447787C0D8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserDevicesSharingStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDevicesSharingStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesSharingStatusResponseBody) SetDevices(v []*UpdateUserDevicesSharingStatusResponseBodyDevices) *UpdateUserDevicesSharingStatusResponseBody {
	s.Devices = v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBody) SetRequestId(v string) *UpdateUserDevicesSharingStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserDevicesSharingStatusResponseBodyDevices struct {
	// example:
	//
	// Online
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// example:
	//
	// 2.2.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// Apple M1
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 2023-05-16 17:18:46
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// Company
	DeviceBelong *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	// example:
	//
	// MacBookPro17,1
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	// example:
	//
	// Online
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// example:
	//
	// Windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 3.5.1
	DeviceVersion *string `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	// example:
	//
	// APPLE SSD AP0512Q Media
	Disk *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// example:
	//
	// Enabled
	DlpStatus *string `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// Enabled
	IaStatus *string `json:"IaStatus,omitempty" xml:"IaStatus,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	InnerIP *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 16
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// Enabled
	NacStatus *string `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	// example:
	//
	// Enabled
	PaStatus *string `json:"PaStatus,omitempty" xml:"PaStatus,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// example:
	//
	// true
	SharingStatus *bool `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	// example:
	//
	// 11.49.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// example:
	//
	// 2023-08-24 19:04:42
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateUserDevicesSharingStatusResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDevicesSharingStatusResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetAppStatus(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.AppStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetAppVersion(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.AppVersion = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetCPU(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.CPU = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetCreateTime(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.CreateTime = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDepartment(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.Department = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDeviceBelong(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DeviceBelong = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDeviceModel(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DeviceModel = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDeviceStatus(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DeviceStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDeviceTag(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DeviceTag = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDeviceType(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DeviceType = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDeviceVersion(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DeviceVersion = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDisk(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.Disk = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetDlpStatus(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.DlpStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetHostname(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.Hostname = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetIaStatus(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.IaStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetInnerIP(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.InnerIP = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetMac(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.Mac = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetMemory(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.Memory = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetNacStatus(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.NacStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetPaStatus(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.PaStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetSaseUserId(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.SaseUserId = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetSharingStatus(v bool) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.SharingStatus = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetSrcIP(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.SrcIP = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetUpdateTime(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.UpdateTime = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponseBodyDevices) SetUsername(v string) *UpdateUserDevicesSharingStatusResponseBodyDevices {
	s.Username = &v
	return s
}

type UpdateUserDevicesSharingStatusResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserDevicesSharingStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserDevicesSharingStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDevicesSharingStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesSharingStatusResponse) SetHeaders(v map[string]*string) *UpdateUserDevicesSharingStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponse) SetStatusCode(v int32) *UpdateUserDevicesSharingStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserDevicesSharingStatusResponse) SetBody(v *UpdateUserDevicesSharingStatusResponseBody) *UpdateUserDevicesSharingStatusResponse {
	s.Body = v
	return s
}

type UpdateUserDevicesStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// Unbound
	DeviceAction *string `json:"DeviceAction,omitempty" xml:"DeviceAction,omitempty"`
	// This parameter is required.
	DeviceTags []*string `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty" type:"Repeated"`
}

func (s UpdateUserDevicesStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDevicesStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesStatusRequest) SetDeviceAction(v string) *UpdateUserDevicesStatusRequest {
	s.DeviceAction = &v
	return s
}

func (s *UpdateUserDevicesStatusRequest) SetDeviceTags(v []*string) *UpdateUserDevicesStatusRequest {
	s.DeviceTags = v
	return s
}

type UpdateUserDevicesStatusResponseBody struct {
	Devices []*UpdateUserDevicesStatusResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	// example:
	//
	// 5FEF5CFA-14CC-5DE5-BD1F-AFFE0996E71D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserDevicesStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDevicesStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesStatusResponseBody) SetDevices(v []*UpdateUserDevicesStatusResponseBodyDevices) *UpdateUserDevicesStatusResponseBody {
	s.Devices = v
	return s
}

func (s *UpdateUserDevicesStatusResponseBody) SetRequestId(v string) *UpdateUserDevicesStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserDevicesStatusResponseBodyDevices struct {
	// example:
	//
	// Online
	AppStatus *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	// example:
	//
	// 2.2.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// Apple M1
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 2023-07-17 18:46:55
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department *string `json:"Department,omitempty" xml:"Department,omitempty"`
	// example:
	//
	// Company
	DeviceBelong *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	// example:
	//
	// MacBookPro17,1
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	// example:
	//
	// Online
	DeviceStatus *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	// example:
	//
	// 36efa42d-2c32-c4dc-e3fc-8541e33a****
	DeviceTag *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	// example:
	//
	// Windows
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 3.5.1
	DeviceVersion *string `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	// example:
	//
	// APPLE SSD AP0512Q Media
	Disk *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	// example:
	//
	// Enabled
	DlpStatus *string `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	// example:
	//
	// win10-64bit
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// Enabled
	IaStatus *string `json:"IaStatus,omitempty" xml:"IaStatus,omitempty"`
	// example:
	//
	// 192.168.XX.XX
	InnerIP *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	// example:
	//
	// 00:16:XX:XX:7c:46
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 16
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// example:
	//
	// Enabled
	NacStatus *string `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	// example:
	//
	// Enabled
	PaStatus *string `json:"PaStatus,omitempty" xml:"PaStatus,omitempty"`
	// example:
	//
	// su_e8f218fb171edd167c2ad917d21f53148bdefc510ca1f3c3cc0249d3643d****
	SaseUserId *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	// example:
	//
	// true
	SharingStatus *bool `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	// example:
	//
	// 11.49.XX.XX
	SrcIP *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	// example:
	//
	// 2023-08-24 19:04:42
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Username   *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s UpdateUserDevicesStatusResponseBodyDevices) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDevicesStatusResponseBodyDevices) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetAppStatus(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.AppStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetAppVersion(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.AppVersion = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetCPU(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.CPU = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetCreateTime(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.CreateTime = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDepartment(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.Department = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDeviceBelong(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DeviceBelong = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDeviceModel(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DeviceModel = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDeviceStatus(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DeviceStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDeviceTag(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DeviceTag = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDeviceType(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DeviceType = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDeviceVersion(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DeviceVersion = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDisk(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.Disk = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetDlpStatus(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.DlpStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetHostname(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.Hostname = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetIaStatus(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.IaStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetInnerIP(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.InnerIP = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetMac(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.Mac = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetMemory(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.Memory = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetNacStatus(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.NacStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetPaStatus(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.PaStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetSaseUserId(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.SaseUserId = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetSharingStatus(v bool) *UpdateUserDevicesStatusResponseBodyDevices {
	s.SharingStatus = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetSrcIP(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.SrcIP = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetUpdateTime(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.UpdateTime = &v
	return s
}

func (s *UpdateUserDevicesStatusResponseBodyDevices) SetUsername(v string) *UpdateUserDevicesStatusResponseBodyDevices {
	s.Username = &v
	return s
}

type UpdateUserDevicesStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserDevicesStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserDevicesStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserDevicesStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserDevicesStatusResponse) SetHeaders(v map[string]*string) *UpdateUserDevicesStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserDevicesStatusResponse) SetStatusCode(v int32) *UpdateUserDevicesStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserDevicesStatusResponse) SetBody(v *UpdateUserDevicesStatusResponseBody) *UpdateUserDevicesStatusResponse {
	s.Body = v
	return s
}

type UpdateUserGroupRequest struct {
	Attributes []*UpdateUserGroupRequestAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// if can be null:
	// true
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Cover
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// usergroup-6f1ef2fc56b6****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s UpdateUserGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupRequest) SetAttributes(v []*UpdateUserGroupRequestAttributes) *UpdateUserGroupRequest {
	s.Attributes = v
	return s
}

func (s *UpdateUserGroupRequest) SetDescription(v string) *UpdateUserGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateUserGroupRequest) SetModifyType(v string) *UpdateUserGroupRequest {
	s.ModifyType = &v
	return s
}

func (s *UpdateUserGroupRequest) SetUserGroupId(v string) *UpdateUserGroupRequest {
	s.UserGroupId = &v
	return s
}

type UpdateUserGroupRequestAttributes struct {
	// example:
	//
	// 12
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Equal
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// department
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateUserGroupRequestAttributes) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupRequestAttributes) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupRequestAttributes) SetIdpId(v int32) *UpdateUserGroupRequestAttributes {
	s.IdpId = &v
	return s
}

func (s *UpdateUserGroupRequestAttributes) SetRelation(v string) *UpdateUserGroupRequestAttributes {
	s.Relation = &v
	return s
}

func (s *UpdateUserGroupRequestAttributes) SetUserGroupType(v string) *UpdateUserGroupRequestAttributes {
	s.UserGroupType = &v
	return s
}

func (s *UpdateUserGroupRequestAttributes) SetValue(v string) *UpdateUserGroupRequestAttributes {
	s.Value = &v
	return s
}

type UpdateUserGroupResponseBody struct {
	// example:
	//
	// FD724DBC-CD76-5235-BF76-59C51B73296D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponseBody) SetRequestId(v string) *UpdateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUserGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponse) SetHeaders(v map[string]*string) *UpdateUserGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserGroupResponse) SetStatusCode(v int32) *UpdateUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserGroupResponse) SetBody(v *UpdateUserGroupResponseBody) *UpdateUserGroupResponse {
	s.Body = v
	return s
}

type UpdateUsersStatusRequest struct {
	// This parameter is required.
	SaseUserIds []*string `json:"SaseUserIds,omitempty" xml:"SaseUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateUsersStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUsersStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateUsersStatusRequest) SetSaseUserIds(v []*string) *UpdateUsersStatusRequest {
	s.SaseUserIds = v
	return s
}

func (s *UpdateUsersStatusRequest) SetStatus(v string) *UpdateUsersStatusRequest {
	s.Status = &v
	return s
}

type UpdateUsersStatusResponseBody struct {
	// example:
	//
	// 47363C2B-1AAA-5954-8847-0E50FCC54117
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUsersStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateUsersStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUsersStatusResponseBody) SetRequestId(v string) *UpdateUsersStatusResponseBody {
	s.RequestId = &v
	return s
}

type UpdateUsersStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUsersStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUsersStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUsersStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateUsersStatusResponse) SetHeaders(v map[string]*string) *UpdateUsersStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateUsersStatusResponse) SetStatusCode(v int32) *UpdateUsersStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUsersStatusResponse) SetBody(v *UpdateUsersStatusResponseBody) *UpdateUsersStatusResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("csas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 挂载connector的应用
//
// @param tmpReq - AttachApplication2ConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AttachApplication2ConnectorResponse
func (client *Client) AttachApplication2ConnectorWithOptions(tmpReq *AttachApplication2ConnectorRequest, runtime *util.RuntimeOptions) (_result *AttachApplication2ConnectorResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AttachApplication2ConnectorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ApplicationIds)) {
		request.ApplicationIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationIds, tea.String("ApplicationIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationIdsShrink)) {
		body["ApplicationIds"] = request.ApplicationIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectorId)) {
		body["ConnectorId"] = request.ConnectorId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AttachApplication2Connector"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AttachApplication2ConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 挂载connector的应用
//
// @param request - AttachApplication2ConnectorRequest
//
// @return AttachApplication2ConnectorResponse
func (client *Client) AttachApplication2Connector(request *AttachApplication2ConnectorRequest) (_result *AttachApplication2ConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AttachApplication2ConnectorResponse{}
	_body, _err := client.AttachApplication2ConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建自定义身份源用户
//
// @param request - CreateClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClientUserResponse
func (client *Client) CreateClientUserWithOptions(request *CreateClientUserRequest, runtime *util.RuntimeOptions) (_result *CreateClientUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.IdpConfigId)) {
		query["IdpConfigId"] = request.IdpConfigId
	}

	if !tea.BoolValue(util.IsUnset(request.MobileNumber)) {
		query["MobileNumber"] = request.MobileNumber
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateClientUser"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClientUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建自定义身份源用户
//
// @param request - CreateClientUserRequest
//
// @return CreateClientUserResponse
func (client *Client) CreateClientUser(request *CreateClientUserRequest) (_result *CreateClientUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClientUserResponse{}
	_body, _err := client.CreateClientUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建动态路由
//
// @param request - CreateDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDynamicRouteResponse
func (client *Client) CreateDynamicRouteWithOptions(request *CreateDynamicRouteRequest, runtime *util.RuntimeOptions) (_result *CreateDynamicRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationIds)) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationType)) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DynamicRouteType)) {
		body["DynamicRouteType"] = request.DynamicRouteType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextHop)) {
		body["NextHop"] = request.NextHop
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RegionIds)) {
		bodyFlat["RegionIds"] = request.RegionIds
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagIds)) {
		bodyFlat["TagIds"] = request.TagIds
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDynamicRoute"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDynamicRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建动态路由
//
// @param request - CreateDynamicRouteRequest
//
// @return CreateDynamicRouteResponse
func (client *Client) CreateDynamicRoute(request *CreateDynamicRouteRequest) (_result *CreateDynamicRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDynamicRouteResponse{}
	_body, _err := client.CreateDynamicRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建自定义身份源部门
//
// @param request - CreateIdpDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdpDepartmentResponse
func (client *Client) CreateIdpDepartmentWithOptions(request *CreateIdpDepartmentRequest, runtime *util.RuntimeOptions) (_result *CreateIdpDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentName)) {
		query["DepartmentName"] = request.DepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.IdpConfigId)) {
		query["IdpConfigId"] = request.IdpConfigId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIdpDepartment"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIdpDepartmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建自定义身份源部门
//
// @param request - CreateIdpDepartmentRequest
//
// @return CreateIdpDepartmentResponse
func (client *Client) CreateIdpDepartment(request *CreateIdpDepartmentRequest) (_result *CreateIdpDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateIdpDepartmentResponse{}
	_body, _err := client.CreateIdpDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建内网访问应用
//
// @param request - CreatePrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivateAccessApplicationResponse
func (client *Client) CreatePrivateAccessApplicationWithOptions(request *CreatePrivateAccessApplicationRequest, runtime *util.RuntimeOptions) (_result *CreatePrivateAccessApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Addresses)) {
		bodyFlat["Addresses"] = request.Addresses
	}

	if !tea.BoolValue(util.IsUnset(request.BrowserAccessStatus)) {
		body["BrowserAccessStatus"] = request.BrowserAccessStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.L7ProxyDomainAutomaticPrefix)) {
		body["L7ProxyDomainAutomaticPrefix"] = request.L7ProxyDomainAutomaticPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.L7ProxyDomainCustom)) {
		body["L7ProxyDomainCustom"] = request.L7ProxyDomainCustom
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PortRanges)) {
		bodyFlat["PortRanges"] = request.PortRanges
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagIds)) {
		bodyFlat["TagIds"] = request.TagIds
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePrivateAccessApplication"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePrivateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建内网访问应用
//
// @param request - CreatePrivateAccessApplicationRequest
//
// @return CreatePrivateAccessApplicationResponse
func (client *Client) CreatePrivateAccessApplication(request *CreatePrivateAccessApplicationRequest) (_result *CreatePrivateAccessApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePrivateAccessApplicationResponse{}
	_body, _err := client.CreatePrivateAccessApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建内网访问策略
//
// @param request - CreatePrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivateAccessPolicyResponse
func (client *Client) CreatePrivateAccessPolicyWithOptions(request *CreatePrivateAccessPolicyRequest, runtime *util.RuntimeOptions) (_result *CreatePrivateAccessPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationIds)) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationType)) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !tea.BoolValue(util.IsUnset(request.CustomUserAttributes)) {
		bodyFlat["CustomUserAttributes"] = request.CustomUserAttributes
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceAttributeAction)) {
		body["DeviceAttributeAction"] = request.DeviceAttributeAction
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceAttributeId)) {
		body["DeviceAttributeId"] = request.DeviceAttributeId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyAction)) {
		body["PolicyAction"] = request.PolicyAction
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagIds)) {
		bodyFlat["TagIds"] = request.TagIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupIds)) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupMode)) {
		body["UserGroupMode"] = request.UserGroupMode
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePrivateAccessPolicy"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建内网访问策略
//
// @param request - CreatePrivateAccessPolicyRequest
//
// @return CreatePrivateAccessPolicyResponse
func (client *Client) CreatePrivateAccessPolicy(request *CreatePrivateAccessPolicyRequest) (_result *CreatePrivateAccessPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePrivateAccessPolicyResponse{}
	_body, _err := client.CreatePrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建内网访问标签
//
// @param request - CreatePrivateAccessTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePrivateAccessTagResponse
func (client *Client) CreatePrivateAccessTagWithOptions(request *CreatePrivateAccessTagRequest, runtime *util.RuntimeOptions) (_result *CreatePrivateAccessTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePrivateAccessTag"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePrivateAccessTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建内网访问标签
//
// @param request - CreatePrivateAccessTagRequest
//
// @return CreatePrivateAccessTagResponse
func (client *Client) CreatePrivateAccessTag(request *CreatePrivateAccessTagRequest) (_result *CreatePrivateAccessTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePrivateAccessTagResponse{}
	_body, _err := client.CreatePrivateAccessTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建设备注册策略
//
// @param tmpReq - CreateRegistrationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRegistrationPolicyResponse
func (client *Client) CreateRegistrationPolicyWithOptions(tmpReq *CreateRegistrationPolicyRequest, runtime *util.RuntimeOptions) (_result *CreateRegistrationPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateRegistrationPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CompanyLimitCount)) {
		request.CompanyLimitCountShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CompanyLimitCount, tea.String("CompanyLimitCount"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PersonalLimitCount)) {
		request.PersonalLimitCountShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PersonalLimitCount, tea.String("PersonalLimitCount"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompanyLimitCountShrink)) {
		body["CompanyLimitCount"] = request.CompanyLimitCountShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CompanyLimitType)) {
		body["CompanyLimitType"] = request.CompanyLimitType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.MatchMode)) {
		body["MatchMode"] = request.MatchMode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PersonalLimitCountShrink)) {
		body["PersonalLimitCount"] = request.PersonalLimitCountShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PersonalLimitType)) {
		body["PersonalLimitType"] = request.PersonalLimitType
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserGroupIds)) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.Whitelist)) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRegistrationPolicy"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRegistrationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建设备注册策略
//
// @param request - CreateRegistrationPolicyRequest
//
// @return CreateRegistrationPolicyResponse
func (client *Client) CreateRegistrationPolicy(request *CreateRegistrationPolicyRequest) (_result *CreateRegistrationPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRegistrationPolicyResponse{}
	_body, _err := client.CreateRegistrationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建用户组
//
// @param request - CreateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroupWithOptions(request *CreateUserGroupRequest, runtime *util.RuntimeOptions) (_result *CreateUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attributes)) {
		bodyFlat["Attributes"] = request.Attributes
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserGroup"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建用户组
//
// @param request - CreateUserGroupRequest
//
// @return CreateUserGroupResponse
func (client *Client) CreateUserGroup(request *CreateUserGroupRequest) (_result *CreateUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserGroupResponse{}
	_body, _err := client.CreateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建数字水印暗水印透明底图
//
// @param request - CreateWmBaseImageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWmBaseImageResponse
func (client *Client) CreateWmBaseImageWithOptions(request *CreateWmBaseImageRequest, runtime *util.RuntimeOptions) (_result *CreateWmBaseImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Height)) {
		body["Height"] = request.Height
	}

	if !tea.BoolValue(util.IsUnset(request.Opacity)) {
		body["Opacity"] = request.Opacity
	}

	if !tea.BoolValue(util.IsUnset(request.Scale)) {
		body["Scale"] = request.Scale
	}

	if !tea.BoolValue(util.IsUnset(request.Width)) {
		body["Width"] = request.Width
	}

	if !tea.BoolValue(util.IsUnset(request.WmInfoBytesB64)) {
		body["WmInfoBytesB64"] = request.WmInfoBytesB64
	}

	if !tea.BoolValue(util.IsUnset(request.WmInfoSize)) {
		body["WmInfoSize"] = request.WmInfoSize
	}

	if !tea.BoolValue(util.IsUnset(request.WmInfoUint)) {
		body["WmInfoUint"] = request.WmInfoUint
	}

	if !tea.BoolValue(util.IsUnset(request.WmType)) {
		body["WmType"] = request.WmType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWmBaseImage"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWmBaseImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数字水印暗水印透明底图
//
// @param request - CreateWmBaseImageRequest
//
// @return CreateWmBaseImageResponse
func (client *Client) CreateWmBaseImage(request *CreateWmBaseImageRequest) (_result *CreateWmBaseImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWmBaseImageResponse{}
	_body, _err := client.CreateWmBaseImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建嵌入水印任务
//
// @param tmpReq - CreateWmEmbedTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWmEmbedTaskResponse
func (client *Client) CreateWmEmbedTaskWithOptions(tmpReq *CreateWmEmbedTaskRequest, runtime *util.RuntimeOptions) (_result *CreateWmEmbedTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateWmEmbedTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CsvControl)) {
		request.CsvControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CsvControl, tea.String("CsvControl"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DocumentControl)) {
		request.DocumentControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DocumentControl, tea.String("DocumentControl"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CsvControlShrink)) {
		query["CsvControl"] = request.CsvControlShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocumentControlShrink)) {
		body["DocumentControl"] = request.DocumentControlShrink
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["FileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Filename)) {
		body["Filename"] = request.Filename
	}

	if !tea.BoolValue(util.IsUnset(request.ImageEmbedJpegQuality)) {
		body["ImageEmbedJpegQuality"] = request.ImageEmbedJpegQuality
	}

	if !tea.BoolValue(util.IsUnset(request.ImageEmbedLevel)) {
		body["ImageEmbedLevel"] = request.ImageEmbedLevel
	}

	if !tea.BoolValue(util.IsUnset(request.VideoBitrate)) {
		body["VideoBitrate"] = request.VideoBitrate
	}

	if !tea.BoolValue(util.IsUnset(request.VideoIsLong)) {
		body["VideoIsLong"] = request.VideoIsLong
	}

	if !tea.BoolValue(util.IsUnset(request.WmInfoBytesB64)) {
		body["WmInfoBytesB64"] = request.WmInfoBytesB64
	}

	if !tea.BoolValue(util.IsUnset(request.WmInfoSize)) {
		body["WmInfoSize"] = request.WmInfoSize
	}

	if !tea.BoolValue(util.IsUnset(request.WmInfoUint)) {
		body["WmInfoUint"] = request.WmInfoUint
	}

	if !tea.BoolValue(util.IsUnset(request.WmType)) {
		body["WmType"] = request.WmType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWmEmbedTask"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWmEmbedTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建嵌入水印任务
//
// @param request - CreateWmEmbedTaskRequest
//
// @return CreateWmEmbedTaskResponse
func (client *Client) CreateWmEmbedTask(request *CreateWmEmbedTaskRequest) (_result *CreateWmEmbedTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWmEmbedTaskResponse{}
	_body, _err := client.CreateWmEmbedTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建文件水印提取任务
//
// @param tmpReq - CreateWmExtractTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWmExtractTaskResponse
func (client *Client) CreateWmExtractTaskWithOptions(tmpReq *CreateWmExtractTaskRequest, runtime *util.RuntimeOptions) (_result *CreateWmExtractTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateWmExtractTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CsvControl)) {
		request.CsvControlShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CsvControl, tea.String("CsvControl"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CsvControlShrink)) {
		query["CsvControl"] = request.CsvControlShrink
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DocumentIsCapture)) {
		body["DocumentIsCapture"] = request.DocumentIsCapture
	}

	if !tea.BoolValue(util.IsUnset(request.FileUrl)) {
		body["FileUrl"] = request.FileUrl
	}

	if !tea.BoolValue(util.IsUnset(request.Filename)) {
		body["Filename"] = request.Filename
	}

	if !tea.BoolValue(util.IsUnset(request.VideoIsLong)) {
		body["VideoIsLong"] = request.VideoIsLong
	}

	if !tea.BoolValue(util.IsUnset(request.VideoSpeed)) {
		body["VideoSpeed"] = request.VideoSpeed
	}

	if !tea.BoolValue(util.IsUnset(request.WmInfoSize)) {
		body["WmInfoSize"] = request.WmInfoSize
	}

	if !tea.BoolValue(util.IsUnset(request.WmType)) {
		body["WmType"] = request.WmType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWmExtractTask"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWmExtractTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文件水印提取任务
//
// @param request - CreateWmExtractTaskRequest
//
// @return CreateWmExtractTaskResponse
func (client *Client) CreateWmExtractTask(request *CreateWmExtractTaskRequest) (_result *CreateWmExtractTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWmExtractTaskResponse{}
	_body, _err := client.CreateWmExtractTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建一条字符串水印信息到数字水印信息的映射记录
//
// @param request - CreateWmInfoMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWmInfoMappingResponse
func (client *Client) CreateWmInfoMappingWithOptions(request *CreateWmInfoMappingRequest, runtime *util.RuntimeOptions) (_result *CreateWmInfoMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WmInfoBytesB64)) {
		body["WmInfoBytesB64"] = request.WmInfoBytesB64
	}

	if !tea.BoolValue(util.IsUnset(request.WmInfoSize)) {
		body["WmInfoSize"] = request.WmInfoSize
	}

	if !tea.BoolValue(util.IsUnset(request.WmType)) {
		body["WmType"] = request.WmType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateWmInfoMapping"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateWmInfoMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建一条字符串水印信息到数字水印信息的映射记录
//
// @param request - CreateWmInfoMappingRequest
//
// @return CreateWmInfoMappingResponse
func (client *Client) CreateWmInfoMapping(request *CreateWmInfoMappingRequest) (_result *CreateWmInfoMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateWmInfoMappingResponse{}
	_body, _err := client.CreateWmInfoMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除自定义身份源指定用户
//
// @param request - DeleteClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientUserResponse
func (client *Client) DeleteClientUserWithOptions(request *DeleteClientUserRequest, runtime *util.RuntimeOptions) (_result *DeleteClientUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteClientUser"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClientUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除自定义身份源指定用户
//
// @param request - DeleteClientUserRequest
//
// @return DeleteClientUserResponse
func (client *Client) DeleteClientUser(request *DeleteClientUserRequest) (_result *DeleteClientUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClientUserResponse{}
	_body, _err := client.DeleteClientUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除动态路由
//
// @param request - DeleteDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDynamicRouteResponse
func (client *Client) DeleteDynamicRouteWithOptions(request *DeleteDynamicRouteRequest, runtime *util.RuntimeOptions) (_result *DeleteDynamicRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DynamicRouteId)) {
		query["DynamicRouteId"] = request.DynamicRouteId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDynamicRoute"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDynamicRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除动态路由
//
// @param request - DeleteDynamicRouteRequest
//
// @return DeleteDynamicRouteResponse
func (client *Client) DeleteDynamicRoute(request *DeleteDynamicRouteRequest) (_result *DeleteDynamicRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDynamicRouteResponse{}
	_body, _err := client.DeleteDynamicRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除指定自定义身份源部门
//
// @param request - DeleteIdpDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIdpDepartmentResponse
func (client *Client) DeleteIdpDepartmentWithOptions(request *DeleteIdpDepartmentRequest, runtime *util.RuntimeOptions) (_result *DeleteIdpDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.IdpConfigId)) {
		query["IdpConfigId"] = request.IdpConfigId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteIdpDepartment"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteIdpDepartmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除指定自定义身份源部门
//
// @param request - DeleteIdpDepartmentRequest
//
// @return DeleteIdpDepartmentResponse
func (client *Client) DeleteIdpDepartment(request *DeleteIdpDepartmentRequest) (_result *DeleteIdpDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteIdpDepartmentResponse{}
	_body, _err := client.DeleteIdpDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除内网访问应用
//
// @param request - DeletePrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateAccessApplicationResponse
func (client *Client) DeletePrivateAccessApplicationWithOptions(request *DeletePrivateAccessApplicationRequest, runtime *util.RuntimeOptions) (_result *DeletePrivateAccessApplicationResponse, _err error) {
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
		Action:      tea.String("DeletePrivateAccessApplication"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePrivateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除内网访问应用
//
// @param request - DeletePrivateAccessApplicationRequest
//
// @return DeletePrivateAccessApplicationResponse
func (client *Client) DeletePrivateAccessApplication(request *DeletePrivateAccessApplicationRequest) (_result *DeletePrivateAccessApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePrivateAccessApplicationResponse{}
	_body, _err := client.DeletePrivateAccessApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除内网访问策略
//
// @param request - DeletePrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateAccessPolicyResponse
func (client *Client) DeletePrivateAccessPolicyWithOptions(request *DeletePrivateAccessPolicyRequest, runtime *util.RuntimeOptions) (_result *DeletePrivateAccessPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		body["PolicyId"] = request.PolicyId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePrivateAccessPolicy"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除内网访问策略
//
// @param request - DeletePrivateAccessPolicyRequest
//
// @return DeletePrivateAccessPolicyResponse
func (client *Client) DeletePrivateAccessPolicy(request *DeletePrivateAccessPolicyRequest) (_result *DeletePrivateAccessPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePrivateAccessPolicyResponse{}
	_body, _err := client.DeletePrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除内网访问标签
//
// @param request - DeletePrivateAccessTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePrivateAccessTagResponse
func (client *Client) DeletePrivateAccessTagWithOptions(request *DeletePrivateAccessTagRequest, runtime *util.RuntimeOptions) (_result *DeletePrivateAccessTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TagId)) {
		body["TagId"] = request.TagId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePrivateAccessTag"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePrivateAccessTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除内网访问标签
//
// @param request - DeletePrivateAccessTagRequest
//
// @return DeletePrivateAccessTagResponse
func (client *Client) DeletePrivateAccessTag(request *DeletePrivateAccessTagRequest) (_result *DeletePrivateAccessTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePrivateAccessTagResponse{}
	_body, _err := client.DeletePrivateAccessTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除设备注册策略
//
// @param request - DeleteRegistrationPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRegistrationPoliciesResponse
func (client *Client) DeleteRegistrationPoliciesWithOptions(request *DeleteRegistrationPoliciesRequest, runtime *util.RuntimeOptions) (_result *DeleteRegistrationPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PolicyIds)) {
		bodyFlat["PolicyIds"] = request.PolicyIds
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRegistrationPolicies"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRegistrationPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除设备注册策略
//
// @param request - DeleteRegistrationPoliciesRequest
//
// @return DeleteRegistrationPoliciesResponse
func (client *Client) DeleteRegistrationPolicies(request *DeleteRegistrationPoliciesRequest) (_result *DeleteRegistrationPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRegistrationPoliciesResponse{}
	_body, _err := client.DeleteRegistrationPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量删除用户非在线设备
//
// @param request - DeleteUserDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserDevicesResponse
func (client *Client) DeleteUserDevicesWithOptions(request *DeleteUserDevicesRequest, runtime *util.RuntimeOptions) (_result *DeleteUserDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceTags)) {
		bodyFlat["DeviceTags"] = request.DeviceTags
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserDevices"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除用户非在线设备
//
// @param request - DeleteUserDevicesRequest
//
// @return DeleteUserDevicesResponse
func (client *Client) DeleteUserDevices(request *DeleteUserDevicesRequest) (_result *DeleteUserDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserDevicesResponse{}
	_body, _err := client.DeleteUserDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除用户组
//
// @param request - DeleteUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroupWithOptions(request *DeleteUserGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		body["UserGroupId"] = request.UserGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserGroup"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除用户组
//
// @param request - DeleteUserGroupRequest
//
// @return DeleteUserGroupResponse
func (client *Client) DeleteUserGroup(request *DeleteUserGroupRequest) (_result *DeleteUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserGroupResponse{}
	_body, _err := client.DeleteUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 卸载connector的应用
//
// @param tmpReq - DetachApplication2ConnectorRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetachApplication2ConnectorResponse
func (client *Client) DetachApplication2ConnectorWithOptions(tmpReq *DetachApplication2ConnectorRequest, runtime *util.RuntimeOptions) (_result *DetachApplication2ConnectorResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DetachApplication2ConnectorShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ApplicationIds)) {
		request.ApplicationIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationIds, tea.String("ApplicationIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationIdsShrink)) {
		body["ApplicationIds"] = request.ApplicationIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ConnectorId)) {
		body["ConnectorId"] = request.ConnectorId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DetachApplication2Connector"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DetachApplication2ConnectorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 卸载connector的应用
//
// @param request - DetachApplication2ConnectorRequest
//
// @return DetachApplication2ConnectorResponse
func (client *Client) DetachApplication2Connector(request *DetachApplication2ConnectorRequest) (_result *DetachApplication2ConnectorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DetachApplication2ConnectorResponse{}
	_body, _err := client.DetachApplication2ConnectorWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询用户设备列表
//
// @param request - ExportUserDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExportUserDevicesResponse
func (client *Client) ExportUserDevicesWithOptions(request *ExportUserDevicesRequest, runtime *util.RuntimeOptions) (_result *ExportUserDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppStatuses)) {
		bodyFlat["AppStatuses"] = request.AppStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.Department)) {
		body["Department"] = request.Department
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceBelong)) {
		body["DeviceBelong"] = request.DeviceBelong
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceStatuses)) {
		bodyFlat["DeviceStatuses"] = request.DeviceStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceTags)) {
		bodyFlat["DeviceTags"] = request.DeviceTags
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceTypes)) {
		bodyFlat["DeviceTypes"] = request.DeviceTypes
	}

	if !tea.BoolValue(util.IsUnset(request.DlpStatuses)) {
		bodyFlat["DlpStatuses"] = request.DlpStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.Hostname)) {
		body["Hostname"] = request.Hostname
	}

	if !tea.BoolValue(util.IsUnset(request.IaStatuses)) {
		bodyFlat["IaStatuses"] = request.IaStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.Mac)) {
		body["Mac"] = request.Mac
	}

	if !tea.BoolValue(util.IsUnset(request.NacStatuses)) {
		bodyFlat["NacStatuses"] = request.NacStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.PaStatuses)) {
		bodyFlat["PaStatuses"] = request.PaStatuses
	}

	if !tea.BoolValue(util.IsUnset(request.SaseUserId)) {
		body["SaseUserId"] = request.SaseUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SharingStatus)) {
		body["SharingStatus"] = request.SharingStatus
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["Username"] = request.Username
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExportUserDevices"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExportUserDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询用户设备列表
//
// @param request - ExportUserDevicesRequest
//
// @return ExportUserDevicesResponse
func (client *Client) ExportUserDevices(request *ExportUserDevicesRequest) (_result *ExportUserDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExportUserDevicesResponse{}
	_body, _err := client.ExportUserDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询已启用的身份源配置
//
// @param request - GetActiveIdpConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetActiveIdpConfigResponse
func (client *Client) GetActiveIdpConfigWithOptions(runtime *util.RuntimeOptions) (_result *GetActiveIdpConfigResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("GetActiveIdpConfig"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetActiveIdpConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询已启用的身份源配置
//
// @return GetActiveIdpConfigResponse
func (client *Client) GetActiveIdpConfig() (_result *GetActiveIdpConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetActiveIdpConfigResponse{}
	_body, _err := client.GetActiveIdpConfigWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询自定义身份源指定用户
//
// @param request - GetClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetClientUserResponse
func (client *Client) GetClientUserWithOptions(request *GetClientUserRequest, runtime *util.RuntimeOptions) (_result *GetClientUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetClientUser"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetClientUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自定义身份源指定用户
//
// @param request - GetClientUserRequest
//
// @return GetClientUserResponse
func (client *Client) GetClientUser(request *GetClientUserRequest) (_result *GetClientUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClientUserResponse{}
	_body, _err := client.GetClientUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询动态路由详情
//
// @param request - GetDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDynamicRouteResponse
func (client *Client) GetDynamicRouteWithOptions(request *GetDynamicRouteRequest, runtime *util.RuntimeOptions) (_result *GetDynamicRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDynamicRoute"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDynamicRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询动态路由详情
//
// @param request - GetDynamicRouteRequest
//
// @return GetDynamicRouteResponse
func (client *Client) GetDynamicRoute(request *GetDynamicRouteRequest) (_result *GetDynamicRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDynamicRouteResponse{}
	_body, _err := client.GetDynamicRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询身份源配置详情
//
// @param request - GetIdpConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdpConfigResponse
func (client *Client) GetIdpConfigWithOptions(request *GetIdpConfigRequest, runtime *util.RuntimeOptions) (_result *GetIdpConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIdpConfig"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIdpConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询身份源配置详情
//
// @param request - GetIdpConfigRequest
//
// @return GetIdpConfigResponse
func (client *Client) GetIdpConfig(request *GetIdpConfigRequest) (_result *GetIdpConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIdpConfigResponse{}
	_body, _err := client.GetIdpConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询内网访问应用详情
//
// @param request - GetPrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrivateAccessApplicationResponse
func (client *Client) GetPrivateAccessApplicationWithOptions(request *GetPrivateAccessApplicationRequest, runtime *util.RuntimeOptions) (_result *GetPrivateAccessApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPrivateAccessApplication"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPrivateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询内网访问应用详情
//
// @param request - GetPrivateAccessApplicationRequest
//
// @return GetPrivateAccessApplicationResponse
func (client *Client) GetPrivateAccessApplication(request *GetPrivateAccessApplicationRequest) (_result *GetPrivateAccessApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPrivateAccessApplicationResponse{}
	_body, _err := client.GetPrivateAccessApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询内网访问策略详情
//
// @param request - GetPrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPrivateAccessPolicyResponse
func (client *Client) GetPrivateAccessPolicyWithOptions(request *GetPrivateAccessPolicyRequest, runtime *util.RuntimeOptions) (_result *GetPrivateAccessPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPrivateAccessPolicy"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询内网访问策略详情
//
// @param request - GetPrivateAccessPolicyRequest
//
// @return GetPrivateAccessPolicyResponse
func (client *Client) GetPrivateAccessPolicy(request *GetPrivateAccessPolicyRequest) (_result *GetPrivateAccessPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPrivateAccessPolicyResponse{}
	_body, _err := client.GetPrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备注册策略详情
//
// @param request - GetRegistrationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRegistrationPolicyResponse
func (client *Client) GetRegistrationPolicyWithOptions(request *GetRegistrationPolicyRequest, runtime *util.RuntimeOptions) (_result *GetRegistrationPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRegistrationPolicy"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRegistrationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备注册策略详情
//
// @param request - GetRegistrationPolicyRequest
//
// @return GetRegistrationPolicyResponse
func (client *Client) GetRegistrationPolicy(request *GetRegistrationPolicyRequest) (_result *GetRegistrationPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRegistrationPolicyResponse{}
	_body, _err := client.GetRegistrationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户设备详情
//
// @param request - GetUserDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserDeviceResponse
func (client *Client) GetUserDeviceWithOptions(request *GetUserDeviceRequest, runtime *util.RuntimeOptions) (_result *GetUserDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserDevice"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户设备详情
//
// @param request - GetUserDeviceRequest
//
// @return GetUserDeviceResponse
func (client *Client) GetUserDevice(request *GetUserDeviceRequest) (_result *GetUserDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserDeviceResponse{}
	_body, _err := client.GetUserDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户组详情
//
// @param request - GetUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroupWithOptions(request *GetUserGroupRequest, runtime *util.RuntimeOptions) (_result *GetUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUserGroup"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户组详情
//
// @param request - GetUserGroupRequest
//
// @return GetUserGroupResponse
func (client *Client) GetUserGroup(request *GetUserGroupRequest) (_result *GetUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserGroupResponse{}
	_body, _err := client.GetUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询嵌入水印任务
//
// @param request - GetWmEmbedTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWmEmbedTaskResponse
func (client *Client) GetWmEmbedTaskWithOptions(request *GetWmEmbedTaskRequest, runtime *util.RuntimeOptions) (_result *GetWmEmbedTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWmEmbedTask"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWmEmbedTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询嵌入水印任务
//
// @param request - GetWmEmbedTaskRequest
//
// @return GetWmEmbedTaskResponse
func (client *Client) GetWmEmbedTask(request *GetWmEmbedTaskRequest) (_result *GetWmEmbedTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWmEmbedTaskResponse{}
	_body, _err := client.GetWmEmbedTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询文件水印提取任务详情
//
// @param request - GetWmExtractTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWmExtractTaskResponse
func (client *Client) GetWmExtractTaskWithOptions(request *GetWmExtractTaskRequest, runtime *util.RuntimeOptions) (_result *GetWmExtractTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetWmExtractTask"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetWmExtractTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文件水印提取任务详情
//
// @param request - GetWmExtractTaskRequest
//
// @return GetWmExtractTaskResponse
func (client *Client) GetWmExtractTask(request *GetWmExtractTaskRequest) (_result *GetWmExtractTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWmExtractTaskResponse{}
	_body, _err := client.GetWmExtractTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询内网访问策略的应用
//
// @param request - ListApplicationsForPrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForPrivateAccessPolicyResponse
func (client *Client) ListApplicationsForPrivateAccessPolicyWithOptions(request *ListApplicationsForPrivateAccessPolicyRequest, runtime *util.RuntimeOptions) (_result *ListApplicationsForPrivateAccessPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplicationsForPrivateAccessPolicy"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationsForPrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问策略的应用
//
// @param request - ListApplicationsForPrivateAccessPolicyRequest
//
// @return ListApplicationsForPrivateAccessPolicyResponse
func (client *Client) ListApplicationsForPrivateAccessPolicy(request *ListApplicationsForPrivateAccessPolicyRequest) (_result *ListApplicationsForPrivateAccessPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationsForPrivateAccessPolicyResponse{}
	_body, _err := client.ListApplicationsForPrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询内网访问标签的应用
//
// @param request - ListApplicationsForPrivateAccessTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApplicationsForPrivateAccessTagResponse
func (client *Client) ListApplicationsForPrivateAccessTagWithOptions(request *ListApplicationsForPrivateAccessTagRequest, runtime *util.RuntimeOptions) (_result *ListApplicationsForPrivateAccessTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApplicationsForPrivateAccessTag"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListApplicationsForPrivateAccessTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问标签的应用
//
// @param request - ListApplicationsForPrivateAccessTagRequest
//
// @return ListApplicationsForPrivateAccessTagResponse
func (client *Client) ListApplicationsForPrivateAccessTag(request *ListApplicationsForPrivateAccessTagRequest) (_result *ListApplicationsForPrivateAccessTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationsForPrivateAccessTagResponse{}
	_body, _err := client.ListApplicationsForPrivateAccessTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询自定义身份源用户
//
// @param request - ListClientUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClientUsersResponse
func (client *Client) ListClientUsersWithOptions(request *ListClientUsersRequest, runtime *util.RuntimeOptions) (_result *ListClientUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClientUsers"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClientUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自定义身份源用户
//
// @param request - ListClientUsersRequest
//
// @return ListClientUsersResponse
func (client *Client) ListClientUsers(request *ListClientUsersRequest) (_result *ListClientUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClientUsersResponse{}
	_body, _err := client.ListClientUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询connector
//
// @param request - ListConnectorsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConnectorsResponse
func (client *Client) ListConnectorsWithOptions(request *ListConnectorsRequest, runtime *util.RuntimeOptions) (_result *ListConnectorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListConnectors"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListConnectorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询connector
//
// @param request - ListConnectorsRequest
//
// @return ListConnectorsResponse
func (client *Client) ListConnectors(request *ListConnectorsRequest) (_result *ListConnectorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListConnectorsResponse{}
	_body, _err := client.ListConnectorsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询动态路由的地域
//
// @param request - ListDynamicRouteRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDynamicRouteRegionsResponse
func (client *Client) ListDynamicRouteRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListDynamicRouteRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListDynamicRouteRegions"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDynamicRouteRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询动态路由的地域
//
// @return ListDynamicRouteRegionsResponse
func (client *Client) ListDynamicRouteRegions() (_result *ListDynamicRouteRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDynamicRouteRegionsResponse{}
	_body, _err := client.ListDynamicRouteRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询动态路由
//
// @param request - ListDynamicRoutesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDynamicRoutesResponse
func (client *Client) ListDynamicRoutesWithOptions(request *ListDynamicRoutesRequest, runtime *util.RuntimeOptions) (_result *ListDynamicRoutesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDynamicRoutes"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDynamicRoutesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询动态路由
//
// @param request - ListDynamicRoutesRequest
//
// @return ListDynamicRoutesResponse
func (client *Client) ListDynamicRoutes(request *ListDynamicRoutesRequest) (_result *ListDynamicRoutesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDynamicRoutesResponse{}
	_body, _err := client.ListDynamicRoutesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询超额注册申请列表
//
// @param request - ListExcessiveDeviceRegistrationApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExcessiveDeviceRegistrationApplicationsResponse
func (client *Client) ListExcessiveDeviceRegistrationApplicationsWithOptions(request *ListExcessiveDeviceRegistrationApplicationsRequest, runtime *util.RuntimeOptions) (_result *ListExcessiveDeviceRegistrationApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListExcessiveDeviceRegistrationApplications"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListExcessiveDeviceRegistrationApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询超额注册申请列表
//
// @param request - ListExcessiveDeviceRegistrationApplicationsRequest
//
// @return ListExcessiveDeviceRegistrationApplicationsResponse
func (client *Client) ListExcessiveDeviceRegistrationApplications(request *ListExcessiveDeviceRegistrationApplicationsRequest) (_result *ListExcessiveDeviceRegistrationApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListExcessiveDeviceRegistrationApplicationsResponse{}
	_body, _err := client.ListExcessiveDeviceRegistrationApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询IDP配置
//
// @param request - ListIdpConfigsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdpConfigsResponse
func (client *Client) ListIdpConfigsWithOptions(request *ListIdpConfigsRequest, runtime *util.RuntimeOptions) (_result *ListIdpConfigsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIdpConfigs"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIdpConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询IDP配置
//
// @param request - ListIdpConfigsRequest
//
// @return ListIdpConfigsResponse
func (client *Client) ListIdpConfigs(request *ListIdpConfigsRequest) (_result *ListIdpConfigsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIdpConfigsResponse{}
	_body, _err := client.ListIdpConfigsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询自定义身份源部门
//
// @param request - ListIdpDepartmentsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdpDepartmentsResponse
func (client *Client) ListIdpDepartmentsWithOptions(request *ListIdpDepartmentsRequest, runtime *util.RuntimeOptions) (_result *ListIdpDepartmentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIdpDepartments"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIdpDepartmentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询自定义身份源部门
//
// @param request - ListIdpDepartmentsRequest
//
// @return ListIdpDepartmentsResponse
func (client *Client) ListIdpDepartments(request *ListIdpDepartmentsRequest) (_result *ListIdpDepartmentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListIdpDepartmentsResponse{}
	_body, _err := client.ListIdpDepartmentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 入网用户列表
//
// @param request - ListNacUserCertRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNacUserCertResponse
func (client *Client) ListNacUserCertWithOptions(request *ListNacUserCertRequest, runtime *util.RuntimeOptions) (_result *ListNacUserCertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.Department)) {
		query["Department"] = request.Department
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		query["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNacUserCert"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNacUserCertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 入网用户列表
//
// @param request - ListNacUserCertRequest
//
// @return ListNacUserCertResponse
func (client *Client) ListNacUserCert(request *ListNacUserCertRequest) (_result *ListNacUserCertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNacUserCertResponse{}
	_body, _err := client.ListNacUserCertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询内网访问应用的策略
//
// @param request - ListPolicesForPrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicesForPrivateAccessApplicationResponse
func (client *Client) ListPolicesForPrivateAccessApplicationWithOptions(request *ListPolicesForPrivateAccessApplicationRequest, runtime *util.RuntimeOptions) (_result *ListPolicesForPrivateAccessApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicesForPrivateAccessApplication"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPolicesForPrivateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问应用的策略
//
// @param request - ListPolicesForPrivateAccessApplicationRequest
//
// @return ListPolicesForPrivateAccessApplicationResponse
func (client *Client) ListPolicesForPrivateAccessApplication(request *ListPolicesForPrivateAccessApplicationRequest) (_result *ListPolicesForPrivateAccessApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPolicesForPrivateAccessApplicationResponse{}
	_body, _err := client.ListPolicesForPrivateAccessApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询内网访问标签的策略
//
// @param request - ListPolicesForPrivateAccessTagRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicesForPrivateAccessTagResponse
func (client *Client) ListPolicesForPrivateAccessTagWithOptions(request *ListPolicesForPrivateAccessTagRequest, runtime *util.RuntimeOptions) (_result *ListPolicesForPrivateAccessTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicesForPrivateAccessTag"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPolicesForPrivateAccessTagResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问标签的策略
//
// @param request - ListPolicesForPrivateAccessTagRequest
//
// @return ListPolicesForPrivateAccessTagResponse
func (client *Client) ListPolicesForPrivateAccessTag(request *ListPolicesForPrivateAccessTagRequest) (_result *ListPolicesForPrivateAccessTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPolicesForPrivateAccessTagResponse{}
	_body, _err := client.ListPolicesForPrivateAccessTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询用户组的策略
//
// @param request - ListPolicesForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPolicesForUserGroupResponse
func (client *Client) ListPolicesForUserGroupWithOptions(request *ListPolicesForUserGroupRequest, runtime *util.RuntimeOptions) (_result *ListPolicesForUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPolicesForUserGroup"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPolicesForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询用户组的策略
//
// @param request - ListPolicesForUserGroupRequest
//
// @return ListPolicesForUserGroupResponse
func (client *Client) ListPolicesForUserGroup(request *ListPolicesForUserGroupRequest) (_result *ListPolicesForUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPolicesForUserGroupResponse{}
	_body, _err := client.ListPolicesForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// pop节点流量统计
//
// @param request - ListPopTrafficStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPopTrafficStatisticsResponse
func (client *Client) ListPopTrafficStatisticsWithOptions(request *ListPopTrafficStatisticsRequest, runtime *util.RuntimeOptions) (_result *ListPopTrafficStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPopTrafficStatistics"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPopTrafficStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// pop节点流量统计
//
// @param request - ListPopTrafficStatisticsRequest
//
// @return ListPopTrafficStatisticsResponse
func (client *Client) ListPopTrafficStatistics(request *ListPopTrafficStatisticsRequest) (_result *ListPopTrafficStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPopTrafficStatisticsResponse{}
	_body, _err := client.ListPopTrafficStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询内网访问应用
//
// @param request - ListPrivateAccessApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessApplicationsResponse
func (client *Client) ListPrivateAccessApplicationsWithOptions(request *ListPrivateAccessApplicationsRequest, runtime *util.RuntimeOptions) (_result *ListPrivateAccessApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPrivateAccessApplications"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrivateAccessApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问应用
//
// @param request - ListPrivateAccessApplicationsRequest
//
// @return ListPrivateAccessApplicationsResponse
func (client *Client) ListPrivateAccessApplications(request *ListPrivateAccessApplicationsRequest) (_result *ListPrivateAccessApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrivateAccessApplicationsResponse{}
	_body, _err := client.ListPrivateAccessApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询动态路由的内网访问应用
//
// @param request - ListPrivateAccessApplicationsForDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessApplicationsForDynamicRouteResponse
func (client *Client) ListPrivateAccessApplicationsForDynamicRouteWithOptions(request *ListPrivateAccessApplicationsForDynamicRouteRequest, runtime *util.RuntimeOptions) (_result *ListPrivateAccessApplicationsForDynamicRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPrivateAccessApplicationsForDynamicRoute"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrivateAccessApplicationsForDynamicRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询动态路由的内网访问应用
//
// @param request - ListPrivateAccessApplicationsForDynamicRouteRequest
//
// @return ListPrivateAccessApplicationsForDynamicRouteResponse
func (client *Client) ListPrivateAccessApplicationsForDynamicRoute(request *ListPrivateAccessApplicationsForDynamicRouteRequest) (_result *ListPrivateAccessApplicationsForDynamicRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrivateAccessApplicationsForDynamicRouteResponse{}
	_body, _err := client.ListPrivateAccessApplicationsForDynamicRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询内网访问策略
//
// @param request - ListPrivateAccessPolicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessPolicesResponse
func (client *Client) ListPrivateAccessPolicesWithOptions(request *ListPrivateAccessPolicesRequest, runtime *util.RuntimeOptions) (_result *ListPrivateAccessPolicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPrivateAccessPolices"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrivateAccessPolicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问策略
//
// @param request - ListPrivateAccessPolicesRequest
//
// @return ListPrivateAccessPolicesResponse
func (client *Client) ListPrivateAccessPolices(request *ListPrivateAccessPolicesRequest) (_result *ListPrivateAccessPolicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrivateAccessPolicesResponse{}
	_body, _err := client.ListPrivateAccessPolicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about all internal access tags within the current Alibaba Cloud account.
//
// @param request - ListPrivateAccessTagsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessTagsResponse
func (client *Client) ListPrivateAccessTagsWithOptions(request *ListPrivateAccessTagsRequest, runtime *util.RuntimeOptions) (_result *ListPrivateAccessTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPrivateAccessTags"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrivateAccessTagsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all internal access tags within the current Alibaba Cloud account.
//
// @param request - ListPrivateAccessTagsRequest
//
// @return ListPrivateAccessTagsResponse
func (client *Client) ListPrivateAccessTags(request *ListPrivateAccessTagsRequest) (_result *ListPrivateAccessTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrivateAccessTagsResponse{}
	_body, _err := client.ListPrivateAccessTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询动态路由的内网访问标签
//
// @param request - ListPrivateAccessTagsForDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPrivateAccessTagsForDynamicRouteResponse
func (client *Client) ListPrivateAccessTagsForDynamicRouteWithOptions(request *ListPrivateAccessTagsForDynamicRouteRequest, runtime *util.RuntimeOptions) (_result *ListPrivateAccessTagsForDynamicRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPrivateAccessTagsForDynamicRoute"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPrivateAccessTagsForDynamicRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询动态路由的内网访问标签
//
// @param request - ListPrivateAccessTagsForDynamicRouteRequest
//
// @return ListPrivateAccessTagsForDynamicRouteResponse
func (client *Client) ListPrivateAccessTagsForDynamicRoute(request *ListPrivateAccessTagsForDynamicRouteRequest) (_result *ListPrivateAccessTagsForDynamicRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPrivateAccessTagsForDynamicRouteResponse{}
	_body, _err := client.ListPrivateAccessTagsForDynamicRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户设备注册策略列表
//
// @param request - ListRegistrationPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegistrationPoliciesResponse
func (client *Client) ListRegistrationPoliciesWithOptions(request *ListRegistrationPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListRegistrationPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRegistrationPolicies"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegistrationPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户设备注册策略列表
//
// @param request - ListRegistrationPoliciesRequest
//
// @return ListRegistrationPoliciesResponse
func (client *Client) ListRegistrationPolicies(request *ListRegistrationPoliciesRequest) (_result *ListRegistrationPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegistrationPoliciesResponse{}
	_body, _err := client.ListRegistrationPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询用户组相关的设备注册策略
//
// @param request - ListRegistrationPoliciesForUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRegistrationPoliciesForUserGroupResponse
func (client *Client) ListRegistrationPoliciesForUserGroupWithOptions(request *ListRegistrationPoliciesForUserGroupRequest, runtime *util.RuntimeOptions) (_result *ListRegistrationPoliciesForUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRegistrationPoliciesForUserGroup"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegistrationPoliciesForUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户组相关的设备注册策略
//
// @param request - ListRegistrationPoliciesForUserGroupRequest
//
// @return ListRegistrationPoliciesForUserGroupResponse
func (client *Client) ListRegistrationPoliciesForUserGroup(request *ListRegistrationPoliciesForUserGroupRequest) (_result *ListRegistrationPoliciesForUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegistrationPoliciesForUserGroupResponse{}
	_body, _err := client.ListRegistrationPoliciesForUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询终端安装软件列表
//
// @param request - ListSoftwareForUserDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSoftwareForUserDeviceResponse
func (client *Client) ListSoftwareForUserDeviceWithOptions(request *ListSoftwareForUserDeviceRequest, runtime *util.RuntimeOptions) (_result *ListSoftwareForUserDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListSoftwareForUserDevice"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSoftwareForUserDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询终端安装软件列表
//
// @param request - ListSoftwareForUserDeviceRequest
//
// @return ListSoftwareForUserDeviceResponse
func (client *Client) ListSoftwareForUserDevice(request *ListSoftwareForUserDeviceRequest) (_result *ListSoftwareForUserDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSoftwareForUserDeviceResponse{}
	_body, _err := client.ListSoftwareForUserDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询内网访问应用的标签
//
// @param request - ListTagsForPrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsForPrivateAccessApplicationResponse
func (client *Client) ListTagsForPrivateAccessApplicationWithOptions(request *ListTagsForPrivateAccessApplicationRequest, runtime *util.RuntimeOptions) (_result *ListTagsForPrivateAccessApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagsForPrivateAccessApplication"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagsForPrivateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问应用的标签
//
// @param request - ListTagsForPrivateAccessApplicationRequest
//
// @return ListTagsForPrivateAccessApplicationResponse
func (client *Client) ListTagsForPrivateAccessApplication(request *ListTagsForPrivateAccessApplicationRequest) (_result *ListTagsForPrivateAccessApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagsForPrivateAccessApplicationResponse{}
	_body, _err := client.ListTagsForPrivateAccessApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询内网访问策略的标签
//
// @param request - ListTagsForPrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsForPrivateAccessPolicyResponse
func (client *Client) ListTagsForPrivateAccessPolicyWithOptions(request *ListTagsForPrivateAccessPolicyRequest, runtime *util.RuntimeOptions) (_result *ListTagsForPrivateAccessPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagsForPrivateAccessPolicy"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagsForPrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问策略的标签
//
// @param request - ListTagsForPrivateAccessPolicyRequest
//
// @return ListTagsForPrivateAccessPolicyResponse
func (client *Client) ListTagsForPrivateAccessPolicy(request *ListTagsForPrivateAccessPolicyRequest) (_result *ListTagsForPrivateAccessPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagsForPrivateAccessPolicyResponse{}
	_body, _err := client.ListTagsForPrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列表查询用户应用权限
//
// @param request - ListUserApplicationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserApplicationsResponse
func (client *Client) ListUserApplicationsWithOptions(request *ListUserApplicationsRequest, runtime *util.RuntimeOptions) (_result *ListUserApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserApplications"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserApplicationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询用户应用权限
//
// @param request - ListUserApplicationsRequest
//
// @return ListUserApplicationsResponse
func (client *Client) ListUserApplications(request *ListUserApplicationsRequest) (_result *ListUserApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserApplicationsResponse{}
	_body, _err := client.ListUserApplicationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询用户设备列表
//
// @param request - ListUserDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserDevicesResponse
func (client *Client) ListUserDevicesWithOptions(request *ListUserDevicesRequest, runtime *util.RuntimeOptions) (_result *ListUserDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserDevices"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询用户设备列表
//
// @param request - ListUserDevicesRequest
//
// @return ListUserDevicesResponse
func (client *Client) ListUserDevices(request *ListUserDevicesRequest) (_result *ListUserDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserDevicesResponse{}
	_body, _err := client.ListUserDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询用户组
//
// @param request - ListUserGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroupsWithOptions(request *ListUserGroupsRequest, runtime *util.RuntimeOptions) (_result *ListUserGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserGroups"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询用户组
//
// @param request - ListUserGroupsRequest
//
// @return ListUserGroupsResponse
func (client *Client) ListUserGroups(request *ListUserGroupsRequest) (_result *ListUserGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserGroupsResponse{}
	_body, _err := client.ListUserGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量查询内网访问策略的用户组
//
// @param request - ListUserGroupsForPrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsForPrivateAccessPolicyResponse
func (client *Client) ListUserGroupsForPrivateAccessPolicyWithOptions(request *ListUserGroupsForPrivateAccessPolicyRequest, runtime *util.RuntimeOptions) (_result *ListUserGroupsForPrivateAccessPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserGroupsForPrivateAccessPolicy"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserGroupsForPrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询内网访问策略的用户组
//
// @param request - ListUserGroupsForPrivateAccessPolicyRequest
//
// @return ListUserGroupsForPrivateAccessPolicyResponse
func (client *Client) ListUserGroupsForPrivateAccessPolicy(request *ListUserGroupsForPrivateAccessPolicyRequest) (_result *ListUserGroupsForPrivateAccessPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserGroupsForPrivateAccessPolicyResponse{}
	_body, _err := client.ListUserGroupsForPrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询设备注册策略相关用户组
//
// @param request - ListUserGroupsForRegistrationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserGroupsForRegistrationPolicyResponse
func (client *Client) ListUserGroupsForRegistrationPolicyWithOptions(request *ListUserGroupsForRegistrationPolicyRequest, runtime *util.RuntimeOptions) (_result *ListUserGroupsForRegistrationPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserGroupsForRegistrationPolicy"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserGroupsForRegistrationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备注册策略相关用户组
//
// @param request - ListUserGroupsForRegistrationPolicyRequest
//
// @return ListUserGroupsForRegistrationPolicyResponse
func (client *Client) ListUserGroupsForRegistrationPolicy(request *ListUserGroupsForRegistrationPolicyRequest) (_result *ListUserGroupsForRegistrationPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserGroupsForRegistrationPolicyResponse{}
	_body, _err := client.ListUserGroupsForRegistrationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列表查询用户零信任策略
//
// @param request - ListUserPrivateAccessPoliciesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserPrivateAccessPoliciesResponse
func (client *Client) ListUserPrivateAccessPoliciesWithOptions(request *ListUserPrivateAccessPoliciesRequest, runtime *util.RuntimeOptions) (_result *ListUserPrivateAccessPoliciesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserPrivateAccessPolicies"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserPrivateAccessPoliciesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询用户零信任策略
//
// @param request - ListUserPrivateAccessPoliciesRequest
//
// @return ListUserPrivateAccessPoliciesResponse
func (client *Client) ListUserPrivateAccessPolicies(request *ListUserPrivateAccessPoliciesRequest) (_result *ListUserPrivateAccessPoliciesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserPrivateAccessPoliciesResponse{}
	_body, _err := client.ListUserPrivateAccessPoliciesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列表查询登陆用户
//
// @param request - ListUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUsersResponse
func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUsers"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列表查询登陆用户
//
// @param request - ListUsersRequest
//
// @return ListUsersResponse
func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据数字水印信息查询字符串水印信息
//
// @param request - LookupWmInfoMappingRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LookupWmInfoMappingResponse
func (client *Client) LookupWmInfoMappingWithOptions(request *LookupWmInfoMappingRequest, runtime *util.RuntimeOptions) (_result *LookupWmInfoMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LookupWmInfoMapping"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LookupWmInfoMappingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据数字水印信息查询字符串水印信息
//
// @param request - LookupWmInfoMappingRequest
//
// @return LookupWmInfoMappingResponse
func (client *Client) LookupWmInfoMapping(request *LookupWmInfoMappingRequest) (_result *LookupWmInfoMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LookupWmInfoMappingResponse{}
	_body, _err := client.LookupWmInfoMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 吊销用户登录会话
//
// @param request - RevokeUserSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RevokeUserSessionResponse
func (client *Client) RevokeUserSessionWithOptions(request *RevokeUserSessionRequest, runtime *util.RuntimeOptions) (_result *RevokeUserSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExternalIds)) {
		query["ExternalIds"] = request.ExternalIds
	}

	if !tea.BoolValue(util.IsUnset(request.IdpId)) {
		query["IdpId"] = request.IdpId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeUserSession"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeUserSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 吊销用户登录会话
//
// @param request - RevokeUserSessionRequest
//
// @return RevokeUserSessionResponse
func (client *Client) RevokeUserSession(request *RevokeUserSessionRequest) (_result *RevokeUserSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RevokeUserSessionResponse{}
	_body, _err := client.RevokeUserSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改自定义身份源指定用户
//
// @param request - UpdateClientUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClientUserResponse
func (client *Client) UpdateClientUserWithOptions(request *UpdateClientUserRequest, runtime *util.RuntimeOptions) (_result *UpdateClientUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.MobileNumber)) {
		query["MobileNumber"] = request.MobileNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateClientUser"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateClientUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改自定义身份源指定用户
//
// @param request - UpdateClientUserRequest
//
// @return UpdateClientUserResponse
func (client *Client) UpdateClientUser(request *UpdateClientUserRequest) (_result *UpdateClientUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClientUserResponse{}
	_body, _err := client.UpdateClientUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改自定义身份源指定用户密码
//
// @param request - UpdateClientUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClientUserPasswordResponse
func (client *Client) UpdateClientUserPasswordWithOptions(request *UpdateClientUserPasswordRequest, runtime *util.RuntimeOptions) (_result *UpdateClientUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		query["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		query["Username"] = request.Username
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateClientUserPassword"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateClientUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改自定义身份源指定用户密码
//
// @param request - UpdateClientUserPasswordRequest
//
// @return UpdateClientUserPasswordResponse
func (client *Client) UpdateClientUserPassword(request *UpdateClientUserPasswordRequest) (_result *UpdateClientUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClientUserPasswordResponse{}
	_body, _err := client.UpdateClientUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改自定义身份源指定用户启用状态
//
// @param request - UpdateClientUserStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateClientUserStatusResponse
func (client *Client) UpdateClientUserStatusWithOptions(request *UpdateClientUserStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateClientUserStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		query["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateClientUserStatus"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateClientUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改自定义身份源指定用户启用状态
//
// @param request - UpdateClientUserStatusRequest
//
// @return UpdateClientUserStatusResponse
func (client *Client) UpdateClientUserStatus(request *UpdateClientUserStatusRequest) (_result *UpdateClientUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClientUserStatusResponse{}
	_body, _err := client.UpdateClientUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改动态路由
//
// @param request - UpdateDynamicRouteRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDynamicRouteResponse
func (client *Client) UpdateDynamicRouteWithOptions(request *UpdateDynamicRouteRequest, runtime *util.RuntimeOptions) (_result *UpdateDynamicRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationIds)) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationType)) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DynamicRouteId)) {
		body["DynamicRouteId"] = request.DynamicRouteId
	}

	if !tea.BoolValue(util.IsUnset(request.DynamicRouteType)) {
		body["DynamicRouteType"] = request.DynamicRouteType
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		body["ModifyType"] = request.ModifyType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NextHop)) {
		body["NextHop"] = request.NextHop
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.RegionIds)) {
		bodyFlat["RegionIds"] = request.RegionIds
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagIds)) {
		bodyFlat["TagIds"] = request.TagIds
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDynamicRoute"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDynamicRouteResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改动态路由
//
// @param request - UpdateDynamicRouteRequest
//
// @return UpdateDynamicRouteResponse
func (client *Client) UpdateDynamicRoute(request *UpdateDynamicRouteRequest) (_result *UpdateDynamicRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateDynamicRouteResponse{}
	_body, _err := client.UpdateDynamicRouteWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量更新超额注册申请状态
//
// @param request - UpdateExcessiveDeviceRegistrationApplicationsStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateExcessiveDeviceRegistrationApplicationsStatusResponse
func (client *Client) UpdateExcessiveDeviceRegistrationApplicationsStatusWithOptions(request *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationIds)) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateExcessiveDeviceRegistrationApplicationsStatus"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateExcessiveDeviceRegistrationApplicationsStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量更新超额注册申请状态
//
// @param request - UpdateExcessiveDeviceRegistrationApplicationsStatusRequest
//
// @return UpdateExcessiveDeviceRegistrationApplicationsStatusResponse
func (client *Client) UpdateExcessiveDeviceRegistrationApplicationsStatus(request *UpdateExcessiveDeviceRegistrationApplicationsStatusRequest) (_result *UpdateExcessiveDeviceRegistrationApplicationsStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateExcessiveDeviceRegistrationApplicationsStatusResponse{}
	_body, _err := client.UpdateExcessiveDeviceRegistrationApplicationsStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改指定自定义身份源部门
//
// @param request - UpdateIdpDepartmentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIdpDepartmentResponse
func (client *Client) UpdateIdpDepartmentWithOptions(request *UpdateIdpDepartmentRequest, runtime *util.RuntimeOptions) (_result *UpdateIdpDepartmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DepartmentId)) {
		query["DepartmentId"] = request.DepartmentId
	}

	if !tea.BoolValue(util.IsUnset(request.DepartmentName)) {
		query["DepartmentName"] = request.DepartmentName
	}

	if !tea.BoolValue(util.IsUnset(request.IdpConfigId)) {
		query["IdpConfigId"] = request.IdpConfigId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIdpDepartment"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIdpDepartmentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改指定自定义身份源部门
//
// @param request - UpdateIdpDepartmentRequest
//
// @return UpdateIdpDepartmentResponse
func (client *Client) UpdateIdpDepartment(request *UpdateIdpDepartmentRequest) (_result *UpdateIdpDepartmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateIdpDepartmentResponse{}
	_body, _err := client.UpdateIdpDepartmentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新NAC User 状态
//
// @param request - UpdateNacUserCertStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNacUserCertStatusResponse
func (client *Client) UpdateNacUserCertStatusWithOptions(request *UpdateNacUserCertStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateNacUserCertStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdList)) {
		bodyFlat["IdList"] = request.IdList
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNacUserCertStatus"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNacUserCertStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新NAC User 状态
//
// @param request - UpdateNacUserCertStatusRequest
//
// @return UpdateNacUserCertStatusResponse
func (client *Client) UpdateNacUserCertStatus(request *UpdateNacUserCertStatusRequest) (_result *UpdateNacUserCertStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNacUserCertStatusResponse{}
	_body, _err := client.UpdateNacUserCertStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改内网访问应用
//
// @param request - UpdatePrivateAccessApplicationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivateAccessApplicationResponse
func (client *Client) UpdatePrivateAccessApplicationWithOptions(request *UpdatePrivateAccessApplicationRequest, runtime *util.RuntimeOptions) (_result *UpdatePrivateAccessApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Addresses)) {
		bodyFlat["Addresses"] = request.Addresses
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationId)) {
		body["ApplicationId"] = request.ApplicationId
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.L7ProxyDomainAutomaticPrefix)) {
		body["L7ProxyDomainAutomaticPrefix"] = request.L7ProxyDomainAutomaticPrefix
	}

	if !tea.BoolValue(util.IsUnset(request.L7ProxyDomainCustom)) {
		body["L7ProxyDomainCustom"] = request.L7ProxyDomainCustom
	}

	if !tea.BoolValue(util.IsUnset(request.L7ProxyDomainPrivate)) {
		body["L7ProxyDomainPrivate"] = request.L7ProxyDomainPrivate
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		body["ModifyType"] = request.ModifyType
	}

	if !tea.BoolValue(util.IsUnset(request.PortRanges)) {
		bodyFlat["PortRanges"] = request.PortRanges
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagIds)) {
		bodyFlat["TagIds"] = request.TagIds
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePrivateAccessApplication"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePrivateAccessApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改内网访问应用
//
// @param request - UpdatePrivateAccessApplicationRequest
//
// @return UpdatePrivateAccessApplicationResponse
func (client *Client) UpdatePrivateAccessApplication(request *UpdatePrivateAccessApplicationRequest) (_result *UpdatePrivateAccessApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePrivateAccessApplicationResponse{}
	_body, _err := client.UpdatePrivateAccessApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改内网访问策略
//
// @param request - UpdatePrivateAccessPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePrivateAccessPolicyResponse
func (client *Client) UpdatePrivateAccessPolicyWithOptions(request *UpdatePrivateAccessPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdatePrivateAccessPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationIds)) {
		bodyFlat["ApplicationIds"] = request.ApplicationIds
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationType)) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !tea.BoolValue(util.IsUnset(request.CustomUserAttributes)) {
		bodyFlat["CustomUserAttributes"] = request.CustomUserAttributes
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceAttributeAction)) {
		body["DeviceAttributeAction"] = request.DeviceAttributeAction
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceAttributeId)) {
		body["DeviceAttributeId"] = request.DeviceAttributeId
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		body["ModifyType"] = request.ModifyType
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyAction)) {
		body["PolicyAction"] = request.PolicyAction
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		body["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagIds)) {
		bodyFlat["TagIds"] = request.TagIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupIds)) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupMode)) {
		body["UserGroupMode"] = request.UserGroupMode
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePrivateAccessPolicy"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePrivateAccessPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改内网访问策略
//
// @param request - UpdatePrivateAccessPolicyRequest
//
// @return UpdatePrivateAccessPolicyResponse
func (client *Client) UpdatePrivateAccessPolicy(request *UpdatePrivateAccessPolicyRequest) (_result *UpdatePrivateAccessPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePrivateAccessPolicyResponse{}
	_body, _err := client.UpdatePrivateAccessPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改设备注册策略
//
// @param tmpReq - UpdateRegistrationPolicyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRegistrationPolicyResponse
func (client *Client) UpdateRegistrationPolicyWithOptions(tmpReq *UpdateRegistrationPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateRegistrationPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdateRegistrationPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CompanyLimitCount)) {
		request.CompanyLimitCountShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CompanyLimitCount, tea.String("CompanyLimitCount"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PersonalLimitCount)) {
		request.PersonalLimitCountShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PersonalLimitCount, tea.String("PersonalLimitCount"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CompanyLimitCountShrink)) {
		body["CompanyLimitCount"] = request.CompanyLimitCountShrink
	}

	if !tea.BoolValue(util.IsUnset(request.CompanyLimitType)) {
		body["CompanyLimitType"] = request.CompanyLimitType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.MatchMode)) {
		body["MatchMode"] = request.MatchMode
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PersonalLimitCountShrink)) {
		body["PersonalLimitCount"] = request.PersonalLimitCountShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PersonalLimitType)) {
		body["PersonalLimitType"] = request.PersonalLimitType
	}

	if !tea.BoolValue(util.IsUnset(request.PolicyId)) {
		body["PolicyId"] = request.PolicyId
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserGroupIds)) {
		bodyFlat["UserGroupIds"] = request.UserGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.Whitelist)) {
		bodyFlat["Whitelist"] = request.Whitelist
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRegistrationPolicy"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRegistrationPolicyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改设备注册策略
//
// @param request - UpdateRegistrationPolicyRequest
//
// @return UpdateRegistrationPolicyResponse
func (client *Client) UpdateRegistrationPolicy(request *UpdateRegistrationPolicyRequest) (_result *UpdateRegistrationPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRegistrationPolicyResponse{}
	_body, _err := client.UpdateRegistrationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量更新用户设备共享状态
//
// @param request - UpdateUserDevicesSharingStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserDevicesSharingStatusResponse
func (client *Client) UpdateUserDevicesSharingStatusWithOptions(request *UpdateUserDevicesSharingStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateUserDevicesSharingStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceTags)) {
		bodyFlat["DeviceTags"] = request.DeviceTags
	}

	if !tea.BoolValue(util.IsUnset(request.SharingStatus)) {
		body["SharingStatus"] = request.SharingStatus
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserDevicesSharingStatus"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserDevicesSharingStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量更新用户设备共享状态
//
// @param request - UpdateUserDevicesSharingStatusRequest
//
// @return UpdateUserDevicesSharingStatusResponse
func (client *Client) UpdateUserDevicesSharingStatus(request *UpdateUserDevicesSharingStatusRequest) (_result *UpdateUserDevicesSharingStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserDevicesSharingStatusResponse{}
	_body, _err := client.UpdateUserDevicesSharingStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量更新用户设备状态
//
// @param request - UpdateUserDevicesStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserDevicesStatusResponse
func (client *Client) UpdateUserDevicesStatusWithOptions(request *UpdateUserDevicesStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateUserDevicesStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceAction)) {
		body["DeviceAction"] = request.DeviceAction
	}

	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceTags)) {
		bodyFlat["DeviceTags"] = request.DeviceTags
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserDevicesStatus"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserDevicesStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量更新用户设备状态
//
// @param request - UpdateUserDevicesStatusRequest
//
// @return UpdateUserDevicesStatusResponse
func (client *Client) UpdateUserDevicesStatus(request *UpdateUserDevicesStatusRequest) (_result *UpdateUserDevicesStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserDevicesStatusResponse{}
	_body, _err := client.UpdateUserDevicesStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改用户组
//
// @param request - UpdateUserGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroupWithOptions(request *UpdateUserGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateUserGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Attributes)) {
		bodyFlat["Attributes"] = request.Attributes
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		body["ModifyType"] = request.ModifyType
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupId)) {
		body["UserGroupId"] = request.UserGroupId
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUserGroup"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改用户组
//
// @param request - UpdateUserGroupRequest
//
// @return UpdateUserGroupResponse
func (client *Client) UpdateUserGroup(request *UpdateUserGroupRequest) (_result *UpdateUserGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserGroupResponse{}
	_body, _err := client.UpdateUserGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量修改登陆用户状态
//
// @param request - UpdateUsersStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUsersStatusResponse
func (client *Client) UpdateUsersStatusWithOptions(request *UpdateUsersStatusRequest, runtime *util.RuntimeOptions) (_result *UpdateUsersStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SaseUserIds)) {
		query["SaseUserIds"] = request.SaseUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUsersStatus"),
		Version:     tea.String("2023-01-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUsersStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量修改登陆用户状态
//
// @param request - UpdateUsersStatusRequest
//
// @return UpdateUsersStatusResponse
func (client *Client) UpdateUsersStatus(request *UpdateUsersStatusRequest) (_result *UpdateUsersStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUsersStatusResponse{}
	_body, _err := client.UpdateUsersStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
