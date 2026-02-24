// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDigitalEmployeeSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDigitalEmployeeSkillsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDigitalEmployeeSkillsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDigitalEmployeeSkillsResponseBody
	GetRequestId() *string
	SetSkills(v []*ListDigitalEmployeeSkillsResponseBodySkills) *ListDigitalEmployeeSkillsResponseBody
	GetSkills() []*ListDigitalEmployeeSkillsResponseBodySkills
	SetTotal(v int64) *ListDigitalEmployeeSkillsResponseBody
	GetTotal() *int64
}

type ListDigitalEmployeeSkillsResponseBody struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// sjC5rekx93Ew7K7VcmI3wkBZBYQ-GphB2ilQu3zJCGxoZuicwyJznfo2riTjr-lq
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0CEC5375-XXXX-XXXX-XXXX-9A629907C1F0
	RequestId *string                                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Skills    []*ListDigitalEmployeeSkillsResponseBodySkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// example:
	//
	// 15
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListDigitalEmployeeSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeeSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeeSkillsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDigitalEmployeeSkillsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDigitalEmployeeSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDigitalEmployeeSkillsResponseBody) GetSkills() []*ListDigitalEmployeeSkillsResponseBodySkills {
	return s.Skills
}

func (s *ListDigitalEmployeeSkillsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListDigitalEmployeeSkillsResponseBody) SetMaxResults(v int32) *ListDigitalEmployeeSkillsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDigitalEmployeeSkillsResponseBody) SetNextToken(v string) *ListDigitalEmployeeSkillsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDigitalEmployeeSkillsResponseBody) SetRequestId(v string) *ListDigitalEmployeeSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDigitalEmployeeSkillsResponseBody) SetSkills(v []*ListDigitalEmployeeSkillsResponseBodySkills) *ListDigitalEmployeeSkillsResponseBody {
	s.Skills = v
	return s
}

func (s *ListDigitalEmployeeSkillsResponseBody) SetTotal(v int64) *ListDigitalEmployeeSkillsResponseBody {
	s.Total = &v
	return s
}

func (s *ListDigitalEmployeeSkillsResponseBody) Validate() error {
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

type ListDigitalEmployeeSkillsResponseBodySkills struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-02-06T14:09:11Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// test
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-02-06T14:09:11Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListDigitalEmployeeSkillsResponseBodySkills) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeeSkillsResponseBodySkills) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeeSkillsResponseBodySkills) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDigitalEmployeeSkillsResponseBodySkills) GetDescription() *string {
	return s.Description
}

func (s *ListDigitalEmployeeSkillsResponseBodySkills) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListDigitalEmployeeSkillsResponseBodySkills) GetEnable() *bool {
	return s.Enable
}

func (s *ListDigitalEmployeeSkillsResponseBodySkills) GetSkillName() *string {
	return s.SkillName
}

func (s *ListDigitalEmployeeSkillsResponseBodySkills) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListDigitalEmployeeSkillsResponseBodySkills) SetCreateTime(v string) *ListDigitalEmployeeSkillsResponseBodySkills {
	s.CreateTime = &v
	return s
}

func (s *ListDigitalEmployeeSkillsResponseBodySkills) SetDescription(v string) *ListDigitalEmployeeSkillsResponseBodySkills {
	s.Description = &v
	return s
}

func (s *ListDigitalEmployeeSkillsResponseBodySkills) SetDisplayName(v string) *ListDigitalEmployeeSkillsResponseBodySkills {
	s.DisplayName = &v
	return s
}

func (s *ListDigitalEmployeeSkillsResponseBodySkills) SetEnable(v bool) *ListDigitalEmployeeSkillsResponseBodySkills {
	s.Enable = &v
	return s
}

func (s *ListDigitalEmployeeSkillsResponseBodySkills) SetSkillName(v string) *ListDigitalEmployeeSkillsResponseBodySkills {
	s.SkillName = &v
	return s
}

func (s *ListDigitalEmployeeSkillsResponseBodySkills) SetUpdateTime(v string) *ListDigitalEmployeeSkillsResponseBodySkills {
	s.UpdateTime = &v
	return s
}

func (s *ListDigitalEmployeeSkillsResponseBodySkills) Validate() error {
	return dara.Validate(s)
}
