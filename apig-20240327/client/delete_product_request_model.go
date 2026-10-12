// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteProductRequest struct {
}

func (s DeleteProductRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductRequest) Validate() error {
	return dara.Validate(s)
}
