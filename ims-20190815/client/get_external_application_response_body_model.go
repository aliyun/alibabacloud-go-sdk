// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExternalApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExternalApplication(v *GetExternalApplicationResponseBodyExternalApplication) *GetExternalApplicationResponseBody
	GetExternalApplication() *GetExternalApplicationResponseBodyExternalApplication
	SetRequestId(v string) *GetExternalApplicationResponseBody
	GetRequestId() *string
}

type GetExternalApplicationResponseBody struct {
	// The information about the external application.
	ExternalApplication *GetExternalApplicationResponseBodyExternalApplication `json:"ExternalApplication,omitempty" xml:"ExternalApplication,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E4C4D1BD-2558-5BD1-8C36-A5D7FB174A55
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExternalApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExternalApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationResponseBody) GetExternalApplication() *GetExternalApplicationResponseBodyExternalApplication {
	return s.ExternalApplication
}

func (s *GetExternalApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExternalApplicationResponseBody) SetExternalApplication(v *GetExternalApplicationResponseBodyExternalApplication) *GetExternalApplicationResponseBody {
	s.ExternalApplication = v
	return s
}

func (s *GetExternalApplicationResponseBody) SetRequestId(v string) *GetExternalApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExternalApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetExternalApplicationResponseBodyExternalApplication struct {
	// The name of the application principal. The value is in the `<app_name>@app.<account_id>.onaliyun.com` format.
	//
	// example:
	//
	// GiteePrd@app.153082740420****.onaliyun.com
	AppPrincipalName *string `json:"AppPrincipalName,omitempty" xml:"AppPrincipalName,omitempty"`
	// The time when the application was installed. The value is a timestamp.
	//
	// example:
	//
	// 1737534146000
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted on the application.
	DelegatedScope *GetExternalApplicationResponseBodyExternalApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
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
	// 407426893752729****
	ForeignAppId *string `json:"ForeignAppId,omitempty" xml:"ForeignAppId,omitempty"`
	// The ID of the Alibaba Cloud account for which the application is installed.
	//
	// example:
	//
	// 173082740420****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The update time of the application. The value is a timestamp.
	//
	// example:
	//
	// 1737534146000
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s GetExternalApplicationResponseBodyExternalApplication) String() string {
	return dara.Prettify(s)
}

func (s GetExternalApplicationResponseBodyExternalApplication) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationResponseBodyExternalApplication) GetAppPrincipalName() *string {
	return s.AppPrincipalName
}

func (s *GetExternalApplicationResponseBodyExternalApplication) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetExternalApplicationResponseBodyExternalApplication) GetDelegatedScope() *GetExternalApplicationResponseBodyExternalApplicationDelegatedScope {
	return s.DelegatedScope
}

func (s *GetExternalApplicationResponseBodyExternalApplication) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetExternalApplicationResponseBodyExternalApplication) GetForeignAppId() *string {
	return s.ForeignAppId
}

func (s *GetExternalApplicationResponseBodyExternalApplication) GetTenantId() *string {
	return s.TenantId
}

func (s *GetExternalApplicationResponseBodyExternalApplication) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetAppPrincipalName(v string) *GetExternalApplicationResponseBodyExternalApplication {
	s.AppPrincipalName = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetCreateDate(v string) *GetExternalApplicationResponseBodyExternalApplication {
	s.CreateDate = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetDelegatedScope(v *GetExternalApplicationResponseBodyExternalApplicationDelegatedScope) *GetExternalApplicationResponseBodyExternalApplication {
	s.DelegatedScope = v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetDisplayName(v string) *GetExternalApplicationResponseBodyExternalApplication {
	s.DisplayName = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetForeignAppId(v string) *GetExternalApplicationResponseBodyExternalApplication {
	s.ForeignAppId = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetTenantId(v string) *GetExternalApplicationResponseBodyExternalApplication {
	s.TenantId = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplication) SetUpdateDate(v string) *GetExternalApplicationResponseBodyExternalApplication {
	s.UpdateDate = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplication) Validate() error {
	return dara.Validate(s)
}

type GetExternalApplicationResponseBodyExternalApplicationDelegatedScope struct {
	// The information about the permissions that are granted on the application.
	PredefinedScopes *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s GetExternalApplicationResponseBodyExternalApplicationDelegatedScope) String() string {
	return dara.Prettify(s)
}

func (s GetExternalApplicationResponseBodyExternalApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScope) GetPredefinedScopes() *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes {
	return s.PredefinedScopes
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScope) SetPredefinedScopes(v *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) *GetExternalApplicationResponseBodyExternalApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScope) Validate() error {
	return dara.Validate(s)
}

type GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) String() string {
	return dara.Prettify(s)
}

func (s GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) GetPredefinedScope() []*GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	return s.PredefinedScope
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopes) Validate() error {
	return dara.Validate(s)
}

type GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	// The description of the permission.
	//
	// example:
	//
	// Obtains the OpenID of the user. This is the default scope and cannot be deleted.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the permission.
	//
	// example:
	//
	// openid
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return dara.Prettify(s)
}

func (s GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) GetDescription() *string {
	return s.Description
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) GetName() *string {
	return s.Name
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) Validate() error {
	return dara.Validate(s)
}
