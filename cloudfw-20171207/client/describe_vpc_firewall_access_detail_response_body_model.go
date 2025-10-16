// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallAccessDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeVpcFirewallAccessDetailResponseBodyDataList) *DescribeVpcFirewallAccessDetailResponseBody
	GetDataList() []*DescribeVpcFirewallAccessDetailResponseBodyDataList
	SetRequestId(v string) *DescribeVpcFirewallAccessDetailResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcFirewallAccessDetailResponseBody
	GetTotalCount() *int32
}

type DescribeVpcFirewallAccessDetailResponseBody struct {
	DataList []*DescribeVpcFirewallAccessDetailResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// example:
	//
	// 8EAC2347-E85A-5DFF-9F49-B8E1BAFB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 9
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcFirewallAccessDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAccessDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAccessDetailResponseBody) GetDataList() []*DescribeVpcFirewallAccessDetailResponseBodyDataList {
	return s.DataList
}

func (s *DescribeVpcFirewallAccessDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallAccessDetailResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcFirewallAccessDetailResponseBody) SetDataList(v []*DescribeVpcFirewallAccessDetailResponseBodyDataList) *DescribeVpcFirewallAccessDetailResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponseBody) SetRequestId(v string) *DescribeVpcFirewallAccessDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallAccessDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponseBody) Validate() error {
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

type DescribeVpcFirewallAccessDetailResponseBodyDataList struct {
	// example:
	//
	// 17845821669.0
	InBytes *int64 `json:"InBytes,omitempty" xml:"InBytes,omitempty"`
	// example:
	//
	// 1123
	OutBytes *int64 `json:"OutBytes,omitempty" xml:"OutBytes,omitempty"`
	// example:
	//
	// 10.125.1.XX
	PeerAssetIP *string `json:"PeerAssetIP,omitempty" xml:"PeerAssetIP,omitempty"`
	// example:
	//
	// i-123451
	PeerAssetInstanceId *string `json:"PeerAssetInstanceId,omitempty" xml:"PeerAssetInstanceId,omitempty"`
	// example:
	//
	// test
	PeerAssetInstanceName *string `json:"PeerAssetInstanceName,omitempty" xml:"PeerAssetInstanceName,omitempty"`
	// example:
	//
	// vpc-123411
	PeerVpcId *string `json:"PeerVpcId,omitempty" xml:"PeerVpcId,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// example:
	//
	// 27
	SessionCount *int64 `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	// example:
	//
	// vpc-test
	PeerVpcName *string `json:"peerVpcName,omitempty" xml:"peerVpcName,omitempty"`
}

func (s DescribeVpcFirewallAccessDetailResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAccessDetailResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) GetInBytes() *int64 {
	return s.InBytes
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) GetOutBytes() *int64 {
	return s.OutBytes
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) GetPeerAssetIP() *string {
	return s.PeerAssetIP
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) GetPeerAssetInstanceId() *string {
	return s.PeerAssetInstanceId
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) GetPeerAssetInstanceName() *string {
	return s.PeerAssetInstanceName
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) GetPeerVpcId() *string {
	return s.PeerVpcId
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) GetSessionCount() *int64 {
	return s.SessionCount
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) GetPeerVpcName() *string {
	return s.PeerVpcName
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) SetInBytes(v int64) *DescribeVpcFirewallAccessDetailResponseBodyDataList {
	s.InBytes = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) SetOutBytes(v int64) *DescribeVpcFirewallAccessDetailResponseBodyDataList {
	s.OutBytes = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) SetPeerAssetIP(v string) *DescribeVpcFirewallAccessDetailResponseBodyDataList {
	s.PeerAssetIP = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) SetPeerAssetInstanceId(v string) *DescribeVpcFirewallAccessDetailResponseBodyDataList {
	s.PeerAssetInstanceId = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) SetPeerAssetInstanceName(v string) *DescribeVpcFirewallAccessDetailResponseBodyDataList {
	s.PeerAssetInstanceName = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) SetPeerVpcId(v string) *DescribeVpcFirewallAccessDetailResponseBodyDataList {
	s.PeerVpcId = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) SetRegionNo(v string) *DescribeVpcFirewallAccessDetailResponseBodyDataList {
	s.RegionNo = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) SetSessionCount(v int64) *DescribeVpcFirewallAccessDetailResponseBodyDataList {
	s.SessionCount = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) SetPeerVpcName(v string) *DescribeVpcFirewallAccessDetailResponseBodyDataList {
	s.PeerVpcName = &v
	return s
}

func (s *DescribeVpcFirewallAccessDetailResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
