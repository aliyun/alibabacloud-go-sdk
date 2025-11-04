// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMediaByHybridRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomFilters(v string) *SearchMediaByHybridRequest
	GetCustomFilters() *string
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
	SetUtcCreate(v string) *SearchMediaByHybridRequest
	GetUtcCreate() *string
}

type SearchMediaByHybridRequest struct {
	// example:
	//
	// {}
	CustomFilters *string `json:"CustomFilters,omitempty" xml:"CustomFilters,omitempty"`
	// The ID of the media asset. If provided, the details of the media asset are returned.
	//
	// example:
	//
	// ****c469e944b5a856828dc2****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The type of media assets. Valid values:
	//
	// - image
	//
	// - video
	//
	// example:
	//
	// video
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The namespace.
	//
	// example:
	//
	// name-1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the search library
	//
	// example:
	//
	// test-1
	SearchLibName *string `json:"SearchLibName,omitempty" xml:"SearchLibName,omitempty"`
	// The natural language search query.
	//
	// example:
	//
	// Two pandas are fighting
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// {}
	UtcCreate *string `json:"UtcCreate,omitempty" xml:"UtcCreate,omitempty"`
}

func (s SearchMediaByHybridRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchMediaByHybridRequest) GoString() string {
	return s.String()
}

func (s *SearchMediaByHybridRequest) GetCustomFilters() *string {
	return s.CustomFilters
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

func (s *SearchMediaByHybridRequest) GetUtcCreate() *string {
	return s.UtcCreate
}

func (s *SearchMediaByHybridRequest) SetCustomFilters(v string) *SearchMediaByHybridRequest {
	s.CustomFilters = &v
	return s
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

func (s *SearchMediaByHybridRequest) SetUtcCreate(v string) *SearchMediaByHybridRequest {
	s.UtcCreate = &v
	return s
}

func (s *SearchMediaByHybridRequest) Validate() error {
	return dara.Validate(s)
}
