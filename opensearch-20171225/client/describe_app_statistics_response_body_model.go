// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppStatisticsResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DescribeAppStatisticsResponseBody
	GetResult() map[string]interface{}
}

type DescribeAppStatisticsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 76FC45F1-4167-D3CD-6737-4F97BA503FA0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The statistics.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DescribeAppStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppStatisticsResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DescribeAppStatisticsResponseBody) SetRequestId(v string) *DescribeAppStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppStatisticsResponseBody) SetResult(v map[string]interface{}) *DescribeAppStatisticsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeAppStatisticsResponseBody) Validate() error {
	return dara.Validate(s)
}
