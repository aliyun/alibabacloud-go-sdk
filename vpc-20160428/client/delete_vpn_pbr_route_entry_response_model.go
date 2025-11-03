// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnPbrRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpnPbrRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpnPbrRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpnPbrRouteEntryResponseBody) *DeleteVpnPbrRouteEntryResponse
	GetBody() *DeleteVpnPbrRouteEntryResponseBody
}

type DeleteVpnPbrRouteEntryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpnPbrRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpnPbrRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnPbrRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpnPbrRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpnPbrRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpnPbrRouteEntryResponse) GetBody() *DeleteVpnPbrRouteEntryResponseBody {
	return s.Body
}

func (s *DeleteVpnPbrRouteEntryResponse) SetHeaders(v map[string]*string) *DeleteVpnPbrRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpnPbrRouteEntryResponse) SetStatusCode(v int32) *DeleteVpnPbrRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpnPbrRouteEntryResponse) SetBody(v *DeleteVpnPbrRouteEntryResponseBody) *DeleteVpnPbrRouteEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteVpnPbrRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
