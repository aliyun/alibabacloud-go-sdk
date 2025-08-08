// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceForIsvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceForIsvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceForIsvResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceForIsvResponseBody) *DescribeInstanceForIsvResponse
	GetBody() *DescribeInstanceForIsvResponseBody
}

type DescribeInstanceForIsvResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceForIsvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceForIsvResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceForIsvResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceForIsvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceForIsvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceForIsvResponse) GetBody() *DescribeInstanceForIsvResponseBody {
	return s.Body
}

func (s *DescribeInstanceForIsvResponse) SetHeaders(v map[string]*string) *DescribeInstanceForIsvResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceForIsvResponse) SetStatusCode(v int32) *DescribeInstanceForIsvResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceForIsvResponse) SetBody(v *DescribeInstanceForIsvResponseBody) *DescribeInstanceForIsvResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceForIsvResponse) Validate() error {
	return dara.Validate(s)
}
