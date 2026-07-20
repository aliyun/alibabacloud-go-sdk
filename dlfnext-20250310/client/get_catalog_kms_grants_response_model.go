// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCatalogKmsGrantsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCatalogKmsGrantsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCatalogKmsGrantsResponse
	GetStatusCode() *int32
	SetBody(v *GetCatalogKmsGrantsResponseBody) *GetCatalogKmsGrantsResponse
	GetBody() *GetCatalogKmsGrantsResponseBody
}

type GetCatalogKmsGrantsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCatalogKmsGrantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCatalogKmsGrantsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCatalogKmsGrantsResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogKmsGrantsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCatalogKmsGrantsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCatalogKmsGrantsResponse) GetBody() *GetCatalogKmsGrantsResponseBody {
	return s.Body
}

func (s *GetCatalogKmsGrantsResponse) SetHeaders(v map[string]*string) *GetCatalogKmsGrantsResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogKmsGrantsResponse) SetStatusCode(v int32) *GetCatalogKmsGrantsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogKmsGrantsResponse) SetBody(v *GetCatalogKmsGrantsResponseBody) *GetCatalogKmsGrantsResponse {
	s.Body = v
	return s
}

func (s *GetCatalogKmsGrantsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
