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
	// The distribution list of audio and video media assets. Statistics are displayed based on the statistical period (calendar hour, day, week, or month) within the specified time range.
	MediaDistributionList []*DescribeMediaDistributionResponseBodyMediaDistributionList `json:"MediaDistributionList,omitempty" xml:"MediaDistributionList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of audio and video media assets.
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
	if s.MediaDistributionList != nil {
		for _, item := range s.MediaDistributionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMediaDistributionResponseBodyMediaDistributionList struct {
	// The number of media assets that match the specified time range within the statistical period.
	//
	// example:
	//
	// 12
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The end time (exclusive) of the statistical period. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
	//
	// example:
	//
	// 2017-11-14T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time (inclusive) of the statistical period. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).
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
