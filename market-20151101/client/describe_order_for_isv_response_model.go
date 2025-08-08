// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrderForIsvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOrderForIsvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOrderForIsvResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOrderForIsvResponseBody) *DescribeOrderForIsvResponse
	GetBody() *DescribeOrderForIsvResponseBody
}

type DescribeOrderForIsvResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOrderForIsvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOrderForIsvResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrderForIsvResponse) GoString() string {
	return s.String()
}

func (s *DescribeOrderForIsvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOrderForIsvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOrderForIsvResponse) GetBody() *DescribeOrderForIsvResponseBody {
	return s.Body
}

func (s *DescribeOrderForIsvResponse) SetHeaders(v map[string]*string) *DescribeOrderForIsvResponse {
	s.Headers = v
	return s
}

func (s *DescribeOrderForIsvResponse) SetStatusCode(v int32) *DescribeOrderForIsvResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOrderForIsvResponse) SetBody(v *DescribeOrderForIsvResponseBody) *DescribeOrderForIsvResponse {
	s.Body = v
	return s
}

func (s *DescribeOrderForIsvResponse) Validate() error {
	return dara.Validate(s)
}
