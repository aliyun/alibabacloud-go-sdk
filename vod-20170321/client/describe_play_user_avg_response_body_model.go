// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayUserAvgResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePlayUserAvgResponseBody
	GetRequestId() *string
	SetUserPlayStatisAvgs(v *DescribePlayUserAvgResponseBodyUserPlayStatisAvgs) *DescribePlayUserAvgResponseBody
	GetUserPlayStatisAvgs() *DescribePlayUserAvgResponseBodyUserPlayStatisAvgs
}

type DescribePlayUserAvgResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6C7F90B2-BDA4-4FAC-****-A38A121DFE19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics on average playback each day.
	UserPlayStatisAvgs *DescribePlayUserAvgResponseBodyUserPlayStatisAvgs `json:"UserPlayStatisAvgs,omitempty" xml:"UserPlayStatisAvgs,omitempty" type:"Struct"`
}

func (s DescribePlayUserAvgResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayUserAvgResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayUserAvgResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlayUserAvgResponseBody) GetUserPlayStatisAvgs() *DescribePlayUserAvgResponseBodyUserPlayStatisAvgs {
	return s.UserPlayStatisAvgs
}

func (s *DescribePlayUserAvgResponseBody) SetRequestId(v string) *DescribePlayUserAvgResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayUserAvgResponseBody) SetUserPlayStatisAvgs(v *DescribePlayUserAvgResponseBodyUserPlayStatisAvgs) *DescribePlayUserAvgResponseBody {
	s.UserPlayStatisAvgs = v
	return s
}

func (s *DescribePlayUserAvgResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePlayUserAvgResponseBodyUserPlayStatisAvgs struct {
	UserPlayStatisAvg []*DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg `json:"UserPlayStatisAvg,omitempty" xml:"UserPlayStatisAvg,omitempty" type:"Repeated"`
}

func (s DescribePlayUserAvgResponseBodyUserPlayStatisAvgs) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayUserAvgResponseBodyUserPlayStatisAvgs) GoString() string {
	return s.String()
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgs) GetUserPlayStatisAvg() []*DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg {
	return s.UserPlayStatisAvg
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgs) SetUserPlayStatisAvg(v []*DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) *DescribePlayUserAvgResponseBodyUserPlayStatisAvgs {
	s.UserPlayStatisAvg = v
	return s
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgs) Validate() error {
	return dara.Validate(s)
}

type DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg struct {
	// The average number of video views.
	//
	// example:
	//
	// 170
	AvgPlayCount *string `json:"AvgPlayCount,omitempty" xml:"AvgPlayCount,omitempty"`
	// The average playback duration. Unit: milliseconds.
	//
	// example:
	//
	// 1035902.8
	AvgPlayDuration *string `json:"AvgPlayDuration,omitempty" xml:"AvgPlayDuration,omitempty"`
	// The date when the statistics were generated. The date follows the *yyyy-MM-dd	- format.
	//
	// example:
	//
	// 20170120
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) GoString() string {
	return s.String()
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) GetAvgPlayCount() *string {
	return s.AvgPlayCount
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) GetAvgPlayDuration() *string {
	return s.AvgPlayDuration
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) GetDate() *string {
	return s.Date
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) SetAvgPlayCount(v string) *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg {
	s.AvgPlayCount = &v
	return s
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) SetAvgPlayDuration(v string) *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg {
	s.AvgPlayDuration = &v
	return s
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) SetDate(v string) *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg {
	s.Date = &v
	return s
}

func (s *DescribePlayUserAvgResponseBodyUserPlayStatisAvgsUserPlayStatisAvg) Validate() error {
	return dara.Validate(s)
}
