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
	// The list of monitoring data.
	//
	// Valid values of N: 1 to 100.
	//
	// This parameter is required.
	MetricList []*PutHybridMonitorMetricDataRequestMetricList `json:"MetricList,omitempty" xml:"MetricList,omitempty" type:"Repeated"`
	// The name of the metric namespace.
	//
	// For information about how to obtain the name of a metric namespace, see [DescribeHybridMonitorNamespaceList](https://help.aliyun.com/document_detail/428880.html).
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

type PutHybridMonitorMetricDataRequestMetricList struct {
	// The list of labels of the metric.
	//
	// Valid values of N: 1 to 100.
	//
	// example:
	//
	// [{"Key":"app","Value":"testApp"},{"Key":"ip","Value":"192.168.XX.XX"},{"Key":"hostName","Value":"host01"}]
	Labels []*PutHybridMonitorMetricDataRequestMetricListLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The name of the metric.
	//
	// Valid values of N: 1 to 100.
	//
	// Format: The name can contain uppercase letters, lowercase letters, digits, and underscores (_). The name must start with an uppercase letter or a lowercase letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// CPU_Usage
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The timestamp when the monitoring data was reported.
	//
	// Valid values of N: 1 to 100.
	//
	// Unit: milliseconds. Default value: the current time.
	//
	// example:
	//
	// 1640776119473
	TS *int64 `json:"TS,omitempty" xml:"TS,omitempty"`
	// The value of the metric.
	//
	// Valid values of N: 1 to 100.
	//
	// Format: an integer or a floating-point number.
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
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutHybridMonitorMetricDataRequestMetricListLabels struct {
	// The key of the label of the metric.
	//
	// Valid values of N: 1 to 100.
	//
	// Format: The key can contain uppercase letters, lowercase letters, digits, and underscores (_). The key must start with an uppercase letter, a lowercase letter, or an underscore (_).
	//
	// > Key and Value must be set at the same time.
	//
	// example:
	//
	// IP
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the label of the metric.
	//
	// Valid values of N: 1 to 100.
	//
	// > Key and Value must be set at the same time.
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
