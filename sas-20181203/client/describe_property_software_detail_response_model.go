// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertySoftwareDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertySoftwareDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertySoftwareDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertySoftwareDetailResponseBody) *DescribePropertySoftwareDetailResponse
	GetBody() *DescribePropertySoftwareDetailResponseBody
}

type DescribePropertySoftwareDetailResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertySoftwareDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertySoftwareDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertySoftwareDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertySoftwareDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertySoftwareDetailResponse) GetBody() *DescribePropertySoftwareDetailResponseBody {
	return s.Body
}

func (s *DescribePropertySoftwareDetailResponse) SetHeaders(v map[string]*string) *DescribePropertySoftwareDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertySoftwareDetailResponse) SetStatusCode(v int32) *DescribePropertySoftwareDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertySoftwareDetailResponse) SetBody(v *DescribePropertySoftwareDetailResponseBody) *DescribePropertySoftwareDetailResponse {
	s.Body = v
	return s
}

func (s *DescribePropertySoftwareDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
