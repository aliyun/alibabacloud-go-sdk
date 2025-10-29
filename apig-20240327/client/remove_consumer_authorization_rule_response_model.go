// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveConsumerAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveConsumerAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveConsumerAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *RemoveConsumerAuthorizationRuleResponseBody) *RemoveConsumerAuthorizationRuleResponse
	GetBody() *RemoveConsumerAuthorizationRuleResponseBody
}

type RemoveConsumerAuthorizationRuleResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveConsumerAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveConsumerAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveConsumerAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *RemoveConsumerAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveConsumerAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveConsumerAuthorizationRuleResponse) GetBody() *RemoveConsumerAuthorizationRuleResponseBody {
	return s.Body
}

func (s *RemoveConsumerAuthorizationRuleResponse) SetHeaders(v map[string]*string) *RemoveConsumerAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *RemoveConsumerAuthorizationRuleResponse) SetStatusCode(v int32) *RemoveConsumerAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveConsumerAuthorizationRuleResponse) SetBody(v *RemoveConsumerAuthorizationRuleResponseBody) *RemoveConsumerAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *RemoveConsumerAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
