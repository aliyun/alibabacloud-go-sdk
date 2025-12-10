// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutoGroupingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAutoGroupingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAutoGroupingRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAutoGroupingRuleResponseBody) *UpdateAutoGroupingRuleResponse
	GetBody() *UpdateAutoGroupingRuleResponseBody
}

type UpdateAutoGroupingRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutoGroupingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAutoGroupingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutoGroupingRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoGroupingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAutoGroupingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAutoGroupingRuleResponse) GetBody() *UpdateAutoGroupingRuleResponseBody {
	return s.Body
}

func (s *UpdateAutoGroupingRuleResponse) SetHeaders(v map[string]*string) *UpdateAutoGroupingRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoGroupingRuleResponse) SetStatusCode(v int32) *UpdateAutoGroupingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutoGroupingRuleResponse) SetBody(v *UpdateAutoGroupingRuleResponseBody) *UpdateAutoGroupingRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateAutoGroupingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
