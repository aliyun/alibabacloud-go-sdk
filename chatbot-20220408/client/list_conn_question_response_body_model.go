// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOutlines(v []*ListConnQuestionResponseBodyOutlines) *ListConnQuestionResponseBody
	GetOutlines() []*ListConnQuestionResponseBodyOutlines
	SetRequestId(v string) *ListConnQuestionResponseBody
	GetRequestId() *string
}

type ListConnQuestionResponseBody struct {
	Outlines []*ListConnQuestionResponseBodyOutlines `json:"Outlines,omitempty" xml:"Outlines,omitempty" type:"Repeated"`
	// example:
	//
	// 92B81548-42B9-4B34-924B-4E778AEB412B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConnQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConnQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnQuestionResponseBody) GetOutlines() []*ListConnQuestionResponseBodyOutlines {
	return s.Outlines
}

func (s *ListConnQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConnQuestionResponseBody) SetOutlines(v []*ListConnQuestionResponseBodyOutlines) *ListConnQuestionResponseBody {
	s.Outlines = v
	return s
}

func (s *ListConnQuestionResponseBody) SetRequestId(v string) *ListConnQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnQuestionResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListConnQuestionResponseBodyOutlines struct {
	// example:
	//
	// 30001979424
	ConnQuestionId *int64 `json:"ConnQuestionId,omitempty" xml:"ConnQuestionId,omitempty"`
	// example:
	//
	// 2022-02-25T02:47:18Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-05-26T10:18:15Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 797
	OutlineId *int64  `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
	Title     *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s ListConnQuestionResponseBodyOutlines) String() string {
	return dara.Prettify(s)
}

func (s ListConnQuestionResponseBodyOutlines) GoString() string {
	return s.String()
}

func (s *ListConnQuestionResponseBodyOutlines) GetConnQuestionId() *int64 {
	return s.ConnQuestionId
}

func (s *ListConnQuestionResponseBodyOutlines) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListConnQuestionResponseBodyOutlines) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListConnQuestionResponseBodyOutlines) GetOutlineId() *int64 {
	return s.OutlineId
}

func (s *ListConnQuestionResponseBodyOutlines) GetTitle() *string {
	return s.Title
}

func (s *ListConnQuestionResponseBodyOutlines) SetConnQuestionId(v int64) *ListConnQuestionResponseBodyOutlines {
	s.ConnQuestionId = &v
	return s
}

func (s *ListConnQuestionResponseBodyOutlines) SetCreateTime(v string) *ListConnQuestionResponseBodyOutlines {
	s.CreateTime = &v
	return s
}

func (s *ListConnQuestionResponseBodyOutlines) SetModifyTime(v string) *ListConnQuestionResponseBodyOutlines {
	s.ModifyTime = &v
	return s
}

func (s *ListConnQuestionResponseBodyOutlines) SetOutlineId(v int64) *ListConnQuestionResponseBodyOutlines {
	s.OutlineId = &v
	return s
}

func (s *ListConnQuestionResponseBodyOutlines) SetTitle(v string) *ListConnQuestionResponseBodyOutlines {
	s.Title = &v
	return s
}

func (s *ListConnQuestionResponseBodyOutlines) Validate() error {
	return dara.Validate(s)
}
