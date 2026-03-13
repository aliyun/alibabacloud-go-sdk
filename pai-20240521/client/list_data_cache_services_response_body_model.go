// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCacheServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCacheServices(v []*CacheService) *ListDataCacheServicesResponseBody
	GetCacheServices() []*CacheService
	SetRequestId(v string) *ListDataCacheServicesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDataCacheServicesResponseBody
	GetTotalCount() *int64
}

type ListDataCacheServicesResponseBody struct {
	CacheServices []*CacheService `json:"CacheServices,omitempty" xml:"CacheServices,omitempty" type:"Repeated"`
	RequestId     *string         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount    *int64          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataCacheServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataCacheServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCacheServicesResponseBody) GetCacheServices() []*CacheService {
	return s.CacheServices
}

func (s *ListDataCacheServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataCacheServicesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDataCacheServicesResponseBody) SetCacheServices(v []*CacheService) *ListDataCacheServicesResponseBody {
	s.CacheServices = v
	return s
}

func (s *ListDataCacheServicesResponseBody) SetRequestId(v string) *ListDataCacheServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCacheServicesResponseBody) SetTotalCount(v int64) *ListDataCacheServicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataCacheServicesResponseBody) Validate() error {
	if s.CacheServices != nil {
		for _, item := range s.CacheServices {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
