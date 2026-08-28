// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMcpRequest struct {
}

func (s GetMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcpRequest) GoString() string {
	return s.String()
}

func (s *GetMcpRequest) Validate() error {
	return dara.Validate(s)
}
