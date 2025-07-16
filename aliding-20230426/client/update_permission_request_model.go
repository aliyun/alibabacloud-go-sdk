// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *UpdatePermissionRequest
	GetDentryUuid() *string
	SetMembers(v []*UpdatePermissionRequestMembers) *UpdatePermissionRequest
	GetMembers() []*UpdatePermissionRequestMembers
	SetOption(v *UpdatePermissionRequestOption) *UpdatePermissionRequest
	GetOption() *UpdatePermissionRequestOption
	SetRoleId(v string) *UpdatePermissionRequest
	GetRoleId() *string
	SetTenantContext(v *UpdatePermissionRequestTenantContext) *UpdatePermissionRequest
	GetTenantContext() *UpdatePermissionRequestTenantContext
}

type UpdatePermissionRequest struct {
	// example:
	//
	// kDnRL6jAJMLgNkw7tBnw5aY4VyMoPYe1
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// This parameter is required.
	Members []*UpdatePermissionRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Option  *UpdatePermissionRequestOption    `json:"Option,omitempty" xml:"Option,omitempty" type:"Struct"`
	// example:
	//
	// READER
	RoleId        *string                               `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	TenantContext *UpdatePermissionRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s UpdatePermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePermissionRequest) GoString() string {
	return s.String()
}

func (s *UpdatePermissionRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *UpdatePermissionRequest) GetMembers() []*UpdatePermissionRequestMembers {
	return s.Members
}

func (s *UpdatePermissionRequest) GetOption() *UpdatePermissionRequestOption {
	return s.Option
}

func (s *UpdatePermissionRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *UpdatePermissionRequest) GetTenantContext() *UpdatePermissionRequestTenantContext {
	return s.TenantContext
}

func (s *UpdatePermissionRequest) SetDentryUuid(v string) *UpdatePermissionRequest {
	s.DentryUuid = &v
	return s
}

func (s *UpdatePermissionRequest) SetMembers(v []*UpdatePermissionRequestMembers) *UpdatePermissionRequest {
	s.Members = v
	return s
}

func (s *UpdatePermissionRequest) SetOption(v *UpdatePermissionRequestOption) *UpdatePermissionRequest {
	s.Option = v
	return s
}

func (s *UpdatePermissionRequest) SetRoleId(v string) *UpdatePermissionRequest {
	s.RoleId = &v
	return s
}

func (s *UpdatePermissionRequest) SetTenantContext(v *UpdatePermissionRequestTenantContext) *UpdatePermissionRequest {
	s.TenantContext = v
	return s
}

func (s *UpdatePermissionRequest) Validate() error {
	return dara.Validate(s)
}

type UpdatePermissionRequestMembers struct {
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

func (s UpdatePermissionRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s UpdatePermissionRequestMembers) GoString() string {
	return s.String()
}

func (s *UpdatePermissionRequestMembers) GetCorpId() *string {
	return s.CorpId
}

func (s *UpdatePermissionRequestMembers) GetId() *string {
	return s.Id
}

func (s *UpdatePermissionRequestMembers) GetType() *string {
	return s.Type
}

func (s *UpdatePermissionRequestMembers) SetCorpId(v string) *UpdatePermissionRequestMembers {
	s.CorpId = &v
	return s
}

func (s *UpdatePermissionRequestMembers) SetId(v string) *UpdatePermissionRequestMembers {
	s.Id = &v
	return s
}

func (s *UpdatePermissionRequestMembers) SetType(v string) *UpdatePermissionRequestMembers {
	s.Type = &v
	return s
}

func (s *UpdatePermissionRequestMembers) Validate() error {
	return dara.Validate(s)
}

type UpdatePermissionRequestOption struct {
	// example:
	//
	// 10
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s UpdatePermissionRequestOption) String() string {
	return dara.Prettify(s)
}

func (s UpdatePermissionRequestOption) GoString() string {
	return s.String()
}

func (s *UpdatePermissionRequestOption) GetDuration() *int64 {
	return s.Duration
}

func (s *UpdatePermissionRequestOption) SetDuration(v int64) *UpdatePermissionRequestOption {
	s.Duration = &v
	return s
}

func (s *UpdatePermissionRequestOption) Validate() error {
	return dara.Validate(s)
}

type UpdatePermissionRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdatePermissionRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdatePermissionRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdatePermissionRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdatePermissionRequestTenantContext) SetTenantId(v string) *UpdatePermissionRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdatePermissionRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
