// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyCronItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePropertyCronItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePropertyCronItemResponse
	GetStatusCode() *int32
	SetBody(v *DescribePropertyCronItemResponseBody) *DescribePropertyCronItemResponse
	GetBody() *DescribePropertyCronItemResponseBody
}

type DescribePropertyCronItemResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePropertyCronItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePropertyCronItemResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyCronItemResponse) GoString() string {
	return s.String()
}

func (s *DescribePropertyCronItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePropertyCronItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePropertyCronItemResponse) GetBody() *DescribePropertyCronItemResponseBody {
	return s.Body
}

func (s *DescribePropertyCronItemResponse) SetHeaders(v map[string]*string) *DescribePropertyCronItemResponse {
	s.Headers = v
	return s
}

func (s *DescribePropertyCronItemResponse) SetStatusCode(v int32) *DescribePropertyCronItemResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePropertyCronItemResponse) SetBody(v *DescribePropertyCronItemResponseBody) *DescribePropertyCronItemResponse {
	s.Body = v
	return s
}

func (s *DescribePropertyCronItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
