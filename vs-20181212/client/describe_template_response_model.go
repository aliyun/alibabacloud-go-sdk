// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTemplateResponseBody) *DescribeTemplateResponse
	GetBody() *DescribeTemplateResponseBody
}

type DescribeTemplateResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTemplateResponse) GetBody() *DescribeTemplateResponseBody {
	return s.Body
}

func (s *DescribeTemplateResponse) SetHeaders(v map[string]*string) *DescribeTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateResponse) SetStatusCode(v int32) *DescribeTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplateResponse) SetBody(v *DescribeTemplateResponseBody) *DescribeTemplateResponse {
	s.Body = v
	return s
}

func (s *DescribeTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
