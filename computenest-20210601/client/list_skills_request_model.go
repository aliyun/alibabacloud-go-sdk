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
	Filter []*ListSkillsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// false
	NeedDownloadUrl *bool `json:"NeedDownloadUrl,omitempty" xml:"NeedDownloadUrl,omitempty"`
	// NextToken
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
	// example:
	//
	// SkillId
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
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
