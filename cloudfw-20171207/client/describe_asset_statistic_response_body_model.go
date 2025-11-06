// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGeneralInstanceSpecStatistic(v *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic) *DescribeAssetStatisticResponseBody
	GetGeneralInstanceSpecStatistic() *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic
	SetRequestId(v string) *DescribeAssetStatisticResponseBody
	GetRequestId() *string
	SetResourceSpecStatistic(v *DescribeAssetStatisticResponseBodyResourceSpecStatistic) *DescribeAssetStatisticResponseBody
	GetResourceSpecStatistic() *DescribeAssetStatisticResponseBodyResourceSpecStatistic
}

type DescribeAssetStatisticResponseBody struct {
	GeneralInstanceSpecStatistic *DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatistic `json:"GeneralInstanceSpecStatistic,omitempty" xml:"GeneralInstanceSpecStatistic,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 850A84******25g4d2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The statistics on specifications.
	ResourceSpecStatistic *DescribeAssetStatisticResponseBodyResourceSpecStatistic `json:"ResourceSpecStatistic,omitempty" xml:"ResourceSpecStatistic,omitempty" type:"Struct"`
}

func (s DescribeAssetStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetStatisticResponseBody) GoString() string {
	return s.String()
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
	CfwGeneralInstanceRegionStatistic      []*DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwGeneralInstanceRegionStatistic      `json:"CfwGeneralInstanceRegionStatistic,omitempty" xml:"CfwGeneralInstanceRegionStatistic,omitempty" type:"Repeated"`
	CfwTotalGeneralInstanceRegionStatistic []*DescribeAssetStatisticResponseBodyGeneralInstanceSpecStatisticCfwTotalGeneralInstanceRegionStatistic `json:"CfwTotalGeneralInstanceRegionStatistic,omitempty" xml:"CfwTotalGeneralInstanceRegionStatistic,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalCfwGeneralInstanceCnt *int32 `json:"TotalCfwGeneralInstanceCnt,omitempty" xml:"TotalCfwGeneralInstanceCnt,omitempty"`
	// example:
	//
	// 1
	TotalCfwGeneralInstanceUsedCnt *int32 `json:"TotalCfwGeneralInstanceUsedCnt,omitempty" xml:"TotalCfwGeneralInstanceUsedCnt,omitempty"`
	// example:
	//
	// 1
	TotalGeneralInstanceUsedCnt *int32 `json:"TotalGeneralInstanceUsedCnt,omitempty" xml:"TotalGeneralInstanceUsedCnt,omitempty"`
	// example:
	//
	// 1
	TotalNatGeneralInstanceCnt *int32 `json:"TotalNatGeneralInstanceCnt,omitempty" xml:"TotalNatGeneralInstanceCnt,omitempty"`
	// example:
	//
	// 1
	TotalNatGeneralInstanceUsedCnt *int32 `json:"TotalNatGeneralInstanceUsedCnt,omitempty" xml:"TotalNatGeneralInstanceUsedCnt,omitempty"`
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
	MemberList []*string `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
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
	MemberList []*string `json:"MemberList,omitempty" xml:"MemberList,omitempty" type:"Repeated"`
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
	// The number of public IP addresses that can be protected.
	//
	// example:
	//
	// 20
	IpNumSpec *int32 `json:"IpNumSpec,omitempty" xml:"IpNumSpec,omitempty"`
	// The number of public IP addresses that are protected.
	//
	// example:
	//
	// 10
	IpNumUsed *int32 `json:"IpNumUsed,omitempty" xml:"IpNumUsed,omitempty"`
	// The number of public IP addresses that can enable data leakage detection.
	//
	// example:
	//
	// 0
	SensitiveDataIpNumSpec *int64 `json:"SensitiveDataIpNumSpec,omitempty" xml:"SensitiveDataIpNumSpec,omitempty"`
	// The number of public IP addresses that enabled data leakage detection.
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
