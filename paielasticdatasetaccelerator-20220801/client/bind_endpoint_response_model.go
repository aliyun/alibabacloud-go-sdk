// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindEndpointResponse
	GetStatusCode() *int32
	SetBody(v *BindEndpointResponseBody) *BindEndpointResponse
	GetBody() *BindEndpointResponseBody
}

type BindEndpointResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s BindEndpointResponse) GoString() string {
	return s.String()
}

func (s *BindEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindEndpointResponse) GetBody() *BindEndpointResponseBody {
	return s.Body
}

func (s *BindEndpointResponse) SetHeaders(v map[string]*string) *BindEndpointResponse {
	s.Headers = v
	return s
}

func (s *BindEndpointResponse) SetStatusCode(v int32) *BindEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *BindEndpointResponse) SetBody(v *BindEndpointResponseBody) *BindEndpointResponse {
	s.Body = v
	return s
}

func (s *BindEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
