// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiGroupVpcWhitelistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeApiGroupVpcWhitelistResponseBody
	GetRequestId() *string
	SetVpcIds(v string) *DescribeApiGroupVpcWhitelistResponseBody
	GetVpcIds() *string
}

type DescribeApiGroupVpcWhitelistResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BZ016
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp11w979o2s9rcr962w25
	VpcIds *string `json:"VpcIds,omitempty" xml:"VpcIds,omitempty"`
}

func (s DescribeApiGroupVpcWhitelistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiGroupVpcWhitelistResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiGroupVpcWhitelistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiGroupVpcWhitelistResponseBody) GetVpcIds() *string {
	return s.VpcIds
}

func (s *DescribeApiGroupVpcWhitelistResponseBody) SetRequestId(v string) *DescribeApiGroupVpcWhitelistResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiGroupVpcWhitelistResponseBody) SetVpcIds(v string) *DescribeApiGroupVpcWhitelistResponseBody {
	s.VpcIds = &v
	return s
}

func (s *DescribeApiGroupVpcWhitelistResponseBody) Validate() error {
	return dara.Validate(s)
}
