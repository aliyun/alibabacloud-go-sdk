// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateAllRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTemplateAllRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTemplateAllRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTemplateAllRulesResponseBody) *DescribeTemplateAllRulesResponse
	GetBody() *DescribeTemplateAllRulesResponseBody
}

type DescribeTemplateAllRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTemplateAllRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTemplateAllRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateAllRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAllRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTemplateAllRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTemplateAllRulesResponse) GetBody() *DescribeTemplateAllRulesResponseBody {
	return s.Body
}

func (s *DescribeTemplateAllRulesResponse) SetHeaders(v map[string]*string) *DescribeTemplateAllRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplateAllRulesResponse) SetStatusCode(v int32) *DescribeTemplateAllRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplateAllRulesResponse) SetBody(v *DescribeTemplateAllRulesResponseBody) *DescribeTemplateAllRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeTemplateAllRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
