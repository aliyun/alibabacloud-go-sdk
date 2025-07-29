// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotNewsRecommendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotNewsRecommendResponseBody
	GetCode() *string
	SetData(v *HotNewsRecommendResponseBodyData) *HotNewsRecommendResponseBody
	GetData() *HotNewsRecommendResponseBodyData
	SetMessage(v string) *HotNewsRecommendResponseBody
	GetMessage() *string
	SetRequestId(v string) *HotNewsRecommendResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotNewsRecommendResponseBody
	GetSuccess() *bool
}

type HotNewsRecommendResponseBody struct {
	// example:
	//
	// 200
	Code *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data *HotNewsRecommendResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 575D5893-01DB-5C81-A899-74F67616762A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s HotNewsRecommendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotNewsRecommendResponseBody) GoString() string {
	return s.String()
}

func (s *HotNewsRecommendResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotNewsRecommendResponseBody) GetData() *HotNewsRecommendResponseBodyData {
	return s.Data
}

func (s *HotNewsRecommendResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotNewsRecommendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotNewsRecommendResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotNewsRecommendResponseBody) SetCode(v string) *HotNewsRecommendResponseBody {
	s.Code = &v
	return s
}

func (s *HotNewsRecommendResponseBody) SetData(v *HotNewsRecommendResponseBodyData) *HotNewsRecommendResponseBody {
	s.Data = v
	return s
}

func (s *HotNewsRecommendResponseBody) SetMessage(v string) *HotNewsRecommendResponseBody {
	s.Message = &v
	return s
}

func (s *HotNewsRecommendResponseBody) SetRequestId(v string) *HotNewsRecommendResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotNewsRecommendResponseBody) SetSuccess(v bool) *HotNewsRecommendResponseBody {
	s.Success = &v
	return s
}

func (s *HotNewsRecommendResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotNewsRecommendResponseBodyData struct {
	News []*HotNewsRecommendResponseBodyDataNews `json:"news,omitempty" xml:"news,omitempty" type:"Repeated"`
}

func (s HotNewsRecommendResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s HotNewsRecommendResponseBodyData) GoString() string {
	return s.String()
}

func (s *HotNewsRecommendResponseBodyData) GetNews() []*HotNewsRecommendResponseBodyDataNews {
	return s.News
}

func (s *HotNewsRecommendResponseBodyData) SetNews(v []*HotNewsRecommendResponseBodyDataNews) *HotNewsRecommendResponseBodyData {
	s.News = v
	return s
}

func (s *HotNewsRecommendResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type HotNewsRecommendResponseBodyDataNews struct {
	// example:
	//
	// xx
	Content   *string   `json:"content,omitempty" xml:"content,omitempty"`
	ImageUrls []*string `json:"imageUrls,omitempty" xml:"imageUrls,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-09-10 15:32:00
	PubTime      *string `json:"pubTime,omitempty" xml:"pubTime,omitempty"`
	SearchSource *string `json:"searchSource,omitempty" xml:"searchSource,omitempty"`
	Source       *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// xx
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// http://xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s HotNewsRecommendResponseBodyDataNews) String() string {
	return dara.Prettify(s)
}

func (s HotNewsRecommendResponseBodyDataNews) GoString() string {
	return s.String()
}

func (s *HotNewsRecommendResponseBodyDataNews) GetContent() *string {
	return s.Content
}

func (s *HotNewsRecommendResponseBodyDataNews) GetImageUrls() []*string {
	return s.ImageUrls
}

func (s *HotNewsRecommendResponseBodyDataNews) GetPubTime() *string {
	return s.PubTime
}

func (s *HotNewsRecommendResponseBodyDataNews) GetSearchSource() *string {
	return s.SearchSource
}

func (s *HotNewsRecommendResponseBodyDataNews) GetSource() *string {
	return s.Source
}

func (s *HotNewsRecommendResponseBodyDataNews) GetTitle() *string {
	return s.Title
}

func (s *HotNewsRecommendResponseBodyDataNews) GetUrl() *string {
	return s.Url
}

func (s *HotNewsRecommendResponseBodyDataNews) SetContent(v string) *HotNewsRecommendResponseBodyDataNews {
	s.Content = &v
	return s
}

func (s *HotNewsRecommendResponseBodyDataNews) SetImageUrls(v []*string) *HotNewsRecommendResponseBodyDataNews {
	s.ImageUrls = v
	return s
}

func (s *HotNewsRecommendResponseBodyDataNews) SetPubTime(v string) *HotNewsRecommendResponseBodyDataNews {
	s.PubTime = &v
	return s
}

func (s *HotNewsRecommendResponseBodyDataNews) SetSearchSource(v string) *HotNewsRecommendResponseBodyDataNews {
	s.SearchSource = &v
	return s
}

func (s *HotNewsRecommendResponseBodyDataNews) SetSource(v string) *HotNewsRecommendResponseBodyDataNews {
	s.Source = &v
	return s
}

func (s *HotNewsRecommendResponseBodyDataNews) SetTitle(v string) *HotNewsRecommendResponseBodyDataNews {
	s.Title = &v
	return s
}

func (s *HotNewsRecommendResponseBodyDataNews) SetUrl(v string) *HotNewsRecommendResponseBodyDataNews {
	s.Url = &v
	return s
}

func (s *HotNewsRecommendResponseBodyDataNews) Validate() error {
	return dara.Validate(s)
}
