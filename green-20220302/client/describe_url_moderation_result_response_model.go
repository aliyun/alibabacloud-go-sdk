// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUrlModerationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUrlModerationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUrlModerationResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUrlModerationResultResponseBody) *DescribeUrlModerationResultResponse
	GetBody() *DescribeUrlModerationResultResponseBody
}

type DescribeUrlModerationResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUrlModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUrlModerationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUrlModerationResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeUrlModerationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUrlModerationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUrlModerationResultResponse) GetBody() *DescribeUrlModerationResultResponseBody {
	return s.Body
}

func (s *DescribeUrlModerationResultResponse) SetHeaders(v map[string]*string) *DescribeUrlModerationResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeUrlModerationResultResponse) SetStatusCode(v int32) *DescribeUrlModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUrlModerationResultResponse) SetBody(v *DescribeUrlModerationResultResponseBody) *DescribeUrlModerationResultResponse {
	s.Body = v
	return s
}

func (s *DescribeUrlModerationResultResponse) Validate() error {
	return dara.Validate(s)
}
