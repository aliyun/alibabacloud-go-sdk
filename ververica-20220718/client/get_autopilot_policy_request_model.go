// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutopilotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAutopilotPolicyRequest struct {
}

func (s GetAutopilotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAutopilotPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetAutopilotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
