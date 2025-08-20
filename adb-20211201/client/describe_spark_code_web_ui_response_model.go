// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSparkCodeWebUiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSparkCodeWebUiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSparkCodeWebUiResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSparkCodeWebUiResponseBody) *DescribeSparkCodeWebUiResponse
	GetBody() *DescribeSparkCodeWebUiResponseBody
}

type DescribeSparkCodeWebUiResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSparkCodeWebUiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSparkCodeWebUiResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSparkCodeWebUiResponse) GoString() string {
	return s.String()
}

func (s *DescribeSparkCodeWebUiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSparkCodeWebUiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSparkCodeWebUiResponse) GetBody() *DescribeSparkCodeWebUiResponseBody {
	return s.Body
}

func (s *DescribeSparkCodeWebUiResponse) SetHeaders(v map[string]*string) *DescribeSparkCodeWebUiResponse {
	s.Headers = v
	return s
}

func (s *DescribeSparkCodeWebUiResponse) SetStatusCode(v int32) *DescribeSparkCodeWebUiResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSparkCodeWebUiResponse) SetBody(v *DescribeSparkCodeWebUiResponseBody) *DescribeSparkCodeWebUiResponse {
	s.Body = v
	return s
}

func (s *DescribeSparkCodeWebUiResponse) Validate() error {
	return dara.Validate(s)
}
