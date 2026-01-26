// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntegrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIntegrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIntegrationResponse
	GetStatusCode() *int32
	SetBody(v *CreateIntegrationResponseBody) *CreateIntegrationResponse
	GetBody() *CreateIntegrationResponseBody
}

type CreateIntegrationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIntegrationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIntegrationResponse) GoString() string {
	return s.String()
}

func (s *CreateIntegrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIntegrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIntegrationResponse) GetBody() *CreateIntegrationResponseBody {
	return s.Body
}

func (s *CreateIntegrationResponse) SetHeaders(v map[string]*string) *CreateIntegrationResponse {
	s.Headers = v
	return s
}

func (s *CreateIntegrationResponse) SetStatusCode(v int32) *CreateIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIntegrationResponse) SetBody(v *CreateIntegrationResponseBody) *CreateIntegrationResponse {
	s.Body = v
	return s
}

func (s *CreateIntegrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
