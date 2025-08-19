// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutConcurrencyInput interface {
	dara.Model
	String() string
	GoString() string
	SetReservedConcurrency(v int64) *PutConcurrencyInput
	GetReservedConcurrency() *int64
}

type PutConcurrencyInput struct {
	// This parameter is required.
	//
	// example:
	//
	// 10
	ReservedConcurrency *int64 `json:"reservedConcurrency,omitempty" xml:"reservedConcurrency,omitempty"`
}

func (s PutConcurrencyInput) String() string {
	return dara.Prettify(s)
}

func (s PutConcurrencyInput) GoString() string {
	return s.String()
}

func (s *PutConcurrencyInput) GetReservedConcurrency() *int64 {
	return s.ReservedConcurrency
}

func (s *PutConcurrencyInput) SetReservedConcurrency(v int64) *PutConcurrencyInput {
	s.ReservedConcurrency = &v
	return s
}

func (s *PutConcurrencyInput) Validate() error {
	return dara.Validate(s)
}
