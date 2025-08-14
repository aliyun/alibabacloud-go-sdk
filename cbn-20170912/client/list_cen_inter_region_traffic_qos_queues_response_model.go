// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCenInterRegionTrafficQosQueuesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCenInterRegionTrafficQosQueuesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCenInterRegionTrafficQosQueuesResponse
	GetStatusCode() *int32
	SetBody(v *ListCenInterRegionTrafficQosQueuesResponseBody) *ListCenInterRegionTrafficQosQueuesResponse
	GetBody() *ListCenInterRegionTrafficQosQueuesResponseBody
}

type ListCenInterRegionTrafficQosQueuesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCenInterRegionTrafficQosQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCenInterRegionTrafficQosQueuesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCenInterRegionTrafficQosQueuesResponse) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosQueuesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCenInterRegionTrafficQosQueuesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCenInterRegionTrafficQosQueuesResponse) GetBody() *ListCenInterRegionTrafficQosQueuesResponseBody {
	return s.Body
}

func (s *ListCenInterRegionTrafficQosQueuesResponse) SetHeaders(v map[string]*string) *ListCenInterRegionTrafficQosQueuesResponse {
	s.Headers = v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponse) SetStatusCode(v int32) *ListCenInterRegionTrafficQosQueuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponse) SetBody(v *ListCenInterRegionTrafficQosQueuesResponseBody) *ListCenInterRegionTrafficQosQueuesResponse {
	s.Body = v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponse) Validate() error {
	return dara.Validate(s)
}
