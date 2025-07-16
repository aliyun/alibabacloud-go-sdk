// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMultiDimTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseId(v string) *AddMultiDimTableRequest
	GetBaseId() *string
	SetFields(v []*AddMultiDimTableRequestFields) *AddMultiDimTableRequest
	GetFields() []*AddMultiDimTableRequestFields
	SetName(v string) *AddMultiDimTableRequest
	GetName() *string
	SetTenantContext(v *AddMultiDimTableRequestTenantContext) *AddMultiDimTableRequest
	GetTenantContext() *AddMultiDimTableRequestTenantContext
}

type AddMultiDimTableRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// r1R7q3QmWew5lo02fxB7nxxxxxxxx
	BaseId        *string                               `json:"BaseId,omitempty" xml:"BaseId,omitempty"`
	Fields        []*AddMultiDimTableRequestFields      `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	Name          *string                               `json:"Name,omitempty" xml:"Name,omitempty"`
	TenantContext *AddMultiDimTableRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s AddMultiDimTableRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMultiDimTableRequest) GoString() string {
	return s.String()
}

func (s *AddMultiDimTableRequest) GetBaseId() *string {
	return s.BaseId
}

func (s *AddMultiDimTableRequest) GetFields() []*AddMultiDimTableRequestFields {
	return s.Fields
}

func (s *AddMultiDimTableRequest) GetName() *string {
	return s.Name
}

func (s *AddMultiDimTableRequest) GetTenantContext() *AddMultiDimTableRequestTenantContext {
	return s.TenantContext
}

func (s *AddMultiDimTableRequest) SetBaseId(v string) *AddMultiDimTableRequest {
	s.BaseId = &v
	return s
}

func (s *AddMultiDimTableRequest) SetFields(v []*AddMultiDimTableRequestFields) *AddMultiDimTableRequest {
	s.Fields = v
	return s
}

func (s *AddMultiDimTableRequest) SetName(v string) *AddMultiDimTableRequest {
	s.Name = &v
	return s
}

func (s *AddMultiDimTableRequest) SetTenantContext(v *AddMultiDimTableRequestTenantContext) *AddMultiDimTableRequest {
	s.TenantContext = v
	return s
}

func (s *AddMultiDimTableRequest) Validate() error {
	return dara.Validate(s)
}

type AddMultiDimTableRequestFields struct {
	// This parameter is required.
	Name     *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Property map[string]interface{} `json:"Property,omitempty" xml:"Property,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AddMultiDimTableRequestFields) String() string {
	return dara.Prettify(s)
}

func (s AddMultiDimTableRequestFields) GoString() string {
	return s.String()
}

func (s *AddMultiDimTableRequestFields) GetName() *string {
	return s.Name
}

func (s *AddMultiDimTableRequestFields) GetProperty() map[string]interface{} {
	return s.Property
}

func (s *AddMultiDimTableRequestFields) GetType() *string {
	return s.Type
}

func (s *AddMultiDimTableRequestFields) SetName(v string) *AddMultiDimTableRequestFields {
	s.Name = &v
	return s
}

func (s *AddMultiDimTableRequestFields) SetProperty(v map[string]interface{}) *AddMultiDimTableRequestFields {
	s.Property = v
	return s
}

func (s *AddMultiDimTableRequestFields) SetType(v string) *AddMultiDimTableRequestFields {
	s.Type = &v
	return s
}

func (s *AddMultiDimTableRequestFields) Validate() error {
	return dara.Validate(s)
}

type AddMultiDimTableRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s AddMultiDimTableRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s AddMultiDimTableRequestTenantContext) GoString() string {
	return s.String()
}

func (s *AddMultiDimTableRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *AddMultiDimTableRequestTenantContext) SetTenantId(v string) *AddMultiDimTableRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *AddMultiDimTableRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
