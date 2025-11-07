// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisNetworkRankingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *GetNisNetworkRankingShrinkRequest
	GetAccountIds() []*string
	SetBeginTime(v int64) *GetNisNetworkRankingShrinkRequest
	GetBeginTime() *int64
	SetDirection(v string) *GetNisNetworkRankingShrinkRequest
	GetDirection() *string
	SetEndTime(v int64) *GetNisNetworkRankingShrinkRequest
	GetEndTime() *int64
	SetFilterShrink(v string) *GetNisNetworkRankingShrinkRequest
	GetFilterShrink() *string
	SetGroupBy(v string) *GetNisNetworkRankingShrinkRequest
	GetGroupBy() *string
	SetOrderBy(v string) *GetNisNetworkRankingShrinkRequest
	GetOrderBy() *string
	SetRegionNo(v string) *GetNisNetworkRankingShrinkRequest
	GetRegionNo() *string
	SetResourceType(v string) *GetNisNetworkRankingShrinkRequest
	GetResourceType() *string
	SetSort(v string) *GetNisNetworkRankingShrinkRequest
	GetSort() *string
	SetTopN(v int32) *GetNisNetworkRankingShrinkRequest
	GetTopN() *int32
	SetUseCrossAccount(v bool) *GetNisNetworkRankingShrinkRequest
	GetUseCrossAccount() *bool
}

type GetNisNetworkRankingShrinkRequest struct {
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1684373600099
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// in
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// example:
	//
	// 1684379093000
	EndTime      *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Protocol
	GroupBy *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bps
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AccessInternetIpV4
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// example:
	//
	// false
	UseCrossAccount *bool `json:"UseCrossAccount,omitempty" xml:"UseCrossAccount,omitempty"`
}

func (s GetNisNetworkRankingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkRankingShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetNisNetworkRankingShrinkRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *GetNisNetworkRankingShrinkRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetNisNetworkRankingShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *GetNisNetworkRankingShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetNisNetworkRankingShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *GetNisNetworkRankingShrinkRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *GetNisNetworkRankingShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetNisNetworkRankingShrinkRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *GetNisNetworkRankingShrinkRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetNisNetworkRankingShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *GetNisNetworkRankingShrinkRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *GetNisNetworkRankingShrinkRequest) GetUseCrossAccount() *bool {
	return s.UseCrossAccount
}

func (s *GetNisNetworkRankingShrinkRequest) SetAccountIds(v []*string) *GetNisNetworkRankingShrinkRequest {
	s.AccountIds = v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetBeginTime(v int64) *GetNisNetworkRankingShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetDirection(v string) *GetNisNetworkRankingShrinkRequest {
	s.Direction = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetEndTime(v int64) *GetNisNetworkRankingShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetFilterShrink(v string) *GetNisNetworkRankingShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetGroupBy(v string) *GetNisNetworkRankingShrinkRequest {
	s.GroupBy = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetOrderBy(v string) *GetNisNetworkRankingShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetRegionNo(v string) *GetNisNetworkRankingShrinkRequest {
	s.RegionNo = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetResourceType(v string) *GetNisNetworkRankingShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetSort(v string) *GetNisNetworkRankingShrinkRequest {
	s.Sort = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetTopN(v int32) *GetNisNetworkRankingShrinkRequest {
	s.TopN = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) SetUseCrossAccount(v bool) *GetNisNetworkRankingShrinkRequest {
	s.UseCrossAccount = &v
	return s
}

func (s *GetNisNetworkRankingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
