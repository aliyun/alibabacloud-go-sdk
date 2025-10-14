// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCancelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderNum(v int64) *ChangeCancelRequest
	GetChangeOrderNum() *int64
}

type ChangeCancelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4988430***950
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
}

func (s ChangeCancelRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeCancelRequest) GoString() string {
	return s.String()
}

func (s *ChangeCancelRequest) GetChangeOrderNum() *int64 {
	return s.ChangeOrderNum
}

func (s *ChangeCancelRequest) SetChangeOrderNum(v int64) *ChangeCancelRequest {
	s.ChangeOrderNum = &v
	return s
}

func (s *ChangeCancelRequest) Validate() error {
	return dara.Validate(s)
}
