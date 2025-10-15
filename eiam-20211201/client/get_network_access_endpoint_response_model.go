// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkAccessEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNetworkAccessEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNetworkAccessEndpointResponse
	GetStatusCode() *int32
	SetBody(v *GetNetworkAccessEndpointResponseBody) *GetNetworkAccessEndpointResponse
	GetBody() *GetNetworkAccessEndpointResponseBody
}

type GetNetworkAccessEndpointResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkAccessEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkAccessEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkAccessEndpointResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkAccessEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNetworkAccessEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNetworkAccessEndpointResponse) GetBody() *GetNetworkAccessEndpointResponseBody {
	return s.Body
}

func (s *GetNetworkAccessEndpointResponse) SetHeaders(v map[string]*string) *GetNetworkAccessEndpointResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkAccessEndpointResponse) SetStatusCode(v int32) *GetNetworkAccessEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkAccessEndpointResponse) SetBody(v *GetNetworkAccessEndpointResponseBody) *GetNetworkAccessEndpointResponse {
	s.Body = v
	return s
}

func (s *GetNetworkAccessEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
