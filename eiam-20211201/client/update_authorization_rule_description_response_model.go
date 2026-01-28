// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAuthorizationRuleDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAuthorizationRuleDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAuthorizationRuleDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAuthorizationRuleDescriptionResponseBody) *UpdateAuthorizationRuleDescriptionResponse
	GetBody() *UpdateAuthorizationRuleDescriptionResponseBody
}

type UpdateAuthorizationRuleDescriptionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAuthorizationRuleDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAuthorizationRuleDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAuthorizationRuleDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateAuthorizationRuleDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAuthorizationRuleDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAuthorizationRuleDescriptionResponse) GetBody() *UpdateAuthorizationRuleDescriptionResponseBody {
	return s.Body
}

func (s *UpdateAuthorizationRuleDescriptionResponse) SetHeaders(v map[string]*string) *UpdateAuthorizationRuleDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateAuthorizationRuleDescriptionResponse) SetStatusCode(v int32) *UpdateAuthorizationRuleDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAuthorizationRuleDescriptionResponse) SetBody(v *UpdateAuthorizationRuleDescriptionResponseBody) *UpdateAuthorizationRuleDescriptionResponse {
	s.Body = v
	return s
}

func (s *UpdateAuthorizationRuleDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
