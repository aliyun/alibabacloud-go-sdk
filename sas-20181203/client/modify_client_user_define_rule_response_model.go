// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClientUserDefineRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyClientUserDefineRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyClientUserDefineRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyClientUserDefineRuleResponseBody) *ModifyClientUserDefineRuleResponse
	GetBody() *ModifyClientUserDefineRuleResponseBody
}

type ModifyClientUserDefineRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyClientUserDefineRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyClientUserDefineRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyClientUserDefineRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyClientUserDefineRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyClientUserDefineRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyClientUserDefineRuleResponse) GetBody() *ModifyClientUserDefineRuleResponseBody {
	return s.Body
}

func (s *ModifyClientUserDefineRuleResponse) SetHeaders(v map[string]*string) *ModifyClientUserDefineRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyClientUserDefineRuleResponse) SetStatusCode(v int32) *ModifyClientUserDefineRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyClientUserDefineRuleResponse) SetBody(v *ModifyClientUserDefineRuleResponseBody) *ModifyClientUserDefineRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyClientUserDefineRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
