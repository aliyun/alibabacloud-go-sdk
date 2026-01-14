// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBasicEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBasicEndpointResponse
	GetStatusCode() *int32
	SetBody(v *GetBasicEndpointResponseBody) *GetBasicEndpointResponse
	GetBody() *GetBasicEndpointResponseBody
}

type GetBasicEndpointResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBasicEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBasicEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBasicEndpointResponse) GoString() string {
	return s.String()
}

func (s *GetBasicEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBasicEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBasicEndpointResponse) GetBody() *GetBasicEndpointResponseBody {
	return s.Body
}

func (s *GetBasicEndpointResponse) SetHeaders(v map[string]*string) *GetBasicEndpointResponse {
	s.Headers = v
	return s
}

func (s *GetBasicEndpointResponse) SetStatusCode(v int32) *GetBasicEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBasicEndpointResponse) SetBody(v *GetBasicEndpointResponseBody) *GetBasicEndpointResponse {
	s.Body = v
	return s
}

func (s *GetBasicEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
