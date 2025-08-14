// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeGroupTrendResponseBody
	GetRequestId() *string
	SetData(v bool) *DescribeGroupTrendResponseBody
	GetData() *bool
}

type DescribeGroupTrendResponseBody struct {
	// Request ID
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	//
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
}

func (s DescribeGroupTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupTrendResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGroupTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGroupTrendResponseBody) GetData() *bool {
	return s.Data
}

func (s *DescribeGroupTrendResponseBody) SetRequestId(v string) *DescribeGroupTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGroupTrendResponseBody) SetData(v bool) *DescribeGroupTrendResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeGroupTrendResponseBody) Validate() error {
	return dara.Validate(s)
}
