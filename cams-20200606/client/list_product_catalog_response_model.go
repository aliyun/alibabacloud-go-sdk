// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductCatalogResponse
	GetStatusCode() *int32
	SetBody(v *ListProductCatalogResponseBody) *ListProductCatalogResponse
	GetBody() *ListProductCatalogResponseBody
}

type ListProductCatalogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductCatalogResponse) GoString() string {
	return s.String()
}

func (s *ListProductCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductCatalogResponse) GetBody() *ListProductCatalogResponseBody {
	return s.Body
}

func (s *ListProductCatalogResponse) SetHeaders(v map[string]*string) *ListProductCatalogResponse {
	s.Headers = v
	return s
}

func (s *ListProductCatalogResponse) SetStatusCode(v int32) *ListProductCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductCatalogResponse) SetBody(v *ListProductCatalogResponseBody) *ListProductCatalogResponse {
	s.Body = v
	return s
}

func (s *ListProductCatalogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
