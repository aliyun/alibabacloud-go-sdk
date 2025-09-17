// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTransportLayerApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMatchType(v string) *ListTransportLayerApplicationsRequest
	GetMatchType() *string
	SetPageNumber(v int32) *ListTransportLayerApplicationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTransportLayerApplicationsRequest
	GetPageSize() *int32
	SetRecordName(v string) *ListTransportLayerApplicationsRequest
	GetRecordName() *string
	SetSiteId(v int64) *ListTransportLayerApplicationsRequest
	GetSiteId() *int64
}

type ListTransportLayerApplicationsRequest struct {
	// example:
	//
	// fuzzy
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 500
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// www.example.com
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456******
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListTransportLayerApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTransportLayerApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListTransportLayerApplicationsRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *ListTransportLayerApplicationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTransportLayerApplicationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTransportLayerApplicationsRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *ListTransportLayerApplicationsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListTransportLayerApplicationsRequest) SetMatchType(v string) *ListTransportLayerApplicationsRequest {
	s.MatchType = &v
	return s
}

func (s *ListTransportLayerApplicationsRequest) SetPageNumber(v int32) *ListTransportLayerApplicationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTransportLayerApplicationsRequest) SetPageSize(v int32) *ListTransportLayerApplicationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTransportLayerApplicationsRequest) SetRecordName(v string) *ListTransportLayerApplicationsRequest {
	s.RecordName = &v
	return s
}

func (s *ListTransportLayerApplicationsRequest) SetSiteId(v int64) *ListTransportLayerApplicationsRequest {
	s.SiteId = &v
	return s
}

func (s *ListTransportLayerApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
