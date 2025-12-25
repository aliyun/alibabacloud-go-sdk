// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserWafRulesetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUserWafRulesetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUserWafRulesetResponse
	GetStatusCode() *int32
	SetBody(v *CreateUserWafRulesetResponseBody) *CreateUserWafRulesetResponse
	GetBody() *CreateUserWafRulesetResponseBody
}

type CreateUserWafRulesetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserWafRulesetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserWafRulesetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUserWafRulesetResponse) GoString() string {
	return s.String()
}

func (s *CreateUserWafRulesetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUserWafRulesetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUserWafRulesetResponse) GetBody() *CreateUserWafRulesetResponseBody {
	return s.Body
}

func (s *CreateUserWafRulesetResponse) SetHeaders(v map[string]*string) *CreateUserWafRulesetResponse {
	s.Headers = v
	return s
}

func (s *CreateUserWafRulesetResponse) SetStatusCode(v int32) *CreateUserWafRulesetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserWafRulesetResponse) SetBody(v *CreateUserWafRulesetResponseBody) *CreateUserWafRulesetResponse {
	s.Body = v
	return s
}

func (s *CreateUserWafRulesetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
