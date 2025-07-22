// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppAgentFunctionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppAgentFunctionStatusResponseBody
	GetRequestId() *string
	SetResult(v string) *DescribeAppAgentFunctionStatusResponseBody
	GetResult() *string
}

type DescribeAppAgentFunctionStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// disable
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeAppAgentFunctionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentFunctionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentFunctionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppAgentFunctionStatusResponseBody) GetResult() *string {
	return s.Result
}

func (s *DescribeAppAgentFunctionStatusResponseBody) SetRequestId(v string) *DescribeAppAgentFunctionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppAgentFunctionStatusResponseBody) SetResult(v string) *DescribeAppAgentFunctionStatusResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeAppAgentFunctionStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
