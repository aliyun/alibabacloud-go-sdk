// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessInstanceTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccessInstanceTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccessInstanceTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccessInstanceTaskResponseBody) *DescribeAccessInstanceTaskResponse
	GetBody() *DescribeAccessInstanceTaskResponseBody
}

type DescribeAccessInstanceTaskResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccessInstanceTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccessInstanceTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessInstanceTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccessInstanceTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccessInstanceTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccessInstanceTaskResponse) GetBody() *DescribeAccessInstanceTaskResponseBody {
	return s.Body
}

func (s *DescribeAccessInstanceTaskResponse) SetHeaders(v map[string]*string) *DescribeAccessInstanceTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccessInstanceTaskResponse) SetStatusCode(v int32) *DescribeAccessInstanceTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccessInstanceTaskResponse) SetBody(v *DescribeAccessInstanceTaskResponseBody) *DescribeAccessInstanceTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeAccessInstanceTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
