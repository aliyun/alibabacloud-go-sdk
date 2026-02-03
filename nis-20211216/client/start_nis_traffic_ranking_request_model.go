// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartNisTrafficRankingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *StartNisTrafficRankingRequest
	GetBeginTime() *int64
	SetDirection(v string) *StartNisTrafficRankingRequest
	GetDirection() *string
	SetEndTime(v int64) *StartNisTrafficRankingRequest
	GetEndTime() *int64
	SetFilter(v []*StartNisTrafficRankingRequestFilter) *StartNisTrafficRankingRequest
	GetFilter() []*StartNisTrafficRankingRequestFilter
	SetGroupBy(v []*string) *StartNisTrafficRankingRequest
	GetGroupBy() []*string
	SetLanguage(v string) *StartNisTrafficRankingRequest
	GetLanguage() *string
	SetMaxResults(v int32) *StartNisTrafficRankingRequest
	GetMaxResults() *int32
	SetNextToken(v string) *StartNisTrafficRankingRequest
	GetNextToken() *string
	SetOrderBy(v string) *StartNisTrafficRankingRequest
	GetOrderBy() *string
	SetRegionNo(v string) *StartNisTrafficRankingRequest
	GetRegionNo() *string
	SetSort(v string) *StartNisTrafficRankingRequest
	GetSort() *string
	SetStorageInterval(v int32) *StartNisTrafficRankingRequest
	GetStorageInterval() *int32
	SetTopN(v int32) *StartNisTrafficRankingRequest
	GetTopN() *int32
	SetTrafficAnalyzerId(v string) *StartNisTrafficRankingRequest
	GetTrafficAnalyzerId() *string
	SetTrafficScenario(v string) *StartNisTrafficRankingRequest
	GetTrafficScenario() *string
	SetTupleDimension(v string) *StartNisTrafficRankingRequest
	GetTupleDimension() *string
}

type StartNisTrafficRankingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1638239092000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Out
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1684373700099
	EndTime *int64                                 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filter  []*StartNisTrafficRankingRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	GroupBy []*string                              `json:"GroupBy,omitempty" xml:"GroupBy,omitempty" type:"Repeated"`
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 2A07PfBPlzmmNi/75Qca9SK73UfY48/+WBiREjfVfXqMQxtV8XckOg5lk7F2bhC+
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bytes
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// Desc
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// example:
	//
	// 10
	StorageInterval *int32 `json:"StorageInterval,omitempty" xml:"StorageInterval,omitempty"`
	// example:
	//
	// 10
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// nta-262****ca07f
	TrafficAnalyzerId *string `json:"TrafficAnalyzerId,omitempty" xml:"TrafficAnalyzerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VpcFlowLogAll
	TrafficScenario *string `json:"TrafficScenario,omitempty" xml:"TrafficScenario,omitempty"`
	// example:
	//
	// Tuple1
	TupleDimension *string `json:"TupleDimension,omitempty" xml:"TupleDimension,omitempty"`
}

func (s StartNisTrafficRankingRequest) String() string {
	return dara.Prettify(s)
}

func (s StartNisTrafficRankingRequest) GoString() string {
	return s.String()
}

func (s *StartNisTrafficRankingRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *StartNisTrafficRankingRequest) GetDirection() *string {
	return s.Direction
}

func (s *StartNisTrafficRankingRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *StartNisTrafficRankingRequest) GetFilter() []*StartNisTrafficRankingRequestFilter {
	return s.Filter
}

func (s *StartNisTrafficRankingRequest) GetGroupBy() []*string {
	return s.GroupBy
}

func (s *StartNisTrafficRankingRequest) GetLanguage() *string {
	return s.Language
}

func (s *StartNisTrafficRankingRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *StartNisTrafficRankingRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *StartNisTrafficRankingRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *StartNisTrafficRankingRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *StartNisTrafficRankingRequest) GetSort() *string {
	return s.Sort
}

func (s *StartNisTrafficRankingRequest) GetStorageInterval() *int32 {
	return s.StorageInterval
}

func (s *StartNisTrafficRankingRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *StartNisTrafficRankingRequest) GetTrafficAnalyzerId() *string {
	return s.TrafficAnalyzerId
}

func (s *StartNisTrafficRankingRequest) GetTrafficScenario() *string {
	return s.TrafficScenario
}

func (s *StartNisTrafficRankingRequest) GetTupleDimension() *string {
	return s.TupleDimension
}

func (s *StartNisTrafficRankingRequest) SetBeginTime(v int64) *StartNisTrafficRankingRequest {
	s.BeginTime = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetDirection(v string) *StartNisTrafficRankingRequest {
	s.Direction = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetEndTime(v int64) *StartNisTrafficRankingRequest {
	s.EndTime = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetFilter(v []*StartNisTrafficRankingRequestFilter) *StartNisTrafficRankingRequest {
	s.Filter = v
	return s
}

func (s *StartNisTrafficRankingRequest) SetGroupBy(v []*string) *StartNisTrafficRankingRequest {
	s.GroupBy = v
	return s
}

func (s *StartNisTrafficRankingRequest) SetLanguage(v string) *StartNisTrafficRankingRequest {
	s.Language = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetMaxResults(v int32) *StartNisTrafficRankingRequest {
	s.MaxResults = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetNextToken(v string) *StartNisTrafficRankingRequest {
	s.NextToken = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetOrderBy(v string) *StartNisTrafficRankingRequest {
	s.OrderBy = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetRegionNo(v string) *StartNisTrafficRankingRequest {
	s.RegionNo = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetSort(v string) *StartNisTrafficRankingRequest {
	s.Sort = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetStorageInterval(v int32) *StartNisTrafficRankingRequest {
	s.StorageInterval = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetTopN(v int32) *StartNisTrafficRankingRequest {
	s.TopN = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetTrafficAnalyzerId(v string) *StartNisTrafficRankingRequest {
	s.TrafficAnalyzerId = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetTrafficScenario(v string) *StartNisTrafficRankingRequest {
	s.TrafficScenario = &v
	return s
}

func (s *StartNisTrafficRankingRequest) SetTupleDimension(v string) *StartNisTrafficRankingRequest {
	s.TupleDimension = &v
	return s
}

func (s *StartNisTrafficRankingRequest) Validate() error {
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

type StartNisTrafficRankingRequestFilter struct {
	// example:
	//
	// FlowAction
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// in
	Operator *string   `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s StartNisTrafficRankingRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s StartNisTrafficRankingRequestFilter) GoString() string {
	return s.String()
}

func (s *StartNisTrafficRankingRequestFilter) GetKey() *string {
	return s.Key
}

func (s *StartNisTrafficRankingRequestFilter) GetOperator() *string {
	return s.Operator
}

func (s *StartNisTrafficRankingRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *StartNisTrafficRankingRequestFilter) SetKey(v string) *StartNisTrafficRankingRequestFilter {
	s.Key = &v
	return s
}

func (s *StartNisTrafficRankingRequestFilter) SetOperator(v string) *StartNisTrafficRankingRequestFilter {
	s.Operator = &v
	return s
}

func (s *StartNisTrafficRankingRequestFilter) SetValue(v []*string) *StartNisTrafficRankingRequestFilter {
	s.Value = v
	return s
}

func (s *StartNisTrafficRankingRequestFilter) Validate() error {
	return dara.Validate(s)
}
