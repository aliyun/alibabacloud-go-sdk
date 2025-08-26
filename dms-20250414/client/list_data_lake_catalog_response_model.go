// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataLakeCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataLakeCatalogResponse
	GetStatusCode() *int32
	SetBody(v *ListDataLakeCatalogResponseBody) *ListDataLakeCatalogResponse
	GetBody() *ListDataLakeCatalogResponseBody
}

type ListDataLakeCatalogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataLakeCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataLakeCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeCatalogResponse) GoString() string {
	return s.String()
}

func (s *ListDataLakeCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataLakeCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataLakeCatalogResponse) GetBody() *ListDataLakeCatalogResponseBody {
	return s.Body
}

func (s *ListDataLakeCatalogResponse) SetHeaders(v map[string]*string) *ListDataLakeCatalogResponse {
	s.Headers = v
	return s
}

func (s *ListDataLakeCatalogResponse) SetStatusCode(v int32) *ListDataLakeCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataLakeCatalogResponse) SetBody(v *ListDataLakeCatalogResponseBody) *ListDataLakeCatalogResponse {
	s.Body = v
	return s
}

func (s *ListDataLakeCatalogResponse) Validate() error {
	return dara.Validate(s)
}
