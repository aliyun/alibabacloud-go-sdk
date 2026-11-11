// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteErrorPagesRedirectRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteErrorPagesRedirectRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteErrorPagesRedirectRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteErrorPagesRedirectRuleResponseBody) *DeleteErrorPagesRedirectRuleResponse
	GetBody() *DeleteErrorPagesRedirectRuleResponseBody
}

type DeleteErrorPagesRedirectRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteErrorPagesRedirectRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteErrorPagesRedirectRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteErrorPagesRedirectRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteErrorPagesRedirectRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteErrorPagesRedirectRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteErrorPagesRedirectRuleResponse) GetBody() *DeleteErrorPagesRedirectRuleResponseBody {
	return s.Body
}

func (s *DeleteErrorPagesRedirectRuleResponse) SetHeaders(v map[string]*string) *DeleteErrorPagesRedirectRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteErrorPagesRedirectRuleResponse) SetStatusCode(v int32) *DeleteErrorPagesRedirectRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteErrorPagesRedirectRuleResponse) SetBody(v *DeleteErrorPagesRedirectRuleResponseBody) *DeleteErrorPagesRedirectRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteErrorPagesRedirectRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
