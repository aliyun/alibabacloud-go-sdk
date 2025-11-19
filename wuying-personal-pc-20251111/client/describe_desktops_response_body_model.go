// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDesktopsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeDesktopsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *DescribeDesktopsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDesktopsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDesktopsResponseBody
	GetRequestId() *string
	SetResult(v interface{}) *DescribeDesktopsResponseBody
	GetResult() interface{}
	SetTraceId(v string) *DescribeDesktopsResponseBody
	GetTraceId() *string
}

type DescribeDesktopsResponseBody struct {
	Code           *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults     *int32      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken      *string     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result         interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
	TraceId        *string     `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeDesktopsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDesktopsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeDesktopsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDesktopsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDesktopsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDesktopsResponseBody) GetResult() interface{} {
	return s.Result
}

func (s *DescribeDesktopsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeDesktopsResponseBody) SetCode(v string) *DescribeDesktopsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetHttpStatusCode(v int32) *DescribeDesktopsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetMaxResults(v int32) *DescribeDesktopsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetNextToken(v string) *DescribeDesktopsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetRequestId(v string) *DescribeDesktopsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDesktopsResponseBody) SetResult(v interface{}) *DescribeDesktopsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeDesktopsResponseBody) SetTraceId(v string) *DescribeDesktopsResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeDesktopsResponseBody) Validate() error {
	return dara.Validate(s)
}
