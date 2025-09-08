// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomBaseRuleCompileResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeCustomBaseRuleCompileResultResponseBody
	GetRequestId() *string
	SetResult(v string) *DescribeCustomBaseRuleCompileResultResponseBody
	GetResult() *string
}

type DescribeCustomBaseRuleCompileResultResponseBody struct {
	// example:
	//
	// 58FDF266-3D56-5DE8-91E0-96A26BAB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeCustomBaseRuleCompileResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomBaseRuleCompileResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomBaseRuleCompileResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomBaseRuleCompileResultResponseBody) GetResult() *string {
	return s.Result
}

func (s *DescribeCustomBaseRuleCompileResultResponseBody) SetRequestId(v string) *DescribeCustomBaseRuleCompileResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomBaseRuleCompileResultResponseBody) SetResult(v string) *DescribeCustomBaseRuleCompileResultResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeCustomBaseRuleCompileResultResponseBody) Validate() error {
	return dara.Validate(s)
}
