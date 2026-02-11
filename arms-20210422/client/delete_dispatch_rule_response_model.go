// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDispatchRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDispatchRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDispatchRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDispatchRuleResponseBody) *DeleteDispatchRuleResponse
	GetBody() *DeleteDispatchRuleResponseBody
}

type DeleteDispatchRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDispatchRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDispatchRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDispatchRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteDispatchRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDispatchRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDispatchRuleResponse) GetBody() *DeleteDispatchRuleResponseBody {
	return s.Body
}

func (s *DeleteDispatchRuleResponse) SetHeaders(v map[string]*string) *DeleteDispatchRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteDispatchRuleResponse) SetStatusCode(v int32) *DeleteDispatchRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDispatchRuleResponse) SetBody(v *DeleteDispatchRuleResponseBody) *DeleteDispatchRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteDispatchRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
