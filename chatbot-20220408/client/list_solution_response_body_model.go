// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSolutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSolutionResponseBody
	GetRequestId() *string
	SetSolutions(v []*ListSolutionResponseBodySolutions) *ListSolutionResponseBody
	GetSolutions() []*ListSolutionResponseBodySolutions
}

type ListSolutionResponseBody struct {
	// example:
	//
	// 5B29DB5E-251D-5A73-84B5-A12DF795F231
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Solutions []*ListSolutionResponseBodySolutions `json:"Solutions,omitempty" xml:"Solutions,omitempty" type:"Repeated"`
}

func (s ListSolutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *ListSolutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSolutionResponseBody) GetSolutions() []*ListSolutionResponseBodySolutions {
	return s.Solutions
}

func (s *ListSolutionResponseBody) SetRequestId(v string) *ListSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSolutionResponseBody) SetSolutions(v []*ListSolutionResponseBodySolutions) *ListSolutionResponseBody {
	s.Solutions = v
	return s
}

func (s *ListSolutionResponseBody) Validate() error {
	if s.Solutions != nil {
		for _, item := range s.Solutions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSolutionResponseBodySolutions struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 1
	ContentType *int32 `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// example:
	//
	// 2022-03-29T03:55:04Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-03-29T06:23:53Z
	ModifyTime       *string   `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	PerspectiveCodes []*string `json:"PerspectiveCodes,omitempty" xml:"PerspectiveCodes,omitempty" type:"Repeated"`
	PlainText        *string   `json:"PlainText,omitempty" xml:"PlainText,omitempty"`
	// example:
	//
	// 496
	SolutionId *int64 `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
}

func (s ListSolutionResponseBodySolutions) String() string {
	return dara.Prettify(s)
}

func (s ListSolutionResponseBodySolutions) GoString() string {
	return s.String()
}

func (s *ListSolutionResponseBodySolutions) GetContent() *string {
	return s.Content
}

func (s *ListSolutionResponseBodySolutions) GetContentType() *int32 {
	return s.ContentType
}

func (s *ListSolutionResponseBodySolutions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSolutionResponseBodySolutions) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListSolutionResponseBodySolutions) GetPerspectiveCodes() []*string {
	return s.PerspectiveCodes
}

func (s *ListSolutionResponseBodySolutions) GetPlainText() *string {
	return s.PlainText
}

func (s *ListSolutionResponseBodySolutions) GetSolutionId() *int64 {
	return s.SolutionId
}

func (s *ListSolutionResponseBodySolutions) SetContent(v string) *ListSolutionResponseBodySolutions {
	s.Content = &v
	return s
}

func (s *ListSolutionResponseBodySolutions) SetContentType(v int32) *ListSolutionResponseBodySolutions {
	s.ContentType = &v
	return s
}

func (s *ListSolutionResponseBodySolutions) SetCreateTime(v string) *ListSolutionResponseBodySolutions {
	s.CreateTime = &v
	return s
}

func (s *ListSolutionResponseBodySolutions) SetModifyTime(v string) *ListSolutionResponseBodySolutions {
	s.ModifyTime = &v
	return s
}

func (s *ListSolutionResponseBodySolutions) SetPerspectiveCodes(v []*string) *ListSolutionResponseBodySolutions {
	s.PerspectiveCodes = v
	return s
}

func (s *ListSolutionResponseBodySolutions) SetPlainText(v string) *ListSolutionResponseBodySolutions {
	s.PlainText = &v
	return s
}

func (s *ListSolutionResponseBodySolutions) SetSolutionId(v int64) *ListSolutionResponseBodySolutions {
	s.SolutionId = &v
	return s
}

func (s *ListSolutionResponseBodySolutions) Validate() error {
	return dara.Validate(s)
}
