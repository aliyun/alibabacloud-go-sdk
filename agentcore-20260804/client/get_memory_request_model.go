// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMemoryRequest struct {
}

func (s GetMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryRequest) Validate() error {
	return dara.Validate(s)
}
