// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceCountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*GetResourceCountsRequestFilter) *GetResourceCountsRequest
	GetFilter() []*GetResourceCountsRequestFilter
	SetGroupByKey(v string) *GetResourceCountsRequest
	GetGroupByKey() *string
}

type GetResourceCountsRequest struct {
	// The filter conditions.
	Filter []*GetResourceCountsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The dimension by which resources are queried. Valid values:
	//
	// 	- ResourceType
	//
	// 	- Region
	//
	// 	- ResourceGroupId
	//
	// 	- TagKey
	//
	// 	- TagValue
	//
	// example:
	//
	// ResourceType
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
}

func (s GetResourceCountsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCountsRequest) GoString() string {
	return s.String()
}

func (s *GetResourceCountsRequest) GetFilter() []*GetResourceCountsRequestFilter {
	return s.Filter
}

func (s *GetResourceCountsRequest) GetGroupByKey() *string {
	return s.GroupByKey
}

func (s *GetResourceCountsRequest) SetFilter(v []*GetResourceCountsRequestFilter) *GetResourceCountsRequest {
	s.Filter = v
	return s
}

func (s *GetResourceCountsRequest) SetGroupByKey(v string) *GetResourceCountsRequest {
	s.GroupByKey = &v
	return s
}

func (s *GetResourceCountsRequest) Validate() error {
	return dara.Validate(s)
}

type GetResourceCountsRequestFilter struct {
	// The key of the filter condition. For more information, see `Supported filter parameters`.
	//
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The matching mode.
	//
	// The value Equals indicates an equal match.
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The values of the filter condition.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s GetResourceCountsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetResourceCountsRequestFilter) GoString() string {
	return s.String()
}

func (s *GetResourceCountsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *GetResourceCountsRequestFilter) GetMatchType() *string {
	return s.MatchType
}

func (s *GetResourceCountsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *GetResourceCountsRequestFilter) SetKey(v string) *GetResourceCountsRequestFilter {
	s.Key = &v
	return s
}

func (s *GetResourceCountsRequestFilter) SetMatchType(v string) *GetResourceCountsRequestFilter {
	s.MatchType = &v
	return s
}

func (s *GetResourceCountsRequestFilter) SetValue(v []*string) *GetResourceCountsRequestFilter {
	s.Value = v
	return s
}

func (s *GetResourceCountsRequestFilter) Validate() error {
	return dara.Validate(s)
}
