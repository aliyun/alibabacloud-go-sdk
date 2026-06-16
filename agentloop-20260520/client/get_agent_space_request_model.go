// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAgentSpaceRequest struct {
}

func (s GetAgentSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpaceRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSpaceRequest) Validate() error {
	return dara.Validate(s)
}
