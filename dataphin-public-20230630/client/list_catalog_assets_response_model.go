// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCatalogAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCatalogAssetsResponse
	GetStatusCode() *int32
	SetBody(v *ListCatalogAssetsResponseBody) *ListCatalogAssetsResponse
	GetBody() *ListCatalogAssetsResponseBody
}

type ListCatalogAssetsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCatalogAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCatalogAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogAssetsResponse) GoString() string {
	return s.String()
}

func (s *ListCatalogAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCatalogAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCatalogAssetsResponse) GetBody() *ListCatalogAssetsResponseBody {
	return s.Body
}

func (s *ListCatalogAssetsResponse) SetHeaders(v map[string]*string) *ListCatalogAssetsResponse {
	s.Headers = v
	return s
}

func (s *ListCatalogAssetsResponse) SetStatusCode(v int32) *ListCatalogAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCatalogAssetsResponse) SetBody(v *ListCatalogAssetsResponseBody) *ListCatalogAssetsResponse {
	s.Body = v
	return s
}

func (s *ListCatalogAssetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
