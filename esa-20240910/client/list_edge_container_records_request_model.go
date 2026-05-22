// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListEdgeContainerRecordsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListEdgeContainerRecordsRequest
	GetPageSize() *int32
	SetRecordMatchType(v string) *ListEdgeContainerRecordsRequest
	GetRecordMatchType() *string
	SetRecordName(v string) *ListEdgeContainerRecordsRequest
	GetRecordName() *string
	SetSiteId(v int64) *ListEdgeContainerRecordsRequest
	GetSiteId() *int64
}

type ListEdgeContainerRecordsRequest struct {
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RecordMatchType *string `json:"RecordMatchType,omitempty" xml:"RecordMatchType,omitempty"`
	RecordName      *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListEdgeContainerRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerRecordsRequest) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerRecordsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEdgeContainerRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEdgeContainerRecordsRequest) GetRecordMatchType() *string {
	return s.RecordMatchType
}

func (s *ListEdgeContainerRecordsRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *ListEdgeContainerRecordsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListEdgeContainerRecordsRequest) SetPageNumber(v int32) *ListEdgeContainerRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEdgeContainerRecordsRequest) SetPageSize(v int32) *ListEdgeContainerRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEdgeContainerRecordsRequest) SetRecordMatchType(v string) *ListEdgeContainerRecordsRequest {
	s.RecordMatchType = &v
	return s
}

func (s *ListEdgeContainerRecordsRequest) SetRecordName(v string) *ListEdgeContainerRecordsRequest {
	s.RecordName = &v
	return s
}

func (s *ListEdgeContainerRecordsRequest) SetSiteId(v int64) *ListEdgeContainerRecordsRequest {
	s.SiteId = &v
	return s
}

func (s *ListEdgeContainerRecordsRequest) Validate() error {
	return dara.Validate(s)
}
