// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyActionEventPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyActionEventPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyActionEventPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyActionEventPolicyResponseBody) *ModifyActionEventPolicyResponse
	GetBody() *ModifyActionEventPolicyResponseBody
}

type ModifyActionEventPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyActionEventPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyActionEventPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyActionEventPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifyActionEventPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyActionEventPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyActionEventPolicyResponse) GetBody() *ModifyActionEventPolicyResponseBody {
	return s.Body
}

func (s *ModifyActionEventPolicyResponse) SetHeaders(v map[string]*string) *ModifyActionEventPolicyResponse {
	s.Headers = v
	return s
}

func (s *ModifyActionEventPolicyResponse) SetStatusCode(v int32) *ModifyActionEventPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyActionEventPolicyResponse) SetBody(v *ModifyActionEventPolicyResponseBody) *ModifyActionEventPolicyResponse {
	s.Body = v
	return s
}

func (s *ModifyActionEventPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
