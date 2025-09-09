// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdsCommodityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DescribeRdsCommodityResponseBody
	GetData() *string
	SetRequestId(v string) *DescribeRdsCommodityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRdsCommodityResponseBody
	GetSuccess() *bool
}

type DescribeRdsCommodityResponseBody struct {
	// Indicates the returned result.
	//
	// example:
	//
	// test
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Indicates the ID of the request.
	//
	// example:
	//
	// DC3ABA3E-0F8A-4596-9104-F5155C34315B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRdsCommodityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdsCommodityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdsCommodityResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeRdsCommodityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRdsCommodityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRdsCommodityResponseBody) SetData(v string) *DescribeRdsCommodityResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeRdsCommodityResponseBody) SetRequestId(v string) *DescribeRdsCommodityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsCommodityResponseBody) SetSuccess(v bool) *DescribeRdsCommodityResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRdsCommodityResponseBody) Validate() error {
	return dara.Validate(s)
}
