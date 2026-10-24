// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientRuleStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClientRuleStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClientRuleStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClientRuleStrategyResponseBody) *DeleteClientRuleStrategyResponse
	GetBody() *DeleteClientRuleStrategyResponseBody
}

type DeleteClientRuleStrategyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClientRuleStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClientRuleStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientRuleStrategyResponse) GoString() string {
	return s.String()
}

func (s *DeleteClientRuleStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClientRuleStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClientRuleStrategyResponse) GetBody() *DeleteClientRuleStrategyResponseBody {
	return s.Body
}

func (s *DeleteClientRuleStrategyResponse) SetHeaders(v map[string]*string) *DeleteClientRuleStrategyResponse {
	s.Headers = v
	return s
}

func (s *DeleteClientRuleStrategyResponse) SetStatusCode(v int32) *DeleteClientRuleStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClientRuleStrategyResponse) SetBody(v *DeleteClientRuleStrategyResponseBody) *DeleteClientRuleStrategyResponse {
	s.Body = v
	return s
}

func (s *DeleteClientRuleStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
