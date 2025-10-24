// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceLogResponseBody) *DescribeServiceLogResponse
	GetBody() *DescribeServiceLogResponseBody
}

type DescribeServiceLogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceLogResponse) GetBody() *DescribeServiceLogResponseBody {
	return s.Body
}

func (s *DescribeServiceLogResponse) SetHeaders(v map[string]*string) *DescribeServiceLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceLogResponse) SetStatusCode(v int32) *DescribeServiceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceLogResponse) SetBody(v *DescribeServiceLogResponseBody) *DescribeServiceLogResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
