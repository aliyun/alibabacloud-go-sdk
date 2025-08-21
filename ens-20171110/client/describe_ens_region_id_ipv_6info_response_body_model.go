// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnsRegionIdIpv6InfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeEnsRegionIdIpv6InfoResponseBody
	GetRequestId() *string
	SetSupportIpv6Info(v *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) *DescribeEnsRegionIdIpv6InfoResponseBody
	GetSupportIpv6Info() *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info
}

type DescribeEnsRegionIdIpv6InfoResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3703C4AC-9396-458C-8F25-1D701334D309
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// IPv6 support information.
	SupportIpv6Info *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info `json:"SupportIpv6Info,omitempty" xml:"SupportIpv6Info,omitempty" type:"Struct"`
}

func (s DescribeEnsRegionIdIpv6InfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionIdIpv6InfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBody) GetSupportIpv6Info() *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info {
	return s.SupportIpv6Info
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBody) SetRequestId(v string) *DescribeEnsRegionIdIpv6InfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBody) SetSupportIpv6Info(v *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) *DescribeEnsRegionIdIpv6InfoResponseBody {
	s.SupportIpv6Info = v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info struct {
	// The ID of the node.
	//
	// example:
	//
	// cn-chengdu-xxxx-4
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// Specifies whether IPv6 is supported. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SupportIpv6 *bool `json:"SupportIpv6,omitempty" xml:"SupportIpv6,omitempty"`
}

func (s DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) GoString() string {
	return s.String()
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) GetSupportIpv6() *bool {
	return s.SupportIpv6
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) SetEnsRegionId(v string) *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) SetSupportIpv6(v bool) *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info {
	s.SupportIpv6 = &v
	return s
}

func (s *DescribeEnsRegionIdIpv6InfoResponseBodySupportIpv6Info) Validate() error {
	return dara.Validate(s)
}
