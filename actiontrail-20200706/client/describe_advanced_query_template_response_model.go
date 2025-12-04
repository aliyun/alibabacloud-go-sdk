// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvancedQueryTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdvancedQueryTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdvancedQueryTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdvancedQueryTemplateResponseBody) *DescribeAdvancedQueryTemplateResponse
	GetBody() *DescribeAdvancedQueryTemplateResponseBody
}

type DescribeAdvancedQueryTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvancedQueryTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvancedQueryTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedQueryTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedQueryTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdvancedQueryTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdvancedQueryTemplateResponse) GetBody() *DescribeAdvancedQueryTemplateResponseBody {
	return s.Body
}

func (s *DescribeAdvancedQueryTemplateResponse) SetHeaders(v map[string]*string) *DescribeAdvancedQueryTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvancedQueryTemplateResponse) SetStatusCode(v int32) *DescribeAdvancedQueryTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvancedQueryTemplateResponse) SetBody(v *DescribeAdvancedQueryTemplateResponseBody) *DescribeAdvancedQueryTemplateResponse {
	s.Body = v
	return s
}

func (s *DescribeAdvancedQueryTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
