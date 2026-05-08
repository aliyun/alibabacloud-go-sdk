// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTextFeedbackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *AddTextFeedbackRequest
	GetContent() *string
	SetQuality(v int32) *AddTextFeedbackRequest
	GetQuality() *int32
	SetTextId(v int64) *AddTextFeedbackRequest
	GetTextId() *int64
}

type AddTextFeedbackRequest struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// 1
	Quality *int32 `json:"quality,omitempty" xml:"quality,omitempty"`
	// example:
	//
	// 8478
	TextId *int64 `json:"textId,omitempty" xml:"textId,omitempty"`
}

func (s AddTextFeedbackRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTextFeedbackRequest) GoString() string {
	return s.String()
}

func (s *AddTextFeedbackRequest) GetContent() *string {
	return s.Content
}

func (s *AddTextFeedbackRequest) GetQuality() *int32 {
	return s.Quality
}

func (s *AddTextFeedbackRequest) GetTextId() *int64 {
	return s.TextId
}

func (s *AddTextFeedbackRequest) SetContent(v string) *AddTextFeedbackRequest {
	s.Content = &v
	return s
}

func (s *AddTextFeedbackRequest) SetQuality(v int32) *AddTextFeedbackRequest {
	s.Quality = &v
	return s
}

func (s *AddTextFeedbackRequest) SetTextId(v int64) *AddTextFeedbackRequest {
	s.TextId = &v
	return s
}

func (s *AddTextFeedbackRequest) Validate() error {
	return dara.Validate(s)
}
