// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoResourceEnable(v bool) *DescribeAssetStatisticResponseBody
	GetAutoResourceEnable() *bool
	SetGeneralInstanceSpecStatistic(v *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) *DescribeAssetStatisticResponseBody
	GetGeneralInstanceSpecStatistic() *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic
	SetRequestId(v string) *DescribeAssetStatisticResponseBody
	GetRequestId() *string
	SetResourceSpecStatistic(v *DescribeAssetStatisticResponseBodyResourceSpecStatistic) *DescribeAssetStatisticResponseBody
	GetResourceSpecStatistic() *DescribeAssetStatisticResponseBodyResourceSpecStatistic
}

type DescribeAssetStatisticResponseBody struct {
	// Whether automatic traffic diversion is enabled. Valid values:- **true**: enabled- **false**: disabled
	//
	// example:
	//
	// true
	AutoResourceEnable *bool `json:"AutoResourceEnable,omitempty" xml:"AutoResourceEnable,omitempty"`
	// Specifications for general instances in version 2.0.
	GeneralInstanceSpecStatistic *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic `json:"GeneralInstanceSpecStatistic,omitempty" xml:"GeneralInstanceSpecStatistic,omitempty" type:"Struct"`
	// ID of the request.
	//
	// example:
	//
	// 850A84******25g4d2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Resource specification usage statistics.
	ResourceSpecStatistic *DescribeAssetStatisticResponseBodyResourceSpecStatistic `json:"ResourceSpecStatistic,omitempty" xml:"ResourceSpecStatistic,omitempty" type:"Struct"`
}

func (s DescribeAssetStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetStatisticResponseBody) GetAutoResourceEnable() *bool {
	return s.AutoResourceEnable
}

func (s *DescribeAssetStatisticResponseBody) GetGeneralInstanceSpecStatistic() *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic {
	return s.GeneralInstanceSpecStatistic
}

func (s *DescribeAssetStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAssetStatisticResponseBody) GetResourceSpecStatistic() *DescribeAssetStatisticResponseBodyResourceSpecStatistic {
	return s.ResourceSpecStatistic
}

func (s *DescribeAssetStatisticResponseBody) SetAutoResourceEnable(v bool) *DescribeAssetStatisticResponseBody {
	s.AutoResourceEnable = &v
	return s
}

func (s *DescribeAssetStatisticResponseBody) SetGeneralInstanceSpecStatistic(v *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) *DescribeAssetStatisticResponseBody {
	s.GeneralInstanceSpecStatistic = v
	return s
}

func (s *DescribeAssetStatisticResponseBody) SetRequestId(v string) *DescribeAssetStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetStatisticResponseBody) SetResourceSpecStatistic(v *DescribeAssetStatisticResponseBodyResourceSpecStatistic) *DescribeAssetStatisticResponseBody {
	s.ResourceSpecStatistic = v
	return s
}

