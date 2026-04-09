// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseSandboxRequest interface {
	dara.Model
	String() string
	GoString() string
}

type PauseSandboxRequest struct {
}

func (s PauseSandboxRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseSandboxRequest) GoString() string {
	return s.String()
}

func (s *PauseSandboxRequest) Validate() error {
	return dara.Validate(s)
}
