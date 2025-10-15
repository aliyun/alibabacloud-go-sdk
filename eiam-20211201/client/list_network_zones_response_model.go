// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetworkZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetworkZonesResponse
	GetStatusCode() *int32
	SetBody(v *ListNetworkZonesResponseBody) *ListNetworkZonesResponse
	GetBody() *ListNetworkZonesResponseBody
}

type ListNetworkZonesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkZonesResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetworkZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetworkZonesResponse) GetBody() *ListNetworkZonesResponseBody {
	return s.Body
}

func (s *ListNetworkZonesResponse) SetHeaders(v map[string]*string) *ListNetworkZonesResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkZonesResponse) SetStatusCode(v int32) *ListNetworkZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkZonesResponse) SetBody(v *ListNetworkZonesResponseBody) *ListNetworkZonesResponse {
	s.Body = v
	return s
}

func (s *ListNetworkZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
