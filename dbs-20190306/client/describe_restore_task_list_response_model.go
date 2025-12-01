// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRestoreTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRestoreTaskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRestoreTaskListResponseBody) *DescribeRestoreTaskListResponse
	GetBody() *DescribeRestoreTaskListResponseBody
}

type DescribeRestoreTaskListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRestoreTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRestoreTaskListResponse) GetBody() *DescribeRestoreTaskListResponseBody {
	return s.Body
}

func (s *DescribeRestoreTaskListResponse) SetHeaders(v map[string]*string) *DescribeRestoreTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreTaskListResponse) SetStatusCode(v int32) *DescribeRestoreTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreTaskListResponse) SetBody(v *DescribeRestoreTaskListResponseBody) *DescribeRestoreTaskListResponse {
	s.Body = v
	return s
}

func (s *DescribeRestoreTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
