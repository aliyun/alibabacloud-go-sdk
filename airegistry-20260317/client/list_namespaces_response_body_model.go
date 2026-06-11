// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListNamespacesResponseBodyData) *ListNamespacesResponseBody
	GetData() *ListNamespacesResponseBodyData
	SetRequestId(v string) *ListNamespacesResponseBody
	GetRequestId() *string
}

type ListNamespacesResponseBody struct {
	Data      *ListNamespacesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNamespacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBody) GetData() *ListNamespacesResponseBodyData {
	return s.Data
}

func (s *ListNamespacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNamespacesResponseBody) SetData(v *ListNamespacesResponseBodyData) *ListNamespacesResponseBody {
	s.Data = v
	return s
}

func (s *ListNamespacesResponseBody) SetRequestId(v string) *ListNamespacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNamespacesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListNamespacesResponseBodyData struct {
	Items      []*ListNamespacesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNamespacesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyData) GetItems() []*ListNamespacesResponseBodyDataItems {
	return s.Items
}

func (s *ListNamespacesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListNamespacesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListNamespacesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNamespacesResponseBodyData) SetItems(v []*ListNamespacesResponseBodyDataItems) *ListNamespacesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListNamespacesResponseBodyData) SetPageNumber(v int32) *ListNamespacesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListNamespacesResponseBodyData) SetPageSize(v int32) *ListNamespacesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListNamespacesResponseBodyData) SetTotalCount(v int32) *ListNamespacesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListNamespacesResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNamespacesResponseBodyDataItems struct {
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	PromptCount *int32  `json:"PromptCount,omitempty" xml:"PromptCount,omitempty"`
	SkillCount  *int32  `json:"SkillCount,omitempty" xml:"SkillCount,omitempty"`
	Source      *string `json:"Source,omitempty" xml:"Source,omitempty"`
	SourceIndex *int32  `json:"SourceIndex,omitempty" xml:"SourceIndex,omitempty"`
	Tags        *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListNamespacesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListNamespacesResponseBodyDataItems) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListNamespacesResponseBodyDataItems) GetDescription() *string {
	return s.Description
}

func (s *ListNamespacesResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListNamespacesResponseBodyDataItems) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListNamespacesResponseBodyDataItems) GetPromptCount() *int32 {
	return s.PromptCount
}

func (s *ListNamespacesResponseBodyDataItems) GetSkillCount() *int32 {
	return s.SkillCount
}

func (s *ListNamespacesResponseBodyDataItems) GetSource() *string {
	return s.Source
}

func (s *ListNamespacesResponseBodyDataItems) GetSourceIndex() *int32 {
	return s.SourceIndex
}

func (s *ListNamespacesResponseBodyDataItems) GetTags() *string {
	return s.Tags
}

func (s *ListNamespacesResponseBodyDataItems) SetCreatedTime(v string) *ListNamespacesResponseBodyDataItems {
	s.CreatedTime = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetDescription(v string) *ListNamespacesResponseBodyDataItems {
	s.Description = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetName(v string) *ListNamespacesResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetNamespaceId(v string) *ListNamespacesResponseBodyDataItems {
	s.NamespaceId = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetPromptCount(v int32) *ListNamespacesResponseBodyDataItems {
	s.PromptCount = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetSkillCount(v int32) *ListNamespacesResponseBodyDataItems {
	s.SkillCount = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetSource(v string) *ListNamespacesResponseBodyDataItems {
	s.Source = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetSourceIndex(v int32) *ListNamespacesResponseBodyDataItems {
	s.SourceIndex = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) SetTags(v string) *ListNamespacesResponseBodyDataItems {
	s.Tags = &v
	return s
}

func (s *ListNamespacesResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
