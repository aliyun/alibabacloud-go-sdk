// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnRouteEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpnRouteEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpnRouteEntryResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpnRouteEntryResponseBody) *CreateVpnRouteEntryResponse
	GetBody() *CreateVpnRouteEntryResponseBody
}

type CreateVpnRouteEntryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpnRouteEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpnRouteEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnRouteEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateVpnRouteEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpnRouteEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpnRouteEntryResponse) GetBody() *CreateVpnRouteEntryResponseBody {
	return s.Body
}

func (s *CreateVpnRouteEntryResponse) SetHeaders(v map[string]*string) *CreateVpnRouteEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateVpnRouteEntryResponse) SetStatusCode(v int32) *CreateVpnRouteEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpnRouteEntryResponse) SetBody(v *CreateVpnRouteEntryResponseBody) *CreateVpnRouteEntryResponse {
	s.Body = v
	return s
}

func (s *CreateVpnRouteEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
