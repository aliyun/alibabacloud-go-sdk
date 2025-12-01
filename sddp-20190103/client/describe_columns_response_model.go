// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeColumnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeColumnsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeColumnsResponseBody) *DescribeColumnsResponse
	GetBody() *DescribeColumnsResponseBody
}

type DescribeColumnsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeColumnsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnsResponse) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeColumnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeColumnsResponse) GetBody() *DescribeColumnsResponseBody {
	return s.Body
}

func (s *DescribeColumnsResponse) SetHeaders(v map[string]*string) *DescribeColumnsResponse {
	s.Headers = v
	return s
}

func (s *DescribeColumnsResponse) SetStatusCode(v int32) *DescribeColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeColumnsResponse) SetBody(v *DescribeColumnsResponseBody) *DescribeColumnsResponse {
	s.Body = v
	return s
}

func (s *DescribeColumnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
