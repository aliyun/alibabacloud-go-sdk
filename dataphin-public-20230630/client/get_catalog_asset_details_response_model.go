// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogAssetDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCatalogAssetDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCatalogAssetDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetCatalogAssetDetailsResponseBody) *GetCatalogAssetDetailsResponse
	GetBody() *GetCatalogAssetDetailsResponseBody
}

type GetCatalogAssetDetailsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCatalogAssetDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCatalogAssetDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogAssetDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogAssetDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCatalogAssetDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCatalogAssetDetailsResponse) GetBody() *GetCatalogAssetDetailsResponseBody {
	return s.Body
}

func (s *GetCatalogAssetDetailsResponse) SetHeaders(v map[string]*string) *GetCatalogAssetDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogAssetDetailsResponse) SetStatusCode(v int32) *GetCatalogAssetDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogAssetDetailsResponse) SetBody(v *GetCatalogAssetDetailsResponseBody) *GetCatalogAssetDetailsResponse {
	s.Body = v
	return s
}

func (s *GetCatalogAssetDetailsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
