// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcFirewallSummaryInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableVpcFirewallQuota(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetAvailableVpcFirewallQuota() *int32
	SetCenExpressConnectVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetCenExpressConnectVpcCount() *int32
	SetCenFirewallVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetCenFirewallVpcCount() *int32
	SetCenTrVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetCenTrVpcCount() *int32
	SetClosedCenFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetClosedCenFirewallCount() *int32
	SetClosedExpressConnectFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetClosedExpressConnectFirewallCount() *int32
	SetClosedVpcFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetClosedVpcFirewallCount() *int32
	SetConfiguredCenFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetConfiguredCenFirewallCount() *int32
	SetConfiguredCenFirewallRegionCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetConfiguredCenFirewallRegionCount() *int32
	SetConfiguredCenFirewallVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetConfiguredCenFirewallVpcCount() *int32
	SetConfiguredCenTrFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetConfiguredCenTrFirewallCount() *int32
	SetConfiguredExpressConnectFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetConfiguredExpressConnectFirewallCount() *int32
	SetConfiguredExpressConnectVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetConfiguredExpressConnectVpcCount() *int32
	SetConfiguredVpcFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetConfiguredVpcFirewallCount() *int32
	SetConfiguredVpcFirewallVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetConfiguredVpcFirewallVpcCount() *int32
	SetExpressConnectVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetExpressConnectVpcCount() *int32
	SetNotConfiguredCenFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetNotConfiguredCenFirewallCount() *int32
	SetNotConfiguredCenTrFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetNotConfiguredCenTrFirewallCount() *int32
	SetNotConfiguredExpressConnectFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetNotConfiguredExpressConnectFirewallCount() *int32
	SetNotConfiguredVpcFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetNotConfiguredVpcFirewallCount() *int32
	SetOpenedCenExpressConnectVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetOpenedCenExpressConnectVpcCount() *int32
	SetOpenedCenFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetOpenedCenFirewallCount() *int32
	SetOpenedCenFirewallVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetOpenedCenFirewallVpcCount() *int32
	SetOpenedCenTrFirewallVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetOpenedCenTrFirewallVpcCount() *int32
	SetOpenedEcrCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetOpenedEcrCount() *int32
	SetOpenedExpressConnectFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetOpenedExpressConnectFirewallCount() *int32
	SetOpenedExpressConnectVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetOpenedExpressConnectVpcCount() *int32
	SetOpenedPeerTrCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetOpenedPeerTrCount() *int32
	SetOpenedVbrCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetOpenedVbrCount() *int32
	SetOpenedVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetOpenedVpcCount() *int32
	SetOpenedVpcFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetOpenedVpcFirewallCount() *int32
	SetOpenedVpnCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetOpenedVpnCount() *int32
	SetRequestId(v string) *DescribeVpcFirewallSummaryInfoResponseBody
	GetRequestId() *string
	SetTotalEcrCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetTotalEcrCount() *int32
	SetTotalPeerTrCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetTotalPeerTrCount() *int32
	SetTotalVbrCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetTotalVbrCount() *int32
	SetTotalVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetTotalVpcCount() *int32
	SetTotalVpcFirewallQuota(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetTotalVpcFirewallQuota() *int32
	SetTotalVpnCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody
	GetTotalVpnCount() *int32
}

