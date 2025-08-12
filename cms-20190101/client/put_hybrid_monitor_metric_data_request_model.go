// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutHybridMonitorMetricDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMetricList(v []*PutHybridMonitorMetricDataRequestMetricList) *PutHybridMonitorMetricDataRequest
	GetMetricList() []*PutHybridMonitorMetricDataRequestMetricList
	SetNamespace(v string) *PutHybridMonitorMetricDataRequest
	GetNamespace() *string
	SetRegionId(v string) *PutHybridMonitorMetricDataRequest
	GetRegionId() *string
}

type PutHybridMonitorMetricDataRequest struct {
	// The monitoring data.
	//
	// Valid values of N: 1 to 100.
	//
	// This parameter is required.
	MetricList []*PutHybridMonitorMetricDataRequestMetricList `json:"MetricList,omitempty" xml:"MetricList,omitempty" type:"Repeated"`
	// The name of the namespace.
	//
	// For information about how to obtain the name of a namespace, see [DescribeHybridMonitorNamespaceList](https://help.aliyun.com/document_detail/428880.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// default-aliyun
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PutHybridMonitorMetricDataRequest) String() string {
	return dara.Prettify(s)
}

func (s PutHybridMonitorMetricDataRequest) GoString() string {
	return s.String()
}

func (s *PutHybridMonitorMetricDataRequest) GetMetricList() []*PutHybridMonitorMetricDataRequestMetricList {
	return s.MetricList
}

func (s *PutHybridMonitorMetricDataRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PutHybridMonitorMetricDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutHybridMonitorMetricDataRequest) SetMetricList(v []*PutHybridMonitorMetricDataRequestMetricList) *PutHybridMonitorMetricDataRequest {
	s.MetricList = v
	return s
}

func (s *PutHybridMonitorMetricDataRequest) SetNamespace(v string) *PutHybridMonitorMetricDataRequest {
	s.Namespace = &v
	return s
}

func (s *PutHybridMonitorMetricDataRequest) SetRegionId(v string) *PutHybridMonitorMetricDataRequest {
	s.RegionId = &v
	return s
}

func (s *PutHybridMonitorMetricDataRequest) Validate() error {
	return dara.Validate(s)
}

type PutHybridMonitorMetricDataRequestMetricList struct {
	// The tags of the metric.
	//
	// Valid values of N: 1 to 100.
	//
	// example:
	//
	// app、ip、hostName等标识信息
	Labels []*PutHybridMonitorMetricDataRequestMetricListLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The metric name.
	//
	// Valid values of N: 1 to 100.
	//
	// The name can contain letters, digits, and underscores (_). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// CPU_Usage
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the monitoring data is imported. The value is a timestamp.
	//
	// Valid values of N: 1 to 100.
	//
	// Unit: milliseconds. By default, the current time is used.
	//
	// example:
	//
	// 1640776119473
	TS *int64 `json:"TS,omitempty" xml:"TS,omitempty"`
	// The value of the metric.
	//
	// Valid values of N: 1 to 100.
	//
	// The value must be an integer or a floating-point number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 90
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutHybridMonitorMetricDataRequestMetricList) String() string {
	return dara.Prettify(s)
}

func (s PutHybridMonitorMetricDataRequestMetricList) GoString() string {
	return s.String()
}

func (s *PutHybridMonitorMetricDataRequestMetricList) GetLabels() []*PutHybridMonitorMetricDataRequestMetricListLabels {
	return s.Labels
}

func (s *PutHybridMonitorMetricDataRequestMetricList) GetName() *string {
	return s.Name
}

func (s *PutHybridMonitorMetricDataRequestMetricList) GetTS() *int64 {
	return s.TS
}

func (s *PutHybridMonitorMetricDataRequestMetricList) GetValue() *string {
	return s.Value
}

func (s *PutHybridMonitorMetricDataRequestMetricList) SetLabels(v []*PutHybridMonitorMetricDataRequestMetricListLabels) *PutHybridMonitorMetricDataRequestMetricList {
	s.Labels = v
	return s
}

func (s *PutHybridMonitorMetricDataRequestMetricList) SetName(v string) *PutHybridMonitorMetricDataRequestMetricList {
	s.Name = &v
	return s
}

func (s *PutHybridMonitorMetricDataRequestMetricList) SetTS(v int64) *PutHybridMonitorMetricDataRequestMetricList {
	s.TS = &v
	return s
}

func (s *PutHybridMonitorMetricDataRequestMetricList) SetValue(v string) *PutHybridMonitorMetricDataRequestMetricList {
	s.Value = &v
	return s
}

func (s *PutHybridMonitorMetricDataRequestMetricList) Validate() error {
	return dara.Validate(s)
}

type PutHybridMonitorMetricDataRequestMetricListLabels struct {
	// The tag key of the metric.
	//
	// Valid values of N: 1 to 100.
	//
	// The key can contain letters, digits, and underscores (_). The key must start with a letter or an underscore (_).
	//
	// >  You must specify both the Key and Value parameters.
	//
	// example:
	//
	// IP
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the metric.
	//
	// Valid values of N: 1 to 100.
	//
	// >  You must specify both the Key and Value parameters.
	//
	// example:
	//
	// 192.168.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutHybridMonitorMetricDataRequestMetricListLabels) String() string {
	return dara.Prettify(s)
}

func (s PutHybridMonitorMetricDataRequestMetricListLabels) GoString() string {
	return s.String()
}

func (s *PutHybridMonitorMetricDataRequestMetricListLabels) GetKey() *string {
	return s.Key
}

func (s *PutHybridMonitorMetricDataRequestMetricListLabels) GetValue() *string {
	return s.Value
}

func (s *PutHybridMonitorMetricDataRequestMetricListLabels) SetKey(v string) *PutHybridMonitorMetricDataRequestMetricListLabels {
	s.Key = &v
	return s
}

func (s *PutHybridMonitorMetricDataRequestMetricListLabels) SetValue(v string) *PutHybridMonitorMetricDataRequestMetricListLabels {
	s.Value = &v
	return s
}

func (s *PutHybridMonitorMetricDataRequestMetricListLabels) Validate() error {
	return dara.Validate(s)
}
