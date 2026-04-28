// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryHealthRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMemoryHealthRequest struct {
}

func (s GetMemoryHealthRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryHealthRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryHealthRequest) Validate() error {
	return dara.Validate(s)
}
