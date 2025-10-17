// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTablePartitionDiagnoseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTablePartitionDiagnoseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTablePartitionDiagnoseResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTablePartitionDiagnoseResponseBody) *DescribeTablePartitionDiagnoseResponse
	GetBody() *DescribeTablePartitionDiagnoseResponseBody
}

type DescribeTablePartitionDiagnoseResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTablePartitionDiagnoseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTablePartitionDiagnoseResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablePartitionDiagnoseResponse) GoString() string {
	return s.String()
}

func (s *DescribeTablePartitionDiagnoseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTablePartitionDiagnoseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTablePartitionDiagnoseResponse) GetBody() *DescribeTablePartitionDiagnoseResponseBody {
	return s.Body
}

func (s *DescribeTablePartitionDiagnoseResponse) SetHeaders(v map[string]*string) *DescribeTablePartitionDiagnoseResponse {
	s.Headers = v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponse) SetStatusCode(v int32) *DescribeTablePartitionDiagnoseResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponse) SetBody(v *DescribeTablePartitionDiagnoseResponseBody) *DescribeTablePartitionDiagnoseResponse {
	s.Body = v
	return s
}

func (s *DescribeTablePartitionDiagnoseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
