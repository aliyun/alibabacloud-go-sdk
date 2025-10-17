// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkAppTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSparkAppTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSparkAppTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSparkAppTypeResponseBody) *DescribeSparkAppTypeResponse
	GetBody() *DescribeSparkAppTypeResponseBody
}

type DescribeSparkAppTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSparkAppTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSparkAppTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkAppTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSparkAppTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSparkAppTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSparkAppTypeResponse) GetBody() *DescribeSparkAppTypeResponseBody {
	return s.Body
}

func (s *DescribeSparkAppTypeResponse) SetHeaders(v map[string]*string) *DescribeSparkAppTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSparkAppTypeResponse) SetStatusCode(v int32) *DescribeSparkAppTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSparkAppTypeResponse) SetBody(v *DescribeSparkAppTypeResponseBody) *DescribeSparkAppTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeSparkAppTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
