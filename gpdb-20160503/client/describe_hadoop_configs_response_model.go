// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHadoopConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHadoopConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHadoopConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHadoopConfigsResponseBody) *DescribeHadoopConfigsResponse
	GetBody() *DescribeHadoopConfigsResponseBody
}

type DescribeHadoopConfigsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHadoopConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHadoopConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHadoopConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeHadoopConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHadoopConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHadoopConfigsResponse) GetBody() *DescribeHadoopConfigsResponseBody {
	return s.Body
}

func (s *DescribeHadoopConfigsResponse) SetHeaders(v map[string]*string) *DescribeHadoopConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeHadoopConfigsResponse) SetStatusCode(v int32) *DescribeHadoopConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHadoopConfigsResponse) SetBody(v *DescribeHadoopConfigsResponseBody) *DescribeHadoopConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeHadoopConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
