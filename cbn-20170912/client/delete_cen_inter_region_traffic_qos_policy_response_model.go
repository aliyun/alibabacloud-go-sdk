// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCenInterRegionTrafficQosPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCenInterRegionTrafficQosPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCenInterRegionTrafficQosPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCenInterRegionTrafficQosPolicyResponseBody) *DeleteCenInterRegionTrafficQosPolicyResponse
	GetBody() *DeleteCenInterRegionTrafficQosPolicyResponseBody
}

type DeleteCenInterRegionTrafficQosPolicyResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCenInterRegionTrafficQosPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCenInterRegionTrafficQosPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCenInterRegionTrafficQosPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteCenInterRegionTrafficQosPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCenInterRegionTrafficQosPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCenInterRegionTrafficQosPolicyResponse) GetBody() *DeleteCenInterRegionTrafficQosPolicyResponseBody {
	return s.Body
}

func (s *DeleteCenInterRegionTrafficQosPolicyResponse) SetHeaders(v map[string]*string) *DeleteCenInterRegionTrafficQosPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyResponse) SetStatusCode(v int32) *DeleteCenInterRegionTrafficQosPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyResponse) SetBody(v *DeleteCenInterRegionTrafficQosPolicyResponseBody) *DeleteCenInterRegionTrafficQosPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteCenInterRegionTrafficQosPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