func (s *DescribeAssetStatisticResponseBody) Validate() error {
	if s.GeneralInstanceSpecStatistic != nil {
		if err := s.GeneralInstanceSpecStatistic.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceSpecStatistic != nil {
		if err := s.ResourceSpecStatistic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic struct {
	// Internet-facing firewall instance usage by region.
	CfwGeneralInstanceRegionStatistic []*DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic `json:"CfwGeneralInstanceRegionStatistic,omitempty" xml:"CfwGeneralInstanceRegionStatistic,omitempty" type:"Repeated"`
	// Total Internet-facing firewall instances by region.
	CfwTotalGeneralInstanceRegionStatistic []*DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic `json:"CfwTotalGeneralInstanceRegionStatistic,omitempty" xml:"CfwTotalGeneralInstanceRegionStatistic,omitempty" type:"Repeated"`
	// Total number of Internet-facing firewall instances.
	//
	// example:
	//
	// 1
	TotalCfwGeneralInstanceCnt *int32 `json:"TotalCfwGeneralInstanceCnt,omitempty" xml:"TotalCfwGeneralInstanceCnt,omitempty"`
	// Number of Internet-facing firewall instances with protection enabled.
	//
	// example:
	//
	// 1
	TotalCfwGeneralInstanceUsedCnt *int32 `json:"TotalCfwGeneralInstanceUsedCnt,omitempty" xml:"TotalCfwGeneralInstanceUsedCnt,omitempty"`
	// Total number of general instances used.
	//
	// example:
	//
	// 1
	TotalGeneralInstanceUsedCnt *int32 `json:"TotalGeneralInstanceUsedCnt,omitempty" xml:"TotalGeneralInstanceUsedCnt,omitempty"`
	// Total number of NAT firewall instances.
	//
	// example:
	//
	// 1
	TotalNatGeneralInstanceCnt *int32 `json:"TotalNatGeneralInstanceCnt,omitempty" xml:"TotalNatGeneralInstanceCnt,omitempty"`
	// Number of NAT firewall instances with protection enabled.
	//
	// example:
	//
	// 1
	TotalNatGeneralInstanceUsedCnt *int32 `json:"TotalNatGeneralInstanceUsedCnt,omitempty" xml:"TotalNatGeneralInstanceUsedCnt,omitempty"`
	// Number of VPC firewall instances with protection enabled.
	//
	// example:
	//
	// 1
	TotalVfwGeneralInstanceUsedCnt *int32 `json:"TotalVfwGeneralInstanceUsedCnt,omitempty" xml:"TotalVfwGeneralInstanceUsedCnt,omitempty"`
}

func (s DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) GoString() string {
	return s.String()
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) GetCfwGeneralInstanceRegionStatistic() []*DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic {
	return s.CfwGeneralInstanceRegionStatistic
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) GetCfwTotalGeneralInstanceRegionStatistic() []*DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic {
	return s.CfwTotalGeneralInstanceRegionStatistic
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) GetTotalCfwGeneralInstanceCnt() *int32 {
	return s.TotalCfwGeneralInstanceCnt
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) GetTotalCfwGeneralInstanceUsedCnt() *int32 {
	return s.TotalCfwGeneralInstanceUsedCnt
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) GetTotalGeneralInstanceUsedCnt() *int32 {
	return s.TotalGeneralInstanceUsedCnt
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) GetTotalNatGeneralInstanceCnt() *int32 {
	return s.TotalNatGeneralInstanceCnt
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) GetTotalNatGeneralInstanceUsedCnt() *int32 {
	return s.TotalNatGeneralInstanceUsedCnt
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) GetTotalVfwGeneralInstanceUsedCnt() *int32 {
	return s.TotalVfwGeneralInstanceUsedCnt
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) SetCfwGeneralInstanceRegionStatistic(v []*DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic) *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic {
	s.CfwGeneralInstanceRegionStatistic = v
	return s
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) SetCfwTotalGeneralInstanceRegionStatistic(v []*DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic) *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic {
	s.CfwTotalGeneralInstanceRegionStatistic = v
	return s
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) SetTotalCfwGeneralInstanceCnt(v int32) *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic {
	s.TotalCfwGeneralInstanceCnt = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) SetTotalCfwGeneralInstanceUsedCnt(v int32) *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic {
	s.TotalCfwGeneralInstanceUsedCnt = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) SetTotalGeneralInstanceUsedCnt(v int32) *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic {
	s.TotalGeneralInstanceUsedCnt = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) SetTotalNatGeneralInstanceCnt(v int32) *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic {
	s.TotalNatGeneralInstanceCnt = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) SetTotalNatGeneralInstanceUsedCnt(v int32) *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic {
	s.TotalNatGeneralInstanceUsedCnt = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) SetTotalVfwGeneralInstanceUsedCnt(v int32) *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic {
	s.TotalVfwGeneralInstanceUsedCnt = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) Validate() error {
	if s.CfwGeneralInstanceRegionStatistic != nil {
		for _, item := range s.CfwGeneralInstanceRegionStatistic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CfwTotalGeneralInstanceRegionStatistic != nil {
		for _, item := range s.CfwTotalGeneralInstanceRegionStatistic {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic struct {
	// List of member accounts in the region.
	MemberList []*string `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// Region information
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic) GoString() string {
	return s.String()
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic) GetMemberList() []*string {
	return s.MemberList
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic) SetMemberList(v []*string) *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic {
	s.MemberList = v
	return s
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic) SetRegionNo(v string) *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic {
	s.RegionNo = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic) Validate() error {
	return dara.Validate(s)
}

type DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic struct {
	// List of member accounts in the region.
	MemberList []*string `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
	// Region information
	//
	// example:
	//
	// cn-shanghai
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic) GoString() string {
	return s.String()
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic) GetMemberList() []*string {
	return s.MemberList
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic) SetMemberList(v []*string) *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic {
	s.MemberList = v
	return s
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic) SetRegionNo(v string) *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic {
	s.RegionNo = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic) Validate() error {
	return dara.Validate(s)
}

type DescribeAssetStatisticResponseBodyResourceSpecStatistic struct {
	// Maximum number of public IP addresses that can be protected.
	//
	// example:
	//
	// 20
	IpNumSpec *int32 `json:"IpNumSpec,omitempty" xml:"IpNumSpec,omitempty"`
	// Number of public IP addresses with protection enabled.
	//
	// example:
	//
	// 10
	IpNumUsed *int32 `json:"IpNumUsed,omitempty" xml:"IpNumUsed,omitempty"`
	// The number of IP specifications for sensitive data.
	//
	// example:
	//
	// 0
	SensitiveDataIpNumSpec *int64 `json:"SensitiveDataIpNumSpec,omitempty" xml:"SensitiveDataIpNumSpec,omitempty"`
	// Number of public IP addresses currently scanned for sensitive data.
	//
	// example:
	//
	// 0
	SensitiveDataIpNumUsed *int64 `json:"SensitiveDataIpNumUsed,omitempty" xml:"SensitiveDataIpNumUsed,omitempty"`
}

func (s DescribeAssetStatisticResponseBodyResourceSpecStatistic) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetStatisticResponseBodyResourceSpecStatistic) GoString() string {
	return s.String()
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) GetIpNumSpec() *int32 {
	return s.IpNumSpec
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) GetIpNumUsed() *int32 {
	return s.IpNumUsed
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) GetSensitiveDataIpNumSpec() *int64 {
	return s.SensitiveDataIpNumSpec
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) GetSensitiveDataIpNumUsed() *int64 {
	return s.SensitiveDataIpNumUsed
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) SetIpNumSpec(v int32) *DescribeAssetStatisticResponseBodyResourceSpecStatistic {
	s.IpNumSpec = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) SetIpNumUsed(v int32) *DescribeAssetStatisticResponseBodyResourceSpecStatistic {
	s.IpNumUsed = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) SetSensitiveDataIpNumSpec(v int64) *DescribeAssetStatisticResponseBodyResourceSpecStatistic {
	s.SensitiveDataIpNumSpec = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) SetSensitiveDataIpNumUsed(v int64) *DescribeAssetStatisticResponseBodyResourceSpecStatistic {
	s.SensitiveDataIpNumUsed = &v
	return s
}

func (s *DescribeAssetStatisticResponseBodyResourceSpecStatistic) Validate() error {
	return dara.Validate(s)
}
