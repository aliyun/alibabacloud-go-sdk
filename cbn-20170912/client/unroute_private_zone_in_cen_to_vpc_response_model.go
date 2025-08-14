// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnroutePrivateZoneInCenToVpcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnroutePrivateZoneInCenToVpcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnroutePrivateZoneInCenToVpcResponse
	GetStatusCode() *int32
	SetBody(v *UnroutePrivateZoneInCenToVpcResponseBody) *UnroutePrivateZoneInCenToVpcResponse
	GetBody() *UnroutePrivateZoneInCenToVpcResponseBody
}

type UnroutePrivateZoneInCenToVpcResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnroutePrivateZoneInCenToVpcResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnroutePrivateZoneInCenToVpcResponse) String() string {
	return dara.Prettify(s)
}

func (s UnroutePrivateZoneInCenToVpcResponse) GoString() string {
	return s.String()
}

func (s *UnroutePrivateZoneInCenToVpcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnroutePrivateZoneInCenToVpcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnroutePrivateZoneInCenToVpcResponse) GetBody() *UnroutePrivateZoneInCenToVpcResponseBody {
	return s.Body
}

func (s *UnroutePrivateZoneInCenToVpcResponse) SetHeaders(v map[string]*string) *UnroutePrivateZoneInCenToVpcResponse {
	s.Headers = v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcResponse) SetStatusCode(v int32) *UnroutePrivateZoneInCenToVpcResponse {
	s.StatusCode = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcResponse) SetBody(v *UnroutePrivateZoneInCenToVpcResponseBody) *UnroutePrivateZoneInCenToVpcResponse {
	s.Body = v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcResponse) Validate() error {
	return dara.Validate(s)
}
