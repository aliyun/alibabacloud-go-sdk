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
	ExternalApplication *GetExternalApplicationResponseBodyExternalApplication `json:"ExternalApplication,omitempty" xml:"ExternalApplication,omitempty" type:"Struct"`
	RequestId           *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	if s.ExternalApplication != nil {
		if err := s.ExternalApplication.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExternalApplicationResponseBodyExternalApplication struct {
	AppPrincipalName *string                                                              `json:"AppPrincipalName,omitempty" xml:"AppPrincipalName,omitempty"`
	CreateDate       *string                                                              `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	DelegatedScope   *GetExternalApplicationResponseBodyExternalApplicationDelegatedScope `json:"DelegatedScope,omitempty" xml:"DelegatedScope,omitempty" type:"Struct"`
	DisplayName      *string                                                              `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ForeignAppId     *string                                                              `json:"ForeignAppId,omitempty" xml:"ForeignAppId,omitempty"`
	TenantId         *string                                                              `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UpdateDate       *string                                                              `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
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
	if s.DelegatedScope != nil {
		if err := s.DelegatedScope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetExternalApplicationResponseBodyExternalApplicationDelegatedScope struct {
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
	if s.PredefinedScopes != nil {
		if err := s.PredefinedScopes.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type GetExternalApplicationResponseBodyExternalApplicationDelegatedScopePredefinedScopesPredefinedScope struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
