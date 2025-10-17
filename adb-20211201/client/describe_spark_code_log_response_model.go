// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkCodeLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSparkCodeLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSparkCodeLogResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSparkCodeLogResponseBody) *DescribeSparkCodeLogResponse
	GetBody() *DescribeSparkCodeLogResponseBody
}

type DescribeSparkCodeLogResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSparkCodeLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSparkCodeLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkCodeLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSparkCodeLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSparkCodeLogResponse) GetBody() *DescribeSparkCodeLogResponseBody {
	return s.Body
}

func (s *DescribeSparkCodeLogResponse) SetHeaders(v map[string]*string) *DescribeSparkCodeLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeSparkCodeLogResponse) SetStatusCode(v int32) *DescribeSparkCodeLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSparkCodeLogResponse) SetBody(v *DescribeSparkCodeLogResponseBody) *DescribeSparkCodeLogResponse {
	s.Body = v
	return s
}

func (s *DescribeSparkCodeLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
