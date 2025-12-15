// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWebSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLatency(v int32) *GetWebSearchResponseBody
	GetLatency() *int32
	SetRequestId(v string) *GetWebSearchResponseBody
	GetRequestId() *string
	SetResult(v *GetWebSearchResponseBodyResult) *GetWebSearchResponseBody
	GetResult() *GetWebSearchResponseBodyResult
	SetUsage(v *GetWebSearchResponseBodyUsage) *GetWebSearchResponseBody
	GetUsage() *GetWebSearchResponseBodyUsage
}

type GetWebSearchResponseBody struct {
	Latency   *int32                          `json:"latency,omitempty" xml:"latency,omitempty"`
	RequestId *string                         `json:"request_id,omitempty" xml:"request_id,omitempty"`
	Result    *GetWebSearchResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	Usage     *GetWebSearchResponseBodyUsage  `json:"usage,omitempty" xml:"usage,omitempty" type:"Struct"`
}

func (s GetWebSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWebSearchResponseBody) GoString() string {
	return s.String()
}

func (s *GetWebSearchResponseBody) GetLatency() *int32 {
	return s.Latency
}

func (s *GetWebSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWebSearchResponseBody) GetResult() *GetWebSearchResponseBodyResult {
	return s.Result
}

func (s *GetWebSearchResponseBody) GetUsage() *GetWebSearchResponseBodyUsage {
	return s.Usage
}

func (s *GetWebSearchResponseBody) SetLatency(v int32) *GetWebSearchResponseBody {
	s.Latency = &v
	return s
}

func (s *GetWebSearchResponseBody) SetRequestId(v string) *GetWebSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWebSearchResponseBody) SetResult(v *GetWebSearchResponseBodyResult) *GetWebSearchResponseBody {
	s.Result = v
	return s
}

func (s *GetWebSearchResponseBody) SetUsage(v *GetWebSearchResponseBodyUsage) *GetWebSearchResponseBody {
	s.Usage = v
	return s
}

