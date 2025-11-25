// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventSrcIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIps(v []*DescribeDDosEventSrcIpResponseBodyIps) *DescribeDDosEventSrcIpResponseBody
	GetIps() []*DescribeDDosEventSrcIpResponseBodyIps
	SetRequestId(v string) *DescribeDDosEventSrcIpResponseBody
	GetRequestId() *string
}

type DescribeDDosEventSrcIpResponseBody struct {
	// An array that consists of information about the source IP address of the volumetric attack.
	Ips []*DescribeDDosEventSrcIpResponseBodyIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 38A0224E-FDBC-4733-A362-B391827FC1E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDDosEventSrcIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventSrcIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventSrcIpResponseBody) GetIps() []*DescribeDDosEventSrcIpResponseBodyIps {
	return s.Ips
}

func (s *DescribeDDosEventSrcIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDosEventSrcIpResponseBody) SetIps(v []*DescribeDDosEventSrcIpResponseBodyIps) *DescribeDDosEventSrcIpResponseBody {
	s.Ips = v
	return s
}

func (s *DescribeDDosEventSrcIpResponseBody) SetRequestId(v string) *DescribeDDosEventSrcIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDosEventSrcIpResponseBody) Validate() error {
	if s.Ips != nil {
		for _, item := range s.Ips {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDDosEventSrcIpResponseBodyIps struct {
	// The code or ID of the source region. For more information, see [Codes of administrative regions in China and codes of countries and areas](https://help.aliyun.com/document_detail/167926.html). For example, **110000*	- indicates Beijing, China, and **us*	- indicates the United States.
	//
	// example:
	//
	// 110000
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The Internet service provider (ISP) for the volumetric attack. Valid values:
	//
	// 	- **100017**: China Telecom
	//
	// 	- **100026**: China Unicom
	//
	// 	- **100025**: China Mobile
	//
	// 	- **100027**: China Education and Research Network
	//
	// 	- **100020**: China Mobile Tietong
	//
	// 	- **1000143**: Dr.Peng Telecom & Media Group
	//
	// 	- **100080**: Beijing Gehua CATV Network
	//
	// 	- **1000139**: National Radio and Television Administration
	//
	// 	- **100023**: Oriental Cable Network
	//
	// 	- **100063**: Founder Broadband
	//
	// 	- **1000337**: China Internet Exchange
	//
	// 	- **100021**: 21Vianet Group
	//
	// 	- **1000333**: Wasu Media Holding
	//
	// 	- **100093**: Wangsu Science & Technology
	//
	// 	- **1000401**: Tencent
	//
	// 	- **100099**: Baidu
	//
	// 	- **1000323**: Alibaba Cloud
	//
	// 	- **100098**: Alibaba
	//
	// example:
	//
	// 100026
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The source IP address of the volumetric attack.
	//
	// example:
	//
	// 218.***.***.24
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
}

func (s DescribeDDosEventSrcIpResponseBodyIps) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventSrcIpResponseBodyIps) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventSrcIpResponseBodyIps) GetAreaId() *string {
	return s.AreaId
}

func (s *DescribeDDosEventSrcIpResponseBodyIps) GetIsp() *string {
	return s.Isp
}

func (s *DescribeDDosEventSrcIpResponseBodyIps) GetSrcIp() *string {
	return s.SrcIp
}

func (s *DescribeDDosEventSrcIpResponseBodyIps) SetAreaId(v string) *DescribeDDosEventSrcIpResponseBodyIps {
	s.AreaId = &v
	return s
}

func (s *DescribeDDosEventSrcIpResponseBodyIps) SetIsp(v string) *DescribeDDosEventSrcIpResponseBodyIps {
	s.Isp = &v
	return s
}

func (s *DescribeDDosEventSrcIpResponseBodyIps) SetSrcIp(v string) *DescribeDDosEventSrcIpResponseBodyIps {
	s.SrcIp = &v
	return s
}

func (s *DescribeDDosEventSrcIpResponseBodyIps) Validate() error {
	return dara.Validate(s)
}
