// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileKey(v string) *GetDocumentDownloadUrlRequest
	GetFileKey() *string
	SetId(v int64) *GetDocumentDownloadUrlRequest
	GetId() *int64
	SetReportType(v string) *GetDocumentDownloadUrlRequest
	GetReportType() *string
}

type GetDocumentDownloadUrlRequest struct {
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// Document management ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 175815
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Report type.
	//
	// example:
	//
	// 5
	ReportType *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
}

func (s GetDocumentDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentDownloadUrlRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *GetDocumentDownloadUrlRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDocumentDownloadUrlRequest) GetReportType() *string {
	return s.ReportType
}

func (s *GetDocumentDownloadUrlRequest) SetFileKey(v string) *GetDocumentDownloadUrlRequest {
	s.FileKey = &v
	return s
}

func (s *GetDocumentDownloadUrlRequest) SetId(v int64) *GetDocumentDownloadUrlRequest {
	s.Id = &v
	return s
}

func (s *GetDocumentDownloadUrlRequest) SetReportType(v string) *GetDocumentDownloadUrlRequest {
	s.ReportType = &v
	return s
}

func (s *GetDocumentDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
