// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Skill) *SearchSkillsResponseBody
	GetData() []*Skill
	SetMaxResults(v int32) *SearchSkillsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *SearchSkillsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *SearchSkillsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *SearchSkillsResponseBody
	GetTotalCount() *int32
}

type SearchSkillsResponseBody struct {
	Data []*Skill `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// AAAAAZjtYxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 53EAEBC0-4DEC-5AF4-AA21-3923D5A819C3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 22
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s SearchSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchSkillsResponseBody) GetData() []*Skill {
	return s.Data
}

func (s *SearchSkillsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchSkillsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchSkillsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchSkillsResponseBody) SetData(v []*Skill) *SearchSkillsResponseBody {
	s.Data = v
	return s
}

func (s *SearchSkillsResponseBody) SetMaxResults(v int32) *SearchSkillsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *SearchSkillsResponseBody) SetNextToken(v string) *SearchSkillsResponseBody {
	s.NextToken = &v
	return s
}

func (s *SearchSkillsResponseBody) SetRequestId(v string) *SearchSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchSkillsResponseBody) SetTotalCount(v int32) *SearchSkillsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchSkillsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
