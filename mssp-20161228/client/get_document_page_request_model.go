// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetDocumentPageRequest
	GetCurrentPage() *int32
	SetDeliveredBy(v string) *GetDocumentPageRequest
	GetDeliveredBy() *string
	SetDocumentName(v string) *GetDocumentPageRequest
	GetDocumentName() *string
	SetDocumentType(v string) *GetDocumentPageRequest
	GetDocumentType() *string
	SetPageSize(v int32) *GetDocumentPageRequest
	GetPageSize() *int32
	SetReportType(v string) *GetDocumentPageRequest
	GetReportType() *string
}

type GetDocumentPageRequest struct {
	// Current page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Delivered by.
	//
	// example:
	//
	// luna
	DeliveredBy *string `json:"DeliveredBy,omitempty" xml:"DeliveredBy,omitempty"`
	// Document name.
	//
	// example:
	//
	// month report
	DocumentName *string `json:"DocumentName,omitempty" xml:"DocumentName,omitempty"`
	// Document type.
	//
	// example:
	//
	// 0
	DocumentType *string `json:"DocumentType,omitempty" xml:"DocumentType,omitempty"`
	// Page size.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Report type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s GetDocumentPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentPageRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetDocumentPageRequest) GetDeliveredBy() *string {
	return s.DeliveredBy
}

func (s *GetDocumentPageRequest) GetDocumentName() *string {
	return s.DocumentName
}

func (s *GetDocumentPageRequest) GetDocumentType() *string {
	return s.DocumentType
}

func (s *GetDocumentPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDocumentPageRequest) GetReportType() *string {
	return s.ReportType
}

func (s *GetDocumentPageRequest) SetCurrentPage(v int32) *GetDocumentPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetDocumentPageRequest) SetDeliveredBy(v string) *GetDocumentPageRequest {
	s.DeliveredBy = &v
	return s
}

func (s *GetDocumentPageRequest) SetDocumentName(v string) *GetDocumentPageRequest {
	s.DocumentName = &v
	return s
}

func (s *GetDocumentPageRequest) SetDocumentType(v string) *GetDocumentPageRequest {
	s.DocumentType = &v
	return s
}

func (s *GetDocumentPageRequest) SetPageSize(v int32) *GetDocumentPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetDocumentPageRequest) SetReportType(v string) *GetDocumentPageRequest {
	s.ReportType = &v
	return s
}

func (s *GetDocumentPageRequest) Validate() error {
	return dara.Validate(s)
}
