// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetUserSsoSettingsRequest struct {
}

func (s GetUserSsoSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserSsoSettingsRequest) GoString() string {
	return s.String()
}

type GetUserSsoSettingsResponse struct {
	RequestId       *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UserSsoSettings *GetUserSsoSettingsResponseUserSsoSettings `json:"UserSsoSettings,omitempty" xml:"UserSsoSettings,omitempty" require:"true" type:"Struct"`
}

func (s GetUserSsoSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserSsoSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetUserSsoSettingsResponse) SetRequestId(v string) *GetUserSsoSettingsResponse {
	s.RequestId = &v
	return s
}

func (s *GetUserSsoSettingsResponse) SetUserSsoSettings(v *GetUserSsoSettingsResponseUserSsoSettings) *GetUserSsoSettingsResponse {
	s.UserSsoSettings = v
	return s
}

type GetUserSsoSettingsResponseUserSsoSettings struct {
	MetadataDocument *string `json:"MetadataDocument,omitempty" xml:"MetadataDocument,omitempty" require:"true"`
	SsoEnabled       *bool   `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty" require:"true"`
	AuxiliaryDomain  *string `json:"AuxiliaryDomain,omitempty" xml:"AuxiliaryDomain,omitempty" require:"true"`
}

func (s GetUserSsoSettingsResponseUserSsoSettings) String() string {
	return tea.Prettify(s)
}

func (s GetUserSsoSettingsResponseUserSsoSettings) GoString() string {
	return s.String()
}

func (s *GetUserSsoSettingsResponseUserSsoSettings) SetMetadataDocument(v string) *GetUserSsoSettingsResponseUserSsoSettings {
	s.MetadataDocument = &v
	return s
}

func (s *GetUserSsoSettingsResponseUserSsoSettings) SetSsoEnabled(v bool) *GetUserSsoSettingsResponseUserSsoSettings {
	s.SsoEnabled = &v
	return s
}

func (s *GetUserSsoSettingsResponseUserSsoSettings) SetAuxiliaryDomain(v string) *GetUserSsoSettingsResponseUserSsoSettings {
	s.AuxiliaryDomain = &v
	return s
}

type SetUserSsoSettingsRequest struct {
	MetadataDocument *string `json:"MetadataDocument,omitempty" xml:"MetadataDocument,omitempty"`
	SsoEnabled       *bool   `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty"`
	AuxiliaryDomain  *string `json:"AuxiliaryDomain,omitempty" xml:"AuxiliaryDomain,omitempty"`
}

func (s SetUserSsoSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s SetUserSsoSettingsRequest) GoString() string {
	return s.String()
}

func (s *SetUserSsoSettingsRequest) SetMetadataDocument(v string) *SetUserSsoSettingsRequest {
	s.MetadataDocument = &v
	return s
}

func (s *SetUserSsoSettingsRequest) SetSsoEnabled(v bool) *SetUserSsoSettingsRequest {
	s.SsoEnabled = &v
	return s
}

func (s *SetUserSsoSettingsRequest) SetAuxiliaryDomain(v string) *SetUserSsoSettingsRequest {
	s.AuxiliaryDomain = &v
	return s
}

type SetUserSsoSettingsResponse struct {
	RequestId       *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	UserSsoSettings *SetUserSsoSettingsResponseUserSsoSettings `json:"UserSsoSettings,omitempty" xml:"UserSsoSettings,omitempty" require:"true" type:"Struct"`
}

func (s SetUserSsoSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s SetUserSsoSettingsResponse) GoString() string {
	return s.String()
}

func (s *SetUserSsoSettingsResponse) SetRequestId(v string) *SetUserSsoSettingsResponse {
	s.RequestId = &v
	return s
}

func (s *SetUserSsoSettingsResponse) SetUserSsoSettings(v *SetUserSsoSettingsResponseUserSsoSettings) *SetUserSsoSettingsResponse {
	s.UserSsoSettings = v
	return s
}

type SetUserSsoSettingsResponseUserSsoSettings struct {
	MetadataDocument *string `json:"MetadataDocument,omitempty" xml:"MetadataDocument,omitempty" require:"true"`
	SsoEnabled       *bool   `json:"SsoEnabled,omitempty" xml:"SsoEnabled,omitempty" require:"true"`
	AuxiliaryDomain  *string `json:"AuxiliaryDomain,omitempty" xml:"AuxiliaryDomain,omitempty" require:"true"`
}

func (s SetUserSsoSettingsResponseUserSsoSettings) String() string {
	return tea.Prettify(s)
}

func (s SetUserSsoSettingsResponseUserSsoSettings) GoString() string {
	return s.String()
}

func (s *SetUserSsoSettingsResponseUserSsoSettings) SetMetadataDocument(v string) *SetUserSsoSettingsResponseUserSsoSettings {
	s.MetadataDocument = &v
	return s
}

func (s *SetUserSsoSettingsResponseUserSsoSettings) SetSsoEnabled(v bool) *SetUserSsoSettingsResponseUserSsoSettings {
	s.SsoEnabled = &v
	return s
}

func (s *SetUserSsoSettingsResponseUserSsoSettings) SetAuxiliaryDomain(v string) *SetUserSsoSettingsResponseUserSsoSettings {
	s.AuxiliaryDomain = &v
	return s
}

type SetDefaultDomainRequest struct {
	DefaultDomainName *string `json:"DefaultDomainName,omitempty" xml:"DefaultDomainName,omitempty" require:"true"`
}

func (s SetDefaultDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultDomainRequest) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainRequest) SetDefaultDomainName(v string) *SetDefaultDomainRequest {
	s.DefaultDomainName = &v
	return s
}

type SetDefaultDomainResponse struct {
	DefaultDomainName *string `json:"DefaultDomainName,omitempty" xml:"DefaultDomainName,omitempty" require:"true"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SetDefaultDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s SetDefaultDomainResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultDomainResponse) SetDefaultDomainName(v string) *SetDefaultDomainResponse {
	s.DefaultDomainName = &v
	return s
}

func (s *SetDefaultDomainResponse) SetRequestId(v string) *SetDefaultDomainResponse {
	s.RequestId = &v
	return s
}

type ListUserBasicInfosRequest struct {
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxItems *int    `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListUserBasicInfosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserBasicInfosRequest) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosRequest) SetMarker(v string) *ListUserBasicInfosRequest {
	s.Marker = &v
	return s
}

func (s *ListUserBasicInfosRequest) SetMaxItems(v int) *ListUserBasicInfosRequest {
	s.MaxItems = &v
	return s
}

type ListUserBasicInfosResponse struct {
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	IsTruncated    *bool                                     `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty" require:"true"`
	Marker         *string                                   `json:"Marker,omitempty" xml:"Marker,omitempty" require:"true"`
	UserBasicInfos *ListUserBasicInfosResponseUserBasicInfos `json:"UserBasicInfos,omitempty" xml:"UserBasicInfos,omitempty" require:"true" type:"Struct"`
}

func (s ListUserBasicInfosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserBasicInfosResponse) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosResponse) SetRequestId(v string) *ListUserBasicInfosResponse {
	s.RequestId = &v
	return s
}

func (s *ListUserBasicInfosResponse) SetIsTruncated(v bool) *ListUserBasicInfosResponse {
	s.IsTruncated = &v
	return s
}

func (s *ListUserBasicInfosResponse) SetMarker(v string) *ListUserBasicInfosResponse {
	s.Marker = &v
	return s
}

func (s *ListUserBasicInfosResponse) SetUserBasicInfos(v *ListUserBasicInfosResponseUserBasicInfos) *ListUserBasicInfosResponse {
	s.UserBasicInfos = v
	return s
}

type ListUserBasicInfosResponseUserBasicInfos struct {
	UserBasicInfo []*ListUserBasicInfosResponseUserBasicInfosUserBasicInfo `json:"UserBasicInfo,omitempty" xml:"UserBasicInfo,omitempty" require:"true" type:"Repeated"`
}

func (s ListUserBasicInfosResponseUserBasicInfos) String() string {
	return tea.Prettify(s)
}

func (s ListUserBasicInfosResponseUserBasicInfos) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosResponseUserBasicInfos) SetUserBasicInfo(v []*ListUserBasicInfosResponseUserBasicInfosUserBasicInfo) *ListUserBasicInfosResponseUserBasicInfos {
	s.UserBasicInfo = v
	return s
}

type ListUserBasicInfosResponseUserBasicInfosUserBasicInfo struct {
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
}

func (s ListUserBasicInfosResponseUserBasicInfosUserBasicInfo) String() string {
	return tea.Prettify(s)
}

func (s ListUserBasicInfosResponseUserBasicInfosUserBasicInfo) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosResponseUserBasicInfosUserBasicInfo) SetUserId(v string) *ListUserBasicInfosResponseUserBasicInfosUserBasicInfo {
	s.UserId = &v
	return s
}

func (s *ListUserBasicInfosResponseUserBasicInfosUserBasicInfo) SetDisplayName(v string) *ListUserBasicInfosResponseUserBasicInfosUserBasicInfo {
	s.DisplayName = &v
	return s
}

func (s *ListUserBasicInfosResponseUserBasicInfosUserBasicInfo) SetUserPrincipalName(v string) *ListUserBasicInfosResponseUserBasicInfosUserBasicInfo {
	s.UserPrincipalName = &v
	return s
}

type UpdateApplicationRequest struct {
	AppId                   *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	NewDisplayName          *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	NewRedirectUris         *string `json:"NewRedirectUris,omitempty" xml:"NewRedirectUris,omitempty"`
	NewPredefinedScopes     *string `json:"NewPredefinedScopes,omitempty" xml:"NewPredefinedScopes,omitempty"`
	NewSecretRequired       *bool   `json:"NewSecretRequired,omitempty" xml:"NewSecretRequired,omitempty"`
	NewAccessTokenValidity  *int    `json:"NewAccessTokenValidity,omitempty" xml:"NewAccessTokenValidity,omitempty"`
	NewRefreshTokenValidity *int    `json:"NewRefreshTokenValidity,omitempty" xml:"NewRefreshTokenValidity,omitempty"`
	NewIsMultiTenant        *bool   `json:"NewIsMultiTenant,omitempty" xml:"NewIsMultiTenant,omitempty"`
}

func (s UpdateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationRequest) GoString() string {
	return s.String()
}

func (s *UpdateApplicationRequest) SetAppId(v string) *UpdateApplicationRequest {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewDisplayName(v string) *UpdateApplicationRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewRedirectUris(v string) *UpdateApplicationRequest {
	s.NewRedirectUris = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewPredefinedScopes(v string) *UpdateApplicationRequest {
	s.NewPredefinedScopes = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewSecretRequired(v bool) *UpdateApplicationRequest {
	s.NewSecretRequired = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewAccessTokenValidity(v int) *UpdateApplicationRequest {
	s.NewAccessTokenValidity = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewRefreshTokenValidity(v int) *UpdateApplicationRequest {
	s.NewRefreshTokenValidity = &v
	return s
}

func (s *UpdateApplicationRequest) SetNewIsMultiTenant(v bool) *UpdateApplicationRequest {
	s.NewIsMultiTenant = &v
	return s
}

type UpdateApplicationResponse struct {
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Application *UpdateApplicationResponseApplication `json:"Application,omitempty" xml:"Application,omitempty" require:"true" type:"Struct"`
}

func (s UpdateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponse) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponse) SetRequestId(v string) *UpdateApplicationResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationResponse) SetApplication(v *UpdateApplicationResponseApplication) *UpdateApplicationResponse {
	s.Application = v
	return s
}

type UpdateApplicationResponseApplication struct {
	UpdateDate           *string                                             `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	AccountId            *string                                             `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	AppId                *string                                             `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	SecretRequired       *bool                                               `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty" require:"true"`
	DisplayName          *string                                             `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	IsMultiTenant        *bool                                               `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty" require:"true"`
	AccessTokenValidity  *int                                                `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty" require:"true"`
	RefreshTokenValidity *int                                                `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty" require:"true"`
	AppType              *string                                             `json:"AppType,omitempty" xml:"AppType,omitempty" require:"true"`
	CreateDate           *string                                             `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	AppName              *string                                             `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	DelegatedScope       *UpdateApplicationResponseApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" require:"true" type:"Struct"`
	RedirectUris         *UpdateApplicationResponseApplicationRedirectUris   `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" require:"true" type:"Struct"`
}

func (s UpdateApplicationResponseApplication) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseApplication) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseApplication) SetUpdateDate(v string) *UpdateApplicationResponseApplication {
	s.UpdateDate = &v
	return s
}

func (s *UpdateApplicationResponseApplication) SetAccountId(v string) *UpdateApplicationResponseApplication {
	s.AccountId = &v
	return s
}

func (s *UpdateApplicationResponseApplication) SetAppId(v string) *UpdateApplicationResponseApplication {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationResponseApplication) SetSecretRequired(v bool) *UpdateApplicationResponseApplication {
	s.SecretRequired = &v
	return s
}

func (s *UpdateApplicationResponseApplication) SetDisplayName(v string) *UpdateApplicationResponseApplication {
	s.DisplayName = &v
	return s
}

func (s *UpdateApplicationResponseApplication) SetIsMultiTenant(v bool) *UpdateApplicationResponseApplication {
	s.IsMultiTenant = &v
	return s
}

func (s *UpdateApplicationResponseApplication) SetAccessTokenValidity(v int) *UpdateApplicationResponseApplication {
	s.AccessTokenValidity = &v
	return s
}

func (s *UpdateApplicationResponseApplication) SetRefreshTokenValidity(v int) *UpdateApplicationResponseApplication {
	s.RefreshTokenValidity = &v
	return s
}

func (s *UpdateApplicationResponseApplication) SetAppType(v string) *UpdateApplicationResponseApplication {
	s.AppType = &v
	return s
}

func (s *UpdateApplicationResponseApplication) SetCreateDate(v string) *UpdateApplicationResponseApplication {
	s.CreateDate = &v
	return s
}

func (s *UpdateApplicationResponseApplication) SetAppName(v string) *UpdateApplicationResponseApplication {
	s.AppName = &v
	return s
}

func (s *UpdateApplicationResponseApplication) SetDelegatedScope(v *UpdateApplicationResponseApplicationDelegatedScope) *UpdateApplicationResponseApplication {
	s.DelegatedScope = v
	return s
}

func (s *UpdateApplicationResponseApplication) SetRedirectUris(v *UpdateApplicationResponseApplicationRedirectUris) *UpdateApplicationResponseApplication {
	s.RedirectUris = v
	return s
}

type UpdateApplicationResponseApplicationDelegatedScope struct {
	PredefinedScopes *UpdateApplicationResponseApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" require:"true" type:"Struct"`
}

func (s UpdateApplicationResponseApplicationDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseApplicationDelegatedScope) SetPredefinedScopes(v *UpdateApplicationResponseApplicationDelegatedScopePredefinedScopes) *UpdateApplicationResponseApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type UpdateApplicationResponseApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*UpdateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateApplicationResponseApplicationDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*UpdateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) *UpdateApplicationResponseApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type UpdateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s UpdateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *UpdateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *UpdateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *UpdateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

type UpdateApplicationResponseApplicationRedirectUris struct {
	// RedirectUri
	RedirectUri []*string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateApplicationResponseApplicationRedirectUris) String() string {
	return tea.Prettify(s)
}

func (s UpdateApplicationResponseApplicationRedirectUris) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseApplicationRedirectUris) SetRedirectUri(v []*string) *UpdateApplicationResponseApplicationRedirectUris {
	s.RedirectUri = v
	return s
}

type ListApplicationsRequest struct {
}

func (s ListApplicationsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

type ListApplicationsResponse struct {
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Applications *ListApplicationsResponseApplications `json:"Applications,omitempty" xml:"Applications,omitempty" require:"true" type:"Struct"`
}

func (s ListApplicationsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponse) SetRequestId(v string) *ListApplicationsResponse {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponse) SetApplications(v *ListApplicationsResponseApplications) *ListApplicationsResponse {
	s.Applications = v
	return s
}

type ListApplicationsResponseApplications struct {
	Application []*ListApplicationsResponseApplicationsApplication `json:"Application,omitempty" xml:"Application,omitempty" require:"true" type:"Repeated"`
}

func (s ListApplicationsResponseApplications) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseApplications) SetApplication(v []*ListApplicationsResponseApplicationsApplication) *ListApplicationsResponseApplications {
	s.Application = v
	return s
}

type ListApplicationsResponseApplicationsApplication struct {
	UpdateDate           *string                                                        `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	AccountId            *string                                                        `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	AppId                *string                                                        `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	SecretRequired       *bool                                                          `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty" require:"true"`
	DisplayName          *string                                                        `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	IsMultiTenant        *bool                                                          `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty" require:"true"`
	AccessTokenValidity  *int                                                           `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty" require:"true"`
	RefreshTokenValidity *int                                                           `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty" require:"true"`
	AppType              *string                                                        `json:"AppType,omitempty" xml:"AppType,omitempty" require:"true"`
	CreateDate           *string                                                        `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	AppName              *string                                                        `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	DelegatedScope       *ListApplicationsResponseApplicationsApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" require:"true" type:"Struct"`
	RedirectUris         *ListApplicationsResponseApplicationsApplicationRedirectUris   `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" require:"true" type:"Struct"`
}

func (s ListApplicationsResponseApplicationsApplication) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseApplicationsApplication) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseApplicationsApplication) SetUpdateDate(v string) *ListApplicationsResponseApplicationsApplication {
	s.UpdateDate = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplication) SetAccountId(v string) *ListApplicationsResponseApplicationsApplication {
	s.AccountId = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplication) SetAppId(v string) *ListApplicationsResponseApplicationsApplication {
	s.AppId = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplication) SetSecretRequired(v bool) *ListApplicationsResponseApplicationsApplication {
	s.SecretRequired = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplication) SetDisplayName(v string) *ListApplicationsResponseApplicationsApplication {
	s.DisplayName = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplication) SetIsMultiTenant(v bool) *ListApplicationsResponseApplicationsApplication {
	s.IsMultiTenant = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplication) SetAccessTokenValidity(v int) *ListApplicationsResponseApplicationsApplication {
	s.AccessTokenValidity = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplication) SetRefreshTokenValidity(v int) *ListApplicationsResponseApplicationsApplication {
	s.RefreshTokenValidity = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplication) SetAppType(v string) *ListApplicationsResponseApplicationsApplication {
	s.AppType = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplication) SetCreateDate(v string) *ListApplicationsResponseApplicationsApplication {
	s.CreateDate = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplication) SetAppName(v string) *ListApplicationsResponseApplicationsApplication {
	s.AppName = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplication) SetDelegatedScope(v *ListApplicationsResponseApplicationsApplicationDelegatedScope) *ListApplicationsResponseApplicationsApplication {
	s.DelegatedScope = v
	return s
}

func (s *ListApplicationsResponseApplicationsApplication) SetRedirectUris(v *ListApplicationsResponseApplicationsApplicationRedirectUris) *ListApplicationsResponseApplicationsApplication {
	s.RedirectUris = v
	return s
}

type ListApplicationsResponseApplicationsApplicationDelegatedScope struct {
	PredefinedScopes *ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" require:"true" type:"Struct"`
}

func (s ListApplicationsResponseApplicationsApplicationDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseApplicationsApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseApplicationsApplicationDelegatedScope) SetPredefinedScopes(v *ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopes) *ListApplicationsResponseApplicationsApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" require:"true" type:"Repeated"`
}

func (s ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) *ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *ListApplicationsResponseApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

type ListApplicationsResponseApplicationsApplicationRedirectUris struct {
	// RedirectUri
	RedirectUri []*string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty" require:"true" type:"Repeated"`
}

func (s ListApplicationsResponseApplicationsApplicationRedirectUris) String() string {
	return tea.Prettify(s)
}

func (s ListApplicationsResponseApplicationsApplicationRedirectUris) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseApplicationsApplicationRedirectUris) SetRedirectUri(v []*string) *ListApplicationsResponseApplicationsApplicationRedirectUris {
	s.RedirectUri = v
	return s
}

type GetApplicationRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s GetApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) SetAppId(v string) *GetApplicationRequest {
	s.AppId = &v
	return s
}

type GetApplicationResponse struct {
	RequestId   *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Application *GetApplicationResponseApplication `json:"Application,omitempty" xml:"Application,omitempty" require:"true" type:"Struct"`
}

func (s GetApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetApplicationResponse) SetRequestId(v string) *GetApplicationResponse {
	s.RequestId = &v
	return s
}

func (s *GetApplicationResponse) SetApplication(v *GetApplicationResponseApplication) *GetApplicationResponse {
	s.Application = v
	return s
}

type GetApplicationResponseApplication struct {
	UpdateDate           *string                                          `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	AccountId            *string                                          `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	AppId                *string                                          `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	SecretRequired       *bool                                            `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty" require:"true"`
	DisplayName          *string                                          `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	IsMultiTenant        *bool                                            `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty" require:"true"`
	AccessTokenValidity  *int                                             `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty" require:"true"`
	RefreshTokenValidity *int                                             `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty" require:"true"`
	AppType              *string                                          `json:"AppType,omitempty" xml:"AppType,omitempty" require:"true"`
	CreateDate           *string                                          `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	AppName              *string                                          `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	DelegatedScope       *GetApplicationResponseApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" require:"true" type:"Struct"`
	RedirectUris         *GetApplicationResponseApplicationRedirectUris   `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" require:"true" type:"Struct"`
}

func (s GetApplicationResponseApplication) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseApplication) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseApplication) SetUpdateDate(v string) *GetApplicationResponseApplication {
	s.UpdateDate = &v
	return s
}

func (s *GetApplicationResponseApplication) SetAccountId(v string) *GetApplicationResponseApplication {
	s.AccountId = &v
	return s
}

func (s *GetApplicationResponseApplication) SetAppId(v string) *GetApplicationResponseApplication {
	s.AppId = &v
	return s
}

func (s *GetApplicationResponseApplication) SetSecretRequired(v bool) *GetApplicationResponseApplication {
	s.SecretRequired = &v
	return s
}

func (s *GetApplicationResponseApplication) SetDisplayName(v string) *GetApplicationResponseApplication {
	s.DisplayName = &v
	return s
}

func (s *GetApplicationResponseApplication) SetIsMultiTenant(v bool) *GetApplicationResponseApplication {
	s.IsMultiTenant = &v
	return s
}

func (s *GetApplicationResponseApplication) SetAccessTokenValidity(v int) *GetApplicationResponseApplication {
	s.AccessTokenValidity = &v
	return s
}

func (s *GetApplicationResponseApplication) SetRefreshTokenValidity(v int) *GetApplicationResponseApplication {
	s.RefreshTokenValidity = &v
	return s
}

func (s *GetApplicationResponseApplication) SetAppType(v string) *GetApplicationResponseApplication {
	s.AppType = &v
	return s
}

func (s *GetApplicationResponseApplication) SetCreateDate(v string) *GetApplicationResponseApplication {
	s.CreateDate = &v
	return s
}

func (s *GetApplicationResponseApplication) SetAppName(v string) *GetApplicationResponseApplication {
	s.AppName = &v
	return s
}

func (s *GetApplicationResponseApplication) SetDelegatedScope(v *GetApplicationResponseApplicationDelegatedScope) *GetApplicationResponseApplication {
	s.DelegatedScope = v
	return s
}

func (s *GetApplicationResponseApplication) SetRedirectUris(v *GetApplicationResponseApplicationRedirectUris) *GetApplicationResponseApplication {
	s.RedirectUris = v
	return s
}

type GetApplicationResponseApplicationDelegatedScope struct {
	PredefinedScopes *GetApplicationResponseApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" require:"true" type:"Struct"`
}

func (s GetApplicationResponseApplicationDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseApplicationDelegatedScope) SetPredefinedScopes(v *GetApplicationResponseApplicationDelegatedScopePredefinedScopes) *GetApplicationResponseApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type GetApplicationResponseApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*GetApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" require:"true" type:"Repeated"`
}

func (s GetApplicationResponseApplicationDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*GetApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) *GetApplicationResponseApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type GetApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s GetApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *GetApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *GetApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *GetApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

type GetApplicationResponseApplicationRedirectUris struct {
	// RedirectUri
	RedirectUri []*string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty" require:"true" type:"Repeated"`
}

func (s GetApplicationResponseApplicationRedirectUris) String() string {
	return tea.Prettify(s)
}

func (s GetApplicationResponseApplicationRedirectUris) GoString() string {
	return s.String()
}

func (s *GetApplicationResponseApplicationRedirectUris) SetRedirectUri(v []*string) *GetApplicationResponseApplicationRedirectUris {
	s.RedirectUri = v
	return s
}

type DeleteApplicationRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s DeleteApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteApplicationRequest) SetAppId(v string) *DeleteApplicationRequest {
	s.AppId = &v
	return s
}

type DeleteApplicationResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteApplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationResponse) SetRequestId(v string) *DeleteApplicationResponse {
	s.RequestId = &v
	return s
}

type GetAppSecretRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppSecretId *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty" require:"true"`
}

func (s GetAppSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAppSecretRequest) GoString() string {
	return s.String()
}

func (s *GetAppSecretRequest) SetAppId(v string) *GetAppSecretRequest {
	s.AppId = &v
	return s
}

func (s *GetAppSecretRequest) SetAppSecretId(v string) *GetAppSecretRequest {
	s.AppSecretId = &v
	return s
}

type GetAppSecretResponse struct {
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AppSecret *GetAppSecretResponseAppSecret `json:"AppSecret,omitempty" xml:"AppSecret,omitempty" require:"true" type:"Struct"`
}

func (s GetAppSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppSecretResponse) GoString() string {
	return s.String()
}

func (s *GetAppSecretResponse) SetRequestId(v string) *GetAppSecretResponse {
	s.RequestId = &v
	return s
}

func (s *GetAppSecretResponse) SetAppSecret(v *GetAppSecretResponseAppSecret) *GetAppSecretResponse {
	s.AppSecret = v
	return s
}

type GetAppSecretResponseAppSecret struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppSecretValue *string `json:"AppSecretValue,omitempty" xml:"AppSecretValue,omitempty" require:"true"`
	AppSecretId    *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty" require:"true"`
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s GetAppSecretResponseAppSecret) String() string {
	return tea.Prettify(s)
}

func (s GetAppSecretResponseAppSecret) GoString() string {
	return s.String()
}

func (s *GetAppSecretResponseAppSecret) SetAppId(v string) *GetAppSecretResponseAppSecret {
	s.AppId = &v
	return s
}

func (s *GetAppSecretResponseAppSecret) SetAppSecretValue(v string) *GetAppSecretResponseAppSecret {
	s.AppSecretValue = &v
	return s
}

func (s *GetAppSecretResponseAppSecret) SetAppSecretId(v string) *GetAppSecretResponseAppSecret {
	s.AppSecretId = &v
	return s
}

func (s *GetAppSecretResponseAppSecret) SetCreateDate(v string) *GetAppSecretResponseAppSecret {
	s.CreateDate = &v
	return s
}

type CreateAppSecretRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s CreateAppSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateAppSecretRequest) SetAppId(v string) *CreateAppSecretRequest {
	s.AppId = &v
	return s
}

type CreateAppSecretResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AppSecret *CreateAppSecretResponseAppSecret `json:"AppSecret,omitempty" xml:"AppSecret,omitempty" require:"true" type:"Struct"`
}

func (s CreateAppSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSecretResponse) GoString() string {
	return s.String()
}

func (s *CreateAppSecretResponse) SetRequestId(v string) *CreateAppSecretResponse {
	s.RequestId = &v
	return s
}

func (s *CreateAppSecretResponse) SetAppSecret(v *CreateAppSecretResponseAppSecret) *CreateAppSecretResponse {
	s.AppSecret = v
	return s
}

type CreateAppSecretResponseAppSecret struct {
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppSecretValue *string `json:"AppSecretValue,omitempty" xml:"AppSecretValue,omitempty" require:"true"`
	AppSecretId    *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty" require:"true"`
	CreateDate     *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s CreateAppSecretResponseAppSecret) String() string {
	return tea.Prettify(s)
}

func (s CreateAppSecretResponseAppSecret) GoString() string {
	return s.String()
}

func (s *CreateAppSecretResponseAppSecret) SetAppId(v string) *CreateAppSecretResponseAppSecret {
	s.AppId = &v
	return s
}

func (s *CreateAppSecretResponseAppSecret) SetAppSecretValue(v string) *CreateAppSecretResponseAppSecret {
	s.AppSecretValue = &v
	return s
}

func (s *CreateAppSecretResponseAppSecret) SetAppSecretId(v string) *CreateAppSecretResponseAppSecret {
	s.AppSecretId = &v
	return s
}

func (s *CreateAppSecretResponseAppSecret) SetCreateDate(v string) *CreateAppSecretResponseAppSecret {
	s.CreateDate = &v
	return s
}

type ListPredefinedScopesRequest struct {
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
}

func (s ListPredefinedScopesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPredefinedScopesRequest) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesRequest) SetAppType(v string) *ListPredefinedScopesRequest {
	s.AppType = &v
	return s
}

type ListPredefinedScopesResponse struct {
	RequestId        *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PredefinedScopes *ListPredefinedScopesResponsePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" require:"true" type:"Struct"`
}

func (s ListPredefinedScopesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPredefinedScopesResponse) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesResponse) SetRequestId(v string) *ListPredefinedScopesResponse {
	s.RequestId = &v
	return s
}

func (s *ListPredefinedScopesResponse) SetPredefinedScopes(v *ListPredefinedScopesResponsePredefinedScopes) *ListPredefinedScopesResponse {
	s.PredefinedScopes = v
	return s
}

type ListPredefinedScopesResponsePredefinedScopes struct {
	PredefinedScope []*ListPredefinedScopesResponsePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" require:"true" type:"Repeated"`
}

func (s ListPredefinedScopesResponsePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s ListPredefinedScopesResponsePredefinedScopes) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesResponsePredefinedScopes) SetPredefinedScope(v []*ListPredefinedScopesResponsePredefinedScopesPredefinedScope) *ListPredefinedScopesResponsePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type ListPredefinedScopesResponsePredefinedScopesPredefinedScope struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s ListPredefinedScopesResponsePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s ListPredefinedScopesResponsePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ListPredefinedScopesResponsePredefinedScopesPredefinedScope) SetDescription(v string) *ListPredefinedScopesResponsePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ListPredefinedScopesResponsePredefinedScopesPredefinedScope) SetName(v string) *ListPredefinedScopesResponsePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

type CreateApplicationRequest struct {
	DisplayName          *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	AppType              *string `json:"AppType,omitempty" xml:"AppType,omitempty" require:"true"`
	RedirectUris         *string `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty"`
	SecretRequired       *bool   `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty"`
	AccessTokenValidity  *int    `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	RefreshTokenValidity *int    `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	PredefinedScopes     *string `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty"`
	IsMultiTenant        *bool   `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty"`
	AppName              *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s CreateApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationRequest) SetDisplayName(v string) *CreateApplicationRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateApplicationRequest) SetAppType(v string) *CreateApplicationRequest {
	s.AppType = &v
	return s
}

func (s *CreateApplicationRequest) SetRedirectUris(v string) *CreateApplicationRequest {
	s.RedirectUris = &v
	return s
}

func (s *CreateApplicationRequest) SetSecretRequired(v bool) *CreateApplicationRequest {
	s.SecretRequired = &v
	return s
}

func (s *CreateApplicationRequest) SetAccessTokenValidity(v int) *CreateApplicationRequest {
	s.AccessTokenValidity = &v
	return s
}

func (s *CreateApplicationRequest) SetRefreshTokenValidity(v int) *CreateApplicationRequest {
	s.RefreshTokenValidity = &v
	return s
}

func (s *CreateApplicationRequest) SetPredefinedScopes(v string) *CreateApplicationRequest {
	s.PredefinedScopes = &v
	return s
}

func (s *CreateApplicationRequest) SetIsMultiTenant(v bool) *CreateApplicationRequest {
	s.IsMultiTenant = &v
	return s
}

func (s *CreateApplicationRequest) SetAppName(v string) *CreateApplicationRequest {
	s.AppName = &v
	return s
}

type CreateApplicationResponse struct {
	RequestId   *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Application *CreateApplicationResponseApplication `json:"Application,omitempty" xml:"Application,omitempty" require:"true" type:"Struct"`
}

func (s CreateApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponse) SetRequestId(v string) *CreateApplicationResponse {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationResponse) SetApplication(v *CreateApplicationResponseApplication) *CreateApplicationResponse {
	s.Application = v
	return s
}

type CreateApplicationResponseApplication struct {
	UpdateDate           *string                                             `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	AccountId            *string                                             `json:"AccountId,omitempty" xml:"AccountId,omitempty" require:"true"`
	AppId                *string                                             `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	SecretRequired       *bool                                               `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty" require:"true"`
	DisplayName          *string                                             `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	IsMultiTenant        *bool                                               `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty" require:"true"`
	AccessTokenValidity  *int                                                `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty" require:"true"`
	RefreshTokenValidity *int                                                `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty" require:"true"`
	AppType              *string                                             `json:"AppType,omitempty" xml:"AppType,omitempty" require:"true"`
	CreateDate           *string                                             `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	AppName              *string                                             `json:"AppName,omitempty" xml:"AppName,omitempty" require:"true"`
	DelegatedScope       *CreateApplicationResponseApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" require:"true" type:"Struct"`
	RedirectUris         *CreateApplicationResponseApplicationRedirectUris   `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" require:"true" type:"Struct"`
}

func (s CreateApplicationResponseApplication) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseApplication) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseApplication) SetUpdateDate(v string) *CreateApplicationResponseApplication {
	s.UpdateDate = &v
	return s
}

func (s *CreateApplicationResponseApplication) SetAccountId(v string) *CreateApplicationResponseApplication {
	s.AccountId = &v
	return s
}

func (s *CreateApplicationResponseApplication) SetAppId(v string) *CreateApplicationResponseApplication {
	s.AppId = &v
	return s
}

func (s *CreateApplicationResponseApplication) SetSecretRequired(v bool) *CreateApplicationResponseApplication {
	s.SecretRequired = &v
	return s
}

func (s *CreateApplicationResponseApplication) SetDisplayName(v string) *CreateApplicationResponseApplication {
	s.DisplayName = &v
	return s
}

func (s *CreateApplicationResponseApplication) SetIsMultiTenant(v bool) *CreateApplicationResponseApplication {
	s.IsMultiTenant = &v
	return s
}

func (s *CreateApplicationResponseApplication) SetAccessTokenValidity(v int) *CreateApplicationResponseApplication {
	s.AccessTokenValidity = &v
	return s
}

func (s *CreateApplicationResponseApplication) SetRefreshTokenValidity(v int) *CreateApplicationResponseApplication {
	s.RefreshTokenValidity = &v
	return s
}

func (s *CreateApplicationResponseApplication) SetAppType(v string) *CreateApplicationResponseApplication {
	s.AppType = &v
	return s
}

func (s *CreateApplicationResponseApplication) SetCreateDate(v string) *CreateApplicationResponseApplication {
	s.CreateDate = &v
	return s
}

func (s *CreateApplicationResponseApplication) SetAppName(v string) *CreateApplicationResponseApplication {
	s.AppName = &v
	return s
}

func (s *CreateApplicationResponseApplication) SetDelegatedScope(v *CreateApplicationResponseApplicationDelegatedScope) *CreateApplicationResponseApplication {
	s.DelegatedScope = v
	return s
}

func (s *CreateApplicationResponseApplication) SetRedirectUris(v *CreateApplicationResponseApplicationRedirectUris) *CreateApplicationResponseApplication {
	s.RedirectUris = v
	return s
}

type CreateApplicationResponseApplicationDelegatedScope struct {
	PredefinedScopes *CreateApplicationResponseApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" require:"true" type:"Struct"`
}

func (s CreateApplicationResponseApplicationDelegatedScope) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseApplicationDelegatedScope) SetPredefinedScopes(v *CreateApplicationResponseApplicationDelegatedScopePredefinedScopes) *CreateApplicationResponseApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

type CreateApplicationResponseApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*CreateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" require:"true" type:"Repeated"`
}

func (s CreateApplicationResponseApplicationDelegatedScopePredefinedScopes) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*CreateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) *CreateApplicationResponseApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

type CreateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
}

func (s CreateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *CreateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *CreateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *CreateApplicationResponseApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

type CreateApplicationResponseApplicationRedirectUris struct {
	// RedirectUri
	RedirectUri []*string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty" require:"true" type:"Repeated"`
}

func (s CreateApplicationResponseApplicationRedirectUris) String() string {
	return tea.Prettify(s)
}

func (s CreateApplicationResponseApplicationRedirectUris) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseApplicationRedirectUris) SetRedirectUri(v []*string) *CreateApplicationResponseApplicationRedirectUris {
	s.RedirectUri = v
	return s
}

type DeleteAppSecretRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppSecretId *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty" require:"true"`
}

func (s DeleteAppSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppSecretRequest) SetAppId(v string) *DeleteAppSecretRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppSecretRequest) SetAppSecretId(v string) *DeleteAppSecretRequest {
	s.AppSecretId = &v
	return s
}

type DeleteAppSecretResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteAppSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppSecretResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppSecretResponse) SetRequestId(v string) *DeleteAppSecretResponse {
	s.RequestId = &v
	return s
}

type ListAppSecretIdsRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
}

func (s ListAppSecretIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppSecretIdsRequest) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsRequest) SetAppId(v string) *ListAppSecretIdsRequest {
	s.AppId = &v
	return s
}

type ListAppSecretIdsResponse struct {
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AppSecrets *ListAppSecretIdsResponseAppSecrets `json:"AppSecrets,omitempty" xml:"AppSecrets,omitempty" require:"true" type:"Struct"`
}

func (s ListAppSecretIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppSecretIdsResponse) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsResponse) SetRequestId(v string) *ListAppSecretIdsResponse {
	s.RequestId = &v
	return s
}

func (s *ListAppSecretIdsResponse) SetAppSecrets(v *ListAppSecretIdsResponseAppSecrets) *ListAppSecretIdsResponse {
	s.AppSecrets = v
	return s
}

type ListAppSecretIdsResponseAppSecrets struct {
	AppSecret []*ListAppSecretIdsResponseAppSecretsAppSecret `json:"AppSecret,omitempty" xml:"AppSecret,omitempty" require:"true" type:"Repeated"`
}

func (s ListAppSecretIdsResponseAppSecrets) String() string {
	return tea.Prettify(s)
}

func (s ListAppSecretIdsResponseAppSecrets) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsResponseAppSecrets) SetAppSecret(v []*ListAppSecretIdsResponseAppSecretsAppSecret) *ListAppSecretIdsResponseAppSecrets {
	s.AppSecret = v
	return s
}

