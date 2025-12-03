// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiZoneAvailableRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMultiZoneAvailableRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMultiZoneAvailableRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMultiZoneAvailableRegionsResponseBody) *DescribeMultiZoneAvailableRegionsResponse
	GetBody() *DescribeMultiZoneAvailableRegionsResponseBody
}

type DescribeMultiZoneAvailableRegionsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMultiZoneAvailableRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMultiZoneAvailableRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMultiZoneAvailableRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMultiZoneAvailableRegionsResponse) GetBody() *DescribeMultiZoneAvailableRegionsResponseBody {
	return s.Body
}

func (s *DescribeMultiZoneAvailableRegionsResponse) SetHeaders(v map[string]*string) *DescribeMultiZoneAvailableRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponse) SetStatusCode(v int32) *DescribeMultiZoneAvailableRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponse) SetBody(v *DescribeMultiZoneAvailableRegionsResponseBody) *DescribeMultiZoneAvailableRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeMultiZoneAvailableRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
