// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiKeyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DescribeApiKeyRequest struct {
}

func (s DescribeApiKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiKeyRequest) Validate() error {
	return dara.Validate(s)
}
