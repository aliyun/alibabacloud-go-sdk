// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEvaluateAndImportTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEvaluateAndImportTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEvaluateAndImportTasksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEvaluateAndImportTasksResponseBody) *DescribeEvaluateAndImportTasksResponse
	GetBody() *DescribeEvaluateAndImportTasksResponseBody
}

type DescribeEvaluateAndImportTasksResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEvaluateAndImportTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEvaluateAndImportTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEvaluateAndImportTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateAndImportTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEvaluateAndImportTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEvaluateAndImportTasksResponse) GetBody() *DescribeEvaluateAndImportTasksResponseBody {
	return s.Body
}

func (s *DescribeEvaluateAndImportTasksResponse) SetHeaders(v map[string]*string) *DescribeEvaluateAndImportTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponse) SetStatusCode(v int32) *DescribeEvaluateAndImportTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponse) SetBody(v *DescribeEvaluateAndImportTasksResponseBody) *DescribeEvaluateAndImportTasksResponse {
	s.Body = v
	return s
}

func (s *DescribeEvaluateAndImportTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