type ListAppSecretIdsResponseAppSecretsAppSecret struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty" require:"true"`
	AppSecretId *string `json:"AppSecretId,omitempty" xml:"AppSecretId,omitempty" require:"true"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s ListAppSecretIdsResponseAppSecretsAppSecret) String() string {
	return tea.Prettify(s)
}

func (s ListAppSecretIdsResponseAppSecretsAppSecret) GoString() string {
	return s.String()
}

func (s *ListAppSecretIdsResponseAppSecretsAppSecret) SetAppId(v string) *ListAppSecretIdsResponseAppSecretsAppSecret {
	s.AppId = &v
	return s
}

func (s *ListAppSecretIdsResponseAppSecretsAppSecret) SetAppSecretId(v string) *ListAppSecretIdsResponseAppSecretsAppSecret {
	s.AppSecretId = &v
	return s
}

func (s *ListAppSecretIdsResponseAppSecretsAppSecret) SetCreateDate(v string) *ListAppSecretIdsResponseAppSecretsAppSecret {
	s.CreateDate = &v
	return s
}

type GenerateCredentialReportRequest struct {
}

func (s GenerateCredentialReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateCredentialReportRequest) GoString() string {
	return s.String()
}

type GenerateCredentialReportResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	State     *string `json:"State,omitempty" xml:"State,omitempty" require:"true"`
}

func (s GenerateCredentialReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateCredentialReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateCredentialReportResponse) SetRequestId(v string) *GenerateCredentialReportResponse {
	s.RequestId = &v
	return s
}

func (s *GenerateCredentialReportResponse) SetState(v string) *GenerateCredentialReportResponse {
	s.State = &v
	return s
}

type GetCredentialReportRequest struct {
}

func (s GetCredentialReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCredentialReportRequest) GoString() string {
	return s.String()
}

type GetCredentialReportResponse struct {
	GeneratedTime *string `json:"GeneratedTime,omitempty" xml:"GeneratedTime,omitempty" require:"true"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Content       *string `json:"Content,omitempty" xml:"Content,omitempty" require:"true"`
}

func (s GetCredentialReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCredentialReportResponse) GoString() string {
	return s.String()
}

func (s *GetCredentialReportResponse) SetGeneratedTime(v string) *GetCredentialReportResponse {
	s.GeneratedTime = &v
	return s
}

func (s *GetCredentialReportResponse) SetRequestId(v string) *GetCredentialReportResponse {
	s.RequestId = &v
	return s
}

func (s *GetCredentialReportResponse) SetContent(v string) *GetCredentialReportResponse {
	s.Content = &v
	return s
}

type UpdateUserRequest struct {
	UserPrincipalName    *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
	UserId               *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	NewUserPrincipalName *string `json:"NewUserPrincipalName,omitempty" xml:"NewUserPrincipalName,omitempty"`
	NewDisplayName       *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	NewMobilePhone       *string `json:"NewMobilePhone,omitempty" xml:"NewMobilePhone,omitempty"`
	NewEmail             *string `json:"NewEmail,omitempty" xml:"NewEmail,omitempty"`
	NewComments          *string `json:"NewComments,omitempty" xml:"NewComments,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetUserPrincipalName(v string) *UpdateUserRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *UpdateUserRequest) SetUserId(v string) *UpdateUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserRequest) SetNewUserPrincipalName(v string) *UpdateUserRequest {
	s.NewUserPrincipalName = &v
	return s
}

func (s *UpdateUserRequest) SetNewDisplayName(v string) *UpdateUserRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateUserRequest) SetNewMobilePhone(v string) *UpdateUserRequest {
	s.NewMobilePhone = &v
	return s
}

func (s *UpdateUserRequest) SetNewEmail(v string) *UpdateUserRequest {
	s.NewEmail = &v
	return s
}

func (s *UpdateUserRequest) SetNewComments(v string) *UpdateUserRequest {
	s.NewComments = &v
	return s
}

type UpdateUserResponse struct {
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	User      *UpdateUserResponseUser `json:"User,omitempty" xml:"User,omitempty" require:"true" type:"Struct"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetRequestId(v string) *UpdateUserResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponse) SetUser(v *UpdateUserResponseUser) *UpdateUserResponse {
	s.User = v
	return s
}

type UpdateUserResponseUser struct {
	UpdateDate        *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty" require:"true"`
	Comments          *string `json:"Comments,omitempty" xml:"Comments,omitempty" require:"true"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	LastLoginDate     *string `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty" require:"true"`
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	CreateDate        *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	MobilePhone       *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty" require:"true"`
}

func (s UpdateUserResponseUser) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponseUser) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseUser) SetUpdateDate(v string) *UpdateUserResponseUser {
	s.UpdateDate = &v
	return s
}

func (s *UpdateUserResponseUser) SetEmail(v string) *UpdateUserResponseUser {
	s.Email = &v
	return s
}

func (s *UpdateUserResponseUser) SetComments(v string) *UpdateUserResponseUser {
	s.Comments = &v
	return s
}

func (s *UpdateUserResponseUser) SetUserId(v string) *UpdateUserResponseUser {
	s.UserId = &v
	return s
}

func (s *UpdateUserResponseUser) SetLastLoginDate(v string) *UpdateUserResponseUser {
	s.LastLoginDate = &v
	return s
}

func (s *UpdateUserResponseUser) SetDisplayName(v string) *UpdateUserResponseUser {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserResponseUser) SetUserPrincipalName(v string) *UpdateUserResponseUser {
	s.UserPrincipalName = &v
	return s
}

func (s *UpdateUserResponseUser) SetCreateDate(v string) *UpdateUserResponseUser {
	s.CreateDate = &v
	return s
}

func (s *UpdateUserResponseUser) SetMobilePhone(v string) *UpdateUserResponseUser {
	s.MobilePhone = &v
	return s
}

type UpdateSAMLProviderRequest struct {
	SAMLProviderName               *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty" require:"true"`
	NewDescription                 *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	NewEncodedSAMLMetadataDocument *string `json:"NewEncodedSAMLMetadataDocument,omitempty" xml:"NewEncodedSAMLMetadataDocument,omitempty"`
}

func (s UpdateSAMLProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSAMLProviderRequest) GoString() string {
	return s.String()
}

func (s *UpdateSAMLProviderRequest) SetSAMLProviderName(v string) *UpdateSAMLProviderRequest {
	s.SAMLProviderName = &v
	return s
}

func (s *UpdateSAMLProviderRequest) SetNewDescription(v string) *UpdateSAMLProviderRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdateSAMLProviderRequest) SetNewEncodedSAMLMetadataDocument(v string) *UpdateSAMLProviderRequest {
	s.NewEncodedSAMLMetadataDocument = &v
	return s
}

type UpdateSAMLProviderResponse struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SAMLProvider *UpdateSAMLProviderResponseSAMLProvider `json:"SAMLProvider,omitempty" xml:"SAMLProvider,omitempty" require:"true" type:"Struct"`
}

func (s UpdateSAMLProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSAMLProviderResponse) GoString() string {
	return s.String()
}

func (s *UpdateSAMLProviderResponse) SetRequestId(v string) *UpdateSAMLProviderResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateSAMLProviderResponse) SetSAMLProvider(v *UpdateSAMLProviderResponseSAMLProvider) *UpdateSAMLProviderResponse {
	s.SAMLProvider = v
	return s
}

type UpdateSAMLProviderResponseSAMLProvider struct {
	UpdateDate       *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty" require:"true"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Arn              *string `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
	CreateDate       *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s UpdateSAMLProviderResponseSAMLProvider) String() string {
	return tea.Prettify(s)
}

func (s UpdateSAMLProviderResponseSAMLProvider) GoString() string {
	return s.String()
}

func (s *UpdateSAMLProviderResponseSAMLProvider) SetUpdateDate(v string) *UpdateSAMLProviderResponseSAMLProvider {
	s.UpdateDate = &v
	return s
}

func (s *UpdateSAMLProviderResponseSAMLProvider) SetSAMLProviderName(v string) *UpdateSAMLProviderResponseSAMLProvider {
	s.SAMLProviderName = &v
	return s
}

func (s *UpdateSAMLProviderResponseSAMLProvider) SetDescription(v string) *UpdateSAMLProviderResponseSAMLProvider {
	s.Description = &v
	return s
}

func (s *UpdateSAMLProviderResponseSAMLProvider) SetArn(v string) *UpdateSAMLProviderResponseSAMLProvider {
	s.Arn = &v
	return s
}

func (s *UpdateSAMLProviderResponseSAMLProvider) SetCreateDate(v string) *UpdateSAMLProviderResponseSAMLProvider {
	s.CreateDate = &v
	return s
}

type UpdateLoginProfileRequest struct {
	UserPrincipalName     *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	Password              *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PasswordResetRequired *bool   `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	MFABindRequired       *bool   `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateLoginProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileRequest) SetUserPrincipalName(v string) *UpdateLoginProfileRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetPassword(v string) *UpdateLoginProfileRequest {
	s.Password = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetPasswordResetRequired(v bool) *UpdateLoginProfileRequest {
	s.PasswordResetRequired = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetMFABindRequired(v bool) *UpdateLoginProfileRequest {
	s.MFABindRequired = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetStatus(v string) *UpdateLoginProfileRequest {
	s.Status = &v
	return s
}

type UpdateLoginProfileResponse struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LoginProfile *UpdateLoginProfileResponseLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" require:"true" type:"Struct"`
}

func (s UpdateLoginProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileResponse) SetRequestId(v string) *UpdateLoginProfileResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateLoginProfileResponse) SetLoginProfile(v *UpdateLoginProfileResponseLoginProfile) *UpdateLoginProfileResponse {
	s.LoginProfile = v
	return s
}

type UpdateLoginProfileResponseLoginProfile struct {
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	UpdateDate            *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	PasswordResetRequired *bool   `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty" require:"true"`
	UserPrincipalName     *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	MFABindRequired       *bool   `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty" require:"true"`
}

func (s UpdateLoginProfileResponseLoginProfile) String() string {
	return tea.Prettify(s)
}

func (s UpdateLoginProfileResponseLoginProfile) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileResponseLoginProfile) SetStatus(v string) *UpdateLoginProfileResponseLoginProfile {
	s.Status = &v
	return s
}

func (s *UpdateLoginProfileResponseLoginProfile) SetUpdateDate(v string) *UpdateLoginProfileResponseLoginProfile {
	s.UpdateDate = &v
	return s
}

func (s *UpdateLoginProfileResponseLoginProfile) SetPasswordResetRequired(v bool) *UpdateLoginProfileResponseLoginProfile {
	s.PasswordResetRequired = &v
	return s
}

func (s *UpdateLoginProfileResponseLoginProfile) SetUserPrincipalName(v string) *UpdateLoginProfileResponseLoginProfile {
	s.UserPrincipalName = &v
	return s
}

func (s *UpdateLoginProfileResponseLoginProfile) SetMFABindRequired(v bool) *UpdateLoginProfileResponseLoginProfile {
	s.MFABindRequired = &v
	return s
}

type UpdateGroupRequest struct {
	NewComments    *string `json:"NewComments,omitempty" xml:"NewComments,omitempty"`
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	NewGroupName   *string `json:"NewGroupName,omitempty" xml:"NewGroupName,omitempty"`
	GroupName      *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s UpdateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) SetNewComments(v string) *UpdateGroupRequest {
	s.NewComments = &v
	return s
}

func (s *UpdateGroupRequest) SetNewDisplayName(v string) *UpdateGroupRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateGroupRequest) SetNewGroupName(v string) *UpdateGroupRequest {
	s.NewGroupName = &v
	return s
}

func (s *UpdateGroupRequest) SetGroupName(v string) *UpdateGroupRequest {
	s.GroupName = &v
	return s
}

type UpdateGroupResponse struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Group     *UpdateGroupResponseGroup `json:"Group,omitempty" xml:"Group,omitempty" require:"true" type:"Struct"`
}

func (s UpdateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponse) SetRequestId(v string) *UpdateGroupResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateGroupResponse) SetGroup(v *UpdateGroupResponseGroup) *UpdateGroupResponse {
	s.Group = v
	return s
}

type UpdateGroupResponseGroup struct {
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
	UpdateDate  *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Comments    *string `json:"Comments,omitempty" xml:"Comments,omitempty" require:"true"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
}

func (s UpdateGroupResponseGroup) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponseGroup) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponseGroup) SetGroupName(v string) *UpdateGroupResponseGroup {
	s.GroupName = &v
	return s
}

func (s *UpdateGroupResponseGroup) SetUpdateDate(v string) *UpdateGroupResponseGroup {
	s.UpdateDate = &v
	return s
}

func (s *UpdateGroupResponseGroup) SetComments(v string) *UpdateGroupResponseGroup {
	s.Comments = &v
	return s
}

func (s *UpdateGroupResponseGroup) SetDisplayName(v string) *UpdateGroupResponseGroup {
	s.DisplayName = &v
	return s
}

func (s *UpdateGroupResponseGroup) SetCreateDate(v string) *UpdateGroupResponseGroup {
	s.CreateDate = &v
	return s
}

func (s *UpdateGroupResponseGroup) SetGroupId(v string) *UpdateGroupResponseGroup {
	s.GroupId = &v
	return s
}

type UpdateAccessKeyRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
	UserAccessKeyId   *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty" require:"true"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
}

func (s UpdateAccessKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *UpdateAccessKeyRequest) SetUserPrincipalName(v string) *UpdateAccessKeyRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *UpdateAccessKeyRequest) SetUserAccessKeyId(v string) *UpdateAccessKeyRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *UpdateAccessKeyRequest) SetStatus(v string) *UpdateAccessKeyRequest {
	s.Status = &v
	return s
}

type UpdateAccessKeyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateAccessKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *UpdateAccessKeyResponse) SetRequestId(v string) *UpdateAccessKeyResponse {
	s.RequestId = &v
	return s
}

type UnbindMFADeviceRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
}

func (s UnbindMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceRequest) SetUserPrincipalName(v string) *UnbindMFADeviceRequest {
	s.UserPrincipalName = &v
	return s
}

type UnbindMFADeviceResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	MFADevice *UnbindMFADeviceResponseMFADevice `json:"MFADevice,omitempty" xml:"MFADevice,omitempty" require:"true" type:"Struct"`
}

func (s UnbindMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceResponse) SetRequestId(v string) *UnbindMFADeviceResponse {
	s.RequestId = &v
	return s
}

func (s *UnbindMFADeviceResponse) SetMFADevice(v *UnbindMFADeviceResponseMFADevice) *UnbindMFADeviceResponse {
	s.MFADevice = v
	return s
}

type UnbindMFADeviceResponseMFADevice struct {
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty" require:"true"`
}

func (s UnbindMFADeviceResponseMFADevice) String() string {
	return tea.Prettify(s)
}

func (s UnbindMFADeviceResponseMFADevice) GoString() string {
	return s.String()
}

func (s *UnbindMFADeviceResponseMFADevice) SetSerialNumber(v string) *UnbindMFADeviceResponseMFADevice {
	s.SerialNumber = &v
	return s
}

type SetSecurityPreferenceRequest struct {
	EnableSaveMFATicket         *bool   `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty"`
	AllowUserToChangePassword   *bool   `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty"`
	AllowUserToManageAccessKeys *bool   `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty"`
	AllowUserToManageMFADevices *bool   `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty"`
	LoginSessionDuration        *int    `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty"`
	LoginNetworkMasks           *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty"`
}

func (s SetSecurityPreferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceRequest) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceRequest) SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceRequest {
	s.EnableSaveMFATicket = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManageAccessKeys = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceRequest {
	s.AllowUserToManageMFADevices = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetLoginSessionDuration(v int) *SetSecurityPreferenceRequest {
	s.LoginSessionDuration = &v
	return s
}

func (s *SetSecurityPreferenceRequest) SetLoginNetworkMasks(v string) *SetSecurityPreferenceRequest {
	s.LoginNetworkMasks = &v
	return s
}

type SetSecurityPreferenceResponse struct {
	RequestId          *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SecurityPreference *SetSecurityPreferenceResponseSecurityPreference `json:"SecurityPreference,omitempty" xml:"SecurityPreference,omitempty" require:"true" type:"Struct"`
}

func (s SetSecurityPreferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponse) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponse) SetRequestId(v string) *SetSecurityPreferenceResponse {
	s.RequestId = &v
	return s
}

func (s *SetSecurityPreferenceResponse) SetSecurityPreference(v *SetSecurityPreferenceResponseSecurityPreference) *SetSecurityPreferenceResponse {
	s.SecurityPreference = v
	return s
}

type SetSecurityPreferenceResponseSecurityPreference struct {
	LoginProfilePreference *SetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference `json:"LoginProfilePreference,omitempty" xml:"LoginProfilePreference,omitempty" require:"true" type:"Struct"`
	AccessKeyPreference    *SetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference    `json:"AccessKeyPreference,omitempty" xml:"AccessKeyPreference,omitempty" require:"true" type:"Struct"`
	MFAPreference          *SetSecurityPreferenceResponseSecurityPreferenceMFAPreference          `json:"MFAPreference,omitempty" xml:"MFAPreference,omitempty" require:"true" type:"Struct"`
}

func (s SetSecurityPreferenceResponseSecurityPreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseSecurityPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseSecurityPreference) SetLoginProfilePreference(v *SetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) *SetSecurityPreferenceResponseSecurityPreference {
	s.LoginProfilePreference = v
	return s
}

func (s *SetSecurityPreferenceResponseSecurityPreference) SetAccessKeyPreference(v *SetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference) *SetSecurityPreferenceResponseSecurityPreference {
	s.AccessKeyPreference = v
	return s
}

func (s *SetSecurityPreferenceResponseSecurityPreference) SetMFAPreference(v *SetSecurityPreferenceResponseSecurityPreferenceMFAPreference) *SetSecurityPreferenceResponseSecurityPreference {
	s.MFAPreference = v
	return s
}

type SetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference struct {
	LoginSessionDuration      *int    `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty" require:"true"`
	LoginNetworkMasks         *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty" require:"true"`
	AllowUserToChangePassword *bool   `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty" require:"true"`
	EnableSaveMFATicket       *bool   `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty" require:"true"`
}

func (s SetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) SetLoginSessionDuration(v int) *SetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference {
	s.LoginSessionDuration = &v
	return s
}

func (s *SetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) SetLoginNetworkMasks(v string) *SetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference {
	s.LoginNetworkMasks = &v
	return s
}

func (s *SetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) SetAllowUserToChangePassword(v bool) *SetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *SetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) SetEnableSaveMFATicket(v bool) *SetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference {
	s.EnableSaveMFATicket = &v
	return s
}

type SetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference struct {
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty" require:"true"`
}

func (s SetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference) SetAllowUserToManageAccessKeys(v bool) *SetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference {
	s.AllowUserToManageAccessKeys = &v
	return s
}

type SetSecurityPreferenceResponseSecurityPreferenceMFAPreference struct {
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty" require:"true"`
}

func (s SetSecurityPreferenceResponseSecurityPreferenceMFAPreference) String() string {
	return tea.Prettify(s)
}

func (s SetSecurityPreferenceResponseSecurityPreferenceMFAPreference) GoString() string {
	return s.String()
}

func (s *SetSecurityPreferenceResponseSecurityPreferenceMFAPreference) SetAllowUserToManageMFADevices(v bool) *SetSecurityPreferenceResponseSecurityPreferenceMFAPreference {
	s.AllowUserToManageMFADevices = &v
	return s
}

