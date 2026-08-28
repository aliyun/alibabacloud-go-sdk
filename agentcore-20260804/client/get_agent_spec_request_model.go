// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpecRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAgentSpecRequest struct {
}

func (s GetAgentSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSpecRequest) Validate() error {
	return dara.Validate(s)
}
