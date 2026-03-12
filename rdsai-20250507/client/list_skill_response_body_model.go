// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListSkillResponseBodyData) *ListSkillResponseBody
	GetData() []*ListSkillResponseBodyData
	SetPageNumber(v int64) *ListSkillResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSkillResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListSkillResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListSkillResponseBody
	GetTotalCount() *int32
}

type ListSkillResponseBody struct {
	// The list of skills.
	Data []*ListSkillResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned records.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillResponseBody) GetData() []*ListSkillResponseBodyData {
	return s.Data
}

func (s *ListSkillResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSkillResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSkillResponseBody) SetData(v []*ListSkillResponseBodyData) *ListSkillResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillResponseBody) SetPageNumber(v int64) *ListSkillResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSkillResponseBody) SetPageSize(v int64) *ListSkillResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSkillResponseBody) SetRequestId(v string) *ListSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillResponseBody) SetTotalCount(v int32) *ListSkillResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSkillResponseBody) Validate() error {
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

type ListSkillResponseBodyData struct {
	// The content of the skill.
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The creation time of the skill.
	//
	// example:
	//
	// 2026-02-04T21:14:45Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The list of database engines.
	Dbtypes []*string `json:"Dbtypes,omitempty" xml:"Dbtypes,omitempty" type:"Repeated"`
	// The description of the skill.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The unique identifier of the skill.
	//
	// example:
	//
	// 9a2ba261-7bb2-41a7-9c6e-1799fb5b****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the skill.
	//
	// example:
	//
	// sql-review
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the skill.
	//
	// example:
	//
	// system
	SkillType *string `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
	// The update time of the skill.
	//
	// example:
	//
	// 2026-02-04T21:14:45Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListSkillResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSkillResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillResponseBodyData) GetContent() map[string]interface{} {
	return s.Content
}

func (s *ListSkillResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListSkillResponseBodyData) GetDbtypes() []*string {
	return s.Dbtypes
}

func (s *ListSkillResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListSkillResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListSkillResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListSkillResponseBodyData) GetSkillType() *string {
	return s.SkillType
}

func (s *ListSkillResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListSkillResponseBodyData) SetContent(v map[string]interface{}) *ListSkillResponseBodyData {
	s.Content = v
	return s
}

func (s *ListSkillResponseBodyData) SetCreatedAt(v string) *ListSkillResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListSkillResponseBodyData) SetDbtypes(v []*string) *ListSkillResponseBodyData {
	s.Dbtypes = v
	return s
}

func (s *ListSkillResponseBodyData) SetDescription(v string) *ListSkillResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListSkillResponseBodyData) SetId(v string) *ListSkillResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListSkillResponseBodyData) SetName(v string) *ListSkillResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListSkillResponseBodyData) SetSkillType(v string) *ListSkillResponseBodyData {
	s.SkillType = &v
	return s
}

func (s *ListSkillResponseBodyData) SetUpdatedAt(v string) *ListSkillResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *ListSkillResponseBodyData) Validate() error {
	return dara.Validate(s)
}
