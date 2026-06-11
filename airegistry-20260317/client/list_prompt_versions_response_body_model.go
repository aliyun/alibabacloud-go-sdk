// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromptVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListPromptVersionsResponseBodyData) *ListPromptVersionsResponseBody
	GetData() *ListPromptVersionsResponseBodyData
	SetRequestId(v string) *ListPromptVersionsResponseBody
	GetRequestId() *string
}

type ListPromptVersionsResponseBody struct {
	Data      *ListPromptVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPromptVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPromptVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPromptVersionsResponseBody) GetData() *ListPromptVersionsResponseBodyData {
	return s.Data
}

func (s *ListPromptVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPromptVersionsResponseBody) SetData(v *ListPromptVersionsResponseBodyData) *ListPromptVersionsResponseBody {
	s.Data = v
	return s
}

func (s *ListPromptVersionsResponseBody) SetRequestId(v string) *ListPromptVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPromptVersionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPromptVersionsResponseBodyData struct {
	PageItems      []*ListPromptVersionsResponseBodyDataPageItems `json:"PageItems,omitempty" xml:"PageItems,omitempty" type:"Repeated"`
	PageNumber     *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PagesAvailable *int32                                         `json:"PagesAvailable,omitempty" xml:"PagesAvailable,omitempty"`
	TotalCount     *int32                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPromptVersionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPromptVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPromptVersionsResponseBodyData) GetPageItems() []*ListPromptVersionsResponseBodyDataPageItems {
	return s.PageItems
}

func (s *ListPromptVersionsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPromptVersionsResponseBodyData) GetPagesAvailable() *int32 {
	return s.PagesAvailable
}

func (s *ListPromptVersionsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPromptVersionsResponseBodyData) SetPageItems(v []*ListPromptVersionsResponseBodyDataPageItems) *ListPromptVersionsResponseBodyData {
	s.PageItems = v
	return s
}

func (s *ListPromptVersionsResponseBodyData) SetPageNumber(v int32) *ListPromptVersionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPromptVersionsResponseBodyData) SetPagesAvailable(v int32) *ListPromptVersionsResponseBodyData {
	s.PagesAvailable = &v
	return s
}

func (s *ListPromptVersionsResponseBodyData) SetTotalCount(v int32) *ListPromptVersionsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListPromptVersionsResponseBodyData) Validate() error {
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

type ListPromptVersionsResponseBodyDataPageItems struct {
	CommitMsg   *string `json:"CommitMsg,omitempty" xml:"CommitMsg,omitempty"`
	GmtModified *int64  `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	PromptKey   *string `json:"PromptKey,omitempty" xml:"PromptKey,omitempty"`
	SrcUser     *string `json:"SrcUser,omitempty" xml:"SrcUser,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Version     *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListPromptVersionsResponseBodyDataPageItems) String() string {
	return dara.Prettify(s)
}

func (s ListPromptVersionsResponseBodyDataPageItems) GoString() string {
	return s.String()
}

func (s *ListPromptVersionsResponseBodyDataPageItems) GetCommitMsg() *string {
	return s.CommitMsg
}

func (s *ListPromptVersionsResponseBodyDataPageItems) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListPromptVersionsResponseBodyDataPageItems) GetPromptKey() *string {
	return s.PromptKey
}

func (s *ListPromptVersionsResponseBodyDataPageItems) GetSrcUser() *string {
	return s.SrcUser
}

func (s *ListPromptVersionsResponseBodyDataPageItems) GetStatus() *string {
	return s.Status
}

func (s *ListPromptVersionsResponseBodyDataPageItems) GetVersion() *string {
	return s.Version
}

func (s *ListPromptVersionsResponseBodyDataPageItems) SetCommitMsg(v string) *ListPromptVersionsResponseBodyDataPageItems {
	s.CommitMsg = &v
	return s
}

func (s *ListPromptVersionsResponseBodyDataPageItems) SetGmtModified(v int64) *ListPromptVersionsResponseBodyDataPageItems {
	s.GmtModified = &v
	return s
}

func (s *ListPromptVersionsResponseBodyDataPageItems) SetPromptKey(v string) *ListPromptVersionsResponseBodyDataPageItems {
	s.PromptKey = &v
	return s
}

func (s *ListPromptVersionsResponseBodyDataPageItems) SetSrcUser(v string) *ListPromptVersionsResponseBodyDataPageItems {
	s.SrcUser = &v
	return s
}

func (s *ListPromptVersionsResponseBodyDataPageItems) SetStatus(v string) *ListPromptVersionsResponseBodyDataPageItems {
	s.Status = &v
	return s
}

func (s *ListPromptVersionsResponseBodyDataPageItems) SetVersion(v string) *ListPromptVersionsResponseBodyDataPageItems {
	s.Version = &v
	return s
}

func (s *ListPromptVersionsResponseBodyDataPageItems) Validate() error {
	return dara.Validate(s)
}
