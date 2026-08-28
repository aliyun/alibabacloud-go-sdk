// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExternalAgentRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetExternalAgentRequest struct {
}

func (s GetExternalAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentRequest) GoString() string {
	return s.String()
}

func (s *GetExternalAgentRequest) Validate() error {
	return dara.Validate(s)
}
