// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListPromptsResponseBodyData) *ListPromptsResponseBody
	GetData() *ListPromptsResponseBodyData
	SetRequestId(v string) *ListPromptsResponseBody
	GetRequestId() *string
}

type ListPromptsResponseBody struct {
	Data      *ListPromptsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPromptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPromptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPromptsResponseBody) GetData() *ListPromptsResponseBodyData {
	return s.Data
}

func (s *ListPromptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPromptsResponseBody) SetData(v *ListPromptsResponseBodyData) *ListPromptsResponseBody {
	s.Data = v
	return s
}

func (s *ListPromptsResponseBody) SetRequestId(v string) *ListPromptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPromptsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPromptsResponseBodyData struct {
	PageItems      []*ListPromptsResponseBodyDataPageItems `json:"PageItems,omitempty" xml:"PageItems,omitempty" type:"Repeated"`
	PageNumber     *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PagesAvailable *int32                                  `json:"PagesAvailable,omitempty" xml:"PagesAvailable,omitempty"`
	TotalCount     *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPromptsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPromptsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPromptsResponseBodyData) GetPageItems() []*ListPromptsResponseBodyDataPageItems {
	return s.PageItems
}

func (s *ListPromptsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPromptsResponseBodyData) GetPagesAvailable() *int32 {
	return s.PagesAvailable
}

func (s *ListPromptsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPromptsResponseBodyData) SetPageItems(v []*ListPromptsResponseBodyDataPageItems) *ListPromptsResponseBodyData {
	s.PageItems = v
	return s
}

func (s *ListPromptsResponseBodyData) SetPageNumber(v int32) *ListPromptsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPromptsResponseBodyData) SetPagesAvailable(v int32) *ListPromptsResponseBodyData {
	s.PagesAvailable = &v
	return s
}

func (s *ListPromptsResponseBodyData) SetTotalCount(v int32) *ListPromptsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListPromptsResponseBodyData) Validate() error {
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

type ListPromptsResponseBodyDataPageItems struct {
	BizTags          []*string          `json:"BizTags,omitempty" xml:"BizTags,omitempty" type:"Repeated"`
	Description      *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	EditingVersion   *string            `json:"EditingVersion,omitempty" xml:"EditingVersion,omitempty"`
	GmtModified      *int64             `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Labels           map[string]*string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	LatestVersion    *string            `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	OnlineCnt        *int32             `json:"OnlineCnt,omitempty" xml:"OnlineCnt,omitempty"`
	PromptKey        *string            `json:"PromptKey,omitempty" xml:"PromptKey,omitempty"`
	ReviewingVersion *string            `json:"ReviewingVersion,omitempty" xml:"ReviewingVersion,omitempty"`
	SchemaVersion    *int32             `json:"SchemaVersion,omitempty" xml:"SchemaVersion,omitempty"`
}

func (s ListPromptsResponseBodyDataPageItems) String() string {
	return dara.Prettify(s)
}

func (s ListPromptsResponseBodyDataPageItems) GoString() string {
	return s.String()
}

func (s *ListPromptsResponseBodyDataPageItems) GetBizTags() []*string {
	return s.BizTags
}

func (s *ListPromptsResponseBodyDataPageItems) GetDescription() *string {
	return s.Description
}

func (s *ListPromptsResponseBodyDataPageItems) GetEditingVersion() *string {
	return s.EditingVersion
}

func (s *ListPromptsResponseBodyDataPageItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListPromptsResponseBodyDataPageItems) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ListPromptsResponseBodyDataPageItems) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *ListPromptsResponseBodyDataPageItems) GetOnlineCnt() *int32 {
	return s.OnlineCnt
}

func (s *ListPromptsResponseBodyDataPageItems) GetPromptKey() *string {
	return s.PromptKey
}

func (s *ListPromptsResponseBodyDataPageItems) GetReviewingVersion() *string {
	return s.ReviewingVersion
}

func (s *ListPromptsResponseBodyDataPageItems) GetSchemaVersion() *int32 {
	return s.SchemaVersion
}

func (s *ListPromptsResponseBodyDataPageItems) SetBizTags(v []*string) *ListPromptsResponseBodyDataPageItems {
	s.BizTags = v
	return s
}

func (s *ListPromptsResponseBodyDataPageItems) SetDescription(v string) *ListPromptsResponseBodyDataPageItems {
	s.Description = &v
	return s
}

func (s *ListPromptsResponseBodyDataPageItems) SetEditingVersion(v string) *ListPromptsResponseBodyDataPageItems {
	s.EditingVersion = &v
	return s
}

func (s *ListPromptsResponseBodyDataPageItems) SetGmtModified(v int64) *ListPromptsResponseBodyDataPageItems {
	s.GmtModified = &v
	return s
}

func (s *ListPromptsResponseBodyDataPageItems) SetLabels(v map[string]*string) *ListPromptsResponseBodyDataPageItems {
	s.Labels = v
	return s
}

func (s *ListPromptsResponseBodyDataPageItems) SetLatestVersion(v string) *ListPromptsResponseBodyDataPageItems {
	s.LatestVersion = &v
	return s
}

func (s *ListPromptsResponseBodyDataPageItems) SetOnlineCnt(v int32) *ListPromptsResponseBodyDataPageItems {
	s.OnlineCnt = &v
	return s
}

func (s *ListPromptsResponseBodyDataPageItems) SetPromptKey(v string) *ListPromptsResponseBodyDataPageItems {
	s.PromptKey = &v
	return s
}

func (s *ListPromptsResponseBodyDataPageItems) SetReviewingVersion(v string) *ListPromptsResponseBodyDataPageItems {
	s.ReviewingVersion = &v
	return s
}

func (s *ListPromptsResponseBodyDataPageItems) SetSchemaVersion(v int32) *ListPromptsResponseBodyDataPageItems {
	s.SchemaVersion = &v
	return s
}

func (s *ListPromptsResponseBodyDataPageItems) Validate() error {
	return dara.Validate(s)
}
