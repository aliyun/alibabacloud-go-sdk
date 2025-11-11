// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInterveneRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInterveneRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInterveneRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInterveneRuleResponseBody) *DeleteInterveneRuleResponse
	GetBody() *DeleteInterveneRuleResponseBody
}

type DeleteInterveneRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInterveneRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInterveneRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInterveneRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteInterveneRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInterveneRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInterveneRuleResponse) GetBody() *DeleteInterveneRuleResponseBody {
	return s.Body
}

func (s *DeleteInterveneRuleResponse) SetHeaders(v map[string]*string) *DeleteInterveneRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteInterveneRuleResponse) SetStatusCode(v int32) *DeleteInterveneRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInterveneRuleResponse) SetBody(v *DeleteInterveneRuleResponseBody) *DeleteInterveneRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteInterveneRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
