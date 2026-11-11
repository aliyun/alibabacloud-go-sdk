// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateErrorPagesRedirectRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateErrorPagesRedirectRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateErrorPagesRedirectRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateErrorPagesRedirectRuleResponseBody) *UpdateErrorPagesRedirectRuleResponse
	GetBody() *UpdateErrorPagesRedirectRuleResponseBody
}

type UpdateErrorPagesRedirectRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateErrorPagesRedirectRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateErrorPagesRedirectRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateErrorPagesRedirectRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateErrorPagesRedirectRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateErrorPagesRedirectRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateErrorPagesRedirectRuleResponse) GetBody() *UpdateErrorPagesRedirectRuleResponseBody {
	return s.Body
}

func (s *UpdateErrorPagesRedirectRuleResponse) SetHeaders(v map[string]*string) *UpdateErrorPagesRedirectRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateErrorPagesRedirectRuleResponse) SetStatusCode(v int32) *UpdateErrorPagesRedirectRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateErrorPagesRedirectRuleResponse) SetBody(v *UpdateErrorPagesRedirectRuleResponseBody) *UpdateErrorPagesRedirectRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateErrorPagesRedirectRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
