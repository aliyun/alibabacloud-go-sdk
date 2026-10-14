// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgenticCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgenticCatalogResponse
	GetStatusCode() *int32
	SetBody(v *ListAgenticCatalogResponseBody) *ListAgenticCatalogResponse
	GetBody() *ListAgenticCatalogResponseBody
}

type ListAgenticCatalogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgenticCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgenticCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticCatalogResponse) GoString() string {
	return s.String()
}

func (s *ListAgenticCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgenticCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgenticCatalogResponse) GetBody() *ListAgenticCatalogResponseBody {
	return s.Body
}

func (s *ListAgenticCatalogResponse) SetHeaders(v map[string]*string) *ListAgenticCatalogResponse {
	s.Headers = v
	return s
}

func (s *ListAgenticCatalogResponse) SetStatusCode(v int32) *ListAgenticCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgenticCatalogResponse) SetBody(v *ListAgenticCatalogResponseBody) *ListAgenticCatalogResponse {
	s.Body = v
	return s
}

func (s *ListAgenticCatalogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
