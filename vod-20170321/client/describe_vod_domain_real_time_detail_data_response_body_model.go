// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainRealTimeDetailDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DescribeVodDomainRealTimeDetailDataResponseBody
	GetData() *string
	SetRequestId(v string) *DescribeVodDomainRealTimeDetailDataResponseBody
	GetRequestId() *string
}

type DescribeVodDomainRealTimeDetailDataResponseBody struct {
	// The returned data details. The data is returned as a JSON string. The following table describes the structure and fields:
	//
	// > If no data exists for a field, the field is not returned.
	//
	// | Field | Type | Description |
	//
	// | ------------- |------------ | ----------- |
	//
	// | domain_name | String | The accelerated domain name. |
	//
	// | isp | String | The ISP name. |
	//
	// | location | String | The region name. |
	//
	// | qps | Long | The queries per second (QPS). |
	//
	// | bps | Long | The bandwidth data. Unit: bit/s. |
	//
	// | http_code | Map | The HTTP status code details. The key is the status code name, and the value is the count of the status code. |
	//
	// | time_stp | String | The data timestamp. The time is in the ISO 8601 standard in UTC. |
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "time_stp": "2018-06-05T20:00:00Z",
	//
	//             "domain_name": "example.com",
	//
	//             "location": "Guangdong",
	//
	//             "isp": "telecom",
	//
	//             "qps": 10
	//
	//       },
	//
	//       {
	//
	//             "time_stp": "2018-06-05T20:00:00Z",
	//
	//             "domain_name": "example.com",
	//
	//             "location": "Jiangsu",
	//
	//             "isp": "unicom",
	//
	//             "qps": 11.1
	//
	//       }
	//
	// ]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1710298E-8AFA-5F6D-A3E9-47103C52177D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainRealTimeDetailDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainRealTimeDetailDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainRealTimeDetailDataResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeVodDomainRealTimeDetailDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainRealTimeDetailDataResponseBody) SetData(v string) *DescribeVodDomainRealTimeDetailDataResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataResponseBody) SetRequestId(v string) *DescribeVodDomainRealTimeDetailDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainRealTimeDetailDataResponseBody) Validate() error {
	return dara.Validate(s)
}
