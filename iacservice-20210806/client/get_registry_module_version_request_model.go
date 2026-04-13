// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistryModuleVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetRegistryModuleVersionRequest struct {
}

func (s GetRegistryModuleVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRegistryModuleVersionRequest) GoString() string {
	return s.String()
}

func (s *GetRegistryModuleVersionRequest) Validate() error {
	return dara.Validate(s)
}
