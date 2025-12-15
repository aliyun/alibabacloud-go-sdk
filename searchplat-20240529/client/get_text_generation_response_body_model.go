// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextGenerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetTextGenerationResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetTextGenerationResponseBody
	GetRequestId() *string
	SetResult(v *GetTextGenerationResponseBodyResult) *GetTextGenerationResponseBody
	GetResult() *GetTextGenerationResponseBodyResult
	SetUsage(v *GetTextGenerationResponseBodyUsage) *GetTextGenerationResponseBody
	GetUsage() *GetTextGenerationResponseBodyUsage
}

type GetTextGenerationResponseBody struct {
	Latency   *int32                               `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                              `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetTextGenerationResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetTextGenerationResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetTextGenerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTextGenerationResponseBody) GoString() string {
	return s.String()
}

func (s *GetTextGenerationResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetTextGenerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTextGenerationResponseBody) GetResult() *GetTextGenerationResponseBodyResult {
	return s.Result
}

func (s *GetTextGenerationResponseBody) GetUsage() *GetTextGenerationResponseBodyUsage {
	return s.Usage
}

func (s *GetTextGenerationResponseBody) SetLatency(v int32) *GetTextGenerationResponseBody {
	s.Latency = &v
	return s
}

func (s *GetTextGenerationResponseBody) SetRequestId(v string) *GetTextGenerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTextGenerationResponseBody) SetResult(v *GetTextGenerationResponseBodyResult) *GetTextGenerationResponseBody {
	s.Result = v
	return s
}

func (s *GetTextGenerationResponseBody) SetUsage(v *GetTextGenerationResponseBodyUsage) *GetTextGenerationResponseBody {
	s.Usage = v
	return s
}

func (s *GetTextGenerationResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTextGenerationResponseBodyResult struct {
	SearchResults []*GetTextGenerationResponseBodyResultSearchResults `json:"search_results,omitempty" xml:"search_results,omitempty" type:"Repeated"`
	Text          *string                                             `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetTextGenerationResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetTextGenerationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTextGenerationResponseBodyResult) GetSearchResults() []*GetTextGenerationResponseBodyResultSearchResults {
	return s.SearchResults
}

func (s *GetTextGenerationResponseBodyResult) GetText() *string {
	return s.Text
}

func (s *GetTextGenerationResponseBodyResult) SetSearchResults(v []*GetTextGenerationResponseBodyResultSearchResults) *GetTextGenerationResponseBodyResult {
	s.SearchResults = v
	return s
}

func (s *GetTextGenerationResponseBodyResult) SetText(v string) *GetTextGenerationResponseBodyResult {
	s.Text = &v
	return s
}

func (s *GetTextGenerationResponseBodyResult) Validate() error {
	if s.SearchResults != nil {
		for _, item := range s.SearchResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTextGenerationResponseBodyResultSearchResults struct {
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	Url   *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetTextGenerationResponseBodyResultSearchResults) String() string {
	return dara.Prettify(s)
}

func (s GetTextGenerationResponseBodyResultSearchResults) GoString() string {
	return s.String()
}

func (s *GetTextGenerationResponseBodyResultSearchResults) GetTitle() *string {
	return s.Title
}

func (s *GetTextGenerationResponseBodyResultSearchResults) GetUrl() *string {
	return s.Url
}

func (s *GetTextGenerationResponseBodyResultSearchResults) SetTitle(v string) *GetTextGenerationResponseBodyResultSearchResults {
	s.Title = &v
	return s
}

func (s *GetTextGenerationResponseBodyResultSearchResults) SetUrl(v string) *GetTextGenerationResponseBodyResultSearchResults {
	s.Url = &v
	return s
}

func (s *GetTextGenerationResponseBodyResultSearchResults) Validate() error {
	return dara.Validate(s)
}

type GetTextGenerationResponseBodyUsage struct {
	InputTokens  *int64 `json:"input_tokens,omitempty" xml:"input_tokens,omitempty"`
	OutputTokens *int64 `json:"output_tokens,omitempty" xml:"output_tokens,omitempty"`
	TotalTokens  *int64 `json:"total_tokens,omitempty" xml:"total_tokens,omitempty"`
}

func (s GetTextGenerationResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetTextGenerationResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetTextGenerationResponseBodyUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetTextGenerationResponseBodyUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetTextGenerationResponseBodyUsage) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetTextGenerationResponseBodyUsage) SetInputTokens(v int64) *GetTextGenerationResponseBodyUsage {
	s.InputTokens = &v
	return s
}

func (s *GetTextGenerationResponseBodyUsage) SetOutputTokens(v int64) *GetTextGenerationResponseBodyUsage {
	s.OutputTokens = &v
	return s
}

func (s *GetTextGenerationResponseBodyUsage) SetTotalTokens(v int64) *GetTextGenerationResponseBodyUsage {
	s.TotalTokens = &v
	return s
}

func (s *GetTextGenerationResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
