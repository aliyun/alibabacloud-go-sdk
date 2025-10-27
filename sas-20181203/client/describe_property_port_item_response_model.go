// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyPortItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyPortItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyPortItemResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyPortItemResponseBody) *DescribePropertyPortItemResponse
	GetBody() *DescribePropertyPortItemResponseBody
}

type DescribePropertyPortItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyPortItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyPortItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyPortItemResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyPortItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyPortItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyPortItemResponse) GetBody() *DescribePropertyPortItemResponseBody {
	return s.Body
}

func (s *DescribePropertyPortItemResponse) SetHeaders(v map[string]*string) *DescribePropertyPortItemResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyPortItemResponse) SetStatusCode(v int32) *DescribePropertyPortItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyPortItemResponse) SetBody(v *DescribePropertyPortItemResponseBody) *DescribePropertyPortItemResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyPortItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
