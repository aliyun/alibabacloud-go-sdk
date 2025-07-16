// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRangeDataByLocateAndIspServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonResult(v string) *DescribeRangeDataByLocateAndIspServiceResponseBody
	GetJsonResult() *string
	SetRequestId(v string) *DescribeRangeDataByLocateAndIspServiceResponseBody
	GetRequestId() *string
}

type DescribeRangeDataByLocateAndIspServiceResponseBody struct {
	// The response parameters in the JSON format. These parameters indicate the following information in sequence: UNIX time, region, ISP, distribution of HTTP status codes, response time, bandwidth (bit/s), average response rate, page views, cache hit ratio, and request hit ratio.
	//
	// example:
	//
	// {"1472659200":{"Tianjin":{"China Telecom":{"http_codes":{"000":0,"200":6,"400":0},"rt":4183,"bandwidth":46639,"avg_speed":7773,"pv":6,"hit_rate":0.93,"request_hit_rate":0.66}}}}
	JsonResult *string `json:"JsonResult,omitempty" xml:"JsonResult,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRangeDataByLocateAndIspServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRangeDataByLocateAndIspServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRangeDataByLocateAndIspServiceResponseBody) GetJsonResult() *string {
	return s.JsonResult
}

func (s *DescribeRangeDataByLocateAndIspServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRangeDataByLocateAndIspServiceResponseBody) SetJsonResult(v string) *DescribeRangeDataByLocateAndIspServiceResponseBody {
	s.JsonResult = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceResponseBody) SetRequestId(v string) *DescribeRangeDataByLocateAndIspServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRangeDataByLocateAndIspServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
