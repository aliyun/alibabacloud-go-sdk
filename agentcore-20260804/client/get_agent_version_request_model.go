// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAgentVersionRequest struct {
}

func (s GetAgentVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentVersionRequest) GoString() string {
	return s.String()
}

func (s *GetAgentVersionRequest) Validate() error {
	return dara.Validate(s)
}
