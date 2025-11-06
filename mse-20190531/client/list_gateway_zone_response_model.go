// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayZoneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGatewayZoneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGatewayZoneResponse
	GetStatusCode() *int32
	SetBody(v *ListGatewayZoneResponseBody) *ListGatewayZoneResponse
	GetBody() *ListGatewayZoneResponseBody
}

type ListGatewayZoneResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGatewayZoneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGatewayZoneResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayZoneResponse) GoString() string {
	return s.String()
}

func (s *ListGatewayZoneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGatewayZoneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGatewayZoneResponse) GetBody() *ListGatewayZoneResponseBody {
	return s.Body
}

func (s *ListGatewayZoneResponse) SetHeaders(v map[string]*string) *ListGatewayZoneResponse {
	s.Headers = v
	return s
}

func (s *ListGatewayZoneResponse) SetStatusCode(v int32) *ListGatewayZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGatewayZoneResponse) SetBody(v *ListGatewayZoneResponseBody) *ListGatewayZoneResponse {
	s.Body = v
	return s
}

func (s *ListGatewayZoneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
