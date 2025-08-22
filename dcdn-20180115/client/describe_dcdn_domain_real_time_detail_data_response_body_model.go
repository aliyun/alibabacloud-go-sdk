// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainRealTimeDetailDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DescribeDcdnDomainRealTimeDetailDataResponseBody
	GetData() *string
	SetRequestId(v string) *DescribeDcdnDomainRealTimeDetailDataResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainRealTimeDetailDataResponseBody struct {
	// The information returned.
	//
	// > The value of this parameter is a JSON string. The following table describes the fields in Data.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//       "time_stp": "2018-06-05T20:00:00Z",
	//
	//       "domain_name": "example.com",
	//
	//       "location": "Guangdong",
	//
	//       "isp": "telecom",
	//
	//       "qps": 10.0
	//
	//     },
	//
	//     {
	//
	//       "time_stp": "2018-06-05T20:00:00Z",
	//
	//       "domain_name": "example.com",
	//
	//       "location": "Jiangsu",
	//
	//       "isp": "unicom",
	//
	//       "qps": 11.1
	//
	//     }
	//
	//   ]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A603F324-7A05-4FB3-ADF3-2563233D26CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainRealTimeDetailDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainRealTimeDetailDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponseBody) SetData(v string) *DescribeDcdnDomainRealTimeDetailDataResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponseBody) SetRequestId(v string) *DescribeDcdnDomainRealTimeDetailDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainRealTimeDetailDataResponseBody) Validate() error {
	return dara.Validate(s)
}
