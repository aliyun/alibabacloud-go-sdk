// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLumaCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLumaCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLumaCatalogResponse
	GetStatusCode() *int32
	SetBody(v *GetLumaCatalogResponseBody) *GetLumaCatalogResponse
	GetBody() *GetLumaCatalogResponseBody
}

type GetLumaCatalogResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLumaCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLumaCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLumaCatalogResponse) GoString() string {
	return s.String()
}

func (s *GetLumaCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLumaCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLumaCatalogResponse) GetBody() *GetLumaCatalogResponseBody {
	return s.Body
}

func (s *GetLumaCatalogResponse) SetHeaders(v map[string]*string) *GetLumaCatalogResponse {
	s.Headers = v
	return s
}

func (s *GetLumaCatalogResponse) SetStatusCode(v int32) *GetLumaCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLumaCatalogResponse) SetBody(v *GetLumaCatalogResponseBody) *GetLumaCatalogResponse {
	s.Body = v
	return s
}

func (s *GetLumaCatalogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
