// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceLogResponseBody) *DescribeInstanceLogResponse
	GetBody() *DescribeInstanceLogResponseBody
}

type DescribeInstanceLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceLogResponse) GetBody() *DescribeInstanceLogResponseBody {
	return s.Body
}

func (s *DescribeInstanceLogResponse) SetHeaders(v map[string]*string) *DescribeInstanceLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceLogResponse) SetStatusCode(v int32) *DescribeInstanceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceLogResponse) SetBody(v *DescribeInstanceLogResponseBody) *DescribeInstanceLogResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
