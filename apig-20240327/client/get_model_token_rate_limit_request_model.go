// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelTokenRateLimitRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetModelTokenRateLimitRequest struct {
}

func (s GetModelTokenRateLimitRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelTokenRateLimitRequest) GoString() string {
	return s.String()
}

func (s *GetModelTokenRateLimitRequest) Validate() error {
	return dara.Validate(s)
}
