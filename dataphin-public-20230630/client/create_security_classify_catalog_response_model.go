// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityClassifyCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecurityClassifyCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecurityClassifyCatalogResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecurityClassifyCatalogResponseBody) *CreateSecurityClassifyCatalogResponse
	GetBody() *CreateSecurityClassifyCatalogResponseBody
}

type CreateSecurityClassifyCatalogResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecurityClassifyCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecurityClassifyCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityClassifyCatalogResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityClassifyCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecurityClassifyCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecurityClassifyCatalogResponse) GetBody() *CreateSecurityClassifyCatalogResponseBody {
	return s.Body
}

func (s *CreateSecurityClassifyCatalogResponse) SetHeaders(v map[string]*string) *CreateSecurityClassifyCatalogResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityClassifyCatalogResponse) SetStatusCode(v int32) *CreateSecurityClassifyCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecurityClassifyCatalogResponse) SetBody(v *CreateSecurityClassifyCatalogResponseBody) *CreateSecurityClassifyCatalogResponse {
	s.Body = v
	return s
}

func (s *CreateSecurityClassifyCatalogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
