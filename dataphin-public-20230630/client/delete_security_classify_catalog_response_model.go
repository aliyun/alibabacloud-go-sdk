// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityClassifyCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecurityClassifyCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecurityClassifyCatalogResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecurityClassifyCatalogResponseBody) *DeleteSecurityClassifyCatalogResponse
	GetBody() *DeleteSecurityClassifyCatalogResponseBody
}

type DeleteSecurityClassifyCatalogResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecurityClassifyCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecurityClassifyCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityClassifyCatalogResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityClassifyCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecurityClassifyCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecurityClassifyCatalogResponse) GetBody() *DeleteSecurityClassifyCatalogResponseBody {
	return s.Body
}

func (s *DeleteSecurityClassifyCatalogResponse) SetHeaders(v map[string]*string) *DeleteSecurityClassifyCatalogResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityClassifyCatalogResponse) SetStatusCode(v int32) *DeleteSecurityClassifyCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityClassifyCatalogResponse) SetBody(v *DeleteSecurityClassifyCatalogResponseBody) *DeleteSecurityClassifyCatalogResponse {
	s.Body = v
	return s
}

func (s *DeleteSecurityClassifyCatalogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
