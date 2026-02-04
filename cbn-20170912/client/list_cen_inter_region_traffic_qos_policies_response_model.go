// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCenInterRegionTrafficQosPoliciesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCenInterRegionTrafficQosPoliciesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCenInterRegionTrafficQosPoliciesResponse
	GetStatusCode() *int32
	SetBody(v *ListCenInterRegionTrafficQosPoliciesResponseBody) *ListCenInterRegionTrafficQosPoliciesResponse
	GetBody() *ListCenInterRegionTrafficQosPoliciesResponseBody
}

type ListCenInterRegionTrafficQosPoliciesResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCenInterRegionTrafficQosPoliciesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCenInterRegionTrafficQosPoliciesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCenInterRegionTrafficQosPoliciesResponse) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosPoliciesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCenInterRegionTrafficQosPoliciesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCenInterRegionTrafficQosPoliciesResponse) GetBody() *ListCenInterRegionTrafficQosPoliciesResponseBody {
	return s.Body
}

func (s *ListCenInterRegionTrafficQosPoliciesResponse) SetHeaders(v map[string]*string) *ListCenInterRegionTrafficQosPoliciesResponse {
	s.Headers = v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponse) SetStatusCode(v int32) *ListCenInterRegionTrafficQosPoliciesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponse) SetBody(v *ListCenInterRegionTrafficQosPoliciesResponseBody) *ListCenInterRegionTrafficQosPoliciesResponse {
	s.Body = v
	return s
}

func (s *ListCenInterRegionTrafficQosPoliciesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
