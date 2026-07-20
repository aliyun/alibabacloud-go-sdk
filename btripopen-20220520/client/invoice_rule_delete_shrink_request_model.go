// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleDeleteShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelAll(v bool) *InvoiceRuleDeleteShrinkRequest
	GetDelAll() *bool
	SetEntitiesShrink(v string) *InvoiceRuleDeleteShrinkRequest
	GetEntitiesShrink() *string
	SetThirdPartId(v string) *InvoiceRuleDeleteShrinkRequest
	GetThirdPartId() *string
}

type InvoiceRuleDeleteShrinkRequest struct {
	// example:
	//
	// false
	DelAll         *bool   `json:"del_all,omitempty" xml:"del_all,omitempty"`
	EntitiesShrink *string `json:"entities,omitempty" xml:"entities,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 340049
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s InvoiceRuleDeleteShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleDeleteShrinkRequest) GoString() string {
	return s.String()
}

func (s *InvoiceRuleDeleteShrinkRequest) GetDelAll() *bool {
	return s.DelAll
}

func (s *InvoiceRuleDeleteShrinkRequest) GetEntitiesShrink() *string {
	return s.EntitiesShrink
}

func (s *InvoiceRuleDeleteShrinkRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *InvoiceRuleDeleteShrinkRequest) SetDelAll(v bool) *InvoiceRuleDeleteShrinkRequest {
	s.DelAll = &v
	return s
}

func (s *InvoiceRuleDeleteShrinkRequest) SetEntitiesShrink(v string) *InvoiceRuleDeleteShrinkRequest {
	s.EntitiesShrink = &v
	return s
}

func (s *InvoiceRuleDeleteShrinkRequest) SetThirdPartId(v string) *InvoiceRuleDeleteShrinkRequest {
	s.ThirdPartId = &v
	return s
}

func (s *InvoiceRuleDeleteShrinkRequest) Validate() error {
	return dara.Validate(s)
}
