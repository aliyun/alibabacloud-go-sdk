// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExternalApplications(v *ListExternalApplicationsResponseBodyExternalApplications) *ListExternalApplicationsResponseBody
	GetExternalApplications() *ListExternalApplicationsResponseBodyExternalApplications
	SetIsTruncated(v bool) *ListExternalApplicationsResponseBody
	GetIsTruncated() *bool
	SetMarker(v string) *ListExternalApplicationsResponseBody
	GetMarker() *string
	SetRequestId(v string) *ListExternalApplicationsResponseBody
	GetRequestId() *string
}

type ListExternalApplicationsResponseBody struct {
	// The information about the external applications.
	ExternalApplications *ListExternalApplicationsResponseBodyExternalApplications `json:"ExternalApplications,omitempty" xml:"ExternalApplications,omitempty" type:"Struct"`
	// Indicates whether the response is truncated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// >  This parameter is returned only when `IsTruncated` is `true`.
	//
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 936E1D9C-157D-45BD-8A3B-81C0716EB077
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExternalApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExternalApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponseBody) GetExternalApplications() *ListExternalApplicationsResponseBodyExternalApplications {
	return s.ExternalApplications
}

func (s *ListExternalApplicationsResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListExternalApplicationsResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *ListExternalApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExternalApplicationsResponseBody) SetExternalApplications(v *ListExternalApplicationsResponseBodyExternalApplications) *ListExternalApplicationsResponseBody {
	s.ExternalApplications = v
	return s
}

func (s *ListExternalApplicationsResponseBody) SetIsTruncated(v bool) *ListExternalApplicationsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListExternalApplicationsResponseBody) SetMarker(v string) *ListExternalApplicationsResponseBody {
	s.Marker = &v
	return s
}

func (s *ListExternalApplicationsResponseBody) SetRequestId(v string) *ListExternalApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExternalApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListExternalApplicationsResponseBodyExternalApplications struct {
	ExternalApplication []*ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication `json:"ExternalApplication,omitempty" xml:"ExternalApplication,omitempty" type:"Repeated"`
}

func (s ListExternalApplicationsResponseBodyExternalApplications) String() string {
	return dara.Prettify(s)
}

func (s ListExternalApplicationsResponseBodyExternalApplications) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponseBodyExternalApplications) GetExternalApplication() []*ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	return s.ExternalApplication
}

func (s *ListExternalApplicationsResponseBodyExternalApplications) SetExternalApplication(v []*ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) *ListExternalApplicationsResponseBodyExternalApplications {
	s.ExternalApplication = v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplications) Validate() error {
	return dara.Validate(s)
}

type ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication struct {
	// The name of the external application principal. The value is in the `<app_name>@app.<account_id>.onaliyun.com` format.
	//
	// example:
	//
	// GiteePrd@app.153082740420****.onaliyun.com
	AppPrincipalName *string `json:"AppPrincipalName,omitempty" xml:"AppPrincipalName,omitempty"`
	// The time when the external application was installed. The value is a timestamp.
	//
	// example:
	//
	// 1603693318000
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted to the external application.
	DelegatedScope *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the external application.
	//
	// example:
	//
	// GiteeAliyun
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the external application.
	//
	// example:
	//
	// 407426893752729****
	ForeignAppId *string `json:"ForeignAppId,omitempty" xml:"ForeignAppId,omitempty"`
	// The ID of the Alibaba Cloud account for which the external application was installed.
	//
	// example:
	//
	// 173082740420****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The update time of the external application. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) String() string {
	return dara.Prettify(s)
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) GetAppPrincipalName() *string {
	return s.AppPrincipalName
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) GetDelegatedScope() *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope {
	return s.DelegatedScope
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) GetForeignAppId() *string {
	return s.ForeignAppId
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) GetTenantId() *string {
	return s.TenantId
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetAppPrincipalName(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.AppPrincipalName = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetCreateDate(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.CreateDate = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetDelegatedScope(v *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.DelegatedScope = v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetDisplayName(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.DisplayName = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetForeignAppId(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.ForeignAppId = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetTenantId(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.TenantId = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) SetUpdateDate(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication {
	s.UpdateDate = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplication) Validate() error {
	return dara.Validate(s)
}

type ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope struct {
	// The information about the permissions that are granted to the external application.
	PredefinedScopes *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope) String() string {
	return dara.Prettify(s)
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope) GetPredefinedScopes() *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes {
	return s.PredefinedScopes
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope) SetPredefinedScopes(v *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope {
	s.PredefinedScopes = v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScope) Validate() error {
	return dara.Validate(s)
}

type ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes struct {
	PredefinedScope []*ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes) String() string {
	return dara.Prettify(s)
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes) GetPredefinedScope() []*ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	return s.PredefinedScope
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes) SetPredefinedScope(v []*ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopes) Validate() error {
	return dara.Validate(s)
}

type ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
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

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return dara.Prettify(s)
}

func (s ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) GetDescription() *string {
	return s.Description
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) GetName() *string {
	return s.Name
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *ListExternalApplicationsResponseBodyExternalApplicationsExternalApplicationDelegatedScopePredefinedScopesPredefinedScope) Validate() error {
	return dara.Validate(s)
}
