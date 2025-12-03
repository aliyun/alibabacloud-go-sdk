// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeInstanceRequest
	GetClusterId() *string
}

type DescribeInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp150tns0sjxs****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeInstanceRequest) SetClusterId(v string) *DescribeInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstanceRequest) Validate() error {
	return dara.Validate(s)
}
