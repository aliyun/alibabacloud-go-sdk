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
	// The list of monitoring data.
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
	if s.MetricList != nil {
		for _, item := range s.MetricList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutCustomMetricRequestMetricList struct {
	// The dimension map, which is used to query monitoring data of a specified resource. Valid values of N: 1 to 21.
	//
	// Format: a collection of key-value pairs. A commonly used key-value pair collection is: `{"Key":"Value"}`.
	//
	// The length of Key and Value is 1 to 64 characters. Characters beyond the first 64 are truncated.
	//
	// The values of Key and Value can contain letters, digits, periods (.), hyphens (-), underscores (_), forward slashes (/), and backslashes (\\).
	//
	// > Dimensions must be passed in as a JSON string that represents the map object, and must be passed in order.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"sampleName1":"value1","sampleName2":"value2"}
	Dimensions *string `json:"Dimensions,omitempty" xml:"Dimensions,omitempty"`
	// The ID of the application group. Valid values of N: 1 to 21.
	//
	// > If the metric does not belong to any application group, enter 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the metric. Valid values of N: 1 to 21. For more information, see [Metrics of cloud services](https://help.aliyun.com/document_detail/163515.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The aggregation period. Valid values of N: 1 to 21. Unit: seconds. Valid values: 60 and 300.
	//
	// > If the type of the reported value is 1, you must set this parameter.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The time when the metric occurred. Valid values of N: 1 to 21. The following two types of values are supported:
	//
	// - UTC time. Format: YYYY-MM-DDThh:mm:ssZ. For example: 20171012T132456.888+0800.
	//
	// - A Long-type timestamp. For example: 1508136760000.
	//
	// example:
	//
	// 1508136760000
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The type of the reported value. Valid values of N: 1 to 21. Valid values:
	//
	// - 0: raw data.
	//
	// - 1: aggregate data.
	//
	// > When you report aggregate data, we recommend that you report both data with a period of 60 seconds and data with a period of 300 seconds. Otherwise, monitoring data cannot be queried for a time span of more than 7 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The collection of metric values. Valid values of N: 1 to 21.
	//
	// > If the type of the reported value is 0, the raw values are reported. CloudMonitor aggregates raw values into multiple values, such as maximum, count, and sum, by period.
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
