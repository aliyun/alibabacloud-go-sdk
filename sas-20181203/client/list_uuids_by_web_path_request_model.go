// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUuidsByWebPathRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListUuidsByWebPathRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListUuidsByWebPathRequest
	GetPageSize() *int32
	SetType(v string) *ListUuidsByWebPathRequest
	GetType() *string
	SetWebPath(v string) *ListUuidsByWebPathRequest
	GetWebPath() *string
}

type ListUuidsByWebPathRequest struct {
	// The page number of the current page to display in a paging query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries to display on each page in a paging query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the web path. Valid values:
	//
	// - **def**: automatically identified by the system
	//
	// - **customize**: manually added.
	//
	// example:
	//
	// def
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The web path.
	//
	// example:
	//
	// /root/www****
	WebPath *string `json:"WebPath,omitempty" xml:"WebPath,omitempty"`
}

func (s ListUuidsByWebPathRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUuidsByWebPathRequest) GoString() string {
	return s.String()
}

func (s *ListUuidsByWebPathRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUuidsByWebPathRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUuidsByWebPathRequest) GetType() *string {
	return s.Type
}

func (s *ListUuidsByWebPathRequest) GetWebPath() *string {
	return s.WebPath
}

func (s *ListUuidsByWebPathRequest) SetCurrentPage(v int32) *ListUuidsByWebPathRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListUuidsByWebPathRequest) SetPageSize(v int32) *ListUuidsByWebPathRequest {
	s.PageSize = &v
	return s
}

func (s *ListUuidsByWebPathRequest) SetType(v string) *ListUuidsByWebPathRequest {
	s.Type = &v
	return s
}

func (s *ListUuidsByWebPathRequest) SetWebPath(v string) *ListUuidsByWebPathRequest {
	s.WebPath = &v
	return s
}

func (s *ListUuidsByWebPathRequest) Validate() error {
	return dara.Validate(s)
}
