// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallTrafficAssetListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeVpcFirewallTrafficAssetListResponseBodyDataList) *DescribeVpcFirewallTrafficAssetListResponseBody
	GetDataList() []*DescribeVpcFirewallTrafficAssetListResponseBodyDataList
	SetRequestId(v string) *DescribeVpcFirewallTrafficAssetListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcFirewallTrafficAssetListResponseBody
	GetTotalCount() *int32
}

type DescribeVpcFirewallTrafficAssetListResponseBody struct {
	// The data list.
	DataList []*DescribeVpcFirewallTrafficAssetListResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// C5BE1AA4-934A-5085-89CC-9AD1CAC3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 132
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcFirewallTrafficAssetListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallTrafficAssetListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBody) GetDataList() []*DescribeVpcFirewallTrafficAssetListResponseBodyDataList {
	return s.DataList
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBody) SetDataList(v []*DescribeVpcFirewallTrafficAssetListResponseBodyDataList) *DescribeVpcFirewallTrafficAssetListResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBody) SetRequestId(v string) *DescribeVpcFirewallTrafficAssetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallTrafficAssetListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBody) Validate() error {
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

type DescribeVpcFirewallTrafficAssetListResponseBodyDataList struct {
	// The IP address of the asset.
	//
	// example:
	//
	// a.com
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The number of IPS hits.
	//
	// example:
	//
	// 10
	IpsHitCnt *int64 `json:"IpsHitCnt,omitempty" xml:"IpsHitCnt,omitempty"`
	// The number of sessions.
	//
	// example:
	//
	// 12
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// The total traffic. Unit: bytes.
	//
	// example:
	//
	// 253023143
	TotalBytes *int64 `json:"TotalBytes,omitempty" xml:"TotalBytes,omitempty"`
	// The ID of the VPC where the asset resides.
	//
	// example:
	//
	// a.com
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the VPC where the asset resides.
	//
	// example:
	//
	// vpc-bp1mos0vhefmx5ah6****
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcFirewallTrafficAssetListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallTrafficAssetListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBodyDataList) GetIP() *string {
	return s.IP
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBodyDataList) GetIpsHitCnt() *int64 {
	return s.IpsHitCnt
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBodyDataList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBodyDataList) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBodyDataList) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBodyDataList) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBodyDataList) SetIP(v string) *DescribeVpcFirewallTrafficAssetListResponseBodyDataList {
	s.IP = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBodyDataList) SetIpsHitCnt(v int64) *DescribeVpcFirewallTrafficAssetListResponseBodyDataList {
	s.IpsHitCnt = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBodyDataList) SetSessionCount(v int64) *DescribeVpcFirewallTrafficAssetListResponseBodyDataList {
	s.SessionCount = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBodyDataList) SetTotalBytes(v int64) *DescribeVpcFirewallTrafficAssetListResponseBodyDataList {
	s.TotalBytes = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBodyDataList) SetVpcId(v string) *DescribeVpcFirewallTrafficAssetListResponseBodyDataList {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBodyDataList) SetVpcName(v string) *DescribeVpcFirewallTrafficAssetListResponseBodyDataList {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcFirewallTrafficAssetListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
