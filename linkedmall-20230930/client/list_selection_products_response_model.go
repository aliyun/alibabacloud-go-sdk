// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSelectionProductsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSelectionProductsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSelectionProductsResponse
	GetStatusCode() *int32
	SetBody(v *ProductPageResult) *ListSelectionProductsResponse
	GetBody() *ProductPageResult
}

type ListSelectionProductsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProductPageResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSelectionProductsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSelectionProductsResponse) GoString() string {
	return s.String()
}

func (s *ListSelectionProductsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSelectionProductsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSelectionProductsResponse) GetBody() *ProductPageResult {
	return s.Body
}

func (s *ListSelectionProductsResponse) SetHeaders(v map[string]*string) *ListSelectionProductsResponse {
	s.Headers = v
	return s
}

func (s *ListSelectionProductsResponse) SetStatusCode(v int32) *ListSelectionProductsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSelectionProductsResponse) SetBody(v *ProductPageResult) *ListSelectionProductsResponse {
	s.Body = v
	return s
}

func (s *ListSelectionProductsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
