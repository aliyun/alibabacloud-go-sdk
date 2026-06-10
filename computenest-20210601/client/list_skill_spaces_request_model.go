// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillSpacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListSkillSpacesRequestFilter) *ListSkillSpacesRequest
	GetFilter() []*ListSkillSpacesRequestFilter
	SetMaxResults(v int32) *ListSkillSpacesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSkillSpacesRequest
	GetNextToken() *string
}

type ListSkillSpacesRequest struct {
	Filter []*ListSkillSpacesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// BBBAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListSkillSpacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillSpacesRequest) GoString() string {
	return s.String()
}

func (s *ListSkillSpacesRequest) GetFilter() []*ListSkillSpacesRequestFilter {
	return s.Filter
}

func (s *ListSkillSpacesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSkillSpacesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSkillSpacesRequest) SetFilter(v []*ListSkillSpacesRequestFilter) *ListSkillSpacesRequest {
	s.Filter = v
	return s
}

func (s *ListSkillSpacesRequest) SetMaxResults(v int32) *ListSkillSpacesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSkillSpacesRequest) SetNextToken(v string) *ListSkillSpacesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSkillSpacesRequest) Validate() error {
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

type ListSkillSpacesRequestFilter struct {
	// example:
	//
	// SkillSpaceId
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListSkillSpacesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListSkillSpacesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListSkillSpacesRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListSkillSpacesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListSkillSpacesRequestFilter) SetName(v string) *ListSkillSpacesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListSkillSpacesRequestFilter) SetValue(v []*string) *ListSkillSpacesRequestFilter {
	s.Value = v
	return s
}

func (s *ListSkillSpacesRequestFilter) Validate() error {
	return dara.Validate(s)
}
