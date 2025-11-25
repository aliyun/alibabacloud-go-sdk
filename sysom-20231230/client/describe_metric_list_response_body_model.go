// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMetricListResponseBody
	GetCode() *string
	SetData(v []*DescribeMetricListResponseBodyData) *DescribeMetricListResponseBody
	GetData() []*DescribeMetricListResponseBodyData
	SetMessage(v string) *DescribeMetricListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMetricListResponseBody
	GetRequestId() *string
}

type DescribeMetricListResponseBody struct {
	// example:
	//
	// Success
	Code *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Data []*DescribeMetricListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 分析ID: 56dae746-ef55-4f77-8373-cb3594c41457
	//
	// 机器i-bp164ns76pzirbwv0snt分析失败, 失败原因: Not get GPU trace data for \"56dae746-ef55-4f77-8373-cb3594c41457\" \"[\"93811\"]\"!
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeMetricListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMetricListResponseBody) GetData() []*DescribeMetricListResponseBodyData {
	return s.Data
}

func (s *DescribeMetricListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMetricListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricListResponseBody) SetCode(v string) *DescribeMetricListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetData(v []*DescribeMetricListResponseBodyData) *DescribeMetricListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeMetricListResponseBody) SetMessage(v string) *DescribeMetricListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetRequestId(v string) *DescribeMetricListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricListResponseBodyData struct {
	// example:
	//
	// {\\"taskExecName\\": \\"build-and-deploy\\", \\"pipelineName\\": \\"pipeline-run-1722909642357\\"}
	Labels *string `json:"labels,omitempty" xml:"labels,omitempty"`
	// example:
	//
	// sysom_cpu_usage_idle
	MetricName *string     `json:"metricName,omitempty" xml:"metricName,omitempty"`
	Values     [][]*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s DescribeMetricListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeMetricListResponseBodyData) GetLabels() *string {
	return s.Labels
}

func (s *DescribeMetricListResponseBodyData) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricListResponseBodyData) GetValues() [][]*string {
	return s.Values
}

func (s *DescribeMetricListResponseBodyData) SetLabels(v string) *DescribeMetricListResponseBodyData {
	s.Labels = &v
	return s
}

func (s *DescribeMetricListResponseBodyData) SetMetricName(v string) *DescribeMetricListResponseBodyData {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricListResponseBodyData) SetValues(v [][]*string) *DescribeMetricListResponseBodyData {
	s.Values = v
	return s
}

func (s *DescribeMetricListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
