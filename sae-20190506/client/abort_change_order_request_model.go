// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortChangeOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *AbortChangeOrderRequest
	GetChangeOrderId() *string
	SetRollback(v bool) *AbortChangeOrderRequest
	GetRollback() *bool
}

type AbortChangeOrderRequest struct {
	// The ID of the change order.
	//
	// This parameter is required.
	//
	// example:
	//
	// be2e1c76-682b-4897-98d3-1d8d6478****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	Rollback      *bool   `json:"Rollback,omitempty" xml:"Rollback,omitempty"`
}

func (s AbortChangeOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s AbortChangeOrderRequest) GoString() string {
	return s.String()
}

func (s *AbortChangeOrderRequest) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *AbortChangeOrderRequest) GetRollback() *bool {
	return s.Rollback
}

func (s *AbortChangeOrderRequest) SetChangeOrderId(v string) *AbortChangeOrderRequest {
	s.ChangeOrderId = &v
	return s
}

func (s *AbortChangeOrderRequest) SetRollback(v bool) *AbortChangeOrderRequest {
	s.Rollback = &v
	return s
}

func (s *AbortChangeOrderRequest) Validate() error {
	return dara.Validate(s)
}
