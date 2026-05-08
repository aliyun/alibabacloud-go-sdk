// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextQueryResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TextQueryResult
	GetRequestId() *string
	SetTexts(v []*Text) *TextQueryResult
	GetTexts() []*Text
	SetTotal(v int32) *TextQueryResult
	GetTotal() *int32
}

type TextQueryResult struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Texts     []*Text `json:"texts,omitempty" xml:"texts,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s TextQueryResult) String() string {
	return dara.Prettify(s)
}

func (s TextQueryResult) GoString() string {
	return s.String()
}

func (s *TextQueryResult) GetRequestId() *string {
	return s.RequestId
}

func (s *TextQueryResult) GetTexts() []*Text {
	return s.Texts
}

func (s *TextQueryResult) GetTotal() *int32 {
	return s.Total
}

func (s *TextQueryResult) SetRequestId(v string) *TextQueryResult {
	s.RequestId = &v
	return s
}

func (s *TextQueryResult) SetTexts(v []*Text) *TextQueryResult {
	s.Texts = v
	return s
}

func (s *TextQueryResult) SetTotal(v int32) *TextQueryResult {
	s.Total = &v
	return s
}

func (s *TextQueryResult) Validate() error {
	if s.Texts != nil {
		for _, item := range s.Texts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
