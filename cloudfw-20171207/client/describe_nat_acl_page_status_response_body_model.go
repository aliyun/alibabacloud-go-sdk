// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatAclPageStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetail(v string) *DescribeNatAclPageStatusResponseBody
	GetDetail() *string
	SetNatAclPageEnable(v bool) *DescribeNatAclPageStatusResponseBody
	GetNatAclPageEnable() *bool
	SetRequestId(v string) *DescribeNatAclPageStatusResponseBody
	GetRequestId() *string
}

type DescribeNatAclPageStatusResponseBody struct {
	// Extra error information.
	//
	// example:
	//
	// proxy_not_exist
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// Indicates whether pagination for access control policies for NAT firewalls is supported.
	//
	// example:
	//
	// True
	NatAclPageEnable *bool `json:"NatAclPageEnable,omitempty" xml:"NatAclPageEnable,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B97F9AD7-A2DB-5F8F-9206-DF89DE0AC9E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNatAclPageStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatAclPageStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatAclPageStatusResponseBody) GetDetail() *string {
	return s.Detail
}

func (s *DescribeNatAclPageStatusResponseBody) GetNatAclPageEnable() *bool {
	return s.NatAclPageEnable
}

func (s *DescribeNatAclPageStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatAclPageStatusResponseBody) SetDetail(v string) *DescribeNatAclPageStatusResponseBody {
	s.Detail = &v
	return s
}

func (s *DescribeNatAclPageStatusResponseBody) SetNatAclPageEnable(v bool) *DescribeNatAclPageStatusResponseBody {
	s.NatAclPageEnable = &v
	return s
}

func (s *DescribeNatAclPageStatusResponseBody) SetRequestId(v string) *DescribeNatAclPageStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatAclPageStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
