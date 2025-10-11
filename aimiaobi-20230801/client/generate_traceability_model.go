// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTraceability interface {
	dara.Model
	String() string
	GoString() string
	SetNews(v []*GenerateTraceabilityNews) *GenerateTraceability
	GetNews() []*GenerateTraceabilityNews
}

type GenerateTraceability struct {
	News []*GenerateTraceabilityNews `json:"News,omitempty" xml:"News,omitempty" type:"Repeated"`
}

func (s GenerateTraceability) String() string {
	return dara.Prettify(s)
}

func (s GenerateTraceability) GoString() string {
	return s.String()
}

func (s *GenerateTraceability) GetNews() []*GenerateTraceabilityNews {
	return s.News
}

func (s *GenerateTraceability) SetNews(v []*GenerateTraceabilityNews) *GenerateTraceability {
	s.News = v
	return s
}

func (s *GenerateTraceability) Validate() error {
	return dara.Validate(s)
}

type GenerateTraceabilityNews struct {
	Index            *int32  `json:"Index,omitempty" xml:"Index,omitempty"`
	PubTime          *string `json:"PubTime,omitempty" xml:"PubTime,omitempty"`
	SearchSource     *string `json:"SearchSource,omitempty" xml:"SearchSource,omitempty"`
	SearchSourceName *string `json:"SearchSourceName,omitempty" xml:"SearchSourceName,omitempty"`
	Title            *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Url              *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GenerateTraceabilityNews) String() string {
	return dara.Prettify(s)
}

func (s GenerateTraceabilityNews) GoString() string {
	return s.String()
}

func (s *GenerateTraceabilityNews) GetIndex() *int32 {
	return s.Index
}

func (s *GenerateTraceabilityNews) GetPubTime() *string {
	return s.PubTime
}

func (s *GenerateTraceabilityNews) GetSearchSource() *string {
	return s.SearchSource
}

func (s *GenerateTraceabilityNews) GetSearchSourceName() *string {
	return s.SearchSourceName
}

func (s *GenerateTraceabilityNews) GetTitle() *string {
	return s.Title
}

func (s *GenerateTraceabilityNews) GetUrl() *string {
	return s.Url
}

func (s *GenerateTraceabilityNews) SetIndex(v int32) *GenerateTraceabilityNews {
	s.Index = &v
	return s
}

func (s *GenerateTraceabilityNews) SetPubTime(v string) *GenerateTraceabilityNews {
	s.PubTime = &v
	return s
}

func (s *GenerateTraceabilityNews) SetSearchSource(v string) *GenerateTraceabilityNews {
	s.SearchSource = &v
	return s
}

func (s *GenerateTraceabilityNews) SetSearchSourceName(v string) *GenerateTraceabilityNews {
	s.SearchSourceName = &v
	return s
}

func (s *GenerateTraceabilityNews) SetTitle(v string) *GenerateTraceabilityNews {
	s.Title = &v
	return s
}

func (s *GenerateTraceabilityNews) SetUrl(v string) *GenerateTraceabilityNews {
	s.Url = &v
	return s
}

func (s *GenerateTraceabilityNews) Validate() error {
	return dara.Validate(s)
}
