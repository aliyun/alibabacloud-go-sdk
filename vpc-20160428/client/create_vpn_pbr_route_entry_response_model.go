// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnPbrRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpnPbrRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpnPbrRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpnPbrRouteEntryResponseBody) *CreateVpnPbrRouteEntryResponse
	GetBody() *CreateVpnPbrRouteEntryResponseBody
}

type CreateVpnPbrRouteEntryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpnPbrRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpnPbrRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnPbrRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateVpnPbrRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpnPbrRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpnPbrRouteEntryResponse) GetBody() *CreateVpnPbrRouteEntryResponseBody {
	return s.Body
}

func (s *CreateVpnPbrRouteEntryResponse) SetHeaders(v map[string]*string) *CreateVpnPbrRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateVpnPbrRouteEntryResponse) SetStatusCode(v int32) *CreateVpnPbrRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpnPbrRouteEntryResponse) SetBody(v *CreateVpnPbrRouteEntryResponseBody) *CreateVpnPbrRouteEntryResponse {
	s.Body = v
	return s
}

func (s *CreateVpnPbrRouteEntryResponse) Validate() error {
	return dara.Validate(s)
}
