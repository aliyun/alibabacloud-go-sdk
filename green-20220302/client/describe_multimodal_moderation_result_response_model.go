// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMultimodalModerationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMultimodalModerationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMultimodalModerationResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMultimodalModerationResultResponseBody) *DescribeMultimodalModerationResultResponse
	GetBody() *DescribeMultimodalModerationResultResponseBody
}

type DescribeMultimodalModerationResultResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMultimodalModerationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMultimodalModerationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMultimodalModerationResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeMultimodalModerationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMultimodalModerationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMultimodalModerationResultResponse) GetBody() *DescribeMultimodalModerationResultResponseBody {
	return s.Body
}

func (s *DescribeMultimodalModerationResultResponse) SetHeaders(v map[string]*string) *DescribeMultimodalModerationResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeMultimodalModerationResultResponse) SetStatusCode(v int32) *DescribeMultimodalModerationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMultimodalModerationResultResponse) SetBody(v *DescribeMultimodalModerationResultResponseBody) *DescribeMultimodalModerationResultResponse {
	s.Body = v
	return s
}

func (s *DescribeMultimodalModerationResultResponse) Validate() error {
	return dara.Validate(s)
}
