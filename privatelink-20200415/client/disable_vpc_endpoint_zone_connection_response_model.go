// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableVpcEndpointZoneConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableVpcEndpointZoneConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableVpcEndpointZoneConnectionResponse
	GetStatusCode() *int32
	SetBody(v *DisableVpcEndpointZoneConnectionResponseBody) *DisableVpcEndpointZoneConnectionResponse
	GetBody() *DisableVpcEndpointZoneConnectionResponseBody
}

type DisableVpcEndpointZoneConnectionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableVpcEndpointZoneConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableVpcEndpointZoneConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableVpcEndpointZoneConnectionResponse) GoString() string {
	return s.String()
}

func (s *DisableVpcEndpointZoneConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableVpcEndpointZoneConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableVpcEndpointZoneConnectionResponse) GetBody() *DisableVpcEndpointZoneConnectionResponseBody {
	return s.Body
}

func (s *DisableVpcEndpointZoneConnectionResponse) SetHeaders(v map[string]*string) *DisableVpcEndpointZoneConnectionResponse {
	s.Headers = v
	return s
}

func (s *DisableVpcEndpointZoneConnectionResponse) SetStatusCode(v int32) *DisableVpcEndpointZoneConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableVpcEndpointZoneConnectionResponse) SetBody(v *DisableVpcEndpointZoneConnectionResponseBody) *DisableVpcEndpointZoneConnectionResponse {
	s.Body = v
	return s
}

func (s *DisableVpcEndpointZoneConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
