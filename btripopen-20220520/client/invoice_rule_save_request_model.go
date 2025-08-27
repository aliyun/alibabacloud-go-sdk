// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleSaveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllEmploye(v bool) *InvoiceRuleSaveRequest
	GetAllEmploye() *bool
	SetEntities(v []*InvoiceRuleSaveRequestEntities) *InvoiceRuleSaveRequest
	GetEntities() []*InvoiceRuleSaveRequestEntities
	SetScope(v int32) *InvoiceRuleSaveRequest
	GetScope() *int32
	SetThirdPartId(v string) *InvoiceRuleSaveRequest
	GetThirdPartId() *string
}

type InvoiceRuleSaveRequest struct {
	// example:
	//
	// true
	AllEmploye *bool                             `json:"all_employe,omitempty" xml:"all_employe,omitempty"`
	Entities   []*InvoiceRuleSaveRequestEntities `json:"entities,omitempty" xml:"entities,omitempty" type:"Repeated"`
	Scope      *int32                            `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i123
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s InvoiceRuleSaveRequest) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleSaveRequest) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveRequest) GetAllEmploye() *bool {
	return s.AllEmploye
}

func (s *InvoiceRuleSaveRequest) GetEntities() []*InvoiceRuleSaveRequestEntities {
	return s.Entities
}

func (s *InvoiceRuleSaveRequest) GetScope() *int32 {
	return s.Scope
}

func (s *InvoiceRuleSaveRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *InvoiceRuleSaveRequest) SetAllEmploye(v bool) *InvoiceRuleSaveRequest {
	s.AllEmploye = &v
	return s
}

func (s *InvoiceRuleSaveRequest) SetEntities(v []*InvoiceRuleSaveRequestEntities) *InvoiceRuleSaveRequest {
	s.Entities = v
	return s
}

func (s *InvoiceRuleSaveRequest) SetScope(v int32) *InvoiceRuleSaveRequest {
	s.Scope = &v
	return s
}

func (s *InvoiceRuleSaveRequest) SetThirdPartId(v string) *InvoiceRuleSaveRequest {
	s.ThirdPartId = &v
	return s
}

func (s *InvoiceRuleSaveRequest) Validate() error {
	return dara.Validate(s)
}

type InvoiceRuleSaveRequestEntities struct {
	// example:
	//
	// 123
	Id   *string `json:"id,omitempty" xml:"id,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s InvoiceRuleSaveRequestEntities) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleSaveRequestEntities) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveRequestEntities) GetId() *string {
	return s.Id
}

func (s *InvoiceRuleSaveRequestEntities) GetName() *string {
	return s.Name
}

func (s *InvoiceRuleSaveRequestEntities) GetType() *int32 {
	return s.Type
}

func (s *InvoiceRuleSaveRequestEntities) SetId(v string) *InvoiceRuleSaveRequestEntities {
	s.Id = &v
	return s
}

func (s *InvoiceRuleSaveRequestEntities) SetName(v string) *InvoiceRuleSaveRequestEntities {
	s.Name = &v
	return s
}

func (s *InvoiceRuleSaveRequestEntities) SetType(v int32) *InvoiceRuleSaveRequestEntities {
	s.Type = &v
	return s
}

func (s *InvoiceRuleSaveRequestEntities) Validate() error {
	return dara.Validate(s)
}
