// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistryModuleRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetRegistryModuleRequest struct {
}

func (s GetRegistryModuleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRegistryModuleRequest) GoString() string {
	return s.String()
}

func (s *GetRegistryModuleRequest) Validate() error {
	return dara.Validate(s)
}
