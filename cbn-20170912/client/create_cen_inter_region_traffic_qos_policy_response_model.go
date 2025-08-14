// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCenInterRegionTrafficQosPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCenInterRegionTrafficQosPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCenInterRegionTrafficQosPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateCenInterRegionTrafficQosPolicyResponseBody) *CreateCenInterRegionTrafficQosPolicyResponse
	GetBody() *CreateCenInterRegionTrafficQosPolicyResponseBody
}

type CreateCenInterRegionTrafficQosPolicyResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCenInterRegionTrafficQosPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCenInterRegionTrafficQosPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCenInterRegionTrafficQosPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateCenInterRegionTrafficQosPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCenInterRegionTrafficQosPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCenInterRegionTrafficQosPolicyResponse) GetBody() *CreateCenInterRegionTrafficQosPolicyResponseBody {
	return s.Body
}

func (s *CreateCenInterRegionTrafficQosPolicyResponse) SetHeaders(v map[string]*string) *CreateCenInterRegionTrafficQosPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyResponse) SetStatusCode(v int32) *CreateCenInterRegionTrafficQosPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyResponse) SetBody(v *CreateCenInterRegionTrafficQosPolicyResponseBody) *CreateCenInterRegionTrafficQosPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateCenInterRegionTrafficQosPolicyResponse) Validate() error {
	return dara.Validate(s)
}
