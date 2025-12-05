// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetRuleResponseBody) *GetRuleResponse
	GetBody() *GetRuleResponseBody
}

type GetRuleResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRuleResponse) GoString() string {
	return s.String()
}

func (s *GetRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRuleResponse) GetBody() *GetRuleResponseBody {
	return s.Body
}

func (s *GetRuleResponse) SetHeaders(v map[string]*string) *GetRuleResponse {
	s.Headers = v
	return s
}

func (s *GetRuleResponse) SetStatusCode(v int32) *GetRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRuleResponse) SetBody(v *GetRuleResponseBody) *GetRuleResponse {
	s.Body = v
	return s
}

func (s *GetRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
