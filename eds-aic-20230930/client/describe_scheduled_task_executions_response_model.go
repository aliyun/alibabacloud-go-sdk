// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduledTaskExecutionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScheduledTaskExecutionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScheduledTaskExecutionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScheduledTaskExecutionsResponseBody) *DescribeScheduledTaskExecutionsResponse
	GetBody() *DescribeScheduledTaskExecutionsResponseBody
}

type DescribeScheduledTaskExecutionsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScheduledTaskExecutionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScheduledTaskExecutionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTaskExecutionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTaskExecutionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScheduledTaskExecutionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScheduledTaskExecutionsResponse) GetBody() *DescribeScheduledTaskExecutionsResponseBody {
	return s.Body
}

func (s *DescribeScheduledTaskExecutionsResponse) SetHeaders(v map[string]*string) *DescribeScheduledTaskExecutionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponse) SetStatusCode(v int32) *DescribeScheduledTaskExecutionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponse) SetBody(v *DescribeScheduledTaskExecutionsResponseBody) *DescribeScheduledTaskExecutionsResponse {
	s.Body = v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
