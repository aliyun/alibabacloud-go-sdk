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
	// The charts.
	ChartList []*DescribeChartListResponseBodyChartList `json:"ChartList,omitempty" xml:"ChartList,omitempty" type:"Repeated"`
	// The request ID.
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
	return dara.Validate(s)
}

type DescribeChartListResponseBodyChartList struct {
	// The name of the business type. Valid values:
	//
	// 	- Overall Operations Metrics
	//
	// 	- Asset Operations Metrics
	//
	// 	- Security Alert Operations Metrics
	//
	// 	- Vulnerability Operations Metrics
	//
	// 	- Baseline Operations Metrics
	//
	// 	- Cloud Product Operations Metrics
	//
	// 	- Honeypot Operations Metrics
	//
	// example:
	//
	// Overall Operation Metrics
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// The business type. Valid values:
	//
	// 	- INDEX_SECURITY_OVERALL_OPERATION
	//
	// 	- INDEX_ASSET_OPERATION
	//
	// 	- INDEX_SUSPICIOUS_OPERATION
	//
	// 	- INDEX_VUL_OPERATION
	//
	// 	- INDEX_BASELINE_CHECK_OPERATION
	//
	// 	- INDEX_CLOUD_ASSET_OPERATION
	//
	// 	- INDEX_HONEYPOT_RISK_OPERATION
	//
	// example:
	//
	// INDEX_SECURITY_OVERALL_OPERATION
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The ID of the chart.
	//
	// example:
	//
	// CID_CLOUD_ASSET_SUMMARY
	ChartId *string `json:"ChartId,omitempty" xml:"ChartId,omitempty"`
	// The name of the chart.
	//
	// example:
	//
	// Security Score
	ChartName *string `json:"ChartName,omitempty" xml:"ChartName,omitempty"`
	// The type of the chart. Valid values:
	//
	// 	- **text**
	//
	// 	- **table**
	//
	// 	- **gauge**
	//
	// 	- **pie**
	//
	// 	- **line**
	//
	// 	- **bar**
	//
	// 	- **timeBar**
	//
	// 	- **timeLine**
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
