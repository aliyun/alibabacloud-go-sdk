// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventLogResponseBody) *DescribeEventLogResponse
	GetBody() *DescribeEventLogResponseBody
}

type DescribeEventLogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventLogResponse) GetBody() *DescribeEventLogResponseBody {
	return s.Body
}

func (s *DescribeEventLogResponse) SetHeaders(v map[string]*string) *DescribeEventLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventLogResponse) SetStatusCode(v int32) *DescribeEventLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventLogResponse) SetBody(v *DescribeEventLogResponseBody) *DescribeEventLogResponse {
	s.Body = v
	return s
}

func (s *DescribeEventLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
