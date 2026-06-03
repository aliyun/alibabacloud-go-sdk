// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteConnectorRequest struct {
}

func (s DeleteConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectorRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnectorRequest) Validate() error {
	return dara.Validate(s)
}
