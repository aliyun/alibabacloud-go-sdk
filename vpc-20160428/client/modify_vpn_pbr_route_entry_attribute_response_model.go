// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnPbrRouteEntryAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpnPbrRouteEntryAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpnPbrRouteEntryAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpnPbrRouteEntryAttributeResponseBody) *ModifyVpnPbrRouteEntryAttributeResponse
	GetBody() *ModifyVpnPbrRouteEntryAttributeResponseBody
}

type ModifyVpnPbrRouteEntryAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpnPbrRouteEntryAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpnPbrRouteEntryAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnPbrRouteEntryAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpnPbrRouteEntryAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpnPbrRouteEntryAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpnPbrRouteEntryAttributeResponse) GetBody() *ModifyVpnPbrRouteEntryAttributeResponseBody {
	return s.Body
}

func (s *ModifyVpnPbrRouteEntryAttributeResponse) SetHeaders(v map[string]*string) *ModifyVpnPbrRouteEntryAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeResponse) SetStatusCode(v int32) *ModifyVpnPbrRouteEntryAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeResponse) SetBody(v *ModifyVpnPbrRouteEntryAttributeResponseBody) *ModifyVpnPbrRouteEntryAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyVpnPbrRouteEntryAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
