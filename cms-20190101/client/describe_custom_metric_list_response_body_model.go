// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomMetricListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCustomMetricListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeCustomMetricListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCustomMetricListResponseBody
	GetRequestId() *string
	SetResult(v string) *DescribeCustomMetricListResponseBody
	GetResult() *string
}

type DescribeCustomMetricListResponseBody struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AF425E4-1DEA-54F2-910A-8117C9686140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The reported custom metrics that are found in the query.
	//
	// example:
	//
	// {\\"all\\":0,\\"size\\":10,\\"param\\":{\\"metric\\":\\"{\\\\\\"metricName\\\\\\":\\\\\\"cpu_total\\\\\\",\\\\\\"groupId\\\\\\":7378****,\\\\\\"project\\\\\\":\\\\\\"acs_customMetric_120886317861****\\\\\\",\\\\\\"dimension\\\\\\":\\\\\\"sampleName1=value1&sampleName2=value2\\\\\\",\\\\\\"status\\\\\\":1}\\",\\"service\\":\\"metric-center.aliyun-inc.com\\"},\\"page\\":1,\\"list\\":[]}
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeCustomMetricListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomMetricListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomMetricListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCustomMetricListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCustomMetricListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomMetricListResponseBody) GetResult() *string {
	return s.Result
}

func (s *DescribeCustomMetricListResponseBody) SetCode(v string) *DescribeCustomMetricListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomMetricListResponseBody) SetMessage(v string) *DescribeCustomMetricListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomMetricListResponseBody) SetRequestId(v string) *DescribeCustomMetricListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomMetricListResponseBody) SetResult(v string) *DescribeCustomMetricListResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeCustomMetricListResponseBody) Validate() error {
	return dara.Validate(s)
}
