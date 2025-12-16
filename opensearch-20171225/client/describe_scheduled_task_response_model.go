// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduledTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScheduledTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScheduledTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScheduledTaskResponseBody) *DescribeScheduledTaskResponse
	GetBody() *DescribeScheduledTaskResponseBody
}

type DescribeScheduledTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScheduledTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScheduledTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScheduledTaskResponse) GetBody() *DescribeScheduledTaskResponseBody {
	return s.Body
}

func (s *DescribeScheduledTaskResponse) SetHeaders(v map[string]*string) *DescribeScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduledTaskResponse) SetStatusCode(v int32) *DescribeScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScheduledTaskResponse) SetBody(v *DescribeScheduledTaskResponseBody) *DescribeScheduledTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeScheduledTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
