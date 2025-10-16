// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceTypeAutoEnableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeResourceTypeAutoEnableResponseBody
	GetRequestId() *string
	SetResourceTypeAutoEnable(v map[string]*bool) *DescribeResourceTypeAutoEnableResponseBody
	GetResourceTypeAutoEnable() map[string]*bool
}

type DescribeResourceTypeAutoEnableResponseBody struct {
	// example:
	//
	// 7447795A-39AB-52CB-8F92-128DF******
	RequestId              *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceTypeAutoEnable map[string]*bool `json:"ResourceTypeAutoEnable,omitempty" xml:"ResourceTypeAutoEnable,omitempty"`
}

func (s DescribeResourceTypeAutoEnableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceTypeAutoEnableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceTypeAutoEnableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceTypeAutoEnableResponseBody) GetResourceTypeAutoEnable() map[string]*bool {
	return s.ResourceTypeAutoEnable
}

func (s *DescribeResourceTypeAutoEnableResponseBody) SetRequestId(v string) *DescribeResourceTypeAutoEnableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceTypeAutoEnableResponseBody) SetResourceTypeAutoEnable(v map[string]*bool) *DescribeResourceTypeAutoEnableResponseBody {
	s.ResourceTypeAutoEnable = v
	return s
}

func (s *DescribeResourceTypeAutoEnableResponseBody) Validate() error {
	return dara.Validate(s)
}
