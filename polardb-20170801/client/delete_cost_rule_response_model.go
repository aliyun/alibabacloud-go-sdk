// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCostRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCostRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCostRuleResponseBody) *DeleteCostRuleResponse
	GetBody() *DeleteCostRuleResponseBody
}

type DeleteCostRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCostRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCostRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCostRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCostRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCostRuleResponse) GetBody() *DeleteCostRuleResponseBody {
	return s.Body
}

func (s *DeleteCostRuleResponse) SetHeaders(v map[string]*string) *DeleteCostRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteCostRuleResponse) SetStatusCode(v int32) *DeleteCostRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCostRuleResponse) SetBody(v *DeleteCostRuleResponseBody) *DeleteCostRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteCostRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
