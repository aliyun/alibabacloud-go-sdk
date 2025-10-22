// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntentionNoteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *ListIntentionNoteResponseBody
	GetCurrentPageNum() *int32
	SetData(v []*ListIntentionNoteResponseBodyData) *ListIntentionNoteResponseBody
	GetData() []*ListIntentionNoteResponseBodyData
	SetPageSize(v int32) *ListIntentionNoteResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListIntentionNoteResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *ListIntentionNoteResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *ListIntentionNoteResponseBody
	GetTotalPageNum() *int32
}

type ListIntentionNoteResponseBody struct {
	// example:
	//
	// 0
	CurrentPageNum *int32                               `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListIntentionNoteResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 6A603AA0-73BA-52B3-AC7D-0F846ECF7A9D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListIntentionNoteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionNoteResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntentionNoteResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListIntentionNoteResponseBody) GetData() []*ListIntentionNoteResponseBodyData {
	return s.Data
}

func (s *ListIntentionNoteResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIntentionNoteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntentionNoteResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListIntentionNoteResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListIntentionNoteResponseBody) SetCurrentPageNum(v int32) *ListIntentionNoteResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListIntentionNoteResponseBody) SetData(v []*ListIntentionNoteResponseBodyData) *ListIntentionNoteResponseBody {
	s.Data = v
	return s
}

func (s *ListIntentionNoteResponseBody) SetPageSize(v int32) *ListIntentionNoteResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListIntentionNoteResponseBody) SetRequestId(v string) *ListIntentionNoteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntentionNoteResponseBody) SetTotalItemNum(v int32) *ListIntentionNoteResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListIntentionNoteResponseBody) SetTotalPageNum(v int32) *ListIntentionNoteResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *ListIntentionNoteResponseBody) Validate() error {
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

type ListIntentionNoteResponseBodyData struct {
	// example:
	//
	// 2022-01-25 10:21:38
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// I20210420142416000001
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	Note           *string `json:"Note,omitempty" xml:"Note,omitempty"`
	// example:
	//
	// 1
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// 1
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListIntentionNoteResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionNoteResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIntentionNoteResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListIntentionNoteResponseBodyData) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *ListIntentionNoteResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *ListIntentionNoteResponseBodyData) GetSource() *int32 {
	return s.Source
}

func (s *ListIntentionNoteResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListIntentionNoteResponseBodyData) SetCreateTime(v string) *ListIntentionNoteResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListIntentionNoteResponseBodyData) SetIntentionBizId(v string) *ListIntentionNoteResponseBodyData {
	s.IntentionBizId = &v
	return s
}

func (s *ListIntentionNoteResponseBodyData) SetNote(v string) *ListIntentionNoteResponseBodyData {
	s.Note = &v
	return s
}

func (s *ListIntentionNoteResponseBodyData) SetSource(v int32) *ListIntentionNoteResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListIntentionNoteResponseBodyData) SetType(v string) *ListIntentionNoteResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListIntentionNoteResponseBodyData) Validate() error {
	return dara.Validate(s)
}
