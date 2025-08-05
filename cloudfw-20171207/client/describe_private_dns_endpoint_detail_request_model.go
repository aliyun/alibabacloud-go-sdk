// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateDnsEndpointDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessInstanceId(v string) *DescribePrivateDnsEndpointDetailRequest
	GetAccessInstanceId() *string
	SetRegionNo(v string) *DescribePrivateDnsEndpointDetailRequest
	GetRegionNo() *string
}

type DescribePrivateDnsEndpointDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pd-12345
	AccessInstanceId *string `json:"AccessInstanceId,omitempty" xml:"AccessInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
}

func (s DescribePrivateDnsEndpointDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsEndpointDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsEndpointDetailRequest) GetAccessInstanceId() *string {
	return s.AccessInstanceId
}

func (s *DescribePrivateDnsEndpointDetailRequest) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribePrivateDnsEndpointDetailRequest) SetAccessInstanceId(v string) *DescribePrivateDnsEndpointDetailRequest {
	s.AccessInstanceId = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailRequest) SetRegionNo(v string) *DescribePrivateDnsEndpointDetailRequest {
	s.RegionNo = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailRequest) Validate() error {
	return dara.Validate(s)
}
