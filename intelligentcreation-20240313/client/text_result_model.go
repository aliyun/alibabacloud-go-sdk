// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextResult interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TextResult
	GetRequestId() *string
	SetText(v *Text) *TextResult
	GetText() *Text
}

type TextResult struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Text *Text `json:"text,omitempty" xml:"text,omitempty"`
}

func (s TextResult) String() string {
	return dara.Prettify(s)
}

func (s TextResult) GoString() string {
	return s.String()
}

func (s *TextResult) GetRequestId() *string {
	return s.RequestId
}

func (s *TextResult) GetText() *Text {
	return s.Text
}

func (s *TextResult) SetRequestId(v string) *TextResult {
	s.RequestId = &v
	return s
}

func (s *TextResult) SetText(v *Text) *TextResult {
	s.Text = v
	return s
}

func (s *TextResult) Validate() error {
	if s.Text != nil {
		if err := s.Text.Validate(); err != nil {
			return err
		}
	}
	return nil
}
