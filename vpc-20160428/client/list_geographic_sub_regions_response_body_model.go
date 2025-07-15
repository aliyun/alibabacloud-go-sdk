// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGeographicSubRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListGeographicSubRegionsResponseBody
	GetCount() *int64
	SetGeographicSubRegions(v []*string) *ListGeographicSubRegionsResponseBody
	GetGeographicSubRegions() []*string
	SetRequestId(v string) *ListGeographicSubRegionsResponseBody
	GetRequestId() *string
}

type ListGeographicSubRegionsResponseBody struct {
	// The number of entries.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The region list.
	GeographicSubRegions []*string `json:"GeographicSubRegions,omitempty" xml:"GeographicSubRegions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGeographicSubRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGeographicSubRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGeographicSubRegionsResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListGeographicSubRegionsResponseBody) GetGeographicSubRegions() []*string {
	return s.GeographicSubRegions
}

func (s *ListGeographicSubRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGeographicSubRegionsResponseBody) SetCount(v int64) *ListGeographicSubRegionsResponseBody {
	s.Count = &v
	return s
}

func (s *ListGeographicSubRegionsResponseBody) SetGeographicSubRegions(v []*string) *ListGeographicSubRegionsResponseBody {
	s.GeographicSubRegions = v
	return s
}

func (s *ListGeographicSubRegionsResponseBody) SetRequestId(v string) *ListGeographicSubRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGeographicSubRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}
