// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationPolicyAddonsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListIntegrationPolicyAddonsRequest struct {
}

func (s ListIntegrationPolicyAddonsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationPolicyAddonsRequest) GoString() string {
	return s.String()
}

func (s *ListIntegrationPolicyAddonsRequest) Validate() error {
	return dara.Validate(s)
}
