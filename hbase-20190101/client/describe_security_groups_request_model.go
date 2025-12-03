// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeSecurityGroupsRequest
	GetClusterId() *string
}

type DescribeSecurityGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp161ax8i03c4uq**
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeSecurityGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeSecurityGroupsRequest) SetClusterId(v string) *DescribeSecurityGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) Validate() error {
	return dara.Validate(s)
}
