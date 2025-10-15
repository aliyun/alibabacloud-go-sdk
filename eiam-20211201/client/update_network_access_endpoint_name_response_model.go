// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNetworkAccessEndpointNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNetworkAccessEndpointNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNetworkAccessEndpointNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNetworkAccessEndpointNameResponseBody) *UpdateNetworkAccessEndpointNameResponse
	GetBody() *UpdateNetworkAccessEndpointNameResponseBody
}

type UpdateNetworkAccessEndpointNameResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNetworkAccessEndpointNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNetworkAccessEndpointNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNetworkAccessEndpointNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateNetworkAccessEndpointNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNetworkAccessEndpointNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNetworkAccessEndpointNameResponse) GetBody() *UpdateNetworkAccessEndpointNameResponseBody {
	return s.Body
}

func (s *UpdateNetworkAccessEndpointNameResponse) SetHeaders(v map[string]*string) *UpdateNetworkAccessEndpointNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateNetworkAccessEndpointNameResponse) SetStatusCode(v int32) *UpdateNetworkAccessEndpointNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNetworkAccessEndpointNameResponse) SetBody(v *UpdateNetworkAccessEndpointNameResponseBody) *UpdateNetworkAccessEndpointNameResponse {
	s.Body = v
	return s
}

func (s *UpdateNetworkAccessEndpointNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
