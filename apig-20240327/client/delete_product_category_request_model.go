// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProductCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteProductCategoryRequest struct {
}

func (s DeleteProductCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProductCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteProductCategoryRequest) Validate() error {
	return dara.Validate(s)
}
