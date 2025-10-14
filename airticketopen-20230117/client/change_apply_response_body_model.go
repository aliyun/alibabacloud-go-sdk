// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeApplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeApplyResponseBody
	GetRequestId() *string
	SetData(v *ChangeApplyResponseBodyData) *ChangeApplyResponseBody
	GetData() *ChangeApplyResponseBodyData
	SetErrorCode(v string) *ChangeApplyResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *ChangeApplyResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *ChangeApplyResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *ChangeApplyResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *ChangeApplyResponseBody
	GetSuccess() *bool
}

type ChangeApplyResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ChangeApplyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ChangeApplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeApplyResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeApplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeApplyResponseBody) GetData() *ChangeApplyResponseBodyData {
	return s.Data
}

func (s *ChangeApplyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChangeApplyResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *ChangeApplyResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ChangeApplyResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *ChangeApplyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeApplyResponseBody) SetRequestId(v string) *ChangeApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeApplyResponseBody) SetData(v *ChangeApplyResponseBodyData) *ChangeApplyResponseBody {
	s.Data = v
	return s
}

func (s *ChangeApplyResponseBody) SetErrorCode(v string) *ChangeApplyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeApplyResponseBody) SetErrorData(v interface{}) *ChangeApplyResponseBody {
	s.ErrorData = v
	return s
}

func (s *ChangeApplyResponseBody) SetErrorMsg(v string) *ChangeApplyResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChangeApplyResponseBody) SetStatus(v int32) *ChangeApplyResponseBody {
	s.Status = &v
	return s
}

func (s *ChangeApplyResponseBody) SetSuccess(v bool) *ChangeApplyResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeApplyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeApplyResponseBodyData struct {
	ChangeOrders []*ChangeApplyResponseBodyDataChangeOrders `json:"change_orders,omitempty" xml:"change_orders,omitempty" type:"Repeated"`
	// example:
	//
	// 4988430***950
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
}

func (s ChangeApplyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChangeApplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeApplyResponseBodyData) GetChangeOrders() []*ChangeApplyResponseBodyDataChangeOrders {
	return s.ChangeOrders
}

func (s *ChangeApplyResponseBodyData) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *ChangeApplyResponseBodyData) SetChangeOrders(v []*ChangeApplyResponseBodyDataChangeOrders) *ChangeApplyResponseBodyData {
	s.ChangeOrders = v
	return s
}

func (s *ChangeApplyResponseBodyData) SetOrderNum(v int64) *ChangeApplyResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *ChangeApplyResponseBodyData) Validate() error {
	if s.ChangeOrders != nil {
		for _, item := range s.ChangeOrders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChangeApplyResponseBodyDataChangeOrders struct {
	// example:
	//
	// 49884*****950
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
	// example:
	//
	// 0
	ChangeOrderStatus *int32 `json:"change_order_status,omitempty" xml:"change_order_status,omitempty"`
	// example:
	//
	// desc reason
	FailReason *string                                              `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	Passengers []*ChangeApplyResponseBodyDataChangeOrdersPassengers `json:"passengers,omitempty" xml:"passengers,omitempty" type:"Repeated"`
}

func (s ChangeApplyResponseBodyDataChangeOrders) String() string {
	return dara.Prettify(s)
}

func (s ChangeApplyResponseBodyDataChangeOrders) GoString() string {
	return s.String()
}

func (s *ChangeApplyResponseBodyDataChangeOrders) GetChangeOrderNum() *int64 {
	return s.ChangeOrderNum
}

func (s *ChangeApplyResponseBodyDataChangeOrders) GetChangeOrderStatus() *int32 {
	return s.ChangeOrderStatus
}

func (s *ChangeApplyResponseBodyDataChangeOrders) GetFailReason() *string {
	return s.FailReason
}

func (s *ChangeApplyResponseBodyDataChangeOrders) GetPassengers() []*ChangeApplyResponseBodyDataChangeOrdersPassengers {
	return s.Passengers
}

func (s *ChangeApplyResponseBodyDataChangeOrders) SetChangeOrderNum(v int64) *ChangeApplyResponseBodyDataChangeOrders {
	s.ChangeOrderNum = &v
	return s
}

func (s *ChangeApplyResponseBodyDataChangeOrders) SetChangeOrderStatus(v int32) *ChangeApplyResponseBodyDataChangeOrders {
	s.ChangeOrderStatus = &v
	return s
}

func (s *ChangeApplyResponseBodyDataChangeOrders) SetFailReason(v string) *ChangeApplyResponseBodyDataChangeOrders {
	s.FailReason = &v
	return s
}

func (s *ChangeApplyResponseBodyDataChangeOrders) SetPassengers(v []*ChangeApplyResponseBodyDataChangeOrdersPassengers) *ChangeApplyResponseBodyDataChangeOrders {
	s.Passengers = v
	return s
}

func (s *ChangeApplyResponseBodyDataChangeOrders) Validate() error {
	if s.Passengers != nil {
		for _, item := range s.Passengers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChangeApplyResponseBodyDataChangeOrdersPassengers struct {
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s ChangeApplyResponseBodyDataChangeOrdersPassengers) String() string {
	return dara.Prettify(s)
}

func (s ChangeApplyResponseBodyDataChangeOrdersPassengers) GoString() string {
	return s.String()
}

func (s *ChangeApplyResponseBodyDataChangeOrdersPassengers) GetDocument() *string {
	return s.Document
}

func (s *ChangeApplyResponseBodyDataChangeOrdersPassengers) GetFirstName() *string {
	return s.FirstName
}

func (s *ChangeApplyResponseBodyDataChangeOrdersPassengers) GetLastName() *string {
	return s.LastName
}

func (s *ChangeApplyResponseBodyDataChangeOrdersPassengers) SetDocument(v string) *ChangeApplyResponseBodyDataChangeOrdersPassengers {
	s.Document = &v
	return s
}

func (s *ChangeApplyResponseBodyDataChangeOrdersPassengers) SetFirstName(v string) *ChangeApplyResponseBodyDataChangeOrdersPassengers {
	s.FirstName = &v
	return s
}

func (s *ChangeApplyResponseBodyDataChangeOrdersPassengers) SetLastName(v string) *ChangeApplyResponseBodyDataChangeOrdersPassengers {
	s.LastName = &v
	return s
}

func (s *ChangeApplyResponseBodyDataChangeOrdersPassengers) Validate() error {
	return dara.Validate(s)
}
