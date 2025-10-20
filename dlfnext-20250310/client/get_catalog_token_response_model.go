// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCatalogTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCatalogTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetCatalogTokenResponseBody) *GetCatalogTokenResponse
	GetBody() *GetCatalogTokenResponseBody
}

type GetCatalogTokenResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCatalogTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCatalogTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogTokenResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCatalogTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCatalogTokenResponse) GetBody() *GetCatalogTokenResponseBody {
	return s.Body
}

func (s *GetCatalogTokenResponse) SetHeaders(v map[string]*string) *GetCatalogTokenResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogTokenResponse) SetStatusCode(v int32) *GetCatalogTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogTokenResponse) SetBody(v *GetCatalogTokenResponseBody) *GetCatalogTokenResponse {
	s.Body = v
	return s
}

func (s *GetCatalogTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
