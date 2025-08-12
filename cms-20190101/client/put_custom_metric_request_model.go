// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCustomMetricRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetricList(v []*PutCustomMetricRequestMetricList) *PutCustomMetricRequest
	GetMetricList() []*PutCustomMetricRequestMetricList
	SetRegionId(v string) *PutCustomMetricRequest
	GetRegionId() *string
}

type PutCustomMetricRequest struct {
	// The monitoring data.
	//
	// This parameter is required.
	MetricList []*PutCustomMetricRequestMetricList `json:"MetricList,omitempty" xml:"MetricList,omitempty" type:"Repeated"`
	RegionId   *string                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PutCustomMetricRequest) String() string {
	return dara.Prettify(s)
}

func (s PutCustomMetricRequest) GoString() string {
	return s.String()
}

func (s *PutCustomMetricRequest) GetMetricList() []*PutCustomMetricRequestMetricList {
	return s.MetricList
}

func (s *PutCustomMetricRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutCustomMetricRequest) SetMetricList(v []*PutCustomMetricRequestMetricList) *PutCustomMetricRequest {
	s.MetricList = v
	return s
}

func (s *PutCustomMetricRequest) SetRegionId(v string) *PutCustomMetricRequest {
	s.RegionId = &v
	return s
}

func (s *PutCustomMetricRequest) Validate() error {
	return dara.Validate(s)
}

type PutCustomMetricRequestMetricList struct {
	// The dimensions based on which the resources are queried. Valid values of N: 1 to 21.
	//
	// Set this parameter to a collection of key-value pairs. Format: `{"Key":"Value"}`.
	//
	// The key or value must be 1 to 64 bytes in length. Excessive characters are truncated.
	//
	// The key or value can contain letters, digits, periods (.), hyphens (-), underscores (_), forward slashes (/), and backslashes (\\\\).
	//
	// >  Dimensions must be formatted as a JSON string in a specified order.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"sampleName1":"value1","sampleName2":"value2"}
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The ID of the application group. Valid values of N: 1 to 21.
	//
	// >  If the metric does not belong to any application group, enter 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The metric name. Valid values of N: 1 to 21. For more information, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The aggregation period. Valid values of N: 1 to 21. Unit: seconds. Valid values: 60 and 300.
	//
	// >  If the Type parameter is set to 1, the Period parameter is required.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The timestamp when the metric data is generated. Valid values of N: 1 to 21. The timestamp can be in one of the following formats:
	//
	// 	- A UTC timestamp in the YYYY-MM-DDThh:mm:ssZ format. Example: 20171012T132456.888+0800.
	//
	// 	- A UNIX timestamp of the LONG type. Example: 1508136760000.
	//
	// example:
	//
	// 1508136760000
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The type of the monitoring data. Valid values of N: 1 to 21. Valid values:
	//
	// 	- 0: raw data
	//
	// 	- 1: aggregate data
	//
	// >  We recommend that you report aggregate data in both the aggregation periods of 60 seconds and 300 seconds. Otherwise, you cannot query monitoring data in a time span that is more than seven days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The collection of metric values. Valid values of N: 1 to 21.
	//
	// >  If the Type parameter is set to 0, the keys in this parameter must be set to the specified value. CloudMonitor aggregates raw data in each aggregation period to generate multiple statistical values, such as the maximum value, the count, and the total value.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"value":10.5}
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s PutCustomMetricRequestMetricList) String() string {
	return dara.Prettify(s)
}

func (s PutCustomMetricRequestMetricList) GoString() string {
	return s.String()
}

func (s *PutCustomMetricRequestMetricList) GetDimensions() *string {
	return s.Dimensions
}

func (s *PutCustomMetricRequestMetricList) GetGroupId() *string {
	return s.GroupId
}

func (s *PutCustomMetricRequestMetricList) GetMetricName() *string {
	return s.MetricName
}

func (s *PutCustomMetricRequestMetricList) GetPeriod() *string {
	return s.Period
}

func (s *PutCustomMetricRequestMetricList) GetTime() *string {
	return s.Time
}

func (s *PutCustomMetricRequestMetricList) GetType() *string {
	return s.Type
}

func (s *PutCustomMetricRequestMetricList) GetValues() *string {
	return s.Values
}

func (s *PutCustomMetricRequestMetricList) SetDimensions(v string) *PutCustomMetricRequestMetricList {
	s.Dimensions = &v
	return s
}

func (s *PutCustomMetricRequestMetricList) SetGroupId(v string) *PutCustomMetricRequestMetricList {
	s.GroupId = &v
	return s
}

func (s *PutCustomMetricRequestMetricList) SetMetricName(v string) *PutCustomMetricRequestMetricList {
	s.MetricName = &v
	return s
}

func (s *PutCustomMetricRequestMetricList) SetPeriod(v string) *PutCustomMetricRequestMetricList {
	s.Period = &v
	return s
}

func (s *PutCustomMetricRequestMetricList) SetTime(v string) *PutCustomMetricRequestMetricList {
	s.Time = &v
	return s
}

func (s *PutCustomMetricRequestMetricList) SetType(v string) *PutCustomMetricRequestMetricList {
	s.Type = &v
	return s
}

func (s *PutCustomMetricRequestMetricList) SetValues(v string) *PutCustomMetricRequestMetricList {
	s.Values = &v
	return s
}

func (s *PutCustomMetricRequestMetricList) Validate() error {
	return dara.Validate(s)
}
