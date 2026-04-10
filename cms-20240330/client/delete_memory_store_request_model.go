// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryStoreRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteMemoryStoreRequest struct {
}

func (s DeleteMemoryStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryStoreRequest) GoString() string {
	return s.String()
}

func (s *DeleteMemoryStoreRequest) Validate() error {
	return dara.Validate(s)
}
