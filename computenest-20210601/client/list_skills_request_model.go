// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListSkillsRequestFilter) *ListSkillsRequest
	GetFilter() []*ListSkillsRequestFilter
	SetMaxResults(v int32) *ListSkillsRequest
	GetMaxResults() *int32
	SetNeedDownloadUrl(v bool) *ListSkillsRequest
	GetNeedDownloadUrl() *bool
	SetNextToken(v string) *ListSkillsRequest
	GetNextToken() *string
}

type ListSkillsRequest struct {
	// The filters for querying Skills.
	Filter []*ListSkillsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The maximum number of entries to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies whether to return the download URL of the Skill package.
	//
	// example:
	//
	// false
	NeedDownloadUrl *bool `json:"NeedDownloadUrl,omitempty" xml:"NeedDownloadUrl,omitempty"`
	// The token for the next page of results. Leave this parameter empty for the first request. For subsequent requests, use the `NextToken` value from the previous response.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsRequest) GoString() string {
	return s.String()
}

func (s *ListSkillsRequest) GetFilter() []*ListSkillsRequestFilter {
	return s.Filter
}

func (s *ListSkillsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSkillsRequest) GetNeedDownloadUrl() *bool {
	return s.NeedDownloadUrl
}

func (s *ListSkillsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSkillsRequest) SetFilter(v []*ListSkillsRequestFilter) *ListSkillsRequest {
	s.Filter = v
	return s
}

func (s *ListSkillsRequest) SetMaxResults(v int32) *ListSkillsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSkillsRequest) SetNeedDownloadUrl(v bool) *ListSkillsRequest {
	s.NeedDownloadUrl = &v
	return s
}

func (s *ListSkillsRequest) SetNextToken(v string) *ListSkillsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSkillsRequest) Validate() error {
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

type ListSkillsRequestFilter struct {
	// The filter name. Valid values:
	//
	// - `SkillId`: The Skill ID. An exact match is performed.
	//
	// - `SkillSpaceId`: The ID of the SkillSpace. An exact match is performed.
	//
	// - `SkillName`: The Skill name.
	//
	// - `MatchType`: The match type for `SkillName`. Valid values: `exact` (exact match), `prefix` (prefix match), and `fuzzy` (fuzzy match).
	//
	// - `SkillType`: The Skill type. Valid values: `official` and `custom`.
	//
	// - `Keyword`: The keyword for a fuzzy match on the Skill name or Skill description.
	//
	// - `SkillLabels`: The Skill labels. A fuzzy match is performed.
	//
	// example:
	//
	// SkillId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The filter values. You can specify a maximum of 10 values.
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListSkillsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListSkillsRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListSkillsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListSkillsRequestFilter) SetName(v string) *ListSkillsRequestFilter {
	s.Name = &v
	return s
}

func (s *ListSkillsRequestFilter) SetValue(v []*string) *ListSkillsRequestFilter {
	s.Value = v
	return s
}

func (s *ListSkillsRequestFilter) Validate() error {
	return dara.Validate(s)
}
