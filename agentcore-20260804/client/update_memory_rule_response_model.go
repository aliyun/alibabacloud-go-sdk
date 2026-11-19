// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMemoryRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMemoryRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMemoryRuleResponseBody) *UpdateMemoryRuleResponse
	GetBody() *UpdateMemoryRuleResponseBody
}

type UpdateMemoryRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMemoryRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMemoryRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateMemoryRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMemoryRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMemoryRuleResponse) GetBody() *UpdateMemoryRuleResponseBody {
	return s.Body
}

func (s *UpdateMemoryRuleResponse) SetHeaders(v map[string]*string) *UpdateMemoryRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateMemoryRuleResponse) SetStatusCode(v int32) *UpdateMemoryRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMemoryRuleResponse) SetBody(v *UpdateMemoryRuleResponseBody) *UpdateMemoryRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateMemoryRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
