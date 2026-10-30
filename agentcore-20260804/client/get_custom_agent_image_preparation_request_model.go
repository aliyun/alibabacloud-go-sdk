// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomAgentImagePreparationRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetCustomAgentImagePreparationRequest struct {
}

func (s GetCustomAgentImagePreparationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentImagePreparationRequest) GoString() string {
	return s.String()
}

func (s *GetCustomAgentImagePreparationRequest) Validate() error {
	return dara.Validate(s)
}
