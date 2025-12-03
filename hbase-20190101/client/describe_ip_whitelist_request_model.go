// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeIpWhitelistRequest
	GetClusterId() *string
}

type DescribeIpWhitelistRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeIpWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpWhitelistRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeIpWhitelistRequest) SetClusterId(v string) *DescribeIpWhitelistRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeIpWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