type SetPasswordPolicyRequest struct {
	MinimumPasswordLength             *int  `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	RequireLowercaseCharacters        *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	RequireUppercaseCharacters        *bool `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
	RequireNumbers                    *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	RequireSymbols                    *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	HardExpire                        *bool `json:"HardExpire,omitempty" xml:"HardExpire,omitempty"`
	MaxLoginAttemps                   *int  `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	PasswordReusePrevention           *int  `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	MaxPasswordAge                    *int  `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	MinimumPasswordDifferentCharacter *int  `json:"MinimumPasswordDifferentCharacter,omitempty" xml:"MinimumPasswordDifferentCharacter,omitempty"`
	PasswordNotContainUserName        *bool `json:"PasswordNotContainUserName,omitempty" xml:"PasswordNotContainUserName,omitempty"`
}

func (s SetPasswordPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyRequest) SetMinimumPasswordLength(v int) *SetPasswordPolicyRequest {
	s.MinimumPasswordLength = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireLowercaseCharacters(v bool) *SetPasswordPolicyRequest {
	s.RequireLowercaseCharacters = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireUppercaseCharacters(v bool) *SetPasswordPolicyRequest {
	s.RequireUppercaseCharacters = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireNumbers(v bool) *SetPasswordPolicyRequest {
	s.RequireNumbers = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireSymbols(v bool) *SetPasswordPolicyRequest {
	s.RequireSymbols = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetHardExpire(v bool) *SetPasswordPolicyRequest {
	s.HardExpire = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMaxLoginAttemps(v int) *SetPasswordPolicyRequest {
	s.MaxLoginAttemps = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetPasswordReusePrevention(v int) *SetPasswordPolicyRequest {
	s.PasswordReusePrevention = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMaxPasswordAge(v int) *SetPasswordPolicyRequest {
	s.MaxPasswordAge = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMinimumPasswordDifferentCharacter(v int) *SetPasswordPolicyRequest {
	s.MinimumPasswordDifferentCharacter = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetPasswordNotContainUserName(v bool) *SetPasswordPolicyRequest {
	s.PasswordNotContainUserName = &v
	return s
}

type SetPasswordPolicyResponse struct {
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PasswordPolicy *SetPasswordPolicyResponsePasswordPolicy `json:"PasswordPolicy,omitempty" xml:"PasswordPolicy,omitempty" require:"true" type:"Struct"`
}

func (s SetPasswordPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponse) SetRequestId(v string) *SetPasswordPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *SetPasswordPolicyResponse) SetPasswordPolicy(v *SetPasswordPolicyResponsePasswordPolicy) *SetPasswordPolicyResponse {
	s.PasswordPolicy = v
	return s
}

type SetPasswordPolicyResponsePasswordPolicy struct {
	MinimumPasswordLength             *int  `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty" require:"true"`
	HardExpire                        *bool `json:"HardExpire,omitempty" xml:"HardExpire,omitempty" require:"true"`
	RequireLowercaseCharacters        *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty" require:"true"`
	RequireNumbers                    *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty" require:"true"`
	MaxLoginAttemps                   *int  `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty" require:"true"`
	MaxPasswordAge                    *int  `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty" require:"true"`
	PasswordNotContainUserName        *bool `json:"PasswordNotContainUserName,omitempty" xml:"PasswordNotContainUserName,omitempty" require:"true"`
	PasswordReusePrevention           *int  `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty" require:"true"`
	RequireUppercaseCharacters        *bool `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty" require:"true"`
	MinimumPasswordDifferentCharacter *int  `json:"MinimumPasswordDifferentCharacter,omitempty" xml:"MinimumPasswordDifferentCharacter,omitempty" require:"true"`
	RequireSymbols                    *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty" require:"true"`
}

func (s SetPasswordPolicyResponsePasswordPolicy) String() string {
	return tea.Prettify(s)
}

func (s SetPasswordPolicyResponsePasswordPolicy) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponsePasswordPolicy) SetMinimumPasswordLength(v int) *SetPasswordPolicyResponsePasswordPolicy {
	s.MinimumPasswordLength = &v
	return s
}

func (s *SetPasswordPolicyResponsePasswordPolicy) SetHardExpire(v bool) *SetPasswordPolicyResponsePasswordPolicy {
	s.HardExpire = &v
	return s
}

func (s *SetPasswordPolicyResponsePasswordPolicy) SetRequireLowercaseCharacters(v bool) *SetPasswordPolicyResponsePasswordPolicy {
	s.RequireLowercaseCharacters = &v
	return s
}

func (s *SetPasswordPolicyResponsePasswordPolicy) SetRequireNumbers(v bool) *SetPasswordPolicyResponsePasswordPolicy {
	s.RequireNumbers = &v
	return s
}

func (s *SetPasswordPolicyResponsePasswordPolicy) SetMaxLoginAttemps(v int) *SetPasswordPolicyResponsePasswordPolicy {
	s.MaxLoginAttemps = &v
	return s
}

func (s *SetPasswordPolicyResponsePasswordPolicy) SetMaxPasswordAge(v int) *SetPasswordPolicyResponsePasswordPolicy {
	s.MaxPasswordAge = &v
	return s
}

func (s *SetPasswordPolicyResponsePasswordPolicy) SetPasswordNotContainUserName(v bool) *SetPasswordPolicyResponsePasswordPolicy {
	s.PasswordNotContainUserName = &v
	return s
}

func (s *SetPasswordPolicyResponsePasswordPolicy) SetPasswordReusePrevention(v int) *SetPasswordPolicyResponsePasswordPolicy {
	s.PasswordReusePrevention = &v
	return s
}

func (s *SetPasswordPolicyResponsePasswordPolicy) SetRequireUppercaseCharacters(v bool) *SetPasswordPolicyResponsePasswordPolicy {
	s.RequireUppercaseCharacters = &v
	return s
}

func (s *SetPasswordPolicyResponsePasswordPolicy) SetMinimumPasswordDifferentCharacter(v int) *SetPasswordPolicyResponsePasswordPolicy {
	s.MinimumPasswordDifferentCharacter = &v
	return s
}

func (s *SetPasswordPolicyResponsePasswordPolicy) SetRequireSymbols(v bool) *SetPasswordPolicyResponsePasswordPolicy {
	s.RequireSymbols = &v
	return s
}

type RemoveUserFromGroupRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	GroupName         *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s RemoveUserFromGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromGroupRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupRequest) SetUserPrincipalName(v string) *RemoveUserFromGroupRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *RemoveUserFromGroupRequest) SetGroupName(v string) *RemoveUserFromGroupRequest {
	s.GroupName = &v
	return s
}

type RemoveUserFromGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RemoveUserFromGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserFromGroupResponse) SetRequestId(v string) *RemoveUserFromGroupResponse {
	s.RequestId = &v
	return s
}

type ListVirtualMFADevicesRequest struct {
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxItems *int    `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListVirtualMFADevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesRequest) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesRequest) SetMarker(v string) *ListVirtualMFADevicesRequest {
	s.Marker = &v
	return s
}

func (s *ListVirtualMFADevicesRequest) SetMaxItems(v int) *ListVirtualMFADevicesRequest {
	s.MaxItems = &v
	return s
}

type ListVirtualMFADevicesResponse struct {
	RequestId         *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	IsTruncated       *bool                                           `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty" require:"true"`
	Marker            *string                                         `json:"Marker,omitempty" xml:"Marker,omitempty" require:"true"`
	VirtualMFADevices *ListVirtualMFADevicesResponseVirtualMFADevices `json:"VirtualMFADevices,omitempty" xml:"VirtualMFADevices,omitempty" require:"true" type:"Struct"`
}

func (s ListVirtualMFADevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponse) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponse) SetRequestId(v string) *ListVirtualMFADevicesResponse {
	s.RequestId = &v
	return s
}

func (s *ListVirtualMFADevicesResponse) SetIsTruncated(v bool) *ListVirtualMFADevicesResponse {
	s.IsTruncated = &v
	return s
}

func (s *ListVirtualMFADevicesResponse) SetMarker(v string) *ListVirtualMFADevicesResponse {
	s.Marker = &v
	return s
}

func (s *ListVirtualMFADevicesResponse) SetVirtualMFADevices(v *ListVirtualMFADevicesResponseVirtualMFADevices) *ListVirtualMFADevicesResponse {
	s.VirtualMFADevices = v
	return s
}

type ListVirtualMFADevicesResponseVirtualMFADevices struct {
	VirtualMFADevice []*ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADevice `json:"VirtualMFADevice,omitempty" xml:"VirtualMFADevice,omitempty" require:"true" type:"Repeated"`
}

func (s ListVirtualMFADevicesResponseVirtualMFADevices) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponseVirtualMFADevices) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseVirtualMFADevices) SetVirtualMFADevice(v []*ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADevice) *ListVirtualMFADevicesResponseVirtualMFADevices {
	s.VirtualMFADevice = v
	return s
}

type ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADevice struct {
	SerialNumber *string                                                             `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty" require:"true"`
	ActivateDate *string                                                             `json:"ActivateDate,omitempty" xml:"ActivateDate,omitempty" require:"true"`
	User         *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADeviceUser `json:"User,omitempty" xml:"User,omitempty" require:"true" type:"Struct"`
}

func (s ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADevice) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADevice) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADevice) SetSerialNumber(v string) *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADevice {
	s.SerialNumber = &v
	return s
}

func (s *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADevice) SetActivateDate(v string) *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADevice {
	s.ActivateDate = &v
	return s
}

func (s *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADevice) SetUser(v *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADeviceUser) *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADevice {
	s.User = v
	return s
}

type ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADeviceUser struct {
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
}

func (s ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADeviceUser) String() string {
	return tea.Prettify(s)
}

func (s ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADeviceUser) GoString() string {
	return s.String()
}

func (s *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADeviceUser) SetUserId(v string) *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADeviceUser {
	s.UserId = &v
	return s
}

func (s *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADeviceUser) SetDisplayName(v string) *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADeviceUser {
	s.DisplayName = &v
	return s
}

func (s *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADeviceUser) SetUserPrincipalName(v string) *ListVirtualMFADevicesResponseVirtualMFADevicesVirtualMFADeviceUser {
	s.UserPrincipalName = &v
	return s
}

type ListUsersForGroupRequest struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Marker    *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxItems  *int    `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListUsersForGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupRequest) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupRequest) SetGroupName(v string) *ListUsersForGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ListUsersForGroupRequest) SetMarker(v string) *ListUsersForGroupRequest {
	s.Marker = &v
	return s
}

func (s *ListUsersForGroupRequest) SetMaxItems(v int) *ListUsersForGroupRequest {
	s.MaxItems = &v
	return s
}

type ListUsersForGroupResponse struct {
	RequestId   *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	IsTruncated *bool                           `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty" require:"true"`
	Marker      *string                         `json:"Marker,omitempty" xml:"Marker,omitempty" require:"true"`
	Users       *ListUsersForGroupResponseUsers `json:"Users,omitempty" xml:"Users,omitempty" require:"true" type:"Struct"`
}

func (s ListUsersForGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponse) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponse) SetRequestId(v string) *ListUsersForGroupResponse {
	s.RequestId = &v
	return s
}

func (s *ListUsersForGroupResponse) SetIsTruncated(v bool) *ListUsersForGroupResponse {
	s.IsTruncated = &v
	return s
}

func (s *ListUsersForGroupResponse) SetMarker(v string) *ListUsersForGroupResponse {
	s.Marker = &v
	return s
}

func (s *ListUsersForGroupResponse) SetUsers(v *ListUsersForGroupResponseUsers) *ListUsersForGroupResponse {
	s.Users = v
	return s
}

type ListUsersForGroupResponseUsers struct {
	User []*ListUsersForGroupResponseUsersUser `json:"User,omitempty" xml:"User,omitempty" require:"true" type:"Repeated"`
}

func (s ListUsersForGroupResponseUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponseUsers) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseUsers) SetUser(v []*ListUsersForGroupResponseUsersUser) *ListUsersForGroupResponseUsers {
	s.User = v
	return s
}

type ListUsersForGroupResponseUsersUser struct {
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	JoinDate          *string `json:"JoinDate,omitempty" xml:"JoinDate,omitempty" require:"true"`
}

func (s ListUsersForGroupResponseUsersUser) String() string {
	return tea.Prettify(s)
}

func (s ListUsersForGroupResponseUsersUser) GoString() string {
	return s.String()
}

func (s *ListUsersForGroupResponseUsersUser) SetUserId(v string) *ListUsersForGroupResponseUsersUser {
	s.UserId = &v
	return s
}

func (s *ListUsersForGroupResponseUsersUser) SetDisplayName(v string) *ListUsersForGroupResponseUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersForGroupResponseUsersUser) SetUserPrincipalName(v string) *ListUsersForGroupResponseUsersUser {
	s.UserPrincipalName = &v
	return s
}

func (s *ListUsersForGroupResponseUsersUser) SetJoinDate(v string) *ListUsersForGroupResponseUsersUser {
	s.JoinDate = &v
	return s
}

type ListUsersRequest struct {
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxItems *int    `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetMarker(v string) *ListUsersRequest {
	s.Marker = &v
	return s
}

func (s *ListUsersRequest) SetMaxItems(v int) *ListUsersRequest {
	s.MaxItems = &v
	return s
}

type ListUsersResponse struct {
	RequestId   *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	IsTruncated *bool                   `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty" require:"true"`
	Marker      *string                 `json:"Marker,omitempty" xml:"Marker,omitempty" require:"true"`
	Users       *ListUsersResponseUsers `json:"Users,omitempty" xml:"Users,omitempty" require:"true" type:"Struct"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetRequestId(v string) *ListUsersResponse {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponse) SetIsTruncated(v bool) *ListUsersResponse {
	s.IsTruncated = &v
	return s
}

func (s *ListUsersResponse) SetMarker(v string) *ListUsersResponse {
	s.Marker = &v
	return s
}

func (s *ListUsersResponse) SetUsers(v *ListUsersResponseUsers) *ListUsersResponse {
	s.Users = v
	return s
}

type ListUsersResponseUsers struct {
	User []*ListUsersResponseUsersUser `json:"User,omitempty" xml:"User,omitempty" require:"true" type:"Repeated"`
}

func (s ListUsersResponseUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseUsers) SetUser(v []*ListUsersResponseUsersUser) *ListUsersResponseUsers {
	s.User = v
	return s
}

type ListUsersResponseUsersUser struct {
	UpdateDate        *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty" require:"true"`
	Comments          *string `json:"Comments,omitempty" xml:"Comments,omitempty" require:"true"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	LastLoginDate     *string `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty" require:"true"`
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	CreateDate        *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	MobilePhone       *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty" require:"true"`
}

func (s ListUsersResponseUsersUser) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseUsersUser) GoString() string {
	return s.String()
}

func (s *ListUsersResponseUsersUser) SetUpdateDate(v string) *ListUsersResponseUsersUser {
	s.UpdateDate = &v
	return s
}

func (s *ListUsersResponseUsersUser) SetEmail(v string) *ListUsersResponseUsersUser {
	s.Email = &v
	return s
}

func (s *ListUsersResponseUsersUser) SetComments(v string) *ListUsersResponseUsersUser {
	s.Comments = &v
	return s
}

func (s *ListUsersResponseUsersUser) SetUserId(v string) *ListUsersResponseUsersUser {
	s.UserId = &v
	return s
}

func (s *ListUsersResponseUsersUser) SetLastLoginDate(v string) *ListUsersResponseUsersUser {
	s.LastLoginDate = &v
	return s
}

func (s *ListUsersResponseUsersUser) SetDisplayName(v string) *ListUsersResponseUsersUser {
	s.DisplayName = &v
	return s
}

func (s *ListUsersResponseUsersUser) SetUserPrincipalName(v string) *ListUsersResponseUsersUser {
	s.UserPrincipalName = &v
	return s
}

func (s *ListUsersResponseUsersUser) SetCreateDate(v string) *ListUsersResponseUsersUser {
	s.CreateDate = &v
	return s
}

func (s *ListUsersResponseUsersUser) SetMobilePhone(v string) *ListUsersResponseUsersUser {
	s.MobilePhone = &v
	return s
}

type ListSAMLProvidersRequest struct {
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxItems *int    `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListSAMLProvidersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSAMLProvidersRequest) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersRequest) SetMarker(v string) *ListSAMLProvidersRequest {
	s.Marker = &v
	return s
}

func (s *ListSAMLProvidersRequest) SetMaxItems(v int) *ListSAMLProvidersRequest {
	s.MaxItems = &v
	return s
}

type ListSAMLProvidersResponse struct {
	RequestId     *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	IsTruncated   *bool                                   `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty" require:"true"`
	Marker        *string                                 `json:"Marker,omitempty" xml:"Marker,omitempty" require:"true"`
	SAMLProviders *ListSAMLProvidersResponseSAMLProviders `json:"SAMLProviders,omitempty" xml:"SAMLProviders,omitempty" require:"true" type:"Struct"`
}

func (s ListSAMLProvidersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSAMLProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersResponse) SetRequestId(v string) *ListSAMLProvidersResponse {
	s.RequestId = &v
	return s
}

func (s *ListSAMLProvidersResponse) SetIsTruncated(v bool) *ListSAMLProvidersResponse {
	s.IsTruncated = &v
	return s
}

func (s *ListSAMLProvidersResponse) SetMarker(v string) *ListSAMLProvidersResponse {
	s.Marker = &v
	return s
}

func (s *ListSAMLProvidersResponse) SetSAMLProviders(v *ListSAMLProvidersResponseSAMLProviders) *ListSAMLProvidersResponse {
	s.SAMLProviders = v
	return s
}

type ListSAMLProvidersResponseSAMLProviders struct {
	SAMLProvider []*ListSAMLProvidersResponseSAMLProvidersSAMLProvider `json:"SAMLProvider,omitempty" xml:"SAMLProvider,omitempty" require:"true" type:"Repeated"`
}

func (s ListSAMLProvidersResponseSAMLProviders) String() string {
	return tea.Prettify(s)
}

func (s ListSAMLProvidersResponseSAMLProviders) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersResponseSAMLProviders) SetSAMLProvider(v []*ListSAMLProvidersResponseSAMLProvidersSAMLProvider) *ListSAMLProvidersResponseSAMLProviders {
	s.SAMLProvider = v
	return s
}

type ListSAMLProvidersResponseSAMLProvidersSAMLProvider struct {
	UpdateDate       *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty" require:"true"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Arn              *string `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
	CreateDate       *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s ListSAMLProvidersResponseSAMLProvidersSAMLProvider) String() string {
	return tea.Prettify(s)
}

func (s ListSAMLProvidersResponseSAMLProvidersSAMLProvider) GoString() string {
	return s.String()
}

func (s *ListSAMLProvidersResponseSAMLProvidersSAMLProvider) SetUpdateDate(v string) *ListSAMLProvidersResponseSAMLProvidersSAMLProvider {
	s.UpdateDate = &v
	return s
}

func (s *ListSAMLProvidersResponseSAMLProvidersSAMLProvider) SetSAMLProviderName(v string) *ListSAMLProvidersResponseSAMLProvidersSAMLProvider {
	s.SAMLProviderName = &v
	return s
}

func (s *ListSAMLProvidersResponseSAMLProvidersSAMLProvider) SetDescription(v string) *ListSAMLProvidersResponseSAMLProvidersSAMLProvider {
	s.Description = &v
	return s
}

func (s *ListSAMLProvidersResponseSAMLProvidersSAMLProvider) SetArn(v string) *ListSAMLProvidersResponseSAMLProvidersSAMLProvider {
	s.Arn = &v
	return s
}

func (s *ListSAMLProvidersResponseSAMLProvidersSAMLProvider) SetCreateDate(v string) *ListSAMLProvidersResponseSAMLProvidersSAMLProvider {
	s.CreateDate = &v
	return s
}

type ListGroupsForUserRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
}

func (s ListGroupsForUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserRequest) SetUserPrincipalName(v string) *ListGroupsForUserRequest {
	s.UserPrincipalName = &v
	return s
}

type ListGroupsForUserResponse struct {
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Groups    *ListGroupsForUserResponseGroups `json:"Groups,omitempty" xml:"Groups,omitempty" require:"true" type:"Struct"`
}

func (s ListGroupsForUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponse) SetRequestId(v string) *ListGroupsForUserResponse {
	s.RequestId = &v
	return s
}

func (s *ListGroupsForUserResponse) SetGroups(v *ListGroupsForUserResponseGroups) *ListGroupsForUserResponse {
	s.Groups = v
	return s
}

type ListGroupsForUserResponseGroups struct {
	Group []*ListGroupsForUserResponseGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" require:"true" type:"Repeated"`
}

func (s ListGroupsForUserResponseGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponseGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseGroups) SetGroup(v []*ListGroupsForUserResponseGroupsGroup) *ListGroupsForUserResponseGroups {
	s.Group = v
	return s
}

