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

type AttachApplication2ConnectorRequest struct {
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// ConnectorID。
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
	ApplicationIdsShrink *string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty"`
	// ConnectorID。
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AttachApplication2ConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateDynamicRouteRequest struct {
	ApplicationIds   []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	ApplicationType  *string   `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	Description      *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DynamicRouteType *string   `json:"DynamicRouteType,omitempty" xml:"DynamicRouteType,omitempty"`
	Name             *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NextHop          *string   `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	Priority         *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RegionIds        []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	Status           *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds           []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
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
	DynamicRouteId *string `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreatePrivateAccessApplicationRequest struct {
	Addresses   []*string                                          `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	Description *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges  []*CreatePrivateAccessApplicationRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Protocol    *string                                            `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Status      *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds      []*string                                          `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
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

func (s *CreatePrivateAccessApplicationRequest) SetDescription(v string) *CreatePrivateAccessApplicationRequest {
	s.Description = &v
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
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End   *int32 `json:"End,omitempty" xml:"End,omitempty"`
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
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApplicationIds       []*string                                               `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	ApplicationType      *string                                                 `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	CustomUserAttributes []*CreatePrivateAccessPolicyRequestCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	Description          *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string                                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyAction         *string                                                 `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	Priority             *int32                                                  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status               *string                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 内网访问标签ID集合。最多可输入100个内网访问标签ID。当**ApplicationType**为**Tag时**，必填。和**ApplicationIds**互斥。
	TagIds       []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// 内网访问策略的用户组类型。取值：
	// - **Normal**：普通用户组。
	// - **Custom**：自定义用户组。
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
	IdpId         *int32  `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	Relation      *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
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
	PolicyId  *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagId     *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePrivateAccessTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CompanyLimitCount  *CreateRegistrationPolicyRequestCompanyLimitCount  `json:"CompanyLimitCount,omitempty" xml:"CompanyLimitCount,omitempty" type:"Struct"`
	CompanyLimitType   *string                                            `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	Description        *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	MatchMode          *string                                            `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	Name               *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PersonalLimitCount *CreateRegistrationPolicyRequestPersonalLimitCount `json:"PersonalLimitCount,omitempty" xml:"PersonalLimitCount,omitempty" type:"Struct"`
	PersonalLimitType  *string                                            `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	Priority           *int64                                             `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status             *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds       []*string                                          `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist          []*string                                          `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
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
	All    *int32 `json:"All,omitempty" xml:"All,omitempty"`
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PC     *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
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
	All    *int32 `json:"All,omitempty" xml:"All,omitempty"`
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PC     *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
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
	CompanyLimitCountShrink  *string   `json:"CompanyLimitCount,omitempty" xml:"CompanyLimitCount,omitempty"`
	CompanyLimitType         *string   `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	Description              *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	MatchMode                *string   `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	Name                     *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PersonalLimitCountShrink *string   `json:"PersonalLimitCount,omitempty" xml:"PersonalLimitCount,omitempty"`
	PersonalLimitType        *string   `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	Priority                 *int64    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status                   *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds             []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist                []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
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
	Policy    *CreateRegistrationPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime   *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail  []*CreateRegistrationPolicyResponseBodyPolicyLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	MatchMode    *string                                                  `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	Name         *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyId     *string                                                  `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Priority     *string                                                  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status       *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string                                                `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string                                                `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
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
	DeviceBelong *string                                                          `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *CreateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	LimitType    *string                                                          `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
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
	All    *int32 `json:"All,omitempty" xml:"All,omitempty"`
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PC     *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRegistrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Attributes  []*CreateUserGroupRequestAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Description *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string                             `json:"Name,omitempty" xml:"Name,omitempty"`
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
	IdpId         *int32  `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	Relation      *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteDynamicRouteRequest struct {
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeletePrivateAccessApplicationRequest struct {
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePrivateAccessTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRegistrationPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteUserGroupRequest struct {
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	// ConnectorID。
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
	ApplicationIdsShrink *string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty"`
	// ConnectorID。
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DetachApplication2ConnectorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetDynamicRouteRequest struct {
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
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ApplicationIds   []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	ApplicationType  *string   `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	CreateTime       *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description      *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DynamicRouteId   *string   `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	DynamicRouteType *string   `json:"DynamicRouteType,omitempty" xml:"DynamicRouteType,omitempty"`
	Name             *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NextHop          *string   `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	Priority         *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RegionIds        []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	Status           *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds           []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetPrivateAccessApplicationRequest struct {
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
	RequestId   *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Addresses     []*string                                                       `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	ApplicationId *string                                                         `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ConnectorIds  []*string                                                       `json:"ConnectorIds,omitempty" xml:"ConnectorIds,omitempty" type:"Repeated"`
	CreateTime    *string                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyIds     []*string                                                       `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	PortRanges    []*GetPrivateAccessApplicationResponseBodyApplicationPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Protocol      *string                                                         `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Status        *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds        []*string                                                       `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
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
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End   *int32 `json:"End,omitempty" xml:"End,omitempty"`
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Policy    *GetPrivateAccessPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ApplicationIds       []*string                                                       `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	ApplicationType      *string                                                         `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	CreateTime           *string                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomUserAttributes []*GetPrivateAccessPolicyResponseBodyPolicyCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	Description          *string                                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string                                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyAction         *string                                                         `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	PolicyId             *string                                                         `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Priority             *int32                                                          `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status               *string                                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds               []*string                                                       `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	UserGroupIds         []*string                                                       `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	UserGroupMode        *string                                                         `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
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
	IdpId         *int32  `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	Relation      *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CreateTime   *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail  []*GetRegistrationPolicyResponseBodyLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	MatchMode    *string                                         `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	Name         *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyId     *string                                         `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Priority     *int64                                          `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string                                       `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string                                       `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
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
	DeviceBelong *string                                                 `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *GetRegistrationPolicyResponseBodyLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	LimitType    *string                                                 `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
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
	All    *int32 `json:"All,omitempty" xml:"All,omitempty"`
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PC     *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRegistrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Device    *GetUserDeviceResponseBodyDevice `json:"Device,omitempty" xml:"Device,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AppStatus     *string                                        `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	AppVersion    *string                                        `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	CPU           *string                                        `json:"CPU,omitempty" xml:"CPU,omitempty"`
	CreateTime    *string                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department    *string                                        `json:"Department,omitempty" xml:"Department,omitempty"`
	DeviceBelong  *string                                        `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	DeviceModel   *string                                        `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceStatus  *string                                        `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	DeviceTag     *string                                        `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	DeviceType    *string                                        `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	DeviceVersion *string                                        `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	Disk          *string                                        `json:"Disk,omitempty" xml:"Disk,omitempty"`
	DlpStatus     *string                                        `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	HistoryUsers  []*GetUserDeviceResponseBodyDeviceHistoryUsers `json:"HistoryUsers,omitempty" xml:"HistoryUsers,omitempty" type:"Repeated"`
	Hostname      *string                                        `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	IaStatus      *string                                        `json:"IaStatus,omitempty" xml:"IaStatus,omitempty"`
	InnerIP       *string                                        `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	Mac           *string                                        `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Memory        *string                                        `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NacStatus     *string                                        `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	PaStatus      *string                                        `json:"PaStatus,omitempty" xml:"PaStatus,omitempty"`
	SaseUserId    *string                                        `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	SharingStatus *bool                                          `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	SrcIP         *string                                        `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	UpdateTime    *string                                        `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Username      *string                                        `json:"Username,omitempty" xml:"Username,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Attributes  []*GetUserGroupResponseBodyUserGroupAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	CreateTime  *string                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	UserGroupId *string                                        `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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
	IdpId         *int32  `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	Relation      *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListApplicationsForPrivateAccessPolicyRequest struct {
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
	Polices   []*ListApplicationsForPrivateAccessPolicyResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	RequestId *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	PolicyId     *string                                                                  `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
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
	Addresses     []*string                                                                          `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	ApplicationId *string                                                                            `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	CreateTime    *string                                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string                                                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string                                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges    []*ListApplicationsForPrivateAccessPolicyResponseBodyPolicesApplicationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Protocol      *string                                                                            `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Status        *string                                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End   *int32 `json:"End,omitempty" xml:"End,omitempty"`
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
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListApplicationsForPrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	TagId        *string                                                            `json:"TagId,omitempty" xml:"TagId,omitempty"`
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
	Addresses     []*string `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	ApplicationId *string   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 内网访问应用创建时间。
	CreateTime  *string                                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string                                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges  []*ListApplicationsForPrivateAccessTagResponseBodyTagsApplicationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Protocol    *string                                                                      `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Status      *string                                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End   *int32 `json:"End,omitempty" xml:"End,omitempty"`
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
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListApplicationsForPrivateAccessTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListConnectorsRequest struct {
	ConnectorIds []*string `json:"ConnectorIds,omitempty" xml:"ConnectorIds,omitempty" type:"Repeated"`
	CurrentPage  *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PageSize     *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status       *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	SwitchStatus *string   `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
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
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum   *int32                                  `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
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
	ConnectorId  *string                                          `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	CreateTime   *string                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Name         *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	RegionId     *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status       *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
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
	ApplicationId   *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
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
	End   *string `json:"End,omitempty" xml:"End,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Regions   []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum  *int32    `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDynamicRouteRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApplicationId   *string   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	CurrentPage     *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DynamicRouteIds []*string `json:"DynamicRouteIds,omitempty" xml:"DynamicRouteIds,omitempty" type:"Repeated"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NextHop         *string   `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	PageSize        *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionIds       []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	Status          *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagId           *string   `json:"TagId,omitempty" xml:"TagId,omitempty"`
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
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum      *int32                                        `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
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
	ApplicationIds   []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	ApplicationType  *string   `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	CreateTime       *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description      *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DynamicRouteId   *string   `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	DynamicRouteType *string   `json:"DynamicRouteType,omitempty" xml:"DynamicRouteType,omitempty"`
	Name             *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NextHop          *string   `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	Priority         *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RegionIds        []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	Status           *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds           []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDynamicRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CurrentPage    *int64    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Department     *string   `json:"Department,omitempty" xml:"Department,omitempty"`
	DeviceTag      *string   `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	Hostname       *string   `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	Mac            *string   `json:"Mac,omitempty" xml:"Mac,omitempty"`
	PageSize       *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SaseUserId     *string   `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Statuses       []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	Username       *string   `json:"Username,omitempty" xml:"Username,omitempty"`
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
	RequestId    *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum     *int64                                                                 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
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
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department    *string `json:"Department,omitempty" xml:"Department,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceTag     *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	Hostname      *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	IsUsed        *bool   `json:"IsUsed,omitempty" xml:"IsUsed,omitempty"`
	Mac           *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	SaseUserId    *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Username      *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	Headers    map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListExcessiveDeviceRegistrationApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListPolicesForPrivateAccessApplicationRequest struct {
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
	RequestId    *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ApplicationType      *string                                                                                       `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	CreateTime           *string                                                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomUserAttributes []*ListPolicesForPrivateAccessApplicationResponseBodyApplicationsPoliciesCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	Description          *string                                                                                       `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string                                                                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyAction         *string                                                                                       `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	PolicyId             *string                                                                                       `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Priority             *int32                                                                                        `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status               *string                                                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupType        *string                                                                                       `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
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
	IdpId         *int32  `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	Relation      *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
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
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPolicesForPrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	TagId   *string                                                  `json:"TagId,omitempty" xml:"TagId,omitempty"`
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
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// 内网访问策略创建时间。
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 自定义用户组属性集合。多个自定义用户组属性之间是或的关系，按照合集生效。
	CustomUserAttributes []*ListPolicesForPrivateAccessTagResponseBodyTagsPolicesCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	Description          *string                                                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string                                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyAction         *string                                                                      `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	PolicyId             *string                                                                      `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Priority             *int32                                                                       `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status               *string                                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupType        *string                                                                      `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
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
	IdpId *int32 `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	// 用户组的关系。取值：
	// - **Equal**：等于。
	// - **Unequal**：不等于。
	Relation *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	// 用户组的类型。取值：
	// - **username**：用户名。
	// - **department**：部门。
	// - **email**：邮箱。
	// - **telephone**：手机。
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	// 用户组属性的值。
	// - 当用户组类型为**username**时，表示用户名的值。长度为1~128个字符，支持中文和大小写英文字母，可包含数字、半角句号（.）、下划线（_）和短划线（-）。
	// - 当用户组类型为**department**时，表示部门的值。如：OU=部门1,OU=SASE钉钉。
	// - 当用户组类型为**email**时，表示邮箱的值。如：username@example.com。
	// - 当用户组类型为**telephone**时，表示手机的值。如：13900001234。
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPolicesForPrivateAccessTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Polices     []*ListPolicesForUserGroupResponseBodyUserGroupsPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	UserGroupId *string                                                 `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyId   *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPolicesForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListPrivateAccessApplicationsRequest struct {
	Address        *string   `json:"Address,omitempty" xml:"Address,omitempty"`
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	ConnectorId    *string   `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	CurrentPage    *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PageSize       *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PolicyId       *string   `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Status         *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagId          *string   `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s ListPrivateAccessApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPrivateAccessApplicationsRequest) GoString() string {
	return s.String()
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
	RequestId    *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum     *int32                                                   `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
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
	Addresses     []*string                                                          `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	ApplicationId *string                                                            `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ConnectorIds  []*string                                                          `json:"ConnectorIds,omitempty" xml:"ConnectorIds,omitempty" type:"Repeated"`
	CreateTime    *string                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string                                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyIds     []*string                                                          `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	PortRanges    []*ListPrivateAccessApplicationsResponseBodyApplicationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Protocol      *string                                                            `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Status        *string                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds        []*string                                                          `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
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
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End   *int32 `json:"End,omitempty" xml:"End,omitempty"`
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPrivateAccessApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId     *string                                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Applications   []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	DynamicRouteId *string                                                                              `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
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
	Addresses     []*string                                                                                      `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	ApplicationId *string                                                                                        `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	CreateTime    *string                                                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description   *string                                                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Name          *string                                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRanges    []*ListPrivateAccessApplicationsForDynamicRouteResponseBodyDynamicRoutesApplicationsPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Protocol      *string                                                                                        `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Status        *string                                                                                        `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End   *int32 `json:"End,omitempty" xml:"End,omitempty"`
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
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPrivateAccessApplicationsForDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApplicationId *string   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	CurrentPage   *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Name          *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PageSize      *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PolicyAction  *string   `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	PolicyIds     []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	Status        *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagId         *string   `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// 用户组ID。取值来源：
	// - [ListUserGroups](~~ListUserGroups~~)：批量查询用户组。
	// - [CreateUserGroup](~~CreateUserGroup~~)：创建用户组。
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

func (s *ListPrivateAccessPolicesRequest) SetUserGroupId(v string) *ListPrivateAccessPolicesRequest {
	s.UserGroupId = &v
	return s
}

type ListPrivateAccessPolicesResponseBody struct {
	Polices   []*ListPrivateAccessPolicesResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum  *int32                                         `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
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
	ApplicationIds       []*string                                                          `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	ApplicationType      *string                                                            `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	CreateTime           *string                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomUserAttributes []*ListPrivateAccessPolicesResponseBodyPolicesCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	Description          *string                                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                 *string                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyAction         *string                                                            `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	PolicyId             *string                                                            `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Priority             *int32                                                             `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status               *string                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds               []*string                                                          `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	UserGroupIds         []*string                                                          `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	UserGroupMode        *string                                                            `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
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
	IdpId         *int32  `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	Relation      *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPrivateAccessPolicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApplicationId *string   `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	CurrentPage   *int32    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Name          *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PageSize      *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PolicyId      *string   `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	SimpleMode    *bool     `json:"SimpleMode,omitempty" xml:"SimpleMode,omitempty"`
	TagIds        []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
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
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tags      []*ListPrivateAccessTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TotalNum  *int32                                   `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
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
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	CreateTime     *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description    *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name           *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyIds      []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	TagId          *string   `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagType        *string   `json:"TagType,omitempty" xml:"TagType,omitempty"`
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPrivateAccessTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId     *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TagId       *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagType     *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
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
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPrivateAccessTagsForDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CompanyLimitType  *string   `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	CurrentPage       *int64    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	MatchMode         *string   `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	Name              *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PageSize          *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PersonalLimitType *string   `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	PolicyIds         []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
	Status            *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupId       *string   `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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
	Policies  []*ListRegistrationPoliciesResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum  *string                                         `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
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
	CreateTime   *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail  []*ListRegistrationPoliciesResponseBodyPoliciesLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	MatchMode    *string                                                    `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	Name         *string                                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyId     *string                                                    `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Priority     *int64                                                     `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status       *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string                                                  `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string                                                  `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
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
	DeviceBelong *string                                                            `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *ListRegistrationPoliciesResponseBodyPoliciesLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	LimitType    *string                                                            `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
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
	All    *int32 `json:"All,omitempty" xml:"All,omitempty"`
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PC     *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRegistrationPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Policies    []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	UserGroupId *string                                                               `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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
	CreateTime  *string                                                                          `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail []*ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	MatchMode   *string                                                                          `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	Name        *string                                                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyId    *string                                                                          `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Priority    *int64                                                                           `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status      *string                                                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Whitelist   []*string                                                                        `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
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
	DeviceBelong *string                                                                                  `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *ListRegistrationPoliciesForUserGroupResponseBodyUserGroupsPoliciesLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	LimitType    *string                                                                                  `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
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
	All    *string `json:"All,omitempty" xml:"All,omitempty"`
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PC     *string `json:"PC,omitempty" xml:"PC,omitempty"`
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
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRegistrationPoliciesForUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CurrentPage *int64  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	DeviceTag   *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Software  []*ListSoftwareForUserDeviceResponseBodySoftware `json:"Software,omitempty" xml:"Software,omitempty" type:"Repeated"`
	TotalNum  *int64                                           `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
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
	Inc         *string   `json:"Inc,omitempty" xml:"Inc,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSoftwareForUserDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	RequestId    *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TagId       *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagType     *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
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
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagsForPrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Polices   []*ListTagsForPrivateAccessPolicyResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TagId       *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TagType     *string `json:"TagType,omitempty" xml:"TagType,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagsForPrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListUserDevicesRequest struct {
	AppStatuses    []*string `json:"AppStatuses,omitempty" xml:"AppStatuses,omitempty" type:"Repeated"`
	CurrentPage    *int64    `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Department     *string   `json:"Department,omitempty" xml:"Department,omitempty"`
	DeviceBelong   *string   `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	DeviceStatuses []*string `json:"DeviceStatuses,omitempty" xml:"DeviceStatuses,omitempty" type:"Repeated"`
	DeviceTags     []*string `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty" type:"Repeated"`
	DeviceTypes    []*string `json:"DeviceTypes,omitempty" xml:"DeviceTypes,omitempty" type:"Repeated"`
	DlpStatuses    []*string `json:"DlpStatuses,omitempty" xml:"DlpStatuses,omitempty" type:"Repeated"`
	Hostname       *string   `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	IaStatuses     []*string `json:"IaStatuses,omitempty" xml:"IaStatuses,omitempty" type:"Repeated"`
	Mac            *string   `json:"Mac,omitempty" xml:"Mac,omitempty"`
	NacStatuses    []*string `json:"NacStatuses,omitempty" xml:"NacStatuses,omitempty" type:"Repeated"`
	PaStatuses     []*string `json:"PaStatuses,omitempty" xml:"PaStatuses,omitempty" type:"Repeated"`
	PageSize       *int64    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SaseUserId     *string   `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	SharingStatus  *bool     `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	SortBy         *string   `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Username       *string   `json:"Username,omitempty" xml:"Username,omitempty"`
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
	Devices   []*ListUserDevicesResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum  *int64                                `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
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
	AppStatus     *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	AppVersion    *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	CPU           *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department    *string `json:"Department,omitempty" xml:"Department,omitempty"`
	DeviceBelong  *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceStatus  *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	DeviceTag     *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	DeviceVersion *string `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	Disk          *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	DlpStatus     *string `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	Hostname      *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	IaStatus      *string `json:"IaStatus,omitempty" xml:"IaStatus,omitempty"`
	InnerIP       *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	Mac           *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Memory        *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NacStatus     *string `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	PaStatus      *string `json:"PaStatus,omitempty" xml:"PaStatus,omitempty"`
	SaseUserId    *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	SharingStatus *bool   `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	SrcIP         *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Username      *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
	CurrentPage    *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// 用户组名称。长度为1~128个字符，支持中文和大小写英文字母，可包含数字、半角句号（.）、下划线（_）和短划线（-）。
	Name         *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PAPolicyId   *string   `json:"PAPolicyId,omitempty" xml:"PAPolicyId,omitempty"`
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
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Attributes  []*ListUserGroupsResponseBodyUserGroupsAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	CreateTime  *string                                           `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string                                           `json:"Name,omitempty" xml:"Name,omitempty"`
	UserGroupId *string                                           `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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
	IdpId         *int32  `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	Relation      *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Polices   []*ListUserGroupsForPrivateAccessPolicyResponseBodyPolices `json:"Polices,omitempty" xml:"Polices,omitempty" type:"Repeated"`
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	IdpId         *int32  `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	Relation      *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
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
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserGroupsForPrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Policies  []*ListUserGroupsForRegistrationPolicyResponseBodyPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Attributes  []*ListUserGroupsForRegistrationPolicyResponseBodyPoliciesUserGroupsAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	CreateTime  *string                                                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                                                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string                                                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	UserGroupId *string                                                                        `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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
	IdpId         *int32  `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	Relation      *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
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
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserGroupsForRegistrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateDynamicRouteRequest struct {
	ApplicationIds   []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	ApplicationType  *string   `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	Description      *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	DynamicRouteId   *string   `json:"DynamicRouteId,omitempty" xml:"DynamicRouteId,omitempty"`
	DynamicRouteType *string   `json:"DynamicRouteType,omitempty" xml:"DynamicRouteType,omitempty"`
	ModifyType       *string   `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	Name             *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NextHop          *string   `json:"NextHop,omitempty" xml:"NextHop,omitempty"`
	Priority         *int32    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	RegionIds        []*string `json:"RegionIds,omitempty" xml:"RegionIds,omitempty" type:"Repeated"`
	Status           *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds           []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDynamicRouteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApplicationIds []*string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	Status         *string   `json:"Status,omitempty" xml:"Status,omitempty"`
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
	RequestId    *string                                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department    *string `json:"Department,omitempty" xml:"Department,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeviceTag     *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	Hostname      *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	IsUsed        *bool   `json:"IsUsed,omitempty" xml:"IsUsed,omitempty"`
	Mac           *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	SaseUserId    *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Username      *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	Headers    map[string]*string                                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateExcessiveDeviceRegistrationApplicationsStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdatePrivateAccessApplicationRequest struct {
	Addresses     []*string                                          `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	ApplicationId *string                                            `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Description   *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifyType    *string                                            `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	PortRanges    []*UpdatePrivateAccessApplicationRequestPortRanges `json:"PortRanges,omitempty" xml:"PortRanges,omitempty" type:"Repeated"`
	Protocol      *string                                            `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Status        *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIds        []*string                                          `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
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
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	End   *int32 `json:"End,omitempty" xml:"End,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePrivateAccessApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ApplicationIds       []*string                                               `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty" type:"Repeated"`
	ApplicationType      *string                                                 `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	CustomUserAttributes []*UpdatePrivateAccessPolicyRequestCustomUserAttributes `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty" type:"Repeated"`
	Description          *string                                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifyType           *string                                                 `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	PolicyAction         *string                                                 `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	PolicyId             *string                                                 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Priority             *int32                                                  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status               *string                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 内网访问标签ID集合。一条策略最多支持100个内网访问标签ID。
	TagIds       []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// 内网访问策略的用户组类型。取值：
	// - **Normal**：普通用户组。
	// - **Custom**：自定义用户组。
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
	IdpId         *int32  `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	Relation      *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CompanyLimitCount  *UpdateRegistrationPolicyRequestCompanyLimitCount  `json:"CompanyLimitCount,omitempty" xml:"CompanyLimitCount,omitempty" type:"Struct"`
	CompanyLimitType   *string                                            `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	Description        *string                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	MatchMode          *string                                            `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	Name               *string                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	PersonalLimitCount *UpdateRegistrationPolicyRequestPersonalLimitCount `json:"PersonalLimitCount,omitempty" xml:"PersonalLimitCount,omitempty" type:"Struct"`
	PersonalLimitType  *string                                            `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	PolicyId           *string                                            `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Priority           *int64                                             `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status             *string                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds       []*string                                          `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist          []*string                                          `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
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
	All    *int32 `json:"All,omitempty" xml:"All,omitempty"`
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PC     *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
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
	All    *int32 `json:"All,omitempty" xml:"All,omitempty"`
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PC     *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
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
	CompanyLimitCountShrink  *string   `json:"CompanyLimitCount,omitempty" xml:"CompanyLimitCount,omitempty"`
	CompanyLimitType         *string   `json:"CompanyLimitType,omitempty" xml:"CompanyLimitType,omitempty"`
	Description              *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	MatchMode                *string   `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	Name                     *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	PersonalLimitCountShrink *string   `json:"PersonalLimitCount,omitempty" xml:"PersonalLimitCount,omitempty"`
	PersonalLimitType        *string   `json:"PersonalLimitType,omitempty" xml:"PersonalLimitType,omitempty"`
	PolicyId                 *string   `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Priority                 *int64    `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status                   *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds             []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist                []*string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
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
	Policy    *UpdateRegistrationPolicyResponseBodyPolicy `json:"Policy,omitempty" xml:"Policy,omitempty" type:"Struct"`
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	CreateTime   *string                                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	LimitDetail  []*UpdateRegistrationPolicyResponseBodyPolicyLimitDetail `json:"LimitDetail,omitempty" xml:"LimitDetail,omitempty" type:"Repeated"`
	MatchMode    *string                                                  `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	Name         *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyId     *string                                                  `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Priority     *string                                                  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status       *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserGroupIds []*string                                                `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	Whitelist    []*string                                                `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
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
	DeviceBelong *string                                                          `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	LimitCount   *UpdateRegistrationPolicyResponseBodyPolicyLimitDetailLimitCount `json:"LimitCount,omitempty" xml:"LimitCount,omitempty" type:"Struct"`
	LimitType    *string                                                          `json:"LimitType,omitempty" xml:"LimitType,omitempty"`
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
	All    *int32 `json:"All,omitempty" xml:"All,omitempty"`
	Mobile *int32 `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PC     *int32 `json:"PC,omitempty" xml:"PC,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRegistrationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DeviceTags    []*string `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty" type:"Repeated"`
	SharingStatus *bool     `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
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
	Devices   []*UpdateUserDevicesSharingStatusResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AppStatus     *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	AppVersion    *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	CPU           *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department    *string `json:"Department,omitempty" xml:"Department,omitempty"`
	DeviceBelong  *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceStatus  *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	DeviceTag     *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	DeviceVersion *string `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	Disk          *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	DlpStatus     *string `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	Hostname      *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	IaStatus      *string `json:"IaStatus,omitempty" xml:"IaStatus,omitempty"`
	InnerIP       *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	Mac           *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Memory        *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NacStatus     *string `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	PaStatus      *string `json:"PaStatus,omitempty" xml:"PaStatus,omitempty"`
	SaseUserId    *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	SharingStatus *bool   `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	SrcIP         *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Username      *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserDevicesSharingStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DeviceAction *string   `json:"DeviceAction,omitempty" xml:"DeviceAction,omitempty"`
	DeviceTags   []*string `json:"DeviceTags,omitempty" xml:"DeviceTags,omitempty" type:"Repeated"`
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
	Devices   []*UpdateUserDevicesStatusResponseBodyDevices `json:"Devices,omitempty" xml:"Devices,omitempty" type:"Repeated"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AppStatus     *string `json:"AppStatus,omitempty" xml:"AppStatus,omitempty"`
	AppVersion    *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	CPU           *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Department    *string `json:"Department,omitempty" xml:"Department,omitempty"`
	DeviceBelong  *string `json:"DeviceBelong,omitempty" xml:"DeviceBelong,omitempty"`
	DeviceModel   *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	DeviceStatus  *string `json:"DeviceStatus,omitempty" xml:"DeviceStatus,omitempty"`
	DeviceTag     *string `json:"DeviceTag,omitempty" xml:"DeviceTag,omitempty"`
	DeviceType    *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	DeviceVersion *string `json:"DeviceVersion,omitempty" xml:"DeviceVersion,omitempty"`
	Disk          *string `json:"Disk,omitempty" xml:"Disk,omitempty"`
	DlpStatus     *string `json:"DlpStatus,omitempty" xml:"DlpStatus,omitempty"`
	Hostname      *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	IaStatus      *string `json:"IaStatus,omitempty" xml:"IaStatus,omitempty"`
	InnerIP       *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	Mac           *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	Memory        *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	NacStatus     *string `json:"NacStatus,omitempty" xml:"NacStatus,omitempty"`
	PaStatus      *string `json:"PaStatus,omitempty" xml:"PaStatus,omitempty"`
	SaseUserId    *string `json:"SaseUserId,omitempty" xml:"SaseUserId,omitempty"`
	SharingStatus *bool   `json:"SharingStatus,omitempty" xml:"SharingStatus,omitempty"`
	SrcIP         *string `json:"SrcIP,omitempty" xml:"SrcIP,omitempty"`
	UpdateTime    *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Username      *string `json:"Username,omitempty" xml:"Username,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserDevicesStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Attributes  []*UpdateUserGroupRequestAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	Description *string                             `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifyType  *string                             `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	UserGroupId *string                             `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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
	IdpId         *int32  `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	Relation      *string `json:"Relation,omitempty" xml:"Relation,omitempty"`
	UserGroupType *string `json:"UserGroupType,omitempty" xml:"UserGroupType,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
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
