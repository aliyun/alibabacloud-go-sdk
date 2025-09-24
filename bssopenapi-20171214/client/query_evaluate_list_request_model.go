// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEvaluateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillCycle(v string) *QueryEvaluateListRequest
	GetBillCycle() *string
	SetBizTypeList(v []*string) *QueryEvaluateListRequest
	GetBizTypeList() []*string
	SetEndAmount(v int64) *QueryEvaluateListRequest
	GetEndAmount() *int64
	SetEndBizTime(v string) *QueryEvaluateListRequest
	GetEndBizTime() *string
	SetEndSearchTime(v string) *QueryEvaluateListRequest
	GetEndSearchTime() *string
	SetOutBizId(v string) *QueryEvaluateListRequest
	GetOutBizId() *string
	SetOwnerId(v int64) *QueryEvaluateListRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *QueryEvaluateListRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryEvaluateListRequest
	GetPageSize() *int32
	SetSortType(v int32) *QueryEvaluateListRequest
	GetSortType() *int32
	SetStartAmount(v int64) *QueryEvaluateListRequest
	GetStartAmount() *int64
	SetStartBizTime(v string) *QueryEvaluateListRequest
	GetStartBizTime() *string
	SetStartSearchTime(v string) *QueryEvaluateListRequest
	GetStartSearchTime() *string
	SetType(v int32) *QueryEvaluateListRequest
	GetType() *int32
}

type QueryEvaluateListRequest struct {
	// The billing cycle.
	//
	// example:
	//
	// 202003
	BillCycle *string `json:"BillCycle,omitempty" xml:"BillCycle,omitempty"`
	// The market types in invoices.
	//
	// >  By default, this parameter is left empty. If this parameter is left empty, all market types are queried.
	//
	// example:
	//
	// ALIYUN
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// The maximum amount to be queried.
	//
	// example:
	//
	// 1000
	EndAmount *int64 `json:"EndAmount,omitempty" xml:"EndAmount,omitempty"`
	// The latest time when an order is paid Specify the time in the yyyy-mm-dd hh:mm:ss format.
	//
	// example:
	//
	// 2020-02-02 15:00:00
	EndBizTime *string `json:"EndBizTime,omitempty" xml:"EndBizTime,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 2020-03-02 12:00:00
	EndSearchTime *string `json:"EndSearchTime,omitempty" xml:"EndSearchTime,omitempty"`
	// The ID of the external order.
	//
	// example:
	//
	// 2387432832696
	OutBizId *string `json:"OutBizId,omitempty" xml:"OutBizId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the sort. Valid values:
	//
	// 	- 1: Sort invoices by ID in descending order.
	//
	// 	- 2: Sort invoices by invoice type in descending order, and then sort invoices of the same type by ID in descending order.
	//
	// 	- 3: Sort invoices by invoice type in ascending order, and then sort invoices of the same type by ID in descending order.
	//
	// example:
	//
	// 1
	SortType *int32 `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// The minimum amount to be queried.
	//
	// example:
	//
	// 100
	StartAmount *int64 `json:"StartAmount,omitempty" xml:"StartAmount,omitempty"`
	// The earliest time when an order is paid. Specify the time in the yyyy-mm-dd hh:mm:ss format.
	//
	// example:
	//
	// 2020-02-02 12:00:00
	StartBizTime *string `json:"StartBizTime,omitempty" xml:"StartBizTime,omitempty"`
	// The beginning of the time range to query.
	//
	// example:
	//
	// 2020-02-02 12:00:00
	StartSearchTime *string `json:"StartSearchTime,omitempty" xml:"StartSearchTime,omitempty"`
	// The type of orders to be queried. Valid values:
	//
	// 	- 1: the orders in which the invoiceable amount is negative.
	//
	// 	- 2: the orders in which the invoiceable amount is positive.
	//
	// 	- 3: the orders in which the invoiceable amount is not 0.
	//
	// 	- 4: the orders in which the amount that has been invoiced is greater than 0.
	//
	// >  By default, this parameter is left empty. If this parameter is left empty, all orders are queried.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryEvaluateListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryEvaluateListRequest) GoString() string {
	return s.String()
}

func (s *QueryEvaluateListRequest) GetBillCycle() *string {
	return s.BillCycle
}

func (s *QueryEvaluateListRequest) GetBizTypeList() []*string {
	return s.BizTypeList
}

func (s *QueryEvaluateListRequest) GetEndAmount() *int64 {
	return s.EndAmount
}

func (s *QueryEvaluateListRequest) GetEndBizTime() *string {
	return s.EndBizTime
}

func (s *QueryEvaluateListRequest) GetEndSearchTime() *string {
	return s.EndSearchTime
}

func (s *QueryEvaluateListRequest) GetOutBizId() *string {
	return s.OutBizId
}

func (s *QueryEvaluateListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryEvaluateListRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryEvaluateListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryEvaluateListRequest) GetSortType() *int32 {
	return s.SortType
}

func (s *QueryEvaluateListRequest) GetStartAmount() *int64 {
	return s.StartAmount
}

func (s *QueryEvaluateListRequest) GetStartBizTime() *string {
	return s.StartBizTime
}

func (s *QueryEvaluateListRequest) GetStartSearchTime() *string {
	return s.StartSearchTime
}

func (s *QueryEvaluateListRequest) GetType() *int32 {
	return s.Type
}

func (s *QueryEvaluateListRequest) SetBillCycle(v string) *QueryEvaluateListRequest {
	s.BillCycle = &v
	return s
}

func (s *QueryEvaluateListRequest) SetBizTypeList(v []*string) *QueryEvaluateListRequest {
	s.BizTypeList = v
	return s
}

func (s *QueryEvaluateListRequest) SetEndAmount(v int64) *QueryEvaluateListRequest {
	s.EndAmount = &v
	return s
}

func (s *QueryEvaluateListRequest) SetEndBizTime(v string) *QueryEvaluateListRequest {
	s.EndBizTime = &v
	return s
}

func (s *QueryEvaluateListRequest) SetEndSearchTime(v string) *QueryEvaluateListRequest {
	s.EndSearchTime = &v
	return s
}

func (s *QueryEvaluateListRequest) SetOutBizId(v string) *QueryEvaluateListRequest {
	s.OutBizId = &v
	return s
}

func (s *QueryEvaluateListRequest) SetOwnerId(v int64) *QueryEvaluateListRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryEvaluateListRequest) SetPageNum(v int32) *QueryEvaluateListRequest {
	s.PageNum = &v
	return s
}

func (s *QueryEvaluateListRequest) SetPageSize(v int32) *QueryEvaluateListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryEvaluateListRequest) SetSortType(v int32) *QueryEvaluateListRequest {
	s.SortType = &v
	return s
}

func (s *QueryEvaluateListRequest) SetStartAmount(v int64) *QueryEvaluateListRequest {
	s.StartAmount = &v
	return s
}

func (s *QueryEvaluateListRequest) SetStartBizTime(v string) *QueryEvaluateListRequest {
	s.StartBizTime = &v
	return s
}

func (s *QueryEvaluateListRequest) SetStartSearchTime(v string) *QueryEvaluateListRequest {
	s.StartSearchTime = &v
	return s
}

func (s *QueryEvaluateListRequest) SetType(v int32) *QueryEvaluateListRequest {
	s.Type = &v
	return s
}

func (s *QueryEvaluateListRequest) Validate() error {
	return dara.Validate(s)
}
