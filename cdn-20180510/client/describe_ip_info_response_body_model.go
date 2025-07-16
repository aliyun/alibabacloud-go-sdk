// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCdnIp(v string) *DescribeIpInfoResponseBody
	GetCdnIp() *string
	SetISP(v string) *DescribeIpInfoResponseBody
	GetISP() *string
	SetIspEname(v string) *DescribeIpInfoResponseBody
	GetIspEname() *string
	SetRegion(v string) *DescribeIpInfoResponseBody
	GetRegion() *string
	SetRegionEname(v string) *DescribeIpInfoResponseBody
	GetRegionEname() *string
	SetRequestId(v string) *DescribeIpInfoResponseBody
	GetRequestId() *string
}

type DescribeIpInfoResponseBody struct {
	// Indicates whether the IP address belongs to an Alibaba Cloud CDN POP.
	//
	// 	- **True**:Yes.
	//
	// 	- **False**:No.
	//
	// example:
	//
	// True
	CdnIp *string `json:"CdnIp,omitempty" xml:"CdnIp,omitempty"`
	// The name of the ISP in Chinese.
	//
	// example:
	//
	// 电信
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The name of the ISP.
	//
	// example:
	//
	// telecom
	IspEname *string `json:"IspEname,omitempty" xml:"IspEname,omitempty"`
	// The name of the region in Chinese.
	//
	// example:
	//
	// 中国-贵州省-贵阳市
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// China-Guizhou-guiyang
	RegionEname *string `json:"RegionEname,omitempty" xml:"RegionEname,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 123847FA-9A00-4426-83B8-B4B45D475930
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeIpInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpInfoResponseBody) GetCdnIp() *string {
	return s.CdnIp
}

func (s *DescribeIpInfoResponseBody) GetISP() *string {
	return s.ISP
}

func (s *DescribeIpInfoResponseBody) GetIspEname() *string {
	return s.IspEname
}

func (s *DescribeIpInfoResponseBody) GetRegion() *string {
	return s.Region
}

func (s *DescribeIpInfoResponseBody) GetRegionEname() *string {
	return s.RegionEname
}

func (s *DescribeIpInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpInfoResponseBody) SetCdnIp(v string) *DescribeIpInfoResponseBody {
	s.CdnIp = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetISP(v string) *DescribeIpInfoResponseBody {
	s.ISP = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetIspEname(v string) *DescribeIpInfoResponseBody {
	s.IspEname = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetRegion(v string) *DescribeIpInfoResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetRegionEname(v string) *DescribeIpInfoResponseBody {
	s.RegionEname = &v
	return s
}

func (s *DescribeIpInfoResponseBody) SetRequestId(v string) *DescribeIpInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
