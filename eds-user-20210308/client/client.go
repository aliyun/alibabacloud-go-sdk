// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GroupResources struct {
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GroupResources) String() string {
	return tea.Prettify(s)
}

func (s GroupResources) GoString() string {
	return s.String()
}

func (s *GroupResources) SetRegion(v string) *GroupResources {
	s.Region = &v
	return s
}

func (s *GroupResources) SetResourceId(v string) *GroupResources {
	s.ResourceId = &v
	return s
}

func (s *GroupResources) SetResourceType(v string) *GroupResources {
	s.ResourceType = &v
	return s
}

type ThirdApp struct {
	AppKey        *string                `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Name          *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	OidcSsoConfig *ThirdAppOidcSsoConfig `json:"OidcSsoConfig,omitempty" xml:"OidcSsoConfig,omitempty" type:"Struct"`
	Secrets       []*ThirdAppSecrets     `json:"Secrets,omitempty" xml:"Secrets,omitempty" type:"Repeated"`
}

func (s ThirdApp) String() string {
	return tea.Prettify(s)
}

func (s ThirdApp) GoString() string {
	return s.String()
}

func (s *ThirdApp) SetAppKey(v string) *ThirdApp {
	s.AppKey = &v
	return s
}

func (s *ThirdApp) SetName(v string) *ThirdApp {
	s.Name = &v
	return s
}

func (s *ThirdApp) SetOidcSsoConfig(v *ThirdAppOidcSsoConfig) *ThirdApp {
	s.OidcSsoConfig = v
	return s
}

func (s *ThirdApp) SetSecrets(v []*ThirdAppSecrets) *ThirdApp {
	s.Secrets = v
	return s
}

type ThirdAppOidcSsoConfig struct {
	AccessTokenEffectiveTime *int32                          `json:"AccessTokenEffectiveTime,omitempty" xml:"AccessTokenEffectiveTime,omitempty"`
	CodeEffectiveTime        *int32                          `json:"CodeEffectiveTime,omitempty" xml:"CodeEffectiveTime,omitempty"`
	EnableAuthLogin          *bool                           `json:"EnableAuthLogin,omitempty" xml:"EnableAuthLogin,omitempty"`
	Endpoints                *ThirdAppOidcSsoConfigEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	GrantScopes              []*string                       `json:"GrantScopes,omitempty" xml:"GrantScopes,omitempty" type:"Repeated"`
	GrantTypes               []*string                       `json:"GrantTypes,omitempty" xml:"GrantTypes,omitempty" type:"Repeated"`
	IdTokenAlgorithmType     *string                         `json:"IdTokenAlgorithmType,omitempty" xml:"IdTokenAlgorithmType,omitempty"`
	IdTokenEffectiveTime     *int32                          `json:"IdTokenEffectiveTime,omitempty" xml:"IdTokenEffectiveTime,omitempty"`
	RedirectUris             []*string                       `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Repeated"`
	RefreshTokenEffective    *int32                          `json:"RefreshTokenEffective,omitempty" xml:"RefreshTokenEffective,omitempty"`
}

func (s ThirdAppOidcSsoConfig) String() string {
	return tea.Prettify(s)
}

func (s ThirdAppOidcSsoConfig) GoString() string {
	return s.String()
}

func (s *ThirdAppOidcSsoConfig) SetAccessTokenEffectiveTime(v int32) *ThirdAppOidcSsoConfig {
	s.AccessTokenEffectiveTime = &v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetCodeEffectiveTime(v int32) *ThirdAppOidcSsoConfig {
	s.CodeEffectiveTime = &v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetEnableAuthLogin(v bool) *ThirdAppOidcSsoConfig {
	s.EnableAuthLogin = &v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetEndpoints(v *ThirdAppOidcSsoConfigEndpoints) *ThirdAppOidcSsoConfig {
	s.Endpoints = v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetGrantScopes(v []*string) *ThirdAppOidcSsoConfig {
	s.GrantScopes = v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetGrantTypes(v []*string) *ThirdAppOidcSsoConfig {
	s.GrantTypes = v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetIdTokenAlgorithmType(v string) *ThirdAppOidcSsoConfig {
	s.IdTokenAlgorithmType = &v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetIdTokenEffectiveTime(v int32) *ThirdAppOidcSsoConfig {
	s.IdTokenEffectiveTime = &v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetRedirectUris(v []*string) *ThirdAppOidcSsoConfig {
	s.RedirectUris = v
	return s
}

func (s *ThirdAppOidcSsoConfig) SetRefreshTokenEffective(v int32) *ThirdAppOidcSsoConfig {
	s.RefreshTokenEffective = &v
	return s
}

type ThirdAppOidcSsoConfigEndpoints struct {
	AuthorizationEndpoint      *string `json:"AuthorizationEndpoint,omitempty" xml:"AuthorizationEndpoint,omitempty"`
	DiscoveryEndpoint          *string `json:"DiscoveryEndpoint,omitempty" xml:"DiscoveryEndpoint,omitempty"`
	GuestAuthorizationEndpoint *string `json:"GuestAuthorizationEndpoint,omitempty" xml:"GuestAuthorizationEndpoint,omitempty"`
	Issuer                     *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	JwksEndpoint               *string `json:"JwksEndpoint,omitempty" xml:"JwksEndpoint,omitempty"`
	LogoutEndpoint             *string `json:"LogoutEndpoint,omitempty" xml:"LogoutEndpoint,omitempty"`
	RevokeEndpoint             *string `json:"RevokeEndpoint,omitempty" xml:"RevokeEndpoint,omitempty"`
	TokenEndpoint              *string `json:"TokenEndpoint,omitempty" xml:"TokenEndpoint,omitempty"`
	UserinfoEndpoint           *string `json:"UserinfoEndpoint,omitempty" xml:"UserinfoEndpoint,omitempty"`
}

func (s ThirdAppOidcSsoConfigEndpoints) String() string {
	return tea.Prettify(s)
}

func (s ThirdAppOidcSsoConfigEndpoints) GoString() string {
	return s.String()
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetAuthorizationEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.AuthorizationEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetDiscoveryEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.DiscoveryEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetGuestAuthorizationEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.GuestAuthorizationEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetIssuer(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.Issuer = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetJwksEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.JwksEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetLogoutEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.LogoutEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetRevokeEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.RevokeEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetTokenEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.TokenEndpoint = &v
	return s
}

func (s *ThirdAppOidcSsoConfigEndpoints) SetUserinfoEndpoint(v string) *ThirdAppOidcSsoConfigEndpoints {
	s.UserinfoEndpoint = &v
	return s
}

type ThirdAppSecrets struct {
	Enable *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
}

func (s ThirdAppSecrets) String() string {
	return tea.Prettify(s)
}

func (s ThirdAppSecrets) GoString() string {
	return s.String()
}

func (s *ThirdAppSecrets) SetEnable(v bool) *ThirdAppSecrets {
	s.Enable = &v
	return s
}

func (s *ThirdAppSecrets) SetSecret(v string) *ThirdAppSecrets {
	s.Secret = &v
	return s
}

type WaIdPermissions struct {
	Code           *string            `json:"Code,omitempty" xml:"Code,omitempty"`
	IsBasicChild   *bool              `json:"IsBasicChild,omitempty" xml:"IsBasicChild,omitempty"`
	Name           *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	SubPermissions []*WaIdPermissions `json:"SubPermissions,omitempty" xml:"SubPermissions,omitempty" type:"Repeated"`
	Type           *string            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s WaIdPermissions) String() string {
	return tea.Prettify(s)
}

func (s WaIdPermissions) GoString() string {
	return s.String()
}

func (s *WaIdPermissions) SetCode(v string) *WaIdPermissions {
	s.Code = &v
	return s
}

func (s *WaIdPermissions) SetIsBasicChild(v bool) *WaIdPermissions {
	s.IsBasicChild = &v
	return s
}

func (s *WaIdPermissions) SetName(v string) *WaIdPermissions {
	s.Name = &v
	return s
}

func (s *WaIdPermissions) SetSubPermissions(v []*WaIdPermissions) *WaIdPermissions {
	s.SubPermissions = v
	return s
}

func (s *WaIdPermissions) SetType(v string) *WaIdPermissions {
	s.Type = &v
	return s
}

type BatchSetDesktopManagerRequest struct {
	// example:
	//
	// 1
	IsDesktopManager *string `json:"IsDesktopManager,omitempty" xml:"IsDesktopManager,omitempty"`
	// This parameter is required.
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s BatchSetDesktopManagerRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchSetDesktopManagerRequest) GoString() string {
	return s.String()
}

func (s *BatchSetDesktopManagerRequest) SetIsDesktopManager(v string) *BatchSetDesktopManagerRequest {
	s.IsDesktopManager = &v
	return s
}

func (s *BatchSetDesktopManagerRequest) SetUsers(v []*string) *BatchSetDesktopManagerRequest {
	s.Users = v
	return s
}

type BatchSetDesktopManagerResponseBody struct {
	// example:
	//
	// 868B8926-2E7A-5BE7-9897-E811E994****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchSetDesktopManagerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchSetDesktopManagerResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSetDesktopManagerResponseBody) SetRequestId(v string) *BatchSetDesktopManagerResponseBody {
	s.RequestId = &v
	return s
}

type BatchSetDesktopManagerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchSetDesktopManagerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchSetDesktopManagerResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchSetDesktopManagerResponse) GoString() string {
	return s.String()
}

func (s *BatchSetDesktopManagerResponse) SetHeaders(v map[string]*string) *BatchSetDesktopManagerResponse {
	s.Headers = v
	return s
}

func (s *BatchSetDesktopManagerResponse) SetStatusCode(v int32) *BatchSetDesktopManagerResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchSetDesktopManagerResponse) SetBody(v *BatchSetDesktopManagerResponseBody) *BatchSetDesktopManagerResponse {
	s.Body = v
	return s
}

type ChangeUserPasswordRequest struct {
	// example:
	//
	// alice***
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// Admin@12***
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty"`
}

func (s ChangeUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ChangeUserPasswordRequest) SetEndUserId(v string) *ChangeUserPasswordRequest {
	s.EndUserId = &v
	return s
}

func (s *ChangeUserPasswordRequest) SetNewPassword(v string) *ChangeUserPasswordRequest {
	s.NewPassword = &v
	return s
}

type ChangeUserPasswordResponseBody struct {
	// example:
	//
	// AA8D67CB-345D-5CDA-986E-FFAC7D0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeUserPasswordResponseBody) SetRequestId(v string) *ChangeUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ChangeUserPasswordResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ChangeUserPasswordResponse) SetHeaders(v map[string]*string) *ChangeUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ChangeUserPasswordResponse) SetStatusCode(v int32) *ChangeUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeUserPasswordResponse) SetBody(v *ChangeUserPasswordResponseBody) *ChangeUserPasswordResponse {
	s.Body = v
	return s
}

type CheckUsedPropertyRequest struct {
	// The ID of the property. You can call the [ListProperty](https://help.aliyun.com/document_detail/410890.html) operation to query the property ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
}

func (s CheckUsedPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUsedPropertyRequest) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyRequest) SetPropertyId(v int64) *CheckUsedPropertyRequest {
	s.PropertyId = &v
	return s
}

type CheckUsedPropertyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 22C97624-2405-54AC-BD44-A63FBE97CC2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of convenience users that are associated with the property.
	//
	// example:
	//
	// 7
	UseCount *int64 `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
}

func (s CheckUsedPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUsedPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyResponseBody) SetRequestId(v string) *CheckUsedPropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUsedPropertyResponseBody) SetUseCount(v int64) *CheckUsedPropertyResponseBody {
	s.UseCount = &v
	return s
}

type CheckUsedPropertyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUsedPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUsedPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUsedPropertyResponse) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyResponse) SetHeaders(v map[string]*string) *CheckUsedPropertyResponse {
	s.Headers = v
	return s
}

func (s *CheckUsedPropertyResponse) SetStatusCode(v int32) *CheckUsedPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUsedPropertyResponse) SetBody(v *CheckUsedPropertyResponseBody) *CheckUsedPropertyResponse {
	s.Body = v
	return s
}

type CheckUsedPropertyValueRequest struct {
	// The property ID. You can call the [ListProperty](~~ListProperty~~) operation to query property ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 380
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The ID of the property value. You can call the [ListProperty](~~ListProperty~~) operation to query the ID of the property value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s CheckUsedPropertyValueRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckUsedPropertyValueRequest) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyValueRequest) SetPropertyId(v int64) *CheckUsedPropertyValueRequest {
	s.PropertyId = &v
	return s
}

func (s *CheckUsedPropertyValueRequest) SetPropertyValueId(v int64) *CheckUsedPropertyValueRequest {
	s.PropertyValueId = &v
	return s
}

type CheckUsedPropertyValueResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of convenience accounts that are associated with the specified custom property value.
	//
	// example:
	//
	// 1
	UseCount *int64 `json:"UseCount,omitempty" xml:"UseCount,omitempty"`
}

func (s CheckUsedPropertyValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckUsedPropertyValueResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyValueResponseBody) SetRequestId(v string) *CheckUsedPropertyValueResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUsedPropertyValueResponseBody) SetUseCount(v int64) *CheckUsedPropertyValueResponseBody {
	s.UseCount = &v
	return s
}

type CheckUsedPropertyValueResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckUsedPropertyValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckUsedPropertyValueResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckUsedPropertyValueResponse) GoString() string {
	return s.String()
}

func (s *CheckUsedPropertyValueResponse) SetHeaders(v map[string]*string) *CheckUsedPropertyValueResponse {
	s.Headers = v
	return s
}

func (s *CheckUsedPropertyValueResponse) SetStatusCode(v int32) *CheckUsedPropertyValueResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckUsedPropertyValueResponse) SetBody(v *CheckUsedPropertyValueResponseBody) *CheckUsedPropertyValueResponse {
	s.Body = v
	return s
}

type CreateGroupRequest struct {
	// example:
	//
	// ENTERPRISE
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// root
	ParentGroupId *string `json:"ParentGroupId,omitempty" xml:"ParentGroupId,omitempty"`
	// example:
	//
	// co-0esnf80jab***
	SolutionId *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetBizType(v string) *CreateGroupRequest {
	s.BizType = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupRequest) SetParentGroupId(v string) *CreateGroupRequest {
	s.ParentGroupId = &v
	return s
}

func (s *CreateGroupRequest) SetSolutionId(v string) *CreateGroupRequest {
	s.SolutionId = &v
	return s
}

type CreateGroupResponseBody struct {
	// example:
	//
	// ug-12341234****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// AA8D67CB-345D-5CDA-986E-FFAC7D0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseBody) SetGroupId(v string) *CreateGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateGroupResponseBody) SetRequestId(v string) *CreateGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetStatusCode(v int32) *CreateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupResponse) SetBody(v *CreateGroupResponseBody) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateOrgRequest struct {
	// This parameter is required.
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// org-evk12ozjvmlxl****
	ParentOrgId *string `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
}

func (s CreateOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrgRequest) GoString() string {
	return s.String()
}

func (s *CreateOrgRequest) SetOrgName(v string) *CreateOrgRequest {
	s.OrgName = &v
	return s
}

func (s *CreateOrgRequest) SetParentOrgId(v string) *CreateOrgRequest {
	s.ParentOrgId = &v
	return s
}

type CreateOrgResponseBody struct {
	// example:
	//
	// org-evk12ozjvmlxl****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// example:
	//
	// 7A2C3803-C975-5871-A232-80A91009****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrgResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrgResponseBody) SetOrgId(v string) *CreateOrgResponseBody {
	s.OrgId = &v
	return s
}

func (s *CreateOrgResponseBody) SetRequestId(v string) *CreateOrgResponseBody {
	s.RequestId = &v
	return s
}

type CreateOrgResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrgResponse) GoString() string {
	return s.String()
}

func (s *CreateOrgResponse) SetHeaders(v map[string]*string) *CreateOrgResponse {
	s.Headers = v
	return s
}

func (s *CreateOrgResponse) SetStatusCode(v int32) *CreateOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrgResponse) SetBody(v *CreateOrgResponseBody) *CreateOrgResponse {
	s.Body = v
	return s
}

type CreatePropertyRequest struct {
	// The property name.
	//
	// This parameter is required.
	//
	// example:
	//
	// department
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The values of the property. You can specify up to 50 values for a property.
	PropertyValues []*string `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
}

func (s CreatePropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyRequest) GoString() string {
	return s.String()
}

func (s *CreatePropertyRequest) SetPropertyKey(v string) *CreatePropertyRequest {
	s.PropertyKey = &v
	return s
}

func (s *CreatePropertyRequest) SetPropertyValues(v []*string) *CreatePropertyRequest {
	s.PropertyValues = v
	return s
}

type CreatePropertyResponseBody struct {
	// The result of creating the property.
	CreateResult *CreatePropertyResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBody) SetCreateResult(v *CreatePropertyResponseBodyCreateResult) *CreatePropertyResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreatePropertyResponseBody) SetRequestId(v string) *CreatePropertyResponseBody {
	s.RequestId = &v
	return s
}

type CreatePropertyResponseBodyCreateResult struct {
	// The ID of the property.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The name of the property.
	//
	// example:
	//
	// department
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The result of creating the property value.
	SavePropertyValueModel *CreatePropertyResponseBodyCreateResultSavePropertyValueModel `json:"SavePropertyValueModel,omitempty" xml:"SavePropertyValueModel,omitempty" type:"Struct"`
}

func (s CreatePropertyResponseBodyCreateResult) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBodyCreateResult) SetPropertyId(v int64) *CreatePropertyResponseBodyCreateResult {
	s.PropertyId = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResult) SetPropertyKey(v string) *CreatePropertyResponseBodyCreateResult {
	s.PropertyKey = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResult) SetSavePropertyValueModel(v *CreatePropertyResponseBodyCreateResultSavePropertyValueModel) *CreatePropertyResponseBodyCreateResult {
	s.SavePropertyValueModel = v
	return s
}

type CreatePropertyResponseBodyCreateResultSavePropertyValueModel struct {
	// The property values that failed to be created.
	FailedPropertyValues []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues `json:"FailedPropertyValues,omitempty" xml:"FailedPropertyValues,omitempty" type:"Repeated"`
	// Details of the property values that were created.
	SavePropertyValues []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues `json:"SavePropertyValues,omitempty" xml:"SavePropertyValues,omitempty" type:"Repeated"`
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModel) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModel) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModel) SetFailedPropertyValues(v []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) *CreatePropertyResponseBodyCreateResultSavePropertyValueModel {
	s.FailedPropertyValues = v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModel) SetSavePropertyValues(v []*CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) *CreatePropertyResponseBodyCreateResultSavePropertyValueModel {
	s.SavePropertyValues = v
	return s
}

type CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues struct {
	// The error code.
	//
	// example:
	//
	// ExistedPropertyValue
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The property value is used by another property.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the property value.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The value of the property.
	//
	// example:
	//
	// HR
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) SetErrorCode(v string) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues {
	s.ErrorCode = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) SetErrorMessage(v string) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues {
	s.ErrorMessage = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) SetPropertyId(v int64) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues {
	s.PropertyId = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues) SetPropertyValue(v string) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelFailedPropertyValues {
	s.PropertyValue = &v
	return s
}

type CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues struct {
	// The value of the property.
	//
	// example:
	//
	// HR
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The ID of the property value.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) SetPropertyValue(v string) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues) SetPropertyValueId(v int64) *CreatePropertyResponseBodyCreateResultSavePropertyValueModelSavePropertyValues {
	s.PropertyValueId = &v
	return s
}

type CreatePropertyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePropertyResponse) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponse) SetHeaders(v map[string]*string) *CreatePropertyResponse {
	s.Headers = v
	return s
}

func (s *CreatePropertyResponse) SetStatusCode(v int32) *CreatePropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePropertyResponse) SetBody(v *CreatePropertyResponseBody) *CreatePropertyResponse {
	s.Body = v
	return s
}

type CreateResourceGroupRequest struct {
	// example:
	//
	// 0
	IsResourceGroupWithOfficeSite *int64 `json:"IsResourceGroupWithOfficeSite,omitempty" xml:"IsResourceGroupWithOfficeSite,omitempty"`
	// example:
	//
	// AliyunConsole
	Platform          *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s CreateResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupRequest) SetIsResourceGroupWithOfficeSite(v int64) *CreateResourceGroupRequest {
	s.IsResourceGroupWithOfficeSite = &v
	return s
}

func (s *CreateResourceGroupRequest) SetPlatform(v string) *CreateResourceGroupRequest {
	s.Platform = &v
	return s
}

func (s *CreateResourceGroupRequest) SetResourceGroupName(v string) *CreateResourceGroupRequest {
	s.ResourceGroupName = &v
	return s
}

type CreateResourceGroupResponseBody struct {
	// example:
	//
	// 868B8926-2E7A-5BE7-9897-E811E994****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-ckf3cx7isinhk***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBody) SetRequestId(v string) *CreateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceGroupResponseBody) SetResourceGroupId(v string) *CreateResourceGroupResponseBody {
	s.ResourceGroupId = &v
	return s
}

type CreateResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponse) SetHeaders(v map[string]*string) *CreateResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceGroupResponse) SetStatusCode(v int32) *CreateResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceGroupResponse) SetBody(v *CreateResourceGroupResponseBody) *CreateResourceGroupResponse {
	s.Body = v
	return s
}

type CreateUsersRequest struct {
	// The date on which the convenience users are automatically locked.
	//
	// example:
	//
	// 2023-03-03
	AutoLockTime *string `json:"AutoLockTime,omitempty" xml:"AutoLockTime,omitempty"`
	IsLocalAdmin *bool   `json:"IsLocalAdmin,omitempty" xml:"IsLocalAdmin,omitempty"`
	// The initial password. If this parameter is left empty, an email for password reset is sent to the specified email address.
	//
	// example:
	//
	// Test123****
	Password           *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PasswordExpireDays *string `json:"PasswordExpireDays,omitempty" xml:"PasswordExpireDays,omitempty"`
	// The information about the convenience user.
	//
	// This parameter is required.
	//
	// example:
	//
	// CreateUsers
	Users []*CreateUsersRequestUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CreateUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersRequest) GoString() string {
	return s.String()
}

func (s *CreateUsersRequest) SetAutoLockTime(v string) *CreateUsersRequest {
	s.AutoLockTime = &v
	return s
}

func (s *CreateUsersRequest) SetIsLocalAdmin(v bool) *CreateUsersRequest {
	s.IsLocalAdmin = &v
	return s
}

func (s *CreateUsersRequest) SetPassword(v string) *CreateUsersRequest {
	s.Password = &v
	return s
}

func (s *CreateUsersRequest) SetPasswordExpireDays(v string) *CreateUsersRequest {
	s.PasswordExpireDays = &v
	return s
}

func (s *CreateUsersRequest) SetUsers(v []*CreateUsersRequestUsers) *CreateUsersRequest {
	s.Users = v
	return s
}

type CreateUsersRequestUsers struct {
	// The email address of the convenience user. The email address is used to receive notifications about events such as desktop assignment. You must specify an email address or a mobile number to receive notifications.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The username of the convenience user. The name can contain lowercase letters, digits, and underscores (_), and must be 3 to 24 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The organization to which the convenience user belongs.
	//
	// example:
	//
	// 1111****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The type of the account ownership.
	//
	// Valid values:
	//
	// 	- CreateFromManager: administrator-activated
	//
	// 	- Normal: user-activated
	//
	// example:
	//
	// Normal
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The user password.
	//
	// >  The password must be at least 10 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters (excluding spaces).
	//
	// example:
	//
	// password1
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Mobile numbers are not supported on the international site (alibabacloud.com).
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The display name of the end user.
	//
	// example:
	//
	// Bean
	RealNickName *string `json:"RealNickName,omitempty" xml:"RealNickName,omitempty"`
	// The remarks on the convenience user.
	//
	// example:
	//
	// remark1
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateUsersRequestUsers) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersRequestUsers) GoString() string {
	return s.String()
}

func (s *CreateUsersRequestUsers) SetEmail(v string) *CreateUsersRequestUsers {
	s.Email = &v
	return s
}

func (s *CreateUsersRequestUsers) SetEndUserId(v string) *CreateUsersRequestUsers {
	s.EndUserId = &v
	return s
}

func (s *CreateUsersRequestUsers) SetOrgId(v string) *CreateUsersRequestUsers {
	s.OrgId = &v
	return s
}

func (s *CreateUsersRequestUsers) SetOwnerType(v string) *CreateUsersRequestUsers {
	s.OwnerType = &v
	return s
}

func (s *CreateUsersRequestUsers) SetPassword(v string) *CreateUsersRequestUsers {
	s.Password = &v
	return s
}

func (s *CreateUsersRequestUsers) SetPhone(v string) *CreateUsersRequestUsers {
	s.Phone = &v
	return s
}

func (s *CreateUsersRequestUsers) SetRealNickName(v string) *CreateUsersRequestUsers {
	s.RealNickName = &v
	return s
}

func (s *CreateUsersRequestUsers) SetRemark(v string) *CreateUsersRequestUsers {
	s.Remark = &v
	return s
}

type CreateUsersResponseBody struct {
	AllSucceed *bool `json:"AllSucceed,omitempty" xml:"AllSucceed,omitempty"`
	// The result of user creation.
	CreateResult *CreateUsersResponseBodyCreateResult `json:"CreateResult,omitempty" xml:"CreateResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBody) SetAllSucceed(v bool) *CreateUsersResponseBody {
	s.AllSucceed = &v
	return s
}

func (s *CreateUsersResponseBody) SetCreateResult(v *CreateUsersResponseBodyCreateResult) *CreateUsersResponseBody {
	s.CreateResult = v
	return s
}

func (s *CreateUsersResponseBody) SetRequestId(v string) *CreateUsersResponseBody {
	s.RequestId = &v
	return s
}

type CreateUsersResponseBodyCreateResult struct {
	// Details of the created convenience users.
	CreatedUsers []*CreateUsersResponseBodyCreateResultCreatedUsers `json:"CreatedUsers,omitempty" xml:"CreatedUsers,omitempty" type:"Repeated"`
	// Details of the convenience users that failed to be created.
	FailedUsers []*CreateUsersResponseBodyCreateResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
}

func (s CreateUsersResponseBodyCreateResult) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersResponseBodyCreateResult) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBodyCreateResult) SetCreatedUsers(v []*CreateUsersResponseBodyCreateResultCreatedUsers) *CreateUsersResponseBodyCreateResult {
	s.CreatedUsers = v
	return s
}

func (s *CreateUsersResponseBodyCreateResult) SetFailedUsers(v []*CreateUsersResponseBodyCreateResultFailedUsers) *CreateUsersResponseBodyCreateResult {
	s.FailedUsers = v
	return s
}

type CreateUsersResponseBodyCreateResultCreatedUsers struct {
	// The email address of the end user.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the end user.
	//
	// example:
	//
	// test1
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The mobile number of the end user.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The display name of the end user.
	//
	// example:
	//
	// Bean
	RealNickName *string `json:"RealNickName,omitempty" xml:"RealNickName,omitempty"`
	// The remarks of the end user.
	//
	// example:
	//
	// remark1
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateUsersResponseBodyCreateResultCreatedUsers) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersResponseBodyCreateResultCreatedUsers) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetEmail(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.Email = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetEndUserId(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.EndUserId = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetPhone(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.Phone = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetRealNickName(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.RealNickName = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultCreatedUsers) SetRemark(v string) *CreateUsersResponseBodyCreateResultCreatedUsers {
	s.Remark = &v
	return s
}

type CreateUsersResponseBodyCreateResultFailedUsers struct {
	// The email address of the end user.
	//
	// example:
	//
	// username2@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the end user.
	//
	// example:
	//
	// test2
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// ExistedEndUserId
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The username test is used by another user.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The mobile number of the end user.
	//
	// example:
	//
	// 1390000****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s CreateUsersResponseBodyCreateResultFailedUsers) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersResponseBodyCreateResultFailedUsers) GoString() string {
	return s.String()
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetEmail(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.Email = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetEndUserId(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetErrorCode(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetErrorMessage(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

func (s *CreateUsersResponseBodyCreateResultFailedUsers) SetPhone(v string) *CreateUsersResponseBodyCreateResultFailedUsers {
	s.Phone = &v
	return s
}

type CreateUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUsersResponse) GoString() string {
	return s.String()
}

func (s *CreateUsersResponse) SetHeaders(v map[string]*string) *CreateUsersResponse {
	s.Headers = v
	return s
}

func (s *CreateUsersResponse) SetStatusCode(v int32) *CreateUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUsersResponse) SetBody(v *CreateUsersResponseBody) *CreateUsersResponse {
	s.Body = v
	return s
}

type DeleteResourceGroupRequest struct {
	// >  The ID of the resource group that you want to delete.
	//
	// 	- If you also specify ResourceGroupIds, both parameters take effect.
	//
	// example:
	//
	// rg-aj01tck67a0szp***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// >  The IDs of the resource groups that you want to delete.
	//
	// 	- If you also specify ResourceGroupId, both parameters take effect.
	ResourceGroupIds []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
}

func (s DeleteResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupRequest) SetResourceGroupId(v string) *DeleteResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteResourceGroupRequest) SetResourceGroupIds(v []*string) *DeleteResourceGroupRequest {
	s.ResourceGroupIds = v
	return s
}

type DeleteResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2463A343-BD32-5803-959E-9A8472A1***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBody) SetRequestId(v string) *DeleteResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponse) SetHeaders(v map[string]*string) *DeleteResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceGroupResponse) SetStatusCode(v int32) *DeleteResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceGroupResponse) SetBody(v *DeleteResourceGroupResponseBody) *DeleteResourceGroupResponse {
	s.Body = v
	return s
}

type DeleteUserPropertyValueRequest struct {
	// The property ID. You can call the [ListProperty](~~ListProperty~~) operation to query the property ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The property value ID. You can call the [ListProperty](~~ListProperty~~) operation to query the property value ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
	// The user ID. You can call the [DescribeUsers](~~DescribeUsers~~) operation to query the user ID, which is the return value of the `Id` parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserPropertyValueRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserPropertyValueRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserPropertyValueRequest) SetPropertyId(v int64) *DeleteUserPropertyValueRequest {
	s.PropertyId = &v
	return s
}

func (s *DeleteUserPropertyValueRequest) SetPropertyValueId(v int64) *DeleteUserPropertyValueRequest {
	s.PropertyValueId = &v
	return s
}

func (s *DeleteUserPropertyValueRequest) SetUserId(v int64) *DeleteUserPropertyValueRequest {
	s.UserId = &v
	return s
}

type DeleteUserPropertyValueResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D6C62E40-F937-5803-B008-92E813399BA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUserPropertyValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserPropertyValueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUserPropertyValueResponseBody) SetRequestId(v string) *DeleteUserPropertyValueResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUserPropertyValueResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserPropertyValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserPropertyValueResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserPropertyValueResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserPropertyValueResponse) SetHeaders(v map[string]*string) *DeleteUserPropertyValueResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserPropertyValueResponse) SetStatusCode(v int32) *DeleteUserPropertyValueResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserPropertyValueResponse) SetBody(v *DeleteUserPropertyValueResponseBody) *DeleteUserPropertyValueResponse {
	s.Body = v
	return s
}