type ListGroupsForUserResponseGroupsGroup struct {
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
	Comments    *string `json:"Comments,omitempty" xml:"Comments,omitempty" require:"true"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	JoinDate    *string `json:"JoinDate,omitempty" xml:"JoinDate,omitempty" require:"true"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
}

func (s ListGroupsForUserResponseGroupsGroup) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsForUserResponseGroupsGroup) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseGroupsGroup) SetGroupName(v string) *ListGroupsForUserResponseGroupsGroup {
	s.GroupName = &v
	return s
}

func (s *ListGroupsForUserResponseGroupsGroup) SetComments(v string) *ListGroupsForUserResponseGroupsGroup {
	s.Comments = &v
	return s
}

func (s *ListGroupsForUserResponseGroupsGroup) SetDisplayName(v string) *ListGroupsForUserResponseGroupsGroup {
	s.DisplayName = &v
	return s
}

func (s *ListGroupsForUserResponseGroupsGroup) SetJoinDate(v string) *ListGroupsForUserResponseGroupsGroup {
	s.JoinDate = &v
	return s
}

func (s *ListGroupsForUserResponseGroupsGroup) SetGroupId(v string) *ListGroupsForUserResponseGroupsGroup {
	s.GroupId = &v
	return s
}

type ListGroupsRequest struct {
	Marker   *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxItems *int    `json:"MaxItems,omitempty" xml:"MaxItems,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) SetMarker(v string) *ListGroupsRequest {
	s.Marker = &v
	return s
}

func (s *ListGroupsRequest) SetMaxItems(v int) *ListGroupsRequest {
	s.MaxItems = &v
	return s
}

type ListGroupsResponse struct {
	RequestId   *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	IsTruncated *bool                     `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty" require:"true"`
	Marker      *string                   `json:"Marker,omitempty" xml:"Marker,omitempty" require:"true"`
	Groups      *ListGroupsResponseGroups `json:"Groups,omitempty" xml:"Groups,omitempty" require:"true" type:"Struct"`
}

func (s ListGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListGroupsResponse) SetRequestId(v string) *ListGroupsResponse {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponse) SetIsTruncated(v bool) *ListGroupsResponse {
	s.IsTruncated = &v
	return s
}

func (s *ListGroupsResponse) SetMarker(v string) *ListGroupsResponse {
	s.Marker = &v
	return s
}

func (s *ListGroupsResponse) SetGroups(v *ListGroupsResponseGroups) *ListGroupsResponse {
	s.Groups = v
	return s
}

type ListGroupsResponseGroups struct {
	Group []*ListGroupsResponseGroupsGroup `json:"Group,omitempty" xml:"Group,omitempty" require:"true" type:"Repeated"`
}

func (s ListGroupsResponseGroups) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseGroups) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseGroups) SetGroup(v []*ListGroupsResponseGroupsGroup) *ListGroupsResponseGroups {
	s.Group = v
	return s
}

type ListGroupsResponseGroupsGroup struct {
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
	UpdateDate  *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Comments    *string `json:"Comments,omitempty" xml:"Comments,omitempty" require:"true"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
}

func (s ListGroupsResponseGroupsGroup) String() string {
	return tea.Prettify(s)
}

func (s ListGroupsResponseGroupsGroup) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseGroupsGroup) SetGroupName(v string) *ListGroupsResponseGroupsGroup {
	s.GroupName = &v
	return s
}

func (s *ListGroupsResponseGroupsGroup) SetUpdateDate(v string) *ListGroupsResponseGroupsGroup {
	s.UpdateDate = &v
	return s
}

func (s *ListGroupsResponseGroupsGroup) SetComments(v string) *ListGroupsResponseGroupsGroup {
	s.Comments = &v
	return s
}

func (s *ListGroupsResponseGroupsGroup) SetDisplayName(v string) *ListGroupsResponseGroupsGroup {
	s.DisplayName = &v
	return s
}

func (s *ListGroupsResponseGroupsGroup) SetCreateDate(v string) *ListGroupsResponseGroupsGroup {
	s.CreateDate = &v
	return s
}

func (s *ListGroupsResponseGroupsGroup) SetGroupId(v string) *ListGroupsResponseGroupsGroup {
	s.GroupId = &v
	return s
}

type ListAccessKeysRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s ListAccessKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysRequest) GoString() string {
	return s.String()
}

func (s *ListAccessKeysRequest) SetUserPrincipalName(v string) *ListAccessKeysRequest {
	s.UserPrincipalName = &v
	return s
}

type ListAccessKeysResponse struct {
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AccessKeys *ListAccessKeysResponseAccessKeys `json:"AccessKeys,omitempty" xml:"AccessKeys,omitempty" require:"true" type:"Struct"`
}

func (s ListAccessKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysResponse) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponse) SetRequestId(v string) *ListAccessKeysResponse {
	s.RequestId = &v
	return s
}

func (s *ListAccessKeysResponse) SetAccessKeys(v *ListAccessKeysResponseAccessKeys) *ListAccessKeysResponse {
	s.AccessKeys = v
	return s
}

type ListAccessKeysResponseAccessKeys struct {
	AccessKey []*ListAccessKeysResponseAccessKeysAccessKey `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" require:"true" type:"Repeated"`
}

func (s ListAccessKeysResponseAccessKeys) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysResponseAccessKeys) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponseAccessKeys) SetAccessKey(v []*ListAccessKeysResponseAccessKeysAccessKey) *ListAccessKeysResponseAccessKeys {
	s.AccessKey = v
	return s
}

type ListAccessKeysResponseAccessKeysAccessKey struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	UpdateDate  *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty" require:"true"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s ListAccessKeysResponseAccessKeysAccessKey) String() string {
	return tea.Prettify(s)
}

func (s ListAccessKeysResponseAccessKeysAccessKey) GoString() string {
	return s.String()
}

func (s *ListAccessKeysResponseAccessKeysAccessKey) SetStatus(v string) *ListAccessKeysResponseAccessKeysAccessKey {
	s.Status = &v
	return s
}

func (s *ListAccessKeysResponseAccessKeysAccessKey) SetUpdateDate(v string) *ListAccessKeysResponseAccessKeysAccessKey {
	s.UpdateDate = &v
	return s
}

func (s *ListAccessKeysResponseAccessKeysAccessKey) SetAccessKeyId(v string) *ListAccessKeysResponseAccessKeysAccessKey {
	s.AccessKeyId = &v
	return s
}

func (s *ListAccessKeysResponseAccessKeysAccessKey) SetCreateDate(v string) *ListAccessKeysResponseAccessKeysAccessKey {
	s.CreateDate = &v
	return s
}

type GetUserMFAInfoRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s GetUserMFAInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoRequest) SetUserPrincipalName(v string) *GetUserMFAInfoRequest {
	s.UserPrincipalName = &v
	return s
}

type GetUserMFAInfoResponse struct {
	RequestId   *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	IsMFAEnable *bool                            `json:"IsMFAEnable,omitempty" xml:"IsMFAEnable,omitempty" require:"true"`
	MFADevice   *GetUserMFAInfoResponseMFADevice `json:"MFADevice,omitempty" xml:"MFADevice,omitempty" require:"true" type:"Struct"`
}

func (s GetUserMFAInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoResponse) SetRequestId(v string) *GetUserMFAInfoResponse {
	s.RequestId = &v
	return s
}

func (s *GetUserMFAInfoResponse) SetIsMFAEnable(v bool) *GetUserMFAInfoResponse {
	s.IsMFAEnable = &v
	return s
}

func (s *GetUserMFAInfoResponse) SetMFADevice(v *GetUserMFAInfoResponseMFADevice) *GetUserMFAInfoResponse {
	s.MFADevice = v
	return s
}

type GetUserMFAInfoResponseMFADevice struct {
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty" require:"true"`
}

func (s GetUserMFAInfoResponseMFADevice) String() string {
	return tea.Prettify(s)
}

func (s GetUserMFAInfoResponseMFADevice) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoResponseMFADevice) SetType(v string) *GetUserMFAInfoResponseMFADevice {
	s.Type = &v
	return s
}

func (s *GetUserMFAInfoResponseMFADevice) SetSerialNumber(v string) *GetUserMFAInfoResponseMFADevice {
	s.SerialNumber = &v
	return s
}

type GetUserRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserAccessKeyId   *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetUserPrincipalName(v string) *GetUserRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

func (s *GetUserRequest) SetUserAccessKeyId(v string) *GetUserRequest {
	s.UserAccessKeyId = &v
	return s
}

type GetUserResponse struct {
	RequestId *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	User      *GetUserResponseUser `json:"User,omitempty" xml:"User,omitempty" require:"true" type:"Struct"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetRequestId(v string) *GetUserResponse {
	s.RequestId = &v
	return s
}

func (s *GetUserResponse) SetUser(v *GetUserResponseUser) *GetUserResponse {
	s.User = v
	return s
}

type GetUserResponseUser struct {
	UpdateDate        *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty" require:"true"`
	Comments          *string `json:"Comments,omitempty" xml:"Comments,omitempty" require:"true"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	LastLoginDate     *string `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty" require:"true"`
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	CreateDate        *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	MobilePhone       *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty" require:"true"`
}

func (s GetUserResponseUser) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponseUser) GoString() string {
	return s.String()
}

func (s *GetUserResponseUser) SetUpdateDate(v string) *GetUserResponseUser {
	s.UpdateDate = &v
	return s
}

func (s *GetUserResponseUser) SetEmail(v string) *GetUserResponseUser {
	s.Email = &v
	return s
}

func (s *GetUserResponseUser) SetComments(v string) *GetUserResponseUser {
	s.Comments = &v
	return s
}

func (s *GetUserResponseUser) SetUserId(v string) *GetUserResponseUser {
	s.UserId = &v
	return s
}

func (s *GetUserResponseUser) SetLastLoginDate(v string) *GetUserResponseUser {
	s.LastLoginDate = &v
	return s
}

func (s *GetUserResponseUser) SetDisplayName(v string) *GetUserResponseUser {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseUser) SetUserPrincipalName(v string) *GetUserResponseUser {
	s.UserPrincipalName = &v
	return s
}

func (s *GetUserResponseUser) SetCreateDate(v string) *GetUserResponseUser {
	s.CreateDate = &v
	return s
}

func (s *GetUserResponseUser) SetMobilePhone(v string) *GetUserResponseUser {
	s.MobilePhone = &v
	return s
}

type GetSecurityPreferenceRequest struct {
}

func (s GetSecurityPreferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceRequest) GoString() string {
	return s.String()
}

type GetSecurityPreferenceResponse struct {
	RequestId          *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SecurityPreference *GetSecurityPreferenceResponseSecurityPreference `json:"SecurityPreference,omitempty" xml:"SecurityPreference,omitempty" require:"true" type:"Struct"`
}

func (s GetSecurityPreferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponse) SetRequestId(v string) *GetSecurityPreferenceResponse {
	s.RequestId = &v
	return s
}

func (s *GetSecurityPreferenceResponse) SetSecurityPreference(v *GetSecurityPreferenceResponseSecurityPreference) *GetSecurityPreferenceResponse {
	s.SecurityPreference = v
	return s
}

type GetSecurityPreferenceResponseSecurityPreference struct {
	LoginProfilePreference *GetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference `json:"LoginProfilePreference,omitempty" xml:"LoginProfilePreference,omitempty" require:"true" type:"Struct"`
	AccessKeyPreference    *GetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference    `json:"AccessKeyPreference,omitempty" xml:"AccessKeyPreference,omitempty" require:"true" type:"Struct"`
	MFAPreference          *GetSecurityPreferenceResponseSecurityPreferenceMFAPreference          `json:"MFAPreference,omitempty" xml:"MFAPreference,omitempty" require:"true" type:"Struct"`
}

func (s GetSecurityPreferenceResponseSecurityPreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseSecurityPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseSecurityPreference) SetLoginProfilePreference(v *GetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) *GetSecurityPreferenceResponseSecurityPreference {
	s.LoginProfilePreference = v
	return s
}

func (s *GetSecurityPreferenceResponseSecurityPreference) SetAccessKeyPreference(v *GetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference) *GetSecurityPreferenceResponseSecurityPreference {
	s.AccessKeyPreference = v
	return s
}

func (s *GetSecurityPreferenceResponseSecurityPreference) SetMFAPreference(v *GetSecurityPreferenceResponseSecurityPreferenceMFAPreference) *GetSecurityPreferenceResponseSecurityPreference {
	s.MFAPreference = v
	return s
}

type GetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference struct {
	LoginSessionDuration      *int    `json:"LoginSessionDuration,omitempty" xml:"LoginSessionDuration,omitempty" require:"true"`
	LoginNetworkMasks         *string `json:"LoginNetworkMasks,omitempty" xml:"LoginNetworkMasks,omitempty" require:"true"`
	AllowUserToChangePassword *bool   `json:"AllowUserToChangePassword,omitempty" xml:"AllowUserToChangePassword,omitempty" require:"true"`
	EnableSaveMFATicket       *bool   `json:"EnableSaveMFATicket,omitempty" xml:"EnableSaveMFATicket,omitempty" require:"true"`
}

func (s GetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) SetLoginSessionDuration(v int) *GetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference {
	s.LoginSessionDuration = &v
	return s
}

func (s *GetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) SetLoginNetworkMasks(v string) *GetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference {
	s.LoginNetworkMasks = &v
	return s
}

func (s *GetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) SetAllowUserToChangePassword(v bool) *GetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference {
	s.AllowUserToChangePassword = &v
	return s
}

func (s *GetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference) SetEnableSaveMFATicket(v bool) *GetSecurityPreferenceResponseSecurityPreferenceLoginProfilePreference {
	s.EnableSaveMFATicket = &v
	return s
}

type GetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference struct {
	AllowUserToManageAccessKeys *bool `json:"AllowUserToManageAccessKeys,omitempty" xml:"AllowUserToManageAccessKeys,omitempty" require:"true"`
}

func (s GetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference) SetAllowUserToManageAccessKeys(v bool) *GetSecurityPreferenceResponseSecurityPreferenceAccessKeyPreference {
	s.AllowUserToManageAccessKeys = &v
	return s
}

type GetSecurityPreferenceResponseSecurityPreferenceMFAPreference struct {
	AllowUserToManageMFADevices *bool `json:"AllowUserToManageMFADevices,omitempty" xml:"AllowUserToManageMFADevices,omitempty" require:"true"`
}

func (s GetSecurityPreferenceResponseSecurityPreferenceMFAPreference) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityPreferenceResponseSecurityPreferenceMFAPreference) GoString() string {
	return s.String()
}

func (s *GetSecurityPreferenceResponseSecurityPreferenceMFAPreference) SetAllowUserToManageMFADevices(v bool) *GetSecurityPreferenceResponseSecurityPreferenceMFAPreference {
	s.AllowUserToManageMFADevices = &v
	return s
}

type GetSAMLProviderRequest struct {
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty" require:"true"`
}

func (s GetSAMLProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSAMLProviderRequest) GoString() string {
	return s.String()
}

func (s *GetSAMLProviderRequest) SetSAMLProviderName(v string) *GetSAMLProviderRequest {
	s.SAMLProviderName = &v
	return s
}

type GetSAMLProviderResponse struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SAMLProvider *GetSAMLProviderResponseSAMLProvider `json:"SAMLProvider,omitempty" xml:"SAMLProvider,omitempty" require:"true" type:"Struct"`
}

func (s GetSAMLProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSAMLProviderResponse) GoString() string {
	return s.String()
}

func (s *GetSAMLProviderResponse) SetRequestId(v string) *GetSAMLProviderResponse {
	s.RequestId = &v
	return s
}

func (s *GetSAMLProviderResponse) SetSAMLProvider(v *GetSAMLProviderResponseSAMLProvider) *GetSAMLProviderResponse {
	s.SAMLProvider = v
	return s
}

type GetSAMLProviderResponseSAMLProvider struct {
	UpdateDate                  *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	SAMLProviderName            *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty" require:"true"`
	Description                 *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	EncodedSAMLMetadataDocument *string `json:"EncodedSAMLMetadataDocument,omitempty" xml:"EncodedSAMLMetadataDocument,omitempty" require:"true"`
	Arn                         *string `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
	CreateDate                  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s GetSAMLProviderResponseSAMLProvider) String() string {
	return tea.Prettify(s)
}

func (s GetSAMLProviderResponseSAMLProvider) GoString() string {
	return s.String()
}

func (s *GetSAMLProviderResponseSAMLProvider) SetUpdateDate(v string) *GetSAMLProviderResponseSAMLProvider {
	s.UpdateDate = &v
	return s
}

func (s *GetSAMLProviderResponseSAMLProvider) SetSAMLProviderName(v string) *GetSAMLProviderResponseSAMLProvider {
	s.SAMLProviderName = &v
	return s
}

func (s *GetSAMLProviderResponseSAMLProvider) SetDescription(v string) *GetSAMLProviderResponseSAMLProvider {
	s.Description = &v
	return s
}

func (s *GetSAMLProviderResponseSAMLProvider) SetEncodedSAMLMetadataDocument(v string) *GetSAMLProviderResponseSAMLProvider {
	s.EncodedSAMLMetadataDocument = &v
	return s
}

func (s *GetSAMLProviderResponseSAMLProvider) SetArn(v string) *GetSAMLProviderResponseSAMLProvider {
	s.Arn = &v
	return s
}

func (s *GetSAMLProviderResponseSAMLProvider) SetCreateDate(v string) *GetSAMLProviderResponseSAMLProvider {
	s.CreateDate = &v
	return s
}

type GetPasswordPolicyRequest struct {
}

func (s GetPasswordPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordPolicyRequest) GoString() string {
	return s.String()
}

type GetPasswordPolicyResponse struct {
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PasswordPolicy *GetPasswordPolicyResponsePasswordPolicy `json:"PasswordPolicy,omitempty" xml:"PasswordPolicy,omitempty" require:"true" type:"Struct"`
}

func (s GetPasswordPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyResponse) SetRequestId(v string) *GetPasswordPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *GetPasswordPolicyResponse) SetPasswordPolicy(v *GetPasswordPolicyResponsePasswordPolicy) *GetPasswordPolicyResponse {
	s.PasswordPolicy = v
	return s
}

type GetPasswordPolicyResponsePasswordPolicy struct {
	MinimumPasswordLength             *int  `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty" require:"true"`
	HardExpire                        *bool `json:"HardExpire,omitempty" xml:"HardExpire,omitempty" require:"true"`
	RequireLowercaseCharacters        *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty" require:"true"`
	RequireNumbers                    *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty" require:"true"`
	MaxLoginAttemps                   *int  `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty" require:"true"`
	MaxPasswordAge                    *int  `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty" require:"true"`
	PasswordNotContainUserName        *bool `json:"PasswordNotContainUserName,omitempty" xml:"PasswordNotContainUserName,omitempty" require:"true"`
	PasswordReusePrevention           *int  `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty" require:"true"`
	RequireUppercaseCharacters        *bool `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty" require:"true"`
	MinimumPasswordDifferentCharacter *int  `json:"MinimumPasswordDifferentCharacter,omitempty" xml:"MinimumPasswordDifferentCharacter,omitempty" require:"true"`
	RequireSymbols                    *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty" require:"true"`
}

func (s GetPasswordPolicyResponsePasswordPolicy) String() string {
	return tea.Prettify(s)
}

func (s GetPasswordPolicyResponsePasswordPolicy) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyResponsePasswordPolicy) SetMinimumPasswordLength(v int) *GetPasswordPolicyResponsePasswordPolicy {
	s.MinimumPasswordLength = &v
	return s
}

func (s *GetPasswordPolicyResponsePasswordPolicy) SetHardExpire(v bool) *GetPasswordPolicyResponsePasswordPolicy {
	s.HardExpire = &v
	return s
}

func (s *GetPasswordPolicyResponsePasswordPolicy) SetRequireLowercaseCharacters(v bool) *GetPasswordPolicyResponsePasswordPolicy {
	s.RequireLowercaseCharacters = &v
	return s
}

