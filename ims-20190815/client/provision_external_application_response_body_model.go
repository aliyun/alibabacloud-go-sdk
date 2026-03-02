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
	ExternalApplication *ProvisionExternalApplicationResponseBodyExternalApplication `json:"ExternalApplication,omitempty" xml:"ExternalApplication,omitempty" type:"Struct"`
	RequestId           *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	if s.ExternalApplication != nil {
		if err := s.ExternalApplication.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProvisionExternalApplicationResponseBodyExternalApplication struct {
	AppPrincipalName *string                                                                    `json:"AppPrincipalName,omitempty" xml:"AppPrincipalName,omitempty"`
	CreateDate       *string                                                                    `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DelegatedScope   *ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	DisplayName      *string                                                                    `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ForeignAppId     *string                                                                    `json:"ForeignAppId,omitempty" xml:"ForeignAppId,omitempty"`
	TenantId         *string                                                                    `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UpdateDate       *string                                                                    `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
	if s.DelegatedScope != nil {
		if err := s.DelegatedScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScope struct {
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
	if s.PredefinedScopes != nil {
		if err := s.PredefinedScopes.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type ProvisionExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
