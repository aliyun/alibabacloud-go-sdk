// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceEndpointApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceEndpointApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceEndpointApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceEndpointApiKeyResponseBody) *GetServiceEndpointApiKeyResponse
	GetBody() *GetServiceEndpointApiKeyResponseBody
}

type GetServiceEndpointApiKeyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceEndpointApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceEndpointApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointApiKeyResponse) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceEndpointApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceEndpointApiKeyResponse) GetBody() *GetServiceEndpointApiKeyResponseBody {
	return s.Body
}

func (s *GetServiceEndpointApiKeyResponse) SetHeaders(v map[string]*string) *GetServiceEndpointApiKeyResponse {
	s.Headers = v
	return s
}

func (s *GetServiceEndpointApiKeyResponse) SetStatusCode(v int32) *GetServiceEndpointApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceEndpointApiKeyResponse) SetBody(v *GetServiceEndpointApiKeyResponseBody) *GetServiceEndpointApiKeyResponse {
	s.Body = v
	return s
}

func (s *GetServiceEndpointApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
