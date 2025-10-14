// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExporterRuleListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeExporterRuleListResponseBody
	GetCode() *string
	SetDatapoints(v *DescribeExporterRuleListResponseBodyDatapoints) *DescribeExporterRuleListResponseBody
	GetDatapoints() *DescribeExporterRuleListResponseBodyDatapoints
	SetMessage(v string) *DescribeExporterRuleListResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeExporterRuleListResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *DescribeExporterRuleListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeExporterRuleListResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeExporterRuleListResponseBody
	GetTotal() *int32
}

type DescribeExporterRuleListResponseBody struct {
	// The HTTP status code.
	//
	// > The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data export rules.
	Datapoints *DescribeExporterRuleListResponseBodyDatapoints `json:"Datapoints,omitempty" xml:"Datapoints,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// susscess
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6BA047CA-8BC6-40BC-BC8F-FBECF35F1993
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1000
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeExporterRuleListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExporterRuleListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExporterRuleListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeExporterRuleListResponseBody) GetDatapoints() *DescribeExporterRuleListResponseBodyDatapoints {
	return s.Datapoints
}

func (s *DescribeExporterRuleListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeExporterRuleListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeExporterRuleListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExporterRuleListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeExporterRuleListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeExporterRuleListResponseBody) SetCode(v string) *DescribeExporterRuleListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeExporterRuleListResponseBody) SetDatapoints(v *DescribeExporterRuleListResponseBodyDatapoints) *DescribeExporterRuleListResponseBody {
	s.Datapoints = v
	return s
}

func (s *DescribeExporterRuleListResponseBody) SetMessage(v string) *DescribeExporterRuleListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeExporterRuleListResponseBody) SetPageNumber(v int32) *DescribeExporterRuleListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeExporterRuleListResponseBody) SetRequestId(v string) *DescribeExporterRuleListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExporterRuleListResponseBody) SetSuccess(v bool) *DescribeExporterRuleListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExporterRuleListResponseBody) SetTotal(v int32) *DescribeExporterRuleListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeExporterRuleListResponseBody) Validate() error {
	if s.Datapoints != nil {
		if err := s.Datapoints.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeExporterRuleListResponseBodyDatapoints struct {
	Datapoint []*DescribeExporterRuleListResponseBodyDatapointsDatapoint `json:"Datapoint,omitempty" xml:"Datapoint,omitempty" type:"Repeated"`
}

func (s DescribeExporterRuleListResponseBodyDatapoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeExporterRuleListResponseBodyDatapoints) GoString() string {
	return s.String()
}

func (s *DescribeExporterRuleListResponseBodyDatapoints) GetDatapoint() []*DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	return s.Datapoint
}

func (s *DescribeExporterRuleListResponseBodyDatapoints) SetDatapoint(v []*DescribeExporterRuleListResponseBodyDatapointsDatapoint) *DescribeExporterRuleListResponseBodyDatapoints {
	s.Datapoint = v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapoints) Validate() error {
	if s.Datapoint != nil {
		for _, item := range s.Datapoint {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeExporterRuleListResponseBodyDatapointsDatapoint struct {
	// The time when the rule was created. The value is a UNIX timestamp.
	//
	// example:
	//
	// 1584024616228
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// Export monitoring data
	Describe *string `json:"Describe,omitempty" xml:"Describe,omitempty"`
	// The associated dimensions.
	//
	// example:
	//
	// {"instanceId":"xxxxx"}
	Dimension *string                                                         `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	DstName   *DescribeExporterRuleListResponseBodyDatapointsDatapointDstName `json:"DstName,omitempty" xml:"DstName,omitempty" type:"Struct"`
	// Indicates whether the rule is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The name of the metric.
	//
	// > For more information, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [Appendix 1: Metrics](https://help.aliyun.com/document_detail/28619.html).
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the service.
	//
	// > For more information, see [DescribeMetricMetaList](https://help.aliyun.com/document_detail/98846.html) or [Appendix 1: Metrics](https://help.aliyun.com/document_detail/28619.html).
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the data export rule.
	//
	// example:
	//
	// myRuleName
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The time window of the exported data.\\
	//
	// Multiple windows are separated with commas (,).
	//
	// > Data in a time window of less than 60 seconds cannot be exported.
	//
	// example:
	//
	// 60,300
	TargetWindows *string `json:"TargetWindows,omitempty" xml:"TargetWindows,omitempty"`
}

func (s DescribeExporterRuleListResponseBodyDatapointsDatapoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeExporterRuleListResponseBodyDatapointsDatapoint) GoString() string {
	return s.String()
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) GetDescribe() *string {
	return s.Describe
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) GetDstName() *DescribeExporterRuleListResponseBodyDatapointsDatapointDstName {
	return s.DstName
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) GetTargetWindows() *string {
	return s.TargetWindows
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetCreateTime(v int64) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.CreateTime = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetDescribe(v string) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.Describe = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetDimension(v string) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.Dimension = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetDstName(v *DescribeExporterRuleListResponseBodyDatapointsDatapointDstName) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.DstName = v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetEnabled(v bool) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.Enabled = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetMetricName(v string) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.MetricName = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetNamespace(v string) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.Namespace = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetRuleName(v string) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.RuleName = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) SetTargetWindows(v string) *DescribeExporterRuleListResponseBodyDatapointsDatapoint {
	s.TargetWindows = &v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapoint) Validate() error {
	if s.DstName != nil {
		if err := s.DstName.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeExporterRuleListResponseBodyDatapointsDatapointDstName struct {
	DstName []*string `json:"DstName,omitempty" xml:"DstName,omitempty" type:"Repeated"`
}

func (s DescribeExporterRuleListResponseBodyDatapointsDatapointDstName) String() string {
	return dara.Prettify(s)
}

func (s DescribeExporterRuleListResponseBodyDatapointsDatapointDstName) GoString() string {
	return s.String()
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapointDstName) GetDstName() []*string {
	return s.DstName
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapointDstName) SetDstName(v []*string) *DescribeExporterRuleListResponseBodyDatapointsDatapointDstName {
	s.DstName = v
	return s
}

func (s *DescribeExporterRuleListResponseBodyDatapointsDatapointDstName) Validate() error {
	return dara.Validate(s)
}
