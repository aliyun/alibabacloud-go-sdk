// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceExportTaskVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportVersion(v string) *ListResourceExportTaskVersionsRequest
	GetExportVersion() *string
	SetKeyword(v string) *ListResourceExportTaskVersionsRequest
	GetKeyword() *string
	SetPageNumber(v int32) *ListResourceExportTaskVersionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceExportTaskVersionsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListResourceExportTaskVersionsRequest
	GetStatus() *string
}

type ListResourceExportTaskVersionsRequest struct {
	// The export version number.
	//
	// example:
	//
	// v1
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// The search keyword. Fuzzy match is supported for export version names.
	//
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of results per page. Default value: 20. Minimum value: 1. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The export status. Valid values:
	//
	// - Queue: queued
	//
	// - Pending: preparing to run
	//
	// - Success: succeeded
	//
	// - Errored: failed
	//
	// - Canceled: canceled.
	//
	// example:
	//
	// Errored
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListResourceExportTaskVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTaskVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsRequest) GetExportVersion() *string {
	return s.ExportVersion
}

func (s *ListResourceExportTaskVersionsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListResourceExportTaskVersionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceExportTaskVersionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceExportTaskVersionsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListResourceExportTaskVersionsRequest) SetExportVersion(v string) *ListResourceExportTaskVersionsRequest {
	s.ExportVersion = &v
	return s
}

func (s *ListResourceExportTaskVersionsRequest) SetKeyword(v string) *ListResourceExportTaskVersionsRequest {
	s.Keyword = &v
	return s
}

func (s *ListResourceExportTaskVersionsRequest) SetPageNumber(v int32) *ListResourceExportTaskVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceExportTaskVersionsRequest) SetPageSize(v int32) *ListResourceExportTaskVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourceExportTaskVersionsRequest) SetStatus(v string) *ListResourceExportTaskVersionsRequest {
	s.Status = &v
	return s
}

func (s *ListResourceExportTaskVersionsRequest) Validate() error {
	return dara.Validate(s)
}
