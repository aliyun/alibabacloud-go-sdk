// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CategoryListQuery) *ListCategoriesRequest
	GetBody() *CategoryListQuery
}

type ListCategoriesRequest struct {
	Body *CategoryListQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCategoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCategoriesRequest) GoString() string {
	return s.String()
}

func (s *ListCategoriesRequest) GetBody() *CategoryListQuery {
	return s.Body
}

func (s *ListCategoriesRequest) SetBody(v *CategoryListQuery) *ListCategoriesRequest {
	s.Body = v
	return s
}

func (s *ListCategoriesRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
