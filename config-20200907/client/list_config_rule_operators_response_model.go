// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRuleOperatorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConfigRuleOperatorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConfigRuleOperatorsResponse
	GetStatusCode() *int32
	SetBody(v *ListConfigRuleOperatorsResponseBody) *ListConfigRuleOperatorsResponse
	GetBody() *ListConfigRuleOperatorsResponseBody
}

type ListConfigRuleOperatorsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConfigRuleOperatorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConfigRuleOperatorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRuleOperatorsResponse) GoString() string {
	return s.String()
}

func (s *ListConfigRuleOperatorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConfigRuleOperatorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConfigRuleOperatorsResponse) GetBody() *ListConfigRuleOperatorsResponseBody {
	return s.Body
}

func (s *ListConfigRuleOperatorsResponse) SetHeaders(v map[string]*string) *ListConfigRuleOperatorsResponse {
	s.Headers = v
	return s
}

func (s *ListConfigRuleOperatorsResponse) SetStatusCode(v int32) *ListConfigRuleOperatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigRuleOperatorsResponse) SetBody(v *ListConfigRuleOperatorsResponseBody) *ListConfigRuleOperatorsResponse {
	s.Body = v
	return s
}

func (s *ListConfigRuleOperatorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
