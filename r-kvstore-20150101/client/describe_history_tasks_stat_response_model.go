// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHistoryTasksStatResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHistoryTasksStatResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHistoryTasksStatResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHistoryTasksStatResponseBody) *DescribeHistoryTasksStatResponse
	GetBody() *DescribeHistoryTasksStatResponseBody
}

type DescribeHistoryTasksStatResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHistoryTasksStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHistoryTasksStatResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHistoryTasksStatResponse) GoString() string {
	return s.String()
}

func (s *DescribeHistoryTasksStatResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHistoryTasksStatResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHistoryTasksStatResponse) GetBody() *DescribeHistoryTasksStatResponseBody {
	return s.Body
}

func (s *DescribeHistoryTasksStatResponse) SetHeaders(v map[string]*string) *DescribeHistoryTasksStatResponse {
	s.Headers = v
	return s
}

func (s *DescribeHistoryTasksStatResponse) SetStatusCode(v int32) *DescribeHistoryTasksStatResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHistoryTasksStatResponse) SetBody(v *DescribeHistoryTasksStatResponseBody) *DescribeHistoryTasksStatResponse {
	s.Body = v
	return s
}

func (s *DescribeHistoryTasksStatResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
