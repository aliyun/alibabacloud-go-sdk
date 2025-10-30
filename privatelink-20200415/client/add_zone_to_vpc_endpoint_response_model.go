// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddZoneToVpcEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddZoneToVpcEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddZoneToVpcEndpointResponse
	GetStatusCode() *int32
	SetBody(v *AddZoneToVpcEndpointResponseBody) *AddZoneToVpcEndpointResponse
	GetBody() *AddZoneToVpcEndpointResponseBody
}

type AddZoneToVpcEndpointResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddZoneToVpcEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddZoneToVpcEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s AddZoneToVpcEndpointResponse) GoString() string {
	return s.String()
}

func (s *AddZoneToVpcEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddZoneToVpcEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddZoneToVpcEndpointResponse) GetBody() *AddZoneToVpcEndpointResponseBody {
	return s.Body
}

func (s *AddZoneToVpcEndpointResponse) SetHeaders(v map[string]*string) *AddZoneToVpcEndpointResponse {
	s.Headers = v
	return s
}

func (s *AddZoneToVpcEndpointResponse) SetStatusCode(v int32) *AddZoneToVpcEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *AddZoneToVpcEndpointResponse) SetBody(v *AddZoneToVpcEndpointResponseBody) *AddZoneToVpcEndpointResponse {
	s.Body = v
	return s
}

func (s *AddZoneToVpcEndpointResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
