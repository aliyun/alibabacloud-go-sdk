// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOutputFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItemType(v string) *ListOutputFilesRequest
	GetItemType() *string
	SetKeyword(v string) *ListOutputFilesRequest
	GetKeyword() *string
	SetOperatingObjectName(v string) *ListOutputFilesRequest
	GetOperatingObjectName() *string
	SetPage(v int64) *ListOutputFilesRequest
	GetPage() *int64
	SetPageSize(v int64) *ListOutputFilesRequest
	GetPageSize() *int64
	SetSharedOnly(v bool) *ListOutputFilesRequest
	GetSharedOnly() *bool
	SetTenantId(v string) *ListOutputFilesRequest
	GetTenantId() *string
}

type ListOutputFilesRequest struct {
	// The type of the output item. Valid values: ppt, html, document, picture, slides, video, audio, email, and others.
	//
	// example:
	//
	// ppt
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// The keyword for searching. Matches output titles or item names.
	//
	// example:
	//
	// string_value
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The name of the digital employee (operating object). Used to filter results by name.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Specifies whether to display only outputs and output items that have sharing enabled.
	//
	// example:
	//
	// False
	SharedOnly *bool `json:"sharedOnly,omitempty" xml:"sharedOnly,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass it explicitly with --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListOutputFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOutputFilesRequest) GoString() string {
	return s.String()
}

func (s *ListOutputFilesRequest) GetItemType() *string {
	return s.ItemType
}

func (s *ListOutputFilesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListOutputFilesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *ListOutputFilesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListOutputFilesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListOutputFilesRequest) GetSharedOnly() *bool {
	return s.SharedOnly
}

func (s *ListOutputFilesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListOutputFilesRequest) SetItemType(v string) *ListOutputFilesRequest {
	s.ItemType = &v
	return s
}

func (s *ListOutputFilesRequest) SetKeyword(v string) *ListOutputFilesRequest {
	s.Keyword = &v
	return s
}

func (s *ListOutputFilesRequest) SetOperatingObjectName(v string) *ListOutputFilesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *ListOutputFilesRequest) SetPage(v int64) *ListOutputFilesRequest {
	s.Page = &v
	return s
}

func (s *ListOutputFilesRequest) SetPageSize(v int64) *ListOutputFilesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOutputFilesRequest) SetSharedOnly(v bool) *ListOutputFilesRequest {
	s.SharedOnly = &v
	return s
}

func (s *ListOutputFilesRequest) SetTenantId(v string) *ListOutputFilesRequest {
	s.TenantId = &v
	return s
}

func (s *ListOutputFilesRequest) Validate() error {
	return dara.Validate(s)
}
