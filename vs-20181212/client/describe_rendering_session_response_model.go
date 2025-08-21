// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRenderingSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRenderingSessionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRenderingSessionResponseBody) *DescribeRenderingSessionResponse
	GetBody() *DescribeRenderingSessionResponseBody
}

type DescribeRenderingSessionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRenderingSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRenderingSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingSessionResponse) GoString() string {
	return s.String()
}

func (s *DescribeRenderingSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRenderingSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRenderingSessionResponse) GetBody() *DescribeRenderingSessionResponseBody {
	return s.Body
}

func (s *DescribeRenderingSessionResponse) SetHeaders(v map[string]*string) *DescribeRenderingSessionResponse {
	s.Headers = v
	return s
}

func (s *DescribeRenderingSessionResponse) SetStatusCode(v int32) *DescribeRenderingSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRenderingSessionResponse) SetBody(v *DescribeRenderingSessionResponseBody) *DescribeRenderingSessionResponse {
	s.Body = v
	return s
}

func (s *DescribeRenderingSessionResponse) Validate() error {
	return dara.Validate(s)
}
