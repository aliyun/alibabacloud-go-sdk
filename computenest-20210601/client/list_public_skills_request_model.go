// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListPublicSkillsRequestFilter) *ListPublicSkillsRequest
	GetFilter() []*ListPublicSkillsRequestFilter
	SetMaxResults(v int32) *ListPublicSkillsRequest
	GetMaxResults() *int32
	SetNeedDownloadUrl(v bool) *ListPublicSkillsRequest
	GetNeedDownloadUrl() *bool
	SetNextToken(v string) *ListPublicSkillsRequest
	GetNextToken() *string
}

type ListPublicSkillsRequest struct {
	// A list of filters.
	Filter []*ListPublicSkillsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies whether to return the download link for the skill package.
	//
	// example:
	//
	// false
	NeedDownloadUrl *bool `json:"NeedDownloadUrl,omitempty" xml:"NeedDownloadUrl,omitempty"`
	// The token to retrieve the next page of results.
	//
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListPublicSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublicSkillsRequest) GoString() string {
	return s.String()
}

func (s *ListPublicSkillsRequest) GetFilter() []*ListPublicSkillsRequestFilter {
	return s.Filter
}

func (s *ListPublicSkillsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPublicSkillsRequest) GetNeedDownloadUrl() *bool {
	return s.NeedDownloadUrl
}

func (s *ListPublicSkillsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPublicSkillsRequest) SetFilter(v []*ListPublicSkillsRequestFilter) *ListPublicSkillsRequest {
	s.Filter = v
	return s
}

func (s *ListPublicSkillsRequest) SetMaxResults(v int32) *ListPublicSkillsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPublicSkillsRequest) SetNeedDownloadUrl(v bool) *ListPublicSkillsRequest {
	s.NeedDownloadUrl = &v
	return s
}

func (s *ListPublicSkillsRequest) SetNextToken(v string) *ListPublicSkillsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPublicSkillsRequest) Validate() error {
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

type ListPublicSkillsRequestFilter struct {
	// The filter criterion. Valid values:
	//
	// - `SkillId`: The ID of the skill. An exact match is performed.
	//
	// - `SkillName`: The name of the skill.
	//
	// - `MatchType`: The match type for the `SkillName` filter. Valid values: `exact` (exact match), `prefix` (prefix match), and `fuzzy` (fuzzy match).
	//
	// - `Keyword`: The keyword used for a fuzzy match on the skill name or skill description.
	//
	// - `SkillLabels`: The skill labels. A fuzzy match is performed.
	//
	// example:
	//
	// SkillId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The filter values. You can specify 1 to 10 values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListPublicSkillsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListPublicSkillsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListPublicSkillsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListPublicSkillsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListPublicSkillsRequestFilter) SetName(v string) *ListPublicSkillsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListPublicSkillsRequestFilter) SetValue(v []*string) *ListPublicSkillsRequestFilter {
	s.Value = v
	return s
}

func (s *ListPublicSkillsRequestFilter) Validate() error {
	return dara.Validate(s)
}
