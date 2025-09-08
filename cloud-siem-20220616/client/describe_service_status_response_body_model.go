// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DescribeServiceStatusResponseBody
	GetData() *bool
	SetRequestId(v string) *DescribeServiceStatusResponseBody
	GetRequestId() *string
}

type DescribeServiceStatusResponseBody struct {
	// Indicates whether the threat analysis feature is authorized to access the resource directory. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *DescribeServiceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceStatusResponseBody) SetData(v bool) *DescribeServiceStatusResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeServiceStatusResponseBody) SetRequestId(v string) *DescribeServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
