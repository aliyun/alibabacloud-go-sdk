// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundDetailListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RefundDetailListResponseBody
	GetRequestId() *string
	SetData(v *RefundDetailListResponseBodyData) *RefundDetailListResponseBody
	GetData() *RefundDetailListResponseBodyData
	SetErrorCode(v string) *RefundDetailListResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *RefundDetailListResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *RefundDetailListResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *RefundDetailListResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *RefundDetailListResponseBody
	GetSuccess() *bool
}

type RefundDetailListResponseBody struct {
	// Request RequestId
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Properly processed return data
	Data *RefundDetailListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// Data carried in error handling
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// Error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// HTTP request successful, status value is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// Whether the operation was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s RefundDetailListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailListResponseBody) GoString() string {
	return s.String()
}

func (s *RefundDetailListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefundDetailListResponseBody) GetData() *RefundDetailListResponseBodyData {
	return s.Data
}

func (s *RefundDetailListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RefundDetailListResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *RefundDetailListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *RefundDetailListResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *RefundDetailListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RefundDetailListResponseBody) SetRequestId(v string) *RefundDetailListResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundDetailListResponseBody) SetData(v *RefundDetailListResponseBodyData) *RefundDetailListResponseBody {
	s.Data = v
	return s
}

func (s *RefundDetailListResponseBody) SetErrorCode(v string) *RefundDetailListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RefundDetailListResponseBody) SetErrorData(v interface{}) *RefundDetailListResponseBody {
	s.ErrorData = v
	return s
}

func (s *RefundDetailListResponseBody) SetErrorMsg(v string) *RefundDetailListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RefundDetailListResponseBody) SetStatus(v int32) *RefundDetailListResponseBody {
	s.Status = &v
	return s
}

func (s *RefundDetailListResponseBody) SetSuccess(v bool) *RefundDetailListResponseBody {
	s.Success = &v
	return s
}

