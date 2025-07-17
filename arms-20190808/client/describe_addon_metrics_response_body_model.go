// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAddonMetricsResponseBody
	GetCode() *int32
	SetData(v []*DescribeAddonMetricsResponseBodyData) *DescribeAddonMetricsResponseBody
	GetData() []*DescribeAddonMetricsResponseBodyData
	SetMessage(v string) *DescribeAddonMetricsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAddonMetricsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAddonMetricsResponseBody
	GetSuccess() *bool
}

type DescribeAddonMetricsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The metric details.
	Data []*DescribeAddonMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B6A00968-82A8-4F14-9D1B-B53827DB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAddonMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAddonMetricsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAddonMetricsResponseBody) GetData() []*DescribeAddonMetricsResponseBodyData {
	return s.Data
}

func (s *DescribeAddonMetricsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAddonMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAddonMetricsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAddonMetricsResponseBody) SetCode(v int32) *DescribeAddonMetricsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAddonMetricsResponseBody) SetData(v []*DescribeAddonMetricsResponseBodyData) *DescribeAddonMetricsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAddonMetricsResponseBody) SetMessage(v string) *DescribeAddonMetricsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAddonMetricsResponseBody) SetRequestId(v string) *DescribeAddonMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAddonMetricsResponseBody) SetSuccess(v bool) *DescribeAddonMetricsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAddonMetricsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAddonMetricsResponseBodyData struct {
	// The metric group.
	//
	// example:
	//
	// Common
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The tags.
	Labels []*DescribeAddonMetricsResponseBodyDataLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The metrics.
	Metrics []*DescribeAddonMetricsResponseBodyDataMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
}

func (s DescribeAddonMetricsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAddonMetricsResponseBodyData) GetGroup() *string {
	return s.Group
}

func (s *DescribeAddonMetricsResponseBodyData) GetLabels() []*DescribeAddonMetricsResponseBodyDataLabels {
	return s.Labels
}

func (s *DescribeAddonMetricsResponseBodyData) GetMetrics() []*DescribeAddonMetricsResponseBodyDataMetrics {
	return s.Metrics
}

func (s *DescribeAddonMetricsResponseBodyData) SetGroup(v string) *DescribeAddonMetricsResponseBodyData {
	s.Group = &v
	return s
}

func (s *DescribeAddonMetricsResponseBodyData) SetLabels(v []*DescribeAddonMetricsResponseBodyDataLabels) *DescribeAddonMetricsResponseBodyData {
	s.Labels = v
	return s
}

func (s *DescribeAddonMetricsResponseBodyData) SetMetrics(v []*DescribeAddonMetricsResponseBodyDataMetrics) *DescribeAddonMetricsResponseBodyData {
	s.Metrics = v
	return s
}

