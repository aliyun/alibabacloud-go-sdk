// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpnConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpnConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpnConnectionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpnConnectionResponseBody) *DeleteVpnConnectionResponse
	GetBody() *DeleteVpnConnectionResponseBody
}

type DeleteVpnConnectionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpnConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpnConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpnConnectionResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpnConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpnConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpnConnectionResponse) GetBody() *DeleteVpnConnectionResponseBody {
	return s.Body
}

func (s *DeleteVpnConnectionResponse) SetHeaders(v map[string]*string) *DeleteVpnConnectionResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpnConnectionResponse) SetStatusCode(v int32) *DeleteVpnConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpnConnectionResponse) SetBody(v *DeleteVpnConnectionResponseBody) *DeleteVpnConnectionResponse {
	s.Body = v
	return s
}

func (s *DeleteVpnConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
