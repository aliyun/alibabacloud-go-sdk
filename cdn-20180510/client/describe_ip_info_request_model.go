// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIP(v string) *DescribeIpInfoRequest
	GetIP() *string
}

type DescribeIpInfoRequest struct {
	// The IP address. You can specify only one IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
}

func (s DescribeIpInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpInfoRequest) GetIP() *string {
	return s.IP
}

func (s *DescribeIpInfoRequest) SetIP(v string) *DescribeIpInfoRequest {
	s.IP = &v
	return s
}

func (s *DescribeIpInfoRequest) Validate() error {
	return dara.Validate(s)
}
