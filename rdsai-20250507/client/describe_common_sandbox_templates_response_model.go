// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonSandboxTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCommonSandboxTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCommonSandboxTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCommonSandboxTemplatesResponseBody) *DescribeCommonSandboxTemplatesResponse
	GetBody() *DescribeCommonSandboxTemplatesResponseBody
}

type DescribeCommonSandboxTemplatesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommonSandboxTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommonSandboxTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonSandboxTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommonSandboxTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCommonSandboxTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCommonSandboxTemplatesResponse) GetBody() *DescribeCommonSandboxTemplatesResponseBody {
	return s.Body
}

func (s *DescribeCommonSandboxTemplatesResponse) SetHeaders(v map[string]*string) *DescribeCommonSandboxTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommonSandboxTemplatesResponse) SetStatusCode(v int32) *DescribeCommonSandboxTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommonSandboxTemplatesResponse) SetBody(v *DescribeCommonSandboxTemplatesResponseBody) *DescribeCommonSandboxTemplatesResponse {
	s.Body = v
	return s
}

func (s *DescribeCommonSandboxTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
