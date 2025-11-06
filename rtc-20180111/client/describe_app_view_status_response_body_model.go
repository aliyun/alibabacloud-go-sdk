// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppViewStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeAppViewStatusResponseBody
	GetRequestId() *string
	SetResult(v string) *DescribeAppViewStatusResponseBody
	GetResult() *string
}

type DescribeAppViewStatusResponseBody struct {
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

func (s DescribeAppViewStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppViewStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppViewStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppViewStatusResponseBody) GetResult() *string {
	return s.Result
}

func (s *DescribeAppViewStatusResponseBody) SetRequestId(v string) *DescribeAppViewStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppViewStatusResponseBody) SetResult(v string) *DescribeAppViewStatusResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeAppViewStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
