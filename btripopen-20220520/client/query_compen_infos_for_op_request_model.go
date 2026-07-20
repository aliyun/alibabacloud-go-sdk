// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCompenInfosForOpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v int32) *QueryCompenInfosForOpRequest
	GetCategory() *int32
	SetCompenId(v string) *QueryCompenInfosForOpRequest
	GetCompenId() *string
	SetOrderId(v string) *QueryCompenInfosForOpRequest
	GetOrderId() *string
}

type QueryCompenInfosForOpRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	Category *int32 `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 82508250249123456
	CompenId *string `json:"compen_id,omitempty" xml:"compen_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1002086203277812345
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s QueryCompenInfosForOpRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCompenInfosForOpRequest) GoString() string {
	return s.String()
}

func (s *QueryCompenInfosForOpRequest) GetCategory() *int32 {
	return s.Category
}

func (s *QueryCompenInfosForOpRequest) GetCompenId() *string {
	return s.CompenId
}

func (s *QueryCompenInfosForOpRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryCompenInfosForOpRequest) SetCategory(v int32) *QueryCompenInfosForOpRequest {
	s.Category = &v
	return s
}

func (s *QueryCompenInfosForOpRequest) SetCompenId(v string) *QueryCompenInfosForOpRequest {
	s.CompenId = &v
	return s
}

func (s *QueryCompenInfosForOpRequest) SetOrderId(v string) *QueryCompenInfosForOpRequest {
	s.OrderId = &v
	return s
}

func (s *QueryCompenInfosForOpRequest) Validate() error {
	return dara.Validate(s)
}
