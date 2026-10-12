// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductRefRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteProductRefRequest struct {
}

func (s DeleteProductRefRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductRefRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductRefRequest) Validate() error {
	return dara.Validate(s)
}
