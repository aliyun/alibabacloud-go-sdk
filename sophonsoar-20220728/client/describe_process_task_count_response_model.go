// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProcessTaskCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProcessTaskCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProcessTaskCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProcessTaskCountResponseBody) *DescribeProcessTaskCountResponse
	GetBody() *DescribeProcessTaskCountResponseBody
}

type DescribeProcessTaskCountResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProcessTaskCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProcessTaskCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProcessTaskCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeProcessTaskCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProcessTaskCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProcessTaskCountResponse) GetBody() *DescribeProcessTaskCountResponseBody {
	return s.Body
}

func (s *DescribeProcessTaskCountResponse) SetHeaders(v map[string]*string) *DescribeProcessTaskCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeProcessTaskCountResponse) SetStatusCode(v int32) *DescribeProcessTaskCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProcessTaskCountResponse) SetBody(v *DescribeProcessTaskCountResponseBody) *DescribeProcessTaskCountResponse {
	s.Body = v
	return s
}

func (s *DescribeProcessTaskCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
