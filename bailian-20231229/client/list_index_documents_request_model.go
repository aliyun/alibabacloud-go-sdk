// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentName(v string) *ListIndexDocumentsRequest
	GetDocumentName() *string
	SetDocumentStatus(v string) *ListIndexDocumentsRequest
	GetDocumentStatus() *string
	SetEnableNameLike(v string) *ListIndexDocumentsRequest
	GetEnableNameLike() *string
	SetIndexId(v string) *ListIndexDocumentsRequest
	GetIndexId() *string
	SetPageNumber(v int32) *ListIndexDocumentsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIndexDocumentsRequest
	GetPageSize() *int32
}

type ListIndexDocumentsRequest struct {
	// The names of the queried documents. The default value is null, which means the names are not used to filter the results.
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// The import status of the documents to be queried. Valid values:
	//
	// 	- INSERT_ERROR
	//
	// 	- RUNNING
	//
	// 	- DELETED
	//
	// 	- FINISH
	//
	// The default value is null, which means the import status is not used to filter the results.
	//
	// example:
	//
	// FINISH
	DocumentStatus *string `json:"DocumentStatus,omitempty" xml:"DocumentStatus,omitempty"`
	EnableNameLike *string `json:"EnableNameLike,omitempty" xml:"EnableNameLike,omitempty"`
	// The primary key ID of the knowledge base, which is the `Data.Id` parameter returned by the [CreateIndex](https://www.alibabacloud.com/help/en/model-studio/developer-reference/api-bailian-2023-12-29-createindex) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0aly8zw
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The page numbers of the pages to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of documents displayed on each page. No maximum value. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIndexDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIndexDocumentsRequest) GoString() string {
	return s.String()
}

func (s *ListIndexDocumentsRequest) GetDocumentName() *string {
	return s.DocumentName
}

func (s *ListIndexDocumentsRequest) GetDocumentStatus() *string {
	return s.DocumentStatus
}

func (s *ListIndexDocumentsRequest) GetEnableNameLike() *string {
	return s.EnableNameLike
}

func (s *ListIndexDocumentsRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *ListIndexDocumentsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIndexDocumentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIndexDocumentsRequest) SetDocumentName(v string) *ListIndexDocumentsRequest {
	s.DocumentName = &v
	return s
}

func (s *ListIndexDocumentsRequest) SetDocumentStatus(v string) *ListIndexDocumentsRequest {
	s.DocumentStatus = &v
	return s
}

func (s *ListIndexDocumentsRequest) SetEnableNameLike(v string) *ListIndexDocumentsRequest {
	s.EnableNameLike = &v
	return s
}

func (s *ListIndexDocumentsRequest) SetIndexId(v string) *ListIndexDocumentsRequest {
	s.IndexId = &v
	return s
}

func (s *ListIndexDocumentsRequest) SetPageNumber(v int32) *ListIndexDocumentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIndexDocumentsRequest) SetPageSize(v int32) *ListIndexDocumentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIndexDocumentsRequest) Validate() error {
	return dara.Validate(s)
}
