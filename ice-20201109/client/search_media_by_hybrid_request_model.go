// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaByHybridRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *SearchMediaByHybridRequest
	GetMediaId() *string
	SetMediaType(v string) *SearchMediaByHybridRequest
	GetMediaType() *string
	SetNamespace(v string) *SearchMediaByHybridRequest
	GetNamespace() *string
	SetPageNo(v int32) *SearchMediaByHybridRequest
	GetPageNo() *int32
	SetPageSize(v int32) *SearchMediaByHybridRequest
	GetPageSize() *int32
	SetSearchLibName(v string) *SearchMediaByHybridRequest
	GetSearchLibName() *string
	SetText(v string) *SearchMediaByHybridRequest
	GetText() *string
}

type SearchMediaByHybridRequest struct {
	// The ID of the media asset. The details of the media asset are returned.
	//
	// example:
	//
	// ****c469e944b5a856828dc2****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// test-1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	Text          *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SearchMediaByHybridRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByHybridRequest) GoString() string {
	return s.String()
}

func (s *SearchMediaByHybridRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *SearchMediaByHybridRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *SearchMediaByHybridRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *SearchMediaByHybridRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *SearchMediaByHybridRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchMediaByHybridRequest) GetSearchLibName() *string {
	return s.SearchLibName
}

func (s *SearchMediaByHybridRequest) GetText() *string {
	return s.Text
}

func (s *SearchMediaByHybridRequest) SetMediaId(v string) *SearchMediaByHybridRequest {
	s.MediaId = &v
	return s
}

func (s *SearchMediaByHybridRequest) SetMediaType(v string) *SearchMediaByHybridRequest {
	s.MediaType = &v
	return s
}

func (s *SearchMediaByHybridRequest) SetNamespace(v string) *SearchMediaByHybridRequest {
	s.Namespace = &v
	return s
}

func (s *SearchMediaByHybridRequest) SetPageNo(v int32) *SearchMediaByHybridRequest {
	s.PageNo = &v
	return s
}

func (s *SearchMediaByHybridRequest) SetPageSize(v int32) *SearchMediaByHybridRequest {
	s.PageSize = &v
	return s
}

func (s *SearchMediaByHybridRequest) SetSearchLibName(v string) *SearchMediaByHybridRequest {
	s.SearchLibName = &v
	return s
}

func (s *SearchMediaByHybridRequest) SetText(v string) *SearchMediaByHybridRequest {
	s.Text = &v
	return s
}

func (s *SearchMediaByHybridRequest) Validate() error {
	return dara.Validate(s)
}
