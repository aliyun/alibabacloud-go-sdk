// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibraryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *GetLibraryListRequest
	GetPage() *int32
	SetPageSize(v int32) *GetLibraryListRequest
	GetPageSize() *int32
	SetQuery(v string) *GetLibraryListRequest
	GetQuery() *string
}

type GetLibraryListRequest struct {
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Query    *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s GetLibraryListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLibraryListRequest) GoString() string {
	return s.String()
}

func (s *GetLibraryListRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetLibraryListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetLibraryListRequest) GetQuery() *string {
	return s.Query
}

func (s *GetLibraryListRequest) SetPage(v int32) *GetLibraryListRequest {
	s.Page = &v
	return s
}

func (s *GetLibraryListRequest) SetPageSize(v int32) *GetLibraryListRequest {
	s.PageSize = &v
	return s
}

func (s *GetLibraryListRequest) SetQuery(v string) *GetLibraryListRequest {
	s.Query = &v
	return s
}

func (s *GetLibraryListRequest) Validate() error {
	return dara.Validate(s)
}
