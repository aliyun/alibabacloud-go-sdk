// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemoryRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemoryRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetMemoryRuleResponseBody) *GetMemoryRuleResponse
	GetBody() *GetMemoryRuleResponseBody
}

type GetMemoryRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemoryRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoryRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryRuleResponse) GoString() string {
	return s.String()
}

func (s *GetMemoryRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemoryRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemoryRuleResponse) GetBody() *GetMemoryRuleResponseBody {
	return s.Body
}

func (s *GetMemoryRuleResponse) SetHeaders(v map[string]*string) *GetMemoryRuleResponse {
	s.Headers = v
	return s
}

func (s *GetMemoryRuleResponse) SetStatusCode(v int32) *GetMemoryRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoryRuleResponse) SetBody(v *GetMemoryRuleResponseBody) *GetMemoryRuleResponse {
	s.Body = v
	return s
}

func (s *GetMemoryRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
