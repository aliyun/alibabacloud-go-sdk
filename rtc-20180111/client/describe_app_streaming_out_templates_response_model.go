// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppStreamingOutTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppStreamingOutTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppStreamingOutTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppStreamingOutTemplatesResponseBody) *DescribeAppStreamingOutTemplatesResponse
	GetBody() *DescribeAppStreamingOutTemplatesResponseBody
}

type DescribeAppStreamingOutTemplatesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppStreamingOutTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppStreamingOutTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppStreamingOutTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppStreamingOutTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppStreamingOutTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppStreamingOutTemplatesResponse) GetBody() *DescribeAppStreamingOutTemplatesResponseBody {
	return s.Body
}

func (s *DescribeAppStreamingOutTemplatesResponse) SetHeaders(v map[string]*string) *DescribeAppStreamingOutTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponse) SetStatusCode(v int32) *DescribeAppStreamingOutTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponse) SetBody(v *DescribeAppStreamingOutTemplatesResponseBody) *DescribeAppStreamingOutTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponse) Validate() error {
	return dara.Validate(s)
}