type DescribeGroupUserRequest struct {
	// example:
	//
	// ENTERPRISE
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// ug-12341234****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// co-0esnf80jab***
	SolutionId *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s DescribeGroupUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupUserRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupUserRequest) SetBizType(v string) *DescribeGroupUserRequest {
	s.BizType = &v
	return s
}

func (s *DescribeGroupUserRequest) SetGroupId(v string) *DescribeGroupUserRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupUserRequest) SetSolutionId(v string) *DescribeGroupUserRequest {
	s.SolutionId = &v
	return s
}

type DescribeGroupUserResponseBody struct {
	Groups []*DescribeGroupUserResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// example:
	//
	// AA8D67CB-345D-5CDA-986E-FFAC7D0****
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*DescribeGroupUserResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribeGroupUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupUserResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupUserResponseBody) SetGroups(v []*DescribeGroupUserResponseBodyGroups) *DescribeGroupUserResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeGroupUserResponseBody) SetRequestId(v string) *DescribeGroupUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupUserResponseBody) SetUsers(v []*DescribeGroupUserResponseBodyUsers) *DescribeGroupUserResponseBody {
	s.Users = v
	return s
}

type DescribeGroupUserResponseBodyGroups struct {
	// example:
	//
	// ug-91mvbosdjsdfh****
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 1
	UserCount *string `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s DescribeGroupUserResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupUserResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeGroupUserResponseBodyGroups) SetGroupId(v string) *DescribeGroupUserResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupUserResponseBodyGroups) SetGroupName(v string) *DescribeGroupUserResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *DescribeGroupUserResponseBodyGroups) SetUserCount(v string) *DescribeGroupUserResponseBodyGroups {
	s.UserCount = &v
	return s
}

type DescribeGroupUserResponseBodyUsers struct {
	// example:
	//
	// xx-xx-xx
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// https://avatar.****.com
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// example:
	//
	// alex****@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// alex****
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// example:
	//
	// 2024-08-26T02:59:22.000+00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2024-08-26T02:59:22.000+00:00
	GmtJoinGroup *string `json:"GmtJoinGroup,omitempty" xml:"GmtJoinGroup,omitempty"`
	// example:
	//
	// 123
	JobNumber *string `json:"JobNumber,omitempty" xml:"JobNumber,omitempty"`
	// example:
	//
	// alex
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// example:
	//
	// 188888****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s DescribeGroupUserResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupUserResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeGroupUserResponseBodyUsers) SetAddress(v string) *DescribeGroupUserResponseBodyUsers {
	s.Address = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetAvatar(v string) *DescribeGroupUserResponseBodyUsers {
	s.Avatar = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetEmail(v string) *DescribeGroupUserResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetEndUserId(v string) *DescribeGroupUserResponseBodyUsers {
	s.EndUserId = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetGmtCreated(v string) *DescribeGroupUserResponseBodyUsers {
	s.GmtCreated = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetGmtJoinGroup(v string) *DescribeGroupUserResponseBodyUsers {
	s.GmtJoinGroup = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetJobNumber(v string) *DescribeGroupUserResponseBodyUsers {
	s.JobNumber = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetNickName(v string) *DescribeGroupUserResponseBodyUsers {
	s.NickName = &v
	return s
}

func (s *DescribeGroupUserResponseBodyUsers) SetPhone(v string) *DescribeGroupUserResponseBodyUsers {
	s.Phone = &v
	return s
}

type DescribeGroupUserResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupUserResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupUserResponse) SetHeaders(v map[string]*string) *DescribeGroupUserResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupUserResponse) SetStatusCode(v int32) *DescribeGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupUserResponse) SetBody(v *DescribeGroupUserResponseBody) *DescribeGroupUserResponse {
	s.Body = v
	return s
}

type DescribeGroupsRequest struct {
	// example:
	//
	// ENTERPRISE
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// ug-12341234****
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// co-0esnf80jab***
	SolutionId *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s DescribeGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupsRequest) SetBizType(v string) *DescribeGroupsRequest {
	s.BizType = &v
	return s
}

func (s *DescribeGroupsRequest) SetGroupId(v string) *DescribeGroupsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupsRequest) SetGroupName(v string) *DescribeGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeGroupsRequest) SetSolutionId(v string) *DescribeGroupsRequest {
	s.SolutionId = &v
	return s
}

type DescribeGroupsResponseBody struct {
	Groups []*DescribeGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBody) SetGroups(v []*DescribeGroupsResponseBodyGroups) *DescribeGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *DescribeGroupsResponseBody) SetRequestId(v string) *DescribeGroupsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeGroupsResponseBodyGroups struct {
	// example:
	//
	// ug-2412ojkwtybd****
	GroupId   *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DescribeGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponseBodyGroups) SetGroupId(v string) *DescribeGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupsResponseBodyGroups) SetGroupName(v string) *DescribeGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

type DescribeGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupsResponse) SetHeaders(v map[string]*string) *DescribeGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupsResponse) SetStatusCode(v int32) *DescribeGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupsResponse) SetBody(v *DescribeGroupsResponseBody) *DescribeGroupsResponse {
	s.Body = v
	return s
}

type DescribeMfaDevicesRequest struct {
	// The domain of the Active Directory (AD) workspace.
	//
	// example:
	//
	// cn.misumi.pri
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The usernames of the convenience users.
	//
	// example:
	//
	// test
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The maximum number of entries to return. Valid values: 1 to 500.\\
	//
	// Default value: 100.
	//
	// example:
	//
	// 100
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Set the value to the token that is obtained from the previous query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The serial numbers of the virtual MFA devices.
	//
	// example:
	//
	// c2d9ae94-a64b-4a0d-8024-9519ca50****
	SerialNumbers []*string `json:"SerialNumbers,omitempty" xml:"SerialNumbers,omitempty" type:"Repeated"`
}

func (s DescribeMfaDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMfaDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeMfaDevicesRequest) SetAdDomain(v string) *DescribeMfaDevicesRequest {
	s.AdDomain = &v
	return s
}

func (s *DescribeMfaDevicesRequest) SetEndUserIds(v []*string) *DescribeMfaDevicesRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribeMfaDevicesRequest) SetMaxResults(v int64) *DescribeMfaDevicesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeMfaDevicesRequest) SetNextToken(v string) *DescribeMfaDevicesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMfaDevicesRequest) SetSerialNumbers(v []*string) *DescribeMfaDevicesRequest {
	s.SerialNumbers = v
	return s
}

type DescribeMfaDevicesResponseBody struct {
	// The information about the virtual MFA devices.
	MfaDevices []*DescribeMfaDevicesResponseBodyMfaDevices `json:"MfaDevices,omitempty" xml:"MfaDevices,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6na6YlN9asMM31MsMcdQNpp
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 33DBB8EC-6E68-4726-91C4-E09C59D9A7D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMfaDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMfaDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMfaDevicesResponseBody) SetMfaDevices(v []*DescribeMfaDevicesResponseBodyMfaDevices) *DescribeMfaDevicesResponseBody {
	s.MfaDevices = v
	return s
}

func (s *DescribeMfaDevicesResponseBody) SetNextToken(v string) *DescribeMfaDevicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeMfaDevicesResponseBody) SetRequestId(v string) *DescribeMfaDevicesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeMfaDevicesResponseBodyMfaDevices struct {
	// The number of consecutive failures to bind the virtual MFA device, or the number of authentication failures based on the virtual MFA device.
	//
	// example:
	//
	// 0
	ConsecutiveFails *int32 `json:"ConsecutiveFails,omitempty" xml:"ConsecutiveFails,omitempty"`
	// The type of the virtual MFA device. The value can only be TOTP_VIRTUAL. This value indicates that the virtual MFA device follows the Time-based One-time Password (TOTP) algorithm.
	//
	// example:
	//
	// TOTP_VIRTUAL
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The username of the convenience user that uses the virtual MFA device.
	//
	// example:
	//
	// test
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The time when the virtual MFA device was enabled. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-22T06:20:49Z
	GmtEnabled *string `json:"GmtEnabled,omitempty" xml:"GmtEnabled,omitempty"`
	// The time when the locked virtual MFA device was automatically unlocked. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-23T06:20:49Z
	GmtUnlock *string `json:"GmtUnlock,omitempty" xml:"GmtUnlock,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 36
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The serial number of the virtual MFA device. The serial number is unique for each device.
	//
	// example:
	//
	// dc856334-446b-4035-bfbc-18af261e****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The status of the virtual MFA device.
	//
	// Valid values:
	//
	// 	- LOCKED
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- UNBOUND
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- NORMAL
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMfaDevicesResponseBodyMfaDevices) String() string {
	return tea.Prettify(s)
}

func (s DescribeMfaDevicesResponseBodyMfaDevices) GoString() string {
	return s.String()
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetConsecutiveFails(v int32) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.ConsecutiveFails = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetDeviceType(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.DeviceType = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetEmail(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.Email = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetEndUserId(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.EndUserId = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetGmtEnabled(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.GmtEnabled = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetGmtUnlock(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.GmtUnlock = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetId(v int64) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.Id = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetSerialNumber(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.SerialNumber = &v
	return s
}

func (s *DescribeMfaDevicesResponseBodyMfaDevices) SetStatus(v string) *DescribeMfaDevicesResponseBodyMfaDevices {
	s.Status = &v
	return s
}

type DescribeMfaDevicesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMfaDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMfaDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMfaDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMfaDevicesResponse) SetHeaders(v map[string]*string) *DescribeMfaDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMfaDevicesResponse) SetStatusCode(v int32) *DescribeMfaDevicesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMfaDevicesResponse) SetBody(v *DescribeMfaDevicesResponseBody) *DescribeMfaDevicesResponse {
	s.Body = v
	return s
}

type DescribeOrgByLayerRequest struct {
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// example:
	//
	// org-11fs****
	ParentOrgId *string `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
}

func (s DescribeOrgByLayerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgByLayerRequest) GoString() string {
	return s.String()
}

func (s *DescribeOrgByLayerRequest) SetOrgName(v string) *DescribeOrgByLayerRequest {
	s.OrgName = &v
	return s
}

func (s *DescribeOrgByLayerRequest) SetParentOrgId(v string) *DescribeOrgByLayerRequest {
	s.ParentOrgId = &v
	return s
}

type DescribeOrgByLayerResponseBody struct {
	Orgs []*DescribeOrgByLayerResponseBodyOrgs `json:"Orgs,omitempty" xml:"Orgs,omitempty" type:"Repeated"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOrgByLayerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgByLayerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOrgByLayerResponseBody) SetOrgs(v []*DescribeOrgByLayerResponseBodyOrgs) *DescribeOrgByLayerResponseBody {
	s.Orgs = v
	return s
}

func (s *DescribeOrgByLayerResponseBody) SetRequestId(v string) *DescribeOrgByLayerResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOrgByLayerResponseBodyOrgs struct {
	// example:
	//
	// org-1mox****
	OrgId   *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// example:
	//
	// org-ezqr****
	ParentOrgId *string `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
}

func (s DescribeOrgByLayerResponseBodyOrgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgByLayerResponseBodyOrgs) GoString() string {
	return s.String()
}

func (s *DescribeOrgByLayerResponseBodyOrgs) SetOrgId(v string) *DescribeOrgByLayerResponseBodyOrgs {
	s.OrgId = &v
	return s
}

func (s *DescribeOrgByLayerResponseBodyOrgs) SetOrgName(v string) *DescribeOrgByLayerResponseBodyOrgs {
	s.OrgName = &v
	return s
}

func (s *DescribeOrgByLayerResponseBodyOrgs) SetParentOrgId(v string) *DescribeOrgByLayerResponseBodyOrgs {
	s.ParentOrgId = &v
	return s
}

type DescribeOrgByLayerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOrgByLayerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOrgByLayerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgByLayerResponse) GoString() string {
	return s.String()
}

func (s *DescribeOrgByLayerResponse) SetHeaders(v map[string]*string) *DescribeOrgByLayerResponse {
	s.Headers = v
	return s
}

func (s *DescribeOrgByLayerResponse) SetStatusCode(v int32) *DescribeOrgByLayerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOrgByLayerResponse) SetBody(v *DescribeOrgByLayerResponseBody) *DescribeOrgByLayerResponse {
	s.Body = v
	return s
}

