// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInvoiceEntityShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelAll(v bool) *DeleteInvoiceEntityShrinkRequest
	GetDelAll() *bool
	SetEntitiesShrink(v string) *DeleteInvoiceEntityShrinkRequest
	GetEntitiesShrink() *string
	SetThirdPartId(v string) *DeleteInvoiceEntityShrinkRequest
	GetThirdPartId() *string
}

type DeleteInvoiceEntityShrinkRequest struct {
	// Specifies whether to delete all applicable personnel. If del_all is set to true, all entities under the invoice header are deleted, and the entity list parameter is not validated.
	//
	// example:
	//
	// false
	DelAll *bool `json:"del_all,omitempty" xml:"del_all,omitempty"`
	// The entity list. This parameter is required when del_all is set to false or null.
	EntitiesShrink *string `json:"entities,omitempty" xml:"entities,omitempty"`
	// The third-party invoice ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 340049
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s DeleteInvoiceEntityShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInvoiceEntityShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteInvoiceEntityShrinkRequest) GetDelAll() *bool {
	return s.DelAll
}

func (s *DeleteInvoiceEntityShrinkRequest) GetEntitiesShrink() *string {
	return s.EntitiesShrink
}

func (s *DeleteInvoiceEntityShrinkRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *DeleteInvoiceEntityShrinkRequest) SetDelAll(v bool) *DeleteInvoiceEntityShrinkRequest {
	s.DelAll = &v
	return s
}

func (s *DeleteInvoiceEntityShrinkRequest) SetEntitiesShrink(v string) *DeleteInvoiceEntityShrinkRequest {
	s.EntitiesShrink = &v
	return s
}

func (s *DeleteInvoiceEntityShrinkRequest) SetThirdPartId(v string) *DeleteInvoiceEntityShrinkRequest {
	s.ThirdPartId = &v
	return s
}

func (s *DeleteInvoiceEntityShrinkRequest) Validate() error {
	return dara.Validate(s)
}
