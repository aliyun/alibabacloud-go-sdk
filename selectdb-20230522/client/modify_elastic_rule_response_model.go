// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyElasticRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyElasticRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyElasticRuleResponseBody) *ModifyElasticRuleResponse
	GetBody() *ModifyElasticRuleResponseBody
}

type ModifyElasticRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyElasticRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyElasticRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyElasticRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyElasticRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyElasticRuleResponse) GetBody() *ModifyElasticRuleResponseBody {
	return s.Body
}

func (s *ModifyElasticRuleResponse) SetHeaders(v map[string]*string) *ModifyElasticRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyElasticRuleResponse) SetStatusCode(v int32) *ModifyElasticRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyElasticRuleResponse) SetBody(v *ModifyElasticRuleResponseBody) *ModifyElasticRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyElasticRuleResponse) Validate() error {
	return dara.Validate(s)
}
