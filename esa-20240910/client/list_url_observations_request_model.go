// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUrlObservationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *ListUrlObservationsRequest
	GetConfigId() *int64
	SetPageNumber(v int32) *ListUrlObservationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUrlObservationsRequest
	GetPageSize() *int32
	SetSiteId(v int64) *ListUrlObservationsRequest
	GetSiteId() *int64
}

type ListUrlObservationsRequest struct {
	// The configuration ID,
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The page number. Page starts from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **500**. Valid values: **1 to 500**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 34003500310****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s ListUrlObservationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUrlObservationsRequest) GoString() string {
	return s.String()
}

func (s *ListUrlObservationsRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListUrlObservationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUrlObservationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUrlObservationsRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ListUrlObservationsRequest) SetConfigId(v int64) *ListUrlObservationsRequest {
	s.ConfigId = &v
	return s
}

func (s *ListUrlObservationsRequest) SetPageNumber(v int32) *ListUrlObservationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUrlObservationsRequest) SetPageSize(v int32) *ListUrlObservationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListUrlObservationsRequest) SetSiteId(v int64) *ListUrlObservationsRequest {
	s.SiteId = &v
	return s
}

func (s *ListUrlObservationsRequest) Validate() error {
	return dara.Validate(s)
}
