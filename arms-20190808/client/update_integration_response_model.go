// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntegrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateIntegrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateIntegrationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateIntegrationResponseBody) *UpdateIntegrationResponse
	GetBody() *UpdateIntegrationResponseBody
}

type UpdateIntegrationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateIntegrationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntegrationResponse) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateIntegrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateIntegrationResponse) GetBody() *UpdateIntegrationResponseBody {
	return s.Body
}

func (s *UpdateIntegrationResponse) SetHeaders(v map[string]*string) *UpdateIntegrationResponse {
	s.Headers = v
	return s
}

func (s *UpdateIntegrationResponse) SetStatusCode(v int32) *UpdateIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateIntegrationResponse) SetBody(v *UpdateIntegrationResponseBody) *UpdateIntegrationResponse {
	s.Body = v
	return s
}

func (s *UpdateIntegrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
