// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveApplicationFromAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveApplicationFromAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveApplicationFromAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *RemoveApplicationFromAuthorizationRuleResponseBody) *RemoveApplicationFromAuthorizationRuleResponse
	GetBody() *RemoveApplicationFromAuthorizationRuleResponseBody
}

type RemoveApplicationFromAuthorizationRuleResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveApplicationFromAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveApplicationFromAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveApplicationFromAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *RemoveApplicationFromAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveApplicationFromAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveApplicationFromAuthorizationRuleResponse) GetBody() *RemoveApplicationFromAuthorizationRuleResponseBody {
	return s.Body
}

func (s *RemoveApplicationFromAuthorizationRuleResponse) SetHeaders(v map[string]*string) *RemoveApplicationFromAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *RemoveApplicationFromAuthorizationRuleResponse) SetStatusCode(v int32) *RemoveApplicationFromAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveApplicationFromAuthorizationRuleResponse) SetBody(v *RemoveApplicationFromAuthorizationRuleResponseBody) *RemoveApplicationFromAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *RemoveApplicationFromAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
