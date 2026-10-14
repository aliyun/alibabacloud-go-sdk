// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgenticCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgenticCatalogResponse
	GetStatusCode() *int32
	SetBody(v *GetAgenticCatalogResponseBody) *GetAgenticCatalogResponse
	GetBody() *GetAgenticCatalogResponseBody
}

type GetAgenticCatalogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgenticCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgenticCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticCatalogResponse) GoString() string {
	return s.String()
}

func (s *GetAgenticCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgenticCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgenticCatalogResponse) GetBody() *GetAgenticCatalogResponseBody {
	return s.Body
}

func (s *GetAgenticCatalogResponse) SetHeaders(v map[string]*string) *GetAgenticCatalogResponse {
	s.Headers = v
	return s
}

func (s *GetAgenticCatalogResponse) SetStatusCode(v int32) *GetAgenticCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgenticCatalogResponse) SetBody(v *GetAgenticCatalogResponseBody) *GetAgenticCatalogResponse {
	s.Body = v
	return s
}

func (s *GetAgenticCatalogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
