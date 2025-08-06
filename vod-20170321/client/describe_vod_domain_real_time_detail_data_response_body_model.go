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
	// The returned results.
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
	// The ID of the request.
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
