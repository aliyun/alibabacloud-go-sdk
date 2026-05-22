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
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RecordMatchType *string `json:"RecordMatchType,omitempty" xml:"RecordMatchType,omitempty"`
	RecordName      *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// This parameter is required.
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
