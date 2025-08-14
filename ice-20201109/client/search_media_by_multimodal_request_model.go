// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaByMultimodalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaType(v string) *SearchMediaByMultimodalRequest
	GetMediaType() *string
	SetNamespace(v string) *SearchMediaByMultimodalRequest
	GetNamespace() *string
	SetPageNo(v int32) *SearchMediaByMultimodalRequest
	GetPageNo() *int32
	SetPageSize(v int32) *SearchMediaByMultimodalRequest
	GetPageSize() *int32
	SetSearchLibName(v string) *SearchMediaByMultimodalRequest
	GetSearchLibName() *string
	SetText(v string) *SearchMediaByMultimodalRequest
	GetText() *string
}

type SearchMediaByMultimodalRequest struct {
	// The type of the media assets.
	//
	// Valid values:
	//
	// 	- image
	//
	// 	- video (default)
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search library.
	//
	// example:
	//
	// test-1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// The content that you want to query. You can describe the content in natural language.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SearchMediaByMultimodalRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByMultimodalRequest) GoString() string {
	return s.String()
}

func (s *SearchMediaByMultimodalRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *SearchMediaByMultimodalRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *SearchMediaByMultimodalRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *SearchMediaByMultimodalRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchMediaByMultimodalRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *SearchMediaByMultimodalRequest) GetText() *string {
	return s.Text
}

func (s *SearchMediaByMultimodalRequest) SetMediaType(v string) *SearchMediaByMultimodalRequest {
	s.MediaType = &v
	return s
}

func (s *SearchMediaByMultimodalRequest) SetNamespace(v string) *SearchMediaByMultimodalRequest {
	s.Namespace = &v
	return s
}

func (s *SearchMediaByMultimodalRequest) SetPageNo(v int32) *SearchMediaByMultimodalRequest {
	s.PageNo = &v
	return s
}

func (s *SearchMediaByMultimodalRequest) SetPageSize(v int32) *SearchMediaByMultimodalRequest {
	s.PageSize = &v
	return s
}

func (s *SearchMediaByMultimodalRequest) SetSearchLibName(v string) *SearchMediaByMultimodalRequest {
	s.SearchLibName = &v
	return s
}

func (s *SearchMediaByMultimodalRequest) SetText(v string) *SearchMediaByMultimodalRequest {
	s.Text = &v
	return s
}

func (s *SearchMediaByMultimodalRequest) Validate() error {
	return dara.Validate(s)
}
