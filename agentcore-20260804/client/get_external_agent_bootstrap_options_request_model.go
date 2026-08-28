// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExternalAgentBootstrapOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetExternalAgentBootstrapOptionsRequest struct {
}

func (s GetExternalAgentBootstrapOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExternalAgentBootstrapOptionsRequest) GoString() string {
	return s.String()
}

func (s *GetExternalAgentBootstrapOptionsRequest) Validate() error {
	return dara.Validate(s)
}
