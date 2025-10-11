// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountResourceCountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*GetMultiAccountResourceCountsRequestFilter) *GetMultiAccountResourceCountsRequest
	GetFilter() []*GetMultiAccountResourceCountsRequestFilter
	SetGroupByKey(v string) *GetMultiAccountResourceCountsRequest
	GetGroupByKey() *string
	SetScope(v string) *GetMultiAccountResourceCountsRequest
	GetScope() *string
}

type GetMultiAccountResourceCountsRequest struct {
	Filter []*GetMultiAccountResourceCountsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// ResourceType
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// example:
	//
	// rd-r4****
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s GetMultiAccountResourceCountsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceCountsRequest) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsRequest) GetFilter() []*GetMultiAccountResourceCountsRequestFilter {
	return s.Filter
}

func (s *GetMultiAccountResourceCountsRequest) GetGroupByKey() *string {
	return s.GroupByKey
}

func (s *GetMultiAccountResourceCountsRequest) GetScope() *string {
	return s.Scope
}

func (s *GetMultiAccountResourceCountsRequest) SetFilter(v []*GetMultiAccountResourceCountsRequestFilter) *GetMultiAccountResourceCountsRequest {
	s.Filter = v
	return s
}

func (s *GetMultiAccountResourceCountsRequest) SetGroupByKey(v string) *GetMultiAccountResourceCountsRequest {
	s.GroupByKey = &v
	return s
}

func (s *GetMultiAccountResourceCountsRequest) SetScope(v string) *GetMultiAccountResourceCountsRequest {
	s.Scope = &v
	return s
}

func (s *GetMultiAccountResourceCountsRequest) Validate() error {
	return dara.Validate(s)
}

type GetMultiAccountResourceCountsRequestFilter struct {
	// example:
	//
	// RegionId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// Equals
	MatchType *string   `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	Value     []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s GetMultiAccountResourceCountsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountResourceCountsRequestFilter) GoString() string {
	return s.String()
}

func (s *GetMultiAccountResourceCountsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *GetMultiAccountResourceCountsRequestFilter) GetMatchType() *string {
	return s.MatchType
}

func (s *GetMultiAccountResourceCountsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *GetMultiAccountResourceCountsRequestFilter) SetKey(v string) *GetMultiAccountResourceCountsRequestFilter {
	s.Key = &v
	return s
}

func (s *GetMultiAccountResourceCountsRequestFilter) SetMatchType(v string) *GetMultiAccountResourceCountsRequestFilter {
	s.MatchType = &v
	return s
}

func (s *GetMultiAccountResourceCountsRequestFilter) SetValue(v []*string) *GetMultiAccountResourceCountsRequestFilter {
	s.Value = v
	return s
}

func (s *GetMultiAccountResourceCountsRequestFilter) Validate() error {
	return dara.Validate(s)
}
