// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSelectionProductSaleInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ProductSaleInfoListQuery) *ListSelectionProductSaleInfosRequest
	GetBody() *ProductSaleInfoListQuery
}

type ListSelectionProductSaleInfosRequest struct {
	// This parameter is required.
	Body *ProductSaleInfoListQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSelectionProductSaleInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSelectionProductSaleInfosRequest) GoString() string {
	return s.String()
}

func (s *ListSelectionProductSaleInfosRequest) GetBody() *ProductSaleInfoListQuery {
	return s.Body
}

func (s *ListSelectionProductSaleInfosRequest) SetBody(v *ProductSaleInfoListQuery) *ListSelectionProductSaleInfosRequest {
	s.Body = v
	return s
}

func (s *ListSelectionProductSaleInfosRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
