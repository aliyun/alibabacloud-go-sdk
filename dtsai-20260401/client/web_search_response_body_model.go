// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *WebSearchResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *WebSearchResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *WebSearchResponseBody
	GetHttpStatusCode() *int32
	SetQuery(v string) *WebSearchResponseBody
	GetQuery() *string
	SetRequestId(v string) *WebSearchResponseBody
	GetRequestId() *string
	SetSearchResult(v []*WebSearchResponseBodySearchResult) *WebSearchResponseBody
	GetSearchResult() []*WebSearchResponseBodySearchResult
	SetSuccess(v bool) *WebSearchResponseBody
	GetSuccess() *bool
	SetTotalResults(v int32) *WebSearchResponseBody
	GetTotalResults() *int32
}

type WebSearchResponseBody struct {
	ErrorCode      *string                              `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage   *string                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Query          *string                              `json:"Query,omitempty" xml:"Query,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SearchResult   []*WebSearchResponseBodySearchResult `json:"SearchResult,omitempty" xml:"SearchResult,omitempty" type:"Repeated"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalResults   *int32                               `json:"TotalResults,omitempty" xml:"TotalResults,omitempty"`
}

func (s WebSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WebSearchResponseBody) GoString() string {
	return s.String()
}

func (s *WebSearchResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *WebSearchResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *WebSearchResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *WebSearchResponseBody) GetQuery() *string {
	return s.Query
}

func (s *WebSearchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WebSearchResponseBody) GetSearchResult() []*WebSearchResponseBodySearchResult {
	return s.SearchResult
}

func (s *WebSearchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *WebSearchResponseBody) GetTotalResults() *int32 {
	return s.TotalResults
}

func (s *WebSearchResponseBody) SetErrorCode(v string) *WebSearchResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *WebSearchResponseBody) SetErrorMessage(v string) *WebSearchResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *WebSearchResponseBody) SetHttpStatusCode(v int32) *WebSearchResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *WebSearchResponseBody) SetQuery(v string) *WebSearchResponseBody {
	s.Query = &v
	return s
}

func (s *WebSearchResponseBody) SetRequestId(v string) *WebSearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *WebSearchResponseBody) SetSearchResult(v []*WebSearchResponseBodySearchResult) *WebSearchResponseBody {
	s.SearchResult = v
	return s
}

func (s *WebSearchResponseBody) SetSuccess(v bool) *WebSearchResponseBody {
	s.Success = &v
	return s
}

func (s *WebSearchResponseBody) SetTotalResults(v int32) *WebSearchResponseBody {
	s.TotalResults = &v
	return s
}

func (s *WebSearchResponseBody) Validate() error {
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

type WebSearchResponseBodySearchResult struct {
	Snippet *string `json:"Snippet,omitempty" xml:"Snippet,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url     *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s WebSearchResponseBodySearchResult) String() string {
	return dara.Prettify(s)
}

func (s WebSearchResponseBodySearchResult) GoString() string {
	return s.String()
}

func (s *WebSearchResponseBodySearchResult) GetSnippet() *string {
	return s.Snippet
}

func (s *WebSearchResponseBodySearchResult) GetTitle() *string {
	return s.Title
}

func (s *WebSearchResponseBodySearchResult) GetUrl() *string {
	return s.Url
}

func (s *WebSearchResponseBodySearchResult) SetSnippet(v string) *WebSearchResponseBodySearchResult {
	s.Snippet = &v
	return s
}

func (s *WebSearchResponseBodySearchResult) SetTitle(v string) *WebSearchResponseBodySearchResult {
	s.Title = &v
	return s
}

func (s *WebSearchResponseBodySearchResult) SetUrl(v string) *WebSearchResponseBodySearchResult {
	s.Url = &v
	return s
}

func (s *WebSearchResponseBodySearchResult) Validate() error {
	return dara.Validate(s)
}
