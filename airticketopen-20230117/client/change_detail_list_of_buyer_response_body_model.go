// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDetailListOfBuyerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeDetailListOfBuyerResponseBody
	GetRequestId() *string
	SetData(v *ChangeDetailListOfBuyerResponseBodyData) *ChangeDetailListOfBuyerResponseBody
	GetData() *ChangeDetailListOfBuyerResponseBodyData
	SetErrorCode(v string) *ChangeDetailListOfBuyerResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *ChangeDetailListOfBuyerResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *ChangeDetailListOfBuyerResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *ChangeDetailListOfBuyerResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *ChangeDetailListOfBuyerResponseBody
	GetSuccess() *bool
}

type ChangeDetailListOfBuyerResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *ChangeDetailListOfBuyerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ChangeDetailListOfBuyerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfBuyerResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfBuyerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeDetailListOfBuyerResponseBody) GetData() *ChangeDetailListOfBuyerResponseBodyData {
	return s.Data
}

func (s *ChangeDetailListOfBuyerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChangeDetailListOfBuyerResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *ChangeDetailListOfBuyerResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ChangeDetailListOfBuyerResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *ChangeDetailListOfBuyerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeDetailListOfBuyerResponseBody) SetRequestId(v string) *ChangeDetailListOfBuyerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBody) SetData(v *ChangeDetailListOfBuyerResponseBodyData) *ChangeDetailListOfBuyerResponseBody {
	s.Data = v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBody) SetErrorCode(v string) *ChangeDetailListOfBuyerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBody) SetErrorData(v interface{}) *ChangeDetailListOfBuyerResponseBody {
	s.ErrorData = v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBody) SetErrorMsg(v string) *ChangeDetailListOfBuyerResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBody) SetStatus(v int32) *ChangeDetailListOfBuyerResponseBody {
	s.Status = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBody) SetSuccess(v bool) *ChangeDetailListOfBuyerResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeDetailListOfBuyerResponseBodyData struct {
	List       []*ChangeDetailListOfBuyerResponseBodyDataList     `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	Pagination *ChangeDetailListOfBuyerResponseBodyDataPagination `json:"pagination,omitempty" xml:"pagination,omitempty" type:"Struct"`
}

func (s ChangeDetailListOfBuyerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfBuyerResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfBuyerResponseBodyData) GetList() []*ChangeDetailListOfBuyerResponseBodyDataList {
	return s.List
}

func (s *ChangeDetailListOfBuyerResponseBodyData) GetPagination() *ChangeDetailListOfBuyerResponseBodyDataPagination {
	return s.Pagination
}

func (s *ChangeDetailListOfBuyerResponseBodyData) SetList(v []*ChangeDetailListOfBuyerResponseBodyDataList) *ChangeDetailListOfBuyerResponseBodyData {
	s.List = v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyData) SetPagination(v *ChangeDetailListOfBuyerResponseBodyDataPagination) *ChangeDetailListOfBuyerResponseBodyData {
	s.Pagination = v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Pagination != nil {
		if err := s.Pagination.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeDetailListOfBuyerResponseBodyDataList struct {
	// example:
	//
	// 4988430***950
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
	// example:
	//
	// 4988430***971
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// example:
	//
	// 2
	OrderStatus *int32 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// example:
	//
	// 2
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
	// example:
	//
	// 1677415274000
	UtcCreateTime *int64 `json:"utc_create_time,omitempty" xml:"utc_create_time,omitempty"`
}

func (s ChangeDetailListOfBuyerResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfBuyerResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) GetChangeOrderNum() *int64 {
	return s.ChangeOrderNum
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) GetTransactionNo() *string {
	return s.TransactionNo
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) GetUtcCreateTime() *int64 {
	return s.UtcCreateTime
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) SetChangeOrderNum(v int64) *ChangeDetailListOfBuyerResponseBodyDataList {
	s.ChangeOrderNum = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) SetOrderNum(v int64) *ChangeDetailListOfBuyerResponseBodyDataList {
	s.OrderNum = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) SetOrderStatus(v int32) *ChangeDetailListOfBuyerResponseBodyDataList {
	s.OrderStatus = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) SetPayStatus(v int32) *ChangeDetailListOfBuyerResponseBodyDataList {
	s.PayStatus = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) SetTransactionNo(v string) *ChangeDetailListOfBuyerResponseBodyDataList {
	s.TransactionNo = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) SetUtcCreateTime(v int64) *ChangeDetailListOfBuyerResponseBodyDataList {
	s.UtcCreateTime = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailListOfBuyerResponseBodyDataPagination struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

func (s ChangeDetailListOfBuyerResponseBodyDataPagination) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfBuyerResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfBuyerResponseBodyDataPagination) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ChangeDetailListOfBuyerResponseBodyDataPagination) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ChangeDetailListOfBuyerResponseBodyDataPagination) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ChangeDetailListOfBuyerResponseBodyDataPagination) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ChangeDetailListOfBuyerResponseBodyDataPagination) SetCurrentPage(v int32) *ChangeDetailListOfBuyerResponseBodyDataPagination {
	s.CurrentPage = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataPagination) SetPageSize(v int32) *ChangeDetailListOfBuyerResponseBodyDataPagination {
	s.PageSize = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataPagination) SetTotalCount(v int32) *ChangeDetailListOfBuyerResponseBodyDataPagination {
	s.TotalCount = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataPagination) SetTotalPage(v int32) *ChangeDetailListOfBuyerResponseBodyDataPagination {
	s.TotalPage = &v
	return s
}

func (s *ChangeDetailListOfBuyerResponseBodyDataPagination) Validate() error {
	return dara.Validate(s)
}
