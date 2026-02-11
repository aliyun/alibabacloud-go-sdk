// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomizeRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomizeRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomizeRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomizeRuleResponseBody) *DeleteCustomizeRuleResponse
	GetBody() *DeleteCustomizeRuleResponseBody
}

type DeleteCustomizeRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomizeRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomizeRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomizeRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomizeRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomizeRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomizeRuleResponse) GetBody() *DeleteCustomizeRuleResponseBody {
	return s.Body
}

func (s *DeleteCustomizeRuleResponse) SetHeaders(v map[string]*string) *DeleteCustomizeRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomizeRuleResponse) SetStatusCode(v int32) *DeleteCustomizeRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomizeRuleResponse) SetBody(v *DeleteCustomizeRuleResponseBody) *DeleteCustomizeRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomizeRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
