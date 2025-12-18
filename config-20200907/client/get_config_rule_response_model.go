// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConfigRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConfigRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetConfigRuleResponseBody) *GetConfigRuleResponse
	GetBody() *GetConfigRuleResponseBody
}

type GetConfigRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfigRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *GetConfigRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConfigRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConfigRuleResponse) GetBody() *GetConfigRuleResponseBody {
	return s.Body
}

func (s *GetConfigRuleResponse) SetHeaders(v map[string]*string) *GetConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *GetConfigRuleResponse) SetStatusCode(v int32) *GetConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigRuleResponse) SetBody(v *GetConfigRuleResponseBody) *GetConfigRuleResponse {
	s.Body = v
	return s
}

func (s *GetConfigRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
