// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaultInjectionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteFaultInjectionRequest struct {
}

func (s DeleteFaultInjectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaultInjectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaultInjectionRequest) Validate() error {
	return dara.Validate(s)
}
