// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnIpInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDcdnIp(v string) *DescribeDcdnIpInfoResponseBody
	GetDcdnIp() *string
	SetISP(v string) *DescribeDcdnIpInfoResponseBody
	GetISP() *string
	SetIspEname(v string) *DescribeDcdnIpInfoResponseBody
	GetIspEname() *string
	SetRegion(v string) *DescribeDcdnIpInfoResponseBody
	GetRegion() *string
	SetRegionEname(v string) *DescribeDcdnIpInfoResponseBody
	GetRegionEname() *string
	SetRequestId(v string) *DescribeDcdnIpInfoResponseBody
	GetRequestId() *string
}

type DescribeDcdnIpInfoResponseBody struct {
	// Indicates whether the specified IP address is assigned to an Alibaba Cloud DCDN POP.
	//
	// 	- True
	//
	// 	- False
	//
	// example:
	//
	// True
	DcdnIp *string `json:"DcdnIp,omitempty" xml:"DcdnIp,omitempty"`
	// The ISP to which the specified IP address belongs.
	//
	// example:
	//
	// China Telecom
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The name of the Internet service provider (ISP).
	//
	// example:
	//
	// telecom
	IspEname *string `json:"IspEname,omitempty" xml:"IspEname,omitempty"`
	// The Chinese name of the region.
	//
	// example:
	//
	// >  The maximum number of times that users can call this operation per second is 50.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The English name of the region.
	//
	// example:
	//
	// China-Guizhou-guiyang
	RegionEname *string `json:"RegionEname,omitempty" xml:"RegionEname,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1B1D0EE7-9559-489D-BC4E-279495EB8FB8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnIpInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnIpInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnIpInfoResponseBody) GetDcdnIp() *string {
	return s.DcdnIp
}

func (s *DescribeDcdnIpInfoResponseBody) GetISP() *string {
	return s.ISP
}

func (s *DescribeDcdnIpInfoResponseBody) GetIspEname() *string {
	return s.IspEname
}

func (s *DescribeDcdnIpInfoResponseBody) GetRegion() *string {
	return s.Region
}

func (s *DescribeDcdnIpInfoResponseBody) GetRegionEname() *string {
	return s.RegionEname
}

func (s *DescribeDcdnIpInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnIpInfoResponseBody) SetDcdnIp(v string) *DescribeDcdnIpInfoResponseBody {
	s.DcdnIp = &v
	return s
}

func (s *DescribeDcdnIpInfoResponseBody) SetISP(v string) *DescribeDcdnIpInfoResponseBody {
	s.ISP = &v
	return s
}

func (s *DescribeDcdnIpInfoResponseBody) SetIspEname(v string) *DescribeDcdnIpInfoResponseBody {
	s.IspEname = &v
	return s
}

func (s *DescribeDcdnIpInfoResponseBody) SetRegion(v string) *DescribeDcdnIpInfoResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeDcdnIpInfoResponseBody) SetRegionEname(v string) *DescribeDcdnIpInfoResponseBody {
	s.RegionEname = &v
	return s
}

func (s *DescribeDcdnIpInfoResponseBody) SetRequestId(v string) *DescribeDcdnIpInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnIpInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
