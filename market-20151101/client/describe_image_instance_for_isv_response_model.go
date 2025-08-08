// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageInstanceForIsvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageInstanceForIsvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageInstanceForIsvResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageInstanceForIsvResponseBody) *DescribeImageInstanceForIsvResponse
	GetBody() *DescribeImageInstanceForIsvResponseBody
}

type DescribeImageInstanceForIsvResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageInstanceForIsvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageInstanceForIsvResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageInstanceForIsvResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageInstanceForIsvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageInstanceForIsvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageInstanceForIsvResponse) GetBody() *DescribeImageInstanceForIsvResponseBody {
	return s.Body
}

func (s *DescribeImageInstanceForIsvResponse) SetHeaders(v map[string]*string) *DescribeImageInstanceForIsvResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageInstanceForIsvResponse) SetStatusCode(v int32) *DescribeImageInstanceForIsvResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageInstanceForIsvResponse) SetBody(v *DescribeImageInstanceForIsvResponseBody) *DescribeImageInstanceForIsvResponse {
	s.Body = v
	return s
}

func (s *DescribeImageInstanceForIsvResponse) Validate() error {
	return dara.Validate(s)
}
