// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistryModuleRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteRegistryModuleRequest struct {
}

func (s DeleteRegistryModuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistryModuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRegistryModuleRequest) Validate() error {
	return dara.Validate(s)
}
