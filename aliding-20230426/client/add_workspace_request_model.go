// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *AddWorkspaceRequest
	GetName() *string
	SetOption(v *AddWorkspaceRequestOption) *AddWorkspaceRequest
	GetOption() *AddWorkspaceRequestOption
	SetTenantContext(v *AddWorkspaceRequestTenantContext) *AddWorkspaceRequest
	GetTenantContext() *AddWorkspaceRequestTenantContext
}

type AddWorkspaceRequest struct {
	// This parameter is required.
	Name          *string                           `json:"Name,omitempty" xml:"Name,omitempty"`
	Option        *AddWorkspaceRequestOption        `json:"Option,omitempty" xml:"Option,omitempty" type:"Struct"`
	TenantContext *AddWorkspaceRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s AddWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceRequest) GetName() *string {
	return s.Name
}

func (s *AddWorkspaceRequest) GetOption() *AddWorkspaceRequestOption {
	return s.Option
}

func (s *AddWorkspaceRequest) GetTenantContext() *AddWorkspaceRequestTenantContext {
	return s.TenantContext
}

func (s *AddWorkspaceRequest) SetName(v string) *AddWorkspaceRequest {
	s.Name = &v
	return s
}

func (s *AddWorkspaceRequest) SetOption(v *AddWorkspaceRequestOption) *AddWorkspaceRequest {
	s.Option = v
	return s
}

func (s *AddWorkspaceRequest) SetTenantContext(v *AddWorkspaceRequestTenantContext) *AddWorkspaceRequest {
	s.TenantContext = v
	return s
}

func (s *AddWorkspaceRequest) Validate() error {
	if s.Option != nil {
		if err := s.Option.Validate(); err != nil {
			return err
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddWorkspaceRequestOption struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// lHiicjNFM2iSFYSdz2iPuI8ZwiEiE
	TeamId *string `json:"TeamId,omitempty" xml:"TeamId,omitempty"`
}

func (s AddWorkspaceRequestOption) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceRequestOption) GoString() string {
	return s.String()
}

func (s *AddWorkspaceRequestOption) GetDescription() *string {
	return s.Description
}

func (s *AddWorkspaceRequestOption) GetTeamId() *string {
	return s.TeamId
}

func (s *AddWorkspaceRequestOption) SetDescription(v string) *AddWorkspaceRequestOption {
	s.Description = &v
	return s
}

func (s *AddWorkspaceRequestOption) SetTeamId(v string) *AddWorkspaceRequestOption {
	s.TeamId = &v
	return s
}

func (s *AddWorkspaceRequestOption) Validate() error {
	return dara.Validate(s)
}

type AddWorkspaceRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s AddWorkspaceRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceRequestTenantContext) GoString() string {
	return s.String()
}

func (s *AddWorkspaceRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *AddWorkspaceRequestTenantContext) SetTenantId(v string) *AddWorkspaceRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *AddWorkspaceRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
