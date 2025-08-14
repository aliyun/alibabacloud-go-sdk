// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCenInterRegionTrafficQosPolicyAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCenInterRegionTrafficQosPolicyAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCenInterRegionTrafficQosPolicyAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody) *UpdateCenInterRegionTrafficQosPolicyAttributeResponse
	GetBody() *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody
}

type UpdateCenInterRegionTrafficQosPolicyAttributeResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCenInterRegionTrafficQosPolicyAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCenInterRegionTrafficQosPolicyAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeResponse) GetBody() *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody {
	return s.Body
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeResponse) SetHeaders(v map[string]*string) *UpdateCenInterRegionTrafficQosPolicyAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeResponse) SetStatusCode(v int32) *UpdateCenInterRegionTrafficQosPolicyAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeResponse) SetBody(v *UpdateCenInterRegionTrafficQosPolicyAttributeResponseBody) *UpdateCenInterRegionTrafficQosPolicyAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateCenInterRegionTrafficQosPolicyAttributeResponse) Validate() error {
	return dara.Validate(s)
}
