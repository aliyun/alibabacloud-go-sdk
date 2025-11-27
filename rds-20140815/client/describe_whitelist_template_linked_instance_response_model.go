// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhitelistTemplateLinkedInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWhitelistTemplateLinkedInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWhitelistTemplateLinkedInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWhitelistTemplateLinkedInstanceResponseBody) *DescribeWhitelistTemplateLinkedInstanceResponse
	GetBody() *DescribeWhitelistTemplateLinkedInstanceResponseBody
}

type DescribeWhitelistTemplateLinkedInstanceResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWhitelistTemplateLinkedInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWhitelistTemplateLinkedInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistTemplateLinkedInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponse) GetBody() *DescribeWhitelistTemplateLinkedInstanceResponseBody {
	return s.Body
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponse) SetHeaders(v map[string]*string) *DescribeWhitelistTemplateLinkedInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponse) SetStatusCode(v int32) *DescribeWhitelistTemplateLinkedInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponse) SetBody(v *DescribeWhitelistTemplateLinkedInstanceResponseBody) *DescribeWhitelistTemplateLinkedInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeWhitelistTemplateLinkedInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
