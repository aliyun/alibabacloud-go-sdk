// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterruptRequest interface {
	dara.Model
	String() string
	GoString() string
}

type InterruptRequest struct {
}

func (s InterruptRequest) String() string {
	return dara.Prettify(s)
}

func (s InterruptRequest) GoString() string {
	return s.String()
}

func (s *InterruptRequest) Validate() error {
	return dara.Validate(s)
}
