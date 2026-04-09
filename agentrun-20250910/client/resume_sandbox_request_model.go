// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeSandboxRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ResumeSandboxRequest struct {
}

func (s ResumeSandboxRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeSandboxRequest) GoString() string {
	return s.String()
}

func (s *ResumeSandboxRequest) Validate() error {
	return dara.Validate(s)
}
