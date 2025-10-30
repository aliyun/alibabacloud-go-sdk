// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventLogPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventLogPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventLogPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventLogPageResponseBody) *DescribeEventLogPageResponse
	GetBody() *DescribeEventLogPageResponseBody
}

type DescribeEventLogPageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventLogPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventLogPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventLogPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventLogPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventLogPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventLogPageResponse) GetBody() *DescribeEventLogPageResponseBody {
	return s.Body
}

func (s *DescribeEventLogPageResponse) SetHeaders(v map[string]*string) *DescribeEventLogPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventLogPageResponse) SetStatusCode(v int32) *DescribeEventLogPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventLogPageResponse) SetBody(v *DescribeEventLogPageResponseBody) *DescribeEventLogPageResponse {
	s.Body = v
	return s
}

func (s *DescribeEventLogPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
