// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayMetricAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetricList(v []*string) *DescribePlayMetricAuthResponseBody
	GetMetricList() []*string
	SetRequestId(v string) *DescribePlayMetricAuthResponseBody
	GetRequestId() *string
}

type DescribePlayMetricAuthResponseBody struct {
	MetricList []*string `json:"MetricList,omitempty" xml:"MetricList,omitempty" type:"Repeated"`
	RequestId  *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePlayMetricAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayMetricAuthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePlayMetricAuthResponseBody) GetMetricList() []*string {
	return s.MetricList
}

func (s *DescribePlayMetricAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePlayMetricAuthResponseBody) SetMetricList(v []*string) *DescribePlayMetricAuthResponseBody {
	s.MetricList = v
	return s
}

func (s *DescribePlayMetricAuthResponseBody) SetRequestId(v string) *DescribePlayMetricAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePlayMetricAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
