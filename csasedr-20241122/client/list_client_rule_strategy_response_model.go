// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientRuleStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClientRuleStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClientRuleStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ListClientRuleStrategyResponseBody) *ListClientRuleStrategyResponse
	GetBody() *ListClientRuleStrategyResponseBody
}

type ListClientRuleStrategyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientRuleStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientRuleStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClientRuleStrategyResponse) GoString() string {
	return s.String()
}

func (s *ListClientRuleStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClientRuleStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClientRuleStrategyResponse) GetBody() *ListClientRuleStrategyResponseBody {
	return s.Body
}

func (s *ListClientRuleStrategyResponse) SetHeaders(v map[string]*string) *ListClientRuleStrategyResponse {
	s.Headers = v
	return s
}

func (s *ListClientRuleStrategyResponse) SetStatusCode(v int32) *ListClientRuleStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientRuleStrategyResponse) SetBody(v *ListClientRuleStrategyResponseBody) *ListClientRuleStrategyResponse {
	s.Body = v
	return s
}

func (s *ListClientRuleStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
