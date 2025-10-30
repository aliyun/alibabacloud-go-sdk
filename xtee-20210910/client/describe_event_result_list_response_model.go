// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventResultListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventResultListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventResultListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventResultListResponseBody) *DescribeEventResultListResponse
	GetBody() *DescribeEventResultListResponseBody
}

type DescribeEventResultListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventResultListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventResultListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventResultListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventResultListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventResultListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventResultListResponse) GetBody() *DescribeEventResultListResponseBody {
	return s.Body
}

func (s *DescribeEventResultListResponse) SetHeaders(v map[string]*string) *DescribeEventResultListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventResultListResponse) SetStatusCode(v int32) *DescribeEventResultListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventResultListResponse) SetBody(v *DescribeEventResultListResponseBody) *DescribeEventResultListResponse {
	s.Body = v
	return s
}

func (s *DescribeEventResultListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
