// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterPrometheusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrometheus(v string) *DescribeClusterPrometheusResponseBody
	GetPrometheus() *string
	SetRequestId(v string) *DescribeClusterPrometheusResponseBody
	GetRequestId() *string
}

type DescribeClusterPrometheusResponseBody struct {
	// The public endpoint of the Prometheus service that is used to monitor a cluster in the ASM instance.
	//
	// example:
	//
	// p.com
	Prometheus *string `json:"Prometheus,omitempty" xml:"Prometheus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BD65C0AD-D3C6-48D3-8D93-38D2015C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterPrometheusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterPrometheusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterPrometheusResponseBody) GetPrometheus() *string {
	return s.Prometheus
}

func (s *DescribeClusterPrometheusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterPrometheusResponseBody) SetPrometheus(v string) *DescribeClusterPrometheusResponseBody {
	s.Prometheus = &v
	return s
}

func (s *DescribeClusterPrometheusResponseBody) SetRequestId(v string) *DescribeClusterPrometheusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterPrometheusResponseBody) Validate() error {
	return dara.Validate(s)
}
