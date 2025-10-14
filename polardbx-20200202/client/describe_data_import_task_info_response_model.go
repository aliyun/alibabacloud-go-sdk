// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataImportTaskInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataImportTaskInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataImportTaskInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataImportTaskInfoResponseBody) *DescribeDataImportTaskInfoResponse
	GetBody() *DescribeDataImportTaskInfoResponseBody
}

type DescribeDataImportTaskInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataImportTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataImportTaskInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataImportTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataImportTaskInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataImportTaskInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataImportTaskInfoResponse) GetBody() *DescribeDataImportTaskInfoResponseBody {
	return s.Body
}

func (s *DescribeDataImportTaskInfoResponse) SetHeaders(v map[string]*string) *DescribeDataImportTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataImportTaskInfoResponse) SetStatusCode(v int32) *DescribeDataImportTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataImportTaskInfoResponse) SetBody(v *DescribeDataImportTaskInfoResponseBody) *DescribeDataImportTaskInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDataImportTaskInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
