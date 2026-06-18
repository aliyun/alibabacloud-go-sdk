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
	// The name of the documents to return. If you do not specify this parameter, the results are not filtered by name.
	//
	// example:
	//
	// 翻译平台运维文档
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// The import status of the documents to return. Valid values:
	//
	// - INSERT_ERROR: The document failed to be imported.
	//
	// - RUNNING: The document is being imported.
	//
	// - DELETED: The document has been deleted.
	//
	// - FINISH: The document was imported successfully.
	//
	// If you do not specify this parameter, the results are not filtered by import status.
	//
	// example:
	//
	// FINISH
	DocumentStatus *string `json:"DocumentStatus,omitempty" xml:"DocumentStatus,omitempty"`
	// Specifies whether to perform a fuzzy search based on the document name. This parameter is used with the `DocumentName` parameter. Valid values:
	//
	// - true: Performs a fuzzy search based on the document name.
	//
	// - false: Performs an exact match based on the document name.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableNameLike *string `json:"EnableNameLike,omitempty" xml:"EnableNameLike,omitempty"`
	// The ID of the knowledge base. This is the value of the `Data.Id` parameter returned by the **CreateIndex*	- operation.
	//
	// example:
	//
	// 79c0alxxxx
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The number of the page to return. The value starts from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of documents to return on each page. Maximum value: 10.
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
