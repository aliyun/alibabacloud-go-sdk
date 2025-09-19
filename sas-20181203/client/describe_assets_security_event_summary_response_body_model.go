// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetsSecurityEventSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssets(v []*DescribeAssetsSecurityEventSummaryResponseBodyAssets) *DescribeAssetsSecurityEventSummaryResponseBody
	GetAssets() []*DescribeAssetsSecurityEventSummaryResponseBodyAssets
	SetRequestId(v string) *DescribeAssetsSecurityEventSummaryResponseBody
	GetRequestId() *string
}

type DescribeAssetsSecurityEventSummaryResponseBody struct {
	// An array that consists of risk information about containers.
	Assets []*DescribeAssetsSecurityEventSummaryResponseBodyAssets `json:"Assets,omitempty" xml:"Assets,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D03DD0FD-6041-5107-AC00-383E28F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAssetsSecurityEventSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetsSecurityEventSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetsSecurityEventSummaryResponseBody) GetAssets() []*DescribeAssetsSecurityEventSummaryResponseBodyAssets {
	return s.Assets
}

func (s *DescribeAssetsSecurityEventSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAssetsSecurityEventSummaryResponseBody) SetAssets(v []*DescribeAssetsSecurityEventSummaryResponseBodyAssets) *DescribeAssetsSecurityEventSummaryResponseBody {
	s.Assets = v
	return s
}

func (s *DescribeAssetsSecurityEventSummaryResponseBody) SetRequestId(v string) *DescribeAssetsSecurityEventSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetsSecurityEventSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAssetsSecurityEventSummaryResponseBodyAssets struct {
	// The type of the asset. Valid values:
	//
	// 	- **namespace**
	//
	// 	- **clusters**
	//
	// 	- **applications**
	//
	// 	- **pods**
	//
	// 	- **containers**
	//
	// 	- **images**
	//
	// 	- **hosts**
	//
	// example:
	//
	// namespace
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The number of potential risky assets.
	//
	// example:
	//
	// 16
	RiskCount *int64 `json:"RiskCount,omitempty" xml:"RiskCount,omitempty"`
	// The total number of assets.
	//
	// example:
	//
	// 30
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAssetsSecurityEventSummaryResponseBodyAssets) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetsSecurityEventSummaryResponseBodyAssets) GoString() string {
	return s.String()
}

func (s *DescribeAssetsSecurityEventSummaryResponseBodyAssets) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeAssetsSecurityEventSummaryResponseBodyAssets) GetRiskCount() *int64 {
	return s.RiskCount
}

func (s *DescribeAssetsSecurityEventSummaryResponseBodyAssets) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAssetsSecurityEventSummaryResponseBodyAssets) SetAssetType(v string) *DescribeAssetsSecurityEventSummaryResponseBodyAssets {
	s.AssetType = &v
	return s
}

func (s *DescribeAssetsSecurityEventSummaryResponseBodyAssets) SetRiskCount(v int64) *DescribeAssetsSecurityEventSummaryResponseBodyAssets {
	s.RiskCount = &v
	return s
}

func (s *DescribeAssetsSecurityEventSummaryResponseBodyAssets) SetTotalCount(v int64) *DescribeAssetsSecurityEventSummaryResponseBodyAssets {
	s.TotalCount = &v
	return s
}

func (s *DescribeAssetsSecurityEventSummaryResponseBodyAssets) Validate() error {
	return dara.Validate(s)
}
