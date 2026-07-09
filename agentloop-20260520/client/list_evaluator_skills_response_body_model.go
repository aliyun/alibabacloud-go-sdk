// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEvaluatorSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListEvaluatorSkillsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListEvaluatorSkillsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEvaluatorSkillsResponseBody
	GetRequestId() *string
	SetSkills(v []*ListEvaluatorSkillsResponseBodySkills) *ListEvaluatorSkillsResponseBody
	GetSkills() []*ListEvaluatorSkillsResponseBodySkills
	SetTotal(v int64) *ListEvaluatorSkillsResponseBody
	GetTotal() *int64
}

type ListEvaluatorSkillsResponseBody struct {
	// The number of entries per page used in this request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// eyJuZXh0IjoiNDAifQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3FE4CD1E-FF41-56BE-B590-7A021D9C1524
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of skill summaries.
	//
	// example:
	//
	// [{"skillName":"trace_context_loader","displayName":"Trace 上下文读取","enable":true}]
	Skills []*ListEvaluatorSkillsResponseBodySkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The total number of skills.
	//
	// example:
	//
	// 3
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListEvaluatorSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluatorSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEvaluatorSkillsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListEvaluatorSkillsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEvaluatorSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEvaluatorSkillsResponseBody) GetSkills() []*ListEvaluatorSkillsResponseBodySkills {
	return s.Skills
}

func (s *ListEvaluatorSkillsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListEvaluatorSkillsResponseBody) SetMaxResults(v int32) *ListEvaluatorSkillsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListEvaluatorSkillsResponseBody) SetNextToken(v string) *ListEvaluatorSkillsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEvaluatorSkillsResponseBody) SetRequestId(v string) *ListEvaluatorSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEvaluatorSkillsResponseBody) SetSkills(v []*ListEvaluatorSkillsResponseBodySkills) *ListEvaluatorSkillsResponseBody {
	s.Skills = v
	return s
}

func (s *ListEvaluatorSkillsResponseBody) SetTotal(v int64) *ListEvaluatorSkillsResponseBody {
	s.Total = &v
	return s
}

func (s *ListEvaluatorSkillsResponseBody) Validate() error {
	if s.Skills != nil {
		for _, item := range s.Skills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEvaluatorSkillsResponseBodySkills struct {
	// The creation time. This field is declared as int64 in CloudSpec, but the backend currently returns the StarOps `createTime` string field.
	//
	// example:
	//
	// 1782816000
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The skill description.
	//
	// example:
	//
	// 读取链路上下文辅助评估
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name.
	//
	// example:
	//
	// Trace 上下文读取
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Indicates whether the skill is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The latest version. This field is declared in CloudSpec but is not currently populated in the backend response.
	//
	// example:
	//
	// 1782816000000
	LatestVersion *string `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
	// The skill name.
	//
	// example:
	//
	// trace_context_loader
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// The update time. This field is declared as int64 in CloudSpec, but the backend currently returns the StarOps `updateTime` string field.
	//
	// example:
	//
	// 1782816600
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListEvaluatorSkillsResponseBodySkills) String() string {
	return dara.Prettify(s)
}

func (s ListEvaluatorSkillsResponseBodySkills) GoString() string {
	return s.String()
}

func (s *ListEvaluatorSkillsResponseBodySkills) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *ListEvaluatorSkillsResponseBodySkills) GetDescription() *string {
	return s.Description
}

func (s *ListEvaluatorSkillsResponseBodySkills) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEvaluatorSkillsResponseBodySkills) GetEnable() *bool {
	return s.Enable
}

func (s *ListEvaluatorSkillsResponseBodySkills) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *ListEvaluatorSkillsResponseBodySkills) GetSkillName() *string {
	return s.SkillName
}

func (s *ListEvaluatorSkillsResponseBodySkills) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *ListEvaluatorSkillsResponseBodySkills) SetCreatedAt(v int64) *ListEvaluatorSkillsResponseBodySkills {
	s.CreatedAt = &v
	return s
}

func (s *ListEvaluatorSkillsResponseBodySkills) SetDescription(v string) *ListEvaluatorSkillsResponseBodySkills {
	s.Description = &v
	return s
}

func (s *ListEvaluatorSkillsResponseBodySkills) SetDisplayName(v string) *ListEvaluatorSkillsResponseBodySkills {
	s.DisplayName = &v
	return s
}

func (s *ListEvaluatorSkillsResponseBodySkills) SetEnable(v bool) *ListEvaluatorSkillsResponseBodySkills {
	s.Enable = &v
	return s
}

func (s *ListEvaluatorSkillsResponseBodySkills) SetLatestVersion(v string) *ListEvaluatorSkillsResponseBodySkills {
	s.LatestVersion = &v
	return s
}

func (s *ListEvaluatorSkillsResponseBodySkills) SetSkillName(v string) *ListEvaluatorSkillsResponseBodySkills {
	s.SkillName = &v
	return s
}

func (s *ListEvaluatorSkillsResponseBodySkills) SetUpdatedAt(v int64) *ListEvaluatorSkillsResponseBodySkills {
	s.UpdatedAt = &v
	return s
}

func (s *ListEvaluatorSkillsResponseBodySkills) Validate() error {
	return dara.Validate(s)
}
