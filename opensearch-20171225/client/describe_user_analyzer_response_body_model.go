// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserAnalyzerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeUserAnalyzerResponseBody
	GetRequestId() *string
	SetResult(v map[string]interface{}) *DescribeUserAnalyzerResponseBody
	GetResult() map[string]interface{}
}

type DescribeUserAnalyzerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FFAEF396-10EF-53C7-1F51-518853880729
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the custom analyzer. For more information, see [UserAnalyzer](https://help.aliyun.com/document_detail/178934.html).
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DescribeUserAnalyzerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserAnalyzerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserAnalyzerResponseBody) GetResult() map[string]interface{} {
	return s.Result
}

func (s *DescribeUserAnalyzerResponseBody) SetRequestId(v string) *DescribeUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *DescribeUserAnalyzerResponseBody {
	s.Result = v
	return s
}

func (s *DescribeUserAnalyzerResponseBody) Validate() error {
	return dara.Validate(s)
}
