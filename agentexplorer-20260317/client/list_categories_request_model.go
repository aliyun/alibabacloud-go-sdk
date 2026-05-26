// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoriesRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListCategoriesRequest struct {
}

func (s ListCategoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCategoriesRequest) GoString() string {
	return s.String()
}

func (s *ListCategoriesRequest) Validate() error {
	return dara.Validate(s)
}
