// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryProcessorAnalyzerResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetText(v string) *ListQueryProcessorAnalyzerResultsRequest
	GetText() *string
}

type ListQueryProcessorAnalyzerResultsRequest struct {
	// The text to be tested.
	//
	// This parameter is required.
	//
	// example:
	//
	// "abcde"
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ListQueryProcessorAnalyzerResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueryProcessorAnalyzerResultsRequest) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorAnalyzerResultsRequest) GetText() *string {
	return s.Text
}

func (s *ListQueryProcessorAnalyzerResultsRequest) SetText(v string) *ListQueryProcessorAnalyzerResultsRequest {
	s.Text = &v
	return s
}

func (s *ListQueryProcessorAnalyzerResultsRequest) Validate() error {
	return dara.Validate(s)
}
