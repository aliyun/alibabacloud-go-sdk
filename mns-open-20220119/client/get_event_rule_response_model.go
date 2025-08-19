// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEventRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEventRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetEventRuleResponseBody) *GetEventRuleResponse
	GetBody() *GetEventRuleResponseBody
}

type GetEventRuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEventRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEventRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEventRuleResponse) GoString() string {
	return s.String()
}

func (s *GetEventRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEventRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEventRuleResponse) GetBody() *GetEventRuleResponseBody {
	return s.Body
}

func (s *GetEventRuleResponse) SetHeaders(v map[string]*string) *GetEventRuleResponse {
	s.Headers = v
	return s
}

func (s *GetEventRuleResponse) SetStatusCode(v int32) *GetEventRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEventRuleResponse) SetBody(v *GetEventRuleResponseBody) *GetEventRuleResponse {
	s.Body = v
	return s
}

func (s *GetEventRuleResponse) Validate() error {
	return dara.Validate(s)
}
