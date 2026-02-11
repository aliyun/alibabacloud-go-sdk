// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegrationTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIntegrationTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIntegrationTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetIntegrationTokenResponseBody) *GetIntegrationTokenResponse
	GetBody() *GetIntegrationTokenResponseBody
}

type GetIntegrationTokenResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIntegrationTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIntegrationTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationTokenResponse) GoString() string {
	return s.String()
}

func (s *GetIntegrationTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIntegrationTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIntegrationTokenResponse) GetBody() *GetIntegrationTokenResponseBody {
	return s.Body
}

func (s *GetIntegrationTokenResponse) SetHeaders(v map[string]*string) *GetIntegrationTokenResponse {
	s.Headers = v
	return s
}

func (s *GetIntegrationTokenResponse) SetStatusCode(v int32) *GetIntegrationTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntegrationTokenResponse) SetBody(v *GetIntegrationTokenResponseBody) *GetIntegrationTokenResponse {
	s.Body = v
	return s
}

func (s *GetIntegrationTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
