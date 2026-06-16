// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextStoreRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteContextStoreRequest struct {
}

func (s DeleteContextStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextStoreRequest) GoString() string {
	return s.String()
}

func (s *DeleteContextStoreRequest) Validate() error {
	return dara.Validate(s)
}
