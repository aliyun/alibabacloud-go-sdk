// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyQosRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyQosRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyQosRuleResponseBody) *ModifyQosRuleResponse
	GetBody() *ModifyQosRuleResponseBody
}

type ModifyQosRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyQosRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyQosRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyQosRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyQosRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyQosRuleResponse) GetBody() *ModifyQosRuleResponseBody {
	return s.Body
}

func (s *ModifyQosRuleResponse) SetHeaders(v map[string]*string) *ModifyQosRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyQosRuleResponse) SetStatusCode(v int32) *ModifyQosRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyQosRuleResponse) SetBody(v *ModifyQosRuleResponseBody) *ModifyQosRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyQosRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
