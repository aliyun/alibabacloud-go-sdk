// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpnRouteEntryWeightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpnRouteEntryWeightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpnRouteEntryWeightResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpnRouteEntryWeightResponseBody) *ModifyVpnRouteEntryWeightResponse
	GetBody() *ModifyVpnRouteEntryWeightResponseBody
}

type ModifyVpnRouteEntryWeightResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpnRouteEntryWeightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpnRouteEntryWeightResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpnRouteEntryWeightResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpnRouteEntryWeightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpnRouteEntryWeightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpnRouteEntryWeightResponse) GetBody() *ModifyVpnRouteEntryWeightResponseBody {
	return s.Body
}

func (s *ModifyVpnRouteEntryWeightResponse) SetHeaders(v map[string]*string) *ModifyVpnRouteEntryWeightResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpnRouteEntryWeightResponse) SetStatusCode(v int32) *ModifyVpnRouteEntryWeightResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpnRouteEntryWeightResponse) SetBody(v *ModifyVpnRouteEntryWeightResponseBody) *ModifyVpnRouteEntryWeightResponse {
	s.Body = v
	return s
}

func (s *ModifyVpnRouteEntryWeightResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
