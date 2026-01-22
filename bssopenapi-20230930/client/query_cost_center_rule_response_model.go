// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostCenterRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCostCenterRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCostCenterRuleResponse
	GetStatusCode() *int32
	SetBody(v *QueryCostCenterRuleResponseBody) *QueryCostCenterRuleResponse
	GetBody() *QueryCostCenterRuleResponseBody
}

type QueryCostCenterRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCostCenterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCostCenterRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterRuleResponse) GoString() string {
	return s.String()
}

func (s *QueryCostCenterRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCostCenterRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCostCenterRuleResponse) GetBody() *QueryCostCenterRuleResponseBody {
	return s.Body
}

func (s *QueryCostCenterRuleResponse) SetHeaders(v map[string]*string) *QueryCostCenterRuleResponse {
	s.Headers = v
	return s
}

func (s *QueryCostCenterRuleResponse) SetStatusCode(v int32) *QueryCostCenterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCostCenterRuleResponse) SetBody(v *QueryCostCenterRuleResponseBody) *QueryCostCenterRuleResponse {
	s.Body = v
	return s
}

func (s *QueryCostCenterRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
