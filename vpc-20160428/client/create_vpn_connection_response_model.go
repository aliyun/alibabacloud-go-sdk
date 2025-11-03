// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpnConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpnConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpnConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpnConnectionResponseBody) *CreateVpnConnectionResponse
	GetBody() *CreateVpnConnectionResponseBody
}

type CreateVpnConnectionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpnConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpnConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpnConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateVpnConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpnConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpnConnectionResponse) GetBody() *CreateVpnConnectionResponseBody {
	return s.Body
}

func (s *CreateVpnConnectionResponse) SetHeaders(v map[string]*string) *CreateVpnConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateVpnConnectionResponse) SetStatusCode(v int32) *CreateVpnConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpnConnectionResponse) SetBody(v *CreateVpnConnectionResponseBody) *CreateVpnConnectionResponse {
	s.Body = v
	return s
}

func (s *CreateVpnConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
