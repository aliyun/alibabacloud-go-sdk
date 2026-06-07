// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListSkillsResponseBodyPagingInfo) *ListSkillsResponseBody
	GetPagingInfo() *ListSkillsResponseBodyPagingInfo
	SetRequestId(v string) *ListSkillsResponseBody
	GetRequestId() *string
}

type ListSkillsResponseBody struct {
	PagingInfo *ListSkillsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBody) GetPagingInfo() *ListSkillsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillsResponseBody) SetPagingInfo(v *ListSkillsResponseBodyPagingInfo) *ListSkillsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListSkillsResponseBody) SetRequestId(v string) *ListSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSkillsResponseBodyPagingInfo struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 5
	NextToken *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Skills    []*ListSkillsResponseBodyPagingInfoSkills `json:"Skills,omitempty" xml:"Skills,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSkillsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBodyPagingInfo) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSkillsResponseBodyPagingInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSkillsResponseBodyPagingInfo) GetSkills() []*ListSkillsResponseBodyPagingInfoSkills {
	return s.Skills
}

func (s *ListSkillsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSkillsResponseBodyPagingInfo) SetMaxResults(v int32) *ListSkillsResponseBodyPagingInfo {
	s.MaxResults = &v
	return s
}

func (s *ListSkillsResponseBodyPagingInfo) SetNextToken(v string) *ListSkillsResponseBodyPagingInfo {
	s.NextToken = &v
	return s
}

func (s *ListSkillsResponseBodyPagingInfo) SetSkills(v []*ListSkillsResponseBodyPagingInfoSkills) *ListSkillsResponseBodyPagingInfo {
	s.Skills = v
	return s
}

func (s *ListSkillsResponseBodyPagingInfo) SetTotalCount(v int32) *ListSkillsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListSkillsResponseBodyPagingInfo) Validate() error {
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

type ListSkillsResponseBodyPagingInfoSkills struct {
	// example:
	//
	// 123456
	CreatorId   *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 123456
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// my-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// TENANT
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListSkillsResponseBodyPagingInfoSkills) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBodyPagingInfoSkills) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBodyPagingInfoSkills) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListSkillsResponseBodyPagingInfoSkills) GetDescription() *string {
	return s.Description
}

func (s *ListSkillsResponseBodyPagingInfoSkills) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListSkillsResponseBodyPagingInfoSkills) GetModifierId() *string {
	return s.ModifierId
}

func (s *ListSkillsResponseBodyPagingInfoSkills) GetName() *string {
	return s.Name
}

func (s *ListSkillsResponseBodyPagingInfoSkills) GetVisibility() *string {
	return s.Visibility
}

func (s *ListSkillsResponseBodyPagingInfoSkills) SetCreatorId(v string) *ListSkillsResponseBodyPagingInfoSkills {
	s.CreatorId = &v
	return s
}

func (s *ListSkillsResponseBodyPagingInfoSkills) SetDescription(v string) *ListSkillsResponseBodyPagingInfoSkills {
	s.Description = &v
	return s
}

func (s *ListSkillsResponseBodyPagingInfoSkills) SetGmtCreateTime(v string) *ListSkillsResponseBodyPagingInfoSkills {
	s.GmtCreateTime = &v
	return s
}

func (s *ListSkillsResponseBodyPagingInfoSkills) SetModifierId(v string) *ListSkillsResponseBodyPagingInfoSkills {
	s.ModifierId = &v
	return s
}

func (s *ListSkillsResponseBodyPagingInfoSkills) SetName(v string) *ListSkillsResponseBodyPagingInfoSkills {
	s.Name = &v
	return s
}

func (s *ListSkillsResponseBodyPagingInfoSkills) SetVisibility(v string) *ListSkillsResponseBodyPagingInfoSkills {
	s.Visibility = &v
	return s
}

func (s *ListSkillsResponseBodyPagingInfoSkills) Validate() error {
	return dara.Validate(s)
}
