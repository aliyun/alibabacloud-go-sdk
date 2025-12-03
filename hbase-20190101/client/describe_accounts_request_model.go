// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeAccountsRequest
	GetClusterId() *string
}

type DescribeAccountsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1uoihlf82e8****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeAccountsRequest) SetClusterId(v string) *DescribeAccountsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeAccountsRequest) Validate() error {
	return dara.Validate(s)
}
