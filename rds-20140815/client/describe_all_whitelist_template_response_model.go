// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllWhitelistTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAllWhitelistTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAllWhitelistTemplateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAllWhitelistTemplateResponseBody) *DescribeAllWhitelistTemplateResponse
	GetBody() *DescribeAllWhitelistTemplateResponseBody
}

type DescribeAllWhitelistTemplateResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllWhitelistTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAllWhitelistTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllWhitelistTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllWhitelistTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAllWhitelistTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAllWhitelistTemplateResponse) GetBody() *DescribeAllWhitelistTemplateResponseBody {
	return s.Body
}

func (s *DescribeAllWhitelistTemplateResponse) SetHeaders(v map[string]*string) *DescribeAllWhitelistTemplateResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllWhitelistTemplateResponse) SetStatusCode(v int32) *DescribeAllWhitelistTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllWhitelistTemplateResponse) SetBody(v *DescribeAllWhitelistTemplateResponseBody) *DescribeAllWhitelistTemplateResponse {
	s.Body = v
	return s
}

func (s *DescribeAllWhitelistTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
