// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDigitalEmployeeSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDigitalEmployeeSkillsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDigitalEmployeeSkillsRequest
	GetNextToken() *string
	SetSkillName(v string) *ListDigitalEmployeeSkillsRequest
	GetSkillName() *string
}

type ListDigitalEmployeeSkillsRequest struct {
	// The maximum number of entries to return for this request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next page. A null value indicates that there are no more pages.
	//
	// example:
	//
	// sjC5rekx93Ew7K7VcmI3wkBZBYQ-GphB2ilQu3zJCGxoZuicwyJznfo2riTjr-lq
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The name of the skill.
	//
	// example:
	//
	// test
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
}

func (s ListDigitalEmployeeSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeeSkillsRequest) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeeSkillsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDigitalEmployeeSkillsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDigitalEmployeeSkillsRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *ListDigitalEmployeeSkillsRequest) SetMaxResults(v int32) *ListDigitalEmployeeSkillsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDigitalEmployeeSkillsRequest) SetNextToken(v string) *ListDigitalEmployeeSkillsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDigitalEmployeeSkillsRequest) SetSkillName(v string) *ListDigitalEmployeeSkillsRequest {
	s.SkillName = &v
	return s
}

func (s *ListDigitalEmployeeSkillsRequest) Validate() error {
	return dara.Validate(s)
}
