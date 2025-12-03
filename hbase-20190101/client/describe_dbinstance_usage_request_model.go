// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeDBInstanceUsageRequest
	GetClusterId() *string
}

type DescribeDBInstanceUsageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp1u0639js2h7****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeDBInstanceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceUsageRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeDBInstanceUsageRequest) SetClusterId(v string) *DescribeDBInstanceUsageRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeDBInstanceUsageRequest) Validate() error {
	return dara.Validate(s)
}
