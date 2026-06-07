// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSkillsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSkillsRequest
	GetNextToken() *string
	SetQ(v string) *ListSkillsRequest
	GetQ() *string
	SetVisibility(v []*string) *ListSkillsRequest
	GetVisibility() []*string
}

type ListSkillsRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 5
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// analysis
	Q          *string   `json:"Q,omitempty" xml:"Q,omitempty"`
	Visibility []*string `json:"Visibility,omitempty" xml:"Visibility,omitempty" type:"Repeated"`
}

func (s ListSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsRequest) GoString() string {
	return s.String()
}

func (s *ListSkillsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSkillsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSkillsRequest) GetQ() *string {
	return s.Q
}

func (s *ListSkillsRequest) GetVisibility() []*string {
	return s.Visibility
}

func (s *ListSkillsRequest) SetMaxResults(v int32) *ListSkillsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSkillsRequest) SetNextToken(v string) *ListSkillsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSkillsRequest) SetQ(v string) *ListSkillsRequest {
	s.Q = &v
	return s
}

func (s *ListSkillsRequest) SetVisibility(v []*string) *ListSkillsRequest {
	s.Visibility = v
	return s
}

func (s *ListSkillsRequest) Validate() error {
	return dara.Validate(s)
}
