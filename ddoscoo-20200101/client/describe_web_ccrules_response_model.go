// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCCRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebCCRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebCCRulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWebCCRulesResponseBody) *DescribeWebCCRulesResponse
	GetBody() *DescribeWebCCRulesResponseBody
}

type DescribeWebCCRulesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebCCRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebCCRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCCRulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebCCRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebCCRulesResponse) GetBody() *DescribeWebCCRulesResponseBody {
	return s.Body
}

func (s *DescribeWebCCRulesResponse) SetHeaders(v map[string]*string) *DescribeWebCCRulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebCCRulesResponse) SetStatusCode(v int32) *DescribeWebCCRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebCCRulesResponse) SetBody(v *DescribeWebCCRulesResponseBody) *DescribeWebCCRulesResponse {
	s.Body = v
	return s
}

func (s *DescribeWebCCRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
