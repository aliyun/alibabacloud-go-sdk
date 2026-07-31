// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProcessListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProcessListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProcessListResponseBody) *DescribeProcessListResponse
	GetBody() *DescribeProcessListResponseBody
}

type DescribeProcessListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProcessListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProcessListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessListResponse) GoString() string {
	return s.String()
}

func (s *DescribeProcessListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProcessListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProcessListResponse) GetBody() *DescribeProcessListResponseBody {
	return s.Body
}

func (s *DescribeProcessListResponse) SetHeaders(v map[string]*string) *DescribeProcessListResponse {
	s.Headers = v
	return s
}

func (s *DescribeProcessListResponse) SetStatusCode(v int32) *DescribeProcessListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProcessListResponse) SetBody(v *DescribeProcessListResponseBody) *DescribeProcessListResponse {
	s.Body = v
	return s
}

func (s *DescribeProcessListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
