// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSandboxRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteSandboxRequest struct {
}

func (s DeleteSandboxRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSandboxRequest) GoString() string {
	return s.String()
}

func (s *DeleteSandboxRequest) Validate() error {
	return dara.Validate(s)
}
