// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInterceptionRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInterceptionRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInterceptionRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInterceptionRuleResponseBody) *DeleteInterceptionRuleResponse
	GetBody() *DeleteInterceptionRuleResponseBody
}

type DeleteInterceptionRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInterceptionRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInterceptionRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInterceptionRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteInterceptionRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInterceptionRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInterceptionRuleResponse) GetBody() *DeleteInterceptionRuleResponseBody {
	return s.Body
}

func (s *DeleteInterceptionRuleResponse) SetHeaders(v map[string]*string) *DeleteInterceptionRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteInterceptionRuleResponse) SetStatusCode(v int32) *DeleteInterceptionRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInterceptionRuleResponse) SetBody(v *DeleteInterceptionRuleResponseBody) *DeleteInterceptionRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteInterceptionRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
