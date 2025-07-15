// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveIpInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetISP(v string) *DescribeLiveIpInfoResponseBody
	GetISP() *string
	SetIspEname(v string) *DescribeLiveIpInfoResponseBody
	GetIspEname() *string
	SetRegion(v string) *DescribeLiveIpInfoResponseBody
	GetRegion() *string
	SetRegionEname(v string) *DescribeLiveIpInfoResponseBody
	GetRegionEname() *string
	SetRequestId(v string) *DescribeLiveIpInfoResponseBody
	GetRequestId() *string
}

type DescribeLiveIpInfoResponseBody struct {
	// The Chinese name of the ISP.
	//
	// example:
	//
	// ChinaTelecom
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The English name of the Internet service provider (ISP).
	//
	// example:
	//
	// telecom
	IspEname *string `json:"IspEname,omitempty" xml:"IspEname,omitempty"`
	// The Chinese name of the region.
	//
	// example:
	//
	// China-Guizhou-guiyang
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

func (s DescribeLiveIpInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveIpInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveIpInfoResponseBody) GetISP() *string {
	return s.ISP
}

func (s *DescribeLiveIpInfoResponseBody) GetIspEname() *string {
	return s.IspEname
}

func (s *DescribeLiveIpInfoResponseBody) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveIpInfoResponseBody) GetRegionEname() *string {
	return s.RegionEname
}

func (s *DescribeLiveIpInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveIpInfoResponseBody) SetISP(v string) *DescribeLiveIpInfoResponseBody {
	s.ISP = &v
	return s
}

func (s *DescribeLiveIpInfoResponseBody) SetIspEname(v string) *DescribeLiveIpInfoResponseBody {
	s.IspEname = &v
	return s
}

func (s *DescribeLiveIpInfoResponseBody) SetRegion(v string) *DescribeLiveIpInfoResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeLiveIpInfoResponseBody) SetRegionEname(v string) *DescribeLiveIpInfoResponseBody {
	s.RegionEname = &v
	return s
}

func (s *DescribeLiveIpInfoResponseBody) SetRequestId(v string) *DescribeLiveIpInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveIpInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
