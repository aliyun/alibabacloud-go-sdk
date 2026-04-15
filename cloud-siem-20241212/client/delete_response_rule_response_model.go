// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResponseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResponseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResponseRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResponseRuleResponseBody) *DeleteResponseRuleResponse
	GetBody() *DeleteResponseRuleResponseBody
}

type DeleteResponseRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResponseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResponseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResponseRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteResponseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResponseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResponseRuleResponse) GetBody() *DeleteResponseRuleResponseBody {
	return s.Body
}

func (s *DeleteResponseRuleResponse) SetHeaders(v map[string]*string) *DeleteResponseRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteResponseRuleResponse) SetStatusCode(v int32) *DeleteResponseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResponseRuleResponse) SetBody(v *DeleteResponseRuleResponseBody) *DeleteResponseRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteResponseRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
