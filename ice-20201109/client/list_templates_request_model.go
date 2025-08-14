// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateSource(v string) *ListTemplatesRequest
	GetCreateSource() *string
	SetKeyword(v string) *ListTemplatesRequest
	GetKeyword() *string
	SetPageNo(v int64) *ListTemplatesRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListTemplatesRequest
	GetPageSize() *int64
	SetSortType(v string) *ListTemplatesRequest
	GetSortType() *string
	SetStatus(v string) *ListTemplatesRequest
	GetStatus() *string
	SetType(v string) *ListTemplatesRequest
	GetType() *string
}

type ListTemplatesRequest struct {
	// The source from which the template was created.
	//
	// Valid values:
	//
	// 	- AliyunConsole
	//
	// 	- WebSDK
	//
	// 	- OpenAPI
	//
	// example:
	//
	// OpenAPI
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// The search keyword. You can use the template ID or title as the keyword to search for templates.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting parameter. By default, the query results are sorted by creation time in descending order.
	//
	// Valid values:
	//
	// 	- CreationTime:Asc: sorted by creation time in ascending order.
	//
	// 	- CreationTime:Desc: sorted by creation time in descending order.
	//
	// example:
	//
	// CreationTime:Desc
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// The template state.
	//
	// Valid values:
	//
	// 	- UploadFailed: Failed to upload the video.
	//
	// 	- ProcessFailed: Failed to process the advanced template.
	//
	// 	- Available: The template is available.
	//
	// 	- Uploading: The video is being uploaded.
	//
	// 	- Created: The template is created but not ready for use.
	//
	// 	- Processing: The advanced template is being processed.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template type.
	//
	// Valid values:
	//
	// 	- Timeline
	//
	// 	- VETemplate
	//
	// example:
	//
	// Timeline
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) GetCreateSource() *string {
	return s.CreateSource
}

func (s *ListTemplatesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTemplatesRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListTemplatesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListTemplatesRequest) GetSortType() *string {
	return s.SortType
}

func (s *ListTemplatesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTemplatesRequest) GetType() *string {
	return s.Type
}

func (s *ListTemplatesRequest) SetCreateSource(v string) *ListTemplatesRequest {
	s.CreateSource = &v
	return s
}

func (s *ListTemplatesRequest) SetKeyword(v string) *ListTemplatesRequest {
	s.Keyword = &v
	return s
}

func (s *ListTemplatesRequest) SetPageNo(v int64) *ListTemplatesRequest {
	s.PageNo = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int64) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetSortType(v string) *ListTemplatesRequest {
	s.SortType = &v
	return s
}

func (s *ListTemplatesRequest) SetStatus(v string) *ListTemplatesRequest {
	s.Status = &v
	return s
}

func (s *ListTemplatesRequest) SetType(v string) *ListTemplatesRequest {
	s.Type = &v
	return s
}

func (s *ListTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