func (s *GetWebSearchResponseBody) Validate() error {
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

type GetWebSearchResponseBodyResult struct {
	SearchResult []*GetWebSearchResponseBodyResultSearchResult `json:"search_result,omitempty" xml:"search_result,omitempty" type:"Repeated"`
}

func (s GetWebSearchResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetWebSearchResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetWebSearchResponseBodyResult) GetSearchResult() []*GetWebSearchResponseBodyResultSearchResult {
	return s.SearchResult
}

func (s *GetWebSearchResponseBodyResult) SetSearchResult(v []*GetWebSearchResponseBodyResultSearchResult) *GetWebSearchResponseBodyResult {
	s.SearchResult = v
	return s
}

func (s *GetWebSearchResponseBodyResult) Validate() error {
	if s.SearchResult != nil {
		for _, item := range s.SearchResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWebSearchResponseBodyResultSearchResult struct {
	Content  *string `json:"content,omitempty" xml:"content,omitempty"`
	Link     *string `json:"link,omitempty" xml:"link,omitempty"`
	Position *int64  `json:"position,omitempty" xml:"position,omitempty"`
	Snippet  *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
	Tilte    *string `json:"tilte,omitempty" xml:"tilte,omitempty"`
}

func (s GetWebSearchResponseBodyResultSearchResult) String() string {
	return dara.Prettify(s)
}

func (s GetWebSearchResponseBodyResultSearchResult) GoString() string {
	return s.String()
}

func (s *GetWebSearchResponseBodyResultSearchResult) GetContent() *string {
	return s.Content
}

func (s *GetWebSearchResponseBodyResultSearchResult) GetLink() *string {
	return s.Link
}

func (s *GetWebSearchResponseBodyResultSearchResult) GetPosition() *int64 {
	return s.Position
}

func (s *GetWebSearchResponseBodyResultSearchResult) GetSnippet() *string {
	return s.Snippet
}

func (s *GetWebSearchResponseBodyResultSearchResult) GetTilte() *string {
	return s.Tilte
}

func (s *GetWebSearchResponseBodyResultSearchResult) SetContent(v string) *GetWebSearchResponseBodyResultSearchResult {
	s.Content = &v
	return s
}

func (s *GetWebSearchResponseBodyResultSearchResult) SetLink(v string) *GetWebSearchResponseBodyResultSearchResult {
	s.Link = &v
	return s
}

func (s *GetWebSearchResponseBodyResultSearchResult) SetPosition(v int64) *GetWebSearchResponseBodyResultSearchResult {
	s.Position = &v
	return s
}

func (s *GetWebSearchResponseBodyResultSearchResult) SetSnippet(v string) *GetWebSearchResponseBodyResultSearchResult {
	s.Snippet = &v
	return s
}

func (s *GetWebSearchResponseBodyResultSearchResult) SetTilte(v string) *GetWebSearchResponseBodyResultSearchResult {
	s.Tilte = &v
	return s
}

func (s *GetWebSearchResponseBodyResultSearchResult) Validate() error {
	return dara.Validate(s)
}

type GetWebSearchResponseBodyUsage struct {
	FilterModel  *GetWebSearchResponseBodyUsageFilterModel  `json:"filter_model,omitempty" xml:"filter_model,omitempty" type:"Struct"`
	RewriteModel *GetWebSearchResponseBodyUsageRewriteModel `json:"rewrite_model,omitempty" xml:"rewrite_model,omitempty" type:"Struct"`
	SearchCount  *int64                                     `json:"search_count,omitempty" xml:"search_count,omitempty"`
}

func (s GetWebSearchResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s GetWebSearchResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *GetWebSearchResponseBodyUsage) GetFilterModel() *GetWebSearchResponseBodyUsageFilterModel {
	return s.FilterModel
}

func (s *GetWebSearchResponseBodyUsage) GetRewriteModel() *GetWebSearchResponseBodyUsageRewriteModel {
	return s.RewriteModel
}

func (s *GetWebSearchResponseBodyUsage) GetSearchCount() *int64 {
	return s.SearchCount
}

func (s *GetWebSearchResponseBodyUsage) SetFilterModel(v *GetWebSearchResponseBodyUsageFilterModel) *GetWebSearchResponseBodyUsage {
	s.FilterModel = v
	return s
}

func (s *GetWebSearchResponseBodyUsage) SetRewriteModel(v *GetWebSearchResponseBodyUsageRewriteModel) *GetWebSearchResponseBodyUsage {
	s.RewriteModel = v
	return s
}

func (s *GetWebSearchResponseBodyUsage) SetSearchCount(v int64) *GetWebSearchResponseBodyUsage {
	s.SearchCount = &v
	return s
}

func (s *GetWebSearchResponseBodyUsage) Validate() error {
	if s.FilterModel != nil {
		if err := s.FilterModel.Validate(); err != nil {
			return err
		}
	}
	if s.RewriteModel != nil {
		if err := s.RewriteModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWebSearchResponseBodyUsageFilterModel struct {
	InputTokens  *int64 `json:"input_tokens,omitempty" xml:"input_tokens,omitempty"`
	OutputTokens *int64 `json:"output_tokens,omitempty" xml:"output_tokens,omitempty"`
	TotalTokens  *int64 `json:"total_tokens,omitempty" xml:"total_tokens,omitempty"`
}

func (s GetWebSearchResponseBodyUsageFilterModel) String() string {
	return dara.Prettify(s)
}

func (s GetWebSearchResponseBodyUsageFilterModel) GoString() string {
	return s.String()
}

func (s *GetWebSearchResponseBodyUsageFilterModel) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetWebSearchResponseBodyUsageFilterModel) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetWebSearchResponseBodyUsageFilterModel) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetWebSearchResponseBodyUsageFilterModel) SetInputTokens(v int64) *GetWebSearchResponseBodyUsageFilterModel {
	s.InputTokens = &v
	return s
}

func (s *GetWebSearchResponseBodyUsageFilterModel) SetOutputTokens(v int64) *GetWebSearchResponseBodyUsageFilterModel {
	s.OutputTokens = &v
	return s
}

func (s *GetWebSearchResponseBodyUsageFilterModel) SetTotalTokens(v int64) *GetWebSearchResponseBodyUsageFilterModel {
	s.TotalTokens = &v
	return s
}

func (s *GetWebSearchResponseBodyUsageFilterModel) Validate() error {
	return dara.Validate(s)
}

type GetWebSearchResponseBodyUsageRewriteModel struct {
	InputTokens  *int64 `json:"input_tokens,omitempty" xml:"input_tokens,omitempty"`
	OutputTokens *int64 `json:"output_tokens,omitempty" xml:"output_tokens,omitempty"`
	TotalTokens  *int64 `json:"total_tokens,omitempty" xml:"total_tokens,omitempty"`
}

func (s GetWebSearchResponseBodyUsageRewriteModel) String() string {
	return dara.Prettify(s)
}

func (s GetWebSearchResponseBodyUsageRewriteModel) GoString() string {
	return s.String()
}

func (s *GetWebSearchResponseBodyUsageRewriteModel) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *GetWebSearchResponseBodyUsageRewriteModel) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *GetWebSearchResponseBodyUsageRewriteModel) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *GetWebSearchResponseBodyUsageRewriteModel) SetInputTokens(v int64) *GetWebSearchResponseBodyUsageRewriteModel {
	s.InputTokens = &v
	return s
}

func (s *GetWebSearchResponseBodyUsageRewriteModel) SetOutputTokens(v int64) *GetWebSearchResponseBodyUsageRewriteModel {
	s.OutputTokens = &v
	return s
}

func (s *GetWebSearchResponseBodyUsageRewriteModel) SetTotalTokens(v int64) *GetWebSearchResponseBodyUsageRewriteModel {
	s.TotalTokens = &v
	return s
}

func (s *GetWebSearchResponseBodyUsageRewriteModel) Validate() error {
	return dara.Validate(s)
}
