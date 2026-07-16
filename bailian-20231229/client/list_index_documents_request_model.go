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
	// Filters the returned file list by file name (without the file extension). Default value: empty, which means no filtering by file name.
	//
	// example:
	//
	// product-overview
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// Filters the returned file list by file import status. Valid values:
	//
	// - INSERT_ERROR: failed to import to the index.
	//
	// - RUNNING: index building in progress.
	//
	// - DELETED: deleted.
	//
	// - FINISH: index building succeeded.
	//
	// - PARSE_FAILED: parsing failed.
	//
	// - DOC_PARSING: parsing in progress.
	//
	// Default value: empty, which means no filtering by file import status.
	//
	// example:
	//
	// FINISH
	DocumentStatus *string `json:"DocumentStatus,omitempty" xml:"DocumentStatus,omitempty"`
	// Specifies whether to enable fuzzy matching for file names. This parameter is used together with the `DocumentName` parameter. Valid values:
	//
	// - true: Performs fuzzy matching on the returned file list based on the file name.
	//
	// - false: Performs exact matching on the returned file list based on the file name.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableNameLike *string `json:"EnableNameLike,omitempty" xml:"EnableNameLike,omitempty"`
	// The knowledge base ID, which is the `Data.Id` returned by the **CreateIndex*	- operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0alxxxx
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of files to display per page in a paged query. No maximum limit.
	//
	// Default value: 10.
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
