// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSelectionSkuSaleInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SkuSaleInfoListQuery) *ListSelectionSkuSaleInfosRequest
	GetBody() *SkuSaleInfoListQuery
}

type ListSelectionSkuSaleInfosRequest struct {
	// This parameter is required.
	Body *SkuSaleInfoListQuery `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSelectionSkuSaleInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSelectionSkuSaleInfosRequest) GoString() string {
	return s.String()
}

func (s *ListSelectionSkuSaleInfosRequest) GetBody() *SkuSaleInfoListQuery {
	return s.Body
}

func (s *ListSelectionSkuSaleInfosRequest) SetBody(v *SkuSaleInfoListQuery) *ListSelectionSkuSaleInfosRequest {
	s.Body = v
	return s
}

func (s *ListSelectionSkuSaleInfosRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
