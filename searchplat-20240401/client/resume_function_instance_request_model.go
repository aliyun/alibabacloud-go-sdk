// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeFunctionInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ResumeFunctionInstanceRequest struct {
}

func (s ResumeFunctionInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeFunctionInstanceRequest) GoString() string {
	return s.String()
}

func (s *ResumeFunctionInstanceRequest) Validate() error {
	return dara.Validate(s)
}