type DescribeOrgsRequest struct {
	// The maximum number of entries to return. Valid values: 1 to 100.\\
	//
	// Default value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the query. The return value is the value of the NextToken response parameter that was returned last time the DescribeOrgs operation was called.
	//
	// example:
	//
	// AAAAAV3MpHK****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The name of the organization.
	//
	// example:
	//
	// org****
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// The parent organization ID.
	//
	// example:
	//
	// org-****
	ParentOrgId *string `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
}

func (s DescribeOrgsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOrgsRequest) SetMaxResults(v int64) *DescribeOrgsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeOrgsRequest) SetNextToken(v string) *DescribeOrgsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeOrgsRequest) SetOrgName(v string) *DescribeOrgsRequest {
	s.OrgName = &v
	return s
}

func (s *DescribeOrgsRequest) SetParentOrgId(v string) *DescribeOrgsRequest {
	s.ParentOrgId = &v
	return s
}

type DescribeOrgsResponseBody struct {
	// The token that determines the start point of the query. The return value is the value of the NextToken response parameter that was returned last time the DescribeOrgs operation was called.
	//
	// example:
	//
	// AAAAAV3MpHK****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The organizations.
	Orgs []*DescribeOrgsResponseBodyOrgs `json:"Orgs,omitempty" xml:"Orgs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0B4BB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOrgsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOrgsResponseBody) SetNextToken(v string) *DescribeOrgsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeOrgsResponseBody) SetOrgs(v []*DescribeOrgsResponseBodyOrgs) *DescribeOrgsResponseBody {
	s.Orgs = v
	return s
}

func (s *DescribeOrgsResponseBody) SetRequestId(v string) *DescribeOrgsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOrgsResponseBodyOrgs struct {
	// The organization ID.
	//
	// example:
	//
	// org-****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The name of the organizational unit.
	//
	// example:
	//
	// org****
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// The parent organization ID.
	//
	// example:
	//
	// org-****
	ParentOrgId *string `json:"ParentOrgId,omitempty" xml:"ParentOrgId,omitempty"`
}

func (s DescribeOrgsResponseBodyOrgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgsResponseBodyOrgs) GoString() string {
	return s.String()
}

func (s *DescribeOrgsResponseBodyOrgs) SetOrgId(v string) *DescribeOrgsResponseBodyOrgs {
	s.OrgId = &v
	return s
}

func (s *DescribeOrgsResponseBodyOrgs) SetOrgName(v string) *DescribeOrgsResponseBodyOrgs {
	s.OrgName = &v
	return s
}

func (s *DescribeOrgsResponseBodyOrgs) SetParentOrgId(v string) *DescribeOrgsResponseBodyOrgs {
	s.ParentOrgId = &v
	return s
}

type DescribeOrgsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOrgsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOrgsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOrgsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOrgsResponse) SetHeaders(v map[string]*string) *DescribeOrgsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOrgsResponse) SetStatusCode(v int32) *DescribeOrgsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOrgsResponse) SetBody(v *DescribeOrgsResponseBody) *DescribeOrgsResponse {
	s.Body = v
	return s
}

type DescribeResourceGroupsRequest struct {
	// example:
	//
	// 0
	NeedContainResourceGroupWithOfficeSite *int64 `json:"NeedContainResourceGroupWithOfficeSite,omitempty" xml:"NeedContainResourceGroupWithOfficeSite,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// AliyunConsole
	Platform          *string   `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ResourceGroupIds  []*string `json:"ResourceGroupIds,omitempty" xml:"ResourceGroupIds,omitempty" type:"Repeated"`
	ResourceGroupName *string   `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s DescribeResourceGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsRequest) SetNeedContainResourceGroupWithOfficeSite(v int64) *DescribeResourceGroupsRequest {
	s.NeedContainResourceGroupWithOfficeSite = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetPageNumber(v int32) *DescribeResourceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetPageSize(v int32) *DescribeResourceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetPlatform(v string) *DescribeResourceGroupsRequest {
	s.Platform = &v
	return s
}

func (s *DescribeResourceGroupsRequest) SetResourceGroupIds(v []*string) *DescribeResourceGroupsRequest {
	s.ResourceGroupIds = v
	return s
}

func (s *DescribeResourceGroupsRequest) SetResourceGroupName(v string) *DescribeResourceGroupsRequest {
	s.ResourceGroupName = &v
	return s
}

type DescribeResourceGroupsResponseBody struct {
	// example:
	//
	// 68BA1DF7-8814-5AED-B844-F8F7F7****
	RequestId     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroup []*DescribeResourceGroupsResponseBodyResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Repeated"`
	// example:
	//
	// 7
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeResourceGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsResponseBody) SetRequestId(v string) *DescribeResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceGroupsResponseBody) SetResourceGroup(v []*DescribeResourceGroupsResponseBodyResourceGroup) *DescribeResourceGroupsResponseBody {
	s.ResourceGroup = v
	return s
}

