// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpecVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAgentSpecVersionRequest struct {
}

func (s GetAgentSpecVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecVersionRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSpecVersionRequest) Validate() error {
	return dara.Validate(s)
}
