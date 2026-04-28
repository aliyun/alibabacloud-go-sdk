// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMemoryTaskRequest struct {
}

func (s GetMemoryTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryTaskRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryTaskRequest) Validate() error {
	return dara.Validate(s)
}
