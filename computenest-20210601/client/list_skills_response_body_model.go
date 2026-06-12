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
	SetRequestId(v string) *ListSkillsResponseBody
	GetRequestId() *string
	SetSkills(v []*ListSkillsResponseBodySkills) *ListSkillsResponseBody
	GetSkills() []*ListSkillsResponseBodySkills
	SetTotalCount(v int32) *ListSkillsResponseBody
	GetTotalCount() *int32
}

type ListSkillsResponseBody struct {
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results. This parameter is empty when all results have been returned.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of Skills.
	Skills []*ListSkillsResponseBodySkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	// The total number of entries that match the query.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillsResponseBody) GetSkills() []*ListSkillsResponseBodySkills {
	return s.Skills
}

func (s *ListSkillsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSkillsResponseBody) SetMaxResults(v int32) *ListSkillsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSkillsResponseBody) SetNextToken(v string) *ListSkillsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSkillsResponseBody) SetRequestId(v string) *ListSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillsResponseBody) SetSkills(v []*ListSkillsResponseBodySkills) *ListSkillsResponseBody {
	s.Skills = v
	return s
}

func (s *ListSkillsResponseBody) SetTotalCount(v int32) *ListSkillsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSkillsResponseBody) Validate() error {
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

type ListSkillsResponseBodySkills struct {
	// The creation time of the Skill.
	//
	// example:
	//
	// 2026-05-10T02:22:18Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The download URL of the Skill package. This parameter is returned only if `NeedDownloadUrl` is set to `true`.
	//
	// example:
	//
	// https://testts-1.oss-cn-beijing.aliyuncs.com/app/yyb_9.1.1.zip
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The Skill description.
	//
	// example:
	//
	// 1111
	SkillDescription *string `json:"SkillDescription,omitempty" xml:"SkillDescription,omitempty"`
	// The Skill ID.
	//
	// example:
	//
	// s-111111
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The Skill labels.
	SkillLabels []*string `json:"SkillLabels,omitempty" xml:"SkillLabels,omitempty" type:"Repeated"`
	// The Skill name.
	//
	// example:
	//
	// reimbursement-print
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The ID of the SkillSpace to which the Skill belongs.
	//
	// example:
	//
	// ss-11111
	SkillSpaceId *string `json:"SkillSpaceId,omitempty" xml:"SkillSpaceId,omitempty"`
	// The last update time of the Skill.
	//
	// example:
	//
	// 2025-11-03T22:58:52Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSkillsResponseBodySkills) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBodySkills) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBodySkills) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSkillsResponseBodySkills) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListSkillsResponseBodySkills) GetSkillDescription() *string {
	return s.SkillDescription
}

func (s *ListSkillsResponseBodySkills) GetSkillId() *string {
	return s.SkillId
}

func (s *ListSkillsResponseBodySkills) GetSkillLabels() []*string {
	return s.SkillLabels
}

func (s *ListSkillsResponseBodySkills) GetSkillName() *string {
	return s.SkillName
}

func (s *ListSkillsResponseBodySkills) GetSkillSpaceId() *string {
	return s.SkillSpaceId
}

func (s *ListSkillsResponseBodySkills) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListSkillsResponseBodySkills) SetCreateTime(v string) *ListSkillsResponseBodySkills {
	s.CreateTime = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetDownloadUrl(v string) *ListSkillsResponseBodySkills {
	s.DownloadUrl = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSkillDescription(v string) *ListSkillsResponseBodySkills {
	s.SkillDescription = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSkillId(v string) *ListSkillsResponseBodySkills {
	s.SkillId = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSkillLabels(v []*string) *ListSkillsResponseBodySkills {
	s.SkillLabels = v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSkillName(v string) *ListSkillsResponseBodySkills {
	s.SkillName = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSkillSpaceId(v string) *ListSkillsResponseBodySkills {
	s.SkillSpaceId = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetUpdateTime(v string) *ListSkillsResponseBodySkills {
	s.UpdateTime = &v
	return s
}

func (s *ListSkillsResponseBodySkills) Validate() error {
	return dara.Validate(s)
}
