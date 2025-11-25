// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallAssetRegionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionNoList(v []*string) *DescribeVpcFirewallAssetRegionListResponseBody
	GetRegionNoList() []*string
	SetRequestId(v string) *DescribeVpcFirewallAssetRegionListResponseBody
	GetRequestId() *string
}

type DescribeVpcFirewallAssetRegionListResponseBody struct {
	RegionNoList []*string `json:"RegionNoList,omitempty" xml:"RegionNoList,omitempty" type:"Repeated"`
	// example:
	//
	// 00933CCB-65A4-5E51-B180-3D154281****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVpcFirewallAssetRegionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallAssetRegionListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallAssetRegionListResponseBody) GetRegionNoList() []*string {
	return s.RegionNoList
}

func (s *DescribeVpcFirewallAssetRegionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallAssetRegionListResponseBody) SetRegionNoList(v []*string) *DescribeVpcFirewallAssetRegionListResponseBody {
	s.RegionNoList = v
	return s
}

func (s *DescribeVpcFirewallAssetRegionListResponseBody) SetRequestId(v string) *DescribeVpcFirewallAssetRegionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallAssetRegionListResponseBody) Validate() error {
	return dara.Validate(s)
}
