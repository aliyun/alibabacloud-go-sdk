// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountFlowListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AccountFlowListResponseBody
	GetRequestId() *string
	SetData(v *AccountFlowListResponseBodyData) *AccountFlowListResponseBody
	GetData() *AccountFlowListResponseBodyData
	SetErrorCode(v string) *AccountFlowListResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *AccountFlowListResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *AccountFlowListResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *AccountFlowListResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *AccountFlowListResponseBody
	GetSuccess() *bool
}

type AccountFlowListResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Data      *AccountFlowListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s AccountFlowListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AccountFlowListResponseBody) GoString() string {
	return s.String()
}

func (s *AccountFlowListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AccountFlowListResponseBody) GetData() *AccountFlowListResponseBodyData {
	return s.Data
}

func (s *AccountFlowListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AccountFlowListResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *AccountFlowListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *AccountFlowListResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *AccountFlowListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AccountFlowListResponseBody) SetRequestId(v string) *AccountFlowListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccountFlowListResponseBody) SetData(v *AccountFlowListResponseBodyData) *AccountFlowListResponseBody {
	s.Data = v
	return s
}

func (s *AccountFlowListResponseBody) SetErrorCode(v string) *AccountFlowListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AccountFlowListResponseBody) SetErrorData(v interface{}) *AccountFlowListResponseBody {
	s.ErrorData = v
	return s
}

func (s *AccountFlowListResponseBody) SetErrorMsg(v string) *AccountFlowListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *AccountFlowListResponseBody) SetStatus(v int32) *AccountFlowListResponseBody {
	s.Status = &v
	return s
}

func (s *AccountFlowListResponseBody) SetSuccess(v bool) *AccountFlowListResponseBody {
	s.Success = &v
	return s
}

func (s *AccountFlowListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AccountFlowListResponseBodyData struct {
	List       []*AccountFlowListResponseBodyDataList     `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	Pagination *AccountFlowListResponseBodyDataPagination `json:"pagination,omitempty" xml:"pagination,omitempty" type:"Struct"`
}

func (s AccountFlowListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AccountFlowListResponseBodyData) GoString() string {
	return s.String()
}

func (s *AccountFlowListResponseBodyData) GetList() []*AccountFlowListResponseBodyDataList {
	return s.List
}

func (s *AccountFlowListResponseBodyData) GetPagination() *AccountFlowListResponseBodyDataPagination {
	return s.Pagination
}

func (s *AccountFlowListResponseBodyData) SetList(v []*AccountFlowListResponseBodyDataList) *AccountFlowListResponseBodyData {
	s.List = v
	return s
}

func (s *AccountFlowListResponseBodyData) SetPagination(v *AccountFlowListResponseBodyDataPagination) *AccountFlowListResponseBodyData {
	s.Pagination = v
	return s
}

func (s *AccountFlowListResponseBodyData) Validate() error {
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

type AccountFlowListResponseBodyDataList struct {
	// example:
	//
	// 1000
	AfterAvailableAmount *float64 `json:"after_available_amount,omitempty" xml:"after_available_amount,omitempty"`
	// example:
	//
	// 1950.5
	BeforeAvailableAmount *float64 `json:"before_available_amount,omitempty" xml:"before_available_amount,omitempty"`
	// example:
	//
	// 49880***971
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
	// example:
	//
	// 1627239841225842666
	FlowId *int64 `json:"flow_id,omitempty" xml:"flow_id,omitempty"`
	// example:
	//
	// 1676799185000
	GmtCreate *int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 1676966530000
	GmtModified *int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// example:
	//
	// 950.5
	OpAmount *float64 `json:"op_amount,omitempty" xml:"op_amount,omitempty"`
	// example:
	//
	// 2
	OpType *int32 `json:"op_type,omitempty" xml:"op_type,omitempty"`
	// example:
	//
	// 4988430***971
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// example:
	//
	// 1
	OrderType *int32 `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// example:
	//
	// 4988430***971
	OutOrderNum *string `json:"out_order_num,omitempty" xml:"out_order_num,omitempty"`
	// example:
	//
	// 48430***971
	RefundOrderNum *int64 `json:"refund_order_num,omitempty" xml:"refund_order_num,omitempty"`
}

func (s AccountFlowListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s AccountFlowListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *AccountFlowListResponseBodyDataList) GetAfterAvailableAmount() *float64 {
	return s.AfterAvailableAmount
}

func (s *AccountFlowListResponseBodyDataList) GetBeforeAvailableAmount() *float64 {
	return s.BeforeAvailableAmount
}

func (s *AccountFlowListResponseBodyDataList) GetChangeOrderNum() *int64 {
	return s.ChangeOrderNum
}

func (s *AccountFlowListResponseBodyDataList) GetFlowId() *int64 {
	return s.FlowId
}

func (s *AccountFlowListResponseBodyDataList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *AccountFlowListResponseBodyDataList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *AccountFlowListResponseBodyDataList) GetOpAmount() *float64 {
	return s.OpAmount
}

func (s *AccountFlowListResponseBodyDataList) GetOpType() *int32 {
	return s.OpType
}

func (s *AccountFlowListResponseBodyDataList) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *AccountFlowListResponseBodyDataList) GetOrderType() *int32 {
	return s.OrderType
}

