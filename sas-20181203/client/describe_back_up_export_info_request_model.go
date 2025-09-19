// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackUpExportInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeBackUpExportInfoRequest
	GetCurrentPage() *int32
	SetExportType(v string) *DescribeBackUpExportInfoRequest
	GetExportType() *string
	SetLang(v string) *DescribeBackUpExportInfoRequest
	GetLang() *string
	SetPageSize(v int32) *DescribeBackUpExportInfoRequest
	GetPageSize() *int32
}

type DescribeBackUpExportInfoRequest struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of archived information. Valid values:
	//
	// 	- **suspiciousExport**: alert event
	//
	// This parameter is required.
	//
	// example:
	//
	// suspiciousExport
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeBackUpExportInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackUpExportInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackUpExportInfoRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeBackUpExportInfoRequest) GetExportType() *string {
	return s.ExportType
}

func (s *DescribeBackUpExportInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeBackUpExportInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackUpExportInfoRequest) SetCurrentPage(v int32) *DescribeBackUpExportInfoRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBackUpExportInfoRequest) SetExportType(v string) *DescribeBackUpExportInfoRequest {
	s.ExportType = &v
	return s
}

func (s *DescribeBackUpExportInfoRequest) SetLang(v string) *DescribeBackUpExportInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBackUpExportInfoRequest) SetPageSize(v int32) *DescribeBackUpExportInfoRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackUpExportInfoRequest) Validate() error {
	return dara.Validate(s)
}
