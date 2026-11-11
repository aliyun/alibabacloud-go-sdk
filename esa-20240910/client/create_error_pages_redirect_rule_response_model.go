// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateErrorPagesRedirectRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateErrorPagesRedirectRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateErrorPagesRedirectRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateErrorPagesRedirectRuleResponseBody) *CreateErrorPagesRedirectRuleResponse
	GetBody() *CreateErrorPagesRedirectRuleResponseBody
}

type CreateErrorPagesRedirectRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateErrorPagesRedirectRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateErrorPagesRedirectRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateErrorPagesRedirectRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateErrorPagesRedirectRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateErrorPagesRedirectRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateErrorPagesRedirectRuleResponse) GetBody() *CreateErrorPagesRedirectRuleResponseBody {
	return s.Body
}

func (s *CreateErrorPagesRedirectRuleResponse) SetHeaders(v map[string]*string) *CreateErrorPagesRedirectRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateErrorPagesRedirectRuleResponse) SetStatusCode(v int32) *CreateErrorPagesRedirectRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleResponse) SetBody(v *CreateErrorPagesRedirectRuleResponseBody) *CreateErrorPagesRedirectRuleResponse {
	s.Body = v
	return s
}

func (s *CreateErrorPagesRedirectRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
