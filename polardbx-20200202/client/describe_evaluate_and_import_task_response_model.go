// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEvaluateAndImportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEvaluateAndImportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEvaluateAndImportTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEvaluateAndImportTaskResponseBody) *DescribeEvaluateAndImportTaskResponse
	GetBody() *DescribeEvaluateAndImportTaskResponseBody
}

type DescribeEvaluateAndImportTaskResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEvaluateAndImportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEvaluateAndImportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEvaluateAndImportTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeEvaluateAndImportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEvaluateAndImportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEvaluateAndImportTaskResponse) GetBody() *DescribeEvaluateAndImportTaskResponseBody {
	return s.Body
}

func (s *DescribeEvaluateAndImportTaskResponse) SetHeaders(v map[string]*string) *DescribeEvaluateAndImportTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponse) SetStatusCode(v int32) *DescribeEvaluateAndImportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponse) SetBody(v *DescribeEvaluateAndImportTaskResponseBody) *DescribeEvaluateAndImportTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeEvaluateAndImportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
