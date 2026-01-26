// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegrationStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIntegrationStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIntegrationStateResponse
	GetStatusCode() *int32
	SetBody(v *GetIntegrationStateResponseBody) *GetIntegrationStateResponse
	GetBody() *GetIntegrationStateResponseBody
}

type GetIntegrationStateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIntegrationStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIntegrationStateResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationStateResponse) GoString() string {
	return s.String()
}

func (s *GetIntegrationStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIntegrationStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIntegrationStateResponse) GetBody() *GetIntegrationStateResponseBody {
	return s.Body
}

func (s *GetIntegrationStateResponse) SetHeaders(v map[string]*string) *GetIntegrationStateResponse {
	s.Headers = v
	return s
}

func (s *GetIntegrationStateResponse) SetStatusCode(v int32) *GetIntegrationStateResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIntegrationStateResponse) SetBody(v *GetIntegrationStateResponseBody) *GetIntegrationStateResponse {
	s.Body = v
	return s
}

func (s *GetIntegrationStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
