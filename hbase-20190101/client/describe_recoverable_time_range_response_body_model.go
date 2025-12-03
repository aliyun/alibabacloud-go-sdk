// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecoverableTimeRangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRecoverableTimeRangeResponseBody
	GetRequestId() *string
	SetTimeBegin(v string) *DescribeRecoverableTimeRangeResponseBody
	GetTimeBegin() *string
	SetTimeEnd(v string) *DescribeRecoverableTimeRangeResponseBody
	GetTimeEnd() *string
}

type DescribeRecoverableTimeRangeResponseBody struct {
	// example:
	//
	// A1A51D18-96DC-465C-9F1B-47180CA22524
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2020-10-26T18:02:03Z
	TimeBegin *string `json:"TimeBegin,omitempty" xml:"TimeBegin,omitempty"`
	// example:
	//
	// 2020-11-05T01:20:31Z
	TimeEnd *string `json:"TimeEnd,omitempty" xml:"TimeEnd,omitempty"`
}

func (s DescribeRecoverableTimeRangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecoverableTimeRangeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableTimeRangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecoverableTimeRangeResponseBody) GetTimeBegin() *string {
	return s.TimeBegin
}

func (s *DescribeRecoverableTimeRangeResponseBody) GetTimeEnd() *string {
	return s.TimeEnd
}

func (s *DescribeRecoverableTimeRangeResponseBody) SetRequestId(v string) *DescribeRecoverableTimeRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecoverableTimeRangeResponseBody) SetTimeBegin(v string) *DescribeRecoverableTimeRangeResponseBody {
	s.TimeBegin = &v
	return s
}

func (s *DescribeRecoverableTimeRangeResponseBody) SetTimeEnd(v string) *DescribeRecoverableTimeRangeResponseBody {
	s.TimeEnd = &v
	return s
}

func (s *DescribeRecoverableTimeRangeResponseBody) Validate() error {
	return dara.Validate(s)
}
