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
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// AAAAAZ9FmxgN6wKfeK/GOKRnnjU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3F976EF8-C10A-57DC-917C-BB7BEB508FFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of skills.
	Skills []*ListPublicSkillsResponseBodySkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	// The total number of entries.
	//
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
	// The time when the skill was created.
	//
	// example:
	//
	// 2025-09-11T02:18:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The download URL of the skill package.
	//
	// example:
	//
	// https://testts-1.oss-cn-beijing.aliyuncs.com/app/yyb_9.1.1.zip
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// if can be null:
	// true
	Locales []*ListPublicSkillsResponseBodySkillsLocales `json:"Locales,omitempty" xml:"Locales,omitempty" type:"Repeated"`
	// The skill description.
	//
	// example:
	//
	// 11111
	SkillDescription *string `json:"SkillDescription,omitempty" xml:"SkillDescription,omitempty"`
	SkillDisplayName *string `json:"SkillDisplayName,omitempty" xml:"SkillDisplayName,omitempty"`
	// Skill ID
	//
	// example:
	//
	// af7e49d9-277f-454a-afc5-1513d41cac31
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The skill label set.
	SkillLabels []*string `json:"SkillLabels,omitempty" xml:"SkillLabels,omitempty" type:"Repeated"`
	// The skill name.
	//
	// example:
	//
	// ziptest
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The ID of the SkillSpace to which the skill belongs.
	//
	// example:
	//
	// ss-1111111
	SkillSpaceId *string `json:"SkillSpaceId,omitempty" xml:"SkillSpaceId,omitempty"`
	// The time when the skill was last updated.
	//
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

func (s *ListPublicSkillsResponseBodySkills) GetLocales() []*ListPublicSkillsResponseBodySkillsLocales {
	return s.Locales
}

func (s *ListPublicSkillsResponseBodySkills) GetSkillDescription() *string {
	return s.SkillDescription
}

func (s *ListPublicSkillsResponseBodySkills) GetSkillDisplayName() *string {
	return s.SkillDisplayName
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

func (s *ListPublicSkillsResponseBodySkills) SetLocales(v []*ListPublicSkillsResponseBodySkillsLocales) *ListPublicSkillsResponseBodySkills {
	s.Locales = v
	return s
}

func (s *ListPublicSkillsResponseBodySkills) SetSkillDescription(v string) *ListPublicSkillsResponseBodySkills {
	s.SkillDescription = &v
	return s
}

func (s *ListPublicSkillsResponseBodySkills) SetSkillDisplayName(v string) *ListPublicSkillsResponseBodySkills {
	s.SkillDisplayName = &v
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
	if s.Locales != nil {
		for _, item := range s.Locales {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPublicSkillsResponseBodySkillsLocales struct {
	EnValue       *string `json:"EnValue,omitempty" xml:"EnValue,omitempty"`
	OriginalValue *string `json:"OriginalValue,omitempty" xml:"OriginalValue,omitempty"`
	ZhValue       *string `json:"ZhValue,omitempty" xml:"ZhValue,omitempty"`
}

func (s ListPublicSkillsResponseBodySkillsLocales) String() string {
	return dara.Prettify(s)
}

func (s ListPublicSkillsResponseBodySkillsLocales) GoString() string {
	return s.String()
}

func (s *ListPublicSkillsResponseBodySkillsLocales) GetEnValue() *string {
	return s.EnValue
}

func (s *ListPublicSkillsResponseBodySkillsLocales) GetOriginalValue() *string {
	return s.OriginalValue
}

func (s *ListPublicSkillsResponseBodySkillsLocales) GetZhValue() *string {
	return s.ZhValue
}

func (s *ListPublicSkillsResponseBodySkillsLocales) SetEnValue(v string) *ListPublicSkillsResponseBodySkillsLocales {
	s.EnValue = &v
	return s
}

func (s *ListPublicSkillsResponseBodySkillsLocales) SetOriginalValue(v string) *ListPublicSkillsResponseBodySkillsLocales {
	s.OriginalValue = &v
	return s
}

func (s *ListPublicSkillsResponseBodySkillsLocales) SetZhValue(v string) *ListPublicSkillsResponseBodySkillsLocales {
	s.ZhValue = &v
	return s
}

func (s *ListPublicSkillsResponseBodySkillsLocales) Validate() error {
	return dara.Validate(s)
}
