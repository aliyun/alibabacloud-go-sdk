// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProductCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetProductCategoryRequest struct {
}

func (s GetProductCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProductCategoryRequest) GoString() string {
	return s.String()
}

func (s *GetProductCategoryRequest) Validate() error {
	return dara.Validate(s)
}
