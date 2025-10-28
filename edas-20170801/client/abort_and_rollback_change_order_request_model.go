// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortAndRollbackChangeOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *AbortAndRollbackChangeOrderRequest
	GetChangeOrderId() *string
}

type AbortAndRollbackChangeOrderRequest struct {
	// The ID of the change process.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6b55106c-c7f2-40f7-a1d4-092870*****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s AbortAndRollbackChangeOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s AbortAndRollbackChangeOrderRequest) GoString() string {
	return s.String()
}

func (s *AbortAndRollbackChangeOrderRequest) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *AbortAndRollbackChangeOrderRequest) SetChangeOrderId(v string) *AbortAndRollbackChangeOrderRequest {
	s.ChangeOrderId = &v
	return s
}

func (s *AbortAndRollbackChangeOrderRequest) Validate() error {
	return dara.Validate(s)
}
