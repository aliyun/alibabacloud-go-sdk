// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationPromptsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApplicationPromptsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApplicationPromptsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApplicationPromptsResponseBody) *DescribeApplicationPromptsResponse
	GetBody() *DescribeApplicationPromptsResponseBody
}

type DescribeApplicationPromptsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApplicationPromptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApplicationPromptsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationPromptsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApplicationPromptsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApplicationPromptsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApplicationPromptsResponse) GetBody() *DescribeApplicationPromptsResponseBody {
	return s.Body
}

func (s *DescribeApplicationPromptsResponse) SetHeaders(v map[string]*string) *DescribeApplicationPromptsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApplicationPromptsResponse) SetStatusCode(v int32) *DescribeApplicationPromptsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApplicationPromptsResponse) SetBody(v *DescribeApplicationPromptsResponseBody) *DescribeApplicationPromptsResponse {
	s.Body = v
	return s
}

func (s *DescribeApplicationPromptsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
