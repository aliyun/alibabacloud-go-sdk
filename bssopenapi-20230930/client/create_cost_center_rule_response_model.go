// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostCenterRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCostCenterRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCostCenterRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateCostCenterRuleResponseBody) *CreateCostCenterRuleResponse
	GetBody() *CreateCostCenterRuleResponseBody
}

type CreateCostCenterRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCostCenterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCostCenterRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCostCenterRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateCostCenterRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCostCenterRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCostCenterRuleResponse) GetBody() *CreateCostCenterRuleResponseBody {
	return s.Body
}

func (s *CreateCostCenterRuleResponse) SetHeaders(v map[string]*string) *CreateCostCenterRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateCostCenterRuleResponse) SetStatusCode(v int32) *CreateCostCenterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCostCenterRuleResponse) SetBody(v *CreateCostCenterRuleResponseBody) *CreateCostCenterRuleResponse {
	s.Body = v
	return s
}

func (s *CreateCostCenterRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
