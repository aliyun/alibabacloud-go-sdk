// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleAddRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntities(v []*InvoiceRuleAddRequestEntities) *InvoiceRuleAddRequest
	GetEntities() []*InvoiceRuleAddRequestEntities
	SetThirdPartId(v string) *InvoiceRuleAddRequest
	GetThirdPartId() *string
}

type InvoiceRuleAddRequest struct {
	// This parameter is required.
	Entities []*InvoiceRuleAddRequestEntities `json:"entities,omitempty" xml:"entities,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 4854821
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s InvoiceRuleAddRequest) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleAddRequest) GoString() string {
	return s.String()
}

func (s *InvoiceRuleAddRequest) GetEntities() []*InvoiceRuleAddRequestEntities {
	return s.Entities
}

func (s *InvoiceRuleAddRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *InvoiceRuleAddRequest) SetEntities(v []*InvoiceRuleAddRequestEntities) *InvoiceRuleAddRequest {
	s.Entities = v
	return s
}

func (s *InvoiceRuleAddRequest) SetThirdPartId(v string) *InvoiceRuleAddRequest {
	s.ThirdPartId = &v
	return s
}

func (s *InvoiceRuleAddRequest) Validate() error {
	return dara.Validate(s)
}

type InvoiceRuleAddRequestEntities struct {
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

func (s InvoiceRuleAddRequestEntities) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleAddRequestEntities) GoString() string {
	return s.String()
}

func (s *InvoiceRuleAddRequestEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *InvoiceRuleAddRequestEntities) GetEntityName() *string {
	return s.EntityName
}

func (s *InvoiceRuleAddRequestEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *InvoiceRuleAddRequestEntities) SetEntityId(v string) *InvoiceRuleAddRequestEntities {
	s.EntityId = &v
	return s
}

func (s *InvoiceRuleAddRequestEntities) SetEntityName(v string) *InvoiceRuleAddRequestEntities {
	s.EntityName = &v
	return s
}

func (s *InvoiceRuleAddRequestEntities) SetEntityType(v string) *InvoiceRuleAddRequestEntities {
	s.EntityType = &v
	return s
}

func (s *InvoiceRuleAddRequestEntities) Validate() error {
	return dara.Validate(s)
}
