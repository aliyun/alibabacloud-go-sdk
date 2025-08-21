// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventIspResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsps(v []*DescribeDDosEventIspResponseBodyIsps) *DescribeDDosEventIspResponseBody
	GetIsps() []*DescribeDDosEventIspResponseBodyIsps
	SetRequestId(v string) *DescribeDDosEventIspResponseBody
	GetRequestId() *string
}

type DescribeDDosEventIspResponseBody struct {
	// The ISPs for the volumetric attack.
	Isps []*DescribeDDosEventIspResponseBodyIsps `json:"Isps,omitempty" xml:"Isps,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C4A3BCD1-4A32-4342-941A-4745AE69508C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDDosEventIspResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventIspResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventIspResponseBody) GetIsps() []*DescribeDDosEventIspResponseBodyIsps {
	return s.Isps
}

func (s *DescribeDDosEventIspResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDosEventIspResponseBody) SetIsps(v []*DescribeDDosEventIspResponseBodyIsps) *DescribeDDosEventIspResponseBody {
	s.Isps = v
	return s
}

func (s *DescribeDDosEventIspResponseBody) SetRequestId(v string) *DescribeDDosEventIspResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDosEventIspResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDDosEventIspResponseBodyIsps struct {
	// The number of request packets that were sent from the ISP.
	//
	// example:
	//
	// 230
	InPkts *int64 `json:"InPkts,omitempty" xml:"InPkts,omitempty"`
	// The code of the ISP. Valid values:
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
	// 1000323
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
}

func (s DescribeDDosEventIspResponseBodyIsps) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventIspResponseBodyIsps) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventIspResponseBodyIsps) GetInPkts() *int64 {
	return s.InPkts
}

func (s *DescribeDDosEventIspResponseBodyIsps) GetIsp() *string {
	return s.Isp
}

func (s *DescribeDDosEventIspResponseBodyIsps) SetInPkts(v int64) *DescribeDDosEventIspResponseBodyIsps {
	s.InPkts = &v
	return s
}

func (s *DescribeDDosEventIspResponseBodyIsps) SetIsp(v string) *DescribeDDosEventIspResponseBodyIsps {
	s.Isp = &v
	return s
}

func (s *DescribeDDosEventIspResponseBodyIsps) Validate() error {
	return dara.Validate(s)
}
