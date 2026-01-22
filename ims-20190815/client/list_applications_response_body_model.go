// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v *ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody
	GetApplications() *ListApplicationsResponseBodyApplications
	SetRequestId(v string) *ListApplicationsResponseBody
	GetRequestId() *string
}

type ListApplicationsResponseBody struct {
	// The information about the application.
	Applications *ListApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CE458B58-8C40-46F7-A9D4-CB85136B0C06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) GetApplications() *ListApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *ListApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsResponseBody) SetApplications(v *ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) Validate() error {
	if s.Applications != nil {
		if err := s.Applications.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationsResponseBodyApplications struct {
	Application []*ListApplicationsResponseBodyApplicationsApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplications) GetApplication() []*ListApplicationsResponseBodyApplicationsApplication {
	return s.Application
}

func (s *ListApplicationsResponseBodyApplications) SetApplication(v []*ListApplicationsResponseBodyApplicationsApplication) *ListApplicationsResponseBodyApplications {
	s.Application = v
	return s
}

func (s *ListApplicationsResponseBodyApplications) Validate() error {
	if s.Application != nil {
		for _, item := range s.Application {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsResponseBodyApplicationsApplication struct {
	// The validity period of the access token. Unit: seconds.
	//
	// example:
	//
	// 3600
	AccessTokenValidity *int32 `json:"AccessTokenValidity,omitempty" xml:"AccessTokenValidity,omitempty"`
	// The ID of the Alibaba Cloud account to which the application belongs.
	//
	// example:
	//
	// 177242285274****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 441442900344560****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// myapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type. Valid values:
	//
	// 	- WebApp: a web application.
	//
	// 	- NativeApp: a native application that runs on an operating system, such as a desktop or mobile operating system.
	//
	// 	- ServerApp: an application that can access Alibaba Cloud services without the need for user logon. Only applications that synchronize user information based on the System for Cross-domain Identity Management (SCIM) protocol are supported.
	//
	// example:
	//
	// WebApp
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-23T09:33:22Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted on the application.
	DelegatedScope *ListApplicationsResponseBodyApplicationsApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// myapp
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Indicates whether the application can be installed by using other Alibaba Cloud accounts.
	//
	// example:
	//
	// true
	IsMultiTenant   *bool   `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty"`
	ProtocolVersion *string `json:"ProtocolVersion,omitempty" xml:"ProtocolVersion,omitempty"`
	// The callback URLs.
	RedirectUris *ListApplicationsResponseBodyApplicationsApplicationRedirectUris `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Struct"`
	// The validity period of the refresh token. Unit: seconds.
	//
	// example:
	//
	// 7776000
	RefreshTokenValidity *int32 `json:"RefreshTokenValidity,omitempty" xml:"RefreshTokenValidity,omitempty"`
	// Indicates whether a secret is required.
	//
	// example:
	//
	// true
	SecretRequired *bool `json:"SecretRequired,omitempty" xml:"SecretRequired,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2020-10-23T09:33:22Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListApplicationsResponseBodyApplicationsApplication) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplication) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetAccessTokenValidity() *int32 {
	return s.AccessTokenValidity
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetAccountId() *string {
	return s.AccountId
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetAppName() *string {
	return s.AppName
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetAppType() *string {
	return s.AppType
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetDelegatedScope() *ListApplicationsResponseBodyApplicationsApplicationDelegatedScope {
	return s.DelegatedScope
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetIsMultiTenant() *bool {
	return s.IsMultiTenant
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetProtocolVersion() *string {
	return s.ProtocolVersion
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetRedirectUris() *ListApplicationsResponseBodyApplicationsApplicationRedirectUris {
	return s.RedirectUris
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetRefreshTokenValidity() *int32 {
	return s.RefreshTokenValidity
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetSecretRequired() *bool {
	return s.SecretRequired
}

func (s *ListApplicationsResponseBodyApplicationsApplication) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetAccessTokenValidity(v int32) *ListApplicationsResponseBodyApplicationsApplication {
	s.AccessTokenValidity = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetAccountId(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.AccountId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetAppId(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.AppId = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetAppName(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.AppName = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetAppType(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.AppType = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetCreateDate(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.CreateDate = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetDelegatedScope(v *ListApplicationsResponseBodyApplicationsApplicationDelegatedScope) *ListApplicationsResponseBodyApplicationsApplication {
	s.DelegatedScope = v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetDisplayName(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.DisplayName = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetIsMultiTenant(v bool) *ListApplicationsResponseBodyApplicationsApplication {
	s.IsMultiTenant = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetProtocolVersion(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.ProtocolVersion = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetRedirectUris(v *ListApplicationsResponseBodyApplicationsApplicationRedirectUris) *ListApplicationsResponseBodyApplicationsApplication {
	s.RedirectUris = v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetRefreshTokenValidity(v int32) *ListApplicationsResponseBodyApplicationsApplication {
	s.RefreshTokenValidity = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetSecretRequired(v bool) *ListApplicationsResponseBodyApplicationsApplication {
	s.SecretRequired = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) SetUpdateDate(v string) *ListApplicationsResponseBodyApplicationsApplication {
	s.UpdateDate = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplication) Validate() error {
	if s.DelegatedScope != nil {
		if err := s.DelegatedScope.Validate(); err != nil {
			return err
		}
	}
	if s.RedirectUris != nil {
		if err := s.RedirectUris.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationsResponseBodyApplicationsApplicationDelegatedScope struct {
	// The information about the permissions that are granted on the application.
	PredefinedScopes *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationDelegatedScope) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScope) GetPredefinedScopes() *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes {
	return s.PredefinedScopes
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScope) SetPredefinedScopes(v *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes) *ListApplicationsResponseBodyApplicationsApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScope) Validate() error {
	if s.PredefinedScopes != nil {
		if err := s.PredefinedScopes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes) GetPredefinedScope() []*ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope {
	return s.PredefinedScope
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopes) Validate() error {
	if s.PredefinedScope != nil {
		for _, item := range s.PredefinedScope {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission.
	//
	// example:
	//
	// Obtain the OpenID of the user. This is the default permission that you cannot remove.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the permission is automatically selected by default when you install the application. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// `openid` is required by default.
	//
	// example:
	//
	// true
	Required *bool `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) GetName() *string {
	return s.Name
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) GetRequired() *bool {
	return s.Required
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) SetRequired(v bool) *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Required = &v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationDelegatedScopePredefinedScopesPredefinedScope) Validate() error {
	return dara.Validate(s)
}

type ListApplicationsResponseBodyApplicationsApplicationRedirectUris struct {
	RedirectUri []*string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty" type:"Repeated"`
}

func (s ListApplicationsResponseBodyApplicationsApplicationRedirectUris) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplicationsApplicationRedirectUris) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplicationsApplicationRedirectUris) GetRedirectUri() []*string {
	return s.RedirectUri
}

func (s *ListApplicationsResponseBodyApplicationsApplicationRedirectUris) SetRedirectUri(v []*string) *ListApplicationsResponseBodyApplicationsApplicationRedirectUris {
	s.RedirectUri = v
	return s
}

func (s *ListApplicationsResponseBodyApplicationsApplicationRedirectUris) Validate() error {
	return dara.Validate(s)
}
