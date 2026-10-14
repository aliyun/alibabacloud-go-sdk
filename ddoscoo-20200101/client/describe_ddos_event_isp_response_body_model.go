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
	// The ISP information of attack sources.
	Isps []*DescribeDDosEventIspResponseBodyIsps `json:"Isps,omitempty" xml:"Isps,omitempty" type:"Repeated"`
	// The request ID.
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
	if s.Isps != nil {
		for _, item := range s.Isps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDDosEventIspResponseBodyIsps struct {
	// The number of request packets from the ISP.
	//
	// example:
	//
	// 230
	InPkts *int64 `json:"InPkts,omitempty" xml:"InPkts,omitempty"`
	// The code of the attack source ISP. Valid values:
	//
	// - **100017**: China Telecom
	//
	// - **100026**: China Unicom
	//
	// - **100025**: China Mobile
	//
	// - **100027**: CERNET
	//
	// - **100020**: China Tietong
	//
	// - **1000143**: Dr. Peng
	//
	// - **100080**: Gehua
	//
	// - **1000139**: China Broadcasting Network
	//
	// - **100023**: Youxiantong
	//
	// - **100063**: Founder Broadband
	//
	// - **1000337**: Haokuan Network
	//
	// - **100021**: 21Vianet
	//
	// - **1000333**: Wasu Media
	//
	// - **100093**: ChinaNetCenter
	//
	// - **1000401**: Tencent
	//
	// - **100099**: Baidu
	//
	// - **1000323**: Alibaba Cloud
	//
	// - **100098**: Alibaba
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
