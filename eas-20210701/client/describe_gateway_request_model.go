// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeGatewayRequest struct {
}

func (s DescribeGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayRequest) GoString() string {
	return s.String()
}

func (s *DescribeGatewayRequest) Validate() error {
	return dara.Validate(s)
}
