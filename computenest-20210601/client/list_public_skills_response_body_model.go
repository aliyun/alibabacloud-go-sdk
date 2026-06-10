// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPublicSkillsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListPublicSkillsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPublicSkillsResponseBody
	GetRequestId() *string
	SetSkills(v []*ListPublicSkillsResponseBodySkills) *ListPublicSkillsResponseBody
	GetSkills() []*ListPublicSkillsResponseBodySkills
	SetTotalCount(v int32) *ListPublicSkillsResponseBody
	GetTotalCount() *int32
}

type ListPublicSkillsResponseBody struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAZ9FmxgN6wKfeK/GOKRnnjU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3F976EF8-C10A-57DC-917C-BB7BEB508FFB
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Skills    []*ListPublicSkillsResponseBodySkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublicSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublicSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicSkillsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPublicSkillsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPublicSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublicSkillsResponseBody) GetSkills() []*ListPublicSkillsResponseBodySkills {
	return s.Skills
}

func (s *ListPublicSkillsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPublicSkillsResponseBody) SetMaxResults(v int32) *ListPublicSkillsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPublicSkillsResponseBody) SetNextToken(v string) *ListPublicSkillsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPublicSkillsResponseBody) SetRequestId(v string) *ListPublicSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicSkillsResponseBody) SetSkills(v []*ListPublicSkillsResponseBodySkills) *ListPublicSkillsResponseBody {
	s.Skills = v
	return s
}

func (s *ListPublicSkillsResponseBody) SetTotalCount(v int32) *ListPublicSkillsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublicSkillsResponseBody) Validate() error {
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

type ListPublicSkillsResponseBodySkills struct {
	// example:
	//
	// 2025-09-11T02:18:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// https://testts-1.oss-cn-beijing.aliyuncs.com/app/yyb_9.1.1.zip
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// example:
	//
	// 11111
	SkillDescription *string `json:"SkillDescription,omitempty" xml:"SkillDescription,omitempty"`
	// Skill ID
	//
	// example:
	//
	// af7e49d9-277f-454a-afc5-1513d41cac31
	SkillId     *string   `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	SkillLabels []*string `json:"SkillLabels,omitempty" xml:"SkillLabels,omitempty" type:"Repeated"`
	// example:
	//
	// ziptest
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// example:
	//
	// ss-1111111
	SkillSpaceId *string `json:"SkillSpaceId,omitempty" xml:"SkillSpaceId,omitempty"`
	// example:
	//
	// 2025-11-03T22:58:52Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListPublicSkillsResponseBodySkills) String() string {
	return dara.Prettify(s)
}

func (s ListPublicSkillsResponseBodySkills) GoString() string {
	return s.String()
}

func (s *ListPublicSkillsResponseBodySkills) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPublicSkillsResponseBodySkills) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListPublicSkillsResponseBodySkills) GetSkillDescription() *string {
	return s.SkillDescription
}

func (s *ListPublicSkillsResponseBodySkills) GetSkillId() *string {
	return s.SkillId
}

func (s *ListPublicSkillsResponseBodySkills) GetSkillLabels() []*string {
	return s.SkillLabels
}

func (s *ListPublicSkillsResponseBodySkills) GetSkillName() *string {
	return s.SkillName
}

func (s *ListPublicSkillsResponseBodySkills) GetSkillSpaceId() *string {
	return s.SkillSpaceId
}

func (s *ListPublicSkillsResponseBodySkills) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListPublicSkillsResponseBodySkills) SetCreateTime(v string) *ListPublicSkillsResponseBodySkills {
	s.CreateTime = &v
	return s
}

func (s *ListPublicSkillsResponseBodySkills) SetDownloadUrl(v string) *ListPublicSkillsResponseBodySkills {
	s.DownloadUrl = &v
	return s
}

func (s *ListPublicSkillsResponseBodySkills) SetSkillDescription(v string) *ListPublicSkillsResponseBodySkills {
	s.SkillDescription = &v
	return s
}

func (s *ListPublicSkillsResponseBodySkills) SetSkillId(v string) *ListPublicSkillsResponseBodySkills {
	s.SkillId = &v
	return s
}

func (s *ListPublicSkillsResponseBodySkills) SetSkillLabels(v []*string) *ListPublicSkillsResponseBodySkills {
	s.SkillLabels = v
	return s
}

func (s *ListPublicSkillsResponseBodySkills) SetSkillName(v string) *ListPublicSkillsResponseBodySkills {
	s.SkillName = &v
	return s
}

func (s *ListPublicSkillsResponseBodySkills) SetSkillSpaceId(v string) *ListPublicSkillsResponseBodySkills {
	s.SkillSpaceId = &v
	return s
}

func (s *ListPublicSkillsResponseBodySkills) SetUpdateTime(v string) *ListPublicSkillsResponseBodySkills {
	s.UpdateTime = &v
	return s
}

func (s *ListPublicSkillsResponseBodySkills) Validate() error {
	return dara.Validate(s)
}
