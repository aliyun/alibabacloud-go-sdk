// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCostCenterShareRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveCostCenterShareRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveCostCenterShareRuleResponse
	GetStatusCode() *int32
	SetBody(v *SaveCostCenterShareRuleResponseBody) *SaveCostCenterShareRuleResponse
	GetBody() *SaveCostCenterShareRuleResponseBody
}

type SaveCostCenterShareRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveCostCenterShareRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveCostCenterShareRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveCostCenterShareRuleResponse) GoString() string {
	return s.String()
}

func (s *SaveCostCenterShareRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveCostCenterShareRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveCostCenterShareRuleResponse) GetBody() *SaveCostCenterShareRuleResponseBody {
	return s.Body
}

func (s *SaveCostCenterShareRuleResponse) SetHeaders(v map[string]*string) *SaveCostCenterShareRuleResponse {
	s.Headers = v
	return s
}

func (s *SaveCostCenterShareRuleResponse) SetStatusCode(v int32) *SaveCostCenterShareRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveCostCenterShareRuleResponse) SetBody(v *SaveCostCenterShareRuleResponseBody) *SaveCostCenterShareRuleResponse {
	s.Body = v
	return s
}

func (s *SaveCostCenterShareRuleResponse) Validate() error {
	return dara.Validate(s)
}
