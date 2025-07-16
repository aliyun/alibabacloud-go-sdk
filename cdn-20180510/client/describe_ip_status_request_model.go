// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIps(v string) *DescribeIpStatusRequest
	GetIps() *string
}

type DescribeIpStatusRequest struct {
	// The IP addresses that you want to query. Separate IP addresses with underscores (_), such as Ips=ip1_ip2.
	//
	// This parameter is required.
	//
	// example:
	//
	// ip1_ip2
	Ips *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
}

func (s DescribeIpStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpStatusRequest) GetIps() *string {
	return s.Ips
}

func (s *DescribeIpStatusRequest) SetIps(v string) *DescribeIpStatusRequest {
	s.Ips = &v
	return s
}

func (s *DescribeIpStatusRequest) Validate() error {
	return dara.Validate(s)
}
