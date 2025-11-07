// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisNetworkRankingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountIds(v []*string) *GetNisNetworkRankingRequest
	GetAccountIds() []*string
	SetBeginTime(v int64) *GetNisNetworkRankingRequest
	GetBeginTime() *int64
	SetDirection(v string) *GetNisNetworkRankingRequest
	GetDirection() *string
	SetEndTime(v int64) *GetNisNetworkRankingRequest
	GetEndTime() *int64
	SetFilter(v []*GetNisNetworkRankingRequestFilter) *GetNisNetworkRankingRequest
	GetFilter() []*GetNisNetworkRankingRequestFilter
	SetGroupBy(v string) *GetNisNetworkRankingRequest
	GetGroupBy() *string
	SetOrderBy(v string) *GetNisNetworkRankingRequest
	GetOrderBy() *string
	SetRegionNo(v string) *GetNisNetworkRankingRequest
	GetRegionNo() *string
	SetResourceType(v string) *GetNisNetworkRankingRequest
	GetResourceType() *string
	SetSort(v string) *GetNisNetworkRankingRequest
	GetSort() *string
	SetTopN(v int32) *GetNisNetworkRankingRequest
	GetTopN() *int32
	SetUseCrossAccount(v bool) *GetNisNetworkRankingRequest
	GetUseCrossAccount() *bool
}

type GetNisNetworkRankingRequest struct {
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
	EndTime *int64                               `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filter  []*GetNisNetworkRankingRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
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

func (s GetNisNetworkRankingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkRankingRequest) GoString() string {
	return s.String()
}

func (s *GetNisNetworkRankingRequest) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *GetNisNetworkRankingRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetNisNetworkRankingRequest) GetDirection() *string {
	return s.Direction
}

func (s *GetNisNetworkRankingRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetNisNetworkRankingRequest) GetFilter() []*GetNisNetworkRankingRequestFilter {
	return s.Filter
}

func (s *GetNisNetworkRankingRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *GetNisNetworkRankingRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetNisNetworkRankingRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *GetNisNetworkRankingRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetNisNetworkRankingRequest) GetSort() *string {
	return s.Sort
}

func (s *GetNisNetworkRankingRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *GetNisNetworkRankingRequest) GetUseCrossAccount() *bool {
	return s.UseCrossAccount
}

func (s *GetNisNetworkRankingRequest) SetAccountIds(v []*string) *GetNisNetworkRankingRequest {
	s.AccountIds = v
	return s
}

func (s *GetNisNetworkRankingRequest) SetBeginTime(v int64) *GetNisNetworkRankingRequest {
	s.BeginTime = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetDirection(v string) *GetNisNetworkRankingRequest {
	s.Direction = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetEndTime(v int64) *GetNisNetworkRankingRequest {
	s.EndTime = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetFilter(v []*GetNisNetworkRankingRequestFilter) *GetNisNetworkRankingRequest {
	s.Filter = v
	return s
}

func (s *GetNisNetworkRankingRequest) SetGroupBy(v string) *GetNisNetworkRankingRequest {
	s.GroupBy = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetOrderBy(v string) *GetNisNetworkRankingRequest {
	s.OrderBy = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetRegionNo(v string) *GetNisNetworkRankingRequest {
	s.RegionNo = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetResourceType(v string) *GetNisNetworkRankingRequest {
	s.ResourceType = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetSort(v string) *GetNisNetworkRankingRequest {
	s.Sort = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetTopN(v int32) *GetNisNetworkRankingRequest {
	s.TopN = &v
	return s
}

func (s *GetNisNetworkRankingRequest) SetUseCrossAccount(v bool) *GetNisNetworkRankingRequest {
	s.UseCrossAccount = &v
	return s
}

func (s *GetNisNetworkRankingRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetNisNetworkRankingRequestFilter struct {
	// example:
	//
	// instanceId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// lb-2zxxxxz1d
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetNisNetworkRankingRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkRankingRequestFilter) GoString() string {
	return s.String()
}

func (s *GetNisNetworkRankingRequestFilter) GetName() *string {
	return s.Name
}

func (s *GetNisNetworkRankingRequestFilter) GetValue() *string {
	return s.Value
}

func (s *GetNisNetworkRankingRequestFilter) SetName(v string) *GetNisNetworkRankingRequestFilter {
	s.Name = &v
	return s
}

func (s *GetNisNetworkRankingRequestFilter) SetValue(v string) *GetNisNetworkRankingRequestFilter {
	s.Value = &v
	return s
}

func (s *GetNisNetworkRankingRequestFilter) Validate() error {
	return dara.Validate(s)
}