func (s *DescribeResourceGroupsResponseBody) SetTotalCount(v string) *DescribeResourceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeResourceGroupsResponseBodyResourceGroup struct {
	AppRules []*DescribeResourceGroupsResponseBodyResourceGroupAppRules `json:"AppRules,omitempty" xml:"AppRules,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	AuthCount *string `json:"AuthCount,omitempty" xml:"AuthCount,omitempty"`
	// example:
	//
	// 2022-11-29T17:25:40.000000000Z
	CreateTime *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Policies   []*DescribeResourceGroupsResponseBodyResourceGroupPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// example:
	//
	// 119
	ResourceCount *string `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// example:
	//
	// rg-cyo0il2pzge1***
	ResourceGroupId   *string                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceGroupName *string                                                  `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	Timers            []*DescribeResourceGroupsResponseBodyResourceGroupTimers `json:"Timers,omitempty" xml:"Timers,omitempty" type:"Repeated"`
}

func (s DescribeResourceGroupsResponseBodyResourceGroup) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceGroupsResponseBodyResourceGroup) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetAppRules(v []*DescribeResourceGroupsResponseBodyResourceGroupAppRules) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.AppRules = v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetAuthCount(v string) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.AuthCount = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetCreateTime(v string) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetPolicies(v []*DescribeResourceGroupsResponseBodyResourceGroupPolicies) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.Policies = v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetResourceCount(v string) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.ResourceCount = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetResourceGroupId(v string) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetResourceGroupName(v string) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetTimers(v []*DescribeResourceGroupsResponseBodyResourceGroupTimers) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.Timers = v
	return s
}

type DescribeResourceGroupsResponseBodyResourceGroupAppRules struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeResourceGroupsResponseBodyResourceGroupAppRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceGroupsResponseBodyResourceGroupAppRules) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupAppRules) SetId(v string) *DescribeResourceGroupsResponseBodyResourceGroupAppRules {
	s.Id = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupAppRules) SetName(v string) *DescribeResourceGroupsResponseBodyResourceGroupAppRules {
	s.Name = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupAppRules) SetType(v int32) *DescribeResourceGroupsResponseBodyResourceGroupAppRules {
	s.Type = &v
	return s
}

type DescribeResourceGroupsResponseBodyResourceGroupPolicies struct {
	// example:
	//
	// pl-a8jnatl8kjasb***
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// False
	IsDefault *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeResourceGroupsResponseBodyResourceGroupPolicies) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceGroupsResponseBodyResourceGroupPolicies) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupPolicies) SetId(v string) *DescribeResourceGroupsResponseBodyResourceGroupPolicies {
	s.Id = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupPolicies) SetIsDefault(v bool) *DescribeResourceGroupsResponseBodyResourceGroupPolicies {
	s.IsDefault = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupPolicies) SetName(v string) *DescribeResourceGroupsResponseBodyResourceGroupPolicies {
	s.Name = &v
	return s
}

type DescribeResourceGroupsResponseBodyResourceGroupTimers struct {
	BindStatus *string `json:"BindStatus,omitempty" xml:"BindStatus,omitempty"`
	// example:
	//
	// t-asdzx0mbjhg***
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TimerStatus *string `json:"TimerStatus,omitempty" xml:"TimerStatus,omitempty"`
}

func (s DescribeResourceGroupsResponseBodyResourceGroupTimers) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceGroupsResponseBodyResourceGroupTimers) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) SetBindStatus(v string) *DescribeResourceGroupsResponseBodyResourceGroupTimers {
	s.BindStatus = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) SetId(v string) *DescribeResourceGroupsResponseBodyResourceGroupTimers {
	s.Id = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) SetName(v string) *DescribeResourceGroupsResponseBodyResourceGroupTimers {
	s.Name = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) SetTimerStatus(v string) *DescribeResourceGroupsResponseBodyResourceGroupTimers {
	s.TimerStatus = &v
	return s
}

type DescribeResourceGroupsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsResponse) SetHeaders(v map[string]*string) *DescribeResourceGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceGroupsResponse) SetStatusCode(v int32) *DescribeResourceGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceGroupsResponse) SetBody(v *DescribeResourceGroupsResponseBody) *DescribeResourceGroupsResponse {
	s.Body = v
	return s
}

type DescribeUsersRequest struct {
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The usernames that must be exactly matched.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The usernames that must be exactly excluded.
	ExcludeEndUserIds []*string `json:"ExcludeEndUserIds,omitempty" xml:"ExcludeEndUserIds,omitempty" type:"Repeated"`
	// The string that is used for fuzzy search. You perform fuzzy search by username (EndUserId) and email address (Email). Wildcard characters (\\*) are supported. For example, if you set this parameter to `a*m`, usernames or email addresses that start with `a` and end with `m` are returned.
	//
	// example:
	//
	// a*m
	Filter                      *string            `json:"Filter,omitempty" xml:"Filter,omitempty"`
	FilterWithAssignedResource  map[string]*string `json:"FilterWithAssignedResource,omitempty" xml:"FilterWithAssignedResource,omitempty"`
	FilterWithAssignedResources map[string]*bool   `json:"FilterWithAssignedResources,omitempty" xml:"FilterWithAssignedResources,omitempty"`
	// The ID of the organization in which you want to query convenience users.
	//
	// example:
	//
	// ug-12341234****
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	IsQueryAllSubOrgs *bool   `json:"IsQueryAllSubOrgs,omitempty" xml:"IsQueryAllSubOrgs,omitempty"`
	// The maximum number of entries per page.
	//
	// 	- Valid values: 1 to 500.
	//
	// 	- Default value: 500.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.\\
	//
	// If not all results are returned in a query, a value is returned for the NextToken parameter. In this case, you can use the return value of NextToken to perform the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the organization in which you want to query users.
	//
	// example:
	//
	// org-4mdgc1cocc59z****
	OrgId      *string                `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ShowExtras map[string]interface{} `json:"ShowExtras,omitempty" xml:"ShowExtras,omitempty"`
	SolutionId *string                `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
	Status     *int32                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersRequest) SetBizType(v string) *DescribeUsersRequest {
	s.BizType = &v
	return s
}

func (s *DescribeUsersRequest) SetEndUserIds(v []*string) *DescribeUsersRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribeUsersRequest) SetExcludeEndUserIds(v []*string) *DescribeUsersRequest {
	s.ExcludeEndUserIds = v
	return s
}

func (s *DescribeUsersRequest) SetFilter(v string) *DescribeUsersRequest {
	s.Filter = &v
	return s
}

func (s *DescribeUsersRequest) SetFilterWithAssignedResource(v map[string]*string) *DescribeUsersRequest {
	s.FilterWithAssignedResource = v
	return s
}

func (s *DescribeUsersRequest) SetFilterWithAssignedResources(v map[string]*bool) *DescribeUsersRequest {
	s.FilterWithAssignedResources = v
	return s
}

func (s *DescribeUsersRequest) SetGroupId(v string) *DescribeUsersRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeUsersRequest) SetIsQueryAllSubOrgs(v bool) *DescribeUsersRequest {
	s.IsQueryAllSubOrgs = &v
	return s
}

func (s *DescribeUsersRequest) SetMaxResults(v int64) *DescribeUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUsersRequest) SetNextToken(v string) *DescribeUsersRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersRequest) SetOrgId(v string) *DescribeUsersRequest {
	s.OrgId = &v
	return s
}

func (s *DescribeUsersRequest) SetShowExtras(v map[string]interface{}) *DescribeUsersRequest {
	s.ShowExtras = v
	return s
}

func (s *DescribeUsersRequest) SetSolutionId(v string) *DescribeUsersRequest {
	s.SolutionId = &v
	return s
}

func (s *DescribeUsersRequest) SetStatus(v int32) *DescribeUsersRequest {
	s.Status = &v
	return s
}

type DescribeUsersShrinkRequest struct {
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The usernames that must be exactly matched.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The usernames that must be exactly excluded.
	ExcludeEndUserIds []*string `json:"ExcludeEndUserIds,omitempty" xml:"ExcludeEndUserIds,omitempty" type:"Repeated"`
	// The string that is used for fuzzy search. You perform fuzzy search by username (EndUserId) and email address (Email). Wildcard characters (\\*) are supported. For example, if you set this parameter to `a*m`, usernames or email addresses that start with `a` and end with `m` are returned.
	//
	// example:
	//
	// a*m
	Filter                            *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	FilterWithAssignedResourceShrink  *string `json:"FilterWithAssignedResource,omitempty" xml:"FilterWithAssignedResource,omitempty"`
	FilterWithAssignedResourcesShrink *string `json:"FilterWithAssignedResources,omitempty" xml:"FilterWithAssignedResources,omitempty"`
	// The ID of the organization in which you want to query convenience users.
	//
	// example:
	//
	// ug-12341234****
	GroupId           *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	IsQueryAllSubOrgs *bool   `json:"IsQueryAllSubOrgs,omitempty" xml:"IsQueryAllSubOrgs,omitempty"`
	// The maximum number of entries per page.
	//
	// 	- Valid values: 1 to 500.
	//
	// 	- Default value: 500.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.\\
	//
	// If not all results are returned in a query, a value is returned for the NextToken parameter. In this case, you can use the return value of NextToken to perform the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the organization in which you want to query users.
	//
	// example:
	//
	// org-4mdgc1cocc59z****
	OrgId            *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	ShowExtrasShrink *string `json:"ShowExtras,omitempty" xml:"ShowExtras,omitempty"`
	SolutionId       *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUsersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsersShrinkRequest) SetBizType(v string) *DescribeUsersShrinkRequest {
	s.BizType = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetEndUserIds(v []*string) *DescribeUsersShrinkRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribeUsersShrinkRequest) SetExcludeEndUserIds(v []*string) *DescribeUsersShrinkRequest {
	s.ExcludeEndUserIds = v
	return s
}

func (s *DescribeUsersShrinkRequest) SetFilter(v string) *DescribeUsersShrinkRequest {
	s.Filter = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetFilterWithAssignedResourceShrink(v string) *DescribeUsersShrinkRequest {
	s.FilterWithAssignedResourceShrink = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetFilterWithAssignedResourcesShrink(v string) *DescribeUsersShrinkRequest {
	s.FilterWithAssignedResourcesShrink = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetGroupId(v string) *DescribeUsersShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetIsQueryAllSubOrgs(v bool) *DescribeUsersShrinkRequest {
	s.IsQueryAllSubOrgs = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetMaxResults(v int64) *DescribeUsersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetNextToken(v string) *DescribeUsersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetOrgId(v string) *DescribeUsersShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetShowExtrasShrink(v string) *DescribeUsersShrinkRequest {
	s.ShowExtrasShrink = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetSolutionId(v string) *DescribeUsersShrinkRequest {
	s.SolutionId = &v
	return s
}

func (s *DescribeUsersShrinkRequest) SetStatus(v int32) *DescribeUsersShrinkRequest {
	s.Status = &v
	return s
}

type DescribeUsersResponseBody struct {
	// The token that determines the start point of the next query. If this parameter is left empty, all results are returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The convenience accounts.
	Users []*DescribeUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribeUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBody) SetNextToken(v string) *DescribeUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersResponseBody) SetRequestId(v string) *DescribeUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsersResponseBody) SetUsers(v []*DescribeUsersResponseBodyUsers) *DescribeUsersResponseBody {
	s.Users = v
	return s
}

type DescribeUsersResponseBodyUsers struct {
	// The work address of the convenience user.
	//
	// example:
	//
	// Hangzhou \\*\\*\\*
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The profile picture of the convenience user.
	//
	// example:
	//
	// https://cdn.*****
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// The email address of the convenience user.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The username of the convenience user.
	//
	// example:
	//
	// alice
	EndUserId    *string                               `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	ExternalName *string                               `json:"ExternalName,omitempty" xml:"ExternalName,omitempty"`
	Extras       *DescribeUsersResponseBodyUsersExtras `json:"Extras,omitempty" xml:"Extras,omitempty" type:"Struct"`
	// The user groups to which the convenience user belongs.
	Groups []*DescribeUsersResponseBodyUsersGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The ID of the convenience user.
	//
	// example:
	//
	// 4205**
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the convenience user is an administrator. If the convenience user is of the administrator-activated type, you must specify a user administrator. Notifications such as password reset on a client are sent to the email address or mobile number of the user administrator. For more information, see [Create a convenience user](https://help.aliyun.com/document_detail/214472.html).
	//
	// example:
	//
	// true
	IsTenantManager *bool `json:"IsTenantManager,omitempty" xml:"IsTenantManager,omitempty"`
	// The employee number of the convenience user.
	//
	// example:
	//
	// A10000**
	JobNumber *string `json:"JobNumber,omitempty" xml:"JobNumber,omitempty"`
	// The nickname of the convenience user.
	//
	// example:
	//
	// Lee
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The ID of the organization to which the convenience user belongs.
	//
	// >  This parameter will be deprecated in the future.
	//
	// example:
	//
	// org-4mdgc1cocc59z****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The organizations to which the convenience user belongs.
	Orgs []*DescribeUsersResponseBodyUsersOrgs `json:"Orgs,omitempty" xml:"Orgs,omitempty" type:"Repeated"`
	// The type of the convenience account.
	//
	// 	- Administrator-activated type: The administrator specifies the username and password of the convenience account. User notifications such as password reset notifications are sent to the email address or mobile number of the administrator.
	//
	// 	- User-activated type: The administrator specifies the username and the email address or mobile number of a convenience user. Notifications such as activation notifications that contain the default password are sent to the email address or mobile number of the convenience user.
	//
	// Valid values:
	//
	// 	- CreateFromManager
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     administrator-activated
	//
	//     <!-- -->
	//
	// 	- Normal
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     user-activated
	//
	//     <!-- -->
	//
	// example:
	//
	// Normal
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The mobile number of the convenience user. If you leave this parameter empty, the value of this parameter is not returned.
	//
	// example:
	//
	// 1381111****
	Phone        *string                                     `json:"Phone,omitempty" xml:"Phone,omitempty"`
	Properties   []*DescribeUsersResponseBodyUsersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	RealNickName *string                                     `json:"RealNickName,omitempty" xml:"RealNickName,omitempty"`
	// The remarks on the convenience user.
	//
	// example:
	//
	// TestUser
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The status of the convenience user.
	//
	// Valid values:
	//
	// 	- 0: The convenience user is normal.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- 9: The convenience user is locked.
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The globally unique ID of the convenience user.
	//
	// example:
	//
	// 41fd1254d8f7****
	WyId *string `json:"WyId,omitempty" xml:"WyId,omitempty"`
}

func (s DescribeUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsers) SetAddress(v string) *DescribeUsersResponseBodyUsers {
	s.Address = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetAvatar(v string) *DescribeUsersResponseBodyUsers {
	s.Avatar = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetEmail(v string) *DescribeUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetEndUserId(v string) *DescribeUsersResponseBodyUsers {
	s.EndUserId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetExternalName(v string) *DescribeUsersResponseBodyUsers {
	s.ExternalName = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetExtras(v *DescribeUsersResponseBodyUsersExtras) *DescribeUsersResponseBodyUsers {
	s.Extras = v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetGroups(v []*DescribeUsersResponseBodyUsersGroups) *DescribeUsersResponseBodyUsers {
	s.Groups = v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetId(v int64) *DescribeUsersResponseBodyUsers {
	s.Id = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetIsTenantManager(v bool) *DescribeUsersResponseBodyUsers {
	s.IsTenantManager = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetJobNumber(v string) *DescribeUsersResponseBodyUsers {
	s.JobNumber = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetNickName(v string) *DescribeUsersResponseBodyUsers {
	s.NickName = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetOrgId(v string) *DescribeUsersResponseBodyUsers {
	s.OrgId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetOrgs(v []*DescribeUsersResponseBodyUsersOrgs) *DescribeUsersResponseBodyUsers {
	s.Orgs = v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetOwnerType(v string) *DescribeUsersResponseBodyUsers {
	s.OwnerType = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetPhone(v string) *DescribeUsersResponseBodyUsers {
	s.Phone = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetProperties(v []*DescribeUsersResponseBodyUsersProperties) *DescribeUsersResponseBodyUsers {
	s.Properties = v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetRealNickName(v string) *DescribeUsersResponseBodyUsers {
	s.RealNickName = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetRemark(v string) *DescribeUsersResponseBodyUsers {
	s.Remark = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetStatus(v int64) *DescribeUsersResponseBodyUsers {
	s.Status = &v
	return s
}

func (s *DescribeUsersResponseBodyUsers) SetWyId(v string) *DescribeUsersResponseBodyUsers {
	s.WyId = &v
	return s
}

type DescribeUsersResponseBodyUsersExtras struct {
	AssignedResourceCount map[string]interface{} `json:"AssignedResourceCount,omitempty" xml:"AssignedResourceCount,omitempty"`
}

func (s DescribeUsersResponseBodyUsersExtras) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBodyUsersExtras) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsersExtras) SetAssignedResourceCount(v map[string]interface{}) *DescribeUsersResponseBodyUsersExtras {
	s.AssignedResourceCount = v
	return s
}

type DescribeUsersResponseBodyUsersGroups struct {
	// The ID of the user group.
	//
	// example:
	//
	// ug-12341234****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the user group.
	//
	// example:
	//
	// User Group 1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DescribeUsersResponseBodyUsersGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBodyUsersGroups) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsersGroups) SetGroupId(v string) *DescribeUsersResponseBodyUsersGroups {
	s.GroupId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersGroups) SetGroupName(v string) *DescribeUsersResponseBodyUsersGroups {
	s.GroupName = &v
	return s
}

type DescribeUsersResponseBodyUsersOrgs struct {
	// The organization ID.
	//
	// example:
	//
	// org-4mdgc1cocc59z****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The organization name.
	//
	// example:
	//
	// Organization 1
	OrgName     *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	OrgNamePath *string `json:"OrgNamePath,omitempty" xml:"OrgNamePath,omitempty"`
}

func (s DescribeUsersResponseBodyUsersOrgs) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBodyUsersOrgs) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsersOrgs) SetOrgId(v string) *DescribeUsersResponseBodyUsersOrgs {
	s.OrgId = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersOrgs) SetOrgName(v string) *DescribeUsersResponseBodyUsersOrgs {
	s.OrgName = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersOrgs) SetOrgNamePath(v string) *DescribeUsersResponseBodyUsersOrgs {
	s.OrgNamePath = &v
	return s
}

type DescribeUsersResponseBodyUsersProperties struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeUsersResponseBodyUsersProperties) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponseBodyUsersProperties) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponseBodyUsersProperties) SetKey(v string) *DescribeUsersResponseBodyUsersProperties {
	s.Key = &v
	return s
}

func (s *DescribeUsersResponseBodyUsersProperties) SetValue(v string) *DescribeUsersResponseBodyUsersProperties {
	s.Value = &v
	return s
}

type DescribeUsersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsersResponse) SetHeaders(v map[string]*string) *DescribeUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsersResponse) SetStatusCode(v int32) *DescribeUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsersResponse) SetBody(v *DescribeUsersResponseBody) *DescribeUsersResponse {
	s.Body = v
	return s
}

type FilterUsersRequest struct {
	// The list of usernames to be precisely excluded.
	ExcludeEndUserIds []*string `json:"ExcludeEndUserIds,omitempty" xml:"ExcludeEndUserIds,omitempty" type:"Repeated"`
	// The string that is used for fuzzy search. You can use usernames and email addresses to perform fuzzy search. Wildcard characters (\\*) are supported for this parameter. For example, if you set this parameter to a\\*m, the usernames or an email addresses that start with a or end with m are returned.
	//
	// example:
	//
	// test
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to return the number of cloud desktops that are assigned to the convenience user.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	IncludeDesktopCount *bool `json:"IncludeDesktopCount,omitempty" xml:"IncludeDesktopCount,omitempty"`
	// Specifies whether to return the number of cloud desktop pools that are assigned to the convenience user.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	IncludeDesktopGroupCount *bool `json:"IncludeDesktopGroupCount,omitempty" xml:"IncludeDesktopGroupCount,omitempty"`
	IncludeOrgInfo           *bool `json:"IncludeOrgInfo,omitempty" xml:"IncludeOrgInfo,omitempty"`
	IncludeSupportIdps       *bool `json:"IncludeSupportIdps,omitempty" xml:"IncludeSupportIdps,omitempty"`
	IsQueryAllSubOrgs        *bool `json:"IsQueryAllSubOrgs,omitempty" xml:"IsQueryAllSubOrgs,omitempty"`
	// The number of entries per page. If you set this parameter to a value greater than 100, the system resets the value to 100.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. If not all results are returned in a query, a value is returned for the NextToken parameter. In this case, you can use the returned NextToken value to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The parameter that might affect the sorting logic.
	OrderParam *FilterUsersRequestOrderParam `json:"OrderParam,omitempty" xml:"OrderParam,omitempty" type:"Struct"`
	// The ID of the organization.
	//
	// example:
	//
	// org-aliyun-wy-org-id
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The activation type of the convenience account.
	//
	// Valid values:
	//
	// 	- CreateFromManager: administrator-activated.
	//
	// 	- Normal: user-activated.
	//
	// example:
	//
	// Normal
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The list of properties for fuzzy search.
	PropertyFilterParam []*FilterUsersRequestPropertyFilterParam `json:"PropertyFilterParam,omitempty" xml:"PropertyFilterParam,omitempty" type:"Repeated"`
	// The list of property names and property values.
	PropertyKeyValueFilterParam []*FilterUsersRequestPropertyKeyValueFilterParam `json:"PropertyKeyValueFilterParam,omitempty" xml:"PropertyKeyValueFilterParam,omitempty" type:"Repeated"`
	Status                      *int32                                           `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s FilterUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersRequest) GoString() string {
	return s.String()
}

func (s *FilterUsersRequest) SetExcludeEndUserIds(v []*string) *FilterUsersRequest {
	s.ExcludeEndUserIds = v
	return s
}

func (s *FilterUsersRequest) SetFilter(v string) *FilterUsersRequest {
	s.Filter = &v
	return s
}

func (s *FilterUsersRequest) SetIncludeDesktopCount(v bool) *FilterUsersRequest {
	s.IncludeDesktopCount = &v
	return s
}

func (s *FilterUsersRequest) SetIncludeDesktopGroupCount(v bool) *FilterUsersRequest {
	s.IncludeDesktopGroupCount = &v
	return s
}

func (s *FilterUsersRequest) SetIncludeOrgInfo(v bool) *FilterUsersRequest {
	s.IncludeOrgInfo = &v
	return s
}

func (s *FilterUsersRequest) SetIncludeSupportIdps(v bool) *FilterUsersRequest {
	s.IncludeSupportIdps = &v
	return s
}

func (s *FilterUsersRequest) SetIsQueryAllSubOrgs(v bool) *FilterUsersRequest {
	s.IsQueryAllSubOrgs = &v
	return s
}

func (s *FilterUsersRequest) SetMaxResults(v int64) *FilterUsersRequest {
	s.MaxResults = &v
	return s
}

func (s *FilterUsersRequest) SetNextToken(v string) *FilterUsersRequest {
	s.NextToken = &v
	return s
}

func (s *FilterUsersRequest) SetOrderParam(v *FilterUsersRequestOrderParam) *FilterUsersRequest {
	s.OrderParam = v
	return s
}

func (s *FilterUsersRequest) SetOrgId(v string) *FilterUsersRequest {
	s.OrgId = &v
	return s
}

func (s *FilterUsersRequest) SetOwnerType(v string) *FilterUsersRequest {
	s.OwnerType = &v
	return s
}

func (s *FilterUsersRequest) SetPropertyFilterParam(v []*FilterUsersRequestPropertyFilterParam) *FilterUsersRequest {
	s.PropertyFilterParam = v
	return s
}

func (s *FilterUsersRequest) SetPropertyKeyValueFilterParam(v []*FilterUsersRequestPropertyKeyValueFilterParam) *FilterUsersRequest {
	s.PropertyKeyValueFilterParam = v
	return s
}

func (s *FilterUsersRequest) SetStatus(v int32) *FilterUsersRequest {
	s.Status = &v
	return s
}

type FilterUsersRequestOrderParam struct {
	// The field that you want to sort by.
	//
	// Valid values:
	//
	// 	- EndUserId: the username.
	//
	// 	- id: the ID of the primary key.
	//
	// 	- gmt_created: the creation time.
	//
	// example:
	//
	// id
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The direction of the sort.
	//
	// Valid values:
	//
	// 	- ASC: the ascending order.
	//
	// 	- DESC (default): the descending order.
	//
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s FilterUsersRequestOrderParam) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersRequestOrderParam) GoString() string {
	return s.String()
}

func (s *FilterUsersRequestOrderParam) SetOrderField(v string) *FilterUsersRequestOrderParam {
	s.OrderField = &v
	return s
}

func (s *FilterUsersRequestOrderParam) SetOrderType(v string) *FilterUsersRequestOrderParam {
	s.OrderType = &v
	return s
}

type FilterUsersRequestPropertyFilterParam struct {
	// The ID of the property.
	//
	// example:
	//
	// 123
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The IDs of the property values.
	//
	// example:
	//
	// test
	PropertyValueIds *string `json:"PropertyValueIds,omitempty" xml:"PropertyValueIds,omitempty"`
}

func (s FilterUsersRequestPropertyFilterParam) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersRequestPropertyFilterParam) GoString() string {
	return s.String()
}

func (s *FilterUsersRequestPropertyFilterParam) SetPropertyId(v int64) *FilterUsersRequestPropertyFilterParam {
	s.PropertyId = &v
	return s
}

func (s *FilterUsersRequestPropertyFilterParam) SetPropertyValueIds(v string) *FilterUsersRequestPropertyFilterParam {
	s.PropertyValueIds = &v
	return s
}

type FilterUsersRequestPropertyKeyValueFilterParam struct {
	// The property name.
	//
	// example:
	//
	// job
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The property values.
	//
	// example:
	//
	// dev
	PropertyValues *string `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty"`
}

func (s FilterUsersRequestPropertyKeyValueFilterParam) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersRequestPropertyKeyValueFilterParam) GoString() string {
	return s.String()
}

func (s *FilterUsersRequestPropertyKeyValueFilterParam) SetPropertyKey(v string) *FilterUsersRequestPropertyKeyValueFilterParam {
	s.PropertyKey = &v
	return s
}

func (s *FilterUsersRequestPropertyKeyValueFilterParam) SetPropertyValues(v string) *FilterUsersRequestPropertyKeyValueFilterParam {
	s.PropertyValues = &v
	return s
}

type FilterUsersShrinkRequest struct {
	// The list of usernames to be precisely excluded.
	ExcludeEndUserIds []*string `json:"ExcludeEndUserIds,omitempty" xml:"ExcludeEndUserIds,omitempty" type:"Repeated"`
	// The string that is used for fuzzy search. You can use usernames and email addresses to perform fuzzy search. Wildcard characters (\\*) are supported for this parameter. For example, if you set this parameter to a\\*m, the usernames or an email addresses that start with a or end with m are returned.
	//
	// example:
	//
	// test
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// Specifies whether to return the number of cloud desktops that are assigned to the convenience user.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	IncludeDesktopCount *bool `json:"IncludeDesktopCount,omitempty" xml:"IncludeDesktopCount,omitempty"`
	// Specifies whether to return the number of cloud desktop pools that are assigned to the convenience user.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// false
	IncludeDesktopGroupCount *bool `json:"IncludeDesktopGroupCount,omitempty" xml:"IncludeDesktopGroupCount,omitempty"`
	IncludeOrgInfo           *bool `json:"IncludeOrgInfo,omitempty" xml:"IncludeOrgInfo,omitempty"`
	IncludeSupportIdps       *bool `json:"IncludeSupportIdps,omitempty" xml:"IncludeSupportIdps,omitempty"`
	IsQueryAllSubOrgs        *bool `json:"IsQueryAllSubOrgs,omitempty" xml:"IsQueryAllSubOrgs,omitempty"`
	// The number of entries per page. If you set this parameter to a value greater than 100, the system resets the value to 100.
	//
	// example:
	//
	// 10
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. If not all results are returned in a query, a value is returned for the NextToken parameter. In this case, you can use the returned NextToken value to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The parameter that might affect the sorting logic.
	OrderParamShrink *string `json:"OrderParam,omitempty" xml:"OrderParam,omitempty"`
	// The ID of the organization.
	//
	// example:
	//
	// org-aliyun-wy-org-id
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The activation type of the convenience account.
	//
	// Valid values:
	//
	// 	- CreateFromManager: administrator-activated.
	//
	// 	- Normal: user-activated.
	//
	// example:
	//
	// Normal
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The list of properties for fuzzy search.
	PropertyFilterParam []*FilterUsersShrinkRequestPropertyFilterParam `json:"PropertyFilterParam,omitempty" xml:"PropertyFilterParam,omitempty" type:"Repeated"`
	// The list of property names and property values.
	PropertyKeyValueFilterParam []*FilterUsersShrinkRequestPropertyKeyValueFilterParam `json:"PropertyKeyValueFilterParam,omitempty" xml:"PropertyKeyValueFilterParam,omitempty" type:"Repeated"`
	Status                      *int32                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s FilterUsersShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersShrinkRequest) GoString() string {
	return s.String()
}

func (s *FilterUsersShrinkRequest) SetExcludeEndUserIds(v []*string) *FilterUsersShrinkRequest {
	s.ExcludeEndUserIds = v
	return s
}

func (s *FilterUsersShrinkRequest) SetFilter(v string) *FilterUsersShrinkRequest {
	s.Filter = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetIncludeDesktopCount(v bool) *FilterUsersShrinkRequest {
	s.IncludeDesktopCount = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetIncludeDesktopGroupCount(v bool) *FilterUsersShrinkRequest {
	s.IncludeDesktopGroupCount = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetIncludeOrgInfo(v bool) *FilterUsersShrinkRequest {
	s.IncludeOrgInfo = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetIncludeSupportIdps(v bool) *FilterUsersShrinkRequest {
	s.IncludeSupportIdps = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetIsQueryAllSubOrgs(v bool) *FilterUsersShrinkRequest {
	s.IsQueryAllSubOrgs = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetMaxResults(v int64) *FilterUsersShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetNextToken(v string) *FilterUsersShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetOrderParamShrink(v string) *FilterUsersShrinkRequest {
	s.OrderParamShrink = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetOrgId(v string) *FilterUsersShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetOwnerType(v string) *FilterUsersShrinkRequest {
	s.OwnerType = &v
	return s
}

func (s *FilterUsersShrinkRequest) SetPropertyFilterParam(v []*FilterUsersShrinkRequestPropertyFilterParam) *FilterUsersShrinkRequest {
	s.PropertyFilterParam = v
	return s
}

func (s *FilterUsersShrinkRequest) SetPropertyKeyValueFilterParam(v []*FilterUsersShrinkRequestPropertyKeyValueFilterParam) *FilterUsersShrinkRequest {
	s.PropertyKeyValueFilterParam = v
	return s
}

func (s *FilterUsersShrinkRequest) SetStatus(v int32) *FilterUsersShrinkRequest {
	s.Status = &v
	return s
}

type FilterUsersShrinkRequestPropertyFilterParam struct {
	// The ID of the property.
	//
	// example:
	//
	// 123
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The IDs of the property values.
	//
	// example:
	//
	// test
	PropertyValueIds *string `json:"PropertyValueIds,omitempty" xml:"PropertyValueIds,omitempty"`
}

func (s FilterUsersShrinkRequestPropertyFilterParam) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersShrinkRequestPropertyFilterParam) GoString() string {
	return s.String()
}

func (s *FilterUsersShrinkRequestPropertyFilterParam) SetPropertyId(v int64) *FilterUsersShrinkRequestPropertyFilterParam {
	s.PropertyId = &v
	return s
}

func (s *FilterUsersShrinkRequestPropertyFilterParam) SetPropertyValueIds(v string) *FilterUsersShrinkRequestPropertyFilterParam {
	s.PropertyValueIds = &v
	return s
}

type FilterUsersShrinkRequestPropertyKeyValueFilterParam struct {
	// The property name.
	//
	// example:
	//
	// job
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The property values.
	//
	// example:
	//
	// dev
	PropertyValues *string `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty"`
}

func (s FilterUsersShrinkRequestPropertyKeyValueFilterParam) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersShrinkRequestPropertyKeyValueFilterParam) GoString() string {
	return s.String()
}

func (s *FilterUsersShrinkRequestPropertyKeyValueFilterParam) SetPropertyKey(v string) *FilterUsersShrinkRequestPropertyKeyValueFilterParam {
	s.PropertyKey = &v
	return s
}

func (s *FilterUsersShrinkRequestPropertyKeyValueFilterParam) SetPropertyValues(v string) *FilterUsersShrinkRequestPropertyKeyValueFilterParam {
	s.PropertyValues = &v
	return s
}

type FilterUsersResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results. If not all results are returned in a query, a value is returned for the NextToken parameter. In this case, you can use the returned NextToken value to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DBD276B5-00FF-5E04-8EF7-5CBA09BF112A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The convenience accounts.
	Users []*FilterUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s FilterUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponseBody) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBody) SetNextToken(v string) *FilterUsersResponseBody {
	s.NextToken = &v
	return s
}

func (s *FilterUsersResponseBody) SetRequestId(v string) *FilterUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *FilterUsersResponseBody) SetUsers(v []*FilterUsersResponseBodyUsers) *FilterUsersResponseBody {
	s.Users = v
	return s
}

type FilterUsersResponseBodyUsers struct {
	// The date when a convenience account is automatically locked.
	//
	// example:
	//
	// 2023-03-03
	AutoLockTime *string `json:"AutoLockTime,omitempty" xml:"AutoLockTime,omitempty"`
	// The number of cloud desktops that are assigned to the convenience user.
	//
	// example:
	//
	// 1
	DesktopCount *int64 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// The number of cloud desktop pools that are assigned to the convenience user. This value is returned if you set `IncludeDesktopGroupCount` to `true`.
	//
	// example:
	//
	// 2
	DesktopGroupCount *int64 `json:"DesktopGroupCount,omitempty" xml:"DesktopGroupCount,omitempty"`
	// The email address of the convenience user.
	//
	// example:
	//
	// testName@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the convenience user is a local administrator.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	EnableAdminAccess *bool `json:"EnableAdminAccess,omitempty" xml:"EnableAdminAccess,omitempty"`
	// The username of the convenience user.
	//
	// example:
	//
	// testName
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The additional information about the convenience user.
	ExternalInfo *FilterUsersResponseBodyUsersExternalInfo `json:"ExternalInfo,omitempty" xml:"ExternalInfo,omitempty" type:"Struct"`
	// The ID of the convenience user.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the convenience user is a tenant administrator.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// true
	IsTenantManager *bool                                  `json:"IsTenantManager,omitempty" xml:"IsTenantManager,omitempty"`
	OrgList         []*FilterUsersResponseBodyUsersOrgList `json:"OrgList,omitempty" xml:"OrgList,omitempty" type:"Repeated"`
	// The type of the account ownership.
	//
	// Valid values:
	//
	// 	- CreateFromManager: administrator-activated
	//
	// 	- Normal: user-activated
	//
	// example:
	//
	// Normal
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// By default, user account passwords do not expire. However, you can set a validity period between 30 and 365 days. Once the period expires, end users must change their password before they can log on to terminals.
	//
	// >  The feature is in invitational preview. If you want to use this feature, submit a ticket.
	//
	// example:
	//
	// 30
	PasswordExpireDays *int32 `json:"PasswordExpireDays,omitempty" xml:"PasswordExpireDays,omitempty"`
	// The number of days remaining until the account password expires.
	//
	// example:
	//
	// 10
	PasswordExpireRestDays *int32 `json:"PasswordExpireRestDays,omitempty" xml:"PasswordExpireRestDays,omitempty"`
	// The mobile number of the convenience user.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The nickname of the convenience user.
	//
	// example:
	//
	// Oliver
	RealNickName *string `json:"RealNickName,omitempty" xml:"RealNickName,omitempty"`
	// The remarks on the convenience user.
	//
	// example:
	//
	// 1
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The remarks on the convenience account.
	//
	// Valid values:
	//
	// 	- 0: The convenience account is normal.
	//
	// 	- 9: The convenience account is locked.
	//
	// example:
	//
	// 0
	Status           *int64                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportLoginIdps []*FilterUsersResponseBodyUsersSupportLoginIdps `json:"SupportLoginIdps,omitempty" xml:"SupportLoginIdps,omitempty" type:"Repeated"`
	// The information about the properties.
	UserSetPropertiesModels []*FilterUsersResponseBodyUsersUserSetPropertiesModels `json:"UserSetPropertiesModels,omitempty" xml:"UserSetPropertiesModels,omitempty" type:"Repeated"`
}

func (s FilterUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsers) SetAutoLockTime(v string) *FilterUsersResponseBodyUsers {
	s.AutoLockTime = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetDesktopCount(v int64) *FilterUsersResponseBodyUsers {
	s.DesktopCount = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetDesktopGroupCount(v int64) *FilterUsersResponseBodyUsers {
	s.DesktopGroupCount = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetEmail(v string) *FilterUsersResponseBodyUsers {
	s.Email = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetEnableAdminAccess(v bool) *FilterUsersResponseBodyUsers {
	s.EnableAdminAccess = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetEndUserId(v string) *FilterUsersResponseBodyUsers {
	s.EndUserId = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetExternalInfo(v *FilterUsersResponseBodyUsersExternalInfo) *FilterUsersResponseBodyUsers {
	s.ExternalInfo = v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetId(v int64) *FilterUsersResponseBodyUsers {
	s.Id = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetIsTenantManager(v bool) *FilterUsersResponseBodyUsers {
	s.IsTenantManager = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetOrgList(v []*FilterUsersResponseBodyUsersOrgList) *FilterUsersResponseBodyUsers {
	s.OrgList = v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetOwnerType(v string) *FilterUsersResponseBodyUsers {
	s.OwnerType = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetPasswordExpireDays(v int32) *FilterUsersResponseBodyUsers {
	s.PasswordExpireDays = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetPasswordExpireRestDays(v int32) *FilterUsersResponseBodyUsers {
	s.PasswordExpireRestDays = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetPhone(v string) *FilterUsersResponseBodyUsers {
	s.Phone = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetRealNickName(v string) *FilterUsersResponseBodyUsers {
	s.RealNickName = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetRemark(v string) *FilterUsersResponseBodyUsers {
	s.Remark = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetStatus(v int64) *FilterUsersResponseBodyUsers {
	s.Status = &v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetSupportLoginIdps(v []*FilterUsersResponseBodyUsersSupportLoginIdps) *FilterUsersResponseBodyUsers {
	s.SupportLoginIdps = v
	return s
}

func (s *FilterUsersResponseBodyUsers) SetUserSetPropertiesModels(v []*FilterUsersResponseBodyUsersUserSetPropertiesModels) *FilterUsersResponseBodyUsers {
	s.UserSetPropertiesModels = v
	return s
}

type FilterUsersResponseBodyUsersExternalInfo struct {
	// The account that is associated with the convenience user.
	//
	// example:
	//
	// test
	ExternalName *string `json:"ExternalName,omitempty" xml:"ExternalName,omitempty"`
	// The account, student ID, or employee ID that is associated with the convenience user.
	//
	// example:
	//
	// 030801
	JobNumber *string `json:"JobNumber,omitempty" xml:"JobNumber,omitempty"`
}

func (s FilterUsersResponseBodyUsersExternalInfo) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponseBodyUsersExternalInfo) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersExternalInfo) SetExternalName(v string) *FilterUsersResponseBodyUsersExternalInfo {
	s.ExternalName = &v
	return s
}

func (s *FilterUsersResponseBodyUsersExternalInfo) SetJobNumber(v string) *FilterUsersResponseBodyUsersExternalInfo {
	s.JobNumber = &v
	return s
}

type FilterUsersResponseBodyUsersOrgList struct {
	OrgId       *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	OrgName     *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	OrgNamePath *string `json:"OrgNamePath,omitempty" xml:"OrgNamePath,omitempty"`
}

func (s FilterUsersResponseBodyUsersOrgList) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponseBodyUsersOrgList) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersOrgList) SetOrgId(v string) *FilterUsersResponseBodyUsersOrgList {
	s.OrgId = &v
	return s
}

func (s *FilterUsersResponseBodyUsersOrgList) SetOrgName(v string) *FilterUsersResponseBodyUsersOrgList {
	s.OrgName = &v
	return s
}

func (s *FilterUsersResponseBodyUsersOrgList) SetOrgNamePath(v string) *FilterUsersResponseBodyUsersOrgList {
	s.OrgNamePath = &v
	return s
}

type FilterUsersResponseBodyUsersSupportLoginIdps struct {
	IdpId   *string `json:"IdpId,omitempty" xml:"IdpId,omitempty"`
	IdpName *string `json:"IdpName,omitempty" xml:"IdpName,omitempty"`
}

func (s FilterUsersResponseBodyUsersSupportLoginIdps) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponseBodyUsersSupportLoginIdps) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersSupportLoginIdps) SetIdpId(v string) *FilterUsersResponseBodyUsersSupportLoginIdps {
	s.IdpId = &v
	return s
}

func (s *FilterUsersResponseBodyUsersSupportLoginIdps) SetIdpName(v string) *FilterUsersResponseBodyUsersSupportLoginIdps {
	s.IdpName = &v
	return s
}

type FilterUsersResponseBodyUsersUserSetPropertiesModels struct {
	// The property ID.
	//
	// example:
	//
	// 12
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The property name.
	//
	// example:
	//
	// department
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The property type.
	//
	// example:
	//
	// 2
	PropertyType *int32 `json:"PropertyType,omitempty" xml:"PropertyType,omitempty"`
	// The property values.
	PropertyValues []*FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
	// The ID of the convenience user that is bound to the property.
	//
	// example:
	//
	// 12345
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the convenience user that is bound to the property.
	//
	// example:
	//
	// testName
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s FilterUsersResponseBodyUsersUserSetPropertiesModels) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponseBodyUsersUserSetPropertiesModels) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetPropertyId(v int64) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.PropertyId = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetPropertyKey(v string) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.PropertyKey = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetPropertyType(v int32) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.PropertyType = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetPropertyValues(v []*FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.PropertyValues = v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetUserId(v int64) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.UserId = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModels) SetUserName(v string) *FilterUsersResponseBodyUsersUserSetPropertiesModels {
	s.UserName = &v
	return s
}

type FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues struct {
	// The property value.
	//
	// example:
	//
	// A
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The ID of the property value.
	//
	// example:
	//
	// 42
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) GoString() string {
	return s.String()
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) SetPropertyValue(v string) *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues) SetPropertyValueId(v int64) *FilterUsersResponseBodyUsersUserSetPropertiesModelsPropertyValues {
	s.PropertyValueId = &v
	return s
}

type FilterUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FilterUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FilterUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s FilterUsersResponse) GoString() string {
	return s.String()
}

func (s *FilterUsersResponse) SetHeaders(v map[string]*string) *FilterUsersResponse {
	s.Headers = v
	return s
}

func (s *FilterUsersResponse) SetStatusCode(v int32) *FilterUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *FilterUsersResponse) SetBody(v *FilterUsersResponseBody) *FilterUsersResponse {
	s.Body = v
	return s
}

type GetManagerInfoByAuthCodeRequest struct {
	// The authorization code.
	//
	// This parameter is required.
	//
	// example:
	//
	// e49cd070452f0044813a467d4743****
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
}

func (s GetManagerInfoByAuthCodeRequest) String() string {
	return tea.Prettify(s)
}

func (s GetManagerInfoByAuthCodeRequest) GoString() string {
	return s.String()
}

func (s *GetManagerInfoByAuthCodeRequest) SetAuthCode(v string) *GetManagerInfoByAuthCodeRequest {
	s.AuthCode = &v
	return s
}

type GetManagerInfoByAuthCodeResponseBody struct {
	// The organization ID.
	//
	// example:
	//
	// 12345678901234****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The mobile number.
	//
	// example:
	//
	// 1301234****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The team name.
	//
	// example:
	//
	// devteam
	TeamName *string `json:"TeamName,omitempty" xml:"TeamName,omitempty"`
	// The tenant name.
	//
	// example:
	//
	// zhangsan
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The ID of the Elastic Desktop Service account.
	//
	// example:
	//
	// 12345678901234****
	WaId *int64 `json:"WaId,omitempty" xml:"WaId,omitempty"`
}

func (s GetManagerInfoByAuthCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetManagerInfoByAuthCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetManagerInfoByAuthCodeResponseBody) SetOrgId(v string) *GetManagerInfoByAuthCodeResponseBody {
	s.OrgId = &v
	return s
}

func (s *GetManagerInfoByAuthCodeResponseBody) SetPhone(v string) *GetManagerInfoByAuthCodeResponseBody {
	s.Phone = &v
	return s
}

func (s *GetManagerInfoByAuthCodeResponseBody) SetRequestId(v string) *GetManagerInfoByAuthCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetManagerInfoByAuthCodeResponseBody) SetTeamName(v string) *GetManagerInfoByAuthCodeResponseBody {
	s.TeamName = &v
	return s
}

func (s *GetManagerInfoByAuthCodeResponseBody) SetUserName(v string) *GetManagerInfoByAuthCodeResponseBody {
	s.UserName = &v
	return s
}

func (s *GetManagerInfoByAuthCodeResponseBody) SetWaId(v int64) *GetManagerInfoByAuthCodeResponseBody {
	s.WaId = &v
	return s
}

type GetManagerInfoByAuthCodeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetManagerInfoByAuthCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetManagerInfoByAuthCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s GetManagerInfoByAuthCodeResponse) GoString() string {
	return s.String()
}

func (s *GetManagerInfoByAuthCodeResponse) SetHeaders(v map[string]*string) *GetManagerInfoByAuthCodeResponse {
	s.Headers = v
	return s
}

func (s *GetManagerInfoByAuthCodeResponse) SetStatusCode(v int32) *GetManagerInfoByAuthCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetManagerInfoByAuthCodeResponse) SetBody(v *GetManagerInfoByAuthCodeResponseBody) *GetManagerInfoByAuthCodeResponse {
	s.Body = v
	return s
}

type InitTenantAliasResponseBody struct {
	AliasInfo *InitTenantAliasResponseBodyAliasInfo `json:"AliasInfo,omitempty" xml:"AliasInfo,omitempty" type:"Struct"`
	// example:
	//
	// WY23***
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitTenantAliasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitTenantAliasResponseBody) GoString() string {
	return s.String()
}

func (s *InitTenantAliasResponseBody) SetAliasInfo(v *InitTenantAliasResponseBodyAliasInfo) *InitTenantAliasResponseBody {
	s.AliasInfo = v
	return s
}

func (s *InitTenantAliasResponseBody) SetData(v string) *InitTenantAliasResponseBody {
	s.Data = &v
	return s
}

func (s *InitTenantAliasResponseBody) SetRequestId(v string) *InitTenantAliasResponseBody {
	s.RequestId = &v
	return s
}

type InitTenantAliasResponseBodyAliasInfo struct {
	// example:
	//
	// FrequencyExceedsLimit
	AliasEditDisabledReason *string `json:"AliasEditDisabledReason,omitempty" xml:"AliasEditDisabledReason,omitempty"`
	// example:
	//
	// False
	AliasEditable *bool `json:"AliasEditable,omitempty" xml:"AliasEditable,omitempty"`
	// example:
	//
	// Customized
	AliasSourceType *string `json:"AliasSourceType,omitempty" xml:"AliasSourceType,omitempty"`
	// example:
	//
	// 2025-04-17 20:31:48
	NextModifyTime *string `json:"NextModifyTime,omitempty" xml:"NextModifyTime,omitempty"`
}

func (s InitTenantAliasResponseBodyAliasInfo) String() string {
	return tea.Prettify(s)
}

func (s InitTenantAliasResponseBodyAliasInfo) GoString() string {
	return s.String()
}

func (s *InitTenantAliasResponseBodyAliasInfo) SetAliasEditDisabledReason(v string) *InitTenantAliasResponseBodyAliasInfo {
	s.AliasEditDisabledReason = &v
	return s
}

func (s *InitTenantAliasResponseBodyAliasInfo) SetAliasEditable(v bool) *InitTenantAliasResponseBodyAliasInfo {
	s.AliasEditable = &v
	return s
}

func (s *InitTenantAliasResponseBodyAliasInfo) SetAliasSourceType(v string) *InitTenantAliasResponseBodyAliasInfo {
	s.AliasSourceType = &v
	return s
}

func (s *InitTenantAliasResponseBodyAliasInfo) SetNextModifyTime(v string) *InitTenantAliasResponseBodyAliasInfo {
	s.NextModifyTime = &v
	return s
}

type InitTenantAliasResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitTenantAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitTenantAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s InitTenantAliasResponse) GoString() string {
	return s.String()
}

func (s *InitTenantAliasResponse) SetHeaders(v map[string]*string) *InitTenantAliasResponse {
	s.Headers = v
	return s
}

func (s *InitTenantAliasResponse) SetStatusCode(v int32) *InitTenantAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *InitTenantAliasResponse) SetBody(v *InitTenantAliasResponseBody) *InitTenantAliasResponse {
	s.Body = v
	return s
}

type ListPropertyResponseBody struct {
	// The token that is used for the next query. If this parameter is empty, all results have been returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The information about the properties.
	Properties []*ListPropertyResponseBodyProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *ListPropertyResponseBody) SetNextToken(v string) *ListPropertyResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPropertyResponseBody) SetProperties(v []*ListPropertyResponseBodyProperties) *ListPropertyResponseBody {
	s.Properties = v
	return s
}

func (s *ListPropertyResponseBody) SetRequestId(v string) *ListPropertyResponseBody {
	s.RequestId = &v
	return s
}

type ListPropertyResponseBodyProperties struct {
	// The ID of the property.
	//
	// example:
	//
	// 30
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The name of the property.
	//
	// example:
	//
	// department
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// Details about the property values.
	PropertyValues []*ListPropertyResponseBodyPropertiesPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
}

func (s ListPropertyResponseBodyProperties) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyResponseBodyProperties) GoString() string {
	return s.String()
}

func (s *ListPropertyResponseBodyProperties) SetPropertyId(v int64) *ListPropertyResponseBodyProperties {
	s.PropertyId = &v
	return s
}

func (s *ListPropertyResponseBodyProperties) SetPropertyKey(v string) *ListPropertyResponseBodyProperties {
	s.PropertyKey = &v
	return s
}

func (s *ListPropertyResponseBodyProperties) SetPropertyValues(v []*ListPropertyResponseBodyPropertiesPropertyValues) *ListPropertyResponseBodyProperties {
	s.PropertyValues = v
	return s
}

type ListPropertyResponseBodyPropertiesPropertyValues struct {
	// The value of the property.
	//
	// example:
	//
	// A
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The ID of the property value.
	//
	// example:
	//
	// 42
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s ListPropertyResponseBodyPropertiesPropertyValues) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyResponseBodyPropertiesPropertyValues) GoString() string {
	return s.String()
}

func (s *ListPropertyResponseBodyPropertiesPropertyValues) SetPropertyValue(v string) *ListPropertyResponseBodyPropertiesPropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *ListPropertyResponseBodyPropertiesPropertyValues) SetPropertyValueId(v int64) *ListPropertyResponseBodyPropertiesPropertyValues {
	s.PropertyValueId = &v
	return s
}

type ListPropertyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyResponse) GoString() string {
	return s.String()
}

func (s *ListPropertyResponse) SetHeaders(v map[string]*string) *ListPropertyResponse {
	s.Headers = v
	return s
}

func (s *ListPropertyResponse) SetStatusCode(v int32) *ListPropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPropertyResponse) SetBody(v *ListPropertyResponseBody) *ListPropertyResponse {
	s.Body = v
	return s
}

type ListPropertyValueRequest struct {
	// The ID of the property. You can call the [ListProperty](https://help.aliyun.com/document_detail/410890.html) operation to query the property ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
}

func (s ListPropertyValueRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyValueRequest) GoString() string {
	return s.String()
}

func (s *ListPropertyValueRequest) SetPropertyId(v int64) *ListPropertyValueRequest {
	s.PropertyId = &v
	return s
}

type ListPropertyValueResponseBody struct {
	// Details about property values.
	PropertyValueInfos []*ListPropertyValueResponseBodyPropertyValueInfos `json:"PropertyValueInfos,omitempty" xml:"PropertyValueInfos,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C52013A5-3422-5D1F-B22C-A57110972AD9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPropertyValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyValueResponseBody) GoString() string {
	return s.String()
}

func (s *ListPropertyValueResponseBody) SetPropertyValueInfos(v []*ListPropertyValueResponseBodyPropertyValueInfos) *ListPropertyValueResponseBody {
	s.PropertyValueInfos = v
	return s
}

func (s *ListPropertyValueResponseBody) SetRequestId(v string) *ListPropertyValueResponseBody {
	s.RequestId = &v
	return s
}

type ListPropertyValueResponseBodyPropertyValueInfos struct {
	// The value of the property.
	//
	// example:
	//
	// HR
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The ID of the property value.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s ListPropertyValueResponseBodyPropertyValueInfos) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyValueResponseBodyPropertyValueInfos) GoString() string {
	return s.String()
}

func (s *ListPropertyValueResponseBodyPropertyValueInfos) SetPropertyValue(v string) *ListPropertyValueResponseBodyPropertyValueInfos {
	s.PropertyValue = &v
	return s
}

func (s *ListPropertyValueResponseBodyPropertyValueInfos) SetPropertyValueId(v int64) *ListPropertyValueResponseBodyPropertyValueInfos {
	s.PropertyValueId = &v
	return s
}

type ListPropertyValueResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPropertyValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPropertyValueResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPropertyValueResponse) GoString() string {
	return s.String()
}

func (s *ListPropertyValueResponse) SetHeaders(v map[string]*string) *ListPropertyValueResponse {
	s.Headers = v
	return s
}

func (s *ListPropertyValueResponse) SetStatusCode(v int32) *ListPropertyValueResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPropertyValueResponse) SetBody(v *ListPropertyValueResponseBody) *ListPropertyValueResponse {
	s.Body = v
	return s
}

type LockMfaDeviceRequest struct {
	// The domain of the Active Directory (AD) workspace.
	//
	// example:
	//
	// pg-jifenn.com
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The serial number of the virtual MFA device. The serial number is unique for each device.
	//
	// example:
	//
	// dc856334-446b-4035-bfbc-18af261e****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s LockMfaDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s LockMfaDeviceRequest) GoString() string {
	return s.String()
}

func (s *LockMfaDeviceRequest) SetAdDomain(v string) *LockMfaDeviceRequest {
	s.AdDomain = &v
	return s
}

func (s *LockMfaDeviceRequest) SetSerialNumber(v string) *LockMfaDeviceRequest {
	s.SerialNumber = &v
	return s
}

type LockMfaDeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 42FE70D8-4336-471B-8314-CCCFCE4159FD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockMfaDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockMfaDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *LockMfaDeviceResponseBody) SetRequestId(v string) *LockMfaDeviceResponseBody {
	s.RequestId = &v
	return s
}

type LockMfaDeviceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockMfaDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LockMfaDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s LockMfaDeviceResponse) GoString() string {
	return s.String()
}

func (s *LockMfaDeviceResponse) SetHeaders(v map[string]*string) *LockMfaDeviceResponse {
	s.Headers = v
	return s
}

func (s *LockMfaDeviceResponse) SetStatusCode(v int32) *LockMfaDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *LockMfaDeviceResponse) SetBody(v *LockMfaDeviceResponseBody) *LockMfaDeviceResponse {
	s.Body = v
	return s
}

type LockUsersRequest struct {
	LogoutSession *bool `json:"LogoutSession,omitempty" xml:"LogoutSession,omitempty"`
	// The usernames of the convenience users that you want to lock.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s LockUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s LockUsersRequest) GoString() string {
	return s.String()
}

func (s *LockUsersRequest) SetLogoutSession(v bool) *LockUsersRequest {
	s.LogoutSession = &v
	return s
}

func (s *LockUsersRequest) SetUsers(v []*string) *LockUsersRequest {
	s.Users = v
	return s
}

type LockUsersResponseBody struct {
	// The result of the locking the convenience user.
	LockUsersResult *LockUsersResponseBodyLockUsersResult `json:"LockUsersResult,omitempty" xml:"LockUsersResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LockUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LockUsersResponseBody) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBody) SetLockUsersResult(v *LockUsersResponseBodyLockUsersResult) *LockUsersResponseBody {
	s.LockUsersResult = v
	return s
}

func (s *LockUsersResponseBody) SetRequestId(v string) *LockUsersResponseBody {
	s.RequestId = &v
	return s
}

type LockUsersResponseBodyLockUsersResult struct {
	// The convenience users that failed to be locked.
	FailedUsers []*LockUsersResponseBodyLockUsersResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
	// The convenience users that were locked.
	LockedUsers []*string `json:"LockedUsers,omitempty" xml:"LockedUsers,omitempty" type:"Repeated"`
}

func (s LockUsersResponseBodyLockUsersResult) String() string {
	return tea.Prettify(s)
}

func (s LockUsersResponseBodyLockUsersResult) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBodyLockUsersResult) SetFailedUsers(v []*LockUsersResponseBodyLockUsersResultFailedUsers) *LockUsersResponseBodyLockUsersResult {
	s.FailedUsers = v
	return s
}

func (s *LockUsersResponseBodyLockUsersResult) SetLockedUsers(v []*string) *LockUsersResponseBodyLockUsersResult {
	s.LockedUsers = v
	return s
}

type LockUsersResponseBodyLockUsersResultFailedUsers struct {
	// The ID of the convenience user that failed to be locked.
	//
	// example:
	//
	// test123
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidUsername
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// test123 is an invalid username.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s LockUsersResponseBodyLockUsersResultFailedUsers) String() string {
	return tea.Prettify(s)
}

func (s LockUsersResponseBodyLockUsersResultFailedUsers) GoString() string {
	return s.String()
}

func (s *LockUsersResponseBodyLockUsersResultFailedUsers) SetEndUserId(v string) *LockUsersResponseBodyLockUsersResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *LockUsersResponseBodyLockUsersResultFailedUsers) SetErrorCode(v string) *LockUsersResponseBodyLockUsersResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *LockUsersResponseBodyLockUsersResultFailedUsers) SetErrorMessage(v string) *LockUsersResponseBodyLockUsersResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

type LockUsersResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LockUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LockUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s LockUsersResponse) GoString() string {
	return s.String()
}

func (s *LockUsersResponse) SetHeaders(v map[string]*string) *LockUsersResponse {
	s.Headers = v
	return s
}

func (s *LockUsersResponse) SetStatusCode(v int32) *LockUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *LockUsersResponse) SetBody(v *LockUsersResponseBody) *LockUsersResponse {
	s.Body = v
	return s
}

type ModifyGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ug-12341234****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	NewGroupName *string `json:"NewGroupName,omitempty" xml:"NewGroupName,omitempty"`
}

func (s ModifyGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyGroupRequest) SetGroupId(v string) *ModifyGroupRequest {
	s.GroupId = &v
	return s
}

func (s *ModifyGroupRequest) SetNewGroupName(v string) *ModifyGroupRequest {
	s.NewGroupName = &v
	return s
}

type ModifyGroupResponseBody struct {
	// example:
	//
	// 868B8926-2E7A-5BE7-9897-E811E994****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGroupResponseBody) SetRequestId(v string) *ModifyGroupResponseBody {
	s.RequestId = &v
	return s
}

type ModifyGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyGroupResponse) SetHeaders(v map[string]*string) *ModifyGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyGroupResponse) SetStatusCode(v int32) *ModifyGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGroupResponse) SetBody(v *ModifyGroupResponseBody) *ModifyGroupResponse {
	s.Body = v
	return s
}

type ModifyOrgRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// org-76joc57khvnhdh***
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// This parameter is required.
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
}

func (s ModifyOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyOrgRequest) GoString() string {
	return s.String()
}

func (s *ModifyOrgRequest) SetOrgId(v string) *ModifyOrgRequest {
	s.OrgId = &v
	return s
}

func (s *ModifyOrgRequest) SetOrgName(v string) *ModifyOrgRequest {
	s.OrgName = &v
	return s
}

type ModifyOrgResponseBody struct {
	// example:
	//
	// 0296EDF8-3C8A-5128-8682-27B29C99****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyOrgResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyOrgResponseBody) SetRequestId(v string) *ModifyOrgResponseBody {
	s.RequestId = &v
	return s
}

type ModifyOrgResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyOrgResponse) GoString() string {
	return s.String()
}

func (s *ModifyOrgResponse) SetHeaders(v map[string]*string) *ModifyOrgResponse {
	s.Headers = v
	return s
}

func (s *ModifyOrgResponse) SetStatusCode(v int32) *ModifyOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyOrgResponse) SetBody(v *ModifyOrgResponseBody) *ModifyOrgResponse {
	s.Body = v
	return s
}

