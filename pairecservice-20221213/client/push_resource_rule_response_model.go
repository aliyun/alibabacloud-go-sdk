// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushResourceRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushResourceRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushResourceRuleResponse
	GetStatusCode() *int32
	SetBody(v *PushResourceRuleResponseBody) *PushResourceRuleResponse
	GetBody() *PushResourceRuleResponseBody
}

type PushResourceRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushResourceRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s PushResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *PushResourceRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushResourceRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushResourceRuleResponse) GetBody() *PushResourceRuleResponseBody {
	return s.Body
}

func (s *PushResourceRuleResponse) SetHeaders(v map[string]*string) *PushResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *PushResourceRuleResponse) SetStatusCode(v int32) *PushResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *PushResourceRuleResponse) SetBody(v *PushResourceRuleResponseBody) *PushResourceRuleResponse {
	s.Body = v
	return s
}

func (s *PushResourceRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
