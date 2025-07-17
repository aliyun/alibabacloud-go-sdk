// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageModerationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageModerationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageModerationResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageModerationResultResponseBody) *DescribeImageModerationResultResponse
	GetBody() *DescribeImageModerationResultResponseBody
}

type DescribeImageModerationResultResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageModerationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageModerationResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageModerationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageModerationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageModerationResultResponse) GetBody() *DescribeImageModerationResultResponseBody {
	return s.Body
}

func (s *DescribeImageModerationResultResponse) SetHeaders(v map[string]*string) *DescribeImageModerationResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageModerationResultResponse) SetStatusCode(v int32) *DescribeImageModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageModerationResultResponse) SetBody(v *DescribeImageModerationResultResponseBody) *DescribeImageModerationResultResponse {
	s.Body = v
	return s
}

func (s *DescribeImageModerationResultResponse) Validate() error {
	return dara.Validate(s)
}
