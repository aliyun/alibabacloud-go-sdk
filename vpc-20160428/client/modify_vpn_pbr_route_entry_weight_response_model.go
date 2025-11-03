// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnPbrRouteEntryWeightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpnPbrRouteEntryWeightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpnPbrRouteEntryWeightResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpnPbrRouteEntryWeightResponseBody) *ModifyVpnPbrRouteEntryWeightResponse
	GetBody() *ModifyVpnPbrRouteEntryWeightResponseBody
}

type ModifyVpnPbrRouteEntryWeightResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpnPbrRouteEntryWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpnPbrRouteEntryWeightResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnPbrRouteEntryWeightResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpnPbrRouteEntryWeightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpnPbrRouteEntryWeightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpnPbrRouteEntryWeightResponse) GetBody() *ModifyVpnPbrRouteEntryWeightResponseBody {
	return s.Body
}

func (s *ModifyVpnPbrRouteEntryWeightResponse) SetHeaders(v map[string]*string) *ModifyVpnPbrRouteEntryWeightResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightResponse) SetStatusCode(v int32) *ModifyVpnPbrRouteEntryWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightResponse) SetBody(v *ModifyVpnPbrRouteEntryWeightResponseBody) *ModifyVpnPbrRouteEntryWeightResponse {
	s.Body = v
	return s
}

func (s *ModifyVpnPbrRouteEntryWeightResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
