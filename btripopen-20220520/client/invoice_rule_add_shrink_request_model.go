// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleAddShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntitiesShrink(v string) *InvoiceRuleAddShrinkRequest
	GetEntitiesShrink() *string
	SetThirdPartId(v string) *InvoiceRuleAddShrinkRequest
	GetThirdPartId() *string
}

type InvoiceRuleAddShrinkRequest struct {
	// This parameter is required.
	EntitiesShrink *string `json:"entities,omitempty" xml:"entities,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4854821
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s InvoiceRuleAddShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleAddShrinkRequest) GoString() string {
	return s.String()
}

func (s *InvoiceRuleAddShrinkRequest) GetEntitiesShrink() *string {
	return s.EntitiesShrink
}

func (s *InvoiceRuleAddShrinkRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *InvoiceRuleAddShrinkRequest) SetEntitiesShrink(v string) *InvoiceRuleAddShrinkRequest {
	s.EntitiesShrink = &v
	return s
}

func (s *InvoiceRuleAddShrinkRequest) SetThirdPartId(v string) *InvoiceRuleAddShrinkRequest {
	s.ThirdPartId = &v
	return s
}

func (s *InvoiceRuleAddShrinkRequest) Validate() error {
	return dara.Validate(s)
}
