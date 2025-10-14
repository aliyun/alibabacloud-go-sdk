// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderNum(v int64) *CancelRequest
	GetOrderNum() *int64
}

type CancelRequest struct {
	// order number
	//
	// This parameter is required.
	//
	// example:
	//
	// 496***2617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s CancelRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelRequest) GoString() string {
	return s.String()
}

func (s *CancelRequest) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *CancelRequest) SetOrderNum(v int64) *CancelRequest {
	s.OrderNum = &v
	return s
}

func (s *CancelRequest) Validate() error {
	return dara.Validate(s)
}
