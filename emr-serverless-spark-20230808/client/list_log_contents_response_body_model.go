// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogContentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListLogContent(v *ListLogContentsResponseBodyListLogContent) *ListLogContentsResponseBody
	GetListLogContent() *ListLogContentsResponseBodyListLogContent
	SetRequestId(v string) *ListLogContentsResponseBody
	GetRequestId() *string
}

type ListLogContentsResponseBody struct {
	// Log content.
	ListLogContent *ListLogContentsResponseBodyListLogContent `json:"listLogContent,omitempty" xml:"listLogContent,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListLogContentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogContentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogContentsResponseBody) GetListLogContent() *ListLogContentsResponseBodyListLogContent {
	return s.ListLogContent
}

func (s *ListLogContentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogContentsResponseBody) SetListLogContent(v *ListLogContentsResponseBodyListLogContent) *ListLogContentsResponseBody {
	s.ListLogContent = v
	return s
}

func (s *ListLogContentsResponseBody) SetRequestId(v string) *ListLogContentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogContentsResponseBody) Validate() error {
	if s.ListLogContent != nil {
		if err := s.ListLogContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLogContentsResponseBodyListLogContent struct {
	// List of log line contents.
	Contents []*ListLogContentsResponseBodyListLogContentContents `json:"contents,omitempty" xml:"contents,omitempty" type:"Repeated"`
	// Total number of log lines.
	//
	// example:
	//
	// 10
	TotalLength *int64 `json:"totalLength,omitempty" xml:"totalLength,omitempty"`
}

func (s ListLogContentsResponseBodyListLogContent) String() string {
	return dara.Prettify(s)
}

func (s ListLogContentsResponseBodyListLogContent) GoString() string {
	return s.String()
}

func (s *ListLogContentsResponseBodyListLogContent) GetContents() []*ListLogContentsResponseBodyListLogContentContents {
	return s.Contents
}

func (s *ListLogContentsResponseBodyListLogContent) GetTotalLength() *int64 {
	return s.TotalLength
}

func (s *ListLogContentsResponseBodyListLogContent) SetContents(v []*ListLogContentsResponseBodyListLogContentContents) *ListLogContentsResponseBodyListLogContent {
	s.Contents = v
	return s
}

func (s *ListLogContentsResponseBodyListLogContent) SetTotalLength(v int64) *ListLogContentsResponseBodyListLogContent {
	s.TotalLength = &v
	return s
}

func (s *ListLogContentsResponseBodyListLogContent) Validate() error {
	if s.Contents != nil {
		for _, item := range s.Contents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLogContentsResponseBodyListLogContentContents struct {
	// Log line content.
	//
	// example:
	//
	// spark pi is 3.14\\n
	LineContent *string `json:"LineContent,omitempty" xml:"LineContent,omitempty"`
}

func (s ListLogContentsResponseBodyListLogContentContents) String() string {
	return dara.Prettify(s)
}

func (s ListLogContentsResponseBodyListLogContentContents) GoString() string {
	return s.String()
}

func (s *ListLogContentsResponseBodyListLogContentContents) GetLineContent() *string {
	return s.LineContent
}

func (s *ListLogContentsResponseBodyListLogContentContents) SetLineContent(v string) *ListLogContentsResponseBodyListLogContentContents {
	s.LineContent = &v
	return s
}

func (s *ListLogContentsResponseBodyListLogContentContents) Validate() error {
	return dara.Validate(s)
}
