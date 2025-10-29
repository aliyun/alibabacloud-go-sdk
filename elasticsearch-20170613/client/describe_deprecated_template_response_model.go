// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeprecatedTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDeprecatedTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDeprecatedTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDeprecatedTemplateResponseBody) *DescribeDeprecatedTemplateResponse
	GetBody() *DescribeDeprecatedTemplateResponseBody
}

type DescribeDeprecatedTemplateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDeprecatedTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDeprecatedTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeprecatedTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeDeprecatedTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDeprecatedTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDeprecatedTemplateResponse) GetBody() *DescribeDeprecatedTemplateResponseBody {
	return s.Body
}

func (s *DescribeDeprecatedTemplateResponse) SetHeaders(v map[string]*string) *DescribeDeprecatedTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeDeprecatedTemplateResponse) SetStatusCode(v int32) *DescribeDeprecatedTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDeprecatedTemplateResponse) SetBody(v *DescribeDeprecatedTemplateResponseBody) *DescribeDeprecatedTemplateResponse {
	s.Body = v
	return s
}

func (s *DescribeDeprecatedTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
