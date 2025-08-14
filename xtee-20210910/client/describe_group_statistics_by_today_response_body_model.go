// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupStatisticsByTodayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeGroupStatisticsByTodayResponseBody
	GetRequestId() *string
	SetData(v bool) *DescribeGroupStatisticsByTodayResponseBody
	GetData() *bool
}

type DescribeGroupStatisticsByTodayResponseBody struct {
	// Request ID
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned data.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DescribeGroupStatisticsByTodayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupStatisticsByTodayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupStatisticsByTodayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupStatisticsByTodayResponseBody) GetData() *bool {
	return s.Data
}

func (s *DescribeGroupStatisticsByTodayResponseBody) SetRequestId(v string) *DescribeGroupStatisticsByTodayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupStatisticsByTodayResponseBody) SetData(v bool) *DescribeGroupStatisticsByTodayResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeGroupStatisticsByTodayResponseBody) Validate() error {
	return dara.Validate(s)
}
