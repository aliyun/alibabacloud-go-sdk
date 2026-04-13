// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetModuleVersionRequest struct {
}

func (s GetModuleVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModuleVersionRequest) GoString() string {
	return s.String()
}

func (s *GetModuleVersionRequest) Validate() error {
	return dara.Validate(s)
}
