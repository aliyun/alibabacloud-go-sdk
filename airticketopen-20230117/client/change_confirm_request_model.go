// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeConfirmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderNum(v int64) *ChangeConfirmRequest
	GetChangeOrderNum() *int64
}

type ChangeConfirmRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4988430***950
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
}

func (s ChangeConfirmRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeConfirmRequest) GoString() string {
	return s.String()
}

func (s *ChangeConfirmRequest) GetChangeOrderNum() *int64 {
	return s.ChangeOrderNum
}

func (s *ChangeConfirmRequest) SetChangeOrderNum(v int64) *ChangeConfirmRequest {
	s.ChangeOrderNum = &v
	return s
}

func (s *ChangeConfirmRequest) Validate() error {
	return dara.Validate(s)
}
