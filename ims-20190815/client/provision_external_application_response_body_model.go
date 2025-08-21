// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProvisionExternalApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExternalApplication(v *ProvisionExternalApplicationResponseBodyExternalApplication) *ProvisionExternalApplicationResponseBody
	GetExternalApplication() *ProvisionExternalApplicationResponseBodyExternalApplication
	SetRequestId(v string) *ProvisionExternalApplicationResponseBody
	GetRequestId() *string
}

type ProvisionExternalApplicationResponseBody struct {
	// The information about the external application.
	ExternalApplication *ProvisionExternalApplicationResponseBodyExternalApplication `json:"ExternalApplication,omitempty" xml:"ExternalApplication,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 55535873-9A6B-5C87-853F-C7CD258826F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ProvisionExternalApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProvisionExternalApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationResponseBody) GetExternalApplication() *ProvisionExternalApplicationResponseBodyExternalApplication {
	return s.ExternalApplication
}

func (s *ProvisionExternalApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProvisionExternalApplicationResponseBody) SetExternalApplication(v *ProvisionExternalApplicationResponseBodyExternalApplication) *ProvisionExternalApplicationResponseBody {
	s.ExternalApplication = v
	return s
}

func (s *ProvisionExternalApplicationResponseBody) SetRequestId(v string) *ProvisionExternalApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type ProvisionExternalApplicationResponseBodyExternalApplication struct {
	// The name of the application principal. The value is in the `<app_name>@app.<account_id>.onaliyun.com` format.
	//
	// example:
	//
	// GiteePrd@app.177242285274****.onaliyun.com
	AppPrincipalName *string `json:"AppPrincipalName,omitempty" xml:"AppPrincipalName,omitempty"`
	// The time when the application was installed. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the scopes of permissions that are granted to the application.
	DelegatedScope *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// GiteeAliyun
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 403550611646604****
	ForeignAppId *string `json:"ForeignAppId,omitempty" xml:"ForeignAppId,omitempty"`
	// The ID of the Alibaba Cloud account to which the external application belongs.
	//
	// example:
	//
	// 157242285274****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The update time. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ProvisionExternalApplicationResponseBodyExternalApplication) String() string {
	return dara.Prettify(s)
}

func (s ProvisionExternalApplicationResponseBodyExternalApplication) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) GetAppPrincipalName() *string {
	return s.AppPrincipalName
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) GetDelegatedScope() *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope {
	return s.DelegatedScope
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) GetForeignAppId() *string {
	return s.ForeignAppId
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) GetTenantId() *string {
	return s.TenantId
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetAppPrincipalName(v string) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.AppPrincipalName = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetCreateDate(v string) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.CreateDate = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetDelegatedScope(v *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.DelegatedScope = v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetDisplayName(v string) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.DisplayName = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetForeignAppId(v string) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.ForeignAppId = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetTenantId(v string) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.TenantId = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) SetUpdateDate(v string) *ProvisionExternalApplicationResponseBodyExternalApplication {
	s.UpdateDate = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplication) Validate() error {
	return dara.Validate(s)
}

type ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope struct {
	// The information about the scopes of permissions that are granted to the application.
	PredefinedScopes *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope) String() string {
	return dara.Prettify(s)
}

func (s ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope) GetPredefinedScopes() *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes {
	return s.PredefinedScopes
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope) SetPredefinedScopes(v *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope) Validate() error {
	return dara.Validate(s)
}

type ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) String() string {
	return dara.Prettify(s)
}

func (s ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) GetPredefinedScope() []*ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	return s.PredefinedScope
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) Validate() error {
	return dara.Validate(s)
}

type ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission scope.
	//
	// example:
	//
	// Obtains the OpenID of the user. This is the default scope and cannot be deleted.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission scope.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return dara.Prettify(s)
}

func (s ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) GetDescription() *string {
	return s.Description
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) GetName() *string {
	return s.Name
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) Validate() error {
	return dara.Validate(s)
}
