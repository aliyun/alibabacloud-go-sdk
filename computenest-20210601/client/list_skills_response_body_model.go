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
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of Skills.
	Skills []*ListSkillsResponseBodySkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	// The total number of entries.
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
	// The time when the Skill was created.
	//
	// example:
	//
	// 2026-05-10T02:22:18Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The download URL for the Skill package.
	//
	// example:
	//
	// https://embedding-pic.oss-cn-beijing-internal.aliyuncs.com/skill-creator.zip
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// if can be null:
	// true
	Locales []*ListSkillsResponseBodySkillsLocales `json:"Locales,omitempty" xml:"Locales,omitempty" type:"Repeated"`
	// The set of connection types that the Skill depends on.
	RequiredConnections []*string `json:"RequiredConnections,omitempty" xml:"RequiredConnections,omitempty" type:"Repeated"`
	// The security scan status of the source file.
	SecurityScanStatus *string `json:"SecurityScanStatus,omitempty" xml:"SecurityScanStatus,omitempty"`
	// The Skill description.
	//
	// example:
	//
	// Create new skills, modify and improve existing skills, and measure skill performance.
	SkillDescription *string `json:"SkillDescription,omitempty" xml:"SkillDescription,omitempty"`
	// The Skill display name.
	SkillDisplayName *string `json:"SkillDisplayName,omitempty" xml:"SkillDisplayName,omitempty"`
	// Skill ID
	//
	// example:
	//
	// s-xxxxx
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The set of Skill labels.
	SkillLabels []*string `json:"SkillLabels,omitempty" xml:"SkillLabels,omitempty" type:"Repeated"`
	// The Skill name.
	//
	// example:
	//
	// skill-creator
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The ID of the SkillSpace to which the Skill belongs.
	//
	// example:
	//
	// ss-xxxxx
	SkillSpaceId *string `json:"SkillSpaceId,omitempty" xml:"SkillSpaceId,omitempty"`
	// The time when the Skill was last updated.
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

func (s *ListSkillsResponseBodySkills) GetLocales() []*ListSkillsResponseBodySkillsLocales {
	return s.Locales
}

func (s *ListSkillsResponseBodySkills) GetRequiredConnections() []*string {
	return s.RequiredConnections
}

func (s *ListSkillsResponseBodySkills) GetSecurityScanStatus() *string {
	return s.SecurityScanStatus
}

func (s *ListSkillsResponseBodySkills) GetSkillDescription() *string {
	return s.SkillDescription
}

func (s *ListSkillsResponseBodySkills) GetSkillDisplayName() *string {
	return s.SkillDisplayName
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

func (s *ListSkillsResponseBodySkills) SetLocales(v []*ListSkillsResponseBodySkillsLocales) *ListSkillsResponseBodySkills {
	s.Locales = v
	return s
}

func (s *ListSkillsResponseBodySkills) SetRequiredConnections(v []*string) *ListSkillsResponseBodySkills {
	s.RequiredConnections = v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSecurityScanStatus(v string) *ListSkillsResponseBodySkills {
	s.SecurityScanStatus = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSkillDescription(v string) *ListSkillsResponseBodySkills {
	s.SkillDescription = &v
	return s
}

func (s *ListSkillsResponseBodySkills) SetSkillDisplayName(v string) *ListSkillsResponseBodySkills {
	s.SkillDisplayName = &v
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

type ListSkillsResponseBodySkillsLocales struct {
	EnValue       *string `json:"EnValue,omitempty" xml:"EnValue,omitempty"`
	OriginalValue *string `json:"OriginalValue,omitempty" xml:"OriginalValue,omitempty"`
	ZhValue       *string `json:"ZhValue,omitempty" xml:"ZhValue,omitempty"`
}

func (s ListSkillsResponseBodySkillsLocales) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBodySkillsLocales) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBodySkillsLocales) GetEnValue() *string {
	return s.EnValue
}

func (s *ListSkillsResponseBodySkillsLocales) GetOriginalValue() *string {
	return s.OriginalValue
}

func (s *ListSkillsResponseBodySkillsLocales) GetZhValue() *string {
	return s.ZhValue
}

func (s *ListSkillsResponseBodySkillsLocales) SetEnValue(v string) *ListSkillsResponseBodySkillsLocales {
	s.EnValue = &v
	return s
}

func (s *ListSkillsResponseBodySkillsLocales) SetOriginalValue(v string) *ListSkillsResponseBodySkillsLocales {
	s.OriginalValue = &v
	return s
}

func (s *ListSkillsResponseBodySkillsLocales) SetZhValue(v string) *ListSkillsResponseBodySkillsLocales {
	s.ZhValue = &v
	return s
}

func (s *ListSkillsResponseBodySkillsLocales) Validate() error {
	return dara.Validate(s)
}
