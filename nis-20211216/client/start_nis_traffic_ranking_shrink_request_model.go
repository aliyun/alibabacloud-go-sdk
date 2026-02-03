// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartNisTrafficRankingShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *StartNisTrafficRankingShrinkRequest
	GetBeginTime() *int64
	SetDirection(v string) *StartNisTrafficRankingShrinkRequest
	GetDirection() *string
	SetEndTime(v int64) *StartNisTrafficRankingShrinkRequest
	GetEndTime() *int64
	SetFilterShrink(v string) *StartNisTrafficRankingShrinkRequest
	GetFilterShrink() *string
	SetGroupByShrink(v string) *StartNisTrafficRankingShrinkRequest
	GetGroupByShrink() *string
	SetLanguage(v string) *StartNisTrafficRankingShrinkRequest
	GetLanguage() *string
	SetMaxResults(v int32) *StartNisTrafficRankingShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *StartNisTrafficRankingShrinkRequest
	GetNextToken() *string
	SetOrderBy(v string) *StartNisTrafficRankingShrinkRequest
	GetOrderBy() *string
	SetRegionNo(v string) *StartNisTrafficRankingShrinkRequest
	GetRegionNo() *string
	SetSort(v string) *StartNisTrafficRankingShrinkRequest
	GetSort() *string
	SetStorageInterval(v int32) *StartNisTrafficRankingShrinkRequest
	GetStorageInterval() *int32
	SetTopN(v int32) *StartNisTrafficRankingShrinkRequest
	GetTopN() *int32
	SetTrafficAnalyzerId(v string) *StartNisTrafficRankingShrinkRequest
	GetTrafficAnalyzerId() *string
	SetTrafficScenario(v string) *StartNisTrafficRankingShrinkRequest
	GetTrafficScenario() *string
	SetTupleDimension(v string) *StartNisTrafficRankingShrinkRequest
	GetTupleDimension() *string
}

type StartNisTrafficRankingShrinkRequest struct {
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
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FilterShrink  *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	GroupByShrink *string `json:"GroupBy,omitempty" xml:"GroupBy,omitempty"`
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

func (s StartNisTrafficRankingShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartNisTrafficRankingShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartNisTrafficRankingShrinkRequest) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *StartNisTrafficRankingShrinkRequest) GetDirection() *string {
	return s.Direction
}

func (s *StartNisTrafficRankingShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *StartNisTrafficRankingShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *StartNisTrafficRankingShrinkRequest) GetGroupByShrink() *string {
	return s.GroupByShrink
}

func (s *StartNisTrafficRankingShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *StartNisTrafficRankingShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *StartNisTrafficRankingShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *StartNisTrafficRankingShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *StartNisTrafficRankingShrinkRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *StartNisTrafficRankingShrinkRequest) GetSort() *string {
	return s.Sort
}

func (s *StartNisTrafficRankingShrinkRequest) GetStorageInterval() *int32 {
	return s.StorageInterval
}

func (s *StartNisTrafficRankingShrinkRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *StartNisTrafficRankingShrinkRequest) GetTrafficAnalyzerId() *string {
	return s.TrafficAnalyzerId
}

func (s *StartNisTrafficRankingShrinkRequest) GetTrafficScenario() *string {
	return s.TrafficScenario
}

func (s *StartNisTrafficRankingShrinkRequest) GetTupleDimension() *string {
	return s.TupleDimension
}

func (s *StartNisTrafficRankingShrinkRequest) SetBeginTime(v int64) *StartNisTrafficRankingShrinkRequest {
	s.BeginTime = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetDirection(v string) *StartNisTrafficRankingShrinkRequest {
	s.Direction = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetEndTime(v int64) *StartNisTrafficRankingShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetFilterShrink(v string) *StartNisTrafficRankingShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetGroupByShrink(v string) *StartNisTrafficRankingShrinkRequest {
	s.GroupByShrink = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetLanguage(v string) *StartNisTrafficRankingShrinkRequest {
	s.Language = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetMaxResults(v int32) *StartNisTrafficRankingShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetNextToken(v string) *StartNisTrafficRankingShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetOrderBy(v string) *StartNisTrafficRankingShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetRegionNo(v string) *StartNisTrafficRankingShrinkRequest {
	s.RegionNo = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetSort(v string) *StartNisTrafficRankingShrinkRequest {
	s.Sort = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetStorageInterval(v int32) *StartNisTrafficRankingShrinkRequest {
	s.StorageInterval = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetTopN(v int32) *StartNisTrafficRankingShrinkRequest {
	s.TopN = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetTrafficAnalyzerId(v string) *StartNisTrafficRankingShrinkRequest {
	s.TrafficAnalyzerId = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetTrafficScenario(v string) *StartNisTrafficRankingShrinkRequest {
	s.TrafficScenario = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) SetTupleDimension(v string) *StartNisTrafficRankingShrinkRequest {
	s.TupleDimension = &v
	return s
}

func (s *StartNisTrafficRankingShrinkRequest) Validate() error {
	return dara.Validate(s)
}
