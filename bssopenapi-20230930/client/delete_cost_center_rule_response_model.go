// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostCenterRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCostCenterRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCostCenterRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCostCenterRuleResponseBody) *DeleteCostCenterRuleResponse
	GetBody() *DeleteCostCenterRuleResponseBody
}

type DeleteCostCenterRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCostCenterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCostCenterRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostCenterRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCostCenterRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCostCenterRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCostCenterRuleResponse) GetBody() *DeleteCostCenterRuleResponseBody {
	return s.Body
}

func (s *DeleteCostCenterRuleResponse) SetHeaders(v map[string]*string) *DeleteCostCenterRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteCostCenterRuleResponse) SetStatusCode(v int32) *DeleteCostCenterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCostCenterRuleResponse) SetBody(v *DeleteCostCenterRuleResponseBody) *DeleteCostCenterRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteCostCenterRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
