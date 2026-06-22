// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeChartListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChartList(v []*DescribeChartListResponseBodyChartList) *DescribeChartListResponseBody
	GetChartList() []*DescribeChartListResponseBodyChartList
	SetRequestId(v string) *DescribeChartListResponseBody
	GetRequestId() *string
}

type DescribeChartListResponseBody struct {
	// The list of charts.
	ChartList []*DescribeChartListResponseBodyChartList `json:"ChartList,omitempty" xml:"ChartList,omitempty" type:"Repeated"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 35B434CC-1615-5937-A04E-A9BC2868DB45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChartListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeChartListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChartListResponseBody) GetChartList() []*DescribeChartListResponseBodyChartList {
	return s.ChartList
}

func (s *DescribeChartListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeChartListResponseBody) SetChartList(v []*DescribeChartListResponseBodyChartList) *DescribeChartListResponseBody {
	s.ChartList = v
	return s
}

func (s *DescribeChartListResponseBody) SetRequestId(v string) *DescribeChartListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChartListResponseBody) Validate() error {
	if s.ChartList != nil {
		for _, item := range s.ChartList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeChartListResponseBodyChartList struct {
	// The business type name. Valid values:
	//
	// - Overall operation metrics
	//
	// - Asset operation metrics
	//
	// - Security alert operation metrics
	//
	// - Vulnerability operation metrics
	//
	// - Baseline operation metrics
	//
	// - Cloud service operation metrics
	//
	// - Cloud honeypot operation metrics.
	//
	// example:
	//
	// Overall Operation Metrics
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// The business type. Valid values:
	//
	// - INDEX_SECURITY_OVERALL_OPERATION
	//
	// - INDEX_ASSET_OPERATION
	//
	// - INDEX_SUSPICIOUS_OPERATION
	//
	// - INDEX_VUL_OPERATION
	//
	// - INDEX_BASELINE_CHECK_OPERATION
	//
	// - INDEX_CLOUD_ASSET_OPERATION
	//
	// - INDEX_HONEYPOT_RISK_OPERATION.
	//
	// example:
	//
	// INDEX_SECURITY_OVERALL_OPERATION
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The chart ID.
	//
	// example:
	//
	// CID_CLOUD_ASSET_SUMMARY
	ChartId *string `json:"ChartId,omitempty" xml:"ChartId,omitempty"`
	// The chart name.
	//
	// example:
	//
	// Security Score
	ChartName *string `json:"ChartName,omitempty" xml:"ChartName,omitempty"`
	// The chart type. Valid values:
	//
	// - **text**: text
	//
	// - **table**: table
	//
	// - **gauge**: gauge chart
	//
	// - **pie**: pie chart
	//
	// - **line**: line chart
	//
	// - **bar**: bar chart
	//
	// - **timeBar**: timeline bar chart
	//
	// - **timeLine**: timeline line chart.
	//
	// example:
	//
	// text
	ChartType *string `json:"ChartType,omitempty" xml:"ChartType,omitempty"`
}

func (s DescribeChartListResponseBodyChartList) String() string {
	return dara.Prettify(s)
}

func (s DescribeChartListResponseBodyChartList) GoString() string {
	return s.String()
}

func (s *DescribeChartListResponseBodyChartList) GetBusinessName() *string {
	return s.BusinessName
}

func (s *DescribeChartListResponseBodyChartList) GetBusinessType() *string {
	return s.BusinessType
}

func (s *DescribeChartListResponseBodyChartList) GetChartId() *string {
	return s.ChartId
}

func (s *DescribeChartListResponseBodyChartList) GetChartName() *string {
	return s.ChartName
}

func (s *DescribeChartListResponseBodyChartList) GetChartType() *string {
	return s.ChartType
}

func (s *DescribeChartListResponseBodyChartList) SetBusinessName(v string) *DescribeChartListResponseBodyChartList {
	s.BusinessName = &v
	return s
}

func (s *DescribeChartListResponseBodyChartList) SetBusinessType(v string) *DescribeChartListResponseBodyChartList {
	s.BusinessType = &v
	return s
}

func (s *DescribeChartListResponseBodyChartList) SetChartId(v string) *DescribeChartListResponseBodyChartList {
	s.ChartId = &v
	return s
}

func (s *DescribeChartListResponseBodyChartList) SetChartName(v string) *DescribeChartListResponseBodyChartList {
	s.ChartName = &v
	return s
}

func (s *DescribeChartListResponseBodyChartList) SetChartType(v string) *DescribeChartListResponseBodyChartList {
	s.ChartType = &v
	return s
}

func (s *DescribeChartListResponseBodyChartList) Validate() error {
	return dara.Validate(s)
}
