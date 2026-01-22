// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeJobResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExitCode(v int64) *DescribeJobResultsResponseBody
	GetExitCode() *int64
	SetOutput(v string) *DescribeJobResultsResponseBody
	GetOutput() *string
	SetRequestId(v string) *DescribeJobResultsResponseBody
	GetRequestId() *string
}

type DescribeJobResultsResponseBody struct {
	// example:
	//
	// 0
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// example:
	//
	// MTU6MzA6MDEK
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BABC742E-04D7-5BA5-8A5F-7D9461D37B19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeJobResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeJobResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobResultsResponseBody) GetExitCode() *int64 {
	return s.ExitCode
}

func (s *DescribeJobResultsResponseBody) GetOutput() *string {
	return s.Output
}

func (s *DescribeJobResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeJobResultsResponseBody) SetExitCode(v int64) *DescribeJobResultsResponseBody {
	s.ExitCode = &v
	return s
}

func (s *DescribeJobResultsResponseBody) SetOutput(v string) *DescribeJobResultsResponseBody {
	s.Output = &v
	return s
}

func (s *DescribeJobResultsResponseBody) SetRequestId(v string) *DescribeJobResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobResultsResponseBody) Validate() error {
	return dara.Validate(s)
}
