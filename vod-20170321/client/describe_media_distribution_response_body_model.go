// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMediaDistributionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaDistributionList(v []*DescribeMediaDistributionResponseBodyMediaDistributionList) *DescribeMediaDistributionResponseBody
	GetMediaDistributionList() []*DescribeMediaDistributionResponseBodyMediaDistributionList
	SetRequestId(v string) *DescribeMediaDistributionResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeMediaDistributionResponseBody
	GetTotal() *int64
}

type DescribeMediaDistributionResponseBody struct {
	// The distribution list of media assets. The data is displayed based on the statistical cycle of the natural hour, day, week, or month of the start and end time.
	MediaDistributionList []*DescribeMediaDistributionResponseBodyMediaDistributionList `json:"MediaDistributionList,omitempty" xml:"MediaDistributionList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of media assets returned.
	//
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMediaDistributionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMediaDistributionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMediaDistributionResponseBody) GetMediaDistributionList() []*DescribeMediaDistributionResponseBodyMediaDistributionList {
	return s.MediaDistributionList
}

func (s *DescribeMediaDistributionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMediaDistributionResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeMediaDistributionResponseBody) SetMediaDistributionList(v []*DescribeMediaDistributionResponseBodyMediaDistributionList) *DescribeMediaDistributionResponseBody {
	s.MediaDistributionList = v
	return s
}

func (s *DescribeMediaDistributionResponseBody) SetRequestId(v string) *DescribeMediaDistributionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMediaDistributionResponseBody) SetTotal(v int64) *DescribeMediaDistributionResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeMediaDistributionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMediaDistributionResponseBodyMediaDistributionList struct {
	// The number of media assets that are queried during the specified time range.
	//
	// example:
	//
	// 12
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The end of the time range during which data is queried (exclusive). The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-14T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start of the time range during which data is queried (inclusive). The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-11-13T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMediaDistributionResponseBodyMediaDistributionList) String() string {
	return dara.Prettify(s)
}

func (s DescribeMediaDistributionResponseBodyMediaDistributionList) GoString() string {
	return s.String()
}

func (s *DescribeMediaDistributionResponseBodyMediaDistributionList) GetCount() *int64 {
	return s.Count
}

func (s *DescribeMediaDistributionResponseBodyMediaDistributionList) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMediaDistributionResponseBodyMediaDistributionList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMediaDistributionResponseBodyMediaDistributionList) SetCount(v int64) *DescribeMediaDistributionResponseBodyMediaDistributionList {
	s.Count = &v
	return s
}

func (s *DescribeMediaDistributionResponseBodyMediaDistributionList) SetEndTime(v string) *DescribeMediaDistributionResponseBodyMediaDistributionList {
	s.EndTime = &v
	return s
}

func (s *DescribeMediaDistributionResponseBodyMediaDistributionList) SetStartTime(v string) *DescribeMediaDistributionResponseBodyMediaDistributionList {
	s.StartTime = &v
	return s
}

func (s *DescribeMediaDistributionResponseBodyMediaDistributionList) Validate() error {
	return dara.Validate(s)
}
