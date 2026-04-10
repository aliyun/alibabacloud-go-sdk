// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryStoreRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMemoryStoreRequest struct {
}

func (s GetMemoryStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryStoreRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryStoreRequest) Validate() error {
	return dara.Validate(s)
}
