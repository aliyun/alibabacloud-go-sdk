// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *DeletePermissionRequest
	GetDentryUuid() *string
	SetMembers(v []*DeletePermissionRequestMembers) *DeletePermissionRequest
	GetMembers() []*DeletePermissionRequestMembers
	SetRoleId(v string) *DeletePermissionRequest
	GetRoleId() *string
	SetTenantContext(v *DeletePermissionRequestTenantContext) *DeletePermissionRequest
	GetTenantContext() *DeletePermissionRequestTenantContext
}

type DeletePermissionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a9E05BDRVQRkezKGCE3nlwPDJ63zgkYA
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// This parameter is required.
	Members []*DeletePermissionRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// MANAGER
	RoleId        *string                               `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	TenantContext *DeletePermissionRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s DeletePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePermissionRequest) GoString() string {
	return s.String()
}

func (s *DeletePermissionRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *DeletePermissionRequest) GetMembers() []*DeletePermissionRequestMembers {
	return s.Members
}

func (s *DeletePermissionRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *DeletePermissionRequest) GetTenantContext() *DeletePermissionRequestTenantContext {
	return s.TenantContext
}

func (s *DeletePermissionRequest) SetDentryUuid(v string) *DeletePermissionRequest {
	s.DentryUuid = &v
	return s
}

func (s *DeletePermissionRequest) SetMembers(v []*DeletePermissionRequestMembers) *DeletePermissionRequest {
	s.Members = v
	return s
}

func (s *DeletePermissionRequest) SetRoleId(v string) *DeletePermissionRequest {
	s.RoleId = &v
	return s
}

func (s *DeletePermissionRequest) SetTenantContext(v *DeletePermissionRequestTenantContext) *DeletePermissionRequest {
	s.TenantContext = v
	return s
}

func (s *DeletePermissionRequest) Validate() error {
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeletePermissionRequestMembers struct {
	// example:
	//
	// 123456
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// ORG
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// ORG
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeletePermissionRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s DeletePermissionRequestMembers) GoString() string {
	return s.String()
}

func (s *DeletePermissionRequestMembers) GetCorpId() *string {
	return s.CorpId
}

func (s *DeletePermissionRequestMembers) GetId() *string {
	return s.Id
}

func (s *DeletePermissionRequestMembers) GetType() *string {
	return s.Type
}

func (s *DeletePermissionRequestMembers) SetCorpId(v string) *DeletePermissionRequestMembers {
	s.CorpId = &v
	return s
}

func (s *DeletePermissionRequestMembers) SetId(v string) *DeletePermissionRequestMembers {
	s.Id = &v
	return s
}

func (s *DeletePermissionRequestMembers) SetType(v string) *DeletePermissionRequestMembers {
	s.Type = &v
	return s
}

func (s *DeletePermissionRequestMembers) Validate() error {
	return dara.Validate(s)
}

type DeletePermissionRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s DeletePermissionRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s DeletePermissionRequestTenantContext) GoString() string {
	return s.String()
}

func (s *DeletePermissionRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *DeletePermissionRequestTenantContext) SetTenantId(v string) *DeletePermissionRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *DeletePermissionRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
