// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpLocationServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInternetIp(v string) *DescribeIpLocationServiceRequest
	GetInternetIp() *string
}

type DescribeIpLocationServiceRequest struct {
	// The IP address of the asset to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 121.199.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
}

func (s DescribeIpLocationServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpLocationServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpLocationServiceRequest) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeIpLocationServiceRequest) SetInternetIp(v string) *DescribeIpLocationServiceRequest {
	s.InternetIp = &v
	return s
}

func (s *DescribeIpLocationServiceRequest) Validate() error {
	return dara.Validate(s)
}
