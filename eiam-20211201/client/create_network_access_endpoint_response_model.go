// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkAccessEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetworkAccessEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetworkAccessEndpointResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetworkAccessEndpointResponseBody) *CreateNetworkAccessEndpointResponse
	GetBody() *CreateNetworkAccessEndpointResponseBody
}

type CreateNetworkAccessEndpointResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetworkAccessEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetworkAccessEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAccessEndpointResponse) GoString() string {
	return s.String()
}

func (s *CreateNetworkAccessEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetworkAccessEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetworkAccessEndpointResponse) GetBody() *CreateNetworkAccessEndpointResponseBody {
	return s.Body
}

func (s *CreateNetworkAccessEndpointResponse) SetHeaders(v map[string]*string) *CreateNetworkAccessEndpointResponse {
	s.Headers = v
	return s
}

func (s *CreateNetworkAccessEndpointResponse) SetStatusCode(v int32) *CreateNetworkAccessEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetworkAccessEndpointResponse) SetBody(v *CreateNetworkAccessEndpointResponseBody) *CreateNetworkAccessEndpointResponse {
	s.Body = v
	return s
}

func (s *CreateNetworkAccessEndpointResponse) Validate() error {
	return dara.Validate(s)
}
