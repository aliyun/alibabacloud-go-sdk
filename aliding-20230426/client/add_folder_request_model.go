// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *AddFolderRequest
	GetName() *string
	SetOption(v *AddFolderRequestOption) *AddFolderRequest
	GetOption() *AddFolderRequestOption
	SetParentId(v string) *AddFolderRequest
	GetParentId() *string
	SetSpaceId(v string) *AddFolderRequest
	GetSpaceId() *string
	SetTenantContext(v *AddFolderRequestTenantContext) *AddFolderRequest
	GetTenantContext() *AddFolderRequestTenantContext
}

type AddFolderRequest struct {
	// This parameter is required.
	Name   *string                 `json:"Name,omitempty" xml:"Name,omitempty"`
	Option *AddFolderRequestOption `json:"Option,omitempty" xml:"Option,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 140822073803
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xPar2SZ63KodG3aV
	SpaceId       *string                        `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContext *AddFolderRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s AddFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFolderRequest) GoString() string {
	return s.String()
}

func (s *AddFolderRequest) GetName() *string {
	return s.Name
}

func (s *AddFolderRequest) GetOption() *AddFolderRequestOption {
	return s.Option
}

func (s *AddFolderRequest) GetParentId() *string {
	return s.ParentId
}

func (s *AddFolderRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *AddFolderRequest) GetTenantContext() *AddFolderRequestTenantContext {
	return s.TenantContext
}

func (s *AddFolderRequest) SetName(v string) *AddFolderRequest {
	s.Name = &v
	return s
}

func (s *AddFolderRequest) SetOption(v *AddFolderRequestOption) *AddFolderRequest {
	s.Option = v
	return s
}

func (s *AddFolderRequest) SetParentId(v string) *AddFolderRequest {
	s.ParentId = &v
	return s
}

func (s *AddFolderRequest) SetSpaceId(v string) *AddFolderRequest {
	s.SpaceId = &v
	return s
}

func (s *AddFolderRequest) SetTenantContext(v *AddFolderRequestTenantContext) *AddFolderRequest {
	s.TenantContext = v
	return s
}

func (s *AddFolderRequest) Validate() error {
	return dara.Validate(s)
}

type AddFolderRequestOption struct {
	AppProperties []*AddFolderRequestOptionAppProperties `json:"AppProperties,omitempty" xml:"AppProperties,omitempty" type:"Repeated"`
	// example:
	//
	// AUTO_RENAME
	ConflictStrategy *string `json:"ConflictStrategy,omitempty" xml:"ConflictStrategy,omitempty"`
}

func (s AddFolderRequestOption) String() string {
	return dara.Prettify(s)
}

func (s AddFolderRequestOption) GoString() string {
	return s.String()
}

func (s *AddFolderRequestOption) GetAppProperties() []*AddFolderRequestOptionAppProperties {
	return s.AppProperties
}

func (s *AddFolderRequestOption) GetConflictStrategy() *string {
	return s.ConflictStrategy
}

func (s *AddFolderRequestOption) SetAppProperties(v []*AddFolderRequestOptionAppProperties) *AddFolderRequestOption {
	s.AppProperties = v
	return s
}

func (s *AddFolderRequestOption) SetConflictStrategy(v string) *AddFolderRequestOption {
	s.ConflictStrategy = &v
	return s
}

func (s *AddFolderRequestOption) Validate() error {
	return dara.Validate(s)
}

type AddFolderRequestOptionAppProperties struct {
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PUBLIC
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s AddFolderRequestOptionAppProperties) String() string {
	return dara.Prettify(s)
}

func (s AddFolderRequestOptionAppProperties) GoString() string {
	return s.String()
}

func (s *AddFolderRequestOptionAppProperties) GetName() *string {
	return s.Name
}

func (s *AddFolderRequestOptionAppProperties) GetValue() *string {
	return s.Value
}

func (s *AddFolderRequestOptionAppProperties) GetVisibility() *string {
	return s.Visibility
}

func (s *AddFolderRequestOptionAppProperties) SetName(v string) *AddFolderRequestOptionAppProperties {
	s.Name = &v
	return s
}

func (s *AddFolderRequestOptionAppProperties) SetValue(v string) *AddFolderRequestOptionAppProperties {
	s.Value = &v
	return s
}

func (s *AddFolderRequestOptionAppProperties) SetVisibility(v string) *AddFolderRequestOptionAppProperties {
	s.Visibility = &v
	return s
}

func (s *AddFolderRequestOptionAppProperties) Validate() error {
	return dara.Validate(s)
}

type AddFolderRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s AddFolderRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s AddFolderRequestTenantContext) GoString() string {
	return s.String()
}

func (s *AddFolderRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *AddFolderRequestTenantContext) SetTenantId(v string) *AddFolderRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *AddFolderRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
