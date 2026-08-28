// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpecLatestRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAgentSpecLatestRequest struct {
}

func (s GetAgentSpecLatestRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecLatestRequest) GoString() string {
	return s.String()
}

func (s *GetAgentSpecLatestRequest) Validate() error {
	return dara.Validate(s)
}
