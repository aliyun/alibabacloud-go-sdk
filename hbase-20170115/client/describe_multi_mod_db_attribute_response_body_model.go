// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiModDbAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DescribeMultiModDbAttributeResponseBody
	GetData() *string
	SetRequestId(v string) *DescribeMultiModDbAttributeResponseBody
	GetRequestId() *string
}

type DescribeMultiModDbAttributeResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMultiModDbAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiModDbAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMultiModDbAttributeResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeMultiModDbAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMultiModDbAttributeResponseBody) SetData(v string) *DescribeMultiModDbAttributeResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeMultiModDbAttributeResponseBody) SetRequestId(v string) *DescribeMultiModDbAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMultiModDbAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
