// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRangeDataByLocateAndIspServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJsonResult(v string) *DescribeVodRangeDataByLocateAndIspServiceResponseBody
	GetJsonResult() *string
	SetRequestId(v string) *DescribeVodRangeDataByLocateAndIspServiceResponseBody
	GetRequestId() *string
}

type DescribeVodRangeDataByLocateAndIspServiceResponseBody struct {
	// The returned result. The value is in the JSON format. These parameters indicate the following information in sequence: UNIX time, region, ISP, distribution of HTTP status codes, response time, bandwidth (bit/s), average response rate, page views, cache hit ratio, and request hit ratio.
	//
	// example:
	//
	// {"1472659200":{"tianjin":{"unicom":{"http_codes":{"000":0,"200":6,"400":0},"rt":4183,"bandwidth":46639,"avg_speed":7773,"pv":6,"hit_rate":0.93,"request_hit_rate":0.66}}}}
	JsonResult *string `json:"JsonResult,omitempty" xml:"JsonResult,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C74802AA-C277-5A80-BDF2-072B05F119C7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodRangeDataByLocateAndIspServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRangeDataByLocateAndIspServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodRangeDataByLocateAndIspServiceResponseBody) GetJsonResult() *string {
	return s.JsonResult
}

func (s *DescribeVodRangeDataByLocateAndIspServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodRangeDataByLocateAndIspServiceResponseBody) SetJsonResult(v string) *DescribeVodRangeDataByLocateAndIspServiceResponseBody {
	s.JsonResult = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceResponseBody) SetRequestId(v string) *DescribeVodRangeDataByLocateAndIspServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodRangeDataByLocateAndIspServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
