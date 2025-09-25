// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibraryId(v string) *GetDocumentListRequest
	GetLibraryId() *string
	SetPage(v int32) *GetDocumentListRequest
	GetPage() *int32
	SetPageSize(v int32) *GetDocumentListRequest
	GetPageSize() *int32
	SetStatus(v string) *GetDocumentListRequest
	GetStatus() *string
}

type GetDocumentListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 3akzl28vap
	LibraryId *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetDocumentListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentListRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentListRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *GetDocumentListRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetDocumentListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDocumentListRequest) GetStatus() *string {
	return s.Status
}

func (s *GetDocumentListRequest) SetLibraryId(v string) *GetDocumentListRequest {
	s.LibraryId = &v
	return s
}

func (s *GetDocumentListRequest) SetPage(v int32) *GetDocumentListRequest {
	s.Page = &v
	return s
}

func (s *GetDocumentListRequest) SetPageSize(v int32) *GetDocumentListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDocumentListRequest) SetStatus(v string) *GetDocumentListRequest {
	s.Status = &v
	return s
}

func (s *GetDocumentListRequest) Validate() error {
	return dara.Validate(s)
}
