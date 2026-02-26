// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSelectionProductSaleInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSelectionProductSaleInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSelectionProductSaleInfosResponse
	GetStatusCode() *int32
	SetBody(v *ProductSaleInfoListResult) *ListSelectionProductSaleInfosResponse
	GetBody() *ProductSaleInfoListResult
}

type ListSelectionProductSaleInfosResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProductSaleInfoListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSelectionProductSaleInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSelectionProductSaleInfosResponse) GoString() string {
	return s.String()
}

func (s *ListSelectionProductSaleInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSelectionProductSaleInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSelectionProductSaleInfosResponse) GetBody() *ProductSaleInfoListResult {
	return s.Body
}

func (s *ListSelectionProductSaleInfosResponse) SetHeaders(v map[string]*string) *ListSelectionProductSaleInfosResponse {
	s.Headers = v
	return s
}

func (s *ListSelectionProductSaleInfosResponse) SetStatusCode(v int32) *ListSelectionProductSaleInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSelectionProductSaleInfosResponse) SetBody(v *ProductSaleInfoListResult) *ListSelectionProductSaleInfosResponse {
	s.Body = v
	return s
}

func (s *ListSelectionProductSaleInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
