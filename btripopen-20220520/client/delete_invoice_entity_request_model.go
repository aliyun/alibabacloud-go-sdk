// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInvoiceEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelAll(v bool) *DeleteInvoiceEntityRequest
	GetDelAll() *bool
	SetEntities(v []*DeleteInvoiceEntityRequestEntities) *DeleteInvoiceEntityRequest
	GetEntities() []*DeleteInvoiceEntityRequestEntities
	SetThirdPartId(v string) *DeleteInvoiceEntityRequest
	GetThirdPartId() *string
}

type DeleteInvoiceEntityRequest struct {
	// example:
	//
	// false
	DelAll   *bool                                 `json:"del_all,omitempty" xml:"del_all,omitempty"`
	Entities []*DeleteInvoiceEntityRequestEntities `json:"entities,omitempty" xml:"entities,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 340049
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s DeleteInvoiceEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInvoiceEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteInvoiceEntityRequest) GetDelAll() *bool {
	return s.DelAll
}

func (s *DeleteInvoiceEntityRequest) GetEntities() []*DeleteInvoiceEntityRequestEntities {
	return s.Entities
}

func (s *DeleteInvoiceEntityRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *DeleteInvoiceEntityRequest) SetDelAll(v bool) *DeleteInvoiceEntityRequest {
	s.DelAll = &v
	return s
}

func (s *DeleteInvoiceEntityRequest) SetEntities(v []*DeleteInvoiceEntityRequestEntities) *DeleteInvoiceEntityRequest {
	s.Entities = v
	return s
}

func (s *DeleteInvoiceEntityRequest) SetThirdPartId(v string) *DeleteInvoiceEntityRequest {
	s.ThirdPartId = &v
	return s
}

func (s *DeleteInvoiceEntityRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteInvoiceEntityRequestEntities struct {
	// example:
	//
	// 12345
	EntityId *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	// example:
	//
	// 1
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
}

func (s DeleteInvoiceEntityRequestEntities) String() string {
	return dara.Prettify(s)
}

func (s DeleteInvoiceEntityRequestEntities) GoString() string {
	return s.String()
}

func (s *DeleteInvoiceEntityRequestEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *DeleteInvoiceEntityRequestEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *DeleteInvoiceEntityRequestEntities) SetEntityId(v string) *DeleteInvoiceEntityRequestEntities {
	s.EntityId = &v
	return s
}

func (s *DeleteInvoiceEntityRequestEntities) SetEntityType(v string) *DeleteInvoiceEntityRequestEntities {
	s.EntityType = &v
	return s
}

func (s *DeleteInvoiceEntityRequestEntities) Validate() error {
	return dara.Validate(s)
}
