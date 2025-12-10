// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoGroupingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAutoGroupingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAutoGroupingRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateAutoGroupingRuleResponseBody) *CreateAutoGroupingRuleResponse
	GetBody() *CreateAutoGroupingRuleResponseBody
}

type CreateAutoGroupingRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAutoGroupingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAutoGroupingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoGroupingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoGroupingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAutoGroupingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAutoGroupingRuleResponse) GetBody() *CreateAutoGroupingRuleResponseBody {
	return s.Body
}

func (s *CreateAutoGroupingRuleResponse) SetHeaders(v map[string]*string) *CreateAutoGroupingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoGroupingRuleResponse) SetStatusCode(v int32) *CreateAutoGroupingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoGroupingRuleResponse) SetBody(v *CreateAutoGroupingRuleResponseBody) *CreateAutoGroupingRuleResponse {
	s.Body = v
	return s
}

func (s *CreateAutoGroupingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
