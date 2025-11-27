// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhitelistTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWhitelistTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWhitelistTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWhitelistTemplateResponseBody) *DescribeWhitelistTemplateResponse
	GetBody() *DescribeWhitelistTemplateResponseBody
}

type DescribeWhitelistTemplateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWhitelistTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWhitelistTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWhitelistTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWhitelistTemplateResponse) GetBody() *DescribeWhitelistTemplateResponseBody {
	return s.Body
}

func (s *DescribeWhitelistTemplateResponse) SetHeaders(v map[string]*string) *DescribeWhitelistTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhitelistTemplateResponse) SetStatusCode(v int32) *DescribeWhitelistTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWhitelistTemplateResponse) SetBody(v *DescribeWhitelistTemplateResponseBody) *DescribeWhitelistTemplateResponse {
	s.Body = v
	return s
}

func (s *DescribeWhitelistTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
