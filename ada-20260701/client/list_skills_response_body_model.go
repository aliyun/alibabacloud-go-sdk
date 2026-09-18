// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSkillsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSkillsResponseBody
	GetNextToken() *string
	SetPageNumber(v int64) *ListSkillsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSkillsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListSkillsResponseBody
	GetRequestId() *string
	SetSkills(v []interface{}) *ListSkillsResponseBody
	GetSkills() []interface{}
	SetTotal(v int64) *ListSkillsResponseBody
	GetTotal() *int64
}

type ListSkillsResponseBody struct {
	// The number of entries per page for the current cursor-based pagination.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page. An empty string is returned if there is no next page.
	//
	// example:
	//
	// ""
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number for compatible page-number-based pagination.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for compatible page-number-based pagination.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of Skill summaries. The current public contract does not define a fixed structure for list items. For common fields, see "Supplementary description of response elements".
	//
	// example:
	//
	// [{"skillId":"skill_example123","name":"code-review","official":false,"description":"A Skill for performing code reviews","iconUrl":"https://example.com/icons/code-review.png","visibility":"user","status":"PUBLISHED","publishedVersion":2,"creatorId":"example-user","createdAt":1760000000000,"canModify":true,"canDelete":true}]
	Skills []interface{} `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	// The total number of Skills that match the current visibility and filter conditions.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSkillsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSkillsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSkillsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillsResponseBody) GetSkills() []interface{} {
	return s.Skills
}

func (s *ListSkillsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListSkillsResponseBody) SetMaxResults(v int32) *ListSkillsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSkillsResponseBody) SetNextToken(v string) *ListSkillsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSkillsResponseBody) SetPageNumber(v int64) *ListSkillsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSkillsResponseBody) SetPageSize(v int64) *ListSkillsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSkillsResponseBody) SetRequestId(v string) *ListSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillsResponseBody) SetSkills(v []interface{}) *ListSkillsResponseBody {
	s.Skills = v
	return s
}

func (s *ListSkillsResponseBody) SetTotal(v int64) *ListSkillsResponseBody {
	s.Total = &v
	return s
}

func (s *ListSkillsResponseBody) Validate() error {
	return dara.Validate(s)
}
