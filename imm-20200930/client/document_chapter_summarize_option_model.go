// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocumentChapterSummarizeOption interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *DocumentChapterSummarizeOption
	GetLimit() *int32
	SetMarker(v int32) *DocumentChapterSummarizeOption
	GetMarker() *int32
	SetVersion(v string) *DocumentChapterSummarizeOption
	GetVersion() *string
}

type DocumentChapterSummarizeOption struct {
	// The maximum number of results to return per request. If you do not specify this parameter, the default value is used.
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The pagination token. It specifies the starting point for the query. To retrieve the next page of results, set this parameter to the marker value from the previous response.
	Marker *int32 `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The summarization model version. If you do not specify this parameter, the default model version is used.
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DocumentChapterSummarizeOption) String() string {
	return dara.Prettify(s)
}

func (s DocumentChapterSummarizeOption) GoString() string {
	return s.String()
}

func (s *DocumentChapterSummarizeOption) GetLimit() *int32 {
	return s.Limit
}

func (s *DocumentChapterSummarizeOption) GetMarker() *int32 {
	return s.Marker
}

func (s *DocumentChapterSummarizeOption) GetVersion() *string {
	return s.Version
}

func (s *DocumentChapterSummarizeOption) SetLimit(v int32) *DocumentChapterSummarizeOption {
	s.Limit = &v
	return s
}

func (s *DocumentChapterSummarizeOption) SetMarker(v int32) *DocumentChapterSummarizeOption {
	s.Marker = &v
	return s
}

func (s *DocumentChapterSummarizeOption) SetVersion(v string) *DocumentChapterSummarizeOption {
	s.Version = &v
	return s
}

func (s *DocumentChapterSummarizeOption) Validate() error {
	return dara.Validate(s)
}
