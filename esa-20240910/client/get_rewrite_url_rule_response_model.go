// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRewriteUrlRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRewriteUrlRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRewriteUrlRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetRewriteUrlRuleResponseBody) *GetRewriteUrlRuleResponse
	GetBody() *GetRewriteUrlRuleResponseBody
}

type GetRewriteUrlRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRewriteUrlRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRewriteUrlRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRewriteUrlRuleResponse) GoString() string {
	return s.String()
}

func (s *GetRewriteUrlRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRewriteUrlRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRewriteUrlRuleResponse) GetBody() *GetRewriteUrlRuleResponseBody {
	return s.Body
}

func (s *GetRewriteUrlRuleResponse) SetHeaders(v map[string]*string) *GetRewriteUrlRuleResponse {
	s.Headers = v
	return s
}

func (s *GetRewriteUrlRuleResponse) SetStatusCode(v int32) *GetRewriteUrlRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRewriteUrlRuleResponse) SetBody(v *GetRewriteUrlRuleResponseBody) *GetRewriteUrlRuleResponse {
	s.Body = v
	return s
}

func (s *GetRewriteUrlRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
