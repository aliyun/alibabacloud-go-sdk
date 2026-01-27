// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImportTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImportTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImportTaskResponseBody) *DescribeImportTaskResponse
	GetBody() *DescribeImportTaskResponseBody
}

type DescribeImportTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImportTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImportTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeImportTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImportTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImportTaskResponse) GetBody() *DescribeImportTaskResponseBody {
	return s.Body
}

func (s *DescribeImportTaskResponse) SetHeaders(v map[string]*string) *DescribeImportTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeImportTaskResponse) SetStatusCode(v int32) *DescribeImportTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImportTaskResponse) SetBody(v *DescribeImportTaskResponseBody) *DescribeImportTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeImportTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