func (s *RefundDetailListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefundDetailListResponseBodyData struct {
	// Data list
	List []*RefundDetailListResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// Pagination information
	Pagination *RefundDetailListResponseBodyDataPagination `json:"pagination,omitempty" xml:"pagination,omitempty" type:"Struct"`
}

func (s RefundDetailListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailListResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefundDetailListResponseBodyData) GetList() []*RefundDetailListResponseBodyDataList {
	return s.List
}

func (s *RefundDetailListResponseBodyData) GetPagination() *RefundDetailListResponseBodyDataPagination {
	return s.Pagination
}

func (s *RefundDetailListResponseBodyData) SetList(v []*RefundDetailListResponseBodyDataList) *RefundDetailListResponseBodyData {
	s.List = v
	return s
}

func (s *RefundDetailListResponseBodyData) SetPagination(v *RefundDetailListResponseBodyDataPagination) *RefundDetailListResponseBodyData {
	s.Pagination = v
	return s
}

func (s *RefundDetailListResponseBodyData) Validate() error {
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

type RefundDetailListResponseBodyDataList struct {
	// Whether it is a supplementary refund
	//
	// example:
	//
	// true
	IsMultiRefund *bool `json:"is_multi_refund,omitempty" xml:"is_multi_refund,omitempty"`
	// Order number （ Ticketing Order Number）
	//
	// example:
	//
	// 49884*****2345
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// Refund order number
	//
	// example:
	//
	// 49884*****950
	RefundOrderNum *int64 `json:"refund_order_num,omitempty" xml:"refund_order_num,omitempty"`
	// Refund order status: 0: Refund application; 1: Refund in progress; 2: Refund failed; 3: Refund succeeded
	//
	// example:
	//
	// 1
	RefundOrderStatus *int32 `json:"refund_order_status,omitempty" xml:"refund_order_status,omitempty"`
	// The original refund order number associated with this supplementary refund. Only present for supplementary refunds, indicating the ID of the original refund order.
	//
	// example:
	//
	// 49884*****2387
	RelatedRefundOrderNum *string `json:"related_refund_order_num,omitempty" xml:"related_refund_order_num,omitempty"`
	// Transaction serial number
	//
	// example:
	//
	// 49884**tde-95za
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
	// Creation time, UTC timestamp
	//
	// example:
	//
	// 1677229002000
	UtcCreateTime *int64 `json:"utc_create_time,omitempty" xml:"utc_create_time,omitempty"`
}

func (s RefundDetailListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *RefundDetailListResponseBodyDataList) GetIsMultiRefund() *bool {
	return s.IsMultiRefund
}

func (s *RefundDetailListResponseBodyDataList) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *RefundDetailListResponseBodyDataList) GetRefundOrderNum() *int64 {
	return s.RefundOrderNum
}

func (s *RefundDetailListResponseBodyDataList) GetRefundOrderStatus() *int32 {
	return s.RefundOrderStatus
}

func (s *RefundDetailListResponseBodyDataList) GetRelatedRefundOrderNum() *string {
	return s.RelatedRefundOrderNum
}

func (s *RefundDetailListResponseBodyDataList) GetTransactionNo() *string {
	return s.TransactionNo
}

func (s *RefundDetailListResponseBodyDataList) GetUtcCreateTime() *int64 {
	return s.UtcCreateTime
}

func (s *RefundDetailListResponseBodyDataList) SetIsMultiRefund(v bool) *RefundDetailListResponseBodyDataList {
	s.IsMultiRefund = &v
	return s
}

func (s *RefundDetailListResponseBodyDataList) SetOrderNum(v int64) *RefundDetailListResponseBodyDataList {
	s.OrderNum = &v
	return s
}

func (s *RefundDetailListResponseBodyDataList) SetRefundOrderNum(v int64) *RefundDetailListResponseBodyDataList {
	s.RefundOrderNum = &v
	return s
}

func (s *RefundDetailListResponseBodyDataList) SetRefundOrderStatus(v int32) *RefundDetailListResponseBodyDataList {
	s.RefundOrderStatus = &v
	return s
}

func (s *RefundDetailListResponseBodyDataList) SetRelatedRefundOrderNum(v string) *RefundDetailListResponseBodyDataList {
	s.RelatedRefundOrderNum = &v
	return s
}

func (s *RefundDetailListResponseBodyDataList) SetTransactionNo(v string) *RefundDetailListResponseBodyDataList {
	s.TransactionNo = &v
	return s
}

func (s *RefundDetailListResponseBodyDataList) SetUtcCreateTime(v int64) *RefundDetailListResponseBodyDataList {
	s.UtcCreateTime = &v
	return s
}

func (s *RefundDetailListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type RefundDetailListResponseBodyDataPagination struct {
	// Current page number
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// Number of items per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// Total count
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// Total pages
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

func (s RefundDetailListResponseBodyDataPagination) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailListResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *RefundDetailListResponseBodyDataPagination) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *RefundDetailListResponseBodyDataPagination) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RefundDetailListResponseBodyDataPagination) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *RefundDetailListResponseBodyDataPagination) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *RefundDetailListResponseBodyDataPagination) SetCurrentPage(v int32) *RefundDetailListResponseBodyDataPagination {
	s.CurrentPage = &v
	return s
}

func (s *RefundDetailListResponseBodyDataPagination) SetPageSize(v int32) *RefundDetailListResponseBodyDataPagination {
	s.PageSize = &v
	return s
}

func (s *RefundDetailListResponseBodyDataPagination) SetTotalCount(v int32) *RefundDetailListResponseBodyDataPagination {
	s.TotalCount = &v
	return s
}

func (s *RefundDetailListResponseBodyDataPagination) SetTotalPage(v int32) *RefundDetailListResponseBodyDataPagination {
	s.TotalPage = &v
	return s
}

func (s *RefundDetailListResponseBodyDataPagination) Validate() error {
	return dara.Validate(s)
}
