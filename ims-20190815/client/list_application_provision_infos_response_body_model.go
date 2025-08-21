// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationProvisionInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationProvisionInfos(v *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) *ListApplicationProvisionInfosResponseBody
	GetApplicationProvisionInfos() *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos
	SetRequestId(v string) *ListApplicationProvisionInfosResponseBody
	GetRequestId() *string
}

type ListApplicationProvisionInfosResponseBody struct {
	// The information about the installed applications.
	ApplicationProvisionInfos *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos `json:"ApplicationProvisionInfos,omitempty" xml:"ApplicationProvisionInfos,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E403EBFD-C997-489D-BFC7-37C05E66D67C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListApplicationProvisionInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBody) GetApplicationProvisionInfos() *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos {
	return s.ApplicationProvisionInfos
}

func (s *ListApplicationProvisionInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationProvisionInfosResponseBody) SetApplicationProvisionInfos(v *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) *ListApplicationProvisionInfosResponseBody {
	s.ApplicationProvisionInfos = v
	return s
}

func (s *ListApplicationProvisionInfosResponseBody) SetRequestId(v string) *ListApplicationProvisionInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos struct {
	ApplicationProvisionInfo []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo `json:"ApplicationProvisionInfo,omitempty" xml:"ApplicationProvisionInfo,omitempty" type:"Repeated"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) GetApplicationProvisionInfo() []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	return s.ApplicationProvisionInfo
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) SetApplicationProvisionInfo(v []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos {
	s.ApplicationProvisionInfo = v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfos) Validate() error {
	return dara.Validate(s)
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 177242285274****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 452392483381546****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// GiteePrd
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The time when the application was installed. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The information about the permissions that are granted to the application.
	DelegatedScope *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	// The display name of the application.
	//
	// example:
	//
	// GiteeAliyun
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The update time. The value is a timestamp.
	//
	// example:
	//
	// 1603693518000
	UpdateDate *string `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetAccountId() *string {
	return s.AccountId
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetAppId() *string {
	return s.AppId
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetAppName() *string {
	return s.AppName
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetDelegatedScope() *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope {
	return s.DelegatedScope
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetAccountId(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.AccountId = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetAppId(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.AppId = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetAppName(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.AppName = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetCreateDate(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.CreateDate = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetDelegatedScope(v *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.DelegatedScope = v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetDisplayName(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.DisplayName = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) SetUpdateDate(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo {
	s.UpdateDate = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfo) Validate() error {
	return dara.Validate(s)
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope struct {
	// The information about the permissions that are granted to the application.
	PredefinedScopes *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) GetPredefinedScopes() *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes {
	return s.PredefinedScopes
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) SetPredefinedScopes(v *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope {
	s.PredefinedScopes = v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScope) Validate() error {
	return dara.Validate(s)
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes struct {
	PredefinedScope []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) GetPredefinedScope() []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	return s.PredefinedScope
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) SetPredefinedScope(v []*ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopes) Validate() error {
	return dara.Validate(s)
}

type ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope struct {
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

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GetName() *string {
	return s.Name
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *ListApplicationProvisionInfosResponseBodyApplicationProvisionInfosApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) Validate() error {
	return dara.Validate(s)
}
