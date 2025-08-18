// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheReserveSpecificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCacheReserveCapacity(v []*string) *GetCacheReserveSpecificationResponseBody
	GetCacheReserveCapacity() []*string
	SetCacheReserveRegion(v []*string) *GetCacheReserveSpecificationResponseBody
	GetCacheReserveRegion() []*string
	SetRequestId(v string) *GetCacheReserveSpecificationResponseBody
	GetRequestId() *string
}

type GetCacheReserveSpecificationResponseBody struct {
	// List of cache retention capacity specifications.
	CacheReserveCapacity []*string `json:"CacheReserveCapacity,omitempty" xml:"CacheReserveCapacity,omitempty" type:"Repeated"`
	// List of cache retention region specifications.
	CacheReserveRegion []*string `json:"CacheReserveRegion,omitempty" xml:"CacheReserveRegion,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCacheReserveSpecificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCacheReserveSpecificationResponseBody) GoString() string {
	return s.String()
}

func (s *GetCacheReserveSpecificationResponseBody) GetCacheReserveCapacity() []*string {
	return s.CacheReserveCapacity
}

func (s *GetCacheReserveSpecificationResponseBody) GetCacheReserveRegion() []*string {
	return s.CacheReserveRegion
}

func (s *GetCacheReserveSpecificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCacheReserveSpecificationResponseBody) SetCacheReserveCapacity(v []*string) *GetCacheReserveSpecificationResponseBody {
	s.CacheReserveCapacity = v
	return s
}

func (s *GetCacheReserveSpecificationResponseBody) SetCacheReserveRegion(v []*string) *GetCacheReserveSpecificationResponseBody {
	s.CacheReserveRegion = v
	return s
}

func (s *GetCacheReserveSpecificationResponseBody) SetRequestId(v string) *GetCacheReserveSpecificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCacheReserveSpecificationResponseBody) Validate() error {
	return dara.Validate(s)
}
