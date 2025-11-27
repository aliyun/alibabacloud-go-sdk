// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCMetricListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeRCMetricListResponseBody
	GetCode() *string
	SetDatapoints(v string) *DescribeRCMetricListResponseBody
	GetDatapoints() *string
	SetMessage(v string) *DescribeRCMetricListResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeRCMetricListResponseBody
	GetNextToken() *string
	SetPeriod(v string) *DescribeRCMetricListResponseBody
	GetPeriod() *string
	SetRequestId(v string) *DescribeRCMetricListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRCMetricListResponseBody
	GetSuccess() *bool
}

type DescribeRCMetricListResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The monitoring data.
	//
	// example:
	//
	// [{\\"timestamp\\":1722909960000,\\"instanceId\\":\\"rc-dh2jf9n6j4s14926****\\",\\"userId\\":\\"1695619988087373\\",\\"Minimum\\":0.097,\\"Maximum\\":0.097,\\"Average\\":0.097},{\\"timestamp\\":1722910020000,\\"instanceId\\":\\"rc-dh2jf9n6j4s14926****\\",\\"userId\\":\\"1695619988087373\\",\\"Minimum\\":0.093,\\"Maximum\\":0.093,\\"Average\\":0.093}]
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// The message that is returned for the request.
	//
	// >  If the request is successful, **Successful*	- is returned. If the request fails, an error message that contains information such as an error code is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// 6178f1825f9fb76ce0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The statistical period of the monitoring data.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EA2D4F34-01A7-46EB-A339-D80882135206
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRCMetricListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCMetricListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRCMetricListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeRCMetricListResponseBody) GetDatapoints() *string {
	return s.Datapoints
}

func (s *DescribeRCMetricListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRCMetricListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRCMetricListResponseBody) GetPeriod() *string {
	return s.Period
}

func (s *DescribeRCMetricListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRCMetricListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRCMetricListResponseBody) SetCode(v string) *DescribeRCMetricListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRCMetricListResponseBody) SetDatapoints(v string) *DescribeRCMetricListResponseBody {
	s.Datapoints = &v
	return s
}

func (s *DescribeRCMetricListResponseBody) SetMessage(v string) *DescribeRCMetricListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRCMetricListResponseBody) SetNextToken(v string) *DescribeRCMetricListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeRCMetricListResponseBody) SetPeriod(v string) *DescribeRCMetricListResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeRCMetricListResponseBody) SetRequestId(v string) *DescribeRCMetricListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRCMetricListResponseBody) SetSuccess(v bool) *DescribeRCMetricListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRCMetricListResponseBody) Validate() error {
	return dara.Validate(s)
}
