// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResolveStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataPoints(v *GetResolveStatisticsResponseBodyDataPoints) *GetResolveStatisticsResponseBody
	GetDataPoints() *GetResolveStatisticsResponseBodyDataPoints
	SetRequestId(v string) *GetResolveStatisticsResponseBody
	GetRequestId() *string
}

type GetResolveStatisticsResponseBody struct {
	DataPoints *GetResolveStatisticsResponseBodyDataPoints `json:"DataPoints,omitempty" xml:"DataPoints,omitempty" type:"Struct"`
	// example:
	//
	// 50F9C40E-188D-B00B-BE2C-7427E531****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResolveStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResolveStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetResolveStatisticsResponseBody) GetDataPoints() *GetResolveStatisticsResponseBodyDataPoints {
	return s.DataPoints
}

func (s *GetResolveStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResolveStatisticsResponseBody) SetDataPoints(v *GetResolveStatisticsResponseBodyDataPoints) *GetResolveStatisticsResponseBody {
	s.DataPoints = v
	return s
}

func (s *GetResolveStatisticsResponseBody) SetRequestId(v string) *GetResolveStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResolveStatisticsResponseBody) Validate() error {
	if s.DataPoints != nil {
		if err := s.DataPoints.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResolveStatisticsResponseBodyDataPoints struct {
	DataPoint []*GetResolveStatisticsResponseBodyDataPointsDataPoint `json:"DataPoint,omitempty" xml:"DataPoint,omitempty" type:"Repeated"`
}

func (s GetResolveStatisticsResponseBodyDataPoints) String() string {
	return dara.Prettify(s)
}

func (s GetResolveStatisticsResponseBodyDataPoints) GoString() string {
	return s.String()
}

func (s *GetResolveStatisticsResponseBodyDataPoints) GetDataPoint() []*GetResolveStatisticsResponseBodyDataPointsDataPoint {
	return s.DataPoint
}

func (s *GetResolveStatisticsResponseBodyDataPoints) SetDataPoint(v []*GetResolveStatisticsResponseBodyDataPointsDataPoint) *GetResolveStatisticsResponseBodyDataPoints {
	s.DataPoint = v
	return s
}

func (s *GetResolveStatisticsResponseBodyDataPoints) Validate() error {
	if s.DataPoint != nil {
		for _, item := range s.DataPoint {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResolveStatisticsResponseBodyDataPointsDataPoint struct {
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	Time  *int32 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetResolveStatisticsResponseBodyDataPointsDataPoint) String() string {
	return dara.Prettify(s)
}

func (s GetResolveStatisticsResponseBodyDataPointsDataPoint) GoString() string {
	return s.String()
}

func (s *GetResolveStatisticsResponseBodyDataPointsDataPoint) GetCount() *int32 {
	return s.Count
}

func (s *GetResolveStatisticsResponseBodyDataPointsDataPoint) GetTime() *int32 {
	return s.Time
}

func (s *GetResolveStatisticsResponseBodyDataPointsDataPoint) SetCount(v int32) *GetResolveStatisticsResponseBodyDataPointsDataPoint {
	s.Count = &v
	return s
}

func (s *GetResolveStatisticsResponseBodyDataPointsDataPoint) SetTime(v int32) *GetResolveStatisticsResponseBodyDataPointsDataPoint {
	s.Time = &v
	return s
}

func (s *GetResolveStatisticsResponseBodyDataPointsDataPoint) Validate() error {
	return dara.Validate(s)
}
