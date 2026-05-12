// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebSearchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *WebSearchResponseBody
	GetCode() *int32
	SetData(v *WebSearchResponseBodyData) *WebSearchResponseBody
	GetData() *WebSearchResponseBodyData
	SetMessage(v string) *WebSearchResponseBody
	GetMessage() *string
	SetTraceId(v string) *WebSearchResponseBody
	GetTraceId() *string
}

type WebSearchResponseBody struct {
	// example:
	//
	// Ok
	Code *int32                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *WebSearchResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// requestId
	//
	// example:
	//
	// 3b5215d417623961959166934d009a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s WebSearchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WebSearchResponseBody) GoString() string {
	return s.String()
}

func (s *WebSearchResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *WebSearchResponseBody) GetData() *WebSearchResponseBodyData {
	return s.Data
}

func (s *WebSearchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *WebSearchResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *WebSearchResponseBody) SetCode(v int32) *WebSearchResponseBody {
	s.Code = &v
	return s
}

func (s *WebSearchResponseBody) SetData(v *WebSearchResponseBodyData) *WebSearchResponseBody {
	s.Data = v
	return s
}

func (s *WebSearchResponseBody) SetMessage(v string) *WebSearchResponseBody {
	s.Message = &v
	return s
}

func (s *WebSearchResponseBody) SetTraceId(v string) *WebSearchResponseBody {
	s.TraceId = &v
	return s
}

func (s *WebSearchResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WebSearchResponseBodyData struct {
	Result []*WebSearchResponseBodyDataResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s WebSearchResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s WebSearchResponseBodyData) GoString() string {
	return s.String()
}

func (s *WebSearchResponseBodyData) GetResult() []*WebSearchResponseBodyDataResult {
	return s.Result
}

func (s *WebSearchResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *WebSearchResponseBodyData) SetResult(v []*WebSearchResponseBodyDataResult) *WebSearchResponseBodyData {
	s.Result = v
	return s
}

func (s *WebSearchResponseBodyData) SetTotal(v int32) *WebSearchResponseBodyData {
	s.Total = &v
	return s
}

func (s *WebSearchResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type WebSearchResponseBodyDataResult struct {
	Snippet *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
	// example:
	//
	// 4567
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// https://pai-dlc-proxy-1-cn-wulanchabu.aliyun.com/ray/dashboard/dlc1a9r0uhfn24cl
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s WebSearchResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s WebSearchResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *WebSearchResponseBodyDataResult) GetSnippet() *string {
	return s.Snippet
}

func (s *WebSearchResponseBodyDataResult) GetTitle() *string {
	return s.Title
}

func (s *WebSearchResponseBodyDataResult) GetUrl() *string {
	return s.Url
}

func (s *WebSearchResponseBodyDataResult) SetSnippet(v string) *WebSearchResponseBodyDataResult {
	s.Snippet = &v
	return s
}

func (s *WebSearchResponseBodyDataResult) SetTitle(v string) *WebSearchResponseBodyDataResult {
	s.Title = &v
	return s
}

func (s *WebSearchResponseBodyDataResult) SetUrl(v string) *WebSearchResponseBodyDataResult {
	s.Url = &v
	return s
}

func (s *WebSearchResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
