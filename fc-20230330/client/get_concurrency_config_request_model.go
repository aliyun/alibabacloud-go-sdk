// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConcurrencyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetConcurrencyConfigRequest struct {
}

func (s GetConcurrencyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConcurrencyConfigRequest) GoString() string {
	return s.String()
}

func (s *GetConcurrencyConfigRequest) Validate() error {
	return dara.Validate(s)
}
