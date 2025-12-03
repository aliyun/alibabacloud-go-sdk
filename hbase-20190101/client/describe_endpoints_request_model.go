// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeEndpointsRequest
	GetClusterId() *string
}

type DescribeEndpointsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeEndpointsRequest) SetClusterId(v string) *DescribeEndpointsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
