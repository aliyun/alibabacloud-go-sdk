// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostCenterShareRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCostCenterShareRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCostCenterShareRuleResponse
	GetStatusCode() *int32
	SetBody(v *QueryCostCenterShareRuleResponseBody) *QueryCostCenterShareRuleResponse
	GetBody() *QueryCostCenterShareRuleResponseBody
}

type QueryCostCenterShareRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCostCenterShareRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCostCenterShareRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterShareRuleResponse) GoString() string {
	return s.String()
}

func (s *QueryCostCenterShareRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCostCenterShareRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCostCenterShareRuleResponse) GetBody() *QueryCostCenterShareRuleResponseBody {
	return s.Body
}

func (s *QueryCostCenterShareRuleResponse) SetHeaders(v map[string]*string) *QueryCostCenterShareRuleResponse {
	s.Headers = v
	return s
}

func (s *QueryCostCenterShareRuleResponse) SetStatusCode(v int32) *QueryCostCenterShareRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCostCenterShareRuleResponse) SetBody(v *QueryCostCenterShareRuleResponseBody) *QueryCostCenterShareRuleResponse {
	s.Body = v
	return s
}

func (s *QueryCostCenterShareRuleResponse) Validate() error {
	return dara.Validate(s)
}
