// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompanyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCompanyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCompanyResponse
	GetStatusCode() *int32
	SetBody(v *GetCompanyResponseBody) *GetCompanyResponse
	GetBody() *GetCompanyResponseBody
}

type GetCompanyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCompanyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCompanyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCompanyResponse) GoString() string {
	return s.String()
}

func (s *GetCompanyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCompanyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCompanyResponse) GetBody() *GetCompanyResponseBody {
	return s.Body
}

func (s *GetCompanyResponse) SetHeaders(v map[string]*string) *GetCompanyResponse {
	s.Headers = v
	return s
}

func (s *GetCompanyResponse) SetStatusCode(v int32) *GetCompanyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCompanyResponse) SetBody(v *GetCompanyResponseBody) *GetCompanyResponse {
	s.Body = v
	return s
}

func (s *GetCompanyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
