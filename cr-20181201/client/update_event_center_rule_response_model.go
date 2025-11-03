// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventCenterRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEventCenterRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEventCenterRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEventCenterRuleResponseBody) *UpdateEventCenterRuleResponse
	GetBody() *UpdateEventCenterRuleResponseBody
}

type UpdateEventCenterRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventCenterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventCenterRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventCenterRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventCenterRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEventCenterRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEventCenterRuleResponse) GetBody() *UpdateEventCenterRuleResponseBody {
	return s.Body
}

func (s *UpdateEventCenterRuleResponse) SetHeaders(v map[string]*string) *UpdateEventCenterRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventCenterRuleResponse) SetStatusCode(v int32) *UpdateEventCenterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventCenterRuleResponse) SetBody(v *UpdateEventCenterRuleResponseBody) *UpdateEventCenterRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateEventCenterRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
