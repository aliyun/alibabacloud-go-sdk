// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRedirectRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRedirectRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRedirectRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetRedirectRuleResponseBody) *GetRedirectRuleResponse
	GetBody() *GetRedirectRuleResponseBody
}

type GetRedirectRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRedirectRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRedirectRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRedirectRuleResponse) GoString() string {
	return s.String()
}

func (s *GetRedirectRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRedirectRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRedirectRuleResponse) GetBody() *GetRedirectRuleResponseBody {
	return s.Body
}

func (s *GetRedirectRuleResponse) SetHeaders(v map[string]*string) *GetRedirectRuleResponse {
	s.Headers = v
	return s
}

func (s *GetRedirectRuleResponse) SetStatusCode(v int32) *GetRedirectRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRedirectRuleResponse) SetBody(v *GetRedirectRuleResponseBody) *GetRedirectRuleResponse {
	s.Body = v
	return s
}

func (s *GetRedirectRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
