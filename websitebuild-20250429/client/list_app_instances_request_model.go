// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ListAppInstancesRequest
	GetBizId() *string
	SetEndTimeBegin(v string) *ListAppInstancesRequest
	GetEndTimeBegin() *string
	SetEndTimeEnd(v string) *ListAppInstancesRequest
	GetEndTimeEnd() *string
	SetExtend(v string) *ListAppInstancesRequest
	GetExtend() *string
	SetMaxResults(v int32) *ListAppInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAppInstancesRequest
	GetNextToken() *string
	SetOrderColumn(v string) *ListAppInstancesRequest
	GetOrderColumn() *string
	SetOrderType(v string) *ListAppInstancesRequest
	GetOrderType() *string
	SetPageNum(v int32) *ListAppInstancesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListAppInstancesRequest
	GetPageSize() *int32
	SetQuery(v string) *ListAppInstancesRequest
	GetQuery() *string
	SetStatusList(v []*string) *ListAppInstancesRequest
	GetStatusList() []*string
}

type ListAppInstancesRequest struct {
	// example:
	//
	// WS20250731233102000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// 2025-12-18T14:30:00Z
	EndTimeBegin *string `json:"EndTimeBegin,omitempty" xml:"EndTimeBegin,omitempty"`
	// example:
	//
	// 2025-12-31T14:30:00Z
	EndTimeEnd *string `json:"EndTimeEnd,omitempty" xml:"EndTimeEnd,omitempty"`
	// example:
	//
	// {}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAARbaCuN6hiD08qrLdwJ9Fh3BFw8paIJ7ylB6A7Qn9JjM
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// gmtCreated
	OrderColumn *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	// example:
	//
	// DOWNGRADE
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 0
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// {\\"RequestId\\":\\"\\"}
	Query      *string   `json:"Query,omitempty" xml:"Query,omitempty"`
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s ListAppInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListAppInstancesRequest) GetBizId() *string {
	return s.BizId
}

func (s *ListAppInstancesRequest) GetEndTimeBegin() *string {
	return s.EndTimeBegin
}

func (s *ListAppInstancesRequest) GetEndTimeEnd() *string {
	return s.EndTimeEnd
}

func (s *ListAppInstancesRequest) GetExtend() *string {
	return s.Extend
}

func (s *ListAppInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppInstancesRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListAppInstancesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListAppInstancesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAppInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppInstancesRequest) GetQuery() *string {
	return s.Query
}

func (s *ListAppInstancesRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListAppInstancesRequest) SetBizId(v string) *ListAppInstancesRequest {
	s.BizId = &v
	return s
}

func (s *ListAppInstancesRequest) SetEndTimeBegin(v string) *ListAppInstancesRequest {
	s.EndTimeBegin = &v
	return s
}

func (s *ListAppInstancesRequest) SetEndTimeEnd(v string) *ListAppInstancesRequest {
	s.EndTimeEnd = &v
	return s
}

func (s *ListAppInstancesRequest) SetExtend(v string) *ListAppInstancesRequest {
	s.Extend = &v
	return s
}

func (s *ListAppInstancesRequest) SetMaxResults(v int32) *ListAppInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppInstancesRequest) SetNextToken(v string) *ListAppInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppInstancesRequest) SetOrderColumn(v string) *ListAppInstancesRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListAppInstancesRequest) SetOrderType(v string) *ListAppInstancesRequest {
	s.OrderType = &v
	return s
}

func (s *ListAppInstancesRequest) SetPageNum(v int32) *ListAppInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *ListAppInstancesRequest) SetPageSize(v int32) *ListAppInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppInstancesRequest) SetQuery(v string) *ListAppInstancesRequest {
	s.Query = &v
	return s
}

func (s *ListAppInstancesRequest) SetStatusList(v []*string) *ListAppInstancesRequest {
	s.StatusList = v
	return s
}

func (s *ListAppInstancesRequest) Validate() error {
	return dara.Validate(s)
}
