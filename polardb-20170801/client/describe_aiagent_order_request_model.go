// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIAgentOrderRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeAIAgentOrderRequest struct {
}

func (s DescribeAIAgentOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIAgentOrderRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIAgentOrderRequest) Validate() error {
	return dara.Validate(s)
}
