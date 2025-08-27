// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInvoiceEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntities(v []*AddInvoiceEntityRequestEntities) *AddInvoiceEntityRequest
	GetEntities() []*AddInvoiceEntityRequestEntities
	SetThirdPartId(v string) *AddInvoiceEntityRequest
	GetThirdPartId() *string
}

type AddInvoiceEntityRequest struct {
	// This parameter is required.
	Entities []*AddInvoiceEntityRequestEntities `json:"entities,omitempty" xml:"entities,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 4854821
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s AddInvoiceEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s AddInvoiceEntityRequest) GoString() string {
	return s.String()
}

func (s *AddInvoiceEntityRequest) GetEntities() []*AddInvoiceEntityRequestEntities {
	return s.Entities
}

func (s *AddInvoiceEntityRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *AddInvoiceEntityRequest) SetEntities(v []*AddInvoiceEntityRequestEntities) *AddInvoiceEntityRequest {
	s.Entities = v
	return s
}

func (s *AddInvoiceEntityRequest) SetThirdPartId(v string) *AddInvoiceEntityRequest {
	s.ThirdPartId = &v
	return s
}

func (s *AddInvoiceEntityRequest) Validate() error {
	return dara.Validate(s)
}

type AddInvoiceEntityRequestEntities struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	EntityId *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	// This parameter is required.
	EntityName *string `json:"entity_name,omitempty" xml:"entity_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
}

func (s AddInvoiceEntityRequestEntities) String() string {
	return dara.Prettify(s)
}

func (s AddInvoiceEntityRequestEntities) GoString() string {
	return s.String()
}

func (s *AddInvoiceEntityRequestEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *AddInvoiceEntityRequestEntities) GetEntityName() *string {
	return s.EntityName
}

func (s *AddInvoiceEntityRequestEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *AddInvoiceEntityRequestEntities) SetEntityId(v string) *AddInvoiceEntityRequestEntities {
	s.EntityId = &v
	return s
}

func (s *AddInvoiceEntityRequestEntities) SetEntityName(v string) *AddInvoiceEntityRequestEntities {
	s.EntityName = &v
	return s
}

func (s *AddInvoiceEntityRequestEntities) SetEntityType(v string) *AddInvoiceEntityRequestEntities {
	s.EntityType = &v
	return s
}

func (s *AddInvoiceEntityRequestEntities) Validate() error {
	return dara.Validate(s)
}
