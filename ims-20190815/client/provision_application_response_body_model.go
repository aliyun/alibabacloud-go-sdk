// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProvisionApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationProvisionInfo(v *ProvisionApplicationResponseBodyApplicationProvisionInfo) *ProvisionApplicationResponseBody
	GetApplicationProvisionInfo() *ProvisionApplicationResponseBodyApplicationProvisionInfo
	SetRequestId(v string) *ProvisionApplicationResponseBody
	GetRequestId() *string
}

type ProvisionApplicationResponseBody struct {
	ApplicationProvisionInfo *ProvisionApplicationResponseBodyApplicationProvisionInfo `json:"ApplicationProvisionInfo,omitempty" xml:"ApplicationProvisionInfo,omitempty" type:"Struct"`
	RequestId                *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ProvisionApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProvisionApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationResponseBody) GetApplicationProvisionInfo() *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	return s.ApplicationProvisionInfo
}

func (s *ProvisionApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProvisionApplicationResponseBody) SetApplicationProvisionInfo(v *ProvisionApplicationResponseBodyApplicationProvisionInfo) *ProvisionApplicationResponseBody {
	s.ApplicationProvisionInfo = v
	return s
}

func (s *ProvisionApplicationResponseBody) SetRequestId(v string) *ProvisionApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProvisionApplicationResponseBody) Validate() error {
	if s.ApplicationProvisionInfo != nil {
		if err := s.ApplicationProvisionInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProvisionApplicationResponseBodyApplicationProvisionInfo struct {
	AccountId        *string                                                                 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AppId            *string                                                                 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName          *string                                                                 `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPrincipalName *string                                                                 `json:"AppPrincipalName,omitempty" xml:"AppPrincipalName,omitempty"`
	CreateDate       *string                                                                 `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DelegatedScope   *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	DisplayName      *string                                                                 `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	TenantId         *string                                                                 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UpdateDate       *string                                                                 `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfo) String() string {
	return dara.Prettify(s)
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfo) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) GetAccountId() *string {
	return s.AccountId
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) GetAppId() *string {
	return s.AppId
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) GetAppName() *string {
	return s.AppName
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) GetAppPrincipalName() *string {
	return s.AppPrincipalName
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) GetDelegatedScope() *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope {
	return s.DelegatedScope
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) GetTenantId() *string {
	return s.TenantId
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) GetUpdateDate() *string {
	return s.UpdateDate
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetAccountId(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.AccountId = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetAppId(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.AppId = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetAppName(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.AppName = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetAppPrincipalName(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.AppPrincipalName = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetCreateDate(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.CreateDate = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetDelegatedScope(v *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.DelegatedScope = v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetDisplayName(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.DisplayName = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetTenantId(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.TenantId = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) SetUpdateDate(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfo {
	s.UpdateDate = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfo) Validate() error {
	if s.DelegatedScope != nil {
		if err := s.DelegatedScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope struct {
	PredefinedScopes *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes `json:"PredefinedScopes,omitempty" xml:"PredefinedScopes,omitempty" type:"Struct"`
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope) String() string {
	return dara.Prettify(s)
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope) GetPredefinedScopes() *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes {
	return s.PredefinedScopes
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope) SetPredefinedScopes(v *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope {
	s.PredefinedScopes = v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScope) Validate() error {
	if s.PredefinedScopes != nil {
		if err := s.PredefinedScopes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes struct {
	PredefinedScope []*ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope `json:"PredefinedScope,omitempty" xml:"PredefinedScope,omitempty" type:"Repeated"`
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) String() string {
	return dara.Prettify(s)
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) GetPredefinedScope() []*ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	return s.PredefinedScope
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) SetPredefinedScope(v []*ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes {
	s.PredefinedScope = v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopes) Validate() error {
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

type ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) String() string {
	return dara.Prettify(s)
}

func (s ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GetDescription() *string {
	return s.Description
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) GetName() *string {
	return s.Name
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetDescription(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Description = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) SetName(v string) *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope {
	s.Name = &v
	return s
}

func (s *ProvisionApplicationResponseBodyApplicationProvisionInfoDelegatedScopePredefinedScopesPredefinedScope) Validate() error {
	return dara.Validate(s)
}
