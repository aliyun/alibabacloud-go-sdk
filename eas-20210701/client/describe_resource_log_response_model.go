// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResourceLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResourceLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResourceLogResponseBody) *DescribeResourceLogResponse
	GetBody() *DescribeResourceLogResponseBody
}

type DescribeResourceLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResourceLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResourceLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResourceLogResponse) GetBody() *DescribeResourceLogResponseBody {
	return s.Body
}

func (s *DescribeResourceLogResponse) SetHeaders(v map[string]*string) *DescribeResourceLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceLogResponse) SetStatusCode(v int32) *DescribeResourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceLogResponse) SetBody(v *DescribeResourceLogResponseBody) *DescribeResourceLogResponse {
	s.Body = v
	return s
}

func (s *DescribeResourceLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
