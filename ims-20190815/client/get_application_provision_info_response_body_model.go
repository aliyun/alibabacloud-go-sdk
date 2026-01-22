// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisionInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationProvisionInfo(v *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) *GetApplicationProvisionInfoResponseBody
	GetApplicationProvisionInfo() *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo
	SetRequestId(v string) *GetApplicationProvisionInfoResponseBody
	GetRequestId() *string
}

type GetApplicationProvisionInfoResponseBody struct {
	// The installation information about the application.
	ApplicationProvisionInfo *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo `json:"ApplicationProvisionInfo,omitempty" xml:"ApplicationProvisionInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 936E1D9C-157D-45BD-8A3B-81C0716EB078
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationProvisionInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisionInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoResponseBody) GetApplicationProvisionInfo() *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	return s.ApplicationProvisionInfo
}

func (s *GetApplicationProvisionInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationProvisionInfoResponseBody) SetApplicationProvisionInfo(v *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) *GetApplicationProvisionInfoResponseBody {
	s.ApplicationProvisionInfo = v
	return s
}

func (s *GetApplicationProvisionInfoResponseBody) SetRequestId(v string) *GetApplicationProvisionInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBody) Validate() error {
	if s.ApplicationProvisionInfo != nil {
		if err := s.ApplicationProvisionInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo struct {
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
	// The information about the scopes of permissions that are granted to the application.
	DelegatedScope *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
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

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) GetAccountId() *string {
	return s.AccountId
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) GetAppId() *string {
	return s.AppId
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) GetAppName() *string {
	return s.AppName
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) GetCreateDate() *string {
	return s.CreateDate
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) GetDelegatedScope() *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope {
	return s.DelegatedScope
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetAccountId(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.AccountId = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetAppId(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.AppId = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetAppName(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.AppName = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetCreateDate(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.CreateDate = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetDelegatedScope(v *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.DelegatedScope = v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetDisplayName(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.DisplayName = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) SetUpdateDate(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo {
	s.UpdateDate = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfo) Validate() error {
	if s.DelegatedScope != nil {
		if err := s.DelegatedScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope struct {
	// The information about the scopes of permissions that are granted to the application.
	PredefinedScopes *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope) GetPredefinedScopes() *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes {
	return s.PredefinedScopes
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope) SetPredefinedScopes(v *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope {
	s.PredefinedScopes = v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScope) Validate() error {
	if s.PredefinedScopes != nil {
		if err := s.PredefinedScopes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes struct {
	PredefinedScope []*GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) GetPredefinedScope() []*GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	return s.PredefinedScope
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) SetPredefinedScope(v []*GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) Validate() error {
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

type GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope struct {
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

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GetDescription() *string {
	return s.Description
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GetName() *string {
	return s.Name
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *GetApplicationProvisionInfoResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) Validate() error {
	return dara.Validate(s)
}
