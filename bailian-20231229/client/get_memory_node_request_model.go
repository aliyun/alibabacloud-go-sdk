// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryNodeRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMemoryNodeRequest struct {
}

func (s GetMemoryNodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryNodeRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryNodeRequest) Validate() error {
	return dara.Validate(s)
}
