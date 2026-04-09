// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSandboxRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetSandboxRequest struct {
}

func (s GetSandboxRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSandboxRequest) GoString() string {
	return s.String()
}

func (s *GetSandboxRequest) Validate() error {
	return dara.Validate(s)
}
