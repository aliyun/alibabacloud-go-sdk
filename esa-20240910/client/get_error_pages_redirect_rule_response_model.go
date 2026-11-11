// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorPagesRedirectRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetErrorPagesRedirectRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetErrorPagesRedirectRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetErrorPagesRedirectRuleResponseBody) *GetErrorPagesRedirectRuleResponse
	GetBody() *GetErrorPagesRedirectRuleResponseBody
}

type GetErrorPagesRedirectRuleResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetErrorPagesRedirectRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetErrorPagesRedirectRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetErrorPagesRedirectRuleResponse) GoString() string {
	return s.String()
}

func (s *GetErrorPagesRedirectRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetErrorPagesRedirectRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetErrorPagesRedirectRuleResponse) GetBody() *GetErrorPagesRedirectRuleResponseBody {
	return s.Body
}

func (s *GetErrorPagesRedirectRuleResponse) SetHeaders(v map[string]*string) *GetErrorPagesRedirectRuleResponse {
	s.Headers = v
	return s
}

func (s *GetErrorPagesRedirectRuleResponse) SetStatusCode(v int32) *GetErrorPagesRedirectRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetErrorPagesRedirectRuleResponse) SetBody(v *GetErrorPagesRedirectRuleResponseBody) *GetErrorPagesRedirectRuleResponse {
	s.Body = v
	return s
}

func (s *GetErrorPagesRedirectRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
