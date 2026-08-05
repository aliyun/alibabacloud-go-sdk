// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCeaseFunctionInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CeaseFunctionInstanceRequest struct {
}

func (s CeaseFunctionInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CeaseFunctionInstanceRequest) GoString() string {
	return s.String()
}

func (s *CeaseFunctionInstanceRequest) Validate() error {
	return dara.Validate(s)
}