func (s *DescribeAddonMetricsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeAddonMetricsResponseBodyDataLabels struct {
	// The description of the tag.
	//
	// example:
	//
	// The number of times a B-tree page of size PAGE_SIZE was successfully compressed.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tag key.
	//
	// example:
	//
	// page_size
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The source of the tag.
	//
	// example:
	//
	// db
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeAddonMetricsResponseBodyDataLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonMetricsResponseBodyDataLabels) GoString() string {
	return s.String()
}

func (s *DescribeAddonMetricsResponseBodyDataLabels) GetDescription() *string {
	return s.Description
}

func (s *DescribeAddonMetricsResponseBodyDataLabels) GetKey() *string {
	return s.Key
}

func (s *DescribeAddonMetricsResponseBodyDataLabels) GetSource() *string {
	return s.Source
}

func (s *DescribeAddonMetricsResponseBodyDataLabels) SetDescription(v string) *DescribeAddonMetricsResponseBodyDataLabels {
	s.Description = &v
	return s
}

func (s *DescribeAddonMetricsResponseBodyDataLabels) SetKey(v string) *DescribeAddonMetricsResponseBodyDataLabels {
	s.Key = &v
	return s
}

func (s *DescribeAddonMetricsResponseBodyDataLabels) SetSource(v string) *DescribeAddonMetricsResponseBodyDataLabels {
	s.Source = &v
	return s
}

func (s *DescribeAddonMetricsResponseBodyDataLabels) Validate() error {
	return dara.Validate(s)
}

type DescribeAddonMetricsResponseBodyDataMetrics struct {
	// The description of the metric.
	//
	// example:
	//
	// The number of times a B-tree page of size PAGE_SIZE was successfully compressed.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tags.
	Labels []*DescribeAddonMetricsResponseBodyDataMetricsLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The metric name.
	//
	// example:
	//
	// mysql_exporter_collector_duration_seconds
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The type of the metric.
	//
	// example:
	//
	// GAUGE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The unit of the metric.
	//
	// example:
	//
	// bytes
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeAddonMetricsResponseBodyDataMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonMetricsResponseBodyDataMetrics) GoString() string {
	return s.String()
}

func (s *DescribeAddonMetricsResponseBodyDataMetrics) GetDescription() *string {
	return s.Description
}

func (s *DescribeAddonMetricsResponseBodyDataMetrics) GetLabels() []*DescribeAddonMetricsResponseBodyDataMetricsLabels {
	return s.Labels
}

func (s *DescribeAddonMetricsResponseBodyDataMetrics) GetMetric() *string {
	return s.Metric
}

func (s *DescribeAddonMetricsResponseBodyDataMetrics) GetType() *string {
	return s.Type
}

func (s *DescribeAddonMetricsResponseBodyDataMetrics) GetUnit() *string {
	return s.Unit
}

func (s *DescribeAddonMetricsResponseBodyDataMetrics) SetDescription(v string) *DescribeAddonMetricsResponseBodyDataMetrics {
	s.Description = &v
	return s
}

func (s *DescribeAddonMetricsResponseBodyDataMetrics) SetLabels(v []*DescribeAddonMetricsResponseBodyDataMetricsLabels) *DescribeAddonMetricsResponseBodyDataMetrics {
	s.Labels = v
	return s
}

func (s *DescribeAddonMetricsResponseBodyDataMetrics) SetMetric(v string) *DescribeAddonMetricsResponseBodyDataMetrics {
	s.Metric = &v
	return s
}

func (s *DescribeAddonMetricsResponseBodyDataMetrics) SetType(v string) *DescribeAddonMetricsResponseBodyDataMetrics {
	s.Type = &v
	return s
}

func (s *DescribeAddonMetricsResponseBodyDataMetrics) SetUnit(v string) *DescribeAddonMetricsResponseBodyDataMetrics {
	s.Unit = &v
	return s
}

func (s *DescribeAddonMetricsResponseBodyDataMetrics) Validate() error {
	return dara.Validate(s)
}

type DescribeAddonMetricsResponseBodyDataMetricsLabels struct {
	// The description of the tag.
	//
	// example:
	//
	// PAGE_SIZE
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The tag key.
	//
	// example:
	//
	// page_size
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The source of the tag.
	//
	// example:
	//
	// db
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeAddonMetricsResponseBodyDataMetricsLabels) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonMetricsResponseBodyDataMetricsLabels) GoString() string {
	return s.String()
}

func (s *DescribeAddonMetricsResponseBodyDataMetricsLabels) GetDescription() *string {
	return s.Description
}

func (s *DescribeAddonMetricsResponseBodyDataMetricsLabels) GetKey() *string {
	return s.Key
}

func (s *DescribeAddonMetricsResponseBodyDataMetricsLabels) GetSource() *string {
	return s.Source
}

func (s *DescribeAddonMetricsResponseBodyDataMetricsLabels) SetDescription(v string) *DescribeAddonMetricsResponseBodyDataMetricsLabels {
	s.Description = &v
	return s
}

func (s *DescribeAddonMetricsResponseBodyDataMetricsLabels) SetKey(v string) *DescribeAddonMetricsResponseBodyDataMetricsLabels {
	s.Key = &v
	return s
}

func (s *DescribeAddonMetricsResponseBodyDataMetricsLabels) SetSource(v string) *DescribeAddonMetricsResponseBodyDataMetricsLabels {
	s.Source = &v
	return s
}

func (s *DescribeAddonMetricsResponseBodyDataMetricsLabels) Validate() error {
	return dara.Validate(s)
}
