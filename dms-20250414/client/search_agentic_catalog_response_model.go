// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAgenticCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchAgenticCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchAgenticCatalogResponse
	GetStatusCode() *int32
	SetBody(v *SearchAgenticCatalogResponseBody) *SearchAgenticCatalogResponse
	GetBody() *SearchAgenticCatalogResponseBody
}

type SearchAgenticCatalogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchAgenticCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchAgenticCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchAgenticCatalogResponse) GoString() string {
	return s.String()
}

func (s *SearchAgenticCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchAgenticCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchAgenticCatalogResponse) GetBody() *SearchAgenticCatalogResponseBody {
	return s.Body
}

func (s *SearchAgenticCatalogResponse) SetHeaders(v map[string]*string) *SearchAgenticCatalogResponse {
	s.Headers = v
	return s
}

func (s *SearchAgenticCatalogResponse) SetStatusCode(v int32) *SearchAgenticCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAgenticCatalogResponse) SetBody(v *SearchAgenticCatalogResponseBody) *SearchAgenticCatalogResponse {
	s.Body = v
	return s
}

func (s *SearchAgenticCatalogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
