// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableWebCCRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableWebCCRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableWebCCRuleResponse
	GetStatusCode() *int32
	SetBody(v *DisableWebCCRuleResponseBody) *DisableWebCCRuleResponse
	GetBody() *DisableWebCCRuleResponseBody
}

type DisableWebCCRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableWebCCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableWebCCRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableWebCCRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableWebCCRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableWebCCRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableWebCCRuleResponse) GetBody() *DisableWebCCRuleResponseBody {
	return s.Body
}

func (s *DisableWebCCRuleResponse) SetHeaders(v map[string]*string) *DisableWebCCRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableWebCCRuleResponse) SetStatusCode(v int32) *DisableWebCCRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableWebCCRuleResponse) SetBody(v *DisableWebCCRuleResponseBody) *DisableWebCCRuleResponse {
	s.Body = v
	return s
}

func (s *DisableWebCCRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
