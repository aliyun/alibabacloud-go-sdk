// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModuleRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteModuleRequest struct {
}

func (s DeleteModuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteModuleRequest) Validate() error {
	return dara.Validate(s)
}