func (s *GetPasswordPolicyResponsePasswordPolicy) SetRequireNumbers(v bool) *GetPasswordPolicyResponsePasswordPolicy {
	s.RequireNumbers = &v
	return s
}

func (s *GetPasswordPolicyResponsePasswordPolicy) SetMaxLoginAttemps(v int) *GetPasswordPolicyResponsePasswordPolicy {
	s.MaxLoginAttemps = &v
	return s
}

func (s *GetPasswordPolicyResponsePasswordPolicy) SetMaxPasswordAge(v int) *GetPasswordPolicyResponsePasswordPolicy {
	s.MaxPasswordAge = &v
	return s
}

func (s *GetPasswordPolicyResponsePasswordPolicy) SetPasswordNotContainUserName(v bool) *GetPasswordPolicyResponsePasswordPolicy {
	s.PasswordNotContainUserName = &v
	return s
}

func (s *GetPasswordPolicyResponsePasswordPolicy) SetPasswordReusePrevention(v int) *GetPasswordPolicyResponsePasswordPolicy {
	s.PasswordReusePrevention = &v
	return s
}

func (s *GetPasswordPolicyResponsePasswordPolicy) SetRequireUppercaseCharacters(v bool) *GetPasswordPolicyResponsePasswordPolicy {
	s.RequireUppercaseCharacters = &v
	return s
}

func (s *GetPasswordPolicyResponsePasswordPolicy) SetMinimumPasswordDifferentCharacter(v int) *GetPasswordPolicyResponsePasswordPolicy {
	s.MinimumPasswordDifferentCharacter = &v
	return s
}

func (s *GetPasswordPolicyResponsePasswordPolicy) SetRequireSymbols(v bool) *GetPasswordPolicyResponsePasswordPolicy {
	s.RequireSymbols = &v
	return s
}

type GetLoginProfileRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
}

func (s GetLoginProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *GetLoginProfileRequest) SetUserPrincipalName(v string) *GetLoginProfileRequest {
	s.UserPrincipalName = &v
	return s
}

type GetLoginProfileResponse struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LoginProfile *GetLoginProfileResponseLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" require:"true" type:"Struct"`
}

func (s GetLoginProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *GetLoginProfileResponse) SetRequestId(v string) *GetLoginProfileResponse {
	s.RequestId = &v
	return s
}

func (s *GetLoginProfileResponse) SetLoginProfile(v *GetLoginProfileResponseLoginProfile) *GetLoginProfileResponse {
	s.LoginProfile = v
	return s
}

type GetLoginProfileResponseLoginProfile struct {
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	LastLoginTime         *string `json:"LastLoginTime,omitempty" xml:"LastLoginTime,omitempty" require:"true"`
	UpdateDate            *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	PasswordResetRequired *bool   `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty" require:"true"`
	UserPrincipalName     *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	MFABindRequired       *bool   `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty" require:"true"`
}

func (s GetLoginProfileResponseLoginProfile) String() string {
	return tea.Prettify(s)
}

func (s GetLoginProfileResponseLoginProfile) GoString() string {
	return s.String()
}

func (s *GetLoginProfileResponseLoginProfile) SetStatus(v string) *GetLoginProfileResponseLoginProfile {
	s.Status = &v
	return s
}

func (s *GetLoginProfileResponseLoginProfile) SetLastLoginTime(v string) *GetLoginProfileResponseLoginProfile {
	s.LastLoginTime = &v
	return s
}

func (s *GetLoginProfileResponseLoginProfile) SetUpdateDate(v string) *GetLoginProfileResponseLoginProfile {
	s.UpdateDate = &v
	return s
}

func (s *GetLoginProfileResponseLoginProfile) SetPasswordResetRequired(v bool) *GetLoginProfileResponseLoginProfile {
	s.PasswordResetRequired = &v
	return s
}

func (s *GetLoginProfileResponseLoginProfile) SetUserPrincipalName(v string) *GetLoginProfileResponseLoginProfile {
	s.UserPrincipalName = &v
	return s
}

func (s *GetLoginProfileResponseLoginProfile) SetMFABindRequired(v bool) *GetLoginProfileResponseLoginProfile {
	s.MFABindRequired = &v
	return s
}

type GetGroupRequest struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s GetGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupRequest) GoString() string {
	return s.String()
}

func (s *GetGroupRequest) SetGroupName(v string) *GetGroupRequest {
	s.GroupName = &v
	return s
}

type GetGroupResponse struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Group     *GetGroupResponseGroup `json:"Group,omitempty" xml:"Group,omitempty" require:"true" type:"Struct"`
}

func (s GetGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponse) GoString() string {
	return s.String()
}

func (s *GetGroupResponse) SetRequestId(v string) *GetGroupResponse {
	s.RequestId = &v
	return s
}

func (s *GetGroupResponse) SetGroup(v *GetGroupResponseGroup) *GetGroupResponse {
	s.Group = v
	return s
}

type GetGroupResponseGroup struct {
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
	UpdateDate  *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Comments    *string `json:"Comments,omitempty" xml:"Comments,omitempty" require:"true"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
}

func (s GetGroupResponseGroup) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponseGroup) GoString() string {
	return s.String()
}

func (s *GetGroupResponseGroup) SetGroupName(v string) *GetGroupResponseGroup {
	s.GroupName = &v
	return s
}

func (s *GetGroupResponseGroup) SetUpdateDate(v string) *GetGroupResponseGroup {
	s.UpdateDate = &v
	return s
}

func (s *GetGroupResponseGroup) SetComments(v string) *GetGroupResponseGroup {
	s.Comments = &v
	return s
}

func (s *GetGroupResponseGroup) SetDisplayName(v string) *GetGroupResponseGroup {
	s.DisplayName = &v
	return s
}

func (s *GetGroupResponseGroup) SetCreateDate(v string) *GetGroupResponseGroup {
	s.CreateDate = &v
	return s
}

func (s *GetGroupResponseGroup) SetGroupId(v string) *GetGroupResponseGroup {
	s.GroupId = &v
	return s
}

type GetDefaultDomainRequest struct {
}

func (s GetDefaultDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultDomainRequest) GoString() string {
	return s.String()
}

type GetDefaultDomainResponse struct {
	DefaultDomainName *string `json:"DefaultDomainName,omitempty" xml:"DefaultDomainName,omitempty" require:"true"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s GetDefaultDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultDomainResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultDomainResponse) SetDefaultDomainName(v string) *GetDefaultDomainResponse {
	s.DefaultDomainName = &v
	return s
}

func (s *GetDefaultDomainResponse) SetRequestId(v string) *GetDefaultDomainResponse {
	s.RequestId = &v
	return s
}

type GetAccountSummaryRequest struct {
}

func (s GetAccountSummaryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSummaryRequest) GoString() string {
	return s.String()
}

type GetAccountSummaryResponse struct {
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SummaryMap *GetAccountSummaryResponseSummaryMap `json:"SummaryMap,omitempty" xml:"SummaryMap,omitempty" require:"true" type:"Struct"`
}

func (s GetAccountSummaryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetAccountSummaryResponse) SetRequestId(v string) *GetAccountSummaryResponse {
	s.RequestId = &v
	return s
}

func (s *GetAccountSummaryResponse) SetSummaryMap(v *GetAccountSummaryResponseSummaryMap) *GetAccountSummaryResponse {
	s.SummaryMap = v
	return s
}

type GetAccountSummaryResponseSummaryMap struct {
	Policies                            *int `json:"Policies,omitempty" xml:"Policies,omitempty" require:"true"`
	GroupsPerUserQuota                  *int `json:"GroupsPerUserQuota,omitempty" xml:"GroupsPerUserQuota,omitempty" require:"true"`
	AttachedPoliciesPerUserQuota        *int `json:"AttachedPoliciesPerUserQuota,omitempty" xml:"AttachedPoliciesPerUserQuota,omitempty" require:"true"`
	Roles                               *int `json:"Roles,omitempty" xml:"Roles,omitempty" require:"true"`
	Users                               *int `json:"Users,omitempty" xml:"Users,omitempty" require:"true"`
	RolesQuota                          *int `json:"RolesQuota,omitempty" xml:"RolesQuota,omitempty" require:"true"`
	VirtualMFADevicesQuota              *int `json:"VirtualMFADevicesQuota,omitempty" xml:"VirtualMFADevicesQuota,omitempty" require:"true"`
	PoliciesQuota                       *int `json:"PoliciesQuota,omitempty" xml:"PoliciesQuota,omitempty" require:"true"`
	AttachedSystemPoliciesPerGroupQuota *int `json:"AttachedSystemPoliciesPerGroupQuota,omitempty" xml:"AttachedSystemPoliciesPerGroupQuota,omitempty" require:"true"`
	MFADevicesInUse                     *int `json:"MFADevicesInUse,omitempty" xml:"MFADevicesInUse,omitempty" require:"true"`
	AccessKeysPerUserQuota              *int `json:"AccessKeysPerUserQuota,omitempty" xml:"AccessKeysPerUserQuota,omitempty" require:"true"`
	VersionsPerPolicyQuota              *int `json:"VersionsPerPolicyQuota,omitempty" xml:"VersionsPerPolicyQuota,omitempty" require:"true"`
	PolicySizeQuota                     *int `json:"PolicySizeQuota,omitempty" xml:"PolicySizeQuota,omitempty" require:"true"`
	AttachedPoliciesPerGroupQuota       *int `json:"AttachedPoliciesPerGroupQuota,omitempty" xml:"AttachedPoliciesPerGroupQuota,omitempty" require:"true"`
	Groups                              *int `json:"Groups,omitempty" xml:"Groups,omitempty" require:"true"`
	AttachedSystemPoliciesPerUserQuota  *int `json:"AttachedSystemPoliciesPerUserQuota,omitempty" xml:"AttachedSystemPoliciesPerUserQuota,omitempty" require:"true"`
	UsersQuota                          *int `json:"UsersQuota,omitempty" xml:"UsersQuota,omitempty" require:"true"`
	AttachedPoliciesPerRoleQuota        *int `json:"AttachedPoliciesPerRoleQuota,omitempty" xml:"AttachedPoliciesPerRoleQuota,omitempty" require:"true"`
	AttachedSystemPoliciesPerRoleQuota  *int `json:"AttachedSystemPoliciesPerRoleQuota,omitempty" xml:"AttachedSystemPoliciesPerRoleQuota,omitempty" require:"true"`
	MFADevices                          *int `json:"MFADevices,omitempty" xml:"MFADevices,omitempty" require:"true"`
	GroupsQuota                         *int `json:"GroupsQuota,omitempty" xml:"GroupsQuota,omitempty" require:"true"`
}

func (s GetAccountSummaryResponseSummaryMap) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSummaryResponseSummaryMap) GoString() string {
	return s.String()
}

func (s *GetAccountSummaryResponseSummaryMap) SetPolicies(v int) *GetAccountSummaryResponseSummaryMap {
	s.Policies = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetGroupsPerUserQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.GroupsPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetAttachedPoliciesPerUserQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.AttachedPoliciesPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetRoles(v int) *GetAccountSummaryResponseSummaryMap {
	s.Roles = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetUsers(v int) *GetAccountSummaryResponseSummaryMap {
	s.Users = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetRolesQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.RolesQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetVirtualMFADevicesQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.VirtualMFADevicesQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetPoliciesQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.PoliciesQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetAttachedSystemPoliciesPerGroupQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.AttachedSystemPoliciesPerGroupQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetMFADevicesInUse(v int) *GetAccountSummaryResponseSummaryMap {
	s.MFADevicesInUse = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetAccessKeysPerUserQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.AccessKeysPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetVersionsPerPolicyQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.VersionsPerPolicyQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetPolicySizeQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.PolicySizeQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetAttachedPoliciesPerGroupQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.AttachedPoliciesPerGroupQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetGroups(v int) *GetAccountSummaryResponseSummaryMap {
	s.Groups = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetAttachedSystemPoliciesPerUserQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.AttachedSystemPoliciesPerUserQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetUsersQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.UsersQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetAttachedPoliciesPerRoleQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.AttachedPoliciesPerRoleQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetAttachedSystemPoliciesPerRoleQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.AttachedSystemPoliciesPerRoleQuota = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetMFADevices(v int) *GetAccountSummaryResponseSummaryMap {
	s.MFADevices = &v
	return s
}

func (s *GetAccountSummaryResponseSummaryMap) SetGroupsQuota(v int) *GetAccountSummaryResponseSummaryMap {
	s.GroupsQuota = &v
	return s
}

type GetAccountSecurityPracticeReportRequest struct {
}

func (s GetAccountSecurityPracticeReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSecurityPracticeReportRequest) GoString() string {
	return s.String()
}

type GetAccountSecurityPracticeReportResponse struct {
	RequestId                   *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AccountSecurityPracticeInfo *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfo `json:"AccountSecurityPracticeInfo,omitempty" xml:"AccountSecurityPracticeInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetAccountSecurityPracticeReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSecurityPracticeReportResponse) GoString() string {
	return s.String()
}

func (s *GetAccountSecurityPracticeReportResponse) SetRequestId(v string) *GetAccountSecurityPracticeReportResponse {
	s.RequestId = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponse) SetAccountSecurityPracticeInfo(v *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfo) *GetAccountSecurityPracticeReportResponse {
	s.AccountSecurityPracticeInfo = v
	return s
}

type GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfo struct {
	Score                           *int                                                                                                `json:"Score,omitempty" xml:"Score,omitempty" require:"true"`
	AccountSecurityPracticeUserInfo *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo `json:"AccountSecurityPracticeUserInfo,omitempty" xml:"AccountSecurityPracticeUserInfo,omitempty" require:"true" type:"Struct"`
}

func (s GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfo) GoString() string {
	return s.String()
}

func (s *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfo) SetScore(v int) *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfo {
	s.Score = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfo) SetAccountSecurityPracticeUserInfo(v *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfo {
	s.AccountSecurityPracticeUserInfo = v
	return s
}

type GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo struct {
	SubUser                    *int    `json:"SubUser,omitempty" xml:"SubUser,omitempty" require:"true"`
	SubUserBindMfa             *int    `json:"SubUserBindMfa,omitempty" xml:"SubUserBindMfa,omitempty" require:"true"`
	SubUserWithUnusedAccessKey *int    `json:"SubUserWithUnusedAccessKey,omitempty" xml:"SubUserWithUnusedAccessKey,omitempty" require:"true"`
	RootWithAccessKey          *int    `json:"RootWithAccessKey,omitempty" xml:"RootWithAccessKey,omitempty" require:"true"`
	SubUserWithOldAccessKey    *int    `json:"SubUserWithOldAccessKey,omitempty" xml:"SubUserWithOldAccessKey,omitempty" require:"true"`
	SubUserPwdLevel            *string `json:"SubUserPwdLevel,omitempty" xml:"SubUserPwdLevel,omitempty" require:"true"`
	OldAkNum                   *int    `json:"OldAkNum,omitempty" xml:"OldAkNum,omitempty" require:"true"`
	UnusedAkNum                *int    `json:"UnusedAkNum,omitempty" xml:"UnusedAkNum,omitempty" require:"true"`
	BindMfa                    *bool   `json:"BindMfa,omitempty" xml:"BindMfa,omitempty" require:"true"`
}

func (s GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) String() string {
	return tea.Prettify(s)
}

func (s GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) GoString() string {
	return s.String()
}

func (s *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUser(v int) *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUser = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUserBindMfa(v int) *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUserBindMfa = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUserWithUnusedAccessKey(v int) *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUserWithUnusedAccessKey = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetRootWithAccessKey(v int) *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.RootWithAccessKey = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUserWithOldAccessKey(v int) *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUserWithOldAccessKey = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetSubUserPwdLevel(v string) *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.SubUserPwdLevel = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetOldAkNum(v int) *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.OldAkNum = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetUnusedAkNum(v int) *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.UnusedAkNum = &v
	return s
}

func (s *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo) SetBindMfa(v bool) *GetAccountSecurityPracticeReportResponseAccountSecurityPracticeInfoAccountSecurityPracticeUserInfo {
	s.BindMfa = &v
	return s
}

type GetAccountMFAInfoRequest struct {
}

func (s GetAccountMFAInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountMFAInfoRequest) GoString() string {
	return s.String()
}

type GetAccountMFAInfoResponse struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	IsMFAEnable *bool   `json:"IsMFAEnable,omitempty" xml:"IsMFAEnable,omitempty" require:"true"`
}

func (s GetAccountMFAInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountMFAInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAccountMFAInfoResponse) SetRequestId(v string) *GetAccountMFAInfoResponse {
	s.RequestId = &v
	return s
}

func (s *GetAccountMFAInfoResponse) SetIsMFAEnable(v bool) *GetAccountMFAInfoResponse {
	s.IsMFAEnable = &v
	return s
}

type GetAccessKeyLastUsedRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
	UserAccessKeyId   *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty" require:"true"`
}

func (s GetAccessKeyLastUsedRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedRequest) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedRequest) SetUserPrincipalName(v string) *GetAccessKeyLastUsedRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *GetAccessKeyLastUsedRequest) SetUserAccessKeyId(v string) *GetAccessKeyLastUsedRequest {
	s.UserAccessKeyId = &v
	return s
}

type GetAccessKeyLastUsedResponse struct {
	RequestId         *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AccessKeyLastUsed *GetAccessKeyLastUsedResponseAccessKeyLastUsed `json:"AccessKeyLastUsed,omitempty" xml:"AccessKeyLastUsed,omitempty" require:"true" type:"Struct"`
}

func (s GetAccessKeyLastUsedResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedResponse) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResponse) SetRequestId(v string) *GetAccessKeyLastUsedResponse {
	s.RequestId = &v
	return s
}

func (s *GetAccessKeyLastUsedResponse) SetAccessKeyLastUsed(v *GetAccessKeyLastUsedResponseAccessKeyLastUsed) *GetAccessKeyLastUsedResponse {
	s.AccessKeyLastUsed = v
	return s
}

type GetAccessKeyLastUsedResponseAccessKeyLastUsed struct {
	LastUsedDate *string `json:"LastUsedDate,omitempty" xml:"LastUsedDate,omitempty" require:"true"`
}

func (s GetAccessKeyLastUsedResponseAccessKeyLastUsed) String() string {
	return tea.Prettify(s)
}

func (s GetAccessKeyLastUsedResponseAccessKeyLastUsed) GoString() string {
	return s.String()
}

func (s *GetAccessKeyLastUsedResponseAccessKeyLastUsed) SetLastUsedDate(v string) *GetAccessKeyLastUsedResponseAccessKeyLastUsed {
	s.LastUsedDate = &v
	return s
}

type DisableVirtualMFARequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
}

func (s DisableVirtualMFARequest) String() string {
	return tea.Prettify(s)
}

func (s DisableVirtualMFARequest) GoString() string {
	return s.String()
}

func (s *DisableVirtualMFARequest) SetUserPrincipalName(v string) *DisableVirtualMFARequest {
	s.UserPrincipalName = &v
	return s
}

type DisableVirtualMFAResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DisableVirtualMFAResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableVirtualMFAResponse) GoString() string {
	return s.String()
}

func (s *DisableVirtualMFAResponse) SetRequestId(v string) *DisableVirtualMFAResponse {
	s.RequestId = &v
	return s
}

type DeleteVirtualMFADeviceRequest struct {
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty" require:"true"`
}

func (s DeleteVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceRequest) SetSerialNumber(v string) *DeleteVirtualMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

type DeleteVirtualMFADeviceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualMFADeviceResponse) SetRequestId(v string) *DeleteVirtualMFADeviceResponse {
	s.RequestId = &v
	return s
}

type DeleteUserRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetUserPrincipalName(v string) *DeleteUserRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
	return s
}

type DeleteUserResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetRequestId(v string) *DeleteUserResponse {
	s.RequestId = &v
	return s
}

