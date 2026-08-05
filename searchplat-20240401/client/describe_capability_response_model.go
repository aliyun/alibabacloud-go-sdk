// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapabilityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCapabilityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCapabilityResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCapabilityResponseBody) *DescribeCapabilityResponse
	GetBody() *DescribeCapabilityResponseBody
}

type DescribeCapabilityResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCapabilityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCapabilityResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapabilityResponse) GoString() string {
	return s.String()
}

func (s *DescribeCapabilityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCapabilityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCapabilityResponse) GetBody() *DescribeCapabilityResponseBody {
	return s.Body
}

func (s *DescribeCapabilityResponse) SetHeaders(v map[string]*string) *DescribeCapabilityResponse {
	s.Headers = v
	return s
}

func (s *DescribeCapabilityResponse) SetStatusCode(v int32) *DescribeCapabilityResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCapabilityResponse) SetBody(v *DescribeCapabilityResponseBody) *DescribeCapabilityResponse {
	s.Body = v
	return s
}

func (s *DescribeCapabilityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
