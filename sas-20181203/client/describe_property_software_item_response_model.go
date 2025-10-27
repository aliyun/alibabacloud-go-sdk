// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertySoftwareItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertySoftwareItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertySoftwareItemResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertySoftwareItemResponseBody) *DescribePropertySoftwareItemResponse
	GetBody() *DescribePropertySoftwareItemResponseBody
}

type DescribePropertySoftwareItemResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertySoftwareItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertySoftwareItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertySoftwareItemResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertySoftwareItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertySoftwareItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertySoftwareItemResponse) GetBody() *DescribePropertySoftwareItemResponseBody {
	return s.Body
}

func (s *DescribePropertySoftwareItemResponse) SetHeaders(v map[string]*string) *DescribePropertySoftwareItemResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertySoftwareItemResponse) SetStatusCode(v int32) *DescribePropertySoftwareItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertySoftwareItemResponse) SetBody(v *DescribePropertySoftwareItemResponseBody) *DescribePropertySoftwareItemResponse {
	s.Body = v
	return s
}

func (s *DescribePropertySoftwareItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
