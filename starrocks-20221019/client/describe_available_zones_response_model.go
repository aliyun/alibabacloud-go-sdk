// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvailableZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvailableZonesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvailableZonesResponseBody) *DescribeAvailableZonesResponse
	GetBody() *DescribeAvailableZonesResponseBody
}

type DescribeAvailableZonesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvailableZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvailableZonesResponse) GetBody() *DescribeAvailableZonesResponseBody {
	return s.Body
}

func (s *DescribeAvailableZonesResponse) SetHeaders(v map[string]*string) *DescribeAvailableZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableZonesResponse) SetStatusCode(v int32) *DescribeAvailableZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableZonesResponse) SetBody(v *DescribeAvailableZonesResponseBody) *DescribeAvailableZonesResponse {
	s.Body = v
	return s
}

func (s *DescribeAvailableZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
