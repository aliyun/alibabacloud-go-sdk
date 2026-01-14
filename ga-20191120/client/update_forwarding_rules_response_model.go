// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateForwardingRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateForwardingRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateForwardingRulesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateForwardingRulesResponseBody) *UpdateForwardingRulesResponse
	GetBody() *UpdateForwardingRulesResponseBody
}

type UpdateForwardingRulesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateForwardingRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateForwardingRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateForwardingRulesResponse) GoString() string {
	return s.String()
}

func (s *UpdateForwardingRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateForwardingRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateForwardingRulesResponse) GetBody() *UpdateForwardingRulesResponseBody {
	return s.Body
}

func (s *UpdateForwardingRulesResponse) SetHeaders(v map[string]*string) *UpdateForwardingRulesResponse {
	s.Headers = v
	return s
}

func (s *UpdateForwardingRulesResponse) SetStatusCode(v int32) *UpdateForwardingRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateForwardingRulesResponse) SetBody(v *UpdateForwardingRulesResponseBody) *UpdateForwardingRulesResponse {
	s.Body = v
	return s
}

func (s *UpdateForwardingRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
