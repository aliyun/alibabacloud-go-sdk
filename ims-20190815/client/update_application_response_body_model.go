// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplication(v *UpdateApplicationResponseBodyApplication) *UpdateApplicationResponseBody
	GetApplication() *UpdateApplicationResponseBodyApplication
	SetRequestId(v string) *UpdateApplicationResponseBody
	GetRequestId() *string
}

type UpdateApplicationResponseBody struct {
	// The information about the application.
	Application *UpdateApplicationResponseBodyApplication `json:"Application,omitempty" xml:"Application,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6616F09B-2768-4C11-8866-A8EE4C4A583E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBody) GetApplication() *UpdateApplicationResponseBodyApplication {
	return s.Application
}

func (s *UpdateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApplicationResponseBody) SetApplication(v *UpdateApplicationResponseBodyApplication) *UpdateApplicationResponseBody {
	s.Application = v
	return s
}

func (s *UpdateApplicationResponseBody) SetRequestId(v string) *UpdateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApplicationResponseBody) Validate() error {
	if s.Application != nil {
		if err := s.Application.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationResponseBodyApplication struct {
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
	// 472457090344041****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// myapp
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type.
	//
	// example:
	//
	// WebApp
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-10-23T08:06:57Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted on the application.
	DelegatedScope *UpdateApplicationResponseBodyApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// NewApp
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Indicates whether the application can be installed by using other Alibaba Cloud accounts.
	//
	// example:
	//
	// true
	IsMultiTenant   *bool   `json:"IsMultiTenant,omitempty" xml:"IsMultiTenant,omitempty"`
	ProtocolVersion *string `json:"ProtocolVersion,omitempty" xml:"ProtocolVersion,omitempty"`
	// The callback URLs.
	RedirectUris *UpdateApplicationResponseBodyApplicationRedirectUris `json:"RedirectUris,omitempty" xml:"RedirectUris,omitempty" type:"Struct"`
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
	// 2020-10-23T08:06:57Z
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s UpdateApplicationResponseBodyApplication) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationResponseBodyApplication) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBodyApplication) GetAccessTokenValidity() *int32 {
	return s.AccessTokenValidity
}

func (s *UpdateApplicationResponseBodyApplication) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateApplicationResponseBodyApplication) GetAppId() *string {
	return s.AppId
}

func (s *UpdateApplicationResponseBodyApplication) GetAppName() *string {
	return s.AppName
}

func (s *UpdateApplicationResponseBodyApplication) GetAppType() *string {
	return s.AppType
}

func (s *UpdateApplicationResponseBodyApplication) GetCreateDate() *string {
	return s.CreateDate
}

func (s *UpdateApplicationResponseBodyApplication) GetDelegatedScope() *UpdateApplicationResponseBodyApplicationDelegatedScope {
	return s.DelegatedScope
}

func (s *UpdateApplicationResponseBodyApplication) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateApplicationResponseBodyApplication) GetIsMultiTenant() *bool {
	return s.IsMultiTenant
}

func (s *UpdateApplicationResponseBodyApplication) GetProtocolVersion() *string {
	return s.ProtocolVersion
}

func (s *UpdateApplicationResponseBodyApplication) GetRedirectUris() *UpdateApplicationResponseBodyApplicationRedirectUris {
	return s.RedirectUris
}

func (s *UpdateApplicationResponseBodyApplication) GetRefreshTokenValidity() *int32 {
	return s.RefreshTokenValidity
}

func (s *UpdateApplicationResponseBodyApplication) GetSecretRequired() *bool {
	return s.SecretRequired
}

func (s *UpdateApplicationResponseBodyApplication) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *UpdateApplicationResponseBodyApplication) SetAccessTokenValidity(v int32) *UpdateApplicationResponseBodyApplication {
	s.AccessTokenValidity = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetAccountId(v string) *UpdateApplicationResponseBodyApplication {
	s.AccountId = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetAppId(v string) *UpdateApplicationResponseBodyApplication {
	s.AppId = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetAppName(v string) *UpdateApplicationResponseBodyApplication {
	s.AppName = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetAppType(v string) *UpdateApplicationResponseBodyApplication {
	s.AppType = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetCreateDate(v string) *UpdateApplicationResponseBodyApplication {
	s.CreateDate = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetDelegatedScope(v *UpdateApplicationResponseBodyApplicationDelegatedScope) *UpdateApplicationResponseBodyApplication {
	s.DelegatedScope = v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetDisplayName(v string) *UpdateApplicationResponseBodyApplication {
	s.DisplayName = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetIsMultiTenant(v bool) *UpdateApplicationResponseBodyApplication {
	s.IsMultiTenant = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetProtocolVersion(v string) *UpdateApplicationResponseBodyApplication {
	s.ProtocolVersion = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetRedirectUris(v *UpdateApplicationResponseBodyApplicationRedirectUris) *UpdateApplicationResponseBodyApplication {
	s.RedirectUris = v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetRefreshTokenValidity(v int32) *UpdateApplicationResponseBodyApplication {
	s.RefreshTokenValidity = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetSecretRequired(v bool) *UpdateApplicationResponseBodyApplication {
	s.SecretRequired = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) SetUpdateDate(v string) *UpdateApplicationResponseBodyApplication {
	s.UpdateDate = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplication) Validate() error {
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

type UpdateApplicationResponseBodyApplicationDelegatedScope struct {
	// The information about the permissions that are granted on the application.
	PredefinedScopes *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s UpdateApplicationResponseBodyApplicationDelegatedScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationResponseBodyApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScope) GetPredefinedScopes() *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes {
	return s.PredefinedScopes
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScope) SetPredefinedScopes(v *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) *UpdateApplicationResponseBodyApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScope) Validate() error {
	if s.PredefinedScopes != nil {
		if err := s.PredefinedScopes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) GetPredefinedScope() []*UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	return s.PredefinedScope
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopes) Validate() error {
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

type UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
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

func (s UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GetDescription() *string {
	return s.Description
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GetName() *string {
	return s.Name
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) GetRequired() *bool {
	return s.Required
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) SetRequired(v bool) *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Required = &v
	return s
}

func (s *UpdateApplicationResponseBodyApplicationDelegatedScopePredefinedScopesPredefinedScope) Validate() error {
	return dara.Validate(s)
}

type UpdateApplicationResponseBodyApplicationRedirectUris struct {
	RedirectUri []*string `json:"RedirectUri,omitempty" xml:"RedirectUri,omitempty" type:"Repeated"`
}

func (s UpdateApplicationResponseBodyApplicationRedirectUris) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationResponseBodyApplicationRedirectUris) GoString() string {
	return s.String()
}

func (s *UpdateApplicationResponseBodyApplicationRedirectUris) GetRedirectUri() []*string {
	return s.RedirectUri
}

func (s *UpdateApplicationResponseBodyApplicationRedirectUris) SetRedirectUri(v []*string) *UpdateApplicationResponseBodyApplicationRedirectUris {
	s.RedirectUri = v
	return s
}

func (s *UpdateApplicationResponseBodyApplicationRedirectUris) Validate() error {
	return dara.Validate(s)
}
