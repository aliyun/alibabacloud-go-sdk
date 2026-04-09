// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSandboxRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StopSandboxRequest struct {
}

func (s StopSandboxRequest) String() string {
	return dara.Prettify(s)
}

func (s StopSandboxRequest) GoString() string {
	return s.String()
}

func (s *StopSandboxRequest) Validate() error {
	return dara.Validate(s)
}