type ModifyUserRequest struct {
	// The email address of the convenience user. For a user-activated convenience user, the email address or mobile number must be verified. You can choose to verify the email address or the mobile number. For an administrator-activated convenience user, the email address and mobile number can be left empty.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The mobile number of the convenience user. For a user-activated convenience user, the email address or mobile number must be verified. You can choose to verify the email address or the mobile number. For an administrator-activated convenience user, the email address and mobile number can be left empty.
	//
	// >  Accounts created on the International site (alibabacloud.com) do not support mobile number-based authentication.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s ModifyUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserRequest) SetEmail(v string) *ModifyUserRequest {
	s.Email = &v
	return s
}

func (s *ModifyUserRequest) SetEndUserId(v string) *ModifyUserRequest {
	s.EndUserId = &v
	return s
}

func (s *ModifyUserRequest) SetPhone(v string) *ModifyUserRequest {
	s.Phone = &v
	return s
}

type ModifyUserResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserResponseBody) SetRequestId(v string) *ModifyUserResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserResponse) SetHeaders(v map[string]*string) *ModifyUserResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserResponse) SetStatusCode(v int32) *ModifyUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserResponse) SetBody(v *ModifyUserResponseBody) *ModifyUserResponse {
	s.Body = v
	return s
}

type MoveOrgRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// org-5yy5icj981xe5****
	NewParentOrgId *string `json:"NewParentOrgId,omitempty" xml:"NewParentOrgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// org-5yy5icj981xe5****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s MoveOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveOrgRequest) GoString() string {
	return s.String()
}

func (s *MoveOrgRequest) SetNewParentOrgId(v string) *MoveOrgRequest {
	s.NewParentOrgId = &v
	return s
}

func (s *MoveOrgRequest) SetOrgId(v string) *MoveOrgRequest {
	s.OrgId = &v
	return s
}

type MoveOrgResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveOrgResponseBody) GoString() string {
	return s.String()
}

func (s *MoveOrgResponseBody) SetRequestId(v string) *MoveOrgResponseBody {
	s.RequestId = &v
	return s
}

type MoveOrgResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveOrgResponse) GoString() string {
	return s.String()
}

func (s *MoveOrgResponse) SetHeaders(v map[string]*string) *MoveOrgResponse {
	s.Headers = v
	return s
}

func (s *MoveOrgResponse) SetStatusCode(v int32) *MoveOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveOrgResponse) SetBody(v *MoveOrgResponseBody) *MoveOrgResponse {
	s.Body = v
	return s
}

type MoveUserOrgRequest struct {
	// This parameter is required.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// org-4mdgc1cocc59z****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s MoveUserOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveUserOrgRequest) GoString() string {
	return s.String()
}

func (s *MoveUserOrgRequest) SetEndUserIds(v []*string) *MoveUserOrgRequest {
	s.EndUserIds = v
	return s
}

func (s *MoveUserOrgRequest) SetOrgId(v string) *MoveUserOrgRequest {
	s.OrgId = &v
	return s
}

type MoveUserOrgResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MoveUserOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveUserOrgResponseBody) GoString() string {
	return s.String()
}

func (s *MoveUserOrgResponseBody) SetRequestId(v string) *MoveUserOrgResponseBody {
	s.RequestId = &v
	return s
}

type MoveUserOrgResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveUserOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveUserOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveUserOrgResponse) GoString() string {
	return s.String()
}

func (s *MoveUserOrgResponse) SetHeaders(v map[string]*string) *MoveUserOrgResponse {
	s.Headers = v
	return s
}

func (s *MoveUserOrgResponse) SetStatusCode(v int32) *MoveUserOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveUserOrgResponse) SetBody(v *MoveUserOrgResponseBody) *MoveUserOrgResponse {
	s.Body = v
	return s
}

type QuerySyncStatusByAliUidResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *QuerySyncStatusByAliUidResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySyncStatusByAliUidResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QuerySyncStatusByAliUidResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySyncStatusByAliUidResponseBody) SetCode(v string) *QuerySyncStatusByAliUidResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetData(v *QuerySyncStatusByAliUidResponseBodyData) *QuerySyncStatusByAliUidResponseBody {
	s.Data = v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetHttpStatusCode(v int32) *QuerySyncStatusByAliUidResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetMessage(v string) *QuerySyncStatusByAliUidResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetRequestId(v string) *QuerySyncStatusByAliUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBody) SetSuccess(v bool) *QuerySyncStatusByAliUidResponseBody {
	s.Success = &v
	return s
}

type QuerySyncStatusByAliUidResponseBodyData struct {
	// example:
	//
	// 131239236086****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// example:
	//
	// cdrs948144195608****
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// 2020-06-30 07:50:42
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2022-03-02 14:27:39
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 18500
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2022-03-02 14:31:39
	LatestBeginTime *string `json:"LatestBeginTime,omitempty" xml:"LatestBeginTime,omitempty"`
	// example:
	//
	// 2022-03-02 16:13:12
	LatestEndTime *string `json:"LatestEndTime,omitempty" xml:"LatestEndTime,omitempty"`
	// example:
	//
	// 2022-03-02 18:24:01
	LatestSuccessTime *string `json:"LatestSuccessTime,omitempty" xml:"LatestSuccessTime,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QuerySyncStatusByAliUidResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s QuerySyncStatusByAliUidResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetAliUid(v int64) *QuerySyncStatusByAliUidResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetCorpId(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetGmtCreated(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetGmtModified(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetId(v int64) *QuerySyncStatusByAliUidResponseBodyData {
	s.Id = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetLatestBeginTime(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.LatestBeginTime = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetLatestEndTime(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.LatestEndTime = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetLatestSuccessTime(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.LatestSuccessTime = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponseBodyData) SetStatus(v string) *QuerySyncStatusByAliUidResponseBodyData {
	s.Status = &v
	return s
}

type QuerySyncStatusByAliUidResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySyncStatusByAliUidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySyncStatusByAliUidResponse) String() string {
	return tea.Prettify(s)
}

func (s QuerySyncStatusByAliUidResponse) GoString() string {
	return s.String()
}

func (s *QuerySyncStatusByAliUidResponse) SetHeaders(v map[string]*string) *QuerySyncStatusByAliUidResponse {
	s.Headers = v
	return s
}

func (s *QuerySyncStatusByAliUidResponse) SetStatusCode(v int32) *QuerySyncStatusByAliUidResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySyncStatusByAliUidResponse) SetBody(v *QuerySyncStatusByAliUidResponseBody) *QuerySyncStatusByAliUidResponse {
	s.Body = v
	return s
}

type RemoveGroupRequest struct {
	// example:
	//
	// ug-12341234****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s RemoveGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupRequest) SetGroupId(v string) *RemoveGroupRequest {
	s.GroupId = &v
	return s
}

type RemoveGroupResponseBody struct {
	// example:
	//
	// 42FE70D8-4336-471B-8314-CCCFCE41****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveGroupResponseBody) SetRequestId(v string) *RemoveGroupResponseBody {
	s.RequestId = &v
	return s
}

type RemoveGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupResponse) SetHeaders(v map[string]*string) *RemoveGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveGroupResponse) SetStatusCode(v int32) *RemoveGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveGroupResponse) SetBody(v *RemoveGroupResponseBody) *RemoveGroupResponse {
	s.Body = v
	return s
}

type RemoveMfaDeviceRequest struct {
	// The domain of the Active Directory (AD) workspace.
	//
	// example:
	//
	// alpha.lftltd.net
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The serial number of the virtual MFA device. The serial number is unique for each device.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc856334-446b-4035-bfbc-18af261e****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s RemoveMfaDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveMfaDeviceRequest) GoString() string {
	return s.String()
}

func (s *RemoveMfaDeviceRequest) SetAdDomain(v string) *RemoveMfaDeviceRequest {
	s.AdDomain = &v
	return s
}

func (s *RemoveMfaDeviceRequest) SetSerialNumber(v string) *RemoveMfaDeviceRequest {
	s.SerialNumber = &v
	return s
}

type RemoveMfaDeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FB550AAB-FB36-4A91-93F6-F4374AF65403
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveMfaDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveMfaDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveMfaDeviceResponseBody) SetRequestId(v string) *RemoveMfaDeviceResponseBody {
	s.RequestId = &v
	return s
}

type RemoveMfaDeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveMfaDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveMfaDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveMfaDeviceResponse) GoString() string {
	return s.String()
}

func (s *RemoveMfaDeviceResponse) SetHeaders(v map[string]*string) *RemoveMfaDeviceResponse {
	s.Headers = v
	return s
}

func (s *RemoveMfaDeviceResponse) SetStatusCode(v int32) *RemoveMfaDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveMfaDeviceResponse) SetBody(v *RemoveMfaDeviceResponseBody) *RemoveMfaDeviceResponse {
	s.Body = v
	return s
}

type RemoveOrgRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// org-5yy5icj981xe5****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
}

func (s RemoveOrgRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveOrgRequest) GoString() string {
	return s.String()
}

func (s *RemoveOrgRequest) SetOrgId(v string) *RemoveOrgRequest {
	s.OrgId = &v
	return s
}

type RemoveOrgResponseBody struct {
	// example:
	//
	// 6C352609-EE7F-5603-B5E6-57C3EDDD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveOrgResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveOrgResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveOrgResponseBody) SetRequestId(v string) *RemoveOrgResponseBody {
	s.RequestId = &v
	return s
}

type RemoveOrgResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveOrgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveOrgResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveOrgResponse) GoString() string {
	return s.String()
}

func (s *RemoveOrgResponse) SetHeaders(v map[string]*string) *RemoveOrgResponse {
	s.Headers = v
	return s
}

func (s *RemoveOrgResponse) SetStatusCode(v int32) *RemoveOrgResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveOrgResponse) SetBody(v *RemoveOrgResponseBody) *RemoveOrgResponse {
	s.Body = v
	return s
}

type RemovePropertyRequest struct {
	// The ID of the property. You can call the [ListProperty](https://help.aliyun.com/document_detail/410890.html) operation to query the property ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
}

func (s RemovePropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s RemovePropertyRequest) GoString() string {
	return s.String()
}

func (s *RemovePropertyRequest) SetPropertyId(v int64) *RemovePropertyRequest {
	s.PropertyId = &v
	return s
}

type RemovePropertyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemovePropertyResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePropertyResponseBody) SetRequestId(v string) *RemovePropertyResponseBody {
	s.RequestId = &v
	return s
}

type RemovePropertyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemovePropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemovePropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s RemovePropertyResponse) GoString() string {
	return s.String()
}

func (s *RemovePropertyResponse) SetHeaders(v map[string]*string) *RemovePropertyResponse {
	s.Headers = v
	return s
}

func (s *RemovePropertyResponse) SetStatusCode(v int32) *RemovePropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *RemovePropertyResponse) SetBody(v *RemovePropertyResponseBody) *RemovePropertyResponse {
	s.Body = v
	return s
}

type RemoveUsersRequest struct {
	// The usernames of the convenience users that you want to remove.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s RemoveUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersRequest) SetUsers(v []*string) *RemoveUsersRequest {
	s.Users = v
	return s
}

type RemoveUsersResponseBody struct {
	// The result of removing the convenience user.
	RemoveUsersResult *RemoveUsersResponseBodyRemoveUsersResult `json:"RemoveUsersResult,omitempty" xml:"RemoveUsersResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBody) SetRemoveUsersResult(v *RemoveUsersResponseBodyRemoveUsersResult) *RemoveUsersResponseBody {
	s.RemoveUsersResult = v
	return s
}

func (s *RemoveUsersResponseBody) SetRequestId(v string) *RemoveUsersResponseBody {
	s.RequestId = &v
	return s
}

type RemoveUsersResponseBodyRemoveUsersResult struct {
	// The convenience users that failed to be removed.
	FailedUsers []*RemoveUsersResponseBodyRemoveUsersResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
	// The convenience users that were removed.
	RemovedUsers []*string `json:"RemovedUsers,omitempty" xml:"RemovedUsers,omitempty" type:"Repeated"`
}

func (s RemoveUsersResponseBodyRemoveUsersResult) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersResponseBodyRemoveUsersResult) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBodyRemoveUsersResult) SetFailedUsers(v []*RemoveUsersResponseBodyRemoveUsersResultFailedUsers) *RemoveUsersResponseBodyRemoveUsersResult {
	s.FailedUsers = v
	return s
}

func (s *RemoveUsersResponseBodyRemoveUsersResult) SetRemovedUsers(v []*string) *RemoveUsersResponseBodyRemoveUsersResult {
	s.RemovedUsers = v
	return s
}

type RemoveUsersResponseBodyRemoveUsersResultFailedUsers struct {
	// The ID of the convenience user that failed to be removed.
	//
	// example:
	//
	// test2
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidUsername
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// test2 is an invalid username.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s RemoveUsersResponseBodyRemoveUsersResultFailedUsers) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersResponseBodyRemoveUsersResultFailedUsers) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBodyRemoveUsersResultFailedUsers) SetEndUserId(v string) *RemoveUsersResponseBodyRemoveUsersResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *RemoveUsersResponseBodyRemoveUsersResultFailedUsers) SetErrorCode(v string) *RemoveUsersResponseBodyRemoveUsersResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *RemoveUsersResponseBodyRemoveUsersResultFailedUsers) SetErrorMessage(v string) *RemoveUsersResponseBodyRemoveUsersResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

type RemoveUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersResponse) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponse) SetHeaders(v map[string]*string) *RemoveUsersResponse {
	s.Headers = v
	return s
}

func (s *RemoveUsersResponse) SetStatusCode(v int32) *RemoveUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUsersResponse) SetBody(v *RemoveUsersResponseBody) *RemoveUsersResponse {
	s.Body = v
	return s
}

type ResetUserPasswordRequest struct {
	// The method to notify the user after the password is reset.
	//
	// > Alibaba Cloud accounts of the international site do not support sending notification through text messages.
	//
	// example:
	//
	// 1
	NotifyType *int32 `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
	// The names of the convenience users whose passwords you want to reset.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ResetUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordRequest) SetNotifyType(v int32) *ResetUserPasswordRequest {
	s.NotifyType = &v
	return s
}

func (s *ResetUserPasswordRequest) SetUsers(v []*string) *ResetUserPasswordRequest {
	s.Users = v
	return s
}

type ResetUserPasswordResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 32D05B39-E6EE-4D7A-9FD0-762A26859D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of resetting the password of the convenience user.
	ResetUsersResult *ResetUserPasswordResponseBodyResetUsersResult `json:"ResetUsersResult,omitempty" xml:"ResetUsersResult,omitempty" type:"Struct"`
}

func (s ResetUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBody) SetRequestId(v string) *ResetUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetResetUsersResult(v *ResetUserPasswordResponseBodyResetUsersResult) *ResetUserPasswordResponseBody {
	s.ResetUsersResult = v
	return s
}

type ResetUserPasswordResponseBodyResetUsersResult struct {
	// The information about the convenience users whose passwords failed to be reset.
	FailedUsers []*ResetUserPasswordResponseBodyResetUsersResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
	// The convenience users to which the system sent a password reset email.
	ResetUsers []*string `json:"ResetUsers,omitempty" xml:"ResetUsers,omitempty" type:"Repeated"`
}

func (s ResetUserPasswordResponseBodyResetUsersResult) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordResponseBodyResetUsersResult) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBodyResetUsersResult) SetFailedUsers(v []*ResetUserPasswordResponseBodyResetUsersResultFailedUsers) *ResetUserPasswordResponseBodyResetUsersResult {
	s.FailedUsers = v
	return s
}

func (s *ResetUserPasswordResponseBodyResetUsersResult) SetResetUsers(v []*string) *ResetUserPasswordResponseBodyResetUsersResult {
	s.ResetUsers = v
	return s
}

type ResetUserPasswordResponseBodyResetUsersResultFailedUsers struct {
	// The ID of the convenience user whose password failed to be reset.
	//
	// example:
	//
	// test123
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidUsername
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// test123 is an invalid username.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s ResetUserPasswordResponseBodyResetUsersResultFailedUsers) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordResponseBodyResetUsersResultFailedUsers) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBodyResetUsersResultFailedUsers) SetEndUserId(v string) *ResetUserPasswordResponseBodyResetUsersResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *ResetUserPasswordResponseBodyResetUsersResultFailedUsers) SetErrorCode(v string) *ResetUserPasswordResponseBodyResetUsersResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *ResetUserPasswordResponseBodyResetUsersResultFailedUsers) SetErrorMessage(v string) *ResetUserPasswordResponseBodyResetUsersResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

type ResetUserPasswordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponse) SetHeaders(v map[string]*string) *ResetUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetUserPasswordResponse) SetStatusCode(v int32) *ResetUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetUserPasswordResponse) SetBody(v *ResetUserPasswordResponseBody) *ResetUserPasswordResponse {
	s.Body = v
	return s
}

type SetUserPropertyValueRequest struct {
	// The property ID. You can call the [ListProperty](~~ListProperty~~) operation to query the property ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The ID of the property value. You can call the [ListProperty](~~ListProperty~~) operation to query the ID of the property value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
	// The ID of the convenience user. You can call the [DescribeUsers](~~DescribeUsers~~) operation to query the user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the convenience user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s SetUserPropertyValueRequest) String() string {
	return tea.Prettify(s)
}

func (s SetUserPropertyValueRequest) GoString() string {
	return s.String()
}

func (s *SetUserPropertyValueRequest) SetPropertyId(v int64) *SetUserPropertyValueRequest {
	s.PropertyId = &v
	return s
}

func (s *SetUserPropertyValueRequest) SetPropertyValueId(v int64) *SetUserPropertyValueRequest {
	s.PropertyValueId = &v
	return s
}

func (s *SetUserPropertyValueRequest) SetUserId(v int64) *SetUserPropertyValueRequest {
	s.UserId = &v
	return s
}

func (s *SetUserPropertyValueRequest) SetUserName(v string) *SetUserPropertyValueRequest {
	s.UserName = &v
	return s
}

type SetUserPropertyValueResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetUserPropertyValueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetUserPropertyValueResponseBody) GoString() string {
	return s.String()
}

func (s *SetUserPropertyValueResponseBody) SetRequestId(v string) *SetUserPropertyValueResponseBody {
	s.RequestId = &v
	return s
}

type SetUserPropertyValueResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetUserPropertyValueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetUserPropertyValueResponse) String() string {
	return tea.Prettify(s)
}

func (s SetUserPropertyValueResponse) GoString() string {
	return s.String()
}

func (s *SetUserPropertyValueResponse) SetHeaders(v map[string]*string) *SetUserPropertyValueResponse {
	s.Headers = v
	return s
}

func (s *SetUserPropertyValueResponse) SetStatusCode(v int32) *SetUserPropertyValueResponse {
	s.StatusCode = &v
	return s
}

func (s *SetUserPropertyValueResponse) SetBody(v *SetUserPropertyValueResponseBody) *SetUserPropertyValueResponse {
	s.Body = v
	return s
}

type SyncAllEduInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncAllEduInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SyncAllEduInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SyncAllEduInfoResponseBody) SetCode(v string) *SyncAllEduInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SyncAllEduInfoResponseBody) SetHttpStatusCode(v int32) *SyncAllEduInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SyncAllEduInfoResponseBody) SetMessage(v string) *SyncAllEduInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SyncAllEduInfoResponseBody) SetRequestId(v string) *SyncAllEduInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncAllEduInfoResponseBody) SetSuccess(v bool) *SyncAllEduInfoResponseBody {
	s.Success = &v
	return s
}

type SyncAllEduInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncAllEduInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncAllEduInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s SyncAllEduInfoResponse) GoString() string {
	return s.String()
}

func (s *SyncAllEduInfoResponse) SetHeaders(v map[string]*string) *SyncAllEduInfoResponse {
	s.Headers = v
	return s
}

func (s *SyncAllEduInfoResponse) SetStatusCode(v int32) *SyncAllEduInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncAllEduInfoResponse) SetBody(v *SyncAllEduInfoResponseBody) *SyncAllEduInfoResponse {
	s.Body = v
	return s
}

type UnlockMfaDeviceRequest struct {
	// The domain of the Active Directory (AD) workspace.
	//
	// example:
	//
	// welab.co.id
	AdDomain *string `json:"AdDomain,omitempty" xml:"AdDomain,omitempty"`
	// The serial number of the virtual MFA device. The serial number is unique for each device.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc856334-446b-4035-bfbc-18af261e****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s UnlockMfaDeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockMfaDeviceRequest) GoString() string {
	return s.String()
}

func (s *UnlockMfaDeviceRequest) SetAdDomain(v string) *UnlockMfaDeviceRequest {
	s.AdDomain = &v
	return s
}

func (s *UnlockMfaDeviceRequest) SetSerialNumber(v string) *UnlockMfaDeviceRequest {
	s.SerialNumber = &v
	return s
}

type UnlockMfaDeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9BD39C60-4E38-43BE-BA2F-69136C6C5190
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnlockMfaDeviceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnlockMfaDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockMfaDeviceResponseBody) SetRequestId(v string) *UnlockMfaDeviceResponseBody {
	s.RequestId = &v
	return s
}

type UnlockMfaDeviceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockMfaDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockMfaDeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockMfaDeviceResponse) GoString() string {
	return s.String()
}

func (s *UnlockMfaDeviceResponse) SetHeaders(v map[string]*string) *UnlockMfaDeviceResponse {
	s.Headers = v
	return s
}

func (s *UnlockMfaDeviceResponse) SetStatusCode(v int32) *UnlockMfaDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockMfaDeviceResponse) SetBody(v *UnlockMfaDeviceResponseBody) *UnlockMfaDeviceResponse {
	s.Body = v
	return s
}

type UnlockUsersRequest struct {
	// The date on which the convenience users are automatically locked.
	//
	// example:
	//
	// 2023-03-03
	AutoLockTime *string `json:"AutoLockTime,omitempty" xml:"AutoLockTime,omitempty"`
	// The usernames of the convenience users that you want to unlock.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s UnlockUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersRequest) GoString() string {
	return s.String()
}

func (s *UnlockUsersRequest) SetAutoLockTime(v string) *UnlockUsersRequest {
	s.AutoLockTime = &v
	return s
}

