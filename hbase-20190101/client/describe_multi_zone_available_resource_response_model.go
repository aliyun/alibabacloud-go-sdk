// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultiZoneAvailableResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMultiZoneAvailableResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMultiZoneAvailableResourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMultiZoneAvailableResourceResponseBody) *DescribeMultiZoneAvailableResourceResponse
	GetBody() *DescribeMultiZoneAvailableResourceResponseBody
}

type DescribeMultiZoneAvailableResourceResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMultiZoneAvailableResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMultiZoneAvailableResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultiZoneAvailableResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultiZoneAvailableResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMultiZoneAvailableResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMultiZoneAvailableResourceResponse) GetBody() *DescribeMultiZoneAvailableResourceResponseBody {
	return s.Body
}

func (s *DescribeMultiZoneAvailableResourceResponse) SetHeaders(v map[string]*string) *DescribeMultiZoneAvailableResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponse) SetStatusCode(v int32) *DescribeMultiZoneAvailableResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponse) SetBody(v *DescribeMultiZoneAvailableResourceResponseBody) *DescribeMultiZoneAvailableResourceResponse {
	s.Body = v
	return s
}

func (s *DescribeMultiZoneAvailableResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
