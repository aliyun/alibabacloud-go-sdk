// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIntegrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddIntegrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddIntegrationResponse
	GetStatusCode() *int32
	SetBody(v *AddIntegrationResponseBody) *AddIntegrationResponse
	GetBody() *AddIntegrationResponseBody
}

type AddIntegrationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddIntegrationResponse) String() string {
	return dara.Prettify(s)
}

func (s AddIntegrationResponse) GoString() string {
	return s.String()
}

func (s *AddIntegrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddIntegrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddIntegrationResponse) GetBody() *AddIntegrationResponseBody {
	return s.Body
}

func (s *AddIntegrationResponse) SetHeaders(v map[string]*string) *AddIntegrationResponse {
	s.Headers = v
	return s
}

func (s *AddIntegrationResponse) SetStatusCode(v int32) *AddIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIntegrationResponse) SetBody(v *AddIntegrationResponseBody) *AddIntegrationResponse {
	s.Body = v
	return s
}

func (s *AddIntegrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
