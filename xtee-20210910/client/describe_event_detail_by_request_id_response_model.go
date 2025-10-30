// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventDetailByRequestIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventDetailByRequestIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventDetailByRequestIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventDetailByRequestIdResponseBody) *DescribeEventDetailByRequestIdResponse
	GetBody() *DescribeEventDetailByRequestIdResponseBody
}

type DescribeEventDetailByRequestIdResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventDetailByRequestIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventDetailByRequestIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventDetailByRequestIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventDetailByRequestIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventDetailByRequestIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventDetailByRequestIdResponse) GetBody() *DescribeEventDetailByRequestIdResponseBody {
	return s.Body
}

func (s *DescribeEventDetailByRequestIdResponse) SetHeaders(v map[string]*string) *DescribeEventDetailByRequestIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventDetailByRequestIdResponse) SetStatusCode(v int32) *DescribeEventDetailByRequestIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventDetailByRequestIdResponse) SetBody(v *DescribeEventDetailByRequestIdResponseBody) *DescribeEventDetailByRequestIdResponse {
	s.Body = v
	return s
}

func (s *DescribeEventDetailByRequestIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
