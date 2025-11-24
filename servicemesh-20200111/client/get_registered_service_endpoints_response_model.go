// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegisteredServiceEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRegisteredServiceEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRegisteredServiceEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *GetRegisteredServiceEndpointsResponseBody) *GetRegisteredServiceEndpointsResponse
	GetBody() *GetRegisteredServiceEndpointsResponseBody
}

type GetRegisteredServiceEndpointsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRegisteredServiceEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRegisteredServiceEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRegisteredServiceEndpointsResponse) GoString() string {
	return s.String()
}

func (s *GetRegisteredServiceEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRegisteredServiceEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRegisteredServiceEndpointsResponse) GetBody() *GetRegisteredServiceEndpointsResponseBody {
	return s.Body
}

func (s *GetRegisteredServiceEndpointsResponse) SetHeaders(v map[string]*string) *GetRegisteredServiceEndpointsResponse {
	s.Headers = v
	return s
}

func (s *GetRegisteredServiceEndpointsResponse) SetStatusCode(v int32) *GetRegisteredServiceEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegisteredServiceEndpointsResponse) SetBody(v *GetRegisteredServiceEndpointsResponseBody) *GetRegisteredServiceEndpointsResponse {
	s.Body = v
	return s
}

func (s *GetRegisteredServiceEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
