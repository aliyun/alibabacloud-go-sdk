// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityClassifyCatalogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSecurityClassifyCatalogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSecurityClassifyCatalogResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSecurityClassifyCatalogResponseBody) *UpdateSecurityClassifyCatalogResponse
	GetBody() *UpdateSecurityClassifyCatalogResponseBody
}

type UpdateSecurityClassifyCatalogResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecurityClassifyCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecurityClassifyCatalogResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityClassifyCatalogResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecurityClassifyCatalogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSecurityClassifyCatalogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSecurityClassifyCatalogResponse) GetBody() *UpdateSecurityClassifyCatalogResponseBody {
	return s.Body
}

func (s *UpdateSecurityClassifyCatalogResponse) SetHeaders(v map[string]*string) *UpdateSecurityClassifyCatalogResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecurityClassifyCatalogResponse) SetStatusCode(v int32) *UpdateSecurityClassifyCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecurityClassifyCatalogResponse) SetBody(v *UpdateSecurityClassifyCatalogResponseBody) *UpdateSecurityClassifyCatalogResponse {
	s.Body = v
	return s
}

func (s *UpdateSecurityClassifyCatalogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
