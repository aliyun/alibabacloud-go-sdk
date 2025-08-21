// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebReportTopIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeWebReportTopIpResponseBodyDataList) *DescribeWebReportTopIpResponseBody
	GetDataList() []*DescribeWebReportTopIpResponseBodyDataList
	SetRequestId(v string) *DescribeWebReportTopIpResponseBody
	GetRequestId() *string
}

type DescribeWebReportTopIpResponseBody struct {
	// The information about the IP addresses.
	DataList []*DescribeWebReportTopIpResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D21BE0C4-8E83-5E32-86C6-AA6BE9B1B5BD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebReportTopIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebReportTopIpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebReportTopIpResponseBody) GetDataList() []*DescribeWebReportTopIpResponseBodyDataList {
	return s.DataList
}

func (s *DescribeWebReportTopIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebReportTopIpResponseBody) SetDataList(v []*DescribeWebReportTopIpResponseBodyDataList) *DescribeWebReportTopIpResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeWebReportTopIpResponseBody) SetRequestId(v string) *DescribeWebReportTopIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebReportTopIpResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWebReportTopIpResponseBodyDataList struct {
	// The ID of the location.
	//
	// example:
	//
	// 90998690
	AreaId *string `json:"AreaId,omitempty" xml:"AreaId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 5
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The Internet service provider (ISP) for the attack. Valid values:
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
	// 100017
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 117.186.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DescribeWebReportTopIpResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebReportTopIpResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeWebReportTopIpResponseBodyDataList) GetAreaId() *string {
	return s.AreaId
}

func (s *DescribeWebReportTopIpResponseBodyDataList) GetCount() *int64 {
	return s.Count
}

func (s *DescribeWebReportTopIpResponseBodyDataList) GetIsp() *string {
	return s.Isp
}

func (s *DescribeWebReportTopIpResponseBodyDataList) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeWebReportTopIpResponseBodyDataList) SetAreaId(v string) *DescribeWebReportTopIpResponseBodyDataList {
	s.AreaId = &v
	return s
}

func (s *DescribeWebReportTopIpResponseBodyDataList) SetCount(v int64) *DescribeWebReportTopIpResponseBodyDataList {
	s.Count = &v
	return s
}

func (s *DescribeWebReportTopIpResponseBodyDataList) SetIsp(v string) *DescribeWebReportTopIpResponseBodyDataList {
	s.Isp = &v
	return s
}

func (s *DescribeWebReportTopIpResponseBodyDataList) SetSourceIp(v string) *DescribeWebReportTopIpResponseBodyDataList {
	s.SourceIp = &v
	return s
}

func (s *DescribeWebReportTopIpResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
