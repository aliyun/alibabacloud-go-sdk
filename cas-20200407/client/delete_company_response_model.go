// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCompanyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCompanyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCompanyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCompanyResponseBody) *DeleteCompanyResponse
	GetBody() *DeleteCompanyResponseBody
}

type DeleteCompanyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCompanyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCompanyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCompanyResponse) GoString() string {
	return s.String()
}

func (s *DeleteCompanyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCompanyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCompanyResponse) GetBody() *DeleteCompanyResponseBody {
	return s.Body
}

func (s *DeleteCompanyResponse) SetHeaders(v map[string]*string) *DeleteCompanyResponse {
	s.Headers = v
	return s
}

func (s *DeleteCompanyResponse) SetStatusCode(v int32) *DeleteCompanyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCompanyResponse) SetBody(v *DeleteCompanyResponseBody) *DeleteCompanyResponse {
	s.Body = v
	return s
}

func (s *DeleteCompanyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
