// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpnRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpnRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpnRouteEntryResponseBody) *DeleteVpnRouteEntryResponse
	GetBody() *DeleteVpnRouteEntryResponseBody
}

type DeleteVpnRouteEntryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpnRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpnRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpnRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpnRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpnRouteEntryResponse) GetBody() *DeleteVpnRouteEntryResponseBody {
	return s.Body
}

func (s *DeleteVpnRouteEntryResponse) SetHeaders(v map[string]*string) *DeleteVpnRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpnRouteEntryResponse) SetStatusCode(v int32) *DeleteVpnRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpnRouteEntryResponse) SetBody(v *DeleteVpnRouteEntryResponseBody) *DeleteVpnRouteEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteVpnRouteEntryResponse) Validate() error {
	return dara.Validate(s)
}