func (s *UnlockUsersRequest) SetUsers(v []*string) *UnlockUsersRequest {
	s.Users = v
	return s
}

type UnlockUsersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 023F4EC4-3602-4A3E-A514-4970847D59DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of unlocking the convenience user.
	UnlockUsersResult *UnlockUsersResponseBodyUnlockUsersResult `json:"UnlockUsersResult,omitempty" xml:"UnlockUsersResult,omitempty" type:"Struct"`
}

func (s UnlockUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersResponseBody) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBody) SetRequestId(v string) *UnlockUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnlockUsersResponseBody) SetUnlockUsersResult(v *UnlockUsersResponseBodyUnlockUsersResult) *UnlockUsersResponseBody {
	s.UnlockUsersResult = v
	return s
}

type UnlockUsersResponseBodyUnlockUsersResult struct {
	// The convenience users that failed to be unlocked.
	FailedUsers []*UnlockUsersResponseBodyUnlockUsersResultFailedUsers `json:"FailedUsers,omitempty" xml:"FailedUsers,omitempty" type:"Repeated"`
	// The convenience users that were unlocked.
	UnlockedUsers []*string `json:"UnlockedUsers,omitempty" xml:"UnlockedUsers,omitempty" type:"Repeated"`
}

func (s UnlockUsersResponseBodyUnlockUsersResult) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersResponseBodyUnlockUsersResult) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBodyUnlockUsersResult) SetFailedUsers(v []*UnlockUsersResponseBodyUnlockUsersResultFailedUsers) *UnlockUsersResponseBodyUnlockUsersResult {
	s.FailedUsers = v
	return s
}

func (s *UnlockUsersResponseBodyUnlockUsersResult) SetUnlockedUsers(v []*string) *UnlockUsersResponseBodyUnlockUsersResult {
	s.UnlockedUsers = v
	return s
}

type UnlockUsersResponseBodyUnlockUsersResultFailedUsers struct {
	// The ID of the convenience user that failed to be unlocked.
	//
	// example:
	//
	// test123
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The error code.
	//
	// example:
	//
	// InvalidUsername
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// test123 is an invalid username.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s UnlockUsersResponseBodyUnlockUsersResultFailedUsers) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersResponseBodyUnlockUsersResultFailedUsers) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponseBodyUnlockUsersResultFailedUsers) SetEndUserId(v string) *UnlockUsersResponseBodyUnlockUsersResultFailedUsers {
	s.EndUserId = &v
	return s
}

func (s *UnlockUsersResponseBodyUnlockUsersResultFailedUsers) SetErrorCode(v string) *UnlockUsersResponseBodyUnlockUsersResultFailedUsers {
	s.ErrorCode = &v
	return s
}

func (s *UnlockUsersResponseBodyUnlockUsersResultFailedUsers) SetErrorMessage(v string) *UnlockUsersResponseBodyUnlockUsersResultFailedUsers {
	s.ErrorMessage = &v
	return s
}

type UnlockUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnlockUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnlockUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s UnlockUsersResponse) GoString() string {
	return s.String()
}

func (s *UnlockUsersResponse) SetHeaders(v map[string]*string) *UnlockUsersResponse {
	s.Headers = v
	return s
}

func (s *UnlockUsersResponse) SetStatusCode(v int32) *UnlockUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *UnlockUsersResponse) SetBody(v *UnlockUsersResponseBody) *UnlockUsersResponse {
	s.Body = v
	return s
}

type UpdatePropertyRequest struct {
	// The ID of the property that you want to modify. You can call the [ListProperty](https://help.aliyun.com/document_detail/410890.html) operation to query the property ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The new property name.
	//
	// This parameter is required.
	//
	// example:
	//
	// testkey
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The values of property.
	PropertyValues []*UpdatePropertyRequestPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
}

func (s UpdatePropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePropertyRequest) SetPropertyId(v int64) *UpdatePropertyRequest {
	s.PropertyId = &v
	return s
}

func (s *UpdatePropertyRequest) SetPropertyKey(v string) *UpdatePropertyRequest {
	s.PropertyKey = &v
	return s
}

func (s *UpdatePropertyRequest) SetPropertyValues(v []*UpdatePropertyRequestPropertyValues) *UpdatePropertyRequest {
	s.PropertyValues = v
	return s
}

type UpdatePropertyRequestPropertyValues struct {
	// The new property value.
	//
	// example:
	//
	// testvalue
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The ID of property value that you want to modify. You can call the [ListProperty](https://help.aliyun.com/document_detail/410890.html) operation to query the property value ID.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s UpdatePropertyRequestPropertyValues) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyRequestPropertyValues) GoString() string {
	return s.String()
}

func (s *UpdatePropertyRequestPropertyValues) SetPropertyValue(v string) *UpdatePropertyRequestPropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *UpdatePropertyRequestPropertyValues) SetPropertyValueId(v int64) *UpdatePropertyRequestPropertyValues {
	s.PropertyValueId = &v
	return s
}

type UpdatePropertyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the modification.
	UpdateResult *UpdatePropertyResponseBodyUpdateResult `json:"UpdateResult,omitempty" xml:"UpdateResult,omitempty" type:"Struct"`
}

func (s UpdatePropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBody) SetRequestId(v string) *UpdatePropertyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePropertyResponseBody) SetUpdateResult(v *UpdatePropertyResponseBodyUpdateResult) *UpdatePropertyResponseBody {
	s.UpdateResult = v
	return s
}

type UpdatePropertyResponseBodyUpdateResult struct {
	// The ID of the property.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The name of the property.
	//
	// example:
	//
	// testkey2
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The result of the property value modification.
	SavePropertyValueModel *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel `json:"SavePropertyValueModel,omitempty" xml:"SavePropertyValueModel,omitempty" type:"Struct"`
}

func (s UpdatePropertyResponseBodyUpdateResult) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyResponseBodyUpdateResult) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBodyUpdateResult) SetPropertyId(v int64) *UpdatePropertyResponseBodyUpdateResult {
	s.PropertyId = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResult) SetPropertyKey(v string) *UpdatePropertyResponseBodyUpdateResult {
	s.PropertyKey = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResult) SetSavePropertyValueModel(v *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) *UpdatePropertyResponseBodyUpdateResult {
	s.SavePropertyValueModel = v
	return s
}

type UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel struct {
	// The property values that failed to be modified.
	FailedPropertyValues []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues `json:"FailedPropertyValues,omitempty" xml:"FailedPropertyValues,omitempty" type:"Repeated"`
	// The property values that were modified.
	SavePropertyValues []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues `json:"SavePropertyValues,omitempty" xml:"SavePropertyValues,omitempty" type:"Repeated"`
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) SetFailedPropertyValues(v []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel {
	s.FailedPropertyValues = v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel) SetSavePropertyValues(v []*UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModel {
	s.SavePropertyValues = v
	return s
}

type UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues struct {
	// The error code.
	//
	// example:
	//
	// ExistedPropertyValue
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The property value is used by another property.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the property.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The value of the property.
	//
	// example:
	//
	// testvalue
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) SetErrorCode(v string) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) SetErrorMessage(v string) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) SetPropertyId(v int64) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues {
	s.PropertyId = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues) SetPropertyValue(v string) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelFailedPropertyValues {
	s.PropertyValue = &v
	return s
}

type UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues struct {
	// The value of the property.
	//
	// example:
	//
	// testvalue2
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The ID of the property value.
	//
	// example:
	//
	// 978
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) SetPropertyValue(v string) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues) SetPropertyValueId(v int64) *UpdatePropertyResponseBodyUpdateResultSavePropertyValueModelSavePropertyValues {
	s.PropertyValueId = &v
	return s
}

type UpdatePropertyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePropertyResponse) GoString() string {
	return s.String()
}

func (s *UpdatePropertyResponse) SetHeaders(v map[string]*string) *UpdatePropertyResponse {
	s.Headers = v
	return s
}

func (s *UpdatePropertyResponse) SetStatusCode(v int32) *UpdatePropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePropertyResponse) SetBody(v *UpdatePropertyResponseBody) *UpdatePropertyResponse {
	s.Body = v
	return s
}

type UserBatchJoinGroupRequest struct {
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// ug-12341234****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s UserBatchJoinGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UserBatchJoinGroupRequest) GoString() string {
	return s.String()
}

func (s *UserBatchJoinGroupRequest) SetEndUserIds(v []*string) *UserBatchJoinGroupRequest {
	s.EndUserIds = v
	return s
}

func (s *UserBatchJoinGroupRequest) SetGroupId(v string) *UserBatchJoinGroupRequest {
	s.GroupId = &v
	return s
}

type UserBatchJoinGroupResponseBody struct {
	// example:
	//
	// 7A2C3803-C975-5871-A232-80A91009****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UserBatchJoinGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UserBatchJoinGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UserBatchJoinGroupResponseBody) SetRequestId(v string) *UserBatchJoinGroupResponseBody {
	s.RequestId = &v
	return s
}

type UserBatchJoinGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UserBatchJoinGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UserBatchJoinGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UserBatchJoinGroupResponse) GoString() string {
	return s.String()
}

func (s *UserBatchJoinGroupResponse) SetHeaders(v map[string]*string) *UserBatchJoinGroupResponse {
	s.Headers = v
	return s
}

func (s *UserBatchJoinGroupResponse) SetStatusCode(v int32) *UserBatchJoinGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UserBatchJoinGroupResponse) SetBody(v *UserBatchJoinGroupResponseBody) *UserBatchJoinGroupResponse {
	s.Body = v
	return s
}

type UserBatchQuitGroupRequest struct {
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// example:
	//
	// ug-lkuvalovhnlxvv****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s UserBatchQuitGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UserBatchQuitGroupRequest) GoString() string {
	return s.String()
}

func (s *UserBatchQuitGroupRequest) SetEndUserIds(v []*string) *UserBatchQuitGroupRequest {
	s.EndUserIds = v
	return s
}

func (s *UserBatchQuitGroupRequest) SetGroupId(v string) *UserBatchQuitGroupRequest {
	s.GroupId = &v
	return s
}

type UserBatchQuitGroupResponseBody struct {
	// example:
	//
	// 890JNJNF-SADASSDFS-SDFSDF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UserBatchQuitGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UserBatchQuitGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UserBatchQuitGroupResponseBody) SetRequestId(v string) *UserBatchQuitGroupResponseBody {
	s.RequestId = &v
	return s
}

type UserBatchQuitGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UserBatchQuitGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UserBatchQuitGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UserBatchQuitGroupResponse) GoString() string {
	return s.String()
}

func (s *UserBatchQuitGroupResponse) SetHeaders(v map[string]*string) *UserBatchQuitGroupResponse {
	s.Headers = v
	return s
}

