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
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **500**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The match mode to filter the record names:
	//
	// 	- **fuzzy**
	//
	// 	- **prefix**
	//
	// 	- **suffix**
	//
	// 	- **exact*	- (default)
	//
	// example:
	//
	// fuzzy
	RecordMatchType *string `json:"RecordMatchType,omitempty" xml:"RecordMatchType,omitempty"`
	// The record name.
	//
	// example:
	//
	// a.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// The website ID. You can call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
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
