// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadTasksRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLoadTasksRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLoadTasksRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLoadTasksRecordsResponseBody) *DescribeLoadTasksRecordsResponse
	GetBody() *DescribeLoadTasksRecordsResponseBody
}

type DescribeLoadTasksRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLoadTasksRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLoadTasksRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadTasksRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLoadTasksRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLoadTasksRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLoadTasksRecordsResponse) GetBody() *DescribeLoadTasksRecordsResponseBody {
	return s.Body
}

func (s *DescribeLoadTasksRecordsResponse) SetHeaders(v map[string]*string) *DescribeLoadTasksRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLoadTasksRecordsResponse) SetStatusCode(v int32) *DescribeLoadTasksRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLoadTasksRecordsResponse) SetBody(v *DescribeLoadTasksRecordsResponseBody) *DescribeLoadTasksRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeLoadTasksRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
