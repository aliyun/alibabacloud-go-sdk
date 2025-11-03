// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnConnectionAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpnConnectionAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpnConnectionAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpnConnectionAttributeResponseBody) *ModifyVpnConnectionAttributeResponse
	GetBody() *ModifyVpnConnectionAttributeResponseBody
}

type ModifyVpnConnectionAttributeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpnConnectionAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpnConnectionAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnConnectionAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpnConnectionAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpnConnectionAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpnConnectionAttributeResponse) GetBody() *ModifyVpnConnectionAttributeResponseBody {
	return s.Body
}

func (s *ModifyVpnConnectionAttributeResponse) SetHeaders(v map[string]*string) *ModifyVpnConnectionAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpnConnectionAttributeResponse) SetStatusCode(v int32) *ModifyVpnConnectionAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpnConnectionAttributeResponse) SetBody(v *ModifyVpnConnectionAttributeResponseBody) *ModifyVpnConnectionAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyVpnConnectionAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
