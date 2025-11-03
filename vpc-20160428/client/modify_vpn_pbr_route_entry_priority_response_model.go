// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnPbrRouteEntryPriorityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpnPbrRouteEntryPriorityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpnPbrRouteEntryPriorityResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpnPbrRouteEntryPriorityResponseBody) *ModifyVpnPbrRouteEntryPriorityResponse
	GetBody() *ModifyVpnPbrRouteEntryPriorityResponseBody
}

type ModifyVpnPbrRouteEntryPriorityResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpnPbrRouteEntryPriorityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpnPbrRouteEntryPriorityResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnPbrRouteEntryPriorityResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpnPbrRouteEntryPriorityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpnPbrRouteEntryPriorityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpnPbrRouteEntryPriorityResponse) GetBody() *ModifyVpnPbrRouteEntryPriorityResponseBody {
	return s.Body
}

func (s *ModifyVpnPbrRouteEntryPriorityResponse) SetHeaders(v map[string]*string) *ModifyVpnPbrRouteEntryPriorityResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityResponse) SetStatusCode(v int32) *ModifyVpnPbrRouteEntryPriorityResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityResponse) SetBody(v *ModifyVpnPbrRouteEntryPriorityResponseBody) *ModifyVpnPbrRouteEntryPriorityResponse {
	s.Body = v
	return s
}

func (s *ModifyVpnPbrRouteEntryPriorityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
