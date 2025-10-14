// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRewriteUrlRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRewriteUrlRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRewriteUrlRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateRewriteUrlRuleResponseBody) *CreateRewriteUrlRuleResponse
	GetBody() *CreateRewriteUrlRuleResponseBody
}

type CreateRewriteUrlRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRewriteUrlRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRewriteUrlRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRewriteUrlRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateRewriteUrlRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRewriteUrlRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRewriteUrlRuleResponse) GetBody() *CreateRewriteUrlRuleResponseBody {
	return s.Body
}

func (s *CreateRewriteUrlRuleResponse) SetHeaders(v map[string]*string) *CreateRewriteUrlRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateRewriteUrlRuleResponse) SetStatusCode(v int32) *CreateRewriteUrlRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRewriteUrlRuleResponse) SetBody(v *CreateRewriteUrlRuleResponseBody) *CreateRewriteUrlRuleResponse {
	s.Body = v
	return s
}

func (s *CreateRewriteUrlRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
