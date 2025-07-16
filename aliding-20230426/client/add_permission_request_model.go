// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryUuid(v string) *AddPermissionRequest
	GetDentryUuid() *string
	SetMembers(v []*AddPermissionRequestMembers) *AddPermissionRequest
	GetMembers() []*AddPermissionRequestMembers
	SetOption(v *AddPermissionRequestOption) *AddPermissionRequest
	GetOption() *AddPermissionRequestOption
	SetRoleId(v string) *AddPermissionRequest
	GetRoleId() *string
	SetTenantContext(v *AddPermissionRequestTenantContext) *AddPermissionRequest
	GetTenantContext() *AddPermissionRequestTenantContext
}

type AddPermissionRequest struct {
	DentryUuid *string `json:"DentryUuid,omitempty" xml:"DentryUuid,omitempty"`
	// This parameter is required.
	Members []*AddPermissionRequestMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	Option  *AddPermissionRequestOption    `json:"Option,omitempty" xml:"Option,omitempty" type:"Struct"`
	// This parameter is required.
	RoleId        *string                            `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	TenantContext *AddPermissionRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s AddPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPermissionRequest) GoString() string {
	return s.String()
}

func (s *AddPermissionRequest) GetDentryUuid() *string {
	return s.DentryUuid
}

func (s *AddPermissionRequest) GetMembers() []*AddPermissionRequestMembers {
	return s.Members
}

func (s *AddPermissionRequest) GetOption() *AddPermissionRequestOption {
	return s.Option
}

func (s *AddPermissionRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *AddPermissionRequest) GetTenantContext() *AddPermissionRequestTenantContext {
	return s.TenantContext
}

func (s *AddPermissionRequest) SetDentryUuid(v string) *AddPermissionRequest {
	s.DentryUuid = &v
	return s
}

func (s *AddPermissionRequest) SetMembers(v []*AddPermissionRequestMembers) *AddPermissionRequest {
	s.Members = v
	return s
}

func (s *AddPermissionRequest) SetOption(v *AddPermissionRequestOption) *AddPermissionRequest {
	s.Option = v
	return s
}

func (s *AddPermissionRequest) SetRoleId(v string) *AddPermissionRequest {
	s.RoleId = &v
	return s
}

func (s *AddPermissionRequest) SetTenantContext(v *AddPermissionRequestTenantContext) *AddPermissionRequest {
	s.TenantContext = v
	return s
}

func (s *AddPermissionRequest) Validate() error {
	return dara.Validate(s)
}

type AddPermissionRequestMembers struct {
	// example:
	//
	// 123456
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// 123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// ORG
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddPermissionRequestMembers) String() string {
	return dara.Prettify(s)
}

func (s AddPermissionRequestMembers) GoString() string {
	return s.String()
}

func (s *AddPermissionRequestMembers) GetCorpId() *string {
	return s.CorpId
}

func (s *AddPermissionRequestMembers) GetId() *string {
	return s.Id
}

func (s *AddPermissionRequestMembers) GetType() *string {
	return s.Type
}

func (s *AddPermissionRequestMembers) SetCorpId(v string) *AddPermissionRequestMembers {
	s.CorpId = &v
	return s
}

func (s *AddPermissionRequestMembers) SetId(v string) *AddPermissionRequestMembers {
	s.Id = &v
	return s
}

func (s *AddPermissionRequestMembers) SetType(v string) *AddPermissionRequestMembers {
	s.Type = &v
	return s
}

func (s *AddPermissionRequestMembers) Validate() error {
	return dara.Validate(s)
}

type AddPermissionRequestOption struct {
	// example:
	//
	// 3600
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s AddPermissionRequestOption) String() string {
	return dara.Prettify(s)
}

func (s AddPermissionRequestOption) GoString() string {
	return s.String()
}

func (s *AddPermissionRequestOption) GetDuration() *int64 {
	return s.Duration
}

func (s *AddPermissionRequestOption) SetDuration(v int64) *AddPermissionRequestOption {
	s.Duration = &v
	return s
}

func (s *AddPermissionRequestOption) Validate() error {
	return dara.Validate(s)
}

type AddPermissionRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s AddPermissionRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s AddPermissionRequestTenantContext) GoString() string {
	return s.String()
}

func (s *AddPermissionRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *AddPermissionRequestTenantContext) SetTenantId(v string) *AddPermissionRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *AddPermissionRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
