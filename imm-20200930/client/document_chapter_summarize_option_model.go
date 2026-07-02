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
	// The number of section-by-section summaries. If neither Marker nor Index is specified, the entire article is summarized by default. Marker and Index must either both be specified or both be omitted.
	//
	// example:
	//
	// 5
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The start position for section-by-section summarization.
	//
	// example:
	//
	// 0
	Marker *int32 `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The version of section-by-section summarization.
	//
	// example:
	//
	// v1
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
