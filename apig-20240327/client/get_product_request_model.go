// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetProductRequest struct {
}

func (s GetProductRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProductRequest) GoString() string {
	return s.String()
}

func (s *GetProductRequest) Validate() error {
	return dara.Validate(s)
}
