// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistryModuleVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteRegistryModuleVersionRequest struct {
}

func (s DeleteRegistryModuleVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistryModuleVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteRegistryModuleVersionRequest) Validate() error {
	return dara.Validate(s)
}
