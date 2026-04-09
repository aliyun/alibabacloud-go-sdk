// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMemoryCollectionRequest struct {
}

func (s GetMemoryCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryCollectionRequest) GoString() string {
	return s.String()
}

func (s *GetMemoryCollectionRequest) Validate() error {
	return dara.Validate(s)
}
