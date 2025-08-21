// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticsearchHealthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeElasticsearchHealthResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeElasticsearchHealthResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeElasticsearchHealthResponseBody
	GetRequestId() *string
	SetResult(v string) *DescribeElasticsearchHealthResponseBody
	GetResult() *string
}

type DescribeElasticsearchHealthResponseBody struct {
	// The response code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0731F217-2C8A-4D42-8BCD-5C352866E3B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The color that indicates the health status of the cluster.
	//
	// example:
	//
	// GREEN
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DescribeElasticsearchHealthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticsearchHealthResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticsearchHealthResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeElasticsearchHealthResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeElasticsearchHealthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticsearchHealthResponseBody) GetResult() *string {
	return s.Result
}

func (s *DescribeElasticsearchHealthResponseBody) SetCode(v string) *DescribeElasticsearchHealthResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeElasticsearchHealthResponseBody) SetMessage(v string) *DescribeElasticsearchHealthResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeElasticsearchHealthResponseBody) SetRequestId(v string) *DescribeElasticsearchHealthResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticsearchHealthResponseBody) SetResult(v string) *DescribeElasticsearchHealthResponseBody {
	s.Result = &v
	return s
}

func (s *DescribeElasticsearchHealthResponseBody) Validate() error {
	return dara.Validate(s)
}
