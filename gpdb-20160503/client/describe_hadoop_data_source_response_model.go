// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHadoopDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHadoopDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHadoopDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHadoopDataSourceResponseBody) *DescribeHadoopDataSourceResponse
	GetBody() *DescribeHadoopDataSourceResponseBody
}

type DescribeHadoopDataSourceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHadoopDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHadoopDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHadoopDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeHadoopDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHadoopDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHadoopDataSourceResponse) GetBody() *DescribeHadoopDataSourceResponseBody {
	return s.Body
}

func (s *DescribeHadoopDataSourceResponse) SetHeaders(v map[string]*string) *DescribeHadoopDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeHadoopDataSourceResponse) SetStatusCode(v int32) *DescribeHadoopDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHadoopDataSourceResponse) SetBody(v *DescribeHadoopDataSourceResponseBody) *DescribeHadoopDataSourceResponse {
	s.Body = v
	return s
}

func (s *DescribeHadoopDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
