// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceInstanceTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceInstanceTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceInstanceTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceInstanceTokenResponseBody) *CreateServiceInstanceTokenResponse
	GetBody() *CreateServiceInstanceTokenResponseBody
}

type CreateServiceInstanceTokenResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceInstanceTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceInstanceTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceInstanceTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceInstanceTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceInstanceTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceInstanceTokenResponse) GetBody() *CreateServiceInstanceTokenResponseBody {
	return s.Body
}

func (s *CreateServiceInstanceTokenResponse) SetHeaders(v map[string]*string) *CreateServiceInstanceTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceInstanceTokenResponse) SetStatusCode(v int32) *CreateServiceInstanceTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceInstanceTokenResponse) SetBody(v *CreateServiceInstanceTokenResponseBody) *CreateServiceInstanceTokenResponse {
	s.Body = v
	return s
}

func (s *CreateServiceInstanceTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
