// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataLakeCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataLakeCatalogResponse
	GetStatusCode() *int32
	SetBody(v *GetDataLakeCatalogResponseBody) *GetDataLakeCatalogResponse
	GetBody() *GetDataLakeCatalogResponseBody
}

type GetDataLakeCatalogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataLakeCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataLakeCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeCatalogResponse) GoString() string {
	return s.String()
}

func (s *GetDataLakeCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataLakeCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataLakeCatalogResponse) GetBody() *GetDataLakeCatalogResponseBody {
	return s.Body
}

func (s *GetDataLakeCatalogResponse) SetHeaders(v map[string]*string) *GetDataLakeCatalogResponse {
	s.Headers = v
	return s
}

func (s *GetDataLakeCatalogResponse) SetStatusCode(v int32) *GetDataLakeCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataLakeCatalogResponse) SetBody(v *GetDataLakeCatalogResponseBody) *GetDataLakeCatalogResponse {
	s.Body = v
	return s
}

func (s *GetDataLakeCatalogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
