// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAlertCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeUserAlertCountResponseBodyData) *DescribeUserAlertCountResponseBody
	GetData() *DescribeUserAlertCountResponseBodyData
	SetRequestId(v string) *DescribeUserAlertCountResponseBody
	GetRequestId() *string
}

type DescribeUserAlertCountResponseBody struct {
	// The data returned.
	Data *DescribeUserAlertCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 90D6CC31-947F-5D8A-BEDC-F312EE9B31EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserAlertCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAlertCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserAlertCountResponseBody) GetData() *DescribeUserAlertCountResponseBodyData {
	return s.Data
}

func (s *DescribeUserAlertCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserAlertCountResponseBody) SetData(v *DescribeUserAlertCountResponseBodyData) *DescribeUserAlertCountResponseBody {
	s.Data = v
	return s
}

func (s *DescribeUserAlertCountResponseBody) SetRequestId(v string) *DescribeUserAlertCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserAlertCountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserAlertCountResponseBodyData struct {
	// The statistics returned.
	Counts []*int64 `json:"Counts,omitempty" xml:"Counts,omitempty" type:"Repeated"`
	// The dates of alerts.
	Dates []*string `json:"Dates,omitempty" xml:"Dates,omitempty" type:"Repeated"`
}

func (s DescribeUserAlertCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAlertCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeUserAlertCountResponseBodyData) GetCounts() []*int64 {
	return s.Counts
}

func (s *DescribeUserAlertCountResponseBodyData) GetDates() []*string {
	return s.Dates
}

func (s *DescribeUserAlertCountResponseBodyData) SetCounts(v []*int64) *DescribeUserAlertCountResponseBodyData {
	s.Counts = v
	return s
}

func (s *DescribeUserAlertCountResponseBodyData) SetDates(v []*string) *DescribeUserAlertCountResponseBodyData {
	s.Dates = v
	return s
}

func (s *DescribeUserAlertCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
