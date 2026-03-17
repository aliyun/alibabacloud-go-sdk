// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPoints(v []*ListAccessPointsResponseBodyAccessPoints) *ListAccessPointsResponseBody
	GetAccessPoints() []*ListAccessPointsResponseBodyAccessPoints
	SetRequestId(v string) *ListAccessPointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAccessPointsResponseBody
	GetTotalCount() *int32
}

type ListAccessPointsResponseBody struct {
	// The information about the access point.
	AccessPoints []*ListAccessPointsResponseBodyAccessPoints `json:"AccessPoints,omitempty" xml:"AccessPoints,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// E26DBAAE-A796-4A48-98B4-B45AFCD1F299
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of access points.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAccessPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessPointsResponseBody) GetAccessPoints() []*ListAccessPointsResponseBodyAccessPoints {
	return s.AccessPoints
}

func (s *ListAccessPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccessPointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAccessPointsResponseBody) SetAccessPoints(v []*ListAccessPointsResponseBodyAccessPoints) *ListAccessPointsResponseBody {
	s.AccessPoints = v
	return s
}

func (s *ListAccessPointsResponseBody) SetRequestId(v string) *ListAccessPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessPointsResponseBody) SetTotalCount(v int32) *ListAccessPointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAccessPointsResponseBody) Validate() error {
	if s.AccessPoints != nil {
		for _, item := range s.AccessPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccessPointsResponseBodyAccessPoints struct {
	// The ID of the access point.
	//
	// example:
	//
	// 401
	AccessPointId *int32 `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The number of available SAG instances in the access point.
	//
	// example:
	//
	// 0
	ActiveSmartAGCount *int32 `json:"ActiveSmartAGCount,omitempty" xml:"ActiveSmartAGCount,omitempty"`
	// The number of offline SAG instances in the access point.
	//
	// example:
	//
	// 7
	InactiveSmartAGCount *int32 `json:"InactiveSmartAGCount,omitempty" xml:"InactiveSmartAGCount,omitempty"`
	// The latitude of the access point.
	//
	// example:
	//
	// 103.81****
	Latitude *string `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// The longitude of the access point.
	//
	// example:
	//
	// 1.35****
	Longitude *string `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
}

func (s ListAccessPointsResponseBodyAccessPoints) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsResponseBodyAccessPoints) GoString() string {
	return s.String()
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetAccessPointId() *int32 {
	return s.AccessPointId
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetActiveSmartAGCount() *int32 {
	return s.ActiveSmartAGCount
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetInactiveSmartAGCount() *int32 {
	return s.InactiveSmartAGCount
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetLatitude() *string {
	return s.Latitude
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetLongitude() *string {
	return s.Longitude
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetAccessPointId(v int32) *ListAccessPointsResponseBodyAccessPoints {
	s.AccessPointId = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetActiveSmartAGCount(v int32) *ListAccessPointsResponseBodyAccessPoints {
	s.ActiveSmartAGCount = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetInactiveSmartAGCount(v int32) *ListAccessPointsResponseBodyAccessPoints {
	s.InactiveSmartAGCount = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetLatitude(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.Latitude = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetLongitude(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.Longitude = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) Validate() error {
	return dara.Validate(s)
}
