// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceCatalogResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceCatalogResponseBody) *ListInstanceCatalogResponse
	GetBody() *ListInstanceCatalogResponseBody
}

type ListInstanceCatalogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceCatalogResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceCatalogResponse) GetBody() *ListInstanceCatalogResponseBody {
	return s.Body
}

func (s *ListInstanceCatalogResponse) SetHeaders(v map[string]*string) *ListInstanceCatalogResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceCatalogResponse) SetStatusCode(v int32) *ListInstanceCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceCatalogResponse) SetBody(v *ListInstanceCatalogResponseBody) *ListInstanceCatalogResponse {
	s.Body = v
	return s
}

func (s *ListInstanceCatalogResponse) Validate() error {
	return dara.Validate(s)
}
