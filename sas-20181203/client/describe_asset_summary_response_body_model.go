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
	// The asset statistics information.
	AssetsSummary *DescribeAssetSummaryResponseBodyAssetsSummary `json:"AssetsSummary,omitempty" xml:"AssetsSummary,omitempty" type:"Struct"`
	// The ID of the request. The ID is a unique identifier that Alibaba Cloud generates for the request. You can use the ID to troubleshoot issues.
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
	if s.AssetsSummary != nil {
		if err := s.AssetsSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAssetSummaryResponseBodyAssetsSummary struct {
	// The total number of assets across all regions.
	//
	// >Security Center uses independent service centers in the Chinese mainland and outside the Chinese mainland. You can check the endpoint to which you are connected to determine the current service region. For more information about the regions included in each service region, see [What is Security Center?](https://help.aliyun.com/document_detail/42302.html).
	//
	// example:
	//
	// 2064
	TotalAssetAllRegion *int32 `json:"TotalAssetAllRegion,omitempty" xml:"TotalAssetAllRegion,omitempty"`
	// The total number of cores of assets across all regions.
	//
	// >Security Center uses independent service centers in the Chinese mainland and outside the Chinese mainland. You can check the endpoint to which you are connected to determine the current service region. For more information about the regions included in each service region, see [What is Security Center?](https://help.aliyun.com/document_detail/42302.html).
	//
	// example:
	//
	// 3200
	TotalCoreAllRegion *int32 `json:"TotalCoreAllRegion,omitempty" xml:"TotalCoreAllRegion,omitempty"`
	// The total number of cores of assets in the current region.
	//
	// >Security Center uses independent service centers in the Chinese mainland and outside the Chinese mainland. You can check the endpoint to which you are connected to determine the current service region. For more information about the regions included in each service region, see [What is Security Center?](https://help.aliyun.com/document_detail/42302.html).
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
