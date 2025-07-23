// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEventRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEventRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateEventRuleResponseBody) *CreateEventRuleResponse
	GetBody() *CreateEventRuleResponseBody
}

type CreateEventRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateEventRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEventRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEventRuleResponse) GetBody() *CreateEventRuleResponseBody {
	return s.Body
}

func (s *CreateEventRuleResponse) SetHeaders(v map[string]*string) *CreateEventRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateEventRuleResponse) SetStatusCode(v int32) *CreateEventRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventRuleResponse) SetBody(v *CreateEventRuleResponseBody) *CreateEventRuleResponse {
	s.Body = v
	return s
}

func (s *CreateEventRuleResponse) Validate() error {
	return dara.Validate(s)
}
