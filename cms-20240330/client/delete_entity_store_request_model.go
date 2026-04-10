// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEntityStoreRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteEntityStoreRequest struct {
}

func (s DeleteEntityStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEntityStoreRequest) GoString() string {
	return s.String()
}

func (s *DeleteEntityStoreRequest) Validate() error {
	return dara.Validate(s)
}
