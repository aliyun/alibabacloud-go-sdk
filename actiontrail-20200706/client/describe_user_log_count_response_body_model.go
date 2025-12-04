// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserLogCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeUserLogCountResponseBodyData) *DescribeUserLogCountResponseBody
	GetData() *DescribeUserLogCountResponseBodyData
	SetRequestId(v string) *DescribeUserLogCountResponseBody
	GetRequestId() *string
}

type DescribeUserLogCountResponseBody struct {
	Data *DescribeUserLogCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 90D6CC31-947F-5D8A-BEDC-F312EE9B31EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserLogCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLogCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserLogCountResponseBody) GetData() *DescribeUserLogCountResponseBodyData {
	return s.Data
}

func (s *DescribeUserLogCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserLogCountResponseBody) SetData(v *DescribeUserLogCountResponseBodyData) *DescribeUserLogCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUserLogCountResponseBody) SetRequestId(v string) *DescribeUserLogCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserLogCountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserLogCountResponseBodyData struct {
	Counts []*int64  `json:"Counts,omitempty" xml:"Counts,omitempty" type:"Repeated"`
	Dates  []*string `json:"Dates,omitempty" xml:"Dates,omitempty" type:"Repeated"`
}

func (s DescribeUserLogCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserLogCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUserLogCountResponseBodyData) GetCounts() []*int64 {
	return s.Counts
}

func (s *DescribeUserLogCountResponseBodyData) GetDates() []*string {
	return s.Dates
}

func (s *DescribeUserLogCountResponseBodyData) SetCounts(v []*int64) *DescribeUserLogCountResponseBodyData {
	s.Counts = v
	return s
}

func (s *DescribeUserLogCountResponseBodyData) SetDates(v []*string) *DescribeUserLogCountResponseBodyData {
	s.Dates = v
	return s
}

func (s *DescribeUserLogCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
