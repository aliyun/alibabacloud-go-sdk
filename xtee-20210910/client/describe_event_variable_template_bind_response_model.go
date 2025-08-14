// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventVariableTemplateBindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventVariableTemplateBindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventVariableTemplateBindResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventVariableTemplateBindResponseBody) *DescribeEventVariableTemplateBindResponse
	GetBody() *DescribeEventVariableTemplateBindResponseBody
}

type DescribeEventVariableTemplateBindResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventVariableTemplateBindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventVariableTemplateBindResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableTemplateBindResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableTemplateBindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventVariableTemplateBindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventVariableTemplateBindResponse) GetBody() *DescribeEventVariableTemplateBindResponseBody {
	return s.Body
}

func (s *DescribeEventVariableTemplateBindResponse) SetHeaders(v map[string]*string) *DescribeEventVariableTemplateBindResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventVariableTemplateBindResponse) SetStatusCode(v int32) *DescribeEventVariableTemplateBindResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventVariableTemplateBindResponse) SetBody(v *DescribeEventVariableTemplateBindResponseBody) *DescribeEventVariableTemplateBindResponse {
	s.Body = v
	return s
}

func (s *DescribeEventVariableTemplateBindResponse) Validate() error {
	return dara.Validate(s)
}