func (s *UserBatchQuitGroupResponse) SetStatusCode(v int32) *UserBatchQuitGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UserBatchQuitGroupResponse) SetBody(v *UserBatchQuitGroupResponseBody) *UserBatchQuitGroupResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eds-user"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 批量设置桌面管理员
//
// @param request - BatchSetDesktopManagerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSetDesktopManagerResponse
func (client *Client) BatchSetDesktopManagerWithOptions(request *BatchSetDesktopManagerRequest, runtime *util.RuntimeOptions) (_result *BatchSetDesktopManagerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsDesktopManager)) {
		body["IsDesktopManager"] = request.IsDesktopManager
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchSetDesktopManager"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchSetDesktopManagerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量设置桌面管理员
//
// @param request - BatchSetDesktopManagerRequest
//
// @return BatchSetDesktopManagerResponse
func (client *Client) BatchSetDesktopManager(request *BatchSetDesktopManagerRequest) (_result *BatchSetDesktopManagerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BatchSetDesktopManagerResponse{}
	_body, _err := client.BatchSetDesktopManagerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 管理员修改用户密码
//
// @param request - ChangeUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeUserPasswordResponse
func (client *Client) ChangeUserPasswordWithOptions(request *ChangeUserPasswordRequest, runtime *util.RuntimeOptions) (_result *ChangeUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		body["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.NewPassword)) {
		body["NewPassword"] = request.NewPassword
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeUserPassword"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 管理员修改用户密码
//
// @param request - ChangeUserPasswordRequest
//
// @return ChangeUserPasswordResponse
func (client *Client) ChangeUserPassword(request *ChangeUserPasswordRequest) (_result *ChangeUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeUserPasswordResponse{}
	_body, _err := client.ChangeUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries whether a property is associated with one or more convenience users.
//
// @param request - CheckUsedPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckUsedPropertyResponse
func (client *Client) CheckUsedPropertyWithOptions(request *CheckUsedPropertyRequest, runtime *util.RuntimeOptions) (_result *CheckUsedPropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		query["PropertyId"] = request.PropertyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckUsedProperty"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUsedPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries whether a property is associated with one or more convenience users.
//
// @param request - CheckUsedPropertyRequest
//
// @return CheckUsedPropertyResponse
func (client *Client) CheckUsedProperty(request *CheckUsedPropertyRequest) (_result *CheckUsedPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUsedPropertyResponse{}
	_body, _err := client.CheckUsedPropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the number of convenience accounts that are associated with the specified custom property value.
//
// Description:
//
// Before you call the operation, you can call the [ListProperty](https://help.aliyun.com/document_detail/410890.html) operation to query the existing user properties and their IDs (PropertyId) and values (PropertyValueId).
//
// @param request - CheckUsedPropertyValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckUsedPropertyValueResponse
func (client *Client) CheckUsedPropertyValueWithOptions(request *CheckUsedPropertyValueRequest, runtime *util.RuntimeOptions) (_result *CheckUsedPropertyValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		query["PropertyId"] = request.PropertyId
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyValueId)) {
		query["PropertyValueId"] = request.PropertyValueId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckUsedPropertyValue"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckUsedPropertyValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the number of convenience accounts that are associated with the specified custom property value.
//
// Description:
//
// Before you call the operation, you can call the [ListProperty](https://help.aliyun.com/document_detail/410890.html) operation to query the existing user properties and their IDs (PropertyId) and values (PropertyValueId).
//
// @param request - CheckUsedPropertyValueRequest
//
// @return CheckUsedPropertyValueResponse
func (client *Client) CheckUsedPropertyValue(request *CheckUsedPropertyValueRequest) (_result *CheckUsedPropertyValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckUsedPropertyValueResponse{}
	_body, _err := client.CheckUsedPropertyValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建角色.
//
// @param request - CreateGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentGroupId)) {
		query["ParentGroupId"] = request.ParentGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionId)) {
		query["SolutionId"] = request.SolutionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建角色.
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建组织
//
// @param request - CreateOrgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrgResponse
func (client *Client) CreateOrgWithOptions(request *CreateOrgRequest, runtime *util.RuntimeOptions) (_result *CreateOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		query["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentOrgId)) {
		query["ParentOrgId"] = request.ParentOrgId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrg"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建组织
//
// @param request - CreateOrgRequest
//
// @return CreateOrgResponse
func (client *Client) CreateOrg(request *CreateOrgRequest) (_result *CreateOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOrgResponse{}
	_body, _err := client.CreateOrgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user property.
//
// @param request - CreatePropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePropertyResponse
func (client *Client) CreatePropertyWithOptions(request *CreatePropertyRequest, runtime *util.RuntimeOptions) (_result *CreatePropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyKey)) {
		body["PropertyKey"] = request.PropertyKey
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyValues)) {
		body["PropertyValues"] = request.PropertyValues
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProperty"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user property.
//
// @param request - CreatePropertyRequest
//
// @return CreatePropertyResponse
func (client *Client) CreateProperty(request *CreatePropertyRequest) (_result *CreatePropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePropertyResponse{}
	_body, _err := client.CreatePropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建资源组
//
// @param request - CreateResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroupWithOptions(request *CreateResourceGroupRequest, runtime *util.RuntimeOptions) (_result *CreateResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsResourceGroupWithOfficeSite)) {
		query["IsResourceGroupWithOfficeSite"] = request.IsResourceGroupWithOfficeSite
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceGroup"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建资源组
//
// @param request - CreateResourceGroupRequest
//
// @return CreateResourceGroupResponse
func (client *Client) CreateResourceGroup(request *CreateResourceGroupRequest) (_result *CreateResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateResourceGroupResponse{}
	_body, _err := client.CreateResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a convenience user.
//
// Description:
//
// Convenience users are dedicated Elastic Desktop Service (EDS) user accounts and are suitable for scenarios in which you do not need to connect to enterprise Active Directory (AD) systems. The information about a convenience user includes the username, email address, and mobile number. You must specify the username or email address.
//
// @param request - CreateUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUsersResponse
func (client *Client) CreateUsersWithOptions(request *CreateUsersRequest, runtime *util.RuntimeOptions) (_result *CreateUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoLockTime)) {
		query["AutoLockTime"] = request.AutoLockTime
	}

	if !tea.BoolValue(util.IsUnset(request.IsLocalAdmin)) {
		query["IsLocalAdmin"] = request.IsLocalAdmin
	}

	if !tea.BoolValue(util.IsUnset(request.PasswordExpireDays)) {
		query["PasswordExpireDays"] = request.PasswordExpireDays
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUsers"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a convenience user.
//
// Description:
//
// Convenience users are dedicated Elastic Desktop Service (EDS) user accounts and are suitable for scenarios in which you do not need to connect to enterprise Active Directory (AD) systems. The information about a convenience user includes the username, email address, and mobile number. You must specify the username or email address.
//
// @param request - CreateUsersRequest
//
// @return CreateUsersResponse
func (client *Client) CreateUsers(request *CreateUsersRequest) (_result *CreateUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUsersResponse{}
	_body, _err := client.CreateUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a resource group.
//
// @param request - DeleteResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroupWithOptions(request *DeleteResourceGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIds)) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceGroup"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a resource group.
//
// @param request - DeleteResourceGroupRequest
//
// @return DeleteResourceGroupResponse
func (client *Client) DeleteResourceGroup(request *DeleteResourceGroupRequest) (_result *DeleteResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteResourceGroupResponse{}
	_body, _err := client.DeleteResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Dissociates a user property from a user.
//
// Description:
//
// Before you call this operation, you can call the FilterUsers operation to query the users that are associated with user properties.
//
// @param request - DeleteUserPropertyValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserPropertyValueResponse
func (client *Client) DeleteUserPropertyValueWithOptions(request *DeleteUserPropertyValueRequest, runtime *util.RuntimeOptions) (_result *DeleteUserPropertyValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		body["PropertyId"] = request.PropertyId
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyValueId)) {
		body["PropertyValueId"] = request.PropertyValueId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUserPropertyValue"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserPropertyValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Dissociates a user property from a user.
//
// Description:
//
// Before you call this operation, you can call the FilterUsers operation to query the users that are associated with user properties.
//
// @param request - DeleteUserPropertyValueRequest
//
// @return DeleteUserPropertyValueResponse
func (client *Client) DeleteUserPropertyValue(request *DeleteUserPropertyValueRequest) (_result *DeleteUserPropertyValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserPropertyValueResponse{}
	_body, _err := client.DeleteUserPropertyValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 全量同步初始化
//
// @param request - DescribeGroupUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupUserResponse
func (client *Client) DescribeGroupUserWithOptions(request *DescribeGroupUserRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionId)) {
		query["SolutionId"] = request.SolutionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGroupUser"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGroupUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 全量同步初始化
//
// @param request - DescribeGroupUserRequest
//
// @return DescribeGroupUserResponse
func (client *Client) DescribeGroupUser(request *DescribeGroupUserRequest) (_result *DescribeGroupUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupUserResponse{}
	_body, _err := client.DescribeGroupUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 全量同步初始化.
//
// @param request - DescribeGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeGroupsResponse
func (client *Client) DescribeGroupsWithOptions(request *DescribeGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		query["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		query["GroupName"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionId)) {
		query["SolutionId"] = request.SolutionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeGroups"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 全量同步初始化.
//
// @param request - DescribeGroupsRequest
//
// @return DescribeGroupsResponse
func (client *Client) DescribeGroups(request *DescribeGroupsRequest) (_result *DescribeGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGroupsResponse{}
	_body, _err := client.DescribeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about virtual multi-factor authentication (MFA) devices that are bound to convenience users.
//
// @param request - DescribeMfaDevicesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMfaDevicesResponse
func (client *Client) DescribeMfaDevicesWithOptions(request *DescribeMfaDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeMfaDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdDomain)) {
		query["AdDomain"] = request.AdDomain
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserIds)) {
		query["EndUserIds"] = request.EndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumbers)) {
		query["SerialNumbers"] = request.SerialNumbers
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMfaDevices"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMfaDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about virtual multi-factor authentication (MFA) devices that are bound to convenience users.
//
// @param request - DescribeMfaDevicesRequest
//
// @return DescribeMfaDevicesResponse
func (client *Client) DescribeMfaDevices(request *DescribeMfaDevicesRequest) (_result *DescribeMfaDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMfaDevicesResponse{}
	_body, _err := client.DescribeMfaDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查找下级组织
//
// @param request - DescribeOrgByLayerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOrgByLayerResponse
func (client *Client) DescribeOrgByLayerWithOptions(request *DescribeOrgByLayerRequest, runtime *util.RuntimeOptions) (_result *DescribeOrgByLayerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		body["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentOrgId)) {
		body["ParentOrgId"] = request.ParentOrgId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOrgByLayer"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOrgByLayerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查找下级组织
//
// @param request - DescribeOrgByLayerRequest
//
// @return DescribeOrgByLayerResponse
func (client *Client) DescribeOrgByLayer(request *DescribeOrgByLayerRequest) (_result *DescribeOrgByLayerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOrgByLayerResponse{}
	_body, _err := client.DescribeOrgByLayerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries organizations.
//
// Description:
//
// An organization is in a tree structure. The root organization ID is in the following format: org-aliyun-wy-org-id.
//
// @param request - DescribeOrgsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeOrgsResponse
func (client *Client) DescribeOrgsWithOptions(request *DescribeOrgsRequest, runtime *util.RuntimeOptions) (_result *DescribeOrgsResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		query["OrgName"] = request.OrgName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentOrgId)) {
		query["ParentOrgId"] = request.ParentOrgId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOrgs"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOrgsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries organizations.
//
// Description:
//
// An organization is in a tree structure. The root organization ID is in the following format: org-aliyun-wy-org-id.
//
// @param request - DescribeOrgsRequest
//
// @return DescribeOrgsResponse
func (client *Client) DescribeOrgs(request *DescribeOrgsRequest) (_result *DescribeOrgsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOrgsResponse{}
	_body, _err := client.DescribeOrgsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看资源组
//
// @param request - DescribeResourceGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeResourceGroupsResponse
func (client *Client) DescribeResourceGroupsWithOptions(request *DescribeResourceGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedContainResourceGroupWithOfficeSite)) {
		query["NeedContainResourceGroupWithOfficeSite"] = request.NeedContainResourceGroupWithOfficeSite
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupIds)) {
		query["ResourceGroupIds"] = request.ResourceGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupName)) {
		query["ResourceGroupName"] = request.ResourceGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceGroups"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看资源组
//
// @param request - DescribeResourceGroupsRequest
//
// @return DescribeResourceGroupsResponse
func (client *Client) DescribeResourceGroups(request *DescribeResourceGroupsRequest) (_result *DescribeResourceGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceGroupsResponse{}
	_body, _err := client.DescribeResourceGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about convenience users. The information of a convenience user includes a username, an email address, and a description.
//
// @param tmpReq - DescribeUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsersResponse
func (client *Client) DescribeUsersWithOptions(tmpReq *DescribeUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeUsersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FilterWithAssignedResource)) {
		request.FilterWithAssignedResourceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterWithAssignedResource, tea.String("FilterWithAssignedResource"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.FilterWithAssignedResources)) {
		request.FilterWithAssignedResourcesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterWithAssignedResources, tea.String("FilterWithAssignedResources"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.ShowExtras)) {
		request.ShowExtrasShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ShowExtras, tea.String("ShowExtras"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.FilterWithAssignedResourceShrink)) {
		query["FilterWithAssignedResource"] = request.FilterWithAssignedResourceShrink
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

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizType)) {
		body["BizType"] = request.BizType
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserIds)) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.ExcludeEndUserIds)) {
		body["ExcludeEndUserIds"] = request.ExcludeEndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.FilterWithAssignedResourcesShrink)) {
		body["FilterWithAssignedResources"] = request.FilterWithAssignedResourcesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.IsQueryAllSubOrgs)) {
		body["IsQueryAllSubOrgs"] = request.IsQueryAllSubOrgs
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowExtrasShrink)) {
		body["ShowExtras"] = request.ShowExtrasShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SolutionId)) {
		body["SolutionId"] = request.SolutionId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsers"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about convenience users. The information of a convenience user includes a username, an email address, and a description.
//
// @param request - DescribeUsersRequest
//
// @return DescribeUsersResponse
func (client *Client) DescribeUsers(request *DescribeUsersRequest) (_result *DescribeUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUsersResponse{}
	_body, _err := client.DescribeUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Filters convenience accounts by property.
//
// @param tmpReq - FilterUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FilterUsersResponse
func (client *Client) FilterUsersWithOptions(tmpReq *FilterUsersRequest, runtime *util.RuntimeOptions) (_result *FilterUsersResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &FilterUsersShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.OrderParam)) {
		request.OrderParamShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.OrderParam, tea.String("OrderParam"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExcludeEndUserIds)) {
		query["ExcludeEndUserIds"] = request.ExcludeEndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeDesktopCount)) {
		query["IncludeDesktopCount"] = request.IncludeDesktopCount
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeDesktopGroupCount)) {
		query["IncludeDesktopGroupCount"] = request.IncludeDesktopGroupCount
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeOrgInfo)) {
		query["IncludeOrgInfo"] = request.IncludeOrgInfo
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeSupportIdps)) {
		query["IncludeSupportIdps"] = request.IncludeSupportIdps
	}

	if !tea.BoolValue(util.IsUnset(request.IsQueryAllSubOrgs)) {
		query["IsQueryAllSubOrgs"] = request.IsQueryAllSubOrgs
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OrderParamShrink)) {
		query["OrderParam"] = request.OrderParamShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerType)) {
		query["OwnerType"] = request.OwnerType
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyFilterParam)) {
		query["PropertyFilterParam"] = request.PropertyFilterParam
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyKeyValueFilterParam)) {
		query["PropertyKeyValueFilterParam"] = request.PropertyKeyValueFilterParam
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FilterUsers"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FilterUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Filters convenience accounts by property.
//
// @param request - FilterUsersRequest
//
// @return FilterUsersResponse
func (client *Client) FilterUsers(request *FilterUsersRequest) (_result *FilterUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FilterUsersResponse{}
	_body, _err := client.FilterUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the information about the current logon administrator based on the authorization code.
//
// @param request - GetManagerInfoByAuthCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetManagerInfoByAuthCodeResponse
func (client *Client) GetManagerInfoByAuthCodeWithOptions(request *GetManagerInfoByAuthCodeRequest, runtime *util.RuntimeOptions) (_result *GetManagerInfoByAuthCodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthCode)) {
		query["AuthCode"] = request.AuthCode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetManagerInfoByAuthCode"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetManagerInfoByAuthCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the information about the current logon administrator based on the authorization code.
//
// @param request - GetManagerInfoByAuthCodeRequest
//
// @return GetManagerInfoByAuthCodeResponse
func (client *Client) GetManagerInfoByAuthCode(request *GetManagerInfoByAuthCodeRequest) (_result *GetManagerInfoByAuthCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetManagerInfoByAuthCodeResponse{}
	_body, _err := client.GetManagerInfoByAuthCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 初始化TenantAlias
//
// @param request - InitTenantAliasRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InitTenantAliasResponse
func (client *Client) InitTenantAliasWithOptions(runtime *util.RuntimeOptions) (_result *InitTenantAliasResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("InitTenantAlias"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InitTenantAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 初始化TenantAlias
//
// @return InitTenantAliasResponse
func (client *Client) InitTenantAlias() (_result *InitTenantAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitTenantAliasResponse{}
	_body, _err := client.InitTenantAliasWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries all user properties within an Alibaba Cloud account.
//
// @param request - ListPropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPropertyResponse
func (client *Client) ListPropertyWithOptions(runtime *util.RuntimeOptions) (_result *ListPropertyResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListProperty"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries all user properties within an Alibaba Cloud account.
//
// @return ListPropertyResponse
func (client *Client) ListProperty() (_result *ListPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPropertyResponse{}
	_body, _err := client.ListPropertyWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries property values of a user property.
//
// @param request - ListPropertyValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPropertyValueResponse
func (client *Client) ListPropertyValueWithOptions(request *ListPropertyValueRequest, runtime *util.RuntimeOptions) (_result *ListPropertyValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		query["PropertyId"] = request.PropertyId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPropertyValue"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPropertyValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries property values of a user property.
//
// @param request - ListPropertyValueRequest
//
// @return ListPropertyValueResponse
func (client *Client) ListPropertyValue(request *ListPropertyValueRequest) (_result *ListPropertyValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPropertyValueResponse{}
	_body, _err := client.ListPropertyValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Locks a virtual multi-factor authentication (MFA) device that is bound to a convenience user.
//
// Description:
//
// After a virtual MFA device is locked, the status of the virtual MFA device changes to LOCKED. The convenience user to which the MFA device is bound cannot log on to the cloud desktop that resides in the workspace with the MFA feature enabled because the identity of the convenience user cannot be verified based on the virtual MFA device. You can call the [UnlockMfaDevice](https://help.aliyun.com/document_detail/286534.html) operation to unlock the virtual MFA device.
//
// @param request - LockMfaDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockMfaDeviceResponse
func (client *Client) LockMfaDeviceWithOptions(request *LockMfaDeviceRequest, runtime *util.RuntimeOptions) (_result *LockMfaDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdDomain)) {
		query["AdDomain"] = request.AdDomain
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("LockMfaDevice"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LockMfaDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Locks a virtual multi-factor authentication (MFA) device that is bound to a convenience user.
//
// Description:
//
// After a virtual MFA device is locked, the status of the virtual MFA device changes to LOCKED. The convenience user to which the MFA device is bound cannot log on to the cloud desktop that resides in the workspace with the MFA feature enabled because the identity of the convenience user cannot be verified based on the virtual MFA device. You can call the [UnlockMfaDevice](https://help.aliyun.com/document_detail/286534.html) operation to unlock the virtual MFA device.
//
// @param request - LockMfaDeviceRequest
//
// @return LockMfaDeviceResponse
func (client *Client) LockMfaDevice(request *LockMfaDeviceRequest) (_result *LockMfaDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockMfaDeviceResponse{}
	_body, _err := client.LockMfaDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Locks one or more convenience users.
//
// @param request - LockUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LockUsersResponse
func (client *Client) LockUsersWithOptions(request *LockUsersRequest, runtime *util.RuntimeOptions) (_result *LockUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogoutSession)) {
		query["LogoutSession"] = request.LogoutSession
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LockUsers"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &LockUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Locks one or more convenience users.
//
// @param request - LockUsersRequest
//
// @return LockUsersResponse
func (client *Client) LockUsers(request *LockUsersRequest) (_result *LockUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LockUsersResponse{}
	_body, _err := client.LockUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改角色.
//
// @param request - ModifyGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyGroupResponse
func (client *Client) ModifyGroupWithOptions(request *ModifyGroupRequest, runtime *util.RuntimeOptions) (_result *ModifyGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.NewGroupName)) {
		query["NewGroupName"] = request.NewGroupName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyGroup"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改角色.
//
// @param request - ModifyGroupRequest
//
// @return ModifyGroupResponse
func (client *Client) ModifyGroup(request *ModifyGroupRequest) (_result *ModifyGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyGroupResponse{}
	_body, _err := client.ModifyGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改组织
//
// @param request - ModifyOrgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyOrgResponse
func (client *Client) ModifyOrgWithOptions(request *ModifyOrgRequest, runtime *util.RuntimeOptions) (_result *ModifyOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		query["OrgId"] = request.OrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgName)) {
		query["OrgName"] = request.OrgName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyOrg"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyOrgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改组织
//
// @param request - ModifyOrgRequest
//
// @return ModifyOrgResponse
func (client *Client) ModifyOrg(request *ModifyOrgRequest) (_result *ModifyOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyOrgResponse{}
	_body, _err := client.ModifyOrgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies user information.
//
// @param request - ModifyUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyUserResponse
func (client *Client) ModifyUserWithOptions(request *ModifyUserRequest, runtime *util.RuntimeOptions) (_result *ModifyUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		query["Email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.EndUserId)) {
		query["EndUserId"] = request.EndUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		query["Phone"] = request.Phone
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyUser"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies user information.
//
// @param request - ModifyUserRequest
//
// @return ModifyUserResponse
func (client *Client) ModifyUser(request *ModifyUserRequest) (_result *ModifyUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserResponse{}
	_body, _err := client.ModifyUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动组织
//
// @param request - MoveOrgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveOrgResponse
func (client *Client) MoveOrgWithOptions(request *MoveOrgRequest, runtime *util.RuntimeOptions) (_result *MoveOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NewParentOrgId)) {
		body["NewParentOrgId"] = request.NewParentOrgId
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveOrg"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveOrgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动组织
//
// @param request - MoveOrgRequest
//
// @return MoveOrgResponse
func (client *Client) MoveOrg(request *MoveOrgRequest) (_result *MoveOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveOrgResponse{}
	_body, _err := client.MoveOrgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移动用户组织
//
// @param request - MoveUserOrgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveUserOrgResponse
func (client *Client) MoveUserOrgWithOptions(request *MoveUserOrgRequest, runtime *util.RuntimeOptions) (_result *MoveUserOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserIds)) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveUserOrg"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveUserOrgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移动用户组织
//
// @param request - MoveUserOrgRequest
//
// @return MoveUserOrgResponse
func (client *Client) MoveUserOrg(request *MoveUserOrgRequest) (_result *MoveUserOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveUserOrgResponse{}
	_body, _err := client.MoveUserOrgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询edu同步信息
//
// @param request - QuerySyncStatusByAliUidRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySyncStatusByAliUidResponse
func (client *Client) QuerySyncStatusByAliUidWithOptions(runtime *util.RuntimeOptions) (_result *QuerySyncStatusByAliUidResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("QuerySyncStatusByAliUid"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &QuerySyncStatusByAliUidResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询edu同步信息
//
// @return QuerySyncStatusByAliUidResponse
func (client *Client) QuerySyncStatusByAliUid() (_result *QuerySyncStatusByAliUidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QuerySyncStatusByAliUidResponse{}
	_body, _err := client.QuerySyncStatusByAliUidWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除角色.
//
// @param request - RemoveGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveGroupResponse
func (client *Client) RemoveGroupWithOptions(request *RemoveGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		query["GroupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveGroup"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除角色.
//
// @param request - RemoveGroupRequest
//
// @return RemoveGroupResponse
func (client *Client) RemoveGroup(request *RemoveGroupRequest) (_result *RemoveGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveGroupResponse{}
	_body, _err := client.RemoveGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a virtual multi-factor authentication (MFA) device that is bound to a convenience user.
//
// Description:
//
// If you remove a virtual MFA device that is bound to a convenience user, the convenience user can no longer use the virtual MFA device to log on to cloud desktops. Before the convenience user can log on to cloud desktops again, a new virtual MFA device must be bound to the convenience user.
//
// @param request - RemoveMfaDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveMfaDeviceResponse
func (client *Client) RemoveMfaDeviceWithOptions(request *RemoveMfaDeviceRequest, runtime *util.RuntimeOptions) (_result *RemoveMfaDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdDomain)) {
		query["AdDomain"] = request.AdDomain
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveMfaDevice"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveMfaDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a virtual multi-factor authentication (MFA) device that is bound to a convenience user.
//
// Description:
//
// If you remove a virtual MFA device that is bound to a convenience user, the convenience user can no longer use the virtual MFA device to log on to cloud desktops. Before the convenience user can log on to cloud desktops again, a new virtual MFA device must be bound to the convenience user.
//
// @param request - RemoveMfaDeviceRequest
//
// @return RemoveMfaDeviceResponse
func (client *Client) RemoveMfaDevice(request *RemoveMfaDeviceRequest) (_result *RemoveMfaDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveMfaDeviceResponse{}
	_body, _err := client.RemoveMfaDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 移除组织
//
// @param request - RemoveOrgRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveOrgResponse
func (client *Client) RemoveOrgWithOptions(request *RemoveOrgRequest, runtime *util.RuntimeOptions) (_result *RemoveOrgResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OrgId)) {
		body["OrgId"] = request.OrgId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveOrg"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveOrgResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 移除组织
//
// @param request - RemoveOrgRequest
//
// @return RemoveOrgResponse
func (client *Client) RemoveOrg(request *RemoveOrgRequest) (_result *RemoveOrgResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveOrgResponse{}
	_body, _err := client.RemoveOrgWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user property.
//
// @param request - RemovePropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemovePropertyResponse
func (client *Client) RemovePropertyWithOptions(request *RemovePropertyRequest, runtime *util.RuntimeOptions) (_result *RemovePropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		body["PropertyId"] = request.PropertyId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveProperty"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemovePropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user property.
//
// @param request - RemovePropertyRequest
//
// @return RemovePropertyResponse
func (client *Client) RemoveProperty(request *RemovePropertyRequest) (_result *RemovePropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemovePropertyResponse{}
	_body, _err := client.RemovePropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes one or more convenience users.
//
// @param request - RemoveUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersResponse
func (client *Client) RemoveUsersWithOptions(request *RemoveUsersRequest, runtime *util.RuntimeOptions) (_result *RemoveUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUsers"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes one or more convenience users.
//
// @param request - RemoveUsersRequest
//
// @return RemoveUsersResponse
func (client *Client) RemoveUsers(request *RemoveUsersRequest) (_result *RemoveUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUsersResponse{}
	_body, _err := client.RemoveUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resets the password for a convenience user. If you call this operation, a token that is used to reset the password is generated, and the system sends a password reset email that includes the token to the email address of the convenience user.
//
// @param request - ResetUserPasswordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetUserPasswordResponse
func (client *Client) ResetUserPasswordWithOptions(request *ResetUserPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NotifyType)) {
		body["NotifyType"] = request.NotifyType
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetUserPassword"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resets the password for a convenience user. If you call this operation, a token that is used to reset the password is generated, and the system sends a password reset email that includes the token to the email address of the convenience user.
//
// @param request - ResetUserPasswordRequest
//
// @return ResetUserPasswordResponse
func (client *Client) ResetUserPassword(request *ResetUserPasswordRequest) (_result *ResetUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetUserPasswordResponse{}
	_body, _err := client.ResetUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates a user property with a convenience user.
//
// @param request - SetUserPropertyValueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetUserPropertyValueResponse
func (client *Client) SetUserPropertyValueWithOptions(request *SetUserPropertyValueRequest, runtime *util.RuntimeOptions) (_result *SetUserPropertyValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		body["PropertyId"] = request.PropertyId
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyValueId)) {
		body["PropertyValueId"] = request.PropertyValueId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["UserId"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SetUserPropertyValue"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SetUserPropertyValueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates a user property with a convenience user.
//
// @param request - SetUserPropertyValueRequest
//
// @return SetUserPropertyValueResponse
func (client *Client) SetUserPropertyValue(request *SetUserPropertyValueRequest) (_result *SetUserPropertyValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetUserPropertyValueResponse{}
	_body, _err := client.SetUserPropertyValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从钉钉手动同步老师学生信息
//
// @param request - SyncAllEduInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SyncAllEduInfoResponse
func (client *Client) SyncAllEduInfoWithOptions(runtime *util.RuntimeOptions) (_result *SyncAllEduInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("SyncAllEduInfo"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SyncAllEduInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从钉钉手动同步老师学生信息
//
// @return SyncAllEduInfoResponse
func (client *Client) SyncAllEduInfo() (_result *SyncAllEduInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SyncAllEduInfoResponse{}
	_body, _err := client.SyncAllEduInfoWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unlocks a virtual multi-factor authentication (MFA) device that is bound to a convenience user.
//
// @param request - UnlockMfaDeviceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockMfaDeviceResponse
func (client *Client) UnlockMfaDeviceWithOptions(request *UnlockMfaDeviceRequest, runtime *util.RuntimeOptions) (_result *UnlockMfaDeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AdDomain)) {
		query["AdDomain"] = request.AdDomain
	}

	if !tea.BoolValue(util.IsUnset(request.SerialNumber)) {
		query["SerialNumber"] = request.SerialNumber
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnlockMfaDevice"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnlockMfaDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlocks a virtual multi-factor authentication (MFA) device that is bound to a convenience user.
//
// @param request - UnlockMfaDeviceRequest
//
// @return UnlockMfaDeviceResponse
func (client *Client) UnlockMfaDevice(request *UnlockMfaDeviceRequest) (_result *UnlockMfaDeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockMfaDeviceResponse{}
	_body, _err := client.UnlockMfaDeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unlocks one or more convenience users.
//
// @param request - UnlockUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnlockUsersResponse
func (client *Client) UnlockUsersWithOptions(request *UnlockUsersRequest, runtime *util.RuntimeOptions) (_result *UnlockUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoLockTime)) {
		query["AutoLockTime"] = request.AutoLockTime
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnlockUsers"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UnlockUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlocks one or more convenience users.
//
// @param request - UnlockUsersRequest
//
// @return UnlockUsersResponse
func (client *Client) UnlockUsers(request *UnlockUsersRequest) (_result *UnlockUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnlockUsersResponse{}
	_body, _err := client.UnlockUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a user property.
//
// @param request - UpdatePropertyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePropertyResponse
func (client *Client) UpdatePropertyWithOptions(request *UpdatePropertyRequest, runtime *util.RuntimeOptions) (_result *UpdatePropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PropertyId)) {
		body["PropertyId"] = request.PropertyId
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyKey)) {
		body["PropertyKey"] = request.PropertyKey
	}

	if !tea.BoolValue(util.IsUnset(request.PropertyValues)) {
		body["PropertyValues"] = request.PropertyValues
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateProperty"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePropertyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a user property.
//
// @param request - UpdatePropertyRequest
//
// @return UpdatePropertyResponse
func (client *Client) UpdateProperty(request *UpdatePropertyRequest) (_result *UpdatePropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdatePropertyResponse{}
	_body, _err := client.UpdatePropertyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户批量分配角色
//
// @param request - UserBatchJoinGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UserBatchJoinGroupResponse
func (client *Client) UserBatchJoinGroupWithOptions(request *UserBatchJoinGroupRequest, runtime *util.RuntimeOptions) (_result *UserBatchJoinGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserIds)) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UserBatchJoinGroup"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UserBatchJoinGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户批量分配角色
//
// @param request - UserBatchJoinGroupRequest
//
// @return UserBatchJoinGroupResponse
func (client *Client) UserBatchJoinGroup(request *UserBatchJoinGroupRequest) (_result *UserBatchJoinGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UserBatchJoinGroupResponse{}
	_body, _err := client.UserBatchJoinGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 用户批量移出角色
//
// @param request - UserBatchQuitGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UserBatchQuitGroupResponse
func (client *Client) UserBatchQuitGroupWithOptions(request *UserBatchQuitGroupRequest, runtime *util.RuntimeOptions) (_result *UserBatchQuitGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndUserIds)) {
		body["EndUserIds"] = request.EndUserIds
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["GroupId"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UserBatchQuitGroup"),
		Version:     tea.String("2021-03-08"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UserBatchQuitGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用户批量移出角色
//
// @param request - UserBatchQuitGroupRequest
//
// @return UserBatchQuitGroupResponse
func (client *Client) UserBatchQuitGroup(request *UserBatchQuitGroupRequest) (_result *UserBatchQuitGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UserBatchQuitGroupResponse{}
	_body, _err := client.UserBatchQuitGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
