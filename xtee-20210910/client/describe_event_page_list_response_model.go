// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventPageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventPageListResponseBody) *DescribeEventPageListResponse
	GetBody() *DescribeEventPageListResponseBody
}

type DescribeEventPageListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventPageListResponse) GetBody() *DescribeEventPageListResponseBody {
	return s.Body
}

func (s *DescribeEventPageListResponse) SetHeaders(v map[string]*string) *DescribeEventPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventPageListResponse) SetStatusCode(v int32) *DescribeEventPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventPageListResponse) SetBody(v *DescribeEventPageListResponseBody) *DescribeEventPageListResponse {
	s.Body = v
	return s
}

func (s *DescribeEventPageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
