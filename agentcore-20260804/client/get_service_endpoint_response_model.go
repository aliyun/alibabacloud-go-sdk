// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceEndpointResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceEndpointResponseBody) *GetServiceEndpointResponse
	GetBody() *GetServiceEndpointResponseBody
}

type GetServiceEndpointResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointResponse) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceEndpointResponse) GetBody() *GetServiceEndpointResponseBody {
	return s.Body
}

func (s *GetServiceEndpointResponse) SetHeaders(v map[string]*string) *GetServiceEndpointResponse {
	s.Headers = v
	return s
}

func (s *GetServiceEndpointResponse) SetStatusCode(v int32) *GetServiceEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceEndpointResponse) SetBody(v *GetServiceEndpointResponseBody) *GetServiceEndpointResponse {
	s.Body = v
	return s
}

func (s *GetServiceEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
