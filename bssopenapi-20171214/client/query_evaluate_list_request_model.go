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
	// 202502
	BillCycle *string `json:"BillCycle,omitempty" xml:"BillCycle,omitempty"`
	// The list of invoicing entities to which the invoiceable objects belong.
	//
	// example:
	//
	// ALIYUN
	BizTypeList []*string `json:"BizTypeList,omitempty" xml:"BizTypeList,omitempty" type:"Repeated"`
	// The maximum amount for the query.
	//
	// example:
	//
	// 1000
	EndAmount *int64 `json:"EndAmount,omitempty" xml:"EndAmount,omitempty"`
	// The end of the time range to query based on the successful payment time. The format is yyyy-mm-dd hh:mm:ss.
	//
	// example:
	//
	// 2025-03-01 00:00:00
	EndBizTime *string `json:"EndBizTime,omitempty" xml:"EndBizTime,omitempty"`
	// The end of the time range to query.
	//
	// example:
	//
	// 2025-03-01 00:00:00
	EndSearchTime *string `json:"EndSearchTime,omitempty" xml:"EndSearchTime,omitempty"`
	// The order ID, bill ID, or contract ID.
	//
	// example:
	//
	// 2387432832696
	OutBizId *string `json:"OutBizId,omitempty" xml:"OutBizId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting type. Valid values:
	//
	// - 1: sorts the results by ID in descending order.
	//
	// - 2: sorts the results by type in descending order, and then by ID in descending order.
	//
	// - 3: sorts the results by type in ascending order, and then by ID in descending order.
	//
	// example:
	//
	// 1
	SortType *int32 `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// The minimum amount for the query.
	//
	// example:
	//
	// 100
	StartAmount *int64 `json:"StartAmount,omitempty" xml:"StartAmount,omitempty"`
	// The beginning of the time range to query based on the successful payment time. The format is yyyy-mm-dd hh:mm:ss.
	//
	// example:
	//
	// 2025-02-01 00:00:00
	StartBizTime *string `json:"StartBizTime,omitempty" xml:"StartBizTime,omitempty"`
	// The start of the time range to query.
	//
	// example:
	//
	// 2025-02-01 00:00:00
	StartSearchTime *string `json:"StartSearchTime,omitempty" xml:"StartSearchTime,omitempty"`
	// The type. Valid values:
	//
	// - 1: queries objects with a negative invoiceable amount.
	//
	// - 2: queries objects with a positive invoiceable amount.
	//
	// - 3: queries all objects with a non-zero invoiceable amount.
	//
	// - 4: queries objects with an invoiced amount greater than 0.
	//
	// > If you leave this parameter empty, all objects are queried.
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
