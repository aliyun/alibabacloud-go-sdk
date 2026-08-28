// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteMcpRequest struct {
}

func (s DeleteMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcpRequest) Validate() error {
	return dara.Validate(s)
}
