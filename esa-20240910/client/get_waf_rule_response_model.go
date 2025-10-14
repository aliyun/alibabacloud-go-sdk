// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWafRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWafRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWafRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetWafRuleResponseBody) *GetWafRuleResponse
	GetBody() *GetWafRuleResponseBody
}

type GetWafRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWafRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWafRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWafRuleResponse) GoString() string {
	return s.String()
}

func (s *GetWafRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWafRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWafRuleResponse) GetBody() *GetWafRuleResponseBody {
	return s.Body
}

func (s *GetWafRuleResponse) SetHeaders(v map[string]*string) *GetWafRuleResponse {
	s.Headers = v
	return s
}

func (s *GetWafRuleResponse) SetStatusCode(v int32) *GetWafRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWafRuleResponse) SetBody(v *GetWafRuleResponseBody) *GetWafRuleResponse {
	s.Body = v
	return s
}

func (s *GetWafRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
