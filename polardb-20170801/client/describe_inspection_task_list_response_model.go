// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionTaskListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInspectionTaskListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInspectionTaskListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInspectionTaskListResponseBody) *DescribeInspectionTaskListResponse
	GetBody() *DescribeInspectionTaskListResponseBody
}

type DescribeInspectionTaskListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInspectionTaskListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInspectionTaskListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTaskListResponse) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTaskListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInspectionTaskListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInspectionTaskListResponse) GetBody() *DescribeInspectionTaskListResponseBody {
	return s.Body
}

func (s *DescribeInspectionTaskListResponse) SetHeaders(v map[string]*string) *DescribeInspectionTaskListResponse {
	s.Headers = v
	return s
}

func (s *DescribeInspectionTaskListResponse) SetStatusCode(v int32) *DescribeInspectionTaskListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInspectionTaskListResponse) SetBody(v *DescribeInspectionTaskListResponseBody) *DescribeInspectionTaskListResponse {
	s.Body = v
	return s
}

func (s *DescribeInspectionTaskListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