type DeleteSAMLProviderRequest struct {
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty" require:"true"`
}

func (s DeleteSAMLProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSAMLProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteSAMLProviderRequest) SetSAMLProviderName(v string) *DeleteSAMLProviderRequest {
	s.SAMLProviderName = &v
	return s
}

type DeleteSAMLProviderResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteSAMLProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSAMLProviderResponse) GoString() string {
	return s.String()
}

func (s *DeleteSAMLProviderResponse) SetRequestId(v string) *DeleteSAMLProviderResponse {
	s.RequestId = &v
	return s
}

type DeleteLoginProfileRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
}

func (s DeleteLoginProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoginProfileRequest) SetUserPrincipalName(v string) *DeleteLoginProfileRequest {
	s.UserPrincipalName = &v
	return s
}

type DeleteLoginProfileResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteLoginProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoginProfileResponse) SetRequestId(v string) *DeleteLoginProfileResponse {
	s.RequestId = &v
	return s
}

type DeleteGroupRequest struct {
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) SetGroupName(v string) *DeleteGroupRequest {
	s.GroupName = &v
	return s
}

type DeleteGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponse) SetRequestId(v string) *DeleteGroupResponse {
	s.RequestId = &v
	return s
}

type DeleteAccessKeyRequest struct {
	UserAccessKeyId   *string `json:"UserAccessKeyId,omitempty" xml:"UserAccessKeyId,omitempty" require:"true"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s DeleteAccessKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyRequest) SetUserAccessKeyId(v string) *DeleteAccessKeyRequest {
	s.UserAccessKeyId = &v
	return s
}

func (s *DeleteAccessKeyRequest) SetUserPrincipalName(v string) *DeleteAccessKeyRequest {
	s.UserPrincipalName = &v
	return s
}

type DeleteAccessKeyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteAccessKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyResponse) SetRequestId(v string) *DeleteAccessKeyResponse {
	s.RequestId = &v
	return s
}

type CreateVirtualMFADeviceRequest struct {
	VirtualMFADeviceName *string `json:"VirtualMFADeviceName,omitempty" xml:"VirtualMFADeviceName,omitempty" require:"true"`
}

func (s CreateVirtualMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceRequest) SetVirtualMFADeviceName(v string) *CreateVirtualMFADeviceRequest {
	s.VirtualMFADeviceName = &v
	return s
}

type CreateVirtualMFADeviceResponse struct {
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VirtualMFADevice *CreateVirtualMFADeviceResponseVirtualMFADevice `json:"VirtualMFADevice,omitempty" xml:"VirtualMFADevice,omitempty" require:"true" type:"Struct"`
}

func (s CreateVirtualMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceResponse) SetRequestId(v string) *CreateVirtualMFADeviceResponse {
	s.RequestId = &v
	return s
}

func (s *CreateVirtualMFADeviceResponse) SetVirtualMFADevice(v *CreateVirtualMFADeviceResponseVirtualMFADevice) *CreateVirtualMFADeviceResponse {
	s.VirtualMFADevice = v
	return s
}

type CreateVirtualMFADeviceResponseVirtualMFADevice struct {
	SerialNumber     *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty" require:"true"`
	QRCodePNG        *string `json:"QRCodePNG,omitempty" xml:"QRCodePNG,omitempty" require:"true"`
	Base32StringSeed *string `json:"Base32StringSeed,omitempty" xml:"Base32StringSeed,omitempty" require:"true"`
}

func (s CreateVirtualMFADeviceResponseVirtualMFADevice) String() string {
	return tea.Prettify(s)
}

func (s CreateVirtualMFADeviceResponseVirtualMFADevice) GoString() string {
	return s.String()
}

func (s *CreateVirtualMFADeviceResponseVirtualMFADevice) SetSerialNumber(v string) *CreateVirtualMFADeviceResponseVirtualMFADevice {
	s.SerialNumber = &v
	return s
}

func (s *CreateVirtualMFADeviceResponseVirtualMFADevice) SetQRCodePNG(v string) *CreateVirtualMFADeviceResponseVirtualMFADevice {
	s.QRCodePNG = &v
	return s
}

func (s *CreateVirtualMFADeviceResponseVirtualMFADevice) SetBase32StringSeed(v string) *CreateVirtualMFADeviceResponseVirtualMFADevice {
	s.Base32StringSeed = &v
	return s
}

type CreateUserRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	MobilePhone       *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty"`
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Comments          *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetUserPrincipalName(v string) *CreateUserRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetMobilePhone(v string) *CreateUserRequest {
	s.MobilePhone = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetComments(v string) *CreateUserRequest {
	s.Comments = &v
	return s
}

type CreateUserResponse struct {
	RequestId *string                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	User      *CreateUserResponseUser `json:"User,omitempty" xml:"User,omitempty" require:"true" type:"Struct"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetRequestId(v string) *CreateUserResponse {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponse) SetUser(v *CreateUserResponseUser) *CreateUserResponse {
	s.User = v
	return s
}

type CreateUserResponseUser struct {
	UpdateDate        *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty" require:"true"`
	Comments          *string `json:"Comments,omitempty" xml:"Comments,omitempty" require:"true"`
	UserId            *string `json:"UserId,omitempty" xml:"UserId,omitempty" require:"true"`
	LastLoginDate     *string `json:"LastLoginDate,omitempty" xml:"LastLoginDate,omitempty" require:"true"`
	DisplayName       *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	CreateDate        *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	MobilePhone       *string `json:"MobilePhone,omitempty" xml:"MobilePhone,omitempty" require:"true"`
}

func (s CreateUserResponseUser) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseUser) GoString() string {
	return s.String()
}

func (s *CreateUserResponseUser) SetUpdateDate(v string) *CreateUserResponseUser {
	s.UpdateDate = &v
	return s
}

func (s *CreateUserResponseUser) SetEmail(v string) *CreateUserResponseUser {
	s.Email = &v
	return s
}

func (s *CreateUserResponseUser) SetComments(v string) *CreateUserResponseUser {
	s.Comments = &v
	return s
}

func (s *CreateUserResponseUser) SetUserId(v string) *CreateUserResponseUser {
	s.UserId = &v
	return s
}

func (s *CreateUserResponseUser) SetLastLoginDate(v string) *CreateUserResponseUser {
	s.LastLoginDate = &v
	return s
}

func (s *CreateUserResponseUser) SetDisplayName(v string) *CreateUserResponseUser {
	s.DisplayName = &v
	return s
}

func (s *CreateUserResponseUser) SetUserPrincipalName(v string) *CreateUserResponseUser {
	s.UserPrincipalName = &v
	return s
}

func (s *CreateUserResponseUser) SetCreateDate(v string) *CreateUserResponseUser {
	s.CreateDate = &v
	return s
}

func (s *CreateUserResponseUser) SetMobilePhone(v string) *CreateUserResponseUser {
	s.MobilePhone = &v
	return s
}

type CreateSAMLProviderRequest struct {
	SAMLProviderName            *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty" require:"true"`
	Description                 *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EncodedSAMLMetadataDocument *string `json:"EncodedSAMLMetadataDocument,omitempty" xml:"EncodedSAMLMetadataDocument,omitempty"`
}

func (s CreateSAMLProviderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSAMLProviderRequest) GoString() string {
	return s.String()
}

func (s *CreateSAMLProviderRequest) SetSAMLProviderName(v string) *CreateSAMLProviderRequest {
	s.SAMLProviderName = &v
	return s
}

func (s *CreateSAMLProviderRequest) SetDescription(v string) *CreateSAMLProviderRequest {
	s.Description = &v
	return s
}

func (s *CreateSAMLProviderRequest) SetEncodedSAMLMetadataDocument(v string) *CreateSAMLProviderRequest {
	s.EncodedSAMLMetadataDocument = &v
	return s
}

type CreateSAMLProviderResponse struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SAMLProvider *CreateSAMLProviderResponseSAMLProvider `json:"SAMLProvider,omitempty" xml:"SAMLProvider,omitempty" require:"true" type:"Struct"`
}

func (s CreateSAMLProviderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSAMLProviderResponse) GoString() string {
	return s.String()
}

func (s *CreateSAMLProviderResponse) SetRequestId(v string) *CreateSAMLProviderResponse {
	s.RequestId = &v
	return s
}

func (s *CreateSAMLProviderResponse) SetSAMLProvider(v *CreateSAMLProviderResponseSAMLProvider) *CreateSAMLProviderResponse {
	s.SAMLProvider = v
	return s
}

type CreateSAMLProviderResponseSAMLProvider struct {
	UpdateDate       *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	SAMLProviderName *string `json:"SAMLProviderName,omitempty" xml:"SAMLProviderName,omitempty" require:"true"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	Arn              *string `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
	CreateDate       *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s CreateSAMLProviderResponseSAMLProvider) String() string {
	return tea.Prettify(s)
}

func (s CreateSAMLProviderResponseSAMLProvider) GoString() string {
	return s.String()
}

func (s *CreateSAMLProviderResponseSAMLProvider) SetUpdateDate(v string) *CreateSAMLProviderResponseSAMLProvider {
	s.UpdateDate = &v
	return s
}

func (s *CreateSAMLProviderResponseSAMLProvider) SetSAMLProviderName(v string) *CreateSAMLProviderResponseSAMLProvider {
	s.SAMLProviderName = &v
	return s
}

func (s *CreateSAMLProviderResponseSAMLProvider) SetDescription(v string) *CreateSAMLProviderResponseSAMLProvider {
	s.Description = &v
	return s
}

func (s *CreateSAMLProviderResponseSAMLProvider) SetArn(v string) *CreateSAMLProviderResponseSAMLProvider {
	s.Arn = &v
	return s
}

func (s *CreateSAMLProviderResponseSAMLProvider) SetCreateDate(v string) *CreateSAMLProviderResponseSAMLProvider {
	s.CreateDate = &v
	return s
}

type CreateLoginProfileRequest struct {
	UserPrincipalName     *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	Password              *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PasswordResetRequired *bool   `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	MFABindRequired       *bool   `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateLoginProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileRequest) SetUserPrincipalName(v string) *CreateLoginProfileRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *CreateLoginProfileRequest) SetPassword(v string) *CreateLoginProfileRequest {
	s.Password = &v
	return s
}

func (s *CreateLoginProfileRequest) SetPasswordResetRequired(v bool) *CreateLoginProfileRequest {
	s.PasswordResetRequired = &v
	return s
}

func (s *CreateLoginProfileRequest) SetMFABindRequired(v bool) *CreateLoginProfileRequest {
	s.MFABindRequired = &v
	return s
}

func (s *CreateLoginProfileRequest) SetStatus(v string) *CreateLoginProfileRequest {
	s.Status = &v
	return s
}

type CreateLoginProfileResponse struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	LoginProfile *CreateLoginProfileResponseLoginProfile `json:"LoginProfile,omitempty" xml:"LoginProfile,omitempty" require:"true" type:"Struct"`
}

func (s CreateLoginProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoginProfileResponse) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileResponse) SetRequestId(v string) *CreateLoginProfileResponse {
	s.RequestId = &v
	return s
}

func (s *CreateLoginProfileResponse) SetLoginProfile(v *CreateLoginProfileResponseLoginProfile) *CreateLoginProfileResponse {
	s.LoginProfile = v
	return s
}

type CreateLoginProfileResponseLoginProfile struct {
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	UpdateDate            *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	PasswordResetRequired *bool   `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty" require:"true"`
	UserPrincipalName     *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	MFABindRequired       *bool   `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty" require:"true"`
}

func (s CreateLoginProfileResponseLoginProfile) String() string {
	return tea.Prettify(s)
}

func (s CreateLoginProfileResponseLoginProfile) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileResponseLoginProfile) SetStatus(v string) *CreateLoginProfileResponseLoginProfile {
	s.Status = &v
	return s
}

func (s *CreateLoginProfileResponseLoginProfile) SetUpdateDate(v string) *CreateLoginProfileResponseLoginProfile {
	s.UpdateDate = &v
	return s
}

func (s *CreateLoginProfileResponseLoginProfile) SetPasswordResetRequired(v bool) *CreateLoginProfileResponseLoginProfile {
	s.PasswordResetRequired = &v
	return s
}

func (s *CreateLoginProfileResponseLoginProfile) SetUserPrincipalName(v string) *CreateLoginProfileResponseLoginProfile {
	s.UserPrincipalName = &v
	return s
}

func (s *CreateLoginProfileResponseLoginProfile) SetMFABindRequired(v bool) *CreateLoginProfileResponseLoginProfile {
	s.MFABindRequired = &v
	return s
}

type CreateGroupRequest struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Comments    *string `json:"Comments,omitempty" xml:"Comments,omitempty"`
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetDisplayName(v string) *CreateGroupRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateGroupRequest) SetComments(v string) *CreateGroupRequest {
	s.Comments = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

type CreateGroupResponse struct {
	RequestId *string                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Group     *CreateGroupResponseGroup `json:"Group,omitempty" xml:"Group,omitempty" require:"true" type:"Struct"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetRequestId(v string) *CreateGroupResponse {
	s.RequestId = &v
	return s
}

func (s *CreateGroupResponse) SetGroup(v *CreateGroupResponseGroup) *CreateGroupResponse {
	s.Group = v
	return s
}

type CreateGroupResponseGroup struct {
	GroupName   *string `json:"GroupName,omitempty" xml:"GroupName,omitempty" require:"true"`
	UpdateDate  *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty" require:"true"`
	Comments    *string `json:"Comments,omitempty" xml:"Comments,omitempty" require:"true"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty" require:"true"`
	CreateDate  *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
	GroupId     *string `json:"GroupId,omitempty" xml:"GroupId,omitempty" require:"true"`
}

func (s CreateGroupResponseGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponseGroup) GoString() string {
	return s.String()
}

func (s *CreateGroupResponseGroup) SetGroupName(v string) *CreateGroupResponseGroup {
	s.GroupName = &v
	return s
}

func (s *CreateGroupResponseGroup) SetUpdateDate(v string) *CreateGroupResponseGroup {
	s.UpdateDate = &v
	return s
}

func (s *CreateGroupResponseGroup) SetComments(v string) *CreateGroupResponseGroup {
	s.Comments = &v
	return s
}

func (s *CreateGroupResponseGroup) SetDisplayName(v string) *CreateGroupResponseGroup {
	s.DisplayName = &v
	return s
}

func (s *CreateGroupResponseGroup) SetCreateDate(v string) *CreateGroupResponseGroup {
	s.CreateDate = &v
	return s
}

func (s *CreateGroupResponseGroup) SetGroupId(v string) *CreateGroupResponseGroup {
	s.GroupId = &v
	return s
}

type CreateAccessKeyRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s CreateAccessKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyRequest) SetUserPrincipalName(v string) *CreateAccessKeyRequest {
	s.UserPrincipalName = &v
	return s
}

type CreateAccessKeyResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	AccessKey *CreateAccessKeyResponseAccessKey `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" require:"true" type:"Struct"`
}

func (s CreateAccessKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyResponse) SetRequestId(v string) *CreateAccessKeyResponse {
	s.RequestId = &v
	return s
}

func (s *CreateAccessKeyResponse) SetAccessKey(v *CreateAccessKeyResponseAccessKey) *CreateAccessKeyResponse {
	s.AccessKey = v
	return s
}

type CreateAccessKeyResponseAccessKey struct {
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty" require:"true"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty" require:"true"`
	CreateDate      *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty" require:"true"`
}

func (s CreateAccessKeyResponseAccessKey) String() string {
	return tea.Prettify(s)
}

func (s CreateAccessKeyResponseAccessKey) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyResponseAccessKey) SetStatus(v string) *CreateAccessKeyResponseAccessKey {
	s.Status = &v
	return s
}

func (s *CreateAccessKeyResponseAccessKey) SetAccessKeyId(v string) *CreateAccessKeyResponseAccessKey {
	s.AccessKeyId = &v
	return s
}

func (s *CreateAccessKeyResponseAccessKey) SetAccessKeySecret(v string) *CreateAccessKeyResponseAccessKey {
	s.AccessKeySecret = &v
	return s
}

func (s *CreateAccessKeyResponseAccessKey) SetCreateDate(v string) *CreateAccessKeyResponseAccessKey {
	s.CreateDate = &v
	return s
}

type ChangePasswordRequest struct {
	OldPassword *string `json:"OldPassword,omitempty" xml:"OldPassword,omitempty" require:"true"`
	NewPassword *string `json:"NewPassword,omitempty" xml:"NewPassword,omitempty" require:"true"`
}

func (s ChangePasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordRequest) GoString() string {
	return s.String()
}

func (s *ChangePasswordRequest) SetOldPassword(v string) *ChangePasswordRequest {
	s.OldPassword = &v
	return s
}

func (s *ChangePasswordRequest) SetNewPassword(v string) *ChangePasswordRequest {
	s.NewPassword = &v
	return s
}

type ChangePasswordResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ChangePasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangePasswordResponse) GoString() string {
	return s.String()
}

func (s *ChangePasswordResponse) SetRequestId(v string) *ChangePasswordResponse {
	s.RequestId = &v
	return s
}

type BindMFADeviceRequest struct {
	SerialNumber        *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	UserPrincipalName   *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	AuthenticationCode1 *string `json:"AuthenticationCode1,omitempty" xml:"AuthenticationCode1,omitempty"`
	AuthenticationCode2 *string `json:"AuthenticationCode2,omitempty" xml:"AuthenticationCode2,omitempty"`
}

func (s BindMFADeviceRequest) String() string {
	return tea.Prettify(s)
}

func (s BindMFADeviceRequest) GoString() string {
	return s.String()
}

func (s *BindMFADeviceRequest) SetSerialNumber(v string) *BindMFADeviceRequest {
	s.SerialNumber = &v
	return s
}

func (s *BindMFADeviceRequest) SetUserPrincipalName(v string) *BindMFADeviceRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *BindMFADeviceRequest) SetAuthenticationCode1(v string) *BindMFADeviceRequest {
	s.AuthenticationCode1 = &v
	return s
}

func (s *BindMFADeviceRequest) SetAuthenticationCode2(v string) *BindMFADeviceRequest {
	s.AuthenticationCode2 = &v
	return s
}

type BindMFADeviceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s BindMFADeviceResponse) String() string {
	return tea.Prettify(s)
}

func (s BindMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *BindMFADeviceResponse) SetRequestId(v string) *BindMFADeviceResponse {
	s.RequestId = &v
	return s
}

type AddUserToGroupRequest struct {
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty" require:"true"`
	GroupName         *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s AddUserToGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUserToGroupRequest) GoString() string {
	return s.String()
}

func (s *AddUserToGroupRequest) SetUserPrincipalName(v string) *AddUserToGroupRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *AddUserToGroupRequest) SetGroupName(v string) *AddUserToGroupRequest {
	s.GroupName = &v
	return s
}

type AddUserToGroupResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AddUserToGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUserToGroupResponse) GoString() string {
	return s.String()
}

