// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteFlowVersionRequest struct {
}

func (s DeleteFlowVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowVersionRequest) Validate() error {
	return dara.Validate(s)
}
