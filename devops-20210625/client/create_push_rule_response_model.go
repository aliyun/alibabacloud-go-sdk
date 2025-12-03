// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePushRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePushRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePushRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreatePushRuleResponseBody) *CreatePushRuleResponse
	GetBody() *CreatePushRuleResponseBody
}

type CreatePushRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePushRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePushRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePushRuleResponse) GoString() string {
	return s.String()
}

func (s *CreatePushRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePushRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePushRuleResponse) GetBody() *CreatePushRuleResponseBody {
	return s.Body
}

func (s *CreatePushRuleResponse) SetHeaders(v map[string]*string) *CreatePushRuleResponse {
	s.Headers = v
	return s
}

func (s *CreatePushRuleResponse) SetStatusCode(v int32) *CreatePushRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePushRuleResponse) SetBody(v *CreatePushRuleResponseBody) *CreatePushRuleResponse {
	s.Body = v
	return s
}

func (s *CreatePushRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
