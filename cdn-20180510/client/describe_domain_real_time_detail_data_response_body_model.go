// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealTimeDetailDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DescribeDomainRealTimeDetailDataResponseBody
	GetData() *string
	SetRequestId(v string) *DescribeDomainRealTimeDetailDataResponseBody
	GetRequestId() *string
}

type DescribeDomainRealTimeDetailDataResponseBody struct {
	// The monitoring data of each ISP in each region.
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
	// B49E6DDA-F413-422B-B58E-2FA23F286726
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainRealTimeDetailDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealTimeDetailDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealTimeDetailDataResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeDomainRealTimeDetailDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainRealTimeDetailDataResponseBody) SetData(v string) *DescribeDomainRealTimeDetailDataResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataResponseBody) SetRequestId(v string) *DescribeDomainRealTimeDetailDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealTimeDetailDataResponseBody) Validate() error {
	return dara.Validate(s)
}
