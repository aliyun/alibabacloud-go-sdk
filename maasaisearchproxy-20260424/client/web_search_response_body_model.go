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
	// example:
	//
	// 1990-01-01 12:00:00
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// snippet
	Snippet *string                                `json:"snippet,omitempty" xml:"snippet,omitempty"`
	Source  *WebSearchResponseBodyDataResultSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
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

func (s *WebSearchResponseBodyDataResult) GetDate() *string {
	return s.Date
}

func (s *WebSearchResponseBodyDataResult) GetSnippet() *string {
	return s.Snippet
}

func (s *WebSearchResponseBodyDataResult) GetSource() *WebSearchResponseBodyDataResultSource {
	return s.Source
}

func (s *WebSearchResponseBodyDataResult) GetTitle() *string {
	return s.Title
}

func (s *WebSearchResponseBodyDataResult) GetUrl() *string {
	return s.Url
}

func (s *WebSearchResponseBodyDataResult) SetDate(v string) *WebSearchResponseBodyDataResult {
	s.Date = &v
	return s
}

func (s *WebSearchResponseBodyDataResult) SetSnippet(v string) *WebSearchResponseBodyDataResult {
	s.Snippet = &v
	return s
}

func (s *WebSearchResponseBodyDataResult) SetSource(v *WebSearchResponseBodyDataResultSource) *WebSearchResponseBodyDataResult {
	s.Source = v
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
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type WebSearchResponseBodyDataResultSource struct {
	// example:
	//
	// domain
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// example:
	//
	// favicon
	Favicon *string `json:"favicon,omitempty" xml:"favicon,omitempty"`
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s WebSearchResponseBodyDataResultSource) String() string {
	return dara.Prettify(s)
}

func (s WebSearchResponseBodyDataResultSource) GoString() string {
	return s.String()
}

func (s *WebSearchResponseBodyDataResultSource) GetDomain() *string {
	return s.Domain
}

func (s *WebSearchResponseBodyDataResultSource) GetFavicon() *string {
	return s.Favicon
}

func (s *WebSearchResponseBodyDataResultSource) GetName() *string {
	return s.Name
}

func (s *WebSearchResponseBodyDataResultSource) SetDomain(v string) *WebSearchResponseBodyDataResultSource {
	s.Domain = &v
	return s
}

func (s *WebSearchResponseBodyDataResultSource) SetFavicon(v string) *WebSearchResponseBodyDataResultSource {
	s.Favicon = &v
	return s
}

func (s *WebSearchResponseBodyDataResultSource) SetName(v string) *WebSearchResponseBodyDataResultSource {
	s.Name = &v
	return s
}

func (s *WebSearchResponseBodyDataResultSource) Validate() error {
	return dara.Validate(s)
}
