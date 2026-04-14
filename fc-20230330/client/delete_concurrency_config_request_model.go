// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConcurrencyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteConcurrencyConfigRequest struct {
}

func (s DeleteConcurrencyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConcurrencyConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteConcurrencyConfigRequest) Validate() error {
	return dara.Validate(s)
}
