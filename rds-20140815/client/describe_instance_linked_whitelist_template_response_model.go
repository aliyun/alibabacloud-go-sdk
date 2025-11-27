// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceLinkedWhitelistTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceLinkedWhitelistTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceLinkedWhitelistTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceLinkedWhitelistTemplateResponseBody) *DescribeInstanceLinkedWhitelistTemplateResponse
	GetBody() *DescribeInstanceLinkedWhitelistTemplateResponseBody
}

type DescribeInstanceLinkedWhitelistTemplateResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceLinkedWhitelistTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceLinkedWhitelistTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLinkedWhitelistTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponse) GetBody() *DescribeInstanceLinkedWhitelistTemplateResponseBody {
	return s.Body
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponse) SetHeaders(v map[string]*string) *DescribeInstanceLinkedWhitelistTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponse) SetStatusCode(v int32) *DescribeInstanceLinkedWhitelistTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponse) SetBody(v *DescribeInstanceLinkedWhitelistTemplateResponseBody) *DescribeInstanceLinkedWhitelistTemplateResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceLinkedWhitelistTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
