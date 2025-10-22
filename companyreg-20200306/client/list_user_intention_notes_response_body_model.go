// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserIntentionNotesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListUserIntentionNotesResponseBodyData) *ListUserIntentionNotesResponseBody
	GetData() []*ListUserIntentionNotesResponseBodyData
	SetPageNum(v int32) *ListUserIntentionNotesResponseBody
	GetPageNum() *int32
	SetPageSize(v int32) *ListUserIntentionNotesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUserIntentionNotesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUserIntentionNotesResponseBody
	GetSuccess() *bool
	SetTotalItemNum(v int32) *ListUserIntentionNotesResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *ListUserIntentionNotesResponseBody
	GetTotalPageNum() *int32
}

type ListUserIntentionNotesResponseBody struct {
	Data []*ListUserIntentionNotesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 5D8BD6E8-28D9-5451-BBA1-3D3DCA6971F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 8
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListUserIntentionNotesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserIntentionNotesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserIntentionNotesResponseBody) GetData() []*ListUserIntentionNotesResponseBodyData {
	return s.Data
}

func (s *ListUserIntentionNotesResponseBody) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListUserIntentionNotesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserIntentionNotesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserIntentionNotesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserIntentionNotesResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListUserIntentionNotesResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListUserIntentionNotesResponseBody) SetData(v []*ListUserIntentionNotesResponseBodyData) *ListUserIntentionNotesResponseBody {
	s.Data = v
	return s
}

func (s *ListUserIntentionNotesResponseBody) SetPageNum(v int32) *ListUserIntentionNotesResponseBody {
	s.PageNum = &v
	return s
}

func (s *ListUserIntentionNotesResponseBody) SetPageSize(v int32) *ListUserIntentionNotesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserIntentionNotesResponseBody) SetRequestId(v string) *ListUserIntentionNotesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserIntentionNotesResponseBody) SetSuccess(v bool) *ListUserIntentionNotesResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserIntentionNotesResponseBody) SetTotalItemNum(v int32) *ListUserIntentionNotesResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListUserIntentionNotesResponseBody) SetTotalPageNum(v int32) *ListUserIntentionNotesResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *ListUserIntentionNotesResponseBody) Validate() error {
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

type ListUserIntentionNotesResponseBodyData struct {
	// example:
	//
	// 2022-01-25 10:21:38
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Note       *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s ListUserIntentionNotesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUserIntentionNotesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserIntentionNotesResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListUserIntentionNotesResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *ListUserIntentionNotesResponseBodyData) SetCreateTime(v string) *ListUserIntentionNotesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListUserIntentionNotesResponseBodyData) SetNote(v string) *ListUserIntentionNotesResponseBodyData {
	s.Note = &v
	return s
}

func (s *ListUserIntentionNotesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
