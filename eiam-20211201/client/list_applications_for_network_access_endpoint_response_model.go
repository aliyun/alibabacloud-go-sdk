// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForNetworkAccessEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationsForNetworkAccessEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationsForNetworkAccessEndpointResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationsForNetworkAccessEndpointResponseBody) *ListApplicationsForNetworkAccessEndpointResponse
	GetBody() *ListApplicationsForNetworkAccessEndpointResponseBody
}

type ListApplicationsForNetworkAccessEndpointResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsForNetworkAccessEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsForNetworkAccessEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForNetworkAccessEndpointResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsForNetworkAccessEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationsForNetworkAccessEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationsForNetworkAccessEndpointResponse) GetBody() *ListApplicationsForNetworkAccessEndpointResponseBody {
	return s.Body
}

func (s *ListApplicationsForNetworkAccessEndpointResponse) SetHeaders(v map[string]*string) *ListApplicationsForNetworkAccessEndpointResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointResponse) SetStatusCode(v int32) *ListApplicationsForNetworkAccessEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointResponse) SetBody(v *ListApplicationsForNetworkAccessEndpointResponseBody) *ListApplicationsForNetworkAccessEndpointResponse {
	s.Body = v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
