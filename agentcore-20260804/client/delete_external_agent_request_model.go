// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExternalAgentRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteExternalAgentRequest struct {
}

func (s DeleteExternalAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExternalAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteExternalAgentRequest) Validate() error {
	return dara.Validate(s)
}
