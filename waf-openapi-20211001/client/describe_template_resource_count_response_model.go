// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateResourceCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTemplateResourceCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTemplateResourceCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTemplateResourceCountResponseBody) *DescribeTemplateResourceCountResponse
	GetBody() *DescribeTemplateResourceCountResponseBody
}

type DescribeTemplateResourceCountResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTemplateResourceCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTemplateResourceCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateResourceCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourceCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTemplateResourceCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTemplateResourceCountResponse) GetBody() *DescribeTemplateResourceCountResponseBody {
	return s.Body
}

func (s *DescribeTemplateResourceCountResponse) SetHeaders(v map[string]*string) *DescribeTemplateResourceCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateResourceCountResponse) SetStatusCode(v int32) *DescribeTemplateResourceCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplateResourceCountResponse) SetBody(v *DescribeTemplateResourceCountResponseBody) *DescribeTemplateResourceCountResponse {
	s.Body = v
	return s
}

func (s *DescribeTemplateResourceCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
