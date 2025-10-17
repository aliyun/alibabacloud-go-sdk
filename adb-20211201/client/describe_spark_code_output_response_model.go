// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkCodeOutputResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSparkCodeOutputResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSparkCodeOutputResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSparkCodeOutputResponseBody) *DescribeSparkCodeOutputResponse
	GetBody() *DescribeSparkCodeOutputResponseBody
}

type DescribeSparkCodeOutputResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSparkCodeOutputResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSparkCodeOutputResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkCodeOutputResponse) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeOutputResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSparkCodeOutputResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSparkCodeOutputResponse) GetBody() *DescribeSparkCodeOutputResponseBody {
	return s.Body
}

func (s *DescribeSparkCodeOutputResponse) SetHeaders(v map[string]*string) *DescribeSparkCodeOutputResponse {
	s.Headers = v
	return s
}

func (s *DescribeSparkCodeOutputResponse) SetStatusCode(v int32) *DescribeSparkCodeOutputResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSparkCodeOutputResponse) SetBody(v *DescribeSparkCodeOutputResponseBody) *DescribeSparkCodeOutputResponse {
	s.Body = v
	return s
}

func (s *DescribeSparkCodeOutputResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
