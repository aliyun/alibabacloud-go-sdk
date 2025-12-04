// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkInterfacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetworkInterfacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetworkInterfacesResponse
	GetStatusCode() *int32
	SetBody(v *ListNetworkInterfacesResponseBody) *ListNetworkInterfacesResponse
	GetBody() *ListNetworkInterfacesResponseBody
}

type ListNetworkInterfacesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkInterfacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkInterfacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkInterfacesResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkInterfacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetworkInterfacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetworkInterfacesResponse) GetBody() *ListNetworkInterfacesResponseBody {
	return s.Body
}

func (s *ListNetworkInterfacesResponse) SetHeaders(v map[string]*string) *ListNetworkInterfacesResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkInterfacesResponse) SetStatusCode(v int32) *ListNetworkInterfacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkInterfacesResponse) SetBody(v *ListNetworkInterfacesResponseBody) *ListNetworkInterfacesResponse {
	s.Body = v
	return s
}

func (s *ListNetworkInterfacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
