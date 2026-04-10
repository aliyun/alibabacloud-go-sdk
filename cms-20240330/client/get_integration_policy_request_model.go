// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegrationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetIntegrationPolicyRequest struct {
}

func (s GetIntegrationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetIntegrationPolicyRequest) Validate() error {
	return dara.Validate(s)
}
