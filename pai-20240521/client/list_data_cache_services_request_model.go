// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCacheServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListDataCacheServicesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListDataCacheServicesRequest
	GetPageSize() *int64
	SetQuotaId(v string) *ListDataCacheServicesRequest
	GetQuotaId() *string
}

type ListDataCacheServicesRequest struct {
	PageNumber *int64  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QuotaId    *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
}

func (s ListDataCacheServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataCacheServicesRequest) GoString() string {
	return s.String()
}

func (s *ListDataCacheServicesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDataCacheServicesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDataCacheServicesRequest) GetQuotaId() *string {
	return s.QuotaId
}

func (s *ListDataCacheServicesRequest) SetPageNumber(v int64) *ListDataCacheServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataCacheServicesRequest) SetPageSize(v int64) *ListDataCacheServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataCacheServicesRequest) SetQuotaId(v string) *ListDataCacheServicesRequest {
	s.QuotaId = &v
	return s
}

func (s *ListDataCacheServicesRequest) Validate() error {
	return dara.Validate(s)
}
