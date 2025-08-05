// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetRiskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetList(v []*DescribeAssetRiskListResponseBodyAssetList) *DescribeAssetRiskListResponseBody
	GetAssetList() []*DescribeAssetRiskListResponseBodyAssetList
	SetRequestId(v string) *DescribeAssetRiskListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAssetRiskListResponseBody
	GetTotalCount() *int64
}

type DescribeAssetRiskListResponseBody struct {
	// The details of the assets.
	AssetList []*DescribeAssetRiskListResponseBodyAssetList `json:"AssetList,omitempty" xml:"AssetList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 443C5781-1C03-5FCD-8EC5-FB9C0B9AC396
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAssetRiskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetRiskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetRiskListResponseBody) GetAssetList() []*DescribeAssetRiskListResponseBodyAssetList {
	return s.AssetList
}

func (s *DescribeAssetRiskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAssetRiskListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAssetRiskListResponseBody) SetAssetList(v []*DescribeAssetRiskListResponseBodyAssetList) *DescribeAssetRiskListResponseBody {
	s.AssetList = v
	return s
}

func (s *DescribeAssetRiskListResponseBody) SetRequestId(v string) *DescribeAssetRiskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetRiskListResponseBody) SetTotalCount(v int64) *DescribeAssetRiskListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAssetRiskListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAssetRiskListResponseBodyAssetList struct {
	// The IP address of the server.
	//
	// example:
	//
	// 39.108.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The IP version of the asset that is protected by Cloud Firewall.
	//
	// Valid values:
	//
	// 	- **4**: IPv4
	//
	// 	- **6**: IPv6
	//
	// example:
	//
	// 4
	IpVersion *int64 `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The reason for the risk.
	//
	// example:
	//
	// other
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **low**
	//
	// 	- **middle**
	//
	// 	- **high**
	//
	// example:
	//
	// low
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s DescribeAssetRiskListResponseBodyAssetList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetRiskListResponseBodyAssetList) GoString() string {
	return s.String()
}

func (s *DescribeAssetRiskListResponseBodyAssetList) GetIp() *string {
	return s.Ip
}

func (s *DescribeAssetRiskListResponseBodyAssetList) GetIpVersion() *int64 {
	return s.IpVersion
}

func (s *DescribeAssetRiskListResponseBodyAssetList) GetReason() *string {
	return s.Reason
}

func (s *DescribeAssetRiskListResponseBodyAssetList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *DescribeAssetRiskListResponseBodyAssetList) SetIp(v string) *DescribeAssetRiskListResponseBodyAssetList {
	s.Ip = &v
	return s
}

func (s *DescribeAssetRiskListResponseBodyAssetList) SetIpVersion(v int64) *DescribeAssetRiskListResponseBodyAssetList {
	s.IpVersion = &v
	return s
}

func (s *DescribeAssetRiskListResponseBodyAssetList) SetReason(v string) *DescribeAssetRiskListResponseBodyAssetList {
	s.Reason = &v
	return s
}

func (s *DescribeAssetRiskListResponseBodyAssetList) SetRiskLevel(v string) *DescribeAssetRiskListResponseBodyAssetList {
	s.RiskLevel = &v
	return s
}

func (s *DescribeAssetRiskListResponseBodyAssetList) Validate() error {
	return dara.Validate(s)
}