func (s *AccountFlowListResponseBodyDataList) GetOutOrderNum() *string {
	return s.OutOrderNum
}

func (s *AccountFlowListResponseBodyDataList) GetRefundOrderNum() *int64 {
	return s.RefundOrderNum
}

func (s *AccountFlowListResponseBodyDataList) SetAfterAvailableAmount(v float64) *AccountFlowListResponseBodyDataList {
	s.AfterAvailableAmount = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetBeforeAvailableAmount(v float64) *AccountFlowListResponseBodyDataList {
	s.BeforeAvailableAmount = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetChangeOrderNum(v int64) *AccountFlowListResponseBodyDataList {
	s.ChangeOrderNum = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetFlowId(v int64) *AccountFlowListResponseBodyDataList {
	s.FlowId = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetGmtCreate(v int64) *AccountFlowListResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetGmtModified(v int64) *AccountFlowListResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetOpAmount(v float64) *AccountFlowListResponseBodyDataList {
	s.OpAmount = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetOpType(v int32) *AccountFlowListResponseBodyDataList {
	s.OpType = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetOrderNum(v int64) *AccountFlowListResponseBodyDataList {
	s.OrderNum = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetOrderType(v int32) *AccountFlowListResponseBodyDataList {
	s.OrderType = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetOutOrderNum(v string) *AccountFlowListResponseBodyDataList {
	s.OutOrderNum = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) SetRefundOrderNum(v int64) *AccountFlowListResponseBodyDataList {
	s.RefundOrderNum = &v
	return s
}

func (s *AccountFlowListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

type AccountFlowListResponseBodyDataPagination struct {
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

func (s AccountFlowListResponseBodyDataPagination) String() string {
	return dara.Prettify(s)
}

func (s AccountFlowListResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *AccountFlowListResponseBodyDataPagination) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *AccountFlowListResponseBodyDataPagination) GetPageSize() *int32 {
	return s.PageSize
}

func (s *AccountFlowListResponseBodyDataPagination) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *AccountFlowListResponseBodyDataPagination) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *AccountFlowListResponseBodyDataPagination) SetCurrentPage(v int32) *AccountFlowListResponseBodyDataPagination {
	s.CurrentPage = &v
	return s
}

func (s *AccountFlowListResponseBodyDataPagination) SetPageSize(v int32) *AccountFlowListResponseBodyDataPagination {
	s.PageSize = &v
	return s
}

func (s *AccountFlowListResponseBodyDataPagination) SetTotalCount(v int32) *AccountFlowListResponseBodyDataPagination {
	s.TotalCount = &v
	return s
}

func (s *AccountFlowListResponseBodyDataPagination) SetTotalPage(v int32) *AccountFlowListResponseBodyDataPagination {
	s.TotalPage = &v
	return s
}

func (s *AccountFlowListResponseBodyDataPagination) Validate() error {
	return dara.Validate(s)
}
