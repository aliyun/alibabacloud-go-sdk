// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventVariableTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventVariableTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventVariableTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventVariableTemplateListResponseBody) *DescribeEventVariableTemplateListResponse
	GetBody() *DescribeEventVariableTemplateListResponseBody
}

type DescribeEventVariableTemplateListResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventVariableTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventVariableTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableTemplateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventVariableTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventVariableTemplateListResponse) GetBody() *DescribeEventVariableTemplateListResponseBody {
	return s.Body
}

func (s *DescribeEventVariableTemplateListResponse) SetHeaders(v map[string]*string) *DescribeEventVariableTemplateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventVariableTemplateListResponse) SetStatusCode(v int32) *DescribeEventVariableTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventVariableTemplateListResponse) SetBody(v *DescribeEventVariableTemplateListResponseBody) *DescribeEventVariableTemplateListResponse {
	s.Body = v
	return s
}

func (s *DescribeEventVariableTemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
