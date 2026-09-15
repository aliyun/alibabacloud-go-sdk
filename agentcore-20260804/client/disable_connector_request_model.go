// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DisableConnectorRequest struct {
}

func (s DisableConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableConnectorRequest) GoString() string {
	return s.String()
}

func (s *DisableConnectorRequest) Validate() error {
	return dara.Validate(s)
}
