// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderNum(v int64) *ChangeDetailRequest
	GetChangeOrderNum() *int64
}

type ChangeDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4988430***950
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
}

func (s ChangeDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailRequest) GoString() string {
	return s.String()
}

func (s *ChangeDetailRequest) GetChangeOrderNum() *int64 {
	return s.ChangeOrderNum
}

func (s *ChangeDetailRequest) SetChangeOrderNum(v int64) *ChangeDetailRequest {
	s.ChangeOrderNum = &v
	return s
}

func (s *ChangeDetailRequest) Validate() error {
	return dara.Validate(s)
}
