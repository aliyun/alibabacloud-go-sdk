// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySkillGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QuerySkillGroupsResponseBody
	GetCurrentPage() *int32
	SetData(v []*QuerySkillGroupsResponseBodyData) *QuerySkillGroupsResponseBody
	GetData() []*QuerySkillGroupsResponseBodyData
	SetOnePageSize(v int32) *QuerySkillGroupsResponseBody
	GetOnePageSize() *int32
	SetRequestId(v string) *QuerySkillGroupsResponseBody
	GetRequestId() *string
	SetTotalPage(v int32) *QuerySkillGroupsResponseBody
	GetTotalPage() *int32
	SetTotalResults(v int32) *QuerySkillGroupsResponseBody
	GetTotalResults() *int32
}

type QuerySkillGroupsResponseBody struct {
	// Current page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Data.
	Data []*QuerySkillGroupsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Number of items per page.
	//
	// example:
	//
	// 20
	OnePageSize *int32 `json:"OnePageSize,omitempty" xml:"OnePageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 4
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 76
	TotalResults *int32 `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
}

func (s QuerySkillGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySkillGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QuerySkillGroupsResponseBody) GetData() []*QuerySkillGroupsResponseBodyData {
	return s.Data
}

func (s *QuerySkillGroupsResponseBody) GetOnePageSize() *int32 {
	return s.OnePageSize
}

func (s *QuerySkillGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySkillGroupsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *QuerySkillGroupsResponseBody) GetTotalResults() *int32 {
	return s.TotalResults
}

func (s *QuerySkillGroupsResponseBody) SetCurrentPage(v int32) *QuerySkillGroupsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetData(v []*QuerySkillGroupsResponseBodyData) *QuerySkillGroupsResponseBody {
	s.Data = v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetOnePageSize(v int32) *QuerySkillGroupsResponseBody {
	s.OnePageSize = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetRequestId(v string) *QuerySkillGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetTotalPage(v int32) *QuerySkillGroupsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) SetTotalResults(v int32) *QuerySkillGroupsResponseBody {
	s.TotalResults = &v
	return s
}

func (s *QuerySkillGroupsResponseBody) Validate() error {
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

type QuerySkillGroupsResponseBodyData struct {
	// Skill group channel type.
	//
	// example:
	//
	// 2
	ChannelType *int32 `json:"ChannelType,omitempty" xml:"ChannelType,omitempty"`
	// Skill group description.
	//
	// example:
	//
	// 自动化技能组
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Skill group display name.
	//
	// example:
	//
	// 自动化技能组
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Skill group ID.
	//
	// example:
	//
	// 123456
	SkillGroupId *int64 `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// Skill group name.
	//
	// example:
	//
	// 自动化技能组
	SkillGroupName *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s QuerySkillGroupsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySkillGroupsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySkillGroupsResponseBodyData) GetChannelType() *int32 {
	return s.ChannelType
}

func (s *QuerySkillGroupsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *QuerySkillGroupsResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *QuerySkillGroupsResponseBodyData) GetSkillGroupId() *int64 {
	return s.SkillGroupId
}

func (s *QuerySkillGroupsResponseBodyData) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *QuerySkillGroupsResponseBodyData) SetChannelType(v int32) *QuerySkillGroupsResponseBodyData {
	s.ChannelType = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetDescription(v string) *QuerySkillGroupsResponseBodyData {
	s.Description = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetDisplayName(v string) *QuerySkillGroupsResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetSkillGroupId(v int64) *QuerySkillGroupsResponseBodyData {
	s.SkillGroupId = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) SetSkillGroupName(v string) *QuerySkillGroupsResponseBodyData {
	s.SkillGroupName = &v
	return s
}

func (s *QuerySkillGroupsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
