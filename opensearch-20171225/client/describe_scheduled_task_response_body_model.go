// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduledTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeScheduledTaskResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DescribeScheduledTaskResponseBody
	GetResult() map[string]interface{}
}

type DescribeScheduledTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 922DC0D9-31B5-45F9-47B7-37DC678D61A8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the scheduled task.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DescribeScheduledTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScheduledTaskResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DescribeScheduledTaskResponseBody) SetRequestId(v string) *DescribeScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduledTaskResponseBody) SetResult(v map[string]interface{}) *DescribeScheduledTaskResponseBody {
	s.Result = v
	return s
}

func (s *DescribeScheduledTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
