// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDefenseTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDefenseTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDefenseTemplateResponseBody) *DescribeDefenseTemplateResponse
	GetBody() *DescribeDefenseTemplateResponseBody
}

type DescribeDefenseTemplateResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDefenseTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDefenseTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDefenseTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDefenseTemplateResponse) GetBody() *DescribeDefenseTemplateResponseBody {
	return s.Body
}

func (s *DescribeDefenseTemplateResponse) SetHeaders(v map[string]*string) *DescribeDefenseTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeDefenseTemplateResponse) SetStatusCode(v int32) *DescribeDefenseTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDefenseTemplateResponse) SetBody(v *DescribeDefenseTemplateResponseBody) *DescribeDefenseTemplateResponse {
	s.Body = v
	return s
}

func (s *DescribeDefenseTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
