// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkInterfacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNetworkInterfacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNetworkInterfacesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNetworkInterfacesResponseBody) *DeleteNetworkInterfacesResponse
	GetBody() *DeleteNetworkInterfacesResponseBody
}

type DeleteNetworkInterfacesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNetworkInterfacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNetworkInterfacesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkInterfacesResponse) GoString() string {
	return s.String()
}

func (s *DeleteNetworkInterfacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNetworkInterfacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNetworkInterfacesResponse) GetBody() *DeleteNetworkInterfacesResponseBody {
	return s.Body
}

func (s *DeleteNetworkInterfacesResponse) SetHeaders(v map[string]*string) *DeleteNetworkInterfacesResponse {
	s.Headers = v
	return s
}

func (s *DeleteNetworkInterfacesResponse) SetStatusCode(v int32) *DeleteNetworkInterfacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNetworkInterfacesResponse) SetBody(v *DeleteNetworkInterfacesResponseBody) *DeleteNetworkInterfacesResponse {
	s.Body = v
	return s
}

func (s *DeleteNetworkInterfacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
