// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListSkillFilesRequestFilter) *ListSkillFilesRequest
	GetFilter() []*ListSkillFilesRequestFilter
	SetMaxResults(v int32) *ListSkillFilesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSkillFilesRequest
	GetNextToken() *string
	SetSkillId(v string) *ListSkillFilesRequest
	GetSkillId() *string
}

type ListSkillFilesRequest struct {
	Filter []*ListSkillFilesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAVz7BQqj2xtiNSC3d3RAD38=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Skill ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 06e9dca2-0ac9-4d2e-a965-e9db9c057e00
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s ListSkillFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillFilesRequest) GoString() string {
	return s.String()
}

func (s *ListSkillFilesRequest) GetFilter() []*ListSkillFilesRequestFilter {
	return s.Filter
}

func (s *ListSkillFilesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSkillFilesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSkillFilesRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *ListSkillFilesRequest) SetFilter(v []*ListSkillFilesRequestFilter) *ListSkillFilesRequest {
	s.Filter = v
	return s
}

func (s *ListSkillFilesRequest) SetMaxResults(v int32) *ListSkillFilesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSkillFilesRequest) SetNextToken(v string) *ListSkillFilesRequest {
	s.NextToken = &v
	return s
}

func (s *ListSkillFilesRequest) SetSkillId(v string) *ListSkillFilesRequest {
	s.SkillId = &v
	return s
}

func (s *ListSkillFilesRequest) Validate() error {
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

type ListSkillFilesRequestFilter struct {
	// example:
	//
	// 111111
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListSkillFilesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListSkillFilesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListSkillFilesRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListSkillFilesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListSkillFilesRequestFilter) SetName(v string) *ListSkillFilesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListSkillFilesRequestFilter) SetValue(v []*string) *ListSkillFilesRequestFilter {
	s.Value = v
	return s
}

func (s *ListSkillFilesRequestFilter) Validate() error {
	return dara.Validate(s)
}
