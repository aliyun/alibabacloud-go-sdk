// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexFileDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocumentName(v string) *ListIndexFileDetailsRequest
	GetDocumentName() *string
	SetDocumentStatus(v string) *ListIndexFileDetailsRequest
	GetDocumentStatus() *string
	SetEnableNameLike(v string) *ListIndexFileDetailsRequest
	GetEnableNameLike() *string
	SetIndexId(v string) *ListIndexFileDetailsRequest
	GetIndexId() *string
	SetPageNumber(v int32) *ListIndexFileDetailsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIndexFileDetailsRequest
	GetPageSize() *int32
}

type ListIndexFileDetailsRequest struct {
	// Filters the returned file details list by file name. Default value: empty, which means no filtering by file name.
	//
	// example:
	//
	// TranslationPlatformO&MDocument.
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// Filters the returned file list by file import status. Valid values:
	//
	// - INSERT_ERROR: Import to index failed.
	//
	// - RUNNING: Index building in progress.
	//
	// - DELETED: Deleted.
	//
	// - FINISH: Index building succeeded.
	//
	// - PARSE_FAILED: Parsing failed.
	//
	// - DOC_PARSING: Parsing in progress.
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
	// example:
	//
	// 79c0alxxxx
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The page number to query. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of files to display per page in a paging query. Maximum value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListIndexFileDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIndexFileDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListIndexFileDetailsRequest) GetDocumentName() *string {
	return s.DocumentName
}

func (s *ListIndexFileDetailsRequest) GetDocumentStatus() *string {
	return s.DocumentStatus
}

func (s *ListIndexFileDetailsRequest) GetEnableNameLike() *string {
	return s.EnableNameLike
}

func (s *ListIndexFileDetailsRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *ListIndexFileDetailsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIndexFileDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIndexFileDetailsRequest) SetDocumentName(v string) *ListIndexFileDetailsRequest {
	s.DocumentName = &v
	return s
}

func (s *ListIndexFileDetailsRequest) SetDocumentStatus(v string) *ListIndexFileDetailsRequest {
	s.DocumentStatus = &v
	return s
}

func (s *ListIndexFileDetailsRequest) SetEnableNameLike(v string) *ListIndexFileDetailsRequest {
	s.EnableNameLike = &v
	return s
}

func (s *ListIndexFileDetailsRequest) SetIndexId(v string) *ListIndexFileDetailsRequest {
	s.IndexId = &v
	return s
}

func (s *ListIndexFileDetailsRequest) SetPageNumber(v int32) *ListIndexFileDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListIndexFileDetailsRequest) SetPageSize(v int32) *ListIndexFileDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIndexFileDetailsRequest) Validate() error {
	return dara.Validate(s)
}
