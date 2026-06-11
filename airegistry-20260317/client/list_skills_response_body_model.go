// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListSkillsResponseBodyData) *ListSkillsResponseBody
	GetData() *ListSkillsResponseBodyData
	SetRequestId(v string) *ListSkillsResponseBody
	GetRequestId() *string
}

type ListSkillsResponseBody struct {
	Data      *ListSkillsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBody) GetData() *ListSkillsResponseBodyData {
	return s.Data
}

func (s *ListSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillsResponseBody) SetData(v *ListSkillsResponseBodyData) *ListSkillsResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillsResponseBody) SetRequestId(v string) *ListSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSkillsResponseBodyData struct {
	PageItems      []*ListSkillsResponseBodyDataPageItems `json:"PageItems,omitempty" xml:"PageItems,omitempty" type:"Repeated"`
	PageNumber     *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PagesAvailable *int32                                 `json:"PagesAvailable,omitempty" xml:"PagesAvailable,omitempty"`
	TotalCount     *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSkillsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBodyData) GetPageItems() []*ListSkillsResponseBodyDataPageItems {
	return s.PageItems
}

func (s *ListSkillsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSkillsResponseBodyData) GetPagesAvailable() *int32 {
	return s.PagesAvailable
}

func (s *ListSkillsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSkillsResponseBodyData) SetPageItems(v []*ListSkillsResponseBodyDataPageItems) *ListSkillsResponseBodyData {
	s.PageItems = v
	return s
}

func (s *ListSkillsResponseBodyData) SetPageNumber(v int32) *ListSkillsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSkillsResponseBodyData) SetPagesAvailable(v int32) *ListSkillsResponseBodyData {
	s.PagesAvailable = &v
	return s
}

func (s *ListSkillsResponseBodyData) SetTotalCount(v int32) *ListSkillsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSkillsResponseBodyData) Validate() error {
	if s.PageItems != nil {
		for _, item := range s.PageItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSkillsResponseBodyDataPageItems struct {
	BizTags          *string            `json:"BizTags,omitempty" xml:"BizTags,omitempty"`
	Description      *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	DownloadCount    *int64             `json:"DownloadCount,omitempty" xml:"DownloadCount,omitempty"`
	EditingVersion   *string            `json:"EditingVersion,omitempty" xml:"EditingVersion,omitempty"`
	Enable           *bool              `json:"Enable,omitempty" xml:"Enable,omitempty"`
	From             *string            `json:"From,omitempty" xml:"From,omitempty"`
	Labels           map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Name             *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	NamespaceId      *string            `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	OnlineCnt        *int32             `json:"OnlineCnt,omitempty" xml:"OnlineCnt,omitempty"`
	Owner            *string            `json:"Owner,omitempty" xml:"Owner,omitempty"`
	ReviewingVersion *string            `json:"ReviewingVersion,omitempty" xml:"ReviewingVersion,omitempty"`
	Scope            *string            `json:"Scope,omitempty" xml:"Scope,omitempty"`
	UpdateTime       *int64             `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Writeable        *bool              `json:"Writeable,omitempty" xml:"Writeable,omitempty"`
}

func (s ListSkillsResponseBodyDataPageItems) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBodyDataPageItems) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBodyDataPageItems) GetBizTags() *string {
	return s.BizTags
}

func (s *ListSkillsResponseBodyDataPageItems) GetDescription() *string {
	return s.Description
}

func (s *ListSkillsResponseBodyDataPageItems) GetDownloadCount() *int64 {
	return s.DownloadCount
}

func (s *ListSkillsResponseBodyDataPageItems) GetEditingVersion() *string {
	return s.EditingVersion
}

func (s *ListSkillsResponseBodyDataPageItems) GetEnable() *bool {
	return s.Enable
}

func (s *ListSkillsResponseBodyDataPageItems) GetFrom() *string {
	return s.From
}

func (s *ListSkillsResponseBodyDataPageItems) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ListSkillsResponseBodyDataPageItems) GetName() *string {
	return s.Name
}

func (s *ListSkillsResponseBodyDataPageItems) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListSkillsResponseBodyDataPageItems) GetOnlineCnt() *int32 {
	return s.OnlineCnt
}

func (s *ListSkillsResponseBodyDataPageItems) GetOwner() *string {
	return s.Owner
}

func (s *ListSkillsResponseBodyDataPageItems) GetReviewingVersion() *string {
	return s.ReviewingVersion
}

func (s *ListSkillsResponseBodyDataPageItems) GetScope() *string {
	return s.Scope
}

func (s *ListSkillsResponseBodyDataPageItems) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListSkillsResponseBodyDataPageItems) GetWriteable() *bool {
	return s.Writeable
}

func (s *ListSkillsResponseBodyDataPageItems) SetBizTags(v string) *ListSkillsResponseBodyDataPageItems {
	s.BizTags = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetDescription(v string) *ListSkillsResponseBodyDataPageItems {
	s.Description = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetDownloadCount(v int64) *ListSkillsResponseBodyDataPageItems {
	s.DownloadCount = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetEditingVersion(v string) *ListSkillsResponseBodyDataPageItems {
	s.EditingVersion = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetEnable(v bool) *ListSkillsResponseBodyDataPageItems {
	s.Enable = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetFrom(v string) *ListSkillsResponseBodyDataPageItems {
	s.From = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetLabels(v map[string]*string) *ListSkillsResponseBodyDataPageItems {
	s.Labels = v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetName(v string) *ListSkillsResponseBodyDataPageItems {
	s.Name = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetNamespaceId(v string) *ListSkillsResponseBodyDataPageItems {
	s.NamespaceId = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetOnlineCnt(v int32) *ListSkillsResponseBodyDataPageItems {
	s.OnlineCnt = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetOwner(v string) *ListSkillsResponseBodyDataPageItems {
	s.Owner = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetReviewingVersion(v string) *ListSkillsResponseBodyDataPageItems {
	s.ReviewingVersion = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetScope(v string) *ListSkillsResponseBodyDataPageItems {
	s.Scope = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetUpdateTime(v int64) *ListSkillsResponseBodyDataPageItems {
	s.UpdateTime = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetWriteable(v bool) *ListSkillsResponseBodyDataPageItems {
	s.Writeable = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) Validate() error {
	return dara.Validate(s)
}
