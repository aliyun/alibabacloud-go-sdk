// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserTrailCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeUserTrailCountResponseBodyData) *DescribeUserTrailCountResponseBody
	GetData() *DescribeUserTrailCountResponseBodyData
	SetRequestId(v string) *DescribeUserTrailCountResponseBody
	GetRequestId() *string
}

type DescribeUserTrailCountResponseBody struct {
	// The returned data.
	Data *DescribeUserTrailCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// EDDEBA6B-FFE2-4EF6-8BAB-2A6B98DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserTrailCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTrailCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserTrailCountResponseBody) GetData() *DescribeUserTrailCountResponseBodyData {
	return s.Data
}

func (s *DescribeUserTrailCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserTrailCountResponseBody) SetData(v *DescribeUserTrailCountResponseBodyData) *DescribeUserTrailCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUserTrailCountResponseBody) SetRequestId(v string) *DescribeUserTrailCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserTrailCountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserTrailCountResponseBodyData struct {
	// An array of trail counts, where each count corresponds to a date in the `Dates` array.
	Counts []*int64 `json:"Counts,omitempty" xml:"Counts,omitempty" type:"Repeated"`
	// A list of dates.
	Dates []*string `json:"Dates,omitempty" xml:"Dates,omitempty" type:"Repeated"`
}

func (s DescribeUserTrailCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserTrailCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUserTrailCountResponseBodyData) GetCounts() []*int64 {
	return s.Counts
}

func (s *DescribeUserTrailCountResponseBodyData) GetDates() []*string {
	return s.Dates
}

func (s *DescribeUserTrailCountResponseBodyData) SetCounts(v []*int64) *DescribeUserTrailCountResponseBodyData {
	s.Counts = v
	return s
}

func (s *DescribeUserTrailCountResponseBodyData) SetDates(v []*string) *DescribeUserTrailCountResponseBodyData {
	s.Dates = v
	return s
}

func (s *DescribeUserTrailCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
