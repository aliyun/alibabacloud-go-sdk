// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetModuleRequest struct {
}

func (s GetModuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModuleRequest) GoString() string {
	return s.String()
}

func (s *GetModuleRequest) Validate() error {
	return dara.Validate(s)
}
