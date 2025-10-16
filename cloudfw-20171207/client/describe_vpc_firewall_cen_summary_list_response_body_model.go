// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallCenSummaryListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCenList(v []*DescribeVpcFirewallCenSummaryListResponseBodyCenList) *DescribeVpcFirewallCenSummaryListResponseBody
	GetCenList() []*DescribeVpcFirewallCenSummaryListResponseBodyCenList
	SetRequestId(v string) *DescribeVpcFirewallCenSummaryListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcFirewallCenSummaryListResponseBody
	GetTotalCount() *int32
}

type DescribeVpcFirewallCenSummaryListResponseBody struct {
	CenList []*DescribeVpcFirewallCenSummaryListResponseBodyCenList `json:"CenList,omitempty" xml:"CenList,omitempty" type:"Repeated"`
	// example:
	//
	// 432D6CCA-5186-5B91-A2B8-10C8994B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVpcFirewallCenSummaryListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenSummaryListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenSummaryListResponseBody) GetCenList() []*DescribeVpcFirewallCenSummaryListResponseBodyCenList {
	return s.CenList
}

func (s *DescribeVpcFirewallCenSummaryListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallCenSummaryListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcFirewallCenSummaryListResponseBody) SetCenList(v []*DescribeVpcFirewallCenSummaryListResponseBodyCenList) *DescribeVpcFirewallCenSummaryListResponseBody {
	s.CenList = v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListResponseBody) SetRequestId(v string) *DescribeVpcFirewallCenSummaryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListResponseBody) SetTotalCount(v int32) *DescribeVpcFirewallCenSummaryListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListResponseBody) Validate() error {
	if s.CenList != nil {
		for _, item := range s.CenList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcFirewallCenSummaryListResponseBodyCenList struct {
	// example:
	//
	// cen-maqfw3abcmjy56****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// example:
	//
	// test
	CenName      *string   `json:"CenName,omitempty" xml:"CenName,omitempty"`
	RegionNoList []*string `json:"RegionNoList,omitempty" xml:"RegionNoList,omitempty" type:"Repeated"`
}

func (s DescribeVpcFirewallCenSummaryListResponseBodyCenList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallCenSummaryListResponseBodyCenList) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallCenSummaryListResponseBodyCenList) GetCenId() *string {
	return s.CenId
}

func (s *DescribeVpcFirewallCenSummaryListResponseBodyCenList) GetCenName() *string {
	return s.CenName
}

func (s *DescribeVpcFirewallCenSummaryListResponseBodyCenList) GetRegionNoList() []*string {
	return s.RegionNoList
}

func (s *DescribeVpcFirewallCenSummaryListResponseBodyCenList) SetCenId(v string) *DescribeVpcFirewallCenSummaryListResponseBodyCenList {
	s.CenId = &v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListResponseBodyCenList) SetCenName(v string) *DescribeVpcFirewallCenSummaryListResponseBodyCenList {
	s.CenName = &v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListResponseBodyCenList) SetRegionNoList(v []*string) *DescribeVpcFirewallCenSummaryListResponseBodyCenList {
	s.RegionNoList = v
	return s
}

func (s *DescribeVpcFirewallCenSummaryListResponseBodyCenList) Validate() error {
	return dara.Validate(s)
}
