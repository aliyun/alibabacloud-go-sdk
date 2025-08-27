// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInvoiceEntityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntitiesShrink(v string) *AddInvoiceEntityShrinkRequest
	GetEntitiesShrink() *string
	SetThirdPartId(v string) *AddInvoiceEntityShrinkRequest
	GetThirdPartId() *string
}

type AddInvoiceEntityShrinkRequest struct {
	// This parameter is required.
	EntitiesShrink *string `json:"entities,omitempty" xml:"entities,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4854821
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s AddInvoiceEntityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddInvoiceEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddInvoiceEntityShrinkRequest) GetEntitiesShrink() *string {
	return s.EntitiesShrink
}

func (s *AddInvoiceEntityShrinkRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *AddInvoiceEntityShrinkRequest) SetEntitiesShrink(v string) *AddInvoiceEntityShrinkRequest {
	s.EntitiesShrink = &v
	return s
}

func (s *AddInvoiceEntityShrinkRequest) SetThirdPartId(v string) *AddInvoiceEntityShrinkRequest {
	s.ThirdPartId = &v
	return s
}

func (s *AddInvoiceEntityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
