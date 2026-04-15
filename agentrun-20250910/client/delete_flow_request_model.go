// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteFlowRequest struct {
}

func (s DeleteFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteFlowRequest) Validate() error {
	return dara.Validate(s)
}
