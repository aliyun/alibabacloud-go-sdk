// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallAssetListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeVpcFirewallAssetListResponseBodyDataList) *DescribeVpcFirewallAssetListResponseBody
	GetDataList() []*DescribeVpcFirewallAssetListResponseBodyDataList
	SetRequestId(v string) *DescribeVpcFirewallAssetListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcFirewallAssetListResponseBody
	GetTotalCount() *int32
}

type DescribeVpcFirewallAssetListResponseBody struct {
	DataList []*DescribeVpcFirewallAssetListResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 924A6CCC-4EAD-5554-8AD0-45F5ED56****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcFirewallAssetListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAssetListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAssetListResponseBody) GetDataList() []*DescribeVpcFirewallAssetListResponseBodyDataList {
	return s.DataList
}

func (s *DescribeVpcFirewallAssetListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallAssetListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcFirewallAssetListResponseBody) SetDataList(v []*DescribeVpcFirewallAssetListResponseBodyDataList) *DescribeVpcFirewallAssetListResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBody) SetRequestId(v string) *DescribeVpcFirewallAssetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallAssetListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcFirewallAssetListResponseBodyDataList struct {
	// example:
	//
	// 192.0.XX.XX
	AssetIP *string `json:"AssetIP,omitempty" xml:"AssetIP,omitempty"`
	// example:
	//
	// i-hp3ez3rs9bxwt034****
	AssetInstanceId *string `json:"AssetInstanceId,omitempty" xml:"AssetInstanceId,omitempty"`
	// example:
	//
	// ecs-test
	AssetInstanceName *string `json:"AssetInstanceName,omitempty" xml:"AssetInstanceName,omitempty"`
	// example:
	//
	// 0.0
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// example:
	//
	// 10
	IpsHitCnt *int64 `json:"IpsHitCnt,omitempty" xml:"IpsHitCnt,omitempty"`
	// example:
	//
	// 0.0
	OutBytes *int64    `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	PortList []*string `json:"PortList,omitempty" xml:"PortList,omitempty" type:"Repeated"`
	// example:
	//
	// cn-beijing
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// 3
	RiskLevel *int32 `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// example:
	//
	// test
	RiskReason *string `json:"RiskReason,omitempty" xml:"RiskReason,omitempty"`
	// example:
	//
	// 27
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// example:
	//
	// 0
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
}

func (s DescribeVpcFirewallAssetListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAssetListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) GetAssetIP() *string {
	return s.AssetIP
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) GetAssetInstanceId() *string {
	return s.AssetInstanceId
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) GetAssetInstanceName() *string {
	return s.AssetInstanceName
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) GetIpsHitCnt() *int64 {
	return s.IpsHitCnt
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) GetPortList() []*string {
	return s.PortList
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) GetRiskLevel() *int32 {
	return s.RiskLevel
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) GetRiskReason() *string {
	return s.RiskReason
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) SetAssetIP(v string) *DescribeVpcFirewallAssetListResponseBodyDataList {
	s.AssetIP = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) SetAssetInstanceId(v string) *DescribeVpcFirewallAssetListResponseBodyDataList {
	s.AssetInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) SetAssetInstanceName(v string) *DescribeVpcFirewallAssetListResponseBodyDataList {
	s.AssetInstanceName = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) SetInBytes(v int64) *DescribeVpcFirewallAssetListResponseBodyDataList {
	s.InBytes = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) SetIpsHitCnt(v int64) *DescribeVpcFirewallAssetListResponseBodyDataList {
	s.IpsHitCnt = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) SetOutBytes(v int64) *DescribeVpcFirewallAssetListResponseBodyDataList {
	s.OutBytes = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) SetPortList(v []*string) *DescribeVpcFirewallAssetListResponseBodyDataList {
	s.PortList = v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) SetRegionNo(v string) *DescribeVpcFirewallAssetListResponseBodyDataList {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) SetRiskLevel(v int32) *DescribeVpcFirewallAssetListResponseBodyDataList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) SetRiskReason(v string) *DescribeVpcFirewallAssetListResponseBodyDataList {
	s.RiskReason = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) SetSessionCount(v int64) *DescribeVpcFirewallAssetListResponseBodyDataList {
	s.SessionCount = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) SetTotalBytes(v int64) *DescribeVpcFirewallAssetListResponseBodyDataList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeVpcFirewallAssetListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
