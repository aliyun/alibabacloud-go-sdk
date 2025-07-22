// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostCenterRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCostCenterRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCostCenterRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCostCenterRuleResponseBody) *ModifyCostCenterRuleResponse
	GetBody() *ModifyCostCenterRuleResponseBody
}

type ModifyCostCenterRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCostCenterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCostCenterRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostCenterRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCostCenterRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCostCenterRuleResponse) GetBody() *ModifyCostCenterRuleResponseBody {
	return s.Body
}

func (s *ModifyCostCenterRuleResponse) SetHeaders(v map[string]*string) *ModifyCostCenterRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyCostCenterRuleResponse) SetStatusCode(v int32) *ModifyCostCenterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCostCenterRuleResponse) SetBody(v *ModifyCostCenterRuleResponseBody) *ModifyCostCenterRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyCostCenterRuleResponse) Validate() error {
	return dara.Validate(s)
}
