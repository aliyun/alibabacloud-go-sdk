// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorSyncPayStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCooperatorOrderId(v string) *CooperatorSyncPayStatusRequest
	GetCooperatorOrderId() *string
	SetCooperatorPayNo(v string) *CooperatorSyncPayStatusRequest
	GetCooperatorPayNo() *string
	SetOrderId(v string) *CooperatorSyncPayStatusRequest
	GetOrderId() *string
	SetPayStatus(v string) *CooperatorSyncPayStatusRequest
	GetPayStatus() *string
	SetPayTime(v int64) *CooperatorSyncPayStatusRequest
	GetPayTime() *int64
}

type CooperatorSyncPayStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	CooperatorOrderId *string `json:"cooperator_order_id,omitempty" xml:"cooperator_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	CooperatorPayNo *string `json:"cooperator_pay_no,omitempty" xml:"cooperator_pay_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1017018197205925373
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	PayStatus *string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1726022215000
	PayTime *int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
}

func (s CooperatorSyncPayStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s CooperatorSyncPayStatusRequest) GoString() string {
	return s.String()
}

func (s *CooperatorSyncPayStatusRequest) GetCooperatorOrderId() *string {
	return s.CooperatorOrderId
}

func (s *CooperatorSyncPayStatusRequest) GetCooperatorPayNo() *string {
	return s.CooperatorPayNo
}

func (s *CooperatorSyncPayStatusRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *CooperatorSyncPayStatusRequest) GetPayStatus() *string {
	return s.PayStatus
}

func (s *CooperatorSyncPayStatusRequest) GetPayTime() *int64 {
	return s.PayTime
}

func (s *CooperatorSyncPayStatusRequest) SetCooperatorOrderId(v string) *CooperatorSyncPayStatusRequest {
	s.CooperatorOrderId = &v
	return s
}

func (s *CooperatorSyncPayStatusRequest) SetCooperatorPayNo(v string) *CooperatorSyncPayStatusRequest {
	s.CooperatorPayNo = &v
	return s
}

func (s *CooperatorSyncPayStatusRequest) SetOrderId(v string) *CooperatorSyncPayStatusRequest {
	s.OrderId = &v
	return s
}

func (s *CooperatorSyncPayStatusRequest) SetPayStatus(v string) *CooperatorSyncPayStatusRequest {
	s.PayStatus = &v
	return s
}

func (s *CooperatorSyncPayStatusRequest) SetPayTime(v int64) *CooperatorSyncPayStatusRequest {
	s.PayTime = &v
	return s
}

func (s *CooperatorSyncPayStatusRequest) Validate() error {
	return dara.Validate(s)
}