type DescribeVpcFirewallSummaryInfoResponseBody struct {
	// example:
	//
	// 4
	AvailableVpcFirewallQuota *int32 `json:"AvailableVpcFirewallQuota,omitempty" xml:"AvailableVpcFirewallQuota,omitempty"`
	// example:
	//
	// 7
	CenExpressConnectVpcCount *int32 `json:"CenExpressConnectVpcCount,omitempty" xml:"CenExpressConnectVpcCount,omitempty"`
	// example:
	//
	// 10
	CenFirewallVpcCount *int32 `json:"CenFirewallVpcCount,omitempty" xml:"CenFirewallVpcCount,omitempty"`
	// example:
	//
	// 13
	CenTrVpcCount *int32 `json:"CenTrVpcCount,omitempty" xml:"CenTrVpcCount,omitempty"`
	// example:
	//
	// 1
	ClosedCenFirewallCount *int32 `json:"ClosedCenFirewallCount,omitempty" xml:"ClosedCenFirewallCount,omitempty"`
	// example:
	//
	// 10
	ClosedExpressConnectFirewallCount *int32 `json:"ClosedExpressConnectFirewallCount,omitempty" xml:"ClosedExpressConnectFirewallCount,omitempty"`
	// example:
	//
	// 5
	ClosedVpcFirewallCount *int32 `json:"ClosedVpcFirewallCount,omitempty" xml:"ClosedVpcFirewallCount,omitempty"`
	// example:
	//
	// 5
	ConfiguredCenFirewallCount *int32 `json:"ConfiguredCenFirewallCount,omitempty" xml:"ConfiguredCenFirewallCount,omitempty"`
	// example:
	//
	// 2
	ConfiguredCenFirewallRegionCount *int32 `json:"ConfiguredCenFirewallRegionCount,omitempty" xml:"ConfiguredCenFirewallRegionCount,omitempty"`
	// example:
	//
	// 18
	ConfiguredCenFirewallVpcCount *int32 `json:"ConfiguredCenFirewallVpcCount,omitempty" xml:"ConfiguredCenFirewallVpcCount,omitempty"`
	// example:
	//
	// 2
	ConfiguredCenTrFirewallCount *int32 `json:"ConfiguredCenTrFirewallCount,omitempty" xml:"ConfiguredCenTrFirewallCount,omitempty"`
	// example:
	//
	// 2
	ConfiguredExpressConnectFirewallCount *int32 `json:"ConfiguredExpressConnectFirewallCount,omitempty" xml:"ConfiguredExpressConnectFirewallCount,omitempty"`
	// example:
	//
	// 2
	ConfiguredExpressConnectVpcCount *int32 `json:"ConfiguredExpressConnectVpcCount,omitempty" xml:"ConfiguredExpressConnectVpcCount,omitempty"`
	// example:
	//
	// 5
	ConfiguredVpcFirewallCount *int32 `json:"ConfiguredVpcFirewallCount,omitempty" xml:"ConfiguredVpcFirewallCount,omitempty"`
	// example:
	//
	// 1
	ConfiguredVpcFirewallVpcCount *int32 `json:"ConfiguredVpcFirewallVpcCount,omitempty" xml:"ConfiguredVpcFirewallVpcCount,omitempty"`
	// example:
	//
	// 2
	ExpressConnectVpcCount *int32 `json:"ExpressConnectVpcCount,omitempty" xml:"ExpressConnectVpcCount,omitempty"`
	// example:
	//
	// 0
	NotConfiguredCenFirewallCount *int32 `json:"NotConfiguredCenFirewallCount,omitempty" xml:"NotConfiguredCenFirewallCount,omitempty"`
	// example:
	//
	// 6
	NotConfiguredCenTrFirewallCount *int32 `json:"NotConfiguredCenTrFirewallCount,omitempty" xml:"NotConfiguredCenTrFirewallCount,omitempty"`
	// example:
	//
	// 7
	NotConfiguredExpressConnectFirewallCount *int32 `json:"NotConfiguredExpressConnectFirewallCount,omitempty" xml:"NotConfiguredExpressConnectFirewallCount,omitempty"`
	// example:
	//
	// 12
	NotConfiguredVpcFirewallCount *int32 `json:"NotConfiguredVpcFirewallCount,omitempty" xml:"NotConfiguredVpcFirewallCount,omitempty"`
	// example:
	//
	// 10
	OpenedCenExpressConnectVpcCount *int32 `json:"OpenedCenExpressConnectVpcCount,omitempty" xml:"OpenedCenExpressConnectVpcCount,omitempty"`
	// example:
	//
	// 4
	OpenedCenFirewallCount *int32 `json:"OpenedCenFirewallCount,omitempty" xml:"OpenedCenFirewallCount,omitempty"`
	// example:
	//
	// 0
	OpenedCenFirewallVpcCount *int32 `json:"OpenedCenFirewallVpcCount,omitempty" xml:"OpenedCenFirewallVpcCount,omitempty"`
	// example:
	//
	// 4
	OpenedCenTrFirewallVpcCount *int32 `json:"OpenedCenTrFirewallVpcCount,omitempty" xml:"OpenedCenTrFirewallVpcCount,omitempty"`
	// example:
	//
	// 0
	OpenedEcrCount *int32 `json:"OpenedEcrCount,omitempty" xml:"OpenedEcrCount,omitempty"`
	// example:
	//
	// 15
	OpenedExpressConnectFirewallCount *int32 `json:"OpenedExpressConnectFirewallCount,omitempty" xml:"OpenedExpressConnectFirewallCount,omitempty"`
	// example:
	//
	// 2
	OpenedExpressConnectVpcCount *int32 `json:"OpenedExpressConnectVpcCount,omitempty" xml:"OpenedExpressConnectVpcCount,omitempty"`
	// example:
	//
	// 3
	OpenedPeerTrCount *int32 `json:"OpenedPeerTrCount,omitempty" xml:"OpenedPeerTrCount,omitempty"`
	// example:
	//
	// 0
	OpenedVbrCount *int32 `json:"OpenedVbrCount,omitempty" xml:"OpenedVbrCount,omitempty"`
	// example:
	//
	// 17
	OpenedVpcCount *int32 `json:"OpenedVpcCount,omitempty" xml:"OpenedVpcCount,omitempty"`
	// example:
	//
	// 9
	OpenedVpcFirewallCount *int32 `json:"OpenedVpcFirewallCount,omitempty" xml:"OpenedVpcFirewallCount,omitempty"`
	// example:
	//
	// 6
	OpenedVpnCount *int32 `json:"OpenedVpnCount,omitempty" xml:"OpenedVpnCount,omitempty"`
	// example:
	//
	// 8AABEF64-7ABF-52CB-BA6C-0598E3DB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalEcrCount *int32 `json:"TotalEcrCount,omitempty" xml:"TotalEcrCount,omitempty"`
	// example:
	//
	// 6
	TotalPeerTrCount *int32 `json:"TotalPeerTrCount,omitempty" xml:"TotalPeerTrCount,omitempty"`
	// example:
	//
	// 5
	TotalVbrCount *int32 `json:"TotalVbrCount,omitempty" xml:"TotalVbrCount,omitempty"`
	// example:
	//
	// 2
	TotalVpcCount *int32 `json:"TotalVpcCount,omitempty" xml:"TotalVpcCount,omitempty"`
	// example:
	//
	// 5
	TotalVpcFirewallQuota *int32 `json:"TotalVpcFirewallQuota,omitempty" xml:"TotalVpcFirewallQuota,omitempty"`
	// example:
	//
	// 1
	TotalVpnCount *int32 `json:"TotalVpnCount,omitempty" xml:"TotalVpnCount,omitempty"`
}

func (s DescribeVpcFirewallSummaryInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcFirewallSummaryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetAvailableVpcFirewallQuota() *int32 {
	return s.AvailableVpcFirewallQuota
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetCenExpressConnectVpcCount() *int32 {
	return s.CenExpressConnectVpcCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetCenFirewallVpcCount() *int32 {
	return s.CenFirewallVpcCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetCenTrVpcCount() *int32 {
	return s.CenTrVpcCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetClosedCenFirewallCount() *int32 {
	return s.ClosedCenFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetClosedExpressConnectFirewallCount() *int32 {
	return s.ClosedExpressConnectFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetClosedVpcFirewallCount() *int32 {
	return s.ClosedVpcFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetConfiguredCenFirewallCount() *int32 {
	return s.ConfiguredCenFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetConfiguredCenFirewallRegionCount() *int32 {
	return s.ConfiguredCenFirewallRegionCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetConfiguredCenFirewallVpcCount() *int32 {
	return s.ConfiguredCenFirewallVpcCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetConfiguredCenTrFirewallCount() *int32 {
	return s.ConfiguredCenTrFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetConfiguredExpressConnectFirewallCount() *int32 {
	return s.ConfiguredExpressConnectFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetConfiguredExpressConnectVpcCount() *int32 {
	return s.ConfiguredExpressConnectVpcCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetConfiguredVpcFirewallCount() *int32 {
	return s.ConfiguredVpcFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetConfiguredVpcFirewallVpcCount() *int32 {
	return s.ConfiguredVpcFirewallVpcCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetExpressConnectVpcCount() *int32 {
	return s.ExpressConnectVpcCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetNotConfiguredCenFirewallCount() *int32 {
	return s.NotConfiguredCenFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetNotConfiguredCenTrFirewallCount() *int32 {
	return s.NotConfiguredCenTrFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetNotConfiguredExpressConnectFirewallCount() *int32 {
	return s.NotConfiguredExpressConnectFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetNotConfiguredVpcFirewallCount() *int32 {
	return s.NotConfiguredVpcFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetOpenedCenExpressConnectVpcCount() *int32 {
	return s.OpenedCenExpressConnectVpcCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetOpenedCenFirewallCount() *int32 {
	return s.OpenedCenFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetOpenedCenFirewallVpcCount() *int32 {
	return s.OpenedCenFirewallVpcCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetOpenedCenTrFirewallVpcCount() *int32 {
	return s.OpenedCenTrFirewallVpcCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetOpenedEcrCount() *int32 {
	return s.OpenedEcrCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetOpenedExpressConnectFirewallCount() *int32 {
	return s.OpenedExpressConnectFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetOpenedExpressConnectVpcCount() *int32 {
	return s.OpenedExpressConnectVpcCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetOpenedPeerTrCount() *int32 {
	return s.OpenedPeerTrCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetOpenedVbrCount() *int32 {
	return s.OpenedVbrCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetOpenedVpcCount() *int32 {
	return s.OpenedVpcCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetOpenedVpcFirewallCount() *int32 {
	return s.OpenedVpcFirewallCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetOpenedVpnCount() *int32 {
	return s.OpenedVpnCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetTotalEcrCount() *int32 {
	return s.TotalEcrCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetTotalPeerTrCount() *int32 {
	return s.TotalPeerTrCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetTotalVbrCount() *int32 {
	return s.TotalVbrCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetTotalVpcCount() *int32 {
	return s.TotalVpcCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetTotalVpcFirewallQuota() *int32 {
	return s.TotalVpcFirewallQuota
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) GetTotalVpnCount() *int32 {
	return s.TotalVpnCount
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetAvailableVpcFirewallQuota(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.AvailableVpcFirewallQuota = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetCenExpressConnectVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.CenExpressConnectVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetCenFirewallVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.CenFirewallVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetCenTrVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.CenTrVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetClosedCenFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.ClosedCenFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetClosedExpressConnectFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.ClosedExpressConnectFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetClosedVpcFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.ClosedVpcFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetConfiguredCenFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.ConfiguredCenFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetConfiguredCenFirewallRegionCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.ConfiguredCenFirewallRegionCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetConfiguredCenFirewallVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.ConfiguredCenFirewallVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetConfiguredCenTrFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.ConfiguredCenTrFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetConfiguredExpressConnectFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.ConfiguredExpressConnectFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetConfiguredExpressConnectVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.ConfiguredExpressConnectVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetConfiguredVpcFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.ConfiguredVpcFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetConfiguredVpcFirewallVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.ConfiguredVpcFirewallVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetExpressConnectVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.ExpressConnectVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetNotConfiguredCenFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.NotConfiguredCenFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetNotConfiguredCenTrFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.NotConfiguredCenTrFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetNotConfiguredExpressConnectFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.NotConfiguredExpressConnectFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetNotConfiguredVpcFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.NotConfiguredVpcFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetOpenedCenExpressConnectVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.OpenedCenExpressConnectVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetOpenedCenFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.OpenedCenFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetOpenedCenFirewallVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.OpenedCenFirewallVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetOpenedCenTrFirewallVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.OpenedCenTrFirewallVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetOpenedEcrCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.OpenedEcrCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetOpenedExpressConnectFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.OpenedExpressConnectFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetOpenedExpressConnectVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.OpenedExpressConnectVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetOpenedPeerTrCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.OpenedPeerTrCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetOpenedVbrCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.OpenedVbrCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetOpenedVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.OpenedVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetOpenedVpcFirewallCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.OpenedVpcFirewallCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetOpenedVpnCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.OpenedVpnCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetRequestId(v string) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetTotalEcrCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.TotalEcrCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetTotalPeerTrCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.TotalPeerTrCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetTotalVbrCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.TotalVbrCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetTotalVpcCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.TotalVpcCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetTotalVpcFirewallQuota(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.TotalVpcFirewallQuota = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) SetTotalVpnCount(v int32) *DescribeVpcFirewallSummaryInfoResponseBody {
	s.TotalVpnCount = &v
	return s
}

func (s *DescribeVpcFirewallSummaryInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