func (s *AddUserToGroupResponse) SetRequestId(v string) *AddUserToGroupResponse {
	s.RequestId = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ims"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetUserSsoSettingsWithOptions(request *GetUserSsoSettingsRequest, runtime *util.RuntimeOptions) (_result *GetUserSsoSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetUserSsoSettingsResponse{}
	_body, _err := client.DoRequest(tea.String("GetUserSsoSettings"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserSsoSettings(request *GetUserSsoSettingsRequest) (_result *GetUserSsoSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserSsoSettingsResponse{}
	_body, _err := client.GetUserSsoSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetUserSsoSettingsWithOptions(request *SetUserSsoSettingsRequest, runtime *util.RuntimeOptions) (_result *SetUserSsoSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetUserSsoSettingsResponse{}
	_body, _err := client.DoRequest(tea.String("SetUserSsoSettings"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetUserSsoSettings(request *SetUserSsoSettingsRequest) (_result *SetUserSsoSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetUserSsoSettingsResponse{}
	_body, _err := client.SetUserSsoSettingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetDefaultDomainWithOptions(request *SetDefaultDomainRequest, runtime *util.RuntimeOptions) (_result *SetDefaultDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetDefaultDomainResponse{}
	_body, _err := client.DoRequest(tea.String("SetDefaultDomain"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetDefaultDomain(request *SetDefaultDomainRequest) (_result *SetDefaultDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetDefaultDomainResponse{}
	_body, _err := client.SetDefaultDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserBasicInfosWithOptions(request *ListUserBasicInfosRequest, runtime *util.RuntimeOptions) (_result *ListUserBasicInfosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListUserBasicInfosResponse{}
	_body, _err := client.DoRequest(tea.String("ListUserBasicInfos"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserBasicInfos(request *ListUserBasicInfosRequest) (_result *ListUserBasicInfosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserBasicInfosResponse{}
	_body, _err := client.ListUserBasicInfosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateApplicationWithOptions(request *UpdateApplicationRequest, runtime *util.RuntimeOptions) (_result *UpdateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateApplicationResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateApplication"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateApplication(request *UpdateApplicationRequest) (_result *UpdateApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateApplicationResponse{}
	_body, _err := client.UpdateApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListApplicationsWithOptions(request *ListApplicationsRequest, runtime *util.RuntimeOptions) (_result *ListApplicationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListApplicationsResponse{}
	_body, _err := client.DoRequest(tea.String("ListApplications"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApplications(request *ListApplicationsRequest) (_result *ListApplicationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListApplicationsResponse{}
	_body, _err := client.ListApplicationsWithOptions(request, runtime)
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
	_result = &GetApplicationResponse{}
	_body, _err := client.DoRequest(tea.String("GetApplication"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteApplicationWithOptions(request *DeleteApplicationRequest, runtime *util.RuntimeOptions) (_result *DeleteApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteApplicationResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteApplication"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) GetAppSecretWithOptions(request *GetAppSecretRequest, runtime *util.RuntimeOptions) (_result *GetAppSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAppSecretResponse{}
	_body, _err := client.DoRequest(tea.String("GetAppSecret"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAppSecret(request *GetAppSecretRequest) (_result *GetAppSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAppSecretResponse{}
	_body, _err := client.GetAppSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppSecretWithOptions(request *CreateAppSecretRequest, runtime *util.RuntimeOptions) (_result *CreateAppSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateAppSecretResponse{}
	_body, _err := client.DoRequest(tea.String("CreateAppSecret"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppSecret(request *CreateAppSecretRequest) (_result *CreateAppSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppSecretResponse{}
	_body, _err := client.CreateAppSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPredefinedScopesWithOptions(request *ListPredefinedScopesRequest, runtime *util.RuntimeOptions) (_result *ListPredefinedScopesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListPredefinedScopesResponse{}
	_body, _err := client.DoRequest(tea.String("ListPredefinedScopes"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPredefinedScopes(request *ListPredefinedScopesRequest) (_result *ListPredefinedScopesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPredefinedScopesResponse{}
	_body, _err := client.ListPredefinedScopesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateApplicationWithOptions(request *CreateApplicationRequest, runtime *util.RuntimeOptions) (_result *CreateApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateApplicationResponse{}
	_body, _err := client.DoRequest(tea.String("CreateApplication"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
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

func (client *Client) DeleteAppSecretWithOptions(request *DeleteAppSecretRequest, runtime *util.RuntimeOptions) (_result *DeleteAppSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteAppSecretResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteAppSecret"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAppSecret(request *DeleteAppSecretRequest) (_result *DeleteAppSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppSecretResponse{}
	_body, _err := client.DeleteAppSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppSecretIdsWithOptions(request *ListAppSecretIdsRequest, runtime *util.RuntimeOptions) (_result *ListAppSecretIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListAppSecretIdsResponse{}
	_body, _err := client.DoRequest(tea.String("ListAppSecretIds"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAppSecretIds(request *ListAppSecretIdsRequest) (_result *ListAppSecretIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAppSecretIdsResponse{}
	_body, _err := client.ListAppSecretIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateCredentialReportWithOptions(request *GenerateCredentialReportRequest, runtime *util.RuntimeOptions) (_result *GenerateCredentialReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GenerateCredentialReportResponse{}
	_body, _err := client.DoRequest(tea.String("GenerateCredentialReport"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateCredentialReport(request *GenerateCredentialReportRequest) (_result *GenerateCredentialReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateCredentialReportResponse{}
	_body, _err := client.GenerateCredentialReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCredentialReportWithOptions(request *GetCredentialReportRequest, runtime *util.RuntimeOptions) (_result *GetCredentialReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetCredentialReportResponse{}
	_body, _err := client.DoRequest(tea.String("GetCredentialReport"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCredentialReport(request *GetCredentialReportRequest) (_result *GetCredentialReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCredentialReportResponse{}
	_body, _err := client.GetCredentialReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateUser"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSAMLProviderWithOptions(request *UpdateSAMLProviderRequest, runtime *util.RuntimeOptions) (_result *UpdateSAMLProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateSAMLProviderResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateSAMLProvider"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSAMLProvider(request *UpdateSAMLProviderRequest) (_result *UpdateSAMLProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSAMLProviderResponse{}
	_body, _err := client.UpdateSAMLProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateLoginProfileWithOptions(request *UpdateLoginProfileRequest, runtime *util.RuntimeOptions) (_result *UpdateLoginProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateLoginProfileResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateLoginProfile"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateLoginProfile(request *UpdateLoginProfileRequest) (_result *UpdateLoginProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLoginProfileResponse{}
	_body, _err := client.UpdateLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateGroupWithOptions(request *UpdateGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateGroupResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateGroup(request *UpdateGroupRequest) (_result *UpdateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateGroupResponse{}
	_body, _err := client.UpdateGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAccessKeyWithOptions(request *UpdateAccessKeyRequest, runtime *util.RuntimeOptions) (_result *UpdateAccessKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateAccessKeyResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateAccessKey"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAccessKey(request *UpdateAccessKeyRequest) (_result *UpdateAccessKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAccessKeyResponse{}
	_body, _err := client.UpdateAccessKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindMFADeviceWithOptions(request *UnbindMFADeviceRequest, runtime *util.RuntimeOptions) (_result *UnbindMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UnbindMFADeviceResponse{}
	_body, _err := client.DoRequest(tea.String("UnbindMFADevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindMFADevice(request *UnbindMFADeviceRequest) (_result *UnbindMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindMFADeviceResponse{}
	_body, _err := client.UnbindMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetSecurityPreferenceWithOptions(request *SetSecurityPreferenceRequest, runtime *util.RuntimeOptions) (_result *SetSecurityPreferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetSecurityPreferenceResponse{}
	_body, _err := client.DoRequest(tea.String("SetSecurityPreference"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetSecurityPreference(request *SetSecurityPreferenceRequest) (_result *SetSecurityPreferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetSecurityPreferenceResponse{}
	_body, _err := client.SetSecurityPreferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetPasswordPolicyWithOptions(request *SetPasswordPolicyRequest, runtime *util.RuntimeOptions) (_result *SetPasswordPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SetPasswordPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("SetPasswordPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetPasswordPolicy(request *SetPasswordPolicyRequest) (_result *SetPasswordPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetPasswordPolicyResponse{}
	_body, _err := client.SetPasswordPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveUserFromGroupWithOptions(request *RemoveUserFromGroupRequest, runtime *util.RuntimeOptions) (_result *RemoveUserFromGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RemoveUserFromGroupResponse{}
	_body, _err := client.DoRequest(tea.String("RemoveUserFromGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (_result *RemoveUserFromGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUserFromGroupResponse{}
	_body, _err := client.RemoveUserFromGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVirtualMFADevicesWithOptions(request *ListVirtualMFADevicesRequest, runtime *util.RuntimeOptions) (_result *ListVirtualMFADevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListVirtualMFADevicesResponse{}
	_body, _err := client.DoRequest(tea.String("ListVirtualMFADevices"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVirtualMFADevices(request *ListVirtualMFADevicesRequest) (_result *ListVirtualMFADevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVirtualMFADevicesResponse{}
	_body, _err := client.ListVirtualMFADevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersForGroupWithOptions(request *ListUsersForGroupRequest, runtime *util.RuntimeOptions) (_result *ListUsersForGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListUsersForGroupResponse{}
	_body, _err := client.DoRequest(tea.String("ListUsersForGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsersForGroup(request *ListUsersForGroupRequest) (_result *ListUsersForGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersForGroupResponse{}
	_body, _err := client.ListUsersForGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListUsersResponse{}
	_body, _err := client.DoRequest(tea.String("ListUsers"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) ListSAMLProvidersWithOptions(request *ListSAMLProvidersRequest, runtime *util.RuntimeOptions) (_result *ListSAMLProvidersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListSAMLProvidersResponse{}
	_body, _err := client.DoRequest(tea.String("ListSAMLProviders"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSAMLProviders(request *ListSAMLProvidersRequest) (_result *ListSAMLProvidersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSAMLProvidersResponse{}
	_body, _err := client.ListSAMLProvidersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGroupsForUserWithOptions(request *ListGroupsForUserRequest, runtime *util.RuntimeOptions) (_result *ListGroupsForUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListGroupsForUserResponse{}
	_body, _err := client.DoRequest(tea.String("ListGroupsForUser"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroupsForUser(request *ListGroupsForUserRequest) (_result *ListGroupsForUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupsForUserResponse{}
	_body, _err := client.ListGroupsForUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListGroupsWithOptions(request *ListGroupsRequest, runtime *util.RuntimeOptions) (_result *ListGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListGroupsResponse{}
	_body, _err := client.DoRequest(tea.String("ListGroups"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListGroups(request *ListGroupsRequest) (_result *ListGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListGroupsResponse{}
	_body, _err := client.ListGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAccessKeysWithOptions(request *ListAccessKeysRequest, runtime *util.RuntimeOptions) (_result *ListAccessKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListAccessKeysResponse{}
	_body, _err := client.DoRequest(tea.String("ListAccessKeys"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAccessKeys(request *ListAccessKeysRequest) (_result *ListAccessKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccessKeysResponse{}
	_body, _err := client.ListAccessKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserMFAInfoWithOptions(request *GetUserMFAInfoRequest, runtime *util.RuntimeOptions) (_result *GetUserMFAInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetUserMFAInfoResponse{}
	_body, _err := client.DoRequest(tea.String("GetUserMFAInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserMFAInfo(request *GetUserMFAInfoRequest) (_result *GetUserMFAInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserMFAInfoResponse{}
	_body, _err := client.GetUserMFAInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserWithOptions(request *GetUserRequest, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetUserResponse{}
	_body, _err := client.DoRequest(tea.String("GetUser"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSecurityPreferenceWithOptions(request *GetSecurityPreferenceRequest, runtime *util.RuntimeOptions) (_result *GetSecurityPreferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSecurityPreferenceResponse{}
	_body, _err := client.DoRequest(tea.String("GetSecurityPreference"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSecurityPreference(request *GetSecurityPreferenceRequest) (_result *GetSecurityPreferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecurityPreferenceResponse{}
	_body, _err := client.GetSecurityPreferenceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSAMLProviderWithOptions(request *GetSAMLProviderRequest, runtime *util.RuntimeOptions) (_result *GetSAMLProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSAMLProviderResponse{}
	_body, _err := client.DoRequest(tea.String("GetSAMLProvider"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSAMLProvider(request *GetSAMLProviderRequest) (_result *GetSAMLProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSAMLProviderResponse{}
	_body, _err := client.GetSAMLProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPasswordPolicyWithOptions(request *GetPasswordPolicyRequest, runtime *util.RuntimeOptions) (_result *GetPasswordPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPasswordPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("GetPasswordPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPasswordPolicy(request *GetPasswordPolicyRequest) (_result *GetPasswordPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPasswordPolicyResponse{}
	_body, _err := client.GetPasswordPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLoginProfileWithOptions(request *GetLoginProfileRequest, runtime *util.RuntimeOptions) (_result *GetLoginProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetLoginProfileResponse{}
	_body, _err := client.DoRequest(tea.String("GetLoginProfile"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLoginProfile(request *GetLoginProfileRequest) (_result *GetLoginProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLoginProfileResponse{}
	_body, _err := client.GetLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGroupWithOptions(request *GetGroupRequest, runtime *util.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetGroupResponse{}
	_body, _err := client.DoRequest(tea.String("GetGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGroup(request *GetGroupRequest) (_result *GetGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGroupResponse{}
	_body, _err := client.GetGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDefaultDomainWithOptions(request *GetDefaultDomainRequest, runtime *util.RuntimeOptions) (_result *GetDefaultDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetDefaultDomainResponse{}
	_body, _err := client.DoRequest(tea.String("GetDefaultDomain"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDefaultDomain(request *GetDefaultDomainRequest) (_result *GetDefaultDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDefaultDomainResponse{}
	_body, _err := client.GetDefaultDomainWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccountSummaryWithOptions(request *GetAccountSummaryRequest, runtime *util.RuntimeOptions) (_result *GetAccountSummaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAccountSummaryResponse{}
	_body, _err := client.DoRequest(tea.String("GetAccountSummary"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccountSummary(request *GetAccountSummaryRequest) (_result *GetAccountSummaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountSummaryResponse{}
	_body, _err := client.GetAccountSummaryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccountSecurityPracticeReportWithOptions(request *GetAccountSecurityPracticeReportRequest, runtime *util.RuntimeOptions) (_result *GetAccountSecurityPracticeReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAccountSecurityPracticeReportResponse{}
	_body, _err := client.DoRequest(tea.String("GetAccountSecurityPracticeReport"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccountSecurityPracticeReport(request *GetAccountSecurityPracticeReportRequest) (_result *GetAccountSecurityPracticeReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountSecurityPracticeReportResponse{}
	_body, _err := client.GetAccountSecurityPracticeReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccountMFAInfoWithOptions(request *GetAccountMFAInfoRequest, runtime *util.RuntimeOptions) (_result *GetAccountMFAInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAccountMFAInfoResponse{}
	_body, _err := client.DoRequest(tea.String("GetAccountMFAInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccountMFAInfo(request *GetAccountMFAInfoRequest) (_result *GetAccountMFAInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountMFAInfoResponse{}
	_body, _err := client.GetAccountMFAInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccessKeyLastUsedWithOptions(request *GetAccessKeyLastUsedRequest, runtime *util.RuntimeOptions) (_result *GetAccessKeyLastUsedResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetAccessKeyLastUsedResponse{}
	_body, _err := client.DoRequest(tea.String("GetAccessKeyLastUsed"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccessKeyLastUsed(request *GetAccessKeyLastUsedRequest) (_result *GetAccessKeyLastUsedResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccessKeyLastUsedResponse{}
	_body, _err := client.GetAccessKeyLastUsedWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableVirtualMFAWithOptions(request *DisableVirtualMFARequest, runtime *util.RuntimeOptions) (_result *DisableVirtualMFAResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DisableVirtualMFAResponse{}
	_body, _err := client.DoRequest(tea.String("DisableVirtualMFA"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableVirtualMFA(request *DisableVirtualMFARequest) (_result *DisableVirtualMFAResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableVirtualMFAResponse{}
	_body, _err := client.DisableVirtualMFAWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVirtualMFADeviceWithOptions(request *DeleteVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteVirtualMFADevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVirtualMFADevice(request *DeleteVirtualMFADeviceRequest) (_result *DeleteVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVirtualMFADeviceResponse{}
	_body, _err := client.DeleteVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteUser"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSAMLProviderWithOptions(request *DeleteSAMLProviderRequest, runtime *util.RuntimeOptions) (_result *DeleteSAMLProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteSAMLProviderResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteSAMLProvider"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSAMLProvider(request *DeleteSAMLProviderRequest) (_result *DeleteSAMLProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSAMLProviderResponse{}
	_body, _err := client.DeleteSAMLProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLoginProfileWithOptions(request *DeleteLoginProfileRequest, runtime *util.RuntimeOptions) (_result *DeleteLoginProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteLoginProfileResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteLoginProfile"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLoginProfile(request *DeleteLoginProfileRequest) (_result *DeleteLoginProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteLoginProfileResponse{}
	_body, _err := client.DeleteLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGroup(request *DeleteGroupRequest) (_result *DeleteGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAccessKeyWithOptions(request *DeleteAccessKeyRequest, runtime *util.RuntimeOptions) (_result *DeleteAccessKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteAccessKeyResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteAccessKey"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAccessKey(request *DeleteAccessKeyRequest) (_result *DeleteAccessKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAccessKeyResponse{}
	_body, _err := client.DeleteAccessKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateVirtualMFADeviceWithOptions(request *CreateVirtualMFADeviceRequest, runtime *util.RuntimeOptions) (_result *CreateVirtualMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateVirtualMFADeviceResponse{}
	_body, _err := client.DoRequest(tea.String("CreateVirtualMFADevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateVirtualMFADevice(request *CreateVirtualMFADeviceRequest) (_result *CreateVirtualMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVirtualMFADeviceResponse{}
	_body, _err := client.CreateVirtualMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserWithOptions(request *CreateUserRequest, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateUserResponse{}
	_body, _err := client.DoRequest(tea.String("CreateUser"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSAMLProviderWithOptions(request *CreateSAMLProviderRequest, runtime *util.RuntimeOptions) (_result *CreateSAMLProviderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateSAMLProviderResponse{}
	_body, _err := client.DoRequest(tea.String("CreateSAMLProvider"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSAMLProvider(request *CreateSAMLProviderRequest) (_result *CreateSAMLProviderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSAMLProviderResponse{}
	_body, _err := client.CreateSAMLProviderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLoginProfileWithOptions(request *CreateLoginProfileRequest, runtime *util.RuntimeOptions) (_result *CreateLoginProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateLoginProfileResponse{}
	_body, _err := client.DoRequest(tea.String("CreateLoginProfile"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLoginProfile(request *CreateLoginProfileRequest) (_result *CreateLoginProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLoginProfileResponse{}
	_body, _err := client.CreateLoginProfileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.DoRequest(tea.String("CreateGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) CreateAccessKeyWithOptions(request *CreateAccessKeyRequest, runtime *util.RuntimeOptions) (_result *CreateAccessKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateAccessKeyResponse{}
	_body, _err := client.DoRequest(tea.String("CreateAccessKey"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccessKey(request *CreateAccessKeyRequest) (_result *CreateAccessKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccessKeyResponse{}
	_body, _err := client.CreateAccessKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangePasswordWithOptions(request *ChangePasswordRequest, runtime *util.RuntimeOptions) (_result *ChangePasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ChangePasswordResponse{}
	_body, _err := client.DoRequest(tea.String("ChangePassword"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangePassword(request *ChangePasswordRequest) (_result *ChangePasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangePasswordResponse{}
	_body, _err := client.ChangePasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindMFADeviceWithOptions(request *BindMFADeviceRequest, runtime *util.RuntimeOptions) (_result *BindMFADeviceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BindMFADeviceResponse{}
	_body, _err := client.DoRequest(tea.String("BindMFADevice"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindMFADevice(request *BindMFADeviceRequest) (_result *BindMFADeviceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindMFADeviceResponse{}
	_body, _err := client.BindMFADeviceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUserToGroupWithOptions(request *AddUserToGroupRequest, runtime *util.RuntimeOptions) (_result *AddUserToGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddUserToGroupResponse{}
	_body, _err := client.DoRequest(tea.String("AddUserToGroup"), tea.String("HTTPS"), tea.String("POST"), tea.String("2019-08-15"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUserToGroup(request *AddUserToGroupRequest) (_result *AddUserToGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUserToGroupResponse{}
	_body, _err := client.AddUserToGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
