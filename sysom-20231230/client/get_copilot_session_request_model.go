// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopilotSessionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetCopilotSessionRequest struct {
}

func (s GetCopilotSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotSessionRequest) GoString() string {
	return s.String()
}

func (s *GetCopilotSessionRequest) Validate() error {
	return dara.Validate(s)
}
