// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleSaveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllEmploye(v bool) *InvoiceRuleSaveShrinkRequest
	GetAllEmploye() *bool
	SetEntitiesShrink(v string) *InvoiceRuleSaveShrinkRequest
	GetEntitiesShrink() *string
	SetScope(v int32) *InvoiceRuleSaveShrinkRequest
	GetScope() *int32
	SetThirdPartId(v string) *InvoiceRuleSaveShrinkRequest
	GetThirdPartId() *string
}

type InvoiceRuleSaveShrinkRequest struct {
	// example:
	//
	// true
	AllEmploye     *bool   `json:"all_employe,omitempty" xml:"all_employe,omitempty"`
	EntitiesShrink *string `json:"entities,omitempty" xml:"entities,omitempty"`
	Scope          *int32  `json:"scope,omitempty" xml:"scope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// i123
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s InvoiceRuleSaveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleSaveShrinkRequest) GoString() string {
	return s.String()
}

func (s *InvoiceRuleSaveShrinkRequest) GetAllEmploye() *bool {
	return s.AllEmploye
}

func (s *InvoiceRuleSaveShrinkRequest) GetEntitiesShrink() *string {
	return s.EntitiesShrink
}

func (s *InvoiceRuleSaveShrinkRequest) GetScope() *int32 {
	return s.Scope
}

func (s *InvoiceRuleSaveShrinkRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *InvoiceRuleSaveShrinkRequest) SetAllEmploye(v bool) *InvoiceRuleSaveShrinkRequest {
	s.AllEmploye = &v
	return s
}

func (s *InvoiceRuleSaveShrinkRequest) SetEntitiesShrink(v string) *InvoiceRuleSaveShrinkRequest {
	s.EntitiesShrink = &v
	return s
}

func (s *InvoiceRuleSaveShrinkRequest) SetScope(v int32) *InvoiceRuleSaveShrinkRequest {
	s.Scope = &v
	return s
}

func (s *InvoiceRuleSaveShrinkRequest) SetThirdPartId(v string) *InvoiceRuleSaveShrinkRequest {
	s.ThirdPartId = &v
	return s
}

func (s *InvoiceRuleSaveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
