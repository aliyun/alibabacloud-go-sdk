// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceVpcEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceVpcEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceVpcEndpointResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceVpcEndpointResponseBody) *GetInstanceVpcEndpointResponse
	GetBody() *GetInstanceVpcEndpointResponseBody
}

type GetInstanceVpcEndpointResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceVpcEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceVpcEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceVpcEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceVpcEndpointResponse) GetBody() *GetInstanceVpcEndpointResponseBody {
	return s.Body
}

func (s *GetInstanceVpcEndpointResponse) SetHeaders(v map[string]*string) *GetInstanceVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceVpcEndpointResponse) SetStatusCode(v int32) *GetInstanceVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceVpcEndpointResponse) SetBody(v *GetInstanceVpcEndpointResponseBody) *GetInstanceVpcEndpointResponse {
	s.Body = v
	return s
}

func (s *GetInstanceVpcEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
