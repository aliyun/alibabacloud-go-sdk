// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLOrder interface {
	dara.Model
	String() string
	GoString() string
	SetCol(v string) *DLOrder
	GetCol() *string
	SetOrder(v int32) *DLOrder
	GetOrder() *int32
}

type DLOrder struct {
	// example:
	//
	// col1
	Col *string `json:"Col,omitempty" xml:"Col,omitempty"`
	// example:
	//
	// 1
	Order *int32 `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s DLOrder) String() string {
	return dara.Prettify(s)
}

func (s DLOrder) GoString() string {
	return s.String()
}

func (s *DLOrder) GetCol() *string {
	return s.Col
}

func (s *DLOrder) GetOrder() *int32 {
	return s.Order
}

func (s *DLOrder) SetCol(v string) *DLOrder {
	s.Col = &v
	return s
}

func (s *DLOrder) SetOrder(v int32) *DLOrder {
	s.Order = &v
	return s
}

func (s *DLOrder) Validate() error {
	return dara.Validate(s)
}
