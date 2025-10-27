// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddClientUserDefineRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddClientUserDefineRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddClientUserDefineRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddClientUserDefineRuleResponseBody) *AddClientUserDefineRuleResponse
	GetBody() *AddClientUserDefineRuleResponseBody
}

type AddClientUserDefineRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddClientUserDefineRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddClientUserDefineRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddClientUserDefineRuleResponse) GoString() string {
	return s.String()
}

func (s *AddClientUserDefineRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddClientUserDefineRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddClientUserDefineRuleResponse) GetBody() *AddClientUserDefineRuleResponseBody {
	return s.Body
}

func (s *AddClientUserDefineRuleResponse) SetHeaders(v map[string]*string) *AddClientUserDefineRuleResponse {
	s.Headers = v
	return s
}

func (s *AddClientUserDefineRuleResponse) SetStatusCode(v int32) *AddClientUserDefineRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddClientUserDefineRuleResponse) SetBody(v *AddClientUserDefineRuleResponseBody) *AddClientUserDefineRuleResponse {
	s.Body = v
	return s
}

func (s *AddClientUserDefineRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
