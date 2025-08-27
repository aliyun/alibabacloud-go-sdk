// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvoiceRuleDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelAll(v bool) *InvoiceRuleDeleteRequest
	GetDelAll() *bool
	SetEntities(v []*InvoiceRuleDeleteRequestEntities) *InvoiceRuleDeleteRequest
	GetEntities() []*InvoiceRuleDeleteRequestEntities
	SetThirdPartId(v string) *InvoiceRuleDeleteRequest
	GetThirdPartId() *string
}

type InvoiceRuleDeleteRequest struct {
	// example:
	//
	// false
	DelAll   *bool                               `json:"del_all,omitempty" xml:"del_all,omitempty"`
	Entities []*InvoiceRuleDeleteRequestEntities `json:"entities,omitempty" xml:"entities,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 340049
	ThirdPartId *string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
}

func (s InvoiceRuleDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleDeleteRequest) GoString() string {
	return s.String()
}

func (s *InvoiceRuleDeleteRequest) GetDelAll() *bool {
	return s.DelAll
}

func (s *InvoiceRuleDeleteRequest) GetEntities() []*InvoiceRuleDeleteRequestEntities {
	return s.Entities
}

func (s *InvoiceRuleDeleteRequest) GetThirdPartId() *string {
	return s.ThirdPartId
}

func (s *InvoiceRuleDeleteRequest) SetDelAll(v bool) *InvoiceRuleDeleteRequest {
	s.DelAll = &v
	return s
}

func (s *InvoiceRuleDeleteRequest) SetEntities(v []*InvoiceRuleDeleteRequestEntities) *InvoiceRuleDeleteRequest {
	s.Entities = v
	return s
}

func (s *InvoiceRuleDeleteRequest) SetThirdPartId(v string) *InvoiceRuleDeleteRequest {
	s.ThirdPartId = &v
	return s
}

func (s *InvoiceRuleDeleteRequest) Validate() error {
	return dara.Validate(s)
}

type InvoiceRuleDeleteRequestEntities struct {
	// example:
	//
	// 12345
	EntityId *string `json:"entity_id,omitempty" xml:"entity_id,omitempty"`
	// example:
	//
	// 1
	EntityType *string `json:"entity_type,omitempty" xml:"entity_type,omitempty"`
}

func (s InvoiceRuleDeleteRequestEntities) String() string {
	return dara.Prettify(s)
}

func (s InvoiceRuleDeleteRequestEntities) GoString() string {
	return s.String()
}

func (s *InvoiceRuleDeleteRequestEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *InvoiceRuleDeleteRequestEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *InvoiceRuleDeleteRequestEntities) SetEntityId(v string) *InvoiceRuleDeleteRequestEntities {
	s.EntityId = &v
	return s
}

func (s *InvoiceRuleDeleteRequestEntities) SetEntityType(v string) *InvoiceRuleDeleteRequestEntities {
	s.EntityType = &v
	return s
}

func (s *InvoiceRuleDeleteRequestEntities) Validate() error {
	return dara.Validate(s)
}
