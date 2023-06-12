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

type CreatePrivateAccessApplicationShrinkRequest struct {
	AddressesShrink  *string `json:"Addresses,omitempty" xml:"Addresses,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PortRangesShrink *string `json:"PortRanges,omitempty" xml:"PortRanges,omitempty"`
	Protocol         *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIdsShrink     *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
}

func (s CreatePrivateAccessApplicationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePrivateAccessApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetAddressesShrink(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.AddressesShrink = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetDescription(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetName(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetPortRangesShrink(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.PortRangesShrink = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetProtocol(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetStatus(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreatePrivateAccessApplicationShrinkRequest) SetTagIdsShrink(v string) *CreatePrivateAccessApplicationShrinkRequest {
	s.TagIdsShrink = &v
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

type CreatePrivateAccessPolicyShrinkRequest struct {
	ApplicationIdsShrink       *string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty"`
	ApplicationType            *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	CustomUserAttributesShrink *string `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty"`
	Description                *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name                       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PolicyAction               *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	Priority                   *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status                     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 内网访问标签ID集合。最多可输入100个内网访问标签ID。当**ApplicationType**为**Tag时**，必填。和**ApplicationIds**互斥。
	TagIdsShrink       *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
	UserGroupIdsShrink *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	// 内网访问策略的用户组类型。取值：
	// - **Normal**：普通用户组。
	// - **Custom**：自定义用户组。
	UserGroupMode *string `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
}

func (s CreatePrivateAccessPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePrivateAccessPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessPolicyShrinkRequest) SetApplicationIdsShrink(v string) *CreatePrivateAccessPolicyShrinkRequest {
	s.ApplicationIdsShrink = &v
	return s
}

func (s *CreatePrivateAccessPolicyShrinkRequest) SetApplicationType(v string) *CreatePrivateAccessPolicyShrinkRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreatePrivateAccessPolicyShrinkRequest) SetCustomUserAttributesShrink(v string) *CreatePrivateAccessPolicyShrinkRequest {
	s.CustomUserAttributesShrink = &v
	return s
}

func (s *CreatePrivateAccessPolicyShrinkRequest) SetDescription(v string) *CreatePrivateAccessPolicyShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePrivateAccessPolicyShrinkRequest) SetName(v string) *CreatePrivateAccessPolicyShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreatePrivateAccessPolicyShrinkRequest) SetPolicyAction(v string) *CreatePrivateAccessPolicyShrinkRequest {
	s.PolicyAction = &v
	return s
}

func (s *CreatePrivateAccessPolicyShrinkRequest) SetPriority(v int32) *CreatePrivateAccessPolicyShrinkRequest {
	s.Priority = &v
	return s
}

func (s *CreatePrivateAccessPolicyShrinkRequest) SetStatus(v string) *CreatePrivateAccessPolicyShrinkRequest {
	s.Status = &v
	return s
}

func (s *CreatePrivateAccessPolicyShrinkRequest) SetTagIdsShrink(v string) *CreatePrivateAccessPolicyShrinkRequest {
	s.TagIdsShrink = &v
	return s
}

func (s *CreatePrivateAccessPolicyShrinkRequest) SetUserGroupIdsShrink(v string) *CreatePrivateAccessPolicyShrinkRequest {
	s.UserGroupIdsShrink = &v
	return s
}

func (s *CreatePrivateAccessPolicyShrinkRequest) SetUserGroupMode(v string) *CreatePrivateAccessPolicyShrinkRequest {
	s.UserGroupMode = &v
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
	Applications []*ListConnectorsResponseBodyConnectorsApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
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

type UpdatePrivateAccessApplicationShrinkRequest struct {
	AddressesShrink  *string `json:"Addresses,omitempty" xml:"Addresses,omitempty"`
	ApplicationId    *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifyType       *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	PortRangesShrink *string `json:"PortRanges,omitempty" xml:"PortRanges,omitempty"`
	Protocol         *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TagIdsShrink     *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
}

func (s UpdatePrivateAccessApplicationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateAccessApplicationShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetAddressesShrink(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.AddressesShrink = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetApplicationId(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetDescription(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetModifyType(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.ModifyType = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetPortRangesShrink(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.PortRangesShrink = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetProtocol(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetStatus(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdatePrivateAccessApplicationShrinkRequest) SetTagIdsShrink(v string) *UpdatePrivateAccessApplicationShrinkRequest {
	s.TagIdsShrink = &v
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

type UpdatePrivateAccessPolicyShrinkRequest struct {
	ApplicationIdsShrink       *string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty"`
	ApplicationType            *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	CustomUserAttributesShrink *string `json:"CustomUserAttributes,omitempty" xml:"CustomUserAttributes,omitempty"`
	Description                *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ModifyType                 *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	PolicyAction               *string `json:"PolicyAction,omitempty" xml:"PolicyAction,omitempty"`
	PolicyId                   *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	Priority                   *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Status                     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 内网访问标签ID集合。一条策略最多支持100个内网访问标签ID。
	TagIdsShrink       *string `json:"TagIds,omitempty" xml:"TagIds,omitempty"`
	UserGroupIdsShrink *string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty"`
	// 内网访问策略的用户组类型。取值：
	// - **Normal**：普通用户组。
	// - **Custom**：自定义用户组。
	UserGroupMode *string `json:"UserGroupMode,omitempty" xml:"UserGroupMode,omitempty"`
}

func (s UpdatePrivateAccessPolicyShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePrivateAccessPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrivateAccessPolicyShrinkRequest) SetApplicationIdsShrink(v string) *UpdatePrivateAccessPolicyShrinkRequest {
	s.ApplicationIdsShrink = &v
	return s
}

func (s *UpdatePrivateAccessPolicyShrinkRequest) SetApplicationType(v string) *UpdatePrivateAccessPolicyShrinkRequest {
	s.ApplicationType = &v
	return s
}

func (s *UpdatePrivateAccessPolicyShrinkRequest) SetCustomUserAttributesShrink(v string) *UpdatePrivateAccessPolicyShrinkRequest {
	s.CustomUserAttributesShrink = &v
	return s
}

func (s *UpdatePrivateAccessPolicyShrinkRequest) SetDescription(v string) *UpdatePrivateAccessPolicyShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdatePrivateAccessPolicyShrinkRequest) SetModifyType(v string) *UpdatePrivateAccessPolicyShrinkRequest {
	s.ModifyType = &v
	return s
}

func (s *UpdatePrivateAccessPolicyShrinkRequest) SetPolicyAction(v string) *UpdatePrivateAccessPolicyShrinkRequest {
	s.PolicyAction = &v
	return s
}

func (s *UpdatePrivateAccessPolicyShrinkRequest) SetPolicyId(v string) *UpdatePrivateAccessPolicyShrinkRequest {
	s.PolicyId = &v
	return s
}

func (s *UpdatePrivateAccessPolicyShrinkRequest) SetPriority(v int32) *UpdatePrivateAccessPolicyShrinkRequest {
	s.Priority = &v
	return s
}

func (s *UpdatePrivateAccessPolicyShrinkRequest) SetStatus(v string) *UpdatePrivateAccessPolicyShrinkRequest {
	s.Status = &v
	return s
}

func (s *UpdatePrivateAccessPolicyShrinkRequest) SetTagIdsShrink(v string) *UpdatePrivateAccessPolicyShrinkRequest {
	s.TagIdsShrink = &v
	return s
}

func (s *UpdatePrivateAccessPolicyShrinkRequest) SetUserGroupIdsShrink(v string) *UpdatePrivateAccessPolicyShrinkRequest {
	s.UserGroupIdsShrink = &v
	return s
}

