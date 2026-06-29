// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeRoutineRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListEdgeRoutineRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeRoutineRecordsRequest
	GetPageSize() *int32
	SetRecordMatchType(v string) *ListEdgeRoutineRecordsRequest
	GetRecordMatchType() *string
	SetRecordName(v string) *ListEdgeRoutineRecordsRequest
	GetRecordName() *string
	SetSiteId(v int64) *ListEdgeRoutineRecordsRequest
	GetSiteId() *int64
}

type ListEdgeRoutineRecordsRequest struct {
	// The page number for a paged query. The value must be greater than or equal to 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page for a paged query. Valid values: 1 to 500. Default value: **500**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The matching mode used to filter by record name. Valid values:
	//
	// - **fuzzy**: fuzzy match.
	//
	// - **prefix**: prefix match.
	//
	// - **suffix**: suffix match.
	//
	// - **exact**: exact match (default).
	//
	// example:
	//
	// fuzzy
	RecordMatchType *string `json:"RecordMatchType,omitempty" xml:"RecordMatchType,omitempty"`
	// Filters by the specified record name.
	//
	// example:
	//
	// a.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The site ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListEdgeRoutineRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeRoutineRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeRoutineRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeRoutineRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeRoutineRecordsRequest) GetRecordMatchType() *string {
	return s.RecordMatchType
}

func (s *ListEdgeRoutineRecordsRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *ListEdgeRoutineRecordsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListEdgeRoutineRecordsRequest) SetPageNumber(v int32) *ListEdgeRoutineRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeRoutineRecordsRequest) SetPageSize(v int32) *ListEdgeRoutineRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEdgeRoutineRecordsRequest) SetRecordMatchType(v string) *ListEdgeRoutineRecordsRequest {
	s.RecordMatchType = &v
	return s
}

func (s *ListEdgeRoutineRecordsRequest) SetRecordName(v string) *ListEdgeRoutineRecordsRequest {
	s.RecordName = &v
	return s
}

func (s *ListEdgeRoutineRecordsRequest) SetSiteId(v int64) *ListEdgeRoutineRecordsRequest {
	s.SiteId = &v
	return s
}

func (s *ListEdgeRoutineRecordsRequest) Validate() error {
	return dara.Validate(s)
}
