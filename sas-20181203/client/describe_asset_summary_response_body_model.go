// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssetsSummary(v *DescribeAssetSummaryResponseBodyAssetsSummary) *DescribeAssetSummaryResponseBody
	GetAssetsSummary() *DescribeAssetSummaryResponseBodyAssetsSummary
	SetRequestId(v string) *DescribeAssetSummaryResponseBody
	GetRequestId() *string
}

type DescribeAssetSummaryResponseBody struct {
	// The statistical information about the assets.
	AssetsSummary *DescribeAssetSummaryResponseBodyAssetsSummary `json:"AssetsSummary,omitempty" xml:"AssetsSummary,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0FA7F1F4-488D-52CA-9BFC-3E47793B49D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAssetSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAssetSummaryResponseBody) GetAssetsSummary() *DescribeAssetSummaryResponseBodyAssetsSummary {
	return s.AssetsSummary
}

func (s *DescribeAssetSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAssetSummaryResponseBody) SetAssetsSummary(v *DescribeAssetSummaryResponseBodyAssetsSummary) *DescribeAssetSummaryResponseBody {
	s.AssetsSummary = v
	return s
}

func (s *DescribeAssetSummaryResponseBody) SetRequestId(v string) *DescribeAssetSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAssetSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAssetSummaryResponseBodyAssetsSummary struct {
	// The total number of protected assets in all regions.
	//
	// >  Security Center supports the Hangzhou and Singapore service centers, which separately correspond to the China and Outside China data management centers. In the Hangzhou service center, Security Center provides protection capabilities for assets that are deployed in the regions covered by the China data management center. In the Singapore service center, Security Center provides protection capabilities for assets that are deployed in the regions covered by the Outside China data management center. You can determine whether the current region is covered by the China data management center or by the Outside China data management center based on the endpoint of Security Center. For more information about the supported regions for each data management center, see [What is Security Center?](https://help.aliyun.com/document_detail/42302.html)
	//
	// example:
	//
	// 2064
	TotalAssetAllRegion *int32 `json:"TotalAssetAllRegion,omitempty" xml:"TotalAssetAllRegion,omitempty"`
	// The total number of cores of protected assets in all regions.
	//
	// >  Security Center supports the Hangzhou and Singapore service centers, which separately correspond to the China and Outside China data management centers. In the Hangzhou service center, Security Center provides protection capabilities for assets that are deployed in the regions covered by the China data management center. In the Singapore service center, Security Center provides protection capabilities for assets that are deployed in the regions covered by the Outside China data management center. You can determine whether the current region is covered by the China data management center or by the Outside China data management center based on the endpoint of Security Center. For more information about the supported regions for each data management center, see [What is Security Center?](https://help.aliyun.com/document_detail/42302.html)
	//
	// example:
	//
	// 3200
	TotalCoreAllRegion *int32 `json:"TotalCoreAllRegion,omitempty" xml:"TotalCoreAllRegion,omitempty"`
	// The total number of cores of protected assets in the current region.
	//
	// >  Security Center supports the Hangzhou and Singapore service centers, which separately correspond to the China and Outside China data management centers. In the Hangzhou service center, Security Center provides protection capabilities for assets that are deployed in the regions covered by the China data management center. In the Singapore service center, Security Center provides protection capabilities for assets that are deployed in the regions covered by the Outside China data management center. You can determine whether the current region is covered by the China data management center or by the Outside China data management center based on the endpoint of Security Center. For more information about the supported regions for each data management center, see [What is Security Center?](https://help.aliyun.com/document_detail/42302.html)
	//
	// example:
	//
	// 1022
	TotalCoreNum *int32 `json:"TotalCoreNum,omitempty" xml:"TotalCoreNum,omitempty"`
}

func (s DescribeAssetSummaryResponseBodyAssetsSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetSummaryResponseBodyAssetsSummary) GoString() string {
	return s.String()
}

func (s *DescribeAssetSummaryResponseBodyAssetsSummary) GetTotalAssetAllRegion() *int32 {
	return s.TotalAssetAllRegion
}

func (s *DescribeAssetSummaryResponseBodyAssetsSummary) GetTotalCoreAllRegion() *int32 {
	return s.TotalCoreAllRegion
}

func (s *DescribeAssetSummaryResponseBodyAssetsSummary) GetTotalCoreNum() *int32 {
	return s.TotalCoreNum
}

func (s *DescribeAssetSummaryResponseBodyAssetsSummary) SetTotalAssetAllRegion(v int32) *DescribeAssetSummaryResponseBodyAssetsSummary {
	s.TotalAssetAllRegion = &v
	return s
}

func (s *DescribeAssetSummaryResponseBodyAssetsSummary) SetTotalCoreAllRegion(v int32) *DescribeAssetSummaryResponseBodyAssetsSummary {
	s.TotalCoreAllRegion = &v
	return s
}

func (s *DescribeAssetSummaryResponseBodyAssetsSummary) SetTotalCoreNum(v int32) *DescribeAssetSummaryResponseBodyAssetsSummary {
	s.TotalCoreNum = &v
	return s
}

func (s *DescribeAssetSummaryResponseBodyAssetsSummary) Validate() error {
	return dara.Validate(s)
}