func (s *UpdatePrivateAccessPolicyShrinkRequest) SetUserGroupMode(v string) *UpdatePrivateAccessPolicyShrinkRequest {
	s.UserGroupMode = &v
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

func (client *Client) CreatePrivateAccessApplicationWithOptions(tmpReq *CreatePrivateAccessApplicationRequest, runtime *util.RuntimeOptions) (_result *CreatePrivateAccessApplicationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePrivateAccessApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Addresses)) {
		request.AddressesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Addresses, tea.String("Addresses"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PortRanges)) {
		request.PortRangesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PortRanges, tea.String("PortRanges"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagIds)) {
		request.TagIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIds, tea.String("TagIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressesShrink)) {
		body["Addresses"] = request.AddressesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PortRangesShrink)) {
		body["PortRanges"] = request.PortRangesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagIdsShrink)) {
		body["TagIds"] = request.TagIdsShrink
	}

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

func (client *Client) CreatePrivateAccessPolicyWithOptions(tmpReq *CreatePrivateAccessPolicyRequest, runtime *util.RuntimeOptions) (_result *CreatePrivateAccessPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePrivateAccessPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ApplicationIds)) {
		request.ApplicationIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationIds, tea.String("ApplicationIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.CustomUserAttributes)) {
		request.CustomUserAttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomUserAttributes, tea.String("CustomUserAttributes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagIds)) {
		request.TagIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIds, tea.String("TagIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserGroupIds)) {
		request.UserGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserGroupIds, tea.String("UserGroupIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationIdsShrink)) {
		body["ApplicationIds"] = request.ApplicationIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationType)) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !tea.BoolValue(util.IsUnset(request.CustomUserAttributesShrink)) {
		body["CustomUserAttributes"] = request.CustomUserAttributesShrink
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

	if !tea.BoolValue(util.IsUnset(request.TagIdsShrink)) {
		body["TagIds"] = request.TagIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupIdsShrink)) {
		body["UserGroupIds"] = request.UserGroupIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupMode)) {
		body["UserGroupMode"] = request.UserGroupMode
	}

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

func (client *Client) UpdatePrivateAccessApplicationWithOptions(tmpReq *UpdatePrivateAccessApplicationRequest, runtime *util.RuntimeOptions) (_result *UpdatePrivateAccessApplicationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdatePrivateAccessApplicationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Addresses)) {
		request.AddressesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Addresses, tea.String("Addresses"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PortRanges)) {
		request.PortRangesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PortRanges, tea.String("PortRanges"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagIds)) {
		request.TagIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIds, tea.String("TagIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressesShrink)) {
		body["Addresses"] = request.AddressesShrink
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

	if !tea.BoolValue(util.IsUnset(request.PortRangesShrink)) {
		body["PortRanges"] = request.PortRangesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Protocol)) {
		body["Protocol"] = request.Protocol
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagIdsShrink)) {
		body["TagIds"] = request.TagIdsShrink
	}

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

func (client *Client) UpdatePrivateAccessPolicyWithOptions(tmpReq *UpdatePrivateAccessPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdatePrivateAccessPolicyResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UpdatePrivateAccessPolicyShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ApplicationIds)) {
		request.ApplicationIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApplicationIds, tea.String("ApplicationIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.CustomUserAttributes)) {
		request.CustomUserAttributesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CustomUserAttributes, tea.String("CustomUserAttributes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagIds)) {
		request.TagIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagIds, tea.String("TagIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.UserGroupIds)) {
		request.UserGroupIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.UserGroupIds, tea.String("UserGroupIds"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ApplicationIdsShrink)) {
		body["ApplicationIds"] = request.ApplicationIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ApplicationType)) {
		body["ApplicationType"] = request.ApplicationType
	}

	if !tea.BoolValue(util.IsUnset(request.CustomUserAttributesShrink)) {
		body["CustomUserAttributes"] = request.CustomUserAttributesShrink
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

	if !tea.BoolValue(util.IsUnset(request.TagIdsShrink)) {
		body["TagIds"] = request.TagIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupIdsShrink)) {
		body["UserGroupIds"] = request.UserGroupIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserGroupMode)) {
		body["UserGroupMode"] = request.UserGroupMode
	}

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
