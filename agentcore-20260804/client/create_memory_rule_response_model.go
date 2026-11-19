// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMemoryRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMemoryRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMemoryRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateMemoryRuleResponseBody) *CreateMemoryRuleResponse
	GetBody() *CreateMemoryRuleResponseBody
}

type CreateMemoryRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMemoryRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMemoryRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMemoryRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateMemoryRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMemoryRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMemoryRuleResponse) GetBody() *CreateMemoryRuleResponseBody {
	return s.Body
}

func (s *CreateMemoryRuleResponse) SetHeaders(v map[string]*string) *CreateMemoryRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateMemoryRuleResponse) SetStatusCode(v int32) *CreateMemoryRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMemoryRuleResponse) SetBody(v *CreateMemoryRuleResponseBody) *CreateMemoryRuleResponse {
	s.Body = v
	return s
}

func (s *CreateMemoryRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
